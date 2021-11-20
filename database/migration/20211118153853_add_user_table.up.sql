create table blockchain.user
(
    id int,
    login NVARCHAR(20) not null,
    password VARCHAR(100) not null,
    wallet VARCHAR(40) not null,
    balance INT(30) null
);

create unique index user_id_uindex
    on blockchain.user (id);

create unique index user_wallet_uindex
    on blockchain.user (wallet);

alter table blockchain.user
    add constraint user_pk
        primary key (id);

alter table blockchain.user modify id int auto_increment;

