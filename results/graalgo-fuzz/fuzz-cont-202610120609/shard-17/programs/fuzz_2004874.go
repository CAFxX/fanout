package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 65535
    var v2 bool = ((v1 % ((v1) | 1)) == v1)
    var v3 uint64 = 1
    var v4 uint64 = 2
    v5 := make(map[int]int)
    var v6 uint16 = v1
    switch 128 {
    case 0:
        v2 = v2
    case 128:
        v2 = (-10 >= -95)
    default:
        var v7 uint64 = 1
        _ = v7
    }
    var v8 int64 = -1
    var v9 int16 = 0
    v3 = v4
    fmt.Println(v1, v8, v3)
    switch "a" {
    case "a":
        var v10 string = "hello"
        _ = v10
    case "":
        if (128 <= -113) {
            delete(v5, 127)
            v3 -= (^(v3))
        }
    }
    _ = len(v5)
    v1 = 1
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float64 = float64(float32(1e10))
        _ = fi11
        _ = v12
    }
    v13 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
