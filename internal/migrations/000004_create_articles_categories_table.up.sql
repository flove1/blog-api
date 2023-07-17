CREATE TABLE IF NOT EXISTS articles_categories (
    id bigserial PRIMARY KEY,
    article_id bigint NOT NULL REFERENCES articles ON DELETE CASCADE,
    category_id bigint NOT NULL REFERENCES categories ON DELETE CASCADE
)