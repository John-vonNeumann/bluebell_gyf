package logic

import (
	"bluebell_gyf/dao/mysql"
	"bluebell_gyf/models"
)

func GetAllCommunityList() ([]*models.Community, error) {
	return mysql.GetAllCommunityList()
}
