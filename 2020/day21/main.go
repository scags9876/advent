package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/scags9876/adventOfCode/lib"
	"sort"
	"strings"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const testInput2Filename = "testinput2.txt"

const verbose = false

func main() {
	fmt.Println("start")
	input := lib.GetInputStrings(inputFilename)
	ingredientMap := part1(input)
	part2(ingredientMap)
}

func part1(input []string) map[string]string {
	foods := parseInput(input)

	if verbose {
		fmt.Printf("input: %s", spew.Sdump(foods))
	}

	allergenSet := getAllergenSet(foods)
	if verbose {
		fmt.Printf("allergens: %s", spew.Sdump(allergenSet))
	}

	// map ingredients to allergens
	ingredientMap := make(map[string]string)
	// map allergens to ingredients
	allergenMap := make(map[string]string)

	possibleIngredients := make(map[string][]string)
	for {
		for _, food := range foods {
			for _, allergen := range food.allergens {
				if ing, ok := allergenMap[allergen]; ok {
					if verbose {
						fmt.Printf("allergen %s already mapped to %s\n", allergen, ing)
					}
					continue
				}
				if set, ok := possibleIngredients[allergen]; ok {
					ingredientList := make([]string, 0)
					for _, ingredient := range food.ingredients {
						if a, ok := ingredientMap[ingredient]; ok {
							if verbose {
								fmt.Printf("ingredient %s already mapped to %s\n", ingredient, a)
							}
							continue
						}
						ingredientList = append(ingredientList, ingredient)
					}
					newSet := intersectSet(set, ingredientList)
					possibleIngredients[allergen] = newSet
					if len(newSet) == 1 {
						allergenMap[allergen] = newSet[0]
						ingredientMap[newSet[0]] = allergen
					}

				} else {
					possibleIngredients[allergen] = food.ingredients
				}
			}
		}
		if len(allergenMap) == len(allergenSet) {
			break
		}
	}

	if verbose {
		fmt.Printf("ingredientMap: %s\nallergenMap: %s", spew.Sdump(ingredientMap), spew.Sdump(allergenMap))
	}

	count := 0
	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if _, ok := ingredientMap[ingredient]; !ok {
				// count if this food does not map to any allergen
				count++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", count)
	return ingredientMap
}

func intersectSet(a, b []string) []string {
	newSet := make([]string, 0)
	for _, valA := range a {
		if stringInSlice(b, valA) {
			newSet = append(newSet, valA)
		}
	}
	return newSet
}

func stringInSlice(set []string, s string) bool {
	for _, el := range set {
		if el == s {
			return true
		}
	}
	return false
}

type food struct {
	ingredients []string
	allergens []string
}

func parseInput(input []string) []food {
	var foods []food
	for _, line := range input {
		food := food{}
		fields := strings.Fields(line)
		mode := "ingredient"
		for _, field := range fields {
			if field == "(contains" {
				mode = "allergen"
				continue
			}
			if mode == "ingredient" {
				food.ingredients = append(food.ingredients, field)
			} else {
				allergen := strings.Trim(field, ",)")
				food.allergens = append(food.allergens, allergen)
			}
		}
		foods = append(foods, food)
	}
	return foods
}

func getAllergenSet(foods []food) []string {
	allergens := make([]string, 0)
	for _, food := range foods {
		for _, allergen := range food.allergens {
			allergens = setInsert(allergens, allergen)
		}
	}

	return allergens
}

func setInsert(ss []string, s string) []string {
	i := sort.SearchStrings(ss, s)
	if i > len(ss)-1 || ss[i] != s {
		ss = append(ss, "")
		copy(ss[i+1:], ss[i:])
		ss[i] = s
	}
	return ss
}

func part2(ingredientMap map[string]string) {
	list := make([]string, 0)

	for ingredient, allergen := range ingredientMap {
		i := sort.Search(len(list), func(idx int) bool {
			return ingredientMap[list[idx]] >= allergen
		})
		if i > len(list)-1 || list[i] != ingredient {
			list = append(list, "")
			copy(list[i+1:], list[i:])
			list[i] = ingredient
		}
	}

	fmt.Println(strings.Join(list, ","))
}
