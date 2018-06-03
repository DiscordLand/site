package database

// Bot represents a bot on the site
type Bot struct {
	ID         string   `json:"id" sql:"unique:id"`
	OwnerID    string   `json:"owner_id"`
	Developers []string `json:"developers"`
	Repository string   `json:"repository"`
	Status     int      `json:"status"`
}

// Uptime stores the time a bot last went offline and when they came online
type Uptime struct {
	ID       string `json:"id" sql:"unique:id"`
	Start    int32  `json:"start"`
	Duration int32  `json:"duration"`
}
