package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 float32 = 1.5
    var v2 uint16 = 32767
    v3 := make(map[uint32]bool)
    v4 := []int{2147483647}
    var v5 int16 = -32768
    var v6 bool = false
    fmt.Println(v5, v6)
    v7 := GenS1{F0: 2, F1: 0, F2: 125}
    v1 -= ((float32(0.1) - (v1 - (0.0 / v1))) + v1)
    var v8 uint8 = 0
    v9 := GenS1{F0: 65535, F1: 2147483647, F2: 18446744073709551615}
    if (v6 || (1 <= 4294967295)) {
        v3[4294967295] = v6
        var v10 int32 = 1
        _ = v10
    }
    if true {
        var v11 uint32 = 229
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 bool = (v1 != v1)
        switch (fi12 + (fi12 % (((fi12 ^ (fi12 & (fi12 * (fi12 >> ((2) & 63)))))) | 1))) {
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 2147483647:
            v13 = false
        case 374:
            if false {
                if (-858 < 9223372036854775807) {
                    if v13 {
                        v14 := GenS2{F0: 42, F1: 1}
                        _ = v14
                    }
                } else {
                    fmt.Println(v13, fi12)
                    for fi15 := 0; fi15 < 15; fi15++ {
                        delete(v3, 65535)
                        var v16 int64 = 332
                        var v17 int = fi12
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
            } else {
                v6 = v13
            }
        default:
            switch 1 {
            case 109:
                var v18 float64 = 1.5
                _ = v18
            default:
                var v19 string = ("\n\t" + ("" + ""))
                _ = v19
            }
        }
        var v20 float32 = (float32(1.5) - v1)
        _ = fi12
        _ = v13
        _ = v20
    }
    v4 = append(v4, -2147483648)
    var v21 int = -1
    v22 := v3[42]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
