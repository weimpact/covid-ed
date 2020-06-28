CREATE TABLE medias (
    id serial,
    title varchar NOT NULL,
    description varchar NOT NULL,
    kind varchar NOT NULL,
    category varchar NOT NULL,
    url varchar NOT NULL,
    created_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
    PRIMARY KEY (title, url)
);
