CREATE TABLE IF NOT EXISTS posts
(
    id        serial PRIMARY KEY,
    content   TEXT     NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    thread_id INTEGER  NOT NULL,
    user_id   INTEGER  NOT NULL,
    parent_id INTEGER  DEFAULT NULL
);

-- Создаем индексы для оптимизации запросов
CREATE INDEX idx_posts_thread_id ON posts (thread_id);
CREATE INDEX idx_posts_user_id ON posts (user_id);
CREATE INDEX idx_posts_create_at ON posts (create_at DESC);
CREATE INDEX idx_posts_parent_id ON posts (parent_id);

-- Добавляем внешние ключи для обеспечения целостности данных
ALTER TABLE posts ADD CONSTRAINT fk_posts_thread_id
    FOREIGN KEY (thread_id) REFERENCES threads(id) ON DELETE CASCADE;

ALTER TABLE posts ADD CONSTRAINT fk_posts_parent_id
    FOREIGN KEY (parent_id) REFERENCES posts(id) ON DELETE CASCADE;
