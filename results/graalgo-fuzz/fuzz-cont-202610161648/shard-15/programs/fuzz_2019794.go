package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 string, p1 int8, p2 int64) int32 {
    switch p0 {
    case "a":
        var v4 int64 = 1
        _ = v4
    case "foo bar":
        v5 := []int{9223372036854775807, 0, 0}
        _ = v5
    default:
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: p1}
            _ = fi6
            _ = v7
        }
    }
    p2 += (p2 / ((0) | 1))
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = 1.5
    var v3 uint64 = 2
    v8 := GenS1{F0: 1}
    v9 := GenS1{F0: -128}
    var v10 uint64 = v3
    var v11 string = "\n\t"
    fmt.Println(v1, v1, v10)
    var v12 bool = false
    var v13 int8 = 1
    switch 2147483647 {
    case 32767:
        v14 := GenS2{F0: (v13 ^ v13)}
        _ = v14
    case -597:
        var v15 bool = (18446744073709551615 < 2)
        _ = v15
    case 127:
        fmt.Println(v3)
    }
    if false {
        v1 = v1
    }
    var v16 uint = 2
    var v17 int8 = v13
    v18 := []int{127, 127}
    v12 = true
    var v19 uint64 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
