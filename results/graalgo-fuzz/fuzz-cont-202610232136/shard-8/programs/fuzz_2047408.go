package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int32; F2 int32 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 uint16; F1 uint32; F2 bool }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int16; F1 int32; F2 int64 }
type GenS13 struct { F0 uint; F1 uint8; F2 bool }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint8; F1 int8 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint32; F1 int8 }
type GenS20 struct { F0 uint32; F1 string }
type GenS21 struct { F0 string; F1 int8; F2 uint }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 int8; F1 uint32 }
func genFn0(p0 int) int {
    v5 := []int{}
    p0 = p0
    fmt.Println(p0)
    v5 = append(v5, p0)
    _ = p0
    _ = v5
    return (-(p0))
}
func genFn1(p0 uint64, p1 uint8, p2 uint16) uint {
    p1 = uint8(p2)
    var v6 bool = (true || (127 == 29))
    p0 += 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 0.0
    var v3 int = 0
    var v4 float32 = (0.1 - v2)
    v4 += 0.1
    var v7 int8 = 0
    var v8 uint8 = 42
    switch v3 {
    case 9223372036854775807:
        var v9 uint64 = 42
        _ = v9
    case 1:
        v7 += (v7 >> ((2) & 63))
    }
    var v10 bool = (0 == 255)
    fmt.Println(v2, v8)
    if v10 {
        v10 = v10
    } else {
        if v10 {
            if (!(-9223372036854775808 != 301)) {
                switch v3 {
                case 128:
                    for fi11 := 0; fi11 < 4; fi11++ {
                        fmt.Println(v2)
                        v12 := []int{(v3 >> ((42) & 63))}
                        for fi13 := 0; fi13 < 18; fi13++ {
                            if v10 {
                                var v14 uint = 65535
                                switch "hello" {
                                case "":
                                    _ = len(v12)
                                case "a":
                                    v3 -= -1
                                default:
                                    v15 := GenS1{F0: 0, F1: 65535, F2: 42}
                                    _ = v15
                                }
                                _ = v14
                            } else {
                                v2 += float32(-2.25)
                            }
                            _ = fi13
                        }
                        _ = fi11
                        _ = v12
                    }
                default:
                    var v16 bool = true
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 5; fi17++ {
                    if true {
                        fmt.Println(fi17)
                        if v10 {
                            v18 := GenS1{F0: 0, F1: 0, F2: 371}
                            var v19 uint64 = 42
                            for fi20 := 0; fi20 < 18; fi20++ {
                                v21 := GenS2{F0: ("foo bar" + "foo bar")}
                                _ = fi20
                                _ = v21
                            }
                            _ = v18
                            _ = v19
                        } else {
                            for fi22 := 0; fi22 < 7; fi22++ {
                                v1[2] = ("hello" + "\n\t")
                                _ = fi22
                            }
                        }
                        v23 := []int{}
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v25 := GenS1{F0: -295, F1: 0, F2: -173}
                        if v10 {
                            switch "foo bar" {
                            case "hello":
                                v2 += float32(0.0)
                            case "foo bar":
                                var v26 uint64 = 2
                                _ = v26
                            case "a":
                                if true {
                                    if ((42 >= 0) || (!v10)) {
                                        var v27 uint8 = (v8 - v8)
                                        var v28 uint = 2
                                        var v29 int16 = -32768
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 int32 = 42
                                    fmt.Println(v8, v30)
                                    _ = v30
                                }
                            }
                        }
                        if true {
                            v31 := GenS1{F0: -834, F1: 255, F2: 0}
                            var v32 string = "hello"
                            v2 -= 3.14159
                            _ = v31
                            _ = v32
                        } else {
                            var v33 float64 = 0.1
                            if ((1.5 / (v2 - v4)) != v4) {
                                var v34 int = genFn0((v3 - genFn0(genFn0(fi17))))
                                v2 = (v4 / float32(1e10))
                                delete(v1, 0)
                                _ = v34
                            } else {
                                if false {
                                    v35 := GenS3{F0: 18446744073709551615, F1: -32768}
                                    _ = v35
                                } else {
                                    v3 = fi24
                                    var v36 string = ""
                                    _ = v36
                                }
                                v37 := GenS4{F0: 0}
                                fmt.Println(fi17, fi24, fi17)
                                _ = v37
                            }
                            _ = v33
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = fi17
                }
            }
            if false {
                v38 := GenS5{F0: (v3 << ((192) & 63)), F1: 729, F2: 2147483647}
                _ = v38
            } else {
                switch v3 {
                case 42:
                    switch "a" {
                    case "a":
                        if (2 > 42) {
                            for fi39 := 0; fi39 < 9; fi39++ {
                                for fi40 := 0; fi40 < 1; fi40++ {
                                    v8 = (v8 * (v8 | v8))
                                    if false {
                                        delete(v1, 472)
                                        var v41 uint16 = 255
                                        _ = v41
                                    }
                                    _ = fi40
                                }
                                v7 += v7
                                _ = fi39
                            }
                        } else {
                            for fi42 := 0; fi42 < 2; fi42++ {
                                v1[65535] = (("hello" + (("hello" + ("a" + ("" + "a"))) + ("a" + "\n\t"))) + (("foo bar" + "") + "\n\t"))
                                v1[2] = ("hello" + (("hello" + "") + "foo bar"))
                                v4 = float32(((-2.25 - -2.25) + 0.1))
                                _ = fi42
                            }
                            v43 := v1[2147483647]
                            switch (v3 + (v3 % (((v3 >> ((genFn1(945, (v8 >> ((1) & 63)), 258)) & 63))) | 1))) {
                            case -889:
                                for fi44 := 0; fi44 < 7; fi44++ {
                                    if (736 != 93) {
                                        v1[42] = ""
                                        var v45 int16 = 252
                                        v45 -= (v45 >> ((1) & 63))
                                        _ = v45
                                    }
                                    for fi46 := 0; fi46 < 14; fi46++ {
                                        switch "foo bar" {
                                        case "a":
                                            var v47 int = -9223372036854775808
                                            v48 := GenS3{F0: 647, F1: 32767}
                                            fmt.Println(v8, v3, v8)
                                            var v49 int64 = 0
                                            if v10 {
                                                v50 := GenS6{F0: 18446744073709551615, F1: 440}
                                                v51 := []int{fi44}
                                                v4 = (v4 - v2)
                                                switch (fi46 / (((v3 & (^(fi44)))) | 1)) {
                                                case -1:
                                                    v4 = v2
                                                case 2147483647:
                                                    fmt.Println(v10, v43)
                                                case -220:
                                                    var v52 int = 2147483647
                                                    _ = v52
                                                default:
                                                    if ((0 >= (v49 ^ (-(v49)))) || true) {
                                                        v53 := v1[843]
                                                        if v10 {
                                                            switch v47 {
                                                            case 246:
                                                                var v54 uint8 = 42
                                                                _ = v54
                                                            }
                                                            if (127 >= (v8 >> ((945) & 63))) {
                                                                v53 = ("a" + "foo bar")
                                                                v2 = 0.0
                                                                v8 = v8
                                                            }
                                                            if false {
                                                                v53 = ("hello" + ("foo bar" + v43))
                                                            } else {
                                                                fmt.Println(v2, v10, v53)
                                                                var v55 string = (("foo bar" + "\n\t") + v43)
                                                                switch (v3 + 42) {
                                                                case -2147483648:
                                                                    v56 := GenS1{F0: int32(v8), F1: 180, F2: v49}
                                                                    _ = v56
                                                                case -377:
                                                                    v49 = (v49 | v49)
                                                                    if (1000 != 32767) {
                                                                        var v57 int16 = 722
                                                                        _ = v57
                                                                    }
                                                                default:
                                                                    switch (v3 >> ((42) & 63)) {
                                                                    case -2147483648:
                                                                        var v58 int = v47
                                                                        _ = v58
                                                                    case 934:
                                                                        var v59 string = v55
                                                                        _ = v59
                                                                    case -552:
                                                                        if ((!(("hello" >= v43) && (77 < 0))) && (2 <= 574)) {
                                                                            fmt.Println(v55)
                                                                            v43 = v53
                                                                            var v60 int8 = v7
                                                                            _ = v60
                                                                        } else {
                                                                            v61 := GenS3{F0: 18446744073709551615, F1: 42}
                                                                            v7 += v7
                                                                            switch (v3 >> ((783) & 63)) {
                                                                            case -1:
                                                                                v10 = (false && false)
                                                                            case -406:
                                                                                fmt.Println(v49, v4, fi46)
                                                                                switch v47 {
                                                                                case 127:
                                                                                    if (v49 == (v49 % (((v49 * v49)) | 1))) {
                                                                                        if false {
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                v62 := GenS7{F0: v49, F1: true}
                                                                                                var v63 bool = v10
                                                                                                _ = v62
                                                                                                _ = v63
                                                                                            }
                                                                                            v64 := GenS5{F0: (v47 + (fi44 << ((124) & 63))), F1: -2147483648, F2: -2147483648}
                                                                                            var v65 uint32 = 2147483647
                                                                                            _ = v64
                                                                                            _ = v65
                                                                                        }
                                                                                        var v66 uint64 = 18446744073709551615
                                                                                        v67 := GenS3{F0: v66, F1: 1000}
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                case 1:
                                                                                    v68 := GenS8{F0: 2, F1: 211, F2: v10}
                                                                                    _ = v68
                                                                                case 32767:
                                                                                    if v10 {
                                                                                        v2 = v4
                                                                                        if v10 {
                                                                                            if false {
                                                                                                switch (-(fi46)) {
                                                                                                case -671:
                                                                                                    fmt.Println(v49, v47, fi46)
                                                                                                default:
                                                                                                    var v69 int = (fi46 + fi44)
                                                                                                    _ = v69
                                                                                                }
                                                                                                if len(v51) >= 2 {
                                                                                                    v51 = v51[1:len(v51)]
                                                                                                }
                                                                                            }
                                                                                            var v70 uint8 = v8
                                                                                            v7 = v7
                                                                                            _ = v70
                                                                                        } else {
                                                                                            v3 -= fi44
                                                                                            v49 = (v49 & -440)
                                                                                            v4 -= (0.0 / (0.0 / ((0.1 * (v2 - 1.5)) * float32(0.1))))
                                                                                        }
                                                                                        if v10 {
                                                                                            var v71 int = (v47 & (v47 / ((fi46) | 1)))
                                                                                            _ = v71
                                                                                        } else {
                                                                                            if v10 {
                                                                                                switch fi46 {
                                                                                                case -491:
                                                                                                    if true {
                                                                                                        var v72 int = int(v7)
                                                                                                        v72 -= -704
                                                                                                        v72 = (v3 - 32768)
                                                                                                        _ = v72
                                                                                                    }
                                                                                                case 127:
                                                                                                    var v73 float32 = float32(1e10)
                                                                                                    _ = v73
                                                                                                case 912:
                                                                                                    if (v2 < float32(0.1)) {
                                                                                                        if (!true) {
                                                                                                            switch v43 {
                                                                                                            case "a":
                                                                                                                delete(v1, 2)
                                                                                                            default:
                                                                                                                var v74 int32 = 286
                                                                                                                _ = v74
                                                                                                            }
                                                                                                            fmt.Println(v49, v4)
                                                                                                        }
                                                                                                    } else {
                                                                                                        v47 += fi44
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 43:
                                                                                v75 := GenS1{F0: 0, F1: 65535, F2: (^(v49))}
                                                                                var v76 int16 = 32767
                                                                                if (true && v10) {
                                                                                    switch ("" + ("" + v55)) {
                                                                                    case "foo bar":
                                                                                        fmt.Println(fi46)
                                                                                    }
                                                                                }
                                                                                _ = v75
                                                                                _ = v76
                                                                            }
                                                                            v8 -= (v8 % (((v8 >> ((440) & 63))) | 1))
                                                                            if v10 {
                                                                                var v77 int = 344
                                                                                if len(v51) > 0 {
                                                                                    v51[0] = (v47 & v77)
                                                                                }
                                                                                if (783 > 946) {
                                                                                    v78 := GenS9{F0: false, F1: v10}
                                                                                    var v79 int8 = (v7 / ((-1) | 1))
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                }
                                                                                fmt.Println(v4)
                                                                                var v80 uint8 = (^(v8))
                                                                                var v81 int32 = 42
                                                                                if ((true && (!v10)) && v10) {
                                                                                    fmt.Println(v49, fi46)
                                                                                    fmt.Println(v77, v55, v80)
                                                                                } else {
                                                                                    var v82 string = "\n\t"
                                                                                    var v83 float32 = v2
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                }
                                                                                if v10 {
                                                                                    if false {
                                                                                        var v84 uint8 = v80
                                                                                        if v10 {
                                                                                            if len(v51) >= 2 {
                                                                                                v51 = v51[1:len(v51)]
                                                                                            }
                                                                                            v85 := GenS5{F0: v77, F1: -1, F2: 2147483647}
                                                                                            _ = v85
                                                                                        } else {
                                                                                            if v10 {
                                                                                                v86 := v1[2147483647]
                                                                                                fmt.Println(fi46)
                                                                                                v49 = (v49 | 707)
                                                                                                _ = v86
                                                                                            }
                                                                                            switch genFn0(-1) {
                                                                                            case 27:
                                                                                                v55 = (v55 + "a")
                                                                                            }
                                                                                            v4 = (v4 - 0.0)
                                                                                            v80 += (v8 * (v8 - v8))
                                                                                            if v10 {
                                                                                                if v10 {
                                                                                                    switch ((((v43 + ("\n\t" + v43)) + "foo bar") + "hello") + ((v43 + "foo bar") + "a")) {
                                                                                                    case "foo bar":
                                                                                                        switch genFn0(genFn0((v47 + (-(fi46))))) {
                                                                                                        case -1:
                                                                                                            if (-105 == v7) {
                                                                                                                var v87 uint = 18446744073709551615
                                                                                                                _ = v87
                                                                                                            }
                                                                                                        case 9223372036854775807:
                                                                                                            var v88 float32 = 1.5
                                                                                                            _ = v88
                                                                                                        }
                                                                                                    }
                                                                                                    var v89 uint64 = 42
                                                                                                    _ = v89
                                                                                                }
                                                                                                v10 = (((fi46 >> ((42) & 63)) > (-(v47))) || ((1 >= 32767) || v10))
                                                                                                v80 -= v84
                                                                                            } else {
                                                                                                v90 := GenS8{F0: 788, F1: 1, F2: false}
                                                                                                v91 := v1[343]
                                                                                                if v10 {
                                                                                                    var v92 int16 = 502
                                                                                                    _ = v92
                                                                                                }
                                                                                                _ = v90
                                                                                                _ = v91
                                                                                            }
                                                                                            v47 -= v3
                                                                                        }
                                                                                        _ = v84
                                                                                    } else {
                                                                                        var v93 string = v53
                                                                                        if (65535 < 8) {
                                                                                            v2 = (v4 / (v4 / ((v4 + v4) * v4)))
                                                                                            v94 := v1[2147483647]
                                                                                            _ = v94
                                                                                        } else {
                                                                                            if v10 {
                                                                                                if (549 != 255) {
                                                                                                    v95 := GenS10{F0: 1}
                                                                                                    _ = v95
                                                                                                }
                                                                                            } else {
                                                                                                var v96 uint = 210
                                                                                                v97 := GenS11{F0: 2147483648}
                                                                                                if len(v51) >= 2 {
                                                                                                    v51 = v51[1:len(v51)]
                                                                                                }
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                            }
                                                                                            v3 = (fi44 | genFn0(fi44))
                                                                                            v98 := GenS2{F0: v55}
                                                                                            _ = v98
                                                                                        }
                                                                                        fmt.Println(v55, v55, v8)
                                                                                        _ = v93
                                                                                    }
                                                                                    switch v53 {
                                                                                    case "":
                                                                                        v53 = (v53 + v43)
                                                                                    default:
                                                                                        if (v2 >= 1.5) {
                                                                                            v53 = v55
                                                                                            if (!v10) {
                                                                                                var v99 float32 = (float32(0.0) / float32(0.1))
                                                                                                v77 = v3
                                                                                                var v100 uint8 = v8
                                                                                                var v101 uint64 = 2
                                                                                                switch (-(fi44)) {
                                                                                                case 128:
                                                                                                    v102 := GenS11{F0: 65535}
                                                                                                    _ = v102
                                                                                                case 32768:
                                                                                                    var v103 uint16 = 2
                                                                                                    _ = v103
                                                                                                }
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                                _ = v99
                                                                                            }
                                                                                            if false {
                                                                                                v104 := GenS12{F0: 42, F1: v81, F2: (v49 | (v49 >> ((909) & 63)))}
                                                                                                if ((v10 || (57 >= 42)) || v10) {
                                                                                                    v105 := GenS13{F0: 255, F1: 42, F2: true}
                                                                                                    _ = v105
                                                                                                }
                                                                                                var v106 bool = v10
                                                                                                _ = v104
                                                                                                _ = v106
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v107 int8 = v7
                                                                                    v55 = ("a" + (v53 + ("hello" + v53)))
                                                                                    _ = v107
                                                                                }
                                                                                var v108 uint8 = v8
                                                                                _ = v108
                                                                                _ = v77
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                switch v43 {
                                                                                case "foo bar":
                                                                                    v109 := GenS13{F0: 2, F1: 132, F2: false}
                                                                                    _ = v109
                                                                                }
                                                                                switch ((v55 + ("a" + v53)) + v55) {
                                                                                case "foo bar":
                                                                                    fmt.Println(v53, v8)
                                                                                case "\n\t":
                                                                                    var v110 int32 = 2147483647
                                                                                    var v111 bool = (1 >= 718)
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                                if (365 < v49) {
                                                                                    var v112 uint = 4294967295
                                                                                    v53 = "hello"
                                                                                    _ = v112
                                                                                } else {
                                                                                    v113 := GenS14{F0: 2147483648}
                                                                                    _ = v113
                                                                                }
                                                                                v114 := v1[65535]
                                                                                if (!((1.5 >= 3.14159) && (!v10))) {
                                                                                    fmt.Println(v43, v43)
                                                                                }
                                                                                _ = v114
                                                                            }
                                                                            _ = v61
                                                                        }
                                                                    default:
                                                                        switch v43 {
                                                                        case "foo bar":
                                                                            fmt.Println(v4, v10)
                                                                        default:
                                                                            v4 = 0.0
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(fi46)
                                                                _ = v55
                                                            }
                                                        } else {
                                                            v115 := GenS15{F0: (v8 * 1), F1: 66}
                                                            v8 += 128
                                                            v1[65535] = (("foo bar" + ("" + "a")) + v43)
                                                            _ = v115
                                                        }
                                                        if len(v51) >= 2 {
                                                            v51 = v51[1:len(v51)]
                                                        }
                                                        var v116 int = v47
                                                        v2 = v4
                                                        _ = v116
                                                        _ = v53
                                                    }
                                                    if (((v8 ^ v8) <= (^(v8))) && v10) {
                                                        if ((v8 >> ((genFn1(42, (v8 - (^(v8))), 65535)) & 63)) != (v8 & v8)) {
                                                            var v117 uint8 = v8
                                                            v47 = (v3 - (^(v3)))
                                                            var v118 uint16 = 65535
                                                            _ = v117
                                                            _ = v118
                                                        }
                                                        var v119 int32 = 1
                                                        var v120 string = (v43 + v43)
                                                        _ = v119
                                                        _ = v120
                                                    } else {
                                                        v121 := GenS5{F0: -346, F1: -494, F2: 339}
                                                        v43 = ((v43 + (v43 + v43)) + ((v43 + (v43 + ("\n\t" + v43))) + v43))
                                                        _ = v121
                                                    }
                                                }
                                                _ = v50
                                                _ = v51
                                            }
                                            if v10 {
                                                v49 = (v49 & (v49 & 9223372036854775807))
                                            }
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        case "foo bar":
                                            switch v43 {
                                            case "a":
                                                delete(v1, 2147483647)
                                            }
                                        default:
                                            v4 += -2.25
                                        }
                                        fmt.Println(v8, v2)
                                        _ = fi46
                                    }
                                    for fi122 := 0; fi122 < 12; fi122++ {
                                        v123 := GenS1{F0: 649, F1: 1, F2: 9223372036854775807}
                                        if (v10 && v10) {
                                            if (0 != -9223372036854775808) {
                                                fmt.Println(v2, v4, v7)
                                                switch (fi44 - (^(fi122))) {
                                                case -174:
                                                    v124 := GenS11{F0: 2}
                                                    _ = v124
                                                case 32768:
                                                    fmt.Println(v4)
                                                case 127:
                                                    v8 = v8
                                                default:
                                                    if (12 == v8) {
                                                        v10 = false
                                                        switch "\n\t" {
                                                        case "":
                                                            var v125 int32 = 1
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v126 uint = 4294967295
                                                                _ = v126
                                                            case "":
                                                                fmt.Println(v3, v7, v10)
                                                                v43 = v43
                                                            case "hello":
                                                                v43 = (v43 + v43)
                                                            }
                                                            v127 := []int{(v3 | fi122), int(v7), (^(fi44)), int(v7)}
                                                            fmt.Println(fi44, v2, fi44)
                                                            _ = v125
                                                            _ = v127
                                                        case "\n\t":
                                                            switch (fi44 ^ fi44) {
                                                            case -2147483648:
                                                                var v128 uint8 = v8
                                                                _ = v128
                                                            case 9223372036854775807:
                                                                var v129 bool = v10
                                                                _ = v129
                                                            case -542:
                                                                switch (v3 << ((489) & 63)) {
                                                                case 41:
                                                                    if (!v10) {
                                                                        v130 := []int{(-(fi122)), (fi44 % ((fi122) | 1))}
                                                                        var v131 float32 = (v4 - (v2 / 1.0))
                                                                        switch fi44 {
                                                                        case 127:
                                                                            switch -9223372036854775808 {
                                                                            case 9223372036854775807:
                                                                                v7 = v7
                                                                            case -807:
                                                                                fmt.Println(v7, v43)
                                                                            }
                                                                        default:
                                                                            var v132 int64 = 359
                                                                            _ = v132
                                                                        }
                                                                        _ = v130
                                                                        _ = v131
                                                                    } else {
                                                                        var v133 float32 = v4
                                                                        _ = v133
                                                                    }
                                                                case 128:
                                                                    v134 := []int{fi44, (fi122 | (v3 % (((^(fi122))) | 1))), fi122, fi122}
                                                                    _ = v134
                                                                case 2147483647:
                                                                    switch fi122 {
                                                                    case 1:
                                                                        fmt.Println(v7)
                                                                    case -9223372036854775808:
                                                                        var v135 float64 = 1e10
                                                                        _ = v135
                                                                    case 0:
                                                                        if (!v10) {
                                                                            if true {
                                                                                var v136 float64 = (0.0 * (3.14159 - -2.25))
                                                                                if v10 {
                                                                                    v137 := []int{fi122, genFn0(fi122), genFn0(fi122), v3}
                                                                                    v138 := GenS12{F0: 0, F1: 2147483647, F2: 42}
                                                                                    _ = v137
                                                                                    _ = v138
                                                                                }
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    fmt.Println(v136)
                                                                                case "":
                                                                                    v139 := []int{genFn0(v3)}
                                                                                    _ = v139
                                                                                }
                                                                                _ = v136
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch v43 {
                                                                        case "foo bar":
                                                                            v2 -= (v4 * (float32(0.1) * v4))
                                                                        case "\n\t":
                                                                            v140 := []int{v3}
                                                                            _ = v140
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v7 += 0
                                                    } else {
                                                        fmt.Println(v4)
                                                        if (556 < 1) {
                                                            var v141 int64 = 576
                                                            v142 := GenS6{F0: 1, F1: 0}
                                                            _ = v141
                                                            _ = v142
                                                        } else {
                                                            var v143 float32 = v2
                                                            v144 := []int{}
                                                            _ = v143
                                                            _ = v144
                                                        }
                                                    }
                                                    if (-2.25 != 1.5) {
                                                        var v145 int64 = 1
                                                        v146 := GenS3{F0: 4294967295, F1: 32767}
                                                        _ = v145
                                                        _ = v146
                                                    } else {
                                                        if v10 {
                                                            if (1 == 4294967295) {
                                                                v147 := []int{}
                                                                _ = v147
                                                            } else {
                                                                v10 = v10
                                                                switch genFn0(genFn0((fi44 & genFn0(fi122)))) {
                                                                case 32767:
                                                                    if ((926 >= 690) || (773 >= 355)) {
                                                                        var v148 uint32 = 4294967295
                                                                        v149 := []int{}
                                                                        if false {
                                                                            var v150 bool = (127 <= (v8 + v8))
                                                                            var v151 int8 = int8(v8)
                                                                            v152 := GenS12{F0: 42, F1: 2147483647, F2: 1}
                                                                            var v153 float32 = 0.0
                                                                            _ = v150
                                                                            _ = v151
                                                                            _ = v152
                                                                            _ = v153
                                                                        }
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                case 127:
                                                                    v154 := []int{genFn0((fi122 & 32767)), (v3 + fi122), (^(fi44)), (-(fi122))}
                                                                    _ = v154
                                                                case 128:
                                                                    v155 := GenS16{F0: 65535}
                                                                    _ = v155
                                                                default:
                                                                    v156 := GenS10{F0: -388}
                                                                    _ = v156
                                                                }
                                                                var v157 uint16 = 953
                                                                var v158 int16 = 1000
                                                                fmt.Println(v4)
                                                                _ = v157
                                                                _ = v158
                                                            }
                                                            switch fi122 {
                                                            case 127:
                                                                v8 -= (v8 >> ((4294967295) & 63))
                                                            }
                                                            v43 = v43
                                                        } else {
                                                            v159 := []int{(v3 - (v3 % ((v3) | 1))), fi122}
                                                            var v160 uint32 = 396
                                                            _ = v159
                                                            _ = v160
                                                        }
                                                        v2 -= (v4 / v4)
                                                    }
                                                }
                                                var v161 float32 = 1e10
                                                fmt.Println(v3, v4, v7)
                                                v161 += (float32(0.0) - v2)
                                                _ = v161
                                            } else {
                                                v162 := GenS2{F0: (v43 + (v43 + v43))}
                                                _ = v162
                                            }
                                            _ = len(v1)
                                            v163 := GenS12{F0: 185, F1: -2147483648, F2: -9223372036854775808}
                                            _ = v163
                                        }
                                        v164 := []int{(fi122 << ((genFn1(934, 104, 844)) & 63)), (fi44 - (v3 + genFn0((^(fi44)))))}
                                        _ = fi122
                                        _ = v123
                                        _ = v164
                                    }
                                    _ = fi44
                                }
                            default:
                                for fi165 := 0; fi165 < 18; fi165++ {
                                    switch int(v8) {
                                    case -9223372036854775808:
                                        for fi166 := 0; fi166 < 7; fi166++ {
                                            v3 = 127
                                            v167 := GenS17{F0: 2, F1: 42, F2: 2}
                                            _ = fi166
                                            _ = v167
                                        }
                                    case 32767:
                                        var v168 uint8 = v8
                                        _ = v168
                                    case 32768:
                                        if (false || v10) {
                                            var v169 int8 = 99
                                            fmt.Println(v7, v3, v2)
                                            _ = v169
                                        }
                                    }
                                    switch (^(v3)) {
                                    case 32767:
                                        for fi170 := 0; fi170 < 8; fi170++ {
                                            if v10 {
                                                if (uint32(v8) < 921) {
                                                    v171 := GenS18{F0: v43}
                                                    switch v43 {
                                                    case "":
                                                        v43 = v43
                                                    default:
                                                        v10 = ("" != "\n\t")
                                                    }
                                                    _ = v171
                                                }
                                            } else {
                                                fmt.Println(v10, fi165)
                                                if (-1 == 1) {
                                                    v172 := []int{}
                                                    fmt.Println(v43, fi170, v8)
                                                    v10 = true
                                                    _ = len(v172)
                                                    if len(v172) >= 2 {
                                                        v172 = v172[1:len(v172)]
                                                    }
                                                    var v173 int64 = -123
                                                    _ = v172
                                                    _ = v173
                                                } else {
                                                    switch (v43 + "\n\t") {
                                                    case "hello":
                                                        v174 := []int{(fi165 << ((318) & 63)), (fi165 << ((2) & 63))}
                                                        _ = v174
                                                    default:
                                                        if v10 {
                                                            v175 := GenS19{F0: 4294967295, F1: (v7 % ((v7) | 1))}
                                                            var v176 int8 = (v7 * (v7 + v7))
                                                            _ = v175
                                                            _ = v176
                                                        } else {
                                                            if (42 > -1) {
                                                                fmt.Println(v7)
                                                                switch v43 {
                                                                case "foo bar":
                                                                    var v177 float64 = (0.0 / 1e10)
                                                                    v43 = v43
                                                                    _ = v177
                                                                case "a":
                                                                    v4 -= (((0.1 * v2) + (((3.14159 + 0.0) / v2) * (v2 - v2))) * (((v4 * v2) - v2) / v4))
                                                                case "hello":
                                                                    fmt.Println(v4, v7)
                                                                }
                                                                if false {
                                                                    fmt.Println(v7, v7, v4)
                                                                    var v178 float64 = 0.0
                                                                    _ = v178
                                                                } else {
                                                                    v3 -= (fi165 % ((-9223372036854775808) | 1))
                                                                    v3 += (fi165 + (fi170 * fi165))
                                                                }
                                                            } else {
                                                                v1[65535] = v43
                                                            }
                                                            v8 -= (v8 + v8)
                                                            v179 := GenS17{F0: 0, F1: 0, F2: uint16(v8)}
                                                            _ = v179
                                                        }
                                                    }
                                                    delete(v1, 42)
                                                    v180 := GenS20{F0: 4294967295, F1: ""}
                                                    _ = v180
                                                }
                                            }
                                            var v181 bool = ((((((v2 + v4) + (-2.25 + float32((float64(v2) * 3.14159)))) * 0.0) / (v4 + 3.14159)) / v4) >= v2)
                                            switch ((v43 + v43) + v43) {
                                            case "foo bar":
                                                if true {
                                                    var v182 uint16 = 255
                                                    _ = v182
                                                }
                                            }
                                            v183 := GenS21{F0: v43, F1: (v7 - (v7 + (v7 | 0))), F2: genFn1(674, (v8 - v8), 32767)}
                                            v10 = true
                                            _ = fi170
                                            _ = v181
                                            _ = v183
                                        }
                                    }
                                    switch v43 {
                                    case "":
                                        var v184 int16 = 32767
                                        _ = v184
                                    case "hello":
                                        if v10 {
                                            v185 := []int{-225}
                                            if (42 > uint(v8)) {
                                                var v186 uint8 = 128
                                                _ = v186
                                            } else {
                                                v2 = ((v4 * float32(0.1)) / -2.25)
                                                if (!(-831 < -2147483648)) {
                                                    fmt.Println(v10, v4)
                                                }
                                            }
                                            _ = v185
                                        } else {
                                            var v187 float64 = (float64(float32((0.0 - 1.5))) * (3.14159 / -2.25))
                                            _ = v187
                                        }
                                    case "\n\t":
                                        v188 := []int{}
                                        _ = v188
                                    default:
                                        if v10 {
                                            v2 = v4
                                            var v189 uint8 = 177
                                            _ = v189
                                        } else {
                                            v43 = ((v43 + "hello") + v43)
                                        }
                                    }
                                    _ = fi165
                                }
                            }
                            _ = v43
                        }
                    case "\n\t":
                        var v190 int32 = 1
                        _ = v190
                    case "":
                        v191 := []int{32767, (^(v3)), v3}
                        _ = v191
                    default:
                        if true {
                            v192 := GenS16{F0: 524}
                            switch genFn0(-474) {
                            case 9223372036854775807:
                                v3 = v3
                            case -9223372036854775808:
                                if true {
                                    for fi193 := 0; fi193 < 15; fi193++ {
                                        v194 := []int{fi193, fi193}
                                        for fi195 := 0; fi195 < 8; fi195++ {
                                            v8 = v8
                                            v194 = append(v194, 4294967295)
                                            var v196 bool = ((false || v10) && (!v10))
                                            if len(v194) >= 2 {
                                                v194 = v194[1:len(v194)]
                                            }
                                            _ = fi195
                                            _ = v196
                                        }
                                        switch -518 {
                                        case 0:
                                            var v197 uint16 = 708
                                            _ = v197
                                        case 32768:
                                            if v10 {
                                                v1[2147483647] = ("foo bar" + ("hello" + "a"))
                                            } else {
                                                var v198 bool = ((!(1 >= -9223372036854775808)) || (v3 >= genFn0(genFn0(-905))))
                                                if true {
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v8, v10)
                                                    }
                                                    if v10 {
                                                        fmt.Println(v7, v3)
                                                    }
                                                    var v199 int = v3
                                                    _ = v199
                                                } else {
                                                    for fi200 := 0; fi200 < 3; fi200++ {
                                                        _ = len(v1)
                                                        var v201 int8 = (v7 * v7)
                                                        v202 := GenS22{F0: 18446744073709551615}
                                                        _ = fi200
                                                        _ = v201
                                                        _ = v202
                                                    }
                                                }
                                                var v203 float32 = v4
                                                _ = v198
                                                _ = v203
                                            }
                                        default:
                                            for fi204 := 0; fi204 < 7; fi204++ {
                                                v7 = v7
                                                v205 := GenS17{F0: 300, F1: 0, F2: 42}
                                                _ = fi204
                                                _ = v205
                                            }
                                        }
                                        _ = fi193
                                        _ = v194
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v206 float64 = 1.5
                                        _ = v206
                                    default:
                                        v1[42] = ""
                                    }
                                }
                            default:
                                fmt.Println(v2, v8, v3)
                            }
                            _ = v192
                        }
                    }
                case 1:
                    v207 := GenS8{F0: 612, F1: 2, F2: (0 <= 138)}
                    _ = v207
                default:
                    for fi208 := 0; fi208 < 5; fi208++ {
                        v3 -= v3
                        switch "\n\t" {
                        case "hello":
                            v10 = (2147483648 > 710)
                        case "\n\t":
                            fmt.Println(fi208)
                        case "foo bar":
                            fmt.Println(v7)
                        default:
                            fmt.Println(v3, v4)
                        }
                        if v10 {
                            switch "a" {
                            case "":
                                fmt.Println(v2, v7)
                            default:
                                v209 := []int{128}
                                _ = v209
                            }
                        }
                        _ = fi208
                    }
                }
                var v210 int32 = 1
                delete(v1, 380)
                _ = v210
            }
            v211 := GenS12{F0: 951, F1: 1, F2: 9223372036854775807}
            _ = v211
        }
        if ((v7 & v7) <= v7) {
            v2 -= v2
            var v212 float64 = 0.1
            _ = v212
        }
        if (18446744073709551615 != 2) {
            if (v10 && (-1 == 32767)) {
                var v213 uint16 = 65535
                var v214 uint64 = 18446744073709551615
                if (2147483648 != 42) {
                    for fi215 := 0; fi215 < 7; fi215++ {
                        fmt.Println(v2, v213)
                        if (false || true) {
                            for fi216 := 0; fi216 < 10; fi216++ {
                                var v217 int64 = -9223372036854775808
                                if v10 {
                                    v1[42] = ("a" + "hello")
                                    v218 := GenS22{F0: 84}
                                    _ = v218
                                }
                                _ = fi216
                                _ = v217
                            }
                        }
                        var v219 int8 = 127
                        _ = fi215
                        _ = v219
                    }
                } else {
                    fmt.Println(v2, v2)
                    v220 := GenS23{F0: 790}
                    _ = v220
                }
                _ = v213
                _ = v214
            }
            switch (("\n\t" + "hello") + "a") {
            case "foo bar":
                fmt.Println(v8, v7, v3)
            default:
                for fi221 := 0; fi221 < 6; fi221++ {
                    var v222 uint = 18446744073709551615
                    v223 := []int{(fi221 - fi221), fi221}
                    _ = fi221
                    _ = v222
                    _ = v223
                }
            }
            v224 := GenS1{F0: 558, F1: 2, F2: 42}
            _ = v224
        } else {
            v225 := GenS21{F0: "", F1: (v7 + (^(v7))), F2: 65535}
            v7 = v7
            for fi226 := 0; fi226 < 12; fi226++ {
                var v227 string = "hello"
                var v228 int16 = 1000
                _ = fi226
                _ = v227
                _ = v228
            }
            _ = v225
        }
    }
    v3 += (v3 >> ((255) & 63))
    v229 := GenS24{F0: v7, F1: 2147483647}
    v4 += (float32(1e10) / v4)
    var v230 int = (v3 | v3)
    var v231 int = v230
    v232 := []int{(v230 + -370), v230, genFn0((v3 * genFn0((^(v230)))))}
    var v233 string = "foo bar"
    var v234 int16 = 0
    for fi235 := 0; fi235 < 18; fi235++ {
        var v236 uint8 = (v8 + (v8 + v8))
        _ = fi235
        _ = v236
    }
    for fi237 := 0; fi237 < 1; fi237++ {
        fmt.Println(fi237, v7, v231)
        _ = fi237
    }
    v231 = int(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v230:", v230)
    fmt.Println("v231:", v231)
    fmt.Println("slicelen v232:", len(v232), v232)
    fmt.Println("v233:", v233)
    fmt.Println("v234:", v234)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v229
    _ = v230
    _ = v231
    _ = v232
    _ = v233
    _ = v234
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
