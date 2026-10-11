package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 771
    var v3 uint32 = (v2 ^ v2)
    var v4 uint64 = 1
    var v5 uint32 = v3
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    switch -517 {
    case 522:
        if true {
            v8 := []int{32767, -9223372036854775808, 564}
            v4 = v4
            _ = v8
        } else {
            v9 := []int{}
            v4 -= v4
            var v10 uint8 = 153
            _ = v10
            _ = v9
        }
    default:
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := v1[v3]
            if (365 >= 18446744073709551615) {
                switch (-(v12)) {
                case -2147483648:
                    if (-2147483648 != 1) {
                        var v13 int32 = 42
                        v3 = (^(v2))
                        v14 := v1[2147483648]
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 255, F1: 108, F2: (^(v4))}
                        _ = v15
                    }
                default:
                    _ = len(v1)
                }
                v2 -= (v5 * 1)
                v12 = v12
            }
            v16 := GenS2{F0: 1, F1: 1}
            _ = fi11
            _ = v12
            _ = v16
        }
    }
    fmt.Println(v3, v5)
    v4 -= (v4 / (((v4 >> ((233) & 63))) | 1))
    v2 -= (v5 * 2147483647)
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS2{F0: -1, F1: 54}
        _ = fi17
        _ = v18
    }
    v2 = (v3 | 42)
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int8 = -128
        _ = fi19
        _ = v20
    }
    var v21 uint32 = 1
    v22 := []int{-754, 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
