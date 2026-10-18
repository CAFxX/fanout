package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    var v3 int8 = -109
    var v4 int16 = -1
    v2 -= 778
    var v5 bool = false
    fmt.Println(v2)
    fmt.Println(v4, v4, v2)
    var v6 int = -1
    var v7 uint8 = 0
    var v8 int64 = v2
    switch v6 {
    case 32767:
        var v9 string = ""
        _ = v9
    default:
        v10 := []int{(v6 / ((v6) | 1)), v6, v6, v6}
        _ = v10
    }
    v8 += (v2 & 0)
    v6 += v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
