package main

import ( "fmt"
"bufio"
"os"
// "io/ioutil"
// "strconv"

)

type Category struct{
	ID int
    name string
	keywords []string
	ignorewords []string 

} 

type Recipe struct{

} 

func main(){

	fmt.Println("Hello World!")
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)
	
	cate1 := Category{ ID : 1, name : "Beef", keywords: []string{" lamb ", " lamb\n", "steak", "beef"}, ignorewords : []string{"stock"}}
	cate2 := Category{ ID : 2, name : "Chicken", keywords : []string{"chicken"}, ignorewords : []string{"stock"}}
	cate3 := Category{ ID : 3, name : "Vegetables", keywords : []string{"parsnip", "beetroot", "broccoli", "cauliflower", "courgette", "cucumber", "lettuce", "spinach", "runner beans"}, ignorewords : []string{"stock"}}

	categories := []Category{cate1, cate2, cate3}
	fmt.Println(categories[1].keywords)

	// recipes = {}

	for i := 1; i < 20; i++ {
		fmt.Println("recipe" + strconv.Itoa(i) + ".txt")

		// Open the file.
    	f, _ := os.Open("recipe" + strconv.Itoa(i) + ".txt")
    	// Create a new Scanner for the file.
    	scanner := bufio.NewScanner(f)
    	// Loop over all lines in the file and print them.
    	for scanner.Scan() {
      		line := scanner.Text()
      		fmt.Println(line)
    	}


		b, err := ioutil.ReadFile("recipe" + strconv.Itoa(i) + ".txt") // just pass the file name
    	if err != nil {
    		//fmt.Print(err)
    	}
    	new_recipe := string(b) // convert content to a 'string'
    	ing_index := 0
    	for i, e in enumerate(new_recipe):
    		if e == "Ingredients\n":
            	ing_index = i
                break
   }

}
