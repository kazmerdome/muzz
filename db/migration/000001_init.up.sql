CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS decisions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    actor_user_id uuid NOT NULL,
    recipient_user_id uuid NOT NULL,
    liked_recipient BOOLEAN NOT NULL,
    recipient_likes_actor BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(actor_user_id, recipient_user_id)
);
