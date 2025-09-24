package model

type Paginate struct {
	PageNo   int `json:"pageNo,omitempty" form:"pageNo"`
	PageSize int `json:"pageSize,omitempty" form:"pageSize"`
}

func (page *Paginate) GetPageNo() int {
	if page.PageNo <= 0 {
		page.PageNo = 1
	}

	return page.PageNo
}

func (page *Paginate) GetPageSize() int {
	if page.PageSize <= 0 {
		page.PageSize = 10
	}

	return page.PageSize
}
