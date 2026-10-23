package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int8; F1 int32; F2 int16 }
type GenS6 struct { F0 int; F1 int32; F2 uint }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string; F1 int32 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int8 }
func genFn0() uint16 {
    v5 := GenS1{F0: 42, F1: -1}
    v6 := []int{923, -851}
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint64, p1 uint16, p2 uint16) int32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 1, F1: int16(fi7)}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        switch (^(fi9)) {
        case 1:
            var v10 int32 = 42
            _ = v10
        case 128:
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS2{F0: 42, F1: 42}
                v13 := GenS3{F0: int8(p2)}
                var v14 uint8 = 1
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []int{-9223372036854775808, -2147483648, 128, 9223372036854775807}
    var v2 int8 = 127
    var v3 int8 = (v2 << ((0) & 63))
    var v4 float32 = -2.25
    v4 = v4
    v15 := GenS4{F0: 128, F1: 127}
    switch -790 {
    case 9223372036854775807:
        var v16 bool = true
        _ = v16
    case 32767:
        if (9223372036854775807 == 814) {
            fmt.Println(v2, v4, v2)
            if len(v1) > 0 {
                v1[0] = 0
            }
        }
    case 42:
        v1 = append(v1, 0)
    default:
        v3 = (v3 + -1)
    }
    var v17 bool = false
    if len(v1) > 0 {
        v1[0] = 0
    }
    if (true && (-1 == (v2 >> ((475) & 63)))) {
        switch "\n\t" {
        case "hello":
            switch "hello" {
            case "hello":
                v1 = append(v1, -1)
            }
        case "a":
            for fi18 := 0; fi18 < 19; fi18++ {
                for fi19 := 0; fi19 < 13; fi19++ {
                    if v17 {
                        var v20 int16 = 0
                        v1 = append(v1, fi18)
                        _ = v20
                    } else {
                        var v21 int16 = -859
                        switch "hello" {
                        case "a":
                            v17 = false
                        case "\n\t":
                            fmt.Println(v3)
                        case "foo bar":
                            if true {
                                v2 = -1
                            } else {
                                var v22 uint16 = 42
                                var v23 int8 = v2
                                var v24 uint64 = 1
                                v1 = append(v1, fi18)
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        default:
                            v21 += (v21 & (-(v21)))
                        }
                        _ = v21
                    }
                    if (71 < 408) {
                        v2 += (v2 * int8(fi19))
                        if (!true) {
                            var v25 int8 = (v2 >> ((255) & 63))
                            switch "a" {
                            case "a":
                                v3 -= (v2 - (v2 ^ v3))
                            case "\n\t":
                                v3 = 0
                            }
                            var v26 int32 = 42
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS2{F0: 1, F1: 2}
                        _ = v27
                    }
                    _ = len(v1)
                    _ = fi19
                }
                var v28 uint = 1
                _ = fi18
                _ = v28
            }
        case "\n\t":
            var v29 bool = (v17 && (((0.0 - ((float64(1.5) * 1e10) - 0.1)) / 3.14159) >= 0.1))
            _ = v29
        }
        var v30 int = -9223372036854775808
        switch (v30 - (v30 >> ((42) & 63))) {
        case 4294967295:
            if len(v1) > 0 {
                v1[0] = (v30 - 586)
            }
        }
        _ = v30
    }
    var v31 uint16 = 0
    v32 := GenS4{F0: 522, F1: 127}
    fmt.Println(v4)
    var v33 uint = 255
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 uint32 = 956
        var v36 int16 = -213
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := GenS5{F0: (v2 % ((42) | 1)), F1: genFn1(61, 65535, genFn0()), F2: 1000}
    v1 = append(v1, 127)
    for fi38 := 0; fi38 < 3; fi38++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi39 := 0; fi39 < 6; fi39++ {
            _ = len(v1)
            if len(v1) > 0 {
                v1[0] = (fi39 & fi39)
            }
            if len(v1) > 0 {
                v1[0] = (fi39 >> (((v33 ^ v33)) & 63))
            }
            if (615 != 42) {
                v4 = float32(0.1)
                var v40 uint = (^(v33))
                if (255 < 98) {
                    if (2 >= 803) {
                        switch (-(fi39)) {
                        case -1:
                            if len(v1) > 0 {
                                v1[0] = (fi38 - fi38)
                            }
                        case 128:
                            var v41 uint64 = uint64(v31)
                            _ = v41
                        }
                        if (!(genFn1(4294967295, genFn0(), 1) < 439)) {
                            if false {
                                var v42 int32 = -2147483648
                                _ = v42
                            } else {
                                fmt.Println(fi38)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = fi38
                            }
                        } else {
                            v43 := GenS6{F0: (fi38 * (fi38 | (fi39 << (((v40 + (v40 * (v33 + v33)))) & 63)))), F1: -823, F2: (v33 & v40)}
                            if (!true) {
                                var v44 uint32 = 1
                                v33 = v40
                                _ = v44
                            }
                            v1 = append(v1, fi38)
                            _ = v43
                        }
                        switch ("a" + "hello") {
                        case "foo bar":
                            switch fi39 {
                            case 2147483647:
                                if (!((128 >= 42) && (!((^(v31)) <= v31)))) {
                                    if ((("a" + "hello") + ("hello" + "a")) == ("" + "\n\t")) {
                                        _ = len(v1)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = len(v1)
                                    } else {
                                        v45 := GenS2{F0: 599, F1: (v33 - 607)}
                                        v46 := GenS5{F0: (v3 << ((v40) & 63)), F1: genFn1(uint64(v33), v31, 32767), F2: 1}
                                        v2 = 119
                                        _ = v45
                                        _ = v46
                                    }
                                    if v17 {
                                        v47 := GenS4{F0: fi39, F1: 68}
                                        if (-598 != 9223372036854775807) {
                                            v48 := GenS7{F0: (fi39 % (((-(fi39))) | 1)), F1: 20}
                                            _ = v48
                                        }
                                        v17 = v17
                                        _ = v47
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    var v49 bool = (!((v40 * 1) != (v33 << (((v33 ^ 4294967295)) & 63))))
                                    _ = v49
                                }
                            case -760:
                                v2 = (v2 & (-(v2)))
                            case -2147483648:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case "a":
                            if ((true && (2147483647 >= 0)) || v17) {
                                v3 -= (v3 % ((-60) | 1))
                                fmt.Println(fi38, v2)
                                v17 = (v17 && v17)
                            }
                        case "hello":
                            if ((42 < -250) || v17) {
                                fmt.Println(v2, v2)
                                if (-9223372036854775808 > -9223372036854775808) {
                                    fmt.Println(v33, fi39, fi38)
                                    if (v17 && (1 == 4294967295)) {
                                        var v50 int16 = 1000
                                        v51 := GenS8{F0: 42}
                                        _ = v50
                                        _ = v51
                                    } else {
                                        if v17 {
                                            fmt.Println(v4)
                                            fmt.Println(v17, v2)
                                            if v17 {
                                                if (!(genFn0() == genFn0())) {
                                                    if true {
                                                        var v52 uint16 = v31
                                                        var v53 int32 = -91
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    var v54 uint = (^(v40))
                                                    _ = v54
                                                } else {
                                                    v3 = (v2 >> (((v40 * (v40 / (((^(v40))) | 1)))) & 63))
                                                }
                                                v31 -= v31
                                                v40 += v40
                                            }
                                            v31 = (^(v31))
                                        } else {
                                            var v55 int = fi38
                                            _ = v55
                                        }
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v31 += v31
                            } else {
                                v17 = (!v17)
                            }
                        default:
                            v1 = append(v1, (fi38 | (fi38 / (((fi38 ^ int(v33))) | 1))))
                        }
                        var v56 int8 = v2
                        _ = v56
                    } else {
                        var v57 uint32 = 635
                        _ = v57
                    }
                    fmt.Println(v17, v3, v17)
                } else {
                    var v58 uint = v33
                    if false {
                        v58 += v40
                        var v59 uint32 = 2
                        _ = v59
                    } else {
                        var v60 int8 = v2
                        _ = v60
                    }
                    _ = v58
                }
                _ = v40
            }
            _ = fi39
        }
        for fi61 := 0; fi61 < 10; fi61++ {
            switch (fi61 | -9223372036854775808) {
            case 32768:
                var v62 int64 = 472
                _ = v62
            case 127:
                if v17 {
                    v33 = 47
                    v3 -= v3
                    fmt.Println(v17)
                } else {
                    if (v31 <= v31) {
                        switch "\n\t" {
                        case "a":
                            if (true || ("a" > "a")) {
                                var v63 int16 = -1
                                v2 -= int8(v31)
                                if v17 {
                                    var v64 float32 = (((v4 / 1e10) * 0.0) + float32(1e10))
                                    switch (fi38 * (fi61 + 32767)) {
                                    case 189:
                                        v64 += (v64 / v64)
                                    case 128:
                                        v1 = append(v1, -2147483648)
                                        var v65 uint32 = 590
                                        var v66 bool = (v17 || (v17 && v17))
                                        var v67 uint = v33
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    default:
                                        switch fi61 {
                                        case 42:
                                            v31 = (v31 & (v31 * (v31 << (((^(v33))) & 63))))
                                        case 2147483647:
                                            v3 = -53
                                            switch "foo bar" {
                                            case "a":
                                                if true {
                                                    v31 -= (v31 ^ genFn0())
                                                    switch (fi61 & (fi38 << ((v33) & 63))) {
                                                    case 661:
                                                        v4 -= (v64 * v64)
                                                    case -227:
                                                        var v68 float64 = (1e10 / (-2.25 * 3.14159))
                                                        _ = v68
                                                    default:
                                                        var v69 int = fi38
                                                        _ = v69
                                                    }
                                                    v2 = v3
                                                }
                                                fmt.Println(fi61, v3, fi38)
                                            case "foo bar":
                                                var v70 int = (fi38 >> ((2) & 63))
                                                _ = v70
                                            default:
                                                v71 := GenS9{F0: 255, F1: true, F2: 65535}
                                                _ = v71
                                            }
                                        case 4294967295:
                                            v72 := GenS9{F0: 0, F1: ((fi38 % ((361) | 1)) != (fi61 & (fi38 % (((-(fi61))) | 1)))), F2: 0}
                                            _ = v72
                                        default:
                                            v3 = v3
                                        }
                                    }
                                    var v73 int8 = (^(v2))
                                    _ = v64
                                    _ = v73
                                } else {
                                    _ = len(v1)
                                    var v74 float64 = 0.1
                                    var v75 string = "foo bar"
                                    _ = v74
                                    _ = v75
                                }
                                _ = v63
                            }
                        case "foo bar":
                            if v17 {
                                v1 = append(v1, (-(fi38)))
                                if true {
                                    v1 = append(v1, (fi61 << (((v33 ^ 825)) & 63)))
                                }
                            } else {
                                switch 9223372036854775807 {
                                case -1:
                                    if v17 {
                                        var v76 int8 = (v2 & v2)
                                        _ = v76
                                    } else {
                                        v31 = v31
                                        v2 += v2
                                    }
                                case 4294967295:
                                    var v77 int64 = 9223372036854775807
                                    _ = v77
                                case -2147483648:
                                    v4 += v4
                                }
                                v4 = 1.5
                                v78 := GenS10{F0: ((v3 | -85) >= v3), F1: 205}
                                var v79 bool = v17
                                if true {
                                    if (!false) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v80 := GenS10{F0: (2 == 257), F1: 42}
                                        if v17 {
                                            if len(v1) > 0 {
                                                v1[0] = fi61
                                            }
                                            fmt.Println(fi38, v33, v33)
                                        }
                                        _ = v80
                                    } else {
                                        v1 = append(v1, (fi61 | 397))
                                        v2 = (v3 << ((18446744073709551615) & 63))
                                        switch fi38 {
                                        case 945:
                                            var v81 uint8 = 127
                                            _ = v81
                                        default:
                                            v82 := GenS11{F0: (v33 % ((42) | 1)), F1: 1}
                                            _ = v82
                                        }
                                    }
                                    var v83 uint16 = v31
                                    _ = v83
                                }
                                _ = v78
                                _ = v79
                            }
                        }
                    } else {
                        v4 = float32(0.0)
                        switch "foo bar" {
                        case "foo bar":
                            switch "\n\t" {
                            case "foo bar":
                                var v84 int = fi61
                                _ = v84
                            default:
                                var v85 bool = (v2 <= 127)
                                _ = v85
                            }
                        default:
                            fmt.Println(v3)
                        }
                        v17 = v17
                        v86 := GenS12{F0: -185, F1: (^(v31)), F2: v31}
                        _ = v86
                    }
                }
            case -1:
                v87 := GenS13{F0: 42}
                _ = v87
            default:
                if (9223372036854775807 == 42) {
                    var v88 int = 387
                    var v89 int64 = -13
                    if (false || (v4 != 0.1)) {
                        if (807 == 2147483647) {
                            if v17 {
                                v90 := GenS8{F0: 42}
                                var v91 bool = ("hello" >= "a")
                                var v92 uint16 = 255
                                _ = v90
                                _ = v91
                                _ = v92
                            }
                        } else {
                            if v17 {
                                v31 = v31
                                if v17 {
                                    var v93 int16 = 366
                                    v17 = v17
                                    _ = v93
                                }
                                var v94 uint16 = v31
                                _ = v94
                            }
                            if (42 <= -1) {
                                if (!true) {
                                    v95 := GenS14{F0: "a", F1: 0}
                                    v96 := GenS14{F0: "hello", F1: -907}
                                    var v97 int64 = int64(fi38)
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                } else {
                                    v31 += (^(v31))
                                    fmt.Println(v4)
                                    var v98 int16 = 1000
                                    v1 = append(v1, -373)
                                    _ = v98
                                }
                                v89 -= 9223372036854775807
                                fmt.Println(v4, v88, v88)
                                var v99 int32 = int32(v33)
                                v88 = v88
                                var v100 int8 = 22
                                _ = v100
                                _ = v99
                            }
                            v101 := GenS15{F0: 1}
                            _ = v101
                        }
                        var v102 int8 = v3
                        _ = v102
                    } else {
                        v103 := GenS13{F0: 255}
                        _ = v103
                    }
                    _ = v88
                    _ = v89
                }
            }
            switch (fi38 ^ (fi38 + (fi61 ^ 855))) {
            case 679:
                v3 = (^(v2))
            case 0:
                var v104 bool = ((!((!false) && v17)) && ("foo bar" < ("\n\t" + "")))
                _ = v104
            }
            var v105 uint64 = uint64(v33)
            _ = fi61
            _ = v105
        }
        _ = fi38
    }
    fmt.Println(v3)
    v1 = append(v1, 128)
    for fi106 := 0; fi106 < 13; fi106++ {
        v107 := GenS16{F0: (v2 * v2)}
        _ = fi106
        _ = v107
    }
    var v108 uint = (v33 / ((42) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v108:", v108)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v108
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v4
}
