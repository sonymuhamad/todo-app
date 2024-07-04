-- migrate:up
CREATE TABLE IF NOT EXISTS user_sessions(
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL REFERENCES users(id),
    login_time      TIMESTAMP WITH TIME ZONE NOT NULL,
    access_token    TEXT NOT NULL,
    expired_at      TIMESTAMP WITH TIME ZONE,
    last_seen_time  TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_user_sessions_user_id ON user_sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_user_sessions_access_token ON user_sessions(access_token);

-- migrate:down
DROP TABLE IF EXISTS user_sessions;

