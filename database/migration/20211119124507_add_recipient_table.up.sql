create table blockchain.recipient
(
    id int,
    user_id int not null,
    constraint recipient_user_id_fk
        foreign key (user_id) references blockchain.user (id)
);

create unique index recipient_id_uindex
    on blockchain.recipient (id);

create unique index recipient_user_id_uindex
    on blockchain.recipient (user_id);

alter table blockchain.recipient
    add constraint recipient_pk
        primary key (id);

alter table blockchain.recipient modify id int auto_increment;

