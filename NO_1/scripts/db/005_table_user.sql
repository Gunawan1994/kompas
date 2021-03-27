BEGIN;

-- CREATE TYPE ENUM "role" ---------------------------------
CREATE TYPE role_type as ENUM('Super Admin');
------------------------------------------------------------

-- CREATE TABLE "user" -------------------------------------
CREATE TABLE "public"."user" (
    "id" Bigint NOT NULL,
    "username" Character Varying( 255 ) NOT NULL,
    "email" Character Varying( 255 ) NOT NULL,
    "password" Character Varying( 255 ) NOT NULL,
    "role" role_type,
    PRIMARY KEY ( "id" ) );
 ;
 -----------------------------------------------------------

-- Password "adrianto123"
INSERT INTO "public"."user" ("id", "username", "email", "password", "role") VALUES
(123456, 'Gunawan', 'gunawan@gmail.com', '$2a$14$W2CPplwNkLuE0Ar3XRFGauWKz2gbS.wNQR4/QXUJRYdafoguGwUFO', 'Super Admin');

COMMIT;