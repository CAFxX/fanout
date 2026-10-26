package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 int16) bool {
    var v4 int16 = -262
    var v5 int = -1
    v6 := []int{-2147483648, (-(v5)), (v5 >> ((255) & 63)), (v5 ^ (v5 << ((443) & 63)))}
    var v7 int64 = -9223372036854775808
    v4 += (p0 << ((0) & 63))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (false || (0 > 42))
}
func genFn1() int16 {
    var v8 uint16 = 351
    if (!(0 == v8)) {
        fmt.Println(v8)
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                v8 = 65535
                v11 := GenS1{F0: -128, F1: -128, F2: 4294967295}
                switch fi10 {
                case -9223372036854775808:
                    switch fi9 {
                    case 127:
                        if true {
                            var v12 uint = 4294967295
                            v12 = 746
                            v12 -= v12
                            var v13 int = fi9
                            fmt.Println(v12, v12, v8)
                            _ = v12
                            _ = v13
                        }
                    }
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        if true {
            var v14 int16 = 1
            v15 := []int{63, -2147483648, 1}
            _ = v14
            _ = v15
        } else {
            v8 = (v8 % ((833) | 1))
        }
    }
    v16 := []int{217, -809, 2147483647}
    _ = v16
    _ = v8
    return 912
}

func main() {
    v1 := []uint32{}
    var v2 int32 = 0
    var v3 uint = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 uint = (v3 << (((v3 / (((^(v3))) | 1))) & 63))
    var v18 uint = v17
    v2 -= (v2 % ((v2) | 1))
    v3 = (v17 ^ v18)
    v3 = 65535
    fmt.Println(v18, v2, v18)
    _ = len(v1)
    v18 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
