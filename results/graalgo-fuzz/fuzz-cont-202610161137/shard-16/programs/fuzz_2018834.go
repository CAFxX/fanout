package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 string }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 int; F2 uint32 }
type GenS5 struct { F0 string }
func genFn0(p0 uint, p1 uint64, p2 bool) uint8 {
    fmt.Println(p1, p1)
    v5 := GenS1{F0: ("\n\t" + "foo bar"), F1: (true && (100 != -1)), F2: ""}
    if (42 > -67) {
        if (42 != 834) {
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := GenS1{F0: "", F1: (true || (!p2)), F2: "\n\t"}
                _ = fi6
                _ = v7
            }
        } else {
            p2 = ((((p1 % (((p1 / ((769) | 1))) | 1)) <= p1) && p2) && (0 <= -9223372036854775808))
            v8 := []int{-829, -2147483648, 32767}
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 float64 = ((0.0 + 3.14159) - 1e10)
                if ((fi9 <= fi9) || p2) {
                    v11 := GenS1{F0: "", F1: true, F2: (("\n\t" + "") + "\n\t")}
                    var v12 bool = (p2 || p2)
                    _ = v11
                    _ = v12
                } else {
                    if len(v8) > 0 {
                        v8[0] = (fi9 << (((p0 | 4294967295)) & 63))
                    }
                    var v13 int16 = 0
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        if (-467 != -1) {
            var v14 bool = (((28 > 431) && (!p2)) || (-669 != 538))
            _ = v14
        } else {
            fmt.Println(p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}
func genFn1(p0 uint64, p1 uint64) int16 {
    p1 = (p1 - 42)
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int8 = -128
        var v17 float32 = float32(0.0)
        p0 += 159
        _ = fi15
        _ = v16
        _ = v17
    }
    p1 += 4294967295
    if (false && ((1 >= 969) || true)) {
        var v18 int16 = 32767
        var v19 uint64 = p0
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := []int{42}
            _ = fi20
            _ = v21
        }
        v22 := []int{}
        p1 += p0
        _ = v22
    }
    p1 -= p0
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []int{710, 32768, 2147483647}
    var v2 int32 = 1
    var v3 uint32 = 65535
    var v4 bool = (42 == 18446744073709551615)
    var v23 int16 = 1
    v24 := GenS2{F0: 32767, F1: 32767, F2: (v3 << ((1) & 63))}
    var v25 float64 = 0.0
    v3 = 42
    fmt.Println(v2, v3, v3)
    var v26 uint = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v23 += int16(v26)
    if len(v1) > 0 {
        v1[0] = 42
    }
    v27 := GenS3{F0: 627, F1: 127}
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    var v28 int16 = v23
    if (false || (float64((float32(v25) - (0.0 / ((0.1 / (1.5 / 1.5)) / 3.14159)))) < (0.1 - 0.0))) {
        v29 := GenS4{F0: (v3 * v3), F1: 251, F2: v3}
        v30 := GenS3{F0: v26, F1: -36}
        _ = v29
        _ = v30
    }
    v31 := GenS5{F0: ("" + ("\n\t" + (("a" + "\n\t") + "")))}
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int16 = (v23 * genFn1(18446744073709551615, 2))
        _ = fi32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
