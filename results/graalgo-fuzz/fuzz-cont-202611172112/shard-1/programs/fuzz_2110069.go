package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint16) uint32 {
    v7 := GenS1{F0: 404, F1: 42}
    v8 := GenS1{F0: 748, F1: 18446744073709551615}
    var v9 float64 = (-2.25 / 1e10)
    v10 := GenS1{F0: 42, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 567
}

func main() {
    var v1 uint8 = 127
    var v2 int = 9223372036854775807
    var v3 float64 = float64(1e10)
    var v4 uint8 = (v1 << ((1) & 63))
    var v5 float64 = v3
    var v6 uint16 = 32767
    v4 = (v1 << ((2) & 63))
    var v11 int32 = 2147483647
    v12 := []int{v2}
    if (v11 != v11) {
        if true {
            for fi13 := 0; fi13 < 2; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 string = "\n\t"
                    _ = fi14
                    _ = v15
                }
                v5 = float64((3.14159 / ((1.5 / 1e10) * 0.1)))
                var v16 int32 = -911
                _ = fi13
                _ = v16
            }
            if true {
                if true {
                    _ = len(v12)
                    var v17 int64 = -1
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 int8 = 86
                        for fi20 := 0; fi20 < 4; fi20++ {
                            var v21 float64 = -2.25
                            var v22 string = ("hello" + "hello")
                            var v23 int8 = (v19 % ((v19) | 1))
                            var v24 float32 = 0.1
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        switch (^(v2)) {
                        case 127:
                            if true {
                                switch "a" {
                                case "a":
                                    for fi25 := 0; fi25 < 4; fi25++ {
                                        if false {
                                            var v26 int = fi18
                                            _ = v26
                                        }
                                        _ = fi25
                                    }
                                case "foo bar":
                                    fmt.Println(v19, v4, v4)
                                default:
                                    fmt.Println(v4)
                                }
                                _ = len(v12)
                            } else {
                                for fi27 := 0; fi27 < 19; fi27++ {
                                    v4 -= (^(v1))
                                    v12 = append(v12, 42)
                                    v19 += (-(v19))
                                    _ = fi27
                                }
                                for fi28 := 0; fi28 < 12; fi28++ {
                                    var v29 float32 = 0.0
                                    fmt.Println(v6)
                                    if true {
                                        _ = len(v12)
                                    } else {
                                        switch "hello" {
                                        case "foo bar":
                                            if (("\n\t" + ("" + (("" + "") + "foo bar"))) != ("foo bar" + "a")) {
                                                v2 = (v2 * fi18)
                                            } else {
                                                v1 -= v4
                                                v11 -= (v11 - v11)
                                                v4 -= v4
                                            }
                                        case "hello":
                                            var v30 string = "foo bar"
                                            _ = v30
                                        case "\n\t":
                                            v5 = ((v5 / v3) - -2.25)
                                        }
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                v1 -= (v1 >> ((18446744073709551615) & 63))
                            }
                        case 14:
                            var v31 bool = ((("a" + "hello") + ("\n\t" + "a")) >= ("a" + ("foo bar" + ("a" + "a"))))
                            _ = v31
                        case 243:
                            if len(v12) > 0 {
                                v12[0] = (fi18 ^ int(v19))
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                    switch v2 {
                    case -1:
                        var v32 uint16 = (v6 + v6)
                        _ = v32
                    default:
                        if ("" > "\n\t") {
                            v2 = v2
                            for fi33 := 0; fi33 < 11; fi33++ {
                                if (-203 <= 0) {
                                    if len(v12) > 0 {
                                        v12[0] = (v2 % ((fi33) | 1))
                                    }
                                    v4 = (v4 + (v4 % ((v4) | 1)))
                                } else {
                                    var v34 int = (v2 << ((18446744073709551615) & 63))
                                    var v35 int8 = -100
                                    v12 = append(v12, (v34 / ((fi33) | 1)))
                                    _ = v34
                                    _ = v35
                                }
                                _ = len(v12)
                                var v36 int16 = 1
                                _ = fi33
                                _ = v36
                            }
                            var v37 uint = 2
                            _ = v37
                        } else {
                            v6 += v6
                            fmt.Println(v1)
                        }
                    }
                }
                for fi38 := 0; fi38 < 9; fi38++ {
                    if true {
                        v2 -= 4294967295
                        if (!false) {
                            v12 = append(v12, 218)
                            var v39 uint64 = 42
                            var v40 int32 = (v11 % (((-(v11))) | 1))
                            _ = v39
                            _ = v40
                        } else {
                            v41 := GenS2{F0: 4294967295, F1: v2}
                            if ((!(1 <= 65535)) || (((v2 | fi38) != 128) || (32767 != 1000))) {
                                fmt.Println(v2, v2)
                                v42 := GenS3{F0: (^(v1)), F1: (v5 >= v3), F2: 42}
                                if true {
                                    if true {
                                        for fi43 := 0; fi43 < 3; fi43++ {
                                            v3 -= v3
                                            _ = fi43
                                        }
                                        for fi44 := 0; fi44 < 14; fi44++ {
                                            v45 := GenS4{F0: 1, F1: 0, F2: "a"}
                                            var v46 uint8 = (v1 ^ (v4 - 0))
                                            _ = fi44
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                    var v47 uint = 2
                                    _ = v47
                                } else {
                                    if false {
                                        if (2 > v4) {
                                            _ = len(v12)
                                        }
                                    } else {
                                        var v48 uint32 = 2
                                        _ = v48
                                    }
                                    switch v2 {
                                    case -2147483648:
                                        _ = len(v12)
                                    case -960:
                                        _ = len(v12)
                                    default:
                                        v5 = v3
                                    }
                                }
                                _ = v42
                            } else {
                                v1 += v1
                                var v49 int64 = 0
                                for fi50 := 0; fi50 < 4; fi50++ {
                                    v51 := GenS3{F0: v1, F1: (1000 <= 100), F2: 124}
                                    v12 = append(v12, (fi38 & 90))
                                    v12 = append(v12, 14)
                                    _ = fi50
                                    _ = v51
                                }
                                _ = v49
                            }
                            _ = v41
                        }
                        switch "\n\t" {
                        case "foo bar":
                            switch "" {
                            case "hello":
                                v12 = append(v12, v2)
                            case "\n\t":
                                var v52 int8 = 29
                                _ = v52
                            case "a":
                                for fi53 := 0; fi53 < 12; fi53++ {
                                    var v54 float32 = 3.14159
                                    if true {
                                        switch "\n\t" {
                                        case "a":
                                            var v55 int16 = 32767
                                            _ = v55
                                        case "":
                                            v4 += (v4 | 127)
                                        default:
                                            v12 = append(v12, 32768)
                                            fmt.Println(v11, v4, v3)
                                        }
                                        fmt.Println(v54)
                                        var v56 uint16 = (v6 ^ 214)
                                        _ = v56
                                    }
                                    v4 += 42
                                    _ = fi53
                                    _ = v54
                                }
                            }
                        }
                    } else {
                        if len(v12) > 0 {
                            v12[0] = fi38
                        }
                        v1 += (v4 & (^(v1)))
                        var v57 uint8 = v4
                        _ = v57
                    }
                    _ = fi38
                }
                var v58 int8 = 100
                _ = v58
            }
        }
        for fi59 := 0; fi59 < 12; fi59++ {
            v4 = (v4 | (v4 & v1))
            switch int(v6) {
            case 894:
                var v60 int64 = -636
                _ = v60
            case 128:
                v4 = v4
            case 32768:
                if (fi59 != fi59) {
                    if ("a" <= (("hello" + "a") + "a")) {
                        if len(v12) > 0 {
                            v12[0] = (fi59 / (((fi59 | v2)) | 1))
                        }
                        fmt.Println(v11, v3)
                    }
                    var v61 uint32 = 65535
                    v62 := GenS5{F0: v4}
                    _ = v61
                    _ = v62
                } else {
                    fmt.Println(v1, v6, v11)
                    if len(v12) > 0 {
                        v12[0] = (^(v2))
                    }
                }
            }
            _ = fi59
        }
    } else {
        var v63 float64 = v3
        _ = v63
    }
    for fi64 := 0; fi64 < 3; fi64++ {
        if (1 == 4294967295) {
            var v65 int = (^(fi64))
            switch v2 {
            case 32767:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            default:
                var v66 float32 = 0.1
                _ = v66
            }
            v67 := GenS5{F0: 127}
            _ = v65
            _ = v67
        } else {
            if (!true) {
                var v68 uint = 439
                if true {
                    for fi69 := 0; fi69 < 10; fi69++ {
                        v6 -= v6
                        switch "\n\t" {
                        case "a":
                            var v70 bool = (0.0 == (((((((-2.25 - 1.5) - 3.14159) / -2.25) * (float32(v3) / 1.0)) / (0.0 * float32((v5 / v5)))) + 1e10) / 1.5))
                            _ = v70
                        case "\n\t":
                            var v71 bool = (v1 <= (v1 % (((v1 + (v4 % (((^(v4))) | 1)))) | 1)))
                            _ = v71
                        default:
                            v2 -= (fi69 & v2)
                        }
                        _ = fi69
                    }
                    switch ("a" + ("a" + ("foo bar" + ("a" + "\n\t")))) {
                    case "hello":
                        fmt.Println(v11, v3, v68)
                    case "a":
                        v2 = v2
                    case "\n\t":
                        if (!(42 >= genFn0(432, (v6 | 32767)))) {
                            var v72 uint32 = 707
                            if false {
                                for fi73 := 0; fi73 < 9; fi73++ {
                                    var v74 float64 = v5
                                    v3 += -2.25
                                    switch -9223372036854775808 {
                                    case 0:
                                        if (((v3 + v3) - (float64(3.14159) - v3)) < v5) {
                                            v75 := GenS5{F0: v1}
                                            _ = v75
                                        }
                                    default:
                                        if len(v12) > 0 {
                                            v12[0] = v2
                                        }
                                    }
                                    _ = fi73
                                    _ = v74
                                }
                            } else {
                                fmt.Println(v4, v6, v4)
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                            var v76 uint8 = (v1 + (v1 | v4))
                            _ = v72
                            _ = v76
                        } else {
                            if (-115 >= 100) {
                                for fi77 := 0; fi77 < 1; fi77++ {
                                    fmt.Println(v68, v5, v4)
                                    if false {
                                        var v78 float64 = v3
                                        v11 -= -1
                                        fmt.Println(v68, v1, fi64)
                                        _ = v78
                                    } else {
                                        fmt.Println(v3, v6)
                                    }
                                    _ = fi77
                                }
                            } else {
                                v4 = 1
                                switch v2 {
                                case 9223372036854775807:
                                    if (!true) {
                                        for fi79 := 0; fi79 < 17; fi79++ {
                                            v1 = v1
                                            v80 := GenS6{F0: int8(v1), F1: (^(v2))}
                                            v12 = append(v12, fi64)
                                            _ = fi79
                                            _ = v80
                                        }
                                        v81 := GenS4{F0: uint16(v1), F1: 2, F2: ("hello" + "")}
                                        _ = v81
                                    } else {
                                        var v82 uint16 = (^(v6))
                                        var v83 uint64 = 42
                                        v12 = append(v12, fi64)
                                        _ = v82
                                        _ = v83
                                    }
                                case 1:
                                    v84 := GenS6{F0: 42, F1: v2}
                                    _ = v84
                                }
                                var v85 string = ("\n\t" + (("hello" + "a") + "\n\t"))
                                _ = v85
                            }
                        }
                    default:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    }
                    switch ("hello" + "hello") {
                    case "":
                        if (!(!((9223372036854775807 <= 0) && (false || false)))) {
                            for fi86 := 0; fi86 < 8; fi86++ {
                                fmt.Println(v11, v6)
                                _ = len(v12)
                                v5 -= (3.14159 + v3)
                                v3 = ((1e10 + 3.14159) - v5)
                                if (2147483647 == genFn0(-1, (v6 ^ v6))) {
                                    v68 += (v68 % (((v68 >> (((^(v68))) & 63))) | 1))
                                } else {
                                    switch (v2 * -2147483648) {
                                    case 32767:
                                        v11 -= v11
                                    }
                                }
                                v68 -= 4294967295
                                v87 := GenS1{F0: 1, F1: v68}
                                var v88 int8 = 100
                                _ = fi86
                                _ = v87
                                _ = v88
                            }
                        } else {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        }
                    }
                } else {
                    if false {
                        for fi89 := 0; fi89 < 7; fi89++ {
                            var v90 int32 = (v11 * v11)
                            fmt.Println(fi64, fi89, v3)
                            if false {
                                v90 = int32(v1)
                            } else {
                                var v91 int32 = (v90 >> ((v68) & 63))
                                var v92 int16 = 42
                                v93 := GenS5{F0: 2}
                                _ = v91
                                _ = v92
                                _ = v93
                            }
                            _ = fi89
                            _ = v90
                        }
                        fmt.Println(v68)
                        var v94 int16 = 1000
                        _ = v94
                    } else {
                        for fi95 := 0; fi95 < 19; fi95++ {
                            fmt.Println(v4, v6)
                            if true {
                                var v96 string = ("a" + "")
                                if len(v12) > 0 {
                                    v12[0] = 92
                                }
                                switch fi64 {
                                case 33:
                                    switch fi95 {
                                    case -241:
                                        v68 = v68
                                    case 503:
                                        v97 := GenS7{F0: -504, F1: ""}
                                        _ = v97
                                    default:
                                        var v98 int16 = -159
                                        _ = v98
                                    }
                                default:
                                    if (!true) {
                                        v2 -= fi95
                                        if (false && false) {
                                            v99 := GenS5{F0: v4}
                                            var v100 uint64 = 796
                                            _ = v100
                                            _ = v99
                                        }
                                    }
                                }
                                _ = v96
                            } else {
                                v2 = 1
                                v5 += (float64(((-2.25 * 1.5) + 1.5)) + float64(0.1))
                                v101 := GenS8{F0: 850}
                                v102 := GenS8{F0: (-(v11))}
                                _ = v101
                                _ = v102
                            }
                            var v103 float64 = v3
                            if true {
                                v104 := GenS6{F0: 127, F1: fi64}
                                var v105 string = (("hello" + "\n\t") + "a")
                                _ = v104
                                _ = v105
                            } else {
                                var v106 int64 = 0
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                _ = v106
                            }
                            var v107 int8 = -71
                            _ = len(v12)
                            var v108 string = (("foo bar" + "hello") + (("hello" + ("foo bar" + (("hello" + "hello") + "hello"))) + "\n\t"))
                            fmt.Println(v108, v6, v107)
                            switch "" {
                            case "\n\t":
                                var v109 uint8 = (v1 / ((v1) | 1))
                                _ = v109
                            }
                            _ = fi95
                            _ = v103
                            _ = v107
                            _ = v108
                        }
                    }
                    v11 = (v11 >> ((v68) & 63))
                }
                _ = v68
            } else {
                if true {
                    var v110 uint8 = (v1 % ((uint8(v6)) | 1))
                    _ = v110
                }
                var v111 uint8 = 2
                _ = v111
            }
            v11 = -2147483648
            var v112 uint16 = (v6 / ((255) | 1))
            _ = v112
        }
        if false {
            v5 += ((((float64((0.0 * ((0.1 + 0.0) - (((1.5 / (float32(v5) * ((0.0 - float32(v3)) / 1.0))) / -2.25) + 1.5)))) / 0.1) / (float64(((3.14159 + -2.25) + 0.0)) + ((v3 - 0.1) + v5))) / ((-2.25 / v3) - ((float64(float32(((v3 - (1.5 + (v3 * ((float64((1.5 * 1.5)) * (1e10 - float64(0.1))) + float64(1.5))))) / 1e10))) / v5) + v5))) + (v5 / v3))
            _ = len(v12)
        } else {
            if true {
                var v113 int64 = 9223372036854775807
                _ = v113
            }
            for fi114 := 0; fi114 < 9; fi114++ {
                v3 = v3
                var v115 float32 = 3.14159
                _ = fi114
                _ = v115
            }
        }
        _ = fi64
    }
    for fi116 := 0; fi116 < 18; fi116++ {
        var v117 uint32 = 91
        v118 := GenS9{F0: 0}
        for fi119 := 0; fi119 < 3; fi119++ {
            if (!true) {
                v3 = (3.14159 / v5)
                v120 := GenS6{F0: -1, F1: fi119}
                v121 := GenS10{F0: int(v117)}
                _ = v120
                _ = v121
            } else {
                switch (("hello" + "foo bar") + ("" + (((("" + (((((((("hello" + "hello") + "") + "a") + (("" + "\n\t") + "foo bar")) + "") + "hello") + "foo bar") + "")) + "hello") + "hello") + "\n\t"))) {
                case "hello":
                    v4 -= v4
                }
                var v122 string = "\n\t"
                _ = v122
            }
            fmt.Println(v1, v11, v6)
            _ = fi119
        }
        _ = fi116
        _ = v117
        _ = v118
    }
    var v123 bool = true
    if v123 {
        v6 += v6
        for fi124 := 0; fi124 < 18; fi124++ {
            if v123 {
                fmt.Println(v5, v3)
                var v125 bool = (!false)
                var v126 uint64 = 1
                _ = v125
                _ = v126
            } else {
                if true {
                    if ((v6 ^ uint16(v1)) != (v6 - v6)) {
                        for fi127 := 0; fi127 < 10; fi127++ {
                            v6 += 894
                            v128 := GenS4{F0: (v6 & uint16(v4)), F1: uint64(v4), F2: ("" + "a")}
                            v4 += 128
                            _ = fi127
                            _ = v128
                        }
                        var v129 int16 = 32767
                        _ = v129
                    } else {
                        if (-2.25 > float64(((-2.25 - 0.1) - ((1.5 / -2.25) * float32((1.5 + (((v5 + 1e10) + v3) + v3))))))) {
                            var v130 uint32 = 3
                            for fi131 := 0; fi131 < 2; fi131++ {
                                switch v2 {
                                case 9223372036854775807:
                                    switch ("\n\t" + "a") {
                                    case "\n\t":
                                        v6 = 1
                                    }
                                default:
                                    v132 := GenS11{F0: (v4 / ((255) | 1))}
                                    _ = v132
                                }
                                _ = fi131
                            }
                            _ = len(v12)
                            _ = v130
                        }
                        var v133 uint = 1
                        _ = v133
                    }
                    v5 = 1e10
                    for fi134 := 0; fi134 < 1; fi134++ {
                        var v135 uint64 = 0
                        var v136 uint16 = 2
                        if true {
                            if false {
                                v1 = v4
                                var v137 int16 = 884
                                if len(v12) > 0 {
                                    v12[0] = (^(fi134))
                                }
                                _ = v137
                            }
                            var v138 int = fi134
                            var v139 uint32 = genFn0(-416, v6)
                            _ = v138
                            _ = v139
                        }
                        var v140 uint8 = (v1 % (((^(v4))) | 1))
                        _ = fi134
                        _ = v135
                        _ = v136
                        _ = v140
                    }
                } else {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                fmt.Println(v2)
            }
            if len(v12) > 0 {
                v12[0] = fi124
            }
            _ = fi124
        }
        for fi141 := 0; fi141 < 10; fi141++ {
            v142 := GenS1{F0: 807, F1: 731}
            var v143 uint = 255
            for fi144 := 0; fi144 < 6; fi144++ {
                var v145 uint64 = 0
                _ = fi144
                _ = v145
            }
            _ = fi141
            _ = v142
            _ = v143
        }
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
