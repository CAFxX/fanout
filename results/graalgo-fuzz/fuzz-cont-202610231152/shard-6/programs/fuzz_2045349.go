package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int16 }
func genFn0(p0 int16) uint16 {
    v6 := []int{}
    if false {
        p0 = (p0 % ((-928) | 1))
        v7 := GenS1{F0: 0, F1: 2, F2: (p0 - p0)}
        _ = v7
    } else {
        var v8 int16 = p0
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 32767
}

func main() {
    var v1 bool = true
    var v2 float64 = (float64(1e10) * (0.0 + (-2.25 * 0.0)))
    v3 := make(map[int]uint32)
    var v4 float64 = 0.0
    v5 := make(map[string]bool)
    delete(v3, 127)
    v9 := []int{32767, 796, 32768}
    v1 = ((2 != 931) || (4294967295 == 585))
    v1 = ((18446744073709551615 < 1) && true)
    fmt.Println(v2, v1)
    for fi10 := 0; fi10 < 11; fi10++ {
        v2 = float64((-2.25 * 0.0))
        v2 -= (((v2 * (((v2 * v2) + v4) / v2)) + float64(((-2.25 + ((-2.25 * 1.5) + (0.0 - 3.14159))) - 0.1))) / (v4 * (v2 * v4)))
        var v11 bool = v1
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: 32767, F2: -32768}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint16 = 48
        var v15 uint64 = 2
        switch ("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + "a")) {
        case "foo bar":
            if true {
                var v16 bool = (v1 || (v1 && (1 == -2147483648)))
                v4 = 1e10
                for fi17 := 0; fi17 < 2; fi17++ {
                    v15 -= 78
                    v18 := v3[fi13]
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
        case "\n\t":
            var v19 int = 128
            _ = v19
        case "a":
            fmt.Println(v1, v1)
        default:
            v20 := GenS1{F0: -27, F1: genFn0(42), F2: 1000}
            _ = v20
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 -= v2
    switch ("hello" + "foo bar") {
    case "":
        v21 := v3[4294967295]
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
