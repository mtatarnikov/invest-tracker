-- +goose Up
INSERT INTO public.users
("name", login, "password")
VALUES('demo', 'demo', '$2a$10$jToRj7.ZxrQ.2k5FlJ4h3ujZFVgx9YmvUadebqqpdCH2HXROo3LFK');
INSERT INTO public.users
("name", login, "password")
VALUES('Mike T', 'Mike', '$2a$10$oh34Bd5dfk05W.fvGe/ERet/61vEUU3ApPRyaHnpLT7ePKRKNzNSG');

INSERT INTO public.assets
(id, user_id, ticker, transaction_type, transaction_date, amount, price, tax, note)
VALUES(1, 1, 'SBER', 'buy', '2024-01-01', 10, 250.0000000000, 0.00, NULL);
INSERT INTO public.assets
(id, user_id, ticker, transaction_type, transaction_date, amount, price, tax, note)
VALUES(2, 1, 'LKOH', 'buy', '2024-01-02', 1, 7000.0000000000, 0.00, 'comment');
INSERT INTO public.assets
(id, user_id, ticker, transaction_type, transaction_date, amount, price, tax, note)
VALUES(3, 1, 'TGLD', 'buy', '2024-06-05', 100, 10.5000000000, 1.00, NULL);
INSERT INTO public.assets
(id, user_id, ticker, transaction_type, transaction_date, amount, price, tax, note)
VALUES(4, 1, 'TITR', 'buy', '2024-06-08', 100, 10.0000000000, 0.00, NULL);
INSERT INTO public.assets
(id, user_id, ticker, transaction_type, transaction_date, amount, price, tax, note)
VALUES(5, 1, 'TITR', 'buy', '2024-06-09', 10, 10.0000000000, 0.00, NULL);

INSERT INTO public.stocks
(ticker, price, pe, ps, market_cap, avg_dividend_5_year)
VALUES('SBER', 317.8600000000, 4.40, 1.09, 6820000000000.00000, 6.95);
INSERT INTO public.stocks
(ticker, price, pe, ps, market_cap, avg_dividend_5_year)
VALUES('LKOH', 7292.0000000000, 4.43, 0.65, 5014000000000.00000, 9.85);

INSERT INTO public.etfs
(ticker, price, fund_tax)
VALUES('TGLD', 8.4100000000, 0.840000);
INSERT INTO public.etfs
(ticker, price, fund_tax)
VALUES('TITR', 9.5100000000, 1.490000);

-- +goose Down
delete from assets where user_id=1
;
delete from users where id=1
;