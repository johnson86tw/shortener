create table users (
    user_id uuid,
    email character varying(200) not null,
    password character varying(1000) not null,
    name varchar(12) varying(255) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);