package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float32 = (0.0 + 0.1)
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 uint64 = 42
    v5 := []string{"foo bar", "foo bar"}
    var v6 float64 = -2.25
    var v7 int = 127
    if (false || (-9223372036854775808 <= 1)) {
        v8 := v3[(^(v7))]
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: 42, F1: 0, F2: -2147483648}
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        if false {
            switch (^(v7)) {
            case -651:
                fmt.Println(v6)
            case 0:
                v4 = (^(v4))
            case 32768:
                if (228 <= 65) {
                    v11 := GenS1{F0: -2147483648, F1: 832, F2: -2147483648}
                    _ = v11
                }
            default:
                if true {
                    delete(v2, (v7 << ((42) & 63)))
                    v12 := GenS1{F0: -599, F1: -1, F2: 739}
                    v13 := GenS2{F0: 2}
                    _ = v12
                    _ = v13
                }
            }
            v1 += v1
            v14 := GenS3{F0: 209, F1: -1}
            _ = v14
        }
    }
    v4 = v4
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            v4 -= 4294967295
            if true {
                v1 = v1
            } else {
                var v17 uint32 = 2147483648
                v18 := GenS1{F0: -1, F1: 507, F2: 596}
                _ = v17
                _ = v18
            }
            if (975 != 18446744073709551615) {
                v19 := GenS1{F0: 0, F1: 42, F2: -2147483648}
                fmt.Println(fi15)
                _ = v19
            }
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println(v4, v7)
    var v20 int32 = -2147483648
    var v21 float64 = (-2.25 - 3.14159)
    v20 = 42
    v22 := GenS4{F0: 255}
    var v23 uint16 = 256
    if false {
        var v24 int32 = v20
        for fi25 := 0; fi25 < 18; fi25++ {
            fmt.Println(v4, v21, fi25)
            var v26 uint = 65535
            switch -477 {
            case 42:
                if true {
                    if ((v1 < ((((v1 * ((v1 / float32(v21)) / 1e10)) + v1) / v1) * -2.25)) && false) {
                        var v27 uint = v26
                        _ = v27
                    }
                    v20 -= v20
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    fmt.Println(v1)
                    fmt.Println(v24, v24)
                }
            case 778:
                v26 = (v26 % ((v26) | 1))
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        if false {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            switch v7 {
            case 0:
                v6 = (v6 + ((v21 * (v6 - (((v6 * v6) * v6) - ((0.1 * float64(v1)) / v6)))) * v6))
            default:
                v21 = (v21 + v21)
            }
            var v28 int32 = (v20 ^ (v20 & 1))
            _ = v28
        } else {
            var v29 uint8 = 128
            _ = v29
        }
    }
    fmt.Println(v4, v1, v23)
    fmt.Println(v23)
    v30 := GenS2{F0: 0}
    v2[1] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
