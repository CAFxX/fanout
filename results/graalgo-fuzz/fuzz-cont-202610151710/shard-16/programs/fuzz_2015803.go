package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int8, p1 int64, p2 int) string {
    var v4 int64 = 9223372036854775807
    var v5 uint8 = 1
    var v6 uint8 = (v5 << ((153) & 63))
    switch p2 {
    case 4294967295:
        if true {
            v7 := []int{(p2 & (p2 * p2)), p2}
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: 561}
                    v11 := GenS1{F0: 98}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                if (v6 == 42) {
                    if ((p0 >> ((18446744073709551615) & 63)) >= p0) {
                        v12 := GenS2{F0: -1, F1: 703}
                        _ = v12
                    }
                }
                _ = fi8
            }
            _ = v7
        }
    case 128:
        v13 := GenS1{F0: 0}
        _ = v13
    case -2147483648:
        var v14 float32 = -2.25
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return ("" + "foo bar")
}

func main() {
    v1 := []uint32{65535, 0, 416}
    v2 := []uint32{645}
    var v3 int = 4294967295
    var v15 uint32 = 85
    v15 = (v15 ^ (^(v15)))
    var v16 int32 = 1
    v16 = (v16 ^ v16)
    var v17 uint8 = 42
    if len(v2) > 0 {
        v2[0] = v15
    }
    var v18 string = "foo bar"
    fmt.Println(v18)
    v3 = v3
    fmt.Println(v15, v16, v17)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
