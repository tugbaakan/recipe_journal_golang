package main

import ( 
	"bufio"
    "fmt"
    "log"
	"os"
	"strconv"
	"strings"
)

// To import the txt files
// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

//To search in a slice
func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}

//Define a new strict for the categories
type Category struct{
	ID int
    name string
	keywords []string
	ignorewords []string 
}

//Define a new struct for the recipes
type Recipe struct{
	name string
	ingredients []string
	category []string
} 

//Define a function for the Recipe struct
func (r *Recipe) setCategory(categories []Category) {
	// loop for all categories
	for _, cate := range categories {
		//loop for all ingredients
		for _, ingr := range r.ingredients {
			//loop for all the words to be ignored
			for _, kw0 := range cate.ignorewords {
				if contains(strings.Fields(ingr), kw0) == false {
					//loop for all keywords
					for _, kw := range cate.keywords{
						if contains(strings.Fields(ingr), kw) {
							r.category = append(r.category, cate.name)
                        	break
                        }
                    }
				}
            }
		}
	}
}

func main(){

	//First of all, greetings
	fmt.Println("Hello World!")

	//Describe the categories
    cate1 := Category{ ID : 1, name : "Beef", keywords: []string{"lamb", " lamb ", " lamb\n", "steak", "beef"}, ignorewords : []string{"stock"}}
    cate2 := Category{ ID : 2, name : "Chicken", keywords : []string{"chicken"}, ignorewords : []string{"stock"}}
    cate3 := Category{ ID : 3, name : "Vegetables", keywords : []string{"parsnip", "parsnips", "beetroot", "broccoli", "cauliflower", "courgette", "courgettes", "cucumber", "lettuce", "spinach", "runner beans"}, ignorewords : []string{"stock"}}

	categories := []Category{cate1, cate2, cate3}

	//import the recipes
	var recipes []Recipe
	for i := 1; i < 10; i++ {
		new_recipe, err := readLines("recipe" + strconv.Itoa(i) + ".txt")
    	if err != nil {
        	log.Fatalf("readLines: %s", err)
    	}
    	 ing_index := 0
        // take the below line of 'Ingredients' for ingredient list
        // find the index and assign it to ing_index
        for i, e := range new_recipe {
        	//fmt.Println(i, e)
    		if e == "Ingredients" {
				ing_index = i
    			//fmt.Println(i, e, ing_index)
    			break
    		}
    	}
    	rec := Recipe{ name : new_recipe[0], ingredients : new_recipe[ing_index:] }
    	rec.setCategory(categories)
        recipes = append ( recipes, rec )
    }

    // We are going to list the recipes of the categories that the user would like to see
    //The user will input a number regarding the categories
    fmt.Println("Type the number of the category you wish to be listed")
    for _, item := range categories {
    	fmt.Println( strconv.Itoa(item.ID) + " for " + item.name )
    }
    var inp_cate_ID int
    fmt.Scan(&inp_cate_ID)
	//fmt.Println("read number", inp_cate_ID, "from stdin")

	//loops through the recipes and categories to find the related recipes
    for _, item := range categories{
        if inp_cate_ID == item.ID {
            fmt.Println(item.name)
            for _, item2 := range recipes {
                for _, item3 := range item2.category {
                    if item3 == item.name {
                    	//found the recipes
                    	//print them
                        fmt.Println( item2.name )
                        break
                    }
                }
            }
        }
    }    


}
