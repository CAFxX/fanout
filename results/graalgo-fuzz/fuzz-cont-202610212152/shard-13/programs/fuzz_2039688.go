package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 bool; F1 int16; F2 uint }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64 }
func genFn0(p0 uint64, p1 int) int32 {
    fmt.Println(p1, p0)
    p0 -= (p0 | p0)
    _ = p0
    _ = p1
    return 471
}

func main() {
    var v1 float32 = 0.0
    var v2 int32 = 2147483647
    var v3 string = "hello"
    var v4 bool = ((!((1.5 * v1) <= 0.1)) && (!false))
    var v5 string = v3
    v6 := []int{0, -2147483648, 32768}
    var v7 uint64 = 42
    var v8 float64 = 3.14159
    fmt.Println(v4)
    var v9 int16 = 755
    var v10 int16 = v9
    v1 += float32((v8 - v8))
    if (!v4) {
        _ = len(v6)
        var v11 int64 = 1
        _ = v11
    } else {
        if true {
            if false {
                for fi12 := 0; fi12 < 14; fi12++ {
                    if len(v6) > 0 {
                        v6[0] = 113
                    }
                    var v13 uint32 = 326
                    v5 = v3
                    _ = fi12
                    _ = v13
                }
                v5 = "a"
                if v4 {
                    if (0 <= -177) {
                        switch "foo bar" {
                        case "a":
                            v4 = (42 < 127)
                        case "\n\t":
                            v14 := GenS1{F0: v3}
                            _ = v14
                        case "hello":
                            fmt.Println(v3)
                        }
                        switch -799 {
                        case -2147483648:
                            _ = len(v6)
                        case 0:
                            var v15 bool = (v4 && v4)
                            _ = v15
                        case 32767:
                            fmt.Println(v10, v1, v5)
                        }
                        switch 32767 {
                        case 317:
                            v2 = v2
                        default:
                            var v16 string = "hello"
                            _ = v16
                        }
                    }
                }
            } else {
                v7 = 823
            }
            if v4 {
                if len(v6) > 0 {
                    v6[0] = -1
                }
                var v17 float32 = v1
                _ = v17
            } else {
                fmt.Println(v2, v8)
                v8 = v8
                _ = len(v6)
            }
            if false {
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := GenS1{F0: (((("hello" + "\n\t") + (v3 + "a")) + v3) + (v3 + ("foo bar" + (v3 + "hello"))))}
                    if (854 < 2147483648) {
                        if (v4 && v4) {
                            var v20 uint = 458
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v22 := GenS2{F0: v10, F1: v7, F2: 32767}
                                _ = len(v6)
                                _ = fi21
                                _ = v22
                            }
                            v7 = (v7 * 879)
                            _ = v20
                        }
                    } else {
                        var v23 uint64 = (v7 & 877)
                        _ = v23
                    }
                    _ = fi18
                    _ = v19
                }
                if v4 {
                    _ = len(v6)
                    for fi24 := 0; fi24 < 5; fi24++ {
                        var v25 string = v3
                        for fi26 := 0; fi26 < 20; fi26++ {
                            if true {
                                v8 += v8
                                v6 = append(v6, (^(fi26)))
                            }
                            _ = fi26
                        }
                        _ = fi24
                        _ = v25
                    }
                }
                _ = len(v6)
            }
        }
    }
    var v27 uint16 = 32767
    var v28 uint32 = 1
    for fi29 := 0; fi29 < 15; fi29++ {
        v30 := GenS2{F0: 42, F1: (v7 + 696), F2: v27}
        switch (-(fi29)) {
        case 1:
            v31 := GenS3{F0: (v9 << ((255) & 63))}
            _ = v31
        case -339:
            if v4 {
                var v32 int8 = 3
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if v4 {
                    fmt.Println(v7)
                    var v33 uint32 = v28
                    _ = v33
                }
                _ = v32
            } else {
                v6 = append(v6, (fi29 / ((fi29) | 1)))
                switch (fi29 % ((fi29) | 1)) {
                case -171:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case 76:
                    v34 := GenS4{F0: 208, F1: (fi29 & fi29)}
                    _ = v34
                default:
                    if v4 {
                        v35 := GenS1{F0: v5}
                        v2 += 683
                        var v36 float64 = float64((v1 * (v1 * float32(1.5))))
                        _ = v35
                        _ = v36
                    }
                }
            }
        case 500:
            for fi37 := 0; fi37 < 2; fi37++ {
                v5 = (v5 + v3)
                fmt.Println(v2, fi37)
                _ = fi37
            }
        default:
            fmt.Println(v8, v7)
        }
        fmt.Println(v3, v1, fi29)
        _ = fi29
        _ = v30
    }
    var v38 int32 = (v2 << ((997) & 63))
    var v39 float32 = v1
    var v40 float64 = (v8 + (1e10 / (-2.25 - (v8 / (float64(1.5) / (float64(v39) / v8))))))
    for fi41 := 0; fi41 < 6; fi41++ {
        for fi42 := 0; fi42 < 14; fi42++ {
            v40 = 1.5
            _ = fi42
        }
        v43 := GenS5{F0: "", F1: v28}
        switch ((("hello" + ((v5 + ((v3 + "foo bar") + v5)) + "hello")) + v3) + "") {
        case "":
            if (v7 < (v7 >> ((405) & 63))) {
                var v44 uint64 = (v7 * (v7 << ((18446744073709551615) & 63)))
                for fi45 := 0; fi45 < 5; fi45++ {
                    v8 = (v8 / ((-2.25 / 1.0) - (v8 - v40)))
                    if len(v6) > 0 {
                        v6[0] = (fi41 & 1)
                    }
                    var v46 int16 = (v10 * v9)
                    v47 := GenS6{F0: (3.14159 <= v40), F1: 42, F2: 65535}
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                for fi48 := 0; fi48 < 4; fi48++ {
                    v49 := GenS6{F0: v4, F1: 1000, F2: 209}
                    if (!(-1 == 83)) {
                        switch int(v28) {
                        case 128:
                            var v50 uint64 = 4294967295
                            _ = v50
                        case 127:
                            v51 := GenS5{F0: "\n\t", F1: (^(v28))}
                            _ = v51
                        case 459:
                            var v52 uint32 = (v28 & v28)
                            _ = v52
                        default:
                            v8 += 1e10
                        }
                        v10 -= v10
                    }
                    _ = fi48
                    _ = v49
                }
                _ = v44
            }
        default:
            if v4 {
                v38 = 0
            }
        }
        _ = fi41
        _ = v43
    }
    var v53 string = "\n\t"
    if (!((((!(42 >= 255)) || true) && ((v5 + (v3 + ((((v3 + v5) + "") + v53) + (v5 + v5)))) >= "a")) || true)) {
        if v4 {
            var v54 uint = 18446744073709551615
            v9 -= (v9 & (-(v10)))
            _ = v54
        } else {
            for fi55 := 0; fi55 < 12; fi55++ {
                switch (fi55 << ((158) & 63)) {
                case 4294967295:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = fi55
            }
            _ = len(v6)
            var v56 uint8 = 42
            _ = v56
        }
        fmt.Println(v5, v2)
    } else {
        var v57 uint16 = 2
        switch "foo bar" {
        case "":
            if v4 {
                v58 := GenS7{F0: uint32(v57), F1: v7}
                var v59 uint32 = (v28 + v28)
                _ = v58
                _ = v59
            }
        case "foo bar":
            switch 294 {
            case -9223372036854775808:
                v60 := GenS8{F0: -982}
                _ = v60
            case 42:
                if ((false || (1 != -2147483648)) || true) {
                    var v61 float32 = 1e10
                    _ = v61
                }
            case 819:
                for fi62 := 0; fi62 < 6; fi62++ {
                    var v63 string = (v53 + "a")
                    if true {
                        var v64 uint64 = v7
                        if v4 {
                            fmt.Println(v38, v7)
                            fmt.Println(v8)
                        }
                        _ = v64
                    }
                    _ = fi62
                    _ = v63
                }
            }
        default:
            for fi65 := 0; fi65 < 18; fi65++ {
                if (((v8 == ((v8 - (-2.25 + 0.0)) / v8)) || ((0.0 > ((float64(-2.25) / v40) - v8)) && (!(!((v57 >> ((4294967295) & 63)) != 538))))) || false) {
                    var v66 int32 = 42
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 1; fi67++ {
                        var v68 uint16 = (v57 * v57)
                        switch fi65 {
                        case 2147483647:
                            switch "a" {
                            case "\n\t":
                                var v69 float64 = 0.1
                                _ = v69
                            case "hello":
                                var v70 int64 = 373
                                _ = v70
                            case "foo bar":
                                var v71 int8 = -87
                                _ = v71
                            }
                            switch fi65 {
                            case 9223372036854775807:
                                v72 := GenS5{F0: (("hello" + ((("" + (v5 + v53)) + (v5 + (v3 + "hello"))) + v53)) + v5), F1: (v28 << ((uint(v7)) & 63))}
                                _ = v72
                            default:
                                var v73 bool = false
                                v74 := GenS1{F0: ("" + v5)}
                                _ = v73
                                _ = v74
                            }
                            v10 -= int16(v28)
                        case 4294967295:
                            var v75 bool = v4
                            _ = v75
                        default:
                            if v4 {
                                var v76 bool = v4
                                if (v53 == v53) {
                                    if (v4 && v76) {
                                        v77 := GenS5{F0: "a", F1: v28}
                                        _ = v77
                                    } else {
                                        v10 = v10
                                        v78 := GenS6{F0: (-9223372036854775808 >= 42), F1: (v9 / (((v10 | v9)) | 1)), F2: 688}
                                        v79 := GenS1{F0: ((v5 + v53) + (v53 + "foo bar"))}
                                        _ = v78
                                        _ = v79
                                    }
                                    var v80 int16 = (v9 / ((32767) | 1))
                                    _ = v80
                                } else {
                                    fmt.Println(fi67, v5, v28)
                                    v27 = (^(v57))
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    var v81 uint = 1
                                    _ = v81
                                }
                                _ = v76
                            } else {
                                if ((v4 && (v7 != (v7 << ((7) & 63)))) || v4) {
                                    v82 := GenS3{F0: v9}
                                    _ = v82
                                } else {
                                    if false {
                                        switch v5 {
                                        case "":
                                            var v83 uint = 391
                                            _ = v83
                                        default:
                                            v7 = (v7 + v7)
                                        }
                                    } else {
                                        v84 := GenS9{F0: -1}
                                        var v85 uint = 1
                                        var v86 float32 = v39
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                    }
                                }
                                v87 := GenS8{F0: 607}
                                if v4 {
                                    if false {
                                        var v88 int32 = genFn0((v7 | v7), (fi67 + (fi65 / ((fi67) | 1))))
                                        _ = v88
                                    } else {
                                        var v89 float32 = (((((v39 + (0.1 * 3.14159)) * v1) - float32((v8 / (v40 - ((3.14159 + -2.25) + v8))))) / ((v1 / 1.5) + (v39 * (v39 * v39)))) + (v39 - v39))
                                        if v4 {
                                            fmt.Println(v68)
                                            if (((v40 + v40) >= v8) || v4) {
                                                v27 = v68
                                                var v90 float64 = (0.0 * 3.14159)
                                                _ = v90
                                            }
                                        }
                                        _ = v89
                                    }
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    switch ("foo bar" + (v3 + v3)) {
                                    case "\n\t":
                                        v4 = v4
                                    case "a":
                                        v10 = 1000
                                    }
                                }
                                _ = v87
                            }
                        }
                        _ = fi67
                        _ = v68
                    }
                }
                v27 -= (^(v27))
                _ = fi65
            }
        }
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v53
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
