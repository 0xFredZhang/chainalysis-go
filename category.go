package chainalysis

// https://support.chainalysis.com/hc/en-us/articles/16492397142292-Category-IDs

const (
	urlRetrieveCategories = "/api/kyt/v2/categories"
)

func (c *ClientImpl) RetrieveCategories() (resp RetrieveCategoriesResp, err error) {
	_, err = c.client.R().
		SetResult(&resp).
		SetError(&resp).
		Get(urlRetrieveCategories)
	return
}

type RetrieveCategoriesResp struct {
	ErrorResp
	Categories []Category `json:"categories"`
}

type Category struct {
	CategoryId   int64  `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
