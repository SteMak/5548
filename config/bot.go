package config

type bot struct {
	GuildID  string   `json:"guild_id,omitempty"`
	Channels []string `json:"channels,omitempty"`
}
