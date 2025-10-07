create table users
(
    id         int not null,
    username      varchar(100) not null unique,
    password      varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    primary key (id)
) engine = InnoDB;