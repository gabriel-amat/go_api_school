DO $$
BEGIN
    -- Verificar e criar o usuário se não existir
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'admin') THEN
        EXECUTE "CREATE USER admin WITH PASSWORD 'admin'";
    END IF;
END $$;