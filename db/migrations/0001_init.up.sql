-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS contacts_id_seq;

-- Table Definition
CREATE TABLE "public"."contacts" (
    "id" int4 NOT NULL DEFAULT nextval('contacts_id_seq'::regclass),
    "firstname" varchar NOT NULL,
    "lastname" varchar NOT NULL,
    "email" varchar NOT NULL,
    "phone" varchar NOT NULL,
    PRIMARY KEY ("id")
);