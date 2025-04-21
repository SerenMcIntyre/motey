
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    background VARCHAR(6) NOT NULL
);

CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    title_name VARCHAR(50) NOT NULL,
    background VARCHAR(6) NOT NULL,
    sticker BYTEA,
    is_measured BOOLEAN NOT NULL,
    measurement_unit VARCHAR(25),
    sticker_value INTEGER,
    user_id UUID NOT NULL REFERENCES users(id)
);

CREATE TABLE frequency_groups (
    id UUID PRIMARY KEY,
    name VARCHAR(25) NOT NULL
);

CREATE TABLE frequencies (
    id UUID PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    frequency_group_id UUID REFERENCES frequency_groups(id)
);

CREATE TABLE task_notifications (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    time TIME NOT NULL,
    task_id UUID NOT NULL REFERENCES tasks(id),
    frequency_id UUID REFERENCES frequencies(id)
);
