package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
func genFn0(p0 int) uint32 {
    var v6 uint = 255
    var v7 string = "foo bar"
    _ = p0
    _ = v6
    _ = v7
    return 410
}

func main() {
    var v1 uint64 = 2
    var v2 float64 = (0.1 + -2.25)
    var v3 uint64 = 825
    v4 := make(map[int]int)
    var v5 uint8 = 164
    v4[786] = 32768
    if false {
        v8 := GenS1{F0: 1000}
        v9 := GenS1{F0: 0}
        v1 += (v1 | 0)
        _ = v8
        _ = v9
    } else {
        v3 = v3
        var v10 string = ""
        _ = v10
    }
    v4[-536] = -989
    var v11 int16 = 649
    v12 := GenS1{F0: -1}
    v2 -= (3.14159 - v2)
    v3 -= (v3 * (v3 >> ((4294967295) & 63)))
    v13 := GenS2{F0: 42, F1: "hello", F2: 4294967295}
    fmt.Println(v1)
    v3 += v3
    if true {
        var v14 int = 984
        var v15 int = v14
        _ = v14
        _ = v15
    }
    var v16 uint32 = 2147483647
    v1 -= v3
    v1 -= 0
    var v17 float64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
