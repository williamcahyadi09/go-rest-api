CREATE TABLE IF NOT EXISTS "book_history_log" (
    id VARCHAR(255) NOT NULL,
    book_id VARCHAR(255) NOT NULL,
    title VARCHAR(255),
    description VARCHAR(255),
    author VARCHAR(255),
    price BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "pk_book_history_log" PRIMARY KEY ("id")
);