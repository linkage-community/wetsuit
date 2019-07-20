package entity

type AlbumFileVariant struct {
	ID        int
	Score     int
	Type      string
	Size      int
	URL       string
	Extension string
	MIME      string
}
type AlbumFile struct {
	ID       int
	Name     string
	Variants []AlbumFileVariant
}
type Application struct {
	ID          int
	Name        string
	IsAutomated bool
}
type User struct {
	ID         int
	Name       string
	ScreenName string
	PostsCount int
	CreatedAt  string
	UpdatedAt  string
	AvatarFile *AlbumFile
}
type Post struct {
	ID          int
	CreatedAt   string
	UpdatedAt   string
	Text        string
	User        User
	Application Application
	Files       []*AlbumFile
}
