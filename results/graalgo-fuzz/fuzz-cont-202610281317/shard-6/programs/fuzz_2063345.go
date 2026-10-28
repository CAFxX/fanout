package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 bool; F1 bool }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS14 struct { F0 uint8; F1 uint64 }
type GenS15 struct { F0 bool; F1 uint32 }
type GenS16 struct { F0 int16; F1 uint64; F2 uint }
type GenS17 struct { F0 int32; F1 uint32 }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    var v6 int64 = -313
    v7 := GenS1{F0: 255}
    _ = p0
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 uint32 = 2
    v2 := []string{("hello" + "\n\t"), ("\n\t" + ((("" + ("a" + "")) + "") + "a")), ("foo bar" + "\n\t")}
    v3 := make(map[string]string)
    var v4 float32 = 1.5
    v5 := make(map[string]string)
    v8 := GenS2{F0: -818}
    if true {
        var v9 uint32 = (v1 + 959)
        _ = v9
    }
    v10 := GenS2{F0: -1}
    v11 := GenS3{F0: 0}
    var v12 int16 = 1000
    switch 128 {
    case 384:
        for fi13 := 0; fi13 < 6; fi13++ {
            v4 = v4
            v2 = append(v2, "foo bar")
            v1 -= (v1 % ((956) | 1))
            _ = fi13
        }
    }
    var v14 int = 9223372036854775807
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS4{F0: 1, F1: "hello"}
        _ = fi15
        _ = v16
    }
    v12 += -743
    for fi17 := 0; fi17 < 6; fi17++ {
        v12 = v12
        switch (fi17 * v14) {
        case -992:
            for fi18 := 0; fi18 < 14; fi18++ {
                v1 += v1
                _ = fi18
            }
        default:
            v19 := v5["foo bar"]
            _ = v19
        }
        v20 := GenS2{F0: 0}
        _ = fi17
        _ = v20
    }
    if (false || false) {
        if (true || (965 >= 0)) {
            if true {
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := GenS4{F0: (v12 + v12), F1: "a"}
                    v23 := GenS4{F0: (v12 >> ((65535) & 63)), F1: "hello"}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 int = v14
                _ = v24
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v14 -= v14
            }
            switch v14 {
            case 42:
                switch ("hello" + "foo bar") {
                case "\n\t":
                    var v25 string = ("a" + ("" + "a"))
                    _ = v25
                }
            case 9223372036854775807:
                switch (v14 * v14) {
                case -127:
                    v12 = (v12 >> ((2) & 63))
                case 0:
                    for fi26 := 0; fi26 < 14; fi26++ {
                        v5["a"] = ("\n\t" + "foo bar")
                        for fi27 := 0; fi27 < 18; fi27++ {
                            v28 := GenS4{F0: v12, F1: ("hello" + ("hello" + ""))}
                            _ = fi27
                            _ = v28
                        }
                        _ = fi26
                    }
                case 9223372036854775807:
                    var v29 uint32 = 292
                    _ = v29
                default:
                    if (true || false) {
                        var v30 uint16 = 0
                        _ = v30
                    } else {
                        var v31 int64 = -537
                        var v32 uint8 = 0
                        _ = v31
                        _ = v32
                    }
                }
            case -927:
                for fi33 := 0; fi33 < 13; fi33++ {
                    if false {
                        switch v14 {
                        case 443:
                            fmt.Println(v14, v12)
                        case -559:
                            if (270 != 0) {
                                var v34 float64 = 0.0
                                v14 = v14
                                if false {
                                    fmt.Println(v34, v34)
                                    var v35 uint16 = 834
                                    var v36 uint32 = v1
                                    _ = v35
                                    _ = v36
                                }
                                _ = v34
                            }
                        case 128:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    _ = fi33
                }
            }
        }
        delete(v5, "")
        var v37 int = (-(v14))
        _ = v37
    }
    v4 += ((float32(-2.25) + ((((-2.25 - (1.5 * v4)) * (3.14159 - v4)) - 0.0) / (v4 / float32(3.14159)))) - ((float32((3.14159 / 0.1)) + (v4 - ((-2.25 + -2.25) - float32((1e10 + (0.0 + 3.14159)))))) + ((1.5 / v4) / (1e10 + v4))))
    var v38 int32 = 0
    if false {
        var v39 bool = (!((v14 >> ((228) & 63)) == v14))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v14 -= v14
        _ = v39
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            v2 = append(v2, (("\n\t" + "a") + "hello"))
            v41 := GenS5{F0: int8(fi40), F1: (v38 - 1)}
            _ = fi40
            _ = v41
        }
        if (2 > 318) {
            fmt.Println(v12, v14)
        }
        fmt.Println(v1, v38)
    }
    var v42 int32 = 2147483647
    if false {
        for fi43 := 0; fi43 < 13; fi43++ {
            fmt.Println(v1, v14)
            var v44 float32 = 0.1
            _ = fi43
            _ = v44
        }
        if false {
            v4 = v4
            v12 = 7
            switch v14 {
            case -218:
                for fi45 := 0; fi45 < 15; fi45++ {
                    switch "" {
                    case "\n\t":
                        switch ("hello" + ("" + ("" + "foo bar"))) {
                        case "a":
                            switch (fi45 - (fi45 << ((978) & 63))) {
                            case -9223372036854775808:
                                v42 = -69
                            case 0:
                                if (((32767 == 32767) && true) && (("hello" <= "foo bar") || true)) {
                                    if (255 <= genFn0(42)) {
                                        var v46 int = (v14 & (^(fi45)))
                                        fmt.Println(v1, v12)
                                        var v47 int = (-(fi45))
                                        _ = v46
                                        _ = v47
                                    } else {
                                        if false {
                                            fmt.Println(v4)
                                            v48 := GenS6{F0: v1}
                                            _ = v48
                                        } else {
                                            v49 := GenS2{F0: v38}
                                            v50 := GenS7{F0: v38, F1: 65535}
                                            switch v14 {
                                            case 42:
                                                fmt.Println(v14, v42, v14)
                                            case -698:
                                                v42 = v38
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                        var v51 uint8 = 0
                                        _ = v51
                                    }
                                    if ((!(fi45 == (-(v14)))) && (127 == 0)) {
                                        var v52 int = (fi45 ^ (fi45 / ((fi45) | 1)))
                                        v53 := GenS8{F0: (-(v12))}
                                        for fi54 := 0; fi54 < 17; fi54++ {
                                            var v55 bool = (!((v38 != (v42 & (v42 + v38))) || false))
                                            var v56 int32 = v38
                                            switch fi45 {
                                            case 2147483647:
                                                if v55 {
                                                    var v57 bool = (271 == 65535)
                                                    _ = v57
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v58 := v5["foo bar"]
                                                    _ = v58
                                                }
                                            case 4294967295:
                                                var v59 uint8 = 127
                                                _ = v59
                                            case 32767:
                                                var v60 uint64 = 65
                                                _ = v60
                                            default:
                                                var v61 float64 = 0.0
                                                _ = v61
                                            }
                                            _ = fi54
                                            _ = v55
                                            _ = v56
                                        }
                                        _ = v52
                                        _ = v53
                                    } else {
                                        switch v14 {
                                        case -780:
                                            for fi62 := 0; fi62 < 20; fi62++ {
                                                var v63 int8 = 127
                                                var v64 int32 = (v42 | (v38 & (v38 << ((18446744073709551615) & 63))))
                                                v12 = 0
                                                _ = fi62
                                                _ = v63
                                                _ = v64
                                            }
                                        default:
                                            for fi65 := 0; fi65 < 11; fi65++ {
                                                var v66 float32 = 1e10
                                                fmt.Println(v38, v12, v14)
                                                var v67 int64 = 9223372036854775807
                                                _ = fi65
                                                _ = v66
                                                _ = v67
                                            }
                                        }
                                    }
                                } else {
                                    if (311 <= -1) {
                                        v3[("" + ("hello" + "hello"))] = "hello"
                                        switch (-(fi45)) {
                                        case 2147483647:
                                            switch (v14 | (fi45 + v14)) {
                                            case 9223372036854775807:
                                                delete(v3, "a")
                                            default:
                                                v1 -= 65535
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                    } else {
                                        v68 := GenS4{F0: v12, F1: "foo bar"}
                                        if false {
                                            if true {
                                                if ((v4 - 1e10) <= (float32(-2.25) - (float32(1e10) - v4))) {
                                                    var v69 uint = 1
                                                    v70 := GenS1{F0: 1}
                                                    _ = v69
                                                    _ = v70
                                                }
                                            }
                                            switch v14 {
                                            case 2147483647:
                                                v71 := GenS2{F0: (v38 + v38)}
                                                _ = v71
                                            case 79:
                                                for fi72 := 0; fi72 < 7; fi72++ {
                                                    if false {
                                                        if false {
                                                            _ = len(v3)
                                                        }
                                                        if true {
                                                            var v73 int32 = v42
                                                            var v74 int32 = v73
                                                            _ = v73
                                                            _ = v74
                                                        } else {
                                                            if (!(v42 != (v38 / (((v38 << ((279) & 63))) | 1)))) {
                                                                switch v14 {
                                                                case -727:
                                                                    if true {
                                                                        if true {
                                                                            var v75 float32 = v4
                                                                            _ = v75
                                                                        }
                                                                        v42 -= v38
                                                                        var v76 float64 = (-2.25 * float64(v4))
                                                                        var v77 uint64 = 1
                                                                        v42 -= v38
                                                                        var v78 string = ("" + ("a" + ("foo bar" + ("" + (("hello" + "a") + "hello")))))
                                                                        if (!((v38 * -1) >= v38)) {
                                                                            var v79 int64 = 9223372036854775807
                                                                            if true {
                                                                                var v80 int = fi72
                                                                                _ = v80
                                                                            } else {
                                                                                v12 = (v12 % ((-32768) | 1))
                                                                            }
                                                                            _ = v79
                                                                        }
                                                                        _ = v76
                                                                        _ = v77
                                                                        _ = v78
                                                                    } else {
                                                                        v3["foo bar"] = "\n\t"
                                                                        v12 = 0
                                                                        var v81 uint8 = 201
                                                                        _ = v81
                                                                    }
                                                                }
                                                            } else {
                                                                v14 += (fi45 + fi72)
                                                                switch ("hello" + "\n\t") {
                                                                case "a":
                                                                    switch (fi45 - -582) {
                                                                    case 2147483647:
                                                                        if (!(32768 != (^(fi72)))) {
                                                                            var v82 int64 = 9223372036854775807
                                                                            var v83 int8 = -108
                                                                            if ((true || (v12 > v12)) || false) {
                                                                                var v84 uint8 = 2
                                                                                _ = v84
                                                                            } else {
                                                                                v14 = (v14 + fi45)
                                                                            }
                                                                            _ = v82
                                                                            _ = v83
                                                                        } else {
                                                                            if (v38 <= v42) {
                                                                                v85 := GenS2{F0: (v38 + (v38 * v42))}
                                                                                var v86 uint64 = 18
                                                                                _ = v85
                                                                                _ = v86
                                                                            } else {
                                                                                v87 := GenS4{F0: 216, F1: "foo bar"}
                                                                                _ = v87
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            var v88 uint32 = (v1 ^ 289)
                                                                            _ = v88
                                                                        }
                                                                    default:
                                                                        v14 = (fi45 % ((910) | 1))
                                                                    }
                                                                case "\n\t":
                                                                    var v89 uint64 = 18446744073709551615
                                                                    _ = v89
                                                                }
                                                            }
                                                            if false {
                                                                var v90 uint = 204
                                                                v38 += (v42 + v42)
                                                                _ = v90
                                                            }
                                                            fmt.Println(v1, v12)
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    } else {
                                                        if ("hello" <= (("a" + "foo bar") + "")) {
                                                            fmt.Println(v38, v4, v1)
                                                            v91 := GenS3{F0: 904}
                                                            v5[("\n\t" + "\n\t")] = ""
                                                            _ = v91
                                                        } else {
                                                            v42 += -508
                                                            if ((352 <= 2) || false) {
                                                                if len(v2) > 0 {
                                                                    v2[0] = ("\n\t" + (("foo bar" + ("hello" + ((("foo bar" + ((("a" + "hello") + "foo bar") + "foo bar")) + "\n\t") + "\n\t"))) + "a"))
                                                                }
                                                                v38 -= v42
                                                            }
                                                            v2 = append(v2, ("a" + "foo bar"))
                                                            var v92 float64 = (-2.25 + float64(0.1))
                                                            delete(v3, (("a" + ("" + (("\n\t" + "\n\t") + (("foo bar" + (("hello" + "hello") + "\n\t")) + "")))) + "\n\t"))
                                                            v93 := GenS4{F0: v12, F1: "foo bar"}
                                                            v4 += (0.1 * (1.5 * (v4 / 1e10)))
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                    }
                                                    v5[("" + "")] = ((("a" + ("a" + (((("foo bar" + "hello") + (("foo bar" + "") + "a")) + ("foo bar" + ("foo bar" + (("hello" + "\n\t") + "hello")))) + ""))) + "\n\t") + "\n\t")
                                                    _ = fi72
                                                }
                                            }
                                        } else {
                                            for fi94 := 0; fi94 < 16; fi94++ {
                                                switch 1 {
                                                case -1:
                                                    var v95 int = (^(v14))
                                                    _ = v95
                                                case -200:
                                                    v96 := GenS9{F0: 4294967295, F1: 350, F2: 42}
                                                    _ = v96
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v12 = (-(v12))
                                                switch ("" + "") {
                                                case "a":
                                                    if (!false) {
                                                        fmt.Println(v1, v14)
                                                        v1 = v1
                                                        var v97 int64 = 9223372036854775807
                                                        _ = v97
                                                    }
                                                case "\n\t":
                                                    v98 := GenS10{F0: (^(v12)), F1: 0}
                                                    _ = v98
                                                default:
                                                    if true {
                                                        if true {
                                                            var v99 uint64 = 42
                                                            _ = v99
                                                        }
                                                        fmt.Println(v38, v12)
                                                    } else {
                                                        v42 = 977
                                                    }
                                                }
                                                v1 = (^(v1))
                                                var v100 float64 = 1e10
                                                _ = fi94
                                                _ = v100
                                            }
                                        }
                                        v101 := GenS6{F0: (v1 ^ 1)}
                                        _ = v101
                                        _ = v68
                                    }
                                    if false {
                                        fmt.Println(fi45, v12)
                                        for fi102 := 0; fi102 < 14; fi102++ {
                                            var v103 int16 = v12
                                            var v104 uint64 = 2
                                            fmt.Println(v12, v12, fi102)
                                            _ = fi102
                                            _ = v103
                                            _ = v104
                                        }
                                        v4 = v4
                                    }
                                    v105 := GenS11{F0: (-55 <= -98), F1: false}
                                    _ = v105
                                }
                            case 328:
                                var v106 string = ("\n\t" + ("hello" + "foo bar"))
                                _ = v106
                            }
                        case "\n\t":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    case "a":
                        fmt.Println(v42, fi45)
                    default:
                        if (1000 == 715) {
                            v107 := GenS12{F0: (v42 & int32(fi45)), F1: 0}
                            v4 -= (v4 - v4)
                            v12 = v12
                            _ = v107
                        } else {
                            for fi108 := 0; fi108 < 9; fi108++ {
                                var v109 int = 730
                                v110 := GenS4{F0: v12, F1: ("\n\t" + "foo bar")}
                                switch "" {
                                case "hello":
                                    v111 := GenS13{F0: 42, F1: 32767, F2: 127}
                                    _ = v111
                                case "":
                                    var v112 int64 = -9223372036854775808
                                    _ = v112
                                case "foo bar":
                                    if ((v4 + float32((-2.25 - -2.25))) < 0.1) {
                                        v14 = fi45
                                    } else {
                                        v113 := GenS4{F0: v12, F1: ""}
                                        _ = v113
                                    }
                                }
                                _ = fi108
                                _ = v109
                                _ = v110
                            }
                        }
                    }
                    if ((v14 << ((988) & 63)) != (fi45 | (v14 / (((v14 << ((2) & 63))) | 1)))) {
                        var v114 float32 = ((0.0 - (0.1 / v4)) - ((((((float32(float64((v4 + -2.25))) * 1.5) + (v4 - v4)) - 1e10) / 1e10) + 1e10) / ((v4 + ((float32(0.1) + 1.5) + ((((3.14159 - 0.0) / (float32(0.1) - (1e10 * (float32(0.1) * float32(float64(((1.5 / 3.14159) - v4))))))) + v4) + v4))) / ((1e10 + v4) + 1.5))))
                        _ = v114
                    } else {
                        var v115 int8 = 100
                        var v116 uint64 = 218
                        for fi117 := 0; fi117 < 12; fi117++ {
                            v116 += (v116 % ((v116) | 1))
                            switch 32767 {
                            case -1:
                                fmt.Println(v116, v116)
                            default:
                                v38 += v42
                            }
                            _ = fi117
                        }
                        _ = v115
                        _ = v116
                    }
                    _ = fi45
                }
            default:
                if ((true || true) || (((true && false) || true) && (!false))) {
                    if false {
                        var v118 int64 = 430
                        var v119 int64 = (v118 & (-(v118)))
                        v120 := GenS13{F0: v118, F1: 255, F2: 42}
                        _ = v118
                        _ = v119
                        _ = v120
                    } else {
                        var v121 float32 = ((float32(1.5) / (v4 * float32(0.0))) + v4)
                        _ = v121
                    }
                    var v122 int64 = 0
                    if (!(!(0 >= 18446744073709551615))) {
                        if (true || (42 > 42)) {
                            v5["hello"] = "foo bar"
                            fmt.Println(v42)
                            if (!((1 >= 65535) || true)) {
                                v12 -= -1
                                v123 := v3["a"]
                                _ = v123
                            }
                        } else {
                            if false {
                                if (!((-9223372036854775808 != (v122 | v122)) && false)) {
                                    switch ("" + "") {
                                    case "hello":
                                        if true {
                                            switch (-(v14)) {
                                            case 2147483647:
                                                _ = len(v5)
                                            case 0:
                                                v124 := GenS14{F0: 1, F1: 0}
                                                _ = v124
                                            case -433:
                                                var v125 string = "hello"
                                                _ = v125
                                            default:
                                                fmt.Println(v4, v4)
                                            }
                                        } else {
                                            v126 := v3["\n\t"]
                                            fmt.Println(v4, v4)
                                            _ = v126
                                        }
                                    }
                                    v127 := GenS2{F0: v38}
                                    _ = v127
                                } else {
                                    v128 := GenS15{F0: (false || ((v1 / ((v1) | 1)) == 239)), F1: (^(v1))}
                                    var v129 string = (("foo bar" + ("hello" + ("a" + "\n\t"))) + (("a" + "") + "\n\t"))
                                    var v130 uint = 2
                                    _ = v128
                                    _ = v129
                                    _ = v130
                                }
                                v131 := GenS16{F0: v12, F1: 765, F2: 694}
                                _ = v131
                            }
                        }
                        _ = len(v2)
                        var v132 float32 = float32(float64(v4))
                        _ = v132
                    }
                    _ = v122
                }
            }
        }
        v14 = v14
    } else {
        v133 := GenS17{F0: v42, F1: (^(v1))}
        _ = v133
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v8
}
