package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 uint8 }
func genFn0() uint8 {
    var v5 int16 = -1
    for fi6 := 0; fi6 < 13; fi6++ {
        v5 = v5
        if (1.5 == float32(3.14159)) {
            fmt.Println(v5, fi6)
        }
        _ = fi6
    }
    if (42 >= 18446744073709551615) {
        var v7 uint64 = 2
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint16 = 42
            var v10 int32 = 424
            fmt.Println(fi8, v7)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    _ = v5
    return 187
}
func genFn1(p0 int16, p1 string, p2 uint32) string {
    var v11 int16 = (^(p0))
    v12 := GenS1{F0: 9223372036854775807, F1: 951, F2: -2147483648}
    v13 := GenS1{F0: 32768, F1: 982, F2: int32(p2)}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return ""
}

func main() {
    var v1 uint32 = 1
    var v2 int8 = 54
    var v3 float64 = 0.0
    var v4 bool = (v3 >= v3)
    switch ("a" + ("a" + "")) {
    case "hello":
        if v4 {
            if v4 {
                var v14 uint16 = 65535
                _ = v14
            }
            v15 := GenS1{F0: 864, F1: -296, F2: 0}
            _ = v15
        }
    case "foo bar":
        if (!((0.1 / (3.14159 * float32((1e10 - v3)))) != (((-2.25 * (0.0 - -2.25)) / 1.0) / ((1.5 / 0.1) + 3.14159)))) {
            v1 = (v1 | (v1 & v1))
            v2 = v2
            for fi16 := 0; fi16 < 12; fi16++ {
                v17 := []int{9223372036854775807, fi16, (fi16 % (((^(fi16))) | 1))}
                v4 = (1 == 42)
                var v18 bool = v4
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            var v19 uint = 92
            var v20 int32 = 108
            v21 := GenS1{F0: 128, F1: 0, F2: v20}
            _ = v19
            _ = v20
            _ = v21
        }
    }
    fmt.Println(v1, v1, v1)
    var v22 int64 = -643
    v23 := []int{1, -2147483648, -446}
    v22 -= v22
    switch 32767 {
    case 824:
        for fi24 := 0; fi24 < 1; fi24++ {
            fmt.Println(v3, v4)
            fmt.Println(v3)
            _ = fi24
        }
    case -9223372036854775808:
        if v4 {
            switch (("a" + "hello") + "a") {
            case "hello":
                for fi25 := 0; fi25 < 14; fi25++ {
                    if (v4 && ((!("hello" < "a")) || v4)) {
                        v26 := GenS1{F0: fi25, F1: 9223372036854775807, F2: 42}
                        v23 = append(v23, (^(fi25)))
                        if len(v23) > 0 {
                            v23[0] = 4294967295
                        }
                        _ = v26
                    }
                    _ = fi25
                }
            default:
                var v27 uint8 = uint8(v1)
                _ = v27
            }
            if ((127 > 2) || (((v1 >> ((18446744073709551615) & 63)) != (v1 * 1)) && v4)) {
                v2 = v2
                var v28 int32 = 549
                _ = v28
            }
        }
    case 9223372036854775807:
        var v29 uint8 = 248
        _ = v29
    }
    if len(v23) > 0 {
        v23[0] = 4294967295
    }
    v30 := GenS2{F0: 242}
    if len(v23) > 0 {
        v23[0] = 32768
    }
    if (462 < 2) {
        fmt.Println(v3)
        if (v22 <= (v22 ^ -687)) {
            _ = len(v23)
            fmt.Println(v2)
            fmt.Println(v22, v2, v4)
        }
        if len(v23) > 0 {
            v23[0] = 128
        }
    } else {
        v3 = (0.1 / (float64(0.1) / (((((1.5 * -2.25) * (v3 / v3)) + (v3 + -2.25)) + (float64((0.0 + (1.5 - 1.5))) / 1e10)) + v3)))
        var v31 int32 = 42
        _ = v31
    }
    v4 = v4
    v4 = v4
    v4 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
}
