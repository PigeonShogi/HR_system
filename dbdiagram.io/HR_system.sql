CREATE TABLE "employees" (
  "id" serial PRIMARY KEY,
  "Identity_Id" serial NOT NULL,
  "code" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "password" varchar NOT NULL DEFAULT 'pa@ss123word',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "identities" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "punches" (
  "id" bigserial PRIMARY KEY,
  "Employee_Id" serial NOT NULL,
  "working_day" timestamptz NOT NULL DEFAULT (now()),
  "working_hours" smallint NOT NULL DEFAULT 0,
  "Status_Id" smallint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "statuses" (
  "id" smallint PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "holidays" (
  "id" serial PRIMARY KEY,
  "date" timestamptz
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "Employee_Id" serial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "From_Employee_Id" serial NOT NULL,
  "To_Employee_Id" serial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "employees" ("code");

CREATE INDEX ON "punches" ("Employee_Id");

CREATE INDEX ON "punches" ("working_day");

CREATE INDEX ON "entries" ("Employee_Id");

CREATE INDEX ON "transfers" ("From_Employee_Id");

CREATE INDEX ON "transfers" ("To_Employee_Id");

CREATE INDEX ON "transfers" ("From_Employee_Id", "To_Employee_Id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "employees" ADD FOREIGN KEY ("Identity_Id") REFERENCES "identities" ("id");

ALTER TABLE "punches" ADD FOREIGN KEY ("Employee_Id") REFERENCES "employees" ("id");

ALTER TABLE "punches" ADD FOREIGN KEY ("Status_Id") REFERENCES "statuses" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("Employee_Id") REFERENCES "employees" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("From_Employee_Id") REFERENCES "employees" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("To_Employee_Id") REFERENCES "employees" ("id");
