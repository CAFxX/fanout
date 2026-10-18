package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int = 0
    var v2 int64 = 0
    var v3 uint = 65535
    var v4 bool = true
    var v5 bool = v4
    v4 = (999 != 751)
    if v5 {
        v2 -= 160
    } else {
        v2 = v2
        var v6 int32 = 0
        var v7 uint8 = 197
        _ = v6
        _ = v7
    }
    if v5 {
        v8 := []int{-507, v1, v1, (-(v1))}
        _ = v8
    } else {
        v9 := GenS1{F0: "hello"}
        v2 = (v2 / ((9223372036854775807) | 1))
        _ = v9
    }
    fmt.Println(v5, v3)
    v3 -= (^(v3))
    var v10 uint16 = 785
    fmt.Println(v10)
    v11 := []int{2147483647, -84, v1, 9223372036854775807}
    v4 = (-93 <= 1)
    fmt.Println(v2, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
