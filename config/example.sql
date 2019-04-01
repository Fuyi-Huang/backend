create database test;

create table appid_key (
    `id` int(11) unsigned not null auto_increment comment '',
    `appid` varchar(30) not null comment '',
    `key` varchar(30) not null comment '',
    `add_time` datetime not null comment '',
    primary key(`id`),
    unique key(`appid`),
    key `idx_add_time`(`add_time`)
) ENGINE=InnoDB default charset=utf8;

