package main

import (
	"context"
	"fmt"
	"time"

	simpconnect "tester/feature_postgres/simp_connect"
	simpsql "tester/feature_postgres/simp_sql"
)

func main() {
	ctx := context.Background()
	con, err := simpconnect.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("норм")

	// err = simpsql.CreateTable(con, ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("все сохдало")
	// err = simpsql.InsertRow(con,
	// 	ctx,
	// 	"Oбед",
	// 	"покушать",
	// 	false,
	// 	time.Now(),
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

	err = simpsql.UpdateTask(ctx, con, simpsql.Task{
		ID:           1,
		Title:        "dump",
		Description:  "go",
		Completed:    false,
		Created_at:   time.Now(),
		Completed_at: nil,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("good boy")
}
