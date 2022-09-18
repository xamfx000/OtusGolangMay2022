create table events
(
    id           text primary key,
    title        text   not null,
    start_time   bigint not null,
    end_time     bigint not null,
    description  text,
    owner_id     int    not null,
    alert_before bigint
);

create index if not exists events_by_owner on events (owner_id, start_time)
