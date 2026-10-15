package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int8; F2 int64 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 string; F1 uint; F2 string }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 uint; F1 string; F2 string }
type GenS9 struct { F0 int64; F1 uint; F2 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int; F1 uint32; F2 int16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int64; F1 uint }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int8; F1 int; F2 uint32 }
type GenS17 struct { F0 uint32; F1 int }
type GenS18 struct { F0 int; F1 uint32 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint8; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{0, 42, 2, 2147483648}
    var v3 uint16 = 806
    v4 := []int{-1, -856, -1, 128}
    var v5 uint = 65535
    v6 := []uint32{847, 65535, 4294967295, 65535}
    if (!false) {
        if ((float64(((1.5 / -2.25) - float32(3.14159))) * -2.25) >= 3.14159) {
            _ = len(v6)
            if false {
                var v7 int8 = 0
                if (false && true) {
                    var v8 string = "hello"
                    _ = v8
                } else {
                    var v9 int32 = -2147483648
                    switch 127 {
                    case -9223372036854775808:
                        var v10 int16 = 1000
                        _ = v10
                    case 51:
                        v9 -= (v9 % (((-(v9))) | 1))
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = v9
                }
                switch "" {
                case "\n\t":
                    if true {
                        if false {
                            switch int(v7) {
                            case 42:
                                fmt.Println(v7)
                            case 32767:
                                if (65535 != 42) {
                                    var v11 float64 = 1.5
                                    fmt.Println(v11)
                                    _ = v11
                                } else {
                                    v12 := GenS1{F0: -9223372036854775808, F1: 0}
                                    fmt.Println(v5)
                                    switch "foo bar" {
                                    case "foo bar":
                                        if false {
                                            v5 = (v5 << (((v5 | v5)) & 63))
                                            if (v3 != (^(v3))) {
                                                v13 := GenS2{F0: true, F1: -136}
                                                var v14 uint32 = 2147483648
                                                _ = v13
                                                _ = v14
                                            } else {
                                                v7 += (v7 / ((127) | 1))
                                                switch 0 {
                                                case 42:
                                                    v1[""] = 2147483647
                                                case 127:
                                                    if (1000 == -740) {
                                                        var v15 bool = true
                                                        _ = v15
                                                    } else {
                                                        if false {
                                                            var v16 uint16 = 630
                                                            v17 := GenS1{F0: -738, F1: -2147483648}
                                                            _ = v16
                                                            _ = v17
                                                        } else {
                                                            var v18 uint16 = (v3 + v3)
                                                            _ = v18
                                                        }
                                                    }
                                                default:
                                                    for fi19 := 0; fi19 < 3; fi19++ {
                                                        var v20 int = (fi19 % (((fi19 + 1)) | 1))
                                                        for fi21 := 0; fi21 < 14; fi21++ {
                                                            v22 := GenS3{F0: int16(v3)}
                                                            v23 := GenS4{F0: 32767, F1: v7, F2: 1}
                                                            _ = fi21
                                                            _ = v22
                                                            _ = v23
                                                        }
                                                        if false {
                                                            _ = len(v1)
                                                        } else {
                                                            switch ("foo bar" + (("foo bar" + "hello") + "foo bar")) {
                                                            case "":
                                                                if false {
                                                                    v5 += (v5 % ((v5) | 1))
                                                                } else {
                                                                    if true {
                                                                        var v24 uint = (v5 * (v5 % (((v5 & v5)) | 1)))
                                                                        v5 = v24
                                                                        _ = v24
                                                                    }
                                                                    var v25 int = -9223372036854775808
                                                                    _ = v25
                                                                }
                                                            case "hello":
                                                                v1["foo bar"] = uint32(v5)
                                                            case "a":
                                                                if false {
                                                                    var v26 int64 = 1
                                                                    v27 := v1[("" + "a")]
                                                                    v20 = v20
                                                                    _ = v26
                                                                    _ = v27
                                                                } else {
                                                                    var v28 uint64 = 1
                                                                    _ = v28
                                                                }
                                                            }
                                                            v29 := GenS4{F0: (v3 >> ((724) & 63)), F1: -2, F2: 0}
                                                            _ = v29
                                                        }
                                                        _ = fi19
                                                        _ = v20
                                                    }
                                                }
                                                fmt.Println(v5, v3)
                                            }
                                        }
                                    case "a":
                                        var v30 int64 = 9223372036854775807
                                        _ = v30
                                    default:
                                        if (1.5 > (1e10 + (3.14159 - 3.14159))) {
                                            var v31 int64 = 120
                                            _ = v31
                                        } else {
                                            switch ("a" + "hello") {
                                            case "\n\t":
                                                if (65535 != 1) {
                                                    switch "" {
                                                    case "":
                                                        v32 := GenS4{F0: (^(v3)), F1: (v7 / ((v7) | 1)), F2: 42}
                                                        _ = v32
                                                    case "\n\t":
                                                        if len(v2) > 0 {
                                                            v2[0] = 643
                                                        }
                                                    case "a":
                                                        var v33 uint32 = 65535
                                                        _ = v33
                                                    }
                                                    var v34 int16 = 42
                                                    _ = v34
                                                }
                                            case "":
                                                if false {
                                                    for fi35 := 0; fi35 < 15; fi35++ {
                                                        v5 = (v5 << ((481) & 63))
                                                        fmt.Println(fi35, v5)
                                                        if true {
                                                            switch fi35 {
                                                            case 767:
                                                                for fi36 := 0; fi36 < 4; fi36++ {
                                                                    v37 := GenS5{F0: v5, F1: int64(v3)}
                                                                    _ = fi36
                                                                    _ = v37
                                                                }
                                                            case 128:
                                                                _ = len(v1)
                                                            case -2147483648:
                                                                var v38 string = "foo bar"
                                                                _ = v38
                                                            }
                                                            var v39 uint64 = 4294967295
                                                            _ = v39
                                                        } else {
                                                            for fi40 := 0; fi40 < 14; fi40++ {
                                                                var v41 int16 = int16(fi40)
                                                                if false {
                                                                    var v42 int16 = -32768
                                                                    v1[("" + "")] = 0
                                                                    v3 = (v3 << ((803) & 63))
                                                                    var v43 uint32 = 178
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v44 uint16 = 1
                                                                    _ = v44
                                                                default:
                                                                    fmt.Println(v5)
                                                                }
                                                                var v45 int32 = 765
                                                                if false {
                                                                    v41 += (v41 >> (((^(v5))) & 63))
                                                                    v3 = 42
                                                                    fmt.Println(v5, v41)
                                                                }
                                                                var v46 uint8 = 225
                                                                if true {
                                                                    fmt.Println(fi35, v46)
                                                                    var v47 string = (("hello" + "foo bar") + "")
                                                                    _ = v47
                                                                } else {
                                                                    if false {
                                                                        var v48 uint16 = v3
                                                                        v49 := v1[("" + "\n\t")]
                                                                        fmt.Println(v7, v7, v3)
                                                                        v3 -= v3
                                                                        switch ("" + "foo bar") {
                                                                        case "a":
                                                                            var v50 uint32 = 65535
                                                                            _ = v50
                                                                        case "hello":
                                                                            if false {
                                                                                v51 := GenS4{F0: v48, F1: v7, F2: 1}
                                                                                switch (fi35 * 128) {
                                                                                case 197:
                                                                                    var v52 float64 = float64(-2.25)
                                                                                    _ = v52
                                                                                case -2147483648:
                                                                                    fmt.Println(v3)
                                                                                case 504:
                                                                                    switch ((("hello" + (("" + "\n\t") + "a")) + "\n\t") + "\n\t") {
                                                                                    case "a":
                                                                                        fmt.Println(fi35, v48)
                                                                                    case "foo bar":
                                                                                        switch ("hello" + "a") {
                                                                                        case "foo bar":
                                                                                            v53 := GenS1{F0: (^(fi40)), F1: (v45 / ((0) | 1))}
                                                                                            _ = v53
                                                                                        case "hello":
                                                                                            v5 += (^(v5))
                                                                                        case "":
                                                                                            switch "" {
                                                                                            case "":
                                                                                                var v54 int16 = -576
                                                                                                _ = v54
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    v48 = (v48 >> (((v5 * v5)) & 63))
                                                                                                    v55 := GenS6{F0: "\n\t", F1: v5, F2: (("foo bar" + (("foo bar" + (((("foo bar" + "foo bar") + (("hello" + "foo bar") + "")) + "a") + ("\n\t" + ("" + "a")))) + "a")) + (("foo bar" + (("foo bar" + ("hello" + (("foo bar" + "") + "\n\t"))) + ("a" + "a"))) + "\n\t"))}
                                                                                                    _ = v55
                                                                                                } else {
                                                                                                    v46 += 42
                                                                                                    var v56 uint = v5
                                                                                                    v46 -= (v46 << (((v5 | (v5 - 748))) & 63))
                                                                                                    _ = v56
                                                                                                }
                                                                                            default:
                                                                                                var v57 uint32 = v49
                                                                                                _ = v57
                                                                                            }
                                                                                        default:
                                                                                            v41 = 656
                                                                                        }
                                                                                    case "\n\t":
                                                                                        switch ("\n\t" + (("\n\t" + "foo bar") + ("foo bar" + "a"))) {
                                                                                        case "a":
                                                                                            fmt.Println(v45, fi40, v49)
                                                                                        case "\n\t":
                                                                                            if true {
                                                                                                fmt.Println(v49, v45, v45)
                                                                                                switch (fi35 | fi35) {
                                                                                                case 127:
                                                                                                    v3 = v48
                                                                                                case -9223372036854775808:
                                                                                                    v5 = v5
                                                                                                case 32768:
                                                                                                    fmt.Println(v46)
                                                                                                default:
                                                                                                    var v58 uint32 = (v49 % (((v49 + 2)) | 1))
                                                                                                    _ = v58
                                                                                                }
                                                                                                v46 = 254
                                                                                                v59 := GenS7{F0: 42, F1: (^(fi35))}
                                                                                                if ((v5 & (v5 << (((^(v5))) & 63))) != (v5 ^ v5)) {
                                                                                                    var v60 uint64 = 726
                                                                                                    _ = v60
                                                                                                } else {
                                                                                                    var v61 float64 = 1.5
                                                                                                    _ = v61
                                                                                                }
                                                                                                _ = v59
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v62 := GenS2{F0: (((v41 | (v41 / ((1) | 1))) < (v41 | v41)) || true), F1: 1}
                                                                                            _ = v62
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    _ = len(v1)
                                                                                }
                                                                                _ = v51
                                                                            } else {
                                                                                var v63 bool = true
                                                                                v64 := GenS8{F0: 255, F1: (("" + "foo bar") + "foo bar"), F2: "hello"}
                                                                                var v65 int64 = 9223372036854775807
                                                                                _ = v63
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                        case "foo bar":
                                                                            if true {
                                                                                if true {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 65535
                                                                                    }
                                                                                } else {
                                                                                    var v66 uint8 = v46
                                                                                    _ = v66
                                                                                }
                                                                            }
                                                                        default:
                                                                            v67 := GenS1{F0: 42, F1: v45}
                                                                            _ = v67
                                                                        }
                                                                        switch (fi35 & -389) {
                                                                        case 128:
                                                                            if false {
                                                                                if true {
                                                                                    v5 = (v5 >> (((v5 + (v5 - v5))) & 63))
                                                                                } else {
                                                                                    var v68 int32 = 183
                                                                                    if (761 == (v48 | v3)) {
                                                                                        v49 -= (v49 - (v49 + uint32(v46)))
                                                                                    } else {
                                                                                        var v69 string = ""
                                                                                        _ = v69
                                                                                    }
                                                                                    _ = v68
                                                                                }
                                                                                var v70 uint8 = uint8(v5)
                                                                                if false {
                                                                                    v3 -= uint16(v49)
                                                                                } else {
                                                                                    switch ("hello" + "foo bar") {
                                                                                    case "\n\t":
                                                                                        var v71 int8 = v7
                                                                                        _ = v71
                                                                                    case "":
                                                                                        v72 := GenS9{F0: 1, F1: v5, F2: (^(v49))}
                                                                                        _ = v72
                                                                                    }
                                                                                    switch 32768 {
                                                                                    case 32768:
                                                                                        v41 -= v41
                                                                                        if false {
                                                                                            v49 = v49
                                                                                            fmt.Println(v45)
                                                                                            var v73 float32 = 3.14159
                                                                                            _ = v73
                                                                                        }
                                                                                        if (!((v5 + v5) <= v5)) {
                                                                                            var v74 bool = (1 == 102)
                                                                                            v75 := GenS8{F0: 65535, F1: "hello", F2: "foo bar"}
                                                                                            var v76 uint32 = 2147483648
                                                                                            _ = v74
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        } else {
                                                                                            v77 := GenS3{F0: (v41 / ((int16(v7)) | 1))}
                                                                                            fmt.Println(v5, v5, fi40)
                                                                                            v41 += (v41 % (((v41 / (((v41 ^ 1000)) | 1))) | 1))
                                                                                            if (true && ((v41 | 546) <= (v41 / ((v41) | 1)))) {
                                                                                                var v78 int64 = 42
                                                                                                fmt.Println(v48, v70)
                                                                                                _ = v78
                                                                                            }
                                                                                            _ = v77
                                                                                        }
                                                                                        var v79 uint64 = 18446744073709551615
                                                                                        var v80 bool = (61 != (v70 >> (((^(v5))) & 63)))
                                                                                        var v81 float64 = (1e10 / -2.25)
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    default:
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    v45 -= 0
                                                                                }
                                                                                _ = v70
                                                                            } else {
                                                                                fmt.Println(v46)
                                                                            }
                                                                        case -1:
                                                                            switch (fi35 & (fi40 & (fi35 & int(v3)))) {
                                                                            case -42:
                                                                                var v82 int64 = int64(v5)
                                                                                _ = v82
                                                                            case -854:
                                                                                fmt.Println(fi40)
                                                                            case 32768:
                                                                                var v83 bool = (uint8(v3) > v46)
                                                                                _ = v83
                                                                            default:
                                                                                v48 = 0
                                                                            }
                                                                        case -2147483648:
                                                                            if (true && (42 <= 62)) {
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    v6 = append(v6, 180)
                                                                                case "\n\t":
                                                                                    fmt.Println(fi35)
                                                                                case "":
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        v41 -= v41
                                                                                    case "hello":
                                                                                        v84 := v1["hello"]
                                                                                        fmt.Println(v7, v46, fi40)
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = (fi35 + int(v5))
                                                                                        }
                                                                                        _ = v84
                                                                                    case "\n\t":
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v85 uint64 = uint64(v3)
                                                                                v85 -= 2
                                                                                v85 -= v85
                                                                                _ = v85
                                                                            } else {
                                                                                var v86 int = 9223372036854775807
                                                                                v87 := v1["hello"]
                                                                                var v88 float64 = (((0.0 + (0.0 * (1e10 + (((float64(0.0) * (0.0 * float64(1e10))) * 1.5) - -2.25)))) - 1e10) - 1e10)
                                                                                _ = v86
                                                                                _ = v87
                                                                                _ = v88
                                                                            }
                                                                        default:
                                                                            fmt.Println(v7, v3, v7)
                                                                        }
                                                                        var v89 string = ("" + "hello")
                                                                        _ = v48
                                                                        _ = v49
                                                                        _ = v89
                                                                    }
                                                                    var v90 int = fi40
                                                                    fmt.Println(fi35, v90)
                                                                    _ = len(v1)
                                                                    v91 := v1[""]
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                                _ = fi40
                                                                _ = v41
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        }
                                                        _ = fi35
                                                    }
                                                    delete(v1, ("\n\t" + "hello"))
                                                    for fi92 := 0; fi92 < 8; fi92++ {
                                                        fmt.Println(v3, v5, v3)
                                                        _ = fi92
                                                    }
                                                } else {
                                                    _ = len(v1)
                                                    var v93 int8 = 1
                                                    v94 := GenS10{F0: (v7 << ((2) & 63))}
                                                    _ = v93
                                                    _ = v94
                                                }
                                            default:
                                                var v95 uint32 = 2147483647
                                                _ = v95
                                            }
                                            var v96 uint8 = 128
                                            _ = v96
                                        }
                                    }
                                    _ = v12
                                }
                            case 128:
                                for fi97 := 0; fi97 < 6; fi97++ {
                                    if false {
                                        v98 := GenS2{F0: (9223372036854775807 <= -1), F1: 1000}
                                        _ = v98
                                    }
                                    _ = fi97
                                }
                            default:
                                v99 := GenS10{F0: 127}
                                _ = v99
                            }
                        } else {
                            if false {
                                var v100 int8 = v7
                                _ = v100
                            } else {
                                for fi101 := 0; fi101 < 9; fi101++ {
                                    if (!(!(624 < 1000))) {
                                        fmt.Println(v7)
                                        for fi102 := 0; fi102 < 6; fi102++ {
                                            v7 -= v7
                                            switch "foo bar" {
                                            case "foo bar":
                                                if false {
                                                    var v103 uint8 = 2
                                                    _ = v103
                                                }
                                            case "a":
                                                var v104 int32 = 2147483647
                                                _ = v104
                                            }
                                            _ = fi102
                                        }
                                        if (!false) {
                                            switch ("foo bar" + "\n\t") {
                                            case "foo bar":
                                                if true {
                                                    v105 := GenS11{F0: (fi101 << ((226) & 63)), F1: 0, F2: 1000}
                                                    if ((1 < int64(v5)) && (!(-776 <= 1))) {
                                                        fmt.Println(v5)
                                                    }
                                                    if ((!(!false)) || false) {
                                                        v106 := GenS4{F0: 685, F1: (v7 / ((-1) | 1)), F2: int64(v5)}
                                                        var v107 uint16 = v3
                                                        switch "" {
                                                        case "foo bar":
                                                            var v108 uint16 = v3
                                                            _ = v108
                                                        case "hello":
                                                            if ((-(v7)) == (v7 << (((^(v5))) & 63))) {
                                                                v109 := v1["a"]
                                                                var v110 uint32 = v109
                                                                _ = v109
                                                                _ = v110
                                                            } else {
                                                                v111 := GenS12{F0: v7}
                                                                _ = v111
                                                            }
                                                        }
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                    _ = v105
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch (fi101 + (fi101 << ((349) & 63))) {
                                                    case 2147483647:
                                                        v7 = 1
                                                    case 872:
                                                        if true {
                                                            for fi112 := 0; fi112 < 12; fi112++ {
                                                                v3 = v3
                                                                v5 += (v5 & v5)
                                                                fmt.Println(fi101, fi101)
                                                                _ = fi112
                                                            }
                                                            _ = len(v6)
                                                            fmt.Println(v3, v3, v3)
                                                        }
                                                    case 1:
                                                        for fi113 := 0; fi113 < 3; fi113++ {
                                                            fmt.Println(v3, fi113, v3)
                                                            var v114 uint16 = (^(v3))
                                                            v5 = 692
                                                            v115 := GenS13{F0: ""}
                                                            var v116 uint64 = 214
                                                            _ = fi113
                                                            _ = v114
                                                            _ = v115
                                                            _ = v116
                                                        }
                                                    }
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                            v117 := GenS1{F0: fi101, F1: -973}
                                            v7 = (v7 >> (((^(v5))) & 63))
                                            _ = v117
                                        } else {
                                            v118 := GenS11{F0: 0, F1: 126, F2: -32768}
                                            if true {
                                                var v119 uint16 = 457
                                                _ = v119
                                            }
                                            _ = v118
                                        }
                                    } else {
                                        v5 -= (v5 + v5)
                                        switch fi101 {
                                        case -9223372036854775808:
                                            if false {
                                                v120 := GenS14{F0: -9223372036854775808, F1: 42}
                                                _ = v120
                                            }
                                        case 634:
                                            var v121 uint64 = 4294967295
                                            _ = v121
                                        default:
                                            v7 += -1
                                        }
                                        var v122 int64 = 1
                                        _ = v122
                                    }
                                    _ = fi101
                                }
                            }
                            fmt.Println(v7, v3)
                        }
                        v123 := GenS7{F0: 6, F1: 2147483647}
                        _ = v123
                    }
                }
                _ = v7
            } else {
                v124 := GenS14{F0: -803, F1: v5}
                v125 := GenS15{F0: -1}
                v126 := GenS16{F0: -114, F1: 32768, F2: 577}
                _ = v124
                _ = v125
                _ = v126
            }
        }
        v6 = append(v6, 65535)
    }
    v5 -= v5
    if (841 <= 2147483647) {
        v5 += v5
        v5 -= v5
        v127 := GenS17{F0: 0, F1: 32767}
        _ = v127
    } else {
        v128 := GenS8{F0: 65535, F1: "foo bar", F2: ""}
        _ = v128
    }
    for fi129 := 0; fi129 < 10; fi129++ {
        v130 := GenS18{F0: fi129, F1: 2}
        var v131 int16 = 42
        _ = fi129
        _ = v130
        _ = v131
    }
    v132 := GenS11{F0: -367, F1: 65535, F2: -1}
    for fi133 := 0; fi133 < 8; fi133++ {
        v6 = append(v6, 301)
        v5 = (^(v5))
        _ = fi133
    }
    var v134 float64 = ((0.1 * ((-2.25 * -2.25) / (1e10 + 0.1))) - 1.5)
    var v135 int8 = 96
    if true {
        v3 += v3
    } else {
        for fi136 := 0; fi136 < 11; fi136++ {
            switch "foo bar" {
            case "\n\t":
                v3 -= (v3 ^ 683)
            case "a":
                fmt.Println(v3)
            case "foo bar":
                v5 = uint(v3)
            default:
                var v137 uint8 = 128
                _ = v137
            }
            for fi138 := 0; fi138 < 17; fi138++ {
                fmt.Println(fi136)
                switch (fi138 << ((v5) & 63)) {
                case 38:
                    var v139 uint32 = 1
                    _ = v139
                default:
                    if (!true) {
                        v135 += v135
                        if (((0.1 + v134) <= v134) || (!(false && (!(!true))))) {
                            fmt.Println(v3, fi136)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v135 += v135
                            v140 := GenS10{F0: v135}
                            switch "hello" {
                            case "hello":
                                v1["\n\t"] = 554
                            case "a":
                                v141 := GenS9{F0: 959, F1: (v5 & v5), F2: 272}
                                _ = v141
                            default:
                                if len(v6) > 0 {
                                    v6[0] = 0
                                }
                            }
                            _ = v140
                        }
                        fmt.Println(fi136)
                        v142 := GenS13{F0: ""}
                        v3 = (v3 + v3)
                        _ = v142
                    }
                }
                var v143 int32 = -1
                _ = fi138
                _ = v143
            }
            var v144 int = (fi136 | (fi136 ^ (fi136 | (-(fi136)))))
            _ = fi136
            _ = v144
        }
    }
    var v145 uint16 = v3
    switch 9223372036854775807 {
    case 127:
        var v146 int32 = 42
        _ = v146
    default:
        if false {
            fmt.Println(v3, v145, v3)
            var v147 bool = (v134 >= 0.0)
            switch "a" {
            case "a":
                v3 = 800
            case "":
                v148 := GenS19{F0: v5}
                _ = v148
            }
            _ = v147
        } else {
            v134 = float64(-2.25)
        }
    }
    for fi149 := 0; fi149 < 16; fi149++ {
        v150 := GenS7{F0: 0, F1: 2147483647}
        fmt.Println(v3, v134, v134)
        switch fi149 {
        case 169:
            var v151 bool = ((v5 | v5) < v5)
            _ = v151
        }
        _ = fi149
        _ = v150
    }
    var v152 bool = (!(!true))
    switch "foo bar" {
    case "hello":
        for fi153 := 0; fi153 < 15; fi153++ {
            var v154 uint8 = 226
            v154 -= (v154 % (((v154 | (^(v154)))) | 1))
            _ = fi153
            _ = v154
        }
    }
    if v152 {
        v5 = (v5 ^ v5)
    } else {
        if ((false || (!(32767 >= -329))) || (1 == v3)) {
            var v155 int = 32767
            v5 += v5
            _ = v155
        } else {
            v135 -= v135
            v145 = (v3 / (((v3 - 746)) | 1))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = len(v2)
        switch 42 {
        case -9223372036854775808:
            if true {
                var v156 int16 = 32767
                _ = v156
            } else {
                var v157 uint16 = v3
                for fi158 := 0; fi158 < 8; fi158++ {
                    v159 := GenS1{F0: fi158, F1: 0}
                    _ = fi158
                    _ = v159
                }
                _ = v157
            }
        }
    }
    v135 = (v135 ^ 75)
    if (!false) {
        if ("foo bar" != "a") {
            v160 := GenS20{F0: 2, F1: 4294967295, F2: 1}
            var v161 int32 = -1
            _ = v160
            _ = v161
        }
    } else {
        switch 0 {
        case 127:
            var v162 float32 = 1.5
            _ = v162
        }
        fmt.Println(v152)
        v3 += (v3 | v145)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v134:", v134)
    fmt.Println("v135:", v135)
    fmt.Println("v145:", v145)
    fmt.Println("v152:", v152)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v132
    _ = v134
    _ = v135
    _ = v145
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
