package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }

func main() {
    var v1 int64 = -1
    var v2 int8 = 100
    var v3 int = 32767
    v4 := []uint32{65535, 2147483647}
    var v5 int8 = (v2 % (((-(v2))) | 1))
    v6 := []string{"\n\t", ("foo bar" + "")}
    v2 = (^(v5))
    fmt.Println(v5, v2, v2)
    v5 = 1
    var v7 int16 = 1000
    var v8 uint32 = 2147483647
    v5 = (^(v5))
    if true {
        var v9 uint8 = 42
        v7 = v7
        _ = v9
    } else {
        switch (v3 / ((v3) | 1)) {
        case 952:
            if (231 <= 4294967295) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                switch v3 {
                case -95:
                    _ = len(v6)
                case -794:
                    if (0.1 > 0.0) {
                        if len(v4) > 0 {
                            v4[0] = (v8 | (v8 - 382))
                        }
                        v10 := GenS1{F0: (^(v7))}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v10
                    }
                case 42:
                    v1 = v1
                }
                var v11 uint16 = 2
                _ = v11
            }
        case -2147483648:
            v12 := GenS2{F0: false, F1: (!true)}
            _ = v12
        default:
            v2 += (v5 & (v2 ^ v2))
        }
        switch (v3 + (v3 - v3)) {
        case 32768:
            if true {
                v13 := GenS3{F0: -2147483648, F1: (v3 - v3), F2: 42}
                var v14 int16 = -678
                if len(v4) > 0 {
                    v4[0] = (v8 * 877)
                }
                _ = v13
                _ = v14
            } else {
                v4 = append(v4, v8)
            }
        case -9223372036854775808:
            for fi15 := 0; fi15 < 8; fi15++ {
                if true {
                    var v16 string = "\n\t"
                    _ = v16
                } else {
                    var v17 uint8 = 42
                    _ = v17
                }
                var v18 float64 = 1e10
                fmt.Println(v2, v5)
                _ = fi15
                _ = v18
            }
        default:
            if false {
                var v19 int = (v3 & (-(v3)))
                var v20 int32 = 274
                v21 := GenS3{F0: (v20 - 2147483647), F1: v19, F2: 255}
                _ = v19
                _ = v20
                _ = v21
            } else {
                var v22 int8 = (v2 | 0)
                _ = v22
            }
        }
        if true {
            v23 := GenS1{F0: (v7 ^ v7)}
            v24 := GenS3{F0: 42, F1: v3, F2: 0}
            v25 := GenS3{F0: 29, F1: v3, F2: 65535}
            _ = v23
            _ = v24
            _ = v25
        }
    }
    var v26 int8 = 100
    v27 := GenS3{F0: int32(v2), F1: 32767, F2: 603}
    v6 = append(v6, "a")
    switch (^(v3)) {
    case -63:
        var v28 bool = (!(846 > 0))
        _ = v28
    case -580:
        v4 = append(v4, (v8 % ((v8) | 1)))
    default:
        v8 += v8
    }
    fmt.Println(v5, v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
