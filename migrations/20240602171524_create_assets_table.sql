-- +goose Up
CREATE TABLE public.assets (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL,
    ticker varchar(16) NOT NULL,
    transaction_type varchar(4) NOT NULL,
    transaction_date date NOT NULL,
    amount integer NOT NULL,
    price numeric (18,10) NOT NULL,
    tax numeric (9,2) DEFAULT 0,
    "note" varchar(255),
    CONSTRAINT fk_assets_user_id FOREIGN KEY (user_id)
    REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE public.assets;