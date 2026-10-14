package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint8; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 int8 }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 bool; F1 int; F2 int64 }
type GenS12 struct { F0 bool; F1 uint }
type GenS13 struct { F0 string; F1 uint8; F2 int32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint16; F1 int; F2 int8 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int32; F1 int16; F2 int64 }
type GenS18 struct { F0 int16; F1 bool; F2 int32 }
type GenS19 struct { F0 int32; F1 int8 }
func genFn0(p0 uint16, p1 uint16, p2 uint8) int64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int8 = -19
        p0 += p1
        var v7 uint = 340
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int64 = -66
    v8 -= 42
    switch 1 {
    case -876:
        if (!(2147483647 < 73)) {
            v9 := GenS1{F0: 18446744073709551615, F1: "hello", F2: 32767}
            v10 := GenS1{F0: 269, F1: "hello", F2: 32768}
            _ = v10
            _ = v9
        } else {
            var v11 uint16 = p1
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    v1 := []string{(("foo bar" + "\n\t") + "\n\t"), "hello", "foo bar"}
    v2 := []uint32{2147483647}
    v3 := []string{(("\n\t" + ("\n\t" + (("a" + "\n\t") + "hello"))) + "\n\t"), "a", "a"}
    var v4 int = 2147483647
    var v12 int16 = -32768
    var v13 float32 = 1e10
    switch (("a" + ("a" + (("foo bar" + "hello") + "a"))) + ("a" + "foo bar")) {
    case "\n\t":
        v14 := GenS1{F0: 69, F1: "\n\t", F2: v4}
        _ = v14
    default:
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: 2, F1: "", F2: (^(v4))}
            var v17 string = "\n\t"
            v18 := GenS1{F0: 683, F1: v17, F2: 4294967295}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            if (!false) {
                switch 32767 {
                case -1:
                    v2 = append(v2, 4294967295)
                case 32768:
                    var v21 int64 = genFn0(2, 502, 64)
                    _ = v21
                case -4:
                    if (51 == 251) {
                        var v22 float64 = (1.5 * (0.1 - 0.1))
                        var v23 uint32 = 2
                        _ = v22
                        _ = v23
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v24 float32 = (v13 * (v13 + float32(((1e10 + -2.25) / 1.0))))
                        v25 := GenS2{F0: 0}
                        _ = v24
                        _ = v25
                    }
                default:
                    fmt.Println(v4, v4, v4)
                }
                v13 += v13
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v26 := GenS3{F0: false, F1: 127, F2: 100}
                    _ = v26
                }
                switch ("" + "\n\t") {
                case "foo bar":
                    var v27 float32 = v13
                    _ = v27
                case "\n\t":
                    v28 := GenS4{F0: ""}
                    _ = v28
                default:
                    v13 = (v13 * v13)
                }
                v13 = v13
                v4 += v4
            } else {
                switch fi19 {
                case 42:
                    _ = len(v1)
                case 128:
                    v4 = v4
                default:
                    v12 = -32768
                }
                var v29 bool = true
                if (true || v29) {
                    var v30 uint64 = 126
                    var v31 int = (fi20 | (-(fi20)))
                    v30 = (v30 | (v30 | 4294967295))
                    _ = v30
                    _ = v31
                }
                var v32 uint8 = 0
                v33 := GenS3{F0: (0 >= 42), F1: v32, F2: 0}
                _ = v29
                _ = v32
                _ = v33
            }
            var v34 int64 = 0
            fmt.Println(fi20, v34, v34)
            _ = fi20
            _ = v34
        }
        fmt.Println(v4)
        _ = fi19
    }
    var v35 int64 = 42
    var v36 int8 = -128
    for fi37 := 0; fi37 < 9; fi37++ {
        if false {
            v13 -= float32(float64(v13))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v36 -= (v36 & v36)
            v38 := GenS3{F0: false, F1: 200, F2: v36}
            var v39 uint32 = 4294967295
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 9; fi40++ {
            v41 := GenS5{F0: v35, F1: 0}
            v42 := GenS6{F0: 2, F1: (-(v35))}
            var v43 bool = (v12 > v12)
            _ = len(v2)
            var v44 uint32 = 284
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
            _ = v44
        }
        _ = fi37
    }
    fmt.Println(v35)
    v3 = append(v3, "\n\t")
    var v45 int8 = v36
    switch "a" {
    case "\n\t":
        switch (-(v4)) {
        case 42:
            for fi46 := 0; fi46 < 18; fi46++ {
                switch "foo bar" {
                case "hello":
                    var v47 uint64 = 0
                    _ = v47
                case "foo bar":
                    if (false && false) {
                        switch "hello" {
                        case "foo bar":
                            v4 = 942
                        default:
                            v36 = 0
                        }
                        v45 = v45
                        switch (fi46 % ((32767) | 1)) {
                        case 127:
                            v48 := GenS6{F0: 65535, F1: v35}
                            _ = v48
                        case 2147483647:
                            v49 := GenS3{F0: false, F1: 127, F2: (v45 >> ((18446744073709551615) & 63))}
                            _ = v49
                        }
                    }
                case "\n\t":
                    var v50 uint = 255
                    _ = v50
                default:
                    if false {
                        fmt.Println(fi46, v45, fi46)
                        v51 := GenS7{F0: v12}
                        _ = v51
                    } else {
                        var v52 float64 = float64(((1.5 * v13) / v13))
                        _ = v52
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi46
            }
        case -424:
            var v53 int32 = 0
            _ = v53
        }
    case "hello":
        if false {
            v54 := GenS5{F0: v35, F1: 242}
            for fi55 := 0; fi55 < 17; fi55++ {
                var v56 float32 = v13
                v57 := GenS8{F0: 65535}
                for fi58 := 0; fi58 < 20; fi58++ {
                    if (!false) {
                        v35 = v35
                        switch 128 {
                        case -1:
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        default:
                            var v59 int8 = v36
                            _ = v59
                        }
                    } else {
                        v12 = (v12 | int16(v35))
                        v60 := GenS4{F0: "foo bar"}
                        var v61 int = (^(v4))
                        fmt.Println(v12, v35)
                        v12 += 0
                        _ = v60
                        _ = v61
                    }
                    if (255 != 1) {
                        switch "\n\t" {
                        case "":
                            v13 = float32((0.1 * float64(1e10)))
                        case "a":
                            if true {
                                var v62 uint64 = 1
                                _ = v62
                            }
                        case "foo bar":
                            switch fi58 {
                            case 2147483647:
                                var v63 uint64 = 2
                                _ = v63
                            case 42:
                                var v64 float32 = ((v56 - v56) - 1e10)
                                _ = v64
                            case 128:
                                v65 := GenS7{F0: v12}
                                _ = v65
                            default:
                                var v66 float32 = v56
                                _ = v66
                            }
                        default:
                            var v67 int16 = v12
                            _ = v67
                        }
                        fmt.Println(v4, v35, v13)
                        v68 := GenS9{F0: (v45 < 42), F1: -128}
                        v69 := GenS9{F0: true, F1: v36}
                        _ = v68
                        _ = v69
                    } else {
                        switch fi55 {
                        case 127:
                            switch "hello" {
                            case "\n\t":
                                _ = len(v2)
                            }
                        case 32767:
                            v12 += (v12 * v12)
                        default:
                            var v70 int32 = -2147483648
                            v12 += int16(v70)
                            v4 += fi55
                            if false {
                                var v71 uint32 = 2147483647
                                v4 = fi58
                                _ = v71
                            } else {
                                v13 = v13
                                v72 := GenS10{F0: 561, F1: 216}
                                var v73 int32 = 1
                                var v74 uint8 = 2
                                v75 := GenS6{F0: 0, F1: (v35 * (^(v35)))}
                                switch -171 {
                                case -945:
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                default:
                                    v3 = append(v3, "\n\t")
                                }
                                if true {
                                    switch int(v45) {
                                    case 32768:
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                    case 42:
                                        if false {
                                            var v76 string = "hello"
                                            var v77 uint32 = 65535
                                            _ = v76
                                            _ = v77
                                        }
                                    default:
                                        if false {
                                            v1 = append(v1, "hello")
                                            var v78 string = "\n\t"
                                            _ = v78
                                        }
                                    }
                                    var v79 uint8 = 42
                                    v70 = 42
                                    var v80 uint8 = (v79 << ((255) & 63))
                                    if ((((v35 & v35) != 1) || (255 <= 544)) && true) {
                                        var v81 uint8 = 198
                                        _ = v81
                                    }
                                    var v82 int = (^(fi55))
                                    _ = len(v3)
                                    if false {
                                        if (715 != 1) {
                                            v83 := GenS6{F0: 0, F1: 0}
                                            v79 -= (v74 % ((95) | 1))
                                            if false {
                                                if (true && (((v35 & genFn0(1, 32767, (v74 & 145))) == (v35 >> ((1) & 63))) && (1e10 < 0.1))) {
                                                    v4 = fi58
                                                    v84 := GenS11{F0: true, F1: -428, F2: genFn0(255, 989, (v74 << ((318) & 63)))}
                                                    if (v36 != (v36 >> ((2) & 63))) {
                                                        v3 = append(v3, "a")
                                                        v4 = fi55
                                                        _ = len(v1)
                                                    } else {
                                                        if (("a" < "") || false) {
                                                            if false {
                                                                v85 := GenS4{F0: "a"}
                                                                _ = v85
                                                            } else {
                                                                var v86 int = (fi55 >> ((65535) & 63))
                                                                _ = v86
                                                            }
                                                            if (!false) {
                                                                fmt.Println(v73)
                                                            } else {
                                                                v87 := GenS10{F0: 65535, F1: 4294967295}
                                                                var v88 bool = true
                                                                v89 := GenS4{F0: (("hello" + "a") + ("a" + "\n\t"))}
                                                                if (!v88) {
                                                                    var v90 uint = 65535
                                                                    _ = v90
                                                                } else {
                                                                    fmt.Println(fi58, v79, v79)
                                                                }
                                                                if (100 > (v45 & 100)) {
                                                                    v91 := GenS4{F0: "foo bar"}
                                                                    var v92 int64 = (^(v35))
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    v74 = (v80 >> ((18446744073709551615) & 63))
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    fmt.Println(v4, v74)
                                                                    if (int64(v70) < -9223372036854775808) {
                                                                        v13 += 3.14159
                                                                    } else {
                                                                        var v93 float64 = -2.25
                                                                        _ = v93
                                                                    }
                                                                }
                                                                v56 = ((v13 * -2.25) * (0.1 / 0.1))
                                                                _ = v87
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        }
                                                        v74 -= (v80 % ((255) | 1))
                                                        if true {
                                                            fmt.Println(fi58)
                                                            v94 := GenS7{F0: (v12 << ((42) & 63))}
                                                            _ = v94
                                                        }
                                                        switch ("a" + ("hello" + "hello")) {
                                                        case "":
                                                            v80 = v80
                                                        case "a":
                                                            v95 := GenS1{F0: 4294967295, F1: "hello", F2: (fi55 / (((^(v82))) | 1))}
                                                            _ = v95
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = "a"
                                                        }
                                                    }
                                                    _ = v84
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    var v96 int16 = v12
                                                    var v97 int32 = 378
                                                    _ = v96
                                                    _ = v97
                                                case "\n\t":
                                                    fmt.Println(v45)
                                                case "hello":
                                                    var v98 uint64 = 2
                                                    if true {
                                                        if false {
                                                            if (393 < (v35 & (^(v35)))) {
                                                                if false {
                                                                    v56 = (v56 / v13)
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = ""
                                                                    }
                                                                    var v99 bool = false
                                                                    fmt.Println(v13, v70, v56)
                                                                    _ = v99
                                                                }
                                                                v100 := GenS12{F0: true, F1: 2}
                                                                _ = v100
                                                            } else {
                                                                if (417 < 425) {
                                                                    if true {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ""
                                                                        }
                                                                        _ = len(v3)
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v74 -= v80
                                                                        case "\n\t":
                                                                            v12 += (-(v12))
                                                                        default:
                                                                            var v101 float64 = -2.25
                                                                            _ = v101
                                                                        }
                                                                    }
                                                                } else {
                                                                    if (("hello" <= "hello") || true) {
                                                                        fmt.Println(v74, fi58, v82)
                                                                        var v102 uint16 = 32767
                                                                        var v103 float32 = 0.0
                                                                        _ = v102
                                                                        _ = v103
                                                                    }
                                                                    var v104 int8 = (v36 ^ -128)
                                                                    v74 -= (v79 / ((v74) | 1))
                                                                    if (0 != 42) {
                                                                        var v105 uint = 4294967295
                                                                        _ = v105
                                                                    }
                                                                    _ = v104
                                                                }
                                                            }
                                                            v4 = (v82 + -9223372036854775808)
                                                            fmt.Println(v70, fi55, v74)
                                                        }
                                                        var v106 float32 = 0.0
                                                        var v107 int8 = v45
                                                        if ("a" >= "a") {
                                                            var v108 uint16 = 2
                                                            _ = v108
                                                        } else {
                                                            if false {
                                                                var v109 uint16 = 2
                                                                v110 := GenS8{F0: 0}
                                                                v111 := GenS1{F0: 18446744073709551615, F1: ("\n\t" + "foo bar"), F2: (fi58 >> ((336) & 63))}
                                                                _ = v109
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                        }
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                    _ = v98
                                                }
                                            } else {
                                                if (427 != 2) {
                                                    fmt.Println(v36)
                                                    switch (fi55 << ((2) & 63)) {
                                                    case 0:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case -1:
                                                        var v112 uint32 = 65535
                                                        _ = v112
                                                    case 4294967295:
                                                        v113 := GenS1{F0: 1, F1: "a", F2: 527}
                                                        _ = v113
                                                    }
                                                    fmt.Println(v80, v74, v70)
                                                    if false {
                                                        v82 += 75
                                                        v12 = (v12 % ((-1) | 1))
                                                        v114 := GenS6{F0: 65535, F1: v35}
                                                        _ = v114
                                                    }
                                                }
                                                var v115 uint16 = 143
                                                if (!true) {
                                                    v116 := GenS13{F0: "", F1: 22, F2: (v73 << ((18446744073709551615) & 63))}
                                                    var v117 bool = (1 == 603)
                                                    v115 = v115
                                                    _ = v116
                                                    _ = v117
                                                } else {
                                                    if true {
                                                        var v118 int = (-(fi58))
                                                        var v119 uint = 353
                                                        var v120 int32 = 1
                                                        v80 = (v79 - (^(v74)))
                                                        if (!((v119 % ((v119) | 1)) < (v119 << ((42) & 63)))) {
                                                            v70 += v70
                                                            if false {
                                                                v35 += (v35 | (-(v35)))
                                                                switch "" {
                                                                case "":
                                                                    v120 += int32(v12)
                                                                case "hello":
                                                                    if true {
                                                                        v121 := GenS10{F0: 18446744073709551615, F1: (v119 + v119)}
                                                                        _ = v121
                                                                    } else {
                                                                        var v122 int16 = int16(v118)
                                                                        _ = len(v1)
                                                                        _ = v122
                                                                    }
                                                                    var v123 uint = (v119 << ((v119) & 63))
                                                                    if false {
                                                                        var v124 uint32 = 0
                                                                        v2 = append(v2, v124)
                                                                        var v125 int16 = (v12 | v12)
                                                                        if (((v119 << ((435) & 63)) > 18446744073709551615) && false) {
                                                                            v56 = (0.0 / 3.14159)
                                                                        } else {
                                                                            switch "a" {
                                                                            case "hello":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                            v126 := GenS14{F0: "a"}
                                                                            var v127 bool = ((!true) && false)
                                                                            _ = v126
                                                                            _ = v127
                                                                        }
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                    _ = v123
                                                                case "a":
                                                                    switch -2147483648 {
                                                                    case 365:
                                                                        if (v118 == (v118 * v82)) {
                                                                            switch (^(v118)) {
                                                                            case 4294967295:
                                                                                var v128 int64 = (v35 & (v35 >> ((v119) & 63)))
                                                                                _ = v128
                                                                            }
                                                                            fmt.Println(v70, v119, v36)
                                                                            var v129 float64 = 1.5
                                                                            _ = v129
                                                                        } else {
                                                                            v36 += 1
                                                                        }
                                                                    }
                                                                default:
                                                                    var v130 uint = (^(v119))
                                                                    _ = v130
                                                                }
                                                            }
                                                        } else {
                                                            var v131 float32 = v13
                                                            _ = v131
                                                        }
                                                        fmt.Println(v79)
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                    }
                                                }
                                                _ = v115
                                            }
                                            _ = v83
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = len(v3)
                                        }
                                    }
                                    v56 -= (0.0 * 0.0)
                                    _ = v79
                                    _ = v80
                                    _ = v82
                                }
                                if ((^(v74)) >= 255) {
                                    v36 -= v36
                                    var v132 float64 = -2.25
                                    _ = v132
                                } else {
                                    var v133 float64 = (1.5 * -2.25)
                                    v45 = int8(v70)
                                    v134 := GenS3{F0: (true && ((^(v35)) == v35)), F1: v74, F2: v36}
                                    v135 := GenS15{F0: 351, F1: 980, F2: 1}
                                    switch ("" + "foo bar") {
                                    case "hello":
                                        var v136 int8 = (v45 % ((v45) | 1))
                                        _ = v136
                                    case "foo bar":
                                        var v137 uint32 = 4294967295
                                        _ = v137
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    if true {
                                        v13 += 0.1
                                        if true {
                                            v133 -= 1.5
                                        }
                                        v138 := GenS16{F0: "foo bar"}
                                        var v139 uint = 65535
                                        fmt.Println(v12)
                                        _ = v138
                                        _ = v139
                                    }
                                    _ = v133
                                    _ = v134
                                    _ = v135
                                }
                                _ = v72
                                _ = v73
                                _ = v74
                                _ = v75
                            }
                            _ = v70
                        }
                    }
                    _ = fi58
                }
                _ = fi55
                _ = v56
                _ = v57
            }
            _ = v54
        }
    case "a":
        for fi140 := 0; fi140 < 14; fi140++ {
            if true {
                var v141 bool = ((!false) && (false && (!false)))
                for fi142 := 0; fi142 < 20; fi142++ {
                    var v143 float32 = v13
                    var v144 uint = 4294967295
                    fmt.Println(fi142)
                    switch fi140 {
                    case 2147483647:
                        var v145 bool = ((((float32(0.0) - v13) + float32(0.0)) >= float32(3.14159)) && (81 > 127))
                        _ = v145
                    case 32768:
                        v3 = append(v3, "foo bar")
                    default:
                        _ = len(v3)
                    }
                    var v146 uint = 65535
                    if ((v36 * -1) < (v36 / ((v36) | 1))) {
                        var v147 float64 = -2.25
                        fmt.Println(v147, v45)
                        _ = v147
                    } else {
                        var v148 bool = v141
                        _ = v148
                    }
                    _ = fi142
                    _ = v143
                    _ = v144
                    _ = v146
                }
                v2 = append(v2, 872)
                _ = v141
            } else {
                fmt.Println(v12, v45, v45)
                var v149 int64 = (v35 << ((188) & 63))
                _ = v149
            }
            for fi150 := 0; fi150 < 2; fi150++ {
                v3 = append(v3, "hello")
                if (1 >= 127) {
                    switch fi150 {
                    case 854:
                        if (!(false || (v36 > (v45 | 0)))) {
                            v151 := GenS17{F0: -2147483648, F1: v12, F2: 430}
                            v35 += 1
                            v12 += v12
                            _ = v151
                        }
                    }
                }
                fmt.Println(v13)
                _ = fi150
            }
            _ = fi140
        }
    default:
        var v152 uint32 = 2147483648
        _ = v152
    }
    if false {
        for fi153 := 0; fi153 < 16; fi153++ {
            var v154 string = ("\n\t" + "\n\t")
            for fi155 := 0; fi155 < 9; fi155++ {
                fmt.Println(fi153, v12)
                _ = fi155
            }
            switch v154 {
            case "hello":
                for fi156 := 0; fi156 < 9; fi156++ {
                    switch 32767 {
                    case 1:
                        fmt.Println(v36, v35, v4)
                        v3 = append(v3, "\n\t")
                        v157 := GenS18{F0: v12, F1: true, F2: 0}
                        _ = v157
                    default:
                        v2 = append(v2, 2147483647)
                    }
                    fmt.Println(v12, fi156, fi153)
                    fmt.Println(fi156, v36, v45)
                    _ = fi156
                }
            case "\n\t":
                var v158 string = "\n\t"
                _ = v158
            case "":
                fmt.Println(v12, v12)
            default:
                var v159 float32 = ((v13 / (v13 / v13)) - v13)
                _ = v159
            }
            _ = fi153
            _ = v154
        }
        var v160 int64 = v35
        _ = v160
    } else {
        switch v4 {
        case 0:
            var v161 uint16 = 32767
            _ = v161
        default:
            _ = len(v1)
        }
        var v162 string = "a"
        var v163 int8 = v36
        _ = v162
        _ = v163
    }
    var v164 int16 = v12
    if (v45 != v36) {
        v2 = append(v2, 825)
        if (!(false && ((v36 / ((v45) | 1)) < 127))) {
            var v165 uint64 = 18446744073709551615
            v166 := GenS10{F0: 1, F1: 4294967295}
            v167 := GenS19{F0: 2147483647, F1: (v45 / ((v45) | 1))}
            _ = v165
            _ = v166
            _ = v167
        }
        var v168 bool = true
        _ = v168
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v164:", v164)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v12
    _ = v13
    _ = v164
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v45
}
