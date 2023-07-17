CREATE TABLE IF NOT EXISTS categories (
	id bigserial PRIMARY KEY,
	category_name varchar(50) UNIQUE NOT NULL CONSTRAINT category_name_not_empty CHECK(length(category_name)>0)
)