CREATE TABLE facts (
    id serial,
    title varchar NOT NULL,
    description varchar NOT NULL,
    created_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
    PRIMARY KEY (id)
);

