package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint16, p1 bool) bool {
    v6 := []int{9223372036854775807, -424, 32767}
    if p1 {
        var v7 float32 = 0.0
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float32 = v7
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return true
}

func main() {
    var v1 float32 = 0.1
    var v2 int8 = -23
    var v3 int64 = 9223372036854775807
    v4 := []int{}
    v5 := []int{-678, 42}
    v4 = append(v4, -626)
    var v10 int32 = 1
    switch ((("\n\t" + "") + "hello") + "\n\t") {
    case "a":
        var v11 float64 = -2.25
        _ = v11
    default:
        v12 := GenS1{F0: (^(v10))}
        _ = v12
    }
    if false {
        var v13 int32 = (-(v10))
        _ = v13
    }
    fmt.Println(v3)
    var v14 uint = 0
    fmt.Println(v14, v10, v2)
    v2 = 0
    v3 -= (v3 >> ((2) & 63))
    var v15 uint32 = 42
    fmt.Println(v3)
    v14 = 1
    fmt.Println(v15, v2, v3)
    fmt.Println(v2, v10)
    v16 := GenS1{F0: (v10 - v10)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
