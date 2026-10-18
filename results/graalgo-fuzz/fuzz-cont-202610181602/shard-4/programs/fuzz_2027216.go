package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 bool; F1 int32; F2 uint }
type GenS6 struct { F0 uint; F1 string; F2 int }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 string; F1 uint32; F2 uint32 }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 int8; F2 int }
type GenS12 struct { F0 uint64; F1 int16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint }
func genFn0(p0 int16) int8 {
    if true {
        var v7 int16 = (p0 - 1000)
        switch -990 {
        case 2147483647:
            v8 := []int{128, -1, 128, -822}
            _ = v8
        case 1:
            v9 := GenS1{F0: 821}
            _ = v9
        default:
            v10 := []int{722, -9223372036854775808, 9223372036854775807, 4294967295}
            _ = v10
        }
        _ = v7
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        p0 = (-(p0))
        if ((0.0 * 0.1) < (-2.25 + ((float64(float32(3.14159)) + -2.25) / 1.0))) {
            if (-1 == 0) {
                v12 := GenS1{F0: 42}
                v13 := []int{fi11, fi11}
                if false {
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v15 := GenS1{F0: 1}
                        v16 := GenS1{F0: 339}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if (1 < 436) {
                        _ = len(v13)
                        p0 -= 1
                        var v17 float32 = 1e10
                        _ = v17
                    } else {
                        if false {
                            v13 = append(v13, (^(fi11)))
                        }
                        var v18 uint8 = 127
                        v19 := GenS2{F0: -100, F1: fi11}
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS1{F0: 65535}
                    _ = v20
                } else {
                    p0 = p0
                }
                _ = v12
                _ = v13
            }
        } else {
            fmt.Println(fi11)
            v21 := GenS3{F0: 4294967295, F1: false}
            for fi22 := 0; fi22 < 19; fi22++ {
                if true {
                    p0 -= -32768
                    v23 := GenS1{F0: 18446744073709551615}
                    v24 := []int{fi22, fi22, (fi11 & fi11), (fi11 >> ((65535) & 63))}
                    _ = v23
                    _ = v24
                }
                _ = fi22
            }
            _ = v21
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int64 = -1
            _ = fi25
            _ = v26
        }
        _ = fi11
    }
    v27 := GenS4{F0: "foo bar", F1: 72}
    v28 := []int{-9223372036854775808}
    v29 := GenS1{F0: 0}
    _ = p0
    _ = v27
    _ = v28
    _ = v29
    return 1
}
func genFn1(p0 uint16, p1 bool, p2 uint64) int32 {
    if p1 {
        switch "a" {
        case "hello":
            fmt.Println(p0, p2, p0)
        case "foo bar":
            p0 = 0
        case "\n\t":
            if (!((-2.25 * (0.1 + 0.1)) >= -2.25)) {
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 int8 = -1
                    _ = fi30
                    _ = v31
                }
                fmt.Println(p2, p2, p0)
            }
        }
    }
    switch "" {
    case "a":
        if p1 {
            v32 := []int{0, 1, -9223372036854775808}
            _ = v32
        }
    case "\n\t":
        p2 += (p2 | p2)
    case "hello":
        var v33 int16 = 1
        _ = v33
    }
    v34 := GenS5{F0: (!p1), F1: -2147483648, F2: 323}
    var v35 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v35
    return 0
}

func main() {
    v1 := []string{"foo bar", "foo bar", "a", "hello"}
    var v2 int64 = -413
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    var v5 int = 783
    var v6 uint64 = 18446744073709551615
    v6 -= (v6 * (v6 << ((2) & 63)))
    for fi36 := 0; fi36 < 7; fi36++ {
        fmt.Println(v6, v6)
        _ = fi36
    }
    v5 -= 4294967295
    v37 := GenS6{F0: 65535, F1: "", F2: (v5 % ((v5) | 1))}
    v2 = (v2 % (((v2 & (v2 % (((v2 / ((207) | 1))) | 1)))) | 1))
    fmt.Println(v2)
    v38 := GenS3{F0: 0, F1: (!(952 == 32767))}
    switch v5 {
    case 356:
        for fi39 := 0; fi39 < 13; fi39++ {
            var v40 int8 = int8(v6)
            v41 := GenS7{F0: genFn0(32767), F1: 170}
            _ = fi39
            _ = v40
            _ = v41
        }
    case -1:
        if true {
            fmt.Println(v5, v6)
            if ((65535 >= 4294967295) && ((v2 % ((1) | 1)) >= (v2 * (v2 % ((0) | 1))))) {
                v3[(v5 % ((-9223372036854775808) | 1))] = (v5 / ((v5) | 1))
            } else {
                switch 32767 {
                case 9223372036854775807:
                    var v42 uint8 = 127
                    _ = v42
                case -2147483648:
                    switch "hello" {
                    case "a":
                        if true {
                            v43 := GenS8{F0: (("" + "a") + "\n\t"), F1: 4294967295, F2: 2}
                            var v44 int8 = -128
                            _ = v43
                            _ = v44
                        }
                    case "foo bar":
                        v6 = v6
                    }
                case 873:
                    switch (v5 / ((v5) | 1)) {
                    case 32768:
                        for fi45 := 0; fi45 < 19; fi45++ {
                            switch fi45 {
                            case 2147483647:
                                v2 -= v2
                            case 127:
                                var v46 int16 = 0
                                _ = v46
                            case -2147483648:
                                switch v5 {
                                case 42:
                                    var v47 uint64 = v6
                                    _ = v47
                                case 4294967295:
                                    v2 = v2
                                }
                            default:
                                for fi48 := 0; fi48 < 20; fi48++ {
                                    v49 := GenS5{F0: false, F1: -247, F2: 685}
                                    if false {
                                        _ = len(v1)
                                    }
                                    v2 += (v2 % (((v2 * v2)) | 1))
                                    _ = fi48
                                    _ = v49
                                }
                            }
                            v6 = v6
                            _ = len(v3)
                            _ = fi45
                        }
                    case -893:
                        v2 -= (v2 << ((2) & 63))
                    case -802:
                        v6 += (v6 + v6)
                    default:
                        if (true || (true && (true || (true && true)))) {
                            v5 += 554
                        }
                    }
                default:
                    if (0.1 > -2.25) {
                        var v50 bool = false
                        if v50 {
                            for fi51 := 0; fi51 < 16; fi51++ {
                                for fi52 := 0; fi52 < 14; fi52++ {
                                    var v53 int = (fi51 & 9223372036854775807)
                                    var v54 int64 = v2
                                    var v55 int32 = -2147483648
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                var v56 bool = true
                                _ = fi51
                                _ = v56
                            }
                        }
                        var v57 int16 = -32768
                        _ = v50
                        _ = v57
                    } else {
                        switch ("" + (("hello" + (("\n\t" + "foo bar") + "foo bar")) + "")) {
                        case "foo bar":
                            for fi58 := 0; fi58 < 15; fi58++ {
                                var v59 int64 = -1
                                v60 := GenS4{F0: "\n\t", F1: (fi58 & -1)}
                                for fi61 := 0; fi61 < 6; fi61++ {
                                    if true {
                                        v62 := GenS9{F0: (^(v6)), F1: 0}
                                        v6 -= v6
                                        _ = v62
                                    } else {
                                        switch 42 {
                                        case 9223372036854775807:
                                            if ((1 >= 67) || false) {
                                                if (885 == 2) {
                                                    var v63 float32 = ((float32(1e10) * 3.14159) - (1e10 * (0.0 - 1.5)))
                                                    var v64 int8 = 84
                                                    switch "" {
                                                    case "foo bar":
                                                        if true {
                                                            var v65 string = (("foo bar" + "foo bar") + "hello")
                                                            _ = v65
                                                        }
                                                    case "":
                                                        v2 += (v59 - -1)
                                                    case "hello":
                                                        if (!true) {
                                                            var v66 uint8 = 127
                                                            v63 = v63
                                                            _ = v66
                                                        } else {
                                                            var v67 uint64 = v6
                                                            _ = v67
                                                        }
                                                    default:
                                                        v2 = v2
                                                    }
                                                    _ = v63
                                                    _ = v64
                                                }
                                                v6 += v6
                                            } else {
                                                switch 0 {
                                                case -511:
                                                    var v68 int32 = -715
                                                    _ = v68
                                                }
                                                switch 1 {
                                                case 0:
                                                    switch fi58 {
                                                    case 128:
                                                        var v69 uint32 = 337
                                                        if true {
                                                            var v70 uint8 = 127
                                                            var v71 bool = (1 > 42)
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            if (455 < 0) {
                                                                fmt.Println(v6, v5)
                                                            } else {
                                                                fmt.Println(v6, v2, v59)
                                                                var v72 float32 = float32((1e10 - -2.25))
                                                                _ = v72
                                                            }
                                                            if (true && false) {
                                                                fmt.Println(v69, v2, v6)
                                                                switch ("" + "") {
                                                                case "\n\t":
                                                                    if true {
                                                                        switch (fi58 % ((833) | 1)) {
                                                                        case -305:
                                                                            if (false || true) {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ""
                                                                                }
                                                                            } else {
                                                                                v73 := GenS10{F0: 32767}
                                                                                _ = v73
                                                                            }
                                                                        case -1:
                                                                            var v74 int16 = -1
                                                                            _ = v74
                                                                        default:
                                                                            if (!false) {
                                                                                fmt.Println(v6, v69)
                                                                                v59 += v59
                                                                                var v75 uint32 = (^(v69))
                                                                                _ = v75
                                                                            } else {
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    var v76 bool = true
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                        }
                                                                        var v77 int16 = int16(fi58)
                                                                        var v78 int32 = 0
                                                                        _ = v77
                                                                        _ = v78
                                                                    } else {
                                                                        _ = len(v1)
                                                                        v69 = 42
                                                                    }
                                                                case "":
                                                                    if (100 == genFn0(1)) {
                                                                        var v79 float64 = 1.5
                                                                        _ = v79
                                                                    }
                                                                case "foo bar":
                                                                    v59 += v2
                                                                }
                                                                var v80 int16 = 1
                                                                _ = v80
                                                            } else {
                                                                if false {
                                                                    switch (fi61 >> ((255) & 63)) {
                                                                    case 679:
                                                                        _ = len(v1)
                                                                    case 32768:
                                                                        v81 := GenS11{F0: 127, F1: genFn0(376), F2: (fi61 ^ 2147483647)}
                                                                        _ = v81
                                                                    case -660:
                                                                        v5 -= (fi61 / ((fi58) | 1))
                                                                    default:
                                                                        if false {
                                                                            fmt.Println(v6)
                                                                            v82 := v4["hello"]
                                                                            v83 := GenS12{F0: (v6 << ((42) & 63)), F1: -243}
                                                                            v4["foo bar"] = (fi58 * (fi61 << ((789) & 63)))
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    v6 -= 0
                                                                    _ = len(v4)
                                                                    switch fi58 {
                                                                    case 4294967295:
                                                                        v1 = append(v1, "a")
                                                                    }
                                                                }
                                                            }
                                                            if false {
                                                                delete(v3, 759)
                                                                if true {
                                                                    v69 = (v69 >> ((18446744073709551615) & 63))
                                                                } else {
                                                                    var v84 int64 = 492
                                                                    v85 := GenS7{F0: -128, F1: (^(v69))}
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                if (!((false && (!("\n\t" == "a"))) || true)) {
                                                                    v1 = append(v1, "")
                                                                    var v86 int = 0
                                                                    v69 += v69
                                                                    _ = v86
                                                                }
                                                            }
                                                        }
                                                        if (!true) {
                                                            var v87 int16 = 42
                                                            v5 = (v5 & 127)
                                                            v6 -= (v6 | (v6 * (v6 & (^(v6)))))
                                                            var v88 int16 = int16(v69)
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        _ = v69
                                                    case -2147483648:
                                                        var v89 float64 = ((3.14159 + float64(1.5)) + 1e10)
                                                        _ = v89
                                                    case 4294967295:
                                                        v59 = v59
                                                    default:
                                                        v3[(fi58 + 4294967295)] = fi58
                                                        switch "hello" {
                                                        case "foo bar":
                                                            switch (("\n\t" + ((("a" + "foo bar") + ((("a" + "hello") + "hello") + "foo bar")) + "foo bar")) + "hello") {
                                                            case "a":
                                                                fmt.Println(fi58)
                                                            }
                                                        default:
                                                            if (!true) {
                                                                if (32767 < 416) {
                                                                    var v90 uint16 = 65535
                                                                    _ = v90
                                                                } else {
                                                                    if ((1 >= 255) && false) {
                                                                        _ = len(v1)
                                                                        if ((false && true) && false) {
                                                                            v5 += v5
                                                                            var v91 uint = 165
                                                                            _ = v91
                                                                        } else {
                                                                            var v92 uint8 = 2
                                                                            v2 = (v2 * v59)
                                                                            v93 := v4[("foo bar" + ("foo bar" + "hello"))]
                                                                            v94 := GenS7{F0: -128, F1: 4294967295}
                                                                            var v95 bool = true
                                                                            switch fi61 {
                                                                            case -196:
                                                                                v6 -= (v6 * (v6 % (((v6 - (v6 + (^(v6))))) | 1)))
                                                                            case 819:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ("hello" + "")
                                                                                }
                                                                            case 0:
                                                                                if ((!(772 > 18446744073709551615)) && v95) {
                                                                                    v96 := GenS5{F0: v95, F1: 736, F2: 4294967295}
                                                                                    if false {
                                                                                        var v97 uint = 2
                                                                                        _ = v97
                                                                                    } else {
                                                                                        var v98 string = "\n\t"
                                                                                        v92 = (v92 % (((v92 | 155)) | 1))
                                                                                        _ = v98
                                                                                    }
                                                                                    _ = v96
                                                                                } else {
                                                                                    fmt.Println(v92)
                                                                                    if (255 >= v92) {
                                                                                        var v99 int32 = 0
                                                                                        v100 := GenS1{F0: 554}
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    } else {
                                                                                        v93 -= fi61
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            switch ("hello" + "foo bar") {
                                                                                            case "\n\t":
                                                                                                v93 -= v5
                                                                                            case "foo bar":
                                                                                                if (v6 > 514) {
                                                                                                    v5 = (v93 * -615)
                                                                                                    fmt.Println(v2, v93, v59)
                                                                                                    _ = len(v3)
                                                                                                }
                                                                                            case "hello":
                                                                                                var v101 uint = 0
                                                                                                _ = v101
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v3[(fi58 + (v93 | fi61))] = 42
                                                                                        }
                                                                                        v92 = 42
                                                                                    }
                                                                                    var v102 uint8 = (v92 & v92)
                                                                                    _ = v102
                                                                                }
                                                                            default:
                                                                                v92 -= v92
                                                                                v5 = fi61
                                                                                switch fi58 {
                                                                                case -2147483648:
                                                                                    var v103 int32 = -2147483648
                                                                                    _ = v103
                                                                                }
                                                                            }
                                                                            _ = v92
                                                                            _ = v93
                                                                            _ = v94
                                                                            _ = v95
                                                                        }
                                                                        if (2 > 1) {
                                                                            v2 += v2
                                                                        }
                                                                        v6 = v6
                                                                    } else {
                                                                        if false {
                                                                            v59 = v59
                                                                        }
                                                                    }
                                                                    var v104 uint = 1
                                                                    _ = len(v1)
                                                                    _ = v104
                                                                }
                                                                var v105 uint64 = 2
                                                                _ = v105
                                                            } else {
                                                                v6 = 0
                                                                fmt.Println(v5, fi58)
                                                                var v106 int32 = -2147483648
                                                                _ = len(v3)
                                                                _ = v106
                                                            }
                                                        }
                                                        v2 += 0
                                                        v107 := GenS12{F0: 0, F1: -1}
                                                        _ = v107
                                                    }
                                                    v4["foo bar"] = (fi61 + 9223372036854775807)
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        fmt.Println(v5, v6, fi61)
                                                        v5 += (v5 / ((fi58) | 1))
                                                        var v108 string = ("\n\t" + ("\n\t" + ("" + "a")))
                                                        _ = v108
                                                    default:
                                                        switch (("" + "\n\t") + ("a" + "")) {
                                                        case "\n\t":
                                                            var v109 int16 = -885
                                                            _ = v109
                                                        default:
                                                            var v110 uint64 = (v6 >> ((0) & 63))
                                                            _ = v110
                                                        }
                                                    }
                                                case 2147483647:
                                                    var v111 float32 = 1e10
                                                    _ = v111
                                                }
                                            }
                                        case 42:
                                            v2 -= v2
                                        case -416:
                                            var v112 int8 = 42
                                            _ = v112
                                        }
                                        var v113 uint16 = 2
                                        _ = v113
                                    }
                                    _ = fi61
                                }
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                        }
                        v114 := GenS13{F0: (^(v6))}
                        _ = v114
                    }
                }
                if true {
                    var v115 float32 = (-2.25 + 0.0)
                    v116 := GenS14{F0: 255}
                    fmt.Println(v5, v115)
                    _ = v115
                    _ = v116
                }
                var v117 int32 = -2147483648
                _ = v117
            }
            var v118 bool = true
            _ = v118
        } else {
            var v119 int = (-(v5))
            _ = v119
        }
    }
    fmt.Println(v2)
    switch v5 {
    case 32768:
        var v120 uint = 220
        _ = v120
    default:
        _ = len(v1)
    }
    delete(v3, (v5 << ((42) & 63)))
    fmt.Println(v5)
    var v121 uint = 42
    fmt.Println(v2, v6)
    var v122 int64 = (v2 ^ (v2 | 489))
    _ = len(v1)
    v5 = int(v6)
    var v123 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v121
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
