package main

import (
	"fmt"

	"tester/feature1"
	"tester/feature2"
	simpconnect "tester/feature_postgres/simp_connect"
)

func main() {
	fmt.Println("dsd")
	feature1.Feature1()
	feature2.Feature2()
	// kyufsds
	// fddgdf
	fmt.Println("cverf")
	con, err := simpconnect.CheckConnection()
	if err == nil {
		fmt.Print("норм")
	}
	_ = con
}
