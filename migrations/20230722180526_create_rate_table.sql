-- +goose Up
-- +goose StatementBegin
CREATE TABLE rates(
  id serial PRIMARY KEY,
  usd NUMERIC(11, 6) not null,  
  eur NUMERIC(11, 6) not null,  
  gbp NUMERIC(11, 6) not null,  
  date timestamp not null UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rates;
-- +goose StatementEnd

