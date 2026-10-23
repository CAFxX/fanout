package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint16) int8 {
    var v6 uint64 = 521
    p0 += (p0 - (p0 << ((721) & 63)))
    var v7 uint32 = 2
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint64) uint64 {
    v8 := GenS1{F0: 514, F1: 42, F2: 902}
    var v9 uint = 65535
    if (!false) {
        p0 += (p0 >> ((v9) & 63))
        var v10 int64 = -575
        _ = v10
    } else {
        p0 = uint64(v9)
        switch 128 {
        case 9223372036854775807:
            fmt.Println(v9, v9)
        default:
            if true {
                v9 = (v9 | 2)
                var v11 uint = v9
                if false {
                    v12 := GenS1{F0: 2147483647, F1: uint8(p0), F2: 776}
                    if (!true) {
                        p0 = (p0 * p0)
                        v9 = 0
                        v13 := GenS2{F0: -51, F1: -1}
                        _ = v13
                    }
                    _ = v12
                }
                _ = v11
            }
        }
    }
    var v14 int16 = 396
    _ = p0
    _ = v14
    _ = v8
    _ = v9
    return (p0 / (((p0 << ((v9) & 63))) | 1))
}

func main() {
    var v1 uint = 285
    var v2 int8 = -1
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 207
    fmt.Println(v3, v1, v2)
    v15 := v4[2147483648]
    fmt.Println(v1, v1, v2)
    if ((1 == 167) && ((v2 - 6) < v2)) {
        v16 := []int{32767, 42, -9223372036854775808, 9223372036854775807}
        _ = v16
    }
    v5 -= (v5 / ((506) | 1))
    if true {
        v17 := []int{42, 682, 504}
        for fi18 := 0; fi18 < 13; fi18++ {
            v2 = v2
            var v19 int8 = v2
            _ = fi18
            _ = v19
        }
        v20 := GenS2{F0: 1, F1: 42}
        _ = v17
        _ = v20
    } else {
        v21 := GenS2{F0: 549, F1: -468}
        switch 9223372036854775807 {
        case 127:
            var v22 int32 = 42
            _ = v22
        case 2147483647:
            var v23 string = "\n\t"
            _ = v23
        default:
            if ((-(v2)) > v2) {
                v3 = false
                switch "hello" {
                case "":
                    v24 := GenS3{F0: 65535}
                    _ = v24
                default:
                    var v25 int8 = 0
                    _ = v25
                }
            }
        }
        v2 = (v2 / (((v2 & (v2 >> (((v1 | v1)) & 63)))) | 1))
        _ = v21
    }
    v26 := []int{-999, 0, -9223372036854775808}
    v15 = (v5 / ((v5) | 1))
    var v27 int8 = genFn0(v2, 451)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
