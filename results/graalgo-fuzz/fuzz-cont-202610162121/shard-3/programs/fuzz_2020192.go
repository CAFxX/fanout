package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 string; F1 int8; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint; F2 uint }
type GenS6 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 string; F1 bool; F2 int64 }
type GenS9 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 uint8; F1 uint }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint; F1 uint8; F2 bool }
func genFn0(p0 bool, p1 int, p2 int) uint {
    p1 = (p2 + 32768)
    var v5 uint64 = 42
    v5 += v5
    var v6 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 42
    var v3 int16 = -794
    v4 := []string{"a"}
    v7 := GenS1{F0: 438, F1: -128}
    var v8 uint8 = uint8(v2)
    v8 -= (v8 - (v8 * v8))
    v2 = v2
    v3 += 1
    switch "" {
    case "hello":
        v3 = (v3 | v3)
    }
    if true {
        if (-1 < 95) {
            var v9 uint = 904
            _ = v9
        }
    }
    var v10 uint16 = (^(v2))
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int16 = v3
        if false {
            for fi13 := 0; fi13 < 6; fi13++ {
                fmt.Println(v3)
                var v14 uint = 4294967295
                v15 := GenS1{F0: 9223372036854775807, F1: int8(v2)}
                var v16 bool = false
                v3 += (v12 >> (((v14 >> ((831) & 63))) & 63))
                switch "" {
                case "foo bar":
                    v2 = (v2 % ((402) | 1))
                case "\n\t":
                    if ((v10 / ((v10) | 1)) <= (v10 ^ (^(v2)))) {
                        if (42 == 1) {
                            if (false || (!(((!(1 > 1)) || ((v8 | 231) <= (^(v8)))) && (-2.25 >= 0.1)))) {
                                if v16 {
                                    v3 -= (v3 - v3)
                                    if (v10 != 65535) {
                                        switch (fi13 - fi11) {
                                        case 4294967295:
                                            if ((-18 != int64(v12)) || (2147483647 == 782)) {
                                                switch (fi13 / ((128) | 1)) {
                                                case 128:
                                                    fmt.Println(v14, v12, v3)
                                                default:
                                                    var v17 uint = (^(v14))
                                                    _ = v17
                                                }
                                            }
                                        case -1:
                                            if (!v16) {
                                                v18 := GenS1{F0: -1, F1: 100}
                                                switch (fi13 | 0) {
                                                case -9223372036854775808:
                                                    if true {
                                                        _ = len(v4)
                                                        switch (fi13 | (^(fi11))) {
                                                        case 32767:
                                                            v12 = v3
                                                            var v19 uint16 = (^(v2))
                                                            _ = v19
                                                        case 9223372036854775807:
                                                            var v20 int = fi13
                                                            _ = v20
                                                        case 0:
                                                            v2 += v10
                                                        default:
                                                            v3 -= v12
                                                        }
                                                        v8 = v8
                                                    } else {
                                                        var v21 int64 = -1
                                                        var v22 string = ("" + "hello")
                                                        if (v16 && ((v21 % (((-(v21))) | 1)) == v21)) {
                                                            var v23 uint = 50
                                                            _ = v23
                                                        } else {
                                                            var v24 int32 = -985
                                                            v12 -= int16(v21)
                                                            switch fi13 {
                                                            case -2147483648:
                                                                switch fi13 {
                                                                case 42:
                                                                    fmt.Println(fi13, fi11)
                                                                    switch fi11 {
                                                                    case 9223372036854775807:
                                                                        var v25 bool = v16
                                                                        _ = v25
                                                                    case 42:
                                                                        v26 := v1[722]
                                                                        _ = v26
                                                                    case -101:
                                                                        var v27 uint32 = 60
                                                                        v28 := GenS1{F0: v21, F1: -67}
                                                                        _ = v27
                                                                        _ = v28
                                                                    }
                                                                    var v29 int64 = (^(v21))
                                                                    _ = v29
                                                                case 2147483647:
                                                                    var v30 int64 = v21
                                                                    _ = v30
                                                                case 32768:
                                                                    var v31 int32 = v24
                                                                    _ = v31
                                                                }
                                                                fmt.Println(v22, fi13)
                                                                if (0.0 > 0.0) {
                                                                    v3 -= (v3 / ((int16(v8)) | 1))
                                                                    if v16 {
                                                                        fmt.Println(v12, v8, v12)
                                                                        v2 = 0
                                                                    }
                                                                    v2 = 795
                                                                    if (((float32(float64(((0.0 - float32(0.1)) / 1e10))) * 0.0) + -2.25) >= 1e10) {
                                                                        v24 += 0
                                                                        v12 = (-(v3))
                                                                    } else {
                                                                        if (v16 && (1e10 >= 0.0)) {
                                                                            if v16 {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            } else {
                                                                                v10 = v2
                                                                                v22 = "foo bar"
                                                                            }
                                                                            var v32 int16 = 615
                                                                            switch v22 {
                                                                            case "foo bar":
                                                                                switch ("hello" + "") {
                                                                                case "hello":
                                                                                    if v16 {
                                                                                        if (v14 <= 2) {
                                                                                            v33 := GenS1{F0: (v21 * 0), F1: int8(v21)}
                                                                                            var v34 uint64 = 354
                                                                                            v12 += (v3 + (^(v12)))
                                                                                            var v35 bool = (-2.25 < 0.1)
                                                                                            _ = v33
                                                                                            _ = v34
                                                                                            _ = v35
                                                                                        }
                                                                                        var v36 uint32 = 575
                                                                                        var v37 int8 = int8(v2)
                                                                                        _ = v36
                                                                                        _ = v37
                                                                                    }
                                                                                case "a":
                                                                                    fmt.Println(v24, fi13)
                                                                                case "\n\t":
                                                                                    v38 := GenS1{F0: (v21 % (((v21 + (^(v21)))) | 1)), F1: 127}
                                                                                    _ = v38
                                                                                }
                                                                            case "hello":
                                                                                var v39 int = fi11
                                                                                _ = v39
                                                                            case "":
                                                                                v40 := GenS2{F0: (-(v24)), F1: 903}
                                                                                var v41 uint16 = v2
                                                                                var v42 int32 = (v24 + v24)
                                                                                switch fi11 {
                                                                                case 32768:
                                                                                    var v43 int8 = 127
                                                                                    _ = v43
                                                                                case 1:
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = "a"
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v10, v22, v41)
                                                                                    fmt.Println(v2, v2, v3)
                                                                                    v12 = 42
                                                                                    if v16 {
                                                                                        var v44 float64 = (-2.25 - 1.5)
                                                                                        var v45 int = fi13
                                                                                        if v16 {
                                                                                            _ = len(v1)
                                                                                            v46 := GenS3{F0: "", F1: 84, F2: 0}
                                                                                            v47 := GenS4{F0: genFn0(((v22 + v22) < v22), 32768, v45)}
                                                                                            _ = v46
                                                                                            _ = v47
                                                                                        }
                                                                                        _ = v44
                                                                                        _ = v45
                                                                                    } else {
                                                                                        switch (fi13 >> (((^(v14))) & 63)) {
                                                                                        case 127:
                                                                                            switch -9223372036854775808 {
                                                                                            case -9223372036854775808:
                                                                                                v48 := GenS5{F0: 42, F1: (v14 << ((genFn0(false, fi13, 334)) & 63)), F2: v14}
                                                                                                _ = v48
                                                                                            case -2147483648:
                                                                                                fmt.Println(v21, v21, fi11)
                                                                                                v49 := GenS4{F0: uint(v41)}
                                                                                                v32 = 5
                                                                                                _ = v49
                                                                                            }
                                                                                        default:
                                                                                            var v50 int8 = int8(v32)
                                                                                            _ = v50
                                                                                        }
                                                                                        v2 = v2
                                                                                    }
                                                                                    v51 := v1[230]
                                                                                    _ = v51
                                                                                }
                                                                                var v52 bool = (1e10 > 1.5)
                                                                                _ = v40
                                                                                _ = v41
                                                                                _ = v42
                                                                                _ = v52
                                                                            default:
                                                                                if false {
                                                                                    if false {
                                                                                        var v53 float64 = 1e10
                                                                                        var v54 uint = (v14 * 1)
                                                                                        var v55 float64 = float64((1.5 - 1e10))
                                                                                        var v56 bool = v16
                                                                                        var v57 int64 = int64(v8)
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                    }
                                                                                    v3 = v32
                                                                                    v22 = (v22 + "a")
                                                                                } else {
                                                                                    delete(v1, 981)
                                                                                }
                                                                            }
                                                                            _ = v32
                                                                        } else {
                                                                            var v58 bool = (0.0 != (0.0 / 3.14159))
                                                                            delete(v1, 42)
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                }
                                                            case 1:
                                                                _ = len(v1)
                                                            }
                                                            if v16 {
                                                                v59 := v1[640]
                                                                _ = len(v1)
                                                                _ = v59
                                                            }
                                                            _ = v24
                                                        }
                                                        if false {
                                                            var v60 bool = v16
                                                            v61 := GenS3{F0: v22, F1: 1, F2: int16(v21)}
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            fmt.Println(v3, v2, fi11)
                                                            v12 += (v3 + (v12 << ((v14) & 63)))
                                                            var v62 uint16 = (v10 / (((v2 - (v2 / ((v10) | 1)))) | 1))
                                                            _ = v62
                                                        }
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                case 1:
                                                    v8 = v8
                                                case 2147483647:
                                                    if (v16 || false) {
                                                        if v16 {
                                                            v63 := GenS1{F0: 9223372036854775807, F1: -63}
                                                            _ = v63
                                                        } else {
                                                            v2 -= (v10 * v2)
                                                            var v64 float32 = -2.25
                                                            _ = v64
                                                        }
                                                    } else {
                                                        if v16 {
                                                            var v65 int8 = 0
                                                            switch fi11 {
                                                            case 2147483647:
                                                                v14 -= v14
                                                                v66 := GenS2{F0: int32(v2), F1: -913}
                                                                var v67 float32 = ((float32((3.14159 - 0.0)) / -2.25) - -2.25)
                                                                _ = v66
                                                                _ = v67
                                                            case -797:
                                                                var v68 uint64 = 1
                                                                _ = v68
                                                            default:
                                                                var v69 int16 = 904
                                                                _ = v69
                                                            }
                                                            var v70 int64 = 1
                                                            if (v16 || (((v14 ^ 42) != 987) && (!(42 == 2)))) {
                                                                switch (^(fi11)) {
                                                                case 1:
                                                                    if ((!false) || v16) {
                                                                        var v71 string = ""
                                                                        v70 = v70
                                                                        v65 = 127
                                                                        _ = v71
                                                                    } else {
                                                                        v8 -= (^(v8))
                                                                    }
                                                                }
                                                                v72 := v1[65535]
                                                                _ = v72
                                                            } else {
                                                                v73 := GenS6{F0: v14, F1: v8, F2: 420}
                                                                _ = len(v1)
                                                                _ = v73
                                                            }
                                                            if v16 {
                                                                v8 = v8
                                                                var v74 uint64 = 211
                                                                var v75 int16 = (v12 >> (((v14 >> ((255) & 63))) & 63))
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            _ = v65
                                                            _ = v70
                                                        } else {
                                                            fmt.Println(fi11)
                                                            var v76 string = "foo bar"
                                                            _ = v76
                                                        }
                                                    }
                                                }
                                                _ = v18
                                            } else {
                                                var v77 bool = v16
                                                switch (fi13 >> ((genFn0(v77, 746, fi13)) & 63)) {
                                                case 32768:
                                                    var v78 uint16 = v10
                                                    _ = v78
                                                case 0:
                                                    _ = len(v1)
                                                case 2147483647:
                                                    if ((!(127 == -27)) && v77) {
                                                        var v79 uint8 = (^(v8))
                                                        v80 := v1[2147483648]
                                                        v14 = (v14 | uint(v10))
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                }
                                                _ = v77
                                            }
                                        }
                                        var v81 uint = (v14 | 1)
                                        _ = v81
                                    } else {
                                        v82 := GenS7{F0: "hello", F1: v14}
                                        var v83 uint = (v14 ^ (v14 % ((genFn0((1.5 != ((0.1 * (1e10 * 0.0)) * ((3.14159 / 0.1) / 1.0))), (fi11 & (fi11 | -402)), (fi13 ^ -9223372036854775808))) | 1)))
                                        _ = v82
                                        _ = v83
                                    }
                                    var v84 bool = (2 != (^(v10)))
                                    _ = v84
                                } else {
                                    v85 := GenS8{F0: "a", F1: ((v14 >> ((132) & 63)) >= genFn0(((0.1 - ((0.0 * float64(-2.25)) + 0.0)) != (-2.25 / (1e10 + -2.25))), fi13, fi13)), F2: -680}
                                    var v86 string = ""
                                    v87 := GenS9{F0: v2, F1: -1, F2: 42}
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                }
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    v14 = v14
                                case "":
                                    v16 = (!(127 != 0))
                                }
                                if false {
                                    v88 := GenS10{F0: uint16(v8), F1: 18}
                                    var v89 uint = genFn0((-128 >= -1), (fi11 + fi11), fi11)
                                    var v90 int16 = v12
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                }
                            }
                            var v91 string = ""
                            _ = v91
                        } else {
                            var v92 int8 = -1
                            switch (fi11 / ((fi13) | 1)) {
                            case 0:
                                fmt.Println(v3)
                            }
                            _ = v92
                        }
                        var v93 int = fi11
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = v93
                    }
                default:
                    v94 := GenS11{F0: 2147483647}
                    _ = v94
                }
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        } else {
            var v95 uint8 = (v8 % ((v8) | 1))
            v96 := GenS12{F0: 560, F1: (v95 + v8)}
            _ = v95
            _ = v96
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi11
        _ = v12
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v97 := v1[739]
    v8 -= (^(v8))
    for fi98 := 0; fi98 < 17; fi98++ {
        if (v8 < v8) {
            for fi99 := 0; fi99 < 2; fi99++ {
                v100 := GenS13{F0: 127, F1: uint(v10)}
                var v101 int32 = 2147483647
                fmt.Println(v2, v101)
                _ = fi99
                _ = v100
                _ = v101
            }
        } else {
            var v102 float32 = (((1.5 / 3.14159) + (float32(1.5) + 0.0)) - 0.0)
            _ = v102
        }
        v1[v97] = v97
        _ = fi98
    }
    if true {
        delete(v1, v97)
    } else {
        fmt.Println(v8)
        v97 -= 2147483647
    }
    v8 = v8
    for fi103 := 0; fi103 < 13; fi103++ {
        for fi104 := 0; fi104 < 6; fi104++ {
            var v105 uint = genFn0(false, (fi104 * (fi104 | fi103)), fi103)
            if (!true) {
                fmt.Println(v10, v10, v10)
                var v106 string = "hello"
                _ = v106
            } else {
                v107 := GenS14{F0: 1}
                _ = v107
            }
            fmt.Println(v2)
            fmt.Println(v8, fi103)
            v3 += v3
            if (!(int64(v105) < 1)) {
                if (-2147483648 > 0) {
                    switch (("" + "a") + "foo bar") {
                    case "foo bar":
                        var v108 int8 = 0
                        _ = v108
                    case "hello":
                        v2 += (v10 | v2)
                    case "\n\t":
                        if false {
                            if (0 != 42) {
                                var v109 string = "hello"
                                _ = len(v4)
                                _ = v109
                            } else {
                                if (true || ((1 != 18446744073709551615) && (!true))) {
                                    v2 = (v10 >> (((v105 - (v105 ^ (v105 ^ (v105 / ((v105) | 1)))))) & 63))
                                } else {
                                    var v110 bool = true
                                    if (false && (!(float64(1.5) != (-2.25 / (float64((1.5 - -2.25)) + 1.5))))) {
                                        switch "" {
                                        case "hello":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    } else {
                                        v111 := GenS10{F0: (^(v10)), F1: v8}
                                        _ = v111
                                    }
                                    _ = v110
                                }
                                var v112 int8 = -91
                                v3 = v3
                                _ = v112
                            }
                            v113 := GenS15{F0: genFn0(true, fi104, -2147483648), F1: (^(v8)), F2: (!true)}
                            v2 = (v10 * (v2 << ((682) & 63)))
                            _ = v113
                        }
                    }
                    if (-1 <= (fi103 << ((4294967295) & 63))) {
                        _ = len(v4)
                    } else {
                        if (65535 < v2) {
                            var v114 uint = (v105 | v105)
                            v4 = append(v4, "a")
                            _ = v114
                        }
                    }
                    v4 = append(v4, "hello")
                } else {
                    v115 := GenS11{F0: 0}
                    _ = v115
                }
                var v116 int8 = 0
                var v117 uint32 = (v97 ^ uint32(v105))
                _ = v116
                _ = v117
            }
            _ = fi104
            _ = v105
        }
        if false {
            switch "a" {
            case "foo bar":
                v10 = (^(v2))
            }
            var v118 uint16 = (^(v2))
            _ = v118
        }
        var v119 int8 = -1
        _ = fi103
        _ = v119
    }
    fmt.Println(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v97
}
