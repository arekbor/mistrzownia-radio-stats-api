package types

import "time"

type Stats struct {
	Id          int       `json:"id"`
	SteamId     string    `json:"steamId"`
	Username    string    `json:"username"`
	AvatarURL   string    `json:"avatarURL"`
	ProfileURL  string    `json:"profileURL"`
	YoutubeURL  string    `json:"youtubeURL"`
	YoutubeName string    `json:"youtubeName"`
	Datetime    time.Time `json:"datetime"`
}

type ResponseStats struct {
	Count int     `json:"count"`
	Data  []Stats `json:"data"`
}
