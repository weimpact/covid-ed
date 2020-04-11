CREATE TABLE myths (
    id serial,
    title varchar NOT NULL,
    description varchar,
    fact_id integer REFERENCES facts (id),
    created_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
    PRIMARY KEY (id)
);

