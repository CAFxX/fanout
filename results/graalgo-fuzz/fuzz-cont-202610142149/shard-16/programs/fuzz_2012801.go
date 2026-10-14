package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int64; F1 uint8; F2 bool }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 int8 }
func genFn0(p0 int16, p1 string, p2 int8) uint16 {
    fmt.Println(p1, p1, p2)
    var v6 string = ""
    var v7 uint64 = 0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}
func genFn1() bool {
    var v8 string = "hello"
    if true {
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                v8 = v8
                var v11 uint16 = 65535
                _ = fi10
                _ = v11
            }
            v8 = "\n\t"
            fmt.Println(fi9, fi9)
            _ = fi9
        }
    }
    _ = v8
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 32767
    var v3 int16 = 1
    v4 := make(map[int]int)
    var v5 uint64 = 18446744073709551615
    var v12 uint16 = 1
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 bool = false
        if (-128 == -7) {
            var v15 bool = (-128 == -1)
            v16 := []int{(fi13 - int(v3))}
            _ = v15
            _ = v16
        } else {
            var v17 uint8 = 186
            v2 += fi13
            var v18 float64 = -2.25
            _ = v17
            _ = v18
        }
        if ("a" <= ("" + ("\n\t" + ""))) {
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 string = ""
                if (!v14) {
                    if (!v14) {
                        v12 -= v12
                        var v21 uint16 = (v12 | v12)
                        _ = v21
                    }
                    _ = len(v4)
                    v2 = 191
                } else {
                    var v22 uint16 = (v12 >> ((2) & 63))
                    fmt.Println(v2, fi13, v14)
                    var v23 uint8 = 0
                    switch "foo bar" {
                    case "a":
                        if true {
                            if (-2.25 >= 0.1) {
                                fmt.Println(v5, fi19)
                            } else {
                                v24 := []int{fi19, 2147483647}
                                switch "" {
                                case "a":
                                    var v25 uint32 = uint32(v12)
                                    _ = v25
                                case "":
                                    if v14 {
                                        v26 := GenS1{F0: 1, F1: uint8(v5)}
                                        var v27 uint = 18446744073709551615
                                        if (-2.25 <= 1e10) {
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                            switch fi13 {
                                            case 32768:
                                                var v28 int8 = 127
                                                _ = v28
                                            case 1:
                                                var v29 uint = (v27 & (v27 | (v27 | uint(v22))))
                                                _ = v29
                                            case 32767:
                                                v12 -= v22
                                            }
                                            v30 := GenS1{F0: 215, F1: v23}
                                            _ = v30
                                        } else {
                                            if (!(v14 || (!(!false)))) {
                                                v1[423] = ((v14 || ((fi13 % ((fi19) | 1)) < fi13)) || (2147483647 >= 42))
                                            } else {
                                                if false {
                                                    v3 = v3
                                                    var v31 uint64 = (v5 & 0)
                                                    _ = v31
                                                }
                                                switch (v20 + ("\n\t" + "foo bar")) {
                                                case "a":
                                                    fmt.Println(v20, v3)
                                                case "":
                                                    v32 := GenS2{F0: (v5 / ((v5) | 1)), F1: -974}
                                                    _ = v32
                                                }
                                                v14 = v14
                                            }
                                            v5 += v5
                                        }
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v33 := GenS3{F0: 65535, F1: (v3 / ((v3) | 1))}
                                        v5 = v5
                                        _ = v33
                                    }
                                case "hello":
                                    switch (v2 / ((-960) | 1)) {
                                    case -9223372036854775808:
                                        var v34 uint16 = genFn0(-32768, ((v20 + "foo bar") + ("a" + v20)), 42)
                                        _ = v34
                                    default:
                                        switch v20 {
                                        case "foo bar":
                                            fmt.Println(v14, v2)
                                        case "\n\t":
                                            fmt.Println(fi19, v12)
                                        default:
                                            v35 := GenS3{F0: 2147483647, F1: int16(v22)}
                                            _ = v35
                                        }
                                    }
                                default:
                                    v36 := GenS2{F0: v5, F1: 1}
                                    _ = v36
                                }
                                _ = v24
                            }
                            _ = len(v1)
                            var v37 uint64 = (v5 - (v5 | 0))
                            _ = v37
                        } else {
                            v38 := GenS1{F0: -1, F1: 128}
                            _ = v38
                        }
                    case "hello":
                        v39 := []int{fi13}
                        _ = v39
                    case "":
                        fmt.Println(v22, v14, fi19)
                    }
                    _ = v22
                    _ = v23
                }
                var v40 int8 = 42
                _ = fi19
                _ = v20
                _ = v40
            }
            fmt.Println(fi13, v12, v2)
            var v41 bool = ((float32(float64(0.0)) / 0.1) <= 0.0)
            _ = v41
        } else {
            var v42 float32 = ((float32((((float64(0.1) - (3.14159 * 1.5)) + 3.14159) + ((-2.25 / -2.25) * 1.5))) + -2.25) - 1.5)
            if (!(0 != 2)) {
                var v43 float64 = -2.25
                v44 := GenS1{F0: -9223372036854775808, F1: 2}
                fmt.Println(fi13)
                _ = v43
                _ = v44
            } else {
                if (1e10 >= 0.0) {
                    for fi45 := 0; fi45 < 3; fi45++ {
                        fmt.Println(v5)
                        v12 = (v12 - 27)
                        _ = fi45
                    }
                    var v46 int32 = 0
                    var v47 int = 2147483647
                    _ = v46
                    _ = v47
                }
            }
            _ = v42
        }
        _ = fi13
        _ = v14
    }
    var v48 int8 = -35
    var v49 string = "foo bar"
    if false {
        v48 = (-(v48))
        _ = len(v4)
    }
    var v50 int8 = v48
    var v51 int16 = 1000
    v12 += genFn0((v51 & (^(v51))), v49, (v48 - -128))
    for fi52 := 0; fi52 < 1; fi52++ {
        var v53 uint16 = (v12 & (v12 - v12))
        if (false || false) {
            v54 := GenS1{F0: 42, F1: 255}
            _ = v54
        } else {
            if ((v5 ^ (^(v5))) <= 715) {
                switch v49 {
                case "a":
                    for fi55 := 0; fi55 < 3; fi55++ {
                        v56 := GenS3{F0: 2147483647, F1: 1000}
                        v48 += -1
                        _ = fi55
                        _ = v56
                    }
                case "\n\t":
                    var v57 int64 = 42
                    _ = v57
                case "foo bar":
                    var v58 uint16 = v12
                    _ = v58
                default:
                    var v59 uint64 = (v5 - 2)
                    _ = v59
                }
                var v60 uint = 1
                _ = v60
            } else {
                if ((true || true) && (0 > 98)) {
                    v5 += 4294967295
                    var v61 int64 = -1
                    _ = v61
                }
                var v62 int8 = v50
                _ = v62
            }
            if (false && true) {
                v1[(v2 ^ 1)] = (1e10 < float64(1.5))
                v63 := GenS4{F0: (v5 % ((v5) | 1))}
                fmt.Println(v50, fi52)
                _ = v63
            }
        }
        _ = fi52
        _ = v53
    }
    fmt.Println(v50, v2)
    var v64 uint = 255
    fmt.Println(v5, v50, v51)
    switch v49 {
    case "foo bar":
        for fi65 := 0; fi65 < 2; fi65++ {
            for fi66 := 0; fi66 < 18; fi66++ {
                v67 := []int{}
                var v68 bool = ((("" + (v49 + (v49 + "\n\t"))) + v49) < v49)
                _ = fi66
                _ = v67
                _ = v68
            }
            _ = fi65
        }
    case "\n\t":
        switch (v2 ^ v2) {
        case 32768:
            v2 = v2
        case 637:
            for fi69 := 0; fi69 < 13; fi69++ {
                if true {
                    v70 := GenS5{F0: 571, F1: 111}
                    for fi71 := 0; fi71 < 7; fi71++ {
                        v2 = (fi71 / (((fi71 - fi71)) | 1))
                        _ = fi71
                    }
                    v72 := []int{-219}
                    _ = v70
                    _ = v72
                } else {
                    var v73 uint16 = genFn0(v51, "hello", v50)
                    var v74 int16 = v51
                    _ = v73
                    _ = v74
                }
                _ = fi69
            }
        }
    case "a":
        v50 -= v50
    default:
        for fi75 := 0; fi75 < 11; fi75++ {
            var v76 int32 = 572
            if ((v5 ^ v5) < v5) {
                var v77 float64 = float64((3.14159 - 0.0))
                var v78 uint16 = genFn0(42, v49, (v48 + v48))
                _ = v77
                _ = v78
            } else {
                v1[fi75] = (v5 <= v5)
                if (!true) {
                    if (42 <= 117) {
                        if (true || (!true)) {
                            fmt.Println(v12, v49)
                            v79 := GenS2{F0: v5, F1: 9223372036854775807}
                            for fi80 := 0; fi80 < 2; fi80++ {
                                fmt.Println(v51)
                                if (0 > (v50 & -1)) {
                                    v3 += v51
                                } else {
                                    if ((v64 * (v64 << ((172) & 63))) <= (v64 | (v64 - v64))) {
                                        v50 = (v48 * v50)
                                        var v81 uint8 = 26
                                        _ = v81
                                    } else {
                                        var v82 uint8 = 211
                                        if false {
                                            if (!(!(true || (!(2147483648 <= 2147483648))))) {
                                                var v83 int16 = (^(v51))
                                                v12 = (^(v12))
                                                v5 -= (v5 & (v5 & uint64(v82)))
                                                _ = v83
                                            }
                                        }
                                        if true {
                                            var v84 int16 = v3
                                            if false {
                                                if (false && (false || (42 < v50))) {
                                                    fmt.Println(v49)
                                                    switch int(v82) {
                                                    case 4294967295:
                                                        v85 := []int{(v2 << ((v64) & 63)), (fi80 | 32768), (v2 / ((32767) | 1))}
                                                        _ = v85
                                                    case -2147483648:
                                                        if true {
                                                            var v86 bool = (v82 < v82)
                                                            if (18446744073709551615 == 537) {
                                                                var v87 uint = uint(v12)
                                                                fmt.Println(v64, v48)
                                                                var v88 uint32 = 2147483647
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                            if (0.0 == 0.0) {
                                                                v89 := []int{(fi80 / (((fi75 / ((-9223372036854775808) | 1))) | 1)), (-(fi80))}
                                                                _ = v89
                                                            }
                                                            _ = v86
                                                        } else {
                                                            if false {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if ((!true) || (9223372036854775807 > fi80)) {
                                                                        if (!true) {
                                                                            switch v49 {
                                                                            case "a":
                                                                                if (true && (true || (!true))) {
                                                                                    switch 2147483647 {
                                                                                    case -1:
                                                                                        v49 = v49
                                                                                        var v90 bool = true
                                                                                        v91 := []int{}
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    case 910:
                                                                                        v64 = v64
                                                                                        var v92 uint64 = (v5 & v5)
                                                                                        _ = v92
                                                                                    }
                                                                                    v82 = (v82 * v82)
                                                                                } else {
                                                                                    delete(v4, 128)
                                                                                    v93 := GenS2{F0: v5, F1: 9223372036854775807}
                                                                                    if false {
                                                                                        v94 := []int{int(v82), (^(fi75))}
                                                                                        _ = v94
                                                                                    }
                                                                                    _ = v93
                                                                                }
                                                                            case "foo bar":
                                                                                v82 -= (v82 | v82)
                                                                            default:
                                                                                if (!true) {
                                                                                    if (int8(v3) < (v50 << ((v64) & 63))) {
                                                                                        var v95 int = 42
                                                                                        var v96 int16 = v3
                                                                                        switch v49 {
                                                                                        case "\n\t":
                                                                                            v97 := GenS3{F0: 959, F1: v51}
                                                                                            _ = v97
                                                                                        case "":
                                                                                            v98 := []int{v2}
                                                                                            _ = v98
                                                                                        }
                                                                                        _ = v95
                                                                                        _ = v96
                                                                                    } else {
                                                                                        var v99 int = (v2 << ((42) & 63))
                                                                                        var v100 int64 = -298
                                                                                        var v101 int = v2
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                        _ = v99
                                                                                    }
                                                                                    if false {
                                                                                        var v102 int8 = (v48 & (-(v48)))
                                                                                        v103 := GenS6{F0: -1, F1: (^(v82)), F2: (!(v48 <= (v50 & (v50 >> ((v64) & 63)))))}
                                                                                        if false {
                                                                                            var v104 uint16 = genFn0(-32768, v49, 127)
                                                                                            _ = v104
                                                                                        }
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                    }
                                                                                }
                                                                            }
                                                                            v105 := v1[2147483647]
                                                                            fmt.Println(v82, v48)
                                                                            _ = v105
                                                                        } else {
                                                                            if false {
                                                                                v51 = (v84 % ((v84) | 1))
                                                                                fmt.Println(v82, v5)
                                                                                var v106 uint32 = 65535
                                                                                v2 = fi80
                                                                                _ = v106
                                                                            } else {
                                                                                fmt.Println(v2, v84)
                                                                                v107 := []int{}
                                                                                _ = v107
                                                                            }
                                                                        }
                                                                        if true {
                                                                            var v108 uint8 = (^(v82))
                                                                            v4[465] = (fi75 ^ int(v108))
                                                                            _ = v108
                                                                        }
                                                                    } else {
                                                                        v76 += (v76 % (((v76 % ((v76) | 1))) | 1))
                                                                        v51 = v84
                                                                    }
                                                                case "a":
                                                                    v49 = (v49 + "a")
                                                                case "hello":
                                                                    var v109 bool = true
                                                                    _ = v109
                                                                }
                                                                fmt.Println(v3, v76, v49)
                                                            } else {
                                                                delete(v1, 0)
                                                            }
                                                            fmt.Println(v64, v50)
                                                        }
                                                        switch (v49 + "\n\t") {
                                                        case "\n\t":
                                                            if (!(((v5 - (v5 * uint64(v64))) <= (v5 + uint64(v12))) && false)) {
                                                                v110 := []int{}
                                                                fmt.Println(v49)
                                                                _ = v110
                                                            } else {
                                                                if (false || (9223372036854775807 >= 42)) {
                                                                    var v111 int8 = (^(v48))
                                                                    var v112 uint16 = v12
                                                                    _ = v111
                                                                    _ = v112
                                                                } else {
                                                                    v113 := GenS7{F0: (v2 >= (^(fi75))), F1: v49}
                                                                    if true {
                                                                        v5 += v5
                                                                    }
                                                                    var v114 int8 = (v48 + v50)
                                                                    _ = v113
                                                                    _ = v114
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v2 += (v2 << ((751) & 63))
                                                        case "":
                                                            v12 += v12
                                                        default:
                                                            var v115 int16 = (v3 % ((int16(v82)) | 1))
                                                            _ = v115
                                                        }
                                                        var v116 uint32 = 1
                                                        _ = v116
                                                    case 261:
                                                        var v117 int64 = -1
                                                        _ = v117
                                                    }
                                                } else {
                                                    switch v49 {
                                                    case "foo bar":
                                                        v48 += (v48 - (v50 ^ (v48 % ((-72) | 1))))
                                                    default:
                                                        v84 = v51
                                                    }
                                                    if (127 >= -95) {
                                                        v48 -= v50
                                                    } else {
                                                        v118 := GenS8{F0: (v64 / ((1) | 1)), F1: (1 >= 42)}
                                                        _ = v118
                                                    }
                                                }
                                                var v119 int64 = -9223372036854775808
                                                var v120 string = ""
                                                _ = v119
                                                _ = v120
                                            } else {
                                                fmt.Println(v5)
                                                var v121 int16 = v3
                                                _ = v121
                                            }
                                            _ = v84
                                        }
                                        fmt.Println(v64)
                                        _ = v82
                                    }
                                    fmt.Println(fi80, v2, v3)
                                    var v122 bool = false
                                    _ = v122
                                }
                                _ = fi80
                            }
                            _ = v79
                        } else {
                            for fi123 := 0; fi123 < 16; fi123++ {
                                v48 = v48
                                var v124 int64 = -1
                                _ = fi123
                                _ = v124
                            }
                            fmt.Println(v2, v3, v12)
                        }
                    } else {
                        var v125 int32 = v76
                        _ = v125
                    }
                    v2 = (^(fi75))
                }
                switch v49 {
                case "foo bar":
                    for fi126 := 0; fi126 < 5; fi126++ {
                        v51 += v3
                        if (521 != v64) {
                            v127 := GenS9{F0: -109}
                            v49 = (v49 + v49)
                            _ = v127
                        }
                        var v128 int8 = (-(v48))
                        _ = fi126
                        _ = v128
                    }
                case "":
                    if true {
                        for fi129 := 0; fi129 < 5; fi129++ {
                            v130 := GenS1{F0: 9223372036854775807, F1: 82}
                            fmt.Println(fi129, fi75)
                            _ = fi129
                            _ = v130
                        }
                    }
                }
            }
            var v131 uint = (^(v64))
            _ = fi75
            _ = v131
            _ = v76
        }
    }
    if true {
        v132 := GenS6{F0: -9223372036854775808, F1: 128, F2: false}
        var v133 uint8 = 0
        _ = v132
        _ = v133
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v64
}
