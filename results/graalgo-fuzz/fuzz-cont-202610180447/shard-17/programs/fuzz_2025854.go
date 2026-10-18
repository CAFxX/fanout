package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 int64; F1 uint16; F2 bool }
type GenS9 struct { F0 uint32; F1 string }
func genFn0() int64 {
    v6 := []int{-2147483648, 2147483647, 4294967295}
    var v7 int16 = 32767
    v8 := GenS1{F0: 42, F1: 1, F2: -18}
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int32 = 1
    v2 := []string{"a", "a", "foo bar", ("a" + ("" + ""))}
    var v3 string = ""
    var v4 int8 = 100
    var v5 float64 = -2.25
    v1 -= -1
    var v9 int8 = 1
    var v10 int = 128
    for fi11 := 0; fi11 < 11; fi11++ {
        v9 = 127
        _ = fi11
    }
    switch v10 {
    case 954:
        for fi12 := 0; fi12 < 20; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                v9 = v4
                fmt.Println(v3, fi13, v9)
                fmt.Println(v9, v3, v4)
                v14 := GenS1{F0: 2, F1: (v1 + -1), F2: 17}
                _ = fi13
                _ = v14
            }
            fmt.Println(v9)
            var v15 int16 = 1000
            _ = fi12
            _ = v15
        }
    case 4294967295:
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 string = v3
            for fi18 := 0; fi18 < 19; fi18++ {
                v10 = (^(v10))
                _ = fi18
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi16
            _ = v17
        }
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            v9 -= 100
            _ = fi19
        }
    }
    _ = len(v2)
    v20 := GenS1{F0: 128, F1: 42, F2: (v9 / ((v9) | 1))}
    if (!true) {
        for fi21 := 0; fi21 < 14; fi21++ {
            fmt.Println(v4, v9)
            switch -170 {
            case 128:
                _ = len(v2)
            case 127:
                switch "hello" {
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case 1:
                v22 := GenS1{F0: 1, F1: v1, F2: v4}
                _ = v22
            }
            var v23 int = fi21
            _ = fi21
            _ = v23
        }
        var v24 int = (^(v10))
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        _ = len(v2)
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        var v26 uint64 = 42
        _ = fi25
        _ = v26
    }
    if (0 < 2) {
        if false {
            var v27 int64 = 854
            for fi28 := 0; fi28 < 11; fi28++ {
                for fi29 := 0; fi29 < 14; fi29++ {
                    fmt.Println(v4, v5, v27)
                    v30 := GenS1{F0: 141, F1: v1, F2: (v9 | v9)}
                    if (("" + v3) < ((("a" + (v3 + ((v3 + v3) + v3))) + "\n\t") + (("\n\t" + ((v3 + "hello") + (((v3 + "\n\t") + "a") + (v3 + v3)))) + (v3 + (v3 + v3))))) {
                        if true {
                            var v31 uint8 = 128
                            var v32 uint16 = 0
                            if ((-(v1)) >= v1) {
                                fmt.Println(fi28)
                                var v33 float32 = 3.14159
                                var v34 uint32 = 2
                                _ = v33
                                _ = v34
                            }
                            if ((-(v4)) < v9) {
                                switch (v3 + ("\n\t" + v3)) {
                                case "foo bar":
                                    var v35 string = ("a" + "hello")
                                    _ = v35
                                case "a":
                                    v1 = v1
                                case "":
                                    if (v5 <= (v5 / v5)) {
                                        var v36 float64 = v5
                                        var v37 string = v3
                                        v38 := GenS1{F0: v31, F1: (-(v1)), F2: v9}
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(v27)
                                        case "hello":
                                            if len(v2) > 0 {
                                                v2[0] = v3
                                            }
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = ((("hello" + "foo bar") + v3) + ("hello" + v3))
                                            }
                                        }
                                        if (-124 >= v4) {
                                            if true {
                                                var v39 uint = 2
                                                fmt.Println(v3, v9)
                                                var v40 int8 = -128
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = v37
                                        }
                                        var v41 uint64 = 18446744073709551615
                                        var v42 uint32 = 21
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v43 := GenS2{F0: 381}
                                        var v44 string = v3
                                        _ = v43
                                        _ = v44
                                    }
                                default:
                                    var v45 uint32 = 2147483648
                                    _ = v45
                                }
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = v3
                                }
                            }
                            _ = v31
                            _ = v32
                        }
                        v46 := GenS3{F0: 2, F1: 1, F2: ((((float64(((1.5 - -2.25) + 1e10)) / 1.5) >= 0.1) || (v5 <= (v5 + v5))) || true)}
                        switch ("" + "a") {
                        case "hello":
                            v47 := GenS4{F0: v3, F1: 127}
                            _ = v47
                        case "\n\t":
                            var v48 float64 = 0.1
                            _ = v48
                        case "":
                            switch "foo bar" {
                            case "a":
                                switch (-(v10)) {
                                case 2147483647:
                                    if true {
                                        if (2 <= 255) {
                                            v49 := GenS1{F0: 128, F1: (v1 % (((v1 ^ -1)) | 1)), F2: (v4 | -108)}
                                            _ = v49
                                        } else {
                                            switch v3 {
                                            case "a":
                                                v2 = append(v2, ((v3 + v3) + v3))
                                            case "hello":
                                                v9 -= (v4 * 65)
                                            case "\n\t":
                                                v3 = ((v3 + v3) + v3)
                                            }
                                        }
                                    } else {
                                        v50 := GenS5{F0: 1, F1: 899}
                                        v51 := GenS6{F0: false, F1: int32(v27)}
                                        _ = v50
                                        _ = v51
                                    }
                                case 0:
                                    fmt.Println(fi29, v1, v10)
                                case 622:
                                    switch 0 {
                                    case 9223372036854775807:
                                        var v52 int64 = genFn0()
                                        _ = v52
                                    }
                                default:
                                    var v53 uint16 = 65535
                                    _ = v53
                                }
                            case "":
                                if false {
                                    var v54 uint64 = 42
                                    if ((42 < 1) && (float32((((v5 + 1e10) + v5) - v5)) == 1e10)) {
                                        v54 = 42
                                        if len(v2) > 0 {
                                            v2[0] = (v3 + (v3 + "\n\t"))
                                        }
                                        if (-1 < 1) {
                                            v2 = append(v2, v3)
                                            if (4294967295 >= fi29) {
                                                if (!false) {
                                                    switch 110 {
                                                    case -156:
                                                        var v55 int8 = v4
                                                        _ = v55
                                                    case -642:
                                                        if len(v2) > 0 {
                                                            v2[0] = (((v3 + "foo bar") + v3) + v3)
                                                        }
                                                    case 257:
                                                        switch 461 {
                                                        case -162:
                                                            v56 := GenS5{F0: 255, F1: v54}
                                                            _ = v56
                                                        case 32768:
                                                            var v57 uint8 = 1
                                                            _ = v57
                                                        case 1:
                                                            v2 = append(v2, (v3 + "a"))
                                                        }
                                                        if true {
                                                            var v58 bool = true
                                                            switch ("foo bar" + v3) {
                                                            case "\n\t":
                                                                v9 -= (-(v9))
                                                            default:
                                                                switch "a" {
                                                                case "a":
                                                                    if v58 {
                                                                        var v59 uint = 0
                                                                        v60 := GenS2{F0: (v59 >> ((v59) & 63))}
                                                                        v61 := GenS7{F0: (("foo bar" + v3) + (v3 + v3)), F1: -1}
                                                                        v62 := GenS1{F0: 1, F1: (v1 / ((-459) | 1)), F2: (v4 / (((v4 - (v9 / (((v4 & v4)) | 1)))) | 1))}
                                                                        _ = v59
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v62
                                                                    } else {
                                                                        var v63 int32 = 51
                                                                        fmt.Println(v1, v5)
                                                                        var v64 int32 = 1
                                                                        var v65 float32 = ((0.1 - 1e10) - 1.5)
                                                                        _ = v63
                                                                        _ = v64
                                                                        _ = v65
                                                                    }
                                                                case "hello":
                                                                    v5 -= v5
                                                                case "\n\t":
                                                                    if v58 {
                                                                        if v58 {
                                                                            switch v10 {
                                                                            case 42:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            case 4294967295:
                                                                                v58 = true
                                                                            case -925:
                                                                                fmt.Println(v9, v10, v9)
                                                                            default:
                                                                                var v66 int8 = (v9 & int8(fi29))
                                                                                _ = v66
                                                                            }
                                                                            v67 := GenS6{F0: (((float32((v5 / (v5 - (((v5 - v5) - v5) * v5)))) - -2.25) * 3.14159) == 1.5), F1: v1}
                                                                            _ = v67
                                                                        } else {
                                                                            v68 := GenS1{F0: 44, F1: (v1 * v1), F2: v4}
                                                                            v9 -= (v9 & v9)
                                                                            _ = v68
                                                                        }
                                                                    } else {
                                                                        v58 = ((-(fi28)) <= fi29)
                                                                    }
                                                                }
                                                            }
                                                            var v69 int64 = genFn0()
                                                            _ = v58
                                                            _ = v69
                                                        }
                                                        v70 := GenS1{F0: 255, F1: v1, F2: v4}
                                                        _ = v70
                                                    default:
                                                        var v71 int16 = 0
                                                        v54 += (v54 << ((4294967295) & 63))
                                                        if true {
                                                            v2 = append(v2, "a")
                                                            v54 += v54
                                                        } else {
                                                            var v72 float32 = 1e10
                                                            if (!(!((("foo bar" + v3) + ((("a" + (v3 + "a")) + "") + (((v3 + v3) + v3) + "hello"))) < "\n\t"))) {
                                                                var v73 uint16 = 910
                                                                _ = v73
                                                            } else {
                                                                if false {
                                                                    fmt.Println(v71, v3, v1)
                                                                    if true {
                                                                        v74 := GenS1{F0: 79, F1: (v1 >> ((266) & 63)), F2: 42}
                                                                        _ = v74
                                                                    }
                                                                } else {
                                                                    v2 = append(v2, (v3 + (((v3 + (("foo bar" + "hello") + "\n\t")) + v3) + v3)))
                                                                    var v75 string = v3
                                                                    if (true || (814 < 2)) {
                                                                        var v76 string = v3
                                                                        _ = v76
                                                                    }
                                                                    _ = v75
                                                                }
                                                                if (2147483648 > 495) {
                                                                    _ = len(v2)
                                                                } else {
                                                                    v3 = v3
                                                                    v77 := GenS2{F0: 42}
                                                                    _ = v77
                                                                }
                                                                v9 = -128
                                                                v1 += (v1 * -2147483648)
                                                            }
                                                            _ = v72
                                                        }
                                                        _ = v71
                                                    }
                                                    if true {
                                                        fmt.Println(fi28, v4)
                                                        if len(v2) > 0 {
                                                            v2[0] = v3
                                                        }
                                                        if false {
                                                            switch v3 {
                                                            case "a":
                                                                fmt.Println(v3, v10, v9)
                                                            case "foo bar":
                                                                fmt.Println(fi29)
                                                            }
                                                            var v78 int = (v10 & (fi28 - 2147483647))
                                                            _ = v78
                                                        }
                                                    } else {
                                                        if true {
                                                            var v79 int32 = (v1 ^ v1)
                                                            _ = len(v2)
                                                            var v80 uint32 = 0
                                                            fmt.Println(v4)
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        var v81 int64 = v27
                                                        if (((float32((v5 / float64(float32(3.14159)))) / 1.5) <= 0.1) && ((v1 - v1) <= v1)) {
                                                            var v82 int16 = 0
                                                            _ = v82
                                                        } else {
                                                            v4 -= 100
                                                        }
                                                        fmt.Println(v5, v10, v5)
                                                        if len(v2) > 0 {
                                                            v2[0] = (v3 + ("" + v3))
                                                        }
                                                        _ = len(v2)
                                                        var v83 uint8 = 2
                                                        _ = v81
                                                        _ = v83
                                                    }
                                                } else {
                                                    var v84 uint = 142
                                                    if true {
                                                        var v85 bool = false
                                                        v86 := GenS6{F0: true, F1: (v1 | 0)}
                                                        var v87 int64 = v27
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                    } else {
                                                        if true {
                                                            fmt.Println(v10, v1)
                                                            var v88 uint8 = 42
                                                            var v89 uint16 = 255
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        v90 := GenS2{F0: v84}
                                                        var v91 uint = (v84 >> ((v84) & 63))
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    _ = v84
                                                }
                                            } else {
                                                v1 -= v1
                                                if false {
                                                    if false {
                                                        v2 = append(v2, "foo bar")
                                                        var v92 int = (fi28 | 1)
                                                        _ = v92
                                                    } else {
                                                        if (true && false) {
                                                            var v93 uint8 = 2
                                                            _ = v93
                                                        } else {
                                                            var v94 uint32 = 268
                                                            _ = v94
                                                        }
                                                        if (0 != 255) {
                                                            var v95 uint16 = 65535
                                                            v96 := GenS7{F0: (v3 + v3), F1: v4}
                                                            v95 += 213
                                                            var v97 int32 = 0
                                                            v10 += (fi28 + (-(fi28)))
                                                            v97 -= v1
                                                            _ = v95
                                                            _ = v96
                                                            _ = v97
                                                        } else {
                                                            var v98 int8 = v4
                                                            _ = v98
                                                        }
                                                    }
                                                }
                                                fmt.Println(v3, v27)
                                            }
                                            v1 = (v1 - (-(v1)))
                                        }
                                    }
                                    _ = v54
                                }
                            case "foo bar":
                                var v99 float64 = (-2.25 * float64(-2.25))
                                _ = v99
                            default:
                                v100 := GenS8{F0: v27, F1: 0, F2: ((-(v1)) == v1)}
                                _ = v100
                            }
                        default:
                            v27 += 1
                        }
                        _ = v46
                    } else {
                        v27 -= v27
                        if (!(fi29 != (^(fi28)))) {
                            fmt.Println(v9)
                        }
                        var v101 uint64 = 4294967295
                        _ = v101
                    }
                    _ = fi29
                    _ = v30
                }
                if ((1e10 / 1.5) != 1e10) {
                    for fi102 := 0; fi102 < 3; fi102++ {
                        var v103 int8 = (v9 & (v4 + v4))
                        _ = fi102
                        _ = v103
                    }
                    for fi104 := 0; fi104 < 16; fi104++ {
                        switch fi28 {
                        case -9223372036854775808:
                            if false {
                                if false {
                                    if ("foo bar" > "hello") {
                                        v105 := GenS9{F0: 4294967295, F1: v3}
                                        fmt.Println(v1, v1, v4)
                                        var v106 int32 = -129
                                        v107 := GenS5{F0: 596, F1: 18446744073709551615}
                                        _ = v105
                                        _ = v106
                                        _ = v107
                                    }
                                    if true {
                                        v27 -= (v27 << ((2) & 63))
                                        if (411 >= 71) {
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                            v9 += (v9 + v4)
                                        }
                                        var v108 string = v3
                                        _ = v108
                                    }
                                }
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = v3
                                    }
                                }
                                var v109 float32 = (1.5 / (1.5 / (1e10 / float32(v5))))
                                _ = v109
                            }
                        case 9223372036854775807:
                            v4 = (^(v4))
                        case 42:
                            var v110 int8 = -1
                            _ = v110
                        }
                        _ = fi104
                    }
                } else {
                    v2 = append(v2, "foo bar")
                    if len(v2) > 0 {
                        v2[0] = v3
                    }
                    switch 32768 {
                    case -1:
                        var v111 int16 = -792
                        _ = v111
                    default:
                        for fi112 := 0; fi112 < 15; fi112++ {
                            v1 = v1
                            _ = fi112
                        }
                    }
                }
                v27 -= (v27 - v27)
                _ = fi28
            }
            v5 = 0.0
            _ = v27
        }
    }
    switch v10 {
    case -9223372036854775808:
        if true {
            var v113 int16 = 1
            _ = v113
        } else {
            for fi114 := 0; fi114 < 16; fi114++ {
                for fi115 := 0; fi115 < 11; fi115++ {
                    switch fi114 {
                    case 432:
                        if ((v10 + fi115) <= fi114) {
                            if true {
                                _ = len(v2)
                            }
                            v4 = 100
                            var v116 string = v3
                            _ = v116
                        }
                    case -1:
                        v5 -= ((((1e10 + (((v5 / float64(3.14159)) / 1.0) - (0.1 / ((3.14159 * -2.25) * v5)))) + float64(((0.0 / (((float32((v5 / v5)) * 3.14159) * 0.1) / 1e10)) + float32(v5)))) / 1.5) - (v5 / v5))
                    default:
                        v117 := GenS2{F0: 2}
                        _ = v117
                    }
                    v3 = v3
                    _ = fi115
                }
                _ = fi114
            }
            var v118 string = (v3 + v3)
            switch (("a" + (v118 + "hello")) + v118) {
            case "a":
                v4 -= 127
            case "foo bar":
                fmt.Println(v9)
            }
            _ = v118
        }
    case 0:
        v5 += 0.1
    default:
        v10 += (v10 ^ 257)
    }
    var v119 string = ("\n\t" + v3)
    v120 := GenS2{F0: 463}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v119:", v119)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v119
    _ = v120
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
