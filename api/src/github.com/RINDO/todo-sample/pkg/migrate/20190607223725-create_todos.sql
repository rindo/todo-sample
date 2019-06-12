-- +migrate Up

CREATE TABLE todos (
  id         SERIAL,
  name       varchar(255) NOT NULL,
  done       boolean NOT NULL DEFAULT FALSE,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id)
);

-- +migrate Down

DROP TABLE IF EXISTS todos;