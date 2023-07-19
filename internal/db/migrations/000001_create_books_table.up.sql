CREATE TABLE IF NOT EXISTS "book" (
    id VARCHAR(255) NOT NULL,
    title VARCHAR(255),
    description VARCHAR(255),
    author VARCHAR(255),
    price BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT "pk_book" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_book_id" ON "book" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_book_title_id" ON "book" USING btree("title");
CREATE INDEX IF NOT EXISTS "ix_book_author_id" ON "book" USING btree("author");
