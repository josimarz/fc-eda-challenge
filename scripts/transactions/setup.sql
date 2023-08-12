use `transactions`;

create table `transaction` (
    `id` char(36) not null,
    `from_id` char(36) not null,
    `to_id` char(36) not null,
    `amount` decimal(10, 5) not null,
    `created_at` datetime not null,
    `updated_at` datetime not null,
    primary key (`id`)
);