package users

import(
	"example.com/amazingmovies/src/pkg/models"
	"example.com/amazingmovies/src/pkg/models/movies"
)

type User struct{
	models.Model
	Username  string   `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	Firstname string   `gorm:"column:firstname;not null;" json:"firstname" form:"firstname"`
	Lastname  string   `gorm:"column:lastname;not null;" json:"lastname" form:"lastname"`
	// Hash for password
	Hash      string   `gorm:"column:hash;not null;" json:"hash"`
	Movies 		[]movies.Movie
}

// func (m *User) BeforeCreate() error {
// 	m.CreatedAt = time.Now()
// 	m.UpdatedAt = time.Now()
// 	return nil
// }

// func (m *User) BeforeUpdate() error {
// 	m.UpdatedAt = time.Now()
// 	return nil
// }