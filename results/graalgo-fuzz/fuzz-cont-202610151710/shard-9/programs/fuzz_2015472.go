package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
func genFn0() uint8 {
    var v7 string = "hello"
    var v8 uint8 = 127
    var v9 float64 = 3.14159
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{2147483647, 0, 33, 2147483647}
    v2 := []uint32{1, 84, 42, 42}
    var v3 uint64 = 0
    v4 := make(map[int]string)
    v5 := []string{"a", "foo bar", "foo bar"}
    var v6 int16 = -1
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(v6, fi10)
        _ = fi10
    }
    if false {
        var v11 int = 846
        _ = v11
    }
    var v12 int64 = 1
    v13 := GenS1{F0: 350, F1: 2147483647, F2: (v12 + v12)}
    switch "a" {
    case "a":
        _ = len(v1)
    case "\n\t":
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int64 = (^(v12))
            _ = fi14
            _ = v15
        }
    case "foo bar":
        v12 += v12
    default:
        if false {
            for fi16 := 0; fi16 < 3; fi16++ {
                v12 -= -1
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 int = (fi17 - (fi17 >> ((18446744073709551615) & 63)))
                    _ = fi17
                    _ = v18
                }
                switch "a" {
                case "hello":
                    v12 -= (v12 | int64(v3))
                case "":
                    var v19 float64 = (0.1 + ((1e10 * ((float64(float32((1e10 * 0.1))) / 1.0) * (0.0 * 1e10))) - ((0.0 * 3.14159) - (1.5 * 0.1))))
                    _ = v19
                default:
                    fmt.Println(v12)
                }
                _ = fi16
            }
            var v20 uint8 = 120
            _ = v20
        }
    }
    v21 := GenS1{F0: 0, F1: 1, F2: (v12 & (-(v12)))}
    var v22 int = -1
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 1; fi24++ {
            v3 += (v3 >> ((51) & 63))
            _ = fi24
        }
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
