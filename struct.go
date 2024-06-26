type TestStruct struct {
    On bool
    Ammo int
    Power int
}

func (t *TestStruct) Shoot() bool {
    if t.On == false || t.Ammo == 0 {
        return false
    }
    t.Ammo--
    return true
}

func (t *TestStruct) RideBike() bool {
    if t.On == false || t.Power == 0 {
        return false
    }
    t.Power--
    return true
}

func main() {

    testStruct := new(TestStruct)
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */

}