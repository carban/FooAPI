package models

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	LastName  string  `json:"lastname"`
	Username  string  `json:"username"`
	BirthDate string  `json:"birthdate"`
	Age       int     `json:"age"`
	Gender    string  `json:"gender"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	Country   string  `json:"country"`
	Height    int     `json:"height"`
	Weight    float32 `json:"weight"`
}

type Small_User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Username string `json:"username"`
}

type Post struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Visibility string     `json:"visibility"`
	Tags       []string   `json:"tags"`
	User       Small_User `json:"user"`
}
