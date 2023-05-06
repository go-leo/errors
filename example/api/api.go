package api

type GetUserReq struct {
	UID int `form:"uid"`
}

type GetUserResp struct {
	UID  int
	Name string
}
