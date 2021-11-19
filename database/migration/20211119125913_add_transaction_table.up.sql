create table blockchain.transaction
(
    id int,
    message NVARCHAR(250) null,
    sender_id int not null,
    recipient_id int not null,
    amount INT(30) not null,
    time datetime default current_timestamp not null,
    prev_hash VARCHAR(300) not null,
    pow INT(30) not null,
    constraint transaction_recipient_id_fk
        foreign key (recipient_id) references blockchain.recipient (id),
    constraint transaction_sender_id_fk
        foreign key (sender_id) references blockchain.sender (id)
);

create unique index transaction_id_uindex
    on blockchain.transaction (id);

create unique index transaction_prev_hash_uindex
    on blockchain.transaction (prev_hash);

alter table blockchain.transaction
    add constraint transaction_pk
        primary key (id);

alter table blockchain.transaction modify id int auto_increment;

