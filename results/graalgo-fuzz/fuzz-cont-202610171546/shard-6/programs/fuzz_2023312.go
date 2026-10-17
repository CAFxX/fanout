package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0() string {
    if (-2147483648 >= 42) {
        var v4 int64 = -633
        if (!true) {
            v4 += v4
            var v5 int = 32768
            _ = v5
        }
        v4 -= (v4 << ((42) & 63))
        _ = v4
    } else {
        v6 := GenS1{F0: (2 != 2147483647)}
        var v7 bool = false
        switch "hello" {
        case "foo bar":
            var v8 uint16 = 0
            _ = v8
        default:
            if v7 {
                var v9 int16 = 42
                var v10 int16 = 1000
                v7 = v7
                _ = v10
                _ = v9
            }
        }
        _ = v6
        _ = v7
    }
    var v11 uint64 = 1
    if (!true) {
        v11 += v11
    }
    _ = v11
    return ""
}

func main() {
    var v1 int8 = 0
    var v2 uint = 687
    var v3 uint64 = 1
    var v12 float64 = ((-2.25 * (3.14159 / float64(1e10))) + (0.0 / 1.5))
    fmt.Println(v2, v12, v2)
    fmt.Println(v12)
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(fi13, v12)
        var v14 int64 = 42
        _ = fi13
        _ = v14
    }
    v12 = float64(1e10)
    v15 := []int{-9223372036854775808}
    fmt.Println(v12, v3)
    var v16 float64 = (v12 / -2.25)
    v17 := GenS2{F0: 255, F1: 1}
    var v18 uint32 = 65535
    v16 -= float64((3.14159 / 0.1))
    if len(v15) > 0 {
        v15[0] = 2147483647
    }
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
