package main

import (
    "fmt"
)

func genFn0() uint16 {
    var v4 uint16 = 0
    if false {
        v4 = v4
    }
    _ = v4
    return (v4 >> ((162) & 63))
}

func main() {
    var v1 uint8 = 230
    var v2 int64 = 1
    var v3 uint64 = 18446744073709551615
    v2 = (v2 % ((42) | 1))
    if true {
        var v5 uint16 = 42
        var v6 float64 = -2.25
        v2 -= 1
        _ = v5
        _ = v6
    } else {
        v7 := []int{32768}
        _ = v7
    }
    fmt.Println(v1)
    v8 := []int{0, -895}
    v8 = append(v8, -9223372036854775808)
    if (2 > 2) {
        var v9 float32 = 1e10
        switch 128 {
        case 32767:
            fmt.Println(v2)
        default:
            v1 -= (^(v1))
        }
        _ = v9
    }
    var v10 float32 = -2.25
    v3 = v3
    var v11 bool = true
    v8 = append(v8, 456)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
}
