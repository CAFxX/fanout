package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 33
    v3 := []int{-619, -1}
    v4 := make(map[int]int)
    var v5 int32 = 201
    v6 := []string{"a", "a", "hello", "hello"}
    for fi7 := 0; fi7 < 8; fi7++ {
        if false {
            var v8 string = "hello"
            _ = v8
        }
        _ = fi7
    }
    fmt.Println(v2, v5)
    var v9 uint32 = 65535
    var v10 bool = ((0.1 * 1.5) > 1e10)
    if (!(!v10)) {
        fmt.Println(v10, v5, v5)
    }
    var v11 uint32 = v9
    v6 = append(v6, "hello")
    var v12 int8 = -75
    var v13 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
