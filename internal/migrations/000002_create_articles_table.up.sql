CREATE TABLE IF NOT EXISTS articles (
	id bigserial PRIMARY KEY,
	title varchar(100) NOT NULL CONSTRAINT category_name_not_empty CHECK(length(category_name)>0),
    description text NOT NULL,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE
)