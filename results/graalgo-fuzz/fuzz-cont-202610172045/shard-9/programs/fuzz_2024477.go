package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 int32; F2 uint }
type GenS4 struct { F0 uint64; F1 int8; F2 int64 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 string; F2 int }
type GenS11 struct { F0 int64; F1 int32; F2 string }
type GenS12 struct { F0 bool; F1 uint; F2 uint8 }
type GenS13 struct { F0 bool; F1 uint32; F2 string }
type GenS14 struct { F0 int; F1 int; F2 uint64 }
type GenS15 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS16 struct { F0 int; F1 int64; F2 uint32 }
type GenS17 struct { F0 uint16; F1 int8; F2 int8 }
type GenS18 struct { F0 int64; F1 int32; F2 int64 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int16; F1 int; F2 uint }
type GenS21 struct { F0 uint8; F1 int16; F2 int64 }
type GenS22 struct { F0 uint32 }
type GenS23 struct { F0 int32; F1 int16; F2 bool }
type GenS24 struct { F0 int }
func genFn0(p0 int16, p1 int32, p2 int8) int {
    p1 += p1
    for fi6 := 0; fi6 < 14; fi6++ {
        switch "a" {
        case "hello":
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := GenS1{F0: int8(p1), F1: 4294967295, F2: -12}
                var v9 float32 = (3.14159 + float32((0.0 * (3.14159 * 1.5))))
                v10 := []int{int(p2), -1, fi7}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        case "":
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int8 = (^(p2))
                var v13 bool = false
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            switch (("foo bar" + "") + ("hello" + ("hello" + ""))) {
            case "foo bar":
                p1 = (p1 & (p1 ^ (p1 + (p1 >> ((42) & 63)))))
            case "hello":
                p1 = 0
            case "a":
                v14 := GenS1{F0: (p2 % (((p2 ^ 42)) | 1)), F1: 63, F2: -128}
                _ = v14
            }
        }
        v15 := GenS2{F0: p1, F1: "hello"}
        _ = fi6
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 bool = (false && (!("a" != ("\n\t" + ("hello" + "\n\t")))))
    var v2 uint64 = 376
    var v3 uint8 = 42
    var v4 uint8 = 58
    v5 := make(map[int]string)
    v1 = (65535 <= 951)
    if v1 {
        if false {
            var v16 int16 = -32768
            switch 32767 {
            case 4294967295:
                for fi17 := 0; fi17 < 4; fi17++ {
                    if ((v1 || (!(!(255 <= v3)))) && (0 == 1)) {
                        if ((65535 < 65535) && (27 < 127)) {
                            v2 = 1
                            delete(v5, fi17)
                            v18 := GenS3{F0: uint32(v2), F1: -1, F2: 18446744073709551615}
                            _ = v18
                        }
                        v4 = (v3 >> ((4294967295) & 63))
                    }
                    _ = fi17
                }
            case -2147483648:
                if false {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v4 -= (v3 - 128)
                        _ = fi19
                    }
                } else {
                    v16 = v16
                    v1 = ((v16 > (v16 + 1000)) && (-1 <= 0))
                    var v20 int16 = -55
                    _ = v20
                }
            }
            if (!((!true) && v1)) {
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS2{F0: -242, F1: "\n\t"}
                    fmt.Println(v3, v4, v2)
                    v3 = (v3 ^ 42)
                    _ = fi21
                    _ = v22
                }
                v3 -= v4
                v23 := []int{}
                _ = v23
            } else {
                v1 = false
            }
            _ = v16
        }
        v24 := GenS4{F0: (v2 * 726), F1: 127, F2: 790}
        v5[-2147483648] = ("a" + "\n\t")
        _ = v24
    } else {
        if (0 >= -9223372036854775808) {
            fmt.Println(v3, v3)
            var v25 float32 = 0.1
            fmt.Println(v4, v25)
            _ = v25
        }
        v2 = (v2 % ((0) | 1))
    }
    v26 := v5[9223372036854775807]
    if true {
        for fi27 := 0; fi27 < 15; fi27++ {
            v1 = (3.14159 != 0.0)
            switch fi27 {
            case 32767:
                if v1 {
                    if v1 {
                        fmt.Println(v3)
                        v2 = (^(v2))
                    } else {
                        v3 += (v3 / (((v3 & (v3 ^ (v4 << ((255) & 63))))) | 1))
                        v28 := []int{(^(fi27)), int(v2), fi27, int(v3)}
                        var v29 string = ("\n\t" + v26)
                        _ = v28
                        _ = v29
                    }
                    v30 := GenS2{F0: 42, F1: v26}
                    for fi31 := 0; fi31 < 17; fi31++ {
                        if v1 {
                            var v32 bool = true
                            if (float64(float32(3.14159)) != 0.0) {
                                v33 := GenS5{F0: 184, F1: -62}
                                if v1 {
                                    if v1 {
                                        var v34 int8 = 42
                                        _ = v34
                                    }
                                    fmt.Println(v32)
                                    v35 := GenS6{F0: (v26 + "\n\t")}
                                    _ = v35
                                } else {
                                    if (42 > 1) {
                                        var v36 float32 = (-2.25 * 0.0)
                                        v37 := v5[(fi27 >> ((255) & 63))]
                                        _ = v36
                                        _ = v37
                                    } else {
                                        v3 = uint8(v2)
                                    }
                                }
                                var v38 int64 = int64(v2)
                                switch genFn0(42, -894, 42) {
                                case -2147483648:
                                    v4 += (v3 - 17)
                                    v3 += v3
                                    fmt.Println(v26)
                                    var v39 bool = v32
                                    _ = v39
                                }
                                _ = v33
                                _ = v38
                            } else {
                                var v40 float32 = 0.0
                                fmt.Println(v32, v26)
                                _ = len(v5)
                                v32 = (!v32)
                                switch v26 {
                                case "":
                                    fmt.Println(fi27, fi27)
                                }
                                _ = v40
                            }
                            _ = v32
                        } else {
                            v41 := GenS7{F0: -1}
                            switch v26 {
                            case "":
                                v26 = "hello"
                            }
                            v26 = (((((((v26 + v26) + ("hello" + v26)) + v26) + v26) + (v26 + ("\n\t" + (v26 + ((v26 + (v26 + v26)) + "a"))))) + v26) + v26)
                            if v1 {
                                v42 := []int{(-(fi31))}
                                if v1 {
                                    if ((!true) || v1) {
                                        var v43 uint32 = 995
                                        var v44 int = 2147483647
                                        if false {
                                            v42 = append(v42, fi31)
                                            var v45 string = ("" + "")
                                            _ = v45
                                        }
                                        _ = v43
                                        _ = v44
                                    } else {
                                        switch (v26 + "hello") {
                                        case "a":
                                            v46 := GenS8{F0: 111, F1: genFn0(1, 1, 100)}
                                            _ = v46
                                        case "":
                                            var v47 float64 = 0.0
                                            _ = v47
                                        case "hello":
                                            v48 := GenS9{F0: -128}
                                            _ = v48
                                        }
                                        switch (fi31 & 32767) {
                                        case 329:
                                            fmt.Println(v3, v2, fi27)
                                        }
                                    }
                                    if len(v42) >= 2 {
                                        v42 = v42[1:len(v42)]
                                    }
                                    v49 := GenS9{F0: 100}
                                    v50 := GenS10{F0: ("foo bar" + ""), F1: v26, F2: 371}
                                    _ = v49
                                    _ = v50
                                } else {
                                    v51 := GenS8{F0: -932, F1: fi31}
                                    v52 := GenS3{F0: 0, F1: 42, F2: 638}
                                    _ = v51
                                    _ = v52
                                }
                                _ = v42
                            }
                            var v53 int8 = int8(v2)
                            v54 := GenS11{F0: 42, F1: -2147483648, F2: "foo bar"}
                            v55 := GenS3{F0: 0, F1: 1, F2: 26}
                            _ = v41
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        _ = fi31
                    }
                    _ = v30
                } else {
                    for fi56 := 0; fi56 < 17; fi56++ {
                        if true {
                            var v57 uint32 = 2147483647
                            v58 := GenS2{F0: -82, F1: (v26 + v26)}
                            v57 = v57
                            _ = v57
                            _ = v58
                        }
                        v5[fi27] = ""
                        var v59 int32 = 42
                        fmt.Println(v26, v26, fi27)
                        var v60 int = genFn0(164, 42, -1)
                        _ = fi56
                        _ = v59
                        _ = v60
                    }
                    v4 -= 127
                }
            case 42:
                switch fi27 {
                case 0:
                    for fi61 := 0; fi61 < 8; fi61++ {
                        if v1 {
                            if (0.0 == (3.14159 - 1.5)) {
                                if true {
                                    if v1 {
                                        var v62 uint64 = (v2 >> ((2) & 63))
                                        if (!(v1 || v1)) {
                                            var v63 uint = 2
                                            _ = v63
                                        }
                                        _ = v62
                                    }
                                    if v1 {
                                        var v64 uint32 = 42
                                        var v65 uint16 = 2
                                        _ = v64
                                        _ = v65
                                    }
                                    v3 += (^(v3))
                                } else {
                                    v4 = v3
                                    v66 := []int{(fi27 - fi61), genFn0(1000, 2147483647, 1)}
                                    fmt.Println(v2)
                                    switch "a" {
                                    case "hello":
                                        if ((0 == -827) || (88 <= 951)) {
                                            var v67 uint64 = v2
                                            fmt.Println(v26, fi61, v3)
                                            var v68 float32 = 1.5
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                    _ = v66
                                }
                                var v69 float64 = (float64(-2.25) - float64(0.1))
                                v70 := []int{(fi27 ^ genFn0(-32768, 818, 127))}
                                _ = v69
                                _ = v70
                            } else {
                                if v1 {
                                    v71 := []int{(fi27 | fi61), genFn0(-1, 2147483647, 100), genFn0(906, 0, -96), (-(fi61))}
                                    v72 := GenS3{F0: 0, F1: 0, F2: 681}
                                    _ = len(v71)
                                    _ = v71
                                    _ = v72
                                } else {
                                    if false {
                                        v73 := GenS12{F0: ((-2147483648 < 42) && (745 == -368)), F1: 1, F2: uint8(v2)}
                                        v1 = v1
                                        _ = v73
                                    } else {
                                        fmt.Println(v1, v2, fi61)
                                        if (true || true) {
                                            v26 = ("foo bar" + "a")
                                            if (!v1) {
                                                v3 += (^(v3))
                                                var v74 uint64 = (v2 - (v2 >> ((65535) & 63)))
                                                if false {
                                                    switch ("a" + "\n\t") {
                                                    case "a":
                                                        if ("a" >= v26) {
                                                            v75 := GenS11{F0: 42, F1: 419, F2: "a"}
                                                            _ = v75
                                                        } else {
                                                            v76 := v5[(fi61 * (fi27 - (fi61 & fi27)))]
                                                            v77 := GenS13{F0: v1, F1: 4294967295, F2: "a"}
                                                            if (v1 || (!(float64(float32((0.1 * (float64(1.5) * (3.14159 + 1e10))))) != 3.14159))) {
                                                                v3 = v4
                                                                v5[873] = "\n\t"
                                                            }
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        var v78 bool = (true || false)
                                                        var v79 uint64 = (v2 % ((822) | 1))
                                                        fmt.Println(v78, v4, fi61)
                                                        if true {
                                                            if false {
                                                                switch (fi27 & int(v79)) {
                                                                case 652:
                                                                    if (71 >= (v4 - v3)) {
                                                                        v80 := GenS14{F0: 768, F1: 32768, F2: v2}
                                                                        if (1000 <= 1) {
                                                                            _ = len(v5)
                                                                            var v81 string = v26
                                                                            if (9223372036854775807 >= 0) {
                                                                                if (v78 || (!v1)) {
                                                                                    delete(v5, genFn0(32767, 215, 127))
                                                                                }
                                                                            }
                                                                            v82 := GenS8{F0: 1000, F1: (^(fi27))}
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        var v83 uint8 = v3
                                                                        _ = v80
                                                                        _ = v83
                                                                    }
                                                                case 32767:
                                                                    v84 := GenS3{F0: 2147483648, F1: 2147483647, F2: 4294967295}
                                                                    _ = v84
                                                                }
                                                                if v78 {
                                                                    var v85 float64 = 3.14159
                                                                    if v78 {
                                                                        var v86 uint8 = v4
                                                                        delete(v5, (fi61 << ((42) & 63)))
                                                                        _ = v86
                                                                    }
                                                                    v79 = (v2 ^ 4294967295)
                                                                    if v78 {
                                                                        switch (v26 + v26) {
                                                                        case "":
                                                                            var v87 float32 = -2.25
                                                                            _ = v87
                                                                        case "a":
                                                                            if false {
                                                                                v88 := v5[(fi61 ^ fi27)]
                                                                                var v89 int8 = -73
                                                                                _ = v88
                                                                                _ = v89
                                                                            }
                                                                        case "hello":
                                                                            switch fi61 {
                                                                            case 373:
                                                                                var v90 uint = 0
                                                                                _ = v90
                                                                            case 1:
                                                                                v74 += 2
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            switch v26 {
                                                                            case "a":
                                                                                delete(v5, (fi27 | fi61))
                                                                            case "foo bar":
                                                                                fmt.Println(v4)
                                                                            case "":
                                                                                fmt.Println(fi61, v78, fi61)
                                                                            }
                                                                        }
                                                                        var v91 uint64 = 1
                                                                        v3 += (^(v3))
                                                                        _ = v91
                                                                    }
                                                                    if true {
                                                                        v92 := GenS11{F0: int64(v4), F1: -587, F2: "hello"}
                                                                        _ = v92
                                                                    }
                                                                    if v1 {
                                                                        v79 += uint64(v3)
                                                                        v93 := GenS15{F0: 42, F1: true, F2: 0}
                                                                        _ = v93
                                                                    } else {
                                                                        v94 := v5[fi61]
                                                                        if (v3 >= (^(v3))) {
                                                                            v79 += 2
                                                                            delete(v5, (-(fi27)))
                                                                            v95 := GenS2{F0: 2147483647, F1: (v26 + v94)}
                                                                            _ = v95
                                                                        } else {
                                                                            v96 := []int{(fi27 / ((-2147483648) | 1))}
                                                                            _ = v96
                                                                        }
                                                                        v97 := GenS13{F0: true, F1: 683, F2: ("hello" + v94)}
                                                                        _ = v94
                                                                        _ = v97
                                                                    }
                                                                    switch (fi61 >> ((255) & 63)) {
                                                                    case -9223372036854775808:
                                                                        var v98 float64 = v85
                                                                        _ = v98
                                                                    case 2147483647:
                                                                        v99 := GenS16{F0: (-(fi61)), F1: -307, F2: 290}
                                                                        _ = v99
                                                                    default:
                                                                        v85 -= (-2.25 / ((v85 / v85) + 1.5))
                                                                        var v100 int = fi61
                                                                        _ = v100
                                                                    }
                                                                    _ = v85
                                                                }
                                                                if v78 {
                                                                    var v101 uint64 = (v2 << ((4294967295) & 63))
                                                                    _ = v101
                                                                }
                                                            }
                                                            v78 = v78
                                                            var v102 string = v26
                                                            _ = v102
                                                        }
                                                        _ = v78
                                                        _ = v79
                                                    case "\n\t":
                                                        if true {
                                                            v26 = ("foo bar" + v26)
                                                            switch fi27 {
                                                            case -1:
                                                                var v103 int = (fi27 - (fi61 / ((fi61) | 1)))
                                                                _ = v103
                                                            case 127:
                                                                switch (-(fi27)) {
                                                                case 128:
                                                                    var v104 float64 = -2.25
                                                                    _ = v104
                                                                case 1:
                                                                    switch v26 {
                                                                    case "foo bar":
                                                                        var v105 uint16 = 255
                                                                        if v1 {
                                                                            v106 := GenS9{F0: -42}
                                                                            _ = v106
                                                                        }
                                                                        v107 := GenS17{F0: v105, F1: -128, F2: -128}
                                                                        fmt.Println(fi61, fi27, v26)
                                                                        _ = v105
                                                                        _ = v107
                                                                    case "hello":
                                                                        v26 = "foo bar"
                                                                    case "\n\t":
                                                                        var v108 float64 = 1.5
                                                                        _ = v108
                                                                    }
                                                                    var v109 int8 = 101
                                                                    v1 = false
                                                                    _ = v109
                                                                case 2147483647:
                                                                    v110 := GenS11{F0: 1, F1: -1, F2: (v26 + v26)}
                                                                    _ = v110
                                                                default:
                                                                    v111 := GenS11{F0: 42, F1: 2147483647, F2: ""}
                                                                    _ = v111
                                                                }
                                                            case -150:
                                                                v112 := []int{fi61}
                                                                _ = v112
                                                            }
                                                        } else {
                                                            fmt.Println(v4, v3, v74)
                                                        }
                                                    case "":
                                                        v26 = v26
                                                        v113 := []int{fi61, (-(fi61)), genFn0(-405, 108, 0)}
                                                        _ = v113
                                                    default:
                                                        fmt.Println(v3, v74, fi27)
                                                    }
                                                    if v1 {
                                                        v4 += (v4 | (v3 & v3))
                                                    }
                                                    fmt.Println(v1, v3)
                                                }
                                                _ = v74
                                            }
                                            var v114 int = fi61
                                            _ = v114
                                        } else {
                                            v4 -= v4
                                            if v1 {
                                                var v115 bool = v1
                                                var v116 int32 = 0
                                                if v115 {
                                                    if v1 {
                                                        fmt.Println(fi27, fi61)
                                                        fmt.Println(v26, v4, v3)
                                                    } else {
                                                        var v117 uint32 = 1
                                                        if (-2.25 <= 1e10) {
                                                            if ((42 < 0) || v115) {
                                                                _ = len(v5)
                                                                if true {
                                                                    var v118 int16 = 1
                                                                    v26 = "foo bar"
                                                                    _ = v118
                                                                } else {
                                                                    if true {
                                                                        var v119 int64 = 9223372036854775807
                                                                        _ = v119
                                                                    }
                                                                    var v120 uint32 = 2
                                                                    _ = v120
                                                                }
                                                                v115 = v1
                                                            } else {
                                                                if v115 {
                                                                    v121 := GenS18{F0: 0, F1: v116, F2: 9223372036854775807}
                                                                    _ = v121
                                                                } else {
                                                                    v122 := GenS13{F0: (!false), F1: v117, F2: ("" + (v26 + v26))}
                                                                    _ = v122
                                                                }
                                                                switch (fi61 << ((4294967295) & 63)) {
                                                                case 127:
                                                                    var v123 uint64 = 0
                                                                    v124 := GenS4{F0: 18446744073709551615, F1: 42, F2: -9223372036854775808}
                                                                    v125 := GenS12{F0: (695 == 0), F1: 4294967295, F2: (v4 + v4)}
                                                                    v2 += v2
                                                                    _ = v123
                                                                    _ = v124
                                                                    _ = v125
                                                                case -9223372036854775808:
                                                                    switch (-(fi61)) {
                                                                    case 0:
                                                                        delete(v5, fi61)
                                                                    case 128:
                                                                        switch fi61 {
                                                                        case -619:
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                if v115 {
                                                                                    v115 = (9223372036854775807 >= -879)
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        var v126 uint16 = uint16(v3)
                                                                                        _ = v126
                                                                                    case "foo bar":
                                                                                        fmt.Println(v117, v1)
                                                                                    default:
                                                                                        v26 = (v26 + ("a" + v26))
                                                                                    }
                                                                                }
                                                                            }
                                                                        case -1:
                                                                            fmt.Println(v116, v4, v26)
                                                                        default:
                                                                            var v127 bool = true
                                                                            _ = v127
                                                                        }
                                                                    default:
                                                                        if (v115 || true) {
                                                                            switch (fi61 << ((65535) & 63)) {
                                                                            case 128:
                                                                                delete(v5, fi27)
                                                                            default:
                                                                                fmt.Println(v3, v1)
                                                                            }
                                                                            switch (v26 + (v26 + "")) {
                                                                            case "\n\t":
                                                                                fmt.Println(v116)
                                                                            case "a":
                                                                                if (9223372036854775807 < -474) {
                                                                                    v128 := GenS3{F0: 0, F1: v116, F2: 796}
                                                                                    _ = v128
                                                                                }
                                                                            }
                                                                            if true {
                                                                                v115 = false
                                                                            } else {
                                                                                if false {
                                                                                    v115 = (true || v1)
                                                                                    var v129 uint8 = v3
                                                                                    fmt.Println(v129, v1, v117)
                                                                                    _ = v129
                                                                                } else {
                                                                                    v4 = v3
                                                                                    v130 := GenS19{F0: v117}
                                                                                    _ = v130
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                case 32767:
                                                                    v116 = v116
                                                                default:
                                                                    fmt.Println(v4)
                                                                }
                                                            }
                                                            switch (^(fi27)) {
                                                            case 4294967295:
                                                                var v131 int32 = (v116 & (v116 + (-(v116))))
                                                                _ = v131
                                                            case 32768:
                                                                fmt.Println(fi27, v115, v115)
                                                            case -2147483648:
                                                                var v132 uint8 = 127
                                                                _ = v132
                                                            }
                                                        }
                                                        _ = v117
                                                    }
                                                    fmt.Println(v2)
                                                    switch fi61 {
                                                    case -825:
                                                        v3 = 42
                                                    case -61:
                                                        v133 := []int{fi61, int(v2), (fi27 * fi61), -9223372036854775808}
                                                        _ = v133
                                                    case 2147483647:
                                                        v2 = 2
                                                    default:
                                                        v134 := []int{genFn0(0, -1, 1), (fi27 % ((fi27) | 1)), (^(fi27))}
                                                        _ = v134
                                                    }
                                                }
                                                _ = v115
                                                _ = v116
                                            } else {
                                                v135 := []int{(fi27 - (fi27 ^ 237)), 32768, fi61, (fi61 << ((65535) & 63))}
                                                _ = v135
                                            }
                                        }
                                        var v136 uint64 = 4294967295
                                        switch ("\n\t" + v26) {
                                        case "hello":
                                            fmt.Println(v26)
                                        }
                                        v137 := []int{(fi61 & (fi61 >> ((255) & 63))), 1, fi61, 32768}
                                        _ = v136
                                        _ = v137
                                    }
                                    switch ("foo bar" + ((v26 + v26) + v26)) {
                                    case "\n\t":
                                        switch fi61 {
                                        case 781:
                                            var v138 int = fi27
                                            _ = v138
                                        case 128:
                                            var v139 int64 = 139
                                            _ = v139
                                        default:
                                            v2 = 476
                                        }
                                        switch (fi27 / ((-9223372036854775808) | 1)) {
                                        case 38:
                                            fmt.Println(v3, v4)
                                        case 42:
                                            v140 := GenS1{F0: 100, F1: 255, F2: 123}
                                            _ = v140
                                        }
                                        var v141 int8 = 42
                                        _ = v141
                                    case "a":
                                        switch (-(fi61)) {
                                        case -1:
                                            v142 := v5[(fi61 - fi61)]
                                            _ = v142
                                        }
                                    default:
                                        var v143 string = v26
                                        _ = v143
                                    }
                                    if true {
                                        var v144 uint64 = v2
                                        _ = v144
                                    } else {
                                        v145 := GenS2{F0: -748, F1: "a"}
                                        v2 -= 18446744073709551615
                                        v146 := []int{fi27, fi27, (fi61 % ((genFn0(0, -489, 42)) | 1)), (fi27 << ((296) & 63))}
                                        v147 := GenS18{F0: -9223372036854775808, F1: -1, F2: -9223372036854775808}
                                        _ = v145
                                        _ = v146
                                        _ = v147
                                    }
                                }
                                v148 := GenS14{F0: (-(fi61)), F1: (fi61 * (-(fi61))), F2: 42}
                                _ = v148
                            }
                            if true {
                                if (v1 || v1) {
                                    var v149 uint32 = uint32(v2)
                                    var v150 float64 = 3.14159
                                    _ = v149
                                    _ = v150
                                } else {
                                    v151 := GenS20{F0: 0, F1: fi61, F2: 1}
                                    _ = v151
                                }
                            }
                        }
                        v26 = "a"
                        _ = fi61
                    }
                case 32767:
                    v3 += 0
                default:
                    for fi152 := 0; fi152 < 6; fi152++ {
                        v2 += (v2 / ((767) | 1))
                        delete(v5, (fi27 >> ((42) & 63)))
                        v26 = ("\n\t" + v26)
                        if false {
                            v153 := GenS3{F0: 4294967295, F1: -2147483648, F2: 42}
                            v26 = (v26 + v26)
                            _ = v153
                        } else {
                            if (2 == (^(v2))) {
                                delete(v5, (fi27 | (fi152 | (-(fi152)))))
                                switch genFn0(-52, -2147483648, 96) {
                                case 4294967295:
                                    var v154 int16 = -1
                                    _ = v154
                                default:
                                    if (v1 || (-32768 != 42)) {
                                        switch v26 {
                                        case "\n\t":
                                            fmt.Println(fi27, v1)
                                        case "":
                                            v155 := GenS21{F0: 1, F1: 613, F2: -9223372036854775808}
                                            _ = v155
                                        case "a":
                                            if v1 {
                                                v156 := []int{(-(fi152)), fi27, (^(fi27))}
                                                var v157 uint16 = 1
                                                v158 := GenS22{F0: 853}
                                                v26 = "\n\t"
                                                _ = v156
                                                _ = v157
                                                _ = v158
                                            } else {
                                                if (!true) {
                                                    var v159 int = fi152
                                                    _ = v159
                                                }
                                            }
                                            if v1 {
                                                var v160 int32 = 2147483647
                                                _ = len(v5)
                                                _ = v160
                                            } else {
                                                if (v1 || (65535 == 2)) {
                                                    v2 = v2
                                                } else {
                                                    v161 := []int{fi152, (fi27 + (fi27 % (((^(fi27))) | 1)))}
                                                    fmt.Println(v1, v26)
                                                    v26 = ("" + v26)
                                                    _ = v161
                                                }
                                                if (-9223372036854775808 >= 42) {
                                                    v162 := []int{fi152, (fi152 & fi27), 0, 128}
                                                    var v163 uint = 128
                                                    v4 = (v4 % ((2) | 1))
                                                    v3 = (v3 ^ 0)
                                                    if true {
                                                        if (uint8(v2) > (v4 | v4)) {
                                                            v5[(fi27 >> (((v163 % (((^(v163))) | 1))) & 63))] = v26
                                                        }
                                                    }
                                                    _ = v162
                                                    _ = v163
                                                }
                                                var v164 float32 = (((0.0 * (0.0 + 0.1)) * 0.0) / 1e10)
                                                switch ("\n\t" + (v26 + "foo bar")) {
                                                case "foo bar":
                                                    v164 = (1e10 * v164)
                                                case "a":
                                                    var v165 uint = 4294967295
                                                    _ = v165
                                                case "":
                                                    var v166 int8 = -128
                                                    _ = v166
                                                default:
                                                    delete(v5, (fi152 % ((32768) | 1)))
                                                }
                                                _ = v164
                                            }
                                        default:
                                            fmt.Println(v4, v1)
                                        }
                                    } else {
                                        var v167 uint64 = 4294967295
                                        v168 := GenS23{F0: 1, F1: 42, F2: (-32768 <= 1000)}
                                        v169 := []int{genFn0(-337, 1, 19), (-(fi27)), fi27, fi27}
                                        v3 += v3
                                        if len(v169) > 0 {
                                            v169[0] = 128
                                        }
                                        if false {
                                            v170 := GenS4{F0: v167, F1: -128, F2: -9223372036854775808}
                                            _ = v170
                                        }
                                        if (v1 || (32767 != 929)) {
                                            fmt.Println(v26, fi152)
                                        } else {
                                            if (!(true || (v1 || false))) {
                                                var v171 int16 = 72
                                                if len(v169) >= 2 {
                                                    v169 = v169[1:len(v169)]
                                                }
                                                v5[(fi152 / ((fi27) | 1))] = ((((v26 + ((v26 + "\n\t") + ("" + "foo bar"))) + (("foo bar" + v26) + ("foo bar" + v26))) + "") + v26)
                                                _ = v171
                                            } else {
                                                switch (fi27 << ((563) & 63)) {
                                                case 127:
                                                    var v172 int32 = int32(v4)
                                                    _ = v172
                                                case 1:
                                                    if (true && (false || (((v3 ^ (v4 ^ (v3 << ((65535) & 63)))) == v3) || v1))) {
                                                        v173 := GenS2{F0: 959, F1: v26}
                                                        if len(v169) > 0 {
                                                            v169[0] = (fi152 >> ((2) & 63))
                                                        }
                                                        if len(v169) >= 2 {
                                                            v169 = v169[1:len(v169)]
                                                        }
                                                        _ = v173
                                                    }
                                                case -605:
                                                    var v174 float32 = ((((-2.25 / 1e10) + ((1e10 + (1.5 / -2.25)) * (0.1 + 1.5))) * 0.1) / 1.5)
                                                    _ = v174
                                                }
                                            }
                                            v169 = append(v169, (fi27 + (-(fi27))))
                                            v2 = v167
                                        }
                                        _ = v167
                                        _ = v168
                                        _ = v169
                                    }
                                }
                                if v1 {
                                    v175 := v5[genFn0(-1, -461, 100)]
                                    var v176 int8 = -128
                                    _ = v175
                                    _ = v176
                                }
                            }
                            var v177 int32 = 2147483647
                            _ = v177
                        }
                        if v1 {
                            v1 = v1
                            if (false && (!v1)) {
                                if v1 {
                                    var v178 uint8 = 128
                                    _ = v178
                                } else {
                                    switch (fi152 >> ((18446744073709551615) & 63)) {
                                    case 2147483647:
                                        var v179 int64 = 1
                                        _ = v179
                                    case 4294967295:
                                        v180 := v5[genFn0(1000, -1, 1)]
                                        _ = v180
                                    }
                                    if true {
                                        v181 := GenS14{F0: 1, F1: (fi152 ^ (fi152 & genFn0(42, 844, -128))), F2: v2}
                                        v1 = v1
                                        switch (fi152 % ((genFn0(32767, 535, 59)) | 1)) {
                                        case 2147483647:
                                            var v182 float32 = 1e10
                                            _ = v182
                                        default:
                                            var v183 int8 = 60
                                            _ = v183
                                        }
                                        _ = v181
                                    }
                                    var v184 uint64 = v2
                                    _ = v184
                                }
                                var v185 int64 = -1
                                _ = v185
                            } else {
                                var v186 int = (fi152 + genFn0(32767, 1, -94))
                                if (v3 > 2) {
                                    switch fi27 {
                                    case 2147483647:
                                        v187 := GenS24{F0: genFn0(-768, 2147483647, 36)}
                                        _ = v187
                                    case -9223372036854775808:
                                        v4 += (^(v4))
                                        if v1 {
                                            v186 = fi152
                                        } else {
                                            v26 = ""
                                            var v188 int8 = -1
                                            _ = v188
                                        }
                                    default:
                                        var v189 int = (fi152 & fi27)
                                        _ = v189
                                    }
                                    v2 = (v2 << ((65535) & 63))
                                    if (fi152 >= (fi27 - (fi152 * (v186 & 32767)))) {
                                        fmt.Println(v4, v2, fi27)
                                        v186 = -1
                                        v190 := GenS20{F0: 32767, F1: (v186 - (-(fi152))), F2: 255}
                                        _ = v190
                                    }
                                } else {
                                    var v191 int32 = 451
                                    if (-1 < 110) {
                                        switch (v186 & -1) {
                                        case -9223372036854775808:
                                            switch 9223372036854775807 {
                                            case 2147483647:
                                                var v192 float64 = 1.5
                                                _ = v192
                                            }
                                        default:
                                            var v193 string = (((v26 + (v26 + v26)) + v26) + v26)
                                            _ = v193
                                        }
                                    } else {
                                        v194 := []int{genFn0(0, (v191 << ((988) & 63)), -33), genFn0(-32768, (v191 | 0), -1)}
                                        v4 += 0
                                        var v195 int8 = -1
                                        _ = v194
                                        _ = v195
                                    }
                                    var v196 bool = ((^(v2)) > v2)
                                    v186 += -2147483648
                                    _ = v191
                                    _ = v196
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    var v197 int64 = -9223372036854775808
                                    _ = v197
                                case "":
                                    switch (v186 + (fi152 >> ((42) & 63))) {
                                    case 127:
                                        v198 := v5[9223372036854775807]
                                        _ = v198
                                    case -369:
                                        fmt.Println(v1)
                                    }
                                case "a":
                                    v3 += v4
                                }
                                _ = v186
                            }
                        } else {
                            v2 = 0
                            v2 -= 2
                            fmt.Println(v4)
                            var v199 float32 = 0.0
                            v200 := GenS15{F0: 480, F1: v1, F2: (v2 & 755)}
                            if ((0 > 0) || v1) {
                                var v201 int64 = 0
                                var v202 string = ("" + v26)
                                v201 = -9223372036854775808
                                _ = v201
                                _ = v202
                            } else {
                                v4 = 2
                                if v1 {
                                    var v203 uint8 = 255
                                    var v204 bool = (1 <= 127)
                                    _ = v203
                                    _ = v204
                                }
                                v205 := GenS21{F0: 139, F1: -32768, F2: 9223372036854775807}
                                _ = v205
                            }
                            if v1 {
                                var v206 int32 = 1
                                var v207 int = 0
                                v208 := GenS2{F0: v206, F1: ""}
                                _ = v206
                                _ = v207
                                _ = v208
                            }
                            _ = v199
                            _ = v200
                        }
                        var v209 uint32 = 0
                        switch genFn0(-889, -265, 42) {
                        case -609:
                            v26 = (v26 + (((("" + (v26 + ((("foo bar" + "hello") + (v26 + "\n\t")) + (v26 + "\n\t")))) + v26) + v26) + v26))
                        case 1:
                            v210 := GenS5{F0: 91, F1: -128}
                            _ = v210
                        case 42:
                            v4 += v4
                        }
                        v5[-124] = "a"
                        _ = fi152
                        _ = v209
                    }
                }
            }
            v211 := GenS13{F0: v1, F1: 2147483647, F2: "\n\t"}
            _ = fi27
            _ = v211
        }
    } else {
        var v212 uint = 4294967295
        fmt.Println(v4)
        for fi213 := 0; fi213 < 1; fi213++ {
            for fi214 := 0; fi214 < 1; fi214++ {
                v215 := []int{}
                var v216 uint16 = 0
                _ = fi214
                _ = v215
                _ = v216
            }
            v212 += (v212 << ((v212) & 63))
            _ = fi213
        }
        _ = v212
    }
    if v1 {
        if v1 {
            var v217 float32 = 1.5
            v1 = false
            _ = v217
        }
        v3 = (^(v3))
    }
    v218 := v5[792]
    if v1 {
        var v219 uint16 = 870
        _ = v219
    } else {
        v220 := []int{}
        _ = v220
    }
    v1 = (true || false)
    var v221 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v218:", v218)
    fmt.Println("v221:", v221)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v218
    _ = v221
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
