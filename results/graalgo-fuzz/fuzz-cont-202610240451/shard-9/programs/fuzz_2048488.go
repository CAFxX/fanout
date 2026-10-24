package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
type GenS3 struct { F0 uint; F1 uint16; F2 int }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint64 = 1
    var v2 string = "a"
    var v3 float32 = 0.0
    var v4 uint16 = 1
    v5 := []int{}
    for fi6 := 0; fi6 < 17; fi6++ {
        v2 = v2
        if ((((-2.25 * (0.1 - 1e10)) <= 1e10) || (18446744073709551615 != v1)) && (187 == v1)) {
            switch (fi6 ^ (fi6 - fi6)) {
            case 2147483647:
                if false {
                    var v7 int = (fi6 & (fi6 & fi6))
                    _ = v7
                }
            case 111:
                v4 -= 2
            default:
                v1 += v1
            }
            var v8 int = (^(fi6))
            fmt.Println(v4)
            _ = v8
        } else {
            v9 := GenS1{F0: 2, F1: (v4 << ((2) & 63)), F2: v4}
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS2{F0: false, F1: (v2 + ((v2 + (v2 + ((v2 + "") + "foo bar"))) + v2)), F2: 2147483647}
    v11 := GenS2{F0: true, F1: v2, F2: 2147483647}
    if (v4 >= (v4 >> ((970) & 63))) {
        v12 := GenS1{F0: 792, F1: 255, F2: v4}
        _ = v12
    }
    var v13 uint32 = 4294967295
    if (!(2147483647 >= -2147483648)) {
        v4 = (v4 - (^(v4)))
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v14 int32 = -2147483648
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v13 -= v13
    v13 = 354
    for fi15 := 0; fi15 < 13; fi15++ {
        v2 = ""
        _ = fi15
    }
    switch -982 {
    case 0:
        switch "hello" {
        case "a":
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 float32 = (float32((-2.25 + 0.0)) * (v3 * v3))
                _ = fi16
                _ = v17
            }
        case "foo bar":
            if len(v5) > 0 {
                v5[0] = -2147483648
            }
        default:
            v14 = v14
        }
    case 551:
        if (-561 == -34) {
            var v18 int = 42
            var v19 uint = 986
            _ = v18
            _ = v19
        } else {
            if (!(-32768 <= 42)) {
                if (32768 < 42) {
                    v20 := GenS3{F0: 4294967295, F1: v4, F2: 639}
                    if ((((0.1 + 1.5) != 3.14159) || (v1 <= (v1 % ((v1) | 1)))) && (false || (1e10 == (float64(v3) / 1e10)))) {
                        v3 -= 1.5
                    } else {
                        var v21 float64 = 1.5
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v21
                    }
                    v1 = v1
                    _ = v20
                }
                var v22 uint = 255
                _ = v22
            } else {
                if false {
                    switch 32767 {
                    case -2147483648:
                        if (1 > -142) {
                            for fi23 := 0; fi23 < 4; fi23++ {
                                v4 += (v4 ^ (v4 & (^(v4))))
                                _ = fi23
                            }
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                    }
                } else {
                    v14 += v14
                    var v24 uint8 = 210
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v24
                }
                v2 = v2
                v25 := GenS4{F0: false}
                _ = v25
            }
            v5 = append(v5, 2147483647)
        }
    default:
        fmt.Println(v3, v13, v2)
    }
    v26 := GenS1{F0: 0, F1: 42, F2: v4}
    v2 = v2
    switch (("" + ("hello" + "hello")) + v2) {
    case "\n\t":
        fmt.Println(v4, v2, v14)
    case "":
        v14 += v14
    case "hello":
        v4 -= 646
    }
    v27 := GenS2{F0: (v4 >= 42), F1: v2, F2: v14}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
