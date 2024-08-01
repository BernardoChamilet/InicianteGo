package main
import "fmt"

func main() {
    var x int
    fmt.Print("Digite um número: ")
    fmt.Scan(&x)
    var n int
    fmt.Print("Digite outro número: ")
    fmt.Scan(&n)
    var resultado int = 1
    for i := 1; i<=n; i++ {
        resultado = resultado * x
    }
    fmt.Print(resultado)
}
