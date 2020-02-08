package pkg_user

type User struct {
	Id       int64  `json:"id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type Comment struct {
	Author *User
	Text   string
}
