package main

import (
	"fmt"

	//"bankingapp/Account"
	admin "bankingapp/Admin"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	admin1 := admin.NewAdmin("SANJEEV")
	fmt.Println(admin1)
	user1, _ := admin1.CreateUser("Sanjeev", "Yadav")
	fmt.Println(user1)
	fmt.Println(admin1)

	//admin1.DeleteCreatedUser("Sanjeev")

	//fmt.Println(admin1)

	bank1, _ := admin1.CreateBank("HDFC")
	fmt.Println(bank1)

	// _, err := user1.CreateNewAccount("HDFC", 6883738, "SADSFWF12")
	// if err != nil {
	// 	panic(err)
	// }

	account1, _ := user1.CreateNewAccount("HDFC", 886876, "sdfsf12")
	fmt.Println(account1)

	err := admin1.UpdateCreatedUser("Sanjeev", "firstName", "SAHIL")
	if err != nil {
		panic(err)
	}
	fmt.Println(user1)

	//admin1.UpdateCreatedBank("HDFC", "bankname", "SBI")

	//admin1.DeleteCreatedBank("HDFC")
	//fmt.Println(admin1.BanksCreatedByMe)
	//fmt.Println(admin1.BanksCreatedByMe)
	// CHECK BANK UPDATE
	//admin1.UpdateCreatedBank("SBI", "bankname", "HDFC")
	fmt.Println(admin1.BanksCreatedByMe[0])

	transfer := user1.TransferMoney("HDFC", 1)
	fmt.Println(transfer)

}
