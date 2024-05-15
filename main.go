package main

import (
	"fmt"
	"sync"
	"time"
)

type CakeBowl struct {
	ingredients map[string]bool
	mu          sync.RWMutex // you might want to use a different mutex
}

func (cb *CakeBowl) AddIngredient(ingredient string) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.ingredients[ingredient] = true
}

func (cb *CakeBowl) ingredientsAdded() bool {
	// modify the code here
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	time.Sleep(1 * time.Second)
	for _, added := range cb.ingredients {
		if !added {
			return false
		}
	}
	return true
}

// Do not change the code below!
func ChefAddsIngredient(bowl *CakeBowl, wg *sync.WaitGroup, ingredient string) {
	defer wg.Done()
	bowl.AddIngredient(ingredient)
}

func checkIngredients(bowl *CakeBowl, wg *sync.WaitGroup, num int) {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bowl.ingredientsAdded()
		}()
	}
}

func main() {
	bowl := &CakeBowl{
		ingredients: map[string]bool{
			"flour":  false,
			"sugar":  false,
			"eggs":   false,
			"butter": false,
		},
	}

	var wg sync.WaitGroup

	ingredients := []string{"flour", "sugar", "eggs", "butter"}
	for _, ingredient := range ingredients {
		wg.Add(1)
		go ChefAddsIngredient(bowl, &wg, ingredient)
	}

	checkIngredients(bowl, &wg, 10)

	wg.Wait()
	fmt.Println("All ingredients are added! Let's make the dough!")
}
