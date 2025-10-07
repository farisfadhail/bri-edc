create table merchants
(
    id         int not null,
    merchant_id       varchar(100) not null unique,
    name   varchar(100) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    primary key (id)
) engine = InnoDB;