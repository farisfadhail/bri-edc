create table settlements
(
    id         int not null,
    batch_id       varchar(100) not null unique,
    total_count       int not null,
    approved int not null,
    declined int not null,
    total_amount int not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    primary key (id)
) engine = InnoDB;