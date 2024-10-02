-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Drop existing tables if they exist
DROP TABLE IF EXISTS interrogation_questions;
DROP TABLE IF EXISTS suspects;
DROP TABLE IF EXISTS evidence;
DROP TABLE IF EXISTS player_cases;
DROP TABLE IF EXISTS cases;
DROP TABLE IF EXISTS players;

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
    location TEXT NOT NULL,
    analysis_result TEXT
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
DO $$
DECLARE
    case_var UUID;
    suspect_id UUID;
BEGIN
    INSERT INTO cases (id, title, description, solution) VALUES 
    (uuid_generate_v4(), 'The Case of the Stolen Golden Fish', 'The Otter Museum''s prized possession, a solid gold statue of an otter holding a fish, has been stolen overnight. The statue, valued at $1 million, was the centerpiece of the museum''s "Aquatic Treasures" exhibit. The theft occurred despite the museum''s state-of-the-art security system. Your task is to investigate the crime scene, gather evidence, and interrogate suspects to solve this fishy case!', 'Finn Waters')
    RETURNING id INTO case_var;

    INSERT INTO evidence (case_id, name, description, location, analysis_result) VALUES 
    (case_var, 'muddy footprint', 'A single muddy footprint found just inside the museum entrance. It appears to be from a size 10 boot.', 'museum entrance', 'The mud contains traces of a rare mineral found only in a nearby construction site where Finn Waters was seen recently.'),
    (case_var, 'deactivated alarm system', 'The security system was deactivated at 2:37 AM using the correct access code.', 'security room', 'Log analysis shows multiple failed attempts before success, suggesting someone familiar with the system but not using it regularly.'),
    (case_var, 'broken glass case', 'The glass case containing the statue has been carefully cut, suggesting professional tools were used.', 'exhibit hall', 'Microscopic analysis reveals the use of a common glass cutter, available at most hardware stores. No specialized knowledge required.'),
    (case_var, 'wet umbrella', 'A damp umbrella found in the corner, despite no rain being reported that night.', 'staff lounge', 'DNA analysis of the umbrella handle matches Finn Waters, contradicting his statement about not using it recently.'),
    (case_var, 'tire tracks', 'Fresh tire tracks leading away from the staff parking area, indicating a hasty departure.', 'parking lot', 'Tire tread analysis matches a standard issue museum vehicle, eliminating the possibility of an outside vehicle being used.');

    INSERT INTO suspects (id, case_id, name, description) VALUES 
    (uuid_generate_v4(), case_var, 'Dr. Olivia Whiskers', 'The museum''s curator, known for her extensive knowledge of aquatic artifacts.'),
    (uuid_generate_v4(), case_var, 'Finn Waters', 'The newly hired night guard, started working at the museum two weeks ago.'),
    (uuid_generate_v4(), case_var, 'Marina Shells', 'A wealthy collector who has previously expressed interest in acquiring the statue.');

    -- Insert interrogation questions for each suspect
    SELECT id INTO suspect_id FROM suspects WHERE name = 'Dr. Olivia Whiskers' AND case_id = case_var;
    INSERT INTO interrogation_questions (suspect_id, question, answer) VALUES
    (suspect_id, 'Where were you on the night of the theft?', 'I was at home, working on my upcoming book about otters in art.'),
    (suspect_id, 'Do you know anyone who might want to steal the statue?', 'No, but there are always collectors willing to pay for such unique pieces.'),
    (suspect_id, 'Have you noticed anything unusual at the museum lately?', 'Now that you mention it, our night guard, Finn, has been asking a lot of questions about the statue''s value.');

    SELECT id INTO suspect_id FROM suspects WHERE name = 'Finn Waters' AND case_id = case_var;
    INSERT INTO interrogation_questions (suspect_id, question, answer) VALUES
    (suspect_id, 'Can you explain your whereabouts during the theft?', 'I was patrolling the museum as usual. I didn''t see or hear anything suspicious.'),
    (suspect_id, 'Did you notice anything out of the ordinary during your shift?', 'No, everything seemed normal. The alarm system was working fine when I checked it.'),
    (suspect_id, 'We found a wet umbrella in the staff lounge. Do you know anything about it?', 'Oh, that''s probably mine. I always keep one handy in case of rain, even if it''s not in the forecast.');

    SELECT id INTO suspect_id FROM suspects WHERE name = 'Marina Shells' AND case_id = case_var;
    INSERT INTO interrogation_questions (suspect_id, question, answer) VALUES
    (suspect_id, 'Where were you on the night of the theft?', 'I was attending a charity gala across town. Hundreds of people can vouch for my presence there.'),
    (suspect_id, 'Have you ever tried to purchase the Golden Fish Statue?', 'Yes, I offered to buy it last year, but the museum refused to sell. It''s a magnificent piece.'),
    (suspect_id, 'Do you know anyone who might have wanted to steal the statue?', 'In my experience, museum insiders are often involved in these kinds of thefts. Have you looked into the staff?');
END $$;
