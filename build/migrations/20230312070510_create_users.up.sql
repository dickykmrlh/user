create type user_role as enum ('super_admin', 'customer');

create table "users"
(
    "id" uuid not null primary key,
    "first_name" text not null,
    "last_name" text not null,
    "role" user_role not null,
    "phone_number" text not null default '',
    "email" text not null default '',
    "created_at" timestamptz not null default now(),
    "updated_at" timestamptz not null default now()
);