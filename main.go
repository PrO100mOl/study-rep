package main

import (
	"fmt"
	"os"
)

func main() {
	// ctx := context.Background()
	// con, err := simpconnect.CheckConnection(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// err = simpsql.InsertRow(con,
	// 	ctx, simpsql.Task{
	// 		Title:       "жить",
	// 		Description: "нужно",
	// 		Completed:   false,
	// 		Created_at:  time.Now(),
	// 	},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = simpsql.UpdateTask(ctx, con, simpsql.Task{
	// 	ID:           1,
	// 	Title:        "walk",
	// 	Description:  "go",
	// 	Completed:    false,
	// 	Created_at:   time.Now(),
	// 	Completed_at: nil,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	r, err := os.Create("out/newfile.txt")
	if err != nil {
		panic(err)
	}
	r.Write([]byte("bkjernkjdgsn. grbsdjkfhbrsgb"))
	fmt.Println("good")
}
