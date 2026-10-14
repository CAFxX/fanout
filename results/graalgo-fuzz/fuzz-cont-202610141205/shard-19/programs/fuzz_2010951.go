package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 918
    var v3 int = 671
    if (true && true) {
        v3 += v3
    }
    _ = len(v1)
    v3 = v3
    var v4 bool = false
    v4 = (-2147483648 != 1)
    fmt.Println(v3, v2, v3)
    _ = len(v1)
    var v5 float64 = 1.5
    var v6 float64 = 3.14159
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
