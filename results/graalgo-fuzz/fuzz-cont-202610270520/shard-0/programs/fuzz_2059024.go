package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 int8, p1 uint64, p2 uint16) uint32 {
    v4 := GenS1{F0: 407, F1: 20}
    var v5 int = 2147483647
    var v6 float64 = 1e10
    v7 := []int{(v5 + 127), (v5 ^ (^(v5)))}
    if false {
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 256
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 704
    v3 := make(map[string]int)
    switch 127 {
    case -2147483648:
        var v8 int8 = 100
        _ = v8
    case -669:
        switch "" {
        case "\n\t":
            v1 = v1
        case "hello":
            v2 += v2
        }
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            v2 = (^(v2))
            _ = fi9
        }
    }
    v10 := []int{}
    var v11 uint8 = 0
    v2 = v2
    switch "" {
    case "":
        v2 = 385
    case "hello":
        var v12 uint16 = 32767
        _ = v12
    default:
        _ = len(v3)
    }
    if false {
        if false {
            switch 115 {
            case 2147483647:
                var v13 int32 = -338
                _ = v13
            case 178:
                fmt.Println(v1, v2)
            default:
                if false {
                    var v14 uint64 = (^(v2))
                    _ = v14
                }
            }
        } else {
            if false {
                fmt.Println(v2, v2, v1)
                var v15 bool = (0 > 9223372036854775807)
                if v15 {
                    v2 = v2
                    for fi16 := 0; fi16 < 12; fi16++ {
                        var v17 uint8 = v11
                        fmt.Println(v1)
                        var v18 uint16 = 42
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                _ = v15
            }
            v11 -= v11
            v19 := GenS1{F0: 0, F1: -2147483648}
            _ = v19
        }
        fmt.Println(v1, v11, v1)
        var v20 bool = (18446744073709551615 <= v2)
        _ = v20
    }
    v21 := v3[(("\n\t" + "hello") + ("foo bar" + ""))]
    var v22 int32 = -636
    switch "a" {
    case "":
        for fi23 := 0; fi23 < 16; fi23++ {
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 uint32 = 1
                _ = fi24
                _ = v25
            }
            v2 -= v2
            _ = fi23
        }
    case "hello":
        switch v21 {
        case 1:
            switch v21 {
            case -2147483648:
                v3[((("\n\t" + "") + (("a" + "hello") + "hello")) + "")] = (v21 >> ((255) & 63))
            case 624:
                _ = len(v10)
            case 32768:
                v22 -= 2147483647
            default:
                v21 = (v21 >> ((1) & 63))
            }
        case -606:
            v1 = v1
        default:
            v1 = (v1 + v1)
        }
    default:
        v11 -= 164
    }
    v21 = (v21 + v21)
    v26 := GenS1{F0: 0, F1: v22}
    v11 = 2
    var v27 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
}
