package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 42
    var v3 uint16 = 32767
    fmt.Println(v1)
    var v4 int8 = -84
    v3 += (^(v3))
    switch 1 {
    case 9223372036854775807:
        v5 := GenS1{F0: "\n\t"}
        _ = v5
    case -153:
        v6 := []int{128, 32768, -987}
        _ = v6
    default:
        for fi7 := 0; fi7 < 18; fi7++ {
            v1 = (v1 & (v1 / ((v1) | 1)))
            v8 := []int{fi7, fi7, int(v1)}
            _ = fi7
            _ = v8
        }
    }
    var v9 int32 = 341
    if false {
        var v10 int = 32767
        v9 = (v9 % ((42) | 1))
        v11 := GenS2{F0: "foo bar", F1: v10}
        _ = v10
        _ = v11
    }
    v12 := []int{32767, -2147483648, 181}
    v2 = v2
    v3 = uint16(v1)
    v13 := GenS3{F0: (^(v4)), F1: 128}
    v1 = (v1 + 1)
    if (!(-879 == -1)) {
        switch 1 {
        case 1:
            if (v3 == (v3 ^ 42)) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 int8 = (v4 % ((12) | 1))
                        if ((((^(v2)) > 691) || false) || false) {
                            if true {
                                var v17 int16 = -877
                                v2 = v2
                                switch 4294967295 {
                                case -9223372036854775808:
                                    var v18 int = (fi14 ^ fi15)
                                    _ = v18
                                }
                                v16 -= v4
                                _ = v17
                            } else {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                fmt.Println(v9, v3)
                                v4 = 90
                            }
                            v12 = append(v12, fi14)
                            v1 -= v1
                        }
                        var v19 int = fi14
                        _ = fi15
                        _ = v16
                        _ = v19
                    }
                    switch (fi14 * (fi14 ^ int(v3))) {
                    case 398:
                        fmt.Println(v4, v1, v3)
                    case -1:
                        v3 -= (v3 ^ v3)
                    default:
                        _ = len(v12)
                    }
                    _ = fi14
                }
                var v20 int = -2147483648
                switch v20 {
                case 2147483647:
                    if len(v12) > 0 {
                        v12[0] = (-(v20))
                    }
                case 1:
                    v21 := GenS4{F0: 2147483647, F1: v4}
                    _ = v21
                }
                _ = v20
            } else {
                v2 += v2
                if false {
                    var v22 int32 = (v9 - v9)
                    if (!((v3 + 32767) < v3)) {
                        var v23 float32 = 3.14159
                        if (v23 > 0.0) {
                            v4 -= (^(v4))
                            v22 -= v22
                        }
                        _ = v23
                    } else {
                        v24 := GenS2{F0: ("" + "foo bar"), F1: 127}
                        fmt.Println(v4, v9)
                        _ = v24
                    }
                    _ = v22
                }
                switch "\n\t" {
                case "a":
                    if ((-32768 > 999) || true) {
                        v25 := GenS4{F0: -245, F1: (v4 / ((-1) | 1))}
                        _ = len(v12)
                        fmt.Println(v1, v4)
                        _ = v25
                    }
                case "\n\t":
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v27 := GenS2{F0: "", F1: int(v2)}
                        _ = fi26
                        _ = v27
                    }
                }
            }
        case 2147483647:
            var v28 uint32 = (v2 + v2)
            _ = v28
        default:
            v2 = (^(v2))
        }
        switch 0 {
        case 624:
            v9 += 463
        }
        v2 -= (v2 / ((v2) | 1))
    } else {
        var v29 uint64 = v1
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint = 18446744073709551615
            for fi32 := 0; fi32 < 14; fi32++ {
                v9 = int32(v3)
                fmt.Println(v1, v2, v1)
                var v33 int32 = v9
                _ = len(v12)
                _ = fi32
                _ = v33
            }
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
