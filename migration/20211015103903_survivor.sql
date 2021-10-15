-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS "survivor" (
 "id"  SERIAL  PRIMARY KEY,
 "name" varchar(50) NOT NULL DEFAULT '',
 "gender" char(1) NOT NULL DEFAULT '',
 "longitude" varchar(255) DEFAULT '',
 "latitude" varchar(255) DEFAULT '',
 "infected" int DEFAULT 0
    );
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table survivor;