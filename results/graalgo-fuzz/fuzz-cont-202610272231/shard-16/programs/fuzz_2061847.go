package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
func genFn0(p0 uint8) uint64 {
    v4 := GenS1{F0: 599}
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{0, fi5}
        var v7 int = (fi5 % (((fi5 >> ((18446744073709551615) & 63))) | 1))
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = 42
    fmt.Println(v8, p0)
    _ = p0
    _ = v4
    _ = v8
    return 846
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = 42
    v3 := make(map[uint32]uint32)
    switch -1 {
    case 32767:
        for fi9 := 0; fi9 < 13; fi9++ {
            if (((-2.25 - 0.0) / 3.14159) >= 1.5) {
                switch int(v2) {
                case 42:
                    v2 -= v2
                case -222:
                    v1 -= v1
                default:
                    v10 := GenS2{F0: v1, F1: (v1 / (((^(v1))) | 1)), F2: 4294967295}
                    _ = v10
                }
            } else {
                switch (fi9 - (-(fi9))) {
                case 799:
                    fmt.Println(v2, fi9, v2)
                case -628:
                    var v11 uint64 = genFn0(255)
                    _ = v11
                default:
                    v12 := GenS2{F0: (v1 * v1), F1: v1, F2: 4294967295}
                    _ = v12
                }
            }
            _ = fi9
        }
    case 4294967295:
        v1 -= (v1 + (v1 << ((1) & 63)))
    default:
        v2 -= v2
    }
    _ = len(v3)
    v2 -= 127
    delete(v3, 4294967295)
    var v13 uint8 = 76
    var v14 int16 = -1
    v1 = (^(v1))
    v15 := []int{}
    v15 = append(v15, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
