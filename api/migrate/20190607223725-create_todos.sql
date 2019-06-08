-- +migrate Up

CREATE TABLE todos (
  id         SERIAL,
  name       varchar(255),
  done       boolean,
  created_at timestamp,
  updated_at timestamp
);

-- +migrate Down

DROP TABLE todos;
