CREATE TABLE IF NOT EXISTS poi
(
    created_on date NOT NULL,
    id serial PRIMARY KEY,
    updated_on date,
    name text
);

CREATE TABLE IF NOT EXISTS events
(
    created_on date NOT NULL,
    date date NOT NULL,
    event_id serial PRIMARY KEY,
    poi_id INTEGER REFERENCES poi(id) ON DELETE RESTRICT,
    updated_on date,
    event_title text
);
