package main

import (
	"fmt"
	"os"
)


type bill struct {
	name string
	items map[string]float32
	tip float32
}

func (b bill) addItem(item string, price float32) {
	b.items[item] = price
}

func (b *bill) addTip(tip float32) {
	b.tip = tip
}

func (b bill) getFormated() string {
	text :=  fmt.Sprintf("Čekis kientui \"%v\"\n\n", b.name )

	var totalPrice float32 = 0;

	for item, price := range b.items {
		text += fmt.Sprintf("%-25v ...%0.2f€\n", item + ":", price)
		totalPrice += price
	}
	
	text +="\n\n";
	text += fmt.Sprintf("%-25v ...%0.2f€\n", "Arbatpinigiai:", b.tip)
	text += fmt.Sprintf("%-25v ...%0.2f€\n", "Viso:", totalPrice + b.tip)

	return text;
}

func (b bill) save() {
	formatedBill := b.getFormated()
	data := []byte(formatedBill)

	path := "bills/" + b.name + ".txt";
	
	err := os.WriteFile(path, data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatedBill)
	fmt.Println("čekis išsaugotas: " + path)
}