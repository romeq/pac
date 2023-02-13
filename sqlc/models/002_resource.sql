CREATE TABLE IF NOT EXISTS resource(
    resource_uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL
);
