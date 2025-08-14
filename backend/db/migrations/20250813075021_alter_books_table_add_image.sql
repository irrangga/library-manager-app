-- migrate:up
ALTER TABLE books
  ADD COLUMN image_url TEXT;

-- migrate:down
ALTER TABLE books
  DROP COLUMN image_url;
