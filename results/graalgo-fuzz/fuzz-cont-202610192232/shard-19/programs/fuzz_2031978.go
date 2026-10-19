package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint32; F2 uint64 }
type GenS5 struct { F0 bool; F1 bool; F2 int64 }
type GenS6 struct { F0 uint16; F1 bool; F2 string }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 int64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int64; F1 uint8; F2 int64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 bool; F1 uint }
type GenS18 struct { F0 int; F1 int16; F2 int }
type GenS19 struct { F0 string; F1 uint }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int64; F1 int; F2 string }
type GenS22 struct { F0 int }
type GenS23 struct { F0 int16; F1 int64; F2 int16 }
type GenS24 struct { F0 int; F1 uint16 }
type GenS25 struct { F0 int8; F1 int16 }

func main() {
    v1 := []string{"\n\t", ("" + "\n\t"), "a"}
    var v2 uint16 = 697
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    v5 := []int{-988, 127, 32767}
    v6 := GenS1{F0: (v2 / ((255) | 1))}
    switch "hello" {
    case "foo bar":
        v4[0] = -259
    case "a":
        v7 := GenS1{F0: (v2 + 42)}
        _ = v7
    }
    v2 = (v2 - (v2 | (v2 * 2)))
    for fi8 := 0; fi8 < 18; fi8++ {
        switch "a" {
        case "a":
            v2 -= v2
        case "":
            fmt.Println(v2)
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi8
    }
    v9 := GenS2{F0: -1}
    if (2147483648 == 301) {
        var v10 int16 = 147
        var v11 int64 = -9223372036854775808
        _ = v10
        _ = v11
    }
    var v12 string = "hello"
    v2 = 2
    v12 = (("foo bar" + (v12 + ((v12 + v12) + v12))) + ((v12 + (((v12 + "hello") + "\n\t") + v12)) + v12))
    v13 := GenS2{F0: 0}
    var v14 int64 = 42
    for fi15 := 0; fi15 < 6; fi15++ {
        if (false || false) {
            fmt.Println(fi15)
        } else {
            switch v12 {
            case "foo bar":
                var v16 uint64 = 1
                _ = v16
            }
            switch fi15 {
            case 2147483647:
                v2 = (^(v2))
            case -926:
                for fi17 := 0; fi17 < 1; fi17++ {
                    v2 -= (v2 - (^(v2)))
                    switch v12 {
                    case "":
                        if true {
                            switch (v12 + v12) {
                            case "foo bar":
                                _ = len(v1)
                            case "hello":
                                v2 -= (^(v2))
                            case "":
                                v12 = ""
                            }
                            if false {
                                switch fi17 {
                                case 106:
                                    v5 = append(v5, (fi17 ^ (fi17 ^ fi17)))
                                case 150:
                                    delete(v4, 2)
                                }
                                if false {
                                    v18 := GenS3{F0: 115}
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    v12 = ""
                                    _ = v18
                                } else {
                                    var v19 int16 = 883
                                    v20 := v3[fi17]
                                    v12 = v12
                                    var v21 bool = true
                                    v5 = append(v5, (-(fi17)))
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            } else {
                                v14 = v14
                                v14 -= (v14 ^ v14)
                                v1 = append(v1, v12)
                            }
                            var v22 int = fi17
                            _ = v22
                        }
                    case "foo bar":
                        v23 := GenS1{F0: 1}
                        _ = v23
                    }
                    v14 += 0
                    _ = fi17
                }
            case 1:
                switch fi15 {
                case 952:
                    var v24 int8 = -75
                    _ = v24
                case 656:
                    if (-128 == 67) {
                        fmt.Println(v12)
                        if true {
                            for fi25 := 0; fi25 < 17; fi25++ {
                                var v26 float64 = 0.1
                                fmt.Println(v12)
                                if false {
                                    v27 := GenS2{F0: 1}
                                    var v28 bool = (!true)
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 float64 = (float64(0.0) * (1e10 + -2.25))
                                    v3[(fi15 << ((90) & 63))] = "foo bar"
                                    _ = len(v3)
                                    _ = v29
                                }
                                _ = fi25
                                _ = v26
                            }
                            for fi30 := 0; fi30 < 14; fi30++ {
                                if (!true) {
                                    switch "hello" {
                                    case "foo bar":
                                        if (2 >= 128) {
                                            if false {
                                                if ((255 == 234) && (false && false)) {
                                                    switch v12 {
                                                    case "":
                                                        var v31 int8 = -128
                                                        _ = v31
                                                    case "hello":
                                                        if len(v1) > 0 {
                                                            v1[0] = v12
                                                        }
                                                    }
                                                    var v32 float32 = (3.14159 / 1e10)
                                                    if ((0 >= 1) && (2 <= (v2 % (((v2 << ((701) & 63))) | 1)))) {
                                                        switch v12 {
                                                        case "hello":
                                                            v33 := v4[247]
                                                            _ = v33
                                                        case "\n\t":
                                                            if true {
                                                                switch (fi30 & fi15) {
                                                                case -51:
                                                                    if ((float32((0.1 / 0.1)) <= ((v32 * v32) / float32((3.14159 / 1.0)))) && true) {
                                                                        v4[588] = (fi30 << ((25) & 63))
                                                                        switch (fi15 + (fi30 & fi30)) {
                                                                        case 4294967295:
                                                                            v12 = ""
                                                                        case 9223372036854775807:
                                                                            v12 = ""
                                                                            v34 := GenS3{F0: 42}
                                                                            var v35 int64 = v14
                                                                            _ = v34
                                                                            _ = v35
                                                                        default:
                                                                            var v36 string = (("" + "") + ((v12 + v12) + ((("hello" + "\n\t") + v12) + v12)))
                                                                            _ = v36
                                                                        }
                                                                        if (v2 <= v2) {
                                                                            v32 = v32
                                                                        } else {
                                                                            switch v12 {
                                                                            case "\n\t":
                                                                                var v37 int32 = 2147483647
                                                                                _ = v37
                                                                            case "hello":
                                                                                var v38 float64 = (3.14159 * 1e10)
                                                                                _ = v38
                                                                            case "foo bar":
                                                                                var v39 int = (-(fi30))
                                                                                _ = v39
                                                                            default:
                                                                                var v40 string = "a"
                                                                                _ = v40
                                                                            }
                                                                            delete(v3, fi30)
                                                                            fmt.Println(fi30, v12, v14)
                                                                        }
                                                                    }
                                                                    switch fi15 {
                                                                    case 127:
                                                                        switch v12 {
                                                                        case "\n\t":
                                                                            v41 := GenS4{F0: v12, F1: 0, F2: 4294967295}
                                                                            _ = v41
                                                                        case "hello":
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                fmt.Println(v32, v12, fi30)
                                                                            default:
                                                                                v42 := GenS5{F0: true, F1: (!true), F2: v14}
                                                                                var v43 int16 = 653
                                                                                _ = v42
                                                                                _ = v43
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                    case -640:
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v44 := GenS6{F0: 65535, F1: (2 != 0), F2: (v12 + "\n\t")}
                                                                            _ = v44
                                                                        case "a":
                                                                            if false {
                                                                                var v45 int64 = 0
                                                                                _ = v45
                                                                            } else {
                                                                                var v46 int32 = 893
                                                                                v47 := GenS4{F0: "foo bar", F1: 2147483648, F2: 416}
                                                                                _ = v46
                                                                                _ = v47
                                                                            }
                                                                        }
                                                                    case -2147483648:
                                                                        switch ((v12 + (v12 + (v12 + v12))) + v12) {
                                                                        case "":
                                                                            var v48 string = v12
                                                                            _ = v48
                                                                        case "hello":
                                                                            v14 = (v14 * 297)
                                                                        }
                                                                    }
                                                                    v2 = (v2 / ((v2) | 1))
                                                                case 128:
                                                                    var v49 int32 = 2147483647
                                                                    _ = v49
                                                                case -62:
                                                                    v50 := GenS7{F0: 255, F1: 0}
                                                                    _ = v50
                                                                }
                                                                var v51 string = (v12 + v12)
                                                                _ = v51
                                                            }
                                                        case "":
                                                            var v52 uint = 255
                                                            _ = v52
                                                        }
                                                        if (-1 < -815) {
                                                            if (2 != uint64(v2)) {
                                                                if (1e10 <= float32(0.1)) {
                                                                    var v53 uint64 = 0
                                                                    if (226 == 1) {
                                                                        v54 := GenS1{F0: v2}
                                                                        switch (v12 + v12) {
                                                                        case "hello":
                                                                            v14 = v14
                                                                        case "a":
                                                                            switch "a" {
                                                                            case "a":
                                                                                switch (fi30 * (fi30 * (fi30 | (fi30 % ((32768) | 1))))) {
                                                                                case 4294967295:
                                                                                    v55 := GenS8{F0: fi15}
                                                                                    _ = v55
                                                                                default:
                                                                                    fmt.Println(v32, v12, v12)
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = fi15
                                                                                    }
                                                                                    if true {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        var v56 uint64 = v53
                                                                                        v53 += 4294967295
                                                                                        _ = v56
                                                                                    }
                                                                                    switch (("a" + "foo bar") + v12) {
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            var v57 int16 = -32768
                                                                                            if (fi30 > (fi15 ^ (fi30 & 156))) {
                                                                                                v2 -= 65535
                                                                                                v2 = (^(v2))
                                                                                                if false {
                                                                                                    switch (fi15 - fi30) {
                                                                                                    case 32767:
                                                                                                        v58 := v4[266]
                                                                                                        _ = v58
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v57)
                                                                                                    var v59 float64 = (-2.25 - 0.1)
                                                                                                    var v60 float64 = 1e10
                                                                                                    _ = v59
                                                                                                    _ = v60
                                                                                                }
                                                                                                v61 := GenS1{F0: (^(v2))}
                                                                                                v62 := GenS9{F0: (1 <= 128), F1: 1}
                                                                                                _ = v61
                                                                                                _ = v62
                                                                                            }
                                                                                            fmt.Println(v14)
                                                                                            var v63 bool = ((645 <= 2147483647) || (true && (true && true)))
                                                                                            _ = v57
                                                                                            _ = v63
                                                                                        }
                                                                                    case "hello":
                                                                                        v2 += 1
                                                                                        var v64 uint = 1
                                                                                        _ = v64
                                                                                    }
                                                                                    v2 -= 2
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (fi15 + fi15)
                                                                                    }
                                                                                    var v65 uint8 = 0
                                                                                    _ = v65
                                                                                }
                                                                                v32 += v32
                                                                                v12 = v12
                                                                            }
                                                                        case "\n\t":
                                                                            var v66 uint64 = 42
                                                                            _ = v66
                                                                        default:
                                                                            switch fi30 {
                                                                            case 2147483647:
                                                                                switch v12 {
                                                                                case "a":
                                                                                    var v67 float32 = v32
                                                                                    _ = v67
                                                                                }
                                                                            default:
                                                                                if ((!((true || true) || false)) || (1000 != 167)) {
                                                                                    v68 := v4[905]
                                                                                    v69 := GenS7{F0: 255, F1: -1}
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                } else {
                                                                                    v32 += v32
                                                                                    delete(v3, -9223372036854775808)
                                                                                }
                                                                            }
                                                                        }
                                                                        v14 -= (v14 * v14)
                                                                        var v70 uint8 = 42
                                                                        _ = v54
                                                                        _ = v70
                                                                    } else {
                                                                        switch fi30 {
                                                                        case 347:
                                                                            v32 = (v32 / v32)
                                                                        case -9223372036854775808:
                                                                            v14 = v14
                                                                        default:
                                                                            if true {
                                                                                if (!(("\n\t" < ((v12 + ((v12 + ((("hello" + "\n\t") + "\n\t") + v12)) + (((v12 + v12) + v12) + (("\n\t" + (v12 + v12)) + (v12 + v12))))) + v12)) && (1.5 < 3.14159))) {
                                                                                    fmt.Println(v2)
                                                                                    var v71 string = v12
                                                                                    _ = v71
                                                                                } else {
                                                                                    v53 = 42
                                                                                    v53 += 1
                                                                                    var v72 int64 = (v14 | v14)
                                                                                    if (3.14159 == (-2.25 / (0.1 * -2.25))) {
                                                                                        v2 += v2
                                                                                        var v73 uint32 = 912
                                                                                        var v74 uint32 = (v73 / (((v73 | v73)) | 1))
                                                                                        _ = v73
                                                                                        _ = v74
                                                                                    } else {
                                                                                        var v75 int8 = 98
                                                                                        v76 := GenS1{F0: 761}
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    }
                                                                                    _ = v72
                                                                                }
                                                                                var v77 uint16 = (v2 ^ v2)
                                                                                _ = v77
                                                                            } else {
                                                                                var v78 uint32 = 140
                                                                                v79 := GenS10{F0: "a"}
                                                                                _ = v78
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                        var v80 uint32 = 0
                                                                        if (!((v12 != ("" + v12)) && false)) {
                                                                            v81 := GenS5{F0: (2147483647 <= 1), F1: (-46 <= -2147483648), F2: v14}
                                                                            var v82 float32 = 0.0
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            if false {
                                                                                v80 = 42
                                                                                v83 := GenS11{F0: 2147483647}
                                                                                _ = len(v4)
                                                                                _ = v83
                                                                            }
                                                                        }
                                                                        _ = v80
                                                                    }
                                                                    v12 = "a"
                                                                    v32 -= (v32 * -2.25)
                                                                    v12 = "hello"
                                                                    _ = v53
                                                                }
                                                                delete(v4, 1)
                                                                v84 := GenS12{F0: -2147483648, F1: (v14 - v14)}
                                                                var v85 int8 = -128
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                v12 = "foo bar"
                                                            }
                                                            v86 := GenS9{F0: ((fi30 * (fi30 & (fi30 | fi30))) >= (fi15 - 128)), F1: 218}
                                                            if true {
                                                                v14 = (v14 & (v14 >> ((42) & 63)))
                                                                var v87 bool = ((3.14159 <= (1e10 - 1e10)) && true)
                                                                var v88 uint32 = 819
                                                                _ = v87
                                                                _ = v88
                                                            } else {
                                                                v89 := v3[-571]
                                                                v90 := v3[4294967295]
                                                                fmt.Println(v32, v2)
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            if (1000 > -498) {
                                                                v91 := GenS10{F0: v12}
                                                                _ = v91
                                                            } else {
                                                                switch ((((v12 + "a") + "foo bar") + "\n\t") + v12) {
                                                                case "":
                                                                    var v92 float64 = 0.0
                                                                    _ = v92
                                                                default:
                                                                    fmt.Println(fi15)
                                                                }
                                                                fmt.Println(v2, v32, v14)
                                                                var v93 float32 = float32((3.14159 / 3.14159))
                                                                var v94 float32 = (v93 / float32(1.5))
                                                                var v95 uint64 = 931
                                                                _ = v93
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                            v96 := GenS13{F0: (^(v2))}
                                                            var v97 uint = 149
                                                            v98 := GenS14{F0: (v2 - (v2 + 65535))}
                                                            _ = v86
                                                            _ = v96
                                                            _ = v97
                                                            _ = v98
                                                        } else {
                                                            var v99 uint64 = 561
                                                            _ = v99
                                                        }
                                                    }
                                                    _ = v32
                                                } else {
                                                    v100 := GenS12{F0: -650, F1: v14}
                                                    _ = v100
                                                }
                                                v14 -= -9223372036854775808
                                            } else {
                                                v101 := GenS13{F0: v2}
                                                v102 := GenS15{F0: v14, F1: 1, F2: (v14 & (v14 % (((v14 >> ((42) & 63))) | 1)))}
                                                if true {
                                                    fmt.Println(fi30, fi15, v14)
                                                    v14 += (v14 % ((v14) | 1))
                                                    v5 = append(v5, (fi30 / (((fi30 % (((fi15 << ((4294967295) & 63))) | 1))) | 1)))
                                                }
                                                _ = v101
                                                _ = v102
                                            }
                                            if (!true) {
                                                v4[646] = (fi30 - (fi15 & fi15))
                                                var v103 uint32 = 4294967295
                                                _ = v103
                                            } else {
                                                v14 = (v14 ^ v14)
                                            }
                                            if true {
                                                v2 = v2
                                            }
                                        } else {
                                            v104 := GenS16{F0: fi30}
                                            v1 = append(v1, "a")
                                            fmt.Println(v2)
                                            _ = v104
                                        }
                                    case "":
                                        v12 = v12
                                        v105 := GenS2{F0: 1}
                                        if (-2.25 > -2.25) {
                                            v14 = -47
                                            switch fi30 {
                                            case 32768:
                                                var v106 uint16 = (v2 & v2)
                                                _ = v106
                                            case 127:
                                                var v107 int = fi30
                                                _ = v107
                                            default:
                                                if (true && (!true)) {
                                                    if true {
                                                        switch fi30 {
                                                        case 463:
                                                            fmt.Println(v14, fi15, v2)
                                                        }
                                                        if (((false || (true && true)) || (1 < 65535)) || (!false)) {
                                                            var v108 uint64 = 0
                                                            var v109 uint = 255
                                                            var v110 float32 = ((-2.25 + 0.0) * 1.5)
                                                            _ = v108
                                                            _ = v109
                                                            _ = v110
                                                        } else {
                                                            v12 = (v12 + v12)
                                                        }
                                                        if false {
                                                            if false {
                                                                if false {
                                                                    v12 = ("\n\t" + (v12 + "foo bar"))
                                                                    if (1 > 1000) {
                                                                        var v111 int = fi30
                                                                        _ = v111
                                                                    } else {
                                                                        if ((3.14159 / 0.1) == ((3.14159 - 1e10) - 0.1)) {
                                                                            v112 := GenS17{F0: ((-128 != 0) || false), F1: 4294967295}
                                                                            _ = len(v1)
                                                                            _ = v112
                                                                        } else {
                                                                            v2 = 394
                                                                            v113 := GenS18{F0: fi15, F1: -124, F2: (fi15 ^ -9223372036854775808)}
                                                                            if (!((!true) || (-856 <= -32768))) {
                                                                                _ = len(v1)
                                                                                v114 := GenS9{F0: ((float64((3.14159 / -2.25)) / 3.14159) < 1e10), F1: uint8(v2)}
                                                                                _ = v114
                                                                            } else {
                                                                                v14 = v14
                                                                                v4[816] = fi15
                                                                                v2 += (v2 + (^(v2)))
                                                                            }
                                                                            _ = v113
                                                                        }
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v14)
                                                                        var v115 int8 = int8(fi15)
                                                                        _ = v115
                                                                    }
                                                                    if len(v5) > 0 {
                                                                        v5[0] = 32768
                                                                    }
                                                                    v116 := GenS14{F0: 42}
                                                                    _ = v116
                                                                }
                                                                var v117 uint8 = 127
                                                                v1 = append(v1, v12)
                                                                _ = v117
                                                            } else {
                                                                fmt.Println(fi30)
                                                                var v118 int8 = 23
                                                                v119 := GenS9{F0: (985 < 273), F1: 42}
                                                                _ = v118
                                                                _ = v119
                                                            }
                                                            v14 += (v14 << ((42) & 63))
                                                            v14 -= (v14 >> ((4294967295) & 63))
                                                        } else {
                                                            fmt.Println(v14)
                                                            fmt.Println(v14, v12, fi15)
                                                            var v120 uint64 = 18446744073709551615
                                                            _ = v120
                                                        }
                                                    }
                                                    if (0 != 2147483648) {
                                                        v121 := GenS19{F0: (("\n\t" + "a") + v12), F1: 18446744073709551615}
                                                        v122 := GenS20{F0: 1000}
                                                        var v123 uint = 255
                                                        v124 := GenS21{F0: 0, F1: fi15, F2: ("foo bar" + (v12 + ""))}
                                                        v123 = v123
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                        _ = v124
                                                    } else {
                                                        v125 := GenS20{F0: 32767}
                                                        if len(v1) > 0 {
                                                            v1[0] = (v12 + "\n\t")
                                                        }
                                                        _ = v125
                                                    }
                                                    var v126 float64 = (-2.25 - 1e10)
                                                    _ = v126
                                                } else {
                                                    var v127 uint16 = v2
                                                    _ = v127
                                                }
                                            }
                                            if false {
                                                v2 -= (v2 * (v2 & 423))
                                            }
                                        }
                                        _ = v105
                                    default:
                                        v128 := GenS14{F0: 32767}
                                        _ = v128
                                    }
                                }
                                _ = fi30
                            }
                            v14 -= (v14 ^ (v14 >> ((4294967295) & 63)))
                        } else {
                            for fi129 := 0; fi129 < 20; fi129++ {
                                var v130 uint8 = 3
                                var v131 string = v12
                                if true {
                                    var v132 uint32 = 42
                                    switch (fi129 / ((fi15) | 1)) {
                                    case -2147483648:
                                        fmt.Println(fi15, v130)
                                    default:
                                        var v133 bool = false
                                        _ = v133
                                    }
                                    v1 = append(v1, ("hello" + v12))
                                    _ = v132
                                } else {
                                    v14 += 1
                                    if ((^(v130)) == 128) {
                                        _ = len(v5)
                                        switch 32767 {
                                        case 0:
                                            v131 = v12
                                        case -589:
                                            var v134 int8 = -94
                                            _ = v134
                                        case 128:
                                            var v135 int8 = 40
                                            _ = v135
                                        default:
                                            var v136 int = (fi129 << ((18446744073709551615) & 63))
                                            _ = v136
                                        }
                                        var v137 int32 = -2147483648
                                        if false {
                                            v2 = 406
                                            if false {
                                                v2 += (v2 % (((^(v2))) | 1))
                                                v138 := GenS9{F0: (1000 != 32767), F1: (^(v130))}
                                                var v139 uint64 = 746
                                                _ = v138
                                                _ = v139
                                            } else {
                                                if false {
                                                    _ = len(v5)
                                                } else {
                                                    delete(v4, 4294967295)
                                                    v14 -= (v14 >> ((2) & 63))
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                if (42 < 1) {
                                                    v140 := GenS17{F0: false, F1: 0}
                                                    v141 := GenS22{F0: fi15}
                                                    var v142 uint8 = (v130 % ((v130) | 1))
                                                    _ = v140
                                                    _ = v141
                                                    _ = v142
                                                } else {
                                                    fmt.Println(v14)
                                                }
                                            }
                                        } else {
                                            v143 := GenS19{F0: v131, F1: 65535}
                                            if false {
                                                _ = len(v1)
                                                _ = len(v4)
                                                delete(v4, 42)
                                                if false {
                                                    v144 := GenS23{F0: 1000, F1: -9223372036854775808, F2: 1}
                                                    v131 = v12
                                                    _ = v144
                                                } else {
                                                    v14 -= (^(v14))
                                                    if len(v5) > 0 {
                                                        v5[0] = fi15
                                                    }
                                                }
                                            }
                                            v1 = append(v1, (v12 + "hello"))
                                            _ = v143
                                        }
                                        _ = v137
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v145 := GenS24{F0: (fi129 + fi129), F1: (v2 + v2)}
                                        var v146 float32 = 1.5
                                        _ = v145
                                        _ = v146
                                    }
                                    _ = len(v3)
                                }
                                _ = fi129
                                _ = v130
                                _ = v131
                            }
                            switch "" {
                            case "hello":
                                switch fi15 {
                                case 32767:
                                    if true {
                                        var v147 uint32 = 42
                                        if (true && ("foo bar" < (v12 + v12))) {
                                            delete(v4, v147)
                                            var v148 float64 = 1.5
                                            _ = v148
                                        }
                                        _ = v147
                                    } else {
                                        v149 := GenS21{F0: (v14 & (^(v14))), F1: fi15, F2: (((v12 + (("" + v12) + (v12 + (v12 + v12)))) + v12) + v12)}
                                        v2 = (v2 - 1)
                                        v5 = append(v5, (fi15 + (-(fi15))))
                                        _ = v149
                                    }
                                case 1:
                                    for fi150 := 0; fi150 < 1; fi150++ {
                                        switch fi15 {
                                        case 1:
                                            var v151 int64 = v14
                                            if (0 != int16(v2)) {
                                                if false {
                                                    var v152 int16 = 963
                                                    v153 := v3[(fi150 | fi150)]
                                                    _ = v152
                                                    _ = v153
                                                } else {
                                                    var v154 float64 = -2.25
                                                    v155 := v4[2147483648]
                                                    v151 -= v14
                                                    _ = v154
                                                    _ = v155
                                                }
                                                v2 = 32767
                                            }
                                            var v156 int32 = 42
                                            _ = v151
                                            _ = v156
                                        case 9223372036854775807:
                                            var v157 int = -551
                                            v12 = v12
                                            _ = v157
                                        }
                                        var v158 int16 = 403
                                        v12 = (("foo bar" + (v12 + v12)) + "foo bar")
                                        var v159 int32 = -1
                                        _ = len(v5)
                                        _ = fi150
                                        _ = v158
                                        _ = v159
                                    }
                                }
                            case "":
                                _ = len(v5)
                            case "\n\t":
                                v160 := v4[2]
                                _ = v160
                            default:
                                var v161 int16 = -557
                                _ = v161
                            }
                        }
                        var v162 uint64 = 2
                        _ = v162
                    }
                case -819:
                    switch -2147483648 {
                    case -9223372036854775808:
                        if false {
                            var v163 bool = true
                            switch (fi15 >> ((18446744073709551615) & 63)) {
                            case 2147483647:
                                if ((v14 ^ v14) >= v14) {
                                    var v164 int32 = 2147483647
                                    for fi165 := 0; fi165 < 14; fi165++ {
                                        var v166 uint8 = 255
                                        if false {
                                            switch (fi165 & fi15) {
                                            case -838:
                                                var v167 bool = (-21 != (v14 & 42))
                                                _ = v167
                                            case -926:
                                                v2 -= v2
                                            default:
                                                var v168 int8 = 100
                                                _ = v168
                                            }
                                        } else {
                                            var v169 float64 = -2.25
                                            fmt.Println(fi165, v163, fi15)
                                            _ = v169
                                        }
                                        var v170 uint32 = 2147483647
                                        v12 = "foo bar"
                                        switch v12 {
                                        case "foo bar":
                                            if (fi165 >= (fi165 / ((fi165) | 1))) {
                                                v171 := GenS2{F0: -128}
                                                if (v2 >= 2) {
                                                    _ = len(v5)
                                                    var v172 int8 = 1
                                                    if false {
                                                        v173 := GenS18{F0: (fi15 >> ((380) & 63)), F1: 1000, F2: fi15}
                                                        _ = v173
                                                    }
                                                    _ = len(v1)
                                                    _ = v172
                                                }
                                                _ = v171
                                            } else {
                                                if (0 == 72) {
                                                    fmt.Println(v12, v14, v170)
                                                    var v174 uint32 = (v170 >> ((909) & 63))
                                                    _ = v174
                                                }
                                            }
                                        case "hello":
                                            v1 = append(v1, v12)
                                        }
                                        _ = fi165
                                        _ = v166
                                        _ = v170
                                    }
                                    v12 = v12
                                    _ = v164
                                } else {
                                    v12 = v12
                                }
                            default:
                                var v175 int32 = 2147483647
                                _ = v175
                            }
                            fmt.Println(v14, fi15)
                            _ = v163
                        }
                    case 4294967295:
                        var v176 uint16 = (v2 >> ((0) & 63))
                        _ = v176
                    case -1:
                        fmt.Println(fi15, v14)
                    }
                }
            default:
                v12 = v12
            }
            v177 := GenS25{F0: 121, F1: int16(v14)}
            _ = v177
        }
        var v178 uint16 = 271
        _ = fi15
        _ = v178
    }
    _ = len(v1)
    v179 := v3[32767]
    _ = len(v5)
    if ("a" == v12) {
        var v180 float64 = (0.0 - 1e10)
        v181 := GenS23{F0: 1, F1: (v14 >> ((40) & 63)), F2: -1}
        _ = v180
        _ = v181
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v179:", v179)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v179
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
