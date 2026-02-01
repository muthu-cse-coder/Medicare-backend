package database

// User queries
const (
	QueryCreateUser = `
		INSERT INTO users (name, email, password, phone, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	QueryGetUserByEmail = `
		SELECT id, name, email, password, phone, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	QueryGetUserByID = `
		SELECT id, name, email, password, phone, role, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	QueryCheckEmailExists = `
		SELECT COUNT(*) FROM users WHERE email = $1
	`

	QueryUpdateUser = `
		UPDATE users
		SET name = $1, email = $2, phone = $3, role = $4, updated_at = NOW()
		WHERE id = $5
	`

	QueryDeleteUser = `
		DELETE FROM users WHERE id = $1
	`

	QueryGetAllUsers = `
		SELECT id, name, email, phone, role, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	QueryCountUsers = `
		SELECT COUNT(*) FROM users
	`

	QueryGetUsersByRole = `
		SELECT id, name, email, phone, role, created_at, updated_at
		FROM users
		WHERE role = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
)
