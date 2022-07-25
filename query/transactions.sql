CREATE TABLE IF NOT EXISTS transactions (
    id  int auto_increment
        constraint transactions_pk
            primary key,
    tanggal_order      timestamp,
    status_pelunasan   varchar(255),
    tanggal_pembayaran timestamp
);

CREATE TABLE IF NOT EXISTS detail_transaction (
      id           int auto_increment
          constraint detail_transaction_pk
            primary key,
      id_transaksi int              null,
      harga        double precision null,
      jumlah       int              null,
      subtotal     double precision null,
);

create unique index transactions_id_uindex
    on transactions (id);

create unique index detail_transaction_id_transaksi_uindex
    on detail_transaction (id_transaksi);

create unique index detail_transaction_id_uindex
    on detail_transaction (id);

INSERT INTO transactions (tanggal_order, status_pelunasan, tanggal_pembayaran) VALUES ('2022-07-25 09:28:00', 'lunas', '2022-07-25 09:28:28');
INSERT INTO transactions (tanggal_order, status_pelunasan, tanggal_pembayaran) VALUES ('2022-07-25 09:28:00', 'pending', '2022-07-25 09:28:28');

INSERT INTO detail_transaction (id_transaksi, harga, jumlah, subtotal) VALUES (1, 1000, 1, 1000);
INSERT INTO detail_transaction (id_transaksi, harga, jumlah, subtotal) VALUES (1, 2000, 2, 4000);

SELECT t.id, t.tanggal_order, t.status_pelunasan as status, t.tanggal_pembayaran, dt.subtotal as total, dt.jumlah as jumlah_barang
from transactions t
         left join detail_transaction dt ON dt.id_transaksi = t.id;