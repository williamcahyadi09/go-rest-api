CREATE TABLE IF NOT EXISTS "tweets" (
    id VARCHAR(255) NOT NULL,
    content VARCHAR(255),
    likes INT,
    user_id VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,

    CONSTRAINT "pk_tweets" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_tweets_id" ON "tweets" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_tweets_user_id" ON "tweets" USING btree("user_id");
