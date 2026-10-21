package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
func genFn0() uint16 {
    var v6 uint8 = 0
    switch "" {
    case "hello":
        var v7 int64 = 0
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float64 = ((((-2.25 - 1e10) - 3.14159) / 1.5) + ((1.5 / -2.25) - 1e10))
            _ = fi8
            _ = v9
        }
    case "a":
        var v10 uint32 = 1
        _ = v10
    default:
        v11 := GenS1{F0: "", F1: (v6 | 0)}
        _ = v11
    }
    v6 += v6
    _ = v6
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{4294967295, 2147483647, 2147483647}
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]bool)
    var v12 uint8 = 2
    v12 += v12
    v13 := GenS1{F0: ("" + "a"), F1: (v12 & (^(v12)))}
    for fi14 := 0; fi14 < 3; fi14++ {
        if (!false) {
            var v15 int8 = 0
            v15 = v15
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            fmt.Println(fi16, fi14, fi14)
            switch "" {
            case "\n\t":
                v12 -= (v12 << ((1) & 63))
            case "hello":
                if (true && false) {
                    v2 = append(v2, 0)
                }
            case "":
                switch fi16 {
                case 71:
                    if false {
                        var v17 uint16 = 407
                        _ = v17
                    } else {
                        v12 += 201
                    }
                }
            default:
                v12 -= v12
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi16
        }
        _ = fi14
    }
    var v18 uint32 = 4294967295
    _ = len(v3)
    var v19 int64 = -1
    var v20 uint16 = 0
    v21 := GenS1{F0: "foo bar", F1: (v12 >> ((18446744073709551615) & 63))}
    _ = len(v2)
    v20 = uint16(v12)
    if len(v3) > 0 {
        v3[0] = v18
    }
    var v22 uint64 = 972
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
