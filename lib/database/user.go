package database

import (
	"golang-crud/config"
	"golang-crud/models"
)

func GetUsersAll() (*[]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return &[]models.User{}, err
	}

	return &users, nil
}

func GetUsersByID(id int) (*models.User, error) {
	var users models.User

	if err := config.DB.First(&users, id).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}

func CreateUsers(users models.User) (*models.User, error) {
	if err := config.DB.Save(&users).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}

func UpdateUsers(id int, userData models.User) (*models.User, error) {
	var users models.User

	if err := config.DB.First(&users, id).Updates(&userData).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}

func DeleteUsers(id int, users models.User) (*models.User, error) {
	if err := config.DB.First(&users, id).Delete(&users).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}

func DeleteUsersPermanent(id int, users models.User) (*models.User, error) {
	if err := config.DB.First(&users, id).Unscoped().Delete(&users).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}
