groupCity := map[int][]string{
	10:   []string{...}, // города с населением 10-99 тыс. человек
	100:  []string{...}, // города с населением 100-999 тыс. человек
	1000: []string{...}, // города с населением 1000 тыс. человек и более
}


cityPopulation100 := make(map[string]int, len(groupCity[100]))
	for pop, city := range groupCity[100] {
		cityPopulation100[city] = pop
	}

	for city, _ := range cityPopulation {
		if _, ok := cityPopulation100[city]; !ok {
			delete(cityPopulation, city)
		}
	}