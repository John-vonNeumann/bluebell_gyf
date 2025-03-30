package logic

import (
	"bluebell_gyf/dao/mysql"
	"bluebell_gyf/models"
)

func GetAllCommunityList() ([]*models.Community, error) {
	return mysql.GetAllCommunityList()
}

func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
