package post

type Service interface {
	GetFeed() ([]Post, error)
	CreatePost(post *Post) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetFeed() ([]Post, error) {
	posts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *service) CreatePost(post *Post) error {
	return s.repo.Create(post)
}
