CREATE TABLE budgets (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,

    currency VARCHAR(3) NOT NULL,
    CHECK (currency IN ('MDL', 'EUR', 'USD')),

    budget_limit NUMERIC(12, 2) NOT NULL,
    CHECK (budget_limit > 0),
    
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    CHECK (period_start <= period_end),

    FOREIGN KEY (user_id)
        REFERENCES users(id), 

    FOREIGN KEY (user_id, category_id)
        REFERENCES categories(user_id, id),

    UNIQUE (user_id, category_id, period_start, period_end)
);