CREATE TABLE IF NOT EXISTS rates (
                              id serial4 NOT NULL,
                              currency varchar(5) NOT NULL,
                              rate float8 NOT NULL,
                              CONSTRAINT rates_currency_key UNIQUE (currency)
);