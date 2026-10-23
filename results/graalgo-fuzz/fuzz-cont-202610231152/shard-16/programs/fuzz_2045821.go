package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int16; F2 int8 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 uint64; F1 string }
type GenS11 struct { F0 uint8; F1 bool; F2 bool }
type GenS12 struct { F0 int; F1 string }
type GenS13 struct { F0 uint8; F1 int; F2 uint }
type GenS14 struct { F0 string; F1 int32; F2 uint8 }
type GenS15 struct { F0 int32; F1 int64; F2 uint64 }
type GenS16 struct { F0 int16; F1 uint8 }
func genFn0() uint64 {
    var v6 uint8 = 127
    switch (("hello" + "a") + "\n\t") {
    case "":
        v7 := []int{128}
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 13; fi8++ {
            switch (fi8 % ((fi8) | 1)) {
            case 127:
                v9 := []int{498, fi8}
                _ = v9
            case 4294967295:
                for fi10 := 0; fi10 < 1; fi10++ {
                    if (-123 <= 0) {
                        var v11 int16 = 1000
                        var v12 uint8 = v6
                        var v13 int16 = (v11 - v11)
                        v14 := GenS1{F0: fi10}
                        switch "" {
                        case "\n\t":
                            v6 = v12
                        case "foo bar":
                            v15 := GenS2{F0: 9223372036854775807, F1: (fi8 + fi10)}
                            _ = v15
                        }
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                }
            }
            var v16 uint16 = 1
            _ = fi8
            _ = v16
        }
    case "a":
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := []int{fi17, 32768}
            v6 += v6
            _ = fi17
            _ = v18
        }
    }
    if (0 <= 634) {
        switch -10 {
        case -2:
            var v19 uint16 = 32767
            _ = v19
        default:
            if true {
                v6 -= (v6 - (v6 / ((v6) | 1)))
            }
        }
        var v20 bool = true
        var v21 uint64 = 0
        _ = v20
        _ = v21
    } else {
        v22 := GenS1{F0: -1}
        for fi23 := 0; fi23 < 18; fi23++ {
            v6 += v6
            var v24 bool = ((((42 < 1) || (!(127 > 20))) || true) && ((v6 > v6) || (42 >= -2147483648)))
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS1{F0: 32768}
        fmt.Println(v6, fi25)
        var v27 int64 = 42
        _ = fi25
        _ = v26
        _ = v27
    }
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 519
    v3 := make(map[string]bool)
    var v4 uint32 = 2147483648
    var v5 uint = 255
    for fi28 := 0; fi28 < 20; fi28++ {
        v3[""] = ((!false) || (421 < -9223372036854775808))
        v4 = (v4 % ((42) | 1))
        var v29 uint32 = v4
        _ = fi28
        _ = v29
    }
    var v30 uint8 = 158
    for fi31 := 0; fi31 < 11; fi31++ {
        if (!false) {
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 uint = v5
                _ = fi32
                _ = v33
            }
            var v34 int32 = 284
            v4 -= (v4 * v4)
            _ = v34
        } else {
            if (2 > v5) {
                for fi35 := 0; fi35 < 8; fi35++ {
                    var v36 float64 = -2.25
                    switch v2 {
                    case -546:
                        v30 = v30
                    case 839:
                        v36 = ((((((float64(0.0) * v36) * (float64((1e10 / -2.25)) * 0.1)) * v36) - (v36 / 1.0)) - ((v36 * ((0.1 / v36) - v36)) - v36)) + v36)
                    case -2147483648:
                        _ = len(v3)
                    default:
                        var v37 uint = uint(v30)
                        _ = v37
                    }
                    _ = fi35
                    _ = v36
                }
            }
        }
        var v38 uint64 = 42
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := GenS3{F0: -1, F1: -1, F2: "a"}
            v41 := []int{v2}
            if ((-9223372036854775808 != -457) && (!(42 > 1000))) {
                if (676 > fi39) {
                    if false {
                        if len(v41) >= 2 {
                            v41 = v41[1:len(v41)]
                        }
                        v30 = v30
                        var v42 int16 = 725
                        switch "foo bar" {
                        case "foo bar":
                            v42 = v42
                        }
                        _ = v42
                    } else {
                        v43 := GenS4{F0: -252}
                        if false {
                            v44 := GenS2{F0: (fi31 * -2147483648), F1: fi39}
                            var v45 float32 = 0.1
                            var v46 int = fi31
                            _ = v44
                            _ = v45
                            _ = v46
                        } else {
                            _ = len(v41)
                            v38 = v38
                        }
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v5, v4, v38)
                        default:
                            var v47 int8 = -128
                            _ = v47
                        }
                        _ = v43
                    }
                    v48 := GenS5{F0: (true && (!true)), F1: 277, F2: 1}
                    v49 := GenS3{F0: int16(v38), F1: 61, F2: "foo bar"}
                    _ = v48
                    _ = v49
                } else {
                    if (1000 != 1000) {
                        fmt.Println(v30, v2)
                        if false {
                            if (((1.5 * 1.5) * 1.5) < 1.5) {
                                v2 = (v2 | (fi31 % ((v2) | 1)))
                            }
                            _ = len(v1)
                        } else {
                            if false {
                                var v50 float64 = 3.14159
                                if false {
                                    var v51 uint16 = 65535
                                    v52 := GenS2{F0: 1, F1: 2147483647}
                                    _ = v51
                                    _ = v52
                                }
                                _ = v50
                            } else {
                                fmt.Println(fi31)
                                if ("foo bar" < "\n\t") {
                                    v2 = v2
                                    var v53 int32 = 42
                                    _ = v53
                                }
                                var v54 bool = (!(true && true))
                                if len(v41) > 0 {
                                    v41[0] = fi31
                                }
                                delete(v1, ("foo bar" + ("foo bar" + "\n\t")))
                                fmt.Println(v5, v38)
                                _ = v54
                            }
                        }
                    }
                    if true {
                        var v55 uint = v5
                        fmt.Println(v2, fi31)
                        _ = v55
                    } else {
                        if (!(1 > int64(v4))) {
                            var v56 uint16 = 32767
                            var v57 float32 = (1.5 - ((1.5 - (1.5 - float32((((1.5 / 1.5) - 1.5) + 3.14159)))) * 1e10))
                            _ = v56
                            _ = v57
                        }
                    }
                }
            } else {
                if true {
                    if (!(-182 > 298)) {
                        switch (v2 / (((fi39 / (((v2 / ((fi31) | 1))) | 1))) | 1)) {
                        case 9223372036854775807:
                            var v58 uint16 = 102
                            _ = v58
                        case -202:
                            if (!(1 <= 2)) {
                                var v59 float64 = 3.14159
                                _ = v59
                            }
                        }
                        v60 := GenS5{F0: (int64(v4) != 1), F1: 1000, F2: 0}
                        if ((^(v2)) <= int(v5)) {
                            v61 := GenS3{F0: 32767, F1: 0, F2: "\n\t"}
                            _ = v61
                        } else {
                            var v62 int8 = -90
                            var v63 int16 = -4
                            _ = v62
                            _ = v63
                        }
                        fmt.Println(v4)
                        _ = v60
                    }
                    var v64 int32 = -1
                    if false {
                        switch fi39 {
                        case 1:
                            var v65 uint32 = (v4 >> (((v5 / ((v5) | 1))) & 63))
                            _ = v65
                        case -662:
                            v5 -= v5
                        default:
                            var v66 int = (-(fi31))
                            _ = v66
                        }
                        if (true || (42 >= 0)) {
                            if (((-337 == -32768) && (false || (!true))) || (!true)) {
                                fmt.Println(fi39)
                                if true {
                                    var v67 int16 = 234
                                    switch ("\n\t" + "a") {
                                    case "\n\t":
                                        v68 := v1["a"]
                                        _ = v68
                                    case "hello":
                                        v69 := GenS1{F0: 4294967295}
                                        _ = v69
                                    case "a":
                                        if ((!(((v67 >> (((v5 + 949)) & 63)) < (^(v67))) || ((0.0 - ((1.5 + (1.5 - (-2.25 - 1e10))) + 1.5)) < 3.14159))) && ((^(v67)) < v67)) {
                                            var v70 uint16 = 255
                                            v71 := GenS4{F0: -823}
                                            var v72 int32 = -955
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v2)
                                        default:
                                            v73 := GenS6{F0: (^(v30)), F1: (!((-(v64)) > v64))}
                                            _ = v73
                                        }
                                    }
                                    _ = v67
                                }
                                fmt.Println(v30, v5)
                            }
                        }
                        v74 := GenS7{F0: v30, F1: v64}
                        _ = v74
                    } else {
                        var v75 int16 = 1
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            v76 := GenS8{F0: v5}
                            _ = v76
                        default:
                            v38 += uint64(v5)
                        }
                        switch ("hello" + "a") {
                        case "\n\t":
                            v75 += 42
                            v64 = v64
                            v5 = 4294967295
                        }
                        _ = v75
                    }
                    var v77 float32 = (1e10 - float32((((0.0 + -2.25) / float64(0.1)) / 1.0)))
                    _ = v64
                    _ = v77
                }
                var v78 bool = true
                v79 := GenS8{F0: (^(v5))}
                _ = v78
                _ = v79
            }
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi31
        _ = v38
    }
    switch v2 {
    case 32768:
        var v80 float32 = 0.0
        _ = v80
    case 127:
        v81 := []int{32768, v2, v2}
        _ = v81
    }
    for fi82 := 0; fi82 < 18; fi82++ {
        switch (v2 >> ((4294967295) & 63)) {
        case 127:
            delete(v3, "")
        case 128:
            v4 -= v4
        case 488:
            for fi83 := 0; fi83 < 13; fi83++ {
                if true {
                    v84 := GenS9{F0: false, F1: 2147483647}
                    v85 := GenS10{F0: 4294967295, F1: ("\n\t" + "foo bar")}
                    _ = v84
                    _ = v85
                } else {
                    var v86 uint16 = 255
                    _ = v86
                }
                _ = fi83
            }
        default:
            if ("a" >= ("\n\t" + ("" + "hello"))) {
                switch "" {
                case "\n\t":
                    v87 := []int{(v2 << ((v5) & 63)), 9223372036854775807, (v2 + fi82), (fi82 << (((v5 & (v5 ^ (v5 * v5)))) & 63))}
                    _ = v87
                case "foo bar":
                    v2 += v2
                case "a":
                    v88 := GenS9{F0: (0.1 > 0.0), F1: -980}
                    _ = v88
                default:
                    for fi89 := 0; fi89 < 17; fi89++ {
                        fmt.Println(fi82, fi89, v2)
                        if false {
                            v90 := GenS11{F0: v30, F1: (!(-543 <= -102)), F2: (!(((!(!(false && false))) && (579 > 4294967295)) || false))}
                            var v91 uint16 = 499
                            v92 := v3["a"]
                            _ = v90
                            _ = v91
                            _ = v92
                        } else {
                            if false {
                                fmt.Println(v4)
                            }
                            switch "a" {
                            case "\n\t":
                                v93 := []int{fi89, (v2 ^ fi89)}
                                _ = v93
                            case "a":
                                if false {
                                    switch ("hello" + "") {
                                    case "a":
                                        v94 := GenS12{F0: (v2 * (^(fi82))), F1: "foo bar"}
                                        _ = v94
                                    case "foo bar":
                                        switch "" {
                                        case "foo bar":
                                            v30 = 1
                                        case "a":
                                            var v95 uint = v5
                                            _ = v95
                                        case "\n\t":
                                            var v96 int32 = 0
                                            _ = v96
                                        default:
                                            if (32767 >= 1000) {
                                                v97 := GenS13{F0: (^(v30)), F1: v2, F2: v5}
                                                _ = v97
                                            } else {
                                                switch ("hello" + ("a" + ("\n\t" + "a"))) {
                                                case "\n\t":
                                                    v98 := []int{(fi89 << ((65535) & 63)), (fi89 + fi82), fi82, fi82}
                                                    _ = v98
                                                default:
                                                    var v99 bool = (v30 >= 255)
                                                    _ = v99
                                                }
                                                var v100 bool = true
                                                v101 := GenS14{F0: "a", F1: 2147483647, F2: v30}
                                                _ = v100
                                                _ = v101
                                            }
                                        }
                                    }
                                    if false {
                                        switch (fi82 / ((v2) | 1)) {
                                        case 32768:
                                            var v102 uint64 = 416
                                            _ = v102
                                        }
                                        fmt.Println(v2, fi82)
                                        var v103 uint64 = 2
                                        _ = v103
                                    } else {
                                        switch 32767 {
                                        case 32768:
                                            switch "hello" {
                                            case "foo bar":
                                                var v104 float32 = 0.1
                                                _ = v104
                                            default:
                                                var v105 uint16 = 2
                                                _ = v105
                                            }
                                        case 803:
                                            v5 -= v5
                                        case -1:
                                            v30 = v30
                                        default:
                                            v106 := GenS15{F0: 799, F1: 0, F2: 4294967295}
                                            _ = v106
                                        }
                                    }
                                    fmt.Println(v5, fi82)
                                } else {
                                    fmt.Println(v2, v2)
                                    var v107 int64 = -1
                                    v30 -= v30
                                    switch "hello" {
                                    case "a":
                                        if false {
                                            v108 := GenS15{F0: 1, F1: 42, F2: 2}
                                            var v109 uint8 = v30
                                            _ = v108
                                            _ = v109
                                        }
                                    case "hello":
                                        v2 = fi82
                                    case "\n\t":
                                        var v110 uint16 = 28
                                        _ = v110
                                    }
                                    _ = v107
                                }
                            }
                            _ = len(v3)
                            v5 += (v5 & 737)
                            v4 = v4
                            v2 = fi82
                            v111 := GenS16{F0: 241, F1: (v30 << ((v5) & 63))}
                            var v112 float32 = (float32(float64((3.14159 - ((1.5 - 1e10) * (-2.25 * ((0.0 - 1.5) * ((3.14159 * 1e10) * 1.5))))))) / 3.14159)
                            v113 := GenS13{F0: v30, F1: (fi82 ^ (fi82 & fi89)), F2: 1}
                            _ = v111
                            _ = v112
                            _ = v113
                        }
                        _ = fi89
                    }
                }
                var v114 uint64 = 545
                for fi115 := 0; fi115 < 14; fi115++ {
                    v116 := GenS1{F0: (fi115 / ((0) | 1))}
                    v117 := []int{}
                    _ = fi115
                    _ = v116
                    _ = v117
                }
                _ = v114
            }
        }
        _ = fi82
    }
    for fi118 := 0; fi118 < 20; fi118++ {
        v30 += (v30 >> ((v5) & 63))
        var v119 int64 = -9223372036854775808
        v5 = v5
        _ = fi118
        _ = v119
    }
    v30 -= 85
    v120 := []int{(v2 / ((-1) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v120:", len(v120), v120)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v120
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
