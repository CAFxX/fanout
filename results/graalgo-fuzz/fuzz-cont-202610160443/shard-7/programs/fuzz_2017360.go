package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8; F1 uint; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int16; F2 int }
type GenS6 struct { F0 uint64; F1 bool; F2 bool }
type GenS7 struct { F0 uint8; F1 int16; F2 uint }
type GenS8 struct { F0 uint64; F1 int16; F2 uint }
type GenS9 struct { F0 int64; F1 uint; F2 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint64, p1 uint32) int64 {
    fmt.Println(p1)
    var v4 uint32 = p1
    var v5 uint8 = uint8(v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return int64(p1)
}
func genFn1(p0 bool) uint32 {
    v6 := GenS1{F0: -2147483648}
    var v7 int32 = 1
    fmt.Println(p0, v7)
    _ = p0
    _ = v6
    _ = v7
    return 614
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    var v3 int8 = 127
    v8 := GenS2{F0: 65535, F1: 2}
    v9 := GenS3{F0: v3, F1: 255, F2: -992}
    var v10 float64 = float64(1e10)
    switch -871 {
    case -9223372036854775808:
        v11 := GenS4{F0: 349}
        _ = v11
    case 0:
        if (v1 && (v1 && v1)) {
            fmt.Println(v3, v10, v1)
            var v12 uint = 0
            _ = v12
        }
    default:
        switch 127 {
        case -1:
            if v1 {
                for fi13 := 0; fi13 < 5; fi13++ {
                    switch "" {
                    case "foo bar":
                        if (v10 < 1.5) {
                            v1 = true
                            var v14 int16 = -1
                            _ = v14
                        } else {
                            var v15 int32 = 0
                            var v16 int16 = 1
                            for fi17 := 0; fi17 < 9; fi17++ {
                                switch fi13 {
                                case 32768:
                                    fmt.Println(v3)
                                case 0:
                                    v3 = -59
                                case 2147483647:
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v16, v15)
                                    case "\n\t":
                                        fmt.Println(v15, v3, v1)
                                    }
                                default:
                                    v3 += (v3 << ((2) & 63))
                                }
                                if (!v1) {
                                    v18 := GenS1{F0: 0}
                                    v1 = v1
                                    v19 := GenS4{F0: 2}
                                    if v1 {
                                        v1 = v1
                                        v15 = 407
                                        var v20 float32 = float32((v10 / 1.5))
                                        fmt.Println(v15, v16, v15)
                                        fmt.Println(fi17, v10)
                                        var v21 float32 = ((v20 / v20) + v20)
                                        _ = v20
                                        _ = v21
                                    }
                                    if v1 {
                                        if v1 {
                                            v22 := GenS5{F0: 0, F1: v16, F2: (^(fi13))}
                                            var v23 int64 = genFn0(42, 2147483647)
                                            v3 -= (v3 / ((-128) | 1))
                                            _ = v22
                                            _ = v23
                                        }
                                        v24 := []int{1, 0, (fi17 + (fi13 & 4294967295)), (fi13 % (((fi13 >> ((4294967295) & 63))) | 1))}
                                        _ = v24
                                    }
                                    v25 := GenS6{F0: 18446744073709551615, F1: true, F2: (v1 && true)}
                                    _ = v18
                                    _ = v19
                                    _ = v25
                                }
                                _ = fi17
                            }
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = fi13
                }
                var v26 int64 = -1
                _ = v26
            } else {
                for fi27 := 0; fi27 < 12; fi27++ {
                    v10 += v10
                    _ = fi27
                }
            }
        default:
            v1 = (834 != 2147483647)
        }
    }
    var v28 uint64 = 4294967295
    switch 127 {
    case 367:
        v29 := v2[4294967295]
        _ = v29
    case -613:
        fmt.Println(v1, v3, v28)
    default:
        var v30 uint = 18446744073709551615
        _ = v30
    }
    v1 = v1
    switch 0 {
    case -9223372036854775808:
        var v31 bool = true
        _ = v31
    case 32767:
        if v1 {
            fmt.Println(v10, v1)
            var v32 int32 = 2147483647
            var v33 uint64 = (v28 | 18446744073709551615)
            _ = v32
            _ = v33
        } else {
            if true {
                v34 := []int{127, 9223372036854775807}
                _ = v34
            }
            for fi35 := 0; fi35 < 5; fi35++ {
                v1 = v1
                _ = fi35
            }
        }
    case 9223372036854775807:
        v2[340] = "hello"
    }
    var v36 int8 = 1
    if v1 {
        v10 = ((v10 * ((-2.25 + v10) - float64((-2.25 * (0.0 + -2.25))))) - (3.14159 / 0.1))
        for fi37 := 0; fi37 < 11; fi37++ {
            switch -459 {
            case 42:
                switch "hello" {
                case "hello":
                    var v38 float64 = (v10 - v10)
                    _ = v38
                case "\n\t":
                    if (!v1) {
                        v39 := GenS2{F0: 1, F1: 65535}
                        _ = v39
                    }
                default:
                    if true {
                        for fi40 := 0; fi40 < 10; fi40++ {
                            v3 -= 127
                            _ = fi40
                        }
                        v41 := GenS4{F0: v28}
                        var v42 string = "foo bar"
                        _ = v41
                        _ = v42
                    } else {
                        var v43 bool = v1
                        v43 = v1
                        _ = v43
                    }
                }
            case 348:
                v44 := []int{}
                _ = v44
            case 32767:
                fmt.Println(v3)
            }
            _ = fi37
        }
    } else {
        v45 := []int{-782, 42, 2147483647, 128}
        var v46 bool = v1
        _ = v45
        _ = v46
    }
    var v47 uint32 = 1
    var v48 bool = (!(!(4294967295 < (v28 << ((132) & 63)))))
    v48 = v1
    if v1 {
        for fi49 := 0; fi49 < 2; fi49++ {
            if v48 {
                for fi50 := 0; fi50 < 2; fi50++ {
                    v48 = false
                    v51 := []int{fi50, (fi50 << ((uint(v28)) & 63)), 32768}
                    _ = fi50
                    _ = v51
                }
                v52 := []int{(fi49 - fi49), -2147483648}
                var v53 int8 = -118
                _ = v52
                _ = v53
            }
            var v54 uint16 = 0
            if v48 {
                if (uint8(v28) >= 128) {
                    if true {
                        _ = len(v2)
                        fmt.Println(v28, fi49, v3)
                        switch "" {
                        case "\n\t":
                            fmt.Println(v10, v28, v3)
                        case "foo bar":
                            delete(v2, genFn1(true))
                        case "":
                            if ((-671 != -1) && false) {
                                switch (-(fi49)) {
                                case 32768:
                                    v55 := []int{(fi49 ^ fi49), 128, fi49, fi49}
                                    _ = v55
                                case 1:
                                    for fi56 := 0; fi56 < 12; fi56++ {
                                        var v57 float64 = 0.0
                                        _ = fi56
                                        _ = v57
                                    }
                                case 392:
                                    v58 := []int{(fi49 >> ((42) & 63))}
                                    _ = v58
                                default:
                                    var v59 uint = 18446744073709551615
                                    _ = v59
                                }
                            }
                        }
                    } else {
                        if v48 {
                            fmt.Println(v1, fi49, v36)
                        }
                        v54 -= 1
                    }
                    var v60 uint64 = v28
                    var v61 int16 = -32768
                    _ = v60
                    _ = v61
                } else {
                    v54 -= 255
                    v62 := []int{fi49, 32767, 4294967295}
                    _ = v62
                }
                if v1 {
                    for fi63 := 0; fi63 < 16; fi63++ {
                        var v64 uint8 = 42
                        _ = fi63
                        _ = v64
                    }
                    v65 := []int{fi49, fi49, (fi49 * fi49)}
                    _ = v65
                } else {
                    switch (fi49 ^ (fi49 % ((-1) | 1))) {
                    case 296:
                        fmt.Println(fi49, v36, v10)
                    case -2147483648:
                        fmt.Println(fi49)
                    case -203:
                        fmt.Println(v3, v10)
                    default:
                        for fi66 := 0; fi66 < 7; fi66++ {
                            v36 += 0
                            _ = fi66
                        }
                    }
                }
                var v67 int64 = -9223372036854775808
                _ = v67
            } else {
                v1 = ((fi49 & (-(fi49))) != (fi49 % ((fi49) | 1)))
                v68 := []int{(fi49 % ((-934) | 1))}
                _ = v68
            }
            _ = fi49
            _ = v54
        }
        if (128 == -980) {
            var v69 float32 = 0.0
            v70 := GenS3{F0: 45, F1: 65535, F2: -216}
            if false {
                v71 := GenS7{F0: 215, F1: 1000, F2: 4294967295}
                var v72 int64 = -9223372036854775808
                var v73 int32 = 2147483647
                _ = v71
                _ = v72
                _ = v73
            } else {
                var v74 uint8 = 2
                fmt.Println(v74, v28)
                _ = v74
            }
            _ = v69
            _ = v70
        } else {
            v75 := GenS8{F0: 42, F1: -203, F2: 65535}
            fmt.Println(v3, v48, v1)
            _ = v75
        }
        var v76 int32 = -1
        _ = v76
    } else {
        if (((v3 != v3) && true) || (true || ((^(v36)) != (v36 ^ v3)))) {
            for fi77 := 0; fi77 < 10; fi77++ {
                v78 := GenS8{F0: (^(v28)), F1: -32768, F2: 65535}
                if (v48 || false) {
                    var v79 uint64 = 1
                    var v80 uint16 = 32767
                    _ = v79
                    _ = v80
                } else {
                    v48 = v1
                    v81 := GenS2{F0: 65535, F1: 255}
                    v82 := GenS7{F0: 42, F1: 1, F2: 1}
                    _ = v81
                    _ = v82
                }
                _ = fi77
                _ = v78
            }
        } else {
            var v83 int64 = -538
            for fi84 := 0; fi84 < 16; fi84++ {
                v85 := GenS9{F0: 410, F1: 1, F2: 0}
                if false {
                    for fi86 := 0; fi86 < 11; fi86++ {
                        fmt.Println(v48, v28, v36)
                        if (-109 >= v36) {
                            if v48 {
                                delete(v2, v47)
                                v87 := []int{}
                                fmt.Println(v10, v83)
                                _ = v87
                            } else {
                                v88 := []int{(fi86 << ((349) & 63))}
                                v47 = 177
                                var v89 int = (fi84 ^ (fi86 % (((fi86 << ((637) & 63))) | 1)))
                                if (-2147483648 >= -2147483648) {
                                    var v90 int = fi84
                                    var v91 int16 = 32767
                                    var v92 int64 = v83
                                    _ = v90
                                    _ = v91
                                    _ = v92
                                }
                                _ = v88
                                _ = v89
                            }
                            v93 := []int{}
                            switch (("foo bar" + "") + "foo bar") {
                            case "\n\t":
                                fmt.Println(v28)
                            case "a":
                                var v94 uint64 = 42
                                _ = v94
                            case "":
                                switch ("hello" + "a") {
                                case "foo bar":
                                    fmt.Println(v28, v36)
                                case "a":
                                    v3 = (^(v36))
                                case "\n\t":
                                    var v95 uint16 = 1
                                    if ((!((v28 % (((^(v28))) | 1)) <= v28)) && false) {
                                        v96 := GenS2{F0: 42, F1: 980}
                                        v47 += 145
                                        var v97 int = (fi84 % ((fi86) | 1))
                                        v36 = v3
                                        _ = v96
                                        _ = v97
                                    }
                                    _ = v95
                                default:
                                    var v98 int8 = v36
                                    _ = v98
                                }
                            default:
                                v99 := GenS4{F0: (v28 % ((v28) | 1))}
                                _ = v99
                            }
                            _ = v93
                        } else {
                            var v100 int = fi84
                            fmt.Println(v83)
                            v36 -= v3
                            fmt.Println(fi86, v3)
                            _ = v100
                        }
                        fmt.Println(v3, v48, v1)
                        var v101 uint16 = 432
                        _ = fi86
                        _ = v101
                    }
                    var v102 float64 = v10
                    _ = v102
                } else {
                    switch "a" {
                    case "":
                        switch "foo bar" {
                        case "\n\t":
                            var v103 float32 = 1e10
                            _ = v103
                        case "foo bar":
                            v47 = 0
                        case "a":
                            if (v83 >= genFn0(v28, 291)) {
                                v83 -= (v83 ^ -942)
                                v104 := GenS10{F0: 65535}
                                _ = v104
                            } else {
                                v105 := GenS4{F0: v28}
                                v106 := GenS9{F0: (v83 ^ genFn0((^(v28)), 918)), F1: 255, F2: 65535}
                                _ = v105
                                _ = v106
                            }
                        }
                    }
                    for fi107 := 0; fi107 < 13; fi107++ {
                        var v108 uint64 = 18446744073709551615
                        _ = fi107
                        _ = v108
                    }
                }
                _ = fi84
                _ = v85
            }
            _ = v83
        }
        for fi109 := 0; fi109 < 5; fi109++ {
            var v110 int32 = 1
            switch "\n\t" {
            case "hello":
                if v1 {
                    var v111 uint8 = 240
                    _ = v111
                }
            default:
                v112 := GenS11{F0: int16(v110), F1: 42}
                _ = v112
            }
            _ = fi109
            _ = v110
        }
        switch "hello" {
        case "hello":
            v113 := []int{32767, -483}
            _ = v113
        case "a":
            v114 := GenS5{F0: 875, F1: -32768, F2: 127}
            _ = v114
        default:
            switch "a" {
            case "":
                var v115 uint8 = 255
                _ = v115
            case "a":
                switch -553 {
                case 857:
                    var v116 int8 = (v3 << ((2) & 63))
                    _ = v116
                case -477:
                    if (!v1) {
                        if v48 {
                            v117 := []int{}
                            v118 := GenS2{F0: 4294967295, F1: 4294967295}
                            v119 := GenS4{F0: v28}
                            _ = v117
                            _ = v118
                            _ = v119
                        } else {
                            var v120 int = -257
                            _ = v120
                        }
                    }
                }
            default:
                fmt.Println(v36, v47, v28)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v47
    _ = v48
    _ = v8
    _ = v9
}
