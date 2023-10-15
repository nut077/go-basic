package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/nut077/go-basic/models"
)

import "fmt"

func main() {
	// var name type = value
	// 8 bit = 1 byte
	// 32 bit = 4 byte
	// var b uint = 2; uint คือไม่มีค่าติดลบ ทำให้มีขนาดกว้างขึ้นกว่าเดิม

	//var a int = 2  มีค่าเหมือนกับ a := 2;
	a := 2
	b := float64(a) // วิธีเปลี่ยน type
	fmt.Println(b)
	fmt.Println("------------------------")

	// if
	i := 8
	if i%2 == 0 {
		fmt.Println("%2")
	} else {
		fmt.Println("not %2")
	}
	fmt.Println("------------------------")

	// for loop
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
	}

	fmt.Println("------------------------")

	i = 1
	// switch
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("Unknown number")
	}
	fmt.Println("------------------------")

	// array
	var names [3]string
	names[0] = "zero"
	names[1] = "one"
	names[2] = "two"
	names2 := [3]string{"zero", "one", "two"}

	var names3 []string
	names3 = append(names3, "zero")
	names3 = append(names3, "two")
	names3 = append(names3, "two")

	names4 := []string{"zero", "one", "two"}

	fmt.Println(names[0])
	fmt.Println("------------------------")
	fmt.Println(names2)
	fmt.Println("------------------------")
	fmt.Println(names3)
	fmt.Println("------------------------")
	fmt.Println(names4)
	fmt.Println("------------------------")
	for i := 0; i < len(names2); i++ {
		fmt.Println(names2[i])
	}
	fmt.Println("------------------------")
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println("------------------------")

	printFullName("freedom", "seven")

	fmt.Println("------------------------")
	fmt.Println(getFullName("get", "freedom"))

	// เพราะว่าฟังก์ชันคืนค่าสองค่า เราจึงประกาศตัวแปรมารองรับได้พร้อมกันสองตัว
	result, err := divide(5, 2)

	// ตรวจสอบก่อนว่ามี error ไหม ถ้ามีก็จบโปรแกรมไปแบบไม่ค่อยสวยด้วย Exit(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
	fmt.Println("------------------------")

	freedom := human{name: "freedom", age: 33}
	// ใช้ & แทนการอ้างถึง reference
	setAdult(&freedom)
	fmt.Println(freedom)
	freedom.printInfoHuman()

	fmt.Println("------------------------")

	talkers := [2]talker{
		human{name: "freedom", age: 33},
		parrot{name: "Kew", age: 2},
	}

	for _, talker := range talkers {
		talker.talk()
	}

	fmt.Println("------------------------")

	search("dog")

	fmt.Println("------------------------")

	resultSum, name := sum("freedom", 2, 5)
	fmt.Println("result =", resultSum)
	fmt.Println("name = " + name)
	fmt.Println("------------------------")

	b1 := book{models.Book{
		Id:     1,
		Title:  "spring",
		Author: "freedom",
	}}

	Debug(b1.GetBook())
	b1.SetTitle("eiei")
	Debug(b1.GetBook())

	fmt.Println("------------------------")
	b2 := NewBook("go lang", "freedom")
	b2.SetTitle("GO LANG")
	Debug(b2.GetBook())

	fmt.Println("------------------------")
	numInt := []int{1, 2, 3, 4, 5}
	numFloat64 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumInt(numInt))
	fmt.Println(sumFloat64(numFloat64))
	fmt.Println(sumIntOrFloat64(numInt))
	fmt.Println(sumIntOrFloat64(numFloat64))

	fmt.Println("------------------------")
	games := []Game{
		{
			Title:    "Game1",
			Platform: "IOS",
			Price:    30,
		},
		{
			Title:    "Game2",
			Platform: "Android",
			Price:    20,
		},
	}

	movies := []Movie{
		{
			Title: "Movie1",
			Price: 10,
		},
		{
			Title: "Movie2",
			Price: 20,
		},
	}
	fmt.Println(sumPriceGameOrMovie(games))
	fmt.Println(sumPriceGameOrMovie(movies))

	fmt.Println("------------------------")

	c := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for i := range c {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%v ", c[i])
		}(i)
	}
	wg.Wait()

	fmt.Println("\n------------------------")
	fmt.Println("ss")
}

type Game struct {
	Title    string
	Platform string
	Price    int
}
type Movie struct {
	Title string
	Price int
}

type GameOrMovie interface {
	getPrice() int
}

func (g Game) getPrice() int {
	return g.Price
}
func (g Movie) getPrice() int {
	return g.Price
}

type Number interface {
	int | float64
}

func sumPriceGameOrMovie[V GameOrMovie](objs []V) int {
	var sum int
	for _, obj := range objs {
		sum += obj.getPrice()
	}
	return sum
}

func sumIntOrFloat64[V Number](nums []V) V {
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}
func sumInt(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
func sumFloat64(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

type IBook interface {
	GetBook() *book
	SetTitle(title string)
}

func NewBook(title, author string) IBook {
	return &book{
		models.Book{
			Id:     1,
			Title:  title,
			Author: author,
		},
	}
}

type book struct {
	models.Book
}

func (b *book) SetTitle(title string) {
	b.Title = title
}

func (b *book) GetBook() *book {
	return b
}

func Debug(obj any) {
	raw, _ := json.MarshalIndent(&obj, "", "\t")
	fmt.Println(string(raw))
}

func printFullName(firstName string, lastName string) {
	fmt.Println(firstName + " " + lastName)
}

func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

// คืนค่า float และ error ออกไปพร้อมกันจากฟังก์ชัน
func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		err := errors.New("Division by zero!")
		return 0.0, err
	}
	return dividend / divisor, nil
}

// แม้ภาษา Go จะไม่มีคลาส แต่เรามี structs ที่สามารถนิยามโครงสร้างของข้อมูลขึ้นมาเองได้
type human struct {
	name    string
	age     int
	isAdult bool
}

type parrot struct {
	name string
	age  int
}

type talker interface {
	talk()
}

func (h human) talk() {
	fmt.Println("Human - I'm talking.")
}

func (p parrot) talk() {
	fmt.Println("Parrot - I'm talking.")
}

func (h human) printInfoHuman() {
	fmt.Println(h.name, h.age)
}

// ใช้ * แทนการ dereference หรือการถอดเอาค่าที่แท้จริงออกมา
func setAdult(h *human) {
	h.isAdult = h.age >= 18
}

func search(keyword string) {
	folders := [3]string{"Document", "Image", "Library"}
	var wg sync.WaitGroup
	// จำนวน goroutines เท่ากับ 3 อันเป็นความยาวของอาร์เรย์ folders
	wg.Add(len(folders))
	// และเพื่อป้องกันไม่ให้ Go หยุดการทำงานไปในทันที เราจึงต้อง Wait จนกว่า Goroutines จะทำงานเสร็จหมด
	for _, folder := range folders {
		// เราสามารถสร้างการทำงานแบบ Concurrency ได้ด้วยการใช้ Goroutines เพียงแค่เติม go เข้าไปข้างหน้าฟังก์ชัน ทุกอย่างก็จะสดชื่น
		// เราต้องส่ง reference ของ wg ไปด้วย เพื่อที่จะสั่ง Done
		go searchFromFolder(keyword, folder, &wg)
	}
	wg.Wait()
}

// คำถามถัดมา Go จะรู้ได้อย่างไรว่า Goroutine ตัวไหนทำงานเสร็จแล้วบ้าง เราจึงต้องสั่ง Done ในแต่ละ routine เพื่อบอกว่าการทำงานของมันเสร็จสิ้นแล้ว
// โปรดสังเกต เราต้องรับพอยเตอร์ของ sync.WaitGroup เข้ามาด้วย
func searchFromFolder(keyword string, folder string, wg *sync.WaitGroup) {
	// ทำการค้นหา
	// เมื่อค้นหาเสร็จ ต้องแจ้งให้ WaitGroup ทราบว่าเราทำงานเสร็จแล้ว
	// WaitGroup จะได้นับถูกว่าเหลือ Goroutines ที่ต้องรออีกกี่ตัว
	wg.Done()
}

func sum(name string, a, b int) (int, string) {
	return a + b, name + "-sum"
}
