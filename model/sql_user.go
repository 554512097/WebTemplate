package model

const sql_user_create = `create table if not exists user (
	id bigint primary key auto_increment,
	nick varchar(30) not null,
	age tinyint default 0,
	phone varchar(20),
	account varchar(50) not null,
	passwrod varchar(30) not null
)`

const sql_user_add = ``
