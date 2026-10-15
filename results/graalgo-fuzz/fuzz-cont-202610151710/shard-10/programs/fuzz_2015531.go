package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int }
type GenS3 struct { F0 int16; F1 string; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 int8, p1 bool, p2 int32) string {
    p1 = ((4294967295 <= 671) && (p2 < -2147483648))
    v6 := GenS1{F0: 32767, F1: 4294967295, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}

func main() {
    v1 := []string{"hello"}
    v2 := []int{-701, 42, -476}
    var v3 int16 = -246
    var v4 int8 = 127
    v5 := make(map[int]bool)
    v7 := GenS1{F0: 32767, F1: 65535, F2: 4294967295}
    fmt.Println(v4, v3)
    var v8 uint64 = 4294967295
    v8 = (^(v8))
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int = (fi9 >> ((2) & 63))
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 4294967295, F1: 204, F2: -996}
    var v12 uint8 = 61
    if ((v12 << ((65535) & 63)) == (v12 / (((v12 * 65)) | 1))) {
        switch 32768 {
        case -70:
            fmt.Println(v12, v3, v3)
        case -1:
            v13 := GenS3{F0: v3, F1: ("foo bar" + "\n\t"), F2: 427}
            _ = v13
        case 32768:
            for fi14 := 0; fi14 < 5; fi14++ {
                for fi15 := 0; fi15 < 16; fi15++ {
                    v3 = v3
                    fmt.Println(v3)
                    switch fi14 {
                    case 553:
                        var v16 int32 = 42
                        _ = v16
                    case -9223372036854775808:
                        if true {
                            v3 -= -950
                            v17 := GenS4{F0: 32767}
                            _ = v17
                        }
                    case 2147483647:
                        var v18 int32 = 766
                        _ = v18
                    default:
                        v19 := GenS5{F0: 0, F1: 42, F2: -1}
                        _ = v19
                    }
                    var v20 bool = (1.5 > 1e10)
                    _ = fi15
                    _ = v20
                }
                _ = fi14
            }
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            v3 = v3
            if len(v2) > 0 {
                v2[0] = fi21
            }
            _ = fi21
        }
    }
    v4 = (v4 % ((v4) | 1))
    fmt.Println(v4)
    var v22 int16 = (v3 & (v3 * v3))
    var v23 uint64 = (v8 ^ v8)
    var v24 uint32 = 42
    var v25 float64 = (1e10 - float64(0.1))
    var v26 uint16 = 42
    var v27 int = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
