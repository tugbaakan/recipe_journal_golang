package main

import ( 
	"bufio"
    "fmt"
    "log"
	"os"
// "io/ioutil"
"strconv"
"strings"

)

type Category struct{
	ID int
    name string
	keywords []string
	ignorewords []string 

} 

type Recipe struct{
	name string
	ingredients []string
	category []string
} 

func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}


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

func main(){

	fmt.Println("Hello World!")
	
	cate1 := Category{ ID : 1, name : "Beef", keywords: []string{" lamb ", " lamb\n", "steak", "beef"}, ignorewords : []string{"stock"}}
	cate2 := Category{ ID : 2, name : "Chicken", keywords : []string{"chicken"}, ignorewords : []string{"stock"}}
	cate3 := Category{ ID : 3, name : "Vegetables", keywords : []string{"parsnip", "beetroot", "broccoli", "cauliflower", "courgette", "cucumber", "lettuce", "spinach", "runner beans"}, ignorewords : []string{"stock"}}

	categories := []Category{cate1, cate2, cate3}
	fmt.Println(categories[1].keywords)

	var recipes []Recipe

	for i := 1; i < 10; i++ {
		//fmt.Println("recipe" + strconv.Itoa(i) + ".txt")

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
        //_ = recipes
    }

    fmt.Println("Type the number of the category you wish to be listed")
    for _, item := range categories {
    	fmt.Println( strconv.Itoa(item.ID) + " for " + item.name )
    }

    reader := bufio.NewReader(os.Stdin)
    inp_cate_ID_0, _ := reader.ReadString('\n')
    inp_cate_ID  , _ := strconv.Atoi(inp_cate_ID_0)

    for _, item := range categories {
    	if item.ID == inp_cate_ID  {
        	fmt.Println( item.name )
        }
        for _, item2 := range recipes {
             for _, item3 := range item2.category {
             	if item3 == item.name {
             		fmt.Println( item2.name )
                    break
                }
            }
        }
    }
}
