create table terminals
(
    id         int auto_increment not null,
    merchant_id       varchar(100) not null,
    terminal_id       varchar(100) not null unique,
    location   varchar(100) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp null,
    primary key (id),

    foreign key (merchant_id) references merchants(merchant_id) on delete cascade
) engine = InnoDB;