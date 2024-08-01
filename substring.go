package main
import (
	"bufio"
	"fmt"
	"os"
)


func main() {
    var palavra string
    fmt.Print("Digite uma palavra: ")
    fmt.Scan(&palavra)
    fmt.Print("Digite uma frase: ")
    reader := bufio.NewReader(os.Stdin)
    frase, _ := reader.ReadString('\n')
    var cont int = 0
    var cont2 int = 0
    var apareceu int = 0
    var tamanhoFrase int = len(frase)
    var tamanhoPalavra int = len(palavra)
    for cont < tamanhoFrase {
        if frase[cont] == palavra[cont2]{
            cont2 = cont2 + 1
            cont = cont + 1
            if cont2 == tamanhoPalavra{
                apareceu = apareceu + 1
                cont2 = 0
                if frase[cont-1] == palavra[cont2]{
                    cont = cont - 1
                }
            }
        } else{
            cont = cont - cont2 + 1
            cont2 = 0
        }
    }
    fmt.Print(apareceu)
}
