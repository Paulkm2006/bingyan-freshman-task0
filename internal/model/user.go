package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;index"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Nickname string `json:"nickname,omitempty"`
}

func AddUser(user *User) error {
	// Add user
	resultUser := db.Model(&User{}).Create(user)
	if resultUser.Error != nil {
		return resultUser.Error
	}
	return nil

}

func UpdateUser(user *User) error {
	// Update user
	var old User
	result := db.Where("id = ?", user.ID).First(&old)
	if result.Error != nil {
		return result.Error
	}
	user.ID = old.ID
	result = db.Save(user)
	return result.Error
}

func DeleteUser(id int) error {
	// Delete user
	result := db.Delete(&User{}, id)
	return result.Error
}

func GetUser(username string) (*User, error) {
	// Get user
	user := &User{}
	result := db.Where("username = ?", username).First(user)
	return user, result.Error
}
