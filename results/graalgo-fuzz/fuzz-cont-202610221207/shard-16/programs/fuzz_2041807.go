package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 uint }
type GenS9 struct { F0 int; F1 uint16; F2 int64 }
type GenS10 struct { F0 uint8; F1 int64; F2 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint64; F1 int; F2 uint64 }
type GenS13 struct { F0 int8; F1 int16; F2 int8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int32; F1 int }
type GenS16 struct { F0 int16; F1 uint16 }
type GenS17 struct { F0 bool; F1 uint }
type GenS18 struct { F0 uint; F1 uint32; F2 string }
type GenS19 struct { F0 string; F1 int32; F2 uint8 }
func genFn0(p0 int32) int16 {
    v6 := GenS1{F0: (("foo bar" + (("hello" + "") + ("a" + ""))) + ("\n\t" + ("a" + "")))}
    if (32767 < -32768) {
        v7 := GenS1{F0: ""}
        p0 = p0
        _ = v7
    } else {
        if false {
            var v8 bool = true
            _ = v8
        } else {
            var v9 uint = 255
            if true {
                var v10 uint64 = 2
                var v11 int64 = 840
                v10 -= v10
                _ = v10
                _ = v11
            }
            fmt.Println(v9, p0)
            _ = v9
        }
        fmt.Println(p0)
    }
    var v12 int8 = 42
    var v13 int16 = 42
    _ = p0
    _ = v12
    _ = v13
    _ = v6
    return v13
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 1
    var v3 bool = (false || (41 == -32768))
    var v4 float64 = -2.25
    v5 := []int{}
    var v14 uint = v2
    delete(v1, 42)
    v15 := GenS1{F0: "a"}
    if true {
        var v16 int = -881
        switch "hello" {
        case "foo bar":
            if (((float64((3.14159 / (0.1 * float32(float64(((-2.25 - (1e10 - (float32(v4) + 0.1))) / 3.14159)))))) / v4) + -2.25) > ((((v4 + -2.25) / v4) / v4) / ((v4 + v4) / 3.14159))) {
                v17 := GenS2{F0: 1, F1: -32768, F2: ("" + "")}
                v18 := GenS3{F0: v3}
                _ = v17
                _ = v18
            } else {
                v3 = v3
                v19 := v1[v16]
                v20 := v1[(v16 * (-(v16)))]
                _ = v19
                _ = v20
            }
        case "\n\t":
            v21 := GenS4{F0: 0, F1: 219, F2: 128}
            _ = v21
        case "a":
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 uint8 = 128
                _ = fi22
                _ = v23
            }
        }
        _ = v16
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        if (845 != v2) {
            v25 := GenS5{F0: 0, F1: -128}
            _ = v25
        }
        _ = fi24
    }
    var v26 float64 = (((v4 / v4) - 3.14159) / v4)
    var v27 uint = v2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v3 {
        if v3 {
            _ = len(v1)
            for fi28 := 0; fi28 < 18; fi28++ {
                for fi29 := 0; fi29 < 15; fi29++ {
                    _ = len(v5)
                    v3 = v3
                    _ = fi29
                }
                v4 += v4
                _ = fi28
            }
        }
    } else {
        if (v3 || (!(!(127 != 100)))) {
            v30 := GenS4{F0: -1, F1: 42, F2: 128}
            v14 -= v14
            if v3 {
                v5 = append(v5, 373)
                v31 := GenS6{F0: (9223372036854775807 != 9223372036854775807)}
                _ = v31
            }
            _ = v30
        }
    }
    v26 = float64(1e10)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = 0.1
    var v32 float32 = 1e10
    v26 = float64(1e10)
    v33 := GenS3{F0: v3}
    if (!(1000 >= 42)) {
        if (((114 >= 128) || (((!(false && true)) && true) || false)) || ("a" >= ((("" + "a") + ("foo bar" + "foo bar")) + ("" + (("hello" + "hello") + ""))))) {
            if false {
                v34 := GenS1{F0: ""}
                if len(v5) > 0 {
                    v5[0] = 788
                }
                _ = v34
            }
            if false {
                for fi35 := 0; fi35 < 18; fi35++ {
                    var v36 int64 = 1
                    v37 := GenS4{F0: 722, F1: 42, F2: 127}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                v38 := GenS4{F0: 0, F1: 2147483648, F2: 0}
                for fi39 := 0; fi39 < 7; fi39++ {
                    var v40 int32 = -437
                    v3 = (true || true)
                    _ = fi39
                    _ = v40
                }
                _ = v38
            } else {
                if (-466 == 0) {
                    if (!(1 < 43)) {
                        var v41 float32 = 1e10
                        fmt.Println(v3)
                        fmt.Println(v27, v4, v3)
                        _ = v41
                    } else {
                        if true {
                            for fi42 := 0; fi42 < 8; fi42++ {
                                v27 += (v2 % (((v14 << ((208) & 63))) | 1))
                                var v43 string = "a"
                                if v3 {
                                    v3 = v3
                                    switch "" {
                                    case "hello":
                                        v26 -= v4
                                    case "foo bar":
                                        _ = len(v1)
                                    default:
                                        v2 += (v14 & 0)
                                    }
                                    v44 := GenS7{F0: 100}
                                    _ = v44
                                }
                                _ = fi42
                                _ = v43
                            }
                            for fi45 := 0; fi45 < 7; fi45++ {
                                v46 := GenS8{F0: 1, F1: (v27 / ((v2) | 1))}
                                v47 := GenS9{F0: (fi45 | -2147483648), F1: 672, F2: -9223372036854775808}
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                            v32 = (v32 * ((((v32 + v32) + (v32 + ((v32 - v32) + v32))) / 1.0) * v32))
                        } else {
                            v2 -= 42
                            v5 = append(v5, 328)
                        }
                        v14 = 65535
                    }
                    v48 := GenS1{F0: ("" + ((("\n\t" + "\n\t") + ("hello" + "a")) + "foo bar"))}
                    for fi49 := 0; fi49 < 5; fi49++ {
                        switch "foo bar" {
                        case "hello":
                            for fi50 := 0; fi50 < 18; fi50++ {
                                v27 = v2
                                if false {
                                    switch (fi50 >> (((v2 + (^(v27)))) & 63)) {
                                    case 1:
                                        var v51 int32 = 42
                                        _ = v51
                                    }
                                } else {
                                    v52 := v1[(fi50 << (((v2 * v27)) & 63))]
                                    fmt.Println(v26, v4, fi49)
                                    switch fi50 {
                                    case -929:
                                        v5 = append(v5, -972)
                                    case 84:
                                        switch (fi49 / (((fi50 << ((v14) & 63))) | 1)) {
                                        case 128:
                                            var v53 uint8 = 127
                                            _ = v53
                                        default:
                                            v3 = v3
                                        }
                                    default:
                                        switch (fi50 << ((v27) & 63)) {
                                        case 0:
                                            v5 = append(v5, fi50)
                                        }
                                    }
                                    _ = v52
                                }
                                v14 = v27
                                switch "a" {
                                case "a":
                                    if false {
                                        var v54 int32 = -2147483648
                                        _ = v54
                                    } else {
                                        if ((0 != 127) && (18446744073709551615 < 0)) {
                                            v32 -= (float32(-2.25) / v32)
                                            v55 := GenS10{F0: 136, F1: 9223372036854775807, F2: (!(false || (true && false)))}
                                            _ = v55
                                        } else {
                                            var v56 string = ("" + ("hello" + "\n\t"))
                                            v26 -= 0.0
                                            if false {
                                                if v3 {
                                                    if len(v5) > 0 {
                                                        v5[0] = fi50
                                                    }
                                                    if v3 {
                                                        if (true && false) {
                                                            var v57 string = "\n\t"
                                                            v56 = v57
                                                            _ = v57
                                                        } else {
                                                            var v58 uint16 = 0
                                                            _ = v58
                                                        }
                                                        v27 = (v14 & (v14 ^ (v14 & 2)))
                                                        switch "" {
                                                        case "":
                                                            v59 := GenS3{F0: false}
                                                            _ = v59
                                                        case "hello":
                                                            var v60 uint64 = 326
                                                            _ = v60
                                                        default:
                                                            if (((float32(v4) / 0.1) - v32) < v32) {
                                                                v1[(fi50 & 2147483647)] = ((("foo bar" + (v56 + v56)) + v56) + (("a" + ((("\n\t" + ("\n\t" + ((((v56 + ("a" + ("\n\t" + v56))) + (("" + v56) + "a")) + v56) + "foo bar"))) + (("foo bar" + (v56 + v56)) + v56)) + (("hello" + "a") + v56))) + v56))
                                                                if v3 {
                                                                    var v61 int = fi49
                                                                    var v62 uint16 = 2
                                                                    switch v56 {
                                                                    case "hello":
                                                                        if len(v5) > 0 {
                                                                            v5[0] = fi49
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(v3)
                                                                    case "foo bar":
                                                                        if (v14 > (v2 | (v27 % (((^(v27))) | 1)))) {
                                                                            var v63 uint64 = 0
                                                                            _ = v63
                                                                        } else {
                                                                            fmt.Println(v27, v62)
                                                                        }
                                                                    }
                                                                    v64 := GenS11{F0: v2}
                                                                    if len(v5) > 0 {
                                                                        v5[0] = fi50
                                                                    }
                                                                    _ = v61
                                                                    _ = v62
                                                                    _ = v64
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(fi50)
                                                        v2 -= (^(v27))
                                                    }
                                                } else {
                                                    v2 -= 18446744073709551615
                                                    var v65 uint16 = 0
                                                    if (!((fi49 > 127) || v3)) {
                                                        switch fi49 {
                                                        case 0:
                                                            if v3 {
                                                                v27 = (v2 << ((255) & 63))
                                                            }
                                                        default:
                                                            v2 = v27
                                                        }
                                                        var v66 string = "hello"
                                                        if true {
                                                            v5 = append(v5, fi49)
                                                            delete(v1, 32768)
                                                            v67 := v1[fi50]
                                                            v56 = "foo bar"
                                                            _ = v67
                                                        } else {
                                                            if (-1 >= -142) {
                                                                _ = len(v1)
                                                                var v68 uint8 = 128
                                                                fmt.Println(v56, v65)
                                                                if (v3 && v3) {
                                                                    if v3 {
                                                                        var v69 bool = ((v56 + v56) < "\n\t")
                                                                        v66 = v66
                                                                        if ((42 <= 18446744073709551615) && v69) {
                                                                            var v70 uint8 = (^(v68))
                                                                            v70 = (v68 >> (((v27 - v2)) & 63))
                                                                            if (-47 <= 63) {
                                                                                v66 = v66
                                                                            } else {
                                                                                _ = len(v5)
                                                                                v71 := GenS12{F0: uint64(v65), F1: fi49, F2: 4294967295}
                                                                                _ = v71
                                                                            }
                                                                            var v72 float32 = v32
                                                                            switch (v66 + (v66 + v66)) {
                                                                            case "":
                                                                                var v73 int64 = 0
                                                                                _ = v73
                                                                            case "foo bar":
                                                                                v74 := GenS13{F0: -1, F1: 1000, F2: 100}
                                                                                _ = v74
                                                                            case "hello":
                                                                                v66 = "hello"
                                                                            default:
                                                                                v75 := GenS14{F0: uint(v68)}
                                                                                _ = v75
                                                                            }
                                                                            _ = v70
                                                                            _ = v72
                                                                        }
                                                                        if v69 {
                                                                            var v76 uint8 = v68
                                                                            var v77 string = v66
                                                                            _ = v76
                                                                            _ = v77
                                                                        } else {
                                                                            var v78 float64 = v4
                                                                            v79 := GenS12{F0: 2, F1: fi50, F2: 433}
                                                                            _ = v78
                                                                            _ = v79
                                                                        }
                                                                        v56 = v66
                                                                        v80 := GenS15{F0: -1, F1: fi50}
                                                                        _ = v69
                                                                        _ = v80
                                                                    } else {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        if (((v32 < -2.25) && (662 >= 0)) || (true || v3)) {
                                                                            switch v66 {
                                                                            case "hello":
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    var v81 bool = v3
                                                                                    _ = v81
                                                                                default:
                                                                                    if v3 {
                                                                                        var v82 uint8 = 39
                                                                                        _ = v82
                                                                                    }
                                                                                    v83 := GenS16{F0: 1, F1: v65}
                                                                                    _ = v83
                                                                                }
                                                                            case "\n\t":
                                                                                fmt.Println(v4)
                                                                            default:
                                                                                var v84 uint32 = 32
                                                                                _ = v84
                                                                            }
                                                                        } else {
                                                                            switch fi50 {
                                                                            case 128:
                                                                                v66 = v56
                                                                            }
                                                                            v85 := v1[4294967295]
                                                                            _ = v85
                                                                        }
                                                                        var v86 int8 = int8(v65)
                                                                        _ = v86
                                                                    }
                                                                }
                                                                _ = v68
                                                            }
                                                            var v87 uint = v27
                                                            fmt.Println(v56, v56)
                                                            _ = v87
                                                        }
                                                        v88 := v1[(fi50 << ((v27) & 63))]
                                                        v66 = v56
                                                        var v89 uint16 = (v65 + 65535)
                                                        if v3 {
                                                            v90 := GenS9{F0: (fi50 >> (((v2 | (^(v14)))) & 63)), F1: 260, F2: -1}
                                                            v91 := GenS1{F0: (v88 + v88)}
                                                            v66 = v56
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        _ = v66
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                    v56 = v56
                                                    var v92 uint = v2
                                                    _ = v65
                                                    _ = v92
                                                }
                                                var v93 uint16 = 77
                                                var v94 uint8 = 2
                                                _ = v93
                                                _ = v94
                                            }
                                            switch fi49 {
                                            case 128:
                                                if (738 > 2147483648) {
                                                    v4 += v4
                                                    var v95 float64 = 1e10
                                                    var v96 int64 = -94
                                                    if len(v5) > 0 {
                                                        v5[0] = fi49
                                                    }
                                                    _ = v95
                                                    _ = v96
                                                } else {
                                                    v26 += (0.1 - 0.0)
                                                    var v97 int16 = -539
                                                    var v98 uint8 = 149
                                                    var v99 int8 = 0
                                                    _ = v97
                                                    _ = v98
                                                    _ = v99
                                                }
                                            case 32767:
                                                v56 = v56
                                                v27 = 1
                                                var v100 uint16 = 166
                                                _ = v100
                                            default:
                                                delete(v1, (fi50 - 32767))
                                            }
                                            _ = v56
                                        }
                                        _ = len(v5)
                                        v101 := GenS17{F0: ((!true) && (9223372036854775807 > 275)), F1: (v2 + v2)}
                                        _ = v101
                                    }
                                case "hello":
                                    v4 = float64(1.5)
                                default:
                                    v102 := GenS18{F0: v14, F1: 353, F2: "a"}
                                    _ = v102
                                }
                                _ = fi50
                            }
                        case "\n\t":
                            if ((v2 % (((v14 ^ (^(v14)))) | 1)) == 0) {
                                for fi103 := 0; fi103 < 17; fi103++ {
                                    fmt.Println(v3, v3, v4)
                                    var v104 int32 = 42
                                    _ = fi103
                                    _ = v104
                                }
                                var v105 int8 = 100
                                var v106 int32 = 1
                                _ = v105
                                _ = v106
                            } else {
                                var v107 bool = (0 == 42)
                                _ = v107
                            }
                        case "":
                            if (0 <= 129) {
                                for fi108 := 0; fi108 < 4; fi108++ {
                                    v4 += 1e10
                                    v109 := GenS19{F0: "", F1: -605, F2: 87}
                                    fmt.Println(v26, fi108)
                                    var v110 int16 = 32767
                                    _ = fi108
                                    _ = v109
                                    _ = v110
                                }
                                v111 := GenS18{F0: v27, F1: 185, F2: ""}
                                _ = v111
                            } else {
                                v32 -= 1.5
                                var v112 uint = 42
                                _ = v112
                            }
                        default:
                            if (18446744073709551615 < 2) {
                                fmt.Println(v26)
                                switch fi49 {
                                case 32767:
                                    var v113 uint16 = 65535
                                    _ = v113
                                case 746:
                                    v5 = append(v5, -807)
                                case 911:
                                    if (!false) {
                                        v3 = ((!v3) || v3)
                                    } else {
                                        v26 -= 3.14159
                                    }
                                default:
                                    v32 += (v32 + float32((0.0 - v4)))
                                }
                                if false {
                                    v27 = v14
                                    v26 = 1e10
                                } else {
                                    v114 := GenS11{F0: 65535}
                                    switch "a" {
                                    case "":
                                        for fi115 := 0; fi115 < 12; fi115++ {
                                            var v116 int8 = 127
                                            fmt.Println(fi115, v26)
                                            _ = fi115
                                            _ = v116
                                        }
                                    case "a":
                                        var v117 uint16 = 255
                                        _ = v117
                                    default:
                                        var v118 uint32 = 4294967295
                                        _ = v118
                                    }
                                    _ = v114
                                }
                            }
                        }
                        _ = fi49
                    }
                    _ = v48
                } else {
                    for fi119 := 0; fi119 < 18; fi119++ {
                        v3 = (!(42 != 1))
                        for fi120 := 0; fi120 < 2; fi120++ {
                            v26 = 3.14159
                            var v121 int32 = 42
                            _ = fi120
                            _ = v121
                        }
                        _ = fi119
                    }
                    if v3 {
                        v3 = true
                        v3 = (-1 >= 100)
                        v122 := v1[964]
                        _ = v122
                    }
                }
                fmt.Println(v26, v4)
            }
            var v123 uint32 = 1
            _ = v123
        }
    } else {
        v2 = v14
        for fi124 := 0; fi124 < 19; fi124++ {
            v32 = (v32 / 3.14159)
            if (!true) {
                if (v3 || (736 == -9223372036854775808)) {
                    v125 := v1[615]
                    v4 = float64(v32)
                    _ = v125
                }
                for fi126 := 0; fi126 < 14; fi126++ {
                    _ = len(v1)
                    _ = fi126
                }
            }
            var v127 int64 = -359
            _ = fi124
            _ = v127
        }
    }
    if ("\n\t" > ((("hello" + "foo bar") + "\n\t") + ("a" + ""))) {
        v2 -= v2
        v2 = (v14 | 984)
        for fi128 := 0; fi128 < 12; fi128++ {
            if true {
                v27 = v14
                var v129 float32 = v32
                _ = v129
            }
            var v130 int64 = 954
            v5 = append(v5, (fi128 ^ fi128))
            _ = fi128
            _ = v130
        }
    } else {
        v27 -= (v2 + (v27 ^ v14))
        var v131 uint16 = 776
        _ = v131
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
