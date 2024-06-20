-- +goose Up
CREATE TABLE stocks (
    ticker varchar(8) PRIMARY KEY,
    price numeric(18, 10) NOT NULL,
    pe numeric(6, 2) DEFAULT 0,
    ps numeric(6, 2) DEFAULT 0,
    market_cap numeric(20, 5) DEFAULT 0,
    avg_dividend_5_year numeric(4, 2) DEFAULT 0
);

CREATE TABLE bonds (
    ticker varchar(16) PRIMARY KEY,
    price numeric(18, 10) NOT NULL,
    profit_percent numeric(5, 2) DEFAULT 0,
    date_expire date,
    coupon_value numeric(5, 2) DEFAULT 0,
    payments_per_year smallint DEFAULT 0
);

CREATE TABLE etfs (
    ticker varchar(8) PRIMARY KEY,
    price numeric(18, 10) NOT NULL,
    fund_tax numeric(8, 6) DEFAULT 0
);

create or replace view v_tickers as 
select s.ticker, 'Акции' as type
from stocks s 
union
select b.ticker, 'Облигации' as type
from bonds b
union
select e.ticker, 'ETF' as type
from etfs e;

-- +goose Down
DROP VIEW IF EXISTS v_tickers;
DROP TABLE IF EXISTS stocks;
DROP TABLE IF EXISTS bonds;
DROP TABLE IF EXISTS etfs;