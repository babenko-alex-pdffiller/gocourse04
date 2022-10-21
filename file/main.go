package main

func main() {
	// 1 write
	// data := []byte("Hello Bold!")
	// text := "Hello Gold!"
	// file, err := os.Create("hello.txt")
	// if err != nil {
	// 	fmt.Println("Unable to create file:", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// _, _ = file.WriteString(text)
	// _, _ = file.Write(data)
	// fmt.Println("Done.")

	// 2.1 read
	// content, err := os.ReadFile("hello.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Print(string(content))

	// 2.2 read
	// file, err := os.Open("hello.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// data := make([]byte, 64)

	// for {
	// 	n, err := file.Read(data)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(string(data[:n]))
	// }

	// 3 remove
	// err := os.Remove("hello.txt")
	// if err != nil {
	// log.Fatal(err)
	// }
	// err = os.RemoveAll("else")
	// if err != nil {
	// log.Fatal(err)
	// }
}
