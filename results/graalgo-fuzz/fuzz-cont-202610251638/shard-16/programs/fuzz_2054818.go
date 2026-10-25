package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int8; F1 int32; F2 bool }
func genFn0(p0 uint8, p1 uint16, p2 uint16) int32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 100}
        _ = fi7
        _ = v8
    }
    p2 = 65535
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 uint32, p1 bool) uint32 {
    switch "hello" {
    case "a":
        var v9 float64 = ((0.0 - 0.0) * float64(0.1))
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS2{F0: 42}
            p0 += p0
            p0 = (p0 & p0)
            _ = fi10
            _ = v11
        }
    case "foo bar":
        var v12 uint16 = 2
        _ = v12
    default:
        var v13 int8 = -1
        _ = v13
    }
    switch 2147483647 {
    case -803:
        v14 := GenS3{F0: 182, F1: 0}
        _ = v14
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 int16 = 0
    var v2 int64 = 42
    var v3 uint64 = 0
    var v4 uint8 = 126
    v5 := []uint32{631, 4294967295}
    var v6 int16 = (v1 | (v1 ^ v1))
    var v15 uint = 2
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint32 = 2
        var v18 string = "foo bar"
        switch "" {
        case "a":
            if (v15 <= v15) {
                if len(v5) > 0 {
                    v5[0] = (v17 * (v17 & (^(v17))))
                }
            }
        case "\n\t":
            v17 = (^(v17))
        case "foo bar":
            var v19 int = (fi16 / ((fi16) | 1))
            _ = v19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    v20 := GenS1{F0: 42}
    v4 = v4
    fmt.Println(v6, v3)
    if false {
        if true {
            var v21 int16 = (v1 % (((v6 << (((v15 ^ (v15 - 458))) & 63))) | 1))
            _ = v21
        }
        v15 = uint(v4)
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(v2, v15)
        _ = fi22
    }
    var v23 int = 76
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    v6 = -864
    if ((((1e10 / 1.0) * 0.1) == float32(3.14159)) && false) {
        for fi24 := 0; fi24 < 14; fi24++ {
            v4 = (^(v4))
            var v25 int64 = 550
            var v26 int64 = -9
            _ = fi24
            _ = v25
            _ = v26
        }
        if (true && (!false)) {
            fmt.Println(v15, v1, v6)
            v4 -= 133
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int32 = -2147483648
                fmt.Println(v2, v1)
                _ = fi27
                _ = v28
            }
        } else {
            v5 = append(v5, 629)
            for fi29 := 0; fi29 < 9; fi29++ {
                fmt.Println(v1)
                switch "a" {
                case "a":
                    for fi30 := 0; fi30 < 20; fi30++ {
                        var v31 int64 = -9223372036854775808
                        switch fi29 {
                        case 128:
                            if (0 >= 4294967295) {
                                var v32 uint64 = (v3 ^ 398)
                                switch -1 {
                                case -363:
                                    if ((((-2.25 < (0.0 + 0.0)) || (-14 < 127)) || false) && ((true && false) && false)) {
                                        var v33 uint64 = (v32 << ((uint(v4)) & 63))
                                        _ = v33
                                    }
                                case -2147483648:
                                    v5 = append(v5, 381)
                                case 32768:
                                    fmt.Println(fi29, v32)
                                default:
                                    fmt.Println(fi30)
                                }
                                _ = v32
                            } else {
                                v34 := GenS2{F0: (^(v1))}
                                v15 = (v15 % ((v15) | 1))
                                var v35 uint = v15
                                _ = v34
                                _ = v35
                            }
                        }
                        _ = fi30
                        _ = v31
                    }
                }
                v4 += (^(v4))
                _ = fi29
            }
            v1 -= (^(v6))
        }
        var v36 uint8 = 53
        _ = v36
    }
    v5 = append(v5, 353)
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    v37 := GenS2{F0: v1}
    switch v23 {
    case 4294967295:
        v38 := GenS3{F0: 18446744073709551615, F1: (v3 + 18446744073709551615)}
        _ = v38
    case 0:
        v39 := GenS4{F0: -1, F1: -1, F2: ((float32((3.14159 - 1.5)) * float32(3.14159)) <= 0.1)}
        _ = v39
    }
    fmt.Println(v1, v23, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
