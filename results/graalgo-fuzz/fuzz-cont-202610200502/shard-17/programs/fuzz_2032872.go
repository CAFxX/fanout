package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint8; F2 int }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 int16; F1 int8; F2 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint; F1 uint64; F2 int8 }
func genFn0(p0 int16) uint64 {
    var v4 float32 = 1.5
    v5 := GenS1{F0: -104, F1: int8(p0)}
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 float64 = 1e10
    var v2 bool = (false && (0 < 1))
    v3 := make(map[string]bool)
    v6 := []int{}
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v1 = v1
            v2 = ((42 > 128) && (!("hello" != ("\n\t" + ("hello" + "foo bar")))))
            _ = fi8
        }
        switch (fi7 | -666) {
        case 4294967295:
            fmt.Println(v2)
        }
        _ = fi7
    }
    delete(v3, "hello")
    v2 = ((v2 && false) && ((!(0.0 >= (0.0 + 3.14159))) || v2))
    var v9 int16 = -62
    if true {
        if (4294967295 == 2147483648) {
            if v2 {
                fmt.Println(v2, v1)
                v10 := GenS1{F0: -9223372036854775808, F1: 100}
                fmt.Println(v9)
                _ = v10
            } else {
                v11 := v3["a"]
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 int32 = 1
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
        } else {
            delete(v3, ("" + ("\n\t" + "foo bar")))
        }
        if true {
            if (725 == 510) {
                switch "hello" {
                case "a":
                    var v14 float64 = v1
                    _ = v14
                }
            }
        }
    } else {
        v2 = v2
        switch "a" {
        case "":
            if false {
                v9 -= 1
            }
        case "hello":
            if v2 {
                v9 = (v9 >> ((1) & 63))
                var v15 string = ("a" + "a")
                var v16 int64 = -677
                _ = v15
                _ = v16
            }
        default:
            fmt.Println(v1)
        }
        var v17 bool = (42 != -240)
        _ = v17
    }
    var v18 int8 = 0
    var v19 uint16 = 264
    switch 128 {
    case 2147483647:
        var v20 string = "a"
        _ = v20
    default:
        for fi21 := 0; fi21 < 4; fi21++ {
            if v2 {
                fmt.Println(fi21, v2, fi21)
            } else {
                var v22 bool = (v2 && v2)
                switch -2147483648 {
                case -1:
                    switch "\n\t" {
                    case "":
                        for fi23 := 0; fi23 < 9; fi23++ {
                            v24 := GenS1{F0: -469, F1: (v18 / ((66) | 1))}
                            var v25 int = fi21
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v26 := GenS1{F0: 1, F1: (v18 % (((v18 % ((v18) | 1))) | 1))}
                            v9 = v9
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        fmt.Println(v9)
                    case "\n\t":
                        var v27 float32 = float32(0.1)
                        _ = v27
                    default:
                        if (!v22) {
                            switch "\n\t" {
                            case "foo bar":
                                v3["foo bar"] = (v2 || (-233 != 1))
                            }
                            v28 := GenS2{F0: 0, F1: 88, F2: -42}
                            if (-1 > 0) {
                                switch (-(fi21)) {
                                case 42:
                                    v29 := GenS1{F0: -9223372036854775808, F1: v18}
                                    _ = v29
                                case -9223372036854775808:
                                    var v30 bool = (fi21 == -9223372036854775808)
                                    _ = v30
                                default:
                                    v6 = append(v6, (-(fi21)))
                                }
                                var v31 int = (fi21 / ((fi21) | 1))
                                _ = v31
                            } else {
                                for fi32 := 0; fi32 < 13; fi32++ {
                                    if (65535 > 787) {
                                        switch fi32 {
                                        case -429:
                                            if v2 {
                                                fmt.Println(v19, v18, v1)
                                                v9 = 827
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            } else {
                                                v22 = false
                                                var v33 int64 = -979
                                                var v34 int64 = 1
                                                _ = v33
                                                _ = v34
                                            }
                                        case 32768:
                                            v2 = false
                                        }
                                        if v22 {
                                            var v35 int16 = (v9 / (((v9 ^ v9)) | 1))
                                            _ = v35
                                        } else {
                                            v36 := GenS3{F0: 848, F1: v18}
                                            var v37 uint16 = 65535
                                            v22 = v2
                                            if false {
                                                var v38 uint = 255
                                                delete(v3, "foo bar")
                                                _ = v38
                                            } else {
                                                var v39 string = "\n\t"
                                                v37 = (v37 | (v37 + (^(v19))))
                                                _ = v39
                                            }
                                            _ = v36
                                            _ = v37
                                        }
                                        v2 = v22
                                    } else {
                                        v22 = (544 != 42)
                                        var v40 int32 = -706
                                        _ = v40
                                    }
                                    _ = fi32
                                }
                            }
                            _ = v28
                        } else {
                            v2 = (42 == 315)
                        }
                    }
                default:
                    for fi41 := 0; fi41 < 17; fi41++ {
                        v2 = true
                        _ = fi41
                    }
                }
                _ = v22
            }
            switch "" {
            case "a":
                var v42 bool = ((381 > 42) && (-235 != -888))
                _ = v42
            default:
                switch (^(fi21)) {
                case -2147483648:
                    var v43 bool = ((2147483647 <= 751) || (((^(v19)) >= 100) && ((1 <= (^(v9))) && ((v2 || (v2 || ((255 < 2) || (v2 && v2)))) && v2))))
                    _ = v43
                default:
                    var v44 int32 = 1
                    _ = v44
                }
            }
            _ = fi21
        }
    }
    if v2 {
        for fi45 := 0; fi45 < 17; fi45++ {
            if false {
                if false {
                    v19 = v19
                } else {
                    switch "hello" {
                    case "a":
                        for fi46 := 0; fi46 < 4; fi46++ {
                            fmt.Println(v2, v2)
                            var v47 float32 = 0.1
                            v1 = v1
                            v18 = (v18 >> ((18446744073709551615) & 63))
                            _ = fi46
                            _ = v47
                        }
                    case "\n\t":
                        v48 := GenS2{F0: 255, F1: 1, F2: v18}
                        _ = v48
                    case "hello":
                        switch (("a" + "") + "foo bar") {
                        case "\n\t":
                            if v2 {
                                if (("" + ("foo bar" + "a")) != "hello") {
                                    for fi49 := 0; fi49 < 1; fi49++ {
                                        var v50 float32 = 0.1
                                        var v51 int32 = 0
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                    var v52 int16 = 848
                                    for fi53 := 0; fi53 < 12; fi53++ {
                                        fmt.Println(v52, v1, fi45)
                                        v52 = v9
                                        _ = fi53
                                    }
                                    _ = v52
                                } else {
                                    var v54 uint16 = v19
                                    if false {
                                        if false {
                                            var v55 uint = 18446744073709551615
                                            for fi56 := 0; fi56 < 17; fi56++ {
                                                if (!v2) {
                                                    if ("a" < "foo bar") {
                                                        v2 = v2
                                                        var v57 int8 = v18
                                                        if v2 {
                                                            v57 = v57
                                                            if (v55 <= 558) {
                                                                v1 += v1
                                                                var v58 uint = v55
                                                                fmt.Println(fi45, v58)
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                _ = v58
                                                            } else {
                                                                fmt.Println(fi56, v55)
                                                            }
                                                        }
                                                        v59 := GenS4{F0: (v9 != (v9 >> ((18446744073709551615) & 63))), F1: (v54 << (((^(v55))) & 63)), F2: "\n\t"}
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        _ = v57
                                                        _ = v59
                                                    }
                                                }
                                                if true {
                                                    if false {
                                                        _ = len(v6)
                                                        v9 = v9
                                                    } else {
                                                        if len(v6) > 0 {
                                                            v6[0] = fi45
                                                        }
                                                        var v60 uint16 = v54
                                                        _ = v60
                                                    }
                                                }
                                                var v61 int = 42
                                                _ = fi56
                                                _ = v61
                                            }
                                            var v62 uint8 = 245
                                            _ = v55
                                            _ = v62
                                        }
                                    }
                                    _ = v54
                                }
                                for fi63 := 0; fi63 < 13; fi63++ {
                                    switch (fi63 | (fi45 >> ((795) & 63))) {
                                    case 127:
                                        var v64 uint8 = 42
                                        _ = v64
                                    case 0:
                                        if ("\n\t" > "hello") {
                                            var v65 int64 = -9223372036854775808
                                            var v66 int16 = 239
                                            var v67 int16 = (v9 - -586)
                                            v2 = true
                                            switch "hello" {
                                            case "":
                                                if (2147483648 > 65535) {
                                                    var v68 uint16 = v19
                                                    _ = v68
                                                }
                                            case "foo bar":
                                                var v69 bool = v2
                                                _ = v69
                                            case "\n\t":
                                                switch "a" {
                                                case "":
                                                    if len(v6) > 0 {
                                                        v6[0] = fi45
                                                    }
                                                case "\n\t":
                                                    v9 = v67
                                                case "a":
                                                    v2 = true
                                                default:
                                                    v6 = append(v6, 0)
                                                }
                                            default:
                                                v70 := GenS5{F0: v67}
                                                _ = v70
                                            }
                                            v19 += v19
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        } else {
                                            switch ("" + "hello") {
                                            case "hello":
                                                v9 = (v9 << ((18446744073709551615) & 63))
                                                if len(v6) > 0 {
                                                    v6[0] = (fi63 + fi63)
                                                }
                                                if (48 != 128) {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                } else {
                                                    var v71 string = "foo bar"
                                                    _ = v71
                                                }
                                            case "\n\t":
                                                v6 = append(v6, fi45)
                                            default:
                                                var v72 string = "a"
                                                _ = v72
                                            }
                                            var v73 uint8 = 12
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            if (618 > 2) {
                                                _ = len(v3)
                                                _ = len(v6)
                                                v74 := v3[""]
                                                var v75 int = (fi63 | (fi45 | -9223372036854775808))
                                                fmt.Println(v9, v19, v75)
                                                v76 := GenS6{F0: 65535, F1: v73, F2: (fi63 % ((v75) | 1))}
                                                v2 = (v73 == (v73 & v73))
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                            if v2 {
                                                fmt.Println(v73)
                                                if true {
                                                    var v77 uint = 1
                                                    var v78 uint8 = (v73 ^ uint8(v77))
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    if len(v6) > 0 {
                                                        v6[0] = (fi45 << ((4294967295) & 63))
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v79 float64 = v1
                                                    if true {
                                                        v80 := GenS3{F0: 777, F1: v18}
                                                        _ = v80
                                                    }
                                                    _ = v79
                                                }
                                            } else {
                                                v81 := GenS7{F0: genFn0(1), F1: v19}
                                                fmt.Println(v2)
                                                v18 = 21
                                                var v82 int32 = 1
                                                v6 = append(v6, (fi63 / ((fi63) | 1)))
                                                _ = v81
                                                _ = v82
                                            }
                                            _ = v73
                                        }
                                    default:
                                        v83 := GenS6{F0: 65535, F1: 1, F2: fi45}
                                        _ = v83
                                    }
                                    _ = fi63
                                }
                            }
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            v84 := GenS8{F0: -9223372036854775808}
                            _ = v84
                        }
                    default:
                        if false {
                            if (1 == 0) {
                                v85 := GenS8{F0: 1}
                                if (v2 && true) {
                                    fmt.Println(v9)
                                    for fi86 := 0; fi86 < 4; fi86++ {
                                        var v87 uint = 1
                                        if false {
                                            if (3.14159 >= (1.5 + (1.5 * -2.25))) {
                                                if false {
                                                    var v88 uint64 = 4294967295
                                                    fmt.Println(v2)
                                                    v89 := GenS9{F0: 1}
                                                    _ = v88
                                                    _ = v89
                                                }
                                                switch fi45 {
                                                case 127:
                                                    fmt.Println(v87, fi86)
                                                default:
                                                    fmt.Println(v18, fi45)
                                                }
                                            } else {
                                                if (!v2) {
                                                    var v90 string = ("foo bar" + "")
                                                    if v2 {
                                                        switch fi45 {
                                                        case -673:
                                                            switch (((v90 + v90) + v90) + v90) {
                                                            case "\n\t":
                                                                v19 -= v19
                                                            case "a":
                                                                v2 = v2
                                                            }
                                                        case 2147483647:
                                                            v91 := GenS10{F0: (!true), F1: ""}
                                                            _ = v91
                                                        case -540:
                                                            v92 := GenS11{F0: v9, F1: 0, F2: 647}
                                                            _ = v92
                                                        }
                                                        switch ((("foo bar" + ("a" + ("foo bar" + (("hello" + "foo bar") + (v90 + v90))))) + v90) + "") {
                                                        case "a":
                                                            var v93 float64 = float64(0.0)
                                                            var v94 float64 = (float64((0.1 / -2.25)) - v93)
                                                            _ = v93
                                                            _ = v94
                                                        case "foo bar":
                                                            fmt.Println(v19, fi45, v18)
                                                        case "hello":
                                                            v95 := GenS9{F0: 2}
                                                            v96 := GenS12{F0: 4294967295}
                                                            v97 := GenS13{F0: (v87 + 65535), F1: 42, F2: v18}
                                                            _ = v95
                                                            _ = v96
                                                            _ = v97
                                                        default:
                                                            v2 = v2
                                                        }
                                                        switch fi86 {
                                                        case 248:
                                                            v87 -= (v87 + uint(v19))
                                                        case 9223372036854775807:
                                                            v98 := GenS10{F0: v2, F1: ("a" + "a")}
                                                            _ = v98
                                                        default:
                                                            v99 := v3[("hello" + (((("hello" + ("a" + (v90 + "a"))) + (v90 + v90)) + v90) + ((v90 + "hello") + (v90 + (v90 + (v90 + v90))))))]
                                                            _ = v99
                                                        }
                                                        var v100 bool = v2
                                                        var v101 int64 = 9223372036854775807
                                                        var v102 bool = (0.0 < v1)
                                                        v90 = v90
                                                        _ = v100
                                                        _ = v101
                                                        _ = v102
                                                    } else {
                                                        v3[v90] = v2
                                                    }
                                                    v1 = ((float64(3.14159) - v1) / float64(((0.0 - 3.14159) - (1.5 * (3.14159 - (0.1 / 1.0))))))
                                                    v103 := GenS1{F0: 1, F1: v18}
                                                    _ = v103
                                                    _ = v90
                                                } else {
                                                    var v104 bool = (v2 || (!(v87 <= v87)))
                                                    var v105 uint8 = 119
                                                    _ = v104
                                                    _ = v105
                                                }
                                                v18 += 127
                                                var v106 float32 = (1.5 / 1.0)
                                                _ = v106
                                            }
                                            var v107 int = -971
                                            _ = v107
                                        } else {
                                            var v108 uint = 1
                                            var v109 int32 = 42
                                            var v110 uint = v87
                                            _ = v108
                                            _ = v109
                                            _ = v110
                                        }
                                        v111 := GenS4{F0: ((((v9 << ((65535) & 63)) == v9) && true) || v2), F1: v19, F2: "\n\t"}
                                        var v112 uint64 = 42
                                        _ = fi86
                                        _ = v111
                                        _ = v112
                                        _ = v87
                                    }
                                } else {
                                    v1 = v1
                                }
                                if (false && v2) {
                                    v2 = (!(v9 > (v9 & v9)))
                                    if (2147483648 <= 42) {
                                        var v113 int64 = -9223372036854775808
                                        for fi114 := 0; fi114 < 2; fi114++ {
                                            v18 = 42
                                            switch (fi45 ^ 2147483647) {
                                            case -925:
                                                var v115 uint8 = 1
                                                _ = v115
                                            case 1:
                                                fmt.Println(v1)
                                            }
                                            _ = fi114
                                        }
                                        var v116 int64 = (v113 >> ((910) & 63))
                                        _ = v113
                                        _ = v116
                                    }
                                    v2 = v2
                                }
                                _ = v85
                            }
                        } else {
                            switch (fi45 - 804) {
                            case 2147483647:
                                var v117 int32 = -619
                                _ = v117
                            default:
                                v19 = 1
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    }
                    switch "" {
                    case "foo bar":
                        var v118 bool = true
                        _ = v118
                    default:
                        v9 += v9
                    }
                    if (4294967295 > fi45) {
                        var v119 bool = false
                        _ = v119
                    }
                }
            }
            v2 = (v2 || (v2 && (true && (-1 == -451))))
            v9 += (^(v9))
            _ = fi45
        }
        switch 4294967295 {
        case -773:
            var v120 int8 = v18
            _ = v120
        case 128:
            v121 := GenS4{F0: (505 >= 2147483648), F1: v19, F2: ""}
            _ = v121
        case -9223372036854775808:
            switch 128 {
            case 127:
                for fi122 := 0; fi122 < 12; fi122++ {
                    var v123 string = "hello"
                    _ = fi122
                    _ = v123
                }
            default:
                var v124 float64 = 3.14159
                _ = v124
            }
        }
        v3["a"] = v2
    } else {
        if (-825 >= int64(v18)) {
            var v125 uint32 = 688
            v18 = (v18 % (((v18 % ((v18) | 1))) | 1))
            _ = v125
        } else {
            switch 127 {
            case 1:
                var v126 uint16 = 42
                _ = v126
            case -9223372036854775808:
                if false {
                    for fi127 := 0; fi127 < 20; fi127++ {
                        for fi128 := 0; fi128 < 10; fi128++ {
                            delete(v3, ((("foo bar" + "\n\t") + (("foo bar" + "") + "a")) + "a"))
                            var v129 float64 = (v1 / v1)
                            if ((v129 / (float64(1.5) + 0.1)) <= v1) {
                                v3[("foo bar" + "")] = false
                            } else {
                                v18 -= -128
                                v1 -= 0.0
                                fmt.Println(v2, fi128)
                            }
                            _ = fi128
                            _ = v129
                        }
                        var v130 uint32 = 65535
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi127
                        _ = v130
                    }
                    if (42 >= 168) {
                        fmt.Println(v2, v18)
                        for fi131 := 0; fi131 < 19; fi131++ {
                            if (!true) {
                                fmt.Println(v2, v18)
                                _ = len(v6)
                                for fi132 := 0; fi132 < 13; fi132++ {
                                    var v133 bool = true
                                    if true {
                                        fmt.Println(v1, v133, fi131)
                                        var v134 uint32 = 2
                                        _ = v134
                                    }
                                    _ = len(v6)
                                    _ = fi132
                                    _ = v133
                                }
                            }
                            _ = fi131
                        }
                    }
                }
            }
        }
        v135 := GenS3{F0: 4294967295, F1: v18}
        var v136 int64 = -1
        _ = v135
        _ = v136
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
