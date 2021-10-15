-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS "inventory" (
    "id"  SERIAL  PRIMARY KEY,
    "survivor_id" int NOT NULL ,
    "inventory_name" varchar(255) NOT NULL DEFAULT '',
    "quantity" varchar(255) DEFAULT '',
    "unit" varchar(255) DEFAULT '',
    FOREIGN KEY (survivor_id) REFERENCES survivor (id)
    );
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table if exists inventory;