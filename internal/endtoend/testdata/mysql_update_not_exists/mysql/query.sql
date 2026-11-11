-- name: UpdateEntities :execrows
UPDATE target_table
SET active = 0
WHERE NOT (
    EXISTS (
        SELECT 1
        FROM source_table s
        WHERE s.updated_at >= sqlc.arg(updated_after)
    )
);
