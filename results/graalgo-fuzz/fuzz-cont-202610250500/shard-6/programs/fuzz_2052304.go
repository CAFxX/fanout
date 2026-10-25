package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string; F1 int8 }
func genFn0() uint32 {
    var v4 uint16 = 1
    fmt.Println(v4)
    var v5 string = "foo bar"
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-257, 1, 1}
    var v3 float64 = (1.5 * -2.25)
    fmt.Println(v3)
    var v6 uint64 = 42
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint32 = 2147483648
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            v3 = v3
        }
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: true, F1: 42}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 32768)
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint32 = 792
        v11 = 42
        _ = fi10
        _ = v11
    }
    switch 1 {
    case -9223372036854775808:
        var v12 float64 = ((((float64((0.1 * float32(((-2.25 + (v3 * (3.14159 + (0.1 / float64((float32((v3 * float64(0.0))) - 1e10)))))) + (-2.25 * float64(1e10)))))) * (((((v3 / ((1.5 + 3.14159) - (v3 - v3))) + -2.25) / ((v3 / 1.0) - v3)) - (float64(-2.25) * v3)) * (1e10 / -2.25))) - 3.14159) - (1e10 * float64(1.5))) / (v3 + 1.5))
        _ = v12
    case 701:
        if (1 < 9223372036854775807) {
            var v13 int8 = -1
            v6 -= v6
            for fi14 := 0; fi14 < 19; fi14++ {
                fmt.Println(v6, v6, v13)
                _ = fi14
            }
            _ = v13
        }
    case -772:
        var v15 int = 4294967295
        _ = v15
    default:
        v16 := GenS2{F0: "foo bar", F1: -128}
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
