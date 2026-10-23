package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 uint; F2 uint32 }
type GenS8 struct { F0 int64; F1 int16; F2 uint64 }
type GenS9 struct { F0 uint64; F1 int64; F2 int }
func genFn0(p0 uint64) int32 {
    if (!(!true)) {
        for fi5 := 0; fi5 < 18; fi5++ {
            p0 += p0
            p0 += (p0 << ((332) & 63))
            _ = fi5
        }
        var v6 float32 = (0.0 / 3.14159)
        v7 := []int{394}
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{32768}
        _ = fi8
        _ = v9
    }
    var v10 int16 = 0
    fmt.Println(v10, v10)
    _ = p0
    _ = v10
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 255
    var v3 uint8 = v2
    var v4 int16 = -116
    switch 127 {
    case 2147483647:
        if (v2 > (v3 * 2)) {
            var v11 int8 = -43
            v12 := []int{}
            _ = v11
            _ = v12
        }
    case -836:
        switch -453 {
        case 801:
            for fi13 := 0; fi13 < 5; fi13++ {
                if (1 <= int64(v4)) {
                    v14 := []int{}
                    _ = v14
                } else {
                    if true {
                        switch "" {
                        case "foo bar":
                            if true {
                                v4 = v4
                            } else {
                                var v15 bool = true
                                v4 = v4
                                _ = v15
                            }
                        case "":
                            fmt.Println(v3, v3)
                        }
                        v16 := GenS1{F0: 42, F1: 281, F2: 277}
                        v3 += v2
                        _ = v16
                    }
                    v17 := GenS2{F0: 1}
                    var v18 int32 = 0
                    _ = v17
                    _ = v18
                }
                v19 := GenS1{F0: 1, F1: 1, F2: genFn0(1)}
                var v20 float32 = 3.14159
                _ = fi13
                _ = v19
                _ = v20
            }
        case -1:
            _ = len(v1)
        default:
            var v21 uint = 18446744073709551615
            _ = v21
        }
    case 42:
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS3{F0: 2147483647, F1: fi22}
            if (true && true) {
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 int = (-(fi24))
                    fmt.Println(v25)
                    v26 := v1[9]
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 string = "a"
                _ = v27
            }
            _ = fi22
            _ = v23
        }
    default:
        switch "a" {
        case "":
            if (65535 <= 1) {
                var v28 int32 = 2147483647
                _ = v28
            }
        case "foo bar":
            v4 = (v4 | (v4 << ((327) & 63)))
        }
    }
    if (-1 > (v4 >> ((0) & 63))) {
        if (((((0.0 - -2.25) != (0.0 + 1.5)) || true) && false) && (735 < 1)) {
            var v29 uint32 = 246
            v30 := []int{-1, -5, 128, -1}
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 17; fi31++ {
                v32 := GenS3{F0: 0, F1: (fi31 % (((fi31 >> ((577) & 63))) | 1))}
                _ = fi31
                _ = v32
            }
            v33 := GenS4{F0: 18446744073709551615}
            var v34 int64 = 0
            _ = v33
            _ = v34
        }
        v35 := []int{486, 42}
        v36 := v1[113]
        _ = v35
        _ = v36
    } else {
        var v37 uint = 18446744073709551615
        for fi38 := 0; fi38 < 9; fi38++ {
            v3 = (v3 | v2)
            for fi39 := 0; fi39 < 13; fi39++ {
                if true {
                    v40 := []int{-21}
                    delete(v1, 282)
                    var v41 int32 = 913
                    if (true && true) {
                        v42 := GenS1{F0: 0, F1: (v37 - v37), F2: v41}
                        _ = v42
                    }
                    switch fi38 {
                    case -9223372036854775808:
                        var v43 int = 1
                        _ = v43
                    case -81:
                        v44 := v1[616]
                        _ = v44
                    }
                    _ = v40
                    _ = v41
                } else {
                    v37 += (v37 * v37)
                }
                _ = fi39
            }
            var v45 float64 = 0.0
            _ = fi38
            _ = v45
        }
        if true {
            v46 := v1[2147483647]
            var v47 float32 = (1.5 * (-2.25 + (1e10 / 1.0)))
            _ = v46
            _ = v47
        } else {
            var v48 float32 = 1e10
            var v49 int = 1
            _ = v48
            _ = v49
        }
        _ = v37
    }
    switch "a" {
    case "\n\t":
        var v50 uint = 2
        _ = v50
    case "hello":
        fmt.Println(v2, v3)
    case "foo bar":
        for fi51 := 0; fi51 < 11; fi51++ {
            v52 := GenS5{F0: (v3 & (^(v2))), F1: 0}
            for fi53 := 0; fi53 < 8; fi53++ {
                var v54 float64 = (-2.25 / 3.14159)
                switch "hello" {
                case "a":
                    v55 := GenS6{F0: (v3 | v2)}
                    _ = v55
                case "hello":
                    if true {
                        fmt.Println(v54)
                        v54 = 1e10
                        switch 2147483647 {
                        case 845:
                            v56 := GenS5{F0: 1, F1: 65535}
                            _ = v56
                        case -2147483648:
                            fmt.Println(fi51, v3, v54)
                        case 9223372036854775807:
                            v57 := GenS6{F0: (v2 / ((255) | 1))}
                            _ = v57
                        default:
                            if ((2 <= 65535) || true) {
                                v58 := GenS7{F0: "", F1: 255, F2: 839}
                                _ = v58
                            } else {
                                v54 = v54
                                v3 -= (v3 / ((1) | 1))
                                if true {
                                    var v59 uint8 = (v2 ^ v3)
                                    var v60 uint8 = (^(v2))
                                    _ = v59
                                    _ = v60
                                } else {
                                    v54 -= (v54 * 0.1)
                                    if true {
                                        var v61 uint8 = v3
                                        fmt.Println(v4, v3, fi51)
                                        _ = v61
                                    } else {
                                        v62 := GenS8{F0: 9223372036854775807, F1: -194, F2: uint64(v3)}
                                        _ = v62
                                    }
                                }
                            }
                        }
                    }
                case "\n\t":
                    v63 := []int{(fi51 & (fi53 & fi51)), (fi53 & fi53)}
                    _ = v63
                }
                _ = fi53
                _ = v54
            }
            var v64 int = (-(fi51))
            _ = fi51
            _ = v52
            _ = v64
        }
    default:
        for fi65 := 0; fi65 < 8; fi65++ {
            var v66 string = ""
            _ = fi65
            _ = v66
        }
    }
    switch (("" + "a") + "") {
    case "":
        v67 := []int{-1, -2147483648, -195}
        _ = v67
    case "\n\t":
        v68 := GenS3{F0: 152, F1: -2147483648}
        _ = v68
    case "foo bar":
        fmt.Println(v4, v3)
    default:
        var v69 uint32 = 65535
        _ = v69
    }
    var v70 uint32 = 2
    for fi71 := 0; fi71 < 4; fi71++ {
        v4 -= v4
        if (v70 != v70) {
            switch "a" {
            case "a":
                var v72 int16 = (-(v4))
                _ = v72
            case "\n\t":
                v73 := GenS9{F0: 8, F1: -19, F2: fi71}
                _ = v73
            case "":
                for fi74 := 0; fi74 < 1; fi74++ {
                    var v75 int16 = 1
                    v2 = 127
                    switch (fi71 - 127) {
                    case -9223372036854775808:
                        v76 := []int{(fi74 ^ (-(fi74))), (^(fi71)), fi74}
                        _ = v76
                    case 32767:
                        v2 = v2
                    }
                    _ = fi74
                    _ = v75
                }
            default:
                for fi77 := 0; fi77 < 1; fi77++ {
                    switch fi71 {
                    case 0:
                        switch "" {
                        case "hello":
                            v3 -= (v2 ^ v3)
                        case "\n\t":
                            v70 = v70
                        default:
                            var v78 uint32 = (v70 % ((v70) | 1))
                            _ = v78
                        }
                    case 9223372036854775807:
                        v79 := GenS5{F0: v2, F1: 42}
                        _ = v79
                    case -9223372036854775808:
                        v80 := []int{fi71}
                        _ = v80
                    }
                    _ = fi77
                }
            }
        } else {
            if (genFn0(2) >= genFn0(1000)) {
                v81 := GenS9{F0: 18446744073709551615, F1: -1, F2: -138}
                if (false && (1 == -2147483648)) {
                    switch -272 {
                    case -9223372036854775808:
                        var v82 int8 = -104
                        _ = v82
                    case 32768:
                        var v83 int64 = 42
                        _ = v83
                    case 1:
                        var v84 bool = ((uint(v70) <= 0) && (fi71 > -1))
                        _ = v84
                    }
                    var v85 uint32 = uint32(v3)
                    _ = v85
                }
                if false {
                    v70 -= 2147483648
                    switch "a" {
                    case "a":
                        v86 := v1[(v70 << ((731) & 63))]
                        _ = v86
                    case "hello":
                        v87 := []int{}
                        _ = v87
                    case "\n\t":
                        fmt.Println(v70, v3)
                    }
                    for fi88 := 0; fi88 < 8; fi88++ {
                        delete(v1, v70)
                        var v89 uint64 = 990
                        v1[4294967295] = (fi88 >> ((65535) & 63))
                        _ = fi88
                        _ = v89
                    }
                } else {
                    if (!false) {
                        v70 -= 2147483647
                        if (!true) {
                            if (!true) {
                                delete(v1, (v70 >> ((0) & 63)))
                            } else {
                                delete(v1, (v70 - 844))
                            }
                        } else {
                            v1[42] = -245
                            var v90 float32 = ((float32((1.5 - (0.0 / 1.0))) + (((float32(0.1) / 1.0) * 0.1) * ((1e10 / (float32(0.1) + 1.5)) * 0.0))) + 0.0)
                            fmt.Println(v2)
                            _ = v90
                        }
                    } else {
                        if ((v4 >> ((0) & 63)) != v4) {
                            for fi91 := 0; fi91 < 11; fi91++ {
                                if false {
                                    v92 := GenS9{F0: 2, F1: 1, F2: (fi71 ^ 9223372036854775807)}
                                    fmt.Println(v2)
                                    _ = v92
                                }
                                _ = fi91
                            }
                        } else {
                            if false {
                                v4 += 797
                            }
                            for fi93 := 0; fi93 < 5; fi93++ {
                                v94 := []int{(fi93 >> ((1) & 63)), (fi71 << ((65535) & 63)), -433}
                                v95 := GenS3{F0: (v70 / (((^(v70))) | 1)), F1: (fi93 + fi71)}
                                v96 := GenS7{F0: "foo bar", F1: 885, F2: 1}
                                _ = fi93
                                _ = v94
                                _ = v95
                                _ = v96
                            }
                            for fi97 := 0; fi97 < 19; fi97++ {
                                if (3.14159 < (-2.25 * 3.14159)) {
                                    var v98 uint32 = 2147483648
                                    _ = v98
                                }
                                var v99 uint8 = v3
                                var v100 int64 = -9223372036854775808
                                _ = fi97
                                _ = v100
                                _ = v99
                            }
                        }
                        v101 := GenS2{F0: 4294967295}
                        for fi102 := 0; fi102 < 11; fi102++ {
                            if (!false) {
                                v103 := GenS6{F0: v2}
                                _ = v103
                            }
                            _ = fi102
                        }
                        _ = v101
                    }
                    v3 += (v3 ^ v3)
                }
                _ = v81
            } else {
                var v104 int8 = 43
                v3 = (v3 % (((v3 * 84)) | 1))
                _ = v104
            }
        }
        v70 -= v70
        _ = fi71
    }
    var v105 int16 = 202
    var v106 bool = true
    fmt.Println(v70)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v70
}
