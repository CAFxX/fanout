package main

import (
    "fmt"
)


func main() {
    v1 := []int{588}
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v4 uint32 = 2147483647
    var v5 int64 = 1
    v5 = -1
    var v6 bool = (127 == 2147483647)
    var v7 float64 = 0.1
    fmt.Println(v7)
    for fi8 := 0; fi8 < 7; fi8++ {
        v7 = float64(1e10)
        v6 = true
        _ = fi8
    }
    v4 = 2
    if (349 >= 0) {
        fmt.Println(v5, v5)
        var v9 float32 = (float32((v7 * 1.5)) / (((1.5 + 1.5) + -2.25) / 1.5))
        var v10 string = ("hello" + "\n\t")
        _ = v10
        _ = v9
    } else {
        var v11 int8 = 42
        _ = v11
    }
    v1 = append(v1, 2147483647)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
