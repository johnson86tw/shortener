create table urls (
	url_id	serial not null unique,
	url	text,
	code varchar(12) unique,
	created_at timestamp default now()
);

create table users (
	user_id uuid,
	name varchar(64),
	email varchar(64) not null unique,
	password varchar(64) not null ,
	created_at timestamp not null default now(),
	updated_at timestamp null,
	deleted_at timestamp null,
	primary key (user_id)
);