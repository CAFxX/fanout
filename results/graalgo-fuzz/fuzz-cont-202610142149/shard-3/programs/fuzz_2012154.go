package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 string; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS9 struct { F0 int; F1 uint8; F2 uint64 }
type GenS10 struct { F0 int64; F1 int16; F2 int8 }
type GenS11 struct { F0 uint8; F1 int64 }
type GenS12 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS13 struct { F0 int32; F1 string }
type GenS14 struct { F0 int8; F1 string; F2 int8 }
type GenS15 struct { F0 uint32; F1 bool; F2 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    var v4 uint32 = 2
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int16 = 1
        v8 := GenS1{F0: -9223372036854775808, F1: 282}
        if false {
            fmt.Println(v7, v3)
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = len(v2)
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v3, fi9)
        _ = fi9
    }
    var v10 int32 = -1
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 14; fi12++ {
        if false {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int64 = 64
                v10 -= (v10 % (((^(v10))) | 1))
                fmt.Println(v4, v14, v10)
                _ = fi13
                _ = v14
            }
            fmt.Println(v11, v3, fi12)
            var v15 uint = 255
            _ = v15
        }
        var v16 uint = 610
        if (-1 <= -128) {
            var v17 uint8 = 233
            if ((^(v10)) == 42) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (v4 != 2147483647) {
                        delete(v1, (v11 + (v11 + (v11 + ""))))
                    } else {
                        if false {
                            v5[uint32(v16)] = 883
                        }
                        v19 := []int{(fi12 | int(v4)), (fi18 * -2147483648), fi18}
                        if len(v19) > 0 {
                            v19[0] = fi18
                        }
                        _ = v19
                    }
                    v10 += (v10 >> (((v16 ^ 140)) & 63))
                    if true {
                        switch v11 {
                        case "":
                            v16 -= (^(v16))
                        case "\n\t":
                            if (true || (962 == v4)) {
                                if (103 == -122) {
                                    switch v11 {
                                    case "a":
                                        v20 := GenS1{F0: -1, F1: v16}
                                        _ = v20
                                    case "hello":
                                        v21 := []int{1, (fi18 | fi18), (fi18 ^ (fi18 ^ fi12))}
                                        _ = v21
                                    case "\n\t":
                                        if false {
                                            var v22 float32 = 0.1
                                            var v23 int16 = int16(v17)
                                            var v24 uint32 = v4
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        } else {
                                            switch (fi18 ^ fi18) {
                                            case 972:
                                                v17 += (v17 / ((172) | 1))
                                            case -1:
                                                v25 := GenS2{F0: (^(v17))}
                                                _ = v25
                                            case 127:
                                                v10 = -2147483648
                                            }
                                            delete(v1, v11)
                                            v3 -= (v3 + 4294967295)
                                        }
                                    default:
                                        v17 = (v17 << ((v16) & 63))
                                    }
                                    var v26 float64 = 0.1
                                    v26 = float64(1e10)
                                    v27 := GenS1{F0: -9223372036854775808, F1: v16}
                                    v28 := GenS2{F0: v17}
                                    if false {
                                        v29 := []int{}
                                        var v30 int = fi18
                                        v31 := GenS2{F0: v17}
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 bool = true
                                        _ = v32
                                    }
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    fmt.Println(v4)
                                    v33 := GenS3{F0: -128, F1: (v17 % ((v17) | 1))}
                                    _ = v33
                                }
                                _ = len(v1)
                                var v34 uint = v16
                                fmt.Println(v17, v34, v10)
                                _ = v34
                            }
                        default:
                            v35 := GenS4{F0: 0, F1: 127, F2: (v17 ^ (v17 >> (((v16 + 255)) & 63)))}
                            _ = v35
                        }
                        if true {
                            v36 := []int{fi18}
                            var v37 int16 = 0
                            v10 += v10
                            var v38 uint = 1
                            var v39 uint16 = 1
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    }
                    fmt.Println(fi18, v10)
                    v11 = v11
                    _ = fi18
                }
            } else {
                for fi40 := 0; fi40 < 18; fi40++ {
                    v4 = (v4 + (v4 + (v4 * (v4 << (((v16 - (v16 ^ (v16 ^ v16)))) & 63)))))
                    var v41 uint64 = (^(v3))
                    if ((v41 + 0) > 2) {
                        var v42 uint = (v16 ^ v16)
                        _ = v42
                    }
                    _ = fi40
                    _ = v41
                }
                if true {
                    if ((v3 >> ((v16) & 63)) < v3) {
                        v4 = (v4 % (((v4 - uint32(v3))) | 1))
                        fmt.Println(v11, v3, v11)
                    }
                    if false {
                        switch ("a" + (v11 + (v11 + "foo bar"))) {
                        case "\n\t":
                            switch (("a" + ((("" + v11) + v11) + v11)) + "\n\t") {
                            case "\n\t":
                                switch "a" {
                                case "hello":
                                    switch fi12 {
                                    case 341:
                                        var v43 bool = ("foo bar" >= (v11 + "\n\t"))
                                        _ = v43
                                    }
                                case "foo bar":
                                    v44 := GenS1{F0: -862, F1: (v16 << ((931) & 63))}
                                    _ = v44
                                case "a":
                                    switch -9223372036854775808 {
                                    case -9223372036854775808:
                                        if ((uint8(v3) >= uint8(v4)) || true) {
                                            for fi45 := 0; fi45 < 9; fi45++ {
                                                v46 := []int{(fi45 & (fi12 ^ fi12)), fi45, 127, (fi45 << (((v16 >> ((1) & 63))) & 63))}
                                                var v47 string = ((("" + (v11 + v11)) + ("foo bar" + ((v11 + ("\n\t" + v11)) + ((v11 + v11) + v11)))) + "hello")
                                                v17 = (v17 * v17)
                                                _ = fi45
                                                _ = v46
                                                _ = v47
                                            }
                                            if (!(v17 > v17)) {
                                                v48 := GenS4{F0: uint16(v3), F1: 1, F2: uint8(v16)}
                                                switch (fi12 | 128) {
                                                case 812:
                                                    fmt.Println(v4)
                                                case 42:
                                                    v16 = v16
                                                case -591:
                                                    for fi49 := 0; fi49 < 10; fi49++ {
                                                        switch v11 {
                                                        case "a":
                                                            v17 = (v17 - v17)
                                                        case "foo bar":
                                                            switch v11 {
                                                            case "":
                                                                var v50 int32 = 798
                                                                _ = v50
                                                            case "\n\t":
                                                                var v51 int32 = int32(v16)
                                                                _ = v51
                                                            case "foo bar":
                                                                var v52 int32 = -1
                                                                if (702 > 0) {
                                                                    fmt.Println(v10)
                                                                    var v53 uint16 = 65535
                                                                    v54 := GenS5{F0: -28}
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                if false {
                                                                    switch (v11 + "hello") {
                                                                    case "\n\t":
                                                                        if ((false || true) || (v16 < v16)) {
                                                                            var v55 uint64 = (v3 - (v3 % (((v3 << ((18446744073709551615) & 63))) | 1)))
                                                                            var v56 uint64 = v3
                                                                            if (int16(fi12) != -1) {
                                                                                fmt.Println(v16)
                                                                                v55 = (v3 >> ((281) & 63))
                                                                                switch (-(fi49)) {
                                                                                case 127:
                                                                                    var v57 string = v11
                                                                                    _ = v57
                                                                                default:
                                                                                    var v58 uint16 = 482
                                                                                    _ = v58
                                                                                }
                                                                            }
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                    case "hello":
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v59 bool = false
                                                                            _ = v59
                                                                        case "":
                                                                            var v60 int64 = 9223372036854775807
                                                                            _ = v60
                                                                        }
                                                                    default:
                                                                        var v61 uint32 = v4
                                                                        _ = v61
                                                                    }
                                                                    var v62 int16 = 1000
                                                                    v3 = 1
                                                                    _ = v62
                                                                }
                                                                v63 := GenS1{F0: 0, F1: (v16 - (v16 & v16))}
                                                                var v64 int64 = 9223372036854775807
                                                                v4 = 733
                                                                _ = v52
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                        default:
                                                            v65 := []int{fi49, fi12}
                                                            _ = v65
                                                        }
                                                        _ = fi49
                                                    }
                                                }
                                                fmt.Println(v3)
                                                _ = v48
                                            } else {
                                                for fi66 := 0; fi66 < 2; fi66++ {
                                                    v67 := GenS5{F0: 584}
                                                    var v68 int32 = 1
                                                    _ = fi66
                                                    _ = v67
                                                    _ = v68
                                                }
                                                var v69 int16 = -364
                                                var v70 int64 = 1
                                                _ = v69
                                                _ = v70
                                            }
                                            var v71 float64 = float64(-2.25)
                                            _ = v71
                                        } else {
                                            var v72 int64 = -1
                                            var v73 bool = true
                                            v4 = (^(v4))
                                            _ = v72
                                            _ = v73
                                        }
                                    case 73:
                                        switch (fi12 % (((fi12 * fi12)) | 1)) {
                                        case 128:
                                            v4 = (v4 & (v4 & v4))
                                        case -2147483648:
                                            v16 = (^(v16))
                                        case -185:
                                            v74 := GenS3{F0: -12, F1: v17}
                                            _ = v74
                                        }
                                    }
                                default:
                                    v75 := []int{fi12, 227, (-(fi12)), fi12}
                                    _ = v75
                                }
                            case "hello":
                                v17 -= (v17 / (((v17 >> ((1) & 63))) | 1))
                            }
                        case "a":
                            for fi76 := 0; fi76 < 13; fi76++ {
                                var v77 uint64 = (v3 << ((v16) & 63))
                                switch 128 {
                                case 1:
                                    v78 := GenS1{F0: 1, F1: (v16 << ((65535) & 63))}
                                    _ = v78
                                case -2147483648:
                                    delete(v2, v4)
                                case 9223372036854775807:
                                    var v79 uint32 = v4
                                    _ = v79
                                }
                                var v80 uint64 = 18446744073709551615
                                v4 -= v4
                                v81 := []int{fi12}
                                _ = fi76
                                _ = v77
                                _ = v80
                                _ = v81
                            }
                        default:
                            var v82 bool = (true && false)
                            _ = v82
                        }
                    }
                    for fi83 := 0; fi83 < 4; fi83++ {
                        fmt.Println(fi83)
                        if false {
                            var v84 int16 = 461
                            fmt.Println(v4)
                            var v85 string = ("" + v11)
                            _ = v84
                            _ = v85
                        } else {
                            if (100 >= -119) {
                                v86 := GenS6{F0: 846, F1: (v11 + v11), F2: (v16 + (v16 - v16))}
                                _ = v86
                            }
                            var v87 int8 = 100
                            var v88 int64 = -946
                            _ = v87
                            _ = v88
                        }
                        v11 = ""
                        var v89 int16 = 42
                        fmt.Println(v4, fi83)
                        _ = fi83
                        _ = v89
                    }
                } else {
                    fmt.Println(v16, v17)
                    v17 += (v17 | 2)
                    fmt.Println(v11, v3, v11)
                }
                if true {
                    v90 := GenS7{F0: 127}
                    _ = v90
                } else {
                    switch ((v11 + ((v11 + v11) + v11)) + ("a" + (((((v11 + v11) + v11) + v11) + (((v11 + ("foo bar" + ((v11 + ((v11 + "\n\t") + v11)) + v11))) + ("a" + ((v11 + v11) + "foo bar"))) + "hello")) + "hello"))) {
                    case "hello":
                        for fi91 := 0; fi91 < 18; fi91++ {
                            if (!true) {
                                v11 = (v11 + "hello")
                                switch (v11 + v11) {
                                case "hello":
                                    var v92 string = v11
                                    _ = v92
                                case "a":
                                    v93 := GenS8{F0: 558, F1: 1, F2: v17}
                                    _ = v93
                                case "foo bar":
                                    var v94 string = (v11 + (v11 + v11))
                                    _ = v94
                                default:
                                    v95 := GenS9{F0: (fi91 - fi12), F1: v17, F2: v3}
                                    _ = v95
                                }
                            } else {
                                switch (fi12 | 858) {
                                case 9223372036854775807:
                                    var v96 bool = false
                                    _ = v96
                                case -91:
                                    if true {
                                        var v97 int8 = -10
                                        var v98 uint8 = v17
                                        _ = v97
                                        _ = v98
                                    }
                                case 128:
                                    if ((0.0 != 0.0) && true) {
                                        v99 := []int{0, fi12, (fi12 << ((0) & 63))}
                                        _ = v99
                                    } else {
                                        var v100 uint = (v16 ^ 65535)
                                        _ = v100
                                    }
                                }
                                v101 := GenS10{F0: int64(fi12), F1: -32768, F2: 13}
                                _ = v101
                            }
                            delete(v1, "hello")
                            _ = fi91
                        }
                    case "\n\t":
                        v102 := GenS8{F0: (v4 & 42), F1: v17, F2: (v17 + 42)}
                        _ = v102
                    case "foo bar":
                        if (!false) {
                            var v103 uint64 = v3
                            _ = v103
                        } else {
                            var v104 int = fi12
                            for fi105 := 0; fi105 < 20; fi105++ {
                                fmt.Println(v10, v16)
                                if (42 > 1) {
                                    if true {
                                        v106 := GenS11{F0: uint8(v16), F1: -1}
                                        v104 = (fi12 % (((fi105 << (((v16 << ((165) & 63))) & 63))) | 1))
                                        _ = v106
                                    }
                                    var v107 int32 = 42
                                    _ = v107
                                } else {
                                    var v108 uint = v16
                                    _ = v108
                                }
                                _ = fi105
                            }
                            var v109 uint32 = v4
                            _ = v104
                            _ = v109
                        }
                    default:
                        v110 := v1["\n\t"]
                        _ = v110
                    }
                }
            }
            _ = v17
        } else {
            switch "a" {
            case "foo bar":
                v111 := GenS12{F0: 2, F1: 877, F2: 1}
                _ = v111
            case "":
                if (704 < 32767) {
                    var v112 uint32 = (^(v4))
                    switch fi12 {
                    case 0:
                        v113 := GenS13{F0: (-(v10)), F1: v11}
                        _ = v113
                    case 9223372036854775807:
                        for fi114 := 0; fi114 < 9; fi114++ {
                            fmt.Println(v4, v3, fi12)
                            _ = fi114
                        }
                    default:
                        fmt.Println(v112, v11, v4)
                    }
                    _ = v112
                }
            case "\n\t":
                v115 := GenS14{F0: int8(v3), F1: v11, F2: int8(v4)}
                _ = v115
            }
            fmt.Println(v16, v16, v4)
        }
        _ = fi12
        _ = v16
    }
    var v116 string = v11
    var v117 bool = true
    if true {
        var v118 int16 = 32767
        v119 := []int{0, 127, 32767, -1}
        switch 32767 {
        case 9223372036854775807:
            v120 := GenS15{F0: (v4 >> ((2) & 63)), F1: v117, F2: 0}
            _ = v120
        }
        _ = v118
        _ = v119
    } else {
        var v121 uint8 = 42
        v122 := GenS8{F0: 2147483648, F1: 0, F2: (v121 & 2)}
        v123 := GenS3{F0: 127, F1: (v121 << ((4294967295) & 63))}
        _ = v121
        _ = v122
        _ = v123
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
