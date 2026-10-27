package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int64; F1 int8; F2 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 uint16; F1 uint64; F2 int }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 string; F1 uint16 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 string; F1 uint8 }
func genFn0() int8 {
    var v7 uint64 = 1
    var v8 uint8 = 2
    v7 = 18446744073709551615
    if false {
        for fi9 := 0; fi9 < 10; fi9++ {
            v7 += 1
            _ = fi9
        }
    }
    if ("\n\t" == "a") {
        v7 = (^(v7))
        v10 := GenS1{F0: v7, F1: 1, F2: 948}
        _ = v10
    }
    _ = v7
    _ = v8
    return -8
}
func genFn1() string {
    v11 := []int{}
    var v12 uint = 699
    _ = v11
    _ = v12
    return ((("foo bar" + "\n\t") + ("foo bar" + (("foo bar" + ("foo bar" + "")) + ""))) + (("" + "\n\t") + "foo bar"))
}

func main() {
    v1 := []uint32{}
    var v2 string = "foo bar"
    var v3 uint = 255
    var v4 int8 = 127
    var v5 uint64 = 584
    var v6 int64 = -9223372036854775808
    v4 -= (v4 & (v4 << (((v3 * (^(v3)))) & 63)))
    v13 := GenS2{F0: -9223372036854775808}
    if true {
        v14 := GenS1{F0: 422, F1: 1, F2: 1000}
        if false {
            for fi15 := 0; fi15 < 14; fi15++ {
                for fi16 := 0; fi16 < 13; fi16++ {
                    var v17 int = (fi16 / (((fi16 + (fi16 - 9223372036854775807))) | 1))
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v18 := GenS3{F0: -2147483648, F1: 42}
            for fi19 := 0; fi19 < 1; fi19++ {
                v4 = (v4 * v4)
                v20 := GenS2{F0: v6}
                if (!(32767 > 0)) {
                    v4 = -1
                    var v21 int8 = (v4 / (((v4 * (v4 - 35))) | 1))
                    v22 := GenS4{F0: v3, F1: 1000}
                    _ = v21
                    _ = v22
                } else {
                    v4 = (v4 * 42)
                    v23 := GenS2{F0: (v6 & v6)}
                    v24 := GenS2{F0: 9223372036854775807}
                    _ = v23
                    _ = v24
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        _ = v14
    }
    if (!(false || (!(-32768 < 42)))) {
        switch ("\n\t" + ((("foo bar" + ("foo bar" + v2)) + v2) + v2)) {
        case "hello":
            if (v3 > v3) {
                var v25 uint32 = 554
                v2 = v2
                var v26 uint = (v3 ^ (^(v3)))
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 15; fi27++ {
                    if (1 != uint8(v5)) {
                        v28 := GenS5{F0: (fi27 * 32768), F1: (v5 << (((v3 * uint(v5))) & 63)), F2: ((v2 + "\n\t") + "hello")}
                        var v29 int8 = 1
                        fmt.Println(fi27, v3, v3)
                        _ = v28
                        _ = v29
                    } else {
                        var v30 uint8 = 1
                        for fi31 := 0; fi31 < 12; fi31++ {
                            v5 -= v5
                            v3 = v3
                            _ = fi31
                        }
                        v32 := GenS1{F0: v5, F1: 0, F2: 42}
                        _ = v30
                        _ = v32
                    }
                    _ = fi27
                }
                if false {
                    if (true && (v3 == (^(v3)))) {
                        for fi33 := 0; fi33 < 12; fi33++ {
                            switch (fi33 | 574) {
                            case 1:
                                switch (fi33 * (fi33 + 280)) {
                                case 32767:
                                    if false {
                                        for fi34 := 0; fi34 < 7; fi34++ {
                                            switch (fi34 << (((v3 / ((v3) | 1))) & 63)) {
                                            case 127:
                                                var v35 float64 = (float64((-2.25 / 1.5)) * ((float64((3.14159 + (0.1 + ((-2.25 + -2.25) * 3.14159)))) / 1.5) + 0.0))
                                                _ = v35
                                            case 866:
                                                if true {
                                                    v4 = v4
                                                }
                                            case 1:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch 32768 {
                                                case -9223372036854775808:
                                                    switch fi33 {
                                                    case 42:
                                                        var v36 uint64 = 4294967295
                                                        _ = v36
                                                    default:
                                                        if (true && false) {
                                                            v1 = append(v1, 65)
                                                        } else {
                                                            if true {
                                                                v37 := GenS6{F0: 0}
                                                                var v38 float64 = -2.25
                                                                var v39 int16 = int16(v4)
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            } else {
                                                                v40 := GenS1{F0: v5, F1: 1, F2: 0}
                                                                _ = v40
                                                            }
                                                        }
                                                    }
                                                default:
                                                    var v41 int64 = int64(v3)
                                                    _ = v41
                                                }
                                            }
                                            var v42 string = v2
                                            _ = fi34
                                            _ = v42
                                        }
                                    } else {
                                        switch 127 {
                                        case -1:
                                            v5 -= (^(v5))
                                        default:
                                            var v43 uint = (^(v3))
                                            _ = v43
                                        }
                                    }
                                default:
                                    switch fi33 {
                                    case 42:
                                        for fi44 := 0; fi44 < 10; fi44++ {
                                            var v45 int8 = (v4 << ((370) & 63))
                                            fmt.Println(v4, v3, fi44)
                                            _ = fi44
                                            _ = v45
                                        }
                                    case 4294967295:
                                        v46 := GenS7{F0: v5, F1: 1}
                                        _ = v46
                                    case 128:
                                        var v47 int16 = -880
                                        _ = v47
                                    default:
                                        v48 := GenS8{F0: 802}
                                        _ = v48
                                    }
                                }
                            default:
                                if true {
                                    switch 32768 {
                                    case 127:
                                        var v49 uint16 = 255
                                        _ = v49
                                    case 32768:
                                        var v50 int8 = (v4 & genFn0())
                                        _ = v50
                                    case -553:
                                        v51 := GenS9{F0: (^(v5)), F1: (v6 ^ v6)}
                                        _ = v51
                                    }
                                }
                            }
                            v52 := GenS5{F0: fi33, F1: v5, F2: v2}
                            _ = fi33
                            _ = v52
                        }
                        var v53 bool = true
                        _ = v53
                    }
                    _ = len(v1)
                    v2 = (v2 + (v2 + v2))
                } else {
                    for fi54 := 0; fi54 < 9; fi54++ {
                        if (0 <= 1) {
                            _ = len(v1)
                            if (true || false) {
                                for fi55 := 0; fi55 < 17; fi55++ {
                                    fmt.Println(v4, fi55, v5)
                                    var v56 uint = v3
                                    if false {
                                        fmt.Println(v4)
                                        if true {
                                            if (0.1 < (float32((float64(1e10) / (3.14159 + (1e10 - (0.0 / (0.1 * (((float64(-2.25) / 1.5) / 1e10) - 1e10))))))) + 1.5)) {
                                                switch (fi55 - fi54) {
                                                case -1:
                                                    _ = len(v1)
                                                case 492:
                                                    var v57 string = ("a" + v2)
                                                    _ = v57
                                                default:
                                                    switch "a" {
                                                    case "":
                                                        var v58 int = fi54
                                                        _ = v58
                                                    default:
                                                        var v59 int16 = -944
                                                        _ = v59
                                                    }
                                                }
                                                v60 := GenS10{F0: 0, F1: (v5 + 42)}
                                                if true {
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(v56, v3)
                                                    case "foo bar":
                                                        var v61 int64 = 1
                                                        _ = v61
                                                    case "\n\t":
                                                        switch v2 {
                                                        case "foo bar":
                                                            v4 -= (v4 | v4)
                                                        case "a":
                                                            switch (v2 + (v2 + "")) {
                                                            case "foo bar":
                                                                if (!(-2147483648 == 1)) {
                                                                    switch fi55 {
                                                                    case 305:
                                                                        v62 := GenS11{F0: v6, F1: (-(v4)), F2: fi55}
                                                                        _ = v62
                                                                    case 128:
                                                                        if true {
                                                                            if false {
                                                                                var v63 bool = ((v6 | (^(v6))) > (-(v6)))
                                                                                _ = v63
                                                                            } else {
                                                                                v64 := GenS9{F0: v5, F1: v6}
                                                                                _ = v64
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v4)
                                                                            if false {
                                                                                v56 = 255
                                                                                var v65 uint8 = 200
                                                                                _ = v65
                                                                            } else {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v3 = 1
                                                                                v66 := GenS9{F0: 0, F1: (^(v6))}
                                                                                v67 := GenS10{F0: 255, F1: uint64(v3)}
                                                                                _ = v66
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (!false) {
                                                                            switch fi55 {
                                                                            case 749:
                                                                                fmt.Println(v4, v5, fi55)
                                                                            case -1:
                                                                                fmt.Println(fi54, fi55)
                                                                            }
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                v68 := GenS4{F0: (v3 << (((^(v3))) & 63)), F1: 1}
                                                                                _ = v68
                                                                            case "a":
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                switch (fi55 * fi54) {
                                                                                case 4294967295:
                                                                                    switch v2 {
                                                                                    case "hello":
                                                                                        v3 = v3
                                                                                    case "a":
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        v69 := GenS12{F0: v2}
                                                                                        _ = v69
                                                                                    case "":
                                                                                        switch (fi54 << ((544) & 63)) {
                                                                                        case -832:
                                                                                            if (v2 > "\n\t") {
                                                                                                var v70 string = v2
                                                                                                fmt.Println(v3)
                                                                                                _ = v70
                                                                                            } else {
                                                                                                var v71 uint = 42
                                                                                                var v72 int8 = (v4 << ((v71) & 63))
                                                                                                _ = v71
                                                                                                _ = v72
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case 32767:
                                                                                    fmt.Println(v56)
                                                                                case 861:
                                                                                    var v73 uint32 = 574
                                                                                    _ = v73
                                                                                default:
                                                                                    var v74 uint32 = 4294967295
                                                                                    _ = v74
                                                                                }
                                                                            case "foo bar":
                                                                                var v75 int16 = 1000
                                                                                _ = v75
                                                                            default:
                                                                                v1 = append(v1, 2147483647)
                                                                            }
                                                                            var v76 string = "\n\t"
                                                                            v6 -= v6
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    fmt.Println(v6, fi54, v5)
                                                                } else {
                                                                    if (683 == v5) {
                                                                        if (v3 == v3) {
                                                                            if true {
                                                                                switch "" {
                                                                                case "":
                                                                                    var v77 uint32 = 2147483648
                                                                                    _ = v77
                                                                                case "\n\t":
                                                                                    switch (fi54 >> ((v3) & 63)) {
                                                                                    case -591:
                                                                                        switch ("" + v2) {
                                                                                        case "a":
                                                                                            switch ("\n\t" + v2) {
                                                                                            case "hello":
                                                                                                v78 := GenS4{F0: uint(v5), F1: -858}
                                                                                                v2 = "foo bar"
                                                                                                var v79 float64 = float64((-2.25 / 1.0))
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            case "a":
                                                                                                var v80 string = ((v2 + ("a" + "")) + (("" + "foo bar") + ""))
                                                                                                _ = v80
                                                                                            case "":
                                                                                                if (!true) {
                                                                                                    v56 = v56
                                                                                                    var v81 bool = false
                                                                                                    _ = v81
                                                                                                }
                                                                                            default:
                                                                                                var v82 int64 = (v6 >> ((v3) & 63))
                                                                                                _ = v82
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v83 := GenS8{F0: 42}
                                                                                            _ = v83
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 65535
                                                                                    }
                                                                                }
                                                                                var v84 int = -2147483648
                                                                                _ = v84
                                                                            }
                                                                        } else {
                                                                            v2 = "foo bar"
                                                                        }
                                                                        fmt.Println(v3, v2)
                                                                        switch v2 {
                                                                        case "foo bar":
                                                                            var v85 uint64 = (v5 + v5)
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        _ = len(v1)
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v4, v4)
                                                            }
                                                        }
                                                    default:
                                                        v86 := GenS9{F0: v5, F1: (^(v6))}
                                                        _ = v86
                                                    }
                                                } else {
                                                    fmt.Println(v6, v2, v6)
                                                    fmt.Println(fi54, v4)
                                                    fmt.Println(v4, v6)
                                                    fmt.Println(v2, v5, fi55)
                                                    var v87 string = ""
                                                    _ = v87
                                                }
                                                _ = v60
                                            }
                                            v2 = v2
                                            v3 = v3
                                            var v88 int32 = -2147483648
                                            v89 := GenS2{F0: v6}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(v2, v3)
                                            _ = v88
                                            _ = v89
                                        } else {
                                            fmt.Println(v6, fi55)
                                            switch -1 {
                                            case 127:
                                                switch fi54 {
                                                case 2147483647:
                                                    switch ("\n\t" + ("a" + "")) {
                                                    case "hello":
                                                        v1 = append(v1, 1)
                                                        v90 := GenS13{F0: false, F1: (-(v6))}
                                                        _ = v90
                                                    case "\n\t":
                                                        v91 := GenS14{F0: 32767, F1: v5, F2: fi55}
                                                        _ = v91
                                                    default:
                                                        v92 := GenS15{F0: 1}
                                                        _ = v92
                                                    }
                                                }
                                            case 9223372036854775807:
                                                if true {
                                                    v93 := GenS15{F0: -1}
                                                    fmt.Println(fi55)
                                                    var v94 float64 = -2.25
                                                    _ = v93
                                                    _ = v94
                                                }
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 2147483647
                                            }
                                        }
                                    } else {
                                        var v95 uint64 = v5
                                        var v96 uint16 = 1
                                        _ = v95
                                        _ = v96
                                    }
                                    if true {
                                        v97 := GenS16{F0: 1}
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v98 := GenS10{F0: 65535, F1: 2}
                                            switch fi54 {
                                            case 127:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 903:
                                                var v99 int8 = int8(v3)
                                                _ = v99
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = 120
                                                }
                                            }
                                            _ = v98
                                        } else {
                                            switch fi54 {
                                            case -9223372036854775808:
                                                var v100 int8 = v4
                                                _ = v100
                                            default:
                                                v4 += (v4 - -128)
                                            }
                                            switch (fi55 | fi55) {
                                            case -1:
                                                if (0 > 0) {
                                                    var v101 float64 = (((float64(1.5) - float64(0.1)) + 3.14159) - 1e10)
                                                    var v102 int16 = 32767
                                                    _ = v101
                                                    _ = v102
                                                } else {
                                                    v103 := GenS17{F0: ("foo bar" + v2), F1: 1}
                                                    v2 = "\n\t"
                                                    fmt.Println(v2, v5)
                                                    _ = v103
                                                }
                                            case 517:
                                                v1 = append(v1, 106)
                                            case -9:
                                                var v104 uint64 = uint64(v56)
                                                _ = v104
                                            }
                                        }
                                        v105 := GenS18{F0: v2}
                                        _ = v105
                                        _ = v97
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                            } else {
                                switch (fi54 * (fi54 / (((fi54 % (((fi54 % ((fi54) | 1))) | 1))) | 1))) {
                                case -1:
                                    var v106 bool = false
                                    _ = v106
                                case 127:
                                    v6 += (v6 & 42)
                                default:
                                    var v107 uint64 = (v5 + v5)
                                    _ = v107
                                }
                                fmt.Println(v4, v4, v2)
                                v6 = (v6 + v6)
                            }
                        }
                        _ = fi54
                    }
                    var v108 int8 = (-(v4))
                    if (0 > 65535) {
                        var v109 uint8 = 254
                        switch 1 {
                        case -1:
                            fmt.Println(v6)
                        case -806:
                            var v110 int8 = (v108 + (v4 & genFn0()))
                            _ = v110
                        case 2147483647:
                            var v111 int64 = -681
                            _ = v111
                        }
                        if false {
                            var v112 int64 = v6
                            v113 := GenS8{F0: 1}
                            var v114 float64 = (float64(((-2.25 * -2.25) + 0.1)) - 3.14159)
                            _ = v112
                            _ = v113
                            _ = v114
                        }
                        _ = v109
                    }
                    _ = v108
                }
            }
        case "a":
            var v115 float64 = 1.5
            _ = v115
        case "\n\t":
            if true {
                v3 += 799
                var v116 uint = (^(v3))
                switch 32768 {
                case 128:
                    var v117 string = (v2 + v2)
                    _ = v117
                default:
                    var v118 int8 = (v4 / (((^(v4))) | 1))
                    _ = v118
                }
                _ = v116
            }
        default:
            if (0 >= 2) {
                fmt.Println(v5, v2)
                var v119 uint = 1
                _ = v119
            }
        }
        switch v2 {
        case "foo bar":
            var v120 uint16 = uint16(v5)
            _ = v120
        default:
            v121 := GenS19{F0: v2, F1: 0}
            _ = v121
        }
        var v122 int8 = (v4 >> ((v3) & 63))
        _ = v122
    }
    v2 = ("" + (v2 + "a"))
    v6 = int64(v4)
    v123 := GenS16{F0: 255}
    var v124 uint64 = 2
    if ((65535 == 727) && false) {
        var v125 uint16 = 2
        var v126 string = ("foo bar" + "hello")
        var v127 float64 = (0.1 * (1.5 / 1.0))
        _ = v125
        _ = v126
        _ = v127
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v123
    _ = v124
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
