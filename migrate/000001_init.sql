create table if not exists account
(
    id serial not null constraint account_pk primary key,
    email text not null,
    password_hash text not null,
    role integer not null default 0,
    created_at timestamp with time zone default now()
);

create unique index if not exists account_email_uindex on account (email);
