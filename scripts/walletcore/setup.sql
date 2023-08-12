use `walletcore`;

create table `customer` (
    `id` char(36) not null,
    `name` varchar(255) not null,
    `email` varchar(255) not null,
    `created_at` datetime not null,
    `updated_at` datetime not null,
    primary key (`id`)
);

create table `account` (
    `id` char(36) not null,
    `customer_id` char(36) not null,
    `balance` decimal(10, 5) not null,
    `created_at` datetime not null,
    `updated_at` datetime not null,
    primary key (`id`),
    foreign key (`customer_id`) references `customer`(`id`)
);