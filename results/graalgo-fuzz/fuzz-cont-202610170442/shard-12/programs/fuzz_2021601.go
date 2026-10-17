package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint16 }

func main() {
    v1 := []int{4294967295, 127, -264}
    var v2 string = ""
    var v3 int64 = -1
    var v4 uint8 = 128
    var v5 string = "\n\t"
    var v6 bool = (!false)
    var v7 int8 = 41
    v7 = v7
    v4 = 127
    var v8 uint64 = 914
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint8 = v4
        var v11 uint = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    switch v2 {
    case "hello":
        v12 := GenS1{F0: v8, F1: (v4 & v4), F2: v4}
        _ = v12
    case "\n\t":
        var v13 uint8 = v4
        _ = v13
    default:
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(v3)
            for fi15 := 0; fi15 < 9; fi15++ {
                fmt.Println(v5, fi14, fi15)
                var v16 float32 = (0.1 / 3.14159)
                v8 -= (v8 + (^(v8)))
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    if (-228 != 1) {
        if v6 {
            if (128 < 127) {
                var v17 bool = v6
                _ = v17
            }
            var v18 uint16 = 805
            v19 := GenS1{F0: v8, F1: uint8(v8), F2: 127}
            _ = v18
            _ = v19
        } else {
            if len(v1) > 0 {
                v1[0] = 32767
            }
            if v6 {
                v1 = append(v1, -1)
            }
            v4 -= (^(v4))
        }
        v3 = (v3 ^ (v3 % (((v3 + v3)) | 1)))
    } else {
        fmt.Println(v2)
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint32 = 363
                _ = fi21
                _ = v22
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                v24 := GenS1{F0: v8, F1: (v4 - (v4 ^ (v4 ^ 2))), F2: v4}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 8; fi25++ {
                v1 = append(v1, (-(fi25)))
                _ = fi25
            }
            _ = fi20
        }
    }
    var v26 int = -520
    if v6 {
        _ = len(v1)
        switch -2147483648 {
        case -1:
            v27 := GenS1{F0: v8, F1: v4, F2: (v4 % (((v4 & v4)) | 1))}
            _ = v27
        case 2147483647:
            switch -9223372036854775808 {
            case 32768:
                var v28 uint16 = 65535
                _ = v28
            case 979:
                v7 = (v7 + 42)
            default:
                var v29 int32 = -35
                _ = v29
            }
        default:
            for fi30 := 0; fi30 < 14; fi30++ {
                for fi31 := 0; fi31 < 11; fi31++ {
                    v26 -= v26
                    _ = fi31
                }
                fmt.Println(v2, v8)
                for fi32 := 0; fi32 < 7; fi32++ {
                    if (42 > 42) {
                        v33 := GenS1{F0: 18446744073709551615, F1: v4, F2: (v4 * uint8(v8))}
                        v34 := GenS2{F0: (v3 / (((v3 - int64(fi30))) | 1)), F1: 2, F2: 582}
                        _ = v33
                        _ = v34
                    }
                    _ = fi32
                }
                _ = fi30
            }
        }
        var v35 float32 = 1.5
        _ = v35
    } else {
        var v36 int16 = -260
        var v37 string = v2
        v4 = v4
        _ = v36
        _ = v37
    }
    v26 += v26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
