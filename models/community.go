package models

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64  `json:"id" db:"community_id"`
	Name         string `json:"name" db:"community_name"`
	Introduction string `json:"introduction,omitempty" db:"introduction"`
	CreteTime    string `json:"create_time" db:"create_time"`
}
