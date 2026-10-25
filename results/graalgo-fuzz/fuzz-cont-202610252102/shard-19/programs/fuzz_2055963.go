package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int; F2 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS8 struct { F0 int32; F1 int16 }

func main() {
    v1 := []int{-166, 2147483647, -2147483648, 32767}
    var v2 uint32 = 65535
    var v3 uint64 = 769
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    var v6 uint16 = 42
    if ((0.0 / 1.5) < 3.14159) {
        v7 := GenS1{F0: 987}
        if true {
            switch 1 {
            case 42:
                v6 = (v6 + (^(v6)))
            case 9223372036854775807:
                if true {
                    v8 := GenS2{F0: 0}
                    v6 = (v6 << ((18446744073709551615) & 63))
                    v9 := GenS2{F0: 0}
                    _ = v8
                    _ = v9
                }
            case 4294967295:
                var v10 int64 = -9223372036854775808
                _ = v10
            default:
                fmt.Println(v2)
            }
        } else {
            fmt.Println(v6)
            var v11 uint32 = (v2 >> ((2) & 63))
            v11 = v11
            _ = v11
        }
        var v12 float32 = (-2.25 - (0.1 * (-2.25 / float32(((1e10 / 1.5) * float64(1.5))))))
        _ = v12
        _ = v7
    } else {
        v2 = 1
        v13 := v5[205]
        _ = v13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 int32 = 42
    if (2 == 255) {
        v2 = v2
        v15 := GenS3{F0: (v14 | v14), F1: 2147483647, F2: -1}
        if true {
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := GenS4{F0: 32767}
                fmt.Println(v14, v14)
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v4[2147483647] = "\n\t"
                    var v20 uint16 = (v6 * 65535)
                    _ = fi19
                    _ = v20
                }
                fmt.Println(v3, fi18, v14)
                _ = fi18
            }
        } else {
            var v21 bool = (!true)
            v2 -= 65535
            v14 += (-(v14))
            _ = v21
        }
        _ = v15
    } else {
        var v22 string = "foo bar"
        v22 = v22
        _ = v22
    }
    if (!(v14 < v14)) {
        v23 := GenS5{F0: 0}
        fmt.Println(v14, v14)
        v6 += (v6 >> ((0) & 63))
        _ = v23
    } else {
        if (false && (v14 <= v14)) {
            v3 = 343
            v6 -= 0
            if ((uint8(v6) == 2) || true) {
                v14 = (^(v14))
            }
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            if len(v1) > 0 {
                v1[0] = int(v6)
            }
            v6 -= 32767
            v25 := GenS6{F0: (fi24 * (fi24 * fi24))}
            _ = fi24
            _ = v25
        }
        v14 = (v14 & (v14 + -1))
    }
    var v26 int64 = 42
    if true {
        var v27 uint32 = 2147483647
        _ = v27
    } else {
        v28 := GenS5{F0: 42}
        v29 := GenS7{F0: 225, F1: -2147483648, F2: (v6 >> ((255) & 63))}
        fmt.Println(v2, v2)
        _ = v28
        _ = v29
    }
    v26 += -663
    v30 := GenS8{F0: (v14 * v14), F1: 32767}
    fmt.Println(v6, v3, v14)
    v2 -= uint32(v6)
    v26 += 226
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
