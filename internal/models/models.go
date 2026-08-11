package models

import "time"

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID            string     `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	URL           string     `json:"url"`
	UserID        string     `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

type FeedFollow struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	FeedID    string    `json:"feed_id"`
}

type Post struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	URL         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      string     `json:"feed_id"`
}
