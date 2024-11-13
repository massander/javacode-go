create table if not exists
    wallets (
        wallet_id uuid primary key,
        balance int not null default 0 check (balance >= 0),
        created_at timestamptz not null default now(),
        updated_at timestamptz not null default now()
    );