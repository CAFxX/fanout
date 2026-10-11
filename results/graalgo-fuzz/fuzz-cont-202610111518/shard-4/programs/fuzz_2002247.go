package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 bool }
type GenS5 struct { F0 int8; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint16; F1 int64; F2 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint64; F1 int32 }
type GenS16 struct { F0 uint; F1 uint32 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint32; F1 uint }
func genFn0(p0 int8, p1 int32) uint64 {
    fmt.Println(p0, p0)
    fmt.Println(p0, p1)
    if false {
        p1 = (p1 << ((18446744073709551615) & 63))
    }
    var v6 uint8 = 1
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 uint8) bool {
    p0 += (p0 / ((226) | 1))
    fmt.Println(p0)
    _ = p0
    return false
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]int)
    var v3 int32 = -704
    var v4 uint16 = (^(v1))
    var v5 string = ""
    var v7 float64 = (float64(-2.25) * ((0.0 / float64(3.14159)) * float64(1e10)))
    switch 128 {
    case 42:
        if ((1000 != -599) || true) {
            fmt.Println(v4)
        }
    case -9223372036854775808:
        var v8 int64 = -787
        _ = v8
    case 127:
        v9 := GenS1{F0: false}
        _ = v9
    default:
        var v10 bool = true
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if (127 >= 0) {
            var v12 uint8 = 128
            v4 += (v4 - v1)
            v13 := GenS1{F0: (!(0 == (v3 ^ (v3 % (((v3 | v3)) | 1)))))}
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    v7 += (((0.1 - (v7 + v7)) * (v7 + (1e10 - v7))) + float64(0.1))
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{fi14, (fi14 + (fi14 << ((255) & 63))), (fi14 >> ((18446744073709551615) & 63)), (fi14 | (fi14 * fi14))}
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := v2[((v5 + (v5 + v5)) + "")]
            switch -437 {
            case 229:
                var v18 int = -715
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            v7 = v7
            v21 := GenS1{F0: true}
            _ = fi20
            _ = v21
        }
        v5 = v5
        if false {
            switch fi19 {
            case 866:
                var v22 uint16 = (v1 ^ 255)
                _ = v22
            case 9223372036854775807:
                for fi23 := 0; fi23 < 5; fi23++ {
                    var v24 float32 = (1.5 * 0.0)
                    var v25 bool = (v7 > v7)
                    var v26 float32 = (((float32((v7 - -2.25)) - ((v24 - v24) * (v24 - 3.14159))) - v24) * -2.25)
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            var v27 string = "foo bar"
            _ = v27
        }
        _ = fi19
    }
    v28 := GenS2{F0: 42, F1: -1}
    for fi29 := 0; fi29 < 2; fi29++ {
        switch -260 {
        case -202:
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 uint32 = 2147483648
                switch (fi30 >> ((18446744073709551615) & 63)) {
                case 2147483647:
                    var v32 uint8 = 127
                    _ = v32
                case 4294967295:
                    fmt.Println(v4, v3, v5)
                case 0:
                    if (v5 == "\n\t") {
                        var v33 bool = false
                        v7 -= v7
                        var v34 int8 = 1
                        _ = v33
                        _ = v34
                    }
                default:
                    switch v5 {
                    case "hello":
                        v35 := GenS1{F0: ((^(v3)) > 2147483647)}
                        _ = v35
                    case "foo bar":
                        var v36 uint = 255
                        var v37 uint32 = v31
                        _ = v36
                        _ = v37
                    default:
                        if (!(128 < 2)) {
                            var v38 float32 = (3.14159 / 0.1)
                            _ = v38
                        } else {
                            if true {
                                var v39 int = 128
                                fmt.Println(fi29)
                                _ = v39
                            }
                            if false {
                                v40 := GenS3{F0: v4, F1: false, F2: 1000}
                                var v41 uint = 0
                                _ = v40
                                _ = v41
                            } else {
                                v3 -= v3
                                v1 = v4
                            }
                            if (1 < 32767) {
                                if (32767 < 1000) {
                                    fmt.Println(v5)
                                    v42 := GenS2{F0: (v3 * 0), F1: -710}
                                    _ = v42
                                } else {
                                    v43 := []int{fi29}
                                    var v44 uint16 = (v4 << ((2) & 63))
                                    v31 += (v31 / ((v31) | 1))
                                    _ = v43
                                    _ = v44
                                }
                                var v45 bool = ((v3 & 2147483647) <= v3)
                                if v45 {
                                    if (fi29 <= fi29) {
                                        var v46 string = v5
                                        fmt.Println(v5)
                                        var v47 string = ("a" + "\n\t")
                                        _ = v46
                                        _ = v47
                                    }
                                    fmt.Println(v4)
                                } else {
                                    var v48 float32 = float32(1e10)
                                    fmt.Println(v5)
                                    _ = v48
                                }
                                v49 := GenS4{F0: v5, F1: 2, F2: (1 > -1)}
                                _ = v45
                                _ = v49
                            } else {
                                if true {
                                    var v50 string = "a"
                                    var v51 float32 = 1e10
                                    _ = v50
                                    _ = v51
                                }
                            }
                            if false {
                                if false {
                                    var v52 float64 = v7
                                    v53 := GenS5{F0: 127, F1: 100, F2: 1}
                                    v52 -= v7
                                    _ = v52
                                    _ = v53
                                }
                                v54 := []int{(fi29 - fi30), (^(fi30)), 0}
                                v55 := GenS2{F0: (v3 ^ v3), F1: 32767}
                                if (fi29 > fi30) {
                                    v56 := GenS1{F0: (true || ((51 < 42) || ((v7 * v7) <= 0.0)))}
                                    v7 -= 0.0
                                    fmt.Println(v4, v7)
                                    _ = v56
                                } else {
                                    if len(v54) > 0 {
                                        v54[0] = 42
                                    }
                                    var v57 int8 = 42
                                    v31 -= (v31 & v31)
                                    _ = v57
                                }
                                _ = v54
                                _ = v55
                            } else {
                                if false {
                                    var v58 int64 = 9223372036854775807
                                    var v59 uint64 = 1
                                    var v60 string = ((("a" + (v5 + v5)) + v5) + v5)
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                } else {
                                    var v61 uint16 = 809
                                    v5 = "\n\t"
                                    _ = v61
                                }
                            }
                            var v62 int8 = 100
                            switch v5 {
                            case "hello":
                                if true {
                                    var v63 uint64 = 712
                                    v64 := GenS6{F0: v4}
                                    var v65 uint16 = v4
                                    if ((true || ((v62 ^ v62) != int8(v4))) && true) {
                                        switch int(v4) {
                                        case 128:
                                            var v66 int16 = 42
                                            fmt.Println(v31, v5, v1)
                                            v1 -= 2
                                            _ = v66
                                        case 127:
                                            fmt.Println(v4, v62, fi29)
                                        case 581:
                                            v3 = 2147483647
                                        default:
                                            var v67 uint = 2
                                            _ = v67
                                        }
                                    }
                                    v68 := []int{(-(fi30)), -230, fi30}
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                    _ = v68
                                } else {
                                    v2[((((v5 + ("\n\t" + "\n\t")) + v5) + "a") + v5)] = fi29
                                    fmt.Println(fi29)
                                    switch fi29 {
                                    case 365:
                                        switch (fi29 / ((fi29) | 1)) {
                                        case 1:
                                            v69 := GenS7{F0: true, F1: 211}
                                            _ = v69
                                        }
                                        v70 := GenS5{F0: (v62 - 100), F1: (^(v62)), F2: uint64(v1)}
                                        _ = v70
                                    case 9223372036854775807:
                                        if true {
                                            fmt.Println(v3)
                                            delete(v2, "\n\t")
                                            v71 := GenS8{F0: 32767, F1: -1, F2: 0}
                                            _ = v71
                                        }
                                    case 4294967295:
                                        if ((fi30 & (fi30 & int(v3))) > fi29) {
                                            v72 := v2["a"]
                                            var v73 int8 = v62
                                            v73 += v73
                                            _ = v72
                                            _ = v73
                                        } else {
                                            v74 := GenS9{F0: -507}
                                            if (42 == fi30) {
                                                switch int(v3) {
                                                case 2147483647:
                                                    if true {
                                                        switch v5 {
                                                        case "foo bar":
                                                            var v75 int8 = v62
                                                            fmt.Println(v5, v1, v5)
                                                            _ = v75
                                                        case "hello":
                                                            switch v5 {
                                                            case "foo bar":
                                                                v1 = 42
                                                            case "hello":
                                                                var v76 float64 = (v7 * v7)
                                                                _ = v76
                                                            }
                                                            v77 := GenS7{F0: (!(9223372036854775807 <= -1)), F1: 82}
                                                            v62 -= -37
                                                            _ = v77
                                                        case "a":
                                                            var v78 uint8 = 142
                                                            v79 := []int{fi29}
                                                            _ = v78
                                                            _ = v79
                                                        default:
                                                            v80 := GenS10{F0: (-1 != 1)}
                                                            _ = v80
                                                        }
                                                        v81 := []int{-2147483648, 0, (fi29 % ((fi30) | 1))}
                                                        v82 := GenS11{F0: -9223372036854775808}
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                default:
                                                    v7 = v7
                                                }
                                                v1 -= 0
                                            }
                                            _ = v74
                                        }
                                    default:
                                        var v83 bool = (42 < 18446744073709551615)
                                        _ = v83
                                    }
                                }
                            case "\n\t":
                                var v84 uint = 18446744073709551615
                                _ = v84
                            case "a":
                                switch ("" + "hello") {
                                case "a":
                                    var v85 int32 = (v3 | v3)
                                    _ = v85
                                case "hello":
                                    if false {
                                        if true {
                                            v86 := GenS12{F0: genFn0(v62, v3)}
                                            switch v5 {
                                            case "hello":
                                                if (false || false) {
                                                    v31 = (v31 >> ((18446744073709551615) & 63))
                                                    v2[v5] = (fi29 % ((128) | 1))
                                                }
                                            case "foo bar":
                                                var v87 uint = 722
                                                _ = v87
                                            }
                                            _ = v86
                                        }
                                        switch (v5 + "hello") {
                                        case "foo bar":
                                            v88 := GenS7{F0: true, F1: 254}
                                            _ = v88
                                        case "\n\t":
                                            switch fi30 {
                                            case 157:
                                                v89 := GenS11{F0: int64(v31)}
                                                _ = v89
                                            case 42:
                                                fmt.Println(v62, v4, v5)
                                            default:
                                                fmt.Println(fi30)
                                            }
                                        case "hello":
                                            fmt.Println(v62)
                                            if ((v3 % ((v3) | 1)) < 0) {
                                                v2[v5] = fi30
                                                switch fi30 {
                                                case 128:
                                                    var v90 uint32 = v31
                                                    _ = v90
                                                case 4294967295:
                                                    v31 += (v31 + 4294967295)
                                                case 1:
                                                    switch (("foo bar" + (v5 + (v5 + ((v5 + ((v5 + ("\n\t" + ("a" + v5))) + "a")) + v5)))) + v5) {
                                                    case "hello":
                                                        switch (v5 + "") {
                                                        case "a":
                                                            var v91 uint16 = (v1 << ((684) & 63))
                                                            _ = v91
                                                        default:
                                                            var v92 uint8 = 1
                                                            switch (-(fi30)) {
                                                            case 9223372036854775807:
                                                                if false {
                                                                    if (true && (!true)) {
                                                                        v93 := GenS13{F0: 1}
                                                                        fmt.Println(v4, v3, v4)
                                                                        _ = v93
                                                                    }
                                                                    switch v5 {
                                                                    case "hello":
                                                                        var v94 float64 = float64(0.0)
                                                                        _ = v94
                                                                    }
                                                                    if (!true) {
                                                                        if false {
                                                                            v62 = (v62 * (v62 * v62))
                                                                        }
                                                                    } else {
                                                                        if (!true) {
                                                                            if ((fi29 - (fi29 >> ((65535) & 63))) == (fi29 >> ((1) & 63))) {
                                                                                switch fi30 {
                                                                                case -2147483648:
                                                                                    v95 := GenS14{F0: (v92 - v92)}
                                                                                    _ = v95
                                                                                }
                                                                                var v96 int64 = -122
                                                                                if (v5 != v5) {
                                                                                    var v97 int8 = (-(v62))
                                                                                    _ = v97
                                                                                } else {
                                                                                    var v98 string = (v5 + "foo bar")
                                                                                    _ = v98
                                                                                }
                                                                                _ = v96
                                                                            } else {
                                                                                v99 := GenS3{F0: v1, F1: (!false), F2: 42}
                                                                                _ = v99
                                                                            }
                                                                            fmt.Println(fi30, v5, v7)
                                                                            if (int(v62) >= (fi30 >> ((18446744073709551615) & 63))) {
                                                                                switch (fi30 * 1) {
                                                                                case 32767:
                                                                                    fmt.Println(fi30)
                                                                                case 1:
                                                                                    var v100 int = (fi29 << ((1) & 63))
                                                                                    v92 = 1
                                                                                    v1 -= (v4 >> ((250) & 63))
                                                                                    _ = v100
                                                                                case -993:
                                                                                    if true {
                                                                                        fmt.Println(v92, v92)
                                                                                        var v101 uint64 = 171
                                                                                        _ = v101
                                                                                    }
                                                                                default:
                                                                                    v102 := GenS15{F0: 42, F1: v3}
                                                                                    _ = v102
                                                                                }
                                                                                var v103 string = v5
                                                                                _ = v103
                                                                            } else {
                                                                                _ = len(v2)
                                                                                switch v5 {
                                                                                case "":
                                                                                    fmt.Println(v62)
                                                                                case "hello":
                                                                                    v104 := v2[v5]
                                                                                    _ = v104
                                                                                default:
                                                                                    var v105 float32 = (((3.14159 / 0.1) + 0.1) + 1e10)
                                                                                    _ = v105
                                                                                }
                                                                                v106 := GenS1{F0: false}
                                                                                _ = v106
                                                                            }
                                                                        }
                                                                        if (true || false) {
                                                                            v107 := GenS16{F0: 42, F1: (v31 | 0)}
                                                                            _ = v107
                                                                        } else {
                                                                            v108 := GenS14{F0: 0}
                                                                            fmt.Println(v62, v31, fi29)
                                                                            v4 = v1
                                                                            v3 -= (v3 + (-(v3)))
                                                                            _ = v108
                                                                        }
                                                                        v109 := GenS17{F0: fi30}
                                                                        _ = v109
                                                                    }
                                                                }
                                                            case 1:
                                                                var v110 int32 = 460
                                                                if true {
                                                                    v4 += v4
                                                                } else {
                                                                    v7 += float64((float32(v7) - float32(((0.0 + v7) * v7))))
                                                                    if (910 != 42) {
                                                                        if (!true) {
                                                                            var v111 int = fi29
                                                                            _ = v111
                                                                        } else {
                                                                            fmt.Println(v3, v3)
                                                                            var v112 float64 = 0.0
                                                                            v1 = v4
                                                                            _ = v112
                                                                        }
                                                                        switch (fi29 - 653) {
                                                                        case -9223372036854775808:
                                                                            v92 = v92
                                                                        case -471:
                                                                            v113 := v2[v5]
                                                                            _ = v113
                                                                        default:
                                                                            var v114 uint = 65535
                                                                            fmt.Println(v3)
                                                                            _ = v114
                                                                        }
                                                                        var v115 float32 = 1e10
                                                                        _ = v115
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            var v116 bool = true
                                                                            _ = v116
                                                                        default:
                                                                            v117 := []int{fi29, (fi30 * fi30), fi29, fi30}
                                                                            _ = v117
                                                                        }
                                                                    }
                                                                    if true {
                                                                        if (false && (((("" + "") + v5) + "") >= v5)) {
                                                                            var v118 uint64 = 18446744073709551615
                                                                            _ = v118
                                                                        } else {
                                                                            fmt.Println(v110)
                                                                        }
                                                                        fmt.Println(v4)
                                                                    }
                                                                }
                                                                _ = v110
                                                            }
                                                            switch fi30 {
                                                            case -1:
                                                                v31 -= (v31 + (v31 >> ((1) & 63)))
                                                            case -2147483648:
                                                                v119 := []int{}
                                                                _ = v119
                                                            case 219:
                                                                var v120 bool = (true && (false || (!(127 <= -126))))
                                                                _ = v120
                                                            }
                                                            _ = v92
                                                        }
                                                    }
                                                default:
                                                    if (!(!true)) {
                                                        fmt.Println(fi29, v62, fi30)
                                                    } else {
                                                        var v121 int32 = v3
                                                        v122 := GenS3{F0: (v1 * 0), F1: ((^(v31)) == (v31 - v31)), F2: -32768}
                                                        var v123 float32 = -2.25
                                                        v124 := GenS13{F0: 530}
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                        _ = v124
                                                    }
                                                    if (1e10 <= v7) {
                                                        v62 = (v62 % ((101) | 1))
                                                    }
                                                }
                                            }
                                            if false {
                                                v125 := []int{(-(fi30)), 4294967295, (^(fi30)), int(v62)}
                                                v2["foo bar"] = 2147483647
                                                _ = v125
                                            }
                                        }
                                    }
                                case "foo bar":
                                    if false {
                                        v62 = 0
                                    }
                                default:
                                    v1 = v4
                                }
                            default:
                                if (true && false) {
                                    var v126 string = ("foo bar" + (("" + v5) + v5))
                                    v127 := []int{(fi29 >> ((578) & 63))}
                                    var v128 uint32 = (v31 * v31)
                                    var v129 uint = 213
                                    _ = v126
                                    _ = v127
                                    _ = v128
                                    _ = v129
                                } else {
                                    var v130 bool = (!false)
                                    _ = v130
                                }
                            }
                            _ = v62
                        }
                    }
                }
                _ = fi30
                _ = v31
            }
        case 128:
            v131 := GenS18{F0: 2147483648, F1: 302}
            _ = v131
        }
        _ = fi29
    }
    var v132 uint16 = v1
    v133 := []int{9223372036854775807, 127, -9223372036854775808}
    fmt.Println(v4, v5, v132)
    var v134 float64 = 1.5
    for fi135 := 0; fi135 < 3; fi135++ {
        for fi136 := 0; fi136 < 4; fi136++ {
            if false {
                v7 = ((v7 / (-2.25 + v7)) + 1.5)
                var v137 int64 = -9223372036854775808
                _ = v137
            } else {
                var v138 int = fi135
                fmt.Println(v5, fi136, v134)
                switch (v5 + v5) {
                case "":
                    v4 -= (v132 / ((769) | 1))
                default:
                    v7 = (((v7 - ((float64(1.5) * v7) * v134)) / (v134 + float64(0.0))) * v7)
                }
                _ = v138
            }
            fmt.Println(fi136, v4, fi136)
            var v139 uint16 = v1
            _ = fi136
            _ = v139
        }
        _ = fi135
    }
    for fi140 := 0; fi140 < 20; fi140++ {
        _ = len(v2)
        if len(v133) >= 2 {
            v133 = v133[1:len(v133)]
        }
        _ = fi140
    }
    v133 = append(v133, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("slicelen v133:", len(v133), v133)
    fmt.Println("v134:", v134)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v132
    _ = v133
    _ = v134
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
