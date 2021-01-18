create table users
(
    id           int auto_increment
        primary key,
    first_name   varchar(45)                        null,
    last_name    varchar(45)                        null,
    email        varchar(45)                        not null,
    username        varchar(45)                        not null,
    date_created datetime default CURRENT_TIMESTAMP null,
    status       varchar(45)                        null,
    password     varchar(45)                        not null,
    constraint email_UNIQUE
        unique (email)
);
