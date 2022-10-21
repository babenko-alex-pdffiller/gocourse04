package main

func main() {
	// 1
	// _ = regexp.MustCompile(`+++`)

	// _, err := regexp.Compile(`+++`)
	// if err != nil {
	// fmt.Println(err)
	// }

	// matched, _ := regexp.MatchString(`ayayay`, "o no")
	// fmt.Println(matched) // false

	// 2 simple matches
	// matched, _ := regexp.MatchString(`I am here`, "I am there")
	// fmt.Println(matched) // false
	//
	//matched, _ := regexp.MatchString(`Banana`, "Banana")
	//fmt.Println(matched) // true
	//
	// matched, _ := regexp.MatchString(`cat`, "black cat meow")
	// fmt.Println(matched) // true
	//
	// re, _ := regexp.Compile(`cat`)
	// res := re.FindAllString("black cat meowcat", -1)
	// fmt.Println(res) // [cat cat]

	// 3 anchors
	// matched, _ := regexp.MatchString(`^I am here`, " I am here")
	// fmt.Println(matched) // false
	//
	// matched, _ := regexp.MatchString(`^cat$`, "Black cat meow")
	// fmt.Println(matched) // false
	//
	// matched, _ := regexp.MatchString(`\bcat\b`, "Black cat meow")
	// fmt.Println(matched) // true
	//
	// matched, _ := regexp.MatchString(`\Bcat\b`, "Blackcat meow")
	// fmt.Println(matched) // true

	// 4 symbol classes
	// matched, _ := regexp.MatchString(`.....`, "qw2 )")
	// fmt.Println(matched) // true
	//
	// matched, _ := regexp.MatchString(`^\w\wow\d\b.\D\Dow\d$`, "meow3_meow4")
	// fmt.Println(matched) // false
	//
	// matched, _ := regexp.MatchString(`^...@_\w\wD$`, "ukr@_UPD")
	// fmt.Println(matched) // true
	//
	// matched, _ := regexp.MatchString(`^GO\s\d.\d\d$`, "GO 1.16")
	// fmt.Println(matched) // true

	// 5 spec symbols and escapes
	// matched, _ := regexp.MatchString(`^I\nam\nhere$`, "I\nam\nhere")
	//fmt.Println(matched) // true
	//
	// matched, _ := regexp.MatchString(`^\x49\nam\nhere$`, "I\nam\nhere")
	// fmt.Println(matched) // true
	//
	// matched, _ := regexp.MatchString(`a+b=c`, "a+b=c")
	// fmt.Println(matched) // false
	//
	// matched, _ := regexp.MatchString(`a\|b=c`, "a|b=c")
	// fmt.Println(matched) // true

	// 6 repetition
	// re, _ := regexp.Compile(`\d+`)
	// res := re.FindAllString("A123AA455AAA2A89", -1)
	// fmt.Println(res) // [123  455  2  89]
	//
	// matched, _ := regexp.MatchString(`^A{1}G{1,3}A{1,}!{,2}$`, "AGGAA!!")
	// fmt.Println(matched) // false (be careful, only python, с {0,2} behavior correct)
	//
	// re, _ := regexp.Compile(`<.*>`)
	// res := re.FindAllString("<p><b>Golang</b> <i>VS</i> <b>Python</b></p>", -1)
	// fmt.Println(res) // [<p><b>Golang</b> <i>VS</i> <b>Python</b></p>] (len=1) :(
	//
	// re, _ := regexp.Compile(`<.*?>`)
	// res := re.FindAllString("<p><b>Golang</b> <i>VS</i> <b>Python</b></p>", -1)
	// fmt.Println(res) // [<p>  <b>  </b>  <i>  </i>  <b>  </b>  </p>] :)

	// 7 square brackets
	// matched, _ := regexp.MatchString(`good|bad|[^ice\s]$`, "work is hmm")
	// fmt.Println(matched) // true
	//
	// matched, _ := regexp.MatchString(`^[aAuUfF]*?go|python [1-3]\.\d$`, "Uf go 1.6")
	// fmt.Println(matched) // false
	//
	// matched, _ := regexp.MatchString(`^[haHA]+$`, "HahaHaaaahaaaaaaa")
	// fmt.Println(matched) // true
	//
	//matched, _ := regexp.MatchString(`^[haHA]+|[goGO]*$`, "")
	//
	//fmt.Println(matched) // true

	// 8 groups
	//re, _ := regexp.Compile(`.(\d+)`)
	//res := re.FindAllStringSubmatch("Funny1 2020 ye12ar", -1)
	//fmt.Println(res) // [[y1  1] [ 2020  2020] [e12  12]]

	//re, _ := regexp.Compile(`(\d{4})-(\d{2})-(\d{2})`)
	//res := re.FindAllStringSubmatch("Now is 2021-01-14", -1)
	//fmt.Println(res) // [[2021-01-14  2021  01  14]]
	//
	//re, _ := regexp.Compile(`.*?(([a-zA-Z\-0-9]+)\\.[a-zA-Z]{2,})`)
	//res := re.FindAllStringSubmatch("version: v2-v3\\\\Go", -1)
	//fmt.Println(res) // [[version: v2-v3\\Go  v2-v3\\Go  v2-v3]]
}
