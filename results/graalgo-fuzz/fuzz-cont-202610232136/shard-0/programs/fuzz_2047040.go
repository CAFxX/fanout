package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 int16; F1 string; F2 uint8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 int; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int16; F1 uint8; F2 int16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint16; F1 int16; F2 int16 }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 int16 }
func genFn0(p0 int, p1 uint) uint {
    v5 := []int{(p0 >> ((410) & 63))}
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return (p1 << ((p1) & 63))
}
func genFn1(p0 uint8, p1 int64, p2 string) int {
    var v6 int64 = -755
    v7 := []int{-908}
    fmt.Println(p0)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 float32 = ((0.0 - (1e10 / float32((float64(1e10) - (-2.25 + ((1.5 + (-2.25 - 1.5)) / 3.14159)))))) * (float32((0.0 - -2.25)) + 0.1))
    var v2 float32 = (((v1 - (v1 * v1)) * (v1 - (1.5 * float32(3.14159)))) + v1)
    var v3 int32 = 2147483647
    var v4 uint = 2
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: -11, F1: "\n\t"}
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: 100, F1: ("" + (("foo bar" + "a") + "foo bar"))}
            fmt.Println(v1, v2, v2)
            fmt.Println(v4, v1, v1)
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := GenS2{F0: genFn1(252, 1, "a"), F1: 205, F2: -1}
    v2 = (((v2 * v1) - float32(3.14159)) / (float32(-2.25) * ((((0.0 + (v2 + float32(((((3.14159 + 1e10) * ((float64(v2) - (0.1 * 1.5)) * (0.1 - ((1e10 + 1e10) - (3.14159 + ((1e10 * (((1.5 * 1e10) * (-2.25 / (((0.0 * 3.14159) - float64(1.5)) + 3.14159))) + (3.14159 / -2.25))) + 1.5)))))) / 1.0) * (0.0 / ((float64(((v2 - v2) - v2)) - (((0.1 - float64(v2)) + (0.0 - (0.0 + 3.14159))) * -2.25)) * (((0.1 - float64((float32(3.14159) + (0.1 / v2)))) - -2.25) * (-2.25 - 1.5)))))))) - (1.5 * (0.0 - v2))) / 3.14159) * v2)))
    if false {
        fmt.Println(v3, v3)
        v4 = (v4 ^ v4)
    }
    if false {
        var v13 uint64 = 46
        _ = v13
    } else {
        v3 = (v3 + v3)
        fmt.Println(v4, v3)
    }
    if true {
        var v14 float32 = v2
        v1 -= v14
        _ = v14
    }
    if true {
        v4 = (^(v4))
        if (0 <= 0) {
            v15 := []int{}
            v16 := GenS1{F0: -128, F1: ""}
            _ = v15
            _ = v16
        } else {
            if (true || (0 >= 2)) {
                var v17 uint = genFn0(genFn1(2, -675, "hello"), genFn0(42, (v4 + 339)))
                _ = v17
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v4 = (v4 % (((v4 % ((65535) | 1))) | 1))
                    var v19 uint64 = 4294967295
                    switch (fi18 & (fi18 & fi18)) {
                    case 9223372036854775807:
                        v1 = v2
                    case -858:
                        var v20 int = fi18
                        _ = v20
                    case -276:
                        v19 = 4294967295
                    default:
                        var v21 uint64 = (v19 - 4294967295)
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
                var v22 float64 = 0.0
                _ = v22
            }
            v23 := GenS2{F0: -1, F1: 117, F2: -32768}
            _ = v23
        }
        if (true || (230 <= 0)) {
            for fi24 := 0; fi24 < 16; fi24++ {
                switch (fi24 & genFn1(128, 1, "hello")) {
                case -9223372036854775808:
                    switch genFn1(0, 736, "a") {
                    case 4294967295:
                        for fi25 := 0; fi25 < 1; fi25++ {
                            v26 := []int{}
                            v2 = -2.25
                            var v27 int16 = 878
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                case 32768:
                    var v28 uint8 = 1
                    _ = v28
                case 9223372036854775807:
                    if (false || true) {
                        for fi29 := 0; fi29 < 20; fi29++ {
                            if false {
                                if (!true) {
                                    v30 := GenS3{F0: 410, F1: (("foo bar" + "a") + ("a" + "")), F2: 223}
                                    v3 += (v3 * v3)
                                    _ = v30
                                }
                                var v31 string = ("foo bar" + (("foo bar" + "hello") + ""))
                                var v32 uint = 835
                                if true {
                                    v33 := []int{4294967295, (fi29 & fi29), fi24, (fi29 | (fi24 >> ((v4) & 63)))}
                                    _ = len(v33)
                                    v3 = (v3 & (v3 ^ v3))
                                    _ = v33
                                }
                                _ = v31
                                _ = v32
                            }
                            v2 = (0.0 / float32(1.5))
                            v34 := []int{(fi29 * fi29), fi24, (fi29 * fi24), fi24}
                            _ = fi29
                            _ = v34
                        }
                        v35 := GenS4{F0: (v4 | v4), F1: -748}
                        _ = v35
                    }
                default:
                    var v36 uint16 = 462
                    _ = v36
                }
                for fi37 := 0; fi37 < 14; fi37++ {
                    if (false || (0 <= 0)) {
                        switch ("" + (("a" + ("" + "hello")) + "foo bar")) {
                        case "foo bar":
                            v4 = 4294967295
                        case "\n\t":
                            if false {
                                if (522 != 65535) {
                                    fmt.Println(v3, v3)
                                    v38 := []int{(fi24 & -9223372036854775808), (fi37 + 128)}
                                    var v39 uint8 = 248
                                    fmt.Println(fi37, v2, v2)
                                    fmt.Println(v3)
                                    v38 = append(v38, -275)
                                    _ = v38
                                    _ = v39
                                }
                                v40 := []int{fi24}
                                _ = v40
                            }
                        }
                    } else {
                        v41 := []int{(fi37 ^ 1), (fi37 >> (((v4 * v4)) & 63)), (fi24 * fi37)}
                        if ((((3.14159 - 1.5) + 1e10) + 1e10) < ((1.5 - 1.5) + (float64(float32(0.1)) / ((0.0 - 1e10) / 1.0)))) {
                            var v42 uint16 = 1
                            v43 := GenS5{F0: fi24, F1: (^(v42))}
                            _ = v42
                            _ = v43
                        } else {
                            var v44 float64 = 1.5
                            _ = v44
                        }
                        _ = v41
                    }
                    _ = fi37
                }
                _ = fi24
            }
            for fi45 := 0; fi45 < 15; fi45++ {
                v46 := []int{-829, fi45}
                if ((!(-346 != -9223372036854775808)) || (!false)) {
                    if len(v46) >= 2 {
                        v46 = v46[1:len(v46)]
                    }
                    for fi47 := 0; fi47 < 5; fi47++ {
                        var v48 bool = true
                        switch "foo bar" {
                        case "hello":
                            if v48 {
                                v49 := GenS6{F0: -679, F1: 42, F2: 1}
                                if v48 {
                                    if v48 {
                                        if true {
                                            var v50 int16 = 1
                                            _ = v50
                                        } else {
                                            if true {
                                                switch genFn1(128, 42, "a") {
                                                case -734:
                                                    _ = len(v46)
                                                }
                                                fmt.Println(v2)
                                                v2 -= ((v1 / v1) - float32(-2.25))
                                            } else {
                                                v48 = false
                                                if len(v46) > 0 {
                                                    v46[0] = genFn1(0, 42, ((("foo bar" + ("hello" + "hello")) + "hello") + "a"))
                                                }
                                                var v51 int64 = -9223372036854775808
                                                _ = v51
                                            }
                                        }
                                        v2 = 1.5
                                        var v52 uint64 = 164
                                        _ = v52
                                    } else {
                                        var v53 bool = true
                                        if len(v46) >= 2 {
                                            v46 = v46[1:len(v46)]
                                        }
                                        _ = v53
                                    }
                                } else {
                                    var v54 uint16 = 213
                                    _ = v54
                                }
                                switch fi45 {
                                case -9223372036854775808:
                                    if (fi47 != 32768) {
                                        if (v48 || true) {
                                            if len(v46) > 0 {
                                                v46[0] = (^(fi47))
                                            }
                                            var v55 bool = (1.5 == ((3.14159 - 0.1) + float64(v2)))
                                            var v56 uint = (v4 ^ 255)
                                            _ = v55
                                            _ = v56
                                        }
                                        if len(v46) >= 2 {
                                            v46 = v46[1:len(v46)]
                                        }
                                    } else {
                                        switch (fi47 & genFn1(2, 37, "a")) {
                                        case 32768:
                                            var v57 uint8 = 42
                                            _ = v57
                                        }
                                        switch (fi47 << ((v4) & 63)) {
                                        case 433:
                                            v48 = v48
                                        case -9223372036854775808:
                                            v58 := GenS5{F0: fi47, F1: 32767}
                                            _ = v58
                                        }
                                    }
                                case 32768:
                                    v3 += v3
                                case -514:
                                    var v59 uint64 = 680
                                    _ = v59
                                }
                                _ = v49
                            }
                        case "foo bar":
                            var v60 float64 = ((((-2.25 + (0.1 * 1.5)) - 0.0) + 0.1) + (((1e10 - 3.14159) * 3.14159) + float64(v1)))
                            _ = v60
                        case "":
                            switch (fi47 % ((fi45) | 1)) {
                            case 1:
                                if false {
                                    var v61 int8 = 1
                                    _ = v61
                                }
                            case 719:
                                if true {
                                    _ = len(v46)
                                    v48 = (-970 >= fi45)
                                } else {
                                    if (1 >= 42) {
                                        _ = len(v46)
                                        var v62 float32 = v2
                                        if (2 < 108) {
                                            var v63 uint = (v4 - v4)
                                            if len(v46) >= 2 {
                                                v46 = v46[1:len(v46)]
                                            }
                                            v3 = (v3 % ((2147483647) | 1))
                                            switch ("hello" + ("\n\t" + ("" + ""))) {
                                            case "hello":
                                                if v48 {
                                                    if true {
                                                        if v48 {
                                                            v64 := GenS7{F0: 255, F1: fi45}
                                                            v63 = (v4 & 65535)
                                                            _ = v64
                                                        } else {
                                                            var v65 bool = false
                                                            switch "" {
                                                            case "a":
                                                                if len(v46) > 0 {
                                                                    v46[0] = (fi47 ^ 1)
                                                                }
                                                            case "hello":
                                                                v66 := GenS8{F0: 1, F1: 84, F2: -945}
                                                                _ = v66
                                                            case "\n\t":
                                                                v46 = append(v46, -946)
                                                            default:
                                                                fmt.Println(v3, v3, v4)
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                if false {
                                                                    fmt.Println(v2, v63, v65)
                                                                    var v67 bool = (!v65)
                                                                    _ = v67
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(fi45, v63, v3)
                                                            case "a":
                                                                v68 := GenS9{F0: 65535}
                                                                _ = v68
                                                            default:
                                                                if len(v46) > 0 {
                                                                    v46[0] = fi45
                                                                }
                                                            }
                                                            _ = v65
                                                        }
                                                        v48 = (0.1 > 1.5)
                                                        var v69 int16 = 1000
                                                        var v70 int64 = 9223372036854775807
                                                        v62 += 0.0
                                                        _ = v69
                                                        _ = v70
                                                    } else {
                                                        v3 -= (v3 / ((-2147483648) | 1))
                                                        var v71 uint16 = 628
                                                        fmt.Println(v62, v48)
                                                        _ = v71
                                                    }
                                                    fmt.Println(v3, v62)
                                                    var v72 uint8 = 223
                                                    _ = v72
                                                } else {
                                                    if (-398 != 0) {
                                                        if true {
                                                            var v73 uint8 = 92
                                                            var v74 string = "hello"
                                                            _ = v73
                                                            _ = v74
                                                        } else {
                                                            v62 -= (v62 / 1.5)
                                                        }
                                                        fmt.Println(v63)
                                                        var v75 uint16 = 0
                                                        var v76 uint = v4
                                                        v2 = ((v1 - v62) - v62)
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        v1 = float32((((1e10 * -2.25) * 1.5) - (-2.25 - 0.1)))
                                                        var v77 float32 = (1e10 + 0.1)
                                                        _ = v77
                                                    }
                                                    if v48 {
                                                        var v78 int8 = -64
                                                        if (!(!v48)) {
                                                            if v48 {
                                                                var v79 float32 = 0.0
                                                                var v80 int16 = 1
                                                                fmt.Println(v62)
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                var v81 int = genFn1(255, -237, ("hello" + (("" + (("" + "hello") + (("" + ("a" + "")) + "hello"))) + ("" + ""))))
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if len(v46) > 0 {
                                                                        v46[0] = (v81 >> (((v4 * 298)) & 63))
                                                                    }
                                                                case "a":
                                                                    var v82 int32 = (v3 >> (((v63 | (^(v63)))) & 63))
                                                                    _ = v82
                                                                default:
                                                                    var v83 int8 = 127
                                                                    var v84 float64 = -2.25
                                                                    switch "hello" {
                                                                    case "":
                                                                        if v48 {
                                                                            var v85 int16 = 0
                                                                            if false {
                                                                                if len(v46) >= 2 {
                                                                                    v46 = v46[1:len(v46)]
                                                                                }
                                                                                if len(v46) > 0 {
                                                                                    v46[0] = genFn1(uint8(v4), 1, "hello")
                                                                                }
                                                                                var v86 int32 = (^(v3))
                                                                                _ = v86
                                                                            } else {
                                                                                v87 := GenS2{F0: genFn1(255, 1, ("" + ("foo bar" + ""))), F1: 1, F2: (v85 - (v85 | (v85 / (((v85 * 916)) | 1))))}
                                                                                v88 := GenS10{F0: "\n\t"}
                                                                                _ = v87
                                                                                _ = v88
                                                                            }
                                                                            _ = v85
                                                                        } else {
                                                                            var v89 float64 = ((0.0 / (v84 * ((0.1 - v84) * v84))) - ((v84 - v84) / 1e10))
                                                                            _ = v89
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(v4)
                                                                    default:
                                                                        v83 = v78
                                                                    }
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                                var v90 int64 = 0
                                                                _ = v81
                                                                _ = v90
                                                            }
                                                        }
                                                        var v91 int8 = v78
                                                        _ = v78
                                                        _ = v91
                                                    } else {
                                                        var v92 float64 = (0.0 / -2.25)
                                                        _ = v92
                                                    }
                                                    if (827 < 1) {
                                                        switch (fi45 + -87) {
                                                        case 0:
                                                            fmt.Println(v1, v63)
                                                        case 130:
                                                            v2 = (v2 * v62)
                                                        default:
                                                            v93 := GenS4{F0: genFn0(fi47, (^(v4))), F1: -464}
                                                            _ = v93
                                                        }
                                                        var v94 uint8 = 2
                                                        _ = v94
                                                    }
                                                }
                                            case "a":
                                                fmt.Println(v1, v62)
                                            case "\n\t":
                                                fmt.Println(fi47, fi45)
                                            }
                                            _ = v63
                                        }
                                        _ = v62
                                    } else {
                                        v4 = v4
                                    }
                                    var v95 uint = (v4 << ((v4) & 63))
                                    _ = v95
                                }
                            case 322:
                                switch (^(fi47)) {
                                case 1:
                                    switch 812 {
                                    case 127:
                                        if len(v46) > 0 {
                                            v46[0] = genFn1(42, -315, "\n\t")
                                        }
                                    case -9223372036854775808:
                                        if len(v46) > 0 {
                                            v46[0] = int(v3)
                                        }
                                    case 2147483647:
                                        v96 := GenS10{F0: "\n\t"}
                                        _ = v96
                                    }
                                case -608:
                                    fmt.Println(v1)
                                default:
                                    switch "" {
                                    case "a":
                                        if (-1 >= 42) {
                                            v2 += 0.0
                                        } else {
                                            fmt.Println(v3, v2)
                                        }
                                    }
                                }
                                switch "" {
                                case "hello":
                                    var v97 int16 = 32767
                                    _ = v97
                                case "foo bar":
                                    var v98 int8 = -78
                                    _ = v98
                                default:
                                    var v99 float32 = (v1 / v2)
                                    _ = v99
                                }
                            default:
                                var v100 uint = 18446744073709551615
                                _ = v100
                            }
                        default:
                            if len(v46) > 0 {
                                v46[0] = fi45
                            }
                        }
                        _ = fi47
                        _ = v48
                    }
                    if ((!(!(50 == 1))) && (((3.14159 / 3.14159) == (0.1 / 0.1)) && true)) {
                        switch "a" {
                        case "foo bar":
                            var v101 float64 = 1.5
                            _ = v101
                        case "hello":
                            switch ("a" + "a") {
                            case "hello":
                                v102 := GenS11{F0: 0, F1: 1000, F2: 464}
                                _ = v102
                            case "foo bar":
                                if true {
                                    v1 -= v2
                                }
                            case "":
                                v1 = float32((0.1 / 1.5))
                            }
                        default:
                            var v103 uint32 = 4294967295
                            _ = v103
                        }
                        for fi104 := 0; fi104 < 10; fi104++ {
                            fmt.Println(v3, v2, fi45)
                            _ = fi104
                        }
                        if true {
                            for fi105 := 0; fi105 < 3; fi105++ {
                                v4 = (v4 | (^(v4)))
                                v106 := GenS1{F0: 127, F1: ""}
                                _ = fi105
                                _ = v106
                            }
                            var v107 float64 = 0.1
                            if (551 <= 9223372036854775807) {
                                if (127 > -128) {
                                    for fi108 := 0; fi108 < 18; fi108++ {
                                        var v109 uint32 = 42
                                        v110 := GenS3{F0: 32767, F1: "foo bar", F2: 30}
                                        v1 += (3.14159 - ((v2 - (3.14159 * (-2.25 - 1.5))) + v1))
                                        var v111 uint = v4
                                        v1 -= (float32(0.1) - float32(v107))
                                        _ = fi108
                                        _ = v109
                                        _ = v110
                                        _ = v111
                                    }
                                }
                                if true {
                                    var v112 int16 = 0
                                    _ = v112
                                }
                            } else {
                                fmt.Println(v1, v4)
                                v113 := GenS7{F0: 128, F1: fi45}
                                _ = v113
                            }
                            _ = v107
                        } else {
                            if (!true) {
                                for fi114 := 0; fi114 < 2; fi114++ {
                                    if len(v46) >= 2 {
                                        v46 = v46[1:len(v46)]
                                    }
                                    v3 = (v3 ^ (v3 + -2147483648))
                                    _ = fi114
                                }
                            } else {
                                var v115 float32 = v2
                                fmt.Println(v1, v115)
                                if (v3 > -814) {
                                    fmt.Println(v2, v115)
                                }
                                _ = v115
                            }
                            var v116 int8 = 42
                            v1 = float32(3.14159)
                            _ = v116
                        }
                    } else {
                        v117 := GenS12{F0: fi45, F1: 1}
                        _ = v117
                    }
                }
                v46 = append(v46, fi45)
                _ = fi45
                _ = v46
            }
            v118 := GenS13{F0: -32768}
            _ = v118
        }
    } else {
        var v119 uint = (v4 % ((212) | 1))
        _ = v119
    }
    v3 += (v3 / ((v3) | 1))
    var v120 uint16 = 441
    v121 := []int{2147483647, -2147483648, 127}
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("slicelen v121:", len(v121), v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
}
