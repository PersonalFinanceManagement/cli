-- migrate:up

CREATE TABLE transactions (
    id TEXT PRIMARY KEY,
    amount INTEGER NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT FALSE,
    type TEXT NOT NULL,
    -- core forieng key fields 
    source_account_id TEXT NOT NULL,
    destination_account_id TEXT NOT NULL,
    -- metadata fields 
    payee TEXT,
    category_id TEXT,
    description TEXT,
    method_of_payment TEXT,

    created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)

-- migrate:down
DROP TABLE transactions;
