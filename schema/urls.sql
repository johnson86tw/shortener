create table urls (
	url_id	serial not null unique,
	url	text,
	code varchar(12) unique,
	created_at	timestamp default now()
);