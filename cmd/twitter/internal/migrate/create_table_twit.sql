create table twit
(
    ID       uuid not null default gen_random_uuid(),
    AuthorID uuid not null,
    Text     text not null,

    unique (Text),
    primary key (ID),
    primary key (AuthorID)
)