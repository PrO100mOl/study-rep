package main

import (
	"context"
	"fmt"

	simpconnect "tester/feature_postgres/simp_connect"
)

func main() {
	ctx := context.Background()
	con, err := simpconnect.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Print("норм")

	// err = simpsql.CreateTable(con, ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("все сохдало")
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
	// if err := simpsql.UpdateRow(ctx, con); err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// if err := simpsql.DeleteRow(ctx, con); err != nil {
	// 	panic(err)
	// }

	// res, err := simpsql.SelectRows(ctx, con)
	// if err != nil {
	// 	panic(err)
	// }
	// pp.Println(res)

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

	// SQLQ := `INSERT INTO users (full_name, phone_number)
	// VALUES ($1, $2);
	// `
	// _, err = con.Exec(ctx, SQLQ, "игорь", "+424242424")
	// if err != nil {
	// 	panic(err)
	// }

	SQLQ := `SELECT * FROM users
	`
	res, err := con.Query(ctx, SQLQ)
	if err != nil {
		panic(err)
	}
	var we, tre string
	err = res.Scan(&we, &tre)
	fmt.Println(res.Next())
	fmt.Println(we, tre)
	fmt.Println("good boy")
}
