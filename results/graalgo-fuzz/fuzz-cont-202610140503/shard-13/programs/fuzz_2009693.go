package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 string; F1 int64; F2 uint16 }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS10 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 uint32, p1 string, p2 uint8) uint16 {
    var v6 int = 1
    var v7 uint32 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 string) uint64 {
    switch 9223372036854775807 {
    case 2147483647:
        v8 := GenS1{F0: 615}
        _ = v8
    default:
        p0 = 196
    }
    var v9 int = 32768
    v10 := GenS2{F0: 74, F1: 4294967295, F2: false}
    v9 = (-(v9))
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int64 = -9223372036854775808
        v13 := []int{v9, 42, (fi11 * v9)}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 bool = true
    v2 := []uint32{65535, 2, 4294967295, 93}
    v3 := make(map[int]string)
    var v4 int8 = -128
    v5 := []uint32{}
    for fi14 := 0; fi14 < 12; fi14++ {
        fmt.Println(v4, v4, v1)
        _ = fi14
    }
    v15 := GenS1{F0: 591}
    var v16 uint16 = 989
    if false {
        var v17 int8 = 100
        if ((((1.5 / float32(-2.25)) - -2.25) >= float32(float64(0.0))) || (!v1)) {
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := v3[(fi18 * fi18)]
                v16 -= v16
                _ = fi18
                _ = v19
            }
        } else {
            if (0.1 > float64(0.1)) {
                if v1 {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if v1 {
                        var v20 uint64 = genFn1(0, "\n\t")
                        _ = v20
                    } else {
                        _ = len(v3)
                        v16 = genFn0(780, ("foo bar" + "foo bar"), 71)
                        if (v1 || false) {
                            v1 = (0 != 255)
                            v17 -= -37
                        }
                    }
                    switch "hello" {
                    case "":
                        v21 := GenS2{F0: v4, F1: 825, F2: v1}
                        _ = v21
                    case "\n\t":
                        _ = len(v2)
                    default:
                        fmt.Println(v1, v16, v17)
                    }
                }
            } else {
                v1 = true
                if true {
                    v22 := GenS2{F0: v17, F1: 1, F2: (!v1)}
                    if (41 <= 2147483648) {
                        delete(v3, -2147483648)
                    } else {
                        fmt.Println(v16)
                    }
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        for fi24 := 0; fi24 < 9; fi24++ {
                            fmt.Println(v17, fi24)
                            _ = fi24
                        }
                        _ = fi23
                    }
                }
            }
            v16 += v16
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
        v4 = v17
        _ = v17
    } else {
        if (-2.25 == -2.25) {
            v25 := GenS2{F0: (v4 | (v4 % ((v4) | 1))), F1: 871, F2: ((!(1 <= 2147483647)) || v1)}
            var v26 int64 = 9223372036854775807
            _ = v25
            _ = v26
        }
        switch 329 {
        case 44:
            var v27 uint16 = (v16 << ((42) & 63))
            _ = v27
        case -391:
            var v28 string = ""
            _ = v28
        case -1:
            var v29 uint8 = 127
            _ = v29
        default:
            v16 = genFn0(497, ("hello" + "\n\t"), 255)
        }
    }
    v4 = v4
    var v30 uint64 = 18446744073709551615
    delete(v3, -1)
    if v1 {
        v16 -= genFn0(256, "\n\t", 127)
    } else {
        if (!(("" + "hello") <= ((("foo bar" + "") + "\n\t") + ""))) {
            switch 1 {
            case 2147483647:
                for fi31 := 0; fi31 < 18; fi31++ {
                    v1 = (v1 && false)
                    if (18446744073709551615 < 65535) {
                        v2 = append(v2, uint32(v30))
                    } else {
                        var v32 uint8 = 1
                        _ = v32
                    }
                    var v33 bool = v1
                    _ = fi31
                    _ = v33
                }
            case -9223372036854775808:
                v34 := GenS1{F0: 42}
                _ = v34
            }
            v4 -= (v4 + int8(v30))
            fmt.Println(v4)
        } else {
            var v35 int8 = v4
            var v36 int16 = -926
            fmt.Println(v16)
            _ = v35
            _ = v36
        }
        v30 = v30
        var v37 int64 = 456
        _ = v37
    }
    v38 := GenS2{F0: (v4 + (^(v4))), F1: 461, F2: (v1 && (-9223372036854775808 < 9223372036854775807))}
    switch (("" + "foo bar") + "\n\t") {
    case "\n\t":
        var v39 uint16 = (^(v16))
        _ = v39
    case "a":
        var v40 float32 = (0.0 / 3.14159)
        _ = v40
    case "":
        _ = len(v3)
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        fmt.Println(v30, v30, v1)
        _ = fi41
    }
    if v1 {
        var v42 uint = 4294967295
        v43 := GenS2{F0: v4, F1: 950, F2: (128 >= 34)}
        _ = v42
        _ = v43
    }
    fmt.Println(v1, v30)
    var v44 bool = (-32768 >= 980)
    for fi45 := 0; fi45 < 5; fi45++ {
        v46 := v3[fi45]
        if v1 {
            switch (fi45 % ((32767) | 1)) {
            case -77:
                for fi47 := 0; fi47 < 4; fi47++ {
                    _ = len(v3)
                    _ = len(v3)
                    fmt.Println(fi47, v4)
                    v30 = 18446744073709551615
                    if v1 {
                        v4 = v4
                    } else {
                        switch -2147483648 {
                        case -2147483648:
                            v46 = v46
                        }
                        v48 := GenS3{F0: 848, F1: -247}
                        if v44 {
                            if (v30 <= (^(v30))) {
                                v49 := GenS4{F0: v16}
                                switch 932 {
                                case 9223372036854775807:
                                    switch fi45 {
                                    case -2147483648:
                                        v44 = v1
                                    case 42:
                                        switch v46 {
                                        case "hello":
                                            v44 = v44
                                        case "foo bar":
                                            if (("a" + (v46 + v46)) >= (("\n\t" + "foo bar") + (v46 + ("foo bar" + (("" + ((v46 + "a") + (v46 + "foo bar"))) + v46))))) {
                                                var v50 float32 = 0.0
                                                var v51 uint32 = 2147483647
                                                _ = v50
                                                _ = v51
                                            }
                                        case "a":
                                            v52 := GenS4{F0: v16}
                                            _ = v52
                                        }
                                    case 32768:
                                        if len(v2) > 0 {
                                            v2[0] = 993
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                case -2147483648:
                                    v53 := GenS5{F0: 2}
                                    _ = v53
                                case 128:
                                    var v54 int = (fi45 * (fi45 & fi47))
                                    _ = v54
                                }
                                v55 := GenS6{F0: 113, F1: int16(v30)}
                                _ = v49
                                _ = v55
                            } else {
                                v46 = (v46 + (v46 + ("" + ("hello" + v46))))
                            }
                            v56 := GenS7{F0: "\n\t", F1: -9223372036854775808, F2: (v16 & (v16 | v16))}
                            _ = v56
                        } else {
                            fmt.Println(v44)
                        }
                        var v57 int64 = -9223372036854775808
                        v57 -= v57
                        switch -2147483648 {
                        case -1:
                            switch v46 {
                            case "\n\t":
                                fmt.Println(v46, v57, v57)
                            case "foo bar":
                                if (2 != 1) {
                                    v58 := GenS8{F0: -1, F1: 42}
                                    v30 -= 30
                                    _ = v58
                                }
                            case "":
                                v1 = true
                            }
                        }
                        if false {
                            v16 -= (v16 & (v16 / ((genFn0(428, v46, 0)) | 1)))
                        }
                        var v59 uint8 = 127
                        _ = v48
                        _ = v57
                        _ = v59
                    }
                    _ = fi47
                }
            default:
                v60 := GenS7{F0: v46, F1: -1, F2: genFn0(1, v46, 2)}
                _ = v60
            }
        }
        _ = fi45
        _ = v46
    }
    if ("\n\t" < "a") {
        var v61 uint8 = 128
        _ = v61
    } else {
        if (65535 < 42) {
            switch 2147483647 {
            case 0:
                if false {
                    v30 += (v30 ^ (v30 ^ 0))
                }
            }
        } else {
            for fi62 := 0; fi62 < 16; fi62++ {
                if len(v5) > 0 {
                    v5[0] = 802
                }
                var v63 uint8 = uint8(v30)
                _ = fi62
                _ = v63
            }
            switch 1 {
            case 32768:
                switch 128 {
                case 1:
                    v4 += -1
                case -472:
                    for fi64 := 0; fi64 < 7; fi64++ {
                        var v65 bool = (v44 && (65535 > 42))
                        if false {
                            v65 = v1
                            v66 := GenS6{F0: v4, F1: int16(v16)}
                            _ = v66
                        } else {
                            var v67 uint = 255
                            v16 = v16
                            _ = v67
                        }
                        v16 = (v16 ^ (v16 / (((v16 * 0)) | 1)))
                        _ = fi64
                        _ = v65
                    }
                }
            case 490:
                v16 += (v16 * genFn0(145, "", uint8(v16)))
            default:
                for fi68 := 0; fi68 < 8; fi68++ {
                    if (true || ((v30 & 2) != v30)) {
                        v44 = ("" > "\n\t")
                    }
                    v1 = v1
                    if (!v1) {
                        if (!(127 >= (v4 - (v4 / (((v4 - (v4 | 42))) | 1))))) {
                            fmt.Println(fi68)
                            if ((-736 <= -1) || v44) {
                                if ((v30 >> ((837) & 63)) < (v30 - uint64(v16))) {
                                    var v69 int = 9223372036854775807
                                    switch fi68 {
                                    case 128:
                                        switch (fi68 / ((v69) | 1)) {
                                        case 164:
                                            v70 := GenS7{F0: "", F1: 0, F2: 1}
                                            _ = v70
                                        case 127:
                                            if v44 {
                                                var v71 uint64 = 4294967295
                                                _ = v71
                                            }
                                        case 225:
                                            if v1 {
                                                v44 = v1
                                                v69 = 0
                                                switch fi68 {
                                                case 0:
                                                    v44 = (691 == 65535)
                                                case 9223372036854775807:
                                                    v72 := GenS4{F0: v16}
                                                    _ = v72
                                                case 128:
                                                    fmt.Println(v44, v44)
                                                }
                                            } else {
                                                v16 = (v16 ^ v16)
                                            }
                                        }
                                    case 570:
                                        switch (fi68 << ((2) & 63)) {
                                        case 42:
                                            fmt.Println(fi68)
                                        case 1:
                                            v1 = (int(v16) > -409)
                                        default:
                                            var v73 string = ""
                                            _ = v73
                                        }
                                    }
                                    for fi74 := 0; fi74 < 8; fi74++ {
                                        v75 := v3[v69]
                                        var v76 string = "foo bar"
                                        _ = fi74
                                        _ = v75
                                        _ = v76
                                    }
                                    _ = v69
                                }
                                fmt.Println(v44)
                                var v77 int16 = 42
                                _ = v77
                            } else {
                                switch (fi68 / (((fi68 / ((fi68) | 1))) | 1)) {
                                case -2147483648:
                                    fmt.Println(v1, v30)
                                default:
                                    var v78 uint16 = (v16 << ((65535) & 63))
                                    _ = v78
                                }
                            }
                        }
                    }
                    _ = fi68
                }
            }
        }
        if (42 < 2147483647) {
            var v79 uint8 = 255
            _ = v79
        } else {
            v80 := GenS8{F0: -766, F1: -32768}
            for fi81 := 0; fi81 < 18; fi81++ {
                switch "a" {
                case "foo bar":
                    v16 -= v16
                case "":
                    var v82 uint32 = 183
                    _ = v82
                case "a":
                    switch ("" + "\n\t") {
                    case "a":
                        fmt.Println(v16, v4, fi81)
                    case "foo bar":
                        for fi83 := 0; fi83 < 19; fi83++ {
                            switch fi83 {
                            case 32767:
                                v44 = (2 > 42)
                            case -1:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                if (157 != 2) {
                                    var v84 uint = 200
                                    switch fi83 {
                                    case 359:
                                        var v85 int16 = -1
                                        _ = v85
                                    case 32768:
                                        v86 := GenS4{F0: (v16 >> (((v84 << ((v84) & 63))) & 63))}
                                        _ = v86
                                    case 127:
                                        var v87 int8 = (v4 | (-(v4)))
                                        _ = v87
                                    }
                                    _ = v84
                                } else {
                                    var v88 uint8 = 140
                                    _ = v88
                                }
                            }
                            _ = fi83
                        }
                    default:
                        if (v30 > (v30 | (^(v30)))) {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v4)
                            case "\n\t":
                                var v89 bool = true
                                _ = v89
                            }
                            v4 = -125
                        }
                    }
                default:
                    fmt.Println(v16, v4, v4)
                }
                if (v44 && false) {
                    var v90 uint = 1
                    if false {
                        var v91 uint8 = 127
                        delete(v3, 390)
                        _ = v91
                    } else {
                        v92 := GenS4{F0: genFn0(4294967295, "foo bar", 16)}
                        _ = v92
                    }
                    _ = v90
                } else {
                    v93 := GenS5{F0: 2}
                    _ = v93
                }
                _ = fi81
            }
            var v94 int = 603
            _ = v80
            _ = v94
        }
        if (true || v44) {
            v44 = v44
            for fi95 := 0; fi95 < 15; fi95++ {
                v96 := v3[-1]
                for fi97 := 0; fi97 < 1; fi97++ {
                    if ((v96 == v96) && false) {
                        var v98 int16 = -82
                        var v99 int64 = 792
                        v100 := GenS4{F0: (v16 + 171)}
                        _ = v100
                        _ = v98
                        _ = v99
                    } else {
                        v30 = v30
                    }
                    if v1 {
                        v101 := GenS9{F0: 2, F1: 1, F2: 382}
                        if (v44 || (872 == 119)) {
                            v1 = ((fi97 >> ((uint(v30)) & 63)) <= (fi95 & (fi97 + (fi97 % (((fi97 >> ((1) & 63))) | 1)))))
                            v102 := GenS9{F0: (v30 << ((1) & 63)), F1: 176, F2: 2147483647}
                            fmt.Println(v4)
                            switch v96 {
                            case "":
                                var v103 uint16 = (^(v16))
                                _ = v103
                            case "foo bar":
                                switch (v96 + (v96 + "a")) {
                                case "hello":
                                    if v44 {
                                        fmt.Println(v16, v16)
                                        var v104 uint64 = v30
                                        v4 += v4
                                        var v105 uint8 = 42
                                        v106 := GenS10{F0: v105, F1: 2, F2: -2147483648}
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                    }
                                case "foo bar":
                                    var v107 int16 = 0
                                    _ = v107
                                default:
                                    var v108 int16 = 42
                                    _ = v108
                                }
                            case "hello":
                                delete(v3, fi97)
                            default:
                                if true {
                                    var v109 float32 = (1.5 + -2.25)
                                    v1 = v44
                                    var v110 uint16 = v16
                                    _ = v109
                                    _ = v110
                                }
                            }
                            var v111 uint32 = 0
                            _ = v102
                            _ = v111
                        }
                        _ = v101
                    }
                    v112 := GenS5{F0: 2147483647}
                    _ = fi97
                    _ = v112
                }
                if v1 {
                    v4 -= v4
                    v2 = append(v2, 4294967295)
                }
                _ = fi95
                _ = v96
            }
        } else {
            var v113 int64 = 9223372036854775807
            _ = v113
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v44
    _ = v5
}
