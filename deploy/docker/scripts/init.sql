-- Create tables
CREATE TABLE IF NOT EXISTS players (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    cases_solved INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS cases (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    solution TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS player_cases (
    player_id UUID REFERENCES players(id),
    case_id UUID REFERENCES cases(id),
    status TEXT NOT NULL,
    PRIMARY KEY (player_id, case_id)
);

CREATE TABLE IF NOT EXISTS evidence (
    id UUID PRIMARY KEY,
    case_id UUID REFERENCES cases(id),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS suspects (
    id UUID PRIMARY KEY,
    case_id UUID REFERENCES cases(id),
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS interrogation_questions (
    id UUID PRIMARY KEY,
    suspect_id UUID REFERENCES suspects(id),
    question TEXT NOT NULL,
    answer TEXT NOT NULL
);

-- Insert sample data for "The Stolen Golden Fish Statue" case
INSERT INTO cases (id, title, description, solution) VALUES 
('123e4567-e89b-12d3-a456-426614174000', 'The Case of the Stolen Golden Fish', 'The Otter Museum''s prized possession, a solid gold statue of an otter holding a fish, has been stolen overnight. The statue, valued at $1 million, was the centerpiece of the museum''s "Aquatic Treasures" exhibit. The theft occurred despite the museum''s state-of-the-art security system. Your task is to investigate the crime scene, gather evidence, and interrogate suspects to solve this fishy case!', 'Finn Waters');

INSERT INTO evidence (id, case_id, name, description, location) VALUES 
('223e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Muddy Footprint', 'A single muddy footprint found just inside the museum entrance. It appears to be from a size 11 boot.', 'Museum Entrance'),
('323e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Deactivated Alarm System', 'The security system was deactivated at 2:37 AM using the correct access code.', 'Security Room'),
('423e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Broken Glass Case', 'The glass case containing the statue has been carefully cut, suggesting professional tools were used.', 'Exhibit Hall'),
('523e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Wet Umbrella', 'A damp umbrella found in the corner, despite no rain being reported that night.', 'Staff Lounge'),
('623e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Tire Tracks', 'Fresh tire tracks leading away from the staff parking area, indicating a hasty departure.', 'Parking Lot');

INSERT INTO suspects (id, case_id, name, description) VALUES 
('723e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Dr. Olivia Whiskers', 'The museum''s curator, known for her extensive knowledge of aquatic artifacts.'),
('823e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Finn Waters', 'The newly hired night guard, started working at the museum two weeks ago.'),
('923e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Marina Shells', 'A wealthy collector who has previously expressed interest in acquiring the statue.');

INSERT INTO interrogation_questions (id, suspect_id, question, answer) VALUES 
('a23e4567-e89b-12d3-a456-426614174000', '723e4567-e89b-12d3-a456-426614174000', 'Where were you on the night of the theft?', 'I was at home, working on my upcoming book about otters in art.'),
('b23e4567-e89b-12d3-a456-426614174000', '723e4567-e89b-12d3-a456-426614174000', 'Do you know anyone who might want to steal the statue?', 'No, but there are always collectors willing to pay for such unique pieces.'),
('c23e4567-e89b-12d3-a456-426614174000', '723e4567-e89b-12d3-a456-426614174000', 'Have you noticed anything unusual at the museum lately?', 'Now that you mention it, our night guard, Finn, has been asking a lot of questions about the statue''s value.'),
('d23e4567-e89b-12d3-a456-426614174000', '823e4567-e89b-12d3-a456-426614174000', 'Can you explain your whereabouts during the theft?', 'I was patrolling the museum as usual. I didn''t see or hear anything suspicious.'),
('e23e4567-e89b-12d3-a456-426614174000', '823e4567-e89b-12d3-a456-426614174000', 'Did you notice anything out of the ordinary during your shift?', 'No, everything seemed normal. The alarm system was working fine when I checked it.'),
('f23e4567-e89b-12d3-a456-426614174000', '823e4567-e89b-12d3-a456-426614174000', 'We found a wet umbrella in the staff lounge. Do you know anything about it?', 'Oh, that''s probably mine. I always keep one handy in case of rain, even if it''s not in the forecast.'),
('g23e4567-e89b-12d3-a456-426614174000', '923e4567-e89b-12d3-a456-426614174000', 'Where were you on the night of the theft?', 'I was attending a charity gala across town. Hundreds of people can vouch for my presence there.'),
('h23e4567-e89b-12d3-a456-426614174000', '923e4567-e89b-12d3-a456-426614174000', 'Have you ever tried to purchase the Golden Fish Statue?', 'Yes, I offered to buy it last year, but the museum refused to sell. It''s a magnificent piece.'),
('i23e4567-e89b-12d3-a456-426614174000', '923e4567-e89b-12d3-a456-426614174000', 'Do you know anyone who might have wanted to steal the statue?', 'In my experience, museum insiders are often involved in these kinds of thefts. Have you looked into the staff?');
