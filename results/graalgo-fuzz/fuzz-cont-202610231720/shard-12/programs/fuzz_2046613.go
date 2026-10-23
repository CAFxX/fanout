package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int; F2 uint32 }
type GenS6 struct { F0 uint64; F1 int32; F2 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint64; F1 uint64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint16; F1 int8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint; F1 int8 }
type GenS18 struct { F0 bool; F1 uint16; F2 int64 }
type GenS19 struct { F0 uint16; F1 bool; F2 int64 }
func genFn0() int {
    if (42 >= 0) {
        v7 := []int{0, 2147483647, 1}
        var v8 uint32 = 163
        if ((-128 > 42) || false) {
            for fi9 := 0; fi9 < 12; fi9++ {
                v8 += v8
                _ = fi9
            }
            switch -231 {
            case 63:
                fmt.Println(v8)
            case 1:
                var v10 uint = 351
                _ = v10
            case 41:
                v11 := GenS1{F0: 1, F1: -932, F2: -32768}
                _ = v11
            default:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        }
        _ = v7
        _ = v8
    } else {
        var v12 string = "\n\t"
        fmt.Println(v12)
        v12 = v12
        _ = v12
    }
    if (!true) {
        v13 := []int{127, -935, -387}
        if (((!(!(false && (true || false)))) && ("foo bar" > "")) && (280 <= 4294967295)) {
            switch 4294967295 {
            case 53:
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 float64 = (0.1 * (1e10 - 1.5))
                    if len(v13) > 0 {
                        v13[0] = (fi14 + 32768)
                    }
                    var v16 uint16 = 65535
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case 127:
                _ = len(v13)
            default:
                var v17 int8 = 67
                _ = v17
            }
            var v18 int8 = -69
            _ = v18
        }
        _ = v13
    } else {
        if false {
            v19 := GenS2{F0: 9223372036854775807, F1: 1, F2: -671}
            _ = v19
        } else {
            if true {
                var v20 bool = true
                v21 := GenS2{F0: 1, F1: 873, F2: 42}
                _ = v20
                _ = v21
            } else {
                var v22 string = (("\n\t" + "hello") + ("a" + ("a" + "")))
                v23 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 232}
                _ = v22
                _ = v23
            }
        }
    }
    var v24 float64 = 1.5
    v25 := GenS3{F0: 314, F1: 0}
    fmt.Println(v24)
    _ = v24
    _ = v25
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 127
    v3 := []int{(-(v2)), v2}
    v4 := make(map[string]uint32)
    var v5 uint = 872
    v6 := []string{"foo bar", "foo bar", "hello", "a"}
    switch ("" + "foo bar") {
    case "":
        fmt.Println(v5, v5)
    case "a":
        switch "hello" {
        case "hello":
            v2 = (v2 << (((^(v5))) & 63))
        case "":
            v26 := GenS4{F0: 0}
            _ = v26
        case "foo bar":
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int8 = 1
                _ = fi27
                _ = v28
            }
        }
    case "foo bar":
        v5 -= v5
    }
    var v29 int64 = -9223372036854775808
    var v30 string = ""
    v31 := GenS5{F0: -1, F1: v2, F2: 4294967295}
    var v32 bool = (float32((float64(1e10) / 0.1)) >= 3.14159)
    if true {
        v33 := GenS4{F0: -19}
        _ = v33
    } else {
        if (v2 <= v2) {
            if ((1e10 / ((float32(1e10) + 0.1) / 1.5)) <= 1e10) {
                delete(v1, 2147483648)
            } else {
                var v34 uint16 = 204
                var v35 int = v2
                switch "\n\t" {
                case "\n\t":
                    v36 := GenS6{F0: 2, F1: 42, F2: -940}
                    _ = v36
                case "hello":
                    delete(v4, v30)
                }
                _ = v34
                _ = v35
            }
            if true {
                if (!v32) {
                    var v37 int = genFn0()
                    if (42 == 100) {
                        var v38 int64 = (v29 / (((v29 ^ v29)) | 1))
                        _ = v38
                    } else {
                        _ = len(v3)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi39 := 0; fi39 < 10; fi39++ {
                            v29 -= -1
                            var v40 string = "hello"
                            for fi41 := 0; fi41 < 15; fi41++ {
                                v5 += v5
                                var v42 int16 = 966
                                _ = fi41
                                _ = v42
                            }
                            _ = fi39
                            _ = v40
                        }
                    }
                    _ = v37
                } else {
                    switch (v2 % (((v2 * (v2 % ((genFn0()) | 1)))) | 1)) {
                    case 2147483647:
                        if len(v3) > 0 {
                            v3[0] = -9223372036854775808
                        }
                    case -373:
                        for fi43 := 0; fi43 < 11; fi43++ {
                            v5 -= v5
                            _ = fi43
                        }
                    default:
                        if (219 != -2147483648) {
                            var v44 uint8 = 34
                            if (v32 || false) {
                                v32 = v32
                                if (-1 == v29) {
                                    if v32 {
                                        fmt.Println(v5, v44)
                                        for fi45 := 0; fi45 < 9; fi45++ {
                                            var v46 uint32 = 2147483647
                                            v5 -= v5
                                            v44 += v44
                                            _ = fi45
                                            _ = v46
                                        }
                                    } else {
                                        for fi47 := 0; fi47 < 14; fi47++ {
                                            var v48 string = v30
                                            _ = fi47
                                            _ = v48
                                        }
                                        var v49 int64 = (v29 ^ v29)
                                        v29 = (v49 & v29)
                                        _ = v49
                                    }
                                } else {
                                    _ = len(v3)
                                    if v32 {
                                        for fi50 := 0; fi50 < 13; fi50++ {
                                            switch (v2 * -216) {
                                            case 127:
                                                var v51 uint32 = 65535
                                                _ = v51
                                            }
                                            if (((v30 + v30) > "foo bar") || ((-2147483648 == -2147483648) && (587 != 964))) {
                                                var v52 int8 = 102
                                                for fi53 := 0; fi53 < 19; fi53++ {
                                                    v5 = v5
                                                    if (false && (!(1 <= 677))) {
                                                        v54 := GenS4{F0: (v52 + -72)}
                                                        if (!v32) {
                                                            switch (fi53 | 493) {
                                                            case 128:
                                                                if ((-(v52)) > v52) {
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                } else {
                                                                    if (v52 == (-(v52))) {
                                                                        v5 = 255
                                                                    } else {
                                                                        v52 = (^(v52))
                                                                        var v55 int16 = 1000
                                                                        _ = v55
                                                                    }
                                                                    var v56 int = v2
                                                                    var v57 uint16 = 754
                                                                    v2 = (-(v2))
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            case 970:
                                                                if v32 {
                                                                    v5 = (v5 + v5)
                                                                    v29 = int64(v52)
                                                                    fmt.Println(v29)
                                                                } else {
                                                                    v32 = v32
                                                                    var v58 uint = 1
                                                                    v1[2147483648] = v30
                                                                    _ = v58
                                                                }
                                                            }
                                                        } else {
                                                            if v32 {
                                                                fmt.Println(v5)
                                                                var v59 int16 = -32768
                                                                if false {
                                                                    var v60 int16 = (v59 & (v59 * v59))
                                                                    var v61 int8 = v52
                                                                    _ = v60
                                                                    _ = v61
                                                                } else {
                                                                    var v62 string = ""
                                                                    if (-413 <= 42) {
                                                                        if v32 {
                                                                            switch v2 {
                                                                            case -831:
                                                                                v63 := GenS3{F0: 0, F1: 2147483647}
                                                                                _ = v63
                                                                            case 9223372036854775807:
                                                                                _ = len(v4)
                                                                                var v64 int = fi53
                                                                                _ = len(v4)
                                                                                _ = v64
                                                                            default:
                                                                                fmt.Println(v30)
                                                                            }
                                                                            v5 += (v5 >> (((v5 % ((255) | 1))) & 63))
                                                                        }
                                                                    } else {
                                                                        v62 = ((v62 + (("hello" + ("" + v62)) + "foo bar")) + ((v30 + v62) + "foo bar"))
                                                                    }
                                                                    v65 := GenS7{F0: (v62 + v62)}
                                                                    _ = v62
                                                                    _ = v65
                                                                }
                                                                _ = v59
                                                            }
                                                            fmt.Println(v32, v5)
                                                        }
                                                        _ = v54
                                                    }
                                                    _ = fi53
                                                }
                                                _ = v52
                                            } else {
                                                v66 := GenS8{F0: 42, F1: 42}
                                                if (((^(v44)) == v44) || v32) {
                                                    v67 := GenS9{F0: (((-2.25 != (3.14159 * 0.1)) && (true || ((v32 && v32) && v32))) && (true || v32))}
                                                    if ((!(!v32)) && false) {
                                                        var v68 int8 = -1
                                                        v69 := GenS6{F0: 42, F1: 385, F2: int16(v68)}
                                                        switch v2 {
                                                        case 9223372036854775807:
                                                            switch ("a" + v30) {
                                                            case "a":
                                                                fmt.Println(v68, v2)
                                                            default:
                                                                for fi70 := 0; fi70 < 16; fi70++ {
                                                                    if true {
                                                                        var v71 int = 9223372036854775807
                                                                        switch fi70 {
                                                                        case -88:
                                                                            var v72 int64 = (v29 - (-(v29)))
                                                                            _ = v72
                                                                        case 153:
                                                                            v73 := GenS10{F0: (v68 * v68)}
                                                                            _ = v73
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        if ((18446744073709551615 != (v5 & (^(v5)))) || (244 == 32767)) {
                                                                            v6 = append(v6, v30)
                                                                            if (996 >= 4294967295) {
                                                                                v44 -= 152
                                                                                _ = len(v6)
                                                                                var v74 uint64 = 824
                                                                                _ = v74
                                                                            } else {
                                                                                v29 = (v29 ^ (v29 << (((v5 & 255)) & 63)))
                                                                                _ = len(v3)
                                                                                v75 := GenS11{F0: v29, F1: 1}
                                                                                _ = v75
                                                                            }
                                                                        } else {
                                                                            var v76 uint64 = 2
                                                                            if ((v29 & v29) != 1) {
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    switch (v30 + (v30 + "a")) {
                                                                                    case "hello":
                                                                                        switch ("foo bar" + v30) {
                                                                                        case "\n\t":
                                                                                            v5 = v5
                                                                                        case "a":
                                                                                            switch 127 {
                                                                                            case 127:
                                                                                                var v77 string = "a"
                                                                                                _ = v77
                                                                                            case -2147483648:
                                                                                                v76 += (v76 ^ v76)
                                                                                            case 32768:
                                                                                                v6 = append(v6, v30)
                                                                                            }
                                                                                        default:
                                                                                            _ = len(v1)
                                                                                        }
                                                                                    default:
                                                                                        v3 = append(v3, fi50)
                                                                                    }
                                                                                case "foo bar":
                                                                                    v78 := GenS5{F0: v68, F1: 289, F2: 2147483647}
                                                                                    _ = v78
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        v5 -= v5
                                                                                        v79 := GenS1{F0: v29, F1: int64(fi50), F2: int16(v76)}
                                                                                        _ = v79
                                                                                    } else {
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    if ((true && (v5 == v5)) || true) {
                                                                                        var v80 int64 = (v29 * v29)
                                                                                        switch (-(fi70)) {
                                                                                        case 9223372036854775807:
                                                                                            if (v32 && (!v32)) {
                                                                                                if (v76 > v76) {
                                                                                                    var v81 int8 = v68
                                                                                                    v68 = (v81 ^ -1)
                                                                                                    var v82 uint = (^(v5))
                                                                                                    _ = v81
                                                                                                    _ = v82
                                                                                                }
                                                                                                _ = len(v4)
                                                                                                var v83 float32 = 0.0
                                                                                                _ = v83
                                                                                            } else {
                                                                                                v30 = (("hello" + v30) + ("hello" + v30))
                                                                                                if (!v32) {
                                                                                                    var v84 string = v30
                                                                                                    _ = v84
                                                                                                }
                                                                                                v29 -= (v80 << (((v5 >> ((v5) & 63))) & 63))
                                                                                            }
                                                                                        default:
                                                                                            var v85 uint32 = 734
                                                                                            var v86 float32 = (0.0 - 0.0)
                                                                                            v87 := GenS12{F0: 2147483648}
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        }
                                                                                        _ = v80
                                                                                    }
                                                                                }
                                                                                if v32 {
                                                                                    v68 = (v68 * v68)
                                                                                }
                                                                                v88 := GenS13{F0: (^(v76)), F1: 963}
                                                                                _ = v88
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (-(fi50))
                                                                    }
                                                                    if (!v32) {
                                                                        v89 := GenS10{F0: v68}
                                                                        var v90 int16 = 32767
                                                                        _ = v89
                                                                        _ = v90
                                                                    }
                                                                    _ = fi70
                                                                }
                                                            }
                                                        }
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = v30
                                                    }
                                                    _ = v67
                                                }
                                                _ = v66
                                            }
                                            _ = fi50
                                        }
                                    } else {
                                        for fi91 := 0; fi91 < 16; fi91++ {
                                            switch genFn0() {
                                            case 128:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            case -42:
                                                var v92 int = 4294967295
                                                _ = v92
                                            case 4294967295:
                                                for fi93 := 0; fi93 < 20; fi93++ {
                                                    var v94 float64 = -2.25
                                                    _ = fi93
                                                    _ = v94
                                                }
                                            }
                                            var v95 uint32 = 4294967295
                                            _ = fi91
                                            _ = v95
                                        }
                                    }
                                    fmt.Println(v29, v30)
                                }
                            } else {
                                if false {
                                    var v96 string = v30
                                    _ = v96
                                }
                                var v97 int64 = v29
                                _ = v97
                            }
                            _ = v44
                        } else {
                            if (v30 <= (v30 + (v30 + "a"))) {
                                v29 = 9223372036854775807
                                var v98 int16 = 1
                                var v99 float64 = -2.25
                                _ = v98
                                _ = v99
                            } else {
                                v1[2147483647] = v30
                                switch "hello" {
                                case "hello":
                                    for fi100 := 0; fi100 < 8; fi100++ {
                                        v32 = v32
                                        if v32 {
                                            for fi101 := 0; fi101 < 14; fi101++ {
                                                switch "a" {
                                                case "hello":
                                                    v102 := GenS14{F0: 505}
                                                    _ = v102
                                                case "\n\t":
                                                    v2 = fi100
                                                default:
                                                    if v32 {
                                                        if (0 > 255) {
                                                            switch v2 {
                                                            case -9223372036854775808:
                                                                v2 += fi100
                                                            default:
                                                                if (0 >= 558) {
                                                                    var v103 uint16 = 0
                                                                    if v32 {
                                                                        fmt.Println(v5)
                                                                        fmt.Println(v32, v30, v5)
                                                                        if (235 == 203) {
                                                                            var v104 int8 = 1
                                                                            switch (^(fi100)) {
                                                                            case -836:
                                                                                if (!false) {
                                                                                    var v105 int8 = (v104 ^ -4)
                                                                                    var v106 uint32 = 653
                                                                                    v2 = 1
                                                                                    _ = v105
                                                                                    _ = v106
                                                                                } else {
                                                                                    if (v5 < v5) {
                                                                                        switch (fi101 & fi100) {
                                                                                        case -577:
                                                                                            var v107 int8 = (v104 / ((v104) | 1))
                                                                                            _ = v107
                                                                                        case 2147483647:
                                                                                            delete(v1, 4294967295)
                                                                                            v29 = v29
                                                                                            var v108 uint8 = 42
                                                                                            _ = v108
                                                                                        case 42:
                                                                                            v2 -= (^(fi101))
                                                                                        }
                                                                                        var v109 bool = (((-2.25 + 0.1) + (((0.0 / 1e10) * -2.25) * 1.5)) >= -2.25)
                                                                                        _ = v109
                                                                                    } else {
                                                                                        v110 := GenS9{F0: (true && (1 >= -1))}
                                                                                        _ = v110
                                                                                    }
                                                                                    var v111 string = ("\n\t" + v30)
                                                                                    v32 = false
                                                                                    _ = v111
                                                                                }
                                                                            }
                                                                            _ = v104
                                                                        } else {
                                                                            if false {
                                                                                if (!false) {
                                                                                    fmt.Println(fi101, fi101, fi100)
                                                                                    fmt.Println(fi101, v30, fi100)
                                                                                    v103 += (^(v103))
                                                                                } else {
                                                                                    var v112 uint = 4294967295
                                                                                    switch (v30 + (v30 + ("a" + "hello"))) {
                                                                                    case "hello":
                                                                                        v29 = v29
                                                                                    case "\n\t":
                                                                                        var v113 uint16 = 736
                                                                                        _ = v113
                                                                                    case "a":
                                                                                        v1[65535] = ((v30 + (v30 + (v30 + v30))) + v30)
                                                                                    default:
                                                                                        var v114 int32 = -2147483648
                                                                                        _ = v114
                                                                                    }
                                                                                    _ = v112
                                                                                }
                                                                            } else {
                                                                                v3 = append(v3, v2)
                                                                                var v115 float32 = 0.0
                                                                                var v116 int = (v2 / (((fi101 / ((genFn0()) | 1))) | 1))
                                                                                var v117 uint16 = (v103 >> ((v5) & 63))
                                                                                var v118 bool = v32
                                                                                _ = v115
                                                                                _ = v116
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                if (391 == 247) {
                                                                                    v119 := GenS15{F0: v103, F1: -99}
                                                                                    if (false || v32) {
                                                                                        v120 := GenS15{F0: v103, F1: 1}
                                                                                        v121 := GenS16{F0: (v29 % ((v29) | 1))}
                                                                                        var v122 uint16 = 42
                                                                                        v122 += v122
                                                                                        switch 0 {
                                                                                        case 42:
                                                                                            if (((0.0 - 1.5) + -2.25) > float32(-2.25)) {
                                                                                                switch (v2 >> ((1) & 63)) {
                                                                                                case 1:
                                                                                                    if ((v30 < "foo bar") && v32) {
                                                                                                        if (float32((((1e10 / 1.0) + float64(3.14159)) / float64((0.1 + ((0.1 + 1.5) / 0.1))))) <= float32(float64((-2.25 / 1.0)))) {
                                                                                                            var v123 bool = v32
                                                                                                            fmt.Println(fi100)
                                                                                                            if (4294967295 >= 2) {
                                                                                                                v124 := GenS7{F0: ("hello" + "")}
                                                                                                                _ = v124
                                                                                                            }
                                                                                                            _ = v123
                                                                                                        } else {
                                                                                                            switch ("\n\t" + "hello") {
                                                                                                            case "a":
                                                                                                                if (v32 || false) {
                                                                                                                    v32 = v32
                                                                                                                    v125 := GenS15{F0: (v122 * (v103 * v103)), F1: 35}
                                                                                                                    switch "foo bar" {
                                                                                                                    case "":
                                                                                                                        v126 := GenS17{F0: v5, F1: 42}
                                                                                                                        _ = v126
                                                                                                                    case "hello":
                                                                                                                        switch fi100 {
                                                                                                                        case 42:
                                                                                                                            v127 := GenS15{F0: 582, F1: 15}
                                                                                                                            switch ((v30 + "\n\t") + ((("foo bar" + v30) + v30) + (v30 + ((v30 + ("a" + "\n\t")) + v30)))) {
                                                                                                                            case "":
                                                                                                                                var v128 float64 = 1e10
                                                                                                                                _ = v128
                                                                                                                            }
                                                                                                                            _ = v127
                                                                                                                        case 127:
                                                                                                                            v2 = (v2 >> ((v5) & 63))
                                                                                                                        default:
                                                                                                                            delete(v4, ("hello" + v30))
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v125
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v129 := v4[v30]
                                                                                                                if false {
                                                                                                                    var v130 int16 = 0
                                                                                                                    _ = v130
                                                                                                                }
                                                                                                                _ = v129
                                                                                                            case "foo bar":
                                                                                                                delete(v4, "hello")
                                                                                                            }
                                                                                                            switch (fi100 << ((0) & 63)) {
                                                                                                            case 1:
                                                                                                                v30 = ("foo bar" + ("hello" + v30))
                                                                                                            default:
                                                                                                                fmt.Println(fi100, v30, v103)
                                                                                                            }
                                                                                                            v131 := GenS9{F0: false}
                                                                                                            _ = v131
                                                                                                        }
                                                                                                    }
                                                                                                case 128:
                                                                                                    var v132 int = fi100
                                                                                                    _ = v132
                                                                                                default:
                                                                                                    var v133 int16 = 774
                                                                                                    _ = v133
                                                                                                }
                                                                                                fmt.Println(v29, v30)
                                                                                                var v134 uint32 = 75
                                                                                                _ = v134
                                                                                            } else {
                                                                                                var v135 uint64 = 199
                                                                                                delete(v4, v30)
                                                                                                v2 -= -9223372036854775808
                                                                                                _ = v135
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            if (1 < 1) {
                                                                                                v32 = true
                                                                                                var v136 string = v30
                                                                                                _ = v136
                                                                                            }
                                                                                        default:
                                                                                            switch genFn0() {
                                                                                            case 302:
                                                                                                v137 := GenS9{F0: true}
                                                                                                _ = v137
                                                                                            }
                                                                                        }
                                                                                        v138 := GenS18{F0: (true && (-179 != 42)), F1: v103, F2: 9223372036854775807}
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                        _ = v122
                                                                                        _ = v138
                                                                                    }
                                                                                    v30 = v30
                                                                                    _ = v119
                                                                                }
                                                                            case "a":
                                                                                v5 += (v5 * v5)
                                                                            default:
                                                                                var v139 float32 = float32(1e10)
                                                                                if (42 > 162) {
                                                                                    fmt.Println(v30, v2)
                                                                                }
                                                                                _ = v139
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v103
                                                                }
                                                            }
                                                            v32 = true
                                                            var v140 uint16 = 255
                                                            switch ("" + v30) {
                                                            case "hello":
                                                                switch v30 {
                                                                case "hello":
                                                                    v141 := v1[909]
                                                                    _ = v141
                                                                case "":
                                                                    if false {
                                                                        v30 = v30
                                                                        var v142 uint8 = 229
                                                                        var v143 bool = v32
                                                                        _ = v142
                                                                        _ = v143
                                                                    } else {
                                                                        v4[(("" + v30) + v30)] = 2147483648
                                                                        var v144 uint8 = 255
                                                                        _ = v144
                                                                    }
                                                                default:
                                                                    delete(v4, (v30 + ((((v30 + v30) + ((v30 + (((v30 + v30) + "hello") + (v30 + (((v30 + "") + v30) + "a")))) + "\n\t")) + v30) + v30)))
                                                                    var v145 uint8 = 42
                                                                    if len(v6) > 0 {
                                                                        v6[0] = ""
                                                                    }
                                                                    _ = v145
                                                                }
                                                            }
                                                            var v146 string = v30
                                                            _ = v140
                                                            _ = v146
                                                        } else {
                                                            var v147 int64 = v29
                                                            _ = v147
                                                        }
                                                        switch (^(fi101)) {
                                                        case -2147483648:
                                                            if v32 {
                                                                if (!(!(!((^(v5)) <= (v5 >> ((2) & 63)))))) {
                                                                    switch v30 {
                                                                    case "hello":
                                                                        v29 += (v29 % ((v29) | 1))
                                                                    default:
                                                                        v30 = ("a" + v30)
                                                                    }
                                                                } else {
                                                                    var v148 uint16 = 32767
                                                                    _ = v148
                                                                }
                                                            } else {
                                                                var v149 float64 = (1.5 / 1e10)
                                                                if v32 {
                                                                    switch (v30 + v30) {
                                                                    case "a":
                                                                        fmt.Println(v149)
                                                                    case "\n\t":
                                                                        v29 = (v29 / (((v29 << ((0) & 63))) | 1))
                                                                    case "hello":
                                                                        v4[(("hello" + "hello") + v30)] = 1
                                                                    }
                                                                    v150 := GenS11{F0: (v29 % ((v29) | 1)), F1: 4294967295}
                                                                    _ = v150
                                                                }
                                                                _ = v149
                                                            }
                                                            fmt.Println(v30, v29, fi100)
                                                            v32 = v32
                                                            v6 = append(v6, "hello")
                                                            var v151 uint32 = 880
                                                            _ = v151
                                                        case -928:
                                                            switch -502 {
                                                            case 0:
                                                                v152 := GenS6{F0: 145, F1: -2147483648, F2: 42}
                                                                _ = v152
                                                            case 32767:
                                                                _ = len(v6)
                                                            }
                                                        default:
                                                            v153 := GenS5{F0: -128, F1: (fi100 - -2147483648), F2: uint32(v5)}
                                                            _ = v153
                                                        }
                                                    }
                                                }
                                                switch 1 {
                                                case -2147483648:
                                                    var v154 bool = v32
                                                    _ = v154
                                                case 127:
                                                    v155 := GenS5{F0: 53, F1: (^(v2)), F2: 2147483647}
                                                    _ = v155
                                                default:
                                                    v156 := GenS19{F0: 255, F1: false, F2: (v29 | -990)}
                                                    _ = v156
                                                }
                                                var v157 uint8 = 2
                                                _ = fi101
                                                _ = v157
                                            }
                                            delete(v1, 65535)
                                        } else {
                                            fmt.Println(v30)
                                            v158 := GenS16{F0: (^(v29))}
                                            var v159 string = v30
                                            _ = v158
                                            _ = v159
                                        }
                                        fmt.Println(v5, v2)
                                        _ = fi100
                                    }
                                case "\n\t":
                                    for fi160 := 0; fi160 < 4; fi160++ {
                                        fmt.Println(v2, v5, v30)
                                        _ = fi160
                                    }
                                }
                            }
                            if (41 <= 128) {
                                var v161 bool = false
                                if true {
                                    v162 := GenS12{F0: 0}
                                    _ = v162
                                } else {
                                    for fi163 := 0; fi163 < 18; fi163++ {
                                        if true {
                                            for fi164 := 0; fi164 < 1; fi164++ {
                                                switch (v30 + v30) {
                                                case "hello":
                                                    var v165 uint64 = 4294967295
                                                    _ = v165
                                                case "\n\t":
                                                    if v161 {
                                                        var v166 uint = (v5 % (((v5 * 964)) | 1))
                                                        var v167 int64 = v29
                                                        var v168 uint8 = 127
                                                        var v169 float64 = (1e10 / 1e10)
                                                        fmt.Println(v29, v30)
                                                        _ = v166
                                                        _ = v167
                                                        _ = v168
                                                        _ = v169
                                                    }
                                                case "":
                                                    var v170 uint = (v5 + 611)
                                                    _ = v170
                                                default:
                                                    v32 = (251 > 65535)
                                                }
                                                _ = fi164
                                            }
                                            var v171 bool = ((v5 ^ v5) >= 555)
                                            var v172 uint64 = 2
                                            _ = v171
                                            _ = v172
                                        } else {
                                            switch "foo bar" {
                                            case "foo bar":
                                                if len(v3) > 0 {
                                                    v3[0] = -1
                                                }
                                            case "\n\t":
                                                var v173 float64 = float64((-2.25 + -2.25))
                                                _ = v173
                                            default:
                                                if true {
                                                    v174 := GenS8{F0: 42, F1: 255}
                                                    v30 = "\n\t"
                                                    if v32 {
                                                        fmt.Println(v30, v5)
                                                        var v175 uint64 = 2
                                                        v2 = fi163
                                                        _ = v175
                                                    }
                                                    _ = v174
                                                } else {
                                                    v176 := GenS15{F0: 611, F1: 42}
                                                    _ = v176
                                                }
                                            }
                                            var v177 uint = (v5 % (((^(v5))) | 1))
                                            for fi178 := 0; fi178 < 4; fi178++ {
                                                var v179 int64 = v29
                                                v180 := GenS3{F0: 807, F1: 783}
                                                _ = fi178
                                                _ = v179
                                                _ = v180
                                            }
                                            _ = v177
                                        }
                                        _ = fi163
                                    }
                                }
                                fmt.Println(v161)
                                _ = v161
                            } else {
                                fmt.Println(v2, v5, v30)
                                fmt.Println(v5, v29, v5)
                            }
                        }
                    }
                    var v181 int32 = 0
                    var v182 float32 = 0.1
                    _ = v181
                    _ = v182
                }
            } else {
                switch v30 {
                case "":
                    v32 = (-77 > -128)
                case "\n\t":
                    switch "foo bar" {
                    case "foo bar":
                        var v183 int8 = 1
                        _ = v183
                    }
                }
                var v184 bool = true
                delete(v1, 4294967295)
                _ = v184
            }
        }
    }
    switch 406 {
    case 722:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println(v30, v2, v32)
    var v185 int32 = 42
    if (42 != 255) {
        var v186 int32 = 1
        _ = v186
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v185:", v185)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v185
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
