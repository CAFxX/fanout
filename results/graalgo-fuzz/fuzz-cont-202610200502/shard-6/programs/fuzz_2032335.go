package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int16; F2 int8 }
func genFn0(p0 uint, p1 uint64, p2 string) int32 {
    v6 := []int{}
    p0 += (p0 | (p0 << (((p0 * (^(p0)))) & 63)))
    p1 += (p1 << ((p0) & 63))
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 634
}

func main() {
    var v1 int8 = 100
    var v2 uint64 = 314
    var v3 string = "a"
    var v4 bool = true
    v5 := []uint32{4294967295, 2, 506, 1}
    if true {
        v7 := GenS1{F0: 255, F1: 65535, F2: genFn0(1, (v2 & (v2 << ((1) & 63))), v3)}
        switch int(v2) {
        case 168:
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: 42, F1: 125, F2: 42}
                var v10 uint32 = 2147483647
                _ = fi8
                _ = v10
                _ = v9
            }
        case 2147483647:
            var v11 string = v3
            _ = v11
        }
        v12 := GenS2{F0: -2147483648}
        _ = v12
        _ = v7
    }
    var v13 int32 = -117
    v3 = "a"
    var v14 bool = v4
    var v15 int64 = 327
    var v16 uint8 = 1
    if v4 {
        v3 = v3
        if false {
            var v17 string = ("" + v3)
            var v18 bool = false
            _ = v17
            _ = v18
        }
        switch -9223372036854775808 {
        case 0:
            for fi19 := 0; fi19 < 18; fi19++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v5 = append(v5, uint32(v2))
                fmt.Println(v14)
                _ = fi19
            }
        case 9223372036854775807:
            v13 += int32(v2)
        case 946:
            switch 127 {
            case 0:
                v16 += v16
            case -258:
                v15 += (v15 % ((v15) | 1))
            case 466:
                v5 = append(v5, 2147483648)
            }
        }
    } else {
        var v20 bool = (1 == 1)
        for fi21 := 0; fi21 < 20; fi21++ {
            for fi22 := 0; fi22 < 16; fi22++ {
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v23 := GenS1{F0: uint16(v2), F1: 65535, F2: -2147483648}
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        v15 = v15
        _ = v20
    }
    v24 := GenS3{F0: -146, F1: 1000, F2: v1}
    if len(v5) > 0 {
        v5[0] = 870
    }
    v25 := GenS3{F0: int16(v1), F1: 1000, F2: 75}
    var v26 int32 = v13
    v4 = (((1.5 / 0.1) > ((1.5 / 0.1) + 1e10)) && (v14 || v4))
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
