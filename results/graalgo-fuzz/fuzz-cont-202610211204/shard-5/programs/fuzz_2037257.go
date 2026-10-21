package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int64; F1 int16; F2 uint32 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 bool; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1e10
    v3 := []int{359}
    if (!(!true)) {
        v4 := v1[2147483647]
        fmt.Println(v2, v2)
        _ = v4
    } else {
        if (!false) {
            v2 -= v2
        }
        for fi5 := 0; fi5 < 4; fi5++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v6 := GenS1{F0: -1}
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            v3 = append(v3, 175)
            for fi8 := 0; fi8 < 10; fi8++ {
                if (0.1 <= -2.25) {
                    if (!false) {
                        v9 := GenS2{F0: 486, F1: 30}
                        v2 = 1.5
                        _ = v9
                    }
                    v2 = v2
                }
                if ((31 <= 1) && true) {
                    if (32767 >= 710) {
                        _ = len(v1)
                        v2 = 0.0
                        _ = len(v3)
                    }
                    var v10 uint8 = 0
                    var v11 float64 = float64(v2)
                    _ = v10
                    _ = v11
                } else {
                    v3 = append(v3, -2147483648)
                }
                _ = fi8
            }
            _ = fi7
        }
    }
    v12 := GenS2{F0: 2147483647, F1: 2}
    if false {
        v2 += float32(0.1)
        var v13 uint32 = 42
        v3 = append(v3, -298)
        _ = v13
    } else {
        v2 -= ((v2 + (1.5 * v2)) / (v2 / v2))
        v2 = (v2 * 3.14159)
        switch -744 {
        case 387:
            v14 := GenS3{F0: -1, F1: 32767, F2: 687}
            _ = v14
        case 128:
            switch 424 {
            case 9223372036854775807:
                v2 += (v2 + 0.0)
            case 2147483647:
                if ((1 <= -1) || true) {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        if false {
                            for fi16 := 0; fi16 < 14; fi16++ {
                                var v17 float32 = 0.1
                                var v18 int64 = 9223372036854775807
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS4{F0: "a", F1: 9223372036854775807}
                            v2 = v2
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v2 = v2
                            var v21 int64 = -1
                            v2 += ((((v2 + float32(((3.14159 * 3.14159) / 0.1))) + 3.14159) - ((v2 * (v2 + (((0.0 * (v2 / v2)) / 3.14159) / v2))) / v2)) / 1e10)
                            _ = fi20
                            _ = v21
                        }
                        if (!false) {
                            v2 = 0.1
                            if (fi15 == (fi15 * 128)) {
                                v22 := v1[fi15]
                                v23 := GenS3{F0: 0, F1: -1, F2: 413}
                                _ = v22
                                _ = v23
                            }
                        }
                        _ = fi15
                    }
                    if false {
                        if (!(32767 >= -32768)) {
                            var v24 int16 = 928
                            fmt.Println(v2, v2)
                            _ = v24
                        }
                    }
                }
            default:
                if (128 != 0) {
                    v25 := GenS5{F0: false, F1: 303, F2: 1}
                    _ = v25
                }
            }
        default:
            for fi26 := 0; fi26 < 15; fi26++ {
                fmt.Println(v2, v2)
                _ = fi26
            }
        }
    }
    var v27 float32 = v2
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS5{F0: true, F1: 42, F2: 2}
        fmt.Println(v2)
        _ = fi28
        _ = v29
    }
    switch ("foo bar" + "") {
    case "\n\t":
        delete(v1, 32768)
    }
    var v30 uint16 = 917
    var v31 uint32 = 65535
    if true {
        v32 := v1[128]
        _ = v32
    } else {
        if len(v3) > 0 {
            v3[0] = -946
        }
        switch 2147483647 {
        case -9223372036854775808:
            v33 := GenS6{F0: -11}
            _ = v33
        }
    }
    _ = len(v3)
    var v34 int64 = -9223372036854775808
    v35 := GenS6{F0: 973}
    v36 := GenS7{F0: -501, F1: ((((false && true) && ((v34 % ((0) | 1)) <= (v34 | (v34 + v34)))) && ("" < "hello")) && true), F2: (v34 - 553)}
    v34 = v34
    v37 := GenS1{F0: -1}
    var v38 float64 = 3.14159
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
