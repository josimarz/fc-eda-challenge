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

-- Customer 1

set @customerId := uuid();

insert into `customer` (
    `id`,
    `name`,
    `email`,
    `created_at`,
    `updated_at`
)
values
    (@customerId, "Josimar Zimermann", "josimarz@yahoo.com.br", current_timestamp, current_timestamp);

insert into `account` (
    `id`,
    `customer_id`,
    `balance`,
    `created_at`,
    `updated_at`
)
values
    (uuid(), @customerId, 2000.0, current_timestamp, current_timestamp);

-- Customer 2

set @customerId := uuid();

insert into `customer` (
    `id`,
    `name`,
    `email`,
    `created_at`,
    `updated_at`
)
values
    (@customerId, "Gustavo Kuerten", "guga@tennis.com", current_timestamp, current_timestamp);

insert into `account` (
    `id`,
    `customer_id`,
    `balance`,
    `created_at`,
    `updated_at`
)
values
    (uuid(), @customerId, 1000.0, current_timestamp, current_timestamp);

-- Customer 3

set @customerId := uuid();

insert into `customer` (
    `id`,
    `name`,
    `email`,
    `created_at`,
    `updated_at`
)
values
    (@customerId, "Ana Ivanovic", "ivanovic@wta.com", current_timestamp, current_timestamp);

insert into `account` (
    `id`,
    `customer_id`,
    `balance`,
    `created_at`,
    `updated_at`
)
values
    (uuid(), @customerId, 5000.0, current_timestamp, current_timestamp);

-- Customer 4

set @customerId := uuid();

insert into `customer` (
    `id`,
    `name`,
    `email`,
    `created_at`,
    `updated_at`
)
values
    (@customerId, "Maria Sharapova", "sharapova@wta.com", current_timestamp, current_timestamp);

insert into `account` (
    `id`,
    `customer_id`,
    `balance`,
    `created_at`,
    `updated_at`
)
values
    (uuid(), @customerId, 500.0, current_timestamp, current_timestamp);