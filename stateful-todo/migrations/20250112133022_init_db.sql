-- +goose Up
-- +goose StatementBegin
create table if not exists tasks
(
    id      bigserial
        constraint tasks_pk
            primary key,
    title   varchar(50) not null,
    description   varchar(200) not null
) partition by hash (id);

create table if not exists packages_1 partition of tasks
    for values with (modulus 3, remainder 0);

create table if not exists packages_2 partition of tasks
    for values with (modulus 3, remainder 1);

create table if not exists packages_3 partition of tasks
    for values with (modulus 3, remainder 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tasks;
-- +goose StatementEnd
