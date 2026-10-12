package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 bool; F2 int16 }
func genFn0(p0 bool) uint16 {
    var v4 int16 = 1000
    if true {
        if p0 {
            for fi5 := 0; fi5 < 1; fi5++ {
                fmt.Println(v4)
                _ = fi5
            }
            v6 := GenS1{F0: 513, F1: false}
            v4 += v4
            _ = v6
        }
    } else {
        fmt.Println(p0, p0)
        v4 = 1000
    }
    fmt.Println(p0)
    var v7 float32 = 0.0
    v8 := GenS1{F0: 9223372036854775807, F1: p0}
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []int{}
    var v2 float64 = ((0.0 + 0.0) + (3.14159 + 0.0))
    v3 := make(map[string]int)
    v9 := GenS1{F0: 9223372036854775807, F1: (42 <= 604)}
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS1{F0: 1, F1: false}
        _ = fi10
        _ = v11
    }
    v2 -= ((((0.0 - 0.1) - ((v2 + 1e10) / 3.14159)) * (1e10 - v2)) * float64(1.5))
    var v12 float64 = ((v2 * v2) - (0.0 + (float64((1.5 + ((1e10 - 1e10) / 1.5))) - v2)))
    var v13 uint = 721
    v2 -= v2
    var v14 float32 = (0.0 / ((1.5 - 0.0) / ((0.1 - ((3.14159 + 3.14159) - 0.0)) * (((1.5 + 1.5) + 1e10) + 0.0))))
    var v15 float64 = 1e10
    switch ("\n\t" + ("" + "hello")) {
    case "foo bar":
        v3[("a" + "hello")] = 2147483647
    case "\n\t":
        var v16 uint16 = 241
        _ = v16
    case "":
        var v17 uint = (^(v13))
        _ = v17
    default:
        v18 := GenS1{F0: 9223372036854775807, F1: false}
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            switch fi20 {
            case 128:
                v14 += ((0.0 + v14) - -2.25)
            case 4294967295:
                v21 := GenS2{F0: ("\n\t" + "a")}
                _ = v21
            case 9223372036854775807:
                v12 += v12
            }
            _ = fi20
        }
        var v22 int32 = 2147483647
        _ = fi19
        _ = v22
    }
    v15 = 0.1
    fmt.Println(v14, v14)
    var v23 int16 = 355
    fmt.Println(v15)
    var v24 uint16 = 65535
    var v25 uint = (v13 | v13)
    _ = len(v3)
    v26 := GenS3{F0: 0, F1: ((127 >= 1) && false), F2: v23}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v9
}
