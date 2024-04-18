create table twit
(
    ID       uuid not null default gen_random_uuid(),
    author_id uuid not null,
    text     text not null,
    createdAt time.Time,
    updatedAt time.Time,

    primary key (ID)
)