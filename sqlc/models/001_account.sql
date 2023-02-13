CREATE TABLE IF NOT EXISTS account(
    account_uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE
);
