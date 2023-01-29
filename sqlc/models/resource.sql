CREATE TABLE IF NOT EXISTS resource(
    resource_uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS resource_to_role(
    resource_uuid UUID NOT NULL,
    role_uuid UUID NOT NULL,
    FOREIGN KEY (resource_uuid) REFERENCES resource,
    FOREIGN KEY (role_uuid) REFERENCES role
);