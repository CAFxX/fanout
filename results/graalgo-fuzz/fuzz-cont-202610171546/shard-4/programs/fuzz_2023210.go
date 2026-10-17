package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
func genFn0(p0 uint) uint {
    var v5 uint = (^(p0))
    var v6 uint = v5
    var v7 float64 = 1e10
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1() int32 {
    var v8 float64 = 1.5
    if true {
        fmt.Println(v8)
    }
    _ = v8
    return 489
}

func main() {
    var v1 float64 = 0.0
    v2 := []uint32{42, 628}
    var v3 uint8 = 127
    v4 := []string{"foo bar"}
    fmt.Println(v1)
    v1 -= 0.1
    var v9 string = "foo bar"
    var v10 string = "a"
    for fi11 := 0; fi11 < 11; fi11++ {
        v2 = append(v2, 2)
        for fi12 := 0; fi12 < 15; fi12++ {
            v9 = v9
            var v13 int = fi11
            v14 := GenS1{F0: 0, F1: v13}
            _ = fi12
            _ = v13
            _ = v14
        }
        if true {
            v1 += v1
        } else {
            var v15 int16 = 1
            v9 = v9
            v15 = (v15 + (-(v15)))
            _ = v15
        }
        _ = fi11
    }
    switch 4294967295 {
    case -2147483648:
        v1 -= 0.1
    case -273:
        fmt.Println(v9)
    case 854:
        _ = len(v4)
    default:
        _ = len(v4)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v10)
    fmt.Println(v9, v1, v10)
    var v16 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
