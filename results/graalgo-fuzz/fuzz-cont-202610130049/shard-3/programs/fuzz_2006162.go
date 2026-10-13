package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int16 }
type GenS5 struct { F0 int32; F1 int16; F2 uint }
type GenS6 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32; F1 uint }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint8; F1 string }
type GenS13 struct { F0 string; F1 uint16 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS17 struct { F0 uint16; F1 bool; F2 int16 }
func genFn0(p0 uint32, p1 string, p2 int32) int8 {
    var v7 int = 127
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 float64 = (-2.25 * 3.14159)
        for fi10 := 0; fi10 < 8; fi10++ {
            if false {
                if false {
                    p2 = p2
                    v11 := GenS1{F0: p2, F1: 42}
                    _ = v11
                } else {
                    p0 -= (p0 ^ p0)
                    v12 := []int{}
                    v12 = append(v12, v7)
                    _ = v12
                }
                v13 := GenS2{F0: (p1 > (p1 + p1)), F1: 1, F2: fi10}
                switch ((p1 + ("" + p1)) + "a") {
                case "a":
                    var v14 int32 = -371
                    _ = v14
                case "\n\t":
                    v15 := GenS3{F0: 42}
                    _ = v15
                case "foo bar":
                    if true {
                        v16 := []int{(fi8 >> ((2) & 63)), (fi10 | fi10)}
                        _ = v16
                    } else {
                        p0 -= (p0 / ((p0) | 1))
                        v17 := GenS1{F0: (p2 | p2), F1: 183}
                        v18 := []int{fi10}
                        _ = v17
                        _ = v18
                    }
                }
                fmt.Println(fi8, v7, p1)
                if true {
                    p1 = (p1 + p1)
                    v19 := GenS2{F0: (1 > 674), F1: 95, F2: (fi10 >> ((825) & 63))}
                    _ = v19
                }
                _ = v13
            }
            v20 := []int{}
            fmt.Println(v7, p0)
            _ = fi10
            _ = v20
        }
        if (937 <= p0) {
            var v21 uint64 = 2
            _ = v21
        }
        _ = fi8
        _ = v9
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        p0 += (p0 & p0)
        var v23 bool = true
        var v24 uint8 = 209
        _ = fi22
        _ = v23
        _ = v24
    }
    if ((1e10 / 0.1) > (1.5 / 3.14159)) {
        if false {
            p1 = p1
            if false {
                for fi25 := 0; fi25 < 3; fi25++ {
                    if true {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            if (((p1 + p1) + p1) < ((p1 + p1) + (p1 + "\n\t"))) {
                                p1 = p1
                                var v27 float32 = float32(0.1)
                                _ = v27
                            } else {
                                var v28 bool = (v7 < v7)
                                var v29 uint8 = 0
                                _ = v28
                                _ = v29
                            }
                            var v30 int8 = 100
                            fmt.Println(p1, p0)
                            _ = fi26
                            _ = v30
                        }
                        switch (p1 + p1) {
                        case "a":
                            v31 := []int{fi25, (fi25 ^ fi25)}
                            _ = v31
                        case "foo bar":
                            var v32 float64 = (0.0 * (1.5 - 0.0))
                            _ = v32
                        case "\n\t":
                            for fi33 := 0; fi33 < 12; fi33++ {
                                v34 := GenS2{F0: true, F1: 219, F2: -794}
                                _ = fi33
                                _ = v34
                            }
                        default:
                            v35 := GenS2{F0: (184 < 127), F1: 255, F2: (v7 << ((1) & 63))}
                            _ = v35
                        }
                        p2 = p2
                    }
                    p1 = (p1 + p1)
                    p2 += (-(p2))
                    _ = fi25
                }
            }
            fmt.Println(v7, p1, p1)
        } else {
            v36 := GenS2{F0: false, F1: 2, F2: v7}
            if (42 == 42) {
                v37 := GenS2{F0: false, F1: 0, F2: (-(v7))}
                _ = v37
            }
            var v38 uint64 = 42
            _ = v36
            _ = v38
        }
        fmt.Println(p0)
    } else {
        var v39 bool = ((1 > 255) || ((2 > 413) || (false && false)))
        var v40 int16 = 32767
        _ = v39
        _ = v40
    }
    v41 := []int{-9223372036854775808, (v7 * v7), (v7 >> ((2) & 63)), (v7 >> ((500) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v41
    _ = v7
    return 127
}

func main() {
    var v1 uint = 2
    v2 := []int{-808}
    v3 := make(map[string]bool)
    var v4 int = 127
    var v5 int32 = 0
    var v6 int = 2147483647
    v4 -= 872
    _ = len(v2)
    var v42 string = (("\n\t" + ("\n\t" + "foo bar")) + "hello")
    for fi43 := 0; fi43 < 17; fi43++ {
        for fi44 := 0; fi44 < 17; fi44++ {
            v3["hello"] = (v5 > v5)
            _ = fi44
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch v4 {
        case 0:
            var v45 int32 = (v5 + v5)
            _ = v45
        default:
            if (false || false) {
                for fi46 := 0; fi46 < 4; fi46++ {
                    if false {
                        delete(v3, ((v42 + v42) + (v42 + v42)))
                        fmt.Println(v42)
                        var v47 uint64 = 651
                        _ = v47
                    } else {
                        fmt.Println(fi46, fi43)
                    }
                    if true {
                        if false {
                            var v48 int = (fi46 ^ (^(fi43)))
                            _ = v48
                        } else {
                            var v49 uint8 = uint8(v1)
                            if true {
                                var v50 uint = v1
                                if (0.0 < (-2.25 / (1.5 * 0.1))) {
                                    fmt.Println(fi46, fi46)
                                } else {
                                    switch v42 {
                                    case "hello":
                                        var v51 int64 = -9223372036854775808
                                        _ = v51
                                    default:
                                        v6 += fi43
                                        if ((!(!true)) && ((845 > 173) || (960 == 0))) {
                                            var v52 int32 = v5
                                            _ = v52
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = v4
                                            }
                                            var v53 float32 = 3.14159
                                            _ = v53
                                        }
                                        v54 := GenS3{F0: 94}
                                        v1 += v50
                                        var v55 uint8 = v49
                                        _ = v54
                                        _ = v55
                                    }
                                    fmt.Println(fi46, v50, v6)
                                    var v56 int = 9223372036854775807
                                    v6 += fi46
                                    _ = v56
                                }
                                v57 := v3[v42]
                                _ = v50
                                _ = v57
                            } else {
                                v58 := GenS4{F0: 855, F1: 221, F2: 1}
                                delete(v3, "\n\t")
                                v42 = "foo bar"
                                _ = v58
                            }
                            v4 -= v4
                            _ = v49
                        }
                        var v59 int = -33
                        v60 := GenS3{F0: 2}
                        _ = v59
                        _ = v60
                    }
                    switch (((v42 + v42) + v42) + "hello") {
                    case "foo bar":
                        switch ((v42 + v42) + v42) {
                        case "a":
                            if false {
                                v61 := GenS3{F0: 59}
                                var v62 float64 = (0.1 + 1.5)
                                _ = v61
                                _ = v62
                            } else {
                                fmt.Println(v1, v1)
                                fmt.Println(v5)
                            }
                        case "foo bar":
                            fmt.Println(v4, v6, fi46)
                        case "hello":
                            v5 = v5
                        default:
                            var v63 uint16 = 1
                            _ = v63
                        }
                    case "\n\t":
                        var v64 int64 = -9223372036854775808
                        _ = v64
                    default:
                        v65 := GenS5{F0: (v5 >> (((v1 % ((v1) | 1))) & 63)), F1: -978, F2: (v1 % (((v1 * 21)) | 1))}
                        _ = v65
                    }
                    _ = fi46
                }
                v66 := GenS5{F0: v5, F1: 32767, F2: (^(v1))}
                if ((true && true) || true) {
                    v67 := GenS2{F0: true, F1: 1, F2: (^(v6))}
                    v6 = 128
                    v68 := GenS2{F0: false, F1: 128, F2: (fi43 * -1)}
                    _ = v67
                    _ = v68
                }
                _ = v66
            }
        }
        _ = fi43
    }
    for fi69 := 0; fi69 < 16; fi69++ {
        if (v4 <= fi69) {
            if false {
                v70 := GenS1{F0: v5, F1: 157}
                fmt.Println(v4)
                _ = v70
            } else {
                _ = len(v2)
            }
            for fi71 := 0; fi71 < 13; fi71++ {
                var v72 uint = (v1 & (v1 & 2))
                v73 := GenS6{F0: 2, F1: 127, F2: int8(v72)}
                if true {
                    switch 127 {
                    case 1:
                        var v74 uint = (v72 + v72)
                        _ = v74
                    }
                    var v75 int16 = 0
                    var v76 uint64 = 4294967295
                    _ = v75
                    _ = v76
                } else {
                    v1 += v1
                }
                v77 := GenS7{F0: 65535}
                _ = fi71
                _ = v72
                _ = v73
                _ = v77
            }
        }
        v78 := GenS8{F0: fi69, F1: 255}
        for fi79 := 0; fi79 < 17; fi79++ {
            _ = len(v3)
            if true {
                switch "hello" {
                case "a":
                    v80 := GenS3{F0: 844}
                    v81 := GenS4{F0: -1, F1: 177, F2: -1}
                    v5 += 2147483647
                    _ = v80
                    _ = v81
                case "":
                    var v82 uint32 = 42
                    _ = v82
                }
                if len(v2) > 0 {
                    v2[0] = fi79
                }
                switch (^(v4)) {
                case -2147483648:
                    switch fi69 {
                    case -264:
                        var v83 string = v42
                        _ = v83
                    case 127:
                        var v84 uint = (v1 / ((v1) | 1))
                        _ = v84
                    case 104:
                        v85 := GenS2{F0: true, F1: 60, F2: (v6 ^ (fi79 - v4))}
                        switch -30 {
                        case 32768:
                            v86 := GenS4{F0: -9223372036854775808, F1: 183, F2: 1}
                            _ = v86
                        case -122:
                            var v87 uint16 = 409
                            _ = v87
                        }
                        _ = v85
                    default:
                        if true {
                            var v88 string = (("hello" + v42) + ("a" + v42))
                            _ = len(v3)
                            var v89 int8 = genFn0(470, (((((v42 + (((v42 + v42) + v42) + ("a" + ("a" + v88)))) + "foo bar") + "foo bar") + v42) + v88), (v5 << (((v1 * v1)) & 63)))
                            _ = v88
                            _ = v89
                        } else {
                            switch (v4 << (((v1 ^ v1)) & 63)) {
                            case -1:
                                if len(v2) > 0 {
                                    v2[0] = v4
                                }
                            case 9223372036854775807:
                                var v90 string = ""
                                _ = v90
                            }
                            if (65535 <= 42) {
                                var v91 string = ((v42 + "hello") + v42)
                                _ = v91
                            }
                            var v92 string = ("" + "\n\t")
                            _ = v92
                        }
                    }
                }
            } else {
                v6 -= (^(v4))
            }
            v93 := GenS9{F0: 433}
            _ = fi79
            _ = v93
        }
        _ = fi69
        _ = v78
    }
    if true {
        v2 = append(v2, v6)
        fmt.Println(v4, v42, v1)
        for fi94 := 0; fi94 < 13; fi94++ {
            v1 -= v1
            switch fi94 {
            case 128:
                switch -216 {
                case -563:
                    fmt.Println(v6)
                case 1:
                    for fi95 := 0; fi95 < 15; fi95++ {
                        v96 := v3[(v42 + "a")]
                        v96 = (v96 && (!false))
                        _ = fi95
                        _ = v96
                    }
                case 4294967295:
                    v97 := GenS8{F0: (v6 % ((-73) | 1)), F1: 42}
                    _ = v97
                default:
                    for fi98 := 0; fi98 < 15; fi98++ {
                        var v99 int32 = v5
                        var v100 int16 = 32767
                        var v101 float32 = 0.1
                        v3[""] = (1.5 >= (0.0 * (0.1 * 0.0)))
                        v102 := GenS10{F0: 2147483647, F1: v1}
                        _ = fi98
                        _ = v100
                        _ = v101
                        _ = v102
                        _ = v99
                    }
                }
            default:
                v4 = v6
            }
            if true {
                v103 := GenS11{F0: 706}
                if (true || (2147483648 < 2147483648)) {
                    var v104 uint = (^(v1))
                    for fi105 := 0; fi105 < 17; fi105++ {
                        fmt.Println(v1, fi94)
                        _ = fi105
                    }
                    _ = len(v3)
                    _ = v104
                } else {
                    for fi106 := 0; fi106 < 1; fi106++ {
                        if (1 <= 0) {
                            fmt.Println(fi94)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(fi106)
                        }
                        fmt.Println(v42)
                        _ = fi106
                    }
                    if ((^(v1)) > v1) {
                        v42 = ("\n\t" + (v42 + ("foo bar" + v42)))
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v103
            } else {
                switch v42 {
                case "foo bar":
                    var v107 float32 = (0.0 / 1.0)
                    _ = v107
                case "hello":
                    if false {
                        var v108 uint16 = 255
                        var v109 float64 = 0.1
                        _ = v108
                        _ = v109
                    } else {
                        for fi110 := 0; fi110 < 17; fi110++ {
                            v111 := GenS12{F0: 255, F1: v42}
                            var v112 uint8 = 60
                            _ = fi110
                            _ = v111
                            _ = v112
                        }
                        var v113 int8 = 100
                        v1 += (v1 & v1)
                        _ = v113
                    }
                case "a":
                    var v114 float32 = 0.1
                    _ = v114
                default:
                    fmt.Println(v4, v6, v1)
                }
                var v115 int16 = -32768
                for fi116 := 0; fi116 < 17; fi116++ {
                    v4 -= (v6 / (((fi94 >> ((v1) & 63))) | 1))
                    _ = fi116
                }
                _ = v115
            }
            _ = fi94
        }
    }
    v5 = (v5 % (((v5 & 0)) | 1))
    v117 := GenS13{F0: "foo bar", F1: 1}
    for fi118 := 0; fi118 < 1; fi118++ {
        fmt.Println(v42, fi118)
        fmt.Println(fi118)
        v119 := GenS14{F0: 0, F1: 128}
        _ = fi118
        _ = v119
    }
    v120 := GenS15{F0: true}
    switch "" {
    case "a":
        fmt.Println(v5)
    case "hello":
        var v121 float32 = float32(0.1)
        _ = v121
    }
    var v122 int = 4294967295
    if ((v5 | v5) == 2147483647) {
        v123 := v3[(((v42 + v42) + "hello") + ("" + "hello"))]
        _ = v123
    }
    v2 = append(v2, 995)
    if (-128 >= 0) {
        var v124 string = ""
        if len(v2) > 0 {
            v2[0] = v6
        }
        var v125 string = (v124 + v124)
        _ = v124
        _ = v125
    } else {
        switch "hello" {
        case "hello":
            v126 := GenS16{F0: v1, F1: 65535, F2: 2147483647}
            _ = v126
        case "a":
            v5 = (v5 / ((v5) | 1))
        case "foo bar":
            var v127 int = (-(v6))
            _ = v127
        default:
            var v128 string = (v42 + (v42 + (((v42 + v42) + ("hello" + ("" + "\n\t"))) + "\n\t")))
            _ = v128
        }
        v1 = 65535
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v5)
    for fi129 := 0; fi129 < 9; fi129++ {
        for fi130 := 0; fi130 < 16; fi130++ {
            switch (v6 << ((503) & 63)) {
            case 9223372036854775807:
                if true {
                    fmt.Println(fi130, v5, fi129)
                    v42 = "a"
                    if false {
                        v1 = v1
                        v131 := GenS17{F0: 2, F1: true, F2: int16(v122)}
                        v132 := v3[(v42 + v42)]
                        _ = v131
                        _ = v132
                    } else {
                        var v133 string = v42
                        delete(v3, ("" + v133))
                        _ = v133
                    }
                } else {
                    var v134 string = v42
                    v122 += (v122 / ((2147483647) | 1))
                    if true {
                        switch -321 {
                        case 0:
                            v4 -= 0
                        case -985:
                            var v135 float64 = float64(-2.25)
                            _ = v135
                        default:
                            var v136 string = v134
                            var v137 bool = (!false)
                            var v138 uint = (v1 + (v1 ^ (v1 << ((1) & 63))))
                            _ = v136
                            _ = v137
                            _ = v138
                        }
                        if (false && (0 != 0)) {
                            v139 := v3[(v134 + "\n\t")]
                            if (!false) {
                                fmt.Println(v5, fi130, v134)
                                v140 := GenS14{F0: uint64(v1), F1: 0}
                                _ = v140
                            } else {
                                var v141 int = v6
                                _ = v141
                            }
                            _ = v139
                        }
                    }
                    _ = v134
                }
            case 42:
                var v142 int32 = v5
                _ = v142
            }
            _ = fi130
        }
        if true {
            if false {
                v1 += 65535
                v5 = 42
            } else {
                if false {
                    var v143 uint16 = 54
                    _ = v143
                } else {
                    v6 -= v122
                }
                fmt.Println(fi129)
            }
            var v144 int32 = v5
            v42 = v42
            _ = v144
        }
        if ((^(v1)) != (v1 % (((v1 * (v1 ^ 65535))) | 1))) {
            fmt.Println(v1)
            if len(v2) > 0 {
                v2[0] = 674
            }
        }
        _ = fi129
    }
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v120
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
