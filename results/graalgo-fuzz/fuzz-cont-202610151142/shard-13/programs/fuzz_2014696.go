package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint16; F2 uint }
type GenS6 struct { F0 int16; F1 bool; F2 int8 }
type GenS7 struct { F0 int16; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 int16 }
type GenS11 struct { F0 int; F1 int16; F2 int8 }
type GenS12 struct { F0 uint64; F1 int16 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int8 = 127
        var v9 int = 0
        var v10 float32 = (float32(0.0) * 0.1)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch "foo bar" {
    case "a":
        v11 := GenS1{F0: "", F1: 0, F2: 32768}
        _ = v11
    }
    return 458
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 608
    var v3 int8 = -128
    v4 := []uint32{754, 291, 104, 860}
    var v5 float32 = 0.1
    var v6 int8 = (v3 + v3)
    if ((174 < 2) || true) {
        v12 := GenS1{F0: "foo bar", F1: 65535, F2: 0}
        _ = v12
    } else {
        fmt.Println(v5, v2)
    }
    v6 -= 1
    var v13 uint8 = 42
    var v14 uint8 = v13
    if false {
        v15 := GenS2{F0: false, F1: 255}
        var v16 uint64 = 4294967295
        if (true && true) {
            switch -1 {
            case 0:
                v17 := GenS3{F0: v14, F1: 42}
                _ = v17
            default:
                for fi18 := 0; fi18 < 15; fi18++ {
                    if true {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            if true {
                                if true {
                                    v2 += v2
                                    if (true && true) {
                                        if true {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = len(v4)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        var v20 bool = (!true)
                                        var v21 uint = 0
                                        _ = v20
                                        _ = v21
                                    }
                                } else {
                                    v22 := GenS4{F0: 4294967295}
                                    _ = v22
                                }
                                if false {
                                    v23 := GenS4{F0: 42}
                                    _ = v23
                                }
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                            }
                            var v24 bool = true
                            _ = fi19
                            _ = v24
                        }
                        v25 := GenS1{F0: "\n\t", F1: 42, F2: int(v3)}
                        var v26 float32 = float32(0.1)
                        _ = v25
                        _ = v26
                    }
                    _ = fi18
                }
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                for fi28 := 0; fi28 < 20; fi28++ {
                    var v29 uint32 = 57
                    v30 := GenS5{F0: -1, F1: 65535, F2: 208}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
        } else {
            for fi31 := 0; fi31 < 10; fi31++ {
                v32 := v1["\n\t"]
                var v33 int32 = -1
                _ = fi31
                _ = v32
                _ = v33
            }
            fmt.Println(v6, v16)
        }
        _ = v15
        _ = v16
    } else {
        switch "" {
        case "":
            v5 = (v5 + v5)
        case "\n\t":
            v3 -= (-(v6))
        case "hello":
            for fi34 := 0; fi34 < 2; fi34++ {
                v13 -= (v14 % (((v14 >> ((273) & 63))) | 1))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v13 -= (v13 << ((255) & 63))
                _ = fi34
            }
        }
        if true {
            v3 = (v6 & (v6 * (v6 | int8(v2))))
            for fi35 := 0; fi35 < 10; fi35++ {
                switch fi35 {
                case 127:
                    if true {
                        switch (fi35 >> ((2) & 63)) {
                        case -9223372036854775808:
                            v36 := GenS2{F0: ((-(v2)) >= (v2 % ((v2) | 1))), F1: 0}
                            _ = v36
                        case 4294967295:
                            fmt.Println(fi35, v5, v6)
                        default:
                            v1["\n\t"] = false
                        }
                    } else {
                        if ((!true) && (("" + ("hello" + "a")) >= "")) {
                            fmt.Println(v5, v6)
                            var v37 string = "foo bar"
                            _ = v37
                        }
                    }
                case 724:
                    v38 := GenS6{F0: 913, F1: (0.0 == -2.25), F2: (v6 * (v3 | 0))}
                    _ = v38
                }
                _ = fi35
            }
            v6 = 113
        }
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        var v40 string = "foo bar"
        var v41 uint8 = v13
        v4 = append(v4, 702)
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi42
    }
    v13 = v13
    _ = len(v4)
    v4 = append(v4, 2147483648)
    if true {
        var v43 bool = true
        if (42 >= 42) {
            v44 := GenS1{F0: "foo bar", F1: 246, F2: 32768}
            if ((!(true || ((v43 && v43) && (v43 && v43)))) || false) {
                _ = len(v1)
                v45 := v1[""]
                for fi46 := 0; fi46 < 11; fi46++ {
                    v5 -= (v5 / v5)
                    _ = fi46
                }
                _ = v45
            } else {
                v47 := GenS6{F0: 0, F1: (int8(v2) >= 127), F2: -121}
                if len(v4) > 0 {
                    v4[0] = 1
                }
                for fi48 := 0; fi48 < 12; fi48++ {
                    var v49 uint = 1
                    _ = fi48
                    _ = v49
                }
                _ = v47
            }
            _ = v44
        } else {
            for fi50 := 0; fi50 < 16; fi50++ {
                var v51 uint8 = v14
                if true {
                    fmt.Println(v2, v3, v5)
                    switch ("" + ("\n\t" + "")) {
                    case "hello":
                        if v43 {
                            for fi52 := 0; fi52 < 3; fi52++ {
                                var v53 uint32 = 42
                                _ = fi52
                                _ = v53
                            }
                            v51 -= 1
                            v3 += 127
                        } else {
                            fmt.Println(fi50, v5, v14)
                        }
                    }
                    _ = len(v4)
                } else {
                    v51 = v51
                    if true {
                        var v54 uint32 = 42
                        var v55 float32 = 1e10
                        switch ("" + "a") {
                        case "a":
                            if ((v54 % ((4294967295) | 1)) <= (v54 >> ((2) & 63))) {
                                var v56 int16 = -32768
                                v57 := GenS5{F0: v2, F1: 32767, F2: 784}
                                _ = v56
                                _ = v57
                            } else {
                                var v58 int = (fi50 >> ((1) & 63))
                                var v59 bool = v43
                                _ = v58
                                _ = v59
                            }
                        case "\n\t":
                            var v60 uint64 = 18446744073709551615
                            _ = v60
                        default:
                            v13 = v14
                        }
                        _ = v54
                        _ = v55
                    } else {
                        fmt.Println(v2)
                    }
                    var v61 bool = true
                    _ = v61
                }
                fmt.Println(v6, v14, v14)
                _ = fi50
                _ = v51
            }
            var v62 string = "\n\t"
            _ = v62
        }
        for fi63 := 0; fi63 < 2; fi63++ {
            if true {
                switch 128 {
                case 128:
                    v6 += (v6 - (-(v3)))
                case 0:
                    v13 += v14
                case -846:
                    var v64 int32 = 2147483647
                    _ = v64
                default:
                    v65 := GenS6{F0: 64, F1: v43, F2: (v3 ^ 42)}
                    _ = v65
                }
                v2 = v2
                var v66 uint32 = 42
                _ = v66
            }
            var v67 uint8 = 2
            _ = fi63
            _ = v67
        }
        _ = v43
    }
    v3 += -34
    v68 := GenS1{F0: "hello", F1: 1, F2: 2147483647}
    var v69 float64 = (float64(3.14159) / -2.25)
    if true {
        var v70 int64 = (v2 + v2)
        _ = v70
    } else {
        delete(v1, ("hello" + "a"))
        v3 = (v3 + -31)
    }
    for fi71 := 0; fi71 < 2; fi71++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v5)
        switch 2147483647 {
        case -748:
            var v72 float32 = v5
            _ = v72
        case 0:
            var v73 uint8 = v14
            _ = v73
        default:
            if ((fi71 / ((fi71) | 1)) > fi71) {
                var v74 string = ("hello" + ("hello" + "hello"))
                for fi75 := 0; fi75 < 12; fi75++ {
                    switch (fi71 / ((fi71) | 1)) {
                    case 1:
                        delete(v1, v74)
                    case 32768:
                        if (2 >= 4294967295) {
                            var v76 uint64 = 4294967295
                            v77 := v1[(v74 + "\n\t")]
                            v76 += uint64(v13)
                            _ = v76
                            _ = v77
                        } else {
                            var v78 float64 = float64(0.0)
                            v79 := GenS7{F0: 60, F1: genFn0(), F2: 1}
                            if true {
                                v14 = 128
                                if false {
                                    v14 -= (^(v13))
                                    if (65535 < 1) {
                                        var v80 float64 = -2.25
                                        _ = v80
                                    } else {
                                        v74 = "a"
                                        if (!((v2 << ((91) & 63)) <= v2)) {
                                            _ = len(v1)
                                            v5 += v5
                                        }
                                        var v81 uint8 = (v14 << ((0) & 63))
                                        _ = v81
                                    }
                                }
                                if (false || true) {
                                    if true {
                                        if (!false) {
                                            if true {
                                                v2 -= v2
                                            }
                                            var v82 float64 = (0.0 / v69)
                                            _ = v82
                                        } else {
                                            v5 = (v5 - 1.5)
                                            switch -766 {
                                            case -922:
                                                if ((!true) && false) {
                                                    v4 = append(v4, 1)
                                                    if len(v4) > 0 {
                                                        v4[0] = 397
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v83 float32 = v5
                                                        _ = v83
                                                    case "\n\t":
                                                        if ((false || (!(false && true))) || (!true)) {
                                                            if ((v13 >> ((0) & 63)) < 180) {
                                                                _ = len(v1)
                                                                var v84 float32 = v5
                                                                _ = v84
                                                            }
                                                        } else {
                                                            if (301 <= -999) {
                                                                switch v74 {
                                                                case "foo bar":
                                                                    if false {
                                                                        v1[v74] = false
                                                                    }
                                                                }
                                                                v2 -= v2
                                                            } else {
                                                                var v85 uint = 0
                                                                v86 := GenS4{F0: v85}
                                                                switch (-(fi75)) {
                                                                case 128:
                                                                    var v87 int16 = -419
                                                                    _ = v87
                                                                case 32768:
                                                                    var v88 uint64 = 2
                                                                    _ = v88
                                                                default:
                                                                    if true {
                                                                        v3 += -1
                                                                        v89 := GenS2{F0: ((fi75 & 55) > (fi71 >> ((777) & 63))), F1: (v85 % ((v85) | 1))}
                                                                        _ = v89
                                                                    } else {
                                                                        if (((!true) && true) && (false || true)) {
                                                                            v90 := GenS2{F0: true, F1: v85}
                                                                            v2 -= int64(v85)
                                                                            _ = v90
                                                                        } else {
                                                                            var v91 uint8 = 211
                                                                            fmt.Println(v74)
                                                                            if (!false) {
                                                                                v78 = v69
                                                                                if false {
                                                                                    var v92 bool = (fi71 != 2147483647)
                                                                                    v93 := GenS8{F0: uint16(v85), F1: 4294967295, F2: 42}
                                                                                    switch v74 {
                                                                                    case "\n\t":
                                                                                        _ = len(v1)
                                                                                    case "hello":
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = 1
                                                                                        }
                                                                                    case "a":
                                                                                        var v94 float32 = float32(v78)
                                                                                        _ = v94
                                                                                    }
                                                                                    v95 := GenS9{F0: false}
                                                                                    _ = v92
                                                                                    _ = v93
                                                                                    _ = v95
                                                                                } else {
                                                                                    v78 -= v78
                                                                                    var v96 int = -1
                                                                                    _ = v96
                                                                                }
                                                                                fmt.Println(v69, v2)
                                                                            } else {
                                                                                fmt.Println(v78, fi75, v85)
                                                                                v13 = 127
                                                                                switch (-(fi71)) {
                                                                                case 1:
                                                                                    v3 += (v6 - (v3 | int8(v2)))
                                                                                case -967:
                                                                                    var v97 string = ((v74 + "\n\t") + v74)
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                }
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                            fmt.Println(fi75)
                                                        }
                                                    }
                                                    if true {
                                                        if (("" + "hello") <= v74) {
                                                            var v98 float64 = 1.5
                                                            v99 := GenS10{F0: v2, F1: -32768}
                                                            switch (v74 + "foo bar") {
                                                            case "":
                                                                var v100 int8 = (v6 >> ((2) & 63))
                                                                _ = v100
                                                            case "foo bar":
                                                                var v101 int16 = 36
                                                                _ = v101
                                                            case "\n\t":
                                                                fmt.Println(v98, v2, v98)
                                                            default:
                                                                v13 += (v13 | v14)
                                                            }
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            switch v74 {
                                                            case "foo bar":
                                                                v6 = v3
                                                            case "":
                                                                _ = len(v4)
                                                                switch fi75 {
                                                                case 127:
                                                                    var v102 int64 = -918
                                                                    _ = v102
                                                                }
                                                                v3 = v3
                                                            case "\n\t":
                                                                fmt.Println(v14, v13)
                                                            default:
                                                                var v103 int8 = (^(v3))
                                                                _ = v103
                                                            }
                                                        }
                                                    } else {
                                                        var v104 uint16 = 978
                                                        if false {
                                                            var v105 bool = false
                                                            v74 = (v74 + "\n\t")
                                                            switch (("" + "") + v74) {
                                                            case "":
                                                                var v106 int64 = v2
                                                                _ = v106
                                                            case "hello":
                                                                if v105 {
                                                                    v104 += v104
                                                                    v107 := GenS11{F0: 0, F1: int16(v3), F2: -72}
                                                                    _ = v107
                                                                }
                                                            default:
                                                                if v105 {
                                                                    var v108 uint16 = v104
                                                                    var v109 uint32 = 2
                                                                    _ = v108
                                                                    _ = v109
                                                                }
                                                            }
                                                            var v110 uint = 411
                                                            _ = v105
                                                            _ = v110
                                                        }
                                                        var v111 int = 32768
                                                        switch v74 {
                                                        case "":
                                                            v112 := GenS12{F0: 909, F1: -1}
                                                            _ = v112
                                                        case "hello":
                                                            if false {
                                                                var v113 int16 = int16(v3)
                                                                v104 = (^(v104))
                                                                v78 = (-2.25 + v69)
                                                                _ = v113
                                                            } else {
                                                                v114 := GenS12{F0: genFn0(), F1: 372}
                                                                var v115 int64 = (-(v2))
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                        case "a":
                                                            v2 += v2
                                                        }
                                                        _ = v104
                                                        _ = v111
                                                    }
                                                }
                                            }
                                            v5 = v5
                                        }
                                        var v116 uint8 = (v14 + v14)
                                        _ = v116
                                    }
                                    v14 += (v14 - 127)
                                } else {
                                    v74 = v74
                                    v14 += v13
                                }
                            } else {
                                var v117 bool = false
                                _ = v117
                            }
                            _ = v78
                            _ = v79
                        }
                    case -2147483648:
                        v118 := GenS1{F0: "", F1: 18446744073709551615, F2: (fi71 & (fi75 * -2147483648))}
                        _ = v118
                    default:
                        v14 = (v13 / ((v14) | 1))
                    }
                    _ = fi75
                }
                _ = v74
            }
        }
        _ = fi71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
}
