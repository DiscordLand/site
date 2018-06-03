package database

// DiscordUser represents a generic Discord user
type DiscordUser struct {
	Username      string `json:"username"`
	Locale        string `json:"locale"`
	ID            string `json:"id" sql:"unique:id"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	Email         string `json:"email,omitempty"`
}

// User represents one of our users
type User struct {
	ID       string `json:"id" sql:"unique:id"`
	Token    string `json:"token"`
	GitHub   string `json:"github"`
	Homepage string `json:"homepage"`
}

// Session represents a user's session whenever they log in
type Session struct {
	ID    string `json:"id"`
	Token string `json:"unique:token"`
}
