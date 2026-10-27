package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool }
func genFn0() int8 {
    switch 0 {
    case -723:
        v6 := []int{}
        _ = v6
    case -2147483648:
        var v7 int8 = -1
        _ = v7
    case -802:
        var v8 float64 = (0.0 + 1e10)
        _ = v8
    default:
        var v9 int32 = -2147483648
        _ = v9
    }
    var v10 int8 = 119
    v10 += (v10 & (-(v10)))
    if (255 > 42) {
        fmt.Println(v10)
        var v11 float64 = -2.25
        switch -9223372036854775808 {
        case 4294967295:
            v11 += v11
        }
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        if (-2147483648 > int32(v10)) {
            v10 += v10
            v10 = -59
            var v13 string = "a"
            _ = v13
        }
        v10 += v10
        _ = fi12
    }
    _ = v10
    return 127
}

func main() {
    var v1 uint64 = 78
    var v2 int64 = -507
    v3 := []int{2147483647, -1}
    var v4 string = "foo bar"
    var v5 float64 = 3.14159
    var v14 float64 = 3.14159
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint64 = v1
        _ = fi15
        _ = v16
    }
    var v17 int8 = 64
    var v18 int16 = 1
    v19 := GenS1{F0: 32767, F1: 144}
    var v20 int = 0
    var v21 string = (v4 + v4)
    fmt.Println(v2, v2)
    var v22 bool = (v1 == v1)
    switch v4 {
    case "\n\t":
        v14 = float64((0.1 / 0.1))
    }
    if (v22 && (9223372036854775807 <= (v20 / ((v20) | 1)))) {
        var v23 string = "a"
        v24 := GenS2{F0: v22}
        fmt.Println(v23)
        _ = v23
        _ = v24
    }
    fmt.Println(v2)
    var v25 int = (-(v20))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
