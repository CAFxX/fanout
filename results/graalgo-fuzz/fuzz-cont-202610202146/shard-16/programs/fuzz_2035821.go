package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 int16; F1 string; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int32; F1 int32; F2 int64 }
type GenS8 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS9 struct { F0 string; F1 uint32; F2 uint }
type GenS10 struct { F0 int; F1 uint }
type GenS11 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS12 struct { F0 uint64; F1 int32; F2 int64 }

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = 1e10
    var v3 bool = true
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    var v6 int16 = 0
    var v7 int8 = 1
    fmt.Println(v6, v1, v3)
    var v8 uint32 = 0
    v9 := GenS1{F0: 1}
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 float64 = float64(1.5)
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: (v6 | (v6 - v6))}
        v14 := GenS1{F0: (v6 / ((v6) | 1))}
        _ = fi10
        _ = v13
        _ = v14
    }
    switch 17 {
    case 128:
        if (v7 != v7) {
            v7 = (^(v7))
            switch 9223372036854775807 {
            case 1:
                var v15 uint = 582
                _ = v15
            case 9223372036854775807:
                for fi16 := 0; fi16 < 16; fi16++ {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if v3 {
                            var v18 bool = v3
                            v19 := v4["foo bar"]
                            switch v19 {
                            case "a":
                                v20 := GenS1{F0: v6}
                                _ = v20
                            default:
                                v21 := []int{(fi17 % ((fi16) | 1)), (fi17 * (fi17 - int(v8))), (-(fi16))}
                                _ = v21
                            }
                            _ = v18
                            _ = v19
                        } else {
                            if true {
                                if (!false) {
                                    if ((v6 ^ (v6 & int16(v7))) < -32768) {
                                        v22 := []int{fi16}
                                        _ = v22
                                    } else {
                                        v23 := GenS1{F0: (v6 / ((v6) | 1))}
                                        _ = v23
                                    }
                                    v6 += (v6 & v6)
                                    if true {
                                        if v3 {
                                            var v24 float64 = (((float64((float32(v2) / (0.0 - (1e10 - 1.5)))) + v1) - v2) - v1)
                                            v7 = 119
                                            v24 = v1
                                            _ = v24
                                        }
                                        v3 = false
                                        fmt.Println(fi16)
                                    } else {
                                        switch fi17 {
                                        case 42:
                                            switch "" {
                                            case "\n\t":
                                                var v25 float32 = 1.5
                                                var v26 float32 = 0.0
                                                _ = v25
                                                _ = v26
                                            }
                                        case -305:
                                            fmt.Println(v7)
                                        case -2147483648:
                                            var v27 uint = 32
                                            _ = v27
                                        }
                                    }
                                    var v28 int = fi16
                                    v8 = v8
                                    _ = v28
                                } else {
                                    var v29 int32 = int32(v8)
                                    v3 = true
                                    _ = v29
                                }
                                v1 = v2
                            } else {
                                var v30 uint32 = v8
                                var v31 int8 = v7
                                v2 = 1e10
                                _ = v30
                                _ = v31
                            }
                            var v32 int16 = 1
                            _ = v32
                        }
                        if v3 {
                            var v33 uint8 = 205
                            _ = v33
                        } else {
                            var v34 int64 = -9223372036854775808
                            v5[(^(fi17))] = (("" + ("foo bar" + "\n\t")) != ("\n\t" + "hello"))
                            v35 := []int{(fi16 + 128), (fi17 & (-(fi16)))}
                            _ = v34
                            _ = v35
                        }
                        fmt.Println(v8, v6, v1)
                        var v36 uint16 = 1
                        _ = fi17
                        _ = v36
                    }
                    _ = fi16
                }
            case 32767:
                switch 128 {
                case -492:
                    if v3 {
                        if v3 {
                            fmt.Println(v7, v2, v2)
                            if (-916 > -392) {
                                var v37 int = 1
                                var v38 bool = (!v3)
                                _ = v37
                                _ = v38
                            } else {
                                v1 -= v1
                                if (((((0.1 - v2) * v2) * ((v1 * ((1e10 / 1.5) / 3.14159)) / v2)) * ((float64(0.1) / (v1 * v1)) * 1e10)) <= v2) {
                                    for fi39 := 0; fi39 < 9; fi39++ {
                                        if v3 {
                                            if (!false) {
                                                var v40 uint32 = 370
                                                _ = v40
                                            } else {
                                                if (v3 || v3) {
                                                    v41 := GenS1{F0: 42}
                                                    _ = v41
                                                }
                                                var v42 string = "\n\t"
                                                if (0 != 1) {
                                                    var v43 float32 = 0.1
                                                    v44 := GenS2{F0: (v8 ^ v8), F1: 194, F2: (((^(fi39)) == 127) || ((v7 - int8(v8)) >= 1))}
                                                    switch "" {
                                                    case "hello":
                                                        switch "" {
                                                        case "a":
                                                            for fi45 := 0; fi45 < 18; fi45++ {
                                                                v8 = 42
                                                                v8 = (v8 >> ((4294967295) & 63))
                                                                v46 := GenS2{F0: 42, F1: 128, F2: v3}
                                                                v47 := GenS1{F0: int16(v8)}
                                                                if v3 {
                                                                    v48 := []int{45, (fi39 << ((945) & 63)), 0, fi39}
                                                                    _ = v48
                                                                } else {
                                                                    var v49 float32 = v43
                                                                    _ = len(v4)
                                                                    v50 := GenS2{F0: (v8 ^ (v8 + v8)), F1: 2, F2: true}
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                                fmt.Println(v6)
                                                                _ = fi45
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        default:
                                                            switch ((("\n\t" + v42) + ((("\n\t" + v42) + "a") + (("\n\t" + v42) + ("\n\t" + ((("foo bar" + (v42 + ("\n\t" + (v42 + v42)))) + (("" + "foo bar") + "foo bar")) + v42))))) + v42) {
                                                            case "hello":
                                                                var v51 int64 = 1
                                                                _ = v51
                                                            case "\n\t":
                                                                for fi52 := 0; fi52 < 20; fi52++ {
                                                                    v53 := GenS3{F0: 18446744073709551615, F1: 80, F2: (v7 + -22)}
                                                                    _ = fi52
                                                                    _ = v53
                                                                }
                                                            default:
                                                                v6 += (v6 << ((0) & 63))
                                                            }
                                                        }
                                                    }
                                                    _ = v43
                                                    _ = v44
                                                }
                                                _ = v42
                                            }
                                        }
                                        _ = fi39
                                    }
                                    delete(v5, 32768)
                                    if false {
                                        v54 := GenS3{F0: 0, F1: 2147483647, F2: (v7 & -128)}
                                        _ = v54
                                    }
                                }
                                var v55 float64 = 1e10
                                _ = v55
                            }
                            var v56 uint8 = 2
                            _ = v56
                        }
                    } else {
                        if v3 {
                            v57 := []int{}
                            fmt.Println(v6, v1)
                            if (255 <= 127) {
                                if (false && true) {
                                    var v58 int8 = v7
                                    var v59 int32 = 0
                                    _ = v58
                                    _ = v59
                                } else {
                                    var v60 int32 = -2147483648
                                    if (v6 != (v6 + v6)) {
                                        switch 9223372036854775807 {
                                        case 0:
                                            var v61 float32 = 0.0
                                            _ = v61
                                        case 127:
                                            var v62 float64 = (float64(3.14159) * v2)
                                            _ = v62
                                        case -905:
                                            if v3 {
                                                fmt.Println(v2, v7, v6)
                                            } else {
                                                fmt.Println(v6, v60)
                                                switch (("foo bar" + "") + ("foo bar" + ("a" + "\n\t"))) {
                                                case "a":
                                                    var v63 uint = 18446744073709551615
                                                    _ = v63
                                                case "\n\t":
                                                    v64 := GenS3{F0: 695, F1: (v60 + int32(v7)), F2: 89}
                                                    _ = v64
                                                case "":
                                                    delete(v4, "hello")
                                                default:
                                                    v65 := GenS4{F0: 32767, F1: "hello", F2: 42}
                                                    _ = v65
                                                }
                                                _ = len(v5)
                                            }
                                        }
                                    } else {
                                        v3 = (false || (-2.25 < (-2.25 - float32((((float64(float32((v2 * v1))) - (3.14159 * (v1 * v2))) - v1) - v2)))))
                                    }
                                    switch 128 {
                                    case 32767:
                                        switch -229 {
                                        case -814:
                                            v66 := GenS1{F0: v6}
                                            _ = v66
                                        default:
                                            for fi67 := 0; fi67 < 17; fi67++ {
                                                v7 += v7
                                                v60 -= -2147483648
                                                _ = fi67
                                            }
                                        }
                                    case 0:
                                        var v68 uint8 = 127
                                        _ = v68
                                    case 32768:
                                        for fi69 := 0; fi69 < 6; fi69++ {
                                            _ = len(v57)
                                            var v70 uint = 2
                                            _ = fi69
                                            _ = v70
                                        }
                                    default:
                                        if (v6 >= -32768) {
                                            var v71 bool = (v60 >= 1)
                                            v72 := GenS5{F0: 42}
                                            _ = v71
                                            _ = v72
                                        } else {
                                            var v73 string = "hello"
                                            fmt.Println(v1)
                                            for fi74 := 0; fi74 < 8; fi74++ {
                                                for fi75 := 0; fi75 < 15; fi75++ {
                                                    if (v1 > 1.5) {
                                                        _ = len(v5)
                                                        v76 := GenS6{F0: (^(v7)), F1: (v6 - 0)}
                                                        var v77 uint64 = 1
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        var v78 int8 = -128
                                                        v73 = v73
                                                        if true {
                                                            v78 = (v78 ^ int8(v6))
                                                            v7 += (v78 / ((42) | 1))
                                                        } else {
                                                            if false {
                                                                var v79 float32 = 1.5
                                                                if ((fi74 ^ fi75) <= (fi75 / (((fi75 % ((1) | 1))) | 1))) {
                                                                    v1 -= v1
                                                                    v8 = 789
                                                                    fmt.Println(fi75)
                                                                } else {
                                                                    var v80 uint32 = v8
                                                                    if len(v57) >= 2 {
                                                                        v57 = v57[1:len(v57)]
                                                                    }
                                                                    _ = len(v57)
                                                                    if false {
                                                                        if v3 {
                                                                            fmt.Println(v79, v2, v1)
                                                                        }
                                                                    } else {
                                                                        v2 += 0.0
                                                                    }
                                                                    if v3 {
                                                                        v2 += (float64((((float32(((1.5 - float64(float32(v2))) + (v1 * (1e10 * ((v1 * (v2 * (v1 * (1.5 - -2.25)))) + 1.5))))) - float32(3.14159)) / (0.1 + (v79 - 1e10))) - v79)) - v1)
                                                                        var v81 int32 = -112
                                                                        var v82 int64 = -296
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                    var v83 string = v73
                                                                    _ = v80
                                                                    _ = v83
                                                                }
                                                                var v84 uint16 = 633
                                                                fmt.Println(v79, v1)
                                                                v85 := GenS6{F0: v7, F1: (v6 ^ -718)}
                                                                if v3 {
                                                                    fmt.Println(fi74)
                                                                    v2 = v1
                                                                    var v86 int8 = 1
                                                                    _ = v86
                                                                } else {
                                                                    v7 = v78
                                                                }
                                                                _ = v79
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                delete(v4, ("" + v73))
                                                            }
                                                            v87 := GenS7{F0: (v60 | v60), F1: (v60 - (^(v60))), F2: -211}
                                                            var v88 int32 = 295
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        _ = v78
                                                    }
                                                    if ((v60 * 1) > 1) {
                                                        var v89 int8 = v7
                                                        v6 = v6
                                                        _ = v89
                                                    }
                                                    _ = fi75
                                                }
                                                v90 := GenS1{F0: v6}
                                                _ = fi74
                                                _ = v90
                                            }
                                            _ = v73
                                        }
                                    }
                                    _ = v60
                                }
                                v91 := GenS8{F0: -524, F1: (v8 % (((v8 << ((4294967295) & 63))) | 1)), F2: 962}
                                var v92 uint = 2
                                _ = v91
                                _ = v92
                            }
                            _ = v57
                        }
                        var v93 bool = true
                        _ = v93
                    }
                case 4294967295:
                    v7 -= v7
                default:
                    v94 := []int{-9223372036854775808}
                    _ = v94
                }
            }
            v95 := GenS9{F0: ("hello" + (("a" + "a") + "foo bar")), F1: v8, F2: 181}
            _ = v95
        }
    case 968:
        delete(v5, 0)
    default:
        if (!(2 < 1)) {
            var v96 uint64 = 724
            fmt.Println(v3)
            v97 := GenS4{F0: v6, F1: ("" + "hello"), F2: 1}
            _ = v96
            _ = v97
        } else {
            v7 -= v7
            if (!v3) {
                _ = len(v5)
                v8 = (v8 ^ 2147483647)
                switch "foo bar" {
                case "":
                    if ((v8 | (^(v8))) < 0) {
                        switch -434 {
                        case -424:
                            var v98 uint16 = 503
                            _ = v98
                        case 1:
                            v7 = v7
                        default:
                            var v99 int64 = -528
                            _ = v99
                        }
                        v100 := GenS5{F0: 457}
                        var v101 int32 = -1
                        _ = v100
                        _ = v101
                    } else {
                        if v3 {
                            var v102 int = 1
                            _ = v102
                        }
                        if (!v3) {
                            delete(v5, 9223372036854775807)
                            var v103 int8 = 42
                            var v104 uint16 = 260
                            _ = v103
                            _ = v104
                        } else {
                            fmt.Println(v3)
                            v105 := GenS10{F0: int(v7), F1: 65535}
                            _ = v105
                        }
                    }
                }
            } else {
                for fi106 := 0; fi106 < 8; fi106++ {
                    v107 := GenS1{F0: (v6 >> ((255) & 63))}
                    _ = fi106
                    _ = v107
                }
                fmt.Println(v8, v3, v1)
                v6 += (^(v6))
            }
        }
    }
    switch 2147483647 {
    case -961:
        for fi108 := 0; fi108 < 3; fi108++ {
            var v109 int32 = -995
            _ = fi108
            _ = v109
        }
    default:
        v1 = 3.14159
    }
    v110 := []int{2147483647, -1, 2147483647, -2147483648}
    var v111 float32 = float32(1.5)
    var v112 bool = (-2147483648 >= -828)
    var v113 uint64 = 18446744073709551615
    for fi114 := 0; fi114 < 4; fi114++ {
        if len(v110) >= 2 {
            v110 = v110[1:len(v110)]
        }
        _ = fi114
    }
    var v115 float64 = (1.5 * (((v2 / v1) / v1) / 1.0))
    _ = len(v110)
    if v112 {
        for fi116 := 0; fi116 < 7; fi116++ {
            var v117 uint16 = 842
            var v118 bool = (1 <= 42)
            var v119 int16 = -127
            _ = fi116
            _ = v117
            _ = v118
            _ = v119
        }
        if false {
            var v120 string = "\n\t"
            for fi121 := 0; fi121 < 15; fi121++ {
                v122 := GenS11{F0: 180, F1: (v7 >> ((1) & 63)), F2: (v8 / ((v8) | 1))}
                v7 = (v7 + v7)
                _ = len(v4)
                _ = fi121
                _ = v122
            }
            var v123 int32 = 340
            _ = v120
            _ = v123
        } else {
            fmt.Println(v111, v111, v112)
            v124 := GenS12{F0: (v113 % ((v113) | 1)), F1: 42, F2: -318}
            v2 -= 3.14159
            _ = v124
        }
        v115 += v115
    }
    v1 = 1.5
    switch ("a" + ("" + "hello")) {
    case "":
        var v125 uint8 = 2
        _ = v125
    default:
        for fi126 := 0; fi126 < 11; fi126++ {
            var v127 int16 = v6
            switch fi126 {
            case 1:
                if v3 {
                    v5[fi126] = (!((^(v7)) <= v7))
                } else {
                    if (((v2 - (3.14159 + v2)) + v2) < ((float64(0.0) * v2) / (v2 * float64(v111)))) {
                        if v112 {
                            v115 = (float64(((((v111 / v111) - (-2.25 + (-2.25 + (v111 - (v111 + (v111 / 1e10)))))) + v111) + v111)) + v1)
                            v128 := GenS4{F0: v6, F1: "", F2: 2}
                            _ = v128
                        }
                    } else {
                        fmt.Println(fi126, v111)
                    }
                    v7 = (-(v7))
                    var v129 int64 = 986
                    _ = v129
                }
            case -2147483648:
                v6 -= (v6 & v6)
            default:
                fmt.Println(v2, v3, v7)
            }
            if ((v112 || v3) || (false && v3)) {
                for fi130 := 0; fi130 < 12; fi130++ {
                    fmt.Println(v6, v1, v2)
                    switch (-(fi126)) {
                    case 32768:
                        var v131 uint32 = (v8 - (v8 + v8))
                        _ = v131
                    case 42:
                        v2 = v115
                        _ = len(v4)
                        var v132 float32 = (3.14159 - v111)
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            _ = len(v110)
                        default:
                            var v133 uint64 = (v113 / (((v113 | 719)) | 1))
                            _ = v133
                        }
                        _ = v132
                    default:
                        var v134 uint32 = v8
                        _ = v134
                    }
                    _ = fi130
                }
                v4[(("\n\t" + "foo bar") + "a")] = ("hello" + ((("hello" + "a") + "hello") + ("\n\t" + "")))
            }
            _ = fi126
            _ = v127
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v110:", len(v110), v110)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v113
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
