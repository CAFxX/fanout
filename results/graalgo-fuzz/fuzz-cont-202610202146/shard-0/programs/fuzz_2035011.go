package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 string; F1 uint8; F2 uint64 }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 bool; F1 int32 }
type GenS17 struct { F0 int64; F1 int8; F2 uint64 }
type GenS18 struct { F0 bool; F1 int8 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint32; F1 int32 }
type GenS21 struct { F0 int16; F1 int }
type GenS22 struct { F0 int16; F1 uint; F2 uint8 }
type GenS23 struct { F0 uint32; F1 bool; F2 int }
type GenS24 struct { F0 int8; F1 uint32 }
type GenS25 struct { F0 int; F1 int; F2 uint8 }
type GenS26 struct { F0 uint }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 uint }
func genFn0() int64 {
    v6 := GenS1{F0: 935, F1: -59}
    var v7 int32 = 1
    v8 := GenS1{F0: 200, F1: 100}
    v7 = v7
    var v9 string = "\n\t"
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 int32, p1 int, p2 int64) int16 {
    v10 := GenS2{F0: -1, F1: p0, F2: 2}
    for fi11 := 0; fi11 < 12; fi11++ {
        if (42 != 42) {
            v12 := GenS2{F0: -128, F1: int32(p1), F2: 223}
            v13 := GenS3{F0: 2}
            v14 := GenS4{F0: p2}
            _ = v12
            _ = v13
            _ = v14
        } else {
            v15 := GenS3{F0: 65535}
            switch "hello" {
            case "hello":
                v16 := GenS5{F0: true}
                _ = v16
            }
            _ = v15
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(fi11, p1, p1)
            v18 := []int{p1, (p1 % ((9223372036854775807) | 1)), (^(fi11)), (^(p1))}
            _ = fi17
            _ = v18
        }
        _ = fi11
    }
    v19 := GenS1{F0: 1, F1: 1}
    v20 := GenS2{F0: 20, F1: 0, F2: 456}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 uint16 = 948
    v2 := []int{}
    v3 := []string{"\n\t", "hello", "hello", "\n\t"}
    var v4 uint = 1
    var v5 uint = (v4 | v4)
    switch ("" + "\n\t") {
    case "hello":
        var v21 int = -644
        _ = v21
    default:
        _ = len(v2)
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            v5 -= 865
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            fmt.Println(v1, fi23, fi22)
            _ = fi23
        }
        v24 := GenS1{F0: v1, F1: -1}
        _ = fi22
        _ = v24
    }
    var v25 float64 = (1.5 * 0.0)
    for fi26 := 0; fi26 < 5; fi26++ {
        switch "\n\t" {
        case "":
            v25 = 3.14159
        case "\n\t":
            v27 := GenS6{F0: genFn1(-807, fi26, 937), F1: 1}
            _ = v27
        case "a":
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 uint32 = 42
                switch fi28 {
                case -2147483648:
                    var v30 uint16 = 344
                    _ = v30
                }
                var v31 float32 = -2.25
                _ = fi28
                _ = v29
                _ = v31
            }
        }
        _ = fi26
    }
    var v32 float32 = 1e10
    v4 -= (v4 - (v5 + v5))
    if ((!(!false)) || ((false || (-9223372036854775808 <= 0)) && (42 <= 42))) {
        for fi33 := 0; fi33 < 19; fi33++ {
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 int16 = 42
                switch fi33 {
                case 2147483647:
                    v25 += 1e10
                default:
                    if (626 >= 1) {
                        var v36 string = ""
                        v3 = append(v3, v36)
                        _ = v36
                    }
                }
                _ = fi34
                _ = v35
            }
            for fi37 := 0; fi37 < 18; fi37++ {
                v38 := GenS1{F0: v1, F1: -115}
                _ = len(v3)
                _ = fi37
                _ = v38
            }
            v1 += 42
            _ = fi33
        }
    } else {
        v5 = 634
    }
    if false {
        var v39 int16 = genFn1(562, 309, -588)
        var v40 float64 = v25
        v32 += v32
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        v5 -= 18446744073709551615
        fmt.Println(fi41)
        _ = len(v3)
        _ = fi41
    }
    if (!(4294967295 <= 42)) {
        v32 -= (v32 * float32(1e10))
        var v42 uint16 = 32767
        v1 = (v42 & v1)
        _ = v42
    }
    v4 = v5
    v43 := GenS7{F0: 9223372036854775807, F1: (!true)}
    for fi44 := 0; fi44 < 19; fi44++ {
        v45 := GenS8{F0: (("hello" + ((("a" + "\n\t") + "") + ("a" + (("a" + "a") + "a")))) + ((("hello" + "hello") + ("\n\t" + (("\n\t" + ("foo bar" + ("foo bar" + ((("" + ("\n\t" + "")) + "\n\t") + "foo bar")))) + ""))) + "\n\t")), F1: 121, F2: 0}
        _ = fi44
        _ = v45
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if (((-2.25 / v32) + -2.25) >= (v32 * (1.5 / 1.0))) {
        if false {
            var v46 uint16 = (v1 / (((v1 * (v1 | (v1 / ((v1) | 1))))) | 1))
            if false {
                v25 = 3.14159
                fmt.Println(v1, v5, v25)
            }
            for fi47 := 0; fi47 < 17; fi47++ {
                if true {
                    v32 = (v32 + (3.14159 - v32))
                    if ((v46 * (v1 - (v1 + uint16(v5)))) == 117) {
                        var v48 float64 = (v25 + float64(v32))
                        switch fi47 {
                        case 9223372036854775807:
                            if (true || true) {
                                fmt.Println(fi47, v46)
                                _ = len(v3)
                                if true {
                                    for fi49 := 0; fi49 < 19; fi49++ {
                                        v50 := GenS8{F0: "a", F1: 166, F2: 42}
                                        var v51 int64 = -908
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = len(v2)
                                    v5 -= 0
                                }
                            }
                        }
                        _ = v48
                    }
                } else {
                    v52 := GenS9{F0: "", F1: v46}
                    v53 := GenS10{F0: -15, F1: (^(v5))}
                    if true {
                        var v54 uint16 = v1
                        v55 := GenS1{F0: (v1 % (((v46 - (^(v46)))) | 1)), F1: 59}
                        v25 = v25
                        _ = v54
                        _ = v55
                    }
                    _ = v52
                    _ = v53
                }
                v56 := GenS6{F0: 165, F1: 65535}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi47
                _ = v56
            }
            _ = v46
        }
        if false {
            var v57 int32 = -2147483648
            v58 := GenS1{F0: (v1 / ((v1) | 1)), F1: 0}
            _ = v57
            _ = v58
        }
        if (0 >= 63) {
            v5 = v5
            v59 := GenS11{F0: true}
            switch (("" + "hello") + (("" + "hello") + "")) {
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v59
        } else {
            var v60 uint8 = 0
            _ = v60
        }
    } else {
        var v61 uint8 = 128
        _ = v61
    }
    if (778 != 127) {
        switch (("" + (("" + "foo bar") + "\n\t")) + "hello") {
        case "a":
            for fi62 := 0; fi62 < 3; fi62++ {
                if ((!false) && (int16(fi62) != -1)) {
                    if (42 < int64(fi62)) {
                        var v63 int64 = 9223372036854775807
                        _ = v63
                    } else {
                        v64 := GenS12{F0: 0, F1: -128}
                        var v65 int64 = 9223372036854775807
                        _ = v64
                        _ = v65
                    }
                    if false {
                        v66 := GenS12{F0: -927, F1: -59}
                        if ((2147483648 < 2) && false) {
                            if (fi62 == (fi62 ^ (fi62 + (fi62 % ((-2147483648) | 1))))) {
                                for fi67 := 0; fi67 < 7; fi67++ {
                                    v1 -= (v1 & (v1 % (((v1 * uint16(v4))) | 1)))
                                    var v68 float32 = v32
                                    var v69 int32 = 1
                                    _ = fi67
                                    _ = v68
                                    _ = v69
                                }
                                if (0 < 0) {
                                    v70 := GenS3{F0: (v1 & v1)}
                                    var v71 int = (fi62 / ((fi62) | 1))
                                    v71 = 1
                                    _ = v70
                                    _ = v71
                                } else {
                                    fmt.Println(v1)
                                }
                                var v72 float64 = float64(((v32 / v32) * (v32 - v32)))
                                _ = v72
                            }
                            if false {
                                if false {
                                    v25 = (1e10 - (v25 / 1.0))
                                    if false {
                                        fmt.Println(v5, fi62)
                                        switch "" {
                                        case "":
                                            var v73 float64 = (((((v25 / v25) - v25) - (v25 / 0.1)) / 3.14159) / ((((-2.25 * v25) * v25) / v25) - ((0.1 - (v25 / (v25 + 3.14159))) * 0.0)))
                                            _ = v73
                                        case "hello":
                                            var v74 int64 = 1
                                            _ = v74
                                        default:
                                            var v75 uint8 = 0
                                            _ = v75
                                        }
                                    }
                                    var v76 float64 = float64(v32)
                                    _ = v76
                                } else {
                                    v25 = v25
                                }
                                switch "a" {
                                case "a":
                                    fmt.Println(fi62)
                                case "\n\t":
                                    v25 = v25
                                case "hello":
                                    v25 = (v25 / 1e10)
                                default:
                                    v5 = v5
                                }
                            } else {
                                var v77 int8 = -1
                                v5 += 255
                                var v78 uint64 = 0
                                _ = v77
                                _ = v78
                            }
                        } else {
                            v25 -= v25
                            v79 := GenS8{F0: ("" + "foo bar"), F1: 42, F2: 42}
                            _ = v79
                        }
                        _ = v66
                    } else {
                        switch fi62 {
                        case 42:
                            _ = len(v3)
                        case 32768:
                            v32 = (3.14159 * (0.1 - v32))
                        case -497:
                            var v80 int = (fi62 * fi62)
                            _ = v80
                        default:
                            var v81 float32 = v32
                            _ = v81
                        }
                        if false {
                            var v82 uint = (^(v5))
                            v83 := GenS13{F0: 0, F1: (v82 | 4294967295)}
                            _ = v82
                            _ = v83
                        } else {
                            v4 = 18446744073709551615
                        }
                    }
                    if (65535 <= (^(v1))) {
                        v84 := GenS3{F0: 255}
                        switch "hello" {
                        case "hello":
                            v4 = 537
                        case "a":
                            v3 = append(v3, ("\n\t" + ""))
                        default:
                            var v85 int64 = 1
                            _ = v85
                        }
                        for fi86 := 0; fi86 < 5; fi86++ {
                            if false {
                                v87 := GenS4{F0: -9223372036854775808}
                                _ = v87
                            } else {
                                var v88 int32 = -1
                                v89 := GenS14{F0: fi62}
                                v32 += ((v32 * 1.5) * 1.5)
                                _ = v88
                                _ = v89
                            }
                            if (!true) {
                                fmt.Println(fi62, v1, fi62)
                                if (v25 == float64(v32)) {
                                    v90 := GenS15{F0: -2147483648}
                                    var v91 uint32 = 2
                                    var v92 int32 = 0
                                    _ = v90
                                    _ = v91
                                    _ = v92
                                } else {
                                    fmt.Println(fi62, fi62, v1)
                                    switch fi62 {
                                    case 32768:
                                        fmt.Println(fi86)
                                    case -32:
                                        var v93 int16 = -32768
                                        _ = v93
                                    default:
                                        v5 += v4
                                    }
                                    v94 := GenS16{F0: (float32(v25) <= v32), F1: 0}
                                    _ = v94
                                }
                            }
                            switch -1 {
                            case 1:
                                fmt.Println(fi62, v32)
                            case -2147483648:
                                if (!true) {
                                    switch ("foo bar" + "") {
                                    case "foo bar":
                                        if (42 >= 18) {
                                            if true {
                                                if true {
                                                    var v95 float32 = 0.1
                                                    _ = v95
                                                } else {
                                                    v4 -= v4
                                                    v4 = v5
                                                }
                                                v32 -= (v32 * float32((v25 + v25)))
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                var v96 uint8 = 14
                                                _ = v96
                                            default:
                                                fmt.Println(v1)
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            v97 := GenS17{F0: genFn0(), F1: 42, F2: 0}
                                            _ = v97
                                        }
                                    case "a":
                                        _ = len(v3)
                                    case "":
                                        v5 -= v5
                                    default:
                                        if true {
                                            v5 -= v5
                                            if false {
                                                if true {
                                                    if false {
                                                        _ = len(v2)
                                                        fmt.Println(v4)
                                                    }
                                                    var v98 int64 = 378
                                                    _ = v98
                                                }
                                            } else {
                                                v99 := GenS9{F0: ("a" + "hello"), F1: v1}
                                                v100 := GenS1{F0: (^(v1)), F1: 122}
                                                _ = v100
                                                _ = v99
                                            }
                                            v1 -= (v1 % ((v1) | 1))
                                        }
                                    }
                                    if (127 <= 100) {
                                        v1 += uint16(v5)
                                    } else {
                                        var v101 int = (fi62 * (fi86 ^ (fi86 - fi62)))
                                        v25 = v25
                                        _ = len(v3)
                                        _ = v101
                                    }
                                    fmt.Println(fi86, fi62, v1)
                                } else {
                                    var v102 int = 17
                                    var v103 uint = v5
                                    v102 -= 128
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = ("" + "foo bar")
                                        }
                                        v104 := GenS18{F0: false, F1: 0}
                                        _ = v104
                                    } else {
                                        v5 = v4
                                        if (!(!false)) {
                                            _ = len(v2)
                                            v3 = append(v3, ((("foo bar" + "a") + ("a" + ((("foo bar" + "a") + "a") + ("" + ((((("a" + ("" + "a")) + "a") + ((((("foo bar" + "hello") + "foo bar") + ("" + ((("foo bar" + "") + "\n\t") + "foo bar"))) + "\n\t") + "a")) + "\n\t") + ("a" + "foo bar")))))) + "hello"))
                                        } else {
                                            switch -91 {
                                            case 32768:
                                                v105 := GenS19{F0: 32767}
                                                _ = v105
                                            case -9223372036854775808:
                                                if (42 < -2147483648) {
                                                    if true {
                                                        var v106 int16 = -852
                                                        var v107 int32 = -967
                                                        _ = len(v2)
                                                        _ = v106
                                                        _ = v107
                                                    } else {
                                                        var v108 uint = v103
                                                        var v109 float32 = -2.25
                                                        fmt.Println(v109)
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                    var v110 int8 = -1
                                                    _ = v110
                                                } else {
                                                    v102 -= 9223372036854775807
                                                    if (uint64(v4) > 750) {
                                                        if (!false) {
                                                            switch (v102 & 9223372036854775807) {
                                                            case -269:
                                                                var v111 uint64 = 18446744073709551615
                                                                _ = v111
                                                            case 127:
                                                                var v112 string = "hello"
                                                                _ = v112
                                                            case 128:
                                                                v25 += (v25 / 0.1)
                                                            }
                                                        } else {
                                                            v102 -= 127
                                                            var v113 uint = v103
                                                            switch "hello" {
                                                            case "\n\t":
                                                                fmt.Println(v102, v103, v25)
                                                            case "":
                                                                v102 = fi62
                                                            case "hello":
                                                                v25 -= (((3.14159 * v25) / (v25 - -2.25)) * (v25 - ((1.5 * v25) - v25)))
                                                            }
                                                            if false {
                                                                var v114 bool = (true && true)
                                                                if len(v2) > 0 {
                                                                    v2[0] = (fi62 >> ((v4) & 63))
                                                                }
                                                                v115 := GenS20{F0: 2, F1: 42}
                                                                fmt.Println(v114)
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                            v116 := GenS17{F0: 42, F1: 100, F2: 2}
                                                            v117 := GenS17{F0: 9223372036854775807, F1: 0, F2: 1}
                                                            if true {
                                                                switch "hello" {
                                                                case "":
                                                                    v118 := GenS21{F0: 42, F1: 257}
                                                                    _ = v118
                                                                case "foo bar":
                                                                    v102 = (-(fi62))
                                                                default:
                                                                    v1 += 900
                                                                }
                                                                switch ("hello" + "a") {
                                                                case "\n\t":
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v113)
                                                                case "hello":
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v102
                                                                    }
                                                                    v3 = append(v3, "foo bar")
                                                                    v32 += 3.14159
                                                                }
                                                                fmt.Println(v102, v32, v113)
                                                                v32 = (1e10 - (v32 - ((v32 + v32) + float32((3.14159 + (v25 - 0.1))))))
                                                                var v119 int32 = -2147483648
                                                                _ = v119
                                                            }
                                                            if true {
                                                                switch 42 {
                                                                case 400:
                                                                    var v120 string = (("hello" + (("a" + "\n\t") + (("hello" + "") + (("" + "hello") + "\n\t")))) + (("hello" + "hello") + "\n\t"))
                                                                    _ = v120
                                                                case 0:
                                                                    if (!(42 != 501)) {
                                                                        var v121 int64 = -1
                                                                        var v122 uint16 = v1
                                                                        _ = v121
                                                                        _ = v122
                                                                    }
                                                                case 32768:
                                                                    v4 += v113
                                                                default:
                                                                    if false {
                                                                        var v123 uint64 = 945
                                                                        _ = v123
                                                                    } else {
                                                                        v124 := GenS22{F0: -178, F1: v113, F2: 255}
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ("" + "foo bar")
                                                                        }
                                                                        _ = v124
                                                                    }
                                                                }
                                                                fmt.Println(v25, v4)
                                                                var v125 uint16 = v1
                                                                _ = v125
                                                            } else {
                                                                v25 -= (float64(0.1) * float64(3.14159))
                                                                v4 = uint(v1)
                                                                fmt.Println(fi62, v1, v5)
                                                            }
                                                            var v126 uint8 = 2
                                                            _ = v113
                                                            _ = v116
                                                            _ = v117
                                                            _ = v126
                                                        }
                                                        v127 := GenS23{F0: 4294967295, F1: false, F2: (fi62 - (v102 << ((4294967295) & 63)))}
                                                        _ = v127
                                                    }
                                                }
                                            default:
                                                v5 += 0
                                            }
                                            if ((1 <= 165) && ((!false) || (!((!((29 <= 42) || (!("a" >= "a")))) && (v1 < v1))))) {
                                                fmt.Println(fi62, v1)
                                                if ((false && false) && (!(!(v32 > 3.14159)))) {
                                                    v128 := GenS6{F0: -789, F1: 0}
                                                    var v129 int64 = 1
                                                    _ = v128
                                                    _ = v129
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            } else {
                                                var v130 uint16 = v1
                                                var v131 string = "a"
                                                _ = v130
                                                _ = v131
                                            }
                                        }
                                    }
                                    v132 := GenS12{F0: genFn0(), F1: 100}
                                    _ = v102
                                    _ = v103
                                    _ = v132
                                }
                            }
                            _ = fi86
                        }
                        _ = v84
                    } else {
                        var v133 uint32 = 4294967295
                        switch "" {
                        case "a":
                            var v134 uint64 = 788
                            _ = v134
                        case "hello":
                            var v135 int = (fi62 | fi62)
                            _ = v135
                        }
                        _ = v133
                    }
                } else {
                    switch (((("" + "a") + "\n\t") + "foo bar") + ("" + "foo bar")) {
                    case "\n\t":
                        if ((((v32 + 1.5) != 0.1) || (!false)) || false) {
                            switch "\n\t" {
                            case "":
                                v25 = v25
                            case "\n\t":
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = fi62
                                    }
                                    v136 := GenS21{F0: -1, F1: 1}
                                    v137 := GenS24{F0: 100, F1: 2147483647}
                                    _ = v136
                                    _ = v137
                                } else {
                                    if (0 > -1) {
                                        for fi138 := 0; fi138 < 8; fi138++ {
                                            if false {
                                                if (v1 == v1) {
                                                    if true {
                                                        if len(v2) > 0 {
                                                            v2[0] = 0
                                                        }
                                                    }
                                                    v139 := GenS17{F0: 28, F1: 1, F2: 42}
                                                    v4 += (v5 | v4)
                                                    var v140 int8 = 100
                                                    _ = v139
                                                    _ = v140
                                                }
                                                v5 = (v4 & 42)
                                                var v141 string = ""
                                                _ = v141
                                            }
                                            var v142 int16 = -810
                                            switch (fi138 | fi138) {
                                            case -388:
                                                switch "\n\t" {
                                                case "a":
                                                    if false {
                                                        switch fi138 {
                                                        case 32767:
                                                            var v143 bool = ((fi62 > (fi62 | fi138)) || true)
                                                            _ = v143
                                                        case -9223372036854775808:
                                                            if len(v3) > 0 {
                                                                v3[0] = ((("hello" + "\n\t") + ("a" + (("" + (("a" + "a") + "")) + ""))) + "hello")
                                                            }
                                                        case 128:
                                                            fmt.Println(v5, fi138, fi138)
                                                        }
                                                        v5 = (^(v4))
                                                    } else {
                                                        v1 += (^(v1))
                                                    }
                                                case "\n\t":
                                                    fmt.Println(fi62, fi138)
                                                }
                                            case -1:
                                                var v144 float32 = float32(v25)
                                                _ = v144
                                            case -499:
                                                switch "foo bar" {
                                                case "hello":
                                                    if false {
                                                        fmt.Println(fi62)
                                                        switch (-(fi138)) {
                                                        case 32767:
                                                            v142 = (v142 << ((v4) & 63))
                                                        }
                                                        v145 := GenS25{F0: 9223372036854775807, F1: (fi62 ^ fi62), F2: 255}
                                                        _ = v145
                                                    } else {
                                                        v146 := GenS26{F0: 65535}
                                                        fmt.Println(fi62, v32, v4)
                                                        if (-2.25 > v32) {
                                                            if (42 >= uint64(v4)) {
                                                                var v147 uint64 = 18446744073709551615
                                                                fmt.Println(v147, v4)
                                                                var v148 bool = (1 > 2)
                                                                fmt.Println(v148)
                                                                _ = v147
                                                                _ = v148
                                                            } else {
                                                                v2 = append(v2, (^(fi138)))
                                                                var v149 bool = true
                                                                _ = v149
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v146
                                                    }
                                                    v1 += 152
                                                    switch (("foo bar" + ("foo bar" + "a")) + "a") {
                                                    case "a":
                                                        v142 -= -32768
                                                    }
                                                case "a":
                                                    var v150 int32 = -1
                                                    _ = v150
                                                default:
                                                    v151 := GenS27{F0: 255}
                                                    _ = v151
                                                }
                                            }
                                            _ = fi138
                                            _ = v142
                                        }
                                        var v152 uint = v4
                                        _ = v152
                                    } else {
                                        v4 -= 255
                                        switch (fi62 - (fi62 * fi62)) {
                                        case -579:
                                            switch (fi62 ^ 740) {
                                            case -257:
                                                fmt.Println(v4, v4, v5)
                                            }
                                        default:
                                            var v153 uint16 = 42
                                            _ = v153
                                        }
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        v1 -= 32767
                                    case "":
                                        var v154 int16 = 1000
                                        _ = v154
                                    case "hello":
                                        var v155 int32 = 0
                                        _ = v155
                                    }
                                    v5 -= (v4 + (v4 >> (((^(v5))) & 63)))
                                }
                            }
                            v156 := GenS2{F0: 1, F1: 2147483647, F2: 214}
                            _ = v156
                        } else {
                            var v157 int8 = 1
                            switch (fi62 % ((1) | 1)) {
                            case 4294967295:
                                v3 = append(v3, (("foo bar" + "foo bar") + ""))
                            }
                            _ = v157
                        }
                    case "":
                        var v158 uint8 = 81
                        _ = v158
                    default:
                        var v159 int = 128
                        _ = v159
                    }
                    fmt.Println(v25)
                }
                if len(v2) > 0 {
                    v2[0] = (fi62 | (fi62 | (fi62 / ((fi62) | 1))))
                }
                v4 = v5
                _ = fi62
            }
        default:
            if (128 >= 785) {
                v1 = 1
                v4 += v5
                if false {
                    var v160 int16 = -1
                    _ = v160
                } else {
                    v161 := GenS21{F0: 435, F1: 0}
                    _ = v161
                }
            } else {
                fmt.Println(v1, v4)
                if len(v3) > 0 {
                    v3[0] = ((("\n\t" + "") + "foo bar") + "\n\t")
                }
                switch 1 {
                case 2147483647:
                    if true {
                        v162 := GenS28{F0: v5}
                        if (42 <= -705) {
                            v3 = append(v3, "hello")
                            v32 -= 1.5
                        } else {
                            switch 563 {
                            case 4294967295:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 1:
                                fmt.Println(v5)
                            }
                            _ = len(v2)
                        }
                        var v163 int16 = -1
                        _ = v162
                        _ = v163
                    }
                case 32768:
                    v5 = 42
                default:
                    var v164 int = 2147483647
                    _ = v164
                }
            }
        }
        var v165 bool = (!(false && (!(-588 < 591))))
        _ = v165
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
}
