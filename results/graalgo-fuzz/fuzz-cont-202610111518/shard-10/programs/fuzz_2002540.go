package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 int }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 4294967295
    var v3 int32 = 42
    if (18446744073709551615 >= 1) {
        var v4 float32 = ((float32(0.0) * 0.1) - (float32(3.14159) / (-2.25 - 0.1)))
        _ = v4
    } else {
        if (0.0 != 1e10) {
            v3 = int32(v2)
        }
        v3 += v3
    }
    if (false || true) {
        var v5 int16 = 674
        _ = v5
    } else {
        switch "a" {
        case "a":
            if false {
                v2 = (v2 * (v2 & (v2 << ((v2) & 63))))
                for fi6 := 0; fi6 < 5; fi6++ {
                    v2 = (v2 | 0)
                    _ = fi6
                }
            } else {
                v3 = v3
            }
        case "hello":
            var v7 uint16 = 415
            _ = v7
        default:
            switch 407 {
            case 42:
                var v8 uint8 = 16
                _ = v8
            case 9223372036854775807:
                v9 := GenS1{F0: -390, F1: -128, F2: v3}
                _ = v9
            case 255:
                if ((((9223372036854775807 > 24) || false) && false) || (392 <= -485)) {
                    v10 := GenS2{F0: "", F1: 373}
                    v2 += v2
                    _ = v10
                }
            }
        }
        delete(v1, "foo bar")
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 float32 = 1.5
                switch fi11 {
                case 477:
                    v14 := []int{}
                    _ = v14
                default:
                    switch (fi11 / ((1) | 1)) {
                    case 127:
                        delete(v1, "")
                        v2 = v2
                        v13 += 1e10
                        var v15 string = "\n\t"
                        _ = v15
                    case 9223372036854775807:
                        if (97 != -1) {
                            fmt.Println(v2, v2)
                            var v16 uint = 255
                            var v17 uint64 = 0
                            _ = v16
                            _ = v17
                        } else {
                            v18 := GenS2{F0: ("" + "hello"), F1: 32767}
                            _ = v18
                        }
                    case 4294967295:
                        switch ("a" + "hello") {
                        case "a":
                            v19 := GenS2{F0: "hello", F1: 103}
                            _ = v19
                        case "foo bar":
                            delete(v1, ("a" + "a"))
                        case "":
                            v20 := GenS1{F0: 42, F1: 42, F2: (v3 - (v3 % ((v3) | 1)))}
                            _ = v20
                        }
                    default:
                        v1["hello"] = (fi12 - fi12)
                    }
                }
                v21 := GenS3{F0: 110, F1: "foo bar", F2: fi11}
                _ = fi12
                _ = v13
                _ = v21
            }
            var v22 int = (fi11 - fi11)
            _ = fi11
            _ = v22
        }
    }
    switch "a" {
    case "\n\t":
        var v23 uint8 = 128
        _ = v23
    default:
        v2 -= (v2 + 36)
    }
    if (!(false || true)) {
        delete(v1, "foo bar")
    }
    var v24 uint32 = 403
    v25 := GenS1{F0: 32767, F1: 42, F2: (-(v3))}
    delete(v1, ("\n\t" + (("a" + "\n\t") + "")))
    if true {
        fmt.Println(v3)
        var v26 int = 53
        _ = v26
    }
    v3 = (^(v3))
    var v27 int8 = 6
    fmt.Println(v2, v3, v27)
    if false {
        fmt.Println(v2)
        if (0.1 > (float32((-2.25 + 3.14159)) / 1e10)) {
            var v28 float32 = 3.14159
            if (!("\n\t" <= ("\n\t" + ("" + "foo bar")))) {
                if (18446744073709551615 >= 18446744073709551615) {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        if ((v3 & (^(v3))) > (-(v3))) {
                            v30 := []int{}
                            for fi31 := 0; fi31 < 12; fi31++ {
                                if ((^(fi29)) != (fi29 & (fi29 | (fi31 / ((fi29) | 1))))) {
                                    v24 = (^(v24))
                                    var v32 bool = (32767 >= 2)
                                    _ = v32
                                }
                                var v33 bool = (fi29 == fi29)
                                _ = fi31
                                _ = v33
                            }
                            if len(v30) > 0 {
                                v30[0] = 267
                            }
                            _ = v30
                        } else {
                            v24 += v24
                        }
                        _ = fi29
                    }
                }
            }
            if false {
                var v34 int64 = 0
                v35 := GenS4{F0: 65535}
                _ = v34
                _ = v35
            }
            _ = v28
        }
    } else {
        delete(v1, "")
        switch -9223372036854775808 {
        case 127:
            v27 += (^(v27))
        case -2147483648:
            fmt.Println(v24, v24, v2)
        case -644:
            if false {
                switch -9223372036854775808 {
                case 287:
                    v24 -= v24
                case 8:
                    for fi36 := 0; fi36 < 13; fi36++ {
                        var v37 int64 = 1
                        _ = fi36
                        _ = v37
                    }
                case -36:
                    fmt.Println(v2, v24)
                }
                v3 = v3
                var v38 bool = (v2 < v2)
                _ = v38
            } else {
                v24 = v24
            }
        }
    }
    v27 = v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
}
