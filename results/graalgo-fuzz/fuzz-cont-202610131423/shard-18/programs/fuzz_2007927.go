package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS8 struct { F0 int8; F1 uint8; F2 string }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 bool; F1 int8 }
type GenS11 struct { F0 int16; F1 string }
func genFn0(p0 int64, p1 int, p2 uint16) uint {
    if false {
        fmt.Println(p1)
    } else {
        if false {
            v6 := GenS1{F0: -74}
            _ = v6
        }
    }
    if false {
        var v7 int32 = 2147483647
        for fi8 := 0; fi8 < 10; fi8++ {
            p1 = fi8
            _ = fi8
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            p1 = p1
            var v10 uint8 = 113
            v11 := GenS1{F0: 1}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        p0 = (^(p0))
        v13 := GenS2{F0: 1, F1: 4294967295, F2: 171}
        _ = fi12
        _ = v13
    }
    if false {
        var v14 bool = true
        v14 = v14
        _ = v14
    } else {
        var v15 float32 = (0.1 + 0.0)
        v16 := GenS1{F0: int8(p2)}
        fmt.Println(p2, v15)
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            if (uint8(p2) > 87) {
                var v19 int8 = 42
                fmt.Println(fi17, p1)
                if (32767 >= -32768) {
                    var v20 float32 = (3.14159 + 3.14159)
                    if (false || false) {
                        v21 := GenS1{F0: v19}
                        switch "hello" {
                        case "":
                            var v22 bool = true
                            _ = v22
                        }
                        var v23 int = p1
                        fmt.Println(v23, p1, fi17)
                        _ = v21
                        _ = v23
                    }
                    _ = v20
                }
                var v24 int64 = (p0 / ((p0) | 1))
                switch p1 {
                case -9223372036854775808:
                    switch (fi18 + (p1 | fi18)) {
                    case -762:
                        v25 := GenS3{F0: p2, F1: (-(p0)), F2: v24}
                        _ = v25
                    case 32768:
                        switch fi17 {
                        case 0:
                            v26 := GenS4{F0: 2}
                            _ = v26
                        default:
                            v27 := []int{}
                            _ = v27
                        }
                    case -1:
                        if false {
                            var v28 int8 = v19
                            switch ("\n\t" + (("hello" + ("hello" + "a")) + "\n\t")) {
                            case "a":
                                switch "\n\t" {
                                case "\n\t":
                                    var v29 string = "foo bar"
                                    _ = v29
                                case "":
                                    fmt.Println(p2)
                                default:
                                    var v30 bool = true
                                    _ = v30
                                }
                                v31 := GenS5{F0: uint32(p2)}
                                _ = v31
                            case "":
                                if false {
                                    var v32 float32 = 1.5
                                    v33 := GenS6{F0: v24}
                                    if (!(true || (0 > 32767))) {
                                        p1 = p1
                                        if true {
                                            if (2 < 194) {
                                                var v34 float64 = 1.5
                                                var v35 int32 = int32(v19)
                                                _ = v34
                                                _ = v35
                                            } else {
                                                switch (^(fi17)) {
                                                case 4294967295:
                                                    v36 := []int{(fi18 / ((p1) | 1)), (^(p1)), 2147483647, fi18}
                                                    v36 = append(v36, -429)
                                                    var v37 uint16 = (p2 ^ (p2 ^ (^(p2))))
                                                    _ = v36
                                                    _ = v37
                                                default:
                                                    var v38 uint32 = 4294967295
                                                    _ = v38
                                                }
                                            }
                                            fmt.Println(p0)
                                            v39 := GenS4{F0: 46}
                                            _ = v39
                                        }
                                    }
                                    _ = v32
                                    _ = v33
                                } else {
                                    if (42 >= 2147483648) {
                                        p1 = (p1 * fi17)
                                        fmt.Println(p1, v28)
                                    }
                                    v24 = (p0 ^ int64(p1))
                                    v19 = v19
                                    if false {
                                        v40 := GenS7{F0: 42, F1: v24, F2: 538}
                                        v41 := []int{-734, fi17, int(v28), (p1 + -881)}
                                        _ = v40
                                        _ = v41
                                    }
                                    p1 = (fi17 | p1)
                                }
                            case "foo bar":
                                v42 := GenS4{F0: 0}
                                _ = v42
                            default:
                                v43 := []int{fi17, (^(fi17)), fi18, -221}
                                _ = v43
                            }
                            var v44 uint64 = 1
                            _ = v28
                            _ = v44
                        } else {
                            switch 281 {
                            case -560:
                                v24 = (p0 * v24)
                            case 9223372036854775807:
                                v45 := []int{}
                                _ = v45
                            case 32767:
                                var v46 float64 = 0.1
                                _ = v46
                            default:
                                fmt.Println(fi18, v24, fi18)
                            }
                            v47 := []int{}
                            p1 -= -997
                            _ = v47
                        }
                        v48 := []int{}
                        var v49 uint64 = 2
                        var v50 string = "hello"
                        _ = v48
                        _ = v49
                        _ = v50
                    default:
                        if (((!false) && true) || false) {
                            switch p1 {
                            case 0:
                                p0 = (v24 >> ((0) & 63))
                            default:
                                var v51 uint32 = 0
                                _ = v51
                            }
                        }
                    }
                }
                p1 += p1
                var v52 uint8 = 255
                _ = v19
                _ = v24
                _ = v52
            } else {
                switch p1 {
                case 477:
                    switch "hello" {
                    case "":
                        v53 := GenS2{F0: 0, F1: 18446744073709551615, F2: 18446744073709551615}
                        _ = v53
                    }
                case 4294967295:
                    var v54 int8 = 113
                    _ = v54
                }
                p1 = (fi18 >> ((510) & 63))
                var v55 int16 = 465
                _ = v55
            }
            _ = fi18
        }
        var v56 int8 = -128
        v57 := []int{(^(fi17))}
        _ = fi17
        _ = v56
        _ = v57
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int = 128
    var v2 float32 = 1e10
    var v3 uint16 = 65535
    v4 := []int{}
    v5 := make(map[string]string)
    for fi58 := 0; fi58 < 19; fi58++ {
        v59 := GenS4{F0: 211}
        switch (fi58 << ((266) & 63)) {
        case 127:
            if ("hello" == (("hello" + "foo bar") + "\n\t")) {
                var v60 uint16 = (v3 + v3)
                _ = v60
            }
        case -2147483648:
            var v61 string = "a"
            _ = v61
        }
        _ = fi58
        _ = v59
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    var v62 int16 = -854
    if false {
        v63 := GenS1{F0: 100}
        _ = v63
    }
    switch "foo bar" {
    case "hello":
        v64 := GenS8{F0: int8(v3), F1: 42, F2: ((("a" + "foo bar") + "foo bar") + ("hello" + "a"))}
        _ = v64
    case "\n\t":
        for fi65 := 0; fi65 < 4; fi65++ {
            for fi66 := 0; fi66 < 12; fi66++ {
                v4 = append(v4, 128)
                _ = fi66
            }
            if false {
                _ = len(v4)
            } else {
                v1 = (v1 + v1)
            }
            for fi67 := 0; fi67 < 10; fi67++ {
                switch (-(v1)) {
                case 42:
                    var v68 int = 9223372036854775807
                    v69 := GenS6{F0: -9223372036854775808}
                    _ = v68
                    _ = v69
                case 4294967295:
                    if (((1e10 / -2.25) + 1e10) == ((-2.25 * (1.5 - (float64(0.0) * (-2.25 * 0.1)))) + 3.14159)) {
                        v70 := GenS2{F0: 0, F1: 22, F2: 610}
                        if ((0 == 372) && true) {
                            switch "" {
                            case "foo bar":
                                fmt.Println(fi67)
                            case "a":
                                if (!true) {
                                    var v71 uint64 = 0
                                    _ = v71
                                } else {
                                    var v72 int16 = (^(v62))
                                    _ = len(v5)
                                    v73 := GenS9{F0: 223, F1: 18446744073709551615}
                                    _ = len(v5)
                                    _ = v72
                                    _ = v73
                                }
                            case "\n\t":
                                v62 = 42
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            if false {
                                if (false || true) {
                                    var v74 uint16 = v3
                                    var v75 int8 = 42
                                    var v76 int = (^(v1))
                                    v77 := GenS1{F0: v75}
                                    v75 += v75
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                } else {
                                    fmt.Println(v2, v3, fi67)
                                    var v78 int = (^(v1))
                                    _ = v78
                                }
                            }
                        } else {
                            v79 := GenS7{F0: 65535, F1: -316, F2: 0}
                            v5[""] = "\n\t"
                            _ = v79
                        }
                        var v80 string = ("" + "a")
                        v5[((v80 + v80) + v80)] = v80
                        _ = v70
                        _ = v80
                    }
                case 0:
                    if (1 == 1) {
                        v2 += v2
                    }
                default:
                    if true {
                        v3 -= (v3 * (v3 | v3))
                        v5["foo bar"] = "foo bar"
                    } else {
                        v81 := GenS9{F0: 255, F1: 18446744073709551615}
                        v3 = (v3 & v3)
                        v2 = 1e10
                        _ = v81
                    }
                }
                if (304 == 0) {
                    if (true || (((int(v62) > 0) || true) && (0.1 > -2.25))) {
                        switch "foo bar" {
                        case "a":
                            switch fi67 {
                            case -39:
                                var v82 int32 = 1
                                _ = v82
                            case 4294967295:
                                switch "\n\t" {
                                case "a":
                                    var v83 int = fi65
                                    _ = v83
                                }
                            case 128:
                                var v84 float32 = 1e10
                                _ = v84
                            default:
                                if (true && true) {
                                    v3 = v3
                                    if (("foo bar" + "") == "") {
                                        v85 := GenS6{F0: 9223372036854775807}
                                        v2 = float32(1.5)
                                        _ = v85
                                    }
                                }
                            }
                        case "foo bar":
                            fmt.Println(v2)
                        case "hello":
                            if true {
                                v86 := GenS4{F0: 231}
                                switch "hello" {
                                case "":
                                    v87 := GenS3{F0: 42, F1: 42, F2: -947}
                                    _ = v87
                                case "\n\t":
                                    var v88 uint = 42
                                    _ = v88
                                case "foo bar":
                                    var v89 int32 = -284
                                    _ = v89
                                }
                                if true {
                                    var v90 float32 = 1e10
                                    _ = v90
                                } else {
                                    v3 += (v3 >> ((2) & 63))
                                }
                                _ = v86
                            } else {
                                if true {
                                    if (0 <= 127) {
                                        var v91 int = (fi67 * (fi65 & (-(v1))))
                                        v92 := GenS10{F0: (0.0 > v2), F1: 0}
                                        v91 = (fi67 + v1)
                                        _ = v91
                                        _ = v92
                                    } else {
                                        var v93 int8 = 42
                                        _ = v93
                                    }
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = v1
                                    }
                                    v94 := GenS11{F0: v62, F1: ((("\n\t" + (("foo bar" + "foo bar") + "foo bar")) + "\n\t") + "")}
                                    switch ("a" + "a") {
                                    case "hello":
                                        if len(v4) > 0 {
                                            v4[0] = 429
                                        }
                                    case "\n\t":
                                        if (!true) {
                                            if true {
                                                fmt.Println(v3, v3)
                                                v1 = (fi65 | fi67)
                                                if true {
                                                    v5["\n\t"] = ""
                                                } else {
                                                    v3 = 65535
                                                }
                                            }
                                            switch v1 {
                                            case 127:
                                                if true {
                                                    fmt.Println(fi67, fi65)
                                                }
                                            case 128:
                                                fmt.Println(fi67, v2)
                                            }
                                            if (true || (1 >= -1)) {
                                                v3 += 1
                                                var v95 int32 = -480
                                                if ("hello" <= (("\n\t" + (("\n\t" + "") + "foo bar")) + "a")) {
                                                    v2 -= (1.5 - -2.25)
                                                }
                                                _ = v95
                                            }
                                        } else {
                                            fmt.Println(fi65, v2)
                                            if true {
                                                var v96 string = "a"
                                                switch (fi67 >> ((18446744073709551615) & 63)) {
                                                case 4294967295:
                                                    v96 = v96
                                                case -2147483648:
                                                    switch "" {
                                                    case "\n\t":
                                                        delete(v5, (v96 + v96))
                                                    case "foo bar":
                                                        v1 = 2147483647
                                                    }
                                                case -116:
                                                    var v97 int64 = 9223372036854775807
                                                    _ = v97
                                                default:
                                                    v98 := GenS7{F0: 2147483648, F1: 9223372036854775807, F2: 0}
                                                    _ = v98
                                                }
                                                _ = v96
                                            } else {
                                                var v99 uint8 = 0
                                                if len(v4) > 0 {
                                                    v4[0] = fi65
                                                }
                                                _ = v99
                                            }
                                            v100 := GenS11{F0: 1, F1: "\n\t"}
                                            _ = v100
                                        }
                                    case "a":
                                        if false {
                                            switch fi65 {
                                            case 127:
                                                v62 -= (^(v62))
                                            case -298:
                                                var v101 int16 = v62
                                                _ = v101
                                            default:
                                                var v102 bool = (-1 >= 42)
                                                _ = v102
                                            }
                                            var v103 uint64 = 42
                                            v104 := v5["foo bar"]
                                            _ = v103
                                            _ = v104
                                        }
                                    default:
                                        var v105 int32 = 2147483647
                                        _ = v105
                                    }
                                    _ = v94
                                }
                                v106 := GenS11{F0: (v62 | 42), F1: (("a" + ("a" + "foo bar")) + ("" + "foo bar"))}
                                _ = v106
                            }
                        default:
                            switch ((("\n\t" + "hello") + "hello") + "foo bar") {
                            case "foo bar":
                                v107 := GenS4{F0: uint8(v3)}
                                _ = v107
                            }
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = fi65
                        }
                        var v108 uint64 = 42
                        _ = v108
                    }
                    var v109 int64 = 42
                    v62 -= v62
                    _ = v109
                }
                v62 = v62
                _ = fi67
            }
            _ = fi65
        }
    }
    var v110 uint8 = 1
    switch "a" {
    case "\n\t":
        if (v62 > v62) {
            for fi111 := 0; fi111 < 4; fi111++ {
                var v112 uint8 = v110
                v112 = 2
                _ = fi111
                _ = v112
            }
            var v113 int32 = 874
            v2 = 1e10
            _ = v113
        }
    case "":
        if true {
            v3 = 710
            v62 += 42
        } else {
            v114 := GenS4{F0: (v110 << ((42) & 63))}
            fmt.Println(v1)
            v1 = (v1 | v1)
            _ = v114
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v3 += 857
    if true {
        var v115 uint16 = v3
        _ = v115
    } else {
        for fi116 := 0; fi116 < 8; fi116++ {
            fmt.Println(v1, v1)
            _ = fi116
        }
    }
    if (958 < 2) {
        var v117 int = (v1 & v1)
        if false {
            v118 := GenS5{F0: 716}
            _ = v118
        } else {
            var v119 uint32 = 2
            var v120 uint = 42
            if ((!false) && (true || (true || (v117 > v1)))) {
                for fi121 := 0; fi121 < 2; fi121++ {
                    var v122 bool = (1 == -9223372036854775808)
                    _ = fi121
                    _ = v122
                }
                _ = len(v5)
                var v123 uint32 = v119
                _ = v123
            }
            _ = v119
            _ = v120
        }
        _ = v117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
}
