create table twit
(
    id      uuid not null default gen_random_uuid(),
    author_id uuid not null,
    text     text not null,
    created_at time.Time,
    updated_at time.Time,

    primary key (id)
)