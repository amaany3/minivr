CREATE SCHEMA IF NOT EXISTS "public";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-----------------------------
-- publishers
-----------------------------
CREATE TABLE "public"."publishers" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "name" character varying(128) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
