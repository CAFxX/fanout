package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int, p1 int8) uint16 {
    v4 := GenS1{F0: 65535, F1: (p1 >> ((18446744073709551615) & 63))}
    v5 := []int{}
    var v6 int16 = 32767
    var v7 float64 = 1e10
    var v8 uint = 668
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    var v3 bool = (false || (0 < 4294967295))
    switch 0 {
    case -2147483648:
        if v3 {
            if (132 <= 32767) {
                v9 := []int{-9223372036854775808, 379}
                v1 -= float64(1e10)
                var v10 string = "\n\t"
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 0, F1: -34}
            _ = v11
        } else {
            _ = len(v2)
            if false {
                var v12 int8 = 127
                var v13 int8 = (-(v12))
                v12 = (^(v13))
                _ = v12
                _ = v13
            }
            var v14 int16 = 1000
            _ = v14
        }
    case 433:
        v2[128] = (((!(6 >= 62)) || ((v3 || false) || false)) || (0 != 0))
    case 9223372036854775807:
        v1 -= v1
    default:
        v3 = (212 >= 0)
    }
    var v15 uint64 = 18446744073709551615
    for fi16 := 0; fi16 < 2; fi16++ {
        v15 = (v15 | v15)
        v3 = (!(!(!v3)))
        v15 += (v15 >> ((179) & 63))
        _ = fi16
    }
    v1 = 0.0
    v2[127] = (-1 <= 127)
    var v17 uint8 = 127
    v18 := GenS1{F0: 2, F1: -128}
    v19 := GenS1{F0: 4294967295, F1: -53}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
