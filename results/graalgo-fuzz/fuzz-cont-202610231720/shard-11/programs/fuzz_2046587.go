package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS7 struct { F0 bool; F1 int; F2 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int; F1 int16 }
type GenS10 struct { F0 string; F1 uint32 }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int64, p1 int16, p2 uint) int32 {
    var v6 int64 = p0
    var v7 uint64 = 42
    v8 := []int{}
    var v9 int32 = 895
    _ = len(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []int{0, 807, -9223372036854775808}
    v2 := []string{"foo bar", "\n\t"}
    v3 := make(map[string]string)
    var v4 uint16 = 65535
    v5 := make(map[uint32]int)
    var v10 bool = ("a" <= ("" + ""))
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int16 = 1000
        _ = fi11
        _ = v12
    }
    var v13 int16 = 1
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v13, v4)
        if true {
            fmt.Println(v10)
            switch -1 {
            case -754:
                var v15 int32 = 2147483647
                _ = v15
            case 334:
                switch "" {
                case "\n\t":
                    v16 := GenS1{F0: 42}
                    _ = v16
                default:
                    if true {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            v18 := GenS1{F0: fi17}
                            if (2 == 2) {
                                var v19 int32 = 2147483647
                                _ = v19
                            }
                            v20 := GenS1{F0: fi14}
                            _ = fi17
                            _ = v18
                            _ = v20
                        }
                        fmt.Println(v10)
                    } else {
                        v10 = (v10 && (-1 > 207))
                        var v21 int16 = int16(v4)
                        _ = v21
                    }
                }
            }
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 uint16 = v4
                v10 = (0.0 != 0.1)
                v24 := GenS1{F0: (fi22 | (fi14 | (fi14 * fi22)))}
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 int16 = (v13 % (((v13 << ((4294967295) & 63))) | 1))
                delete(v5, 457)
                v27 := GenS2{F0: 1}
                v28 := GenS3{F0: v4, F1: 2}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            v2 = append(v2, "hello")
        }
        if (!false) {
            if ((fi14 & 128) != (fi14 + int(v13))) {
                v10 = v10
                var v29 bool = v10
                if (4294967295 < 42) {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                } else {
                    var v30 int16 = (v13 + -967)
                    var v31 uint8 = 2
                    _ = v30
                    _ = v31
                }
                _ = v29
            } else {
                if (v10 && v10) {
                    v32 := GenS3{F0: (^(v4)), F1: 154}
                    _ = len(v2)
                    if v10 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        if (42 < 941) {
                            v13 = v13
                        }
                        v33 := GenS4{F0: ("" + "hello"), F1: 483}
                        fmt.Println(v10, v4, fi14)
                        _ = v33
                    }
                    _ = v32
                }
                fmt.Println(v13, v4)
                if (!(100 > -82)) {
                    delete(v5, 391)
                    var v34 int64 = -308
                    _ = v34
                }
            }
            v10 = v10
            v35 := GenS5{F0: v13, F1: (fi14 / ((4294967295) | 1))}
            _ = v35
        } else {
            v36 := GenS3{F0: v4, F1: 1}
            _ = v36
        }
        _ = fi14
    }
    fmt.Println(v10, v4, v4)
    switch "\n\t" {
    case "foo bar":
        v10 = false
    case "\n\t":
        switch 127 {
        case 32767:
            v37 := GenS3{F0: 42, F1: 479}
            _ = v37
        case 89:
            switch 32768 {
            case 127:
                var v38 float64 = 0.0
                _ = v38
            case 42:
                for fi39 := 0; fi39 < 8; fi39++ {
                    _ = len(v5)
                    for fi40 := 0; fi40 < 6; fi40++ {
                        v3["a"] = "a"
                        v4 -= 42
                        v13 += 32767
                        _ = fi40
                    }
                    v41 := GenS4{F0: "", F1: -19}
                    _ = fi39
                    _ = v41
                }
            case 1:
                var v42 float64 = (((0.1 / 0.1) / float64((3.14159 - ((-2.25 + (0.1 + ((1e10 - 3.14159) + (-2.25 / 1.0)))) * ((1e10 / (3.14159 * -2.25)) - 1.5))))) * (1e10 + -2.25))
                _ = v42
            default:
                if (0 >= 0) {
                    if (!v10) {
                        var v43 int64 = -1
                        _ = v43
                    }
                    delete(v5, 1)
                }
            }
        }
    case "hello":
        var v44 int16 = v13
        _ = v44
    default:
        v13 += (v13 | 522)
    }
    if v10 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        for fi45 := 0; fi45 < 12; fi45++ {
            var v46 uint = 42
            for fi47 := 0; fi47 < 19; fi47++ {
                v3[("a" + "\n\t")] = ("foo bar" + ((("hello" + ("\n\t" + "foo bar")) + "") + ""))
                var v48 int64 = 42
                if ("a" == ("" + "foo bar")) {
                    if ((("foo bar" + ("" + "\n\t")) + ("\n\t" + "\n\t")) > "hello") {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v13 = 42
                        var v49 uint16 = v4
                        var v50 uint8 = 0
                        _ = v49
                        _ = v50
                    } else {
                        v51 := v3[("hello" + ("hello" + ""))]
                        _ = v51
                    }
                    if (!false) {
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(fi45, v10)
                        case "":
                            var v52 bool = (fi47 == (fi45 | fi47))
                            _ = v52
                        case "hello":
                            var v53 uint8 = 255
                            _ = v53
                        }
                        var v54 uint64 = 0
                        switch (("\n\t" + "hello") + "\n\t") {
                        case "foo bar":
                            var v55 float64 = 1.5
                            _ = v55
                        case "\n\t":
                            _ = len(v5)
                        case "hello":
                            var v56 int64 = (v48 - -40)
                            _ = v56
                        }
                        _ = v54
                    }
                }
                _ = fi47
                _ = v48
            }
            v57 := GenS3{F0: v4, F1: 2147483647}
            _ = fi45
            _ = v46
            _ = v57
        }
        var v58 uint = 255
        _ = v58
    }
    if (0.0 == (-2.25 - -2.25)) {
        switch 42 {
        case -1:
            v10 = (2147483647 >= 42)
        case 2147483647:
            v10 = (0 != genFn0(42, 0, 18446744073709551615))
        }
    } else {
        switch -1 {
        case 128:
            if v10 {
                for fi59 := 0; fi59 < 11; fi59++ {
                    for fi60 := 0; fi60 < 7; fi60++ {
                        var v61 uint16 = (v4 * v4)
                        switch fi60 {
                        case 754:
                            v61 = v4
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v61 = v61
                        case "hello":
                            _ = len(v3)
                        default:
                            var v62 float64 = (1e10 * 1.5)
                            _ = v62
                        }
                        v4 += (v61 >> ((65535) & 63))
                        _ = fi60
                        _ = v61
                    }
                    _ = fi59
                }
                var v63 int16 = 736
                _ = v63
            }
        case -87:
            switch "a" {
            case "a":
                if (3.14159 < 1e10) {
                    fmt.Println(v13, v13, v10)
                    var v64 int32 = genFn0(-848, -32768, 255)
                    _ = v64
                }
            }
        default:
            v4 -= (^(v4))
        }
        v65 := GenS2{F0: -9223372036854775808}
        if true {
            for fi66 := 0; fi66 < 11; fi66++ {
                var v67 uint = 42
                var v68 int16 = v13
                _ = fi66
                _ = v67
                _ = v68
            }
            fmt.Println(v13, v10)
            var v69 uint8 = 42
            _ = v69
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v70 := GenS6{F0: 18446744073709551615, F1: 4294967295, F2: (v4 | (^(v4)))}
            _ = v70
        }
        _ = v65
    }
    if ((-2.25 >= -2.25) && v10) {
        v71 := GenS3{F0: (v4 % (((v4 ^ (v4 ^ (v4 | (v4 + v4))))) | 1)), F1: 42}
        var v72 int32 = 0
        if v10 {
            fmt.Println(v13, v4, v4)
            v4 = v4
        } else {
            for fi73 := 0; fi73 < 16; fi73++ {
                v10 = v10
                switch fi73 {
                case -1:
                    var v74 int8 = 17
                    _ = v74
                case 0:
                    v5[0] = (fi73 << ((42) & 63))
                case -9223372036854775808:
                    var v75 int64 = -254
                    _ = v75
                default:
                    for fi76 := 0; fi76 < 4; fi76++ {
                        v77 := GenS4{F0: "\n\t", F1: -148}
                        fmt.Println(fi73)
                        if v10 {
                            v78 := GenS7{F0: v10, F1: 712, F2: -1}
                            var v79 float32 = (1e10 + (1.5 + (float32(((0.1 * 0.1) / float64(float32(3.14159)))) + 1e10)))
                            v2 = append(v2, "foo bar")
                            _ = v78
                            _ = v79
                        }
                        _ = fi76
                        _ = v77
                    }
                }
                _ = fi73
            }
        }
        _ = v71
        _ = v72
    } else {
        v10 = ((-994 < -1) || v10)
        for fi80 := 0; fi80 < 13; fi80++ {
            if (300 != 1) {
                var v81 bool = (!(159 == 42))
                _ = v81
            } else {
                switch (fi80 & fi80) {
                case -97:
                    var v82 bool = (!(108 > 0))
                    _ = v82
                case 1000:
                    if (0.0 > (0.0 / 3.14159)) {
                        v13 += v13
                        fmt.Println(fi80, fi80)
                        var v83 int8 = 0
                        _ = v83
                    }
                case 4294967295:
                    switch "a" {
                    case "a":
                        v13 += (v13 % ((-1) | 1))
                    case "foo bar":
                        fmt.Println(v4, fi80)
                    case "\n\t":
                        fmt.Println(fi80, fi80, v13)
                    }
                default:
                    if (572 != (v4 & (v4 << ((94) & 63)))) {
                        var v84 int8 = 1
                        if true {
                            var v85 int64 = -306
                            switch (fi80 * int(v4)) {
                            case 128:
                                v5[4294967295] = fi80
                            case -937:
                                v86 := GenS7{F0: ((^(v13)) != v13), F1: fi80, F2: 100}
                                _ = v86
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(fi80)
                                default:
                                    for fi87 := 0; fi87 < 7; fi87++ {
                                        switch "" {
                                        case "\n\t":
                                            var v88 uint = 1
                                            _ = v88
                                        case "":
                                            if (!("foo bar" == "")) {
                                                var v89 uint64 = 0
                                                fmt.Println(v85, v13, v4)
                                                fmt.Println(v85, v13, v10)
                                                _ = v89
                                            }
                                        }
                                        _ = fi87
                                    }
                                }
                            }
                            switch "" {
                            case "a":
                                for fi90 := 0; fi90 < 11; fi90++ {
                                    fmt.Println(v10)
                                    switch fi90 {
                                    case 4294967295:
                                        var v91 int16 = v13
                                        _ = v91
                                    }
                                    v3[(("foo bar" + ("foo bar" + "foo bar")) + ("" + ""))] = ("foo bar" + "a")
                                    var v92 bool = false
                                    var v93 int64 = (v85 - (v85 | (v85 ^ 9223372036854775807)))
                                    _ = fi90
                                    _ = v92
                                    _ = v93
                                }
                            default:
                                v4 -= v4
                            }
                            _ = v85
                        }
                        var v94 string = "hello"
                        _ = v84
                        _ = v94
                    } else {
                        if v10 {
                            var v95 int = fi80
                            var v96 bool = (false && v10)
                            if true {
                                var v97 int8 = 72
                                _ = v97
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = fi80
                                }
                                v10 = (42 >= -554)
                                v98 := GenS8{F0: 1000}
                                _ = v98
                            }
                            _ = v95
                            _ = v96
                        } else {
                            for fi99 := 0; fi99 < 10; fi99++ {
                                _ = len(v5)
                                fmt.Println(fi80, v13)
                                v13 -= (v13 ^ (v13 | (v13 & v13)))
                                _ = fi99
                            }
                            if v10 {
                                v13 = (v13 << ((341) & 63))
                            }
                            v13 = v13
                        }
                        if (-728 <= 799) {
                            var v100 int32 = -883
                            _ = v100
                        }
                    }
                }
            }
            if v10 {
                switch "a" {
                case "a":
                    var v101 int64 = 0
                    _ = v101
                default:
                    if false {
                        switch 925 {
                        case -9223372036854775808:
                            v102 := GenS1{F0: (fi80 | fi80)}
                            _ = v102
                        default:
                            var v103 uint = 0
                            _ = v103
                        }
                    }
                }
                var v104 uint16 = 1
                _ = v104
            } else {
                var v105 uint32 = 4294967295
                if true {
                    var v106 int = fi80
                    fmt.Println(v10, v13)
                    _ = v106
                }
                _ = v105
            }
            var v107 uint32 = 764
            _ = fi80
            _ = v107
        }
    }
    if ((160 <= 460) || v10) {
        for fi108 := 0; fi108 < 17; fi108++ {
            var v109 uint32 = 853
            switch fi108 {
            case 929:
                for fi110 := 0; fi110 < 6; fi110++ {
                    var v111 float64 = (float64(0.0) / 1.5)
                    _ = fi110
                    _ = v111
                }
            case 2147483647:
                for fi112 := 0; fi112 < 6; fi112++ {
                    var v113 int8 = 127
                    _ = fi112
                    _ = v113
                }
            case -2147483648:
                fmt.Println(v10, v10)
            default:
                var v114 int8 = 0
                _ = v114
            }
            _ = fi108
            _ = v109
        }
    }
    if (-1 >= -73) {
        if v10 {
            var v115 int = 32767
            if true {
                if (!((v4 >> ((622) & 63)) != (v4 >> ((18446744073709551615) & 63)))) {
                    v115 = (v115 / ((v115) | 1))
                    var v116 int64 = 582
                    var v117 int8 = 1
                    _ = v116
                    _ = v117
                }
                v10 = v10
            }
            v118 := GenS2{F0: 1}
            _ = v115
            _ = v118
        } else {
            v10 = false
        }
        v4 += v4
    }
    v119 := GenS9{F0: 32768, F1: v13}
    for fi120 := 0; fi120 < 8; fi120++ {
        fmt.Println(v13, v4)
        switch "" {
        case "a":
            v4 -= v4
        default:
            for fi121 := 0; fi121 < 3; fi121++ {
                fmt.Println(v10)
                if v10 {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                }
                switch "foo bar" {
                case "hello":
                    v122 := GenS10{F0: "\n\t", F1: 793}
                    _ = v122
                }
                _ = fi121
            }
        }
        fmt.Println(v10, v10)
        _ = fi120
    }
    v123 := GenS11{F0: uint8(v4)}
    switch ("\n\t" + ("\n\t" + ("a" + ("foo bar" + "a")))) {
    case "hello":
        v124 := v5[0]
        _ = v124
    case "":
        v125 := GenS7{F0: ((0.0 > 3.14159) && (v13 <= v13)), F1: -9223372036854775808, F2: 36}
        _ = v125
    }
    var v126 uint32 = 65535
    fmt.Println(v126)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v126:", v126)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v119
    _ = v123
    _ = v126
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
