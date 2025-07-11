package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

type Websites struct {
}

func maps() {
	websites := map[string]string{
		"Google\n":              "https://google.com\n",
		"Amazon Web Services\n": "https://aws.com\n",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedIn\n"] = "https://linkedin.com\n"
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)
}
