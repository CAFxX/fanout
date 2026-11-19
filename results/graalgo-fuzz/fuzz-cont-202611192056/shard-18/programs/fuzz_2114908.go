package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int16; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64; F1 int8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64 }
func genFn0(p0 uint64) uint {
    if ((!(true || false)) || false) {
        v6 := []int{}
        fmt.Println(p0)
        _ = v6
    } else {
        v7 := GenS1{F0: 60, F1: -179, F2: 2147483648}
        p0 += p0
        _ = v7
    }
    fmt.Println(p0)
    p0 += p0
    p0 = p0
    _ = p0
    return 1
}
func genFn1(p0 uint16) uint32 {
    var v8 int32 = 1
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint64 = 0
        switch 9223372036854775807 {
        case 60:
            var v11 float32 = float32(1e10)
            _ = v11
        case 0:
            var v12 int32 = v8
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    if (0 <= 559) {
        switch 2147483647 {
        case -169:
            v13 := GenS2{F0: genFn0(4294967295), F1: v8}
            _ = v13
        default:
            var v14 int8 = 127
            _ = v14
        }
        v8 = (v8 * 42)
    }
    _ = p0
    _ = v8
    return 665
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    var v3 uint = 42
    var v4 float64 = float64(3.14159)
    var v5 uint16 = 26
    if (32767 != 1) {
        fmt.Println(v1, v3)
    }
    var v15 uint = 65535
    if false {
        v16 := GenS3{F0: false, F1: 737, F2: 2}
        _ = len(v2)
        _ = v16
    } else {
        v17 := GenS4{F0: v5, F1: 65535}
        _ = v17
    }
    var v18 int32 = (-(v1))
    var v19 int16 = 432
    fmt.Println(v1, v4)
    var v20 int = 140
    fmt.Println(v3)
    v21 := GenS5{F0: 18446744073709551615, F1: v19, F2: 42}
    for fi22 := 0; fi22 < 9; fi22++ {
        fmt.Println(v15)
        if false {
            if (false || false) {
                if true {
                    if ((v19 == v19) && ("foo bar" <= "hello")) {
                        v1 = v1
                    }
                    v15 -= 1
                } else {
                    var v23 int = v20
                    v1 += -2147483648
                    _ = v23
                }
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS6{F0: 27, F1: 2, F2: (v19 >> (((v3 << (((v3 % (((v3 | v15)) | 1))) & 63))) & 63))}
                    _ = fi24
                    _ = v25
                }
            } else {
                if (-9223372036854775808 < 42) {
                    v26 := GenS7{F0: genFn0(0)}
                    v27 := []int{(fi22 >> ((genFn0(uint64(v5))) & 63)), (v20 * fi22), (fi22 | (v20 * (fi22 * fi22)))}
                    _ = v26
                    _ = v27
                } else {
                    if false {
                        for fi28 := 0; fi28 < 2; fi28++ {
                            var v29 uint16 = 366
                            var v30 int32 = v1
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        v4 = v4
                    }
                    v31 := GenS8{F0: (-(v18)), F1: v3}
                    for fi32 := 0; fi32 < 13; fi32++ {
                        v3 += v3
                        v15 += 655
                        _ = fi32
                    }
                    _ = v31
                }
                fmt.Println(v20, fi22)
            }
            if true {
                v33 := []int{(fi22 / ((42) | 1)), (fi22 * (fi22 / (((v20 * -2147483648)) | 1)))}
                switch v20 {
                case 729:
                    switch (("" + "a") + ("\n\t" + (((("a" + "\n\t") + "a") + "a") + "a"))) {
                    case "hello":
                        v15 += (v3 + (v3 % ((v3) | 1)))
                    case "\n\t":
                        var v34 bool = true
                        _ = v34
                    default:
                        var v35 uint8 = 255
                        _ = v35
                    }
                default:
                    fmt.Println(v1)
                }
                _ = v33
            } else {
                v36 := GenS9{F0: 1, F1: 841}
                if false {
                    v37 := []int{fi22, v20, (fi22 << (((v15 * (v15 % (((^(v15))) | 1)))) & 63))}
                    _ = v37
                }
                _ = v36
            }
        }
        _ = fi22
    }
    switch "" {
    case "foo bar":
        var v38 bool = (!false)
        _ = v38
    case "":
        v39 := GenS7{F0: (^(v3))}
        _ = v39
    case "a":
        if true {
            var v40 uint = 4294967295
            if true {
                if false {
                    var v41 bool = (!false)
                    v4 = (v4 + v4)
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 13; fi42++ {
                        if true {
                            if true {
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    switch int(v18) {
                                    case -27:
                                        switch fi43 {
                                        case -570:
                                            if (!(true || true)) {
                                                switch "a" {
                                                case "hello":
                                                    v18 += (v18 * (v18 / ((-415) | 1)))
                                                case "":
                                                    switch int(v40) {
                                                    case 85:
                                                        v1 -= v18
                                                    case -764:
                                                        var v44 int8 = -1
                                                        _ = v44
                                                    }
                                                default:
                                                    v20 = (fi43 >> ((v15) & 63))
                                                    var v45 string = ""
                                                    v46 := GenS9{F0: v1, F1: 2}
                                                    _ = v45
                                                    _ = v46
                                                }
                                                v20 = fi42
                                            } else {
                                                v2[(v20 % (((-(fi42))) | 1))] = false
                                                if ((^(v5)) <= (v5 ^ v5)) {
                                                    v47 := GenS2{F0: genFn0(34), F1: (v1 - v1)}
                                                    if (false || true) {
                                                        if true {
                                                            var v48 int8 = 113
                                                            _ = v48
                                                        } else {
                                                            var v49 uint32 = 0
                                                            fmt.Println(v1)
                                                            _ = v49
                                                        }
                                                    }
                                                    v3 = (v3 + v40)
                                                    _ = v47
                                                } else {
                                                    v50 := GenS3{F0: false, F1: (v40 / (((^(v3))) | 1)), F2: 92}
                                                    var v51 int64 = 9223372036854775807
                                                    _ = v50
                                                    _ = v51
                                                }
                                                switch ("a" + ("foo bar" + "")) {
                                                case "":
                                                    var v52 uint32 = 42
                                                    _ = v52
                                                default:
                                                    v53 := []int{}
                                                    _ = v53
                                                }
                                                fmt.Println(v3, v1)
                                                fmt.Println(v5)
                                                var v54 float64 = (float64((3.14159 + 0.0)) / float64(3.14159))
                                                v54 -= (v4 + v4)
                                                var v55 int32 = (v18 >> ((genFn0(4294967295)) & 63))
                                                v1 = -2147483648
                                                if true {
                                                    if false {
                                                        var v56 int8 = -19
                                                        var v57 uint32 = 526
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        delete(v2, 4294967295)
                                                        _ = len(v2)
                                                        if (!true) {
                                                            v58 := []int{fi43, fi43, (fi43 | fi43), fi43}
                                                            var v59 int64 = -9223372036854775808
                                                            v20 -= -9223372036854775808
                                                            v60 := GenS7{F0: 42}
                                                            v55 = 226
                                                            if true {
                                                                var v61 uint16 = (v5 / (((v5 >> (((^(v40))) & 63))) | 1))
                                                                switch "a" {
                                                                case "a":
                                                                    v15 = (v3 & (v3 & v3))
                                                                case "\n\t":
                                                                    fmt.Println(v1, fi42, v54)
                                                                default:
                                                                    var v62 bool = true
                                                                    switch 128 {
                                                                    case 128:
                                                                        switch fi43 {
                                                                        case 0:
                                                                            var v63 int8 = -107
                                                                            _ = v63
                                                                        case 77:
                                                                            if (128 < 1) {
                                                                                v4 += (1e10 * 0.0)
                                                                                var v64 int8 = 8
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                    case -1:
                                                                        if v62 {
                                                                            v54 -= (((v4 * v54) * v4) / (1e10 * float64(0.1)))
                                                                            v3 += uint(v61)
                                                                        }
                                                                    case 9223372036854775807:
                                                                        v65 := GenS10{F0: (^(v20)), F1: 255}
                                                                        _ = v65
                                                                    default:
                                                                        v58 = append(v58, (fi42 & fi43))
                                                                    }
                                                                    if (v62 || v62) {
                                                                        v61 += (v5 % (((v61 - v61)) | 1))
                                                                        v15 = (v40 ^ (^(v40)))
                                                                        v5 = 241
                                                                    } else {
                                                                        v59 = int64(v5)
                                                                        var v66 int8 = 20
                                                                        _ = v66
                                                                    }
                                                                    _ = v62
                                                                }
                                                                _ = v61
                                                            }
                                                            v67 := GenS5{F0: 269, F1: v19, F2: v59}
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                            _ = v67
                                                        } else {
                                                            if true {
                                                                if true {
                                                                    var v68 uint = v3
                                                                    var v69 bool = true
                                                                    switch (v20 / (((v20 << (((^(v68))) & 63))) | 1)) {
                                                                    case 42:
                                                                        v1 = v1
                                                                    case 127:
                                                                        if (42 < 1) {
                                                                            v70 := []int{fi42, v20}
                                                                            _ = v70
                                                                        }
                                                                    case -9:
                                                                        v3 = v15
                                                                    }
                                                                    switch fi42 {
                                                                    case -1:
                                                                        if v69 {
                                                                            v71 := GenS7{F0: (v3 - 77)}
                                                                            var v72 int32 = (v55 | v1)
                                                                            fmt.Println(fi42, v54, fi43)
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                    case 890:
                                                                        var v73 uint64 = 2
                                                                        _ = v73
                                                                    case 184:
                                                                        fmt.Println(v3, fi43, fi42)
                                                                    default:
                                                                        if (v69 || false) {
                                                                            v74 := v2[fi42]
                                                                            v74 = v74
                                                                            v55 -= (v55 | v18)
                                                                            _ = v74
                                                                        }
                                                                        if (!v69) {
                                                                            v75 := GenS8{F0: v55, F1: 42}
                                                                            var v76 float64 = (v4 * (1e10 + v54))
                                                                            var v77 uint16 = 210
                                                                            _ = v75
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                    }
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                                var v78 int16 = v19
                                                                _ = v78
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v79 := []int{9223372036854775807}
                                                                _ = v79
                                                            case "hello":
                                                                v4 += (v54 - ((float64(0.1) / (v4 + float64(0.0))) + v4))
                                                            }
                                                            v54 = (1e10 / (v4 / (v54 * 1.5)))
                                                            v54 = 3.14159
                                                            v80 := GenS2{F0: 143, F1: (-(v55))}
                                                            var v81 uint32 = 2
                                                            fmt.Println(v15)
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        v54 += float64(1.5)
                                                        if (v54 < v54) {
                                                            switch (("\n\t" + ("" + "foo bar")) + "hello") {
                                                            case "":
                                                                switch int(v19) {
                                                                case -25:
                                                                    fmt.Println(v3, v40)
                                                                    var v82 bool = (fi43 <= fi42)
                                                                    var v83 int64 = 9223372036854775807
                                                                    _ = v82
                                                                    _ = v83
                                                                case 266:
                                                                    if ((((215 < 742) && true) || (true && true)) || true) {
                                                                        var v84 int8 = 42
                                                                        _ = v84
                                                                    } else {
                                                                        v85 := GenS11{F0: 114}
                                                                        var v86 uint = (v40 | v3)
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    v87 := []int{(fi42 / ((v20) | 1))}
                                                                    _ = v87
                                                                default:
                                                                    if (false && false) {
                                                                        if true {
                                                                            v19 -= v19
                                                                        }
                                                                        var v88 uint16 = (v5 & v5)
                                                                        if ((fi42 / ((fi42) | 1)) == (fi42 | fi42)) {
                                                                            v89 := GenS12{F0: 2, F1: 0}
                                                                            _ = v89
                                                                        }
                                                                        _ = v88
                                                                    } else {
                                                                        fmt.Println(v19, v5)
                                                                        v19 = int16(v5)
                                                                        v4 -= float64((3.14159 - 0.1))
                                                                        fmt.Println(fi43)
                                                                        if (true || true) {
                                                                            v55 = (v18 << (((v40 - (v3 / ((v40) | 1)))) & 63))
                                                                            var v90 bool = (false || ((169 < v15) && ("" > "\n\t")))
                                                                            _ = v90
                                                                        } else {
                                                                            var v91 bool = (1 > 4294967295)
                                                                            if ("\n\t" <= "a") {
                                                                                var v92 int32 = -2147483648
                                                                                _ = v92
                                                                            } else {
                                                                                if v91 {
                                                                                    if v91 {
                                                                                        v93 := []int{(fi42 >> ((genFn0(1)) & 63))}
                                                                                        switch 127 {
                                                                                        case -1:
                                                                                            switch (v20 << ((genFn0(18446744073709551615)) & 63)) {
                                                                                            case -478:
                                                                                                if len(v93) >= 2 {
                                                                                                    v93 = v93[1:len(v93)]
                                                                                                }
                                                                                            default:
                                                                                                v15 = (v40 + v40)
                                                                                            }
                                                                                        case 2147483647:
                                                                                            fmt.Println(v19, v5, v4)
                                                                                        case 372:
                                                                                            v3 += (v3 % ((v3) | 1))
                                                                                        }
                                                                                        var v94 uint16 = v5
                                                                                        var v95 int16 = (v19 ^ v19)
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            v2[fi42] = (-685 > (v95 % ((v95) | 1)))
                                                                                        case "a":
                                                                                            v91 = false
                                                                                        case "hello":
                                                                                            v19 = 0
                                                                                        }
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    } else {
                                                                                        v18 += -1
                                                                                    }
                                                                                    v5 = v5
                                                                                    v96 := GenS13{F0: v20}
                                                                                    _ = v96
                                                                                } else {
                                                                                    v4 = 0.0
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v97 int16 = (v19 / ((v19) | 1))
                                                                                var v98 int64 = -48
                                                                                if true {
                                                                                    fmt.Println(v55)
                                                                                    var v99 int64 = (v98 >> ((v40) & 63))
                                                                                    v1 = v18
                                                                                    _ = v99
                                                                                } else {
                                                                                    var v100 string = "hello"
                                                                                    _ = v100
                                                                                }
                                                                                var v101 uint8 = 255
                                                                                _ = v101
                                                                                _ = v97
                                                                                _ = v98
                                                                            } else {
                                                                                v15 += 0
                                                                            }
                                                                            v20 -= fi43
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v102 bool = true
                                                            if (!v102) {
                                                                v103 := GenS14{F0: fi43}
                                                                _ = v103
                                                            }
                                                            _ = v102
                                                        }
                                                    }
                                                }
                                                _ = v54
                                                _ = v55
                                            }
                                        case -940:
                                            v4 += -2.25
                                        case -1000:
                                            v5 -= v5
                                        default:
                                            var v104 int = (^(v20))
                                            _ = v104
                                        }
                                    case 4294967295:
                                        _ = len(v2)
                                    case 128:
                                        fmt.Println(v4, v5)
                                    default:
                                        if false {
                                            fmt.Println(fi42)
                                            delete(v2, fi42)
                                            v15 = 118
                                        }
                                    }
                                    v105 := []int{(fi42 % (((v20 * fi43)) | 1))}
                                    v40 += v15
                                    _ = fi43
                                    _ = v105
                                }
                            } else {
                                var v106 uint8 = 100
                                v2[(fi42 >> (((^(v40))) & 63))] = (!true)
                                v3 = genFn0(4294967295)
                                _ = v106
                            }
                            v5 = (v5 + (v5 | uint16(v3)))
                        }
                        for fi107 := 0; fi107 < 6; fi107++ {
                            v108 := GenS12{F0: 337, F1: 42}
                            _ = fi107
                            _ = v108
                        }
                        _ = fi42
                    }
                }
                for fi109 := 0; fi109 < 13; fi109++ {
                    v110 := []int{931, -786}
                    for fi111 := 0; fi111 < 14; fi111++ {
                        v40 += v3
                        _ = fi111
                    }
                    switch 42 {
                    case 128:
                        var v112 float64 = v4
                        _ = v112
                    }
                    _ = fi109
                    _ = v110
                }
            }
            _ = v40
        }
    default:
        v113 := GenS10{F0: (v20 >> (((v3 % (((^(v3))) | 1))) & 63)), F1: v5}
        _ = v113
    }
    if false {
        var v114 int = v20
        v15 = (v3 >> (((v3 / ((v3) | 1))) & 63))
        var v115 float64 = ((((v4 - float64(float32(v4))) * 1.5) - ((0.0 / v4) / 1e10)) / v4)
        _ = v114
        _ = v115
    } else {
        switch "hello" {
        case "\n\t":
            v116 := v2[(v20 & 32768)]
            _ = v116
        case "foo bar":
            for fi117 := 0; fi117 < 7; fi117++ {
                for fi118 := 0; fi118 < 1; fi118++ {
                    var v119 string = "hello"
                    var v120 uint32 = 2147483648
                    if (uint16(v15) < (v5 % ((65535) | 1))) {
                        v121 := GenS8{F0: 42, F1: v15}
                        switch "foo bar" {
                        case "\n\t":
                            var v122 uint64 = uint64(v5)
                            _ = v122
                        }
                        switch ((v119 + v119) + "\n\t") {
                        case "\n\t":
                            fmt.Println(v18)
                        case "a":
                            v20 -= (v20 & (fi117 % (((-(fi117))) | 1)))
                        default:
                            v123 := GenS2{F0: (^(v15)), F1: (v1 << (((v3 % ((v3) | 1))) & 63))}
                            _ = v123
                        }
                        _ = v121
                    } else {
                        _ = len(v2)
                        delete(v2, (-(fi118)))
                    }
                    _ = fi118
                    _ = v119
                    _ = v120
                }
                fmt.Println(v5, v18, v18)
                _ = fi117
            }
        default:
            if (!false) {
                for fi124 := 0; fi124 < 13; fi124++ {
                    switch fi124 {
                    case 127:
                        v1 -= v1
                    case 9223372036854775807:
                        if true {
                            v19 -= (v19 * (v19 >> (((v15 / ((v3) | 1))) & 63)))
                            v5 += (v5 & (v5 << ((genFn0(18446744073709551615)) & 63)))
                        }
                    }
                    _ = fi124
                }
                if false {
                    switch v20 {
                    case -716:
                        switch "foo bar" {
                        case "":
                            var v125 int = v20
                            _ = v125
                        }
                    }
                }
                var v126 float32 = -2.25
                _ = v126
            }
        }
        v127 := GenS5{F0: 325, F1: -1, F2: 0}
        _ = v127
    }
    v20 = v20
    v18 = v18
    if true {
        v4 += (1.5 - v4)
        v5 = (v5 & (v5 / (((v5 + (v5 | 42))) | 1)))
    }
    var v128 int64 = 1
    v129 := []int{32767, int(v1), (v20 | v20)}
    for fi130 := 0; fi130 < 12; fi130++ {
        v131 := GenS15{F0: 2}
        fmt.Println(fi130)
        var v132 uint = genFn0(1)
        _ = fi130
        _ = v131
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("slicelen v129:", len(v129), v129)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v128
    _ = v129
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
