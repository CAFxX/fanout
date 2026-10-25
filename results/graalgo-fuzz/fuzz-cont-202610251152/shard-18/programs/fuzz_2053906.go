package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int8 }
func genFn0(p0 int, p1 uint16, p2 int) int32 {
    if true {
        var v5 int8 = int8(p2)
        var v6 bool = false
        if (!v6) {
            p0 += (p0 >> ((4294967295) & 63))
        }
        _ = v5
        _ = v6
    } else {
        var v7 int64 = 42
        _ = v7
    }
    p2 = p2
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int = 128
        _ = fi8
        _ = v9
    }
    switch "foo bar" {
    case "a":
        if true {
            var v10 uint64 = 1
            var v11 uint8 = 98
            fmt.Println(p0, p0)
            _ = v10
            _ = v11
        }
    case "foo bar":
        v12 := []int{(p2 - p0)}
        _ = v12
    }
    switch "\n\t" {
    case "hello":
        if (71 != 198) {
            p0 -= p0
            for fi13 := 0; fi13 < 17; fi13++ {
                p1 += 861
                _ = fi13
            }
            v14 := []int{}
            _ = v14
        } else {
            var v15 string = "hello"
            for fi16 := 0; fi16 < 10; fi16++ {
                switch p0 {
                case 127:
                    p2 += fi16
                case 32767:
                    v17 := []int{(^(fi16))}
                    _ = v17
                }
                if (3.14159 == (0.1 / -2.25)) {
                    v18 := GenS1{F0: -9223372036854775808}
                    var v19 float32 = -2.25
                    p1 += p1
                    _ = v18
                    _ = v19
                } else {
                    p1 += p1
                }
                v15 = ((v15 + v15) + "")
                _ = fi16
            }
            _ = v15
        }
    case "foo bar":
        fmt.Println(p0, p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 uint32, p1 uint16, p2 uint8) int64 {
    v20 := GenS1{F0: -9223372036854775808}
    switch 341 {
    case -1:
        var v21 int16 = 42
        _ = v21
    case 32767:
        var v22 string = ("" + "")
        _ = v22
    case 278:
        for fi23 := 0; fi23 < 10; fi23++ {
            p0 += (^(p0))
            switch -1 {
            case 1:
                v24 := GenS1{F0: 42}
                _ = v24
            case 128:
                if (true || false) {
                    var v25 int16 = 0
                    _ = v25
                }
            default:
                var v26 int16 = 781
                _ = v26
            }
            _ = fi23
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := []int{(-(fi27)), (fi27 * fi27), (fi27 + fi27)}
        switch (-(fi27)) {
        case -1:
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(p2)
                _ = fi29
            }
        case 2147483647:
            if false {
                p2 = (p2 | (^(p2)))
                p1 = (p1 + (^(p1)))
            } else {
                if len(v28) > 0 {
                    v28[0] = -604
                }
            }
        case 20:
            p0 = 2147483647
        default:
            for fi30 := 0; fi30 < 7; fi30++ {
                p0 = p0
                var v31 string = ""
                if len(v28) > 0 {
                    v28[0] = 42
                }
                switch (fi30 ^ fi30) {
                case -211:
                    switch "hello" {
                    case "":
                        if ((-(fi30)) < (fi30 << ((65535) & 63))) {
                            var v32 int16 = 1
                            p2 = p2
                            if (!(4294967295 == 129)) {
                                switch 32767 {
                                case 0:
                                    p1 = 32767
                                case 32767:
                                    p2 -= uint8(p1)
                                default:
                                    if true {
                                        fmt.Println(fi30)
                                        var v33 float32 = -2.25
                                        switch (v31 + "foo bar") {
                                        case "":
                                            var v34 uint64 = 1
                                            _ = v34
                                        case "hello":
                                            v35 := GenS1{F0: -9223372036854775808}
                                            _ = v35
                                        default:
                                            fmt.Println(v32)
                                        }
                                        _ = v33
                                    } else {
                                        var v36 bool = false
                                        v37 := GenS1{F0: 1}
                                        _ = v36
                                        _ = v37
                                    }
                                }
                            } else {
                                if true {
                                    v28 = append(v28, fi30)
                                    switch -9223372036854775808 {
                                    case 0:
                                        var v38 int8 = -1
                                        _ = v38
                                    case 32767:
                                        v39 := GenS1{F0: -9223372036854775808}
                                        _ = v39
                                    case -9223372036854775808:
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    default:
                                        v40 := GenS2{F0: p0}
                                        _ = v40
                                    }
                                } else {
                                    if len(v28) > 0 {
                                        v28[0] = (-(fi30))
                                    }
                                    _ = len(v28)
                                    var v41 int32 = 0
                                    _ = v41
                                }
                                var v42 uint64 = 42
                                _ = v42
                            }
                            _ = v32
                        } else {
                            if len(v28) > 0 {
                                v28[0] = 1
                            }
                            if true {
                                switch v31 {
                                case "foo bar":
                                    v43 := GenS2{F0: p0}
                                    _ = v43
                                default:
                                    v44 := GenS2{F0: p0}
                                    _ = v44
                                }
                            }
                            var v45 int32 = 42
                            if true {
                                var v46 float32 = -2.25
                                _ = v46
                            }
                            switch fi30 {
                            case 32768:
                                if false {
                                    switch (fi27 & 847) {
                                    case 543:
                                        if true {
                                            _ = len(v28)
                                            v31 = v31
                                            switch fi30 {
                                            case 0:
                                                var v47 int64 = 1
                                                _ = v47
                                            case -164:
                                                if len(v28) > 0 {
                                                    v28[0] = 9223372036854775807
                                                }
                                            }
                                        }
                                    }
                                    switch (fi30 | (fi27 ^ (fi27 ^ int(v45)))) {
                                    case 127:
                                        var v48 uint16 = p1
                                        var v49 float64 = 0.0
                                        var v50 int8 = 1
                                        var v51 float64 = (v49 + (0.1 / float64((1.5 / (1.5 / ((3.14159 + 0.0) + -2.25))))))
                                        v49 = v49
                                        switch ("hello" + (("foo bar" + (("a" + v31) + v31)) + "foo bar")) {
                                        case "foo bar":
                                            if true {
                                                if ((p2 ^ 42) != (p2 | 126)) {
                                                    var v52 uint8 = p2
                                                    _ = v52
                                                } else {
                                                    if len(v28) > 0 {
                                                        v28[0] = int(v45)
                                                    }
                                                    if false {
                                                        v51 -= (((v49 - v51) / (v51 + v51)) * v51)
                                                        if true {
                                                            v53 := GenS3{F0: (p0 / ((p0) | 1)), F1: 748}
                                                            _ = len(v28)
                                                            p1 += 970
                                                            _ = v53
                                                        }
                                                    } else {
                                                        if ((v31 >= (v31 + "")) || (187 >= 608)) {
                                                            var v54 int8 = (v50 + (v50 % (((v50 << ((42) & 63))) | 1)))
                                                            _ = v54
                                                        }
                                                    }
                                                    if ((p0 != (^(p0))) || (9223372036854775807 <= int64(v48))) {
                                                        if (1 > -832) {
                                                            var v55 bool = (!false)
                                                            _ = v55
                                                        }
                                                        fmt.Println(v51, fi27, v31)
                                                    } else {
                                                        var v56 uint = 65535
                                                        v31 = "hello"
                                                        if true {
                                                            v50 = 42
                                                            fmt.Println(v48)
                                                            v56 = (v56 | 657)
                                                        } else {
                                                            if true {
                                                                if len(v28) >= 2 {
                                                                    v28 = v28[1:len(v28)]
                                                                }
                                                            } else {
                                                                if true {
                                                                    p0 += p0
                                                                } else {
                                                                    switch -9223372036854775808 {
                                                                    case 352:
                                                                        if len(v28) > 0 {
                                                                            v28[0] = 127
                                                                        }
                                                                    case -9223372036854775808:
                                                                        var v57 uint16 = p1
                                                                        _ = v57
                                                                    case 9223372036854775807:
                                                                        v58 := GenS3{F0: p0, F1: -9223372036854775808}
                                                                        if (-1 > -1) {
                                                                            var v59 uint8 = p2
                                                                            fmt.Println(v45, v49)
                                                                            v60 := GenS2{F0: uint32(v56)}
                                                                            _ = v59
                                                                            _ = v60
                                                                        }
                                                                        v61 := GenS4{F0: (v56 + v56), F1: -137, F2: -146}
                                                                        v48 = (^(p1))
                                                                        fmt.Println(v48)
                                                                        _ = v58
                                                                        _ = v61
                                                                    default:
                                                                        _ = len(v28)
                                                                    }
                                                                    switch (fi27 >> ((228) & 63)) {
                                                                    case 1:
                                                                        v62 := GenS5{F0: 0, F1: 1}
                                                                        _ = v62
                                                                    case -863:
                                                                        if true {
                                                                            if true {
                                                                                p2 -= p2
                                                                                p0 = p0
                                                                            } else {
                                                                                v50 -= (v50 >> (((^(v56))) & 63))
                                                                                if true {
                                                                                    var v63 uint64 = uint64(p2)
                                                                                    if false {
                                                                                        var v64 bool = true
                                                                                        if len(v28) > 0 {
                                                                                            v28[0] = (-(fi30))
                                                                                        }
                                                                                        switch v31 {
                                                                                        case "a":
                                                                                            v28 = append(v28, (fi30 << ((v56) & 63)))
                                                                                        case "foo bar":
                                                                                            if len(v28) > 0 {
                                                                                                v28[0] = (fi27 << ((uint(p1)) & 63))
                                                                                            }
                                                                                        case "hello":
                                                                                            var v65 uint32 = p0
                                                                                            _ = v65
                                                                                        }
                                                                                        _ = v64
                                                                                    } else {
                                                                                        v66 := GenS4{F0: v56, F1: 346, F2: 1000}
                                                                                        fmt.Println(v56, v63)
                                                                                        _ = v66
                                                                                    }
                                                                                    v67 := GenS4{F0: (v56 + (v56 + v56)), F1: 0, F2: 1}
                                                                                    _ = v63
                                                                                    _ = v67
                                                                                }
                                                                            }
                                                                            var v68 string = v31
                                                                            _ = v68
                                                                        }
                                                                        v28 = append(v28, 9223372036854775807)
                                                                        var v69 uint32 = (p0 % ((65535) | 1))
                                                                        _ = v69
                                                                    }
                                                                    fmt.Println(v50, v49)
                                                                    fmt.Println(v31)
                                                                }
                                                            }
                                                            v56 -= 2
                                                        }
                                                        if (true || (42 == (p2 << ((4294967295) & 63)))) {
                                                            fmt.Println(v51, v51)
                                                            var v70 int8 = v50
                                                            v51 += float64((((((((1.5 / ((1e10 + (1e10 / 0.1)) * 0.1)) - 0.1) / 1.0) * (0.0 - (0.0 + -2.25))) * 0.0) / 3.14159) / 1.0))
                                                            _ = v70
                                                        } else {
                                                            v71 := GenS6{F0: p0, F1: v50}
                                                            if true {
                                                                var v72 float32 = float32(0.0)
                                                                switch 9223372036854775807 {
                                                                case 4294967295:
                                                                    var v73 uint32 = 753
                                                                    _ = v73
                                                                default:
                                                                    v74 := GenS6{F0: (p0 * p0), F1: v50}
                                                                    _ = v74
                                                                }
                                                                _ = v72
                                                            }
                                                            _ = v71
                                                        }
                                                        v45 -= (v45 + (v45 >> (((v56 ^ (v56 << ((v56) & 63)))) & 63)))
                                                        _ = v56
                                                    }
                                                }
                                            } else {
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                            }
                                        default:
                                            if (false && false) {
                                                p1 = p1
                                            } else {
                                                if false {
                                                    var v75 float64 = v51
                                                    switch v31 {
                                                    case "\n\t":
                                                        fmt.Println(v49, fi30, p0)
                                                    default:
                                                        var v76 int32 = v45
                                                        _ = v76
                                                    }
                                                    _ = v75
                                                } else {
                                                    var v77 int32 = v45
                                                    if len(v28) >= 2 {
                                                        v28 = v28[1:len(v28)]
                                                    }
                                                    _ = v77
                                                }
                                                var v78 uint8 = (p2 / ((135) | 1))
                                                fmt.Println(v45, v49)
                                                v79 := GenS3{F0: p0, F1: 124}
                                                _ = v78
                                                _ = v79
                                            }
                                        }
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    case 4294967295:
                                        switch (^(fi27)) {
                                        case 1:
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                        case 32767:
                                            var v80 uint32 = (p0 + 2147483648)
                                            _ = v80
                                        default:
                                            v81 := GenS7{F0: (fi30 | (fi30 ^ (fi27 * fi30))), F1: p1, F2: (p2 % ((1) | 1))}
                                            _ = v81
                                        }
                                    default:
                                        v28 = append(v28, fi30)
                                    }
                                    var v82 int8 = 62
                                    _ = v82
                                }
                            }
                            switch "hello" {
                            case "a":
                                p1 = p1
                            case "\n\t":
                                var v83 uint = 18446744073709551615
                                _ = v83
                            case "foo bar":
                                v84 := GenS3{F0: (^(p0)), F1: 42}
                                _ = v84
                            default:
                                var v85 string = (v31 + ("\n\t" + "hello"))
                                _ = v85
                            }
                            p1 = (p1 | 42)
                            _ = v45
                        }
                    case "\n\t":
                        var v86 bool = (false || false)
                        _ = v86
                    default:
                        var v87 int8 = 106
                        _ = v87
                    }
                default:
                    p2 = 76
                    if (4294967295 > 537) {
                        v31 = ("hello" + v31)
                    } else {
                        fmt.Println(v31)
                    }
                    v28 = append(v28, (fi30 - fi27))
                }
                var v88 int32 = 42
                _ = fi30
                _ = v31
                _ = v88
            }
        }
        p0 += p0
        _ = fi27
        _ = v28
    }
    var v89 int = -9223372036854775808
    if ((p0 != 425) || true) {
        fmt.Println(p1)
        var v90 uint16 = 42
        fmt.Println(p0)
        _ = v90
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v89
    return 9223372036854775807
}

func main() {
    var v1 int64 = 42
    var v2 int16 = 157
    var v3 int64 = (v1 << ((373) & 63))
    var v4 bool = (-2.25 >= 1e10)
    v91 := []int{1}
    for fi92 := 0; fi92 < 12; fi92++ {
        if (false && (668 == -813)) {
            var v93 int32 = genFn0(int(v2), 42, (fi92 - (fi92 ^ 127)))
            v1 += genFn1(0, 1, 0)
            var v94 int64 = 42
            _ = v93
            _ = v94
        } else {
            switch (fi92 * 1) {
            case 1:
                switch fi92 {
                case 1:
                    for fi95 := 0; fi95 < 5; fi95++ {
                        if v4 {
                            var v96 uint = 2
                            var v97 uint8 = 2
                            var v98 int8 = 1
                            _ = v96
                            _ = v97
                            _ = v98
                        } else {
                            v99 := GenS4{F0: 2, F1: v1, F2: (-(v2))}
                            if false {
                                v91 = append(v91, fi95)
                            } else {
                                if true {
                                    v100 := GenS2{F0: 2}
                                    v4 = v4
                                    _ = v100
                                }
                                if len(v91) >= 2 {
                                    v91 = v91[1:len(v91)]
                                }
                            }
                            switch fi95 {
                            case 32768:
                                if true {
                                    fmt.Println(v4)
                                    if (!(-1 == 83)) {
                                        fmt.Println(v3)
                                        v4 = (0.1 > 1e10)
                                        if (2 != 9) {
                                            fmt.Println(fi95)
                                            v3 = v3
                                            fmt.Println(v3)
                                        }
                                    }
                                    _ = len(v91)
                                } else {
                                    v91 = append(v91, (fi92 ^ (fi92 / ((fi92) | 1))))
                                }
                            case 127:
                                switch 2147483647 {
                                case 642:
                                    _ = len(v91)
                                }
                            default:
                                var v101 float64 = 0.0
                                _ = v101
                            }
                            if v4 {
                                v2 += int16(fi95)
                                _ = len(v91)
                                v4 = false
                                v2 += (v2 - (^(v2)))
                                v1 = (v1 % ((-3) | 1))
                            }
                            v102 := GenS7{F0: 0, F1: 65535, F2: 42}
                            _ = v102
                            _ = v99
                        }
                        v1 = genFn1(42, 32767, 0)
                        var v103 float32 = 1e10
                        v103 -= (float32(0.1) - (0.0 + v103))
                        var v104 int32 = 0
                        fmt.Println(fi95, v4, v104)
                        _ = fi95
                        _ = v103
                        _ = v104
                    }
                default:
                    if (fi92 <= (fi92 % (((fi92 - fi92)) | 1))) {
                        var v105 uint32 = 2147483647
                        v1 = (^(v1))
                        _ = v105
                    } else {
                        v106 := GenS8{F0: 2, F1: (v3 & (^(v3)))}
                        for fi107 := 0; fi107 < 3; fi107++ {
                            v3 = genFn1(65535, 32767, 2)
                            _ = fi107
                        }
                        v3 -= (v1 >> ((4294967295) & 63))
                        _ = v106
                    }
                }
            case 42:
                for fi108 := 0; fi108 < 2; fi108++ {
                    v109 := GenS4{F0: 4294967295, F1: (v1 | 0), F2: -32768}
                    if len(v91) > 0 {
                        v91[0] = fi92
                    }
                    var v110 int64 = v3
                    _ = fi108
                    _ = v109
                    _ = v110
                }
            default:
                v2 = 996
            }
            if (2 < 2) {
                for fi111 := 0; fi111 < 11; fi111++ {
                    switch "a" {
                    case "a":
                        v3 = v1
                    }
                    _ = fi111
                }
            } else {
                if len(v91) > 0 {
                    v91[0] = -1
                }
                for fi112 := 0; fi112 < 16; fi112++ {
                    var v113 float32 = -2.25
                    v114 := GenS6{F0: 65535, F1: 127}
                    v1 = (v1 / ((v1) | 1))
                    _ = fi112
                    _ = v113
                    _ = v114
                }
            }
            fmt.Println(fi92, v1, v2)
        }
        v1 = v1
        var v115 uint16 = 65535
        _ = fi92
        _ = v115
    }
    var v116 int8 = 42
    var v117 uint32 = 0
    v91 = append(v91, 1)
    v118 := GenS9{F0: v116}
    for fi119 := 0; fi119 < 5; fi119++ {
        var v120 uint64 = 18446744073709551615
        var v121 bool = false
        _ = fi119
        _ = v120
        _ = v121
    }
    v116 = (v116 / (((v116 - v116)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v91:", len(v91), v91)
    _ = v1
    _ = v116
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v91
}
