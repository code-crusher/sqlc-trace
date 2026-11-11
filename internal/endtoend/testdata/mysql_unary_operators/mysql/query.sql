-- name: GetNegative :many
SELECT id FROM widgets WHERE amount > -sqlc.arg(floor);

-- name: GetNot :many
SELECT id FROM widgets WHERE NOT active = sqlc.arg(active);

-- name: GetBitNegated :many
SELECT id FROM widgets WHERE amount > ~sqlc.arg(mask);

-- name: GetNegativeSum :many
SELECT id FROM widgets WHERE amount > -(amount + 1);

-- name: GetNotEither :many
SELECT id FROM widgets WHERE NOT (active = 1 OR amount > 2);
