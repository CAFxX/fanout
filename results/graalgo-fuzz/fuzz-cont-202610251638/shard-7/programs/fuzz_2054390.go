package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 float32 = 1e10
    v3 := []int{}
    v4 := make(map[int]int)
    var v5 bool = (1 >= 1)
    v6 := v4[9223372036854775807]
    var v7 float64 = ((((3.14159 / 1.0) + (0.1 - 0.0)) * 3.14159) - -2.25)
    for fi8 := 0; fi8 < 18; fi8++ {
        v5 = true
        _ = fi8
    }
    fmt.Println(v2)
    fmt.Println(v7, v2, v7)
    v5 = true
    var v9 int32 = 965
    var v10 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
