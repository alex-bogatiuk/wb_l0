-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                        order_uid VARCHAR(40) NOT NULL,
                        track_number VARCHAR(16) NOT NULL,
                        entry VARCHAR(10) NOT NULL,
                        locale VARCHAR(2) NOT NULL,
                        internal_signature VARCHAR(50) NOT NULL,
                        customer_id VARCHAR(22) NOT NULL,
                        delivery_service VARCHAR(10) NOT NULL,
                        shardkey VARCHAR(10) NOT NULL,
                        sm_id INT NOT NULL,
                        date_created TIMESTAMP NOT NULL,
                        oof_shard VARCHAR(10) NOT NULL
);


INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES
    ('b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26 06:22:19', '1');

CREATE TABLE order_delivery (
                                order_uid VARCHAR(20) NOT NULL,
                                name VARCHAR(50) NOT NULL,
                                phone VARCHAR(12) NOT NULL,
                                zip VARCHAR(10) NOT NULL,
                                city VARCHAR(20) NOT NULL,
                                address VARCHAR(100) NOT NULL,
                                region VARCHAR(50) NOT NULL,
                                email VARCHAR(50) NOT NULL
);


INSERT INTO order_delivery (order_uid, name, phone, zip, city, address, region, email) VALUES
    ('b563feb7b2b84b6test', 'Test Testov"', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com');


CREATE TABLE order_items (
                             order_uid VARCHAR(20) NOT NULL,
                             chrt_id INT NOT NULL,
                             track_number VARCHAR(16) NOT NULL,
                             price INT NOT NULL,
                             rid VARCHAR(22) NOT NULL,
                             name VARCHAR(100) NOT NULL,
                             sale INT NOT NULL,
                             size VARCHAR(4) NOT NULL,
                             total_price INT NOT NULL,
                             nm_id INT NOT NULL,
                             brand VARCHAR(50) NOT NULL,
                             status INT NOT NULL
) ;


INSERT INTO order_items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES
    ('b563feb7b2b84b6test', 9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);


CREATE TABLE order_payment (
                               order_uid VARCHAR(20) NOT NULL,
                               transaction VARCHAR(20) NOT NULL,
                               request_id VARCHAR(20) NOT NULL,
                               currency VARCHAR(6) NOT NULL,
                               provider VARCHAR(10) NOT NULL,
                               amount INT NOT NULL,
                               payment_dt INT NOT NULL,
                               bank VARCHAR(30) NOT NULL,
                               delivery_cost INT NOT NULL,
                               goods_total INT NOT NULL,
                               custom_fee INT NOT NULL
);


INSERT INTO order_payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES
    ('b563feb7b2b84b6test', 'b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0);

ALTER TABLE orders
    ADD PRIMARY KEY (order_uid);

ALTER TABLE order_delivery
    ADD CONSTRAINT order_delivery FOREIGN KEY (order_uid) REFERENCES orders (order_uid) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE order_items
    ADD CONSTRAINT order_items FOREIGN KEY (order_uid) REFERENCES orders (order_uid) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE order_payment
    ADD CONSTRAINT order_payment FOREIGN KEY (order_uid) REFERENCES orders (order_uid) ON DELETE CASCADE ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_payment, order_items, order_delivery, orders;
-- +goose StatementEnd
