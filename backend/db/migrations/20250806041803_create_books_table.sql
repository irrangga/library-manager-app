-- migrate:up
CREATE TABLE books(
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  publisher VARCHAR(255),
  year INTEGER
);

-- migrate:down
DROP TABLE IF EXISTS books;
