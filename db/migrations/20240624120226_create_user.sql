-- migrate:up
CREATE TABLE IF NOT EXISTS users (
    id           BIGSERIAL                NOT NULL PRIMARY KEY,
    name         TEXT                     NOT NULL,
    email        TEXT                     NOT NULL,
    password     TEXT                     ,
    role         TEXT                     NOT NULL,
    is_active    BOOLEAN                  NOT NULL DEFAULT true,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_uniq_users_email ON users (email);

-- migrate:down
DROP TABLE IF EXISTS users;
