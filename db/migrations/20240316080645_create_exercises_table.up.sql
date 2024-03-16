CREATE TABLE IF NOT EXISTS exercises(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    video_url VARCHAR,
    notes TEXT,
    primary_muscle_group VARCHAR(100),
    secondary_muscle_groups VARCHAR(100)[],
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
