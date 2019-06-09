-- +migrate Up

CREATE TABLE todos (
  id         SERIAL,
  name       varchar(255) NOT NULL,
  done       boolean DEFAULT FALSE,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

-- +migrate Down

DROP TABLE todos;
