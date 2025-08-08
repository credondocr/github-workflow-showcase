package models

import (
	"errors"
	"time"
)

// User represents the user model in the application
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required,min=2,max=100"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"required,min=1,max=120"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the persistence operations for users
type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (*User, error)
	Create(user *User) error
	Update(id int, user *User) error
	Delete(id int) error
}

// InMemoryUserRepository implements UserRepository in memory for this example
type InMemoryUserRepository struct {
	users  []User
	nextID int
}

// NewInMemoryUserRepository creates a new instance of the in-memory repository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  []User{},
		nextID: 1,
	}
}

// GetAll returns all users
func (r *InMemoryUserRepository) GetAll() ([]User, error) {
	return r.users, nil
}

// GetByID returns a user by their ID
func (r *InMemoryUserRepository) GetByID(id int) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// Create adds a new user
func (r *InMemoryUserRepository) Create(user *User) error {
	user.ID = r.nextID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.users = append(r.users, *user)
	r.nextID++
	return nil
}

// Update updates an existing user
func (r *InMemoryUserRepository) Update(id int, user *User) error {
	for i, u := range r.users {
		if u.ID == id {
			user.ID = id
			user.CreatedAt = u.CreatedAt
			user.UpdatedAt = time.Now()
			r.users[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

// Delete removes a user by their ID
func (r *InMemoryUserRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

// Validate validates the user fields
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if len(u.Name) < 2 {
		return errors.New("name must be at least 2 characters long")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Age <= 0 {
		return errors.New("age must be greater than 0")
	}
	if u.Age > 120 {
		return errors.New("age must be less than 120")
	}
	return nil
}
