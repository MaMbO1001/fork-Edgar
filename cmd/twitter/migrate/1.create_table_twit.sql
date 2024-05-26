-- up
create table twit
(
    id         uuid      not null default gen_random_uuid(),
    author_id  uuid      not null,
    text       text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    is_banned  boolean   not null default true,
    primary key (id)
);

-- down
drop table twit;