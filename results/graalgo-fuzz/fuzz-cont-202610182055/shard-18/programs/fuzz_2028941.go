package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint8; F1 int; F2 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{996, 622}
    v3 := []uint32{509, 629, 2, 2147483648}
    v4 := GenS1{F0: -906}
    var v5 uint8 = 138
    v6 := GenS2{F0: 18446744073709551615, F1: 913}
    v5 += v5
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint = 456
        var v9 bool = false
        v5 = 128
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v5)
    var v10 uint = 2
    var v11 int64 = 42
    v12 := GenS2{F0: 4294967295, F1: -1}
    for fi13 := 0; fi13 < 4; fi13++ {
        v5 += v5
        v14 := GenS1{F0: -165}
        for fi15 := 0; fi15 < 5; fi15++ {
            v10 -= (v10 * (v10 | (v10 << ((v10) & 63))))
            var v16 int64 = 42
            switch "foo bar" {
            case "\n\t":
                if ((-(fi15)) <= 9223372036854775807) {
                    v17 := GenS1{F0: 854}
                    var v18 int64 = (v11 | v16)
                    _ = v17
                    _ = v18
                }
            case "hello":
                v19 := GenS1{F0: 1}
                var v20 string = "foo bar"
                _ = v19
                _ = v20
            case "":
                var v21 float64 = 3.14159
                fmt.Println(fi15, v5)
                switch (^(fi13)) {
                case 2147483647:
                    var v22 uint16 = 464
                    _ = v22
                case 32768:
                    var v23 uint64 = 124
                    _ = v23
                case 32767:
                    if false {
                        v24 := GenS3{F0: 42}
                        v16 += 42
                        v25 := v1[fi15]
                        _ = v24
                        _ = v25
                    }
                }
                if false {
                    if false {
                        var v26 float32 = 3.14159
                        if (((false && false) && (((2147483647 > 207) || false) && ((false && ((fi15 <= fi13) && (fi13 <= fi13))) || (!false)))) && false) {
                            v1[fi15] = 493
                            var v27 bool = ("hello" > "")
                            if v27 {
                                fmt.Println(v26)
                                if ((2 != 943) && (1000 != -1)) {
                                    switch (fi15 - -1) {
                                    case 0:
                                        if v27 {
                                            switch (("a" + "a") + "\n\t") {
                                            case "":
                                                if false {
                                                    fmt.Println(v21, v16, fi15)
                                                    v5 = v5
                                                    var v28 uint8 = v5
                                                    _ = v28
                                                } else {
                                                    if (((v5 << (((v10 & 0)) & 63)) == v5) && true) {
                                                        if false {
                                                            switch ("foo bar" + "foo bar") {
                                                            case "foo bar":
                                                                switch 2147483647 {
                                                                case 128:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case 32767:
                                                                    fmt.Println(v5, fi13, fi15)
                                                                case 42:
                                                                    v26 = (v26 + v26)
                                                                }
                                                            case "hello":
                                                                if false {
                                                                    v1[int(v16)] = (fi13 & -192)
                                                                    var v29 float64 = v21
                                                                    if v27 {
                                                                        var v30 uint32 = 4294967295
                                                                        if (-1 <= 0) {
                                                                            if (v11 <= v11) {
                                                                                if v27 {
                                                                                    v16 = (-(v16))
                                                                                }
                                                                            } else {
                                                                                v29 += (v21 / v21)
                                                                            }
                                                                            var v31 uint32 = v30
                                                                            _ = v31
                                                                        } else {
                                                                            if v27 {
                                                                                var v32 int16 = -32768
                                                                                v33 := GenS2{F0: 2, F1: (v32 << ((v10) & 63))}
                                                                                v27 = (1 <= -1)
                                                                                var v34 int64 = 576
                                                                                if ((-1 >= v32) && false) {
                                                                                    var v35 float64 = ((v21 + v29) + (float64(v26) / (v21 - v29)))
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = v30
                                                                                    }
                                                                                    _ = v35
                                                                                }
                                                                                _ = v32
                                                                                _ = v33
                                                                                _ = v34
                                                                            }
                                                                            v36 := GenS2{F0: 702, F1: 506}
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (v30 ^ v30)
                                                                            }
                                                                            _ = v36
                                                                        }
                                                                        _ = v30
                                                                    } else {
                                                                        v27 = (v27 || v27)
                                                                        v26 = -2.25
                                                                        v2 = append(v2, 0)
                                                                    }
                                                                    fmt.Println(v16, v11)
                                                                    v10 -= 42
                                                                    _ = v29
                                                                } else {
                                                                    v27 = (-823 != -300)
                                                                }
                                                            default:
                                                                var v37 float32 = v26
                                                                _ = v37
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v38 bool = (!((v5 | 42) > (v5 / ((v5) | 1))))
                                                                _ = v38
                                                            }
                                                            switch ("\n\t" + (("foo bar" + "a") + ((("\n\t" + (("" + ("hello" + ("hello" + "foo bar"))) + ("\n\t" + "a"))) + "hello") + ""))) {
                                                            case "hello":
                                                                var v39 string = "\n\t"
                                                                _ = v39
                                                            case "a":
                                                                v27 = (2147483648 >= 2147483647)
                                                            default:
                                                                var v40 string = "a"
                                                                v1[(fi15 % ((fi13) | 1))] = 32768
                                                                v1[(fi13 * 42)] = -228
                                                                _ = v40
                                                            }
                                                        }
                                                        fmt.Println(v21)
                                                        var v41 uint = v10
                                                        v42 := GenS1{F0: 1000}
                                                        v43 := GenS1{F0: 10}
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                            case "\n\t":
                                                var v44 uint64 = 857
                                                _ = v44
                                            default:
                                                var v45 bool = true
                                                switch -9223372036854775808 {
                                                case 1:
                                                    v1[4294967295] = (fi13 ^ (fi15 ^ fi13))
                                                }
                                                delete(v1, fi15)
                                                _ = v45
                                            }
                                            v16 = v11
                                        }
                                    case -335:
                                        v16 -= (v16 - -9223372036854775808)
                                    default:
                                        switch -2147483648 {
                                        case -9223372036854775808:
                                            if v27 {
                                                var v46 uint32 = 2147483647
                                                switch (fi13 * -9223372036854775808) {
                                                case 128:
                                                    var v47 int8 = -96
                                                    _ = v47
                                                case -612:
                                                    v16 = v11
                                                }
                                                var v48 int16 = -599
                                                _ = v46
                                                _ = v48
                                            }
                                        }
                                    }
                                    _ = len(v1)
                                } else {
                                    v11 = v16
                                    v27 = v27
                                    v49 := v1[-449]
                                    fmt.Println(v10, fi15)
                                    _ = v49
                                }
                            }
                            _ = v27
                        } else {
                            v5 = 1
                        }
                        fmt.Println(v5, fi13)
                        _ = v26
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v50 int32 = 0
                        delete(v1, 897)
                        if true {
                            v51 := v1[-1]
                            var v52 uint64 = 4294967295
                            v1[fi13] = -933
                            _ = v51
                            _ = v52
                        } else {
                            var v53 int8 = 127
                            _ = v53
                        }
                        _ = v50
                    }
                    v16 -= v16
                } else {
                    switch "a" {
                    case "hello":
                        if true {
                            v2 = append(v2, 991)
                            var v54 uint8 = 255
                            if true {
                                if (65535 == v10) {
                                    var v55 uint32 = 4294967295
                                    _ = v55
                                } else {
                                    if (!(false || false)) {
                                        switch (fi15 % ((fi15) | 1)) {
                                        case 137:
                                            v16 -= 1
                                        case -600:
                                            fmt.Println(v11, fi15, v16)
                                            v10 += (v10 % ((4294967295) | 1))
                                            fmt.Println(v11, v16, v16)
                                        case -758:
                                            v16 -= -1
                                        }
                                    } else {
                                        v56 := GenS2{F0: 605, F1: 32767}
                                        fmt.Println(fi15, fi15)
                                        delete(v1, fi13)
                                        _ = v56
                                    }
                                    v5 += (v54 << ((18446744073709551615) & 63))
                                    if true {
                                        if (612 <= 567) {
                                            switch (fi13 * fi15) {
                                            case -9223372036854775808:
                                                v57 := GenS4{F0: 1, F1: 165}
                                                _ = v57
                                            case 0:
                                                fmt.Println(v21)
                                            case 2147483647:
                                                if ((v10 == v10) || false) {
                                                    v58 := v1[127]
                                                    _ = v58
                                                } else {
                                                    v59 := GenS5{F0: -1, F1: 235}
                                                    var v60 int = fi13
                                                    var v61 uint8 = (v5 >> (((v10 >> ((0) & 63))) & 63))
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                            fmt.Println(fi13, v10)
                                            var v62 float64 = 3.14159
                                            var v63 uint8 = (v5 - 127)
                                            v10 = 0
                                            v64 := GenS2{F0: 1, F1: -1}
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                        if (177 == (v54 % ((127) | 1))) {
                                            v21 += (v21 * v21)
                                        } else {
                                            var v65 bool = (2 == 1)
                                            _ = v65
                                        }
                                        switch fi13 {
                                        case 32768:
                                            v66 := GenS5{F0: 37, F1: v54}
                                            _ = v66
                                        default:
                                            var v67 int32 = 2147483647
                                            _ = v67
                                        }
                                    } else {
                                        if (4294967295 >= 18446744073709551615) {
                                            if false {
                                                if (!((857 != 0) || (1 != 106))) {
                                                    if false {
                                                        v54 -= 128
                                                        switch fi15 {
                                                        case 0:
                                                            _ = len(v1)
                                                        }
                                                    } else {
                                                        v21 -= v21
                                                        if len(v2) > 0 {
                                                            v2[0] = 1
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            if true {
                                                v21 += v21
                                                delete(v1, 32767)
                                                v11 -= v11
                                            } else {
                                                var v68 uint16 = 2
                                                _ = v68
                                            }
                                        }
                                        if true {
                                            var v69 uint8 = (v54 << ((2) & 63))
                                            v69 = (v5 << (((v10 - (v10 << ((328) & 63)))) & 63))
                                            v70 := GenS5{F0: 0, F1: v54}
                                            _ = v69
                                            _ = v70
                                        }
                                        switch (-(fi13)) {
                                        case 1:
                                            v71 := GenS2{F0: 42, F1: 46}
                                            _ = v71
                                        case -1:
                                            var v72 float64 = v21
                                            _ = v72
                                        case 218:
                                            _ = len(v2)
                                        }
                                        var v73 uint = (v10 ^ 4294967295)
                                        _ = v73
                                    }
                                    v5 = 128
                                }
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                                delete(v1, (-(fi15)))
                            } else {
                                var v74 uint64 = 1
                                v75 := GenS4{F0: 77, F1: (v74 / (((v74 % ((v74) | 1))) | 1))}
                                v76 := GenS1{F0: 279}
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                            _ = v54
                        } else {
                            if true {
                                var v77 bool = (86 < 1)
                                _ = v77
                            }
                        }
                    }
                }
                v10 += v10
                _ = v21
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v11, v11, v11)
    v5 = 213
    if (!((0.0 - 0.0) != 1.5)) {
        for fi78 := 0; fi78 < 10; fi78++ {
            var v79 int64 = v11
            if (1.5 == (1.5 + 3.14159)) {
                var v80 uint = v10
                _ = v80
            } else {
                switch "" {
                case "\n\t":
                    if true {
                        var v81 int32 = -977
                        _ = v81
                    } else {
                        if ((327 > -316) && true) {
                            if (100 > int8(v11)) {
                                v79 = v79
                                var v82 string = "\n\t"
                                _ = v82
                            }
                            fmt.Println(v79)
                        } else {
                            switch "\n\t" {
                            case "a":
                                var v83 bool = ((!false) || false)
                                _ = v83
                            case "":
                                switch fi78 {
                                case 127:
                                    var v84 int8 = 72
                                    _ = v84
                                case 701:
                                    switch 127 {
                                    case -923:
                                        v85 := GenS6{F0: 0, F1: (("a" + "") + "hello")}
                                        _ = v85
                                    case -9223372036854775808:
                                        if ((!true) || false) {
                                            var v86 bool = false
                                            v87 := GenS3{F0: v5}
                                            switch ("hello" + "foo bar") {
                                            case "foo bar":
                                                var v88 int = 174
                                                _ = v88
                                            }
                                            _ = v86
                                            _ = v87
                                        } else {
                                            fmt.Println(v5, v79, fi78)
                                        }
                                    case 32767:
                                        v11 = (v79 - (v79 | 687))
                                    default:
                                        if true {
                                            v10 += 346
                                            var v89 uint32 = 1
                                            var v90 uint = (v10 & (v10 * (v10 ^ (v10 ^ (v10 / ((v10) | 1))))))
                                            _ = v89
                                            _ = v90
                                        } else {
                                            var v91 uint64 = 321
                                            _ = v91
                                        }
                                    }
                                default:
                                    switch "hello" {
                                    case "hello":
                                        var v92 int64 = v11
                                        _ = v92
                                    default:
                                        var v93 string = ("a" + "")
                                        _ = v93
                                    }
                                }
                            case "\n\t":
                                if true {
                                    var v94 string = (("hello" + "a") + ("foo bar" + ("hello" + "hello")))
                                    fmt.Println(v94)
                                    _ = v94
                                }
                            }
                            fmt.Println(fi78, v11, v11)
                        }
                    }
                case "a":
                    v5 = 52
                default:
                    var v95 float64 = 1.5
                    _ = v95
                }
                for fi96 := 0; fi96 < 8; fi96++ {
                    var v97 uint8 = 189
                    _ = fi96
                    _ = v97
                }
            }
            v98 := GenS4{F0: 48, F1: 18446744073709551615}
            _ = fi78
            _ = v79
            _ = v98
        }
        fmt.Println(v10, v5, v10)
        var v99 uint16 = 32767
        _ = v99
    }
    if (!((true && true) || (!(3.14159 <= -2.25)))) {
        switch "foo bar" {
        case "":
            switch -171 {
            case 4294967295:
                v5 = v5
            case 0:
                var v100 float64 = 0.0
                _ = v100
            case -1:
                var v101 int64 = (v11 & (-(v11)))
                _ = v101
            }
        case "\n\t":
            if true {
                for fi102 := 0; fi102 < 9; fi102++ {
                    if len(v3) > 0 {
                        v3[0] = 88
                    }
                    _ = fi102
                }
            } else {
                v10 = (v10 / ((42) | 1))
                if false {
                    fmt.Println(v11, v5, v11)
                    v103 := GenS6{F0: 438, F1: "hello"}
                    var v104 float32 = 1.5
                    _ = v103
                    _ = v104
                } else {
                    v105 := GenS4{F0: 100, F1: 42}
                    _ = v105
                }
            }
        case "hello":
            fmt.Println(v5, v5)
        }
    }
    fmt.Println(v10)
    if (-2147483648 < -146) {
        v106 := GenS4{F0: 100, F1: 4294967295}
        v107 := GenS7{F0: (v5 | 152), F1: -2147483648, F2: 1}
        _ = v106
        _ = v107
    }
    v108 := v1[2147483647]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v108:", v108)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v108
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
