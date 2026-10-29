-- Add a username column so accounts can log in without an email address.
-- Existing users get a username derived from the local part of their email,
-- sanitized to [a-z0-9_-] and deduplicated with a numeric suffix.
ALTER TABLE users ADD COLUMN username TEXT;

UPDATE users u
SET username = sub.uname
FROM (
    SELECT id,
           base || CASE WHEN rn > 1 THEN '-' || rn::text ELSE '' END AS uname
    FROM (
        SELECT id,
               COALESCE(NULLIF(regexp_replace(lower(split_part(email, '@', 1)), '[^a-z0-9_-]', '', 'g'), ''), 'user') AS base,
               row_number() OVER (
                   PARTITION BY lower(split_part(email, '@', 1))
                   ORDER BY created_at
               ) AS rn
        FROM users
    ) t
) sub
WHERE u.id = sub.id;

CREATE UNIQUE INDEX users_username_key ON users (username);
