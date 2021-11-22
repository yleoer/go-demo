package productrepo

import "fmt"

type mysqlProductRepo struct{}

func (m mysqlProductRepo) StoreProduct(name string, id int) {
	fmt.Println("mysqlProductRepo: mocking the StoreProduct func")
}

func (m mysqlProductRepo) FindProductByID(id int) {
	fmt.Println("mysqlProductRepo: mocking the FindProductByID func")
}
