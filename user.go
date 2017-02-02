package sandpeople

type key int

const userIdKey key = 9876

type User struct {
	ID          string   // unique
	Name        string   // the User's (display) name, e.g. "Andrew Chilton"
	Permissions []string // e.g. []string{"admin", "edit"}
	Pronoun     string   // named in the 'singular', not plural (unlike the header)
	Handle      string   // like a Unix username, e.g. "chilts" (not unique)
	Avatar      string   // a URL
}
