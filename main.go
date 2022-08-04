package main
import (
	"fmt"
	"github.com/shinvdu/go_module_poc/config"
)

func main() {
	a := config.GetConfig()
	c := config.GetDaoConfig()
	var b config.DaoConfig
	b.Username = "BCD"

    fmt.Println("hello world")
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    config.Xie()
    fmt.Println("hello world2")
}
