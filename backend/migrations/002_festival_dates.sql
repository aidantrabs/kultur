-- +goose Up

-- create festival_dates table
CREATE TABLE festival_dates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    festival_id UUID NOT NULL REFERENCES festivals(id) ON DELETE CASCADE,
    year INTEGER NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    is_tentative BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(festival_id, year)
);

CREATE INDEX idx_festival_dates_year ON festival_dates(year);
CREATE INDEX idx_festival_dates_start ON festival_dates(start_date);

-- remove date columns from festivals
ALTER TABLE festivals DROP COLUMN IF EXISTS date_2026_start;
ALTER TABLE festivals DROP COLUMN IF EXISTS date_2026_end;
ALTER TABLE festivals DROP COLUMN IF EXISTS usual_month;

-- +goose Down

-- restore columns
ALTER TABLE festivals ADD COLUMN usual_month VARCHAR(50);
ALTER TABLE festivals ADD COLUMN date_2026_start DATE;
ALTER TABLE festivals ADD COLUMN date_2026_end DATE;

-- drop festival_dates table
DROP TABLE IF EXISTS festival_dates;
