package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getUserInput("Įveskite užsakovo vardą: ", reader)
	newBill := bill{name: name, items: map[string]float32{}, tip: 0}
	promptActions(reader, &newBill)
}


func promptActions(reader *bufio.Reader, bill *bill) {
	action, _ := getUserInput("p - pridėti prekę, s - išsaugoti čekį, a - arbatpinigiai: ", reader);

	switch action {
	case "p":	
		title, _ := getUserInput("Prekė: ", reader);
		price, _ := getUserInput("Kaina: ", reader);

		priceFloat, error := strconv.ParseFloat(price, 32)

		if error != nil {
			fmt.Println("Neteisinga kaina")
			promptActions(reader, bill)
		} else {
			bill.addItem(title, float32(priceFloat))
		}

		promptActions(reader, bill)
	case "a":	
		tip, _ := getUserInput("Arbatpinigiai: ", reader);
		tipFloat, error := strconv.ParseFloat(tip, 32)

		if error != nil {
			fmt.Println("Neteisinga kaina")
			promptActions(reader, bill)
		} else {
			bill.addTip(float32(tipFloat))
		}

		promptActions(reader, bill)
	case "s":	
		bill.save()
	default:
		fmt.Println("Neteisingas pasirinkimas");
		promptActions(reader, bill)
	}
}