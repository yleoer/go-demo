package productrepo

import "fmt"

type aliCloudProductRepo struct{}

func (a aliCloudProductRepo) StoreProduct(name string, id int) {
	fmt.Println("aliCloudProductRepo: mocking the StoreProduct func")
}

func (a aliCloudProductRepo) FindProductByID(id int) {
	fmt.Println("aliCloudProductRepo: mocking the FindProductByID func")
}



