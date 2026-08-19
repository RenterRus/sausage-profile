-- +goose Up
create table if not exists users (
  	user_login text primary key,
    uuid uuid not null default gen_random_uuid(),
    otp_hash text not null,
    otp_link text not null,
    
    created_at timestamp default now(),
    confirmed boolean default false    
)


-- +goose Down

drop table if exists users;
