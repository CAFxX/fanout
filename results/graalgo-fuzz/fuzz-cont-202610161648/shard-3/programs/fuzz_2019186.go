package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 42
    var v3 float64 = 1.5
    v2 = 42
    v2 = 2
    fmt.Println(v2)
    v3 = v3
    for fi4 := 0; fi4 < 16; fi4++ {
        fmt.Println(fi4, fi4, v2)
        _ = fi4
    }
    var v5 uint32 = 2
    _ = len(v1)
    fmt.Println(v2)
    var v6 int16 = 42
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
