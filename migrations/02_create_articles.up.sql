CREATE TABLE articles (
    id serial,
    title varchar NOT NULL,
    url varchar NOT NULL,
    fact_id integer REFERENCES facts (id),
    created_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
    PRIMARY KEY (id)
);

