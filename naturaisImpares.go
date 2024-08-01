package main
import "fmt"

func main() {
    var num int
    fmt.Print("Digite um número: ")
    fmt.Scan(&num)
    var cont int = 1
    var cont2 int = 0
    for cont2 < num {
        if cont % 2 != 0 {
            fmt.Print(cont)
            cont2 = cont2 + 1
        }
        cont = cont + 1
    }
}
