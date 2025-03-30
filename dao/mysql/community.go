package mysql

import (
	"bluebell_gyf/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetAllCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {
	detail := new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	if err := db.Get(detail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = ErrorInvalidID
		}
		return nil, err
	}
	return detail, nil
}
