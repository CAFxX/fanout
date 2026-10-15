package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int16; F1 uint8 }
type GenS15 struct { F0 string; F1 bool; F2 int64 }
type GenS16 struct { F0 uint16; F1 uint32 }
type GenS17 struct { F0 int64; F1 int8; F2 string }

func main() {
    v1 := []string{"\n\t"}
    var v2 uint32 = 65535
    var v3 float64 = float64(float32(3.14159))
    var v4 uint16 = 32767
    fmt.Println(v2, v2)
    fmt.Println(v2, v4)
    if true {
        if ((!false) && false) {
            for fi5 := 0; fi5 < 20; fi5++ {
                for fi6 := 0; fi6 < 8; fi6++ {
                    if true {
                        v2 -= v2
                    } else {
                        var v7 int16 = 15
                        if true {
                            v8 := GenS1{F0: 9223372036854775807, F1: 1}
                            fmt.Println(v2)
                            v3 = ((v3 / 3.14159) * v3)
                            _ = v8
                        } else {
                            var v9 uint8 = 168
                            _ = v9
                        }
                        _ = v7
                    }
                    _ = fi6
                }
                v4 = 20
                if false {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 string = "a"
                        fmt.Println(v2)
                        _ = fi10
                        _ = v11
                    }
                }
                _ = fi5
            }
            v4 = v4
            v2 = v2
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: -364, F1: 9223372036854775807}
                for fi14 := 0; fi14 < 9; fi14++ {
                    if (-521 >= 1) {
                        switch ("foo bar" + ("\n\t" + "foo bar")) {
                        case "a":
                            switch 0 {
                            case -9223372036854775808:
                                var v15 int64 = -9223372036854775808
                                _ = v15
                            }
                        case "":
                            v16 := GenS1{F0: 42, F1: 0}
                            _ = v16
                        case "hello":
                            v2 -= (v2 << ((1) & 63))
                        default:
                            v17 := GenS1{F0: -9223372036854775808, F1: -569}
                            _ = v17
                        }
                        v18 := GenS2{F0: (v2 >> ((65535) & 63))}
                        var v19 float64 = v3
                        if true {
                            fmt.Println(v4, v19, fi14)
                        }
                        switch "\n\t" {
                        case "a":
                            v20 := GenS1{F0: -675, F1: -1}
                            _ = v20
                        }
                        _ = v18
                        _ = v19
                    }
                    fmt.Println(v4, v3, v2)
                    v21 := GenS2{F0: 569}
                    _ = fi14
                    _ = v21
                }
                var v22 float64 = ((((-2.25 * ((float64(-2.25) / v3) + v3)) + (1e10 * v3)) / 3.14159) / ((float64(1e10) * v3) - 0.0))
                _ = fi12
                _ = v13
                _ = v22
            }
            for fi23 := 0; fi23 < 4; fi23++ {
                v2 = 2147483648
                if (3.14159 >= (3.14159 * 0.1)) {
                    if (true || (4294967295 != 4294967295)) {
                        v4 += (v4 % ((v4) | 1))
                    }
                } else {
                    if (-499 < int64(v2)) {
                        var v24 uint64 = 42
                        _ = v24
                    } else {
                        var v25 float64 = -2.25
                        if false {
                            for fi26 := 0; fi26 < 10; fi26++ {
                                switch "foo bar" {
                                case "a":
                                    v4 -= v4
                                case "hello":
                                    v27 := GenS1{F0: 0, F1: -9223372036854775808}
                                    _ = v27
                                case "foo bar":
                                    v28 := GenS3{F0: 2147483648, F1: 0}
                                    _ = len(v1)
                                    _ = v28
                                default:
                                    var v29 int8 = 127
                                    _ = v29
                                }
                                fmt.Println(v4)
                                _ = fi26
                            }
                        }
                        fmt.Println(v2, v2)
                        _ = v25
                    }
                    if true {
                        if (false || false) {
                            for fi30 := 0; fi30 < 5; fi30++ {
                                v31 := GenS2{F0: 1}
                                v1 = append(v1, ("\n\t" + "\n\t"))
                                _ = fi30
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 9; fi32++ {
                                if ("hello" < "hello") {
                                    if true {
                                        if true {
                                            v33 := GenS1{F0: 9223372036854775807, F1: 846}
                                            _ = v33
                                        }
                                        var v34 uint32 = v2
                                        _ = v34
                                    }
                                    if (float64(1.5) <= float64(((float32(v3) + (3.14159 * (0.1 + -2.25))) * 3.14159))) {
                                        var v35 uint64 = 0
                                        v4 = (v4 & (^(v4)))
                                        fmt.Println(v4, v2)
                                        _ = v35
                                    } else {
                                        if false {
                                            v4 -= (v4 * 115)
                                            _ = len(v1)
                                        } else {
                                            v2 = 2147483648
                                            if (-586 == -9223372036854775808) {
                                                var v36 uint32 = v2
                                                var v37 int32 = -1
                                                var v38 float64 = (v3 - v3)
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            } else {
                                                fmt.Println(fi32)
                                                v4 = 335
                                                if false {
                                                    fmt.Println(v4, fi32, fi23)
                                                    if false {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v3 = 1.5
                                                        fmt.Println(v4, v3, v3)
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = ""
                                                        }
                                                        var v39 int16 = -621
                                                        var v40 bool = (!false)
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                } else {
                                                    v41 := GenS3{F0: 2, F1: 42}
                                                    _ = v41
                                                }
                                            }
                                        }
                                        v4 = v4
                                        if false {
                                            if true {
                                                var v42 int8 = 100
                                                switch (fi23 / ((fi23) | 1)) {
                                                case 42:
                                                    var v43 uint32 = (v2 + v2)
                                                    _ = v43
                                                case 2147483647:
                                                    fmt.Println(v3, v3, v4)
                                                default:
                                                    fmt.Println(v4, fi23)
                                                }
                                                if (0.0 > 0.0) {
                                                    var v44 uint64 = 2
                                                    v44 = (v44 ^ 42)
                                                    if len(v1) > 0 {
                                                        v1[0] = "a"
                                                    }
                                                    _ = v44
                                                } else {
                                                    v1 = append(v1, (("foo bar" + "hello") + "foo bar"))
                                                }
                                                _ = v42
                                            }
                                            var v45 int = -2147483648
                                            _ = v45
                                        } else {
                                            v3 = v3
                                        }
                                    }
                                }
                                _ = fi32
                            }
                        } else {
                            v1 = append(v1, "")
                            for fi46 := 0; fi46 < 3; fi46++ {
                                fmt.Println(fi46, v3, v2)
                                _ = fi46
                            }
                        }
                        v47 := GenS3{F0: (v2 & 2147483648), F1: 0}
                        switch (fi23 % ((fi23) | 1)) {
                        case 32768:
                            fmt.Println(v3, v4, fi23)
                        case 304:
                            var v48 int32 = 333
                            _ = v48
                        case -678:
                            var v49 int8 = -128
                            _ = v49
                        default:
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                if true {
                                    fmt.Println(v4)
                                    var v50 int16 = -32768
                                    _ = v50
                                }
                                fmt.Println(fi23, fi23, fi23)
                            } else {
                                v51 := GenS3{F0: 42, F1: 9223372036854775807}
                                for fi52 := 0; fi52 < 13; fi52++ {
                                    if (true || true) {
                                        v4 -= (v4 / ((v4) | 1))
                                        switch ("" + "") {
                                        case "hello":
                                            if false {
                                                v3 = ((v3 * v3) - 0.1)
                                                var v53 string = "a"
                                                v54 := GenS2{F0: (v2 - 121)}
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        v3 = float64(0.0)
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                        var v55 uint32 = v2
                                        var v56 string = ("hello" + (("a" + "a") + (("a" + "foo bar") + "a")))
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v57 := GenS3{F0: (^(v2)), F1: -9223372036854775808}
                                        _ = v57
                                    }
                                    _ = fi52
                                }
                                for fi58 := 0; fi58 < 16; fi58++ {
                                    v4 = v4
                                    v59 := GenS4{F0: 2147483647, F1: (^(v2))}
                                    var v60 int8 = -128
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                                _ = v51
                            }
                        }
                        _ = v47
                    }
                }
                _ = fi23
            }
        }
    } else {
        v61 := GenS5{F0: (v4 ^ v4), F1: uint8(v4)}
        fmt.Println(v3, v3)
        _ = v61
    }
    v4 = v4
    v2 += v2
    if (false && false) {
        var v62 int64 = 749
        v4 -= (v4 << ((42) & 63))
        v63 := GenS6{F0: 10}
        _ = v62
        _ = v63
    } else {
        for fi64 := 0; fi64 < 10; fi64++ {
            var v65 uint16 = (v4 & v4)
            var v66 int8 = -104
            v4 = (^(v65))
            _ = fi64
            _ = v65
            _ = v66
        }
        v1 = append(v1, "\n\t")
        switch 2147483647 {
        case -711:
            if true {
                var v67 string = ("hello" + "foo bar")
                _ = v67
            } else {
                v4 -= (^(v4))
                switch "a" {
                case "\n\t":
                    var v68 uint32 = v2
                    _ = v68
                case "foo bar":
                    v4 = v4
                }
                switch "\n\t" {
                case "foo bar":
                    if true {
                        var v69 string = ("\n\t" + ("a" + ("foo bar" + "a")))
                        _ = v69
                    } else {
                        v70 := GenS7{F0: 2, F1: 202, F2: 2}
                        if false {
                            if ((38 >= 42) || false) {
                                switch 2147483647 {
                                case 0:
                                    var v71 uint16 = 42
                                    _ = v71
                                case 32768:
                                    v2 = (v2 << ((0) & 63))
                                }
                                var v72 string = "foo bar"
                                v73 := GenS7{F0: 255, F1: 42, F2: 1}
                                _ = v72
                                _ = v73
                            }
                        } else {
                            switch ("foo bar" + "") {
                            case "\n\t":
                                fmt.Println(v4)
                            default:
                                var v74 uint8 = 132
                                _ = v74
                            }
                            var v75 uint = 595
                            _ = v75
                        }
                        for fi76 := 0; fi76 < 15; fi76++ {
                            if ((0 >= 626) || true) {
                                if (!(977 >= 560)) {
                                    v77 := GenS5{F0: v4, F1: 42}
                                    var v78 int16 = 220
                                    _ = v77
                                    _ = v78
                                } else {
                                    for fi79 := 0; fi79 < 17; fi79++ {
                                        var v80 uint = 289
                                        _ = fi79
                                        _ = v80
                                    }
                                    var v81 int = fi76
                                    _ = v81
                                }
                            }
                            _ = fi76
                        }
                        _ = v70
                    }
                case "":
                    switch "" {
                    case "a":
                        var v82 string = "foo bar"
                        _ = v82
                    default:
                        v4 -= v4
                    }
                case "\n\t":
                    for fi83 := 0; fi83 < 2; fi83++ {
                        var v84 float32 = 0.0
                        var v85 uint = 255
                        fmt.Println(fi83, v2, v3)
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                default:
                    fmt.Println(v4, v4)
                }
            }
        case 685:
            for fi86 := 0; fi86 < 9; fi86++ {
                _ = len(v1)
                switch fi86 {
                case -711:
                    v87 := GenS8{F0: 0, F1: (127 >= -53)}
                    _ = v87
                case 42:
                    for fi88 := 0; fi88 < 9; fi88++ {
                        var v89 uint16 = v4
                        var v90 bool = (v4 < (v4 & v4))
                        v90 = v90
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                case 4294967295:
                    var v91 string = "hello"
                    _ = v91
                }
                _ = fi86
            }
        case 440:
            if (2 != 357) {
                fmt.Println(v2, v4, v2)
                v92 := GenS9{F0: 101}
                if (!false) {
                    if (42 <= -128) {
                        v93 := GenS10{F0: ("\n\t" + "a")}
                        v3 -= v3
                        v94 := GenS11{F0: v2, F1: (v2 - v2), F2: 117}
                        _ = v93
                        _ = v94
                    }
                    v4 = 42
                } else {
                    for fi95 := 0; fi95 < 4; fi95++ {
                        if len(v1) > 0 {
                            v1[0] = ((("hello" + (("foo bar" + "a") + "")) + "hello") + ("" + "a"))
                        }
                        var v96 uint8 = 1
                        switch "hello" {
                        case "foo bar":
                            v97 := GenS3{F0: v2, F1: -297}
                            _ = v97
                        case "\n\t":
                            v2 += v2
                        case "a":
                            if (!true) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v98 uint64 = 18446744073709551615
                                v4 = 1
                                _ = v98
                            } else {
                                for fi99 := 0; fi99 < 11; fi99++ {
                                    if (((1.5 + (v3 - -2.25)) / (float64(1.5) + ((v3 / v3) / float64((-2.25 - 1e10))))) != v3) {
                                        v100 := GenS7{F0: 4294967295, F1: (v96 / ((v96) | 1)), F2: 4294967295}
                                        _ = v100
                                    }
                                    fmt.Println(fi99, v96, v4)
                                    v101 := GenS10{F0: "\n\t"}
                                    _ = fi99
                                    _ = v101
                                }
                                v2 = (v2 ^ (v2 & (v2 + 182)))
                                v1 = append(v1, "foo bar")
                            }
                        default:
                            for fi102 := 0; fi102 < 9; fi102++ {
                                if true {
                                    fmt.Println(v3)
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        switch (fi95 ^ fi102) {
                                        case 4294967295:
                                            fmt.Println(fi95, v96, fi95)
                                        }
                                    default:
                                        var v103 uint8 = (v96 / (((v96 + 255)) | 1))
                                        _ = v103
                                    }
                                    var v104 uint = 499
                                    _ = v104
                                }
                                switch (fi102 * (fi102 / ((int(v4)) | 1))) {
                                case 879:
                                    var v105 uint16 = v4
                                    _ = v105
                                default:
                                    var v106 int16 = -1
                                    _ = v106
                                }
                                v107 := GenS12{F0: 2147483647, F1: 255}
                                _ = fi102
                                _ = v107
                            }
                        }
                        _ = fi95
                        _ = v96
                    }
                }
                _ = v92
            }
        default:
            for fi108 := 0; fi108 < 11; fi108++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v109 bool = false
                switch fi108 {
                case 1:
                    _ = len(v1)
                case -2147483648:
                    if v109 {
                        v109 = (!(356 <= 2))
                        v3 = 1.5
                    } else {
                        if (-9223372036854775808 >= -1) {
                            v109 = true
                            switch (fi108 ^ fi108) {
                            case 4294967295:
                                fmt.Println(v2)
                            case 822:
                                if ((!(0 != 243)) && ((v4 % (((^(v4))) | 1)) <= v4)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if v109 {
                                        var v110 bool = v109
                                        _ = v110
                                    } else {
                                        var v111 uint = uint(v4)
                                        fmt.Println(v4, v3, v111)
                                        _ = v111
                                    }
                                }
                            case -692:
                                v3 = (v3 - v3)
                            }
                        } else {
                            for fi112 := 0; fi112 < 5; fi112++ {
                                v3 = 3.14159
                                _ = fi112
                            }
                            fmt.Println(v109, v2)
                            v113 := GenS13{F0: -158}
                            _ = v113
                        }
                        if v109 {
                            switch "\n\t" {
                            case "\n\t":
                                v114 := GenS1{F0: 0, F1: -9223372036854775808}
                                _ = v114
                            case "a":
                                v115 := GenS14{F0: 114, F1: 42}
                                _ = v115
                            case "":
                                for fi116 := 0; fi116 < 10; fi116++ {
                                    v117 := GenS15{F0: ((("foo bar" + ((((("a" + (("a" + ("a" + "")) + ("\n\t" + ""))) + "hello") + "") + "") + (("foo bar" + ("foo bar" + "\n\t")) + "hello"))) + "\n\t") + (("\n\t" + "a") + "hello")), F1: (v109 || v109), F2: 9223372036854775807}
                                    v3 = float64(float32(float64(0.0)))
                                    if v109 {
                                        v2 += v2
                                        v109 = (v3 >= v3)
                                        v118 := GenS16{F0: 0, F1: 2}
                                        _ = v118
                                    }
                                    v119 := GenS16{F0: v4, F1: (v2 - (v2 + (v2 ^ (v2 % ((65) | 1)))))}
                                    _ = fi116
                                    _ = v117
                                    _ = v119
                                }
                            }
                            var v120 int32 = 757
                            v121 := GenS17{F0: 1, F1: 127, F2: "hello"}
                            _ = v120
                            _ = v121
                        }
                    }
                case 568:
                    _ = len(v1)
                }
                _ = fi108
                _ = v109
            }
        }
    }
    var v122 int16 = 378
    var v123 int = 575
    var v124 uint16 = (v4 | v4)
    v123 = (-(v123))
    v124 = v124
    for fi125 := 0; fi125 < 9; fi125++ {
        v126 := GenS6{F0: 0}
        v127 := GenS8{F0: 42, F1: ((1e10 * 0.1) != (-2.25 / float32(v3)))}
        var v128 float32 = (3.14159 / 0.1)
        _ = fi125
        _ = v126
        _ = v127
        _ = v128
    }
    for fi129 := 0; fi129 < 6; fi129++ {
        var v130 uint8 = 1
        v131 := GenS14{F0: 42, F1: v130}
        _ = len(v1)
        _ = fi129
        _ = v130
        _ = v131
    }
    switch v123 {
    case 32767:
        var v132 int = v123
        _ = v132
    default:
        v124 = v4
    }
    v4 -= v124
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v122
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v4
}
