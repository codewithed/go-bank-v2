CREATE TABLE IF NOT EXISTS "users" (
  "username" varchar PRIMARY KEY,
  "password_hash" varchar NOT NULL,
  "email" varchar UNIQUE,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency")
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
