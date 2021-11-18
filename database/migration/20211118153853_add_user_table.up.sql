create table user
(
    id int,
    login NVARCHAR(20) not null,
    password VARCHAR(100) not null,
    balance INT(30) null
);

create unique index user_id_uindex
    on user (id);

create unique index user_login_uindex
    on user (login);

alter table user
    add constraint user_pk
        primary key (id);

alter table user modify id int auto_increment;

