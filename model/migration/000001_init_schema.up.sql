CREATE TABLE "chain_event" (
                               "id_chain_event" bigserial PRIMARY KEY,
                               "wallet_id" serial NOT NULL,
                               "action_type" varchar,
                               "from_metamask_wallet_id" varchar NOT NULL,
                               "to_metamask_wallet_id" varchar NOT NULL,
                               "value" varchar NULL,
                               "log_event" text,
                               "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "line_event" (
                              "id_line_event" bigserial PRIMARY KEY,
                              "id_line_user" varchar NOT NULL,
                              "id_use" serial NOT NULL,
                              "request_log_event" text,
                              "response_log_event" text,
                              "error" boolean NOT NULL DEFAULT false,
                              "error_text" text,
                              "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "event" (
                         "id_event" bigserial PRIMARY KEY,
                         "id_line_event" serial NOT NULL,
                         "id_chain_event" serial NOT NULL,
                         "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "line_owner_validation" (
                                         "id_line_owner_validation" serial PRIMARY KEY,
                                         "code" varchar NOT NULL,
                                         "id_user" int NOT NULL,
                                         "created_at" timestamp NOT NULL DEFAULT 'now()',
                                         "deleted" timestamp
);

CREATE TABLE "chain" (
                         "id_chain" serial PRIMARY KEY,
                         "chain_code" varchar NOT NULL,
                         "chain_name" varchar NOT NULL,
                         "url_api" text,
                         "created_at" timestamp NOT NULL DEFAULT 'now()',
                         "modified" timestamp NOT NULL DEFAULT 'now()',
                         "deleted" timestamp
);

INSERT INTO chain (id_chain, chain_code, chain_name, url_api )
VALUES (1, 'bsc-testnet', 'binance smart chain test net', 'https://api-testnet.bscscan.com/api');
INSERT INTO chain (id_chain, chain_code, chain_name, url_api )
VALUES (2, 'bsc', 'binance smart chain', 'https://api.bscscan.com/api');

CREATE TABLE "users" (
                         "id_user" serial PRIMARY KEY,
                         "username" varchar NOT NULL,
                         "password" varchar NOT NULL,
                         "id_line" varchar,
                         "owner_validation" boolean NOT NULL DEFAULT false,
                         "created_at" timestamp NOT NULL DEFAULT 'now()',
                         "modified" timestamp NOT NULL DEFAULT 'now()',
                         "deleted" timestamp
);

CREATE TABLE "wallet" (
                          "wallet_id" serial PRIMARY KEY,
                          "metamask_wallet_id" varchar NOT NULL,
                          "follow_wallet" boolean NOT NULL DEFAULT false,
                          "id_user" int NOT NULL,
                          "id_chain" int DEFAULT 1,
                          "last_block_number" int NOT NULL DEFAULT 0,
                          "created_at" timestamp NOT NULL DEFAULT 'now()',
                          "modified" timestamp NOT NULL DEFAULT 'now()',
                          "deleted" timestamp
);

CREATE INDEX ON "chain_event" ("from_metamask_wallet_id");

CREATE INDEX ON "chain_event" ("to_metamask_wallet_id");

CREATE INDEX ON "line_event" ("id_line_user");

CREATE INDEX ON "event" ("id_line_event");

CREATE INDEX ON "event" ("id_chain_event");

CREATE INDEX ON "line_owner_validation" ("id_user");

CREATE INDEX ON "users" ("username");

ALTER TABLE "chain_event" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallet" ("wallet_id");

ALTER TABLE "line_event" ADD FOREIGN KEY ("id_use") REFERENCES "users" ("id_user");

ALTER TABLE "event" ADD FOREIGN KEY ("id_line_event") REFERENCES "line_event" ("id_line_event");

ALTER TABLE "event" ADD FOREIGN KEY ("id_chain_event") REFERENCES "chain_event" ("id_chain_event");

ALTER TABLE "line_owner_validation" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id_user");

ALTER TABLE "wallet" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id_user");

ALTER TABLE "wallet" ADD FOREIGN KEY ("id_chain") REFERENCES "chain" ("id_chain");