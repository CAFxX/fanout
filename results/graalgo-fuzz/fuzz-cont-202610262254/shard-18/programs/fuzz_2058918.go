package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 bool }
func genFn0() uint8 {
    v4 := GenS1{F0: 642, F1: 785}
    if true {
        var v5 bool = false
        _ = v5
    } else {
        var v6 uint64 = 42
        v7 := []int{-1, -2147483648, 73, 1}
        _ = len(v7)
        _ = v6
        _ = v7
    }
    var v8 string = "a"
    _ = v4
    _ = v8
    return 127
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]uint32)
    var v3 uint = 250
    var v9 uint64 = 18446744073709551615
    var v10 string = ("foo bar" + "hello")
    for fi11 := 0; fi11 < 20; fi11++ {
        v9 += (v9 - v9)
        _ = fi11
    }
    v12 := GenS2{F0: -712, F1: false}
    v2[-1] = 42
    v9 = 0
    v3 += v3
    if false {
        v1 += 8
        v13 := GenS1{F0: 2, F1: 42}
        _ = len(v2)
        _ = v13
    }
    v3 += (^(v3))
    v9 += 787
    v10 = v10
    v14 := GenS2{F0: 32767, F1: (!false)}
    switch 32768 {
    case 2147483647:
        fmt.Println(v10)
    }
    v15 := []int{-1, 127}
    v1 -= (^(v1))
    var v16 string = v10
    v17 := GenS1{F0: v9, F1: 971}
    var v18 int16 = -71
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
