package productrepo

// ProductRepository 存储库接口
type ProductRepository interface {
	// StoreProduct 存储产品信息
	StoreProduct(name string, id int)
	// FindProductByID 通过产品 ID 查找产品信息
	FindProductByID(id int)
}

func New(environment string) ProductRepository {
	switch environment {
	case "aliCloud":
		return aliCloudProductRepo{}
	case "local-mysql":
		return mysqlProductRepo{}
	default:
		return mockProductRepo{}
	}
}
