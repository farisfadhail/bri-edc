create table transactions
(
    id         int auto_increment not null,
    transaction_id       varchar(100) not null unique,
    merchant_id       varchar(100) not null,
    terminal_id       varchar(100) not null,
    amount int not null,
    card_number varchar(20) not null,
    status varchar(50) not null,
    timestamp timestamp not null,
    is_settled BOOLEAN DEFAULT false,
    hmac varchar(255) not null,
    created_at timestamp default now(),
    primary key (id),

    foreign key (merchant_id) references merchants(merchant_id) on delete cascade,
    foreign key (terminal_id) references terminals(terminal_id) on delete cascade
) engine = InnoDB;