package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
func genFn0() bool {
    v6 := []int{-2147483648, 1}
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: false, F1: false}
        _ = fi7
        _ = v8
    }
    var v9 float32 = 1e10
    _ = v6
    _ = v9
    return (4294967295 <= 128)
}

func main() {
    var v1 uint8 = 2
    v2 := []int{}
    var v3 uint8 = 1
    var v4 float64 = 1e10
    var v5 int64 = -9223372036854775808
    switch "a" {
    case "foo bar":
        v10 := GenS1{F0: ((!false) || (!false)), F1: false}
        _ = v10
    case "hello":
        var v11 int16 = 32767
        _ = v11
    case "":
        v1 -= v3
    default:
        fmt.Println(v1)
    }
    v3 = v1
    for fi12 := 0; fi12 < 14; fi12++ {
        v5 = (-(v5))
        fmt.Println(v3, v5, fi12)
        _ = fi12
    }
    v2 = append(v2, 9223372036854775807)
    fmt.Println(v3, v1)
    if ((0.0 * 0.1) == 0.0) {
        v13 := GenS1{F0: (!(uint(v1) == 65535)), F1: (v5 < (v5 + (v5 & (^(v5)))))}
        fmt.Println(v1, v5, v3)
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 3; fi14++ {
            if true {
                var v15 uint32 = 1
                if (-433 <= 478) {
                    var v16 float64 = ((((0.0 * v4) / (1e10 * (float64(-2.25) / -2.25))) + ((1e10 - (float64(3.14159) + v4)) - v4)) * float64(((1e10 + 1e10) + 3.14159)))
                    v15 -= 615
                    _ = v16
                }
                _ = v15
            } else {
                if true {
                    fmt.Println(fi14, v3, v5)
                    fmt.Println(v4, v4, v3)
                } else {
                    var v17 float64 = 3.14159
                    v2 = append(v2, (fi14 | (fi14 + (fi14 << ((255) & 63)))))
                    v18 := GenS1{F0: true, F1: (1 >= 1)}
                    _ = v17
                    _ = v18
                }
                v3 -= (v3 << ((4294967295) & 63))
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int16 = 42
                _ = fi19
                _ = v20
            }
            _ = fi14
        }
    }
    var v21 int32 = 0
    if len(v2) > 0 {
        v2[0] = 606
    }
    switch "a" {
    case "foo bar":
        v3 = v1
    default:
        v3 = (v1 * v1)
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint32 = 1
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: false, F1: true}
    v1 = v3
    v2 = append(v2, int(v1))
    var v25 bool = (!(881 <= 42))
    switch "" {
    case "\n\t":
        var v26 int32 = (v21 ^ 0)
        _ = v26
    case "foo bar":
        for fi27 := 0; fi27 < 6; fi27++ {
            v4 = 1e10
            v28 := GenS1{F0: (!((v5 ^ v5) == v5)), F1: (!v25)}
            v25 = (42 < 1)
            _ = fi27
            _ = v28
        }
    case "a":
        fmt.Println(v3)
    }
    v3 -= v3
    var v29 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
