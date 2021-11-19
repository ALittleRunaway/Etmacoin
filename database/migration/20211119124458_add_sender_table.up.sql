create table blockchain.sender
(
    id int,
    user_id int not null,
    constraint sender_user_id_fk
        foreign key (user_id) references blockchain.user (id)
);

create unique index sender_id_uindex
    on blockchain.sender (id);

create unique index sender_user_id_uindex
    on blockchain.sender (user_id);

alter table blockchain.sender
    add constraint sender_pk
        primary key (id);

alter table blockchain.sender modify id int auto_increment;
