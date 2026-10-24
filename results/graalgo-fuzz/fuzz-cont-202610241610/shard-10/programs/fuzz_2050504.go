package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int; F1 string; F2 int32 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int) int16 {
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint32 = 387
            if true {
                v9 := GenS1{F0: 454, F1: -459, F2: 2}
                var v10 int8 = -71
                var v11 float64 = 3.14159
                v12 := GenS2{F0: "foo bar", F1: p0}
                v13 := []int{-2147483648, (fi6 | fi6), 9223372036854775807, p0}
                fmt.Println(fi6, v8)
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v9
            }
            v8 = (^(v8))
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v14 := []int{p0, 4294967295, (p0 / ((p0) | 1)), p0}
    switch (p0 - 9223372036854775807) {
    case 9223372036854775807:
        if (1.5 < ((0.0 + (0.0 - (-2.25 * (-2.25 / 1e10)))) * -2.25)) {
            p0 = 2147483647
            switch ("" + "foo bar") {
            case "foo bar":
                var v15 float64 = ((0.1 - (3.14159 + 1.5)) - 1e10)
                _ = v15
            case "\n\t":
                if false {
                    var v16 uint64 = 2
                    p0 = -224
                    var v17 int = (p0 >> ((762) & 63))
                    _ = v16
                    _ = v17
                }
            default:
                var v18 int16 = 1
                _ = v18
            }
        } else {
            switch -102 {
            case 254:
                fmt.Println(p0)
            case 128:
                if (0 < 2) {
                    if (643 <= 32767) {
                        for fi19 := 0; fi19 < 9; fi19++ {
                            p0 = p0
                            switch "hello" {
                            case "":
                                p0 = fi19
                            }
                            if len(v14) > 0 {
                                v14[0] = 42
                            }
                            _ = fi19
                        }
                        v20 := GenS3{F0: -1, F1: -1}
                        if len(v14) > 0 {
                            v14[0] = p0
                        }
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 19; fi21++ {
                            p0 += (p0 / ((fi21) | 1))
                            _ = fi21
                        }
                    }
                } else {
                    fmt.Println(p0)
                }
            default:
                for fi22 := 0; fi22 < 2; fi22++ {
                    v23 := GenS4{F0: -1, F1: 373}
                    _ = fi22
                    _ = v23
                }
            }
            var v24 int64 = 0
            _ = v24
        }
    case 2147483647:
        p0 = (p0 << ((4294967295) & 63))
    case 269:
        var v25 uint64 = 2
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        p0 = -415
        p0 = (-(fi26))
        var v27 int8 = -1
        _ = fi26
        _ = v27
    }
    p0 = (p0 - 4294967295)
    _ = p0
    _ = v14
    return -32768
}

func main() {
    var v1 int8 = 53
    var v2 int = -1
    var v3 int8 = (v1 * (v1 & 127))
    v4 := []uint32{42, 2147483647, 42}
    v5 := make(map[uint32]bool)
    v3 += (^(v3))
    var v28 int16 = 753
    v3 = (v1 & v1)
    _ = len(v4)
    if (-2147483648 < -243) {
        v29 := GenS5{F0: int(v1), F1: "a", F2: -2147483648}
        if true {
            v30 := GenS3{F0: 0, F1: v28}
            _ = len(v5)
            for fi31 := 0; fi31 < 18; fi31++ {
                switch 9223372036854775807 {
                case -2147483648:
                    var v32 int = (v2 - v2)
                    _ = v32
                default:
                    switch v2 {
                    case 127:
                        var v33 float32 = ((float32(1.5) - 1.5) / 0.1)
                        _ = v33
                    case 128:
                        switch "\n\t" {
                        case "foo bar":
                            v34 := v5[2147483648]
                            _ = v34
                        }
                    case 46:
                        v35 := GenS6{F0: (v1 ^ (v1 * int8(fi31))), F1: 4294967295, F2: 4294967295}
                        _ = v35
                    default:
                        v36 := GenS7{F0: "hello"}
                        _ = v36
                    }
                }
                v1 += (v1 >> ((65535) & 63))
                _ = fi31
            }
            _ = v30
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v29
    } else {
        for fi37 := 0; fi37 < 2; fi37++ {
            for fi38 := 0; fi38 < 15; fi38++ {
                fmt.Println(v1)
                _ = fi38
            }
            var v39 uint64 = 0
            _ = fi37
            _ = v39
        }
        delete(v5, 243)
    }
    v40 := GenS1{F0: 673, F1: 928, F2: 1}
    var v41 uint64 = 665
    for fi42 := 0; fi42 < 3; fi42++ {
        if true {
            v5[42] = true
            var v43 uint32 = 252
            v44 := GenS4{F0: v28, F1: 42}
            _ = v43
            _ = v44
        } else {
            for fi45 := 0; fi45 < 12; fi45++ {
                fmt.Println(v41, v2, v41)
                if (true && true) {
                    if (1.5 >= 3.14159) {
                        switch (v2 * (-(fi45))) {
                        case 128:
                            var v46 uint8 = 181
                            _ = v46
                        case 0:
                            if false {
                                var v47 float64 = 3.14159
                                _ = v47
                            }
                        case -461:
                            if (((("" + ((("hello" + ("a" + "\n\t")) + "") + "")) + "\n\t") + "hello") < "\n\t") {
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                                var v48 uint32 = 65535
                                _ = v48
                            } else {
                                if (!(float32(float64(float32(0.0))) > float32(0.1))) {
                                    v49 := GenS7{F0: ("hello" + "\n\t")}
                                    v41 = v41
                                    v2 = fi45
                                    v2 = (^(v2))
                                    if (42 <= v41) {
                                        if (true || (!false)) {
                                            v41 = (v41 - (v41 | 252))
                                            v2 = (fi42 >> ((4294967295) & 63))
                                        }
                                        if (!false) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v50 := GenS5{F0: v2, F1: ((("foo bar" + "hello") + "") + ("\n\t" + "\n\t")), F2: 656}
                                            _ = len(v5)
                                            if true {
                                                var v51 bool = false
                                                if v51 {
                                                    if len(v4) > 0 {
                                                        v4[0] = 42
                                                    }
                                                    var v52 float64 = ((0.0 + (3.14159 * 1e10)) - 0.0)
                                                    fmt.Println(v51)
                                                    _ = v52
                                                } else {
                                                    fmt.Println(v2, v3)
                                                    switch (v2 / ((v2) | 1)) {
                                                    case 774:
                                                        fmt.Println(fi45, v51, v41)
                                                    case 32768:
                                                        var v53 int16 = (v28 + (v28 << ((65535) & 63)))
                                                        _ = v53
                                                    case 580:
                                                        v2 = fi45
                                                    default:
                                                        v54 := GenS5{F0: fi42, F1: "", F2: -1}
                                                        _ = v54
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v55 int32 = 2147483647
                                                        _ = v55
                                                    }
                                                }
                                                _ = v51
                                            } else {
                                                v56 := GenS6{F0: v3, F1: v41, F2: 42}
                                                v57 := GenS5{F0: fi45, F1: "", F2: -8}
                                                var v58 float64 = (0.0 + 3.14159)
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                            v28 = v28
                                            v59 := GenS5{F0: fi42, F1: "", F2: 0}
                                            var v60 uint16 = 819
                                            _ = v50
                                            _ = v59
                                            _ = v60
                                        } else {
                                            var v61 uint32 = 65535
                                            _ = v61
                                        }
                                        var v62 int16 = (-(v28))
                                        _ = v62
                                    } else {
                                        switch "" {
                                        case "a":
                                            v41 -= (v41 / ((v41) | 1))
                                        case "\n\t":
                                            var v63 float64 = (0.1 * float64(1.5))
                                            _ = v63
                                        case "":
                                            var v64 uint32 = 584
                                            _ = v64
                                        default:
                                            v4 = append(v4, 2147483647)
                                        }
                                    }
                                    _ = v49
                                }
                            }
                        default:
                            v65 := GenS6{F0: v1, F1: (v41 | v41), F2: 2}
                            _ = v65
                        }
                        var v66 string = ((("foo bar" + "") + "") + "\n\t")
                        var v67 uint64 = 0
                        var v68 uint8 = 2
                        _ = v66
                        _ = v67
                        _ = v68
                    } else {
                        fmt.Println(v28)
                    }
                    var v69 string = ("\n\t" + "a")
                    v41 = v41
                    var v70 int16 = (v28 - 0)
                    _ = v69
                    _ = v70
                }
                v71 := GenS2{F0: ("hello" + ("" + ((("foo bar" + "hello") + (("\n\t" + ("foo bar" + "\n\t")) + (("" + ("\n\t" + "hello")) + "hello"))) + "foo bar"))), F1: (v2 + (fi42 * (fi42 / (((^(fi42))) | 1))))}
                _ = fi45
                _ = v71
            }
            for fi72 := 0; fi72 < 15; fi72++ {
                v41 = (v41 >> ((245) & 63))
                if (!false) {
                    v2 = fi42
                    switch int(v28) {
                    case -173:
                        _ = len(v5)
                    case -2147483648:
                        if false {
                            var v73 uint = 255
                            _ = v73
                        } else {
                            v74 := GenS7{F0: (("foo bar" + "hello") + "\n\t")}
                            v28 += -284
                            _ = v74
                        }
                    case 32768:
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                    }
                    var v75 int32 = 2147483647
                    fmt.Println(v1, v28)
                    var v76 uint16 = 705
                    fmt.Println(v75)
                    _ = len(v5)
                    v3 = v1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v75
                    _ = v76
                } else {
                    var v77 uint64 = v41
                    var v78 bool = (true && false)
                    var v79 int16 = 39
                    _ = v77
                    _ = v78
                    _ = v79
                }
                var v80 uint = 255
                _ = fi72
                _ = v80
            }
            for fi81 := 0; fi81 < 18; fi81++ {
                v82 := GenS2{F0: "foo bar", F1: (fi42 | fi81)}
                _ = fi81
                _ = v82
            }
        }
        var v83 float64 = (0.0 / 1.0)
        v83 -= float64(float32(v83))
        _ = fi42
        _ = v83
    }
    for fi84 := 0; fi84 < 19; fi84++ {
        v85 := GenS8{F0: 65535, F1: 65535}
        if ((float64(1.5) >= 1.5) && (false || (true || ((true || true) && (0 > 2))))) {
            _ = len(v4)
        } else {
            switch "" {
            case "hello":
                for fi86 := 0; fi86 < 17; fi86++ {
                    v87 := GenS5{F0: (v2 << ((0) & 63)), F1: "\n\t", F2: 935}
                    if (740 <= 4294967295) {
                        v28 -= (v28 << ((255) & 63))
                    } else {
                        v28 -= 32767
                    }
                    fmt.Println(v3)
                    _ = fi86
                    _ = v87
                }
            default:
                for fi88 := 0; fi88 < 15; fi88++ {
                    switch (fi84 / (((fi84 % ((128) | 1))) | 1)) {
                    case 0:
                        if len(v4) > 0 {
                            v4[0] = 67
                        }
                    case 127:
                        switch (("a" + "foo bar") + "a") {
                        case "foo bar":
                            fmt.Println(fi88, v28, v3)
                        case "a":
                            var v89 bool = ((float32((-2.25 / (1.5 + float64(1e10)))) * -2.25) == 0.1)
                            _ = v89
                        }
                    case -996:
                        var v90 int64 = 9223372036854775807
                        _ = v90
                    default:
                        if (true && false) {
                            v28 = v28
                            v91 := GenS5{F0: (v2 % ((fi88) | 1)), F1: ("" + ""), F2: 569}
                            var v92 uint8 = 255
                            _ = v91
                            _ = v92
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    v41 += (v41 | v41)
                    switch 0 {
                    case 32768:
                        fmt.Println(fi88)
                    default:
                        switch 0 {
                        case 4294967295:
                            if false {
                                v2 = fi88
                            } else {
                                v93 := GenS6{F0: (v3 << ((2) & 63)), F1: v41, F2: 502}
                                _ = v93
                            }
                        case -236:
                            if true {
                                var v94 uint32 = 2
                                fmt.Println(v28, fi84)
                                _ = v94
                            }
                        case 42:
                            v95 := GenS2{F0: ("" + "\n\t"), F1: 32767}
                            _ = v95
                        }
                    }
                    _ = fi88
                }
            }
        }
        v1 += (v3 ^ v3)
        _ = fi84
        _ = v85
    }
    fmt.Println(v41, v41, v41)
    var v96 int64 = 9223372036854775807
    v41 = v41
    if false {
        fmt.Println(v2, v28)
        fmt.Println(v96)
        v97 := GenS1{F0: (v41 + v41), F1: v28, F2: (v41 % (((^(v41))) | 1))}
        _ = v97
    }
    var v98 int8 = v3
    v99 := GenS2{F0: "a", F1: (v2 & 366)}
    var v100 int32 = int32(v41)
    switch ("a" + ("foo bar" + "\n\t")) {
    case "foo bar":
        v100 = (v100 >> ((255) & 63))
    default:
        if false {
            v28 -= genFn0(v2)
        } else {
            v3 = (v98 >> ((2) & 63))
            for fi101 := 0; fi101 < 3; fi101++ {
                switch ("\n\t" + "foo bar") {
                case "hello":
                    var v102 uint16 = 42
                    _ = v102
                default:
                    for fi103 := 0; fi103 < 15; fi103++ {
                        v98 += (v1 / ((61) | 1))
                        var v104 int16 = genFn0(v2)
                        _ = fi103
                        _ = v104
                    }
                }
                if (("foo bar" + "hello") <= ((("foo bar" + ("hello" + "\n\t")) + ("" + "foo bar")) + "")) {
                    var v105 int16 = (v28 % ((genFn0(v2)) | 1))
                    v96 -= int64(v100)
                    for fi106 := 0; fi106 < 14; fi106++ {
                        v107 := GenS3{F0: (v96 / ((int64(v3)) | 1)), F1: 1000}
                        _ = fi106
                        _ = v107
                    }
                    _ = v105
                }
                if true {
                    fmt.Println(v28, v1, v41)
                    for fi108 := 0; fi108 < 19; fi108++ {
                        var v109 int64 = -1
                        v110 := GenS7{F0: ((("hello" + "a") + "a") + ("\n\t" + "a"))}
                        _ = fi108
                        _ = v109
                        _ = v110
                    }
                }
                _ = fi101
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v96:", v96)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v100
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v96
    _ = v98
    _ = v99
}
