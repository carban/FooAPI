package models

// album represents data about a record album.
type AlbumReq struct {
	ID               string `json:"id"`
	Name             string `json:"name" binding:"required"`
	Artists          string `json:"artists" binding:"required"`
	IsExplicit       bool   `json:"isExplicit"`
	DurationMs       int    `json:"durationMs" binding:"required"`
	AlbumName        string `json:"albumName" binding:"required"`
	AlbumReleaseDate string `json:"albumReleaseDate" binding:"required"`
}

type UserReq struct {
	ID        string  `json:"id"`
	Name      string  `json:"name" binding:"required"`
	LastName  string  `json:"lastname" binding:"required"`
	Username  string  `json:"username" binding:"required"`
	BirthDate string  `json:"birthdate" binding:"required"`
	Age       int     `json:"age" binding:"required"`
	Gender    string  `json:"gender" binding:"required"`
	Phone     string  `json:"phone" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Country   string  `json:"country" binding:"required"`
	Height    int     `json:"height" binding:"required"`
	Weight    float32 `json:"weight" binding:"required"`
}

type Small_UserReq struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type PostReq struct {
	ID         string        `json:"id"`
	Title      string        `json:"title" binding:"required"`
	Content    string        `json:"content" binding:"required"`
	Visibility string        `json:"visibility" binding:"required"`
	Tags       []string      `json:"tags" binding:"required"`
	Reactions  int           `json:"reactions"`
	User       Small_UserReq `json:"user" binding:"required"`
}

type CommentReq struct {
	ID        string        `json:"id"`
	Comment   string        `json:"comment" binding:"required"`
	Reactions int           `json:"reactions"`
	PostId    string        `json:"postId" binding:"required"`
	User      Small_UserReq `json:"user" binding:"required"`
}

type ProductReq struct {
	ID          string  `json:"id"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Brand       string  `json:"brand" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Rating      float32 `json:"rating" binding:"required"`
	Stock       int     `json:"stock"`
}

type TodoReq struct {
	ID     string `json:"id"`
	Todo   string `json:"todo" binding:"required"`
	State  string `json:"state" binding:"required"`
	Closed bool   `json:"closed"`
	UserId string `json:"userId" binding:"required"`
}

type MovieReq struct {
	ID         string `json:"id"`
	Title      string `json:"title" binding:"required"`
	Year       string `json:"year" binding:"required"`
	Rated      string `json:"rated" binding:"required"`
	Released   string `json:"released" binding:"required"`
	Runtime    string `json:"runtime" binding:"required"`
	Genre      string `json:"genre" binding:"required"`
	Director   string `json:"director" binding:"required"`
	Writer     string `json:"writer" binding:"required"`
	Actors     string `json:"actors" binding:"required"`
	Plot       string `json:"plot" binding:"required"`
	Language   string `json:"language" binding:"required"`
	Country    string `json:"country" binding:"required"`
	Awards     string `json:"awards" binding:"required"`
	Poster     string `json:"poster" binding:"required"`
	ImdbRating string `json:"imdbRating" binding:"required"`
	ImdbID     string `json:"imdbId" binding:"required"`
	BoxOffice  string `json:"boxOffice" binding:"required"`
}
