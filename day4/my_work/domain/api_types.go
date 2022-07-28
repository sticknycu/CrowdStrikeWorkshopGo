package domain

type User struct {
	Name      string   `json:"name"`
	Mail      string   `json:"mail"`
	Age       int      `json:"age"`
	Interests []string `json:"interests"`
}

func main() {

}
