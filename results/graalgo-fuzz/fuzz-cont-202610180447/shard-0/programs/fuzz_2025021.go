package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint16; F2 int }
func genFn0(p0 uint8) int16 {
    p0 = 94
    if (1.5 <= 1e10) {
        if (2 >= 18446744073709551615) {
            v7 := []int{-1, -247}
            if (128 > p0) {
                var v8 uint32 = 42
                if false {
                    v8 = 306
                }
                _ = v8
            }
            var v9 uint8 = 255
            _ = v7
            _ = v9
        }
        if (2 < 18446744073709551615) {
            fmt.Println(p0)
            for fi10 := 0; fi10 < 8; fi10++ {
                p0 -= p0
                _ = fi10
            }
        }
        if (!(!(!false))) {
            var v11 float64 = 0.1
            if true {
                v11 -= (0.1 - ((((0.1 / v11) - float64((3.14159 * 3.14159))) + float64(float32(1e10))) * v11))
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := GenS1{F0: 255, F1: 59, F2: (fi12 ^ fi12)}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    }
    v14 := []int{-597}
    _ = p0
    _ = v14
    return 855
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 32767
    var v3 int32 = 517
    v4 := make(map[int]uint32)
    var v5 uint32 = 212
    v6 := make(map[int]uint32)
    v3 = v3
    var v15 uint32 = v5
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v2, v5)
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    v2 -= (^(v2))
    v18 := []int{-9223372036854775808, -1, 127, 9223372036854775807}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS1{F0: v2, F1: 4294967295, F2: 32767}
        _ = fi19
        _ = v20
    }
    fmt.Println(v2, v3, v2)
    var v21 uint8 = 225
    var v22 uint = 4294967295
    switch 2147483647 {
    case 32767:
        if false {
            var v23 float64 = (float64(((-2.25 - 3.14159) + 1e10)) / (3.14159 + 1e10))
            v22 = 935
            _ = v23
        } else {
            if true {
                for fi24 := 0; fi24 < 17; fi24++ {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        v21 += 0
                        delete(v6, fi25)
                        _ = fi25
                    }
                    v18 = append(v18, (fi24 | (fi24 >> (((^(v22))) & 63))))
                    if true {
                        v26 := GenS2{F0: 0}
                        _ = v26
                    } else {
                        if len(v18) > 0 {
                            v18[0] = fi24
                        }
                    }
                    _ = fi24
                }
                var v27 int64 = 42
                _ = v27
            }
            fmt.Println(v22)
        }
    case 729:
        var v28 int = 533
        _ = v28
    default:
        v29 := GenS2{F0: int64(v3)}
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        v31 := GenS1{F0: v2, F1: (v22 >> ((v22) & 63)), F2: fi30}
        switch "\n\t" {
        case "foo bar":
            fmt.Println(fi30, fi30)
        case "\n\t":
            if ((v2 * v2) <= (^(v2))) {
                v32 := GenS2{F0: -1}
                v15 -= (v15 + 2)
                _ = v32
            } else {
                v15 = (v15 % (((v15 << ((1) & 63))) | 1))
                for fi33 := 0; fi33 < 2; fi33++ {
                    switch 9223372036854775807 {
                    case 32767:
                        v2 = (v2 % (((v2 + uint16(v15))) | 1))
                    case 9223372036854775807:
                        v34 := GenS3{F0: v3, F1: 65535, F2: fi30}
                        _ = v34
                    case 32768:
                        delete(v1, ("a" + "foo bar"))
                    }
                    _ = fi33
                }
            }
        default:
            fmt.Println(v21, v22, v2)
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
