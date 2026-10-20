package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int; F1 bool; F2 uint32 }
type GenS8 struct { F0 int16; F1 uint64; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 string; F1 string; F2 int16 }
type GenS11 struct { F0 uint16; F1 string; F2 uint }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int16; F1 int8; F2 uint32 }
type GenS14 struct { F0 uint64; F1 int; F2 uint32 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint8; F1 uint32 }
func genFn0(p0 string, p1 int) int8 {
    var v7 uint64 = 671
    var v8 int64 = -484
    if (0 >= 49) {
        var v9 string = p0
        v10 := GenS1{F0: 128, F1: 1, F2: -1}
        _ = v10
        _ = v9
    }
    if (false && true) {
        v11 := []int{}
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(p1, v8)
            _ = fi12
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -88
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 854
    var v3 int32 = -95
    var v4 float32 = 3.14159
    v5 := []string{}
    v6 := []string{}
    if (65535 > 42) {
        var v13 int = -785
        if ((int16(v3) >= 988) || (true || ((9223372036854775807 <= 42) && false))) {
            v14 := GenS2{F0: (v2 % ((v2) | 1)), F1: 754}
            if false {
                v4 = (v4 - 1.5)
                v4 += v4
                if (9223372036854775807 != -1) {
                    v15 := GenS2{F0: 0, F1: 65535}
                    if (458 >= -1) {
                        fmt.Println(v2, v13, v4)
                    } else {
                        switch ((("hello" + "") + (("\n\t" + "a") + "a")) + "foo bar") {
                        case "a":
                            if true {
                                switch v13 {
                                case 17:
                                    v16 := GenS1{F0: 42, F1: -1, F2: 127}
                                    _ = v16
                                }
                            }
                        }
                    }
                    var v17 uint16 = 132
                    _ = v15
                    _ = v17
                } else {
                    v18 := GenS3{F0: v13, F1: true}
                    _ = v18
                }
            } else {
                switch v13 {
                case 32767:
                    v1[-2147483648] = 2
                case 182:
                    v3 = 669
                case 42:
                    v19 := GenS1{F0: 156, F1: genFn0("", (v13 ^ v13)), F2: genFn0("hello", v13)}
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v4 = v4
                        v13 = v13
                        for fi21 := 0; fi21 < 6; fi21++ {
                            v22 := GenS1{F0: 255, F1: 100, F2: genFn0("foo bar", (^(fi21)))}
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                }
                v23 := v1[v13]
                _ = v23
            }
            _ = v14
        }
        _ = v13
    } else {
        v24 := GenS3{F0: -673, F1: ((!false) && (!(false && false)))}
        _ = v24
    }
    var v25 bool = true
    if v25 {
        for fi26 := 0; fi26 < 3; fi26++ {
            if (!v25) {
                switch "a" {
                case "":
                    v27 := GenS2{F0: v2, F1: 42}
                    _ = v27
                case "hello":
                    v28 := GenS1{F0: 255, F1: -88, F2: 100}
                    _ = v28
                case "a":
                    v2 += (v2 + v2)
                }
            } else {
                switch ("\n\t" + ("a" + ("foo bar" + "\n\t"))) {
                case "a":
                    if v25 {
                        if v25 {
                            var v29 int8 = -1
                            _ = v29
                        }
                    } else {
                        if v25 {
                            v6 = append(v6, "foo bar")
                        } else {
                            v3 += (v3 - v3)
                        }
                        v1[fi26] = 946
                        v25 = ((v25 || (v25 && v25)) || v25)
                    }
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
            _ = fi26
        }
        if (42 != 67) {
            v4 += float32(3.14159)
            if true {
                for fi30 := 0; fi30 < 9; fi30++ {
                    if ((v25 || false) || v25) {
                        if v25 {
                            v31 := GenS4{F0: 100}
                            v32 := GenS5{F0: 4294967295, F1: (v2 + 581)}
                            _ = v31
                            _ = v32
                        }
                        if v25 {
                            var v33 uint32 = 2147483648
                            for fi34 := 0; fi34 < 13; fi34++ {
                                var v35 int8 = 0
                                var v36 uint32 = (v33 >> ((937) & 63))
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 13; fi37++ {
                                v33 += 0
                                fmt.Println(fi30)
                                _ = fi37
                            }
                            _ = v33
                        } else {
                            switch "\n\t" {
                            case "hello":
                                if (false && false) {
                                    for fi38 := 0; fi38 < 17; fi38++ {
                                        v2 = v2
                                        _ = fi38
                                    }
                                    if v25 {
                                        if false {
                                            v25 = v25
                                            v4 -= 1e10
                                            if v25 {
                                                if (255 != 1) {
                                                    for fi39 := 0; fi39 < 18; fi39++ {
                                                        var v40 uint64 = (v2 / ((v2) | 1))
                                                        if (v25 && ((v3 != v3) && (-9223372036854775808 == 0))) {
                                                            v5 = append(v5, ((("\n\t" + "hello") + "foo bar") + "\n\t"))
                                                            v3 = v3
                                                        }
                                                        if (137 < 42) {
                                                            switch fi30 {
                                                            case -162:
                                                                switch 32767 {
                                                                case -827:
                                                                    var v41 float64 = 0.0
                                                                    _ = v41
                                                                case -619:
                                                                    fmt.Println(v3, fi30)
                                                                }
                                                            case 685:
                                                                var v42 uint64 = 2
                                                                _ = v42
                                                            default:
                                                                v2 = 76
                                                            }
                                                        } else {
                                                            v4 += v4
                                                        }
                                                        if true {
                                                            fmt.Println(v25, fi39, v3)
                                                            if (1000 <= 944) {
                                                                var v43 bool = (v40 != v2)
                                                                var v44 int32 = (v3 % ((-2147483648) | 1))
                                                                _ = v43
                                                                _ = v44
                                                            } else {
                                                                v2 -= v40
                                                                if (((v25 && v25) && (!((^(fi39)) == 128))) && (v40 != (v40 + (v40 << ((0) & 63))))) {
                                                                    if (!(132 == 255)) {
                                                                        v3 -= v3
                                                                        if true {
                                                                            var v45 string = "a"
                                                                            var v46 uint16 = 32767
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            _ = v45
                                                                            _ = v46
                                                                        } else {
                                                                            if (1 < 65535) {
                                                                                var v47 uint32 = 2147483648
                                                                                switch "a" {
                                                                                case "":
                                                                                    v2 = v40
                                                                                }
                                                                                if v25 {
                                                                                    switch fi39 {
                                                                                    case 432:
                                                                                        var v48 string = ("a" + "a")
                                                                                        _ = v48
                                                                                    case 2147483647:
                                                                                        switch fi39 {
                                                                                        case -965:
                                                                                            var v49 uint = 444
                                                                                            _ = v49
                                                                                        case 180:
                                                                                            v47 += (^(v47))
                                                                                        default:
                                                                                            var v50 uint8 = 255
                                                                                            switch fi39 {
                                                                                            case 965:
                                                                                                if true {
                                                                                                    var v51 int8 = 100
                                                                                                    switch "" {
                                                                                                    case "hello":
                                                                                                        if (false && v25) {
                                                                                                            switch "\n\t" {
                                                                                                            case "foo bar":
                                                                                                                var v52 int64 = -1
                                                                                                                _ = v52
                                                                                                            }
                                                                                                            switch "\n\t" {
                                                                                                            case "":
                                                                                                                if ((v3 >> ((4294967295) & 63)) == (v3 >> ((779) & 63))) {
                                                                                                                    if ((0.1 / (3.14159 * (1e10 - (0.0 + ((-2.25 * (1.5 - float64((v4 * v4)))) + (0.1 * 1.5)))))) == 1e10) {
                                                                                                                        v2 = (v2 | 42)
                                                                                                                    } else {
                                                                                                                        fmt.Println(v4, v2)
                                                                                                                        if (v3 >= (v3 % ((v3) | 1))) {
                                                                                                                            var v53 uint64 = (v2 ^ (v2 - 436))
                                                                                                                            _ = v53
                                                                                                                        } else {
                                                                                                                            v54 := GenS2{F0: (v40 & (v40 ^ (v2 + v2))), F1: 93}
                                                                                                                            _ = v54
                                                                                                                        }
                                                                                                                        v50 = v50
                                                                                                                    }
                                                                                                                    fmt.Println(v50, v25)
                                                                                                                    var v55 uint = 65535
                                                                                                                    _ = v55
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                switch (fi30 - fi39) {
                                                                                                                case -531:
                                                                                                                    var v56 uint32 = (v47 * v47)
                                                                                                                    _ = v56
                                                                                                                }
                                                                                                            case "a":
                                                                                                                if (v3 >= (v3 | v3)) {
                                                                                                                    fmt.Println(v47, v40, v2)
                                                                                                                } else {
                                                                                                                    if (v3 <= v3) {
                                                                                                                        v57 := GenS6{F0: 42, F1: "hello"}
                                                                                                                        _ = v57
                                                                                                                    }
                                                                                                                    var v58 uint32 = (^(v47))
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v40)
                                                                                                        } else {
                                                                                                            var v59 uint = 65535
                                                                                                            var v60 uint = (v59 | 65535)
                                                                                                            if (v3 > 0) {
                                                                                                                var v61 int = 32767
                                                                                                                _ = v61
                                                                                                            } else {
                                                                                                                fmt.Println(v3)
                                                                                                                v60 = (v59 % ((0) | 1))
                                                                                                            }
                                                                                                            var v62 uint8 = v50
                                                                                                            if true {
                                                                                                                var v63 uint32 = v47
                                                                                                                _ = v63
                                                                                                            }
                                                                                                            _ = v59
                                                                                                            _ = v60
                                                                                                            _ = v62
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v25, fi30)
                                                                                                    default:
                                                                                                        v64 := GenS7{F0: (fi30 ^ fi39), F1: false, F2: (v47 % ((v47) | 1))}
                                                                                                        _ = v64
                                                                                                    }
                                                                                                    _ = v51
                                                                                                } else {
                                                                                                    var v65 int8 = 94
                                                                                                    switch 32768 {
                                                                                                    case 9223372036854775807:
                                                                                                        if (!v25) {
                                                                                                            var v66 int8 = (v65 % (((v65 - (v65 ^ v65))) | 1))
                                                                                                            _ = v66
                                                                                                        }
                                                                                                    case 0:
                                                                                                        if (v25 || false) {
                                                                                                            v67 := v1[fi30]
                                                                                                            v68 := GenS4{F0: genFn0("hello", fi30)}
                                                                                                            v69 := GenS6{F0: -91, F1: "foo bar"}
                                                                                                            _ = v67
                                                                                                            _ = v68
                                                                                                            _ = v69
                                                                                                        } else {
                                                                                                            switch fi30 {
                                                                                                            case 460:
                                                                                                                v47 -= uint32(v2)
                                                                                                            case 127:
                                                                                                                v25 = v25
                                                                                                            case -1:
                                                                                                                if false {
                                                                                                                    switch 0 {
                                                                                                                    case 544:
                                                                                                                        fmt.Println(v25, v2, fi39)
                                                                                                                    case 128:
                                                                                                                        fmt.Println(fi39, v25)
                                                                                                                    }
                                                                                                                    var v70 int16 = 958
                                                                                                                    _ = v70
                                                                                                                } else {
                                                                                                                    v71 := GenS1{F0: v50, F1: 1, F2: (v65 / ((v65) | 1))}
                                                                                                                    if v25 {
                                                                                                                        v72 := GenS8{F0: 32767, F1: (v2 | v2), F2: 9223372036854775807}
                                                                                                                        _ = v72
                                                                                                                    } else {
                                                                                                                        if v25 {
                                                                                                                            switch ("hello" + ((("\n\t" + "a") + "") + "a")) {
                                                                                                                            case "foo bar":
                                                                                                                                var v73 float32 = v4
                                                                                                                                v50 = (v50 * 182)
                                                                                                                                _ = v73
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v50 -= (v50 / ((v50) | 1))
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v74 int64 = -909
                                                                                                                    _ = v71
                                                                                                                    _ = v74
                                                                                                                }
                                                                                                            default:
                                                                                                                var v75 uint = 539
                                                                                                                _ = v75
                                                                                                            }
                                                                                                            if ((1.5 + ((((float64(float32(float64(0.1))) + 0.1) - (-2.25 - (0.1 - -2.25))) / 1.0) - (3.14159 * 3.14159))) > 1e10) {
                                                                                                                v47 = v47
                                                                                                                var v76 float32 = ((1e10 - v4) + -2.25)
                                                                                                                _ = v76
                                                                                                            } else {
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v77 := GenS8{F0: 42, F1: (v2 % ((v40) | 1)), F2: -9223372036854775808}
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    var v78 int = (fi39 * 626)
                                                                                                    _ = v65
                                                                                                    _ = v78
                                                                                                }
                                                                                                v79 := GenS9{F0: ("foo bar" + "\n\t")}
                                                                                                _ = v79
                                                                                            case -287:
                                                                                                var v80 string = (("a" + "foo bar") + "a")
                                                                                                _ = v80
                                                                                            case 127:
                                                                                                fmt.Println(fi30)
                                                                                            default:
                                                                                                var v81 uint8 = (v50 * uint8(v47))
                                                                                                _ = v81
                                                                                            }
                                                                                            _ = v50
                                                                                        }
                                                                                    case 1:
                                                                                        v2 = v40
                                                                                    }
                                                                                    var v82 int64 = 1
                                                                                    var v83 uint8 = 74
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                } else {
                                                                                    v3 -= v3
                                                                                    var v84 int16 = -985
                                                                                    v85 := GenS10{F0: "foo bar", F1: ("hello" + "foo bar"), F2: -153}
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                }
                                                                                _ = v47
                                                                            }
                                                                            v25 = v25
                                                                            v86 := GenS10{F0: "\n\t", F1: "\n\t", F2: 42}
                                                                            _ = v86
                                                                        }
                                                                        if v25 {
                                                                            if v25 {
                                                                                v2 -= v40
                                                                            } else {
                                                                                if (!v25) {
                                                                                    var v87 int32 = v3
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v88 := GenS11{F0: 65535, F1: "hello", F2: uint(v2)}
                                                                                        _ = v88
                                                                                    case "a":
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    v3 = 0
                                                                                    _ = v87
                                                                                } else {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (("" + "\n\t") + "\n\t")
                                                                                    }
                                                                                    if (!v25) {
                                                                                        if v25 {
                                                                                            fmt.Println(v2, v25, v3)
                                                                                        }
                                                                                        var v89 int = (fi30 << ((65535) & 63))
                                                                                        _ = v89
                                                                                    } else {
                                                                                        v25 = v25
                                                                                        var v90 bool = (!v25)
                                                                                        var v91 uint16 = uint16(v2)
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    }
                                                                                    v92 := GenS12{F0: ((1 >= -176) || ((v25 || ((false || ((fi30 % ((fi39) | 1)) <= fi39)) && (0 <= 1))) && (((fi39 & 128) >= fi30) || (2 > 2147483647))))}
                                                                                    _ = v92
                                                                                }
                                                                                if (13 >= -58) {
                                                                                    v2 = 4294967295
                                                                                    switch (fi30 / ((-153) | 1)) {
                                                                                    case 95:
                                                                                        v25 = (true || true)
                                                                                    case 32768:
                                                                                        var v93 uint16 = 42
                                                                                        _ = v93
                                                                                    default:
                                                                                        if true {
                                                                                            v3 = v3
                                                                                            _ = len(v5)
                                                                                        } else {
                                                                                            var v94 bool = v25
                                                                                            var v95 uint = 75
                                                                                            v25 = (!((^(v3)) == v3))
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v96 uint16 = 2
                                                                        _ = v96
                                                                    }
                                                                } else {
                                                                    var v97 int32 = v3
                                                                    if ((("a" >= "foo bar") || (("foo bar" + (("a" + "") + "a")) == "a")) || false) {
                                                                        v97 += 361
                                                                    } else {
                                                                        fmt.Println(v97)
                                                                    }
                                                                    _ = v97
                                                                }
                                                            }
                                                        }
                                                        _ = fi39
                                                        _ = v40
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        v98 := GenS13{F0: 1000, F1: 42, F2: 924}
                                                        _ = v98
                                                    case "\n\t":
                                                        fmt.Println(v3)
                                                    case "":
                                                        var v99 uint32 = 2147483648
                                                        _ = v99
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v2)
                                                var v100 int = fi30
                                                _ = v100
                                            }
                                        } else {
                                            if true {
                                                v4 += v4
                                            } else {
                                                switch fi30 {
                                                case 4294967295:
                                                    v25 = v25
                                                case 405:
                                                    var v101 float32 = v4
                                                    _ = v101
                                                case -2147483648:
                                                    for fi102 := 0; fi102 < 2; fi102++ {
                                                        v2 += (v2 % (((v2 & (v2 / ((v2) | 1)))) | 1))
                                                        switch "\n\t" {
                                                        case "":
                                                            var v103 int64 = 9223372036854775807
                                                            _ = v103
                                                        case "hello":
                                                            var v104 float32 = 1.5
                                                            _ = v104
                                                        case "a":
                                                            v4 = v4
                                                            fmt.Println(fi30, fi102)
                                                            var v105 uint32 = 816
                                                            _ = v105
                                                        }
                                                        v25 = false
                                                        v25 = (!v25)
                                                        if true {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v3 += -668
                                                            case "":
                                                                v4 = float32(((((-2.25 + 1e10) + (float64(v4) - ((1e10 + ((0.0 * (1e10 * 0.0)) * (0.1 / float64(1e10)))) / ((0.0 + -2.25) + (0.1 * 0.1))))) - (0.1 / (0.0 - float64(v4)))) / 1.0))
                                                            }
                                                            v25 = (42 != 1)
                                                            _ = len(v6)
                                                        } else {
                                                            v4 += float32(((float64(v4) * -2.25) - 1e10))
                                                            var v106 uint = 903
                                                            _ = len(v1)
                                                            _ = v106
                                                        }
                                                        _ = fi102
                                                    }
                                                }
                                                var v107 int32 = v3
                                                _ = v107
                                            }
                                            switch 0 {
                                            case 23:
                                                fmt.Println(v4, v4)
                                            default:
                                                var v108 float32 = (v4 - 3.14159)
                                                _ = v108
                                            }
                                            delete(v1, fi30)
                                        }
                                    }
                                    for fi109 := 0; fi109 < 19; fi109++ {
                                        fmt.Println(fi109)
                                        switch "hello" {
                                        case "hello":
                                            var v110 float32 = ((v4 + (float32(((((0.1 - float64(v4)) * 0.1) - 0.1) / 1e10)) - -2.25)) * 1e10)
                                            _ = v110
                                        }
                                        var v111 uint16 = 529
                                        v3 = 691
                                        v6 = append(v6, "foo bar")
                                        fmt.Println(fi109, fi30)
                                        _ = fi109
                                        _ = v111
                                    }
                                }
                            case "a":
                                var v112 uint64 = (^(v2))
                                _ = v112
                            case "\n\t":
                                switch 127 {
                                case 128:
                                    var v113 uint32 = 197
                                    _ = v113
                                case 32768:
                                    if v25 {
                                        v114 := v1[fi30]
                                        _ = v114
                                    }
                                }
                            default:
                                for fi115 := 0; fi115 < 19; fi115++ {
                                    v2 += v2
                                    var v116 float64 = 0.0
                                    _ = fi115
                                    _ = v116
                                }
                            }
                            switch "a" {
                            case "":
                                v117 := GenS3{F0: (^(fi30)), F1: v25}
                                _ = v117
                            case "foo bar":
                                v118 := v1[(fi30 + (-(fi30)))]
                                _ = v118
                            case "\n\t":
                                switch "" {
                                case "hello":
                                    var v119 uint32 = 42
                                    _ = v119
                                case "foo bar":
                                    for fi120 := 0; fi120 < 1; fi120++ {
                                        var v121 bool = false
                                        var v122 bool = (v2 == v2)
                                        v121 = (586 < 65)
                                        var v123 int8 = -1
                                        _ = fi120
                                        _ = v121
                                        _ = v122
                                        _ = v123
                                    }
                                case "":
                                    for fi124 := 0; fi124 < 5; fi124++ {
                                        var v125 uint32 = 2147483648
                                        var v126 int32 = (v3 + (v3 ^ (v3 << ((4294967295) & 63))))
                                        _ = fi124
                                        _ = v125
                                        _ = v126
                                    }
                                }
                            }
                        }
                    }
                    _ = fi30
                }
            }
            _ = len(v6)
        } else {
            if (v2 <= v2) {
                v127 := GenS14{F0: (v2 - 224), F1: 2147483647, F2: 3}
                _ = v127
            }
            switch -802 {
            case 4294967295:
                delete(v1, 601)
            }
        }
    }
    var v128 int = 6
    if v25 {
        switch "a" {
        case "hello":
            var v129 uint8 = 127
            _ = v129
        case "a":
            for fi130 := 0; fi130 < 4; fi130++ {
                v4 -= (1.5 * float32(float64(v4)))
                _ = fi130
            }
        default:
            for fi131 := 0; fi131 < 5; fi131++ {
                switch 127 {
                case -836:
                    v132 := GenS10{F0: ((("foo bar" + "hello") + (("hello" + ("\n\t" + "a")) + "")) + ("a" + "\n\t")), F1: "", F2: 1000}
                    _ = v132
                default:
                    for fi133 := 0; fi133 < 3; fi133++ {
                        v2 = v2
                        v134 := GenS1{F0: 1, F1: -90, F2: 127}
                        var v135 int16 = 32767
                        var v136 float32 = (1e10 + (v4 * 0.1))
                        switch (fi133 * (fi133 ^ fi131)) {
                        case 2147483647:
                            v2 = v2
                        case 4294967295:
                            var v137 float64 = (1e10 - 0.1)
                            _ = v137
                        case -389:
                            fmt.Println(fi131, v4)
                        default:
                            if (173 > 250) {
                                v3 += (v3 | (v3 % ((v3) | 1)))
                                var v138 uint = 2
                                _ = v138
                            } else {
                                var v139 uint32 = 65535
                                delete(v1, (^(fi131)))
                                v128 -= int(v135)
                                _ = v139
                            }
                        }
                        _ = fi133
                        _ = v134
                        _ = v135
                        _ = v136
                    }
                }
                var v140 uint16 = 32767
                _ = fi131
                _ = v140
            }
        }
    }
    if v25 {
        for fi141 := 0; fi141 < 10; fi141++ {
            if v25 {
                v128 = v128
                for fi142 := 0; fi142 < 10; fi142++ {
                    var v143 string = "foo bar"
                    if (0 >= -1) {
                        switch fi142 {
                        case -2147483648:
                            var v144 int8 = genFn0((v143 + (v143 + (((v143 + "a") + ((((v143 + v143) + (("" + v143) + v143)) + ("hello" + v143)) + v143)) + v143))), 9223372036854775807)
                            _ = v144
                        default:
                            switch v143 {
                            case "a":
                                v145 := GenS15{F0: fi141}
                                _ = v145
                            }
                        }
                        fmt.Println(fi141)
                        v5 = append(v5, "foo bar")
                    }
                    _ = fi142
                    _ = v143
                }
                switch (v128 & fi141) {
                case -1:
                    for fi146 := 0; fi146 < 2; fi146++ {
                        fmt.Println(v128)
                        _ = fi146
                    }
                case 34:
                    v147 := GenS16{F0: -128}
                    _ = v147
                }
            }
            v148 := GenS14{F0: (v2 / ((v2) | 1)), F1: (^(v128)), F2: 2147483647}
            _ = fi141
            _ = v148
        }
        fmt.Println(v128)
        v2 -= 28
    } else {
        var v149 uint8 = 56
        v150 := GenS17{F0: (v149 * (v149 >> ((80) & 63))), F1: 42}
        _ = v149
        _ = v150
    }
    v4 = v4
    for fi151 := 0; fi151 < 19; fi151++ {
        if len(v6) > 0 {
            v6[0] = "a"
        }
        _ = fi151
    }
    v2 -= v2
    for fi152 := 0; fi152 < 3; fi152++ {
        var v153 uint16 = 32767
        _ = fi152
        _ = v153
    }
    if v25 {
        switch ("\n\t" + ("a" + (("a" + ("hello" + (("\n\t" + "hello") + ("hello" + "hello")))) + "foo bar"))) {
        case "":
            delete(v1, 637)
        case "a":
            v25 = ((v4 >= v4) && v25)
        default:
            switch (^(v128)) {
            case 32767:
                var v154 int32 = v3
                _ = v154
            case 496:
                delete(v1, v128)
            case -9223372036854775808:
                v4 = (v4 * v4)
            }
        }
    }
    var v155 int32 = (-(v3))
    var v156 int8 = 67
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v128:", v128)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v128
    _ = v155
    _ = v156
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
