package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool) int32 {
    v7 := []int{4294967295, 1, 128, 32767}
    var v8 bool = p0
    if ((255 == 683) || false) {
        v9 := GenS1{F0: 2147483647}
        _ = v9
    }
    fmt.Println(p0, p0)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    _ = v8
    return 16
}
func genFn1() uint64 {
    var v10 float64 = -2.25
    v11 := []int{42}
    var v12 uint32 = 2147483647
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint = 2
        v12 -= (^(v12))
        _ = fi13
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 uint16 = 42
    v2 := []uint32{42, 181}
    var v3 uint = 2
    var v4 uint32 = 0
    v5 := make(map[string]string)
    var v6 float64 = 1.5
    var v15 int16 = -1
    var v16 int = -9223372036854775808
    var v17 uint32 = 0
    var v18 uint = (v3 * v3)
    v19 := GenS1{F0: uint32(v18)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 128 {
    case 127:
        fmt.Println(v18)
    case 2147483647:
        var v20 uint = v18
        _ = v20
    case 4294967295:
        var v21 uint8 = 0
        _ = v21
    default:
        v3 = (v18 & (v3 - 632))
    }
    if true {
        if (!(0 != -118)) {
            var v22 int = (v16 << ((588) & 63))
            v15 += (-(v15))
            switch (v16 + 184) {
            case -46:
                v23 := GenS2{F0: v18}
                _ = v23
            case -2147483648:
                for fi24 := 0; fi24 < 12; fi24++ {
                    v25 := GenS1{F0: 2}
                    v18 -= (v3 % ((561) | 1))
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 uint32 = v4
                        var v28 uint = 4294967295
                        var v29 float64 = v6
                        _ = len(v2)
                        _ = len(v2)
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = fi24
                    _ = v25
                }
            }
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
