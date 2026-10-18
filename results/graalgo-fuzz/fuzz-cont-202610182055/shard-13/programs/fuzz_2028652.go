package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 string; F1 uint16; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 string; F1 uint; F2 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 uint64; F2 int8 }
type GenS11 struct { F0 uint; F1 string; F2 uint8 }
type GenS12 struct { F0 int; F1 int64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint; F1 uint }
func genFn0(p0 uint, p1 int16, p2 int32) uint {
    fmt.Println(p0)
    v5 := GenS1{F0: -598, F1: "hello"}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 string = "foo bar"
        var v8 uint64 = 22
        v9 := []int{(fi6 ^ (fi6 % ((32768) | 1))), (fi6 % (((fi6 + -2147483648)) | 1)), 575, -1}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{611, 308, 128, 2147483647}
    if false {
        fmt.Println(p1, p0, p0)
    } else {
        p1 = 617
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return p0
}
func genFn1() int8 {
    v11 := GenS2{F0: 11}
    var v12 uint16 = 645
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 bool = (65535 == 1)
    v2 := []string{"", ("a" + "foo bar"), (("a" + ("\n\t" + ("\n\t" + ("foo bar" + "foo bar")))) + "hello")}
    v3 := make(map[string]uint32)
    v4 := []string{"hello"}
    if v1 {
        if true {
            var v13 uint = 4294967295
            _ = v13
        }
    }
    v3[(("foo bar" + (("\n\t" + "hello") + ("hello" + (("hello" + "foo bar") + "hello")))) + "a")] = 494
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if (!(799 <= -1)) {
        var v14 float64 = -2.25
        if (1 <= -1) {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint64 = 911
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                if v1 {
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        var v18 int = (fi17 ^ 128)
                        _ = v18
                    }
                    if v1 {
                        var v19 int32 = -485
                        v14 = v14
                        _ = v19
                    }
                    if v1 {
                        if true {
                            switch (fi17 << ((95) & 63)) {
                            case 127:
                                switch "hello" {
                                case "a":
                                    v14 += v14
                                case "foo bar":
                                    for fi20 := 0; fi20 < 15; fi20++ {
                                        switch "a" {
                                        case "foo bar":
                                            _ = len(v4)
                                        case "\n\t":
                                            v21 := GenS1{F0: 127, F1: ("foo bar" + "")}
                                            _ = v21
                                        default:
                                            var v22 int8 = 57
                                            _ = v22
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v23 uint = 4294967295
                                        _ = fi20
                                        _ = v23
                                    }
                                }
                            case -2147483648:
                                var v24 bool = v1
                                _ = v24
                            default:
                                if v1 {
                                    var v25 uint16 = 42
                                    fmt.Println(v25, v25)
                                    if len(v4) > 0 {
                                        v4[0] = "hello"
                                    }
                                    _ = v25
                                } else {
                                    for fi26 := 0; fi26 < 15; fi26++ {
                                        var v27 float32 = -2.25
                                        if (v14 > v14) {
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v14)
                                                v14 += (0.1 / 0.1)
                                                fmt.Println(fi26, v27, v27)
                                            case "a":
                                                if false {
                                                    v1 = (!true)
                                                } else {
                                                    fmt.Println(v1, fi26)
                                                    fmt.Println(fi26, fi26)
                                                }
                                            default:
                                                v27 = v27
                                            }
                                            if v1 {
                                                var v28 string = "\n\t"
                                                if (true && v1) {
                                                    v1 = v1
                                                    if (200 <= 557) {
                                                        v1 = true
                                                        if v1 {
                                                            v28 = "hello"
                                                            v28 = ""
                                                            var v29 int32 = -608
                                                            _ = v29
                                                        }
                                                    } else {
                                                        v14 += v14
                                                        v30 := GenS3{F0: (-(fi17)), F1: 0}
                                                        _ = v30
                                                    }
                                                }
                                                if (!true) {
                                                    var v31 int = fi17
                                                    v32 := GenS2{F0: 255}
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    var v33 uint16 = 208
                                                    _ = v33
                                                }
                                                _ = v28
                                            }
                                            v34 := GenS4{F0: "", F1: 227, F2: 18446744073709551615}
                                            _ = v34
                                        }
                                        _ = fi26
                                        _ = v27
                                    }
                                }
                            }
                            v35 := GenS1{F0: (fi17 / (((fi17 * (fi17 - (fi17 | (fi17 >> ((1) & 63)))))) | 1)), F1: ((("\n\t" + "\n\t") + "\n\t") + ("hello" + (("" + "a") + "hello")))}
                            _ = v35
                        }
                        var v36 uint16 = 1
                        _ = v36
                    } else {
                        v1 = (4294967295 == 42)
                        _ = len(v3)
                        v14 += 1.5
                    }
                } else {
                    var v37 int64 = 9223372036854775807
                    for fi38 := 0; fi38 < 20; fi38++ {
                        fmt.Println(v37)
                        switch fi17 {
                        case 380:
                            v2 = append(v2, "hello")
                        case 42:
                            if (v1 && true) {
                                if len(v4) > 0 {
                                    v4[0] = ("\n\t" + ("hello" + ("" + "")))
                                }
                                var v39 int16 = 1000
                                if v1 {
                                    var v40 float64 = v14
                                    v41 := v3[("" + ("\n\t" + (("\n\t" + "a") + "")))]
                                    _ = v40
                                    _ = v41
                                }
                                _ = v39
                            } else {
                                var v42 bool = (!(v1 && false))
                                v1 = (1 <= 1)
                                _ = v42
                            }
                            fmt.Println(v14, v14, v1)
                            v43 := GenS5{F0: v1}
                            _ = v43
                        case 4294967295:
                            var v44 uint32 = 42
                            _ = v44
                        default:
                            v45 := GenS4{F0: "hello", F1: 32767, F2: genFn0(285, 32767, 1)}
                            _ = v45
                        }
                        switch "" {
                        case "a":
                            if (4294967295 == 2147483647) {
                                v1 = v1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = len(v3)
                                var v46 int32 = 1
                                _ = v46
                            } else {
                                v37 = v37
                                fmt.Println(v37, v37)
                                var v47 uint64 = 18446744073709551615
                                _ = v47
                            }
                        case "hello":
                            switch int(v37) {
                            case 2147483647:
                                fmt.Println(fi17, fi17)
                            case 4294967295:
                                if true {
                                    v48 := GenS1{F0: fi17, F1: ""}
                                    _ = v48
                                } else {
                                    v1 = (1.5 >= 0.0)
                                    if (!(1 <= 42)) {
                                        var v49 float32 = (0.0 * 1.5)
                                        if (!(-113 >= -128)) {
                                            switch 2147483647 {
                                            case 127:
                                                if (-351 < 0) {
                                                    switch (fi17 % ((127) | 1)) {
                                                    case -807:
                                                        _ = len(v3)
                                                        var v50 int64 = (v37 / ((v37) | 1))
                                                        if v1 {
                                                            var v51 uint32 = 42
                                                            fmt.Println(v49, v1)
                                                            _ = v51
                                                        } else {
                                                            v50 -= (v50 / ((v37) | 1))
                                                        }
                                                        _ = v50
                                                    default:
                                                        if false {
                                                            v1 = (42 > 339)
                                                            if (394 == 0) {
                                                                v52 := GenS6{F0: "hello", F1: 18446744073709551615}
                                                                delete(v3, (("\n\t" + "") + ""))
                                                                _ = v52
                                                            }
                                                        } else {
                                                            v37 = (v37 | (v37 + 9223372036854775807))
                                                            v49 = 0.1
                                                            fmt.Println(v1)
                                                        }
                                                    }
                                                } else {
                                                    delete(v3, ("hello" + "foo bar"))
                                                    var v53 uint16 = 329
                                                    _ = v53
                                                }
                                            case 1:
                                                if (v1 && (271 != -2147483648)) {
                                                    var v54 int = 1
                                                    v2 = append(v2, "")
                                                    _ = v54
                                                }
                                            default:
                                                fmt.Println(v37)
                                            }
                                            var v55 uint64 = 18446744073709551615
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v1 = v1
                                            _ = v55
                                        } else {
                                            v56 := v3[((("" + ("hello" + "foo bar")) + "\n\t") + ("\n\t" + ""))]
                                            var v57 uint = 502
                                            if (false && v1) {
                                                if v1 {
                                                    var v58 uint16 = 774
                                                    v59 := GenS6{F0: ("hello" + (("hello" + "") + "")), F1: genFn0((v57 / ((genFn0(genFn0((v57 >> ((v57) & 63)), 0, -104), int16(v56), -2147483648)) | 1)), -607, 0)}
                                                    v49 = (1.5 + float32(((0.0 + float64(v49)) + v14)))
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = len(v4)
                                            } else {
                                                var v60 bool = v1
                                                v60 = v60
                                                _ = v60
                                            }
                                            _ = v56
                                            _ = v57
                                        }
                                        if v1 {
                                            var v61 uint8 = 127
                                            var v62 bool = (v14 != v14)
                                            fmt.Println(v62, v61)
                                            _ = v61
                                            _ = v62
                                        }
                                        fmt.Println(fi17, fi17)
                                        if v1 {
                                            var v63 uint16 = 255
                                            fmt.Println(fi17)
                                            _ = v63
                                        } else {
                                            v37 += v37
                                        }
                                        _ = v49
                                    } else {
                                        _ = len(v2)
                                        var v64 int64 = (v37 & (v37 / ((v37) | 1)))
                                        var v65 uint32 = 2147483647
                                        _ = v64
                                        _ = v65
                                    }
                                    var v66 int8 = genFn1()
                                    if (!(!true)) {
                                        var v67 int64 = v37
                                        var v68 int32 = -1
                                        _ = v67
                                        _ = v68
                                    }
                                    _ = v66
                                }
                            case 239:
                                if (2147483648 == 2147483647) {
                                    delete(v3, "")
                                    v1 = true
                                } else {
                                    v69 := v3[("a" + "foo bar")]
                                    _ = v69
                                }
                            default:
                                if false {
                                    v37 += v37
                                    fmt.Println(v1, v14, fi17)
                                } else {
                                    var v70 int64 = 42
                                    var v71 int16 = 1
                                    if v1 {
                                        var v72 uint = 42
                                        v37 = 0
                                        var v73 int = -9223372036854775808
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = v70
                                    _ = v71
                                }
                            }
                        default:
                            v14 = -2.25
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (v1 && (false && (0 != 74))) {
                            var v74 uint8 = 1
                            _ = len(v3)
                            v74 += (^(v74))
                            if false {
                                v37 = (v37 / (((v37 * (v37 % ((-9223372036854775808) | 1)))) | 1))
                            } else {
                                fmt.Println(fi38, v37)
                                v74 -= 227
                            }
                            v74 = v74
                            _ = v74
                        } else {
                            if v1 {
                                if (!false) {
                                    v75 := GenS6{F0: ("foo bar" + ""), F1: 2}
                                    var v76 uint8 = 42
                                    _ = v75
                                    _ = v76
                                }
                            }
                            var v77 int64 = v37
                            v78 := GenS5{F0: false}
                            _ = v77
                            _ = v78
                        }
                        _ = fi38
                    }
                    for fi79 := 0; fi79 < 9; fi79++ {
                        v14 = v14
                        if (((("foo bar" + "") >= ("foo bar" + "")) || v1) || (129 > 28)) {
                            switch fi79 {
                            case 128:
                                v14 = float64(0.1)
                            case -1:
                                delete(v3, (("" + "\n\t") + "foo bar"))
                            case 369:
                                var v80 float64 = ((0.0 - 0.1) + (v14 + (v14 / v14)))
                                _ = v80
                            }
                            fmt.Println(fi79, fi17)
                        } else {
                            switch fi79 {
                            case 9223372036854775807:
                                fmt.Println(v14)
                            case 1:
                                v37 = v37
                            case 42:
                                var v81 int32 = -482
                                _ = v81
                            }
                            if v1 {
                                var v82 uint = 428
                                _ = v82
                            }
                            v14 = (v14 / -2.25)
                        }
                        _ = fi79
                    }
                    _ = v37
                }
                _ = fi17
            }
            v1 = (false && v1)
        } else {
            if (87 == -67) {
                var v83 uint16 = 0
                fmt.Println(v83)
                switch 9223372036854775807 {
                case 310:
                    var v84 int32 = 42
                    _ = v84
                default:
                    if true {
                        fmt.Println(v83)
                        var v85 int64 = -1
                        _ = v85
                    }
                }
                _ = v83
            } else {
                for fi86 := 0; fi86 < 18; fi86++ {
                    var v87 uint64 = 133
                    switch "hello" {
                    case "foo bar":
                        if (v1 || v1) {
                            if false {
                                var v88 int8 = 42
                                if v1 {
                                    for fi89 := 0; fi89 < 14; fi89++ {
                                        if v1 {
                                            if true {
                                                fmt.Println(fi89)
                                            } else {
                                                fmt.Println(v1, v87, fi89)
                                                fmt.Println(fi86, fi89)
                                            }
                                            v90 := v3["a"]
                                            _ = v90
                                        } else {
                                            var v91 uint8 = 255
                                            delete(v3, "hello")
                                            if (32767 > 320) {
                                                if true {
                                                    v92 := GenS7{F0: (!v1), F1: 1000}
                                                    var v93 uint16 = 42
                                                    _ = v92
                                                    _ = v93
                                                }
                                            } else {
                                                v94 := GenS8{F0: "hello", F1: genFn0(1, 1, 0), F2: (("\n\t" + "\n\t") + ("foo bar" + ("" + "hello")))}
                                                _ = v94
                                            }
                                            _ = v91
                                        }
                                        v95 := GenS9{F0: 18446744073709551615}
                                        _ = fi89
                                        _ = v95
                                    }
                                    _ = len(v3)
                                    if (!true) {
                                        var v96 float64 = float64(1e10)
                                        var v97 uint32 = 2147483648
                                        _ = v96
                                        _ = v97
                                    }
                                }
                                _ = v88
                            } else {
                                switch "a" {
                                case "foo bar":
                                    var v98 int = fi86
                                    _ = v98
                                }
                                if ((!v1) || false) {
                                    var v99 int64 = -9223372036854775808
                                    _ = v99
                                }
                                fmt.Println(v87)
                            }
                            var v100 bool = (781 >= 42)
                            fmt.Println(v1)
                            _ = v100
                        } else {
                            var v101 bool = true
                            for fi102 := 0; fi102 < 20; fi102++ {
                                fmt.Println(v87)
                                v103 := GenS3{F0: (fi102 * fi86), F1: 1}
                                v104 := GenS1{F0: 128, F1: "a"}
                                _ = fi102
                                _ = v103
                                _ = v104
                            }
                            switch "hello" {
                            case "foo bar":
                                v105 := GenS3{F0: 1, F1: 42}
                                _ = v105
                            default:
                                v106 := GenS10{F0: 255, F1: 452, F2: 42}
                                _ = v106
                            }
                            _ = v101
                        }
                    case "\n\t":
                        v14 -= v14
                    }
                    switch 32768 {
                    case 32768:
                        if (42 <= 250) {
                            var v107 int = fi86
                            _ = v107
                        }
                    case -9223372036854775808:
                        if (1 > 2) {
                            v108 := GenS2{F0: 127}
                            _ = v108
                        }
                    default:
                        var v109 int32 = 2147483647
                        _ = v109
                    }
                    _ = fi86
                    _ = v87
                }
                var v110 uint8 = 164
                v3["\n\t"] = 1
                _ = v110
            }
        }
        v111 := GenS1{F0: -142, F1: ("\n\t" + "hello")}
        _ = v111
        _ = v14
    }
    v1 = ((!true) || v1)
    v4 = append(v4, "\n\t")
    var v112 int32 = 0
    for fi113 := 0; fi113 < 19; fi113++ {
        if (1 > 42) {
            var v114 string = ""
            _ = v114
        } else {
            v1 = (!false)
            fmt.Println(v112)
            switch "\n\t" {
            case "":
                fmt.Println(fi113, v1)
            }
        }
        _ = fi113
    }
    for fi115 := 0; fi115 < 20; fi115++ {
        for fi116 := 0; fi116 < 1; fi116++ {
            var v117 string = "hello"
            fmt.Println(v1, v1, v112)
            v117 = v117
            v4 = append(v4, ("\n\t" + ((v117 + v117) + v117)))
            v118 := GenS11{F0: 65535, F1: (("" + (v117 + v117)) + (("a" + ((v117 + v117) + "\n\t")) + (v117 + "hello"))), F2: 102}
            _ = fi116
            _ = v117
            _ = v118
        }
        _ = fi115
    }
    for fi119 := 0; fi119 < 20; fi119++ {
        if v1 {
            if true {
                v112 = (v112 & (v112 | v112))
                if (0 < v112) {
                    fmt.Println(v1)
                    v1 = v1
                    fmt.Println(v112)
                }
            } else {
                for fi120 := 0; fi120 < 19; fi120++ {
                    if (v1 && (!((361 < -32768) && (1.5 >= (1.5 / (1.5 - ((0.1 / (0.1 * (3.14159 + ((1e10 - (0.1 / (1.5 / ((3.14159 + float64(1e10)) - 0.1)))) * (3.14159 * (1e10 * 0.1)))))) / ((float64(0.1) * 3.14159) - 0.0)))))))) {
                        var v121 string = ("a" + "a")
                        if (1.5 <= 1e10) {
                            if true {
                                v112 = 937
                            } else {
                                v122 := GenS4{F0: ("a" + (v121 + "hello")), F1: 65535, F2: 2}
                                if ((fi119 | int(v112)) <= fi120) {
                                    v123 := GenS1{F0: fi119, F1: v121}
                                    _ = v123
                                } else {
                                    fmt.Println(v112, v112, v1)
                                    v121 = (v121 + "foo bar")
                                }
                                _ = v122
                            }
                            v121 = v121
                        } else {
                            var v124 bool = (v1 && false)
                            fmt.Println(v1)
                            _ = v124
                        }
                        var v125 int8 = 42
                        _ = v121
                        _ = v125
                    }
                    switch "a" {
                    case "a":
                        switch ("\n\t" + "foo bar") {
                        case "\n\t":
                            v112 += (v112 << ((42) & 63))
                        case "hello":
                            var v126 uint8 = 127
                            _ = v126
                        case "":
                            v112 += v112
                        default:
                            v127 := GenS9{F0: 0}
                            _ = v127
                        }
                    case "hello":
                        v2 = append(v2, ("" + ("foo bar" + "\n\t")))
                        if v1 {
                            var v128 float64 = -2.25
                            _ = v128
                        } else {
                            var v129 uint16 = 896
                            _ = v129
                        }
                    case "":
                        var v130 uint64 = 648
                        _ = v130
                    default:
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        v112 += v112
                    }
                    v131 := GenS12{F0: (fi119 / ((127) | 1)), F1: 746}
                    _ = fi120
                    _ = v131
                }
                if false {
                    v132 := GenS13{F0: -15}
                    _ = v132
                } else {
                    v1 = (43 >= 9223372036854775807)
                    v1 = ((fi119 >= -386) || v1)
                }
            }
            v1 = (v1 && v1)
        }
        _ = fi119
    }
    v1 = v1
    var v133 bool = v1
    for fi134 := 0; fi134 < 7; fi134++ {
        var v135 int = (fi134 / (((fi134 * 402)) | 1))
        switch "a" {
        case "a":
            if false {
                v136 := GenS10{F0: 4294967295, F1: 470, F2: 51}
                _ = v136
            } else {
                switch "a" {
                case "hello":
                    v137 := GenS14{F0: (v112 << ((2) & 63))}
                    _ = v137
                case "":
                    var v138 string = "foo bar"
                    _ = v138
                }
                var v139 uint8 = 214
                _ = v139
            }
        case "hello":
            fmt.Println(fi134)
        case "\n\t":
            switch (v135 | fi134) {
            case 4294967295:
                var v140 int = v135
                _ = v140
            case -9223372036854775808:
                var v141 float32 = (0.1 + ((0.0 / 1.5) + 0.0))
                _ = v141
            case -961:
                var v142 uint64 = 42
                _ = v142
            }
        }
        var v143 int64 = -918
        _ = fi134
        _ = v135
        _ = v143
    }
    v133 = false
    if true {
        fmt.Println(v1, v112, v133)
    } else {
        for fi144 := 0; fi144 < 4; fi144++ {
            for fi145 := 0; fi145 < 4; fi145++ {
                v112 -= v112
                _ = fi145
            }
            _ = fi144
        }
    }
    var v146 uint16 = 880
    var v147 bool = false
    if (v1 || true) {
        for fi148 := 0; fi148 < 19; fi148++ {
            switch "foo bar" {
            case "hello":
                fmt.Println(fi148)
            case "a":
                if v147 {
                    if v133 {
                        if (v1 && v133) {
                            fmt.Println(fi148, v146)
                            var v149 uint8 = 162
                            v1 = v147
                            _ = v149
                        }
                    }
                    v150 := GenS15{F0: 547, F1: 65535}
                    _ = v150
                }
            }
            v147 = (false || false)
            v151 := GenS10{F0: 2, F1: 42, F2: 100}
            _ = fi148
            _ = v151
        }
        if v133 {
            v1 = ("" <= "foo bar")
            _ = len(v3)
        }
        if (v1 || false) {
            if v147 {
                v147 = (42 <= 42)
                v4 = append(v4, "hello")
            }
            fmt.Println(v146, v1, v146)
        }
    } else {
        _ = len(v4)
        var v152 uint64 = 18446744073709551615
        fmt.Println(v133)
        _ = v152
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v133:", v133)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v112
    _ = v133
    _ = v146
    _ = v147
    _ = v2
    _ = v3
    _ = v4
}
