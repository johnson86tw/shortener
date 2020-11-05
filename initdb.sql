\c postgres

create table urls (
	url_id	serial not null unique,
	url	text,
	code varchar(12) unique,
	created_at timestamp default now()
);

create extension "uuid-ossp";

create table users (
	user_id uuid not null default uuid_generate_v4(),
	name varchar(64) not null,
	email varchar(64) not null unique,
	password varchar(64) not null ,
	created_at timestamp not null default now(),
	updated_at timestamp null,
	deleted_at timestamp null,
	primary key (user_id)
);


create table user_urls (
	id serial not null unique,
	url text,
	code varchar(12) unique,
	created_at timestamp not null default now(),
	updated_at timestamp null,
	deleted_at timestamp null,
	total_click integer not null default 0,
	user_id uuid,
	constraint fk_owner foreign key(user_id) references users(user_id)
)