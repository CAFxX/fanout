package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS9 struct { F0 int32; F1 string; F2 string }
type GenS10 struct { F0 int32; F1 uint16; F2 string }
type GenS11 struct { F0 int64; F1 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int64; F1 uint64; F2 int }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 uint8; F1 int64; F2 int8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS18 struct { F0 bool; F1 uint64; F2 uint }
type GenS19 struct { F0 uint32; F1 uint }
type GenS20 struct { F0 bool; F1 uint32; F2 int32 }
type GenS21 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS22 struct { F0 uint8 }
type GenS23 struct { F0 int32; F1 uint32; F2 string }
type GenS24 struct { F0 bool; F1 int64; F2 uint16 }
type GenS25 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS26 struct { F0 uint8; F1 uint }
type GenS27 struct { F0 string; F1 uint64 }
type GenS28 struct { F0 bool }
type GenS29 struct { F0 uint16; F1 uint32 }
type GenS30 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS31 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS32 struct { F0 string; F1 uint64; F2 bool }
type GenS33 struct { F0 uint; F1 int64; F2 bool }
type GenS34 struct { F0 int8; F1 uint64 }
type GenS35 struct { F0 uint64 }
type GenS36 struct { F0 int64 }
func genFn0(p0 bool) bool {
    p0 = false
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 = true
        _ = fi6
    }
    p0 = true
    _ = p0
    return p0
}
func genFn1(p0 bool) uint8 {
    v7 := GenS1{F0: 127, F1: -1, F2: 18446744073709551615}
    v8 := GenS1{F0: 100, F1: -119, F2: 1}
    p0 = (!p0)
    if true {
        switch "\n\t" {
        case "":
            fmt.Println(p0)
        case "\n\t":
            if p0 {
                v9 := []int{32768}
                if len(v9) > 0 {
                    v9[0] = -645
                }
                _ = v9
            }
        case "foo bar":
            switch -2147483648 {
            case 0:
                if ((float32(3.14159) - 3.14159) < 3.14159) {
                    switch (("" + ("hello" + "")) + (("\n\t" + ("hello" + "")) + (("hello" + ("hello" + "")) + "foo bar"))) {
                    case "":
                        var v10 int16 = -430
                        _ = v10
                    case "hello":
                        for fi11 := 0; fi11 < 13; fi11++ {
                            p0 = p0
                            var v12 uint16 = 32767
                            _ = fi11
                            _ = v12
                        }
                    case "a":
                        var v13 int = -2147483648
                        _ = v13
                    }
                } else {
                    var v14 int32 = 2147483647
                    _ = v14
                }
            default:
                if p0 {
                    var v15 int32 = 42
                    var v16 float64 = 0.0
                    v17 := []int{}
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        p0 = (true && true)
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v20 := []int{}
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = len(v20)
                            _ = fi19
                            _ = v20
                        }
                        _ = fi18
                    }
                }
            }
        }
        if true {
            if true {
                fmt.Println(p0)
            }
        } else {
            var v21 uint32 = 42
            var v22 int32 = 1
            v21 -= v21
            _ = v21
            _ = v22
        }
    } else {
        switch 32768 {
        case 128:
            for fi23 := 0; fi23 < 11; fi23++ {
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 uint64 = 18446744073709551615
                    _ = fi24
                    _ = v25
                }
                if (p0 || (p0 && false)) {
                    p0 = p0
                    p0 = (!(float32((3.14159 + 3.14159)) >= 0.0))
                } else {
                    p0 = (((true && (((p0 && false) || (("\n\t" + "a") >= "\n\t")) && false)) && false) || true)
                }
                _ = fi23
            }
        }
        if (p0 || (!(42 != 147))) {
            v26 := GenS1{F0: 1, F1: -25, F2: 42}
            p0 = (((!(565 <= 65535)) && p0) && p0)
            p0 = false
            _ = v26
        } else {
            for fi27 := 0; fi27 < 17; fi27++ {
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 uint = 1
                    if false {
                        switch (fi27 % ((fi28) | 1)) {
                        case 128:
                            v29 -= 2
                        default:
                            v30 := []int{(fi28 | fi28), fi27}
                            _ = v30
                        }
                        if p0 {
                            v31 := []int{(fi27 >> ((4294967295) & 63)), 128}
                            v29 -= v29
                            _ = len(v31)
                            _ = v31
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                if p0 {
                                    p0 = true
                                    v29 = 891
                                    v32 := []int{-1, fi28, (fi28 << (((^(v29))) & 63)), fi27}
                                    _ = v32
                                }
                            default:
                                if ((1 > v29) && p0) {
                                    var v33 int64 = 1
                                    _ = v33
                                } else {
                                    if (0 <= 80) {
                                        v34 := []int{fi28, (fi27 * fi27), (fi27 & 42), (^(fi28))}
                                        _ = v34
                                    }
                                    var v35 uint8 = 255
                                    _ = v35
                                }
                            }
                            v29 += v29
                        }
                        if p0 {
                            if true {
                                if true {
                                    switch -1 {
                                    case 32768:
                                        var v36 uint64 = 262
                                        _ = v36
                                    case 901:
                                        var v37 string = "hello"
                                        _ = v37
                                    default:
                                        v29 -= v29
                                    }
                                    var v38 int = fi28
                                    _ = v38
                                } else {
                                    v39 := GenS1{F0: 42, F1: -128, F2: v29}
                                    v40 := []int{(fi28 - (-(fi28))), fi28, (-(fi27))}
                                    _ = v39
                                    _ = v40
                                }
                                v29 = 42
                                if true {
                                    v41 := []int{}
                                    if len(v41) > 0 {
                                        v41[0] = fi28
                                    }
                                    switch "a" {
                                    case "":
                                        if false {
                                            var v42 int16 = 32767
                                            _ = v42
                                        }
                                    case "hello":
                                        v41 = append(v41, (fi28 ^ fi28))
                                    case "\n\t":
                                        fmt.Println(fi28)
                                    }
                                    _ = v41
                                } else {
                                    v43 := GenS1{F0: 42, F1: 71, F2: v29}
                                    _ = v43
                                }
                                p0 = false
                                fmt.Println(v29, fi27)
                                v44 := []int{(fi28 | (fi27 / (((fi27 & 9223372036854775807)) | 1)))}
                                var v45 uint64 = uint64(v29)
                                _ = v44
                                _ = v45
                            }
                            v46 := []int{-504}
                            v47 := GenS1{F0: 0, F1: -21, F2: v29}
                            switch "a" {
                            case "foo bar":
                                fmt.Println(fi28)
                            }
                            _ = v46
                            _ = v47
                        }
                        v48 := GenS1{F0: 127, F1: 100, F2: (^(v29))}
                        _ = v48
                    } else {
                        v29 = v29
                        if p0 {
                            if (true && (9223372036854775807 != -127)) {
                                v49 := []int{(fi27 - (fi28 % ((fi27) | 1))), (fi27 | (fi28 << ((v29) & 63))), (fi28 * fi27)}
                                _ = v49
                            } else {
                                if p0 {
                                    var v50 bool = p0
                                    v51 := GenS2{F0: -1}
                                    var v52 int = fi28
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 int = (fi28 % ((4294967295) | 1))
                                    fmt.Println(p0, fi27)
                                    fmt.Println(v29, v53)
                                    if ((0 > 0) || ((18446744073709551615 > 513) && (!((2147483647 < 2147483648) || (42 >= 42))))) {
                                        switch "hello" {
                                        case "\n\t":
                                            if ((("a" + ("" + "\n\t")) + "\n\t") > "foo bar") {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v54 := GenS1{F0: int8(v29), F1: 100, F2: v29}
                                                    _ = v54
                                                case "a":
                                                    fmt.Println(v53)
                                                case "":
                                                    v55 := []int{(fi27 | fi27), (fi28 - 128), v53}
                                                    _ = v55
                                                }
                                                fmt.Println(v29)
                                            } else {
                                                v56 := GenS1{F0: 100, F1: 100, F2: v29}
                                                _ = v56
                                            }
                                        case "a":
                                            if ((v53 % (((-(v53))) | 1)) <= (fi28 >> ((v29) & 63))) {
                                                p0 = (184 != 87)
                                                var v57 int8 = 1
                                                _ = v57
                                            } else {
                                                v29 += v29
                                                v58 := []int{(fi28 ^ (fi27 >> (((v29 ^ v29)) & 63)))}
                                                _ = v58
                                            }
                                        case "":
                                            p0 = false
                                        default:
                                            if (1 < 255) {
                                                if ((true || ((v29 / ((261) | 1)) == 872)) && ((v53 % (((fi27 >> ((463) & 63))) | 1)) >= 0)) {
                                                    if p0 {
                                                        var v59 int16 = -304
                                                        var v60 int8 = 100
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        v61 := []int{fi28}
                                                        if len(v61) > 0 {
                                                            v61[0] = (fi27 ^ 0)
                                                        }
                                                        _ = v61
                                                    }
                                                    v62 := []int{-2147483648, fi28, fi28}
                                                    if (((0.0 * (0.1 / 1.0)) * (3.14159 * ((1e10 + 3.14159) / 1.0))) < 3.14159) {
                                                        v63 := GenS1{F0: -62, F1: 127, F2: (v29 / ((v29) | 1))}
                                                        _ = v63
                                                    }
                                                    _ = v62
                                                }
                                                p0 = ((v29 >> (((v29 & 1)) & 63)) > v29)
                                                var v64 int16 = -32768
                                                _ = v64
                                            } else {
                                                var v65 uint32 = 113
                                                _ = v65
                                            }
                                        }
                                    } else {
                                        if (521 >= 1000) {
                                            var v66 int32 = 1
                                            switch "\n\t" {
                                            case "hello":
                                                if false {
                                                    fmt.Println(v29)
                                                    switch (("\n\t" + ("hello" + ("\n\t" + "hello"))) + "a") {
                                                    case "":
                                                        if (100 >= 101) {
                                                            var v67 bool = (4294967295 < 330)
                                                            v66 = 495
                                                            _ = v67
                                                        } else {
                                                            v29 = 65535
                                                            p0 = (0 < -1)
                                                            if p0 {
                                                                p0 = (p0 || (v66 == (v66 >> (((v29 >> ((2) & 63))) & 63))))
                                                                if (p0 && true) {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v53 = (v53 + fi27)
                                                                    case "hello":
                                                                        v66 = -2147483648
                                                                    default:
                                                                        p0 = true
                                                                    }
                                                                } else {
                                                                    v68 := []int{(fi27 % (((v53 >> ((65535) & 63))) | 1)), (v53 >> ((1) & 63)), (v53 ^ (-(v53))), fi28}
                                                                    _ = v68
                                                                }
                                                                fmt.Println(p0, p0, v29)
                                                            } else {
                                                                switch (v53 * fi28) {
                                                                case 127:
                                                                    if ((((9223372036854775807 > -9223372036854775808) && p0) && p0) && true) {
                                                                        if p0 {
                                                                            var v69 uint8 = 127
                                                                            p0 = p0
                                                                            v29 = (v29 | (^(v29)))
                                                                            if p0 {
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v70 := []int{}
                                                                                    if (((p0 && (42 == -9223372036854775808)) || (-2.25 != (((3.14159 / float64(1.5)) * 1.5) - 0.0))) && p0) {
                                                                                        v66 += 2147483647
                                                                                        if (v69 != (v69 % ((v69) | 1))) {
                                                                                            var v71 float64 = 3.14159
                                                                                            v72 := GenS2{F0: 42}
                                                                                            v29 += (v29 & (v29 ^ (v29 * (v29 - 97))))
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                        }
                                                                                        if (("hello" + "a") == "foo bar") {
                                                                                            if len(v70) >= 2 {
                                                                                                v70 = v70[1:len(v70)]
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if p0 {
                                                                                            v53 = fi28
                                                                                            if len(v70) >= 2 {
                                                                                                v70 = v70[1:len(v70)]
                                                                                            }
                                                                                            v29 += v29
                                                                                            v73 := GenS3{F0: v29}
                                                                                            if len(v70) > 0 {
                                                                                                v70[0] = (^(fi27))
                                                                                            }
                                                                                            _ = v73
                                                                                        }
                                                                                        v74 := GenS4{F0: 1, F1: v29, F2: 127}
                                                                                        switch fi27 {
                                                                                        case 32767:
                                                                                            v75 := GenS1{F0: 0, F1: 127, F2: (v29 % (((v29 * v29)) | 1))}
                                                                                            _ = v75
                                                                                        case -176:
                                                                                            switch int(v29) {
                                                                                            case -2147483648:
                                                                                                var v76 uint = (v29 - 0)
                                                                                                _ = v76
                                                                                            case 251:
                                                                                                var v77 int16 = 32767
                                                                                                _ = v77
                                                                                            default:
                                                                                                v70 = append(v70, (^(fi28)))
                                                                                            }
                                                                                        default:
                                                                                            var v78 string = "\n\t"
                                                                                            _ = v78
                                                                                        }
                                                                                        _ = v74
                                                                                    }
                                                                                    fmt.Println(v29)
                                                                                    _ = v70
                                                                                }
                                                                                var v79 string = "foo bar"
                                                                                _ = v79
                                                                            } else {
                                                                                v80 := GenS5{F0: 0}
                                                                                _ = v80
                                                                            }
                                                                            switch fi28 {
                                                                            case 322:
                                                                                fmt.Println(v29, fi27, p0)
                                                                            case 32768:
                                                                                switch 280 {
                                                                                case -2147483648:
                                                                                    switch (fi27 / ((fi27) | 1)) {
                                                                                    case 32768:
                                                                                        v81 := []int{(-(fi28)), fi27, 0}
                                                                                        _ = v81
                                                                                    case -801:
                                                                                        v82 := GenS6{F0: -260}
                                                                                        _ = v82
                                                                                    default:
                                                                                        var v83 float64 = 3.14159
                                                                                        _ = v83
                                                                                    }
                                                                                }
                                                                                if (42 == 9223372036854775807) {
                                                                                    fmt.Println(p0, v29)
                                                                                }
                                                                            }
                                                                            v53 = fi28
                                                                            var v84 float64 = 1.5
                                                                            _ = v69
                                                                            _ = v84
                                                                        } else {
                                                                            var v85 int16 = -32768
                                                                            _ = v85
                                                                        }
                                                                        switch (-(v53)) {
                                                                        case -9223372036854775808:
                                                                            var v86 int = (-(fi28))
                                                                            _ = v86
                                                                        }
                                                                        var v87 uint16 = 65535
                                                                        var v88 int64 = 1
                                                                        var v89 uint8 = 0
                                                                        _ = v87
                                                                        _ = v88
                                                                        _ = v89
                                                                    } else {
                                                                        var v90 uint32 = 65535
                                                                        _ = v90
                                                                    }
                                                                case 1:
                                                                    var v91 int64 = -1
                                                                    _ = v91
                                                                case -1:
                                                                    var v92 int32 = v66
                                                                    _ = v92
                                                                default:
                                                                    v93 := GenS7{F0: int16(v66), F1: 18446744073709551615}
                                                                    _ = v93
                                                                }
                                                                if (false && (v29 > v29)) {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if (875 < 287) {
                                                                            fmt.Println(p0, v29, v29)
                                                                            var v94 int64 = int64(v29)
                                                                            _ = v94
                                                                        }
                                                                    case "foo bar":
                                                                        var v95 int = (v53 / ((fi28) | 1))
                                                                        _ = v95
                                                                    case "\n\t":
                                                                        if (-2.25 != (1.5 - -2.25)) {
                                                                            v66 = 0
                                                                            var v96 bool = (!p0)
                                                                            var v97 uint32 = 585
                                                                            _ = v96
                                                                            _ = v97
                                                                        } else {
                                                                            if p0 {
                                                                                p0 = false
                                                                                if (0.0 == 0.1) {
                                                                                    fmt.Println(fi28)
                                                                                    v29 -= (v29 | (^(v29)))
                                                                                }
                                                                                v98 := GenS8{F0: 530, F1: 1, F2: 127}
                                                                                _ = v98
                                                                            } else {
                                                                                v99 := []int{}
                                                                                _ = v99
                                                                            }
                                                                            v66 = 42
                                                                            v53 -= -2147483648
                                                                            var v100 float32 = 0.0
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                    if p0 {
                                                                        v101 := []int{(-(fi27)), fi27, (^(fi28))}
                                                                        switch ("hello" + "hello") {
                                                                        case "a":
                                                                            if p0 {
                                                                                v101 = append(v101, fi28)
                                                                            } else {
                                                                                var v102 int16 = -32768
                                                                                v66 = (^(v66))
                                                                                if len(v101) > 0 {
                                                                                    v101[0] = v53
                                                                                }
                                                                                v103 := GenS9{F0: 1, F1: "hello", F2: ("" + ("foo bar" + "a"))}
                                                                                var v104 uint64 = 42
                                                                                _ = v102
                                                                                _ = v103
                                                                                _ = v104
                                                                            }
                                                                        default:
                                                                            var v105 float32 = (1.5 - float32((0.1 - float64(3.14159))))
                                                                            _ = v105
                                                                        }
                                                                        var v106 uint64 = 430
                                                                        _ = v101
                                                                        _ = v106
                                                                    }
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        if false {
                                                                            v107 := GenS2{F0: v66}
                                                                            _ = v107
                                                                        }
                                                                    case "a":
                                                                        if (-2.25 == (1e10 / -2.25)) {
                                                                            switch "" {
                                                                            case "a":
                                                                                if true {
                                                                                    switch 32768 {
                                                                                    case 473:
                                                                                        v29 = v29
                                                                                    case 42:
                                                                                        if (("\n\t" + ("\n\t" + ("a" + ("\n\t" + "a")))) < ((("foo bar" + ((("" + (("" + "\n\t") + ("hello" + ("hello" + "")))) + ("" + "foo bar")) + "foo bar")) + "") + "hello")) {
                                                                                            if false {
                                                                                                switch 128 {
                                                                                                case 42:
                                                                                                    v66 -= v66
                                                                                                case -2147483648:
                                                                                                    v108 := GenS10{F0: v66, F1: 0, F2: ("a" + "")}
                                                                                                    _ = v108
                                                                                                }
                                                                                                fmt.Println(v53, v53, fi28)
                                                                                                v109 := []int{(-(v53)), v53, v53}
                                                                                                _ = v109
                                                                                            }
                                                                                            v110 := GenS7{F0: 1000, F1: 80}
                                                                                            fmt.Println(v66, fi28)
                                                                                            _ = v110
                                                                                        }
                                                                                    case 2147483647:
                                                                                        p0 = p0
                                                                                    default:
                                                                                        var v111 uint = (v29 | v29)
                                                                                        _ = v111
                                                                                    }
                                                                                } else {
                                                                                    p0 = (3.14159 == 3.14159)
                                                                                    var v112 string = ("hello" + "hello")
                                                                                    _ = v112
                                                                                }
                                                                            default:
                                                                                var v113 uint = (^(v29))
                                                                                _ = v113
                                                                            }
                                                                        } else {
                                                                            fmt.Println(p0)
                                                                            v66 -= 42
                                                                            v29 = v29
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        if (((v66 >> (((v29 | v29)) & 63)) <= int32(v53)) && p0) {
                                                            var v114 uint8 = 85
                                                            v66 += -911
                                                            var v115 int16 = 42
                                                            var v116 float64 = 1.5
                                                            v115 = v115
                                                            _ = v114
                                                            _ = v115
                                                            _ = v116
                                                        }
                                                    case "hello":
                                                        if p0 {
                                                            var v117 int16 = -1
                                                            fmt.Println(v53, fi28, fi27)
                                                            _ = v117
                                                        } else {
                                                            if p0 {
                                                                v66 = (-(v66))
                                                                v118 := []int{(^(fi27)), 2147483647, 430}
                                                                switch ("" + "foo bar") {
                                                                case "a":
                                                                    p0 = true
                                                                case "hello":
                                                                    v53 = v53
                                                                case "":
                                                                    v119 := GenS4{F0: (v66 / (((v66 | 82)) | 1)), F1: 0, F2: 127}
                                                                    _ = v119
                                                                default:
                                                                    if false {
                                                                        p0 = false
                                                                    }
                                                                }
                                                                _ = v118
                                                            }
                                                            var v120 int64 = 1
                                                            _ = v120
                                                        }
                                                    default:
                                                        if (18446744073709551615 < 967) {
                                                            var v121 int8 = 127
                                                            p0 = true
                                                            _ = v121
                                                        } else {
                                                            v29 -= 568
                                                            switch "foo bar" {
                                                            case "a":
                                                                switch "" {
                                                                case "hello":
                                                                    v122 := []int{fi28}
                                                                    _ = v122
                                                                default:
                                                                    if p0 {
                                                                        if (p0 && p0) {
                                                                            var v123 int16 = -225
                                                                            v124 := GenS2{F0: v66}
                                                                            if (255 != 820) {
                                                                                v125 := []int{}
                                                                                var v126 float64 = ((3.14159 / float64(1.5)) / (1e10 - 3.14159))
                                                                                _ = v125
                                                                                _ = v126
                                                                            }
                                                                            if p0 {
                                                                                v29 = 958
                                                                            } else {
                                                                                if (!(1e10 <= 3.14159)) {
                                                                                    var v127 uint16 = 106
                                                                                    fmt.Println(fi28)
                                                                                    v128 := GenS11{F0: -9223372036854775808, F1: 42}
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                }
                                                                                if p0 {
                                                                                    v129 := []int{919, (fi28 >> ((v29) & 63)), (-(fi28)), 9223372036854775807}
                                                                                    if (4294967295 < 65535) {
                                                                                        if len(v129) > 0 {
                                                                                            v129[0] = (^(v53))
                                                                                        }
                                                                                        v130 := GenS8{F0: 578, F1: 1, F2: int8(v123)}
                                                                                        v29 = 916
                                                                                        _ = v130
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            if (true || (p0 || (false && p0))) {
                                                                                                if false {
                                                                                                    v131 := GenS10{F0: 1, F1: 0, F2: "foo bar"}
                                                                                                    _ = v131
                                                                                                } else {
                                                                                                    v66 = (v66 / ((-1) | 1))
                                                                                                }
                                                                                            } else {
                                                                                                if p0 {
                                                                                                    fmt.Println(v66, v123, v53)
                                                                                                    var v132 int64 = -593
                                                                                                    switch (fi27 % ((128) | 1)) {
                                                                                                    case 127:
                                                                                                        if p0 {
                                                                                                            var v133 bool = p0
                                                                                                            var v134 int8 = -81
                                                                                                            _ = v133
                                                                                                            _ = v134
                                                                                                        } else {
                                                                                                            v129 = append(v129, fi27)
                                                                                                        }
                                                                                                        v53 = 42
                                                                                                        var v135 uint64 = 0
                                                                                                        switch ("hello" + "") {
                                                                                                        case "\n\t":
                                                                                                            switch fi28 {
                                                                                                            case -9223372036854775808:
                                                                                                                if (v29 != v29) {
                                                                                                                    v136 := GenS12{F0: -1}
                                                                                                                    _ = v136
                                                                                                                } else {
                                                                                                                    var v137 int64 = (v132 * (v132 * int64(v29)))
                                                                                                                    _ = v137
                                                                                                                }
                                                                                                            default:
                                                                                                                fmt.Println(fi27, v123)
                                                                                                            }
                                                                                                        case "a":
                                                                                                            var v138 float64 = -2.25
                                                                                                            _ = v138
                                                                                                        }
                                                                                                        _ = v135
                                                                                                    case -9223372036854775808:
                                                                                                        if p0 {
                                                                                                            v29 -= (^(v29))
                                                                                                            if (!p0) {
                                                                                                                if (!p0) {
                                                                                                                    v139 := GenS5{F0: 32767}
                                                                                                                    switch fi27 {
                                                                                                                    case 314:
                                                                                                                        if p0 {
                                                                                                                            v140 := GenS13{F0: (-(v132)), F1: 923, F2: (v53 & (fi28 & (fi28 ^ (fi28 - fi28))))}
                                                                                                                            _ = v140
                                                                                                                        } else {
                                                                                                                            var v141 uint16 = 422
                                                                                                                            fmt.Println(v141)
                                                                                                                            _ = v141
                                                                                                                        }
                                                                                                                    case -609:
                                                                                                                        p0 = p0
                                                                                                                    }
                                                                                                                    v53 -= (-(fi27))
                                                                                                                    _ = v139
                                                                                                                }
                                                                                                                v132 += (v132 ^ (v132 - v132))
                                                                                                            }
                                                                                                            v142 := GenS14{F0: fi28, F1: 2147483648}
                                                                                                            _ = v142
                                                                                                        }
                                                                                                        var v143 int = -9223372036854775808
                                                                                                        v144 := GenS15{F0: 69, F1: (^(v132)), F2: -33}
                                                                                                        v123 = 42
                                                                                                        var v145 float32 = -2.25
                                                                                                        if (!(v29 > (v29 + (v29 * (v29 & v29))))) {
                                                                                                            var v146 int32 = -715
                                                                                                            v143 -= (fi27 | (v143 - v53))
                                                                                                            _ = v146
                                                                                                        } else {
                                                                                                            fmt.Println(v123)
                                                                                                            var v147 bool = ((2 > 714) && p0)
                                                                                                            _ = v147
                                                                                                        }
                                                                                                        _ = v143
                                                                                                        _ = v144
                                                                                                        _ = v145
                                                                                                    case 32767:
                                                                                                        v123 = (v123 | (^(v123)))
                                                                                                    }
                                                                                                    _ = v132
                                                                                                } else {
                                                                                                    var v148 uint32 = 2147483648
                                                                                                    _ = v148
                                                                                                }
                                                                                            }
                                                                                        case "\n\t":
                                                                                            fmt.Println(v123)
                                                                                        case "hello":
                                                                                            v149 := GenS12{F0: 2147483647}
                                                                                            _ = v149
                                                                                        default:
                                                                                            v150 := GenS7{F0: v123, F1: 18446744073709551615}
                                                                                            _ = v150
                                                                                        }
                                                                                    }
                                                                                    var v151 float32 = 3.14159
                                                                                    fmt.Println(fi27)
                                                                                    _ = v129
                                                                                    _ = v151
                                                                                } else {
                                                                                    v152 := []int{(fi27 / (((fi28 >> ((v29) & 63))) | 1)), (fi28 | (fi27 & (v53 * 42))), 42}
                                                                                    var v153 uint8 = 132
                                                                                    if len(v152) > 0 {
                                                                                        v152[0] = 4294967295
                                                                                    }
                                                                                    _ = v152
                                                                                    _ = v153
                                                                                }
                                                                                v154 := GenS16{F0: 9223372036854775807}
                                                                                switch ("foo bar" + "foo bar") {
                                                                                case "hello":
                                                                                    var v155 int16 = v123
                                                                                    if false {
                                                                                        fmt.Println(fi28, v155)
                                                                                        v156 := GenS17{F0: 1, F1: (v66 * v66), F2: 42}
                                                                                        v155 -= -562
                                                                                        _ = v156
                                                                                    } else {
                                                                                        v157 := GenS5{F0: -32768}
                                                                                        v158 := GenS18{F0: false, F1: 4294967295, F2: 42}
                                                                                        _ = v157
                                                                                        _ = v158
                                                                                    }
                                                                                    v159 := []int{}
                                                                                    _ = v155
                                                                                    _ = v159
                                                                                case "foo bar":
                                                                                    var v160 int32 = 2147483647
                                                                                    _ = v160
                                                                                case "a":
                                                                                    var v161 bool = p0
                                                                                    v161 = (-1 < 100)
                                                                                    v162 := GenS17{F0: 2147483647, F1: v66, F2: 419}
                                                                                    _ = v161
                                                                                    _ = v162
                                                                                default:
                                                                                    if (p0 && (2 < 273)) {
                                                                                        v66 = (v66 & v66)
                                                                                        var v163 int64 = -9223372036854775808
                                                                                        _ = v163
                                                                                    }
                                                                                }
                                                                                var v164 int8 = 100
                                                                                _ = v154
                                                                                _ = v164
                                                                            }
                                                                            _ = v123
                                                                            _ = v124
                                                                        }
                                                                        switch fi28 {
                                                                        case -183:
                                                                            switch fi28 {
                                                                            case 2147483647:
                                                                                var v165 string = "foo bar"
                                                                                _ = v165
                                                                            case 128:
                                                                                if p0 {
                                                                                    v166 := []int{}
                                                                                    if p0 {
                                                                                        var v167 uint16 = 497
                                                                                        var v168 string = ""
                                                                                        switch v168 {
                                                                                        case "\n\t":
                                                                                            v167 += v167
                                                                                        case "a":
                                                                                            var v169 int64 = 42
                                                                                            _ = v169
                                                                                        case "hello":
                                                                                            v166 = append(v166, v53)
                                                                                        default:
                                                                                            v168 = v168
                                                                                        }
                                                                                        if (!true) {
                                                                                            var v170 float64 = 3.14159
                                                                                            _ = v170
                                                                                        } else {
                                                                                            fmt.Println(v29, v66)
                                                                                            v66 -= (v66 * (v66 ^ v66))
                                                                                            v171 := GenS19{F0: 0, F1: (^(v29))}
                                                                                            v168 = ""
                                                                                            _ = v171
                                                                                        }
                                                                                        _ = v167
                                                                                        _ = v168
                                                                                    }
                                                                                    v29 += (^(v29))
                                                                                    _ = v166
                                                                                } else {
                                                                                    if p0 {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            if (p0 && (-211 <= 333)) {
                                                                                                v172 := []int{(fi28 + 127)}
                                                                                                if p0 {
                                                                                                    fmt.Println(p0, v53)
                                                                                                    v172 = append(v172, fi27)
                                                                                                } else {
                                                                                                    v173 := GenS20{F0: (true || p0), F1: 178, F2: v66}
                                                                                                    v53 -= (-(fi28))
                                                                                                    var v174 float64 = -2.25
                                                                                                    v175 := GenS1{F0: -100, F1: 1, F2: v29}
                                                                                                    switch fi27 {
                                                                                                    case -1:
                                                                                                        var v176 int64 = 9223372036854775807
                                                                                                        _ = v176
                                                                                                    default:
                                                                                                        if p0 {
                                                                                                            var v177 uint32 = 2147483648
                                                                                                            v174 = (-2.25 - (v174 / (v174 * (v174 * float64(0.0)))))
                                                                                                            _ = v177
                                                                                                        }
                                                                                                    }
                                                                                                    v66 = v66
                                                                                                    var v178 int8 = -29
                                                                                                    _ = v173
                                                                                                    _ = v174
                                                                                                    _ = v175
                                                                                                    _ = v178
                                                                                                }
                                                                                                _ = v172
                                                                                            } else {
                                                                                                v179 := GenS21{F0: 1, F1: 9223372036854775807, F2: 18446744073709551615}
                                                                                                _ = v179
                                                                                            }
                                                                                        case "a":
                                                                                            p0 = (9223372036854775807 > 0)
                                                                                        case "\n\t":
                                                                                            fmt.Println(v66, p0, fi28)
                                                                                        default:
                                                                                            var v180 int8 = 42
                                                                                            _ = v180
                                                                                        }
                                                                                        var v181 int16 = 32767
                                                                                        fmt.Println(fi28, p0)
                                                                                        _ = v181
                                                                                    }
                                                                                }
                                                                            case 4294967295:
                                                                                var v182 string = "hello"
                                                                                _ = v182
                                                                            default:
                                                                                fmt.Println(p0, v66, fi28)
                                                                            }
                                                                        default:
                                                                            var v183 int16 = -1
                                                                            _ = v183
                                                                        }
                                                                    } else {
                                                                        v29 -= v29
                                                                        var v184 int16 = 32767
                                                                        fmt.Println(fi27)
                                                                        _ = v184
                                                                    }
                                                                }
                                                                v185 := []int{v53, (v53 % (((v53 ^ -2147483648)) | 1))}
                                                                v29 -= (^(v29))
                                                                var v186 int16 = 32767
                                                                v53 += fi27
                                                                if (!(-1 > 171)) {
                                                                    p0 = p0
                                                                    switch -2147483648 {
                                                                    case -788:
                                                                        v187 := GenS22{F0: 127}
                                                                        _ = v187
                                                                    case 42:
                                                                        var v188 int64 = 0
                                                                        _ = v188
                                                                    default:
                                                                        if false {
                                                                            v189 := GenS9{F0: v66, F1: "hello", F2: "a"}
                                                                            v190 := GenS17{F0: 1, F1: (v66 & (-(v66))), F2: 1}
                                                                            _ = v189
                                                                            _ = v190
                                                                        }
                                                                    }
                                                                } else {
                                                                    v186 = v186
                                                                    if p0 {
                                                                        if false {
                                                                            var v191 int16 = 1000
                                                                            _ = v191
                                                                        }
                                                                        var v192 bool = (!p0)
                                                                        _ = v192
                                                                    }
                                                                    v193 := GenS23{F0: int32(v29), F1: 2147483648, F2: "hello"}
                                                                    _ = v193
                                                                }
                                                                _ = v185
                                                                _ = v186
                                                            case "\n\t":
                                                                if (!p0) {
                                                                    var v194 int8 = -128
                                                                    switch ("foo bar" + "") {
                                                                    case "hello":
                                                                        fmt.Println(v66, p0, v194)
                                                                    case "a":
                                                                        v194 += (^(v194))
                                                                    }
                                                                    _ = v194
                                                                } else {
                                                                    fmt.Println(v66)
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v195 int8 = 100
                                                                        _ = v195
                                                                    default:
                                                                        var v196 uint = (v29 % ((v29) | 1))
                                                                        _ = v196
                                                                    }
                                                                }
                                                            default:
                                                                var v197 string = ("\n\t" + ("\n\t" + ((("foo bar" + "") + "\n\t") + "foo bar")))
                                                                _ = v197
                                                            }
                                                            var v198 int64 = 458
                                                            _ = v198
                                                        }
                                                    }
                                                    v53 = fi28
                                                }
                                            }
                                            if p0 {
                                                v53 -= 0
                                                if ((v66 % ((v66) | 1)) == v66) {
                                                    if (58 == 42) {
                                                        v53 += 42
                                                    } else {
                                                        v199 := GenS24{F0: (-1 >= int8(v29)), F1: -1, F2: 88}
                                                        _ = v199
                                                    }
                                                    p0 = ((3.14159 == ((((-2.25 - 1.5) * (0.1 - (1e10 - float64(0.0)))) / (3.14159 * float64(float32(((-2.25 + 1.5) / 3.14159))))) - 3.14159)) || false)
                                                }
                                                if false {
                                                    v200 := []int{}
                                                    if p0 {
                                                        p0 = p0
                                                        var v201 bool = p0
                                                        var v202 uint32 = 4294967295
                                                        _ = v201
                                                        _ = v202
                                                    }
                                                    if len(v200) >= 2 {
                                                        v200 = v200[1:len(v200)]
                                                    }
                                                    _ = v200
                                                } else {
                                                    var v203 uint64 = 790
                                                    if true {
                                                        fmt.Println(v203, p0, fi28)
                                                    } else {
                                                        if (v66 < v66) {
                                                            v53 = fi27
                                                        }
                                                        var v204 uint8 = 42
                                                        var v205 string = ("hello" + "foo bar")
                                                        _ = v204
                                                        _ = v205
                                                    }
                                                    _ = v203
                                                }
                                            }
                                            v206 := []int{fi28}
                                            _ = v206
                                            _ = v66
                                        } else {
                                            v29 += (v29 & v29)
                                            v207 := GenS6{F0: 1}
                                            if (!(0.0 < 3.14159)) {
                                                v53 = (fi27 - 9223372036854775807)
                                                switch (^(fi27)) {
                                                case -1:
                                                    if (1 > 0) {
                                                        switch (fi27 >> ((v29) & 63)) {
                                                        case 32767:
                                                            v208 := []int{}
                                                            _ = v208
                                                        case -588:
                                                            var v209 string = ""
                                                            _ = v209
                                                        case 0:
                                                            var v210 float64 = float64(0.0)
                                                            _ = v210
                                                        }
                                                        switch 9223372036854775807 {
                                                        case 666:
                                                            switch (("hello" + "foo bar") + "hello") {
                                                            case "hello":
                                                                v53 = (fi27 | v53)
                                                            case "foo bar":
                                                                p0 = (!(fi27 >= fi28))
                                                            case "\n\t":
                                                                v53 = 2147483647
                                                            default:
                                                                fmt.Println(p0, v29, fi27)
                                                            }
                                                        case 9223372036854775807:
                                                            var v211 bool = p0
                                                            _ = v211
                                                        case 32768:
                                                            v212 := []int{}
                                                            _ = v212
                                                        }
                                                        var v213 int32 = 2147483647
                                                        _ = v213
                                                    } else {
                                                        if (0 < 0) {
                                                            if p0 {
                                                                var v214 string = "a"
                                                                _ = v214
                                                            }
                                                        }
                                                    }
                                                case -224:
                                                    v29 -= (v29 + 1)
                                                case 42:
                                                    if true {
                                                        v215 := []int{v53, (fi27 - v53), (fi28 | 128), (fi27 | v53)}
                                                        _ = v215
                                                    }
                                                }
                                                var v216 uint16 = 32767
                                                _ = v216
                                            } else {
                                                if (("a" + ("\n\t" + ((("hello" + ("a" + ("hello" + "hello"))) + (("hello" + "a") + ("\n\t" + "a"))) + (("foo bar" + "hello") + "foo bar")))) > ("\n\t" + ("hello" + ("a" + "foo bar")))) {
                                                    v217 := []int{v53, -1}
                                                    var v218 int8 = -107
                                                    _ = v217
                                                    _ = v218
                                                } else {
                                                    v29 = v29
                                                }
                                                switch -2147483648 {
                                                case -388:
                                                    v53 += (fi27 | 818)
                                                }
                                                if p0 {
                                                    v29 += (v29 / (((v29 + v29)) | 1))
                                                } else {
                                                    v29 = v29
                                                    v219 := GenS25{F0: 42, F1: 99, F2: 1}
                                                    _ = v219
                                                }
                                            }
                                            _ = v207
                                        }
                                    }
                                    _ = v53
                                }
                                var v220 uint8 = 42
                                _ = v220
                            }
                            if (true && (18446744073709551615 == 18446744073709551615)) {
                                if p0 {
                                    if p0 {
                                        v29 += v29
                                        if (!true) {
                                            var v221 bool = (!(127 <= 42))
                                            v222 := []int{(fi27 + -2147483648), fi28, fi27}
                                            if p0 {
                                                p0 = ((!(18446744073709551615 != 2)) || (42 > 845))
                                                if len(v222) > 0 {
                                                    v222[0] = (fi28 & 937)
                                                }
                                                v223 := GenS14{F0: (fi28 + fi27), F1: 65535}
                                                _ = v223
                                            } else {
                                                if true {
                                                    var v224 float32 = (3.14159 + (1.5 - ((-2.25 / 3.14159) + float32(3.14159))))
                                                    if (521 != 2) {
                                                        v29 = (v29 ^ v29)
                                                    }
                                                    _ = v224
                                                } else {
                                                    switch "" {
                                                    case "":
                                                        v221 = (1 < 128)
                                                    case "a":
                                                        switch 42 {
                                                        case 9223372036854775807:
                                                            if (2 >= 1) {
                                                                fmt.Println(fi28, fi28, fi27)
                                                                _ = len(v222)
                                                                if len(v222) >= 2 {
                                                                    v222 = v222[1:len(v222)]
                                                                }
                                                            } else {
                                                                p0 = v221
                                                                if (p0 && true) {
                                                                    fmt.Println(p0)
                                                                    switch fi27 {
                                                                    case 1:
                                                                        v225 := GenS8{F0: 42, F1: 194, F2: 127}
                                                                        _ = v225
                                                                    case 2147483647:
                                                                        var v226 string = "\n\t"
                                                                        _ = v226
                                                                    }
                                                                }
                                                                var v227 uint16 = 2
                                                                _ = v227
                                                            }
                                                        case -9223372036854775808:
                                                            v228 := GenS26{F0: 42, F1: (v29 * (v29 | 962))}
                                                            v222 = append(v222, (fi28 % ((-9223372036854775808) | 1)))
                                                            _ = v228
                                                        default:
                                                            if len(v222) >= 2 {
                                                                v222 = v222[1:len(v222)]
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v221
                                            _ = v222
                                        }
                                    }
                                    fmt.Println(fi27)
                                } else {
                                    v229 := GenS24{F0: ((100 > 13) || p0), F1: 9223372036854775807, F2: 0}
                                    _ = v229
                                }
                            }
                        } else {
                            p0 = p0
                            v29 -= 18446744073709551615
                        }
                        v29 = 4294967295
                    }
                    switch fi27 {
                    case 2147483647:
                        v230 := GenS16{F0: 9223372036854775807}
                        _ = v230
                    default:
                        fmt.Println(v29, fi27, fi27)
                        if p0 {
                            if (!p0) {
                                v231 := GenS10{F0: 296, F1: 1, F2: ""}
                                v29 -= v29
                                switch "a" {
                                case "\n\t":
                                    var v232 float64 = float64(1e10)
                                    _ = v232
                                }
                                _ = v231
                            } else {
                                var v233 int64 = -9223372036854775808
                                v234 := GenS3{F0: 18446744073709551615}
                                v235 := GenS4{F0: -2147483648, F1: (^(v29)), F2: 127}
                                _ = v233
                                _ = v234
                                _ = v235
                            }
                            p0 = p0
                        }
                    }
                    _ = fi28
                    _ = v29
                }
                fmt.Println(fi27, fi27)
                var v236 int = fi27
                _ = fi27
                _ = v236
            }
            v237 := GenS27{F0: "\n\t", F1: 42}
            _ = v237
        }
        for fi238 := 0; fi238 < 9; fi238++ {
            for fi239 := 0; fi239 < 17; fi239++ {
                var v240 float32 = 3.14159
                _ = fi239
                _ = v240
            }
            p0 = p0
            _ = fi238
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = -722
    var v2 int16 = -1
    v3 := make(map[string]string)
    var v4 int = -598
    v5 := make(map[int]int)
    v1 = v1
    switch ("foo bar" + ("\n\t" + "\n\t")) {
    case "\n\t":
        v241 := []int{-9223372036854775808, v4}
        _ = v241
    default:
        fmt.Println(v4)
    }
    var v242 string = "\n\t"
    v242 = ((v242 + (("foo bar" + ("" + (v242 + v242))) + v242)) + v242)
    fmt.Println(v2, v2, v242)
    var v243 float32 = 0.1
    var v244 float32 = (3.14159 / ((float32(1e10) - (v243 / 1e10)) + (v243 - -2.25)))
    if (true || true) {
        for fi245 := 0; fi245 < 6; fi245++ {
            v4 = v4
            _ = fi245
        }
    } else {
        if false {
            switch v242 {
            case "foo bar":
                for fi246 := 0; fi246 < 17; fi246++ {
                    for fi247 := 0; fi247 < 6; fi247++ {
                        var v248 float64 = 1e10
                        if ((2 == 65535) || (v248 > (v248 + (0.0 + (1.5 * v248))))) {
                            var v249 int64 = 1
                            v1 += v1
                            switch v242 {
                            case "\n\t":
                                switch fi246 {
                                case 42:
                                    var v250 uint64 = 4294967295
                                    _ = v250
                                case 0:
                                    v244 -= 3.14159
                                }
                            case "":
                                var v251 uint = 18446744073709551615
                                _ = v251
                            default:
                                var v252 bool = (-128 == 127)
                                _ = v252
                            }
                            _ = v249
                        } else {
                            v5[(fi247 + (fi247 % (((fi246 / ((fi247) | 1))) | 1)))] = (^(v4))
                            switch v242 {
                            case "foo bar":
                                if ((!(true || true)) && true) {
                                    v253 := []int{(fi246 ^ (fi247 | v4)), 42, 959, v4}
                                    var v254 uint32 = 77
                                    _ = v253
                                    _ = v254
                                } else {
                                    v2 -= (v2 / (((v2 | 1)) | 1))
                                    v255 := []int{-954, 0}
                                    var v256 float32 = ((3.14159 * ((((v244 - 0.1) * v244) - (float32((v248 - v248)) + v243)) - ((((v244 + v244) - (v244 * float32(v248))) + ((0.1 * float32(v248)) + v244)) + (v244 - v243)))) * (v243 * v244))
                                    _ = v255
                                    _ = v256
                                }
                            case "hello":
                                v257 := v3[(("foo bar" + "a") + "hello")]
                                _ = v257
                            }
                            var v258 uint16 = 783
                            v244 = v244
                            v259 := GenS17{F0: 0, F1: (v1 / ((v1) | 1)), F2: 42}
                            switch v242 {
                            case "a":
                                var v260 uint64 = 2
                                _ = v260
                            }
                            v258 += 676
                            _ = v258
                            _ = v259
                        }
                        _ = fi247
                        _ = v248
                    }
                    _ = fi246
                }
            case "hello":
                var v261 uint = 2
                _ = v261
            case "":
                var v262 int = 4294967295
                _ = v262
            default:
                fmt.Println(v2, v4, v244)
            }
            v263 := []int{}
            for fi264 := 0; fi264 < 9; fi264++ {
                if false {
                    for fi265 := 0; fi265 < 8; fi265++ {
                        if false {
                            _ = len(v263)
                            switch (^(fi265)) {
                            case 32768:
                                var v266 uint8 = 60
                                _ = v266
                            case -647:
                                switch fi265 {
                                case 32768:
                                    v267 := GenS24{F0: true, F1: 925, F2: 1}
                                    _ = v267
                                case -9223372036854775808:
                                    v242 = v242
                                case 9223372036854775807:
                                    fmt.Println(v242, v1)
                                    if true {
                                        var v268 uint8 = 251
                                        delete(v3, v242)
                                        if true {
                                            var v269 bool = true
                                            v270 := GenS5{F0: (v2 << ((2) & 63))}
                                            _ = v269
                                            _ = v270
                                        }
                                        _ = v268
                                    }
                                }
                            default:
                                var v271 uint16 = 237
                                _ = v271
                            }
                            fmt.Println(v243)
                        } else {
                            if false {
                                var v272 float32 = (((((1e10 + 3.14159) - (v243 + (0.0 - 0.1))) * (v243 + v243)) * ((v243 / v244) * v243)) / float32(1e10))
                                _ = v272
                            } else {
                                fmt.Println(v244)
                            }
                            var v273 int8 = -121
                            delete(v5, v4)
                            _ = v273
                        }
                        switch (v4 & (fi264 % ((840) | 1))) {
                        case 4294967295:
                            var v274 uint = 42
                            var v275 int = fi265
                            v263 = append(v263, v4)
                            _ = v274
                            _ = v275
                        }
                        _ = fi265
                    }
                }
                fmt.Println(v242)
                _ = fi264
            }
            _ = v263
        } else {
            switch v242 {
            case "":
                for fi276 := 0; fi276 < 16; fi276++ {
                    for fi277 := 0; fi277 < 6; fi277++ {
                        switch v242 {
                        case "":
                            fmt.Println(v2, fi276)
                        case "\n\t":
                            v243 = v243
                        default:
                            v278 := v3[v242]
                            v2 -= v2
                            v278 = "\n\t"
                            _ = v278
                        }
                        _ = fi277
                    }
                    v279 := []int{(fi276 << ((2) & 63)), -721}
                    _ = fi276
                    _ = v279
                }
            case "foo bar":
                var v280 uint = 0
                _ = v280
            case "hello":
                v4 = (v4 * 1)
            default:
                v281 := GenS28{F0: (1 >= 77)}
                _ = v281
            }
        }
        v242 = v242
    }
    fmt.Println(v1)
    for fi282 := 0; fi282 < 15; fi282++ {
        fmt.Println(v2, v4, v1)
        if (true || (0 < 760)) {
            var v283 uint8 = 255
            switch 235 {
            case 4294967295:
                var v284 int8 = int8(v2)
                _ = v284
            case -383:
                v285 := GenS29{F0: 971, F1: 264}
                _ = v285
            case -367:
                v244 = -2.25
            default:
                switch fi282 {
                case 127:
                    var v286 int64 = 732
                    _ = v286
                case 84:
                    var v287 uint16 = 2
                    _ = v287
                default:
                    var v288 int = (fi282 - 4294967295)
                    _ = v288
                }
            }
            _ = v283
        } else {
            v289 := GenS28{F0: false}
            switch v242 {
            case "hello":
                switch (v242 + v242) {
                case "a":
                    v3[v242] = (v242 + ((("foo bar" + "a") + "a") + v242))
                case "foo bar":
                    v290 := []int{-1, (fi282 / ((v4) | 1)), 1}
                    _ = v290
                case "":
                    var v291 uint32 = 2147483648
                    _ = v291
                default:
                    var v292 int32 = int32(fi282)
                    _ = v292
                }
            case "foo bar":
                for fi293 := 0; fi293 < 13; fi293++ {
                    v2 = (v2 % ((v2) | 1))
                    var v294 uint16 = 42
                    if (32768 < 9223372036854775807) {
                        fmt.Println(v242)
                        var v295 int32 = v1
                        _ = v295
                    } else {
                        v1 = 642
                        var v296 float64 = 0.1
                        if (false || true) {
                            var v297 uint = 255
                            switch "" {
                            case "a":
                                v298 := GenS30{F0: v294, F1: 211, F2: v1}
                                _ = v298
                            case "\n\t":
                                var v299 string = "hello"
                                _ = v299
                            case "hello":
                                if (v243 < (((float32(float64(((v243 / v243) + 1e10))) * (0.1 + v244)) * float32(v296)) - 1e10)) {
                                    v300 := GenS31{F0: (v294 >> ((v297) & 63)), F1: 0, F2: 0}
                                    var v301 uint8 = 0
                                    _ = v300
                                    _ = v301
                                }
                            default:
                                var v302 uint64 = 18446744073709551615
                                _ = v302
                            }
                            var v303 int32 = (v1 + (v1 * v1))
                            _ = v297
                            _ = v303
                        } else {
                            switch (("foo bar" + v242) + "\n\t") {
                            case "foo bar":
                                if false {
                                    var v304 int16 = (v2 << ((785) & 63))
                                    fmt.Println(v242, fi293, v1)
                                    _ = v304
                                }
                            case "\n\t":
                                var v305 int32 = (v1 >> ((255) & 63))
                                _ = v305
                            case "a":
                                switch (v4 * v4) {
                                case 32767:
                                    v243 += float32(v296)
                                case 128:
                                    if (-410 < 42) {
                                        v306 := GenS4{F0: v1, F1: 2, F2: -1}
                                        var v307 int = (v4 & (fi293 << ((18446744073709551615) & 63)))
                                        _ = v306
                                        _ = v307
                                    }
                                default:
                                    var v308 bool = true
                                    _ = v308
                                }
                            default:
                                if (float32(1e10) < (3.14159 / float32((1.5 / v296)))) {
                                    var v309 float32 = (-2.25 + v243)
                                    _ = v309
                                } else {
                                    var v310 float64 = (v296 / float64(((1e10 + 1e10) / (v243 / v243))))
                                    if (100 > 127) {
                                        var v311 uint8 = genFn1(true)
                                        _ = v311
                                    }
                                    v244 -= 0.0
                                    _ = v310
                                }
                            }
                            if (110 > 70) {
                                v242 = ""
                                if (18446744073709551615 >= 4294967295) {
                                    switch 32768 {
                                    case 4294967295:
                                        v244 += v243
                                    case 53:
                                        v312 := GenS11{F0: 1, F1: 270}
                                        _ = v312
                                    case 9223372036854775807:
                                        v313 := v3[v242]
                                        _ = v313
                                    default:
                                        var v314 int8 = 127
                                        _ = v314
                                    }
                                    v315 := []int{fi293, (-(fi293)), (fi282 ^ (fi282 << ((2) & 63))), (fi282 << ((65535) & 63))}
                                    var v316 uint64 = 42
                                    _ = v315
                                    _ = v316
                                }
                                switch v242 {
                                case "\n\t":
                                    v317 := []int{(v4 >> ((2) & 63)), int(v294)}
                                    v317 = append(v317, 2147483647)
                                    var v318 int64 = 42
                                    _ = v317
                                    _ = v318
                                default:
                                    var v319 uint64 = 256
                                    _ = v319
                                }
                            } else {
                                if true {
                                    delete(v5, 2147483647)
                                    var v320 string = "hello"
                                    var v321 int64 = 0
                                    var v322 float64 = v296
                                    _ = v320
                                    _ = v321
                                    _ = v322
                                }
                                switch "hello" {
                                case "":
                                    if true {
                                        switch (v4 * fi293) {
                                        case 42:
                                            v323 := GenS30{F0: (^(v294)), F1: 42, F2: v1}
                                            if true {
                                                var v324 int8 = 42
                                                _ = v324
                                            }
                                            var v325 uint = 0
                                            if (1 < -1) {
                                                v244 += v244
                                                v326 := GenS24{F0: true, F1: -9223372036854775808, F2: v294}
                                                if ((v294 ^ (^(v294))) <= v294) {
                                                    v327 := []int{-2147483648, 42, fi293, 9223372036854775807}
                                                    v328 := GenS32{F0: "foo bar", F1: 397, F2: true}
                                                    switch int(v1) {
                                                    case 9223372036854775807:
                                                        if false {
                                                            var v329 uint8 = 128
                                                            _ = v329
                                                        }
                                                    case -9223372036854775808:
                                                        v330 := GenS9{F0: v1, F1: ("hello" + ""), F2: ((("\n\t" + "a") + (v242 + "foo bar")) + v242)}
                                                        _ = v330
                                                    }
                                                    _ = v327
                                                    _ = v328
                                                }
                                                _ = v326
                                            } else {
                                                var v331 int64 = 42
                                                v332 := GenS33{F0: (v325 - 393), F1: v331, F2: (((v296 + 0.1) <= (-2.25 + v296)) || (true || true))}
                                                var v333 bool = (true && true)
                                                _ = v331
                                                _ = v332
                                                _ = v333
                                            }
                                            _ = v323
                                            _ = v325
                                        case 17:
                                            v334 := GenS34{F0: 0, F1: 868}
                                            _ = v334
                                        default:
                                            var v335 float64 = 3.14159
                                            _ = v335
                                        }
                                        v336 := GenS35{F0: 4294967295}
                                        if (-2.25 != (v243 * v244)) {
                                            v242 = "\n\t"
                                            switch "" {
                                            case "a":
                                                var v337 uint8 = 197
                                                _ = v337
                                            case "foo bar":
                                                _ = len(v5)
                                            }
                                            v338 := GenS2{F0: 0}
                                            _ = v338
                                        }
                                        _ = v336
                                    } else {
                                        v294 = (v294 >> ((648) & 63))
                                        fmt.Println(v1, fi293)
                                    }
                                case "hello":
                                    var v339 int8 = -1
                                    var v340 int32 = 541
                                    v339 = (v339 + (v339 ^ (^(v339))))
                                    _ = v339
                                    _ = v340
                                }
                            }
                        }
                        _ = v296
                    }
                    _ = fi293
                    _ = v294
                }
            default:
                fmt.Println(v243, v242)
            }
            for fi341 := 0; fi341 < 6; fi341++ {
                if (v2 > (v2 * (^(v2)))) {
                    var v342 int = fi341
                    _ = v342
                } else {
                    v4 += fi341
                }
                _ = fi341
            }
            _ = v289
        }
        _ = fi282
    }
    switch (v4 % (((v4 + 128)) | 1)) {
    case 972:
        for fi343 := 0; fi343 < 17; fi343++ {
            fmt.Println(fi343, v242, v243)
            v243 = 3.14159
            _ = fi343
        }
    case 1:
        var v344 uint64 = 4294967295
        _ = v344
    }
    v345 := []int{v4}
    switch int(v2) {
    case -9223372036854775808:
        var v346 int = 128
        _ = v346
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v4)
    case "\n\t":
        delete(v5, (v4 / ((4294967295) | 1)))
    }
    v347 := GenS36{F0: 9223372036854775807}
    switch v4 {
    case 2147483647:
        switch (-(v4)) {
        case 1:
            var v348 string = ("a" + v242)
            _ = v348
        case 9223372036854775807:
            var v349 int8 = -41
            _ = v349
        case -309:
            var v350 bool = false
            _ = v350
        default:
            var v351 int16 = (v2 | (v2 % ((int16(v1)) | 1)))
            _ = v351
        }
    case 829:
        var v352 string = "\n\t"
        _ = v352
    case -808:
        if len(v345) > 0 {
            v345[0] = v4
        }
    }
    v242 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v242:", v242)
    fmt.Println("v243:", v243)
    fmt.Println("v244:", v244)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v345:", len(v345), v345)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v242
    _ = v243
    _ = v244
    _ = v3
    _ = v345
    _ = v347
    _ = v4
    _ = v5
}
