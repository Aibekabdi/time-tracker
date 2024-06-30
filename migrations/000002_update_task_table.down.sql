ALTER TABLE users DROP CONSTRAINT users_passport_serie_passport_number_key;
ALTER TABLE users ALTER COLUMN passport_serie TYPE VARCHAR(50) USING passport_serie::VARCHAR;
ALTER TABLE users ALTER COLUMN passport_number TYPE VARCHAR(50) USING passport_number::VARCHAR;
ALTER TABLE users ADD CONSTRAINT users_passport_serie_passport_number_key UNIQUE (passport_serie, passport_number);
