package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 bool; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 int32; F2 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int8; F1 uint8 }
type GenS12 struct { F0 bool; F1 string }
type GenS13 struct { F0 int64; F1 int32; F2 uint8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int; F1 uint8; F2 string }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint64; F1 uint32; F2 string }
type GenS18 struct { F0 bool; F1 int8 }
type GenS19 struct { F0 uint32; F1 int16; F2 int8 }
type GenS20 struct { F0 int8; F1 int64 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 int16 }
type GenS24 struct { F0 int8; F1 uint; F2 uint }
type GenS25 struct { F0 uint64; F1 bool }

func main() {
    var v1 uint8 = 214
    var v2 int16 = 32767
    var v3 uint16 = 610
    v4 := make(map[int]int)
    v5 := []string{"foo bar"}
    var v6 bool = true
    v6 = true
    var v7 uint8 = v1
    v2 += -1
    _ = len(v4)
    switch (("a" + "hello") + ("foo bar" + (("" + "") + ("\n\t" + "a")))) {
    case "":
        if true {
            if v6 {
                v1 += 237
                var v8 int32 = 2147483647
                _ = v8
            } else {
                v3 = (v3 * (^(v3)))
                v6 = true
            }
        }
    case "a":
        v7 = 127
    default:
        if (!(!((v3 <= v3) && true))) {
            v3 += v3
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 int64 = 1
                v11 := GenS1{F0: 0, F1: 2, F2: 249}
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    }
    v12 := GenS1{F0: 0, F1: 65535, F2: 18446744073709551615}
    switch 2147483647 {
    case 2147483647:
        v3 -= v3
    case 32767:
        for fi13 := 0; fi13 < 11; fi13++ {
            if v6 {
                for fi14 := 0; fi14 < 6; fi14++ {
                    if v6 {
                        v15 := v4[(fi13 | fi14)]
                        _ = v15
                    }
                    v5 = append(v5, "foo bar")
                    _ = fi14
                }
            } else {
                var v16 float32 = float32(0.0)
                if ((v1 & 81) > (v1 << ((735) & 63))) {
                    v16 = 1.5
                    fmt.Println(v3, v6, v3)
                    for fi17 := 0; fi17 < 12; fi17++ {
                        _ = len(v4)
                        fmt.Println(v1, v2, v3)
                        if (true && v6) {
                            v18 := GenS1{F0: int32(v3), F1: 691, F2: 251}
                            if v6 {
                                if (!false) {
                                    v5 = append(v5, ("\n\t" + "\n\t"))
                                    v19 := GenS1{F0: -610, F1: 2147483648, F2: 42}
                                    if v6 {
                                        v20 := GenS1{F0: -1, F1: 2, F2: 255}
                                        v2 = (-(v2))
                                        _ = v20
                                    } else {
                                        v16 = float32(1.5)
                                        if (true || v6) {
                                            v1 += v1
                                            v21 := GenS1{F0: 987, F1: 1, F2: 0}
                                            v4[fi17] = (fi13 & fi17)
                                            _ = v21
                                        } else {
                                            v22 := GenS1{F0: 0, F1: 833, F2: 255}
                                            if ((uint8(v3) <= (v1 / ((v1) | 1))) || (v6 && ((v7 > v7) && true))) {
                                                var v23 int16 = (v2 - v2)
                                                v24 := GenS1{F0: 1, F1: 2, F2: 42}
                                                _ = v23
                                                _ = v24
                                            }
                                            var v25 float64 = ((0.1 + float64(v16)) * (-2.25 * (3.14159 / 0.1)))
                                            _ = v22
                                            _ = v25
                                        }
                                        v16 -= v16
                                        var v26 string = "hello"
                                        _ = v26
                                    }
                                    _ = v19
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        if false {
                                            v27 := GenS1{F0: 42, F1: 65535, F2: 2}
                                            var v28 uint64 = 18446744073709551615
                                            v6 = false
                                            _ = v27
                                            _ = v28
                                        } else {
                                            if v6 {
                                                var v29 int16 = v2
                                                if ((!v6) || (((v7 + 42) > (^(v1))) || ((v7 / (((^(v1))) | 1)) == (v1 << ((137) & 63))))) {
                                                    if v6 {
                                                        var v30 uint16 = (^(v3))
                                                        _ = v30
                                                    } else {
                                                        if false {
                                                            v7 = v1
                                                            var v31 int8 = -32
                                                            _ = v31
                                                        } else {
                                                            fmt.Println(v16, v6, v2)
                                                            v4[(fi17 - -9223372036854775808)] = fi17
                                                            v2 -= (v2 | 42)
                                                        }
                                                        switch -1 {
                                                        case 931:
                                                            var v32 float64 = 1.5
                                                            _ = v32
                                                        case -385:
                                                            v3 = (^(v3))
                                                            if v6 {
                                                                v2 -= (v29 | v2)
                                                            } else {
                                                                var v33 float64 = 1.5
                                                                _ = v33
                                                            }
                                                        case -1:
                                                            v6 = v6
                                                        default:
                                                            if true {
                                                                v29 -= (-(v29))
                                                            } else {
                                                                var v34 uint64 = 1
                                                                if true {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        switch (fi13 << ((18446744073709551615) & 63)) {
                                                                        case -2147483648:
                                                                            var v35 int = 0
                                                                            _ = v35
                                                                        case 0:
                                                                            v36 := GenS1{F0: 1, F1: 2147483647, F2: 783}
                                                                            _ = v36
                                                                        case 42:
                                                                            if v6 {
                                                                                v37 := GenS1{F0: -199, F1: 2, F2: 42}
                                                                                _ = v37
                                                                            } else {
                                                                                if (1e10 >= (1.5 / 1.5)) {
                                                                                    var v38 bool = v6
                                                                                    _ = v38
                                                                                }
                                                                                fmt.Println(v34, v6, v29)
                                                                            }
                                                                            if len(v5) > 0 {
                                                                                v5[0] = "hello"
                                                                            }
                                                                        default:
                                                                            v39 := GenS2{F0: -1}
                                                                            _ = v39
                                                                        }
                                                                    case "\n\t":
                                                                        if ((42 != v3) || (v6 || v6)) {
                                                                            if false {
                                                                                v40 := GenS3{F0: v3, F1: 128, F2: ""}
                                                                                var v41 float64 = (-2.25 + (1e10 - float64(float32((0.1 / 1.0)))))
                                                                                _ = v40
                                                                                _ = v41
                                                                            }
                                                                            _ = len(v5)
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                if true {
                                                                                    v4[(fi17 % ((fi17) | 1))] = (-(fi17))
                                                                                    if v6 {
                                                                                        var v42 int8 = 51
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            v43 := v4[(fi17 + (fi13 << ((255) & 63)))]
                                                                                            var v44 int64 = 9223372036854775807
                                                                                            var v45 uint32 = 1
                                                                                            _ = v43
                                                                                            _ = v44
                                                                                            _ = v45
                                                                                        case "":
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                        default:
                                                                                            if (!(fi17 <= 4294967295)) {
                                                                                                v1 += v1
                                                                                                var v46 int64 = 345
                                                                                                if ("a" >= "\n\t") {
                                                                                                    v3 = v3
                                                                                                    switch (fi17 * (fi17 & (^(fi13)))) {
                                                                                                    case 32768:
                                                                                                        if v6 {
                                                                                                            var v47 bool = false
                                                                                                            v48 := GenS4{F0: ("" + "hello")}
                                                                                                            _ = v47
                                                                                                            _ = v48
                                                                                                        }
                                                                                                        v49 := v4[(^(fi13))]
                                                                                                        _ = v49
                                                                                                    case 127:
                                                                                                        var v50 int64 = v46
                                                                                                        _ = v50
                                                                                                    case 4294967295:
                                                                                                        v16 -= v16
                                                                                                        var v51 int16 = v2
                                                                                                        var v52 int16 = int16(v3)
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            var v53 uint64 = (v34 | (v34 ^ v34))
                                                                                                            fmt.Println(v7)
                                                                                                            v2 = (v29 << ((970) & 63))
                                                                                                            _ = v53
                                                                                                        case "\n\t":
                                                                                                            if (!(false && (v42 != v42))) {
                                                                                                                var v54 uint32 = 2147483648
                                                                                                                _ = v54
                                                                                                            }
                                                                                                        case "a":
                                                                                                            var v55 int64 = (v46 & v46)
                                                                                                            var v56 uint64 = (^(v34))
                                                                                                            _ = v55
                                                                                                            _ = v56
                                                                                                        }
                                                                                                        switch (-(fi13)) {
                                                                                                        case 9223372036854775807:
                                                                                                            var v57 float64 = (3.14159 + 0.1)
                                                                                                            _ = v57
                                                                                                        }
                                                                                                        _ = v51
                                                                                                        _ = v52
                                                                                                    }
                                                                                                    if (("" >= (("a" + ("foo bar" + ("" + ("hello" + "a")))) + ("hello" + "foo bar"))) && v6) {
                                                                                                        if (1 >= (v34 & (v34 & 18446744073709551615))) {
                                                                                                            var v58 uint8 = (v1 - 2)
                                                                                                            _ = v58
                                                                                                        }
                                                                                                        v46 = -71
                                                                                                        if true {
                                                                                                            if ((v42 * 1) == v42) {
                                                                                                                v59 := GenS5{F0: -78, F1: false, F2: (v29 & (v29 & int16(fi17)))}
                                                                                                                var v60 float64 = 1e10
                                                                                                                var v61 int = (fi13 >> ((261) & 63))
                                                                                                                switch ("\n\t" + ((((("foo bar" + "hello") + ("hello" + "foo bar")) + "\n\t") + "a") + "a")) {
                                                                                                                case "foo bar":
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v46, v29, v16)
                                                                                                                }
                                                                                                                _ = v59
                                                                                                                _ = v60
                                                                                                                _ = v61
                                                                                                            } else {
                                                                                                                var v62 uint64 = v34
                                                                                                                if v6 {
                                                                                                                    if (403 > 65535) {
                                                                                                                        v7 = v1
                                                                                                                        var v63 bool = false
                                                                                                                        _ = v63
                                                                                                                    }
                                                                                                                    switch (fi17 + (^(fi17))) {
                                                                                                                    case -2147483648:
                                                                                                                        if v6 {
                                                                                                                            v64 := GenS6{F0: (v42 >> ((758) & 63))}
                                                                                                                            switch -2147483648 {
                                                                                                                            case 127:
                                                                                                                                v46 -= (v46 ^ (v46 / ((42) | 1)))
                                                                                                                            case 32767:
                                                                                                                                var v65 float32 = (v16 / 0.1)
                                                                                                                                _ = v65
                                                                                                                            default:
                                                                                                                                var v66 uint = uint(v1)
                                                                                                                                _ = v66
                                                                                                                            }
                                                                                                                            var v67 int8 = (v42 - (v42 << ((42) & 63)))
                                                                                                                            v68 := v4[32768]
                                                                                                                            _ = v64
                                                                                                                            _ = v67
                                                                                                                            _ = v68
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch "foo bar" {
                                                                                                                    case "\n\t":
                                                                                                                        var v69 string = "\n\t"
                                                                                                                        _ = v69
                                                                                                                    default:
                                                                                                                        if len(v5) >= 2 {
                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(v46)
                                                                                                                    var v70 bool = false
                                                                                                                    _ = v70
                                                                                                                }
                                                                                                                var v71 int = -2147483648
                                                                                                                _ = v62
                                                                                                                _ = v71
                                                                                                            }
                                                                                                            switch (fi17 << ((4294967295) & 63)) {
                                                                                                            case 9223372036854775807:
                                                                                                                var v72 float32 = (v16 / 1.5)
                                                                                                                _ = v72
                                                                                                            default:
                                                                                                                v42 = v42
                                                                                                            }
                                                                                                            v29 -= 32767
                                                                                                        } else {
                                                                                                            if (((v34 | (^(v34))) <= 18446744073709551615) && ((!true) && v6)) {
                                                                                                                v46 = (-(v46))
                                                                                                                if v6 {
                                                                                                                    fmt.Println(v42)
                                                                                                                }
                                                                                                                switch "" {
                                                                                                                case "\n\t":
                                                                                                                    if (!v6) {
                                                                                                                        var v73 uint16 = v3
                                                                                                                        _ = v73
                                                                                                                    } else {
                                                                                                                        fmt.Println(v3, v3)
                                                                                                                        var v74 float64 = 1.5
                                                                                                                        var v75 int64 = v46
                                                                                                                        _ = v74
                                                                                                                        _ = v75
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    v76 := GenS6{F0: (v42 << ((18446744073709551615) & 63))}
                                                                                                                    fmt.Println(v2)
                                                                                                                    var v77 string = "a"
                                                                                                                    _ = v76
                                                                                                                    _ = v77
                                                                                                                case "a":
                                                                                                                    if true {
                                                                                                                        v78 := GenS7{F0: (v29 - (v2 + (v29 | v29))), F1: 1, F2: v3}
                                                                                                                        var v79 int = 32767
                                                                                                                        _ = v78
                                                                                                                        _ = v79
                                                                                                                    } else {
                                                                                                                        var v80 string = "\n\t"
                                                                                                                        _ = v80
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v81 float32 = float32(-2.25)
                                                                                                                    _ = v81
                                                                                                                }
                                                                                                            } else {
                                                                                                                v46 += (v46 | int64(v1))
                                                                                                                v82 := GenS1{F0: -1, F1: 42, F2: 0}
                                                                                                                var v83 bool = (!(357 < 0))
                                                                                                                _ = v82
                                                                                                                _ = v83
                                                                                                            }
                                                                                                            var v84 int = fi17
                                                                                                            v46 += -9223372036854775808
                                                                                                            _ = len(v4)
                                                                                                            _ = v84
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(fi13, v46)
                                                                                                        v42 = 127
                                                                                                    }
                                                                                                }
                                                                                                v29 = (v2 + v2)
                                                                                                v42 -= (v42 * v42)
                                                                                                _ = v46
                                                                                            }
                                                                                        }
                                                                                        v5 = append(v5, "\n\t")
                                                                                        if false {
                                                                                            switch ((("a" + "hello") + "hello") + ("hello" + ((("foo bar" + ("a" + ((("foo bar" + "") + "a") + "\n\t"))) + ("foo bar" + ("" + ("foo bar" + "foo bar")))) + "hello"))) {
                                                                                            case "a":
                                                                                                if (uint8(v34) != v7) {
                                                                                                    var v85 uint = 220
                                                                                                    var v86 float64 = float64(v16)
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        var v87 uint8 = v1
                                                                                                        v16 = v16
                                                                                                        v1 = uint8(v3)
                                                                                                        fmt.Println(v16, v6)
                                                                                                        if false {
                                                                                                            var v88 uint64 = v34
                                                                                                            _ = v88
                                                                                                        }
                                                                                                        _ = v87
                                                                                                    case "\n\t":
                                                                                                        var v89 int32 = 0
                                                                                                        _ = v89
                                                                                                    case "foo bar":
                                                                                                        v90 := GenS4{F0: "foo bar"}
                                                                                                        _ = v90
                                                                                                    default:
                                                                                                        v91 := GenS1{F0: 265, F1: 2147483647, F2: (v85 & 662)}
                                                                                                        v92 := v4[fi17]
                                                                                                        var v93 int64 = 700
                                                                                                        v29 = v2
                                                                                                        if true {
                                                                                                            v2 += -1
                                                                                                            if (-830 == (-(v2))) {
                                                                                                                switch (fi17 | (v92 / ((fi13) | 1))) {
                                                                                                                case 9223372036854775807:
                                                                                                                    if (2147483648 != 4294967295) {
                                                                                                                        if false {
                                                                                                                            var v94 uint64 = v34
                                                                                                                            _ = v94
                                                                                                                        } else {
                                                                                                                            if v6 {
                                                                                                                                v93 -= (v93 ^ 42)
                                                                                                                                v2 = int16(v85)
                                                                                                                            } else {
                                                                                                                                fmt.Println(v1, v7, v93)
                                                                                                                            }
                                                                                                                            fmt.Println(fi17)
                                                                                                                        }
                                                                                                                        var v95 uint8 = (^(v7))
                                                                                                                        v5 = append(v5, "\n\t")
                                                                                                                        _ = v95
                                                                                                                    } else {
                                                                                                                        if v6 {
                                                                                                                            var v96 uint32 = 0
                                                                                                                            var v97 int32 = 42
                                                                                                                            _ = v96
                                                                                                                            _ = v97
                                                                                                                        } else {
                                                                                                                            if (0 > (fi17 ^ (fi17 * fi13))) {
                                                                                                                                var v98 int16 = (v29 / (((v2 << ((v85) & 63))) | 1))
                                                                                                                                var v99 int64 = int64(v92)
                                                                                                                                var v100 uint = (v85 / (((v85 & uint(v7))) | 1))
                                                                                                                                _ = v100
                                                                                                                                _ = v98
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                            switch (fi13 % ((v92) | 1)) {
                                                                                                                            case 2147483647:
                                                                                                                                v16 = (-2.25 * (v16 * (v16 / float32(v86))))
                                                                                                                                v1 = v1
                                                                                                                                v101 := GenS5{F0: v42, F1: (0 <= -2147483648), F2: v29}
                                                                                                                                _ = v101
                                                                                                                            case -344:
                                                                                                                                v102 := GenS8{F0: v29}
                                                                                                                                _ = v102
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if ((v93 - v93) > int64(v7)) {
                                                                                                                        v16 = 1e10
                                                                                                                        var v103 uint32 = 922
                                                                                                                        _ = v103
                                                                                                                    } else {
                                                                                                                        var v104 float64 = ((v86 + (float64((0.1 / (v16 * v16))) - 3.14159)) - (-2.25 + float64((3.14159 + v16))))
                                                                                                                        v16 -= (v16 / v16)
                                                                                                                        v105 := GenS9{F0: 1, F1: -1}
                                                                                                                        _ = v104
                                                                                                                        _ = v105
                                                                                                                    }
                                                                                                                }
                                                                                                                v106 := GenS10{F0: (v34 | (v34 * 2))}
                                                                                                                _ = v106
                                                                                                            } else {
                                                                                                                v7 -= uint8(v34)
                                                                                                                var v107 uint8 = (v7 | 198)
                                                                                                                var v108 int32 = 873
                                                                                                                v109 := GenS11{F0: int8(v2), F1: (v7 << ((v85) & 63))}
                                                                                                                _ = v107
                                                                                                                _ = v108
                                                                                                                _ = v109
                                                                                                            }
                                                                                                            if (((!v6) || (!(uint32(v1) != 4294967295))) && v6) {
                                                                                                                if (930 == -2147483648) {
                                                                                                                    v3 = (^(v3))
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = len(v4)
                                                                                                            } else {
                                                                                                                fmt.Println(v86, v85)
                                                                                                                if (v93 != (v93 + (v93 << ((909) & 63)))) {
                                                                                                                    if ((-(v92)) != (^(v92))) {
                                                                                                                        v42 = 0
                                                                                                                    }
                                                                                                                    v110 := v4[(fi17 ^ (v92 % (((fi13 + v92)) | 1)))]
                                                                                                                    switch "hello" {
                                                                                                                    case "a":
                                                                                                                        v85 += 440
                                                                                                                    case "":
                                                                                                                        v111 := GenS12{F0: v6, F1: ("" + (("hello" + "hello") + (("foo bar" + "") + "foo bar")))}
                                                                                                                        _ = v111
                                                                                                                    default:
                                                                                                                        v16 += (v16 + 3.14159)
                                                                                                                    }
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                                v5 = append(v5, "\n\t")
                                                                                                            }
                                                                                                        }
                                                                                                        v112 := GenS11{F0: v42, F1: v7}
                                                                                                        _ = v112
                                                                                                        _ = v91
                                                                                                        _ = v92
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                }
                                                                                            }
                                                                                            var v113 string = ("" + "a")
                                                                                            _ = v113
                                                                                        } else {
                                                                                            var v114 uint16 = v3
                                                                                            v2 -= v29
                                                                                            _ = v114
                                                                                        }
                                                                                        var v115 float64 = 3.14159
                                                                                        var v116 int = (fi17 - fi17)
                                                                                        var v117 float64 = v115
                                                                                        v118 := GenS9{F0: v2, F1: 1}
                                                                                        var v119 uint16 = v3
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                        _ = v117
                                                                                        _ = v118
                                                                                        _ = v119
                                                                                        _ = v42
                                                                                    }
                                                                                    v34 = v34
                                                                                } else {
                                                                                    fmt.Println(v29, fi13, v34)
                                                                                    v29 += (v2 & v29)
                                                                                    v4[-816] = -9223372036854775808
                                                                                    v29 += v2
                                                                                }
                                                                            case "a":
                                                                                if v6 {
                                                                                    switch ("foo bar" + "foo bar") {
                                                                                    case "hello":
                                                                                        if (v6 || (!v6)) {
                                                                                            if (-39 != 127) {
                                                                                                _ = len(v5)
                                                                                            } else {
                                                                                                if ((^(v29)) == -1) {
                                                                                                    v5 = append(v5, "hello")
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = "a"
                                                                                                    }
                                                                                                }
                                                                                                var v120 uint = 81
                                                                                                var v121 bool = false
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                            }
                                                                                            switch "a" {
                                                                                            case "hello":
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    var v122 float64 = 0.1
                                                                                                    _ = v122
                                                                                                case "foo bar":
                                                                                                    switch ("foo bar" + ((("hello" + (("" + "a") + "")) + "a") + "\n\t")) {
                                                                                                    case "\n\t":
                                                                                                        if (784 > 65535) {
                                                                                                            fmt.Println(v2)
                                                                                                            switch 127 {
                                                                                                            case 32767:
                                                                                                                v123 := GenS5{F0: 1, F1: false, F2: (v2 % ((v2) | 1))}
                                                                                                                _ = v123
                                                                                                            case 2147483647:
                                                                                                                var v124 uint16 = 65535
                                                                                                                _ = v124
                                                                                                            case 4294967295:
                                                                                                                switch 32767 {
                                                                                                                case 32767:
                                                                                                                    v34 = (v34 + 42)
                                                                                                                default:
                                                                                                                    if (4294967295 > 390) {
                                                                                                                        _ = len(v5)
                                                                                                                    } else {
                                                                                                                        v6 = (!((^(v7)) == v1))
                                                                                                                        v125 := v4[907]
                                                                                                                        var v126 uint32 = 4294967295
                                                                                                                        _ = v125
                                                                                                                        _ = v126
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            v127 := GenS1{F0: -822, F1: 393, F2: 454}
                                                                                                            _ = v127
                                                                                                        }
                                                                                                    case "":
                                                                                                        v6 = (42 >= 42)
                                                                                                    default:
                                                                                                        var v128 int = -1
                                                                                                        _ = v128
                                                                                                    }
                                                                                                default:
                                                                                                    if v6 {
                                                                                                        v129 := v4[(fi13 * (fi13 >> ((1) & 63)))]
                                                                                                        var v130 float64 = 1.5
                                                                                                        if v6 {
                                                                                                            switch ("hello" + (("hello" + "\n\t") + "\n\t")) {
                                                                                                            case "":
                                                                                                                if (0.1 < v16) {
                                                                                                                    switch 42 {
                                                                                                                    case 100:
                                                                                                                        var v131 uint8 = v7
                                                                                                                        _ = v131
                                                                                                                    }
                                                                                                                    if v6 {
                                                                                                                        v130 -= v130
                                                                                                                        v132 := GenS13{F0: -94, F1: 42, F2: 255}
                                                                                                                        delete(v4, 690)
                                                                                                                        _ = v132
                                                                                                                    }
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                _ = len(v4)
                                                                                                            case "\n\t":
                                                                                                                v4[fi17] = int(v3)
                                                                                                            default:
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                            }
                                                                                                            switch 32768 {
                                                                                                            case 32768:
                                                                                                                var v133 float32 = (1.5 - (float32(v130) - (0.0 - ((3.14159 - v16) * ((v16 * 1e10) - 3.14159)))))
                                                                                                                _ = v133
                                                                                                            case 359:
                                                                                                                v1 += 128
                                                                                                                var v134 string = ("foo bar" + "foo bar")
                                                                                                                var v135 int32 = 0
                                                                                                                v7 += (v7 % (((v1 - v1)) | 1))
                                                                                                                _ = v134
                                                                                                                _ = v135
                                                                                                            default:
                                                                                                                v29 += v29
                                                                                                            }
                                                                                                            var v136 uint32 = 2147483647
                                                                                                            _ = v136
                                                                                                        } else {
                                                                                                            if ((v29 >> ((4294967295) & 63)) > v29) {
                                                                                                                fmt.Println(v7, v129)
                                                                                                                var v137 int32 = 1
                                                                                                                var v138 int32 = (v137 | v137)
                                                                                                                _ = v137
                                                                                                                _ = v138
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v129
                                                                                                        _ = v130
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v16 -= v16
                                                                                                var v139 int64 = -9223372036854775808
                                                                                                switch (fi13 + fi17) {
                                                                                                case 128:
                                                                                                    _ = len(v5)
                                                                                                case 4294967295:
                                                                                                    var v140 uint8 = (^(v1))
                                                                                                    _ = v140
                                                                                                case 2147483647:
                                                                                                    var v141 int16 = -654
                                                                                                    _ = v141
                                                                                                }
                                                                                                fmt.Println(v6, v29, v29)
                                                                                                v16 = (0.0 * (3.14159 * ((v16 * (v16 / v16)) + 0.1)))
                                                                                                _ = v139
                                                                                            }
                                                                                            switch ("\n\t" + ("" + "a")) {
                                                                                            case "foo bar":
                                                                                                var v142 int32 = 0
                                                                                                _ = v142
                                                                                            case "hello":
                                                                                                if ("hello" == ("hello" + "hello")) {
                                                                                                    v34 += v34
                                                                                                    v1 -= 42
                                                                                                }
                                                                                            case "\n\t":
                                                                                                var v143 bool = v6
                                                                                                _ = v143
                                                                                            }
                                                                                        }
                                                                                        v2 += v2
                                                                                    case "\n\t":
                                                                                        v144 := GenS14{F0: 917}
                                                                                        _ = v144
                                                                                    case "":
                                                                                        if v6 {
                                                                                            if (0 > (^(v3))) {
                                                                                                v6 = (42 >= -1)
                                                                                                v145 := GenS12{F0: ((fi17 + (^(fi17))) == (fi13 / ((fi17) | 1))), F1: "hello"}
                                                                                                var v146 uint = 255
                                                                                                _ = v145
                                                                                                _ = v146
                                                                                            }
                                                                                            var v147 int16 = (v2 & v2)
                                                                                            _ = v147
                                                                                        } else {
                                                                                            if v6 {
                                                                                                if v6 {
                                                                                                    fmt.Println(v1, v29)
                                                                                                } else {
                                                                                                    switch ("foo bar" + (("\n\t" + "\n\t") + "a")) {
                                                                                                    case "foo bar":
                                                                                                        v148 := GenS11{F0: -99, F1: 42}
                                                                                                        _ = v148
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v34)
                                                                                                var v149 int = (fi13 / ((127) | 1))
                                                                                                _ = v149
                                                                                            }
                                                                                            var v150 string = ""
                                                                                            _ = v150
                                                                                        }
                                                                                    default:
                                                                                        if v6 {
                                                                                            var v151 float32 = ((v16 - v16) - v16)
                                                                                            var v152 int16 = (v2 << ((782) & 63))
                                                                                            switch fi13 {
                                                                                            case -1:
                                                                                                fmt.Println(v3, v7)
                                                                                            case 1:
                                                                                                v153 := GenS15{F0: (fi17 + (fi13 >> ((255) & 63))), F1: v1, F2: "hello"}
                                                                                                _ = v153
                                                                                            default:
                                                                                                if (-1 == 1) {
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = "\n\t"
                                                                                                    }
                                                                                                } else {
                                                                                                    v3 = v3
                                                                                                    switch fi17 {
                                                                                                    case 4294967295:
                                                                                                        var v154 int8 = 75
                                                                                                        _ = v154
                                                                                                    case -9223372036854775808:
                                                                                                        var v155 string = ""
                                                                                                        switch "foo bar" {
                                                                                                        case "foo bar":
                                                                                                            var v156 string = v155
                                                                                                            v152 += (v152 % ((1) | 1))
                                                                                                            if v6 {
                                                                                                                fmt.Println(v152, v152, v16)
                                                                                                                if (v6 || (-715 <= -615)) {
                                                                                                                    if v6 {
                                                                                                                        switch v155 {
                                                                                                                        case "\n\t":
                                                                                                                            if (v3 >= (v3 & 2)) {
                                                                                                                                var v157 float32 = ((0.0 / v151) / ((float32(1e10) + ((v151 * v151) / 1e10)) - v151))
                                                                                                                                if len(v5) >= 2 {
                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                }
                                                                                                                                if (0 < 42) {
                                                                                                                                    var v158 float32 = float32(1.5)
                                                                                                                                    if false {
                                                                                                                                        var v159 uint16 = v3
                                                                                                                                        if v6 {
                                                                                                                                            var v160 float64 = (1.5 / (0.1 + 1e10))
                                                                                                                                            fmt.Println(v159, v34, v151)
                                                                                                                                            delete(v4, (fi13 - fi13))
                                                                                                                                            _ = v160
                                                                                                                                        }
                                                                                                                                        v2 += 1
                                                                                                                                        v161 := GenS2{F0: 0}
                                                                                                                                        _ = v159
                                                                                                                                        _ = v161
                                                                                                                                    } else {
                                                                                                                                        if (-1 == v29) {
                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                v5[0] = v156
                                                                                                                                            }
                                                                                                                                            if v6 {
                                                                                                                                                var v162 int8 = 0
                                                                                                                                                _ = v162
                                                                                                                                            } else {
                                                                                                                                                var v163 uint8 = (v1 / ((v7) | 1))
                                                                                                                                                if ((1.5 > 1.5) || v6) {
                                                                                                                                                    var v164 uint64 = (v34 - 18446744073709551615)
                                                                                                                                                    _ = v164
                                                                                                                                                }
                                                                                                                                                var v165 float32 = v151
                                                                                                                                                _ = v163
                                                                                                                                                _ = v165
                                                                                                                                            }
                                                                                                                                            switch fi17 {
                                                                                                                                            case -556:
                                                                                                                                                v16 -= -2.25
                                                                                                                                            case 0:
                                                                                                                                                v2 -= v29
                                                                                                                                            default:
                                                                                                                                                v151 += v151
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            if (true && (!v6)) {
                                                                                                                                                v151 = v151
                                                                                                                                                if true {
                                                                                                                                                    if v6 {
                                                                                                                                                        v166 := GenS16{F0: (^(fi17))}
                                                                                                                                                        if (!(0 >= fi17)) {
                                                                                                                                                            var v167 bool = false
                                                                                                                                                            var v168 int16 = (v2 << ((1) & 63))
                                                                                                                                                            var v169 uint16 = (v3 | v3)
                                                                                                                                                            var v170 uint8 = (^(v7))
                                                                                                                                                            if v167 {
                                                                                                                                                                fmt.Println(v6)
                                                                                                                                                                v34 += (v34 * v34)
                                                                                                                                                                var v171 uint = 4294967295
                                                                                                                                                                _ = v171
                                                                                                                                                            } else {
                                                                                                                                                                v172 := GenS17{F0: v34, F1: 407, F2: v156}
                                                                                                                                                                if true {
                                                                                                                                                                    var v173 bool = ((v156 + v155) < v156)
                                                                                                                                                                    switch "hello" {
                                                                                                                                                                    case "hello":
                                                                                                                                                                        v152 = v152
                                                                                                                                                                    case "\n\t":
                                                                                                                                                                        var v174 float32 = float32(0.0)
                                                                                                                                                                        _ = v174
                                                                                                                                                                    default:
                                                                                                                                                                        v155 = "a"
                                                                                                                                                                    }
                                                                                                                                                                    fmt.Println(v173, v151)
                                                                                                                                                                    if v167 {
                                                                                                                                                                        v173 = v173
                                                                                                                                                                        v151 -= float32(1.5)
                                                                                                                                                                    }
                                                                                                                                                                    fmt.Println(v3, v7, v173)
                                                                                                                                                                    _ = v173
                                                                                                                                                                } else {
                                                                                                                                                                    switch (("" + v155) + "foo bar") {
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        v1 = 2
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                _ = v172
                                                                                                                                                            }
                                                                                                                                                            v175 := GenS6{F0: -102}
                                                                                                                                                            _ = v167
                                                                                                                                                            _ = v168
                                                                                                                                                            _ = v169
                                                                                                                                                            _ = v170
                                                                                                                                                            _ = v175
                                                                                                                                                        }
                                                                                                                                                        _ = v166
                                                                                                                                                    } else {
                                                                                                                                                        if v6 {
                                                                                                                                                            v16 -= v151
                                                                                                                                                            v7 = v1
                                                                                                                                                        } else {
                                                                                                                                                            if (v1 <= v7) {
                                                                                                                                                                fmt.Println(v157, fi17)
                                                                                                                                                                fmt.Println(v151, fi13)
                                                                                                                                                                if v6 {
                                                                                                                                                                    v2 -= -32768
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                var v176 int32 = 42
                                                                                                                                                                _ = v176
                                                                                                                                                            }
                                                                                                                                                            v177 := GenS5{F0: 1, F1: (!true), F2: v29}
                                                                                                                                                            _ = v177
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    v178 := GenS18{F0: ((v34 ^ v34) > v34), F1: -128}
                                                                                                                                                    _ = v178
                                                                                                                                                }
                                                                                                                                                v156 = v155
                                                                                                                                                var v179 string = (v155 + v156)
                                                                                                                                                fmt.Println(v157)
                                                                                                                                                _ = v179
                                                                                                                                            } else {
                                                                                                                                                delete(v4, (fi17 - 1))
                                                                                                                                                v180 := GenS1{F0: -507, F1: uint32(v3), F2: 18446744073709551615}
                                                                                                                                                v158 = (-2.25 - (v151 - v151))
                                                                                                                                                fmt.Println(v156, v155, fi13)
                                                                                                                                                fmt.Println(v7, v155, v34)
                                                                                                                                                _ = v180
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        v6 = v6
                                                                                                                                    }
                                                                                                                                    _ = v158
                                                                                                                                } else {
                                                                                                                                    if len(v5) > 0 {
                                                                                                                                        v5[0] = (("foo bar" + "hello") + v156)
                                                                                                                                    }
                                                                                                                                    var v181 uint64 = 18446744073709551615
                                                                                                                                    _ = v181
                                                                                                                                }
                                                                                                                                _ = v157
                                                                                                                            }
                                                                                                                        case "":
                                                                                                                            fmt.Println(v29, v34)
                                                                                                                        default:
                                                                                                                            fmt.Println(v16)
                                                                                                                        }
                                                                                                                        fmt.Println(v155, v6, v151)
                                                                                                                        var v182 int32 = 0
                                                                                                                        _ = v182
                                                                                                                    } else {
                                                                                                                        var v183 uint64 = v34
                                                                                                                        v16 = (v16 * 1e10)
                                                                                                                        _ = v183
                                                                                                                    }
                                                                                                                    if len(v5) > 0 {
                                                                                                                        v5[0] = v155
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v3, v29)
                                                                                                                if (v3 != 3) {
                                                                                                                    v155 = (v155 + ("\n\t" + v156))
                                                                                                                    var v184 bool = (v6 && false)
                                                                                                                    _ = len(v5)
                                                                                                                    _ = v184
                                                                                                                }
                                                                                                            }
                                                                                                            switch (fi17 >> ((0) & 63)) {
                                                                                                            case -2147483648:
                                                                                                                if v6 {
                                                                                                                    v185 := GenS7{F0: int16(v34), F1: -437, F2: (v3 % ((v3) | 1))}
                                                                                                                    v2 = v152
                                                                                                                    _ = v185
                                                                                                                } else {
                                                                                                                    v16 = 1.5
                                                                                                                    v186 := v4[fi17]
                                                                                                                    _ = v186
                                                                                                                }
                                                                                                            case -226:
                                                                                                                v152 -= (v152 >> ((391) & 63))
                                                                                                            case -909:
                                                                                                                var v187 int8 = -128
                                                                                                                _ = v187
                                                                                                            }
                                                                                                            _ = v156
                                                                                                        case "a":
                                                                                                            fmt.Println(v155, v3)
                                                                                                        default:
                                                                                                            var v188 float64 = -2.25
                                                                                                            _ = v188
                                                                                                        }
                                                                                                        _ = v155
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v151
                                                                                            _ = v152
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v1, v6, v16)
                                                                                    v2 = v29
                                                                                    v16 += v16
                                                                                } else {
                                                                                    var v189 float32 = v16
                                                                                    fmt.Println(v189, v1, v16)
                                                                                    _ = v189
                                                                                }
                                                                            }
                                                                            v190 := GenS16{F0: fi13}
                                                                            v16 = v16
                                                                            var v191 float64 = -2.25
                                                                            if true {
                                                                                v16 -= ((((v16 / v16) / v16) * 0.1) * v16)
                                                                            } else {
                                                                                var v192 int64 = 42
                                                                                var v193 uint64 = 967
                                                                                v194 := GenS11{F0: -1, F1: (v7 >> ((206) & 63))}
                                                                                if (v192 != v192) {
                                                                                    fmt.Println(v193)
                                                                                    if (v192 <= 1) {
                                                                                        v3 = uint16(v193)
                                                                                        v195 := GenS7{F0: v29, F1: -1, F2: 834}
                                                                                        v196 := GenS9{F0: v29, F1: -1}
                                                                                        _ = v195
                                                                                        _ = v196
                                                                                    } else {
                                                                                        fmt.Println(v6, v6)
                                                                                    }
                                                                                    v6 = v6
                                                                                }
                                                                                v6 = v6
                                                                                if v6 {
                                                                                    var v197 bool = true
                                                                                    if false {
                                                                                        if false {
                                                                                            switch fi17 {
                                                                                            case 32767:
                                                                                                fmt.Println(v34)
                                                                                            case -9223372036854775808:
                                                                                                if v197 {
                                                                                                    fmt.Println(v16, v1)
                                                                                                    if false {
                                                                                                        var v198 bool = v6
                                                                                                        _ = v198
                                                                                                    } else {
                                                                                                        var v199 int64 = (^(v192))
                                                                                                        _ = v199
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v7)
                                                                                                }
                                                                                            default:
                                                                                                var v200 uint8 = v1
                                                                                                _ = v200
                                                                                            }
                                                                                            if ((v7 == (v1 + (v7 << ((2) & 63)))) || (2147483647 >= 2147483648)) {
                                                                                                v7 = v1
                                                                                            } else {
                                                                                                var v201 uint64 = (v34 - 42)
                                                                                                v191 += 3.14159
                                                                                                _ = v201
                                                                                            }
                                                                                            fmt.Println(v1, v34, v197)
                                                                                        } else {
                                                                                            v2 = (v29 | (v29 % (((v2 ^ int16(v193))) | 1)))
                                                                                            v202 := GenS4{F0: "a"}
                                                                                            _ = v202
                                                                                        }
                                                                                    } else {
                                                                                        if v197 {
                                                                                            v203 := GenS9{F0: 99, F1: 42}
                                                                                            fmt.Println(v34, v29, v7)
                                                                                            var v204 string = "\n\t"
                                                                                            _ = v203
                                                                                            _ = v204
                                                                                        }
                                                                                        if (376 >= (v192 ^ int64(v1))) {
                                                                                            var v205 int32 = -2147483648
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = ("foo bar" + ("" + "hello"))
                                                                                            }
                                                                                            if v197 {
                                                                                                if true {
                                                                                                    if false {
                                                                                                        if ((v192 * int64(v205)) < (v192 ^ 42)) {
                                                                                                            var v206 int = fi13
                                                                                                            var v207 int8 = 1
                                                                                                            _ = v206
                                                                                                            _ = v207
                                                                                                        } else {
                                                                                                            fmt.Println(v16, v16, v193)
                                                                                                            if (1 >= (v7 / ((v1) | 1))) {
                                                                                                                var v208 uint64 = (v34 / ((v193) | 1))
                                                                                                                v209 := GenS18{F0: true, F1: -1}
                                                                                                                v208 = 2
                                                                                                                _ = v208
                                                                                                                _ = v209
                                                                                                            }
                                                                                                            var v210 int16 = (v2 % ((32767) | 1))
                                                                                                            var v211 float64 = v191
                                                                                                            _ = len(v5)
                                                                                                            v1 = (v1 / ((v7) | 1))
                                                                                                            _ = v210
                                                                                                            _ = v211
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v212 int32 = v205
                                                                                                        var v213 uint8 = v7
                                                                                                        _ = v212
                                                                                                        _ = v213
                                                                                                    }
                                                                                                } else {
                                                                                                    var v214 int32 = (v205 >> ((129) & 63))
                                                                                                    var v215 int8 = 0
                                                                                                    _ = v214
                                                                                                    _ = v215
                                                                                                }
                                                                                                var v216 float32 = (v16 + v16)
                                                                                                _ = len(v4)
                                                                                                _ = v216
                                                                                            }
                                                                                            _ = v205
                                                                                        } else {
                                                                                            if (v16 > v16) {
                                                                                                var v217 float32 = float32(((v191 + v191) + float64(float32(3.14159))))
                                                                                                v218 := GenS19{F0: 1, F1: -32768, F2: -53}
                                                                                                fmt.Println(v6, v16, v217)
                                                                                                v3 += (v3 | 255)
                                                                                                _ = v217
                                                                                                _ = v218
                                                                                            } else {
                                                                                                if (114 > 127) {
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            var v219 int64 = (v192 % ((v192) | 1))
                                                                                                            _ = v219
                                                                                                        case "foo bar":
                                                                                                            v220 := GenS20{F0: -69, F1: 0}
                                                                                                            _ = v220
                                                                                                        case "\n\t":
                                                                                                            var v221 bool = (!((fi13 ^ 2147483647) < fi13))
                                                                                                            var v222 int32 = 1
                                                                                                            var v223 uint16 = 0
                                                                                                            _ = v221
                                                                                                            _ = v222
                                                                                                            _ = v223
                                                                                                        default:
                                                                                                            if (("a" + ("hello" + ("hello" + (("" + "hello") + "a")))) != "") {
                                                                                                                v224 := GenS21{F0: v193}
                                                                                                                var v225 int16 = (v2 % ((-1) | 1))
                                                                                                                if (!(!true)) {
                                                                                                                    switch fi13 {
                                                                                                                    case 829:
                                                                                                                        v226 := GenS13{F0: int64(v7), F1: 874, F2: v7}
                                                                                                                        _ = v226
                                                                                                                    case 32767:
                                                                                                                        var v227 int16 = (v29 / ((v225) | 1))
                                                                                                                        _ = v227
                                                                                                                    case -149:
                                                                                                                        v1 = (v1 / (((v1 * (v7 / ((v1) | 1)))) | 1))
                                                                                                                    default:
                                                                                                                        var v228 uint32 = 2147483648
                                                                                                                        _ = v228
                                                                                                                    }
                                                                                                                }
                                                                                                                var v229 uint16 = (v3 - (v3 & 727))
                                                                                                                if v197 {
                                                                                                                    if false {
                                                                                                                        v1 -= 42
                                                                                                                    } else {
                                                                                                                        var v230 uint32 = 2147483647
                                                                                                                        _ = len(v5)
                                                                                                                        var v231 bool = v6
                                                                                                                        _ = v230
                                                                                                                        _ = v231
                                                                                                                    }
                                                                                                                }
                                                                                                                if v197 {
                                                                                                                    v192 = (-(v192))
                                                                                                                    var v232 uint = 689
                                                                                                                    _ = v232
                                                                                                                }
                                                                                                                _ = v224
                                                                                                                _ = v225
                                                                                                                _ = v229
                                                                                                            }
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v233 uint32 = 2
                                                                                                        _ = v233
                                                                                                    default:
                                                                                                        v1 -= v1
                                                                                                    }
                                                                                                    fmt.Println(v1, v193, v197)
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v234 float32 = 3.14159
                                                                                        _ = v234
                                                                                    }
                                                                                    var v235 int16 = (v2 / (((v29 / ((v2) | 1))) | 1))
                                                                                    if v197 {
                                                                                        v197 = (true && (v197 && (((("foo bar" + "") + "") + ("" + "a")) != (("hello" + "\n\t") + ("a" + "foo bar")))))
                                                                                        var v236 bool = (((((^(v3)) <= v3) || true) && false) || (32767 == v3))
                                                                                        if true {
                                                                                            if true {
                                                                                                var v237 int = fi13
                                                                                                _ = v237
                                                                                            } else {
                                                                                                v191 = v191
                                                                                                v193 = (v34 - v193)
                                                                                                fmt.Println(v16)
                                                                                            }
                                                                                            fmt.Println(v2)
                                                                                            var v238 int = fi17
                                                                                            _ = v238
                                                                                        } else {
                                                                                            var v239 bool = v197
                                                                                            fmt.Println(v3, fi17)
                                                                                            _ = v239
                                                                                        }
                                                                                        fmt.Println(v192)
                                                                                        var v240 uint64 = 650
                                                                                        _ = v236
                                                                                        _ = v240
                                                                                    } else {
                                                                                        if (!v197) {
                                                                                            var v241 float64 = (float64(0.0) - v191)
                                                                                            v5 = append(v5, ("a" + "\n\t"))
                                                                                            var v242 uint32 = 65535
                                                                                            _ = v241
                                                                                            _ = v242
                                                                                        }
                                                                                    }
                                                                                    var v243 uint32 = 4294967295
                                                                                    _ = v197
                                                                                    _ = v235
                                                                                    _ = v243
                                                                                } else {
                                                                                    fmt.Println(v34, fi17, v1)
                                                                                }
                                                                                _ = v192
                                                                                _ = v193
                                                                                _ = v194
                                                                            }
                                                                            _ = v190
                                                                            _ = v191
                                                                        } else {
                                                                            v244 := GenS22{F0: "hello"}
                                                                            v29 += (v29 * (v2 + v2))
                                                                            _ = v244
                                                                        }
                                                                    case "":
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                }
                                                                _ = v34
                                                            }
                                                        }
                                                        v245 := GenS10{F0: 466}
                                                        _ = v245
                                                    }
                                                    v6 = (((v1 - (v7 << ((255) & 63))) == uint8(v3)) && (255 < (v1 << ((255) & 63))))
                                                }
                                                var v246 float64 = float64((v16 / v16))
                                                _ = v246
                                                _ = v29
                                            }
                                            v5 = append(v5, "\n\t")
                                        }
                                    }
                                    v247 := GenS23{F0: v2}
                                    v248 := GenS3{F0: (^(v3)), F1: (v1 % (((v1 ^ 127)) | 1)), F2: "hello"}
                                    _ = v247
                                    _ = v248
                                }
                                var v249 int16 = (v2 * v2)
                                _ = v249
                            }
                            v1 = v1
                            _ = v18
                        } else {
                            v250 := v4[fi13]
                            var v251 uint8 = (^(v7))
                            if v6 {
                                v6 = (v1 <= (v251 * v1))
                                if true {
                                    v252 := GenS19{F0: 2147483647, F1: 1000, F2: -66}
                                    var v253 uint32 = 2147483647
                                    _ = v252
                                    _ = v253
                                }
                            } else {
                                var v254 uint8 = v7
                                var v255 bool = (v6 || (1 > (fi13 >> ((42) & 63))))
                                if (false && false) {
                                    if ((v3 / ((v3) | 1)) == (v3 & v3)) {
                                        if (v16 >= -2.25) {
                                            fmt.Println(fi17, v251)
                                        }
                                        v256 := GenS10{F0: 2}
                                        _ = v256
                                    } else {
                                        v257 := GenS24{F0: -128, F1: uint(v7), F2: 2}
                                        var v258 int64 = 817
                                        v254 = v1
                                        _ = v257
                                        _ = v258
                                    }
                                    var v259 int16 = (v2 % ((v2) | 1))
                                    var v260 uint = 255
                                    _ = v259
                                    _ = v260
                                } else {
                                    if (!(65535 != 451)) {
                                        var v261 uint = 223
                                        var v262 uint32 = 99
                                        _ = v261
                                        _ = v262
                                    }
                                    switch (fi17 + 9223372036854775807) {
                                    case 127:
                                        if (!((!(!(((!false) && v255) || v255))) || false)) {
                                            if true {
                                                var v263 int64 = 1
                                                _ = v263
                                            }
                                        }
                                    case -2147483648:
                                        if v6 {
                                            switch 1 {
                                            case -798:
                                                var v264 bool = v6
                                                _ = v264
                                            default:
                                                switch "" {
                                                case "foo bar":
                                                    var v265 bool = v6
                                                    _ = v265
                                                case "":
                                                    var v266 uint8 = 255
                                                    _ = v266
                                                case "hello":
                                                    delete(v4, fi13)
                                                }
                                            }
                                        }
                                    }
                                    var v267 int = (-(fi13))
                                    _ = v267
                                }
                                var v268 string = "foo bar"
                                _ = v254
                                _ = v255
                                _ = v268
                            }
                            _ = v250
                            _ = v251
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v269 uint64 = 42
                            _ = v269
                        case "hello":
                            v7 += v7
                        }
                        _ = fi17
                    }
                }
                delete(v4, (fi13 << ((4294967295) & 63)))
                _ = v16
            }
            var v270 string = "\n\t"
            v271 := GenS9{F0: (v2 << ((42) & 63)), F1: -291}
            _ = fi13
            _ = v270
            _ = v271
        }
    }
    if v6 {
        switch int(v7) {
        case -962:
            _ = len(v4)
        }
    }
    var v272 float64 = 0.0
    var v273 uint32 = 1
    if (1 >= -1) {
        for fi274 := 0; fi274 < 7; fi274++ {
            fmt.Println(v272, v1)
            for fi275 := 0; fi275 < 13; fi275++ {
                if (-1 > -1) {
                    var v276 float64 = (v272 - (float64(3.14159) + float64(0.1)))
                    switch "hello" {
                    case "hello":
                        var v277 uint32 = 42
                        _ = v277
                    default:
                        switch fi274 {
                        case 128:
                            v276 = 3.14159
                        case -1:
                            fmt.Println(v2, v1, v6)
                        default:
                            var v278 string = (("\n\t" + "hello") + "hello")
                            _ = v278
                        }
                    }
                    v2 = -307
                    _ = v276
                } else {
                    var v279 int8 = 1
                    v273 += 2147483648
                    v280 := GenS6{F0: v279}
                    v4[(fi274 * (fi275 % ((fi275) | 1)))] = fi274
                    _ = v279
                    _ = v280
                }
                v281 := GenS22{F0: ""}
                _ = fi275
                _ = v281
            }
            _ = fi274
        }
        var v282 uint64 = 972
        _ = v282
    } else {
        v3 -= uint16(v1)
        v283 := GenS25{F0: 1, F1: (!false)}
        v1 = v1
        _ = v283
    }
    for fi284 := 0; fi284 < 16; fi284++ {
        fmt.Println(v7, v3, v273)
        _ = fi284
    }
    switch (("\n\t" + "foo bar") + "") {
    case "a":
        v6 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v272:", v272)
    fmt.Println("v273:", v273)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v272
    _ = v273
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
