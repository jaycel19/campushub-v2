package post

import (
	"testing"
	"time"
)

func TestRankPosts_OrderByScore(t *testing.T) {
	now := time.Now()

	posts := []Post{
		{
			Content:   "Low",
			Likes:     1,
			Comments:  1,
			CreatedAt: now,
		},
		{
			Content:   "High",
			Likes:     50,
			Comments:  20,
			CreatedAt: now.Add(-10 * time.Hour),
		},
		{
			Content:   "Mid",
			Likes:     10,
			Comments:  5,
			CreatedAt: now.Add(-2 * time.Hour),
		},
	}

	result := rankPosts(posts)

	if result[0].Content != "High" {
		t.Errorf("Expected 'High' first, got %s", result[0].Content)
	}

	if result[1].Content != "Mid" {
		t.Errorf("Expected 'Mid' second, got %s", result[1].Content)
	}

	if result[2].Content != "Low" {
		t.Errorf("Expected 'Low' last, got %s", result[2].Content)
	}
}
