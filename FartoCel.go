package main
import "fmt"

func main() {
    var far int
    fmt.Print("Digite a temperatura: ")
    fmt.Scan(&far)
    var cel int = (far-32)*5/9
    fmt.Print(cel)
}
