package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint8; F2 int64 }
type GenS5 struct { F0 bool; F1 uint; F2 int8 }
type GenS6 struct { F0 string; F1 bool; F2 string }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 uint; F1 uint16 }
type GenS11 struct { F0 int8; F1 int16; F2 int32 }
type GenS12 struct { F0 uint8; F1 int32; F2 bool }
type GenS13 struct { F0 string; F1 int64; F2 int }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 string; F1 int16; F2 int }
func genFn0(p0 uint8, p1 int) int64 {
    v7 := GenS1{F0: "", F1: 42, F2: 32767}
    if (1e10 == (-2.25 / -2.25)) {
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 100
    var v2 int8 = -9
    var v3 uint32 = 65535
    var v4 int32 = 422
    var v5 bool = false
    var v6 uint32 = (v3 % (((v3 & v3)) | 1))
    v8 := GenS1{F0: "a", F1: v1, F2: 0}
    switch -362 {
    case 0:
        var v9 bool = (!(162 > 326))
        _ = v9
    case 817:
        v10 := GenS2{F0: v4, F1: 184}
        _ = v10
    case 991:
        fmt.Println(v6, v5, v1)
    default:
        v11 := GenS3{F0: 493, F1: v3}
        _ = v11
    }
    switch -1 {
    case 2147483647:
        v1 = (^(v2))
    case 4294967295:
        switch "" {
        case "":
            if (460 <= 2) {
                v12 := GenS4{F0: (^(v2)), F1: 42, F2: -845}
                if v5 {
                    if (((!(!false)) && (v2 != (v1 / ((v1) | 1)))) || v5) {
                        for fi13 := 0; fi13 < 14; fi13++ {
                            if v5 {
                                var v14 uint32 = (^(v6))
                                if (!("\n\t" <= ("" + (((("" + "") + "a") + "") + "")))) {
                                    v5 = true
                                    v15 := []int{}
                                    _ = v15
                                }
                                v4 -= (v4 - 2147483647)
                                _ = v14
                            }
                            v16 := GenS5{F0: false, F1: 42, F2: (^(v2))}
                            _ = fi13
                            _ = v16
                        }
                    }
                    v1 = 0
                    fmt.Println(v1, v6)
                }
                _ = v12
            }
        case "hello":
            v6 -= (^(v6))
        case "a":
            switch "" {
            case "\n\t":
                var v17 string = "\n\t"
                _ = v17
            case "a":
                v3 += 4294967295
            }
        default:
            v4 = v4
        }
    default:
        for fi18 := 0; fi18 < 20; fi18++ {
            if v5 {
                v6 = v6
            }
            _ = fi18
        }
    }
    fmt.Println(v6, v2, v3)
    fmt.Println(v4, v1, v4)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := []int{}
        switch "\n\t" {
        case "a":
            if false {
                fmt.Println(v3, v6)
            } else {
                var v21 uint = 18446744073709551615
                _ = v21
            }
        case "foo bar":
            switch fi19 {
            case 42:
                var v22 float32 = ((3.14159 * ((0.1 + 1.5) / float32(1.5))) / 0.1)
                _ = v22
            case -1:
                var v23 uint8 = 151
                _ = v23
            }
        }
        _ = fi19
        _ = v20
    }
    v24 := GenS3{F0: 1, F1: 0}
    if v5 {
        for fi25 := 0; fi25 < 17; fi25++ {
            if (128 != 128) {
                var v26 int64 = 42
                _ = v26
            }
            v5 = v5
            _ = fi25
        }
    } else {
        if (42 < 78) {
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := []int{-9223372036854775808, fi27, -530}
                var v29 uint8 = 1
                _ = fi27
                _ = v28
                _ = v29
            }
            v5 = true
            if v5 {
                var v30 float64 = float64((1.5 / 0.1))
                if (408 >= 142) {
                    if v5 {
                        var v31 bool = v5
                        _ = v31
                    } else {
                        v32 := []int{32767, 267, 32768}
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 11; fi33++ {
                        v34 := GenS6{F0: "hello", F1: (v4 > v4), F2: ((("foo bar" + (("\n\t" + "") + "foo bar")) + "a") + ("a" + "\n\t"))}
                        v35 := GenS4{F0: (v2 ^ v1), F1: 1, F2: 9223372036854775807}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    v30 = (v30 + (float64(-2.25) * ((v30 + v30) * ((1e10 / (v30 + v30)) - ((v30 / (v30 / v30)) * v30)))))
                }
                _ = v30
            }
        } else {
            fmt.Println(v1, v1, v5)
            v36 := GenS1{F0: ("a" + "foo bar"), F1: 127, F2: 1}
            _ = v36
        }
        v1 = 122
    }
    var v37 int8 = v1
    switch "foo bar" {
    case "hello":
        if ((1 >= 0) || (0 < -9223372036854775808)) {
            switch "a" {
            case "foo bar":
                if false {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 int8 = 42
                        v40 := GenS7{F0: 901, F1: 155}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    var v41 int64 = 513
                    _ = v41
                } else {
                    switch ("" + (("\n\t" + "\n\t") + (("" + ("\n\t" + "")) + ""))) {
                    case "a":
                        v42 := []int{-9223372036854775808, -9223372036854775808, 32767}
                        _ = v42
                    case "foo bar":
                        v43 := []int{2147483647, -1, 280}
                        _ = v43
                    }
                    v44 := GenS4{F0: int8(v4), F1: 2, F2: -138}
                    v45 := GenS5{F0: (false || v5), F1: 255, F2: v2}
                    _ = v44
                    _ = v45
                }
            }
        } else {
            switch (("hello" + "hello") + "\n\t") {
            case "hello":
                v37 = (v37 + v37)
            case "\n\t":
                var v46 int16 = 1000
                _ = v46
            case "":
                for fi47 := 0; fi47 < 14; fi47++ {
                    for fi48 := 0; fi48 < 17; fi48++ {
                        v4 += -417
                        v49 := []int{-855, fi48, (fi48 >> ((817) & 63))}
                        v1 = (v37 - (v37 * -1))
                        _ = fi48
                        _ = v49
                    }
                    var v50 uint = 255
                    if (false || (!(float32(1e10) != (0.1 / 3.14159)))) {
                        for fi51 := 0; fi51 < 20; fi51++ {
                            var v52 uint64 = 0
                            _ = fi51
                            _ = v52
                        }
                        fmt.Println(v2)
                    } else {
                        switch "a" {
                        case "foo bar":
                            if v5 {
                                for fi53 := 0; fi53 < 10; fi53++ {
                                    v54 := []int{}
                                    v54 = append(v54, (fi47 >> (((v50 & v50)) & 63)))
                                    _ = fi53
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 18; fi55++ {
                                    var v56 int8 = v37
                                    var v57 uint8 = 1
                                    v58 := []int{fi55}
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                            } else {
                                switch ("a" + "foo bar") {
                                case "hello":
                                    v59 := GenS8{F0: 42}
                                    _ = v59
                                case "":
                                    var v60 int = 999
                                    _ = v60
                                case "a":
                                    var v61 uint32 = (v6 - (^(v6)))
                                    _ = v61
                                default:
                                    v62 := []int{(fi47 % (((fi47 - -1)) | 1)), 128, (fi47 << ((0) & 63)), 643}
                                    _ = v62
                                }
                                var v63 int8 = (v2 / (((^(v37))) | 1))
                                v64 := []int{(^(fi47))}
                                _ = v63
                                _ = v64
                            }
                        case "a":
                            for fi65 := 0; fi65 < 19; fi65++ {
                                var v66 bool = ((true || v5) && true)
                                if true {
                                    v67 := GenS1{F0: "\n\t", F1: (v2 / (((v2 << ((42) & 63))) | 1)), F2: 42}
                                    switch ("a" + ("foo bar" + "a")) {
                                    case "":
                                        if v66 {
                                            v2 = (-(v2))
                                            if v5 {
                                                switch (fi65 * 775) {
                                                case 127:
                                                    var v68 float32 = 0.1
                                                    _ = v68
                                                case -2147483648:
                                                    v1 -= 42
                                                case 32768:
                                                    if v66 {
                                                        v69 := GenS8{F0: -9223372036854775808}
                                                        _ = v69
                                                    } else {
                                                        switch "foo bar" {
                                                        case "":
                                                            if ((v3 / (((v6 % ((583) | 1))) | 1)) <= 0) {
                                                                fmt.Println(v3)
                                                                v70 := []int{fi65, (^(fi47)), fi65}
                                                                var v71 float64 = 1.5
                                                                _ = v70
                                                                _ = v71
                                                            } else {
                                                                v66 = false
                                                                switch 487 {
                                                                case 2147483647:
                                                                    v72 := []int{(fi47 + (fi65 / (((^(fi47))) | 1))), (fi65 ^ (fi47 ^ (fi47 % (((fi47 >> ((255) & 63))) | 1)))), (fi47 << (((v50 / ((v50) | 1))) & 63)), 32768}
                                                                    _ = v72
                                                                case 0:
                                                                    if false {
                                                                        var v73 int16 = 0
                                                                        if v66 {
                                                                            v2 -= v2
                                                                        } else {
                                                                            v73 = v73
                                                                        }
                                                                        _ = v73
                                                                    } else {
                                                                        switch (fi47 ^ fi47) {
                                                                        case 798:
                                                                            v74 := GenS9{F0: ("\n\t" + ""), F1: 18446744073709551615}
                                                                            _ = v74
                                                                        default:
                                                                            if (127 >= 233) {
                                                                                v75 := GenS5{F0: ((3.14159 / 3.14159) <= 0.1), F1: 65535, F2: v37}
                                                                                v2 = (^(v1))
                                                                                _ = v75
                                                                            } else {
                                                                                v5 = (42 <= 85)
                                                                                v66 = false
                                                                                v76 := []int{}
                                                                                _ = v76
                                                                            }
                                                                        }
                                                                        var v77 uint16 = 32767
                                                                        _ = v77
                                                                    }
                                                                }
                                                                if v66 {
                                                                    v78 := GenS6{F0: ("hello" + ("\n\t" + "foo bar")), F1: false, F2: ("hello" + ("a" + ((("foo bar" + ((("" + ("" + "hello")) + ("a" + "hello")) + "foo bar")) + "hello") + ("hello" + ("\n\t" + "a")))))}
                                                                    _ = v78
                                                                }
                                                                v2 -= v2
                                                                v37 -= -128
                                                            }
                                                        case "\n\t":
                                                            v79 := GenS5{F0: v66, F1: (v50 - 18446744073709551615), F2: v37}
                                                            _ = v79
                                                        case "hello":
                                                            v6 = (^(v3))
                                                        }
                                                        fmt.Println(v6)
                                                    }
                                                    var v80 string = ("\n\t" + "")
                                                    _ = v80
                                                default:
                                                    v2 -= (v37 & v37)
                                                    switch "" {
                                                    case "foo bar":
                                                        var v81 int16 = -32768
                                                        _ = v81
                                                    case "":
                                                        switch (-(fi65)) {
                                                        case 42:
                                                            var v82 int64 = -846
                                                            _ = v82
                                                        case -609:
                                                            v1 += (v1 ^ (v37 / (((-(v2))) | 1)))
                                                        }
                                                    case "a":
                                                        if v66 {
                                                            v66 = v66
                                                        }
                                                    }
                                                    if (2 < 42) {
                                                        v50 -= (v50 >> ((326) & 63))
                                                        switch "hello" {
                                                        case "":
                                                            fmt.Println(fi65)
                                                            v66 = v5
                                                        case "a":
                                                            if (v66 && v66) {
                                                                v66 = ((fi47 / ((42) | 1)) <= (fi47 & 128))
                                                            } else {
                                                                v37 += -1
                                                                v83 := []int{(fi65 & 4294967295), fi47, 32768}
                                                                v3 -= 2147483647
                                                                v66 = v66
                                                                _ = v83
                                                            }
                                                        case "hello":
                                                            var v84 uint16 = 32767
                                                            _ = v84
                                                        }
                                                        fmt.Println(v50, v1, v3)
                                                    }
                                                }
                                            } else {
                                                var v85 int16 = -511
                                                if v66 {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v86 int = fi47
                                                        _ = v86
                                                    case "":
                                                        fmt.Println(v50, fi65)
                                                    case "a":
                                                        fmt.Println(v1)
                                                        var v87 string = ((((("hello" + "") + "\n\t") + ("\n\t" + "a")) + ((("hello" + "") + "hello") + "\n\t")) + (("\n\t" + "\n\t") + ""))
                                                        _ = v87
                                                    }
                                                    var v88 uint = v50
                                                    v89 := []int{(fi65 ^ 4294967295), (fi47 - (fi65 >> ((1) & 63))), fi65, (fi65 | fi47)}
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    if v5 {
                                                        v90 := GenS2{F0: v4, F1: uint(v6)}
                                                        var v91 int8 = -128
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        v92 := GenS4{F0: (v2 << (((v50 + v50)) & 63)), F1: 40, F2: 198}
                                                        if (true || true) {
                                                            v5 = (269 >= 557)
                                                            if ((v4 - 0) == v4) {
                                                                v93 := []int{(-(fi65)), fi65, fi65, (^(fi65))}
                                                                _ = v93
                                                            }
                                                            if false {
                                                                v66 = v66
                                                                var v94 uint64 = uint64(v3)
                                                                var v95 uint32 = 2147483648
                                                                _ = v94
                                                                _ = v95
                                                            } else {
                                                                v4 += -361
                                                            }
                                                        }
                                                        if false {
                                                            v2 = (v2 ^ int8(v3))
                                                        } else {
                                                            var v96 int8 = (v1 / (((v2 + (v37 | int8(v50)))) | 1))
                                                            v97 := GenS10{F0: (v50 / ((963) | 1)), F1: 0}
                                                            v85 -= (^(v85))
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        _ = v92
                                                    }
                                                    v98 := []int{(fi65 | int(v50)), fi47, (fi47 & (fi65 - (fi47 / (((fi65 % ((779) | 1))) | 1)))), fi47}
                                                    fmt.Println(v6, v85)
                                                    _ = v98
                                                }
                                                v99 := GenS3{F0: 32767, F1: v3}
                                                _ = v85
                                                _ = v99
                                            }
                                        } else {
                                            v3 = v6
                                            v100 := GenS7{F0: 942, F1: 42}
                                            _ = v100
                                        }
                                    case "foo bar":
                                        if true {
                                            switch (("foo bar" + "hello") + "") {
                                            case "hello":
                                                if false {
                                                    v50 += 2
                                                    var v101 int16 = 1
                                                    if ("a" >= "\n\t") {
                                                        if (v4 >= v4) {
                                                            if v5 {
                                                                v102 := []int{fi65}
                                                                fmt.Println(v4, v5, fi65)
                                                                _ = len(v102)
                                                                _ = v102
                                                            } else {
                                                                var v103 uint = 289
                                                                v104 := GenS11{F0: (v2 & 42), F1: v101, F2: (v4 | (v4 * -2147483648))}
                                                                if v66 {
                                                                    var v105 int8 = v1
                                                                    var v106 uint32 = v6
                                                                    if (fi65 <= -9223372036854775808) {
                                                                        v106 = (v106 * 65535)
                                                                        if (65535 != 32767) {
                                                                            v107 := GenS10{F0: (^(v50)), F1: 749}
                                                                            _ = v107
                                                                        }
                                                                        var v108 int64 = -9223372036854775808
                                                                        fmt.Println(v6, v101)
                                                                        _ = v108
                                                                    }
                                                                    _ = v105
                                                                    _ = v106
                                                                } else {
                                                                    v6 -= (v6 | (v6 ^ (v6 * v6)))
                                                                    v50 = v50
                                                                }
                                                                switch -95 {
                                                                case 32768:
                                                                    v5 = ((0.1 + -2.25) >= ((((0.1 / -2.25) + (1.5 - 0.0)) + 1.5) / 1.0))
                                                                    var v109 int = fi47
                                                                    v110 := GenS7{F0: 1, F1: 42}
                                                                    var v111 uint = (v50 % ((973) | 1))
                                                                    _ = v109
                                                                    _ = v110
                                                                    _ = v111
                                                                case 127:
                                                                    v3 = v3
                                                                    fmt.Println(v1, fi47, v66)
                                                                    var v112 int16 = (v101 * v101)
                                                                    _ = v112
                                                                case 951:
                                                                    var v113 float64 = (1e10 * (1.5 - float64(1e10)))
                                                                    _ = v113
                                                                default:
                                                                    v50 = (^(v103))
                                                                }
                                                                v6 += (v3 - 2147483647)
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            if v66 {
                                                                var v114 uint64 = uint64(v6)
                                                                v2 = v37
                                                                var v115 int64 = 42
                                                                _ = v114
                                                                _ = v115
                                                            } else {
                                                                var v116 uint32 = v3
                                                                _ = v116
                                                            }
                                                            v2 = (-(v1))
                                                        }
                                                        v117 := GenS12{F0: 184, F1: (v4 & (v4 ^ (v4 ^ v4))), F2: true}
                                                        if (!(v37 < v1)) {
                                                            if v66 {
                                                                var v118 int = 331
                                                                v119 := []int{}
                                                                fmt.Println(v66)
                                                                var v120 uint = (v50 % (((^(v50))) | 1))
                                                                _ = v118
                                                                _ = v119
                                                                _ = v120
                                                            } else {
                                                                var v121 uint8 = 255
                                                                var v122 int = (fi65 << ((v50) & 63))
                                                                _ = v121
                                                                _ = v122
                                                            }
                                                            v123 := []int{fi47, fi65, fi47}
                                                            _ = v123
                                                        } else {
                                                            if false {
                                                                switch 634 {
                                                                case 4294967295:
                                                                    v101 = v101
                                                                case -9223372036854775808:
                                                                    v124 := []int{}
                                                                    _ = v124
                                                                }
                                                                v50 += 1
                                                            }
                                                            var v125 float32 = (1.5 - ((-2.25 - float32(0.0)) / ((0.0 - 1e10) + 1.5)))
                                                            _ = v125
                                                        }
                                                        _ = v117
                                                    } else {
                                                        var v126 int16 = (v101 - v101)
                                                        _ = v126
                                                    }
                                                    _ = v101
                                                }
                                            case "a":
                                                var v127 bool = (v5 && ((!v5) || v5))
                                                _ = v127
                                            default:
                                                v50 = (v50 & v50)
                                            }
                                        }
                                    default:
                                        switch (fi47 >> ((v50) & 63)) {
                                        case 4294967295:
                                            v128 := GenS8{F0: -9223372036854775808}
                                            _ = v128
                                        default:
                                            if v66 {
                                                if v66 {
                                                    v129 := []int{(fi65 ^ (fi47 << ((1) & 63))), 453, fi47}
                                                    _ = v129
                                                }
                                            }
                                        }
                                    }
                                    _ = v67
                                } else {
                                    v4 -= int32(v50)
                                }
                                v130 := []int{(fi65 - 4294967295), fi65, (^(fi65)), fi47}
                                _ = fi65
                                _ = v130
                                _ = v66
                            }
                        default:
                            if v5 {
                                for fi131 := 0; fi131 < 18; fi131++ {
                                    fmt.Println(v4, fi47, v4)
                                    _ = fi131
                                }
                            }
                        }
                    }
                    _ = fi47
                    _ = v50
                }
            }
            if false {
                for fi132 := 0; fi132 < 13; fi132++ {
                    for fi133 := 0; fi133 < 10; fi133++ {
                        fmt.Println(fi132)
                        if (v5 || v5) {
                            v37 = v2
                        } else {
                            v4 = -421
                        }
                        v3 = 520
                        _ = fi133
                    }
                    for fi134 := 0; fi134 < 19; fi134++ {
                        var v135 uint = 42
                        _ = fi134
                        _ = v135
                    }
                    _ = fi132
                }
                v136 := GenS13{F0: "a", F1: 42, F2: -9223372036854775808}
                for fi137 := 0; fi137 < 7; fi137++ {
                    v5 = ("hello" <= "a")
                    if v5 {
                        var v138 bool = (!false)
                        v139 := GenS14{F0: v3}
                        for fi140 := 0; fi140 < 12; fi140++ {
                            fmt.Println(v2, v3, fi140)
                            if v5 {
                                var v141 int16 = 0
                                v142 := GenS7{F0: 214, F1: 1}
                                _ = v141
                                _ = v142
                            } else {
                                if (212 == -1) {
                                    var v143 uint = 18446744073709551615
                                    _ = v143
                                }
                                if (v37 != (v1 / ((1) | 1))) {
                                    if (32767 == 42) {
                                        switch ("hello" + (("hello" + "\n\t") + "")) {
                                        case "\n\t":
                                            fmt.Println(v3, v5, v2)
                                        case "foo bar":
                                            fmt.Println(v37, v4, v2)
                                        default:
                                            var v144 int8 = -92
                                            _ = v144
                                        }
                                        var v145 uint64 = 4294967295
                                        _ = v145
                                    }
                                    v146 := []int{(fi140 | 32768), fi140, 127, fi137}
                                    _ = v146
                                } else {
                                    v6 += 1
                                    switch (fi140 ^ 9223372036854775807) {
                                    case -761:
                                        v2 += v2
                                    case -9223372036854775808:
                                        v6 -= (v3 << ((485) & 63))
                                    case 0:
                                        var v147 int8 = v1
                                        var v148 uint = 255
                                        fmt.Println(v4)
                                        _ = v147
                                        _ = v148
                                    default:
                                        switch ("" + "hello") {
                                        case "":
                                            v149 := GenS6{F0: "\n\t", F1: ((216 == 32767) || false), F2: ("" + ("" + "a"))}
                                            v150 := GenS3{F0: 552, F1: v6}
                                            v151 := []int{(fi140 ^ (fi140 | fi140)), fi140, fi137, (fi140 & -140)}
                                            switch ("foo bar" + ("foo bar" + "")) {
                                            case "\n\t":
                                                v5 = (!v138)
                                            case "a":
                                                var v152 uint16 = 0
                                                _ = v152
                                            }
                                            v153 := GenS14{F0: 65535}
                                            if (!(!((v2 >> ((0) & 63)) > (v37 | (v1 + (v37 | 0)))))) {
                                                switch "" {
                                                case "foo bar":
                                                    var v154 string = ("foo bar" + "")
                                                    _ = v154
                                                case "":
                                                    v151 = append(v151, fi137)
                                                }
                                                if len(v151) >= 2 {
                                                    v151 = v151[1:len(v151)]
                                                }
                                            }
                                            _ = v149
                                            _ = v150
                                            _ = v151
                                            _ = v153
                                        }
                                    }
                                    if (!(2 <= 432)) {
                                        if (!false) {
                                            switch "" {
                                            case "foo bar":
                                                if v5 {
                                                    v37 = v2
                                                    fmt.Println(v3, fi137, v5)
                                                    if false {
                                                        v155 := []int{fi137, 9223372036854775807, fi140, (fi137 % ((-2147483648) | 1))}
                                                        _ = len(v155)
                                                        if len(v155) >= 2 {
                                                            v155 = v155[1:len(v155)]
                                                        }
                                                        if false {
                                                            v1 = v2
                                                            switch "hello" {
                                                            case "hello":
                                                                fmt.Println(fi140, fi137)
                                                            case "":
                                                                v4 = v4
                                                            default:
                                                                var v156 int16 = 0
                                                                _ = v156
                                                            }
                                                            if false {
                                                                v157 := GenS5{F0: v5, F1: 1, F2: (v1 ^ (v1 ^ (^(v1))))}
                                                                if len(v155) >= 2 {
                                                                    v155 = v155[1:len(v155)]
                                                                }
                                                                _ = v157
                                                            } else {
                                                                v3 = v6
                                                            }
                                                            var v158 uint64 = 929
                                                            _ = v158
                                                        }
                                                        var v159 int = fi137
                                                        _ = v155
                                                        _ = v159
                                                    } else {
                                                        v160 := GenS13{F0: "hello", F1: 9223372036854775807, F2: fi137}
                                                        fmt.Println(v6, fi137, v5)
                                                        v161 := []int{-46, fi140}
                                                        switch "\n\t" {
                                                        case "hello":
                                                            fmt.Println(v3, v2, v3)
                                                        case "foo bar":
                                                            var v162 int = -434
                                                            _ = v162
                                                        case "\n\t":
                                                            var v163 uint64 = 18446744073709551615
                                                            _ = v163
                                                        }
                                                        _ = v160
                                                        _ = v161
                                                    }
                                                    var v164 int = fi137
                                                    _ = v164
                                                } else {
                                                    if v138 {
                                                        if (0.0 < 1e10) {
                                                            switch (fi140 - -419) {
                                                            case 9223372036854775807:
                                                                v165 := GenS15{F0: 0}
                                                                _ = v165
                                                            }
                                                        } else {
                                                            var v166 uint = 230
                                                            fmt.Println(v1, v6, v138)
                                                            _ = v166
                                                        }
                                                    } else {
                                                        v167 := []int{(fi137 + (fi140 & fi137))}
                                                        _ = len(v167)
                                                        var v168 int = -957
                                                        if v138 {
                                                            v3 -= (v3 % ((v3) | 1))
                                                        } else {
                                                            var v169 float32 = 3.14159
                                                            _ = v169
                                                        }
                                                        v168 += (-(v168))
                                                        _ = v167
                                                        _ = v168
                                                    }
                                                    v170 := GenS10{F0: 323, F1: 65535}
                                                    _ = v170
                                                }
                                            case "":
                                                var v171 uint16 = 42
                                                _ = v171
                                            case "a":
                                                if (v138 || (2 <= 42)) {
                                                    v172 := GenS10{F0: 852, F1: 255}
                                                    _ = v172
                                                }
                                            default:
                                                switch "" {
                                                case "a":
                                                    v173 := GenS7{F0: 32767, F1: 42}
                                                    _ = v173
                                                case "\n\t":
                                                    if false {
                                                        v174 := GenS4{F0: v37, F1: 42, F2: 0}
                                                        _ = v174
                                                    }
                                                case "":
                                                    fmt.Println(v37)
                                                default:
                                                    v138 = false
                                                    v5 = (((4294967295 > 4294967295) || v5) || v138)
                                                    switch int(v3) {
                                                    case 1:
                                                        v175 := GenS12{F0: 1, F1: 0, F2: (v138 || (!(!(0.1 >= (1e10 / (0.1 - 3.14159))))))}
                                                        _ = v175
                                                    case 42:
                                                        var v176 uint16 = 165
                                                        _ = v176
                                                    default:
                                                        v177 := []int{(^(fi140)), fi137, -953}
                                                        v6 -= (v6 ^ v3)
                                                        if len(v177) >= 2 {
                                                            v177 = v177[1:len(v177)]
                                                        }
                                                        _ = v177
                                                    }
                                                }
                                            }
                                            var v178 uint = 255
                                            switch "hello" {
                                            case "hello":
                                                if (-351 != genFn0(1, fi137)) {
                                                    var v179 uint8 = 255
                                                    fmt.Println(v1, v3)
                                                    _ = v179
                                                } else {
                                                    if ((!("foo bar" >= "\n\t")) && v5) {
                                                        var v180 float64 = 3.14159
                                                        _ = v180
                                                    }
                                                    v2 -= v2
                                                    v181 := GenS13{F0: ("\n\t" + "hello"), F1: 0, F2: 32767}
                                                    _ = v181
                                                }
                                            default:
                                                var v182 int = (fi137 / ((-2147483648) | 1))
                                                _ = v182
                                            }
                                            var v183 int16 = -32768
                                            v184 := GenS11{F0: int8(v4), F1: v183, F2: v4}
                                            _ = v178
                                            _ = v183
                                            _ = v184
                                        } else {
                                            v185 := GenS1{F0: "foo bar", F1: v1, F2: 669}
                                            var v186 uint32 = v3
                                            _ = v185
                                            _ = v186
                                        }
                                        fmt.Println(v37, v3)
                                    } else {
                                        var v187 int64 = 0
                                        v188 := GenS16{F0: ("hello" + ("foo bar" + "a")), F1: 1000, F2: 9223372036854775807}
                                        v189 := []int{42, 0, fi140, fi137}
                                        _ = v187
                                        _ = v188
                                        _ = v189
                                    }
                                }
                                v6 += (v3 - v3)
                            }
                            _ = fi140
                        }
                        _ = v138
                        _ = v139
                    }
                    _ = fi137
                }
                _ = v136
            }
        }
    case "\n\t":
        v3 -= v3
    }
    var v190 string = (("foo bar" + ((("a" + ("foo bar" + "a")) + "hello") + ((("\n\t" + ("hello" + ((("hello" + "\n\t") + "foo bar") + ("hello" + "hello")))) + "\n\t") + "foo bar"))) + (("hello" + "\n\t") + "a"))
    for fi191 := 0; fi191 < 18; fi191++ {
        for fi192 := 0; fi192 < 4; fi192++ {
            var v193 float64 = (1e10 - -2.25)
            switch "foo bar" {
            case "foo bar":
                var v194 float64 = (0.0 - v193)
                _ = v194
            case "\n\t":
                v193 += 3.14159
            default:
                v1 += -127
            }
            var v195 int32 = 2147483647
            _ = fi192
            _ = v193
            _ = v195
        }
        _ = fi191
    }
    var v196 int32 = v4
    v2 -= (v37 * v37)
    v2 = v37
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v190:", v190)
    fmt.Println("v196:", v196)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v190
    _ = v196
    _ = v2
    _ = v24
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
