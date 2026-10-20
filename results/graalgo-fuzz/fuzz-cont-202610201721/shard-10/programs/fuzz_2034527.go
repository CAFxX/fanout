package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8; F1 int16; F2 uint32 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 uint8; F1 bool; F2 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 int; F2 string }
type GenS10 struct { F0 bool; F1 uint; F2 uint }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32; F1 uint8 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 bool; F1 uint8; F2 uint }
type GenS17 struct { F0 string; F1 uint16 }
type GenS18 struct { F0 uint64; F1 int64 }
type GenS19 struct { F0 int8; F1 uint }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 bool }

func main() {
    var v1 int16 = 32767
    var v2 uint = 1
    v3 := []int{20, 9223372036854775807, int(v2)}
    v4 := []uint32{33, 2, 648, 4294967295}
    var v5 string = ("\n\t" + (("foo bar" + "\n\t") + "hello"))
    var v6 uint8 = 42
    fmt.Println(v5, v2)
    var v7 uint8 = (v6 >> ((v2) & 63))
    var v8 int = int(v1)
    switch "foo bar" {
    case "a":
        v1 = v1
    case "":
        v9 := GenS1{F0: (v7 <= v6)}
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 13; fi10++ {
            v1 = -1
            fmt.Println(v8, v2, v1)
            if false {
                v4 = append(v4, 2)
            }
            _ = fi10
        }
    }
    var v11 int64 = 0
    for fi12 := 0; fi12 < 2; fi12++ {
        switch v5 {
        case "foo bar":
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(fi13, fi12, fi13)
                switch -604 {
                case -9223372036854775808:
                    if (((-2.25 / ((((float64((1.5 / (3.14159 + 0.0))) + 3.14159) - 3.14159) / 1e10) + 1e10)) * 3.14159) <= 0.0) {
                        var v14 uint16 = 0
                        v15 := GenS1{F0: false}
                        _ = v14
                        _ = v15
                    } else {
                        if (false || (!(int64(v8) == (v11 << (((v2 * v2)) & 63))))) {
                            v8 += v8
                        } else {
                            fmt.Println(v2)
                            fmt.Println(v2, v6, v2)
                        }
                        var v16 float32 = (0.1 / (-2.25 * -2.25))
                        v3 = append(v3, (fi12 - fi12))
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                v17 := GenS2{F0: v2, F1: 42, F2: 486}
                                var v18 int64 = 1
                                _ = len(v3)
                                _ = v17
                                _ = v18
                            }
                        } else {
                            v19 := GenS2{F0: (v2 << (((v2 & (v2 * v2))) & 63)), F1: -930, F2: 390}
                            _ = v19
                        }
                        _ = v16
                    }
                case 32768:
                    v20 := GenS1{F0: true}
                    _ = v20
                case 42:
                    v4 = append(v4, 42)
                }
                switch "" {
                case "\n\t":
                    v6 -= 67
                case "hello":
                    if true {
                        v21 := GenS1{F0: (65535 >= 619)}
                        v3 = append(v3, v8)
                        _ = v21
                    } else {
                        v11 -= (v11 | (^(v11)))
                        v22 := GenS1{F0: (3.14159 > float64(((1.5 + -2.25) * 3.14159)))}
                        _ = v22
                    }
                case "":
                    v4 = append(v4, 804)
                default:
                    v7 -= 255
                }
                _ = fi13
            }
        case "hello":
            v4 = append(v4, 2147483647)
        case "":
            if (true && (v6 > 128)) {
                v23 := GenS1{F0: false}
                var v24 bool = true
                _ = v23
                _ = v24
            } else {
                if (true || (!true)) {
                    fmt.Println(v7, v5, v6)
                }
                var v25 uint = v2
                fmt.Println(v1, v25, fi12)
                _ = v25
            }
        default:
            v4 = append(v4, 65535)
        }
        switch v8 {
        case -935:
            v7 -= v7
        default:
            if (((v2 % ((163) | 1)) > v2) && false) {
                v6 = (v6 % ((128) | 1))
                v26 := GenS3{F0: -78, F1: (^(v1)), F2: 405}
                v8 -= v8
                _ = v26
            } else {
                if true {
                    v6 -= v7
                    var v27 float32 = -2.25
                    v28 := GenS2{F0: (v2 * v2), F1: -2147483648, F2: 2}
                    _ = v27
                    _ = v28
                }
            }
        }
        var v29 int64 = 964
        _ = fi12
        _ = v29
    }
    var v30 string = "a"
    for fi31 := 0; fi31 < 9; fi31++ {
        v6 = 28
        v32 := GenS2{F0: v2, F1: -1, F2: 431}
        v33 := GenS3{F0: 45, F1: v1, F2: 2147483648}
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        v5 = "a"
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 float64 = float64(-2.25)
            _ = fi35
            _ = v36
        }
        var v37 int8 = 82
        _ = fi34
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 string = v5
        if (201 > 18446744073709551615) {
            for fi40 := 0; fi40 < 1; fi40++ {
                v41 := GenS4{F0: "hello", F1: v2}
                if (v11 != (v11 | (v11 ^ (v11 & (-(v11)))))) {
                    if (-2.25 >= float64(((3.14159 + (0.1 + (((0.0 + -2.25) + (0.1 + -2.25)) + 0.1))) + (1e10 - float32(float64(((1.5 + 0.1) * 1e10))))))) {
                        if true {
                            var v42 uint32 = 42
                            v43 := GenS2{F0: (v2 >> (((v2 ^ 182)) & 63)), F1: 2147483647, F2: v42}
                            _ = v42
                            _ = v43
                        }
                        v4 = append(v4, 974)
                        if (42 <= -128) {
                            if false {
                                switch (^(fi38)) {
                                case 32767:
                                    if false {
                                        if true {
                                            if len(v3) > 0 {
                                                v3[0] = (fi38 << (((v2 >> ((v2) & 63))) & 63))
                                            }
                                        }
                                        v30 = v30
                                        v39 = v5
                                        v2 = (v2 & (^(v2)))
                                        switch (fi38 | 32768) {
                                        case -9223372036854775808:
                                            v6 = (v6 % (((v7 + (v7 - v6))) | 1))
                                        default:
                                            v44 := GenS4{F0: ((v39 + ("foo bar" + v39)) + (v39 + "\n\t")), F1: 657}
                                            _ = v44
                                        }
                                    } else {
                                        var v45 uint64 = 884
                                        _ = len(v4)
                                        var v46 uint8 = v6
                                        switch v8 {
                                        case 32767:
                                            v1 = int16(fi40)
                                        case 32768:
                                            v1 += (v1 * (v1 / (((v1 - (v1 << ((2) & 63)))) | 1)))
                                            var v47 bool = false
                                            _ = v47
                                        case 128:
                                            switch v39 {
                                            case "":
                                                v4 = append(v4, 0)
                                            case "a":
                                                fmt.Println(v5, v7)
                                            case "hello":
                                                var v48 uint = (v2 >> ((18446744073709551615) & 63))
                                                if true {
                                                    var v49 int16 = int16(v11)
                                                    fmt.Println(fi40, v39, fi40)
                                                    _ = v49
                                                } else {
                                                    v45 += (v45 >> ((393) & 63))
                                                    fmt.Println(v5, fi40, v46)
                                                    if false {
                                                        fmt.Println(v30, v11)
                                                        switch (((("\n\t" + v30) + ("hello" + v30)) + "") + ("foo bar" + (((v39 + "") + ("\n\t" + v39)) + v5))) {
                                                        case "\n\t":
                                                            var v50 int16 = (v1 & v1)
                                                            _ = v50
                                                        default:
                                                            switch v5 {
                                                            case "foo bar":
                                                                if true {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = uint32(v48)
                                                                    }
                                                                    if false {
                                                                        var v51 int16 = -522
                                                                        _ = v51
                                                                    }
                                                                } else {
                                                                    v4 = append(v4, 2147483648)
                                                                    if false {
                                                                        v8 = 217
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 874
                                                                        }
                                                                    }
                                                                    v52 := GenS5{F0: v48}
                                                                    _ = v52
                                                                }
                                                            case "":
                                                                v2 -= (^(v2))
                                                            case "a":
                                                                if ((fi40 | fi38) == (fi38 * (fi40 / (((fi40 << (((v2 % (((v48 << ((v2) & 63))) | 1))) & 63))) | 1)))) {
                                                                    var v53 float64 = 0.1
                                                                    fmt.Println(v7)
                                                                    v54 := GenS6{F0: 42, F1: v45}
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    v30 = v30
                                                                }
                                                            }
                                                        }
                                                        var v55 int64 = -9223372036854775808
                                                        v56 := GenS7{F0: v7, F1: false, F2: 0}
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        v7 = v46
                                                        v57 := GenS8{F0: v30}
                                                        v1 = 42
                                                        _ = v57
                                                    }
                                                    var v58 uint8 = v46
                                                    _ = v58
                                                }
                                                fmt.Println(v8, v46)
                                                _ = v48
                                            }
                                            switch fi38 {
                                            case 128:
                                                switch 737 {
                                                case 127:
                                                    if false {
                                                        if ((0.1 * 3.14159) < (3.14159 + (0.0 / 0.1))) {
                                                            if true {
                                                                if false {
                                                                    var v59 uint = v2
                                                                    if (true || (2147483648 > 42)) {
                                                                        var v60 uint32 = 146
                                                                        var v61 int32 = 42
                                                                        v11 += v11
                                                                        var v62 bool = (!(true && false))
                                                                        fmt.Println(v7, v46)
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v62
                                                                    } else {
                                                                        v7 += uint8(v45)
                                                                        if (728 > 65535) {
                                                                            v4 = append(v4, 2147483647)
                                                                        }
                                                                        var v63 uint32 = 1
                                                                        _ = v63
                                                                    }
                                                                    switch v5 {
                                                                    case "hello":
                                                                        fmt.Println(fi40, v2, v1)
                                                                        if true {
                                                                            var v64 int16 = v1
                                                                            fmt.Println(v2)
                                                                            v59 -= 1
                                                                            _ = v64
                                                                        }
                                                                        if true {
                                                                            if ((!(true || ((^(v59)) >= v2))) && false) {
                                                                                v65 := GenS9{F0: "", F1: (^(fi40)), F2: v39}
                                                                                v7 = v7
                                                                                _ = v65
                                                                            } else {
                                                                                fmt.Println(v11, v1, v46)
                                                                            }
                                                                            switch int(v45) {
                                                                            case -759:
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v66 uint8 = 1
                                                                                fmt.Println(v6, v1)
                                                                                if (13 != -128) {
                                                                                    var v67 uint32 = 56
                                                                                    _ = v67
                                                                                } else {
                                                                                    switch ("hello" + "a") {
                                                                                    case "hello":
                                                                                        if false {
                                                                                            v5 = "hello"
                                                                                        } else {
                                                                                            v59 = (v59 | v59)
                                                                                            var v68 int = fi40
                                                                                            _ = v68
                                                                                        }
                                                                                    case "":
                                                                                        switch v8 {
                                                                                        case 0:
                                                                                            v69 := GenS10{F0: true, F1: v2, F2: (v59 % ((136) | 1))}
                                                                                            _ = v69
                                                                                        case 4294967295:
                                                                                            var v70 uint = (v2 ^ v2)
                                                                                            _ = v70
                                                                                        case 239:
                                                                                            var v71 uint = (v2 / ((v59) | 1))
                                                                                            _ = v71
                                                                                        }
                                                                                    default:
                                                                                        var v72 int16 = v1
                                                                                        _ = v72
                                                                                    }
                                                                                    fmt.Println(fi38, v7, v6)
                                                                                    v73 := GenS11{F0: fi38, F1: (v8 + fi40)}
                                                                                    _ = v73
                                                                                }
                                                                                _ = v66
                                                                            case 32767:
                                                                                var v74 uint64 = (v45 | v45)
                                                                                fmt.Println(v59, fi40, v45)
                                                                                v3 = append(v3, fi40)
                                                                                _ = v74
                                                                            case -2147483648:
                                                                                var v75 uint64 = (v45 % (((v45 / ((1) | 1))) | 1))
                                                                                _ = v75
                                                                            }
                                                                            var v76 uint16 = 255
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    _ = v59
                                                                }
                                                                var v77 string = v30
                                                                if false {
                                                                    var v78 uint8 = v46
                                                                    _ = v78
                                                                } else {
                                                                    if true {
                                                                        var v79 string = ((v5 + "\n\t") + "hello")
                                                                        _ = v79
                                                                    } else {
                                                                        if true {
                                                                            v80 := GenS5{F0: v2}
                                                                            fmt.Println(v7)
                                                                            v7 = v7
                                                                            _ = v80
                                                                        } else {
                                                                            var v81 uint32 = 540
                                                                            v82 := GenS8{F0: (((v77 + "\n\t") + v30) + v30)}
                                                                            v8 += v8
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        var v83 float64 = 0.1
                                                                        var v84 string = v77
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                }
                                                                _ = v77
                                                            } else {
                                                                v1 = -32768
                                                                v39 = (((("\n\t" + v5) + v30) + ("" + v39)) + "")
                                                            }
                                                            if (!true) {
                                                                if true {
                                                                    fmt.Println(v46, v30, fi40)
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (fi40 << (((v2 + v2)) & 63))
                                                                    }
                                                                }
                                                                switch ((v39 + v39) + v39) {
                                                                case "":
                                                                    var v85 int = fi40
                                                                    _ = v85
                                                                case "hello":
                                                                    var v86 bool = false
                                                                    _ = v86
                                                                default:
                                                                    v46 += 42
                                                                }
                                                                switch -472 {
                                                                case 32767:
                                                                    if false {
                                                                        if false {
                                                                            var v87 string = v30
                                                                            _ = v87
                                                                        } else {
                                                                            v3 = append(v3, (^(v8)))
                                                                            var v88 string = v30
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                case -255:
                                                                    v45 += (v45 + v45)
                                                                case -692:
                                                                    if false {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v45 += v45
                                                                        var v89 bool = true
                                                                        _ = v89
                                                                    } else {
                                                                        if (853 < 375) {
                                                                            if false {
                                                                                fmt.Println(v2, v1, v2)
                                                                            } else {
                                                                                var v90 uint32 = 0
                                                                                _ = v90
                                                                            }
                                                                            if (v39 != "\n\t") {
                                                                                v3 = append(v3, fi40)
                                                                            }
                                                                            v91 := GenS3{F0: 42, F1: 1000, F2: 0}
                                                                            _ = v91
                                                                        }
                                                                        switch (fi40 | fi38) {
                                                                        case 4294967295:
                                                                            v6 -= (^(v46))
                                                                        default:
                                                                            var v92 uint8 = v7
                                                                            v6 -= v7
                                                                            _ = v92
                                                                        }
                                                                        fmt.Println(v1)
                                                                    }
                                                                }
                                                                v39 = ("hello" + v30)
                                                                var v93 int8 = 28
                                                                switch 58 {
                                                                case 666:
                                                                    _ = len(v3)
                                                                case -583:
                                                                    v45 = (v45 - v45)
                                                                case 128:
                                                                    _ = len(v3)
                                                                }
                                                                _ = v93
                                                            }
                                                            var v94 uint16 = 32767
                                                            _ = v94
                                                        }
                                                        v95 := GenS12{F0: 1}
                                                        v8 += fi40
                                                        switch ((v39 + v39) + ("a" + "hello")) {
                                                        case "":
                                                            switch fi38 {
                                                            case 32767:
                                                                v3 = append(v3, 32767)
                                                            default:
                                                                v96 := GenS7{F0: (v6 >> ((65535) & 63)), F1: false, F2: -32768}
                                                                _ = v96
                                                            }
                                                            fmt.Println(v5, v5)
                                                        default:
                                                            v97 := GenS2{F0: (v2 >> (((v2 - (v2 * v2))) & 63)), F1: 732, F2: 943}
                                                            _ = v97
                                                        }
                                                        var v98 int = fi38
                                                        _ = v95
                                                        _ = v98
                                                    } else {
                                                        var v99 int = fi38
                                                        v99 = 127
                                                        _ = v99
                                                    }
                                                case 140:
                                                    var v100 int8 = -17
                                                    _ = v100
                                                default:
                                                    switch (fi38 + v8) {
                                                    case 4294967295:
                                                        if false {
                                                            if true {
                                                                v6 -= v6
                                                                fmt.Println(v11, v45)
                                                                v101 := GenS2{F0: (v2 - v2), F1: 1, F2: 856}
                                                                _ = v101
                                                            }
                                                            switch fi40 {
                                                            case 912:
                                                                if ((true || false) || false) {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v102 int8 = -115
                                                                    v30 = "\n\t"
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 0
                                                                    }
                                                                    var v103 uint32 = 237
                                                                    var v104 uint8 = (v6 % (((v46 | (v6 & 128))) | 1))
                                                                    _ = v102
                                                                    _ = v103
                                                                    _ = v104
                                                                } else {
                                                                    switch "" {
                                                                    case "hello":
                                                                        if (!false) {
                                                                            fmt.Println(v1, fi38, fi40)
                                                                            switch v39 {
                                                                            case "hello":
                                                                                var v105 bool = (v5 >= ("hello" + "a"))
                                                                                _ = v105
                                                                            case "foo bar":
                                                                                var v106 string = v5
                                                                                _ = v106
                                                                            case "\n\t":
                                                                                v1 -= v1
                                                                            default:
                                                                                var v107 bool = (-1 > 351)
                                                                                v108 := GenS6{F0: -1, F1: 227}
                                                                                v8 = (-(fi40))
                                                                                _ = v107
                                                                                _ = v108
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v109 uint32 = 2147483648
                                                                        _ = v109
                                                                    }
                                                                }
                                                            case 42:
                                                                if true {
                                                                    if (-1 > 86) {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 4294967295
                                                                        }
                                                                        if true {
                                                                            v1 -= (v1 & int16(v7))
                                                                            fmt.Println(v8, v11, v6)
                                                                        } else {
                                                                            fmt.Println(fi40, fi38)
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v6 = 127
                                                                        v8 += fi38
                                                                        v30 = "a"
                                                                    } else {
                                                                        switch (fi38 | (v8 & (fi38 << ((v2) & 63)))) {
                                                                        case -858:
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                v110 := GenS13{F0: 2147483647, F1: (^(v7))}
                                                                                _ = v110
                                                                            default:
                                                                                fmt.Println(v7, v6, v5)
                                                                            }
                                                                        case 683:
                                                                            var v111 uint64 = 42
                                                                            _ = v111
                                                                        case 32768:
                                                                            var v112 float64 = 1e10
                                                                            _ = v112
                                                                        default:
                                                                            var v113 string = (v30 + "foo bar")
                                                                            if ((false && false) || false) {
                                                                                var v114 uint64 = uint64(v6)
                                                                                if false {
                                                                                    v7 = (v46 + (v6 << ((18446744073709551615) & 63)))
                                                                                }
                                                                                switch fi38 {
                                                                                case -522:
                                                                                    switch v39 {
                                                                                    case "":
                                                                                        v115 := GenS14{F0: v1}
                                                                                        _ = v115
                                                                                    default:
                                                                                        var v116 int64 = (-(v11))
                                                                                        _ = v116
                                                                                    }
                                                                                }
                                                                                _ = v114
                                                                            } else {
                                                                                fmt.Println(v39)
                                                                                v46 = v6
                                                                            }
                                                                            fmt.Println(v11)
                                                                            _ = v113
                                                                        }
                                                                        if false {
                                                                            switch 42 {
                                                                            case 128:
                                                                                fmt.Println(v11, v45)
                                                                            case -925:
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    var v117 uint8 = (v46 / (((v6 & v7)) | 1))
                                                                                    _ = v117
                                                                                case "a":
                                                                                    if true {
                                                                                        v8 += 42
                                                                                        if (-2.25 == -2.25) {
                                                                                            switch v39 {
                                                                                            case "":
                                                                                                v7 = v7
                                                                                            case "foo bar":
                                                                                                v118 := GenS15{F0: 255}
                                                                                                v5 = v39
                                                                                                fmt.Println(fi40)
                                                                                                _ = v118
                                                                                            default:
                                                                                                if true {
                                                                                                    if (!(((1.5 == (3.14159 + -2.25)) || true) && false)) {
                                                                                                        v119 := GenS7{F0: (v46 % ((v7) | 1)), F1: false, F2: (v1 - -388)}
                                                                                                        _ = v119
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        switch (fi38 | (v8 % (((fi40 % (((fi38 << ((v2) & 63))) | 1))) | 1))) {
                                                                                        case -258:
                                                                                            v11 += 471
                                                                                        case 593:
                                                                                            fmt.Println(fi38)
                                                                                        default:
                                                                                            if false {
                                                                                                var v120 uint64 = 321
                                                                                                var v121 uint16 = 2
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                            } else {
                                                                                                v7 = 255
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case -818:
                                                                                v7 = v46
                                                                            default:
                                                                                if true {
                                                                                    v30 = v39
                                                                                    v122 := GenS14{F0: 42}
                                                                                    _ = v122
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v45 += 2
                                                                        }
                                                                        v123 := GenS5{F0: (v2 | v2)}
                                                                        _ = v123
                                                                    }
                                                                    _ = len(v4)
                                                                }
                                                            }
                                                        } else {
                                                            if true {
                                                                var v124 int64 = (-(v11))
                                                                switch ((v39 + "foo bar") + "hello") {
                                                                case "hello":
                                                                    if (v46 != 1) {
                                                                        var v125 uint8 = (v46 / ((2) | 1))
                                                                        _ = v125
                                                                    } else {
                                                                        v124 += 0
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v4 = append(v4, 645)
                                                                    }
                                                                case "":
                                                                    fmt.Println(v30, v39, v11)
                                                                default:
                                                                    var v126 int = 513
                                                                    _ = v126
                                                                }
                                                                _ = v124
                                                            }
                                                            if true {
                                                                var v127 uint64 = 42
                                                                v128 := GenS16{F0: true, F1: v46, F2: v2}
                                                                v1 = 32767
                                                                _ = v127
                                                                _ = v128
                                                            } else {
                                                                var v129 int16 = v1
                                                                _ = v129
                                                            }
                                                        }
                                                    case 11:
                                                        v45 += 911
                                                    default:
                                                        var v130 int64 = (v11 >> ((v2) & 63))
                                                        _ = v130
                                                    }
                                                }
                                            case -915:
                                                var v131 uint32 = 888
                                                _ = v131
                                            case 928:
                                                v39 = ("foo bar" + "\n\t")
                                            }
                                            var v132 uint32 = 2
                                            _ = v132
                                        default:
                                            switch fi40 {
                                            case 302:
                                                v45 -= (v45 + v45)
                                            default:
                                                v133 := GenS17{F0: (v5 + v30), F1: 260}
                                                _ = v133
                                            }
                                        }
                                        v6 -= (v46 * (^(v6)))
                                        _ = v45
                                        _ = v46
                                    }
                                case 32768:
                                    var v134 int = (v8 * (fi40 ^ (v8 >> ((v2) & 63))))
                                    _ = v134
                                case -1:
                                    v5 = v39
                                default:
                                    var v135 int32 = 1
                                    _ = v135
                                }
                                v30 = v5
                                v1 += (v1 | 0)
                            } else {
                                v30 = v5
                                var v136 uint32 = 1
                                var v137 uint = v2
                                _ = v136
                                _ = v137
                            }
                        } else {
                            v138 := GenS9{F0: "\n\t", F1: (^(v8)), F2: ("\n\t" + v39)}
                            var v139 string = (v5 + v5)
                            _ = v138
                            _ = v139
                        }
                        var v140 uint16 = 0
                        _ = v140
                    } else {
                        v39 = v30
                        if true {
                            v11 += (v11 * 1)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v2, fi40)
                            switch fi38 {
                            case 168:
                                fmt.Println(v5, fi38)
                            case 9223372036854775807:
                                if true {
                                    var v141 uint32 = 65535
                                    _ = v141
                                }
                            }
                            var v142 int8 = 0
                            if (true && false) {
                                if false {
                                    var v143 uint8 = v6
                                    _ = v143
                                } else {
                                    var v144 uint64 = 4294967295
                                    _ = v144
                                }
                                if len(v3) > 0 {
                                    v3[0] = (fi38 * 32767)
                                }
                            }
                            var v145 string = ("a" + ((v5 + "") + (v5 + "a")))
                            v146 := GenS18{F0: 0, F1: (v11 * (v11 | (v11 - int64(v2))))}
                            if (("\n\t" + v5) == (v5 + v5)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                _ = len(v3)
                                var v147 uint32 = 65535
                                _ = v147
                            }
                            _ = v142
                            _ = v145
                            _ = v146
                        } else {
                            fmt.Println(fi38, v30, v2)
                            var v148 int64 = v11
                            _ = v148
                        }
                    }
                }
                _ = fi40
                _ = v41
            }
            if false {
                for fi149 := 0; fi149 < 5; fi149++ {
                    var v150 int64 = -1
                    switch fi38 {
                    case 2147483647:
                        switch ((v5 + (("foo bar" + ("foo bar" + "a")) + v5)) + ("a" + v39)) {
                        case "":
                            v6 = (^(v7))
                        }
                    case 9223372036854775807:
                        var v151 uint64 = 4294967295
                        _ = v151
                    case 4294967295:
                        switch ((v39 + ("a" + ((v30 + "a") + v30))) + ("\n\t" + v39)) {
                        case "foo bar":
                            fmt.Println(v1, v1, v1)
                        case "a":
                            switch (-(v8)) {
                            case -9223372036854775808:
                                v152 := GenS19{F0: 33, F1: (v2 / (((v2 + (v2 - v2))) | 1))}
                                _ = v152
                            case 42:
                                var v153 float64 = float64(3.14159)
                                _ = v153
                            case -2:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                fmt.Println(v30, v7)
                            }
                        }
                    default:
                        v4 = append(v4, 861)
                    }
                    _ = fi149
                    _ = v150
                }
                var v154 int16 = (v1 & (v1 % ((v1) | 1)))
                if true {
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    fmt.Println(v154, v39, v1)
                    if (!((((((("" + (v39 + v39)) + ((v39 + "\n\t") + v5)) + "foo bar") + v5) != ((("foo bar" + v39) + "hello") + "hello")) && (true && false)) || (854 != 0))) {
                        for fi155 := 0; fi155 < 7; fi155++ {
                            var v156 int8 = 42
                            v30 = (v30 + v39)
                            if false {
                                switch (^(fi155)) {
                                case 2147483647:
                                    fmt.Println(v30, v8, fi155)
                                }
                                _ = len(v3)
                                switch (fi155 / ((-2147483648) | 1)) {
                                case 32768:
                                    v11 = (v11 | v11)
                                case 0:
                                    if len(v4) > 0 {
                                        v4[0] = 534
                                    }
                                }
                            }
                            fmt.Println(v156)
                            v157 := GenS17{F0: v39, F1: 1}
                            _ = fi155
                            _ = v156
                            _ = v157
                        }
                    }
                } else {
                    fmt.Println(v30)
                    v11 += 0
                }
                _ = v154
            } else {
                if len(v4) > 0 {
                    v4[0] = 425
                }
            }
            var v158 string = "a"
            _ = v158
        } else {
            var v159 int64 = 0
            for fi160 := 0; fi160 < 15; fi160++ {
                v159 += (v159 / ((v11) | 1))
                v2 += (v2 - 358)
                v161 := GenS7{F0: v7, F1: (1 <= 0), F2: v1}
                _ = fi160
                _ = v161
            }
            _ = v159
        }
        _ = fi38
        _ = v39
    }
    switch v30 {
    case "\n\t":
        v162 := GenS20{F0: true}
        _ = v162
    case "hello":
        v2 += v2
    case "":
        if false {
            fmt.Println(v11, v8)
            var v163 int32 = 0
            if true {
                v164 := GenS21{F0: ((!(true || (1 > 203))) || false)}
                _ = v164
            } else {
                for fi165 := 0; fi165 < 9; fi165++ {
                    v11 = (v11 | v11)
                    fmt.Println(v30, v163)
                    _ = fi165
                }
                for fi166 := 0; fi166 < 14; fi166++ {
                    v167 := GenS6{F0: v163, F1: 1}
                    var v168 uint = (v2 ^ 255)
                    _ = fi166
                    _ = v167
                    _ = v168
                }
            }
            _ = v163
        } else {
            var v169 uint64 = 42
            var v170 uint16 = uint16(v6)
            _ = len(v4)
            _ = v169
            _ = v170
        }
    default:
        _ = len(v3)
    }
    if false {
        for fi171 := 0; fi171 < 2; fi171++ {
            v3 = append(v3, fi171)
            _ = fi171
        }
        switch "\n\t" {
        case "a":
            v172 := GenS12{F0: 2}
            _ = v172
        case "hello":
            if false {
                for fi173 := 0; fi173 < 15; fi173++ {
                    v174 := GenS8{F0: v30}
                    _ = fi173
                    _ = v174
                }
            } else {
                var v175 bool = true
                _ = v175
            }
        case "foo bar":
            if false {
                for fi176 := 0; fi176 < 2; fi176++ {
                    switch ("foo bar" + v5) {
                    case "foo bar":
                        switch (fi176 ^ (v8 + fi176)) {
                        case -2147483648:
                            v7 = v6
                        default:
                            var v177 uint8 = (v6 >> (((v2 + v2)) & 63))
                            _ = v177
                        }
                    case "":
                        if false {
                            switch "hello" {
                            case "":
                                for fi178 := 0; fi178 < 12; fi178++ {
                                    v4 = append(v4, 0)
                                    _ = fi178
                                }
                            case "hello":
                                var v179 float32 = 0.1
                                _ = v179
                            default:
                                for fi180 := 0; fi180 < 19; fi180++ {
                                    var v181 uint64 = 82
                                    var v182 int64 = (^(v11))
                                    var v183 int = 0
                                    _ = fi180
                                    _ = v181
                                    _ = v182
                                    _ = v183
                                }
                            }
                        } else {
                            var v184 int64 = (v11 / ((v11) | 1))
                            _ = v184
                        }
                    default:
                        v5 = v30
                    }
                    v7 = (v6 + v7)
                    _ = fi176
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        default:
            v185 := GenS7{F0: 171, F1: ((v6 & 0) < (v6 << ((42) & 63))), F2: (v1 ^ (v1 * v1))}
            _ = v185
        }
    } else {
        var v186 string = "hello"
        _ = v186
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    switch v8 {
    case -9223372036854775808:
        v2 = (v2 / ((1) | 1))
    case 128:
        if false {
            for fi187 := 0; fi187 < 10; fi187++ {
                for fi188 := 0; fi188 < 1; fi188++ {
                    var v189 uint32 = uint32(v2)
                    if false {
                        fmt.Println(v30)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v190 int = v8
                    _ = fi188
                    _ = v189
                    _ = v190
                }
                if len(v3) > 0 {
                    v3[0] = -1
                }
                switch (v8 * 737) {
                case -648:
                    v6 += v7
                case 1:
                    for fi191 := 0; fi191 < 7; fi191++ {
                        var v192 bool = false
                        _ = fi191
                        _ = v192
                    }
                }
                _ = fi187
            }
            for fi193 := 0; fi193 < 4; fi193++ {
                var v194 uint64 = 18446744073709551615
                _ = len(v4)
                v194 = v194
                _ = fi193
                _ = v194
            }
            var v195 uint = 2
            _ = v195
        }
    case 4294967295:
        for fi196 := 0; fi196 < 20; fi196++ {
            if (100 > 0) {
                if true {
                    if (-340 < 42) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(fi196, v2, v2)
                        v3 = append(v3, v8)
                    } else {
                        var v197 uint32 = 860
                        switch (-(v8)) {
                        case 127:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 895:
                            var v198 int16 = (v1 - (v1 | v1))
                            _ = v198
                        default:
                            if (373 >= 1) {
                                v8 = -374
                                switch (v30 + "") {
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = v8
                                    }
                                case "hello":
                                    v199 := GenS16{F0: (547 > 4294967295), F1: v6, F2: v2}
                                    _ = v199
                                }
                                var v200 uint64 = 303
                                _ = v200
                            }
                        }
                        var v201 int64 = (v11 * v11)
                        _ = v197
                        _ = v201
                    }
                }
            } else {
                var v202 uint16 = 2
                _ = v202
            }
            _ = fi196
        }
    default:
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
