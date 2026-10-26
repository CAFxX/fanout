package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 string; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 uint16; F2 int8 }
type GenS7 struct { F0 string; F1 int64; F2 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 uint; F1 uint64; F2 uint }
type GenS13 struct { F0 uint; F1 uint32 }
func genFn0(p0 bool, p1 int32) uint32 {
    var v6 int8 = -57
    if (("\n\t" == "a") || (1e10 != 0.1)) {
        var v7 uint = 255
        switch 0 {
        case -9223372036854775808:
            if p0 {
                v6 += v6
                for fi8 := 0; fi8 < 14; fi8++ {
                    p1 -= (p1 << ((v7) & 63))
                    _ = fi8
                }
                v9 := []int{}
                _ = v9
            } else {
                v10 := []int{-1}
                p0 = p0
                _ = v10
            }
        case 32767:
            var v11 uint = 65535
            _ = v11
        case 1:
            fmt.Println(v7, p0)
        }
        switch "\n\t" {
        case "\n\t":
            switch ("foo bar" + "") {
            case "\n\t":
                v7 = (v7 >> (((^(v7))) & 63))
            case "hello":
                for fi12 := 0; fi12 < 8; fi12++ {
                    v13 := GenS1{F0: 42, F1: 42, F2: 0}
                    v14 := GenS1{F0: 1000, F1: 0, F2: 396}
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v16 := []int{(fi12 * fi15), fi15}
                        switch (fi15 ^ (fi12 << ((v7) & 63))) {
                        case 128:
                            p0 = false
                        }
                        if len(v16) > 0 {
                            v16[0] = (-(fi12))
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case "foo bar":
                fmt.Println(p0, p0, v6)
            }
        default:
            v17 := GenS2{F0: ("foo bar" + "foo bar")}
            _ = v17
        }
        _ = v7
    } else {
        if true {
            v6 = v6
            var v18 uint = 1
            v19 := GenS3{F0: (!true), F1: "a", F2: 127}
            _ = v18
            _ = v19
        }
        if p0 {
            fmt.Println(p1)
            var v20 float32 = (float32(3.14159) - (1e10 / 1.0))
            _ = v20
        } else {
            var v21 uint64 = 1
            v21 += (v21 | (v21 ^ 0))
            _ = v21
        }
        var v22 float64 = 0.1
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v6
    return 45
}
func genFn1() uint8 {
    for fi23 := 0; fi23 < 2; fi23++ {
        switch "a" {
        case "foo bar":
            v24 := GenS4{F0: 42, F1: int32(fi23), F2: 2147483648}
            _ = v24
        case "a":
            if (!true) {
                v25 := []int{(-(fi23))}
                _ = v25
            } else {
                fmt.Println(fi23)
                v26 := []int{(^(fi23)), 0}
                var v27 uint32 = genFn0((982 == 513), -61)
                _ = v26
                _ = v27
            }
        case "hello":
            var v28 int8 = 127
            _ = v28
        default:
            v29 := []int{42, fi23}
            _ = v29
        }
        fmt.Println(fi23)
        var v30 int = (fi23 ^ fi23)
        _ = fi23
        _ = v30
    }
    switch 127 {
    case 32767:
        var v31 int8 = 127
        _ = v31
    default:
        var v32 uint = 622
        _ = v32
    }
    v33 := []int{2147483647, 2147483647}
    _ = v33
    return 127
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 2
    var v3 int64 = -488
    var v4 int64 = (v3 << ((65535) & 63))
    var v5 bool = true
    if v5 {
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := GenS5{F0: ("\n\t" + "\n\t")}
            var v36 bool = (false || v5)
            var v37 string = (("hello" + "hello") + "")
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        if (3.14159 == ((1.5 + 0.1) * (0.0 * 1e10))) {
            v5 = (!v5)
            var v38 int8 = 1
            v39 := []int{127, 4294967295, 9223372036854775807}
            _ = v38
            _ = v39
        } else {
            v40 := v1["a"]
            _ = v40
        }
        if v5 {
            if v5 {
                var v41 int64 = -1
                for fi42 := 0; fi42 < 5; fi42++ {
                    if (v5 && ((genFn1() != 2) && (v2 > (v2 | 32767)))) {
                        var v43 int16 = 904
                        _ = v43
                    } else {
                        v41 -= v4
                    }
                    _ = fi42
                }
                v44 := GenS6{F0: v3, F1: v2, F2: 127}
                _ = v41
                _ = v44
            } else {
                if v5 {
                    for fi45 := 0; fi45 < 9; fi45++ {
                        v4 -= v4
                        _ = fi45
                    }
                    fmt.Println(v4, v5)
                    switch "a" {
                    case "\n\t":
                        for fi46 := 0; fi46 < 15; fi46++ {
                            v5 = ((!(0.1 >= 0.0)) || true)
                            v1["\n\t"] = (-924 != (v4 + (v4 | v3)))
                            var v47 float64 = 3.14159
                            _ = fi46
                            _ = v47
                        }
                    case "foo bar":
                        for fi48 := 0; fi48 < 9; fi48++ {
                            if v5 {
                                if v5 {
                                    switch fi48 {
                                    case 1:
                                        v1[""] = ((v5 || (0 < 278)) && false)
                                    case 4294967295:
                                        v1[""] = (false || v5)
                                    case -2147483648:
                                        fmt.Println(v3, fi48)
                                    default:
                                        if (1.5 > float32(3.14159)) {
                                            var v49 uint8 = 1
                                            v50 := []int{42, fi48, (fi48 | 2147483647)}
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    var v51 uint8 = 42
                                    _ = v51
                                } else {
                                    v5 = true
                                    var v52 int32 = 0
                                    _ = v52
                                }
                                v4 -= v3
                            } else {
                                var v53 bool = false
                                _ = v53
                            }
                            _ = fi48
                        }
                    case "":
                        var v54 string = ("hello" + "foo bar")
                        _ = v54
                    default:
                        switch "a" {
                        case "foo bar":
                            v4 = (v4 - v3)
                        default:
                            v55 := GenS2{F0: "hello"}
                            _ = v55
                        }
                    }
                }
            }
        }
    }
    v56 := []int{1, -1}
    if (v5 && ((1 > 58) && ("hello" == ("\n\t" + (("\n\t" + "hello") + ""))))) {
        for fi57 := 0; fi57 < 3; fi57++ {
            switch "hello" {
            case "":
                for fi58 := 0; fi58 < 10; fi58++ {
                    v1[("foo bar" + "a")] = ((2 > 120) || false)
                    if (!(42 >= 100)) {
                        if v5 {
                            _ = len(v1)
                            v59 := GenS1{F0: -929, F1: 42, F2: 2}
                            v4 = v4
                            _ = v59
                        }
                        v5 = v5
                    } else {
                        v4 = (v4 & v3)
                    }
                    if v5 {
                        v60 := GenS7{F0: ((("foo bar" + ("hello" + "foo bar")) + "hello") + ""), F1: (v4 ^ 42), F2: fi58}
                        switch (fi58 - (fi58 - fi58)) {
                        case -296:
                            if v5 {
                                v61 := GenS3{F0: (1 != 15), F1: ("hello" + "\n\t"), F2: genFn1()}
                                switch -36 {
                                case 127:
                                    v3 = (v4 * (v3 << ((2) & 63)))
                                case 4294967295:
                                    if v5 {
                                        v62 := GenS3{F0: (3.14159 <= -2.25), F1: (("\n\t" + "") + "hello"), F2: 127}
                                        v63 := GenS4{F0: 255, F1: 0, F2: 42}
                                        v4 -= v3
                                        _ = v62
                                        _ = v63
                                    } else {
                                        v3 = (-(v4))
                                        _ = len(v56)
                                    }
                                default:
                                    if (!(42 > 0)) {
                                        var v64 uint8 = 242
                                        switch (fi57 & 4294967295) {
                                        case 97:
                                            var v65 uint16 = 255
                                            _ = v65
                                        case -9223372036854775808:
                                            var v66 float32 = ((1.5 + 0.1) * float32(3.14159))
                                            _ = v66
                                        case 42:
                                            v4 = v3
                                        }
                                        _ = v64
                                    } else {
                                        var v67 int32 = -1
                                        v5 = v5
                                        _ = v67
                                    }
                                }
                                _ = v61
                            } else {
                                if false {
                                    var v68 int32 = 0
                                    _ = v68
                                } else {
                                    v3 += (v3 - (v4 + v4))
                                    v2 -= (v2 >> ((2) & 63))
                                }
                                v4 = v4
                                fmt.Println(v5, v5)
                                if v5 {
                                    v3 = v3
                                    var v69 uint8 = 28
                                    if v5 {
                                        if true {
                                            v4 -= v4
                                            v70 := GenS1{F0: -387, F1: v3, F2: 374}
                                            v71 := GenS8{F0: uint(v69)}
                                            if v5 {
                                                switch fi58 {
                                                case -922:
                                                    switch -2147483648 {
                                                    case 1:
                                                        switch "hello" {
                                                        case "hello":
                                                            v72 := GenS9{F0: 884}
                                                            _ = v72
                                                        case "foo bar":
                                                            var v73 uint = 0
                                                            _ = v73
                                                        }
                                                    case -522:
                                                        v2 = (v2 - 105)
                                                    }
                                                case 42:
                                                    var v74 uint16 = (v2 | 65535)
                                                    _ = v74
                                                }
                                                if (0 < (v3 | v4)) {
                                                    v2 -= v2
                                                } else {
                                                    var v75 int = (fi57 ^ fi57)
                                                    v76 := GenS10{F0: 42}
                                                    _ = v75
                                                    _ = v76
                                                }
                                            }
                                            _ = v70
                                            _ = v71
                                        } else {
                                            if v5 {
                                                v5 = true
                                                v2 -= (v2 + (v2 & (v2 & 2)))
                                                v77 := GenS7{F0: ((("a" + "hello") + (("hello" + "") + "a")) + ""), F1: v4, F2: fi58}
                                                if ((!(42 < 4294967295)) || (!false)) {
                                                    v3 -= 9223372036854775807
                                                    _ = len(v56)
                                                    var v78 float64 = float64(3.14159)
                                                    fmt.Println(v3, v5)
                                                    var v79 uint = 961
                                                    _ = v78
                                                    _ = v79
                                                } else {
                                                    if v5 {
                                                        if (0 > 781) {
                                                            delete(v1, "a")
                                                            var v80 float64 = 3.14159
                                                            var v81 uint32 = 42
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            v82 := GenS5{F0: "a"}
                                                            v83 := v1["foo bar"]
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        v4 += (v4 & v4)
                                                    } else {
                                                        if v5 {
                                                            if (!(true && (4294967295 > 941))) {
                                                                if len(v56) >= 2 {
                                                                    v56 = v56[1:len(v56)]
                                                                }
                                                            }
                                                            var v84 int = fi57
                                                            _ = v84
                                                        }
                                                        v85 := GenS4{F0: (v2 % (((v2 | (v2 + v2))) | 1)), F1: -2147483648, F2: 1}
                                                        _ = len(v56)
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v86 := GenS11{F0: 1, F1: 42}
                                                            _ = v86
                                                        case "\n\t":
                                                            var v87 uint8 = (v69 >> ((2) & 63))
                                                            _ = v87
                                                        case "":
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                switch (fi57 - -602) {
                                                                case 32768:
                                                                    var v88 float64 = 0.0
                                                                    _ = v88
                                                                case 109:
                                                                    v5 = v5
                                                                default:
                                                                    v2 = (v2 / ((v2) | 1))
                                                                }
                                                            case "hello":
                                                                v4 = 42
                                                            }
                                                        }
                                                        v3 = (v3 & (v4 ^ -559))
                                                        _ = v85
                                                    }
                                                    v3 -= 210
                                                    var v89 int8 = 1
                                                    _ = v89
                                                }
                                                if (("" <= "hello") && (19 <= 122)) {
                                                    var v90 int64 = (v3 % ((1) | 1))
                                                    var v91 uint8 = v69
                                                    v4 = v90
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    switch 32768 {
                                                    case 32767:
                                                        var v92 int64 = (-(v3))
                                                        _ = v92
                                                    case 128:
                                                        v5 = true
                                                    case 216:
                                                        var v93 int16 = 1000
                                                        _ = v93
                                                    default:
                                                        if ((v5 || false) || (v5 && true)) {
                                                            fmt.Println(fi57, fi57)
                                                            var v94 int = 0
                                                            v2 += (v2 + 42)
                                                            _ = v94
                                                        }
                                                        var v95 uint8 = (^(v69))
                                                        _ = v95
                                                    }
                                                    var v96 uint16 = v2
                                                    if len(v56) >= 2 {
                                                        v56 = v56[1:len(v56)]
                                                    }
                                                    _ = v96
                                                }
                                                _ = v77
                                            } else {
                                                var v97 uint64 = 2
                                                v4 -= 42
                                                _ = v97
                                            }
                                        }
                                        var v98 float64 = (-2.25 + -2.25)
                                        _ = v98
                                    } else {
                                        v69 = v69
                                        switch (fi58 ^ fi58) {
                                        case 4294967295:
                                            v99 := GenS2{F0: (("hello" + "") + "a")}
                                            _ = v99
                                        }
                                        var v100 float32 = -2.25
                                        v100 = v100
                                        _ = v100
                                    }
                                    _ = v69
                                }
                            }
                        case 32767:
                            v5 = true
                        case 4294967295:
                            v2 = (v2 % ((v2) | 1))
                            v101 := GenS6{F0: (v4 - (v4 / (((v4 + v3)) | 1))), F1: v2, F2: 100}
                            _ = v101
                        }
                        fmt.Println(fi58)
                        delete(v1, "hello")
                        _ = v60
                    }
                    v4 = (^(v4))
                    delete(v1, ("\n\t" + ("" + ("hello" + ("\n\t" + "")))))
                    fmt.Println(fi57, v5, v2)
                    var v102 uint8 = 128
                    v4 = -9223372036854775808
                    var v103 uint32 = 0
                    _ = fi58
                    _ = v102
                    _ = v103
                }
            case "foo bar":
                switch (fi57 & fi57) {
                case 9223372036854775807:
                    v104 := v1[("hello" + "hello")]
                    _ = v104
                case 0:
                    v105 := GenS12{F0: 0, F1: 1, F2: 324}
                    _ = v105
                }
            }
            _ = fi57
        }
        if (18446744073709551615 >= 1) {
            if len(v56) >= 2 {
                v56 = v56[1:len(v56)]
            }
            var v106 uint32 = 183
            _ = v106
        }
        v107 := GenS3{F0: (v5 || false), F1: "foo bar", F2: 42}
        _ = v107
    } else {
        _ = len(v1)
        var v108 int64 = (v4 ^ (^(v4)))
        _ = v108
    }
    v5 = (42 < (^(v4)))
    if v5 {
        fmt.Println(v4)
        switch 42 {
        case 127:
            var v109 uint32 = 2147483647
            _ = v109
        }
        var v110 int8 = -68
        _ = v110
    }
    var v111 bool = (765 != 1)
    for fi112 := 0; fi112 < 2; fi112++ {
        v113 := GenS6{F0: 42, F1: (v2 * 65535), F2: 127}
        var v114 uint64 = 0
        switch -9223372036854775808 {
        case 0:
            v115 := GenS5{F0: "hello"}
            _ = v115
        case -895:
            fmt.Println(v111)
        case -9223372036854775808:
            v111 = ((102 <= -1) || (("foo bar" + "a") < "foo bar"))
        }
        _ = fi112
        _ = v113
        _ = v114
    }
    v4 -= v3
    var v116 int16 = 1
    v4 += v3
    for fi117 := 0; fi117 < 17; fi117++ {
        delete(v1, "foo bar")
        v2 += v2
        _ = fi117
    }
    var v118 float32 = 1e10
    _ = len(v1)
    v4 -= 1
    v119 := GenS13{F0: 1, F1: 2}
    if (!(724 <= 0)) {
        fmt.Println(v2)
        var v120 int = 9223372036854775807
        var v121 int8 = 100
        _ = v120
        _ = v121
    } else {
        v111 = v111
    }
    fmt.Println(v118)
    if false {
        if (1 == uint64(v2)) {
            var v122 int16 = v116
            v123 := GenS7{F0: "foo bar", F1: (v3 % (((v3 & (v3 << ((627) & 63)))) | 1)), F2: -689}
            _ = v122
            _ = v123
        } else {
            var v124 bool = (!v5)
            _ = v124
        }
        v116 -= v116
        fmt.Println(v116)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v111:", v111)
    fmt.Println("v116:", v116)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v111
    _ = v116
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
