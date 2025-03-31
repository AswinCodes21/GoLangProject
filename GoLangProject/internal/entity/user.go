package entity

type User struct {
	ID       int    `db:"id"`
	FullName string `db:"full_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
