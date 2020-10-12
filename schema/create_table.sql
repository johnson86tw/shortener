create table urls (
	url_id	serial not null unique,
	url	text,
	code varchar(12) unique,
	created_at timestamp default now()
);

-- 新增 url
INSERT INTO public.urls (url, code) VALUES('http://www.fake.com', 'abc');

create table users (
	user_id uuid not null default uuid_generate_v4(),
	name varchar(64),
	email varchar(64) not null unique,
	password varchar(64) not null ,
	created_at timestamp not null default now(),
	updated_at timestamp null,
	deleted_at timestamp null,
	primary key (user_id)
);

-- 安裝產 uuid 的套件
create extension "uuid-ossp";

-- 新增 user
INSERT INTO public.users
(name, email, password)
VALUES('ken', 'ken@gmail.com', '123');

-- 更改欄位預設值
ALTER TABLE users ALTER COLUMN user_id SET DEFAULT uuid_generate_v4();

