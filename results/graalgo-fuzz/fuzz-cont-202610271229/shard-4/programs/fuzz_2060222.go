package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 int64; F1 string; F2 bool }
type GenS9 struct { F0 uint16; F1 uint; F2 int }
type GenS10 struct { F0 string; F1 bool }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int64; F1 int8 }
type GenS15 struct { F0 uint64 }
func genFn0(p0 bool) string {
    if p0 {
        p0 = p0
        v6 := GenS1{F0: false}
        v7 := []int{2147483647, 42}
        _ = v6
        _ = v7
    } else {
        v8 := []int{-2147483648}
        var v9 int64 = 100
        _ = v8
        _ = v9
    }
    var v10 uint = 926
    if ((848 > 2147483647) && p0) {
        if ((3.14159 * 0.1) <= (0.1 * (0.1 / 1.0))) {
            v11 := GenS2{F0: 255, F1: 689}
            var v12 int8 = -1
            _ = v11
            _ = v12
        } else {
            fmt.Println(v10, v10)
        }
        p0 = p0
        v10 = 508
    } else {
        var v13 uint64 = 4294967295
        p0 = (!(!false))
        var v14 uint8 = 1
        _ = v13
        _ = v14
    }
    v15 := []int{243}
    if len(v15) > 0 {
        v15[0] = 127
    }
    _ = p0
    _ = v10
    _ = v15
    return "a"
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]string)
    var v3 uint = 553
    var v4 bool = false
    var v5 float32 = -2.25
    fmt.Println(v1, v3, v1)
    if true {
        v5 -= (0.0 - v5)
    }
    var v16 string = "hello"
    var v17 uint16 = 458
    v18 := GenS2{F0: 42, F1: 2}
    var v19 string = ""
    for fi20 := 0; fi20 < 15; fi20++ {
        if true {
            if (v1 < 1) {
                if (!(0 >= 42)) {
                    v21 := GenS3{F0: -1}
                    var v22 bool = (32767 == 160)
                    if (v5 == v5) {
                        v5 += (((((1.5 + -2.25) + v5) * ((v5 / float32(((-2.25 * 3.14159) * float64(v5)))) + v5)) / 3.14159) + v5)
                    }
                    _ = v21
                    _ = v22
                }
                v17 = 42
                var v23 int32 = v1
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 5; fi24++ {
                v4 = v4
                v25 := GenS4{F0: v1, F1: v3}
                var v26 int64 = 9223372036854775807
                _ = fi24
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(v16)
                _ = fi27
            }
            v28 := []int{fi20, (-(fi20)), (fi20 & -859)}
            _ = v28
        }
        fmt.Println(v4)
        _ = fi20
    }
    v5 = -2.25
    switch 127 {
    case 9223372036854775807:
        for fi29 := 0; fi29 < 11; fi29++ {
            for fi30 := 0; fi30 < 17; fi30++ {
                v4 = true
                _ = fi30
            }
            switch "\n\t" {
            case "a":
                v31 := GenS5{F0: 772, F1: uint8(v17), F2: 43}
                _ = v31
            case "":
                v17 -= (v17 ^ v17)
            case "foo bar":
                var v32 int16 = 0
                _ = v32
            default:
                var v33 uint32 = 42
                _ = v33
            }
            _ = fi29
        }
    case -2147483648:
        v16 = v16
    case 128:
        v34 := GenS6{F0: 1000, F1: 18446744073709551615}
        _ = v34
    }
    if v4 {
        if false {
            for fi35 := 0; fi35 < 8; fi35++ {
                v5 = float32(float64(1e10))
                for fi36 := 0; fi36 < 17; fi36++ {
                    var v37 int8 = -128
                    v16 = v16
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
            switch 128 {
            case 2147483647:
                for fi38 := 0; fi38 < 9; fi38++ {
                    var v39 int64 = 1
                    switch 42 {
                    case 2147483647:
                        if true {
                            switch (fi38 % ((fi38) | 1)) {
                            case 128:
                                for fi40 := 0; fi40 < 7; fi40++ {
                                    var v41 float32 = 1e10
                                    v42 := GenS7{F0: ("" + "a"), F1: (!(v4 || false))}
                                    var v43 int32 = (v1 ^ -1)
                                    switch 1 {
                                    case -412:
                                        v44 := GenS6{F0: 1000, F1: (v3 | v3)}
                                        _ = v44
                                    default:
                                        switch v19 {
                                        case "":
                                            v41 -= ((v41 / 1.0) + v41)
                                        case "\n\t":
                                            var v45 int64 = (v39 << (((v3 | (v3 >> ((uint(v17)) & 63)))) & 63))
                                            _ = v45
                                        case "hello":
                                            if (false || v4) {
                                                if ((false && v4) && v4) {
                                                    var v46 int32 = (v1 | v43)
                                                    var v47 uint8 = 128
                                                    v47 = v47
                                                    _ = v46
                                                    _ = v47
                                                }
                                                var v48 string = v16
                                                v49 := GenS8{F0: 0, F1: (((((v19 + "") + v19) + v16) + v16) + v19), F2: v4}
                                                _ = v48
                                                _ = v49
                                            } else {
                                                v41 = 1e10
                                                if (!v4) {
                                                    var v50 float64 = 1.5
                                                    switch v19 {
                                                    case "":
                                                        v41 = v41
                                                        if false {
                                                            var v51 int64 = (v39 / (((v39 & v39)) | 1))
                                                            _ = v51
                                                        } else {
                                                            var v52 float32 = 1.5
                                                            fmt.Println(v41)
                                                            _ = v52
                                                        }
                                                    case "\n\t":
                                                        v3 -= v3
                                                    }
                                                    v53 := []int{int(v39)}
                                                    _ = v50
                                                    _ = v53
                                                }
                                            }
                                        default:
                                            switch ("" + "foo bar") {
                                            case "":
                                                if (-1 < 1) {
                                                    switch (-(fi38)) {
                                                    case -1:
                                                        delete(v2, (fi38 & fi40))
                                                    case -190:
                                                        if (-22 < 0) {
                                                            v54 := []int{}
                                                            _ = v54
                                                        } else {
                                                            switch "foo bar" {
                                                            case "a":
                                                                if v4 {
                                                                    fmt.Println(fi38, v41)
                                                                }
                                                            case "hello":
                                                                var v55 uint32 = 503
                                                                _ = v55
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        var v56 uint16 = (v17 - 65535)
                                                        _ = v56
                                                    }
                                                }
                                            case "foo bar":
                                                fmt.Println(v43, v17)
                                            }
                                        }
                                    }
                                    if (!(42 > 42)) {
                                        v5 -= (v5 + float32((((1.5 - (((-2.25 / 1.0) + ((1e10 * 3.14159) - 1e10)) + 3.14159)) / (float64(v41) / float64(v41))) + 0.1)))
                                        if false {
                                            v57 := GenS5{F0: uint64(v17), F1: 159, F2: 34}
                                            v58 := v2[fi40]
                                            fmt.Println(v17, fi38, v39)
                                            fmt.Println(v3, v5)
                                            _ = v57
                                            _ = v58
                                        }
                                        v41 = -2.25
                                    } else {
                                        var v59 uint = v3
                                        fmt.Println(v5, v43, fi40)
                                        var v60 int64 = v39
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            case 42:
                                for fi61 := 0; fi61 < 15; fi61++ {
                                    if false {
                                        v4 = (!v4)
                                        v62 := []int{fi38, fi38, fi38}
                                        _ = v62
                                    } else {
                                        var v63 int16 = -32768
                                        v5 -= float32(3.14159)
                                        switch 128 {
                                        case 682:
                                            v63 = v63
                                        case -2147483648:
                                            var v64 float64 = 0.0
                                            var v65 float32 = v5
                                            var v66 string = (v19 + v16)
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        case 4294967295:
                                            v17 += (v17 ^ (v17 & 640))
                                        }
                                        _ = v63
                                    }
                                    v67 := GenS9{F0: (^(v17)), F1: 42, F2: (fi61 ^ (fi38 | -503))}
                                    v5 = v5
                                    _ = fi61
                                    _ = v67
                                }
                            default:
                                v68 := GenS10{F0: (v19 + ""), F1: (-1 < (v39 + (v39 * v39)))}
                                _ = v68
                            }
                            var v69 int = int(v39)
                            _ = v69
                        }
                    case -1:
                        var v70 bool = v4
                        _ = v70
                    case -447:
                        var v71 int64 = (^(v39))
                        _ = v71
                    }
                    var v72 uint32 = 154
                    _ = fi38
                    _ = v39
                    _ = v72
                }
            default:
                fmt.Println(v1)
            }
        }
        v5 = ((v5 * v5) + (v5 + v5))
        v73 := GenS6{F0: 32767, F1: v3}
        _ = v73
    }
    switch -1 {
    case 918:
        if (!v4) {
            var v74 uint64 = 18446744073709551615
            if v4 {
                v4 = false
                var v75 uint8 = 1
                if v4 {
                    fmt.Println(v16, v4, v3)
                }
                _ = v75
            } else {
                v76 := GenS11{F0: 32767}
                v2[-943] = "foo bar"
                _ = v76
            }
            v16 = ((("foo bar" + (("a" + ("foo bar" + v19)) + ("hello" + ("\n\t" + v19)))) + v16) + "\n\t")
            _ = v74
        } else {
            v77 := []int{-982}
            _ = v77
        }
    case 4294967295:
        v4 = (!v4)
    case 9223372036854775807:
        if v4 {
            v16 = ((v16 + v16) + ((v16 + (v16 + ((v16 + v16) + ((("" + "\n\t") + "hello") + "hello")))) + v16))
            switch "foo bar" {
            case "hello":
                for fi78 := 0; fi78 < 2; fi78++ {
                    v79 := []int{9223372036854775807, (-(fi78))}
                    switch int(v3) {
                    case 4294967295:
                        if (v5 < v5) {
                            var v80 uint8 = 127
                            fmt.Println(v1, v17, v4)
                            if true {
                                fmt.Println(v3)
                            } else {
                                v16 = (v19 + ("" + v19))
                                v3 = v3
                            }
                            _ = v80
                        } else {
                            v3 = v3
                            var v81 int = (fi78 ^ (fi78 / ((-2147483648) | 1)))
                            v82 := GenS2{F0: 2, F1: 65535}
                            _ = v81
                            _ = v82
                        }
                    case 364:
                        if (false && ((v4 || v4) || false)) {
                            v83 := GenS7{F0: "\n\t", F1: v4}
                            v16 = v19
                            _ = v83
                        }
                    }
                    v84 := GenS4{F0: v1, F1: (v3 % ((42) | 1))}
                    _ = fi78
                    _ = v79
                    _ = v84
                }
            case "a":
                v85 := GenS2{F0: 2, F1: 1}
                _ = v85
            case "\n\t":
                if v4 {
                    fmt.Println(v1, v19, v5)
                    fmt.Println(v4, v19, v1)
                    v86 := GenS11{F0: 1}
                    _ = v86
                } else {
                    for fi87 := 0; fi87 < 9; fi87++ {
                        var v88 int64 = 9223372036854775807
                        switch -585 {
                        case 128:
                            for fi89 := 0; fi89 < 17; fi89++ {
                                v19 = "hello"
                                if false {
                                    v90 := v2[fi87]
                                    fmt.Println(v19)
                                    _ = v90
                                }
                                if true {
                                    fmt.Println(fi89, v5)
                                    var v91 int64 = -995
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(fi89)
                                    case "":
                                        var v92 float64 = 0.0
                                        _ = v92
                                    case "foo bar":
                                        switch (fi89 * (fi87 - (fi87 / ((-397) | 1)))) {
                                        case 4294967295:
                                            v17 = v17
                                        case -1:
                                            v4 = v4
                                        case 595:
                                            if (2 > 42) {
                                                v93 := GenS12{F0: (47 < 127)}
                                                if (v16 >= v19) {
                                                    v94 := []int{fi87, (fi87 / (((fi87 & fi89)) | 1)), int(v1), (fi87 * 1)}
                                                    if (!v4) {
                                                        if len(v94) > 0 {
                                                            v94[0] = (-(fi87))
                                                        }
                                                        var v95 string = v19
                                                        _ = v95
                                                    } else {
                                                        v3 += (v3 % ((18446744073709551615) | 1))
                                                    }
                                                    _ = v94
                                                }
                                                _ = v93
                                            }
                                        default:
                                            v17 -= (v17 ^ (v17 ^ v17))
                                            var v96 int8 = 1
                                            v97 := []int{(fi89 << ((v3) & 63)), (fi87 / ((fi89) | 1)), -580}
                                            _ = v96
                                            _ = v97
                                        }
                                    }
                                    _ = v91
                                } else {
                                    v4 = false
                                }
                                _ = fi89
                            }
                        case 1:
                            var v98 uint64 = 462
                            _ = v98
                        case -115:
                            v99 := GenS13{F0: 0}
                            _ = v99
                        }
                        fmt.Println(v3, v16, fi87)
                        _ = fi87
                        _ = v88
                    }
                }
            }
            var v100 bool = (v4 && (4294967295 <= 2147483647))
            _ = v100
        } else {
            v101 := []int{4294967295}
            v1 += (v1 << (((v3 << ((0) & 63))) & 63))
            v102 := GenS14{F0: 939, F1: -23}
            _ = v101
            _ = v102
        }
    default:
        fmt.Println(v3, v16)
    }
    if v4 {
        for fi103 := 0; fi103 < 9; fi103++ {
            v16 = ""
            var v104 int64 = -1
            v105 := GenS15{F0: uint64(v3)}
            _ = fi103
            _ = v104
            _ = v105
        }
    }
    v2[1] = v19
    v106 := GenS10{F0: ("foo bar" + v16), F1: (668 == (v3 | v3))}
    var v107 uint8 = 234
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v106
    _ = v107
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
