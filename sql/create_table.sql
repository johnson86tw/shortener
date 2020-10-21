create table urls (
	url_id	serial not null unique,
	url	text,
	code varchar(12) unique,
	created_at timestamp default now()
);

-- 新增 url
INSERT INTO public.urls (url, code) VALUES('http://www.fake.com', 'abc');


-- 安裝產 uuid 的套件
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

-- 更改 users name 為 not null
update users set name = '' where name is null;
alter table users alter column name set not null;

-- 新增 user
INSERT INTO public.users
(name, email, password)
VALUES('ken', 'ken@gmail.com', '123');

-- 更改欄位預設值
ALTER TABLE users ALTER COLUMN user_id SET DEFAULT uuid_generate_v4();


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

-- 新增 user_urls
INSERT INTO user_urls (url, code, user_id) 
VALUES('http://www.facebook.com', 'aiofjf', '4425ff13-354f-4e45-897f-ac76476305d5');

