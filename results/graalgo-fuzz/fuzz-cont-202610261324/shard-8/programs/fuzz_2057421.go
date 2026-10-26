package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 string }
func genFn0(p0 bool, p1 int, p2 bool) int {
    v7 := GenS1{F0: 0}
    v8 := GenS1{F0: 2}
    p2 = (2 == 207)
    var v9 int32 = -1
    var v10 int8 = -37
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p1
}
func genFn1() int8 {
    var v11 float64 = 1.5
    fmt.Println(v11)
    _ = v11
    return 100
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]uint32)
    var v3 uint8 = 0
    var v4 uint = 2
    var v5 float32 = (3.14159 * 3.14159)
    v6 := []string{"\n\t", "a"}
    fmt.Println(v1, v3, v3)
    if true {
        if true {
            if (!false) {
                if true {
                    var v12 int8 = v1
                    _ = len(v2)
                    var v13 uint = 2
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(v4, v3)
                    v5 += (float32(float64(v5)) - v5)
                    if false {
                        v14 := GenS2{F0: 333}
                        if false {
                            if false {
                                v1 += 1
                                v5 += float32(1e10)
                                v2[32767] = 65535
                            }
                            _ = len(v6)
                            var v15 float64 = 0.1
                            _ = v15
                        } else {
                            var v16 string = ("\n\t" + "")
                            v3 = (v3 ^ 17)
                            if (!(-267 == 2147483647)) {
                                _ = len(v2)
                                var v17 int32 = 1
                                v18 := GenS1{F0: 1}
                                _ = v17
                                _ = v18
                            }
                            _ = v16
                        }
                        _ = v14
                    } else {
                        if ("\n\t" != ("hello" + "a")) {
                            if false {
                                var v19 float64 = 3.14159
                                v4 = 42
                                v3 = v3
                                _ = v19
                            } else {
                                switch (("a" + "a") + ("a" + "")) {
                                case "\n\t":
                                    v20 := GenS2{F0: 224}
                                    _ = v20
                                }
                                var v21 float64 = float64(3.14159)
                                _ = v21
                            }
                        }
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v3, v4)
                        case "a":
                            switch "" {
                            case "\n\t":
                                if (true && (1000 <= -32768)) {
                                    v6 = append(v6, "foo bar")
                                    if (true && (32767 == 245)) {
                                        var v22 uint8 = v3
                                        v23 := GenS2{F0: 18446744073709551615}
                                        _ = v22
                                        _ = v23
                                    } else {
                                        fmt.Println(v1, v5)
                                    }
                                    fmt.Println(v5, v4, v1)
                                } else {
                                    _ = len(v6)
                                    v1 = v1
                                }
                            case "foo bar":
                                var v24 int64 = 308
                                _ = v24
                            }
                        case "foo bar":
                            v25 := GenS2{F0: 1}
                            _ = v25
                        default:
                            v26 := GenS2{F0: 4294967295}
                            _ = v26
                        }
                    }
                }
                fmt.Println(v3, v4)
                switch 128 {
                case 4294967295:
                    v1 -= (v1 >> (((v4 >> (((v4 / ((v4) | 1))) & 63))) & 63))
                case 127:
                    for fi27 := 0; fi27 < 20; fi27++ {
                        var v28 int16 = 42
                        _ = fi27
                        _ = v28
                    }
                case 32767:
                    var v29 float32 = v5
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 9; fi30++ {
                        var v31 uint32 = 27
                        _ = fi30
                        _ = v31
                    }
                }
            } else {
                var v32 uint8 = v3
                fmt.Println(v5, v4, v4)
                _ = v32
            }
            var v33 int = 477
            _ = v33
        } else {
            v4 += uint(v3)
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            _ = len(v6)
            for fi35 := 0; fi35 < 3; fi35++ {
                v1 -= (-(v1))
                var v36 float32 = (((v5 / (((v5 + v5) * v5) - 0.0)) * (v5 / 1.0)) / (((-2.25 / ((v5 * 1e10) - (1e10 * v5))) / 1.0) + float32(0.1)))
                _ = len(v6)
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            v38 := GenS1{F0: 2}
            _ = len(v6)
            v5 = 0.1
            _ = fi37
            _ = v38
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v4 -= v4
    }
    if (213 > 255) {
        v6 = append(v6, "\n\t")
        var v39 uint = v4
        v40 := GenS2{F0: 878}
        _ = v39
        _ = v40
    } else {
        if len(v6) > 0 {
            v6[0] = "\n\t"
        }
        if (0 == 42) {
            for fi41 := 0; fi41 < 8; fi41++ {
                v4 += (v4 | v4)
                _ = len(v2)
                v1 = (^(v1))
                _ = fi41
            }
            fmt.Println(v1)
        } else {
            v5 = v5
        }
    }
    v4 = 2
    if false {
        for fi42 := 0; fi42 < 6; fi42++ {
            fmt.Println(v3, fi42)
            _ = fi42
        }
    } else {
        delete(v2, -260)
        for fi43 := 0; fi43 < 2; fi43++ {
            if (!true) {
                v1 = v1
            }
            var v44 int8 = -2
            _ = fi43
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        if true {
            if ((-2.25 <= (3.14159 - float32(3.14159))) || false) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v4 -= 644
                var v46 int32 = -1
                _ = v46
            } else {
                fmt.Println(v1, fi45)
            }
            switch (fi45 + (fi45 ^ fi45)) {
            case 9223372036854775807:
                v47 := GenS1{F0: 255}
                _ = v47
            default:
                switch (fi45 + 1) {
                case -877:
                    var v48 int64 = 930
                    _ = v48
                case 834:
                    if (false || false) {
                        v1 = (v1 % ((int8(v3)) | 1))
                        if false {
                            var v49 uint16 = 65535
                            for fi50 := 0; fi50 < 9; fi50++ {
                                switch ((("" + "") + "hello") + (("hello" + ("hello" + "")) + "foo bar")) {
                                case "foo bar":
                                    if true {
                                        v51 := GenS2{F0: 0}
                                        if (993 != 758) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                if true {
                                                    v3 = (^(v3))
                                                } else {
                                                    var v52 int16 = -1
                                                    if (int16(v3) < (v52 / ((v52) | 1))) {
                                                        var v53 string = "\n\t"
                                                        _ = len(v6)
                                                        switch genFn0((uint64(v4) != 102), (^(fi50)), ((v49 % ((v49) | 1)) >= v49)) {
                                                        case -9223372036854775808:
                                                            delete(v2, (fi45 << ((v4) & 63)))
                                                        default:
                                                            v54 := GenS1{F0: (^(v49))}
                                                            _ = v54
                                                        }
                                                        v53 = v53
                                                        v3 -= v3
                                                        _ = v53
                                                    } else {
                                                        v55 := v2[374]
                                                        var v56 string = (("foo bar" + "") + ((("a" + "a") + "foo bar") + "a"))
                                                        v5 = (3.14159 * (((float32(0.1) - v5) / -2.25) - (((3.14159 / float32(1.5)) / v5) + v5)))
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    v57 := GenS1{F0: v49}
                                                    _ = v52
                                                    _ = v57
                                                }
                                            case "":
                                                v58 := GenS1{F0: v49}
                                                _ = v58
                                            }
                                        } else {
                                            var v59 uint64 = 782
                                            v60 := GenS3{F0: v49, F1: v3, F2: (v3 << ((255) & 63))}
                                            _ = v59
                                            _ = v60
                                        }
                                        _ = v51
                                    } else {
                                        if ((int16(v4) > -1) || true) {
                                            var v61 bool = (!(908 == 65535))
                                            _ = v61
                                        } else {
                                            v5 = (3.14159 + v5)
                                        }
                                        if ((v4 % ((v4) | 1)) <= (v4 / ((v4) | 1))) {
                                            var v62 int8 = -1
                                            _ = v62
                                        }
                                        v1 = (v1 * (v1 + (v1 & (v1 ^ v1))))
                                    }
                                case "a":
                                    var v63 float32 = (((v5 + 1e10) - v5) + v5)
                                    _ = v63
                                case "":
                                    fmt.Println(v5)
                                }
                                var v64 string = (("\n\t" + ("foo bar" + ("\n\t" + "a"))) + "a")
                                v65 := GenS4{F0: -32768}
                                _ = fi50
                                _ = v64
                                _ = v65
                            }
                            v2[fi45] = 879
                            _ = v49
                        } else {
                            switch ("foo bar" + "") {
                            case "a":
                                if true {
                                    for fi66 := 0; fi66 < 4; fi66++ {
                                        var v67 float64 = 1.5
                                        _ = len(v6)
                                        _ = fi66
                                        _ = v67
                                    }
                                } else {
                                    for fi68 := 0; fi68 < 1; fi68++ {
                                        var v69 int = 32768
                                        switch fi68 {
                                        case 42:
                                            switch "" {
                                            case "\n\t":
                                                v1 += v1
                                            }
                                        default:
                                            _ = len(v2)
                                            v5 = 1e10
                                            fmt.Println(v1)
                                        }
                                        _ = fi68
                                        _ = v69
                                    }
                                    switch -788 {
                                    case 115:
                                        v3 += (v3 / ((238) | 1))
                                    case -2147483648:
                                        v70 := GenS5{F0: 255, F1: 65535}
                                        _ = v70
                                    }
                                    switch 0 {
                                    case 128:
                                        for fi71 := 0; fi71 < 20; fi71++ {
                                            v72 := GenS6{F0: v3, F1: v3, F2: -119}
                                            switch "\n\t" {
                                            case "hello":
                                                v4 = v4
                                            case "":
                                                v73 := GenS2{F0: 0}
                                                _ = v73
                                            case "a":
                                                v1 -= genFn1()
                                            default:
                                                v74 := GenS3{F0: 1, F1: 128, F2: (^(v3))}
                                                _ = v74
                                            }
                                            _ = fi71
                                            _ = v72
                                        }
                                    case 32767:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    case 32768:
                                        for fi75 := 0; fi75 < 12; fi75++ {
                                            if len(v6) > 0 {
                                                v6[0] = (("foo bar" + "hello") + ("" + "a"))
                                            }
                                            v1 += (v1 * (v1 | 42))
                                            v1 = (v1 | 0)
                                            delete(v2, -9223372036854775808)
                                            if true {
                                                if (!(92 > -1)) {
                                                    var v76 bool = (-1 >= 42)
                                                    _ = v76
                                                } else {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v77 int = (^(fi45))
                                                        _ = v77
                                                    case "a":
                                                        v78 := GenS4{F0: -32768}
                                                        _ = v78
                                                    case "":
                                                        var v79 int = (fi45 - (fi45 + fi45))
                                                        _ = v79
                                                    }
                                                }
                                            }
                                            _ = fi75
                                        }
                                    default:
                                        var v80 uint8 = v3
                                        _ = v80
                                    }
                                }
                            case "hello":
                                v81 := GenS7{F0: 9223372036854775807, F1: (0 != 680)}
                                _ = v81
                            default:
                                for fi82 := 0; fi82 < 15; fi82++ {
                                    fmt.Println(v4, v1)
                                    var v83 uint64 = 0
                                    _ = fi82
                                    _ = v83
                                }
                            }
                        }
                    }
                case -1:
                    for fi84 := 0; fi84 < 8; fi84++ {
                        var v85 bool = (!(!((v1 / ((-128) | 1)) == (v1 << (((^(v4))) & 63)))))
                        var v86 int = fi45
                        _ = fi84
                        _ = v85
                        _ = v86
                    }
                default:
                    var v87 float32 = 0.0
                    _ = v87
                }
            }
        } else {
            var v88 int = (fi45 % ((genFn0(false, genFn0(true, (fi45 * fi45), false), ((v4 == (v4 | v4)) || ("\n\t" != "")))) | 1))
            _ = v88
        }
        for fi89 := 0; fi89 < 7; fi89++ {
            if true {
                _ = len(v6)
                v90 := GenS3{F0: 819, F1: 241, F2: v3}
                _ = v90
            }
            _ = fi89
        }
        _ = fi45
    }
    var v91 string = "hello"
    var v92 int64 = 1
    if false {
        if true {
            fmt.Println(v91, v92)
            for fi93 := 0; fi93 < 7; fi93++ {
                var v94 float64 = float64((3.14159 - 0.1))
                _ = fi93
                _ = v94
            }
        } else {
            for fi95 := 0; fi95 < 2; fi95++ {
                v6 = append(v6, (v91 + v91))
                v5 += (v5 + float32(0.0))
                switch (^(fi95)) {
                case 0:
                    var v96 uint32 = 65535
                    _ = v96
                case -9223372036854775808:
                    for fi97 := 0; fi97 < 14; fi97++ {
                        switch 9223372036854775807 {
                        case -48:
                            if true {
                                switch 724 {
                                case 128:
                                    v4 = v4
                                case -1:
                                    v2[(fi95 % ((fi95) | 1))] = 825
                                case 535:
                                    fmt.Println(v4, fi95, v91)
                                }
                                v98 := GenS7{F0: (v92 >> ((v4) & 63)), F1: (float64(v5) <= 1e10)}
                                v91 = ""
                                v99 := GenS1{F0: 850}
                                _ = v98
                                _ = v99
                            } else {
                                v91 = ""
                                v91 = "hello"
                                var v100 uint8 = v3
                                var v101 uint8 = 0
                                _ = v100
                                _ = v101
                            }
                        }
                        v4 = v4
                        _ = fi97
                    }
                case 9223372036854775807:
                    if (346 < -2147483648) {
                        v6 = append(v6, "a")
                        v102 := GenS1{F0: 2}
                        v2[(^(fi95))] = 42
                        _ = v102
                    } else {
                        if (false && false) {
                            v2[genFn0((!(0.1 >= 1.5)), fi95, false)] = 0
                        }
                        if true {
                            fmt.Println(v92)
                            var v103 int8 = (^(v1))
                            fmt.Println(v5, v91)
                            _ = v103
                        }
                    }
                default:
                    var v104 uint = v4
                    _ = v104
                }
                _ = fi95
            }
            switch v91 {
            case "a":
                v3 = (v3 - (v3 + (v3 - v3)))
            }
        }
    } else {
        if false {
            v1 += v1
        }
        v5 = v5
        var v105 int16 = 895
        _ = v105
    }
    if true {
        if true {
            if ((v1 << (((v4 << ((uint(v3)) & 63))) & 63)) == (v1 * (v1 * (^(v1))))) {
                if (false && (v91 > (v91 + ((v91 + (v91 + v91)) + v91)))) {
                    v106 := GenS7{F0: (v92 >> ((v4) & 63)), F1: (318 != 0)}
                    v1 += genFn1()
                    v5 = v5
                    _ = v106
                } else {
                    if (0 > uint32(v3)) {
                        fmt.Println(v92)
                    } else {
                        if false {
                            fmt.Println(v91, v3, v92)
                            for fi107 := 0; fi107 < 18; fi107++ {
                                var v108 uint32 = 535
                                _ = fi107
                                _ = v108
                            }
                        }
                    }
                    fmt.Println(v1)
                    v91 = ((v91 + (v91 + (v91 + (v91 + "\n\t")))) + v91)
                }
                v1 = v1
                var v109 uint = (^(v4))
                _ = v109
            }
            fmt.Println(v91, v91)
        } else {
            fmt.Println(v4, v91)
            v110 := GenS8{F0: v91}
            v3 = v3
            _ = v110
        }
        var v111 int32 = -967
        _ = v111
    } else {
        v4 -= v4
        _ = len(v6)
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v92
}
