package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 int }
func genFn0() uint {
    if ((3.14159 * 1e10) < 1.5) {
        var v6 int64 = 984
        var v7 int32 = 1
        v8 := []int{127, 712}
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{-207, -811}
    _ = v9
    return 255
}
func genFn1(p0 string) uint8 {
    var v10 int32 = 2147483647
    p0 = (p0 + p0)
    switch "\n\t" {
    case "hello":
        for fi11 := 0; fi11 < 11; fi11++ {
            p0 = p0
            v10 -= (^(v10))
            v12 := GenS1{F0: 18446744073709551615, F1: (1 == 0)}
            _ = fi11
            _ = v12
        }
    case "\n\t":
        fmt.Println(v10, p0)
    case "":
        v10 += v10
    default:
        switch "" {
        case "hello":
            var v13 int16 = 42
            _ = v13
        case "foo bar":
            v14 := GenS2{F0: 81, F1: 1, F2: -2147483648}
            _ = v14
        }
    }
    v10 = (v10 & v10)
    _ = p0
    _ = v10
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{798, 418, -2147483648, -1}
    var v3 int64 = -9223372036854775808
    var v4 int = -406
    v5 := make(map[int]string)
    v4 += -868
    for fi15 := 0; fi15 < 16; fi15++ {
        v3 = v3
        var v16 string = ("a" + "foo bar")
        v17 := v1[4294967295]
        _ = fi15
        _ = v16
        _ = v17
    }
    v3 = (v3 & v3)
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 int8 = 0
            v21 := GenS2{F0: 380, F1: 954, F2: (fi18 / ((9223372036854775807) | 1))}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    v4 = -9223372036854775808
    var v22 uint64 = 1
    var v23 uint = 1
    v4 -= (v4 | v4)
    var v24 int32 = 744
    var v25 float32 = (0.1 * ((3.14159 + 1e10) / 0.1))
    v26 := GenS1{F0: v23, F1: false}
    var v27 int = -462
    for fi28 := 0; fi28 < 19; fi28++ {
        v24 -= -2147483648
        _ = fi28
    }
    var v29 uint = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
