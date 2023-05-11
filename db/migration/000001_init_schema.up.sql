CREATE TABLE "employees" (
  "id" serial PRIMARY KEY,
  "identity_id" serial NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "password" varchar NOT NULL DEFAULT 'pa@ss123word',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "identities" (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "punches" (
  "id" bigserial PRIMARY KEY,
  "employee_id" serial NOT NULL,
  "working_day" timestamptz NOT NULL DEFAULT (now()),
  "working_hours" smallint NOT NULL DEFAULT 0,
  "status_id" smallint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "statuses" (
  "id" smallint PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "holidays" (
  "id" serial PRIMARY KEY,
  "date" timestamptz UNIQUE NOT NULL
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "employee_id" serial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_employee_id" serial NOT NULL,
  "to_employee_id" serial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "employees" ("code");

CREATE INDEX ON "punches" ("employee_id");

CREATE INDEX ON "punches" ("working_day");

CREATE INDEX ON "entries" ("employee_id");

CREATE INDEX ON "transfers" ("from_employee_id");

CREATE INDEX ON "transfers" ("to_employee_id");

CREATE INDEX ON "transfers" ("from_employee_id", "to_employee_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "employees" ADD FOREIGN KEY ("identity_id") REFERENCES "identities" ("id");

ALTER TABLE "punches" ADD FOREIGN KEY ("employee_id") REFERENCES "employees" ("id");

ALTER TABLE "punches" ADD FOREIGN KEY ("status_id") REFERENCES "statuses" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("employee_id") REFERENCES "employees" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_employee_id") REFERENCES "employees" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_employee_id") REFERENCES "employees" ("id");
