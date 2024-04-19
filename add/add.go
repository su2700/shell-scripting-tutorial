package main
import ( "fmt" )

func main() {
	var X int
	var Y int
	fmt.Println("add of X AND Y")
	fmt.Print("Enter X: ")
    _, err := fmt.Scan(&X)
    if err != nil {
        fmt.Println("Error reading X:", err)
        return
    }

    // Asking for input of Y
    fmt.Print("Enter Y: ")
    _, err = fmt.Scan(&Y)
    if err != nil {
        fmt.Println("Error reading Y:", err)
        return
    }

    // Calculating and displaying the result
    fmt.Printf("X + Y = %d + %d = %d\n", X, Y, X+Y)

}