-- +migrate Up

CREATE TABLE todos (
  id         SERIAL,
  name       VARCHAR(255) NOT NULL,
  done       BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id)
);

-- +migrate Down

DROP TABLE IF EXISTS todos;
