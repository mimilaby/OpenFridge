-- name: AddFood :one
INSERT INTO food (
  name,
  description,
  price,
  calories,
  weight,
  amount,
  weight_per_amount,
  photo,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetFoodAvailable :many
SELECT * FROM food
WHERE amount > 0;
