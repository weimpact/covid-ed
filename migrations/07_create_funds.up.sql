CREATE TABLE funds (
    id serial,
    title varchar NOT NULL,
    description varchar NOT NULL,
    website varchar NOT NULL,
    donate_url varchar not null,
    image_url varchar,
    created_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
    PRIMARY KEY (website)
);

