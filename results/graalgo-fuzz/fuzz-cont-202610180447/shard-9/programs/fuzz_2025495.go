package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint8; F1 string; F2 uint32 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool }
func genFn0() string {
    switch "a" {
    case "":
        var v5 uint32 = 0
        _ = v5
    default:
        var v6 int64 = -900
        _ = v6
    }
    var v7 int32 = -2147483648
    v8 := []int{}
    for fi9 := 0; fi9 < 7; fi9++ {
        v7 -= (v7 & 0)
        v10 := GenS1{F0: 46}
        var v11 float64 = (0.1 + 1e10)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    return ""
}
func genFn1() int16 {
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS2{F0: 255, F1: fi13, F2: 0}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 bool = (false || false)
    v16 := []int{32767, 127}
    _ = v15
    _ = v16
    return 444
}

func main() {
    var v1 int16 = 100
    v2 := make(map[int]uint32)
    var v3 bool = (-9223372036854775808 == 32768)
    var v4 uint8 = 128
    switch 127 {
    case 42:
        v17 := GenS2{F0: (v4 - v4), F1: 127, F2: 2147483647}
        _ = v17
    case -2147483648:
        fmt.Println(v1, v3)
    case 2147483647:
        var v18 float32 = 1e10
        _ = v18
    default:
        if ("" <= ("hello" + "a")) {
            v1 += (^(v1))
        } else {
            var v19 uint32 = 65535
            var v20 float64 = (0.1 / 0.1)
            switch "a" {
            case "foo bar":
                var v21 int16 = (^(v1))
                _ = v21
            case "hello":
                switch 9223372036854775807 {
                case -130:
                    if (((0 == 255) || (65535 == 42)) && v3) {
                        if (913 >= -9223372036854775808) {
                            var v22 int8 = 124
                            _ = v22
                        } else {
                            v23 := GenS3{F0: 255, F1: 32767, F2: true}
                            for fi24 := 0; fi24 < 9; fi24++ {
                                if true {
                                    var v25 int32 = 2147483647
                                    switch ("a" + "hello") {
                                    case "":
                                        v4 += v4
                                    case "\n\t":
                                        v26 := []int{fi24}
                                        _ = v26
                                    default:
                                        fmt.Println(fi24)
                                    }
                                    if (!(v3 || false)) {
                                        fmt.Println(v3, v25, v25)
                                        if ((^(v25)) != v25) {
                                            v2[1] = v19
                                            v27 := []int{9223372036854775807, -825}
                                            v28 := GenS2{F0: v4, F1: fi24, F2: v19}
                                            _ = v27
                                            _ = v28
                                        }
                                        switch 42 {
                                        case 737:
                                            v19 = (v19 & v19)
                                        case 1:
                                            v29 := GenS4{F0: v4, F1: "foo bar", F2: uint32(v4)}
                                            _ = v29
                                        }
                                    }
                                    _ = v25
                                } else {
                                    v4 -= 42
                                    v19 = (^(v19))
                                }
                                _ = fi24
                            }
                            _ = v23
                        }
                        switch ("" + ("hello" + "a")) {
                        case "hello":
                            v30 := []int{2147483647}
                            _ = v30
                        case "":
                            var v31 int64 = 178
                            _ = v31
                        }
                        switch 1 {
                        case 886:
                            v32 := []int{-527, -605}
                            _ = v32
                        case 82:
                            v33 := []int{203, 9223372036854775807}
                            _ = v33
                        }
                    } else {
                        v34 := []int{-1, 32767, 4294967295, -2147483648}
                        _ = v34
                    }
                case 413:
                    v35 := GenS2{F0: v4, F1: 127, F2: 838}
                    _ = v35
                case 32768:
                    var v36 uint32 = 0
                    _ = v36
                default:
                    var v37 uint16 = 2
                    _ = v37
                }
            default:
                switch "foo bar" {
                case "hello":
                    v38 := GenS1{F0: 255}
                    _ = v38
                case "foo bar":
                    v39 := GenS2{F0: (v4 % ((v4) | 1)), F1: -804, F2: (v19 & v19)}
                    _ = v39
                default:
                    v4 = 148
                }
            }
            _ = v19
            _ = v20
        }
    }
    var v40 uint8 = 1
    var v41 uint64 = 1
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 bool = (2147483647 == 1)
        _ = fi42
        _ = v43
    }
    v40 = (v40 ^ 130)
    if true {
        fmt.Println(v40, v3)
        if false {
            v3 = (!(4294967295 >= 1))
            var v44 bool = (true || (false && (!v3)))
            fmt.Println(v1, v40, v40)
            _ = v44
        } else {
            var v45 int = -381
            for fi46 := 0; fi46 < 13; fi46++ {
                fmt.Println(v3, v1, v45)
                var v47 bool = (v3 && v3)
                fmt.Println(v1, v3)
                _ = fi46
                _ = v47
            }
            var v48 string = ("foo bar" + ("a" + (("" + "") + "")))
            _ = v45
            _ = v48
        }
    } else {
        if true {
            switch "\n\t" {
            case "a":
                switch "\n\t" {
                case "\n\t":
                    for fi49 := 0; fi49 < 9; fi49++ {
                        if ((fi49 / ((fi49) | 1)) >= fi49) {
                            v50 := []int{}
                            switch 367 {
                            case -2147483648:
                                _ = len(v50)
                            case -1:
                                switch "a" {
                                case "":
                                    v1 -= 42
                                case "foo bar":
                                    fmt.Println(v40, v40, v40)
                                }
                            case 9223372036854775807:
                                var v51 uint64 = v41
                                _ = v51
                            default:
                                v1 = (v1 << ((313) & 63))
                            }
                            if v3 {
                                v52 := GenS3{F0: 1, F1: 789, F2: (v3 && (42 >= 5))}
                                v40 = (v4 << ((255) & 63))
                                switch (fi49 - (fi49 ^ fi49)) {
                                case 820:
                                    fmt.Println(v1)
                                }
                                _ = v52
                            }
                            _ = v50
                        }
                        _ = fi49
                    }
                case "":
                    if (542 == 139) {
                        if false {
                            v53 := GenS4{F0: v40, F1: "a", F2: 76}
                            _ = len(v2)
                            v4 = 0
                            _ = v53
                        } else {
                            var v54 int32 = -632
                            if (!v3) {
                                var v55 uint64 = 2
                                var v56 int64 = 9223372036854775807
                                _ = v55
                                _ = v56
                            }
                            _ = v54
                        }
                        v4 -= (^(v4))
                        if v3 {
                            for fi57 := 0; fi57 < 16; fi57++ {
                                fmt.Println(v41, v3)
                                v58 := v2[int(v1)]
                                _ = fi57
                                _ = v58
                            }
                            v3 = (333 != 59)
                        }
                    } else {
                        v40 = (v40 >> ((42) & 63))
                    }
                case "hello":
                    if true {
                        v41 -= 758
                    } else {
                        v59 := GenS4{F0: 193, F1: "", F2: 488}
                        v1 = (v1 >> ((2) & 63))
                        _ = len(v2)
                        _ = v59
                    }
                default:
                    v40 = 128
                }
            case "foo bar":
                v60 := []int{}
                _ = v60
            default:
                switch "\n\t" {
                case "":
                    v3 = v3
                default:
                    for fi61 := 0; fi61 < 20; fi61++ {
                        for fi62 := 0; fi62 < 1; fi62++ {
                            fmt.Println(v3)
                            switch ("" + "\n\t") {
                            case "foo bar":
                                var v63 int = 1
                                _ = v63
                            case "":
                                var v64 float64 = 1e10
                                _ = v64
                            case "hello":
                                if true {
                                    switch 1 {
                                    case 281:
                                        v65 := v2[-2147483648]
                                        _ = v65
                                    case 4294967295:
                                        if v3 {
                                            if (!(((!(-2147483648 != 0)) && v3) || false)) {
                                                var v66 int32 = 42
                                                v67 := GenS3{F0: (v4 % ((v40) | 1)), F1: 1, F2: v3}
                                                v40 = v4
                                                _ = v66
                                                _ = v67
                                            }
                                            v68 := GenS5{F0: 42, F1: fi62}
                                            _ = v68
                                        } else {
                                            v40 += 2
                                            var v69 uint32 = 2147483647
                                            v70 := v2[fi62]
                                            _ = v69
                                            _ = v70
                                        }
                                    case 2147483647:
                                        if (100 > -128) {
                                            if false {
                                                v71 := []int{(-(fi61)), 2147483647}
                                                var v72 int8 = -1
                                                switch fi62 {
                                                case -87:
                                                    if len(v71) >= 2 {
                                                        v71 = v71[1:len(v71)]
                                                    }
                                                case -959:
                                                    var v73 int32 = -2147483648
                                                    _ = v73
                                                case 32767:
                                                    v41 -= (v41 % ((v41) | 1))
                                                }
                                                _ = v71
                                                _ = v72
                                            }
                                            _ = len(v2)
                                            v3 = true
                                            var v74 int8 = 8
                                            _ = v74
                                        }
                                    }
                                    v4 = v4
                                }
                            default:
                                if true {
                                    v41 -= 18446744073709551615
                                    v4 += 202
                                } else {
                                    v4 += (v4 * (^(v40)))
                                    switch fi62 {
                                    case 4294967295:
                                        if (((0.0 * (3.14159 / (3.14159 + (0.0 * -2.25)))) - (float64((1e10 - 1e10)) * 3.14159)) != (float64(1e10) - 0.0)) {
                                            var v75 int = fi62
                                            if (212 > 138) {
                                                v4 += 2
                                                var v76 bool = v3
                                                _ = v76
                                            } else {
                                                if (!true) {
                                                    v77 := v2[(-(fi62))]
                                                    v3 = (-356 != 0)
                                                    fmt.Println(v75, v3)
                                                    _ = v77
                                                } else {
                                                    v78 := []int{9223372036854775807, (fi62 ^ (^(fi62))), (fi61 % ((v75) | 1))}
                                                    v79 := GenS2{F0: v4, F1: (fi61 >> ((0) & 63)), F2: 0}
                                                    v80 := GenS2{F0: v4, F1: v75, F2: 2147483648}
                                                    v40 = v40
                                                    v2[v75] = 88
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                v75 += v75
                                            }
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v41, v75)
                                            default:
                                                v40 = (v4 >> ((uint(v41)) & 63))
                                            }
                                            fmt.Println(v75)
                                            _ = v75
                                        }
                                    }
                                    var v81 uint = 255
                                    _ = v81
                                }
                            }
                            if (4294967295 == 937) {
                                v82 := []int{fi61, 404, 127}
                                _ = v82
                            } else {
                                if (!false) {
                                    v83 := []int{fi61, (fi62 / ((fi61) | 1)), 0}
                                    fmt.Println(v41, fi62)
                                    _ = v83
                                } else {
                                    v40 = (v4 - v4)
                                }
                                var v84 string = (("a" + "") + "\n\t")
                                v84 = (v84 + "")
                                _ = v84
                            }
                            _ = fi62
                        }
                        _ = fi61
                    }
                }
            }
            v85 := GenS2{F0: v40, F1: -593, F2: 4294967295}
            v1 -= v1
            _ = v85
        }
        var v86 uint8 = (^(v40))
        fmt.Println(v40, v3)
        _ = v86
    }
    v87 := GenS6{F0: 14}
    v4 -= v40
    v1 = int16(v4)
    var v88 bool = v3
    if (("\n\t" + "foo bar") >= "") {
        var v89 uint16 = 42
        _ = v89
    } else {
        v2[32768] = 422
        v90 := GenS7{F0: v88}
        v1 = genFn1()
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v87
    _ = v88
}
