-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create tables
CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    cases_solved INTEGER DEFAULT 0
);

CREATE TABLE cases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    solution TEXT NOT NULL
);

CREATE TABLE player_cases (
    player_id UUID REFERENCES players(id),
    case_id UUID REFERENCES cases(id),
    status TEXT NOT NULL,
    PRIMARY KEY (player_id, case_id)
);

CREATE TABLE evidence (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    case_id UUID REFERENCES cases(id),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL
);

CREATE TABLE suspects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    case_id UUID REFERENCES cases(id),
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE interrogation_questions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    suspect_id UUID REFERENCES suspects(id),
    question TEXT NOT NULL,
    answer TEXT NOT NULL
);

-- Insert sample data for "The Stolen Golden Fish Statue" case
INSERT INTO cases (id, title, description, solution) VALUES 
(uuid_generate_v4(), 'The Case of the Stolen Golden Fish', 'The Otter Museum''s prized possession, a solid gold statue of an otter holding a fish, has been stolen overnight. The statue, valued at $1 million, was the centerpiece of the museum''s "Aquatic Treasures" exhibit. The theft occurred despite the museum''s state-of-the-art security system. Your task is to investigate the crime scene, gather evidence, and interrogate suspects to solve this fishy case!', 'Finn Waters');

-- Get the case ID for use in subsequent inserts
DO $$
DECLARE
    case_id UUID;
BEGIN
    SELECT id INTO case_id FROM cases WHERE title = 'The Case of the Stolen Golden Fish';

    INSERT INTO evidence (case_id, name, description, location) VALUES 
    (case_id, 'Muddy Footprint', 'A single muddy footprint found just inside the museum entrance. It appears to be from a size 11 boot.', 'Museum Entrance'),
    (case_id, 'Deactivated Alarm System', 'The security system was deactivated at 2:37 AM using the correct access code.', 'Security Room'),
    (case_id, 'Broken Glass Case', 'The glass case containing the statue has been carefully cut, suggesting professional tools were used.', 'Exhibit Hall'),
    (case_id, 'Wet Umbrella', 'A damp umbrella found in the corner, despite no rain being reported that night.', 'Staff Lounge'),
    (case_id, 'Tire Tracks', 'Fresh tire tracks leading away from the staff parking area, indicating a hasty departure.', 'Parking Lot');

    INSERT INTO suspects (id, case_id, name, description) VALUES 
    (uuid_generate_v4(), case_id, 'Dr. Olivia Whiskers', 'The museum''s curator, known for her extensive knowledge of aquatic artifacts.'),
    (uuid_generate_v4(), case_id, 'Finn Waters', 'The newly hired night guard, started working at the museum two weeks ago.'),
    (uuid_generate_v4(), case_id, 'Marina Shells', 'A wealthy collector who has previously expressed interest in acquiring the statue.');

    -- Get suspect IDs for use in interrogation questions
    WITH suspect_ids AS (
        SELECT id, name FROM suspects WHERE case_id = case_id
    )
    INSERT INTO interrogation_questions (suspect_id, question, answer)
    SELECT 
        id,
        CASE 
            WHEN name = 'Dr. Olivia Whiskers' THEN unnest(ARRAY[
                'Where were you on the night of the theft?',
                'Do you know anyone who might want to steal the statue?',
                'Have you noticed anything unusual at the museum lately?'
            ])
            WHEN name = 'Finn Waters' THEN unnest(ARRAY[
                'Can you explain your whereabouts during the theft?',
                'Did you notice anything out of the ordinary during your shift?',
                'We found a wet umbrella in the staff lounge. Do you know anything about it?'
            ])
            WHEN name = 'Marina Shells' THEN unnest(ARRAY[
                'Where were you on the night of the theft?',
                'Have you ever tried to purchase the Golden Fish Statue?',
                'Do you know anyone who might have wanted to steal the statue?'
            ])
        END,
        CASE 
            WHEN name = 'Dr. Olivia Whiskers' THEN unnest(ARRAY[
                'I was at home, working on my upcoming book about otters in art.',
                'No, but there are always collectors willing to pay for such unique pieces.',
                'Now that you mention it, our night guard, Finn, has been asking a lot of questions about the statue''s value.'
            ])
            WHEN name = 'Finn Waters' THEN unnest(ARRAY[
                'I was patrolling the museum as usual. I didn''t see or hear anything suspicious.',
                'No, everything seemed normal. The alarm system was working fine when I checked it.',
                'Oh, that''s probably mine. I always keep one handy in case of rain, even if it''s not in the forecast.'
            ])
            WHEN name = 'Marina Shells' THEN unnest(ARRAY[
                'I was attending a charity gala across town. Hundreds of people can vouch for my presence there.',
                'Yes, I offered to buy it last year, but the museum refused to sell. It''s a magnificent piece.',
                'In my experience, museum insiders are often involved in these kinds of thefts. Have you looked into the staff?'
            ])
        END
    FROM suspect_ids;
END $$;
