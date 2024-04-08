create table twit
(
    ID       uuid not null default gen_random_uuid(),
    AuthorID uuid not null,
    Text     text not null,
    CreatedAt time.Time,
    UpdatedAt time.Time,

    primary key (ID)
)