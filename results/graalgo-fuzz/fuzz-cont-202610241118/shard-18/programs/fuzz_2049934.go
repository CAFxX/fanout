package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 int; F2 uint }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint8; F1 uint }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int64; F1 string }
type GenS15 struct { F0 int64; F1 uint8; F2 uint8 }
func genFn0(p0 uint64) int32 {
    var v5 int32 = 887
    var v6 int = -9223372036854775808
    v6 = -615
    fmt.Println(v6, v6, v6)
    _ = p0
    _ = v5
    _ = v6
    return (-(v5))
}

func main() {
    var v1 int = -111
    var v2 bool = false
    v3 := []string{"foo bar"}
    var v4 float64 = (((0.1 + (1.5 - 0.1)) + 0.1) / 3.14159)
    var v7 int8 = 1
    for fi8 := 0; fi8 < 2; fi8++ {
        v7 = v7
        _ = fi8
    }
    v2 = v2
    v9 := GenS1{F0: -1, F1: 1}
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int = v1
        switch fi10 {
        case 114:
            v12 := GenS2{F0: 0}
            _ = v12
        case -9223372036854775808:
            fmt.Println(v11, v4)
        case -2147483648:
            if (-104 <= (v7 / ((v7) | 1))) {
                if v2 {
                    switch ("a" + ("\n\t" + (("\n\t" + "foo bar") + "a"))) {
                    case "\n\t":
                        v3 = append(v3, ((("\n\t" + "foo bar") + "hello") + "a"))
                    }
                    if (255 == 552) {
                        v3 = append(v3, "foo bar")
                        v13 := GenS1{F0: -1, F1: 0}
                        for fi14 := 0; fi14 < 17; fi14++ {
                            if len(v3) > 0 {
                                v3[0] = ("\n\t" + ("a" + ("hello" + "hello")))
                            }
                            v4 += (((float64((1.5 + 1.5)) / float64(((0.1 - (0.0 + 1e10)) / float32(1e10)))) * v4) + v4)
                            if true {
                                switch (v11 - v11) {
                                case -128:
                                    if v2 {
                                        var v15 int16 = -32768
                                        if v2 {
                                            var v16 float64 = v4
                                            v2 = v2
                                            v17 := GenS2{F0: 0}
                                            _ = v16
                                            _ = v17
                                        } else {
                                            v2 = (v2 || (((1e10 - float32(v4)) >= ((((1.5 / 1.0) * 3.14159) + (1e10 * 0.1)) - (0.1 / 1e10))) || ((42 != 1) || (v4 > v4))))
                                            var v18 uint32 = 0
                                            var v19 int32 = 2147483647
                                            _ = v18
                                            _ = v19
                                        }
                                        var v20 uint64 = 481
                                        _ = v15
                                        _ = v20
                                    }
                                case 32767:
                                    if false {
                                        var v21 int16 = 1000
                                        _ = v21
                                    }
                                }
                            }
                            _ = fi14
                        }
                        _ = v13
                    }
                } else {
                    var v22 float32 = (-2.25 * float32(v4))
                    if ((v2 && (1 < 2)) && (42 != 1)) {
                        _ = len(v3)
                        v23 := GenS3{F0: 42}
                        v24 := GenS3{F0: -1}
                        _ = v23
                        _ = v24
                    } else {
                        v7 = 42
                        v22 = v22
                        for fi25 := 0; fi25 < 2; fi25++ {
                            _ = len(v3)
                            v11 -= (fi25 << ((525) & 63))
                            _ = fi25
                        }
                    }
                    fmt.Println(v2, v2, v4)
                    _ = v22
                }
                switch ("hello" + ("foo bar" + "hello")) {
                case "foo bar":
                    switch "hello" {
                    case "a":
                        if v2 {
                            var v26 int = v11
                            v7 -= v7
                            if ((^(v26)) < (v26 % (((fi10 - (v11 - (v1 % ((fi10) | 1))))) | 1))) {
                                v27 := GenS4{F0: 42}
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    var v29 uint8 = 0
                                    if (v29 != (v29 >> ((2) & 63))) {
                                        v1 -= v11
                                    } else {
                                        v26 -= (v11 ^ (fi10 >> ((255) & 63)))
                                        var v30 float64 = v4
                                        if v2 {
                                            var v31 int8 = (^(v7))
                                            _ = v31
                                        }
                                        _ = v30
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                v32 := GenS5{F0: -9223372036854775808}
                                _ = v27
                                _ = v32
                            }
                            _ = v26
                        } else {
                            v33 := GenS2{F0: 1}
                            for fi34 := 0; fi34 < 7; fi34++ {
                                switch ("" + (("hello" + "foo bar") + (("foo bar" + "foo bar") + ("foo bar" + ("" + "a"))))) {
                                case "foo bar":
                                    var v35 uint16 = 758
                                    _ = v35
                                case "":
                                    if false {
                                        fmt.Println(v7)
                                    }
                                    _ = len(v3)
                                    v4 += (3.14159 / (v4 / -2.25))
                                    var v36 float32 = -2.25
                                    _ = v36
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                }
                                _ = fi34
                            }
                            if (65535 <= 1) {
                                fmt.Println(v7)
                                v37 := GenS5{F0: 33}
                                _ = v37
                            } else {
                                var v38 float64 = 0.0
                                for fi39 := 0; fi39 < 6; fi39++ {
                                    v11 -= 9223372036854775807
                                    switch -2147483648 {
                                    case 9223372036854775807:
                                        fmt.Println(v1)
                                    case -2147483648:
                                        v40 := GenS6{F0: (^(v7)), F1: v11, F2: 18446744073709551615}
                                        _ = v40
                                    case 2147483647:
                                        var v41 uint = 18446744073709551615
                                        v42 := GenS1{F0: -1, F1: 1}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = fi39
                                }
                                _ = v38
                            }
                            _ = v33
                        }
                    case "foo bar":
                        for fi43 := 0; fi43 < 10; fi43++ {
                            if v2 {
                                var v44 uint = 1
                                _ = v44
                            }
                            var v45 uint64 = 18446744073709551615
                            if (!true) {
                                v2 = true
                            } else {
                                fmt.Println(fi10)
                                var v46 int16 = -491
                                _ = v46
                            }
                            fmt.Println(v4, v2)
                            _ = fi43
                            _ = v45
                        }
                    default:
                        var v47 uint = 838
                        _ = v47
                    }
                case "\n\t":
                    if (((v1 >> ((970) & 63)) != 42) && true) {
                        if (210 <= 42) {
                            v7 += 1
                            for fi48 := 0; fi48 < 18; fi48++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi48
                            }
                            for fi49 := 0; fi49 < 7; fi49++ {
                                switch fi10 {
                                case 127:
                                    v2 = false
                                }
                                _ = fi49
                            }
                        } else {
                            if v2 {
                                var v50 float32 = 0.1
                                var v51 uint8 = 63
                                for fi52 := 0; fi52 < 4; fi52++ {
                                    fmt.Println(v2, v4)
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "foo bar")
                                    }
                                    if true {
                                        v53 := GenS1{F0: -277, F1: 525}
                                        switch (fi10 - v11) {
                                        case -952:
                                            if v2 {
                                                _ = len(v3)
                                            } else {
                                                if (false || v2) {
                                                    v2 = v2
                                                    v54 := GenS7{F0: (float64(((v50 - v50) * (0.0 * 1e10))) <= ((v4 + (((0.0 * v4) / float64(v50)) / float64(((0.0 - v50) * ((1.5 + (((((v50 * 0.1) + ((-2.25 * v50) / v50)) * float32((1.5 * float64((float32(float64((v50 * ((0.0 - v50) + (3.14159 + v50))))) + 0.0))))) + ((v50 / (v50 * v50)) * v50)) / float32(((v4 / (3.14159 - v4)) - v4)))) * v50))))) / v4)), F1: 1}
                                                    _ = v54
                                                } else {
                                                    if v2 {
                                                        v51 = v51
                                                        var v55 int8 = 127
                                                        _ = v55
                                                    }
                                                }
                                            }
                                        case 4294967295:
                                            var v56 int8 = (v7 << ((1) & 63))
                                            _ = v56
                                        }
                                        v57 := GenS3{F0: 1}
                                        var v58 uint = 0
                                        _ = v53
                                        _ = v57
                                        _ = v58
                                    }
                                    _ = fi52
                                }
                                _ = v50
                                _ = v51
                            } else {
                                v59 := GenS8{F0: 127, F1: 42}
                                _ = v59
                            }
                            var v60 float64 = ((v4 * 1e10) * (v4 + (v4 / (0.0 - ((v4 - 1.5) * 1.5)))))
                            _ = v60
                        }
                        _ = len(v3)
                    }
                }
                v3 = append(v3, "hello")
            } else {
                fmt.Println(v2)
            }
        default:
            switch "foo bar" {
            case "\n\t":
                for fi61 := 0; fi61 < 2; fi61++ {
                    var v62 bool = (-263 >= -332)
                    switch "" {
                    case "hello":
                        if len(v3) > 0 {
                            v3[0] = ("a" + "a")
                        }
                    case "a":
                        switch (fi61 * (v1 % ((v1) | 1))) {
                        case 1:
                            v1 += (^(fi10))
                        case 32767:
                            if ((!(v62 && v62)) || (32767 == 918)) {
                                v63 := GenS9{F0: 2}
                                v64 := GenS9{F0: 4294967295}
                                if (-2147483648 == -2147483648) {
                                    var v65 uint8 = 201
                                    v1 -= (fi10 - fi10)
                                    if len(v3) > 0 {
                                        v3[0] = ("hello" + (("foo bar" + "hello") + ""))
                                    }
                                    fmt.Println(v11, v65, v2)
                                    _ = v65
                                }
                                _ = v63
                                _ = v64
                            } else {
                                var v66 int32 = 2147483647
                                fmt.Println(v62, v62, fi61)
                                _ = v66
                            }
                        case 4294967295:
                            fmt.Println(v2, v4, fi10)
                        }
                    }
                    var v67 int64 = 9223372036854775807
                    _ = fi61
                    _ = v62
                    _ = v67
                }
            }
        }
        if true {
            for fi68 := 0; fi68 < 6; fi68++ {
                v69 := GenS9{F0: 42}
                v4 -= (v4 + v4)
                _ = fi68
                _ = v69
            }
        } else {
            v3 = append(v3, "foo bar")
        }
        _ = fi10
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = (("foo bar" + "") + "")
    }
    var v70 string = ("hello" + "")
    v1 -= v1
    switch v1 {
    case -611:
        v71 := GenS9{F0: 572}
        _ = v71
    case 545:
        var v72 int16 = 1
        _ = v72
    case -131:
        var v73 int = (-(v1))
        _ = v73
    }
    v70 = (v70 + v70)
    var v74 bool = v2
    v2 = (v74 || v74)
    if v2 {
        var v75 bool = true
        fmt.Println(v7)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v75
    } else {
        v74 = (!v2)
        for fi76 := 0; fi76 < 14; fi76++ {
            if (!(255 != 0)) {
                for fi77 := 0; fi77 < 20; fi77++ {
                    v78 := GenS10{F0: v74}
                    var v79 uint16 = 32767
                    v3 = append(v3, "\n\t")
                    _ = fi77
                    _ = v78
                    _ = v79
                }
                if v2 {
                    switch v70 {
                    case "hello":
                        fmt.Println(fi76, v70)
                    default:
                        var v80 uint64 = 0
                        _ = v80
                    }
                    var v81 int64 = 0
                    _ = len(v3)
                    _ = v81
                } else {
                    var v82 float64 = v4
                    var v83 int32 = -2147483648
                    _ = v82
                    _ = v83
                }
                if v74 {
                    var v84 int32 = 476
                    _ = v84
                } else {
                    if v74 {
                        for fi85 := 0; fi85 < 4; fi85++ {
                            if true {
                                if true {
                                    var v86 int32 = -2147483648
                                    _ = v86
                                } else {
                                    var v87 string = (("hello" + ("a" + (v70 + (v70 + ("hello" + v70))))) + (v70 + v70))
                                    v4 -= v4
                                    fmt.Println(fi76)
                                    switch v1 {
                                    case -845:
                                        v88 := GenS11{F0: 304}
                                        _ = v88
                                    case -1:
                                        fmt.Println(v74)
                                    default:
                                        v89 := GenS10{F0: (255 <= 217)}
                                        _ = v89
                                    }
                                    if ((-9223372036854775808 == -702) || (42 < 0)) {
                                        var v90 int16 = 151
                                        var v91 int16 = v90
                                        v92 := GenS12{F0: ((v74 && (v2 && (!(2147483648 <= 2147483647)))) && (65535 != 4294967295))}
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    }
                                    v93 := GenS9{F0: 0}
                                    if v74 {
                                        if true {
                                            v94 := GenS8{F0: 1, F1: 18446744073709551615}
                                            _ = v94
                                        } else {
                                            var v95 int32 = genFn0(18446744073709551615)
                                            switch v70 {
                                            case "foo bar":
                                                var v96 uint64 = 246
                                                _ = v96
                                            case "":
                                                v70 = v70
                                            case "hello":
                                                if false {
                                                    var v97 float32 = float32((v4 / 1.0))
                                                    var v98 int = (fi85 | 1)
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    fmt.Println(v2, v1)
                                                    v95 -= (v95 << ((653) & 63))
                                                    if true {
                                                        v99 := GenS5{F0: fi76}
                                                        _ = v99
                                                    } else {
                                                        v74 = v2
                                                        v3 = append(v3, "hello")
                                                        if ((((("hello" + v70) < v70) || v74) || true) && v74) {
                                                            v100 := GenS1{F0: (-(v95)), F1: 18446744073709551615}
                                                            var v101 int64 = 9223372036854775807
                                                            _ = v100
                                                            _ = v101
                                                        } else {
                                                            fmt.Println(fi76)
                                                            if len(v3) > 0 {
                                                                v3[0] = v70
                                                            }
                                                        }
                                                    }
                                                    v87 = v70
                                                    v1 = v1
                                                }
                                            }
                                            v74 = v74
                                            _ = v95
                                        }
                                    } else {
                                        var v102 int32 = 284
                                        fmt.Println(v2, v7)
                                        if (v87 == v70) {
                                            v103 := GenS13{F0: 468}
                                            v2 = (3.14159 == ((0.1 * 1.5) + (float32(v4) + ((1e10 * 1.5) / (1e10 * 3.14159)))))
                                            var v104 uint16 = 2
                                            _ = v103
                                            _ = v104
                                        } else {
                                            fmt.Println(v2, v74)
                                            v105 := GenS1{F0: -2147483648, F1: 18446744073709551615}
                                            _ = v105
                                        }
                                        _ = v102
                                    }
                                    _ = v87
                                    _ = v93
                                }
                                var v106 float64 = float64((1.5 / (3.14159 + 1e10)))
                                v106 -= (1e10 * (-2.25 + (float64(3.14159) + v106)))
                                var v107 uint16 = 42
                                _ = v106
                                _ = v107
                            } else {
                                v2 = ((("foo bar" + v70) + ((v70 + "a") + (v70 + v70))) == "a")
                                v108 := GenS12{F0: v74}
                                var v109 int64 = 42
                                _ = v108
                                _ = v109
                            }
                            v110 := GenS12{F0: (v2 || false)}
                            _ = fi85
                            _ = v110
                        }
                        v3 = append(v3, (v70 + (v70 + (("\n\t" + v70) + v70))))
                    } else {
                        switch v1 {
                        case 0:
                            v2 = v2
                        case 42:
                            fmt.Println(v70)
                        }
                    }
                }
            } else {
                var v111 int32 = -1
                fmt.Println(v7)
                v3 = append(v3, "")
                _ = v111
            }
            var v112 uint16 = 740
            if v2 {
                v7 += (v7 >> ((912) & 63))
                var v113 bool = true
                v114 := GenS14{F0: 1, F1: (v70 + "a")}
                _ = v113
                _ = v114
            } else {
                for fi115 := 0; fi115 < 8; fi115++ {
                    if true {
                        if v74 {
                            var v116 uint8 = 42
                            v4 = v4
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v116
                        } else {
                            if (v74 && ((fi76 == v1) && v74)) {
                                v1 += fi76
                            } else {
                                v117 := GenS7{F0: (!(v4 == v4)), F1: v112}
                                if (0 <= 0) {
                                    switch "a" {
                                    case "":
                                        v4 -= float64(-2.25)
                                    case "a":
                                        switch v70 {
                                        case "foo bar":
                                            v118 := GenS10{F0: (65535 >= 42)}
                                            _ = v118
                                        default:
                                            _ = len(v3)
                                        }
                                    default:
                                        switch v70 {
                                        case "hello":
                                            var v119 uint32 = 2147483648
                                            _ = v119
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = "\n\t"
                                            }
                                        }
                                    }
                                } else {
                                    if ((130 >= 156) && (v2 && (!true))) {
                                        _ = len(v3)
                                    }
                                }
                                v120 := GenS15{F0: -1, F1: 42, F2: 22}
                                _ = v117
                                _ = v120
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v121 uint32 = 65535
                        _ = v121
                    }
                    var v122 int16 = -1
                    var v123 uint16 = v112
                    _ = fi115
                    _ = v122
                    _ = v123
                }
                var v124 float64 = (0.1 + float64((float32(v4) - 0.1)))
                _ = v124
            }
            _ = fi76
            _ = v112
        }
    }
    switch v1 {
    case 4294967295:
        if (!v74) {
            if len(v3) > 0 {
                v3[0] = ("\n\t" + (v70 + v70))
            }
        } else {
            switch 0 {
            case 256:
                v70 = v70
            default:
                v7 = v7
            }
        }
    case 1:
        for fi125 := 0; fi125 < 18; fi125++ {
            for fi126 := 0; fi126 < 6; fi126++ {
                v3 = append(v3, (v70 + (v70 + (v70 + ((v70 + (v70 + (v70 + v70))) + v70)))))
                _ = fi126
            }
            _ = fi125
        }
    default:
        if (-1 != 0) {
            var v127 uint8 = 234
            switch v70 {
            case "a":
                v128 := GenS9{F0: 0}
                _ = v128
            }
            _ = v127
        } else {
            for fi129 := 0; fi129 < 15; fi129++ {
                v74 = v2
                _ = fi129
            }
            var v130 string = v70
            var v131 bool = ((!(!false)) || v2)
            _ = v130
            _ = v131
        }
    }
    for fi132 := 0; fi132 < 8; fi132++ {
        fmt.Println(fi132)
        _ = fi132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v70
    _ = v74
    _ = v9
}
