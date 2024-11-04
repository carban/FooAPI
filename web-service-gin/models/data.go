package models

type Song struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Artists          string `json:"artists"`
	IsExplicit       bool   `json:"isExplicit"`
	DurationMs       int    `json:"durationMs"`
	AlbumName        string `json:"albumName"`
	AlbumReleaseDate string `json:"albumReleaseDate"`
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
	Reactions  int        `json:"reactions"`
	User       Small_User `json:"user"`
}

type Comment struct {
	ID        string     `json:"id"`
	Comment   string     `json:"comment"`
	Reactions int        `json:"reactions"`
	PostId    string     `json:"postId"`
	User      Small_User `json:"user"`
}

type Product struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Rating      float32 `json:"rating"`
	Stock       int     `json:"stock"`
}

type Todo struct {
	ID     string `json:"id"`
	Todo   string `json:"todo"`
	State  string `json:"state"`
	Closed bool   `json:"closed"`
	UserId string `json:"userId"`
}

type Movie struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Year       string `json:"year"`
	Rated      string `json:"rated"`
	Released   string `json:"released"`
	Runtime    string `json:"runtime"`
	Genre      string `json:"genre"`
	Director   string `json:"director"`
	Writer     string `json:"writer"`
	Actors     string `json:"actors"`
	Plot       string `json:"plot"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	Awards     string `json:"awards"`
	Poster     string `json:"poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbId"`
	BoxOffice  string `json:"boxOffice"`
}
