package main

import (
    "fmt"
)


func main() {
    var v1 float32 = float32(0.1)
    var v2 uint32 = 920
    v3 := make(map[uint32]int)
    var v4 string = "\n\t"
    var v5 float32 = float32(((0.1 / float64((v1 + v1))) * (float64((v1 + v1)) * 0.1)))
    var v6 uint = 603
    fmt.Println(v6, v2)
    fmt.Println(v4, v1)
    v7 := v3[v2]
    v4 = v4
    v4 = "hello"
    var v8 uint32 = 2147483647
    v9 := []int{v7}
    var v10 int = v7
    _ = len(v9)
    fmt.Println(v4)
    if (!false) {
        fmt.Println(v7, v4, v8)
    }
    switch v4 {
    case "foo bar":
        var v11 int16 = 475
        _ = v11
    case "a":
        fmt.Println(v4)
    default:
        v5 = v1
    }
    v4 = (v4 + (((v4 + "") + "a") + ((v4 + v4) + "\n\t")))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
