package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 bool = (0 >= 23)
    var v4 int8 = 35
    var v5 uint64 = 216
    v6 := GenS1{F0: 2, F1: 0}
    if (v5 != v5) {
        var v7 int32 = 0
        v7 = v7
        _ = v7
    }
    fmt.Println(v5, v4, v5)
    v2[346] = "a"
    fmt.Println(v4)
    if false {
        v3 = true
        if (!(("\n\t" > ("\n\t" + "\n\t")) && (false || v3))) {
            v8 := GenS1{F0: 42, F1: -2147483648}
            v9 := []int{0, 32767}
            v10 := GenS1{F0: 1, F1: -531}
            _ = v10
            _ = v8
            _ = v9
        }
    } else {
        v3 = v3
        var v11 int32 = -32
        var v12 int8 = 127
        _ = v11
        _ = v12
    }
    var v13 float64 = float64(((3.14159 / (3.14159 / (0.1 + 1e10))) / 1e10))
    v14 := []int{-2147483648, -371, 9223372036854775807, 2147483647}
    v15 := v1[42]
    if false {
        v16 := GenS2{F0: 255}
        fmt.Println(v3)
        _ = v16
    } else {
        var v17 string = "foo bar"
        delete(v1, 42)
        _ = v17
    }
    v4 += (v4 | v4)
    v4 -= (v4 & v4)
    v1[int(v4)] = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
