package main

import (
	"fmt"
	"sort"
)

func main() {

	hero1 := Hero{"Vital", 8, 250, nil}
	fmt.Printf("Герой %s (уровень %v), в кармане %v золота\n", hero1.Name, hero1.Level, hero1.Gold)

	hero1.AddItem("меч", 1)
	hero1.AddItem("зелье", 2)
	hero1.PrintInventory()
	hero1.RemoveItem("меч", 1)
	hero1.PrintInventory()
	hero1.AddItem("посох", 1)
	hero1.AddItem("сфера некроманта", 1)
	hero1.AddItem("осколок души", 3)
	hero1.PrintInventory()

	chest := map[string]int{
		"корона лича": 1,
		"лохмотья":    3,
		"ухо гоблина": 5,
		"перо гарпии": 0,
		"":            2,
		"Жезл":        -2,
	}

	a, s := hero1.MergeInventory(chest)

	fmt.Println("Added", a, "Skipped:", s)

	hero1.PrintInventory()

}

type Hero struct {
	Name      string
	Level     int
	Gold      int
	Inventory map[string]int
}

func (h *Hero) AddItem(name string, qty int) {
	if qty <= 0 {
		return
	}

	if h.Inventory == nil {
		h.Inventory = make(map[string]int)
	}

	if v, ok := h.Inventory[name]; ok {
		h.Inventory[name] = v + qty
	} else {
		h.Inventory[name] = qty
	}
}

func (h *Hero) RemoveItem(name string, qty int) bool {

	if qty <= 0 {
		fmt.Println("Некоректное кол-во предметов")
		return false
	}

	if v, ok := h.Inventory[name]; ok {
		if qty > v {
			fmt.Printf("В инвентаре нет столько пердметов: %s, всего их %v\n ", name, h.Inventory[name])
			return false
		}
	}

	if v, ok := h.Inventory[name]; ok {
		h.Inventory[name] = v - qty
		fmt.Printf("Предмет: %s, удалено: %v, осталось: %v\n", name, qty, h.Inventory[name])
		if h.Inventory[name] == 0 {
			delete(h.Inventory, name)
		}
		return true
	}
	return false
}

func (h *Hero) Count(name string) int {

	if v, ok := h.Inventory[name]; ok {
		return v
	}
	return 0
}

func (h *Hero) PrintInventory() {
	if len(h.Inventory) == 0 {
		fmt.Println("Инвентарь пуст")
		return
	}
	fmt.Println("Содержание инвентаря:")
	keys := make([]string, 0, len(h.Inventory))
	for k := range h.Inventory {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i := 0; i < len(keys); i++ {
		fmt.Printf("Предмет %s - %v штука(и)\n", keys[i], h.Inventory[keys[i]])
	}
	fmt.Println("___")
}

func (h *Hero) AddGold(gold int) {

	if gold <= 0 {
		fmt.Println("Недопустимое значение кол-ва золота")
		return
	}
	h.Gold += gold
}

func (h *Hero) SpendGold(gold int) bool {

	if gold <= 0 {
		fmt.Println("Некорректное количество золота")
		return false
	}

	if gold > h.Gold {
		fmt.Printf("У героя %s не хватает золота. Всего у него: %v\n", h.Name, h.Gold)
		return false
	}

	h.Gold -= gold
	fmt.Printf("Герой %s потратил %v золота, у него осталось %v", h.Name, gold, h.Gold)
	return true
}

func (h *Hero) MergeInventory(other map[string]int) (added int, skipped int) {

	if other == nil || len(other) == 0 {
		return 0, 0
	}

	for name, qty := range other {

		if name == "" {

			if qty > 0 {
				skipped += qty
			}
			continue
		}
		if qty <= 0 {
			continue
		}

		h.AddItem(name, qty)
		added += qty

	}
	return added, skipped
}
