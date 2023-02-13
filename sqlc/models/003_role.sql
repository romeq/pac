CREATE TABLE IF NOT EXISTS role(
    role_uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS account_to_role(
    account_uuid UUID NOT NULL,
    role_uuid UUID NOT NULL,
    FOREIGN KEY (account_uuid) REFERENCES account,
    FOREIGN KEY (role_uuid) REFERENCES role,
    CONSTRAINT no_duplicates UNIQUE (account_uuid, role_uuid)
);

CREATE TABLE IF NOT EXISTS resource_to_role(
    resource_uuid UUID NOT NULL,
    role_uuid UUID NOT NULL,
    FOREIGN KEY (resource_uuid) REFERENCES resource,
    FOREIGN KEY (role_uuid) REFERENCES role
);
