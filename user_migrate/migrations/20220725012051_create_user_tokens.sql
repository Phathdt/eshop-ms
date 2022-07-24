-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_tokens" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" bigint,
    "token" text,
    "expired_at" timestamp(0) DEFAULT now(),
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

CREATE UNIQUE INDEX ON "user_tokens" ("user_id", "token");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user_tokens";
-- +goose StatementEnd
