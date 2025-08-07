-- migrate:up
ALTER TABLE books
  ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

-- migrate:down
ALTER TABLE books
  DROP COLUMN created_at,
  DROP COLUMN updated_at;
