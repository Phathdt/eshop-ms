-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" text,
    "password" text,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

CREATE UNIQUE INDEX ON "users" ("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
