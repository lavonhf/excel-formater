package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"github.com/xuri/excelize/v2"
)

func main() {
    f, err := excelize.OpenFile("./go.xlsx")
    if err != nil {
        log.Fatal(err)
    }
for i := 0; i < 2; i++ {    
 cell := fmt.Sprintf("C%d", i+2)
  value, err := f.GetCellValue("sheet1", cell)
    if err != nil {
        log.Fatalf("Failed to get cell value: %v", err)
    }
newValue :=spaceRemover(value)
result,err:= strconv.Atoi(newValue)
if err != nil {
    panic(err)
}

  f.SetCellValue("sheet1", cell, result)
}
f.Save()
}

func spaceRemover(s string) string {
    return strings.ReplaceAll(s, " ", "")
}






