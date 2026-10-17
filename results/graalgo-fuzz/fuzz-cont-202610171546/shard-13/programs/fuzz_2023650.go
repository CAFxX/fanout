package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint; F1 bool; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 uint64; F2 int }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8; F1 int64 }
type GenS11 struct { F0 int32; F1 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 uint32; F1 uint; F2 int32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint16; F1 uint; F2 int8 }
type GenS17 struct { F0 uint16; F1 uint64 }
func genFn0() string {
    v6 := GenS1{F0: 434, F1: 0}
    var v7 uint8 = 42
    if (-9223372036854775808 >= 726) {
        v8 := GenS1{F0: 32767, F1: 1}
        v7 = 128
        for fi9 := 0; fi9 < 9; fi9++ {
            if true {
                var v10 int32 = -2147483648
                _ = v10
            } else {
                if ((27 != -128) && false) {
                    v7 += v7
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v7 -= 127
                        switch "a" {
                        case "hello":
                            switch ((("foo bar" + "foo bar") + ("a" + "\n\t")) + "\n\t") {
                            case "":
                                v12 := []int{fi9, 490}
                                _ = v12
                            }
                        case "foo bar":
                            if ((false || true) && false) {
                                var v13 int64 = 0
                                var v14 bool = true
                                v15 := GenS1{F0: 1, F1: (v13 * v13)}
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                v7 = (v7 + 2)
                            }
                        }
                        v7 = v7
                        _ = fi11
                    }
                    fmt.Println(fi9)
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 float64 = (((0.0 - (-2.25 / -2.25)) / 3.14159) / 1.5)
                    v18 := []int{}
                    v7 += (v7 % (((v7 << ((536) & 63))) | 1))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 float32 = (((1e10 + (1.5 + float32(float64(1e10)))) / (float32(3.14159) + (((-2.25 + ((3.14159 / (-2.25 + 3.14159)) + 3.14159)) - 0.0) - 0.1))) + (0.0 - (1.5 / 1.0)))
                v21 := GenS2{F0: 255, F1: (!false), F2: 274}
                v20 += 1.5
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi9
        }
        _ = v8
    } else {
        var v22 uint64 = 0
        v23 := GenS1{F0: 21, F1: 1}
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS1{F0: 255, F1: -1}
        _ = fi24
        _ = v25
    }
    var v26 int32 = 42
    _ = v26
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 uint8) uint64 {
    var v27 bool = (2147483647 != 1)
    v28 := []int{}
    _ = p0
    _ = v27
    _ = v28
    return 529
}

func main() {
    v1 := []string{"a", (((("hello" + "hello") + ("a" + (("a" + "hello") + "hello"))) + ("a" + (("hello" + "hello") + "hello"))) + "foo bar"), ("foo bar" + "a"), ((((("hello" + ((("" + "\n\t") + "") + "a")) + "a") + "foo bar") + "") + ("foo bar" + ("hello" + "")))}
    var v2 string = (((("a" + "a") + "a") + "hello") + (((("" + "") + "a") + "hello") + "a"))
    v3 := []string{v2, "hello"}
    v4 := make(map[string]int)
    var v5 uint8 = 127
    switch v2 {
    case "hello":
        var v29 float32 = 1e10
        _ = v29
    default:
        v5 = (v5 << ((569) & 63))
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(fi31)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi31
        }
        _ = fi30
    }
    if false {
        v5 += (v5 & 42)
        for fi32 := 0; fi32 < 16; fi32++ {
            v5 = v5
            var v33 uint8 = 255
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := GenS2{F0: 554, F1: (0 < -1), F2: 2}
        for fi36 := 0; fi36 < 20; fi36++ {
            v2 = (((("a" + (v2 + "hello")) + (v2 + "foo bar")) + v2) + (v2 + "hello"))
            v37 := GenS3{F0: 1}
            fmt.Println(v2, v5)
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    v38 := GenS3{F0: 42}
    v39 := GenS1{F0: 286, F1: 1}
    v2 = ((v2 + v2) + v2)
    v2 = "hello"
    if true {
        fmt.Println(v5, v5)
        v40 := GenS4{F0: 42, F1: 128, F2: 1}
        switch 0 {
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 127:
            v2 = v2
        }
        _ = v40
    } else {
        if false {
            if false {
                if ((!false) && false) {
                    if false {
                        switch v2 {
                        case "hello":
                            v5 += v5
                        default:
                            fmt.Println(v5)
                        }
                        for fi41 := 0; fi41 < 17; fi41++ {
                            for fi42 := 0; fi42 < 12; fi42++ {
                                fmt.Println(fi41, fi41)
                                _ = fi42
                            }
                            for fi43 := 0; fi43 < 5; fi43++ {
                                switch (v2 + (v2 + "")) {
                                case "a":
                                    var v44 int = fi41
                                    _ = v44
                                default:
                                    var v45 uint16 = 32767
                                    _ = v45
                                }
                                _ = fi43
                            }
                            _ = fi41
                        }
                    } else {
                        fmt.Println(v5, v5)
                        if false {
                            fmt.Println(v5, v5)
                            v2 = ""
                            var v46 float32 = -2.25
                            _ = v46
                        }
                        fmt.Println(v5, v2)
                    }
                    for fi47 := 0; fi47 < 2; fi47++ {
                        v2 = v2
                        for fi48 := 0; fi48 < 18; fi48++ {
                            v2 = v2
                            _ = fi48
                        }
                        _ = fi47
                    }
                    var v49 int32 = -2147483648
                    _ = v49
                }
                if (true || false) {
                    var v50 float64 = 1e10
                    var v51 uint64 = 1
                    for fi52 := 0; fi52 < 19; fi52++ {
                        var v53 int = fi52
                        for fi54 := 0; fi54 < 18; fi54++ {
                            switch (fi52 >> ((4294967295) & 63)) {
                            case 128:
                                fmt.Println(v51, fi54, fi52)
                            case 816:
                                v53 -= (fi54 >> ((4294967295) & 63))
                            case -1:
                                fmt.Println(fi54, v53)
                            }
                            if len(v1) > 0 {
                                v1[0] = (v2 + (v2 + "\n\t"))
                            }
                            v55 := GenS2{F0: 255, F1: false, F2: v51}
                            _ = fi54
                            _ = v55
                        }
                        v53 -= (v53 ^ 42)
                        _ = fi52
                        _ = v53
                    }
                    _ = v50
                    _ = v51
                } else {
                    v56 := GenS5{F0: 2147483647}
                    var v57 float64 = 1e10
                    for fi58 := 0; fi58 < 14; fi58++ {
                        fmt.Println(v5, v5, fi58)
                        var v59 uint64 = 18446744073709551615
                        if len(v1) > 0 {
                            v1[0] = (("\n\t" + v2) + ("a" + v2))
                        }
                        _ = fi58
                        _ = v59
                    }
                    _ = v56
                    _ = v57
                }
            } else {
                v2 = v2
                switch 4294967295 {
                case -363:
                    v60 := GenS3{F0: 1000}
                    _ = v60
                case -57:
                    var v61 int = 32768
                    _ = v61
                default:
                    switch -9223372036854775808 {
                    case -785:
                        var v62 float32 = -2.25
                        _ = v62
                    case 127:
                        var v63 int16 = -784
                        _ = v63
                    case -2147483648:
                        var v64 int8 = -116
                        _ = v64
                    }
                }
            }
        }
    }
    switch 128 {
    case -1:
        if len(v3) > 0 {
            v3[0] = v2
        }
    }
    var v65 uint8 = v5
    switch 127 {
    case 128:
        for fi66 := 0; fi66 < 18; fi66++ {
            for fi67 := 0; fi67 < 6; fi67++ {
                switch fi66 {
                case -1:
                    v1 = append(v1, "hello")
                case 0:
                    v68 := GenS6{F0: 2147483647, F1: 18446744073709551615, F2: 894}
                    _ = v68
                }
                var v69 uint32 = 1
                _ = fi67
                _ = v69
            }
            switch v2 {
            case "\n\t":
                if true {
                    v70 := GenS7{F0: 12, F1: 1}
                    switch "foo bar" {
                    case "a":
                        if (!true) {
                            v65 -= v65
                            v71 := v4[(((v2 + (("foo bar" + v2) + v2)) + ((v2 + ("\n\t" + "a")) + v2)) + ((((("foo bar" + v2) + "hello") + ((v2 + "a") + v2)) + v2) + ""))]
                            _ = v71
                        }
                    }
                    switch (fi66 - (fi66 >> ((0) & 63))) {
                    case 2147483647:
                        for fi72 := 0; fi72 < 9; fi72++ {
                            switch (-(fi66)) {
                            case 1:
                                if false {
                                    v73 := GenS1{F0: 1, F1: 0}
                                    fmt.Println(v5)
                                    _ = v73
                                }
                            case 0:
                                v2 = v2
                            default:
                                v74 := GenS2{F0: 0, F1: false, F2: 0}
                                _ = v74
                            }
                            _ = fi72
                        }
                    case 42:
                        switch v2 {
                        case "a":
                            v75 := GenS3{F0: -1}
                            _ = v75
                        case "foo bar":
                            for fi76 := 0; fi76 < 8; fi76++ {
                                v5 = 0
                                v2 = v2
                                switch fi66 {
                                case 128:
                                    v2 = ((v2 + ((v2 + v2) + "")) + v2)
                                case 1:
                                    fmt.Println(v5, v65, fi76)
                                }
                                v2 = "hello"
                                v2 = v2
                                v65 += v65
                                _ = fi76
                            }
                        case "hello":
                            var v77 int64 = int64(fi66)
                            _ = v77
                        }
                    case 548:
                        fmt.Println(v2)
                    }
                    _ = v70
                }
            case "hello":
                if true {
                    var v78 uint16 = 65535
                    fmt.Println(v78, v2, v2)
                    var v79 uint8 = (v5 >> ((204) & 63))
                    _ = v78
                    _ = v79
                }
            case "":
                fmt.Println(v5, v2)
            }
            _ = fi66
        }
    case 613:
        if false {
            v80 := GenS8{F0: 1000, F1: 784}
            fmt.Println(v5)
            var v81 float32 = (-2.25 - 0.1)
            _ = v80
            _ = v81
        }
    case -2147483648:
        if false {
            v82 := GenS3{F0: 32767}
            switch -359 {
            case 0:
                for fi83 := 0; fi83 < 19; fi83++ {
                    v84 := GenS2{F0: 1, F1: false, F2: 280}
                    fmt.Println(v2, v2)
                    for fi85 := 0; fi85 < 18; fi85++ {
                        var v86 bool = (!true)
                        _ = fi85
                        _ = v86
                    }
                    _ = fi83
                    _ = v84
                }
            case 9223372036854775807:
                if (!false) {
                    if (114 != 1) {
                        v87 := GenS2{F0: 1, F1: true, F2: 275}
                        _ = v87
                    }
                    v5 = v5
                    v88 := GenS9{F0: 22}
                    _ = v88
                }
            default:
                v65 = 255
            }
            switch 32767 {
            case 425:
                switch v2 {
                case "foo bar":
                    switch v2 {
                    case "foo bar":
                        fmt.Println(v5)
                    case "":
                        fmt.Println(v65)
                    default:
                        v89 := GenS10{F0: v5, F1: 451}
                        _ = v89
                    }
                case "":
                    if true {
                        v2 = ("\n\t" + "hello")
                    }
                }
            }
            _ = v82
        } else {
            var v90 int16 = 443
            v5 += v5
            switch "a" {
            case "foo bar":
                v90 += (v90 >> ((4294967295) & 63))
            case "a":
                for fi91 := 0; fi91 < 9; fi91++ {
                    for fi92 := 0; fi92 < 2; fi92++ {
                        var v93 int64 = 1
                        if false {
                            v65 = v65
                            var v94 bool = (v2 >= v2)
                            var v95 bool = true
                            _ = v94
                            _ = v95
                        }
                        delete(v4, (("" + "foo bar") + v2))
                        _ = fi92
                        _ = v93
                    }
                    v65 -= v5
                    if false {
                        switch (fi91 - (fi91 | fi91)) {
                        case -304:
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                        default:
                            if (0 > 1) {
                                v65 = v65
                                v96 := GenS11{F0: 426, F1: 0}
                                var v97 uint64 = 0
                                _ = v96
                                _ = v97
                            } else {
                                v98 := GenS8{F0: 0, F1: 2}
                                fmt.Println(v2, v90)
                                fmt.Println(v5, v2, v90)
                                _ = v98
                            }
                        }
                        if false {
                            delete(v4, "hello")
                            switch (-(fi91)) {
                            case 753:
                                fmt.Println(v90, v90)
                            case -2147483648:
                                v2 = "a"
                            case 680:
                                var v99 float64 = 0.1
                                _ = v99
                            default:
                                _ = len(v1)
                            }
                            for fi100 := 0; fi100 < 13; fi100++ {
                                v5 -= v5
                                if (!true) {
                                    v101 := GenS12{F0: 495}
                                    _ = v101
                                }
                                _ = fi100
                            }
                        }
                        if (!(!(v2 <= ""))) {
                            switch fi91 {
                            case 128:
                                switch fi91 {
                                case 9223372036854775807:
                                    for fi102 := 0; fi102 < 5; fi102++ {
                                        v103 := GenS13{F0: 42, F1: 0}
                                        v65 = (v65 * v5)
                                        v104 := GenS14{F0: 65535, F1: 1, F2: -514}
                                        _ = fi102
                                        _ = v103
                                        _ = v104
                                    }
                                }
                            case 0:
                                var v105 int32 = -1
                                _ = v105
                            case 32767:
                                v106 := GenS15{F0: 4294967295}
                                _ = v106
                            default:
                                if (234 == 1) {
                                    v90 = v90
                                    for fi107 := 0; fi107 < 12; fi107++ {
                                        v90 += (v90 >> ((1) & 63))
                                        _ = fi107
                                    }
                                } else {
                                    var v108 int32 = 761
                                    v109 := GenS2{F0: 390, F1: true, F2: 0}
                                    _ = v108
                                    _ = v109
                                }
                            }
                            if false {
                                fmt.Println(v65, v5, fi91)
                                v5 = (v5 >> ((84) & 63))
                            }
                            v110 := GenS16{F0: 255, F1: 2, F2: 80}
                            _ = v110
                        } else {
                            switch fi91 {
                            case -903:
                                switch (fi91 ^ -2147483648) {
                                case 2147483647:
                                    fmt.Println(v2)
                                default:
                                    var v111 int = (fi91 + 893)
                                    _ = v111
                                }
                            }
                            v90 -= 42
                            var v112 uint16 = 255
                            _ = v112
                        }
                    } else {
                        var v113 int64 = 9223372036854775807
                        fmt.Println(fi91, fi91, v65)
                        _ = v113
                    }
                    _ = fi91
                }
            default:
                switch 9223372036854775807 {
                case 128:
                    var v114 string = (("" + "hello") + "a")
                    _ = v114
                case -573:
                    if len(v3) > 0 {
                        v3[0] = (v2 + (v2 + ("" + v2)))
                    }
                case 467:
                    v90 += (v90 ^ (v90 / ((1000) | 1)))
                }
            }
            _ = v90
        }
    }
    v5 = v5
    v1 = append(v1, v2)
    if (-2.25 >= (0.1 / 1.0)) {
        var v115 uint16 = 1
        for fi116 := 0; fi116 < 16; fi116++ {
            v65 = 127
            for fi117 := 0; fi117 < 18; fi117++ {
                _ = len(v4)
                _ = len(v1)
                var v118 uint16 = v115
                if (1 <= 2147483648) {
                    if (false && false) {
                        var v119 float32 = 0.1
                        v118 += 28
                        fmt.Println(v115, v65)
                        var v120 int64 = 9223372036854775807
                        _ = v119
                        _ = v120
                    }
                    switch v2 {
                    case "foo bar":
                        var v121 uint = 62
                        _ = v121
                    default:
                        var v122 uint64 = 1
                        _ = v122
                    }
                    switch (v2 + v2) {
                    case "\n\t":
                        fmt.Println(v65)
                    }
                } else {
                    v65 = 1
                }
                _ = fi117
                _ = v118
            }
            _ = fi116
        }
        v5 -= (v5 * (v5 / ((0) | 1)))
        _ = v115
    } else {
        var v123 int16 = 1000
        _ = v123
    }
    v124 := GenS17{F0: 42, F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v124
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v65
}
