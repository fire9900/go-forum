CREATE TABLE IF NOT EXISTS threads
(
    id        serial PRIMARY KEY,
    title     TEXT     NOT NULL,
    content   TEXT     NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id   INTEGER  NOT NULL
);

CREATE INDEX idx_threads_user_id ON threads (user_id);
CREATE INDEX idx_threads_create_at ON threads (create_at DESC);