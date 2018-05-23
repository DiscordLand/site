package database

// DiscordUser represents a generic Discord user
type DiscordUser struct {
	ID            string `json:"id" sql:"unique"`
	Avatar        string `json:"avatar"`
	Email         string `json:"email,omitempty"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

// User represents one of our users
type User struct {
	ID       string `json:"id" sql:"unique"`
	Token    string `json:"token"`
	GitHub   string `json:"github"`
	Homepage string `json:"homepage"`
}

// WebSession represents a users session whenever they log in
type WebSession struct {
	ID    string
	Token string
}
