package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8; F1 string; F2 int8 }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64; F1 int }

func main() {
    var v1 int8 = 100
    var v2 int8 = 100
    var v3 uint64 = 42
    var v4 uint64 = (^(v3))
    var v5 uint = 18446744073709551615
    var v6 int16 = 42
    if true {
        var v7 int32 = 2147483647
        if false {
            v3 = v3
            if ((((-2.25 * 0.1) * 3.14159) * 3.14159) <= 1.5) {
                var v8 string = "foo bar"
                v9 := GenS1{F0: 2, F1: 182, F2: -40}
                v4 += v4
                _ = v8
                _ = v9
            }
            v2 = v1
        } else {
            v5 -= (^(v5))
            for fi10 := 0; fi10 < 12; fi10++ {
                if (true || false) {
                    if true {
                        v11 := []int{(fi10 >> ((v5) & 63)), (^(fi10)), (fi10 << ((v5) & 63))}
                        _ = v11
                    }
                } else {
                    var v12 uint16 = 76
                    _ = v12
                }
                v4 -= (v3 - v4)
                _ = fi10
            }
            v2 = 1
        }
        _ = v7
    } else {
        v13 := GenS1{F0: (v5 ^ v5), F1: 81, F2: 2147483647}
        var v14 int = -995
        v3 -= v3
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := []int{fi15}
            _ = fi16
            _ = v17
        }
        var v18 uint32 = 990
        _ = fi15
        _ = v18
    }
    var v19 float32 = float32(1.5)
    v5 += (v5 + v5)
    v5 = v5
    v20 := []int{1, -2147483648, 144}
    _ = len(v20)
    var v21 int32 = 316
    v5 = (v5 >> (((v5 << ((v5) & 63))) & 63))
    fmt.Println(v19)
    var v22 int = 345
    var v23 uint16 = 0
    v20 = append(v20, (v22 / ((2147483647) | 1)))
    v6 -= int16(v22)
    if ((!(false || true)) && true) {
        switch (v22 + v22) {
        case -9223372036854775808:
            var v24 float32 = (1.5 * (float32(1.5) / v19))
            _ = v24
        case -340:
            if false {
                if false {
                    var v25 float64 = 0.1
                    fmt.Println(v6, v5, v25)
                    _ = v25
                } else {
                    if (65535 >= 434) {
                        v2 = (^(v1))
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                    if ((v19 < v19) && false) {
                        if len(v20) > 0 {
                            v20[0] = v22
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v26 int = int(v2)
                            _ = v26
                        case "a":
                            if len(v20) > 0 {
                                v20[0] = 615
                            }
                        case "":
                            v23 = (v23 % ((v23) | 1))
                        }
                        if ((-(v22)) <= int(v6)) {
                            v21 = 437
                        } else {
                            v23 = v23
                            v27 := GenS1{F0: uint(v4), F1: 127, F2: (v22 * 4294967295)}
                            if true {
                                if ((1.5 / (-2.25 - ((1.5 + 0.1) - (float32(0.0) * float32((1.5 / 1.0)))))) <= -2.25) {
                                    if false {
                                        v28 := GenS1{F0: 65535, F1: 127, F2: v22}
                                        _ = len(v20)
                                        _ = v28
                                    }
                                    if false {
                                        var v29 float64 = (float64(0.1) - 0.1)
                                        for fi30 := 0; fi30 < 14; fi30++ {
                                            if len(v20) > 0 {
                                                v20[0] = 128
                                            }
                                            fmt.Println(v5)
                                            var v31 int = -825
                                            _ = fi30
                                            _ = v31
                                        }
                                        for fi32 := 0; fi32 < 15; fi32++ {
                                            v2 -= (^(v2))
                                            if (1 != v23) {
                                                if false {
                                                    for fi33 := 0; fi33 < 10; fi33++ {
                                                        v34 := GenS1{F0: v5, F1: 224, F2: (fi33 | 42)}
                                                        v35 := GenS2{F0: v23, F1: 128}
                                                        _ = fi33
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    var v36 string = ((("foo bar" + "foo bar") + "\n\t") + "")
                                                    _ = v36
                                                }
                                                var v37 uint8 = 24
                                                if len(v20) >= 2 {
                                                    v20 = v20[1:len(v20)]
                                                }
                                                _ = v37
                                            } else {
                                                v20 = append(v20, fi32)
                                                for fi38 := 0; fi38 < 14; fi38++ {
                                                    var v39 bool = ((!false) && false)
                                                    var v40 uint = 65535
                                                    v39 = v39
                                                    _ = fi38
                                                    _ = v39
                                                    _ = v40
                                                }
                                                v29 += 3.14159
                                            }
                                            _ = fi32
                                        }
                                        _ = v29
                                    } else {
                                        v6 -= (^(v6))
                                        v41 := GenS3{F0: v2, F1: "a", F2: (v1 >> ((v5) & 63))}
                                        fmt.Println(v6, v22, v2)
                                        _ = v41
                                    }
                                    if true {
                                        for fi42 := 0; fi42 < 18; fi42++ {
                                            switch fi42 {
                                            case 9223372036854775807:
                                                switch fi42 {
                                                case -372:
                                                    if ((!false) && false) {
                                                        v43 := GenS2{F0: (^(v23)), F1: 76}
                                                        _ = v43
                                                    } else {
                                                        if (true && true) {
                                                            v20 = append(v20, (v22 & fi42))
                                                            if false {
                                                                for fi44 := 0; fi44 < 6; fi44++ {
                                                                    var v45 int64 = 42
                                                                    if (v45 < (v45 - -325)) {
                                                                        if len(v20) > 0 {
                                                                            v20[0] = -1
                                                                        }
                                                                    } else {
                                                                        var v46 int32 = 2147483647
                                                                        _ = v46
                                                                    }
                                                                    _ = fi44
                                                                    _ = v45
                                                                }
                                                                v22 = fi42
                                                            } else {
                                                                var v47 uint32 = 975
                                                                v23 = (v23 << (((v5 / (((^(v5))) | 1))) & 63))
                                                                if (fi42 != v22) {
                                                                    var v48 string = (("hello" + "\n\t") + "foo bar")
                                                                    for fi49 := 0; fi49 < 3; fi49++ {
                                                                        _ = len(v20)
                                                                        v50 := GenS3{F0: v1, F1: "a", F2: (v2 / ((v2) | 1))}
                                                                        _ = fi49
                                                                        _ = v50
                                                                    }
                                                                    _ = len(v20)
                                                                    _ = v48
                                                                } else {
                                                                    v21 = int32(v3)
                                                                    v23 = v23
                                                                }
                                                                _ = v47
                                                            }
                                                        }
                                                    }
                                                }
                                            case -1:
                                                switch fi42 {
                                                case -650:
                                                    v51 := GenS3{F0: v1, F1: "", F2: (-(v1))}
                                                    _ = v51
                                                case 32767:
                                                    if (1 > (v5 - v5)) {
                                                        switch "" {
                                                        case "foo bar":
                                                            if (!false) {
                                                                switch (-(fi42)) {
                                                                case -730:
                                                                    switch (v22 / ((664) | 1)) {
                                                                    case -464:
                                                                        v23 += uint16(v5)
                                                                    case 870:
                                                                        v2 = (v1 * v1)
                                                                    default:
                                                                        v52 := GenS4{F0: (-2.25 != 3.14159), F1: v5, F2: (^(v5))}
                                                                        _ = v52
                                                                    }
                                                                case -1:
                                                                    for fi53 := 0; fi53 < 20; fi53++ {
                                                                        v6 += int16(fi42)
                                                                        _ = fi53
                                                                    }
                                                                }
                                                                for fi54 := 0; fi54 < 20; fi54++ {
                                                                    var v55 string = "\n\t"
                                                                    if (v22 != fi54) {
                                                                        v56 := GenS5{F0: (fi54 & (fi54 | 127)), F1: (v5 / ((v5) | 1))}
                                                                        v20 = append(v20, v22)
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            if false {
                                                                                if true {
                                                                                    v57 := GenS6{F0: v3}
                                                                                    var v58 float64 = (((((float64(-2.25) / 1.5) * 1.5) + 0.0) - -2.25) + (1e10 - 0.1))
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                                var v59 int64 = -809
                                                                                var v60 int = -1
                                                                                _ = v59
                                                                                _ = v60
                                                                            } else {
                                                                                v5 = v5
                                                                                switch fi54 {
                                                                                case 0:
                                                                                    v6 = -253
                                                                                case 717:
                                                                                    var v61 uint = (^(v5))
                                                                                    _ = v61
                                                                                case 586:
                                                                                    v5 += 65535
                                                                                }
                                                                                if true {
                                                                                    if len(v20) >= 2 {
                                                                                        v20 = v20[1:len(v20)]
                                                                                    }
                                                                                    if len(v20) >= 2 {
                                                                                        v20 = v20[1:len(v20)]
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v6, v5, v3)
                                                                                    v62 := GenS7{F0: 42, F1: "foo bar"}
                                                                                    switch (fi42 / ((127) | 1)) {
                                                                                    case 0:
                                                                                        switch fi42 {
                                                                                        case -1:
                                                                                            fmt.Println(v4, v6)
                                                                                        case 127:
                                                                                            v63 := GenS3{F0: (v1 % (((v2 ^ -128)) | 1)), F1: "a", F2: v1}
                                                                                            _ = v63
                                                                                        default:
                                                                                            if true {
                                                                                                v6 = 32767
                                                                                                v64 := GenS7{F0: v23, F1: v55}
                                                                                                if false {
                                                                                                    var v65 int16 = v6
                                                                                                    if len(v20) > 0 {
                                                                                                        v20[0] = fi42
                                                                                                    }
                                                                                                    v20 = append(v20, (fi42 & fi42))
                                                                                                    var v66 bool = ((v1 % (((v1 + v2)) | 1)) != (^(v1)))
                                                                                                    v20 = append(v20, v22)
                                                                                                    v2 = v2
                                                                                                    _ = v65
                                                                                                    _ = v66
                                                                                                } else {
                                                                                                    if len(v20) > 0 {
                                                                                                        v20[0] = v22
                                                                                                    }
                                                                                                    var v67 int64 = int64(v1)
                                                                                                    _ = v67
                                                                                                }
                                                                                                _ = v64
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v68 int16 = (^(v6))
                                                                                        _ = v68
                                                                                    }
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v56
                                                                    } else {
                                                                        v69 := GenS8{F0: 127, F1: v23}
                                                                        _ = v69
                                                                    }
                                                                    _ = fi54
                                                                    _ = v55
                                                                }
                                                            }
                                                        default:
                                                            v6 -= (v6 % ((v6) | 1))
                                                        }
                                                    }
                                                }
                                            case 128:
                                                v70 := GenS8{F0: 108, F1: v23}
                                                _ = v70
                                            }
                                            var v71 string = "\n\t"
                                            _ = fi42
                                            _ = v71
                                        }
                                    } else {
                                        if (float64(3.14159) <= ((0.1 - 1.5) * 0.0)) {
                                            v19 += (v19 - v19)
                                        }
                                        for fi72 := 0; fi72 < 12; fi72++ {
                                            v73 := GenS2{F0: (v23 + uint16(v5)), F1: 255}
                                            _ = fi72
                                            _ = v73
                                        }
                                        var v74 float32 = float32(float64(v19))
                                        _ = v74
                                    }
                                } else {
                                    v3 -= v3
                                    if false {
                                        v75 := GenS9{F0: v4}
                                        var v76 int = v22
                                        _ = v75
                                        _ = v76
                                    } else {
                                        var v77 int8 = (v1 << (((v5 ^ v5)) & 63))
                                        if len(v20) > 0 {
                                            v20[0] = (v22 % ((-507) | 1))
                                        }
                                        _ = v77
                                    }
                                }
                                v78 := GenS7{F0: (v23 / ((v23) | 1)), F1: "a"}
                                var v79 string = ((((("a" + ("hello" + "hello")) + ((("" + ("hello" + "")) + "\n\t") + ("" + "\n\t"))) + "foo bar") + ("hello" + "foo bar")) + "foo bar")
                                _ = v78
                                _ = v79
                            } else {
                                v6 -= (v6 >> ((v5) & 63))
                            }
                            _ = v27
                        }
                    } else {
                        var v80 uint8 = 128
                        v4 += (v4 * 4294967295)
                        var v81 float64 = 3.14159
                        _ = v80
                        _ = v81
                    }
                    if len(v20) > 0 {
                        v20[0] = v22
                    }
                }
            } else {
                v22 = 32767
            }
        default:
            _ = len(v20)
        }
        if true {
            for fi82 := 0; fi82 < 15; fi82++ {
                switch "foo bar" {
                case "\n\t":
                    v5 = (^(v5))
                default:
                    var v83 uint = v5
                    _ = v83
                }
                _ = fi82
            }
        }
        v1 -= v2
    } else {
        for fi84 := 0; fi84 < 20; fi84++ {
            v19 += 0.1
            v85 := GenS10{F0: (^(v3)), F1: (fi84 >> ((v5) & 63))}
            var v86 uint = (v5 << (((^(v5))) & 63))
            _ = fi84
            _ = v85
            _ = v86
        }
        var v87 bool = (-20 < (v2 + v1))
        _ = v87
    }
    v2 = (v1 * 0)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
