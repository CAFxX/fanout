package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int32; F2 uint }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 int8; F2 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 int8; F2 int64 }
type GenS10 struct { F0 uint8; F1 int8 }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint32; F1 uint8 }
type GenS14 struct { F0 uint16; F1 int32 }
type GenS15 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS16 struct { F0 int64; F1 uint8 }
func genFn0(p0 bool, p1 uint64, p2 int8) string {
    v4 := []int{0, -1, -55, 804}
    var v5 uint16 = 868
    fmt.Println(v5, v5)
    v6 := GenS1{F0: -9223372036854775808}
    p0 = (!(false && p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1(p0 string) int32 {
    p0 = "hello"
    if true {
        var v7 string = "a"
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        p0 = ("foo bar" + p0)
        p0 = ""
        _ = fi8
    }
    var v9 uint8 = 42
    if (((p0 + p0) == p0) || ((!(42 <= 42)) && true)) {
        var v10 string = (p0 + p0)
        v11 := GenS2{F0: 1, F1: 4294967295, F2: 32767}
        for fi12 := 0; fi12 < 8; fi12++ {
            p0 = p0
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v9
    return 0
}

func main() {
    v1 := []string{"foo bar", "foo bar", "a"}
    v2 := []uint32{2147483648, 1}
    var v3 uint8 = 167
    switch "a" {
    case "\n\t":
        v2 = append(v2, uint32(v3))
    case "":
        if len(v1) > 0 {
            v1[0] = ("" + "\n\t")
        }
    default:
        var v13 uint = 0
        _ = v13
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += 255
    v3 -= (v3 + (v3 - (^(v3))))
    for fi14 := 0; fi14 < 19; fi14++ {
        _ = len(v1)
        for fi15 := 0; fi15 < 14; fi15++ {
            v3 = v3
            _ = fi15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            if (2 <= 42) {
                fmt.Println(v3, v3)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi16
        }
        _ = fi14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            switch "a" {
            case "\n\t":
                if ((fi17 <= (fi18 | (fi18 + fi18))) && true) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch (fi18 & fi17) {
                    case 9223372036854775807:
                        if (true && (698 != 1)) {
                            var v19 float64 = 0.1
                            var v20 uint16 = 452
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v19
                            _ = v20
                        } else {
                            v3 += 127
                            if (479 < 977) {
                                v3 = v3
                            } else {
                                v21 := GenS2{F0: 1, F1: 255, F2: 765}
                                _ = v21
                            }
                        }
                    case 2147483647:
                        var v22 float32 = 0.1
                        _ = v22
                    default:
                        var v23 int64 = 0
                        _ = v23
                    }
                    if false {
                        var v24 int64 = 412
                        var v25 uint16 = 936
                        _ = v24
                        _ = v25
                    }
                } else {
                    switch 9223372036854775807 {
                    case 0:
                        fmt.Println(fi18, v3)
                    case 750:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -201:
                        v3 -= (^(v3))
                    }
                }
            }
            if true {
                var v26 float32 = ((-2.25 / 1.5) + 3.14159)
                if false {
                    var v27 uint = 0
                    _ = v27
                }
                _ = v26
            }
            _ = fi18
        }
        var v28 uint = 2
        _ = fi17
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 float64 = float64(float32((((3.14159 * 1.5) - (3.14159 - 0.0)) - float64(((1.5 - -2.25) + 1e10)))))
        _ = fi29
        _ = v30
    }
    fmt.Println(v3)
    for fi31 := 0; fi31 < 3; fi31++ {
        v3 = 66
        var v32 float32 = 0.0
        switch "\n\t" {
        case "\n\t":
            if (!(474 == 172)) {
                v1 = append(v1, "a")
                fmt.Println(v32, v32, fi31)
            }
        case "foo bar":
            for fi33 := 0; fi33 < 9; fi33++ {
                v34 := GenS2{F0: 496, F1: 457, F2: 32767}
                v32 -= ((((3.14159 - v32) - 0.1) / v32) - v32)
                v3 = v3
                if (false && true) {
                    var v35 int32 = -649
                    if false {
                        var v36 int32 = (v35 / ((751) | 1))
                        _ = v36
                    } else {
                        switch "hello" {
                        case "a":
                            if (!(true || false)) {
                                v37 := GenS3{F0: fi33}
                                v35 -= v35
                                switch -9223372036854775808 {
                                case 128:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    var v38 float32 = (v32 / 1.5)
                                    _ = v38
                                }
                                _ = v37
                            }
                        default:
                            if (100 < -128) {
                                var v39 float32 = v32
                                if false {
                                    v3 = 1
                                    v40 := GenS2{F0: 522, F1: 42, F2: 2}
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v35, v39, v3)
                                    case "\n\t":
                                        v3 = 2
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        var v41 uint64 = 0
                                        _ = v41
                                    case "a":
                                        fmt.Println(v35, v32, fi31)
                                    case "\n\t":
                                        switch "" {
                                        case "foo bar":
                                            switch "a" {
                                            case "foo bar":
                                                v32 += (1.5 + float32((1e10 * (1.5 + -2.25))))
                                            }
                                        case "\n\t":
                                            var v42 float32 = ((v32 * (v39 + 0.1)) + (v32 - -2.25))
                                            _ = v42
                                        case "a":
                                            fmt.Println(v32)
                                        default:
                                            if (42 >= 32767) {
                                                v43 := GenS2{F0: 32767, F1: 2, F2: 255}
                                                _ = v43
                                            } else {
                                                v2 = append(v2, 1)
                                                v44 := GenS2{F0: 195, F1: 88, F2: 877}
                                                v32 -= 3.14159
                                                _ = v44
                                            }
                                        }
                                    default:
                                        v2 = append(v2, 154)
                                    }
                                    _ = v40
                                } else {
                                    switch fi31 {
                                    case 499:
                                        if false {
                                            if true {
                                                var v45 bool = (32767 > 0)
                                                _ = v45
                                            } else {
                                                var v46 int16 = 42
                                                v1 = append(v1, "a")
                                                _ = v46
                                            }
                                        } else {
                                            var v47 bool = (-1 >= 0)
                                            v48 := GenS4{F0: 39, F1: v35, F2: 689}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                    case -553:
                                        v39 = v32
                                    default:
                                        var v49 uint = 0
                                        _ = v49
                                    }
                                    var v50 int32 = 1
                                    _ = v50
                                }
                                _ = v39
                            }
                        }
                        var v51 int8 = 42
                        switch "foo bar" {
                        case "":
                            if true {
                                switch "" {
                                case "":
                                    var v52 float32 = 0.1
                                    _ = v52
                                case "\n\t":
                                    var v53 int32 = (v35 % (((v35 / (((v35 * 1)) | 1))) | 1))
                                    _ = v53
                                default:
                                    v35 += v35
                                }
                                var v54 uint32 = 77
                                _ = v54
                            } else {
                                fmt.Println(fi31, fi31)
                                var v55 bool = true
                                var v56 bool = v55
                                if (false || v56) {
                                    var v57 int16 = 42
                                    _ = v57
                                } else {
                                    if false {
                                        if v55 {
                                            v51 += (v51 + 42)
                                            v56 = (float32(0.0) <= v32)
                                        } else {
                                            var v58 float32 = 0.1
                                            v55 = (!v55)
                                            if len(v2) > 0 {
                                                v2[0] = 2
                                            }
                                            _ = v58
                                        }
                                        fmt.Println(v35)
                                    }
                                    if v55 {
                                        if ((v35 - v35) >= 894) {
                                            v51 += v51
                                            if v55 {
                                                v1 = append(v1, "\n\t")
                                            }
                                            if false {
                                                v32 = 3.14159
                                                var v59 float64 = ((-2.25 / 1e10) + 0.0)
                                                _ = v59
                                            } else {
                                                v60 := GenS1{F0: -1}
                                                fmt.Println(v32, v55)
                                                if (false || (65535 <= 0)) {
                                                    var v61 int8 = 127
                                                    switch (fi31 + (-(fi33))) {
                                                    case 32768:
                                                        var v62 float64 = 3.14159
                                                        v1 = append(v1, "\n\t")
                                                        var v63 int8 = (v51 & v51)
                                                        var v64 uint = 687
                                                        fmt.Println(v35)
                                                        if (160 > -2147483648) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            fmt.Println(fi31, fi31)
                                                        } else {
                                                            v65 := GenS4{F0: v64, F1: 951, F2: v64}
                                                            v64 -= v64
                                                            _ = v65
                                                        }
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    case 2147483647:
                                                        v35 = 1
                                                    }
                                                    var v66 int32 = genFn1("")
                                                    _ = v61
                                                    _ = v66
                                                } else {
                                                    var v67 uint16 = 627
                                                    var v68 float64 = (0.1 - 1.5)
                                                    _ = v67
                                                    _ = v68
                                                }
                                                _ = v60
                                            }
                                        }
                                        fmt.Println(fi31, v3, v55)
                                        v69 := GenS5{F0: 923, F1: (v3 ^ 128)}
                                        _ = v69
                                    }
                                    switch 32768 {
                                    case 127:
                                        v70 := GenS6{F0: (fi33 ^ -1)}
                                        _ = v70
                                    case -1:
                                        v32 += v32
                                    case -2147483648:
                                        var v71 uint16 = 32767
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v71, fi31)
                                        fmt.Println(v55)
                                        if v56 {
                                            v32 -= v32
                                            var v72 float32 = 1.5
                                            var v73 string = "hello"
                                            _ = v72
                                            _ = v73
                                        } else {
                                            fmt.Println(v71, v35)
                                            if v56 {
                                                v74 := GenS5{F0: 2147483647, F1: v3}
                                                _ = v74
                                            } else {
                                                var v75 int32 = (-(v35))
                                                if v55 {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if v55 {
                                                        if v56 {
                                                            if false {
                                                                v75 += v35
                                                                var v76 float32 = v32
                                                                var v77 int = -1
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        }
                                                        v32 = float32(1.5)
                                                        if true {
                                                            var v78 int32 = (^(v35))
                                                            v35 += v78
                                                            _ = len(v2)
                                                            _ = v78
                                                        }
                                                    } else {
                                                        v79 := GenS7{F0: v75, F1: 0, F2: 770}
                                                        _ = v79
                                                    }
                                                    var v80 int = 387
                                                    _ = v80
                                                }
                                                var v81 int = (-(fi33))
                                                _ = v75
                                                _ = v81
                                            }
                                        }
                                        if (!(18446744073709551615 == 0)) {
                                            v56 = (32767 <= 42)
                                            switch 2147483647 {
                                            case 831:
                                                fmt.Println(fi33, v71)
                                                switch "" {
                                                case "hello":
                                                    if (true || v56) {
                                                        if (fi33 != fi33) {
                                                            var v82 float32 = (((0.0 + v32) / (v32 * 3.14159)) / 1.0)
                                                            var v83 int16 = 0
                                                            var v84 float32 = v32
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                        } else {
                                                            v35 += int32(v51)
                                                            v32 = float32(0.0)
                                                            _ = len(v1)
                                                            if len(v2) > 0 {
                                                                v2[0] = 469
                                                            }
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = "a"
                                                        }
                                                        if v55 {
                                                            v56 = v55
                                                            var v85 string = (("foo bar" + "\n\t") + "hello")
                                                            v86 := GenS6{F0: fi33}
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        if v55 {
                                                            if len(v1) > 0 {
                                                                v1[0] = ("foo bar" + "hello")
                                                            }
                                                            v35 = (-(v35))
                                                            var v87 int32 = v35
                                                            _ = v87
                                                        }
                                                    } else {
                                                        fmt.Println(fi31, v32)
                                                        var v88 uint32 = 857
                                                        _ = v88
                                                    }
                                                case "":
                                                    if (!(v32 < v32)) {
                                                        if v56 {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v35 = (v35 | 926)
                                                            _ = len(v1)
                                                        }
                                                        var v89 float32 = (v32 * (float32(float64(v32)) / v32))
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if (42 >= 1) {
                                                            var v90 int32 = (^(v35))
                                                            v56 = true
                                                            if v55 {
                                                                v91 := GenS8{F0: 42}
                                                                _ = v91
                                                            }
                                                            v2 = append(v2, 42)
                                                            _ = v90
                                                        } else {
                                                            v92 := GenS8{F0: (fi33 | 32768)}
                                                            var v93 bool = ("\n\t" != "foo bar")
                                                            var v94 int8 = (v51 << ((459) & 63))
                                                            _ = v92
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                        if ((v3 ^ v3) <= v3) {
                                                            fmt.Println(v55, v51)
                                                        } else {
                                                            v89 = 0.0
                                                            var v95 int8 = 97
                                                            _ = v95
                                                        }
                                                        _ = v89
                                                    }
                                                case "a":
                                                    var v96 float32 = v32
                                                    _ = v96
                                                }
                                                var v97 bool = v56
                                                _ = v97
                                            case 42:
                                                if (42 < (v71 & (v71 / (((v71 * v71)) | 1)))) {
                                                    v2 = append(v2, 42)
                                                    v55 = (-1 >= (-(v51)))
                                                }
                                            default:
                                                v3 = v3
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                switch "hello" {
                                                case "":
                                                    if (v32 != v32) {
                                                        v51 -= -1
                                                    }
                                                case "foo bar":
                                                    if v55 {
                                                        if true {
                                                            switch "" {
                                                            case "":
                                                                v98 := GenS2{F0: (v71 >> ((255) & 63)), F1: 1, F2: (v71 / ((v71) | 1))}
                                                                _ = v98
                                                            }
                                                            var v99 float32 = (3.14159 * (v32 * v32))
                                                            _ = v99
                                                        } else {
                                                            if false {
                                                                v51 = 0
                                                                v3 -= v3
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v56 = true
                                                                case "hello":
                                                                    switch 331 {
                                                                    case 32767:
                                                                        v100 := GenS5{F0: 429, F1: (v3 >> ((42) & 63))}
                                                                        _ = v100
                                                                    case 9223372036854775807:
                                                                        v35 += genFn1("a")
                                                                    }
                                                                case "":
                                                                    var v101 uint32 = 2
                                                                    _ = v101
                                                                default:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 1
                                                                    }
                                                                }
                                                            }
                                                            v102 := GenS9{F0: 255, F1: -1, F2: int64(v51)}
                                                            v71 = (v71 >> ((943) & 63))
                                                            switch "foo bar" {
                                                            case "a":
                                                                v32 += (v32 + (v32 / (v32 - float32((-2.25 - (-2.25 / 3.14159))))))
                                                            default:
                                                                fmt.Println(v32, v3)
                                                            }
                                                            _ = v102
                                                        }
                                                        switch 32767 {
                                                        case 0:
                                                            if v55 {
                                                                v103 := GenS10{F0: v3, F1: v51}
                                                                _ = v103
                                                            } else {
                                                                var v104 bool = v56
                                                                v105 := GenS8{F0: -2147483648}
                                                                _ = v104
                                                                _ = v105
                                                            }
                                                        case -224:
                                                            if v56 {
                                                                v71 = 2
                                                            } else {
                                                                v51 = (v51 * v51)
                                                                v71 -= v71
                                                                if (true && (1 == v51)) {
                                                                    v56 = true
                                                                }
                                                            }
                                                            _ = len(v1)
                                                            v51 = (v51 ^ 127)
                                                        default:
                                                            var v106 uint32 = 0
                                                            _ = v106
                                                        }
                                                        fmt.Println(v71, fi31, fi31)
                                                    } else {
                                                        fmt.Println(v55)
                                                        var v107 int8 = -106
                                                        _ = v107
                                                    }
                                                }
                                            case "foo bar":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "hello":
                                                var v108 int8 = (v51 << ((2) & 63))
                                                _ = v108
                                            }
                                        } else {
                                            v1 = append(v1, "")
                                            switch 32767 {
                                            case -9223372036854775808:
                                                var v109 float64 = float64(v32)
                                                _ = v109
                                            }
                                            var v110 int8 = 1
                                            _ = v110
                                        }
                                        var v111 int64 = int64(v71)
                                        _ = v111
                                        _ = v71
                                    default:
                                        v112 := GenS3{F0: 42}
                                        _ = v112
                                    }
                                }
                                _ = v55
                                _ = v56
                            }
                        case "a":
                            var v113 uint16 = 42
                            _ = v113
                        case "foo bar":
                            var v114 float32 = v32
                            var v115 uint64 = 2
                            v35 -= v35
                            _ = v114
                            _ = v115
                        default:
                            fmt.Println(fi31, v32)
                        }
                        _ = v51
                    }
                    v32 += (-2.25 / 0.1)
                    _ = v35
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi33
                _ = v34
            }
        default:
            for fi116 := 0; fi116 < 7; fi116++ {
                v117 := GenS11{F0: 1, F1: 9223372036854775807}
                if true {
                    fmt.Println(fi116, v32)
                    v118 := GenS3{F0: fi31}
                    _ = v118
                } else {
                    var v119 int8 = 27
                    v119 -= v119
                    _ = v119
                }
                v32 = v32
                _ = fi116
                _ = v117
            }
        }
        _ = fi31
        _ = v32
    }
    v120 := GenS4{F0: 18446744073709551615, F1: genFn1("\n\t"), F2: 255}
    v3 = (v3 / ((v3) | 1))
    var v121 uint16 = 2
    var v122 uint = 144
    if len(v2) > 0 {
        v2[0] = 65535
    }
    if true {
        v2 = append(v2, 2147483648)
        for fi123 := 0; fi123 < 4; fi123++ {
            if true {
                if (32767 < 1000) {
                    v122 -= (v122 ^ 18446744073709551615)
                }
                var v124 int32 = -2147483648
                _ = v124
            } else {
                v121 = (v121 & (v121 | 855))
                if false {
                    v125 := GenS12{F0: -9223372036854775808}
                    _ = v125
                } else {
                    for fi126 := 0; fi126 < 1; fi126++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch ((("foo bar" + "a") + "foo bar") + ("hello" + "\n\t")) {
                        case "foo bar":
                            v127 := GenS6{F0: (fi126 - 42)}
                            _ = v127
                        case "a":
                            v121 += v121
                            if false {
                                var v128 bool = (((!false) && (!true)) && (v122 != (v122 ^ v122)))
                                _ = v128
                            }
                            v122 = v122
                        case "\n\t":
                            switch 127 {
                            case 988:
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                            case 32768:
                                v122 += v122
                            case 127:
                                fmt.Println(fi123)
                            }
                        default:
                            if false {
                                v1 = append(v1, ("foo bar" + "a"))
                            } else {
                                fmt.Println(v3)
                                var v129 int = 32767
                                fmt.Println(v3, v121)
                                _ = v129
                            }
                        }
                        _ = len(v2)
                        _ = fi126
                    }
                }
                if (9223372036854775807 > 582) {
                    v122 -= 0
                    switch (fi123 >> ((v122) & 63)) {
                    case 4294967295:
                        for fi130 := 0; fi130 < 17; fi130++ {
                            var v131 float64 = -2.25
                            v122 -= (^(v122))
                            if false {
                                v132 := GenS13{F0: 1, F1: (v3 - 55)}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v132
                            }
                            _ = fi130
                            _ = v131
                        }
                    case -2147483648:
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                switch 686 {
                                case 128:
                                    for fi133 := 0; fi133 < 14; fi133++ {
                                        var v134 uint32 = 2147483647
                                        if (0 <= 42) {
                                            if len(v2) > 0 {
                                                v2[0] = (v134 << (((v122 - v122)) & 63))
                                            }
                                        } else {
                                            switch fi133 {
                                            case 127:
                                                var v135 uint8 = 127
                                                _ = v135
                                            default:
                                                if false {
                                                    if false {
                                                        v122 = 65535
                                                        if (-2.25 == 3.14159) {
                                                            v1 = append(v1, ("foo bar" + ("foo bar" + "\n\t")))
                                                            fmt.Println(v122, v121)
                                                            _ = len(v2)
                                                        } else {
                                                            v122 += 4294967295
                                                            fmt.Println(v121, v121, v122)
                                                        }
                                                        if (1 <= v3) {
                                                            switch 2147483647 {
                                                            case 2147483647:
                                                                if false {
                                                                    v121 += 65535
                                                                    var v136 int = 127
                                                                    v122 = (^(v122))
                                                                    _ = len(v2)
                                                                    v137 := GenS14{F0: 42, F1: 621}
                                                                    _ = v136
                                                                    _ = v137
                                                                }
                                                                v138 := GenS9{F0: 18446744073709551615, F1: 1, F2: -1}
                                                                v139 := GenS1{F0: 42}
                                                                _ = v138
                                                                _ = v139
                                                            case 213:
                                                                v3 = v3
                                                            case -2147483648:
                                                                if false {
                                                                    var v140 uint64 = 540
                                                                    fmt.Println(v134, fi123)
                                                                    v141 := GenS7{F0: 1, F1: 1, F2: -1}
                                                                    _ = v140
                                                                    _ = v141
                                                                }
                                                            }
                                                            var v142 int = fi133
                                                            _ = v142
                                                        }
                                                    }
                                                } else {
                                                    if (-972 > 9223372036854775807) {
                                                        var v143 bool = (479 < -543)
                                                        _ = v143
                                                    }
                                                    v122 += v122
                                                }
                                            }
                                            var v144 uint16 = 255
                                            var v145 float32 = -2.25
                                            _ = v144
                                            _ = v145
                                        }
                                        var v146 int8 = -1
                                        _ = fi133
                                        _ = v134
                                        _ = v146
                                    }
                                }
                            }
                        }
                    case -9223372036854775808:
                        if true {
                            fmt.Println(v122, fi123, fi123)
                            v147 := GenS2{F0: v121, F1: v122, F2: v121}
                            _ = v147
                        } else {
                            for fi148 := 0; fi148 < 10; fi148++ {
                                v149 := GenS7{F0: -1, F1: -128, F2: 1}
                                switch ("a" + (("\n\t" + "a") + "foo bar")) {
                                case "\n\t":
                                    fmt.Println(fi148, v122, v122)
                                case "a":
                                    var v150 int64 = 461
                                    _ = v150
                                }
                                v151 := GenS15{F0: 42, F1: v3, F2: v121}
                                _ = fi148
                                _ = v149
                                _ = v151
                            }
                        }
                    default:
                        for fi152 := 0; fi152 < 18; fi152++ {
                            if true {
                                v121 += v121
                                v121 += v121
                            } else {
                                switch "" {
                                case "foo bar":
                                    v1 = append(v1, ("a" + (("" + (("hello" + "\n\t") + "a")) + "hello")))
                                case "a":
                                    v1 = append(v1, ("foo bar" + "foo bar"))
                                default:
                                    v121 += (v121 / ((v121) | 1))
                                }
                                _ = len(v1)
                            }
                            v3 = v3
                            v121 += (v121 / (((^(v121))) | 1))
                            _ = fi152
                        }
                    }
                }
            }
            switch (fi123 * (-(fi123))) {
            case 711:
                fmt.Println(fi123, v121, v122)
            case 0:
                for fi153 := 0; fi153 < 9; fi153++ {
                    v122 -= (v122 - v122)
                    _ = fi153
                }
            case 9223372036854775807:
                if (float32((float64((0.0 + ((1.5 * (3.14159 * 0.0)) * -2.25))) / 1.0)) < 1.5) {
                    var v154 int16 = int16(fi123)
                    v155 := GenS16{F0: -9223372036854775808, F1: v3}
                    var v156 bool = ((("foo bar" + (("foo bar" + ("a" + "\n\t")) + "a")) > ("hello" + ("foo bar" + "a"))) && ((v154 - v154) >= (v154 & v154)))
                    _ = v154
                    _ = v155
                    _ = v156
                }
            }
            _ = fi123
        }
        v157 := GenS16{F0: 9223372036854775807, F1: v3}
        _ = v157
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v120
    _ = v121
    _ = v122
    _ = v2
    _ = v3
}
