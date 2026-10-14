package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int16; F1 string; F2 int64 }
type GenS4 struct { F0 int32; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 int; F1 int16 }
type GenS10 struct { F0 int32; F1 int32; F2 int64 }
type GenS11 struct { F0 uint32; F1 int32; F2 uint32 }
func genFn0(p0 int, p1 int, p2 int64) bool {
    v4 := []int{}
    if (1 > (p2 >> ((42) & 63))) {
        p2 -= (p2 - p2)
        v5 := GenS1{F0: 0, F1: 1, F2: 42}
        p0 = p0
        _ = v5
    } else {
        p2 += p2
        var v6 bool = false
        for fi7 := 0; fi7 < 14; fi7++ {
            p1 += p1
            _ = fi7
        }
        _ = v6
    }
    _ = len(v4)
    if (false || (!(true && false))) {
        var v8 uint16 = 65535
        _ = len(v4)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS2{F0: (p2 | 0), F1: 128, F2: 1}
            _ = fi9
            _ = v10
        }
        _ = len(v4)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (!(42 < -118))
}
func genFn1(p0 bool) int16 {
    p0 = (1 != 338)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int64 = 9223372036854775807
        v12 = (^(v12))
        _ = fi11
        _ = v12
    }
    _ = p0
    return 544
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint = 42
    v13 := GenS3{F0: 42, F1: "", F2: 42}
    if false {
        if ((true && true) || (!(!(0 > 44)))) {
            if (2 >= 255) {
                var v14 uint = v3
                for fi15 := 0; fi15 < 16; fi15++ {
                    switch "hello" {
                    case "a":
                        v14 = v14
                    case "\n\t":
                        v3 += (^(v14))
                    default:
                        if false {
                            if false {
                                v3 = v3
                            }
                        }
                    }
                    _ = fi15
                }
                _ = v14
            } else {
                switch (("" + ("" + ((((("\n\t" + "hello") + "") + "foo bar") + "foo bar") + ""))) + "\n\t") {
                case "":
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if false {
                            if false {
                                switch fi16 {
                                case 1:
                                    switch "hello" {
                                    case "foo bar":
                                        v17 := GenS3{F0: 1, F1: "a", F2: -1}
                                        _ = v17
                                    }
                                default:
                                    v3 = 65535
                                }
                            } else {
                                v18 := GenS4{F0: 377, F1: true, F2: 564}
                                var v19 float64 = (-2.25 - (0.0 + (1e10 / (1e10 * (float64(3.14159) * ((-2.25 - 1.5) + 0.0))))))
                                if (4294967295 < 965) {
                                    v20 := GenS3{F0: -32768, F1: ("\n\t" + ("hello" + "")), F2: 1}
                                    v21 := GenS3{F0: -1, F1: ("foo bar" + (("foo bar" + ("foo bar" + "\n\t")) + "foo bar")), F2: int64(v3)}
                                    _ = v20
                                    _ = v21
                                }
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = fi16
                    }
                case "foo bar":
                    v22 := []int{1, 0}
                    _ = v22
                default:
                    v23 := GenS5{F0: 42}
                    _ = v23
                }
                if false {
                    var v24 string = "foo bar"
                    switch (v24 + v24) {
                    case "foo bar":
                        if (true && true) {
                            var v25 uint64 = 809
                            if true {
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    var v27 int32 = 42
                                    v28 := v2[2147483648]
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                                v25 = 29
                            } else {
                                v29 := GenS2{F0: 1, F1: 255, F2: 537}
                                v30 := GenS1{F0: v3, F1: 1, F2: v3}
                                _ = v29
                                _ = v30
                            }
                            v25 -= v25
                            _ = v25
                        } else {
                            fmt.Println(v3, v24)
                            v31 := GenS3{F0: 1000, F1: v24, F2: 0}
                            if false {
                                v3 = 42
                            } else {
                                v24 = "a"
                                switch v24 {
                                case "hello":
                                    for fi32 := 0; fi32 < 16; fi32++ {
                                        delete(v2, 335)
                                        v33 := GenS6{F0: 127}
                                        var v34 int = -676
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                default:
                                    v3 = v3
                                }
                            }
                            _ = v31
                        }
                    case "a":
                        for fi35 := 0; fi35 < 12; fi35++ {
                            v3 = 18446744073709551615
                            v36 := []int{-1, (fi35 ^ fi35), 32768}
                            var v37 uint32 = 1
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    default:
                        if ((!(!((924 > 42) && false))) && (v24 == "foo bar")) {
                            v38 := GenS4{F0: -988, F1: (false && (true && true)), F2: 1}
                            _ = v38
                        } else {
                            fmt.Println(v24, v24)
                        }
                    }
                    var v39 int = 127
                    _ = v24
                    _ = v39
                }
                v40 := GenS7{F0: "foo bar", F1: 577}
                _ = v40
            }
            v41 := []int{}
            var v42 uint = (^(v3))
            _ = v41
            _ = v42
        } else {
            if ((!(((false && (true && true)) || true) || (!(!false)))) || true) {
                var v43 float64 = -2.25
                v44 := GenS4{F0: -172, F1: true, F2: 42}
                v43 = v43
                _ = v43
                _ = v44
            } else {
                v3 = (v3 / ((2) | 1))
                switch -33 {
                case 227:
                    v3 = v3
                default:
                    v3 -= v3
                }
            }
            switch -1 {
            case -9223372036854775808:
                for fi45 := 0; fi45 < 13; fi45++ {
                    if (!false) {
                        v46 := GenS2{F0: -1, F1: 103, F2: 1}
                        _ = v46
                    } else {
                        if false {
                            switch (fi45 / ((fi45) | 1)) {
                            case 128:
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    if (1 < 41) {
                                        var v48 float32 = (float32(float64(1e10)) + 0.0)
                                        var v49 int64 = -1
                                        v50 := v2[2]
                                        v51 := []int{(-(fi47)), fi45, fi45, fi47}
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                    var v52 bool = (((-2.25 - -2.25) / 3.14159) == 1e10)
                                    _ = fi47
                                    _ = v52
                                }
                            default:
                                v3 -= v3
                            }
                        }
                        v3 += 0
                        v3 += v3
                    }
                    v3 -= 621
                    v53 := []int{fi45, 650, fi45}
                    _ = fi45
                    _ = v53
                }
            case 597:
                var v54 int = 2147483647
                _ = v54
            default:
                for fi55 := 0; fi55 < 19; fi55++ {
                    for fi56 := 0; fi56 < 9; fi56++ {
                        var v57 uint32 = 42
                        v58 := GenS8{F0: 100, F1: (v57 & uint32(v3))}
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    var v59 uint16 = 587
                    _ = fi55
                    _ = v59
                }
            }
            v60 := GenS4{F0: -2147483648, F1: false, F2: 571}
            _ = v60
        }
        var v61 int16 = -148
        v3 -= (v3 << (((^(v3))) & 63))
        _ = v61
    } else {
        var v62 int64 = 42
        _ = v62
    }
    var v63 int32 = 2147483647
    var v64 int = 42
    if false {
        fmt.Println(v3, v64)
        var v65 int = (v64 << (((v3 * v3)) & 63))
        _ = v65
    } else {
        for fi66 := 0; fi66 < 12; fi66++ {
            var v67 float64 = -2.25
            _ = fi66
            _ = v67
        }
        if true {
            fmt.Println(v64, v3, v3)
        } else {
            var v68 uint8 = 149
            _ = v68
        }
    }
    for fi69 := 0; fi69 < 1; fi69++ {
        var v70 uint32 = 71
        v70 = 2
        _ = fi69
        _ = v70
    }
    for fi71 := 0; fi71 < 6; fi71++ {
        if false {
            v72 := GenS4{F0: (v63 / (((-(v63))) | 1)), F1: false, F2: 2}
            v73 := GenS9{F0: v64, F1: -240}
            for fi74 := 0; fi74 < 17; fi74++ {
                fmt.Println(v3, fi71)
                _ = fi74
            }
            _ = v72
            _ = v73
        } else {
            for fi75 := 0; fi75 < 5; fi75++ {
                v63 += 2147483647
                _ = fi75
            }
            v76 := GenS8{F0: -1, F1: 65535}
            var v77 uint32 = 2147483647
            _ = v76
            _ = v77
        }
        _ = fi71
    }
    if (true && true) {
        v3 = 0
    } else {
        fmt.Println(v63, v3, v3)
        for fi78 := 0; fi78 < 19; fi78++ {
            v79 := GenS2{F0: 9223372036854775807, F1: 217, F2: 4294967295}
            var v80 int64 = 470
            var v81 int8 = 1
            _ = fi78
            _ = v79
            _ = v80
            _ = v81
        }
        if true {
            v82 := []int{(v64 % ((v64) | 1)), (v64 * (v64 % (((v64 ^ (^(v64)))) | 1))), v64, 127}
            fmt.Println(v3, v63, v64)
            fmt.Println(v63, v64, v63)
            _ = v82
        }
    }
    v3 -= v3
    v3 += v3
    v83 := []int{v64, v64}
    if (3.14159 < (float32(1.5) + 0.0)) {
        fmt.Println(v64, v63)
    }
    v63 -= 41
    switch ("foo bar" + "hello") {
    case "":
        if (v64 != (v64 << (((^(v3))) & 63))) {
            switch "hello" {
            case "a":
                switch "foo bar" {
                case "hello":
                    _ = len(v83)
                case "foo bar":
                    switch -429 {
                    case 32767:
                        switch -174 {
                        case 127:
                            for fi84 := 0; fi84 < 17; fi84++ {
                                if len(v83) > 0 {
                                    v83[0] = -429
                                }
                                for fi85 := 0; fi85 < 3; fi85++ {
                                    var v86 int8 = 29
                                    if (-339 > 0) {
                                        var v87 uint16 = 986
                                        v63 = (v63 | (v63 + -163))
                                        v88 := GenS6{F0: v86}
                                        _ = v87
                                        _ = v88
                                    }
                                    if ("a" == "foo bar") {
                                        v63 = v63
                                    } else {
                                        delete(v1, fi84)
                                    }
                                    _ = fi85
                                    _ = v86
                                }
                                v89 := GenS10{F0: v63, F1: (-(v63)), F2: 1}
                                _ = fi84
                                _ = v89
                            }
                        case 4294967295:
                            var v90 bool = false
                            _ = v90
                        case 32768:
                            var v91 int64 = -634
                            _ = v91
                        }
                    case 128:
                        v64 += v64
                    }
                case "a":
                    for fi92 := 0; fi92 < 3; fi92++ {
                        for fi93 := 0; fi93 < 5; fi93++ {
                            v3 = v3
                            v3 += 1
                            _ = fi93
                        }
                        _ = fi92
                    }
                }
            case "":
                if (2 != 802) {
                    for fi94 := 0; fi94 < 9; fi94++ {
                        for fi95 := 0; fi95 < 19; fi95++ {
                            v64 -= (fi95 % ((fi94) | 1))
                            v83 = append(v83, v64)
                            _ = fi95
                        }
                        fmt.Println(v63, v3, v63)
                        for fi96 := 0; fi96 < 16; fi96++ {
                            var v97 uint64 = 18446744073709551615
                            switch (("\n\t" + "") + (((("hello" + ("\n\t" + ("hello" + "hello"))) + "hello") + "a") + ("hello" + "\n\t"))) {
                            case "foo bar":
                                var v98 float64 = (float64(1.5) - 0.1)
                                _ = v98
                            }
                            _ = fi96
                            _ = v97
                        }
                        _ = fi94
                    }
                } else {
                    _ = len(v2)
                }
            case "\n\t":
                fmt.Println(v63, v63, v64)
            }
        }
    default:
        switch "" {
        case "foo bar":
            var v99 int = (v64 & 9223372036854775807)
            _ = v99
        case "a":
            _ = len(v83)
        case "":
            var v100 uint8 = 127
            _ = v100
        default:
            var v101 uint64 = 1
            _ = v101
        }
    }
    v102 := GenS11{F0: 510, F1: (v63 % (((v63 * v63)) | 1)), F2: 581}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v83:", len(v83), v83)
    _ = v1
    _ = v102
    _ = v13
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v83
}
