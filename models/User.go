package models

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Interests []Interest
}
type Interest struct {
	ID   int
	Name string
}
type Page struct {
	ID          int
	Name        string
	Description string
	URI         string
}

type UserViewModel struct {
	Page  *Page // Page yapısına pointer olarak referans tanımlanıyor
	Users []User
}

type InterestMapping struct {
	UserID     int
	InterestID int
}
