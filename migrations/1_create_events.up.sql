CREATE TABLE IF NOT EXISTS events(
    uuid          text,
    title         text      NOT NULL,
    start_at      timestamp NOT NULL,
    duration      bigint    NOT NULL,
    descr         text      NOT NULL,
    user_name     text      NOT NULL,
    notify_at     timestamp,
    notified      boolean default false,
    CONSTRAINT events_pkey PRIMARY KEY (uuid)
);

CREATE INDEX ON events (start_at);
CREATE index ON events USING HASH(user_name);
