package post

import (
	"sort"
	"strconv"
	"time"
)

type Service interface {
	GetFeed(page, limit string) ([]Post, error)
	CreatePost(post *Post) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetFeed(pageStr, limitStr string) ([]Post, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	} else if limit > 50 {
		limit = 50
	}

	offset := (page - 1) * limit
	posts, err := s.repo.GetPaginated(offset, limit)
	if err != nil {
		return nil, err
	}

	return rankPosts(posts), nil
}

func (s *service) CreatePost(post *Post) error {
	return s.repo.Create(post)
}

// Feed ranking
func rankPosts(posts []Post) []Post {
	type ScoredPost struct {
		Post  Post
		Score float64
	}

	var scored []ScoredPost

	now := time.Now()

	for _, p := range posts {
		hours := now.Sub(p.CreatedAt).Hours()

		score := float64(p.Likes*2+p.Comments*3) - hours*0.1

		scored = append(scored, ScoredPost{
			Post:  p,
			Score: score,
		})
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	var result []Post
	for _, s := range scored {
		result = append(result, s.Post)
	}

	return result
}
