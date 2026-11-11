CREATE TABLE target_table (
  id     BIGINT PRIMARY KEY,
  active BOOLEAN NOT NULL
);

CREATE TABLE source_table (
  id         BIGINT PRIMARY KEY,
  updated_at TIMESTAMP NOT NULL
);
