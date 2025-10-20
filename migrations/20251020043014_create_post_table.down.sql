-- Удаляем внешние ключи
ALTER TABLE posts DROP CONSTRAINT IF EXISTS fk_posts_parent_id;
ALTER TABLE posts DROP CONSTRAINT IF EXISTS fk_posts_thread_id;

-- Удаляем индексы
DROP INDEX IF EXISTS idx_posts_parent_id;
DROP INDEX IF EXISTS idx_posts_create_at;
DROP INDEX IF EXISTS idx_posts_user_id;
DROP INDEX IF EXISTS idx_posts_thread_id;

-- Удаляем таблицу
DROP TABLE IF EXISTS posts;
