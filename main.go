package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"math"
)
type Stock struct {
	Ticker string
	Gap float64
	Openingprice float64
}
func load(path string)([]Stock,error){
	f,err :=os.Open(path)

	if err !=nil {
		fmt.Println(err)
		return nil,err
	}
	defer f.Close()
	r:=csv.NewReader(f)
	rows,err :=r.ReadAll()

	if err !=nil {
		fmt.Println(err)
		return nil,err
	}
	rows = slices.Delete(rows, 0,1)

	var stocks[]Stock
	for _,row := range rows{
		ticker := row[0]
		// gap := row[1]
		gap,err := strconv.ParseFloat(row[1],64)
		if err !=nil {
			continue
		}
		// Openingprice := row[2]
		openingprice,err := strconv.ParseFloat(row[2],64)
		
		stocks = append(stocks,Stock{
			Ticker: ticker,
			Gap: gap,
			Openingprice: openingprice,
		})

	}
	return stocks,nil
}
func main() {
	stocks,err := load("./opg.csv")
	if err !=nil {
		fmt.Println(err)
		return
	}

	slices.DeleteFunc(stocks, func(s Stock) bool {
		return math.Abs(s.Gap)	< .1
	})

}