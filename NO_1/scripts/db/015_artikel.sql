BEGIN;

-- CREATE SEQUENCE "id" -------------------------------
CREATE SEQUENCE "public"."id"
INCREMENT 1
MINVALUE 1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
-- -------------------------------------------------------------

-- CREATE TABLE "artikel" -------------------------------------
CREATE TABLE "public"."artikel" (
    "id" Bigint DEFAULT nextval('id'::regclass) NOT NULL,
    "judul" Character Varying( 255 ) NOT NULL,
	"body" Text NOT NULL,
	"author" Bigint NOT NULL,
    PRIMARY KEY ( "id" ) );
 ;
 -----------------------------------------------------------


COMMIT;