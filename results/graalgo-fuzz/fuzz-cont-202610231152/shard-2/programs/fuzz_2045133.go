package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint16, p1 string) int {
    fmt.Println(p0, p1)
    var v6 bool = true
    var v7 int = 32768
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (v7 & (v7 - (v7 * (v7 << ((718) & 63)))))
}

func main() {
    v1 := []uint32{65535, 1, 1}
    v2 := make(map[string]int)
    var v3 uint8 = 25
    var v4 int16 = int16(v3)
    var v5 uint8 = v3
    fmt.Println(v3, v3)
    v5 -= (v3 >> ((2) & 63))
    switch "a" {
    case "":
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
    case "hello":
        var v9 int8 = -128
        _ = v9
    }
    fmt.Println(v5)
    v10 := GenS1{F0: 171, F1: (^(v4))}
    v1 = append(v1, 42)
    if (756 == 42) {
        switch genFn0(32767, "\n\t") {
        case -806:
            var v11 uint32 = 65535
            _ = v11
        case -2147483648:
            var v12 int32 = -2147483648
            _ = v12
        case 127:
            switch 32767 {
            case 754:
                fmt.Println(v3, v5, v3)
            default:
                var v13 uint32 = 65535
                _ = v13
            }
        default:
            v14 := GenS2{F0: "\n\t"}
            _ = v14
        }
    } else {
        var v15 float32 = (((0.0 - 0.0) / 0.1) + 1e10)
        var v16 bool = ((589 != 42) && true)
        _ = v15
        _ = v16
    }
    v5 += (v3 << ((0) & 63))
    var v17 int = -2147483648
    v4 = (v4 % (((^(v4))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
