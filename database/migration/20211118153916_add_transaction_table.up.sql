create table transaction
(
    id int,
    sender_id INT not null,
    recipient_id int not null,
    amount INT(30) not null,
    message NVARCHAR(100) null,
    time DATETIME not null,
    prev_hash VARCHAR(100) not null,
    pow INT(30) not null
);

create unique index transaction_id_uindex
    on transaction (id);

alter table transaction
    add constraint transaction_pk
        primary key (id);

alter table transaction modify id int auto_increment;

alter table transaction
    add constraint transaction_user_id_fk
        foreign key (sender_id) references user (id);

alter table transaction
    add constraint transaction_user_id_fk_2
        foreign key (recipient_id) references user (id);