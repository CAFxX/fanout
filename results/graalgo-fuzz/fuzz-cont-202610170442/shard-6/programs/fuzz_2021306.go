package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint8; F1 uint16; F2 int }
type GenS9 struct { F0 bool; F1 int32; F2 int16 }
type GenS10 struct { F0 int64; F1 uint; F2 bool }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 int64 }
type GenS14 struct { F0 uint32; F1 string }
type GenS15 struct { F0 int; F1 string }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 bool; F1 uint }
type GenS19 struct { F0 int8; F1 uint32 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    v3 := []uint32{65535, 1}
    var v4 float64 = -2.25
    v5 := make(map[uint32]string)
    v6 := make(map[int]uint32)
    if ((((3.14159 / (v4 + (v4 - float64(((1e10 + (1.5 + 1.5)) + (0.1 + ((-2.25 * -2.25) * float32(v4)))))))) < (v4 - 0.1)) || (42 >= 42)) && false) {
        switch ((("hello" + "foo bar") + "a") + "foo bar") {
        case "\n\t":
            fmt.Println(v1)
        case "hello":
            for fi7 := 0; fi7 < 18; fi7++ {
                fmt.Println(fi7)
                fmt.Println(v4, fi7, fi7)
                switch fi7 {
                case 42:
                    fmt.Println(v1)
                }
                _ = fi7
            }
        case "a":
            _ = len(v3)
        }
    }
    if v1 {
        v8 := v6[42]
        if (-59 >= -9223372036854775808) {
            v9 := GenS1{F0: (v8 | v8)}
            v10 := GenS1{F0: (v8 + (v8 & 4294967295))}
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                switch "a" {
                case "a":
                    v4 = v4
                case "foo bar":
                    fmt.Println(v8, fi12, fi12)
                }
                _ = fi12
            }
            v8 -= 42
            for fi13 := 0; fi13 < 11; fi13++ {
                v8 = 4294967295
                var v14 uint16 = 65535
                _ = fi13
                _ = v14
            }
            _ = fi11
        }
        _ = v8
    }
    var v15 uint32 = 0
    v15 += (v15 & 2147483648)
    v16 := GenS1{F0: 1}
    var v17 float64 = (v4 - (v4 - v4))
    switch "\n\t" {
    case "foo bar":
        for fi18 := 0; fi18 < 20; fi18++ {
            if true {
                if ((v1 && (255 >= 2)) || (!true)) {
                    if false {
                        if ((32767 < 32767) && false) {
                            if v1 {
                                var v19 uint8 = uint8(v15)
                                if (((true && v1) || v1) || false) {
                                    if (!(!false)) {
                                        if len(v3) > 0 {
                                            v3[0] = v15
                                        }
                                        v20 := GenS2{F0: (("a" + (("hello" + "hello") + ("a" + ""))) + "a")}
                                        _ = v20
                                    } else {
                                        switch "a" {
                                        case "a":
                                            v21 := GenS3{F0: (v15 - (v15 & 795)), F1: 42, F2: v15}
                                            _ = v21
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "foo bar":
                                                for fi22 := 0; fi22 < 2; fi22++ {
                                                    fmt.Println(v1)
                                                    _ = fi22
                                                }
                                            case "hello":
                                                var v23 int64 = 9223372036854775807
                                                _ = v23
                                            }
                                        case "hello":
                                            switch "\n\t" {
                                            case "hello":
                                                var v24 uint8 = 223
                                                _ = v24
                                            default:
                                                switch fi18 {
                                                case 1:
                                                    switch ("hello" + "") {
                                                    case "\n\t":
                                                        for fi25 := 0; fi25 < 17; fi25++ {
                                                            var v26 string = ("hello" + "a")
                                                            _ = fi25
                                                            _ = v26
                                                        }
                                                    case "foo bar":
                                                        v19 += v19
                                                    case "a":
                                                        v27 := GenS3{F0: (v15 + 2), F1: 362, F2: (v15 << ((1) & 63))}
                                                        _ = v27
                                                    default:
                                                        v1 = (!(-894 != 121))
                                                    }
                                                case -564:
                                                    var v28 string = "\n\t"
                                                    _ = v28
                                                case 42:
                                                    fmt.Println(v15, v15, v1)
                                                default:
                                                    if ((1 > 1) || v1) {
                                                        v29 := GenS3{F0: 4294967295, F1: -1, F2: (^(v15))}
                                                        _ = v29
                                                    }
                                                }
                                            }
                                        default:
                                            var v30 uint16 = 607
                                            _ = v30
                                        }
                                        var v31 int64 = 127
                                        _ = v31
                                    }
                                } else {
                                    var v32 bool = (!v1)
                                    delete(v6, (fi18 - fi18))
                                    fmt.Println(v19, v4)
                                    _ = v32
                                }
                                delete(v5, (v15 / (((^(v15))) | 1)))
                                _ = v19
                            } else {
                                var v33 uint16 = 0
                                v34 := GenS1{F0: v15}
                                v5[v15] = "foo bar"
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v15)
                            v1 = ((-1 != -78) || (32767 == 1000))
                        } else {
                            fmt.Println(fi18)
                            if ((0 >= -1) || (int32(fi18) >= 430)) {
                                if v1 {
                                    if (-1 != -804) {
                                        fmt.Println(v1, v15)
                                    }
                                    v35 := GenS4{F0: "foo bar", F1: (^(v15)), F2: 1}
                                    _ = v35
                                } else {
                                    for fi36 := 0; fi36 < 11; fi36++ {
                                        v37 := GenS4{F0: "\n\t", F1: 2, F2: 853}
                                        v38 := GenS5{F0: 4294967295, F1: 578}
                                        if v1 {
                                            v39 := GenS6{F0: 4294967295, F1: fi18}
                                            switch 143 {
                                            case 1:
                                                fmt.Println(fi36, fi36)
                                            case 32768:
                                                var v40 int64 = -738
                                                _ = v40
                                            default:
                                                var v41 uint8 = 127
                                                _ = v41
                                            }
                                            if false {
                                                var v42 string = ((((("" + "\n\t") + "\n\t") + "") + (("hello" + ("\n\t" + ("\n\t" + "foo bar"))) + "\n\t")) + "")
                                                v3 = append(v3, (^(v15)))
                                                var v43 int16 = 32767
                                                _ = v42
                                                _ = v43
                                            } else {
                                                switch (fi18 & -134) {
                                                case -1:
                                                    v44 := GenS3{F0: (v15 / ((42) | 1)), F1: 42, F2: v15}
                                                    _ = v44
                                                }
                                                fmt.Println(v1)
                                                var v45 int8 = 45
                                                v45 = 118
                                                _ = v45
                                            }
                                            _ = v39
                                        } else {
                                            v17 = 0.1
                                            var v46 int32 = 42
                                            if (!false) {
                                                if (27 == 127) {
                                                    v15 -= (^(v15))
                                                    var v47 uint8 = 47
                                                    v17 = (v4 / v17)
                                                    _ = v47
                                                } else {
                                                    fmt.Println(v15, v46, v1)
                                                    var v48 uint64 = 957
                                                    if v1 {
                                                        if (!v1) {
                                                            if false {
                                                                if (v1 && (!false)) {
                                                                    if ((!(128 < 1)) && (100 > int8(v15))) {
                                                                        var v49 uint64 = (v48 * 2)
                                                                        v4 = v17
                                                                        _ = v49
                                                                    } else {
                                                                        fmt.Println(v15, v1, v17)
                                                                        v48 -= v48
                                                                        delete(v5, uint32(v48))
                                                                    }
                                                                    _ = len(v6)
                                                                }
                                                                v50 := GenS7{F0: 165, F1: 2}
                                                                _ = v50
                                                            } else {
                                                                if ((1.5 != 0.0) && v1) {
                                                                    var v51 int64 = -1
                                                                    _ = v51
                                                                } else {
                                                                    var v52 uint16 = 773
                                                                    v17 = -2.25
                                                                    _ = v52
                                                                }
                                                                if v1 {
                                                                    v53 := GenS6{F0: v48, F1: (fi36 << ((42) & 63))}
                                                                    if false {
                                                                        var v54 uint32 = (^(v15))
                                                                        _ = v54
                                                                    } else {
                                                                        v17 -= ((float64(3.14159) - (float64((((1.5 / 0.1) + 3.14159) + 0.0)) * float64(-2.25))) - v17)
                                                                        v4 += v17
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v55 := GenS6{F0: (v48 % (((^(v48))) | 1)), F1: -2147483648}
                                                                        if true {
                                                                            var v56 int32 = (^(v46))
                                                                            _ = v56
                                                                        }
                                                                        _ = v55
                                                                    }
                                                                    var v57 uint64 = 183
                                                                    fmt.Println(v15, v4, fi18)
                                                                    v57 = (v57 << ((2) & 63))
                                                                    _ = v53
                                                                    _ = v57
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v4 += 3.14159
                                                                    case "\n\t":
                                                                        v3 = append(v3, (v15 & (v15 >> ((255) & 63))))
                                                                    case "hello":
                                                                        fmt.Println(fi18)
                                                                    default:
                                                                        _ = len(v5)
                                                                    }
                                                                    v17 -= ((((0.1 - (float64(1.5) / v17)) * (v4 / 1.5)) / ((v4 * (((1e10 - v4) / (float64((((0.0 + 1e10) + 1e10) - (3.14159 + -2.25))) * (float64(1.5) * 1.5))) * 1e10)) * v17)) + (v4 / ((v4 - v4) - v17)))
                                                                    if v1 {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v58 := GenS6{F0: v48, F1: (fi18 % (((fi36 ^ (fi18 << ((42) & 63)))) | 1))}
                                                                        var v59 uint16 = 65535
                                                                        _ = v58
                                                                        _ = v59
                                                                    } else {
                                                                        fmt.Println(v48, v48, v46)
                                                                        var v60 float32 = -2.25
                                                                        switch (fi18 - fi36) {
                                                                        case -858:
                                                                            v46 -= 0
                                                                        case 9223372036854775807:
                                                                            if false {
                                                                                delete(v2, 835)
                                                                                v17 = float64(v60)
                                                                                if (2 >= 158) {
                                                                                    delete(v5, v15)
                                                                                    fmt.Println(v46, v15, v46)
                                                                                    v61 := GenS8{F0: 127, F1: 255, F2: fi36}
                                                                                    _ = v61
                                                                                }
                                                                            } else {
                                                                                if ((^(v48)) <= (v48 << ((444) & 63))) {
                                                                                    if ((v1 || true) || v1) {
                                                                                        v3 = append(v3, 0)
                                                                                        v60 -= ((v60 + v60) + float32(1e10))
                                                                                    } else {
                                                                                        delete(v5, 2)
                                                                                    }
                                                                                    var v62 int32 = 2147483647
                                                                                    if ((0 != -1) || true) {
                                                                                        fmt.Println(v46)
                                                                                        if v1 {
                                                                                            v63 := GenS5{F0: 32767, F1: 1}
                                                                                            v60 += 0.1
                                                                                            v64 := GenS3{F0: (v15 * v15), F1: -2147483648, F2: v15}
                                                                                            _ = v63
                                                                                            _ = v64
                                                                                        }
                                                                                    } else {
                                                                                        v46 += 42
                                                                                    }
                                                                                    fmt.Println(v17, v48, v46)
                                                                                    _ = v62
                                                                                } else {
                                                                                    switch fi36 {
                                                                                    case 128:
                                                                                        var v65 int64 = 9223372036854775807
                                                                                        _ = v65
                                                                                    case 205:
                                                                                        var v66 int64 = 42
                                                                                        _ = v66
                                                                                    default:
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v46 = v46
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                v60 = v60
                                                                                            case "\n\t":
                                                                                                switch (fi36 - -604) {
                                                                                                case 32768:
                                                                                                    if (false && true) {
                                                                                                        if true {
                                                                                                            fmt.Println(v46, v60)
                                                                                                            var v67 int = (-(fi36))
                                                                                                            v6[v67] = 4294967295
                                                                                                            switch (v67 * (fi18 - fi18)) {
                                                                                                            case 128:
                                                                                                                v17 = v4
                                                                                                            }
                                                                                                            _ = v67
                                                                                                        }
                                                                                                        fmt.Println(fi36, v46, fi36)
                                                                                                        fmt.Println(v15, v48)
                                                                                                    } else {
                                                                                                        v68 := GenS5{F0: (fi18 >> ((266) & 63)), F1: 0}
                                                                                                        _ = v68
                                                                                                    }
                                                                                                case 0:
                                                                                                    if v1 {
                                                                                                        switch (fi18 + (fi36 | (^(fi18)))) {
                                                                                                        case 1:
                                                                                                            if (0 >= 1000) {
                                                                                                                if (!(fi36 < (-(fi18)))) {
                                                                                                                    v69 := GenS9{F0: (-1 <= -508), F1: (v46 % ((2147483647) | 1)), F2: -1}
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = 4294967295
                                                                                                                    }
                                                                                                                    _ = v69
                                                                                                                }
                                                                                                                v46 += -2147483648
                                                                                                                var v70 uint16 = 1
                                                                                                                fmt.Println(v46, v60, fi18)
                                                                                                                _ = v70
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            v48 = 31
                                                                                                        case 32767:
                                                                                                            var v71 float64 = (-2.25 / 0.1)
                                                                                                            switch (^(fi18)) {
                                                                                                            case 32768:
                                                                                                                v72 := GenS10{F0: 9223372036854775807, F1: 430, F2: v1}
                                                                                                                _ = v72
                                                                                                            default:
                                                                                                                v48 -= v48
                                                                                                            }
                                                                                                            v3 = append(v3, (v15 + 4294967295))
                                                                                                            _ = v71
                                                                                                        default:
                                                                                                            v46 = v46
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v73 int64 = 9223372036854775807
                                                                                                        if true {
                                                                                                            v74 := GenS11{F0: (v73 * (v73 << ((634) & 63)))}
                                                                                                            _ = v74
                                                                                                        } else {
                                                                                                            fmt.Println(fi36)
                                                                                                            if false {
                                                                                                                v46 -= (-(v46))
                                                                                                                var v75 uint16 = 42
                                                                                                                if v1 {
                                                                                                                    fmt.Println(v46)
                                                                                                                    var v76 int64 = (v73 - 42)
                                                                                                                    _ = v76
                                                                                                                }
                                                                                                                _ = v75
                                                                                                            }
                                                                                                            v77 := GenS12{F0: ((("" + ("" + "hello")) + "") + "a")}
                                                                                                            switch ("hello" + "a") {
                                                                                                            case "hello":
                                                                                                                _ = len(v3)
                                                                                                            case "a":
                                                                                                                if (42 > 128) {
                                                                                                                    var v78 int16 = 1000
                                                                                                                    _ = v78
                                                                                                                } else {
                                                                                                                    v79 := GenS13{F0: v73, F1: v73}
                                                                                                                    _ = v79
                                                                                                                }
                                                                                                            default:
                                                                                                                v80 := GenS14{F0: v15, F1: "a"}
                                                                                                                _ = v80
                                                                                                            }
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            _ = v77
                                                                                                        }
                                                                                                        _ = v73
                                                                                                    }
                                                                                                case 128:
                                                                                                    v81 := GenS13{F0: -605, F1: 0}
                                                                                                    _ = v81
                                                                                                default:
                                                                                                    v82 := GenS15{F0: 2147483647, F1: ("hello" + (("foo bar" + "") + ("\n\t" + "a")))}
                                                                                                    _ = v82
                                                                                                }
                                                                                            case "":
                                                                                                v83 := GenS16{F0: (^(v48))}
                                                                                                _ = v83
                                                                                            default:
                                                                                                v17 = 1e10
                                                                                            }
                                                                                            delete(v2, v15)
                                                                                        case "\n\t":
                                                                                            _ = len(v2)
                                                                                        case "a":
                                                                                            var v84 int32 = v46
                                                                                            _ = v84
                                                                                        }
                                                                                    }
                                                                                    if v1 {
                                                                                        v85 := GenS15{F0: (fi18 << ((2) & 63)), F1: "hello"}
                                                                                        var v86 int32 = v46
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    }
                                                                                }
                                                                                v60 = float32(v17)
                                                                            }
                                                                        }
                                                                        _ = v60
                                                                    }
                                                                }
                                                                if v1 {
                                                                    var v87 uint32 = 65535
                                                                    _ = v87
                                                                } else {
                                                                    fmt.Println(v1, v17, v15)
                                                                }
                                                                v17 -= v17
                                                                v88 := GenS17{F0: 19}
                                                                _ = v88
                                                            }
                                                        }
                                                    }
                                                    _ = v48
                                                }
                                                fmt.Println(v46, fi36)
                                                var v89 uint16 = 42
                                                _ = v89
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v15
                                            }
                                            fmt.Println(v46)
                                            var v90 int64 = 1
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v91 float64 = float64(1.5)
                                            _ = v46
                                            _ = v90
                                            _ = v91
                                        }
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                    switch (fi18 + (^(fi18))) {
                                    case 127:
                                        v15 = (v15 & 2)
                                    case -538:
                                        switch (^(fi18)) {
                                        case 128:
                                            var v92 int64 = -9223372036854775808
                                            _ = v92
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                _ = len(v3)
                                            case "a":
                                                var v93 int8 = 1
                                                _ = v93
                                            case "":
                                                v4 -= v4
                                            default:
                                                v17 -= 1.5
                                            }
                                        }
                                    default:
                                        v15 -= 1
                                    }
                                }
                            }
                            switch "foo bar" {
                            case "":
                                for fi94 := 0; fi94 < 4; fi94++ {
                                    v1 = (0.1 != v17)
                                    var v95 int8 = 0
                                    _ = fi94
                                    _ = v95
                                }
                            case "\n\t":
                                delete(v5, v15)
                            default:
                                var v96 uint8 = 2
                                _ = v96
                            }
                        }
                        var v97 int64 = 101
                        var v98 uint = 2
                        _ = v97
                        _ = v98
                    }
                } else {
                    v1 = v1
                    var v99 bool = (1 != 1)
                    if (676 <= 989) {
                        v5[(v15 - (v15 * (v15 % (((^(v15))) | 1))))] = ("foo bar" + "")
                        v3 = append(v3, 786)
                    } else {
                        v1 = false
                        var v100 uint = 4294967295
                        _ = v100
                    }
                    _ = v99
                }
                v17 += (v4 - (v17 + v4))
                for fi101 := 0; fi101 < 9; fi101++ {
                    var v102 float64 = (1.5 + (v4 * 1e10))
                    var v103 float64 = float64(1.5)
                    _ = fi101
                    _ = v102
                    _ = v103
                }
            }
            _ = len(v5)
            if v1 {
                var v104 uint = 0
                v17 = 0.1
                _ = v104
            }
            _ = fi18
        }
    default:
        v105 := v6[9223372036854775807]
        _ = v105
    }
    if (42 < 32767) {
        v17 = v4
        v1 = true
        v106 := GenS11{F0: -9223372036854775808}
        _ = v106
    } else {
        v17 = v4
        v107 := GenS3{F0: v15, F1: 943, F2: 0}
        _ = v107
    }
    var v108 int16 = 32767
    v109 := GenS18{F0: (0 >= 0), F1: 830}
    v110 := GenS19{F0: 0, F1: (v15 << ((2) & 63))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (("" + "foo bar") + "a") {
    case "foo bar":
        var v111 int64 = 1
        _ = v111
    case "":
        v112 := GenS3{F0: (v15 >> ((42) & 63)), F1: 1, F2: 42}
        _ = v112
    }
    switch -175 {
    case -140:
        var v113 uint64 = 1
        _ = v113
    case 127:
        v108 = -309
    case 1:
        v15 = 521
    default:
        switch -9223372036854775808 {
        case 32768:
            var v114 int64 = 9223372036854775807
            _ = v114
        case 980:
            switch -387 {
            case 1:
                v4 = v4
            case -328:
                if v1 {
                    v115 := GenS14{F0: (v15 / ((65535) | 1)), F1: ""}
                    v116 := GenS17{F0: 100}
                    _ = v115
                    _ = v116
                } else {
                    var v117 uint16 = 0
                    _ = v117
                }
            case 128:
                var v118 uint8 = 128
                _ = v118
            }
        case 42:
            v15 = (v15 * (v15 << ((2) & 63)))
        }
    }
    for fi119 := 0; fi119 < 2; fi119++ {
        var v120 int64 = 1
        var v121 uint = 868
        _ = fi119
        _ = v120
        _ = v121
    }
    v1 = v1
    if len(v3) > 0 {
        v3[0] = (v15 | 806)
    }
    var v122 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v122:", v122)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v122
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
