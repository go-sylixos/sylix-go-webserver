package main

/*
应用
状态
集群名称
部署时间
更新人
*/
type App struct {
	Name       string `json:"name"`
	Status     uint   `json:"status"`
	Cluster    string `json:"cluster"`
	DeployTime string `json:"deployTime"`
	UpdateBy   string `json:"updateBy"`
}

/*
size: 10,
  list: [ ],
  total: 32,
  pageSize: 10,
  pageNum: 1,
  pages: 4,
  hasPreviousPage: false,
  hasNextPage: true
*/

type ListResponse struct {
	Size       int   `json:"size"`
	List       []App `json:"list"`
	Total      int   `json:"total"`
	PageSize   int   `json:"pageSize"`
	PageNum    int   `json:"pageNum"`
	Pages      int   `json:"pages"`
	HasPrePage bool  `json:"hasPreviousPage"`
	HasNext    bool  `json:"hasNextPage"`
}

type ListRequest struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Code     int `json:"code"`
}
