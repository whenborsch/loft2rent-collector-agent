package model

type Loft struct {
	ID       string `db:"loft_id"`
	Title    string `db:"title"`
	Address  string `db:"address"`
	Website  string `db:"website"`
	Phone    string `db:"phone"`
	WhatsApp string `db:"whatsapp"`
	Telegram string `db:"telegram"`
	Email    string `db:"email"`
}
