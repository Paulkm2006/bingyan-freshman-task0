package model

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func AddAdmin(adminUser *Admin) error {
	result := db.Model(&Admin{}).Create(adminUser)
	return result.Error
}

func UpdateAdmin(adminUser *Admin) error {
	result := db.First(&Admin{}, adminUser.ID)
	if result.Error != nil {
		return result.Error
	}
	result = db.Save(adminUser)
	return result.Error
}

func DeleteAdmin(id int) error {
	result := db.Delete(&Admin{}, id)
	return result.Error
}

func GetAdmin(username string) (*Admin, error) {
	adminUser := &Admin{}
	result := db.Where("username = ?, role = ?", username, "admin").First(adminUser)
	return adminUser, result.Error
}
