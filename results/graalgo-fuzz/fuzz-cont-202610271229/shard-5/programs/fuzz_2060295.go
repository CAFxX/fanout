package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 string; F1 string; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int64; F1 uint; F2 bool }
type GenS10 struct { F0 int16; F1 string }
type GenS11 struct { F0 int8; F1 string; F2 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int8; F1 int8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS17 struct { F0 uint; F1 uint32 }
type GenS18 struct { F0 bool; F1 int32; F2 uint8 }
type GenS19 struct { F0 uint; F1 int16; F2 int16 }
type GenS20 struct { F0 string; F1 uint32 }
type GenS21 struct { F0 uint8 }
func genFn0(p0 uint16) bool {
    switch "hello" {
    case "a":
        var v5 int64 = 300
        _ = v5
    case "\n\t":
        fmt.Println(p0)
    case "foo bar":
        v6 := []int{}
        _ = v6
    }
    p0 = (^(p0))
    p0 += 42
    v7 := []int{42}
    _ = p0
    _ = v7
    return false
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 1
    var v3 uint = 970
    var v4 int16 = 32767
    v8 := []int{0, 128, -519}
    switch ("foo bar" + "hello") {
    case "\n\t":
        if len(v8) > 0 {
            v8[0] = int(v3)
        }
    case "hello":
        switch -1 {
        case 0:
            fmt.Println(v3, v2, v1)
        case -535:
            var v9 string = ("\n\t" + "a")
            _ = v9
        case -1:
            fmt.Println(v3)
        }
    case "foo bar":
        if true {
            var v10 int = 128
            _ = v10
        } else {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int64 = -9223372036854775808
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                if (!(true || false)) {
                    var v14 uint16 = uint16(v2)
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 int16 = v4
                        if len(v8) > 0 {
                            v8[0] = 4294967295
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
                if false {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v8 = append(v8, 141)
                        switch fi13 {
                        case 4294967295:
                            v8 = append(v8, (fi17 * -808))
                        case 0:
                            v3 = v3
                        }
                        _ = fi17
                    }
                    fmt.Println(v4, v4)
                } else {
                    if len(v8) > 0 {
                        v8[0] = (-(fi13))
                    }
                    for fi18 := 0; fi18 < 3; fi18++ {
                        if true {
                            if (!((-2.25 * (float32(3.14159) * (1.5 / 1e10))) == -2.25)) {
                                var v19 int = 1
                                var v20 uint64 = 745
                                if len(v8) > 0 {
                                    v8[0] = 127
                                }
                                _ = v19
                                _ = v20
                            } else {
                                if true {
                                    if (v1 < (v1 | 1)) {
                                        var v21 float32 = 3.14159
                                        switch (fi13 % (((-(fi13))) | 1)) {
                                        case -9223372036854775808:
                                            if (65535 != 65535) {
                                                var v22 string = "a"
                                                _ = v22
                                            }
                                        case 32:
                                            fmt.Println(v3, v2)
                                            if (9223372036854775807 <= 594) {
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                            } else {
                                                var v23 uint8 = (v2 << (((v3 << ((v3) & 63))) & 63))
                                                _ = v23
                                            }
                                            v24 := GenS1{F0: 2}
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            var v25 int = fi13
                                            v21 -= float32(((1.5 / float64(float32(0.1))) / 1.0))
                                            var v26 int64 = 466
                                            if true {
                                                if true {
                                                    if ("" <= (("foo bar" + "hello") + "hello")) {
                                                        v27 := GenS1{F0: 244}
                                                        var v28 uint = (^(v3))
                                                        var v29 string = "foo bar"
                                                        switch (v29 + v29) {
                                                        case "\n\t":
                                                            v1 = (-(v1))
                                                        default:
                                                            fmt.Println(v3)
                                                        }
                                                        _ = v27
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        var v30 uint64 = 242
                                                        var v31 uint32 = 2147483648
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    v25 += 32767
                                                    fmt.Println(fi13, v25)
                                                } else {
                                                    var v32 float32 = 0.1
                                                    switch fi18 {
                                                    case 4294967295:
                                                        fmt.Println(v32, v2, v32)
                                                    case 748:
                                                        var v33 uint32 = 2147483647
                                                        _ = v33
                                                    }
                                                    var v34 bool = (false && false)
                                                    _ = v32
                                                    _ = v34
                                                }
                                            }
                                            _ = len(v8)
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        default:
                                            v35 := GenS1{F0: uint8(v3)}
                                            _ = v35
                                        }
                                        fmt.Println(v3, fi13, v21)
                                        fmt.Println(fi18)
                                        _ = v21
                                    }
                                    if (104 > -9223372036854775808) {
                                        v36 := GenS2{F0: -1, F1: 884}
                                        var v37 int64 = -898
                                        if false {
                                            if false {
                                                v38 := GenS1{F0: 128}
                                                v8 = append(v8, -2147483648)
                                                _ = v38
                                            }
                                        } else {
                                            var v39 int8 = 100
                                            if true {
                                                switch "hello" {
                                                case "hello":
                                                    fmt.Println(v2)
                                                case "a":
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                case "":
                                                    if (v4 >= int16(v2)) {
                                                        var v40 int32 = v1
                                                        var v41 int = (fi18 * fi18)
                                                        _ = v40
                                                        _ = v41
                                                    } else {
                                                        if false {
                                                            var v42 uint = 824
                                                            v43 := GenS2{F0: (v37 * v37), F1: v37}
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            v44 := GenS2{F0: (v37 % ((v37) | 1)), F1: v37}
                                                            var v45 uint32 = 2
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        if ((v4 + (v4 >> ((v3) & 63))) != (v4 & 32767)) {
                                                            fmt.Println(v1, v4)
                                                        }
                                                    }
                                                default:
                                                    var v46 uint = (^(v3))
                                                    _ = v46
                                                }
                                                v39 += v39
                                                switch 898 {
                                                case 9223372036854775807:
                                                    if false {
                                                        var v47 uint64 = 0
                                                        v37 += 1
                                                        if (true || false) {
                                                            v48 := GenS3{F0: "hello", F1: "hello", F2: 42}
                                                            _ = v48
                                                        }
                                                        _ = v47
                                                    } else {
                                                        if false {
                                                            v49 := GenS4{F0: 4294967295}
                                                            v50 := GenS5{F0: "", F1: false, F2: 42}
                                                            switch fi13 {
                                                            case -27:
                                                                if len(v8) > 0 {
                                                                    v8[0] = (fi13 % ((-809) | 1))
                                                                }
                                                            case -220:
                                                                switch -9223372036854775808 {
                                                                case 128:
                                                                    var v51 string = ""
                                                                    _ = v51
                                                                case 596:
                                                                    if false {
                                                                        if ((^(v3)) < v3) {
                                                                            var v52 uint16 = 0
                                                                            if (((0.0 - (0.0 - 3.14159)) * 0.1) >= 3.14159) {
                                                                                v52 -= v52
                                                                                v8 = append(v8, (-(fi13)))
                                                                            }
                                                                            v53 := GenS3{F0: "", F1: ("a" + ("foo bar" + "")), F2: 88}
                                                                            var v54 float64 = 1.5
                                                                            _ = v52
                                                                            _ = v53
                                                                            _ = v54
                                                                        }
                                                                        var v55 string = "\n\t"
                                                                        v37 = (v37 << ((1) & 63))
                                                                        _ = v55
                                                                    } else {
                                                                        if false {
                                                                            v1 += 1
                                                                            fmt.Println(fi18, v1, v3)
                                                                            var v56 int16 = v4
                                                                            _ = v56
                                                                        }
                                                                        var v57 int64 = v37
                                                                        _ = v57
                                                                    }
                                                                }
                                                            case -9223372036854775808:
                                                                if false {
                                                                    if true {
                                                                        switch (fi13 | 127) {
                                                                        case 127:
                                                                            var v58 int64 = -9223372036854775808
                                                                            _ = v58
                                                                        default:
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                if len(v8) >= 2 {
                                                                                    v8 = v8[1:len(v8)]
                                                                                }
                                                                            case "hello":
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = fi18
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if len(v8) > 0 {
                                                                            v8[0] = fi18
                                                                        }
                                                                        if len(v8) > 0 {
                                                                            v8[0] = 196
                                                                        }
                                                                    }
                                                                    v3 = 255
                                                                } else {
                                                                    v8 = append(v8, fi13)
                                                                }
                                                            default:
                                                                fmt.Println(v2, v2, v2)
                                                            }
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        v8 = append(v8, int(v4))
                                                    }
                                                }
                                                fmt.Println(v4, v1, fi18)
                                                v59 := GenS4{F0: 757}
                                                switch "hello" {
                                                case "hello":
                                                    v60 := GenS3{F0: ((("" + "foo bar") + ("\n\t" + "hello")) + "hello"), F1: (("hello" + "foo bar") + ("a" + "hello")), F2: 2}
                                                    _ = v60
                                                case "a":
                                                    fmt.Println(fi13, v37)
                                                case "":
                                                    v1 = (v1 ^ v1)
                                                }
                                                v61 := GenS2{F0: -236, F1: v37}
                                                var v62 uint16 = 827
                                                var v63 float32 = (1.5 - float32(1.5))
                                                _ = v59
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            } else {
                                                fmt.Println(v1, v4, v39)
                                            }
                                            switch fi18 {
                                            case 9223372036854775807:
                                                v2 = v2
                                            case 4294967295:
                                                switch "" {
                                                case "\n\t":
                                                    switch (fi13 << ((v3) & 63)) {
                                                    case 0:
                                                        var v64 int64 = 746
                                                        _ = v64
                                                    case -2147483648:
                                                        v2 = v2
                                                    case -631:
                                                        if ("hello" <= ("foo bar" + "hello")) {
                                                            _ = len(v8)
                                                            switch "foo bar" {
                                                            case "":
                                                                if true {
                                                                    if (3.14159 <= -2.25) {
                                                                        v8 = append(v8, fi13)
                                                                        v65 := GenS6{F0: 2, F1: v3, F2: 127}
                                                                        _ = v65
                                                                    }
                                                                    switch "hello" {
                                                                    case "a":
                                                                        var v66 string = "a"
                                                                        _ = v66
                                                                    case "":
                                                                        fmt.Println(fi18)
                                                                    default:
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            var v67 uint32 = 896
                                                                            _ = v67
                                                                        case "hello":
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                var v68 bool = false
                                                                                _ = v68
                                                                            case "":
                                                                                v69 := GenS5{F0: "a", F1: (65535 > 615), F2: 2147483647}
                                                                                _ = v69
                                                                            default:
                                                                                fmt.Println(v3)
                                                                            }
                                                                        case "\n\t":
                                                                            v70 := GenS7{F0: 18446744073709551615, F1: 80}
                                                                            var v71 int = -873
                                                                            var v72 int64 = (^(v37))
                                                                            _ = v70
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "hello":
                                                                        v8 = append(v8, 9223372036854775807)
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                    }
                                                                }
                                                            case "hello":
                                                                fmt.Println(v1)
                                                            default:
                                                                if (3.14159 < 1.5) {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if (v1 != v1) {
                                                                            v37 -= v37
                                                                            if (0.0 < 0.0) {
                                                                                var v73 uint16 = 840
                                                                                _ = v73
                                                                            } else {
                                                                                v8 = append(v8, (fi13 & -385))
                                                                            }
                                                                        } else {
                                                                            var v74 float64 = 0.0
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                            _ = v74
                                                                        }
                                                                    case "":
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                    default:
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        v4 = -40
                                                                        v3 += (^(v3))
                                                                    } else {
                                                                        v75 := GenS8{F0: -524, F1: ((^(fi13)) >= fi18)}
                                                                        if (true || (0.1 != float64(float32(1e10)))) {
                                                                            if false {
                                                                                fmt.Println(v1, v3, v39)
                                                                            }
                                                                            var v76 float64 = 3.14159
                                                                            _ = v76
                                                                        }
                                                                        _ = v75
                                                                    }
                                                                    v77 := GenS9{F0: int64(v1), F1: v3, F2: (365 == 892)}
                                                                    _ = v77
                                                                } else {
                                                                    var v78 uint32 = 0
                                                                    _ = v78
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v79 uint = (v3 * v3)
                                                    if (!(int8(fi13) < 42)) {
                                                        if false {
                                                            var v80 bool = (!((!(!true)) || true))
                                                            _ = v80
                                                        } else {
                                                            var v81 string = "\n\t"
                                                            if (v39 == (v39 << ((330) & 63))) {
                                                                if ((((0.1 - 1.5) + ((0.0 + (float64((((3.14159 + float32(((0.0 - ((3.14159 / (-2.25 + -2.25)) + 0.1)) + (1e10 - 0.0)))) * 1.5) - 0.1)) + float64(0.0))) - 1.5)) / 1e10) <= (3.14159 * -2.25)) {
                                                                    if (fi18 == fi18) {
                                                                        fmt.Println(fi13)
                                                                    }
                                                                    v2 -= v2
                                                                    _ = len(v8)
                                                                } else {
                                                                    v82 := GenS3{F0: v81, F1: (v81 + "a"), F2: 2}
                                                                    _ = v82
                                                                }
                                                                v81 = "foo bar"
                                                                v39 += 100
                                                            } else {
                                                                v83 := GenS3{F0: v81, F1: (v81 + "foo bar"), F2: 329}
                                                                _ = v83
                                                            }
                                                            if (849 != 13) {
                                                                var v84 uint = v79
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                _ = v84
                                                            } else {
                                                                if (!(v79 > v79)) {
                                                                    v8 = append(v8, fi18)
                                                                    fmt.Println(v39, v2)
                                                                } else {
                                                                    v79 -= (v3 | (v3 / (((v79 >> (((v79 - v79)) & 63))) | 1)))
                                                                    v85 := GenS8{F0: fi13, F1: true}
                                                                    _ = v85
                                                                }
                                                                v39 -= (v39 * v39)
                                                                var v86 uint8 = v2
                                                                _ = v86
                                                            }
                                                            _ = v81
                                                        }
                                                        v87 := GenS8{F0: fi18, F1: true}
                                                        v4 += (v4 << ((18446744073709551615) & 63))
                                                        _ = v87
                                                    } else {
                                                        var v88 uint8 = (v2 >> ((v3) & 63))
                                                        switch 32768 {
                                                        case 1:
                                                            v89 := GenS2{F0: (-(v37)), F1: (^(v37))}
                                                            _ = v89
                                                        }
                                                        if (!(42 > 481)) {
                                                            v90 := GenS10{F0: v4, F1: ((("a" + "foo bar") + "\n\t") + ("foo bar" + "hello"))}
                                                            if false {
                                                                if len(v8) > 0 {
                                                                    v8[0] = fi18
                                                                }
                                                                fmt.Println(fi13, v37)
                                                                var v91 uint64 = 0
                                                                var v92 int16 = (v4 & (v4 % (((v4 + (v4 & (v4 * -1)))) | 1)))
                                                                _ = v91
                                                                _ = v92
                                                            } else {
                                                                var v93 int16 = (v4 + v4)
                                                                fmt.Println(v3, v2)
                                                                _ = v93
                                                            }
                                                            switch fi18 {
                                                            case 582:
                                                                v39 -= (v39 & (v39 + 0))
                                                                v88 -= v2
                                                                var v94 string = ("a" + "\n\t")
                                                                v8 = append(v8, 173)
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                _ = v94
                                                            case 32768:
                                                                if ((1.5 / 1e10) > (1.5 - float32(0.0))) {
                                                                    fmt.Println(v4, v39)
                                                                }
                                                            case -9223372036854775808:
                                                                var v95 int16 = v4
                                                                v96 := GenS1{F0: v88}
                                                                if len(v8) > 0 {
                                                                    v8[0] = fi13
                                                                }
                                                                v97 := GenS7{F0: 0, F1: 49}
                                                                v98 := GenS8{F0: int(v39), F1: (937 != 255)}
                                                                v99 := GenS1{F0: 0}
                                                                var v100 bool = (int8(fi18) <= (v39 >> ((711) & 63)))
                                                                if (100 > -128) {
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    v95 -= v4
                                                                    var v101 uint = v3
                                                                    _ = v101
                                                                } else {
                                                                    var v102 int16 = (v4 - (v95 | (v4 | v4)))
                                                                    v103 := GenS11{F0: 42, F1: ("foo bar" + ("" + "foo bar")), F2: 0}
                                                                    var v104 uint64 = 1
                                                                    _ = v102
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                                _ = v100
                                                                _ = v95
                                                                _ = v96
                                                                _ = v97
                                                                _ = v98
                                                                _ = v99
                                                            default:
                                                                var v105 uint = (^(v79))
                                                                _ = v105
                                                            }
                                                            _ = v90
                                                        }
                                                        _ = v88
                                                    }
                                                    _ = v79
                                                case "":
                                                    v39 -= (v39 & (v39 % ((v39) | 1)))
                                                default:
                                                    var v106 uint8 = 0
                                                    _ = v106
                                                }
                                            }
                                            _ = v39
                                        }
                                        _ = v36
                                        _ = v37
                                    }
                                    switch 4294967295 {
                                    case -2147483648:
                                        v8 = append(v8, fi13)
                                    }
                                }
                                if len(v8) > 0 {
                                    v8[0] = 861
                                }
                                var v107 int = fi18
                                var v108 uint8 = (v2 * 1)
                                _ = v107
                                _ = v108
                            }
                        }
                        switch 9223372036854775807 {
                        case 0:
                            v2 += 42
                        }
                        _ = fi18
                    }
                }
                _ = fi13
            }
        }
    }
    var v109 int16 = 1
    v109 = (v109 - (v109 + v4))
    if (390 != -9223372036854775808) {
        fmt.Println(v1, v1, v1)
        var v110 int32 = (v1 / ((v1) | 1))
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v110
    } else {
        var v111 uint = (^(v3))
        fmt.Println(v4)
        _ = v111
    }
    if true {
        v4 = 32767
    }
    v112 := GenS12{F0: 42}
    fmt.Println(v4, v2, v1)
    var v113 uint64 = 42
    var v114 int8 = 1
    if ("a" < "a") {
        _ = len(v8)
        var v115 int8 = (v114 ^ (v114 * 1))
        _ = v115
    } else {
        switch -2147483648 {
        case 127:
            v116 := GenS12{F0: -1}
            _ = v116
        case -9223372036854775808:
            fmt.Println(v4)
        case 1:
            if ((2147483647 <= 42) || (240 < 864)) {
                fmt.Println(v109, v113)
                for fi117 := 0; fi117 < 14; fi117++ {
                    var v118 uint8 = 127
                    v113 = v113
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi117
                    _ = v118
                }
            } else {
                switch "foo bar" {
                case "a":
                    for fi119 := 0; fi119 < 20; fi119++ {
                        var v120 int8 = v114
                        v4 = (v4 / (((-(v4))) | 1))
                        var v121 bool = true
                        _ = fi119
                        _ = v120
                        _ = v121
                    }
                case "foo bar":
                    if len(v8) > 0 {
                        v8[0] = 803
                    }
                default:
                    for fi122 := 0; fi122 < 7; fi122++ {
                        v113 = 274
                        v123 := GenS13{F0: v114, F1: 42}
                        if (!true) {
                            if false {
                                switch (fi122 % (((-(fi122))) | 1)) {
                                case -278:
                                    v3 -= uint(v2)
                                case 2147483647:
                                    if len(v8) > 0 {
                                        v8[0] = (fi122 ^ (fi122 << (((v3 >> ((v3) & 63))) & 63)))
                                    }
                                }
                            } else {
                                if len(v8) > 0 {
                                    v8[0] = (fi122 + (fi122 - fi122))
                                }
                            }
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        } else {
                            for fi124 := 0; fi124 < 2; fi124++ {
                                v109 += (v4 | 753)
                                _ = fi124
                            }
                            v2 -= uint8(v3)
                        }
                        _ = fi122
                        _ = v123
                    }
                }
                switch int(v2) {
                case 808:
                    fmt.Println(v2, v109, v2)
                case -635:
                    switch 128 {
                    case 2147483647:
                        v125 := GenS14{F0: ((^(v109)) >= 0)}
                        _ = v125
                    case 128:
                        _ = len(v8)
                    }
                default:
                    v3 = 1
                }
                for fi126 := 0; fi126 < 15; fi126++ {
                    fmt.Println(v4, v114)
                    if false {
                        var v127 bool = (472 == uint32(v113))
                        _ = v127
                    }
                    _ = fi126
                }
            }
        default:
            if true {
                for fi128 := 0; fi128 < 17; fi128++ {
                    if ((v4 / (((v109 / (((v109 / ((v109) | 1))) | 1))) | 1)) != v109) {
                        switch "\n\t" {
                        case "foo bar":
                            if (2 < 446) {
                                v129 := GenS15{F0: 2147483647}
                                _ = v129
                            }
                        case "a":
                            var v130 uint = v3
                            _ = v130
                        case "\n\t":
                            var v131 float64 = 0.0
                            _ = v131
                        default:
                            for fi132 := 0; fi132 < 8; fi132++ {
                                var v133 bool = true
                                switch "" {
                                case "":
                                    var v134 float64 = (1.5 - (0.1 * (0.0 + (0.1 * 3.14159))))
                                    _ = v134
                                case "\n\t":
                                    v113 -= v113
                                    var v135 int8 = (v114 | (v114 / ((1) | 1)))
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    switch "\n\t" {
                                    case "":
                                        v136 := GenS15{F0: v1}
                                        _ = v136
                                    case "\n\t":
                                        v8 = append(v8, (^(fi128)))
                                    case "a":
                                        var v137 float64 = 1.5
                                        _ = v137
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    case "hello":
                                        v138 := GenS9{F0: -1, F1: 0, F2: (42 < 1)}
                                        _ = v138
                                    }
                                    _ = v135
                                case "foo bar":
                                    if (2 == 1) {
                                        v114 = v114
                                        var v139 int = (fi128 / ((fi132) | 1))
                                        v140 := GenS15{F0: v1}
                                        if len(v8) > 0 {
                                            v8[0] = 666
                                        }
                                        _ = v139
                                        _ = v140
                                    }
                                default:
                                    v141 := GenS4{F0: (^(v113))}
                                    _ = v141
                                }
                                _ = len(v8)
                                _ = fi132
                                _ = v133
                            }
                        }
                    }
                    var v142 int64 = -9223372036854775808
                    _ = fi128
                    _ = v142
                }
                v109 -= (v4 >> (((^(v3))) & 63))
            }
        }
    }
    for fi143 := 0; fi143 < 14; fi143++ {
        v144 := GenS16{F0: v113, F1: 974, F2: (v1 | (-(v1)))}
        _ = fi143
        _ = v144
    }
    v109 += 0
    if false {
        fmt.Println(v3)
    } else {
        v8 = append(v8, -945)
        if (!(v4 >= v109)) {
            var v145 bool = (0.0 <= 1e10)
            v146 := GenS16{F0: v113, F1: 4294967295, F2: v1}
            if v145 {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            } else {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v2, v145, v4)
                case "foo bar":
                    switch (("\n\t" + "") + "") {
                    case "a":
                        var v147 uint64 = v113
                        _ = v147
                    }
                default:
                    fmt.Println(v113)
                }
                if v145 {
                    fmt.Println(v2, v113)
                }
                fmt.Println(v3)
            }
            _ = v145
            _ = v146
        } else {
            var v148 uint32 = 2
            fmt.Println(v113, v1)
            _ = v148
        }
        _ = len(v8)
    }
    if true {
        v149 := GenS17{F0: 65535, F1: 2147483648}
        for fi150 := 0; fi150 < 2; fi150++ {
            var v151 uint32 = 0
            _ = len(v8)
            for fi152 := 0; fi152 < 20; fi152++ {
                if (v113 <= (v113 * (v113 >> ((uint(v151)) & 63)))) {
                    v151 -= 42
                    switch (("a" + ("" + "hello")) + ("\n\t" + ("hello" + ("" + "hello")))) {
                    case "\n\t":
                        var v153 uint = (v3 * v3)
                        _ = v153
                    case "a":
                        var v154 float32 = 0.0
                        _ = v154
                    }
                    if (false && ((v113 - v113) != (v113 | (v113 / ((18446744073709551615) | 1))))) {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        var v155 float64 = 0.0
                        v156 := GenS10{F0: v109, F1: "a"}
                        _ = v155
                        _ = v156
                    } else {
                        var v157 uint32 = v151
                        _ = v157
                    }
                } else {
                    switch 128 {
                    case 451:
                        v109 = int16(v3)
                    default:
                        if (true || ((int8(v1) < (v114 - -92)) || false)) {
                            var v158 uint8 = v2
                            var v159 float64 = float64(0.1)
                            _ = v158
                            _ = v159
                        }
                    }
                }
                v160 := GenS18{F0: true, F1: v1, F2: (v2 % ((v2) | 1))}
                _ = fi152
                _ = v160
            }
            _ = fi150
            _ = v151
        }
        _ = v149
    }
    if true {
        var v161 int = -9223372036854775808
        switch (v161 * v161) {
        case -25:
            if false {
                var v162 uint8 = (v2 >> ((v3) & 63))
                var v163 int16 = 898
                var v164 int = 1
                _ = v162
                _ = v163
                _ = v164
            } else {
                if (false && true) {
                    for fi165 := 0; fi165 < 11; fi165++ {
                        v8 = append(v8, v161)
                        fmt.Println(v1, v161)
                        _ = fi165
                    }
                    var v166 int64 = 519
                    v167 := GenS9{F0: (v166 | 9223372036854775807), F1: 1, F2: (2147483647 >= uint32(v113))}
                    _ = v166
                    _ = v167
                } else {
                    if false {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        if (1 != 4294967295) {
                            var v168 uint32 = 2
                            _ = v168
                        } else {
                            v1 = (v1 * -1)
                        }
                        if false {
                            for fi169 := 0; fi169 < 10; fi169++ {
                                v2 += v2
                                _ = fi169
                            }
                            if (634 == (v3 ^ (v3 * (v3 >> ((v3) & 63))))) {
                                for fi170 := 0; fi170 < 10; fi170++ {
                                    switch -795 {
                                    case 32768:
                                        var v171 int32 = v1
                                        _ = v171
                                    case 634:
                                        v8 = append(v8, (-(fi170)))
                                    case -2147483648:
                                        fmt.Println(v109, v114, v3)
                                    default:
                                        if false {
                                            switch (fi170 % (((v161 / ((-963) | 1))) | 1)) {
                                            case 0:
                                                var v172 int64 = 0
                                                _ = v172
                                            default:
                                                var v173 float64 = 1.5
                                                _ = v173
                                            }
                                            fmt.Println(v114, v4, fi170)
                                            var v174 uint64 = v113
                                            _ = v174
                                        }
                                    }
                                    _ = fi170
                                }
                                _ = len(v8)
                                fmt.Println(v4)
                            } else {
                                v113 = v113
                            }
                            for fi175 := 0; fi175 < 7; fi175++ {
                                for fi176 := 0; fi176 < 18; fi176++ {
                                    var v177 float32 = (1.5 / 0.1)
                                    _ = fi176
                                    _ = v177
                                }
                                v178 := GenS4{F0: (v113 >> (((v3 >> ((0) & 63))) & 63))}
                                for fi179 := 0; fi179 < 20; fi179++ {
                                    v180 := GenS14{F0: true}
                                    if (!(false || false)) {
                                        if false {
                                            var v181 int8 = (v114 << ((v3) & 63))
                                            fmt.Println(v161)
                                            _ = v181
                                        } else {
                                            v8 = append(v8, 4294967295)
                                            var v182 bool = ((621 == 2) && true)
                                            _ = v182
                                        }
                                        v161 = (fi179 / ((fi175) | 1))
                                        if ((937 <= 255) || false) {
                                            v1 = -2147483648
                                            var v183 uint = (^(v3))
                                            _ = v183
                                        }
                                        v114 = int8(v2)
                                        fmt.Println(v114, v1)
                                        v109 = (v4 << ((v3) & 63))
                                        switch 32768 {
                                        case 398:
                                            if (842 <= 2) {
                                                v184 := GenS15{F0: v1}
                                                if (false || (!false)) {
                                                    if false {
                                                        v1 = (v1 << ((v3) & 63))
                                                        v185 := GenS4{F0: v113}
                                                        v186 := GenS19{F0: (v3 << ((255) & 63)), F1: 32767, F2: (v4 >> ((v3) & 63))}
                                                        _ = v185
                                                        _ = v186
                                                    }
                                                    v114 = -31
                                                    if (-1 >= -9223372036854775808) {
                                                        fmt.Println(v3, fi179, v4)
                                                        v187 := GenS19{F0: v3, F1: v109, F2: int16(v2)}
                                                        _ = v187
                                                    } else {
                                                        var v188 uint64 = 609
                                                        _ = v188
                                                    }
                                                } else {
                                                    v109 -= v109
                                                    _ = len(v8)
                                                    v114 -= 42
                                                }
                                                var v189 bool = true
                                                v161 += 1
                                                _ = v184
                                                _ = v189
                                            }
                                        }
                                        var v190 string = (("hello" + ("" + "a")) + "foo bar")
                                        _ = v190
                                    } else {
                                        var v191 int64 = int64(v1)
                                        _ = v191
                                    }
                                    _ = fi179
                                    _ = v180
                                }
                                _ = fi175
                                _ = v178
                            }
                        }
                    }
                    v161 -= int(v1)
                    var v192 int = (v161 | v161)
                    _ = v192
                }
                if ((v2 <= (v2 / ((127) | 1))) || false) {
                    v161 -= (v161 + v161)
                    fmt.Println(v113, v161)
                    _ = len(v8)
                } else {
                    fmt.Println(v4, v113)
                }
                var v193 int8 = v114
                _ = v193
            }
        case -2147483648:
            if (9223372036854775807 < 1) {
                if (18446744073709551615 >= 0) {
                    for fi194 := 0; fi194 < 6; fi194++ {
                        switch ("hello" + "") {
                        case "a":
                            for fi195 := 0; fi195 < 8; fi195++ {
                                var v196 uint8 = v2
                                fmt.Println(v109)
                                var v197 float64 = ((1.5 - 0.1) + (1.5 + 0.1))
                                _ = fi195
                                _ = v196
                                _ = v197
                            }
                        case "":
                            var v198 float32 = float32(((-2.25 * 3.14159) + 0.0))
                            _ = v198
                        case "foo bar":
                            var v199 uint32 = 2147483648
                            _ = v199
                        }
                        _ = fi194
                    }
                } else {
                    var v200 uint64 = (v113 * 4294967295)
                    v201 := GenS20{F0: ("foo bar" + ("" + ("\n\t" + ("" + "hello")))), F1: 1}
                    if ((!(!(!true))) && (v200 >= 620)) {
                        v114 += 0
                        v202 := GenS21{F0: v2}
                        _ = v202
                    }
                    _ = v200
                    _ = v201
                }
                v3 += v3
            }
        default:
            var v203 uint16 = 2
            _ = v203
        }
        _ = v161
    }
    v3 = uint(v113)
    var v204 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v204:", v204)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v109
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v204
    _ = v3
    _ = v4
    _ = v8
}
