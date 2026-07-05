package user

import "sync"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserService keeps users in memory so the example runs with no database.
type UserService struct {
	once   sync.Once
	mu     sync.Mutex
	users  []User
	nextID int
}

func (s *UserService) seed() {
	s.once.Do(func() {
		s.users = []User{
			{ID: 1, Name: "Grace Hopper", Email: "grace@example.com"},
			{ID: 2, Name: "Alan Turing", Email: "alan@example.com"},
			{ID: 3, Name: "Ada Lovelace", Email: "ada@example.com"},
		}
		s.nextID = 4
	})
}

func (s *UserService) List() []User {
	s.seed()
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]User, len(s.users))
	copy(users, s.users)
	return users
}

func (s *UserService) Get(id int) (User, bool) {
	s.seed()
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range s.users {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

func (s *UserService) Create(name, email string) User {
	s.seed()
	s.mu.Lock()
	defer s.mu.Unlock()

	u := User{ID: s.nextID, Name: name, Email: email}
	s.nextID++
	s.users = append(s.users, u)
	return u
}
