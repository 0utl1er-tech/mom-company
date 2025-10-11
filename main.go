package main

import (
	"context"
	"net/http"
	"os"
	"syscall"

	companyv1connect "github.com/0utl1er-tech/mom-company/gen/pb/company/v1/companyv1connect"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/0utl1er-tech/mom-company/internal/handler"
	"github.com/0utl1er-tech/mom-company/internal/service/company"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	connPool, err := pgxpool.New(context.Background(), cfg.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create connection pool")
	}

	queries := db.New(connPool)
	companyService := company.NewService(queries, connPool)

	// HTTPサーバーの設定
	mux := http.NewServeMux()

	// Connect-Goハンドラーを登録
	companyHandler := handler.NewCompanyServiceHandler(companyService)
	path, handler := companyv1connect.NewCompanyServiceHandler(companyHandler)
	mux.Handle(path, handler)

	// HTTP/2対応のサーバーを作成
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	// サーバー起動とGraceful Shutdown
	waitGroup, ctx := errgroup.WithContext(context.Background())

	waitGroup.Go(func() error {
		log.Info().Msgf("Start Connect-Go server at %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Connect-Go server failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("Graceful shutdown Connect-Go server")
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Failed to shutdown server gracefully")
			return err
		}
		log.Info().Msg("Connect-Go server is stopped")
		return nil
	})

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to wait")
	}
}
