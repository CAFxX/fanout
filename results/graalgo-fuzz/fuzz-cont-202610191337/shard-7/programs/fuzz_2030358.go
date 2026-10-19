package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 int64; F1 uint16; F2 string }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint }
func genFn0() int32 {
    v6 := GenS1{F0: "foo bar"}
    var v7 int = 4294967295
    for fi8 := 0; fi8 < 14; fi8++ {
        v7 -= fi8
        if (18446744073709551615 < 4294967295) {
            var v9 bool = (v7 != fi8)
            if true {
                v10 := GenS1{F0: "hello"}
                switch (fi8 * 4294967295) {
                case 128:
                    v11 := GenS1{F0: "\n\t"}
                    _ = v11
                case -9223372036854775808:
                    v12 := GenS1{F0: "foo bar"}
                    _ = v12
                case 32768:
                    if v9 {
                        var v13 int = v7
                        _ = v13
                    }
                }
                _ = v10
            } else {
                fmt.Println(fi8, fi8, v7)
                if ("\n\t" != "hello") {
                    v7 = (v7 * -142)
                    v14 := GenS1{F0: ""}
                    switch fi8 {
                    case 1:
                        v7 = (v7 * 32768)
                    }
                    _ = v14
                } else {
                    v9 = v9
                }
                if (4294967295 > 2) {
                    var v15 int64 = int64(fi8)
                    v16 := GenS1{F0: "hello"}
                    fmt.Println(v9, v9, v15)
                    _ = v15
                    _ = v16
                }
            }
            _ = v9
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int64, p1 uint8, p2 uint64) uint {
    v17 := GenS1{F0: "hello"}
    fmt.Println(p2, p0, p1)
    fmt.Println(p2, p0)
    v18 := GenS1{F0: ("foo bar" + "a")}
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 2147483647
    var v3 int8 = 1
    v4 := make(map[string]bool)
    var v5 uint = 18446744073709551615
    v2 = (v2 ^ genFn0())
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            delete(v4, (("\n\t" + "") + ((("foo bar" + (((("foo bar" + (("foo bar" + "\n\t") + "hello")) + "a") + "hello") + "a")) + (("a" + "foo bar") + "foo bar")) + "foo bar")))
            v21 := GenS2{F0: 42, F1: 65535, F2: 42}
            _ = fi20
            _ = v21
        }
        v22 := GenS3{F0: -1}
        var v23 uint32 = 2
        _ = fi19
        _ = v22
        _ = v23
    }
    switch 32768 {
    case 128:
        switch -35 {
        case -9223372036854775808:
            var v24 int32 = v2
            _ = v24
        }
    }
    v3 = v3
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := []int{32767, 128, fi25}
        fmt.Println(v3, v5, v2)
        for fi27 := 0; fi27 < 15; fi27++ {
            delete(v4, "\n\t")
            var v28 int32 = v2
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    if false {
        delete(v4, "foo bar")
        if (!(100 < (-(v3)))) {
            v29 := GenS2{F0: 42, F1: (^(v5)), F2: -935}
            v30 := GenS2{F0: 199, F1: 0, F2: 174}
            fmt.Println(v2, v5, v5)
            _ = v29
            _ = v30
        } else {
            var v31 int8 = v3
            v32 := []int{4294967295}
            _ = len(v4)
            _ = v31
            _ = v32
        }
        v2 = genFn0()
    } else {
        if true {
            for fi33 := 0; fi33 < 4; fi33++ {
                for fi34 := 0; fi34 < 20; fi34++ {
                    v3 = -1
                    _ = fi34
                }
                for fi35 := 0; fi35 < 6; fi35++ {
                    if false {
                        var v36 string = "\n\t"
                        _ = v36
                    } else {
                        v2 = (v2 ^ v2)
                        if (-2.25 <= 1.5) {
                            switch fi33 {
                            case 908:
                                switch (("foo bar" + (("\n\t" + (("a" + ("a" + "")) + (("\n\t" + "foo bar") + (("\n\t" + ("a" + "hello")) + (("foo bar" + "hello") + "\n\t"))))) + (("hello" + "\n\t") + (("" + "hello") + "foo bar")))) + ("hello" + "")) {
                                case "a":
                                    if (!true) {
                                        var v37 int64 = -137
                                        if (!(-1 < v2)) {
                                            v38 := []int{fi33, 1, (^(fi35))}
                                            var v39 int8 = -70
                                            _ = v38
                                            _ = v39
                                        }
                                        v40 := []int{fi35}
                                        _ = v37
                                        _ = v40
                                    }
                                case "hello":
                                    fmt.Println(v5, fi35, fi35)
                                case "":
                                    if false {
                                        v2 += v2
                                        v41 := GenS4{F0: v3}
                                        _ = v41
                                    }
                                }
                            case 694:
                                var v42 int16 = 46
                                _ = v42
                            default:
                                v43 := GenS5{F0: 54}
                                _ = v43
                            }
                            switch (("hello" + "\n\t") + "\n\t") {
                            case "foo bar":
                                if true {
                                    switch (fi33 / ((-49) | 1)) {
                                    case 788:
                                        switch ("" + "foo bar") {
                                        case "hello":
                                            v44 := GenS1{F0: "foo bar"}
                                            _ = v44
                                        }
                                        v45 := GenS5{F0: 0}
                                        v46 := GenS4{F0: 65}
                                        if true {
                                            v47 := []int{fi35, 32767, (^(fi35)), fi35}
                                            if true {
                                                v48 := GenS4{F0: v3}
                                                var v49 bool = false
                                                _ = v48
                                                _ = v49
                                            } else {
                                                if (2147483648 != 43) {
                                                    var v50 int8 = (v3 + (v3 ^ v3))
                                                    _ = v50
                                                } else {
                                                    v5 += (v5 % ((v5) | 1))
                                                }
                                                if false {
                                                    var v51 float32 = (3.14159 / -2.25)
                                                    _ = v51
                                                } else {
                                                    fmt.Println(v5, v5, fi35)
                                                }
                                                _ = len(v1)
                                            }
                                            v47 = append(v47, 4294967295)
                                            _ = v47
                                        } else {
                                            if false {
                                                v5 = 42
                                            }
                                            if false {
                                                var v52 uint32 = 966
                                                switch "hello" {
                                                case "a":
                                                    v53 := GenS6{F0: 1000, F1: (v2 & v2)}
                                                    _ = v53
                                                case "foo bar":
                                                    var v54 bool = (18446744073709551615 >= 0)
                                                    _ = v54
                                                }
                                                var v55 float32 = 0.1
                                                var v56 float64 = float64(3.14159)
                                                _ = v52
                                                _ = v55
                                                _ = v56
                                            } else {
                                                if ((!true) || true) {
                                                    var v57 uint = 255
                                                    var v58 string = ((("\n\t" + "foo bar") + ("a" + "hello")) + "\n\t")
                                                    _ = v57
                                                    _ = v58
                                                }
                                                if (-108 <= 815) {
                                                    var v59 float32 = (0.0 / (-2.25 * ((((1e10 - 0.1) / -2.25) - 0.0) / 1.0)))
                                                    v5 = (v5 % ((v5) | 1))
                                                    _ = v59
                                                } else {
                                                    fmt.Println(fi33, v5, v5)
                                                }
                                            }
                                        }
                                        _ = len(v1)
                                        _ = v45
                                        _ = v46
                                    case -1:
                                        var v60 string = (("" + ("" + "hello")) + ((("a" + "foo bar") + "foo bar") + ("hello" + (("" + "foo bar") + ""))))
                                        _ = v60
                                    default:
                                        _ = len(v4)
                                    }
                                    v5 = v5
                                } else {
                                    v61 := GenS5{F0: 128}
                                    var v62 int32 = (v2 * genFn0())
                                    if true {
                                        v63 := GenS7{F0: 0}
                                        _ = v63
                                    } else {
                                        v5 = genFn1(46, 1, 2)
                                        var v64 uint32 = 42
                                        if (!false) {
                                            v5 += v5
                                            var v65 uint32 = (^(v64))
                                            switch (^(fi33)) {
                                            case 2147483647:
                                                v5 = v5
                                            case 962:
                                                var v66 int = fi35
                                                _ = v66
                                            default:
                                                if false {
                                                    var v67 float32 = 3.14159
                                                    _ = v67
                                                } else {
                                                    if true {
                                                        v3 -= (v3 / ((v3) | 1))
                                                    }
                                                }
                                                v5 -= 42
                                            }
                                            if false {
                                                var v68 int8 = -46
                                                v69 := []int{fi35}
                                                fmt.Println(fi33, v68)
                                                _ = v68
                                                _ = v69
                                            } else {
                                                v70 := GenS8{F0: v3}
                                                var v71 uint64 = 512
                                                var v72 string = ""
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                            v65 += 2147483647
                                            var v73 int16 = -32768
                                            var v74 uint = (^(v5))
                                            var v75 uint64 = 918
                                            if (!false) {
                                                fmt.Println(v73, v5, v2)
                                                if false {
                                                    v76 := []int{-9223372036854775808}
                                                    v77 := GenS9{F0: v64}
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    fmt.Println(fi33)
                                                    v74 = (v5 ^ v5)
                                                    v78 := []int{fi33}
                                                    _ = v78
                                                }
                                                var v79 int64 = 626
                                                _ = v79
                                            } else {
                                                var v80 int16 = 1000
                                                v81 := GenS4{F0: v3}
                                                _ = v80
                                                _ = v81
                                            }
                                            _ = v65
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        } else {
                                            v82 := []int{56, (fi33 / (((-(fi33))) | 1)), (fi35 / (((fi33 | fi35)) | 1))}
                                            _ = v82
                                        }
                                        _ = v64
                                    }
                                    _ = v61
                                    _ = v62
                                }
                            default:
                                v5 += genFn1(9223372036854775807, 77, 1)
                            }
                            v83 := GenS6{F0: -815, F1: (v2 + v2)}
                            _ = v83
                        } else {
                            switch (fi33 ^ (-(fi35))) {
                            case 749:
                                var v84 bool = true
                                _ = v84
                            }
                            var v85 int8 = v3
                            _ = v85
                        }
                        switch "hello" {
                        case "hello":
                            v2 = v2
                            v5 = genFn1(-298, 63, 1)
                        default:
                            switch (fi35 & (fi33 % (((fi35 - int(v5))) | 1))) {
                            case -806:
                                v86 := []int{}
                                _ = v86
                            case 9223372036854775807:
                                fmt.Println(v2)
                            case -56:
                                v3 = (-(v3))
                            default:
                                var v87 string = "foo bar"
                                if false {
                                    _ = len(v4)
                                    if (-189 == -433) {
                                        v3 -= 27
                                        if true {
                                            v2 = (v2 * int32(v3))
                                        } else {
                                            var v88 bool = false
                                            if (v5 >= v5) {
                                                var v89 uint = v5
                                                var v90 string = "foo bar"
                                                switch fi33 {
                                                case 2147483647:
                                                    v91 := []int{}
                                                    _ = v91
                                                case 857:
                                                    v2 = (^(v2))
                                                    v87 = (v90 + v87)
                                                    v92 := GenS6{F0: -176, F1: v2}
                                                    _ = v92
                                                case 42:
                                                    v1[v90] = (((v87 + v87) + v90) + ("foo bar" + v87))
                                                default:
                                                    v93 := GenS10{F0: 0, F1: (v89 - genFn1(int64(fi35), 42, 4294967295))}
                                                    v94 := GenS9{F0: 2}
                                                    _ = v93
                                                    _ = v94
                                                }
                                                fmt.Println(v2, v89)
                                                v3 = (v3 | 0)
                                                fmt.Println(v89, fi33)
                                                v95 := GenS6{F0: int16(v5), F1: (v2 >> (((v5 >> ((v89) & 63))) & 63))}
                                                switch v90 {
                                                case "":
                                                    fmt.Println(v3)
                                                }
                                                var v96 int8 = (v3 % ((90) | 1))
                                                v97 := []int{fi33}
                                                _ = v89
                                                _ = v90
                                                _ = v95
                                                _ = v96
                                                _ = v97
                                            } else {
                                                if (127 != 214) {
                                                    v5 = v5
                                                    var v98 int32 = 2147483647
                                                    var v99 int8 = (v3 << ((18446744073709551615) & 63))
                                                    var v100 int8 = (v3 % ((v99) | 1))
                                                    _ = v100
                                                    _ = v98
                                                    _ = v99
                                                }
                                            }
                                            var v101 int8 = (v3 ^ v3)
                                            _ = v101
                                            _ = v88
                                        }
                                        var v102 uint64 = 677
                                        _ = v102
                                    } else {
                                        v5 -= genFn1(42, 170, 2)
                                        if true {
                                            fmt.Println(fi35, v3)
                                            v3 = (v3 | (v3 - v3))
                                            var v103 uint16 = 65535
                                            _ = v103
                                        } else {
                                            switch 4294967295 {
                                            case -9223372036854775808:
                                                switch v87 {
                                                case "foo bar":
                                                    v3 -= (v3 * v3)
                                                case "\n\t":
                                                    var v104 float64 = 1.5
                                                    v87 = (((v87 + v87) + ("" + v87)) + "a")
                                                    _ = v104
                                                case "":
                                                    var v105 uint8 = 0
                                                    _ = v105
                                                default:
                                                    v106 := GenS9{F0: 0}
                                                    v5 += v5
                                                    v5 -= 18446744073709551615
                                                    _ = v106
                                                }
                                            case 127:
                                                v5 = (^(v5))
                                                v2 -= v2
                                                var v107 uint8 = 0
                                                v108 := GenS1{F0: ""}
                                                var v109 uint64 = 1
                                                if true {
                                                    v2 -= v2
                                                } else {
                                                    v110 := GenS9{F0: 110}
                                                    v111 := []int{4294967295}
                                                    _ = v110
                                                    _ = v111
                                                }
                                                _ = v107
                                                _ = v108
                                                _ = v109
                                            }
                                            if false {
                                                switch v87 {
                                                case "hello":
                                                    var v112 uint32 = 4294967295
                                                    _ = v112
                                                default:
                                                    var v113 int = (fi33 * (fi35 / (((fi35 ^ fi33)) | 1)))
                                                    _ = v113
                                                }
                                                v5 -= 1
                                            }
                                            switch (v87 + "") {
                                            case "hello":
                                                v114 := GenS11{F0: 752, F1: 0, F2: v87}
                                                _ = v114
                                            }
                                        }
                                        v115 := GenS8{F0: v3}
                                        _ = v115
                                    }
                                }
                                var v116 float32 = 0.1
                                _ = len(v1)
                                _ = v116
                                _ = v87
                            }
                        }
                    }
                    switch ("hello" + "foo bar") {
                    case "hello":
                        v117 := v4[("foo bar" + ("hello" + (("a" + ("a" + "foo bar")) + "\n\t")))]
                        _ = v117
                    }
                    v5 -= (v5 >> ((v5) & 63))
                    v3 = (v3 / ((v3) | 1))
                    _ = fi35
                }
                _ = fi33
            }
        } else {
            v118 := v4["\n\t"]
            _ = v118
        }
    }
    v3 += 100
    fmt.Println(v2, v3, v2)
    _ = len(v4)
    v119 := GenS2{F0: 255, F1: v5, F2: 0}
    for fi120 := 0; fi120 < 1; fi120++ {
        if true {
            v5 -= v5
        }
        for fi121 := 0; fi121 < 14; fi121++ {
            if true {
                var v122 int64 = 9223372036854775807
                _ = v122
            } else {
                if false {
                    switch "hello" {
                    case "hello":
                        v123 := GenS12{F0: 4294967295}
                        _ = v123
                    }
                    var v124 float64 = 3.14159
                    v125 := GenS6{F0: 1, F1: (v2 | (v2 - 1))}
                    _ = v124
                    _ = v125
                }
                v5 = v5
                switch (fi120 / ((fi120) | 1)) {
                case 69:
                    switch (("foo bar" + ("\n\t" + "")) + "foo bar") {
                    case "foo bar":
                        var v126 bool = true
                        _ = v126
                    case "a":
                        v2 = 0
                    default:
                        v3 -= v3
                    }
                }
            }
            _ = fi121
        }
        _ = fi120
    }
    v5 = 4294967295
    v127 := GenS8{F0: v3}
    var v128 uint64 = 42
    fmt.Println(v2)
    v3 -= v3
    if false {
        var v129 uint64 = v128
        for fi130 := 0; fi130 < 1; fi130++ {
            for fi131 := 0; fi131 < 18; fi131++ {
                v129 += (v128 & (v128 / (((v128 >> ((v5) & 63))) | 1)))
                v132 := []int{}
                _ = fi131
                _ = v132
            }
            var v133 bool = ((42 > 0) && true)
            _ = fi130
            _ = v133
        }
        v2 = -1
        _ = v129
    }
    if ((false && false) || false) {
        if false {
            fmt.Println(v128, v3, v128)
            for fi134 := 0; fi134 < 7; fi134++ {
                v5 += (v5 ^ (v5 % ((genFn1(-9223372036854775808, 127, 1)) | 1)))
                _ = fi134
            }
            v4["\n\t"] = false
        } else {
            if (2147483648 > uint32(v128)) {
                var v135 int32 = int32(v3)
                _ = v135
            } else {
                var v136 uint8 = 198
                v137 := v1["foo bar"]
                if (!false) {
                    switch 810 {
                    case -2147483648:
                        if (true || (!((1e10 + 0.1) < (3.14159 / -2.25)))) {
                            _ = len(v1)
                            v138 := GenS13{F0: 2}
                            _ = v138
                        } else {
                            switch 9223372036854775807 {
                            case -478:
                                for fi139 := 0; fi139 < 1; fi139++ {
                                    fmt.Println(v2)
                                    v128 = (^(v128))
                                    var v140 uint8 = 127
                                    _ = fi139
                                    _ = v140
                                }
                            case 4294967295:
                                for fi141 := 0; fi141 < 15; fi141++ {
                                    v3 += (v3 & v3)
                                    v128 = 18446744073709551615
                                    _ = fi141
                                }
                            }
                            var v142 uint32 = 0
                            _ = v142
                        }
                    default:
                        for fi143 := 0; fi143 < 9; fi143++ {
                            if false {
                                fmt.Println(v128)
                                v144 := []int{(fi143 + (fi143 / ((32768) | 1))), (fi143 >> (((v5 * genFn1(1, 128, v128))) & 63)), fi143, (fi143 * fi143)}
                                _ = v144
                            }
                            fmt.Println(v128, v136)
                            _ = fi143
                        }
                    }
                    var v145 int32 = (v2 - (v2 & genFn0()))
                    _ = v145
                }
                _ = v136
                _ = v137
            }
            if false {
                if (("" + ("" + ("hello" + (((("foo bar" + ("hello" + ((("foo bar" + "") + ("" + (("a" + "hello") + "hello"))) + "a"))) + ("" + "a")) + "a") + (("foo bar" + "") + "hello"))))) > ((("\n\t" + "a") + "foo bar") + ((("foo bar" + "a") + ("a" + ("\n\t" + "a"))) + "a"))) {
                    switch int(v2) {
                    case -817:
                        fmt.Println(v128)
                    case -2147483648:
                        _ = len(v4)
                    case 200:
                        var v146 float64 = 0.1
                        _ = v146
                    default:
                        var v147 uint64 = v128
                        _ = v147
                    }
                    var v148 uint16 = 65535
                    _ = v148
                }
            }
            var v149 uint8 = 255
            _ = v149
        }
    } else {
        var v150 uint8 = 130
        _ = v150
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v119
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
