package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint; F2 bool }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int16; F1 string; F2 uint64 }
type GenS8 struct { F0 int; F1 string; F2 int32 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 string; F2 string }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint8; F1 int }
type GenS15 struct { F0 int8 }
func genFn0(p0 uint32, p1 int32, p2 int8) uint8 {
    fmt.Println(p2)
    if true {
        v7 := []int{128, -2147483648}
        v7 = append(v7, 848)
        _ = v7
    } else {
        fmt.Println(p0, p1)
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int = fi8
            if (2 > 255) {
                fmt.Println(fi8, p1, p0)
                if (42 != 2) {
                    v10 := GenS1{F0: 42}
                    _ = v10
                } else {
                    v11 := []int{66, v9}
                    var v12 uint64 = 2
                    v13 := GenS2{F0: p0, F1: ("\n\t" + ("\n\t" + (("foo bar" + "\n\t") + "hello"))), F2: "hello"}
                    _ = v11
                    _ = v12
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v15 := GenS3{F0: p2, F1: false, F2: 128}
                    var v16 int64 = 0
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                fmt.Println(fi8)
                switch 127 {
                case 42:
                    if (-2.25 == 0.1) {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v9 += (v9 * -228)
                            _ = fi17
                        }
                    } else {
                        v18 := GenS2{F0: (^(p0)), F1: (("hello" + "hello") + ""), F2: "\n\t"}
                        v9 += (fi8 + fi8)
                        if true {
                            var v19 int = (fi8 | (v9 & v9))
                            _ = v19
                        }
                        _ = v18
                    }
                default:
                    v20 := []int{v9, 9223372036854775807, fi8}
                    _ = v20
                }
            }
            _ = fi8
            _ = v9
        }
    }
    p0 = p0
    if true {
        v21 := []int{1}
        _ = len(v21)
        var v22 float32 = 0.0
        _ = v21
        _ = v22
    } else {
        if (1 <= 42) {
            var v23 int8 = p2
            switch "a" {
            case "a":
                var v24 int64 = 628
                _ = v24
            case "":
                for fi25 := 0; fi25 < 15; fi25++ {
                    if true {
                        for fi26 := 0; fi26 < 4; fi26++ {
                            var v27 uint32 = (^(p0))
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        var v28 int8 = 9
                        _ = v28
                    }
                    _ = fi25
                }
            }
            _ = v23
        } else {
            if false {
                var v29 int8 = (-(p2))
                _ = v29
            } else {
                v30 := []int{127}
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                _ = v30
            }
            if (!(2147483647 >= 32767)) {
                var v31 uint8 = 215
                v32 := []int{42}
                _ = v31
                _ = v32
            }
        }
        if false {
            var v33 int32 = p1
            if ((p1 ^ 1) != (-(p1))) {
                if false {
                    var v34 int64 = -934
                    v35 := GenS3{F0: p2, F1: false, F2: 2}
                    _ = v34
                    _ = v35
                } else {
                    v36 := []int{42, -511, -1}
                    _ = v36
                }
            }
            _ = v33
        }
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []string{("\n\t" + ""), ("" + "\n\t")}
    var v2 uint = 2
    var v3 int16 = 42
    var v4 int = 127
    v5 := []int{(v4 / ((4294967295) | 1)), v4}
    v6 := make(map[string]string)
    if false {
        for fi37 := 0; fi37 < 17; fi37++ {
            var v38 uint64 = 4294967295
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 18; fi39++ {
            v40 := GenS1{F0: 841}
            _ = fi39
            _ = v40
        }
    }
    v2 = v2
    for fi41 := 0; fi41 < 15; fi41++ {
        v2 -= 65535
        v42 := GenS4{F0: 2147483648, F1: -122}
        if (!(139 >= int64(fi41))) {
            fmt.Println(v2, fi41, v4)
            if false {
                v4 += (^(v4))
                v43 := GenS1{F0: 0}
                v4 += 127
                _ = v43
            }
        }
        _ = fi41
        _ = v42
    }
    var v44 string = (("\n\t" + "foo bar") + "\n\t")
    v2 = (^(v2))
    if (2 < 32) {
        v45 := GenS2{F0: 196, F1: (v44 + v44), F2: (v44 + ("hello" + v44))}
        for fi46 := 0; fi46 < 10; fi46++ {
            if ((v4 >= fi46) || (42 > 127)) {
                if (106 != 108) {
                    if false {
                        for fi47 := 0; fi47 < 20; fi47++ {
                            v48 := GenS2{F0: 885, F1: "\n\t", F2: "hello"}
                            _ = fi47
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 6; fi49++ {
                            switch "" {
                            case "":
                                var v50 uint64 = 1
                                _ = v50
                            case "a":
                                v2 -= v2
                            case "\n\t":
                                v4 = fi46
                            default:
                                v51 := GenS5{F0: 65535, F1: v2, F2: true}
                                _ = v51
                            }
                            _ = len(v5)
                            _ = fi49
                        }
                    }
                    for fi52 := 0; fi52 < 3; fi52++ {
                        var v53 string = ""
                        var v54 bool = true
                        if (!(v4 == (fi46 << (((v2 << (((v2 & 2)) & 63))) & 63)))) {
                            switch (fi52 + -1) {
                            case 4294967295:
                                v55 := GenS6{F0: 304, F1: 53}
                                _ = v55
                            default:
                                if false {
                                    if true {
                                        fmt.Println(v2)
                                        fmt.Println(v53)
                                    }
                                    var v56 bool = (!(!(!((0.0 / (((1e10 - ((1.5 - 1e10) + (3.14159 + 0.0))) - 0.0) / ((float32(0.1) + ((((1.5 - -2.25) + 0.0) * 1e10) / 1.0)) - -2.25))) != (-2.25 / (1.5 - 3.14159))))))
                                    var v57 int64 = 697
                                    _ = v56
                                    _ = v57
                                } else {
                                    _ = len(v6)
                                    switch "" {
                                    case "":
                                        switch (v44 + ("hello" + v44)) {
                                        case "":
                                            v53 = ((v44 + "") + v44)
                                        }
                                    default:
                                        if v54 {
                                            var v58 string = v53
                                            switch (fi46 - (-(fi46))) {
                                            case 32768:
                                                v3 = (v3 | v3)
                                            case 388:
                                                var v59 int32 = 0
                                                _ = v59
                                            case 127:
                                                v2 -= 968
                                            }
                                            var v60 bool = ((0.0 - 1.5) > (((float32(-2.25) - 0.1) - -2.25) / 1e10))
                                            _ = v58
                                            _ = v60
                                        }
                                    }
                                }
                            }
                            switch v53 {
                            case "hello":
                                v6["hello"] = (("" + ("foo bar" + "hello")) + v53)
                            case "a":
                                var v61 int32 = -65
                                _ = v61
                            case "\n\t":
                                var v62 uint8 = 128
                                _ = v62
                            }
                        } else {
                            v2 += v2
                        }
                        fmt.Println(v44, v44, v44)
                        var v63 float64 = (-2.25 - 0.0)
                        _ = fi52
                        _ = v53
                        _ = v54
                        _ = v63
                    }
                    var v64 int64 = 9223372036854775807
                    _ = v64
                } else {
                    if false {
                        if false {
                            v65 := v6[v44]
                            _ = v65
                        }
                        var v66 uint8 = uint8(v2)
                        _ = v66
                    } else {
                        var v67 float64 = -2.25
                        if (255 != 255) {
                            var v68 float64 = v67
                            if true {
                                switch v44 {
                                case "foo bar":
                                    v69 := GenS6{F0: 332, F1: 2}
                                    _ = v69
                                default:
                                    v44 = ("hello" + (((v44 + ("a" + v44)) + "hello") + ((v44 + ("foo bar" + ((("hello" + "\n\t") + (v44 + v44)) + (v44 + v44)))) + "foo bar")))
                                }
                                switch fi46 {
                                case 0:
                                    if false {
                                        var v70 uint16 = 2
                                        _ = v70
                                    } else {
                                        for fi71 := 0; fi71 < 5; fi71++ {
                                            var v72 int32 = -1
                                            var v73 int = 128
                                            _ = fi71
                                            _ = v72
                                            _ = v73
                                        }
                                        switch -814 {
                                        case 1:
                                            var v74 int64 = 9223372036854775807
                                            _ = v74
                                        case 32768:
                                            v75 := v6[((v44 + v44) + ((v44 + "foo bar") + (v44 + v44)))]
                                            _ = v75
                                        default:
                                            switch "" {
                                            case "a":
                                                v76 := GenS6{F0: 4294967295, F1: 2}
                                                _ = v76
                                            case "foo bar":
                                                v2 = (v2 & (^(v2)))
                                            case "":
                                                if (2147483647 <= -347) {
                                                    v44 = ("hello" + (((v44 + v44) + ("\n\t" + "foo bar")) + v44))
                                                    switch (((v44 + (("\n\t" + "\n\t") + (((v44 + v44) + (v44 + v44)) + ("a" + (((v44 + "a") + ("\n\t" + "\n\t")) + v44))))) + (v44 + v44)) + ("hello" + v44)) {
                                                    case "":
                                                        if false {
                                                            fmt.Println(v4, v2, v4)
                                                            v2 += (^(v2))
                                                        }
                                                    case "foo bar":
                                                        for fi77 := 0; fi77 < 17; fi77++ {
                                                            fmt.Println(v67, v44, v4)
                                                            if false {
                                                                if (1 > 228) {
                                                                    fmt.Println(v68)
                                                                    switch ("hello" + (v44 + v44)) {
                                                                    case "hello":
                                                                        v6[("a" + v44)] = (("\n\t" + (((v44 + (((v44 + v44) + ("a" + "a")) + ("\n\t" + v44))) + "a") + "a")) + "\n\t")
                                                                    case "\n\t":
                                                                        var v78 string = (v44 + "a")
                                                                        _ = v78
                                                                    case "":
                                                                        var v79 string = "\n\t"
                                                                        switch fi77 {
                                                                        case 757:
                                                                            var v80 int = 128
                                                                            _ = v80
                                                                        case 42:
                                                                            var v81 uint8 = 112
                                                                            _ = v81
                                                                        case 32767:
                                                                            if (34 > 127) {
                                                                                var v82 uint32 = 733
                                                                                _ = v82
                                                                            }
                                                                        default:
                                                                            if false {
                                                                                delete(v6, (v79 + v44))
                                                                                switch (v4 / (((v4 >> ((255) & 63))) | 1)) {
                                                                                case -376:
                                                                                    v83 := GenS6{F0: 0, F1: 2147483647}
                                                                                    _ = v83
                                                                                case -804:
                                                                                    if false {
                                                                                        if true {
                                                                                            v6[(("foo bar" + v79) + (("\n\t" + (v44 + v44)) + v44))] = ("a" + (v44 + v44))
                                                                                            v3 += int16(v4)
                                                                                        }
                                                                                    } else {
                                                                                        switch v44 {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v67)
                                                                                        }
                                                                                        v3 = v3
                                                                                        var v84 int = fi77
                                                                                        _ = v84
                                                                                    }
                                                                                case -395:
                                                                                    if (false && (!false)) {
                                                                                        if (v68 >= v68) {
                                                                                            v85 := GenS7{F0: v3, F1: v44, F2: 289}
                                                                                            if (!false) {
                                                                                                switch (fi77 / (((fi77 + fi77)) | 1)) {
                                                                                                case 0:
                                                                                                    var v86 uint64 = 0
                                                                                                    _ = v86
                                                                                                case 1:
                                                                                                    switch (v79 + ((v79 + ("a" + v79)) + v44)) {
                                                                                                    case "hello":
                                                                                                        v68 -= float64(1.5)
                                                                                                    }
                                                                                                case 32767:
                                                                                                    var v87 int64 = 400
                                                                                                    _ = v87
                                                                                                }
                                                                                            } else {
                                                                                                v88 := v6[""]
                                                                                                _ = v88
                                                                                            }
                                                                                            v67 += float64(float32(1e10))
                                                                                            v89 := GenS3{F0: -1, F1: ((^(v3)) != 42), F2: genFn0(2, 42, -1)}
                                                                                            _ = v85
                                                                                            _ = v89
                                                                                        } else {
                                                                                            v2 -= 3
                                                                                            var v90 uint = (v2 + (v2 + v2))
                                                                                            _ = v90
                                                                                        }
                                                                                    } else {
                                                                                        v2 = v2
                                                                                    }
                                                                                default:
                                                                                    var v91 float32 = -2.25
                                                                                    _ = v91
                                                                                }
                                                                            } else {
                                                                                var v92 uint = (v2 >> (((v2 & 983)) & 63))
                                                                                var v93 uint8 = 128
                                                                                _ = v92
                                                                                _ = v93
                                                                            }
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                    fmt.Println(fi46, v68)
                                                                    v94 := GenS8{F0: (fi77 | 9223372036854775807), F1: ((v44 + v44) + (v44 + v44)), F2: 783}
                                                                    _ = v94
                                                                } else {
                                                                    var v95 uint32 = uint32(v2)
                                                                    _ = v95
                                                                }
                                                            } else {
                                                                v68 = v67
                                                                if (false && (1 == 42)) {
                                                                    var v96 float32 = float32(0.0)
                                                                    if true {
                                                                        var v97 uint = v2
                                                                        fmt.Println(v4, v44)
                                                                        v98 := GenS9{F0: 42, F1: (v3 * v3)}
                                                                        v99 := GenS10{F0: v3}
                                                                        _ = len(v5)
                                                                        _ = v97
                                                                        _ = v98
                                                                        _ = v99
                                                                    } else {
                                                                        _ = len(v1)
                                                                        var v100 uint32 = 65535
                                                                        var v101 int8 = 100
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                    v2 = v2
                                                                    _ = v96
                                                                } else {
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (fi77 + 127)
                                                                    }
                                                                    v102 := GenS1{F0: 1}
                                                                    if (!(!true)) {
                                                                        v44 = ""
                                                                        fmt.Println(v67, v2, v2)
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "\n\t"
                                                                        }
                                                                        var v103 uint16 = 65535
                                                                        v104 := GenS2{F0: 42, F1: v44, F2: (v44 + "a")}
                                                                        _ = v103
                                                                        _ = v104
                                                                    }
                                                                    _ = v102
                                                                }
                                                                v105 := GenS2{F0: 552, F1: "\n\t", F2: ("a" + v44)}
                                                                _ = v105
                                                            }
                                                            if true {
                                                                var v106 float64 = 1.5
                                                                v107 := GenS11{F0: v3, F1: ("\n\t" + (v44 + "hello")), F2: v44}
                                                                _ = v106
                                                                _ = v107
                                                            } else {
                                                                var v108 uint = (v2 & (v2 << (((v2 / (((v2 * v2)) | 1))) & 63)))
                                                                switch v4 {
                                                                case 2147483647:
                                                                    v68 = v67
                                                                    var v109 bool = (((((1.5 * v68) + (3.14159 / v68)) - ((1e10 / v67) - v67)) <= v68) && (("hello" <= v44) || ((float32(v68) == (-2.25 - 0.1)) || true)))
                                                                    _ = v109
                                                                case -778:
                                                                    v110 := GenS8{F0: (fi46 << ((709) & 63)), F1: (v44 + v44), F2: -2147483648}
                                                                    _ = v110
                                                                default:
                                                                    var v111 int16 = v3
                                                                    _ = v111
                                                                }
                                                                _ = v108
                                                            }
                                                            _ = fi77
                                                        }
                                                    default:
                                                        if false {
                                                            fmt.Println(v44, v3, v2)
                                                            for fi112 := 0; fi112 < 11; fi112++ {
                                                                if false {
                                                                    fmt.Println(v3, fi46, v68)
                                                                    v67 += v67
                                                                    v113 := GenS7{F0: v3, F1: ((v44 + v44) + "a"), F2: 1}
                                                                    if true {
                                                                        var v114 int8 = -1
                                                                        v115 := GenS4{F0: 2147483647, F1: -128}
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                    _ = v113
                                                                } else {
                                                                    var v116 uint8 = 161
                                                                    _ = v116
                                                                }
                                                                var v117 uint16 = 475
                                                                v118 := GenS7{F0: 1, F1: (v44 + v44), F2: 4294967295}
                                                                _ = fi112
                                                                _ = v117
                                                                _ = v118
                                                            }
                                                            var v119 string = v44
                                                            _ = v119
                                                        }
                                                    }
                                                    var v120 int16 = 727
                                                    _ = v120
                                                }
                                            default:
                                                v67 = -2.25
                                            }
                                        }
                                        v2 = 65535
                                    }
                                }
                                for fi121 := 0; fi121 < 5; fi121++ {
                                    v44 = "a"
                                    var v122 uint32 = 313
                                    var v123 string = ("\n\t" + v44)
                                    v5 = append(v5, 32768)
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    v6[v44] = v44
                                    _ = fi121
                                    _ = v122
                                    _ = v123
                                }
                            } else {
                                var v124 int8 = 42
                                _ = v124
                            }
                            _ = v68
                        }
                        _ = v67
                    }
                }
                var v125 int8 = -128
                if true {
                    switch v4 {
                    case 294:
                        _ = len(v6)
                    default:
                        v4 = (^(fi46))
                    }
                } else {
                    switch (v44 + ("a" + v44)) {
                    case "a":
                        v126 := GenS2{F0: uint32(v2), F1: ("foo bar" + ("a" + "\n\t")), F2: "hello"}
                        _ = v126
                    case "hello":
                        var v127 int = (fi46 - (v4 / (((v4 << ((v2) & 63))) | 1)))
                        _ = v127
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    switch "" {
                    case "hello":
                        if (0 < 0) {
                            v4 = (-(v4))
                            var v128 uint32 = 1
                            _ = v128
                        }
                    default:
                        v129 := GenS1{F0: 4294967295}
                        _ = v129
                    }
                }
                _ = v125
            }
            _ = fi46
        }
        var v130 uint8 = 102
        _ = v130
        _ = v45
    }
    fmt.Println(v2, v2, v2)
    if len(v5) > 0 {
        v5[0] = v4
    }
    _ = len(v1)
    v131 := GenS12{F0: v2, F1: 255}
    var v132 float32 = 3.14159
    v133 := GenS13{F0: v2}
    if ((!true) || true) {
        fmt.Println(v4, v2)
        fmt.Println(v132, v132, v4)
    } else {
        var v134 uint = v2
        _ = v134
    }
    for fi135 := 0; fi135 < 8; fi135++ {
        var v136 int8 = 0
        v136 = 0
        _ = fi135
        _ = v136
    }
    for fi137 := 0; fi137 < 8; fi137++ {
        if len(v1) > 0 {
            v1[0] = v44
        }
        _ = len(v6)
        v138 := GenS10{F0: -1}
        _ = fi137
        _ = v138
    }
    v139 := GenS14{F0: 42, F1: (v4 << (((v2 / (((v2 + 42)) | 1))) & 63))}
    switch "" {
    case "\n\t":
        v140 := GenS15{F0: 59}
        _ = v140
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v131
    _ = v132
    _ = v133
    _ = v139
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
