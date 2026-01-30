-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE festivals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    date_type VARCHAR(20) NOT NULL,
    usual_month VARCHAR(50),
    date_2026_start DATE,
    date_2026_end DATE,
    region VARCHAR(50) NOT NULL,
    heritage_type VARCHAR(50) NOT NULL,
    festival_type VARCHAR(50) NOT NULL,
    summary TEXT NOT NULL,
    story TEXT,
    what_to_expect TEXT,
    how_to_participate TEXT,
    practical_info TEXT,
    cover_image_url TEXT,
    gallery_images JSONB DEFAULT '[]',
    video_embeds JSONB DEFAULT '[]',
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE memories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    festival_id UUID REFERENCES festivals(id) ON DELETE CASCADE,
    author_name VARCHAR(100),
    author_email VARCHAR(255),
    content TEXT NOT NULL,
    year_of_memory VARCHAR(20),
    status VARCHAR(20) DEFAULT 'pending',
    submitted_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    digest_weekly BOOLEAN DEFAULT FALSE,
    festival_reminders JSONB DEFAULT '[]',
    confirmed BOOLEAN DEFAULT FALSE,
    confirmation_token VARCHAR(100),
    unsubscribe_token VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS memories;
DROP TABLE IF EXISTS festivals;
