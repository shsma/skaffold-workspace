CREATE TABLE projects
(
    id           UUID PRIMARY KEY,
    display_name TEXT                     NOT NULL,
    description  TEXT,
    created      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CLOCK_TIMESTAMP(),
    updated      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CLOCK_TIMESTAMP()
);