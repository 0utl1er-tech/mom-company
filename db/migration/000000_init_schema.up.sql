-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2025-10-11T04:58:40.136Z

CREATE TYPE "type" AS ENUM (
  'kabu',
  'yugen',
  'godo',
  'goshi',
  'gomei',
  'other'
);

CREATE TYPE "presuf" AS ENUM (
  'prefix',
  'suffix'
);

CREATE TABLE "company" (
  "id" uuid PRIMARY KEY,
  "ceo" uuid UNIQUE NOT NULL,
  "trademark" varchar NOT NULL,
  "type" type NOT NULL,
  "position" presuf NOT NULL,
  "address" varchar NOT NULL,
  "company_code" varchar UNIQUE NOT NULL,
  "contact_id" uuid UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "staff" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "role" varchar NOT NULL DEFAULT '一般社員',
  "contact_id" uuid UNIQUE NOT NULL,
  "company_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "contact" (
  "id" uuid PRIMARY KEY,
  "email" varchar NOT NULL DEFAULT '',
  "phone" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "company" ("trademark", "position", "type", "address");

CREATE INDEX ON "company" ("company_code");

COMMENT ON COLUMN "company"."trademark" IS '商標';

COMMENT ON COLUMN "company"."position" IS '前株か後株か';

COMMENT ON COLUMN "company"."company_code" IS '法人番号';

COMMENT ON COLUMN "staff"."role" IS '役職';

ALTER TABLE "staff" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "staff" ADD FOREIGN KEY ("id") REFERENCES "company" ("ceo");

ALTER TABLE "company" ADD FOREIGN KEY ("contact_id") REFERENCES "contact" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("contact_id") REFERENCES "contact" ("id");
