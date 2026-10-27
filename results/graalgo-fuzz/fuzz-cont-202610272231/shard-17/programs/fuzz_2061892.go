package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint }
func genFn0() int64 {
    var v5 float64 = ((((0.1 + 0.0) / (1e10 / 1.5)) + 0.0) + 3.14159)
    var v6 string = "hello"
    v7 := []int{128, -586, 127, 379}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: -993}
        v5 = 0.0
        switch "foo bar" {
        case "a":
            var v10 uint8 = 255
            _ = v10
        case "hello":
            fmt.Println(v5, v5, fi8)
        case "":
            v7 = append(v7, fi8)
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                v6 = v6
                var v12 string = ""
                v6 = "a"
                _ = fi11
                _ = v12
            }
        }
        _ = fi8
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int32 = -1
    var v2 float64 = 1e10
    var v3 int64 = -553
    var v4 float64 = 0.0
    var v13 int = 4294967295
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{fi14}
        var v16 int8 = 100
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: (v3 & (v3 ^ v3))}
    v18 := []int{(v13 * v13), (v13 % ((v13) | 1))}
    fmt.Println(v4, v13)
    v18 = append(v18, -1)
    v19 := GenS2{F0: 42, F1: -472, F2: 42}
    v18 = append(v18, -621)
    var v20 string = ("" + "\n\t")
    if (2 > 255) {
        var v21 float64 = v4
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS1{F0: -9223372036854775808}
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        _ = len(v18)
        for fi24 := 0; fi24 < 8; fi24++ {
            v18 = append(v18, fi24)
            v25 := GenS1{F0: (v3 ^ (-(v3)))}
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
