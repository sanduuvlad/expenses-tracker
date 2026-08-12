CREATE TABLE expenses (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,

    amount NUMERIC(12,2) NOT NULL,
    CHECK (amount > 0),

    currency VARCHAR(3) NOT NULL,
    CHECK (currency IN ('MDL', 'EUR', 'USD')),

    expense_date DATE NOT NULL, 
    expense_description VARCHAR(200),

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,

    FOREIGN KEY (user_id)
        REFERENCES users(id),
    
    FOREIGN KEY (user_id, category_id)
        REFERENCES categories(user_id, id)
);