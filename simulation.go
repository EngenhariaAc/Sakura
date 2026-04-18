package main

import "fmt"

func main() {
    petals := 100

    fmt.Println("🌸 Simulating sakura bloom...")

    for i := 0; i < 5; i++ {
        petals -= 10
        fmt.Printf("Petals remaining: %d\n", petals)
    }

    fmt.Println("🍃 Petals have fallen gracefully.")
}
