package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 bool; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 int; F2 uint }
type GenS8 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS14 struct { F0 int64; F1 int; F2 uint16 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 string; F1 int16 }
type GenS18 struct { F0 bool; F1 int; F2 int8 }
type GenS19 struct { F0 int; F1 int8 }
type GenS20 struct { F0 int; F1 uint; F2 uint }
func genFn0() int16 {
    var v4 int = 939
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        if (!("\n\t" > "a")) {
            if false {
                fmt.Println(v4)
            }
        }
    }
    if (1 == 2) {
        v5 := []int{127, v4}
        var v6 uint = 957
        v7 := GenS1{F0: 762, F1: 4294967295}
        _ = v5
        _ = v6
        _ = v7
    }
    _ = v4
    return -571
}
func genFn1(p0 uint32) int32 {
    v8 := []int{9223372036854775807, 32768, 1}
    var v9 int8 = 100
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float32 = float32(float64((0.0 * -2.25)))
        fmt.Println(fi10, v9)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        _ = len(v8)
        _ = fi12
    }
    _ = p0
    _ = v8
    _ = v9
    return 124
}

func main() {
    v1 := []string{"", ("" + "a"), "\n\t"}
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    if false {
        switch "a" {
        case "foo bar":
            var v13 bool = (!false)
            _ = v13
        case "hello":
            fmt.Println(v2)
        }
        v2 = 1.5
    } else {
        v1 = append(v1, "foo bar")
    }
    v14 := GenS2{F0: 65535}
    for fi15 := 0; fi15 < 8; fi15++ {
        v2 -= 1e10
        if (!(false && (!false))) {
            var v16 int64 = -1
            delete(v3, 356)
            _ = v16
        } else {
            v1 = append(v1, "a")
            v17 := GenS3{F0: true, F1: false, F2: 892}
            var v18 int32 = genFn1(573)
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    switch "foo bar" {
    case "a":
        var v19 int64 = -1
        _ = v19
    case "":
        v3[2147483647] = (18446744073709551615 < 555)
    case "\n\t":
        fmt.Println(v2)
    }
    var v20 int32 = 42
    if (genFn1(2147483648) >= v20) {
        v2 += (v2 + (float32(1e10) / 1.5))
    } else {
        if false {
            for fi21 := 0; fi21 < 10; fi21++ {
                if (1 >= 1) {
                    v2 += v2
                } else {
                    if len(v1) > 0 {
                        v1[0] = ("" + (("foo bar" + "") + ("hello" + ("a" + (("foo bar" + "\n\t") + "a")))))
                    }
                    var v22 uint = 65535
                    _ = v22
                }
                _ = fi21
            }
            var v23 uint16 = 2
            _ = v23
        }
        switch "\n\t" {
        case "foo bar":
            if ((!(("hello" != "hello") || (3.14159 < 3.14159))) && false) {
                v2 = 0.0
                var v24 bool = (2147483647 < v20)
                if v24 {
                    v25 := GenS4{F0: -25}
                    v26 := GenS2{F0: 4294967295}
                    if (!(2 != 2)) {
                        for fi27 := 0; fi27 < 16; fi27++ {
                            var v28 bool = (v24 || false)
                            _ = fi27
                            _ = v28
                        }
                        v2 = (-2.25 + ((0.0 + v2) - v2))
                        for fi29 := 0; fi29 < 2; fi29++ {
                            v2 = (v2 * -2.25)
                            if v24 {
                                v20 = (v20 + (-(v20)))
                                switch "" {
                                case "hello":
                                    for fi30 := 0; fi30 < 4; fi30++ {
                                        var v31 float32 = 0.1
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        _ = fi30
                                        _ = v31
                                    }
                                default:
                                    v32 := GenS4{F0: -128}
                                    _ = v32
                                }
                                v24 = (32767 != 255)
                            } else {
                                switch 42 {
                                case -2147483648:
                                    switch (fi29 | (fi29 / ((fi29) | 1))) {
                                    case -975:
                                        v33 := GenS3{F0: (!true), F1: ((fi29 | fi29) >= (-(fi29))), F2: 0}
                                        _ = v33
                                    case -2147483648:
                                        if (0.1 >= (0.1 + 0.0)) {
                                            if (!(4294967295 <= 255)) {
                                                _ = len(v3)
                                            }
                                        } else {
                                            v2 = v2
                                            var v34 float64 = 1e10
                                            v35 := GenS1{F0: 1000, F1: 678}
                                            _ = v34
                                            _ = v35
                                        }
                                    case -1:
                                        v24 = true
                                    default:
                                        for fi36 := 0; fi36 < 10; fi36++ {
                                            v37 := GenS5{F0: 737, F1: 2147483647, F2: -1}
                                            if v24 {
                                                if true {
                                                    if false {
                                                        v38 := GenS6{F0: ("hello" + "hello")}
                                                        _ = v38
                                                    }
                                                    v39 := GenS1{F0: 619, F1: 0}
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v2 = v2
                                                    }
                                                    _ = v39
                                                }
                                                fmt.Println(fi36, fi36, fi29)
                                            } else {
                                                fmt.Println(fi29, fi29)
                                                if (!false) {
                                                    v40 := GenS7{F0: (!v24), F1: fi36, F2: 18446744073709551615}
                                                    if (v24 && false) {
                                                        v24 = ((!v24) || v24)
                                                        v41 := GenS8{F0: 102, F1: 1, F2: 77}
                                                        _ = v41
                                                    } else {
                                                        fmt.Println(v2, v24, fi36)
                                                        var v42 uint32 = 175
                                                        if false {
                                                            v43 := GenS2{F0: (^(v42))}
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v20 -= 595
                                                            _ = v43
                                                        } else {
                                                            v44 := GenS6{F0: ((((("a" + ("a" + "foo bar")) + "\n\t") + "") + "a") + "")}
                                                            v1 = append(v1, "")
                                                            switch ((("\n\t" + (("foo bar" + "hello") + "hello")) + "foo bar") + "") {
                                                            case "\n\t":
                                                                var v45 uint32 = v42
                                                                _ = v45
                                                            default:
                                                                var v46 int = fi29
                                                                _ = v46
                                                            }
                                                            v2 += (((float32(float64((v2 + (3.14159 - (v2 - (v2 / v2)))))) * v2) * ((0.1 / ((float32(-2.25) - (v2 / v2)) / 1.0)) * v2)) + v2)
                                                            v24 = (519 <= 32767)
                                                            _ = v44
                                                        }
                                                        _ = v42
                                                    }
                                                    fmt.Println(v2, fi29)
                                                    var v47 int8 = 127
                                                    _ = v40
                                                    _ = v47
                                                } else {
                                                    if false {
                                                        if (!true) {
                                                            v24 = v24
                                                            v24 = true
                                                            var v48 float64 = -2.25
                                                            _ = v48
                                                        }
                                                    } else {
                                                        fmt.Println(v2, fi29)
                                                        v49 := GenS9{F0: "a", F1: 1}
                                                        if (v24 || (542 == 676)) {
                                                            if (1e10 != 1.5) {
                                                                fmt.Println(fi36)
                                                                v50 := GenS10{F0: 0, F1: 191}
                                                                if (!v24) {
                                                                    fmt.Println(fi36, v2)
                                                                    v24 = (32767 == 32767)
                                                                } else {
                                                                    v2 = (v2 + v2)
                                                                    _ = len(v3)
                                                                    fmt.Println(v20, v2, fi36)
                                                                    if false {
                                                                        v2 = 0.1
                                                                    }
                                                                }
                                                                _ = v50
                                                            } else {
                                                                if (3.14159 > -2.25) {
                                                                    var v51 int64 = 42
                                                                    fmt.Println(v24, v2)
                                                                    _ = v51
                                                                }
                                                            }
                                                            var v52 float32 = float32(1e10)
                                                            v20 = (-(v20))
                                                            v24 = (2147483647 != 2)
                                                            if (2 < 65535) {
                                                                _ = len(v1)
                                                                v3[2] = (!(-2.25 <= (-2.25 * (0.0 + (((float64(0.0) * float64(-2.25)) - 3.14159) * 3.14159)))))
                                                                v52 += v52
                                                                v53 := GenS2{F0: 374}
                                                                _ = v53
                                                            } else {
                                                                var v54 float64 = ((0.1 * (3.14159 - (3.14159 * 1e10))) - (((1e10 * (float64(v52) - 3.14159)) * 0.1) + 3.14159))
                                                                if (42 == 0) {
                                                                    v24 = (!v24)
                                                                    _ = len(v1)
                                                                    v52 += (1e10 + float32(((v54 / (v54 / v54)) - v54)))
                                                                    v24 = (!false)
                                                                }
                                                                _ = v54
                                                            }
                                                            if v24 {
                                                                var v55 int8 = 1
                                                                var v56 int32 = v20
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                            fmt.Println(fi29, fi29)
                                                            _ = v52
                                                        }
                                                        _ = v49
                                                    }
                                                }
                                                v24 = (("" + ("hello" + "foo bar")) <= "a")
                                            }
                                            var v57 string = ("\n\t" + ((("" + "hello") + "a") + ("foo bar" + "hello")))
                                            _ = fi36
                                            _ = v37
                                            _ = v57
                                        }
                                    }
                                default:
                                    for fi58 := 0; fi58 < 10; fi58++ {
                                        var v59 string = ("\n\t" + "a")
                                        v60 := GenS11{F0: 857, F1: 679}
                                        _ = fi58
                                        _ = v59
                                        _ = v60
                                    }
                                }
                            }
                            v61 := GenS12{F0: 255}
                            _ = fi29
                            _ = v61
                        }
                    }
                    _ = v25
                    _ = v26
                } else {
                    v2 -= v2
                    if (true || v24) {
                        for fi62 := 0; fi62 < 13; fi62++ {
                            for fi63 := 0; fi63 < 6; fi63++ {
                                v64 := GenS5{F0: 255, F1: 0, F2: -57}
                                switch fi62 {
                                case 4294967295:
                                    if v24 {
                                        _ = len(v3)
                                        fmt.Println(fi63, v2)
                                    }
                                case 32768:
                                    v65 := v3[203]
                                    _ = v65
                                }
                                switch "foo bar" {
                                case "a":
                                    v66 := GenS5{F0: 590, F1: 1, F2: 1}
                                    _ = v66
                                case "":
                                    if true {
                                        if ((!((42 >= 255) && (v24 && ((-(fi62)) != (fi62 - (fi62 | (fi62 | fi63))))))) && (42 > 127)) {
                                            v67 := GenS13{F0: 2, F1: 99, F2: -1}
                                            v68 := GenS3{F0: v24, F1: false, F2: 1}
                                            _ = v67
                                            _ = v68
                                        }
                                        fmt.Println(v24, v20)
                                        if (2 <= 2) {
                                            v24 = (!(1 <= 381))
                                            v69 := GenS14{F0: 0, F1: 9223372036854775807, F2: 255}
                                            var v70 int32 = v20
                                            var v71 int16 = genFn0()
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                        } else {
                                            var v72 uint = 255
                                            v73 := GenS15{F0: 94}
                                            _ = v72
                                            _ = v73
                                        }
                                    } else {
                                        var v74 int8 = 100
                                        _ = v74
                                    }
                                case "\n\t":
                                    if v24 {
                                        switch fi63 {
                                        case -9223372036854775808:
                                            var v75 float64 = 1e10
                                            _ = v75
                                        case -630:
                                            var v76 int16 = -1
                                            _ = v76
                                        case 385:
                                            v3[1] = (!v24)
                                            if v24 {
                                                switch int(v20) {
                                                case 214:
                                                    if (!false) {
                                                        var v77 int16 = -32768
                                                        if ((fi62 << ((255) & 63)) != (fi62 & fi62)) {
                                                            v77 += 42
                                                            if true {
                                                                var v78 int16 = -1
                                                                delete(v3, 65535)
                                                                _ = v78
                                                            }
                                                        }
                                                        _ = v77
                                                    } else {
                                                        if ((42 <= 1) || (!((fi63 * fi63) == 127))) {
                                                            if true {
                                                                v79 := v3[2147483647]
                                                                _ = v79
                                                            } else {
                                                                v24 = v24
                                                                var v80 int64 = -751
                                                                v24 = (false && (1e10 <= 1e10))
                                                                _ = v80
                                                            }
                                                            v2 = v2
                                                            var v81 bool = false
                                                            _ = v81
                                                        } else {
                                                            switch ("" + "") {
                                                            case "hello":
                                                                var v82 int = (fi62 & 1)
                                                                _ = v82
                                                            case "":
                                                                fmt.Println(v24, fi63, v20)
                                                            case "a":
                                                                if v24 {
                                                                    v83 := GenS10{F0: 32767, F1: 2}
                                                                    v84 := GenS16{F0: 452}
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v24)
                                                        if (true || v24) {
                                                            v2 -= 0.1
                                                        }
                                                    }
                                                case 42:
                                                    v2 = 1e10
                                                case 128:
                                                    switch (fi63 + fi62) {
                                                    case 0:
                                                        if v24 {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    case 128:
                                                        if (!(0 >= -1)) {
                                                            var v85 uint16 = 255
                                                            _ = v85
                                                        } else {
                                                            var v86 float32 = (float32((3.14159 - 1e10)) / (-2.25 * (v2 / v2)))
                                                            _ = v86
                                                        }
                                                    }
                                                default:
                                                    var v87 int = 2147483647
                                                    _ = v87
                                                }
                                                if ((false && false) || v24) {
                                                    v20 = genFn1(100)
                                                    v20 += v20
                                                    v88 := GenS2{F0: 65535}
                                                    _ = v88
                                                } else {
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "a":
                                                            fmt.Println(v24, v24, v20)
                                                        case "\n\t":
                                                            var v89 int16 = int16(fi63)
                                                            _ = v89
                                                        case "":
                                                            if (v24 && (!v24)) {
                                                                if false {
                                                                    v90 := GenS17{F0: "foo bar", F1: 1}
                                                                    v24 = ((-62 > 0) && (v20 >= -2147483648))
                                                                    _ = v90
                                                                } else {
                                                                    v91 := GenS18{F0: (v24 && (v24 || true)), F1: (fi62 | fi62), F2: 0}
                                                                    v92 := GenS9{F0: "", F1: 42}
                                                                    v20 -= (v20 ^ v20)
                                                                    v20 += v20
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v93 := GenS6{F0: ""}
                                                                    switch fi62 {
                                                                    case 4294967295:
                                                                        v24 = (1 == 9223372036854775807)
                                                                    case 1:
                                                                        v20 -= v20
                                                                    default:
                                                                        v20 += genFn1(830)
                                                                    }
                                                                    delete(v3, 1)
                                                                    _ = v91
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                                delete(v3, 2147483647)
                                                            }
                                                        default:
                                                            _ = len(v3)
                                                        }
                                                    case "hello":
                                                        if true {
                                                            v94 := GenS16{F0: 2}
                                                            _ = v94
                                                        }
                                                    }
                                                    v2 = v2
                                                }
                                                v3[2] = v24
                                            }
                                            v95 := GenS19{F0: fi63, F1: -128}
                                            _ = v95
                                        default:
                                            if (127 == 42) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v96 int16 = -401
                                                if (1 < 2) {
                                                    v96 = (v96 >> ((738) & 63))
                                                }
                                                _ = v96
                                            }
                                        }
                                        if (("foo bar" + "\n\t") == "\n\t") {
                                            switch ("" + "") {
                                            case "hello":
                                                switch 260 {
                                                case -9223372036854775808:
                                                    v20 += (v20 & -2147483648)
                                                case 729:
                                                    v97 := GenS20{F0: 42, F1: 18446744073709551615, F2: 850}
                                                    _ = v97
                                                case 1:
                                                    var v98 uint = 85
                                                    _ = v98
                                                default:
                                                    v20 += v20
                                                    v2 += ((((v2 * (float32(3.14159) + (1.5 - float32(float64((v2 * v2)))))) - v2) / float32(1e10)) / float32((-2.25 * 1.5)))
                                                }
                                                var v99 int8 = 100
                                                v100 := GenS16{F0: 330}
                                                _ = v100
                                                _ = v99
                                            default:
                                                fmt.Println(v20, fi62, fi62)
                                            }
                                        } else {
                                            v2 = (((v2 + (1e10 - v2)) / 1.0) * v2)
                                        }
                                    } else {
                                        var v101 int = (fi63 * (fi62 - (fi62 ^ (fi62 + fi63))))
                                        v2 = v2
                                        _ = v101
                                    }
                                    var v102 uint = 848
                                    _ = v102
                                }
                                _ = fi63
                                _ = v64
                            }
                            if true {
                                v3[0] = v24
                            }
                            v103 := GenS11{F0: 65535, F1: 1}
                            _ = fi62
                            _ = v103
                        }
                        var v104 int64 = -529
                        v104 = 58
                        _ = v104
                    } else {
                        var v105 int64 = -9223372036854775808
                        v106 := GenS16{F0: 347}
                        v20 = v20
                        _ = v105
                        _ = v106
                    }
                    var v107 string = "a"
                    _ = v107
                }
                _ = v24
            }
        case "":
            v2 = (1.5 - v2)
        case "a":
            for fi108 := 0; fi108 < 16; fi108++ {
                var v109 int64 = -1
                _ = fi108
                _ = v109
            }
        }
    }
    var v110 uint8 = 68
    switch ("" + (("" + "\n\t") + ("" + "a"))) {
    case "hello":
        var v111 int8 = 1
        _ = v111
    case "":
        for fi112 := 0; fi112 < 5; fi112++ {
            for fi113 := 0; fi113 < 9; fi113++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v20, fi113, fi113)
                _ = fi113
            }
            var v114 string = ("" + "")
            v115 := GenS1{F0: -770, F1: 2}
            _ = fi112
            _ = v114
            _ = v115
        }
    case "a":
        switch "hello" {
        case "a":
            for fi116 := 0; fi116 < 1; fi116++ {
                var v117 bool = ((0.1 + (float64(v2) / -2.25)) >= 1.5)
                _ = fi116
                _ = v117
            }
        default:
            v110 -= v110
        }
    default:
        _ = len(v3)
    }
    v3[1] = ((false && ((!false) || false)) || (-74 >= 0))
    for fi118 := 0; fi118 < 13; fi118++ {
        v119 := v3[42]
        for fi120 := 0; fi120 < 12; fi120++ {
            v20 = v20
            var v121 int = -9223372036854775808
            _ = fi120
            _ = v121
        }
        _ = fi118
        _ = v119
    }
    if len(v1) > 0 {
        v1[0] = (("foo bar" + ("a" + "hello")) + "foo bar")
    }
    v20 -= v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v110
    _ = v14
    _ = v2
    _ = v20
    _ = v3
}
