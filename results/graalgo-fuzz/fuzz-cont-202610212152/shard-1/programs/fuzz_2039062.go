package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int8; F2 string }
type GenS5 struct { F0 int; F1 int; F2 uint }
type GenS6 struct { F0 uint32; F1 uint; F2 int64 }

func main() {
    var v1 uint8 = 136
    var v2 uint8 = v1
    v3 := make(map[uint32]bool)
    var v4 int8 = 42
    v5 := make(map[string]uint32)
    var v6 int = 220
    var v7 int16 = -1
    v8 := GenS1{F0: 0, F1: v6, F2: (v4 % ((-128) | 1))}
    v6 -= v6
    if true {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{(fi9 ^ (v6 % ((v6) | 1)))}
            switch v6 {
            case 722:
                fmt.Println(v7, v2, v2)
            }
            switch ("" + "") {
            case "":
                v11 := GenS1{F0: v1, F1: (v6 & 9223372036854775807), F2: -1}
                _ = v11
            default:
                fmt.Println(v4, fi9, fi9)
            }
            _ = fi9
            _ = v10
        }
    } else {
        var v12 uint64 = 164
        _ = v12
    }
    fmt.Println(v7)
    var v13 float32 = (((float32(float64(1.5)) / 3.14159) + (((0.1 / 1e10) + (((((3.14159 * float32(1e10)) + -2.25) - 0.1) / 1.0) / ((1.5 * 0.1) + 1.5))) + float32(((1.5 + 3.14159) + (((0.0 + (1e10 / 1.0)) + (((((3.14159 / float64(3.14159)) + 3.14159) + (0.1 * (0.0 - (1e10 / 0.1)))) + 3.14159) - (0.1 + ((0.0 * (float64(1e10) / 1.5)) - ((float64(3.14159) + float64(0.1)) / (1.5 + 3.14159)))))) + -2.25))))) * (1.5 + 1.5))
    v14 := []int{int(v7), 32767, (^(v6)), (v6 | v6)}
    v2 += v1
    v5["hello"] = 421
    if (true && (((1e10 - -2.25) <= 0.1) && (false && (v4 != (v4 % (((v4 << ((235) & 63))) | 1)))))) {
        if (1000 > v7) {
            switch "" {
            case "\n\t":
                v15 := v3[2147483648]
                _ = v15
            case "hello":
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint64 = 458
                    v18 := GenS1{F0: (v1 * v2), F1: (-(fi16)), F2: v4}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            if ((v7 + v7) == v7) {
                v6 += -999
                v19 := GenS2{F0: 168, F1: v4}
                _ = v19
            }
            switch "foo bar" {
            case "a":
                var v20 int16 = v7
                _ = v20
            default:
                var v21 int8 = (-(v4))
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            switch (("\n\t" + (("a" + ("hello" + "a")) + ("" + (("" + "a") + "hello")))) + ("" + "hello")) {
            case "a":
                if len(v14) > 0 {
                    v14[0] = fi22
                }
            case "foo bar":
                fmt.Println(v2, v2)
            case "hello":
                fmt.Println(v7, v7)
            }
            var v23 uint64 = 1
            if true {
                switch "foo bar" {
                case "hello":
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v7 = v7
                        _ = fi24
                    }
                case "a":
                    switch ("foo bar" + "foo bar") {
                    case "":
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    case "hello":
                        var v25 int64 = -541
                        _ = v25
                    default:
                        var v26 uint32 = 42
                        _ = v26
                    }
                default:
                    if true {
                        v27 := GenS1{F0: (v1 >> ((141) & 63)), F1: int(v23), F2: v4}
                        _ = v27
                    }
                }
            }
            _ = fi22
            _ = v23
        }
        if false {
            if ("\n\t" != "foo bar") {
                v28 := GenS1{F0: (v1 | 24), F1: (-(v6)), F2: v4}
                var v29 uint = 2
                fmt.Println(v29)
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 10; fi30++ {
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 float64 = 0.1
                    if true {
                        var v33 string = ("foo bar" + "\n\t")
                        _ = v33
                    } else {
                        v13 = v13
                        var v34 uint32 = 2
                        v2 -= v2
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                var v35 uint16 = 255
                if (uint64(v35) == 4294967295) {
                    for fi36 := 0; fi36 < 5; fi36++ {
                        v35 -= (v35 / (((v35 - (v35 << ((4294967295) & 63)))) | 1))
                        v4 = -1
                        var v37 uint64 = 4294967295
                        _ = fi36
                        _ = v37
                    }
                }
                _ = fi30
                _ = v35
            }
            switch "a" {
            case "":
                switch 9223372036854775807 {
                case 32768:
                    v38 := GenS2{F0: (^(v2)), F1: 42}
                    _ = v38
                case 127:
                    var v39 uint8 = 128
                    _ = v39
                case -327:
                    if (-181 < -1) {
                        v40 := GenS3{F0: 565}
                        var v41 int8 = v4
                        _ = v40
                        _ = v41
                    }
                default:
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
            case "\n\t":
                v7 += v7
            case "a":
                fmt.Println(v13, v13)
            default:
                switch v6 {
                case 184:
                    v5["hello"] = 241
                case 238:
                    for fi42 := 0; fi42 < 3; fi42++ {
                        for fi43 := 0; fi43 < 7; fi43++ {
                            v44 := GenS4{F0: (v2 * (v1 + 128)), F1: v4, F2: ""}
                            v13 = (0.0 * v13)
                            _ = fi43
                            _ = v44
                        }
                        v45 := GenS5{F0: (v6 ^ v6), F1: 488, F2: 32}
                        switch (("foo bar" + ((("\n\t" + "") + "foo bar") + ("" + "a"))) + "") {
                        case "a":
                            _ = len(v14)
                        }
                        _ = fi42
                        _ = v45
                    }
                case 4294967295:
                    if false {
                        if len(v14) > 0 {
                            v14[0] = (v6 & (-(v6)))
                        }
                        if ("\n\t" < "foo bar") {
                            v46 := GenS5{F0: int(v2), F1: (^(v6)), F2: 89}
                            v6 -= (v6 - v6)
                            _ = v46
                        } else {
                            var v47 int = (^(v6))
                            _ = v47
                        }
                        fmt.Println(v4, v1, v4)
                    } else {
                        for fi48 := 0; fi48 < 2; fi48++ {
                            v49 := GenS5{F0: 4294967295, F1: fi48, F2: 18446744073709551615}
                            var v50 uint32 = 0
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        var v51 float64 = float64((-2.25 + v13))
                        _ = v51
                    }
                default:
                    if false {
                        if (v13 > (-2.25 + v13)) {
                            var v52 int64 = -9223372036854775808
                            switch (v6 + v6) {
                            case 1:
                                _ = len(v5)
                            case 924:
                                v53 := GenS6{F0: 2147483648, F1: 216, F2: (v52 << ((4294967295) & 63))}
                                _ = v53
                            default:
                                var v54 int8 = v4
                                _ = v54
                            }
                            v55 := GenS4{F0: 164, F1: 100, F2: ("hello" + "foo bar")}
                            _ = v52
                            _ = v55
                        }
                        if len(v14) > 0 {
                            v14[0] = -283
                        }
                        var v56 int8 = (^(v4))
                        _ = v56
                    }
                }
            }
        } else {
            if (-2147483648 != v6) {
                v57 := GenS5{F0: (v6 * (v6 / (((v6 >> ((2) & 63))) | 1))), F1: (v6 % (((v6 | -9223372036854775808)) | 1)), F2: 42}
                for fi58 := 0; fi58 < 12; fi58++ {
                    for fi59 := 0; fi59 < 13; fi59++ {
                        var v60 string = ""
                        var v61 uint16 = 42
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    _ = fi58
                }
                _ = v57
            } else {
                v14 = append(v14, 4294967295)
                v14 = append(v14, (v6 ^ v6))
            }
            var v62 bool = (!((v2 << ((4294967295) & 63)) != (v2 + v2)))
            _ = v62
        }
    }
    if ((!true) && false) {
        delete(v5, "")
        var v63 int32 = -2147483648
        _ = v63
    }
    var v64 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v7
    _ = v8
}
