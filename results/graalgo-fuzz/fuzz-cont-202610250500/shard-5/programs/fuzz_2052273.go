package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint32 }
func genFn0() uint16 {
    v5 := GenS1{F0: "", F1: 42}
    switch 292 {
    case 0:
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: ("hello" + "foo bar"), F1: 1000}
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
    case -523:
        var v9 int16 = 42
        _ = v9
    }
    var v10 float64 = 3.14159
    v11 := []int{}
    _ = v10
    _ = v11
    _ = v5
    return 1
}

func main() {
    var v1 int = -107
    var v2 int32 = 0
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    if false {
        v3[v1] = (0 <= 0)
        switch 616 {
        case -9223372036854775808:
            v12 := []int{}
            _ = v12
        case 248:
            var v13 int = v1
            _ = v13
        case 976:
            if (!false) {
                v2 += 1
            } else {
                if (4294967295 < 846) {
                    v4[354] = 1
                } else {
                    var v14 int16 = -1
                    _ = v14
                }
                v2 += -2147483648
                for fi15 := 0; fi15 < 3; fi15++ {
                    if ((false && (32767 != 65535)) && (true && (111 <= 255))) {
                        var v16 int8 = 1
                        v17 := []int{}
                        var v18 bool = false
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 uint16 = 1
                        var v21 float32 = (0.0 + 0.1)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    var v22 bool = true
                    _ = fi15
                    _ = v22
                }
            }
        default:
            var v23 int32 = (v2 + (v2 | v2))
            _ = v23
        }
        fmt.Println(v2)
    } else {
        switch v1 {
        case -13:
            switch -1 {
            case 42:
                if (1.5 <= (0.1 + float32(((1.5 - 0.1) - 0.0)))) {
                    switch v1 {
                    case 127:
                        var v24 uint64 = 18446744073709551615
                        _ = v24
                    default:
                        v25 := GenS1{F0: ("foo bar" + "a"), F1: 1000}
                        _ = v25
                    }
                } else {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v1 -= (v1 & (fi26 * fi26))
                        if true {
                            v27 := GenS2{F0: 2, F1: -1, F2: -131}
                            switch (fi26 << ((18446744073709551615) & 63)) {
                            case -112:
                                var v28 uint64 = 704
                                _ = v28
                            case -226:
                                v2 += (v2 - v2)
                            case -2147483648:
                                v2 += (-(v2))
                            }
                            _ = v27
                        }
                        _ = fi26
                    }
                    v1 += v1
                }
            default:
                v1 -= -9223372036854775808
            }
        case 32768:
            if false {
                v29 := []int{(v1 / (((v1 % ((32768) | 1))) | 1)), (v1 & (v1 & (-(v1))))}
                _ = v29
            } else {
                var v30 uint32 = 1
                var v31 float32 = float32(3.14159)
                for fi32 := 0; fi32 < 10; fi32++ {
                    v1 = (v1 | v1)
                    _ = fi32
                }
                _ = v30
                _ = v31
            }
        case -408:
            v33 := GenS3{F0: (v2 << ((255) & 63)), F1: 176}
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 uint = 65535
        _ = fi34
        _ = v35
    }
    v4[0] = 2147483647
    var v36 int16 = 242
    var v37 uint8 = 248
    delete(v4, 1)
    if false {
        if false {
            v38 := []int{(v1 >> ((0) & 63)), v1}
            if (((v37 - v37) != (v37 / (((v37 ^ (^(v37)))) | 1))) && false) {
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 uint64 = 237
                    switch (v1 | (v1 & v1)) {
                    case 32768:
                        switch (v1 & (v1 + (-(v1)))) {
                        case -2147483648:
                            v40 += 4294967295
                        }
                    default:
                        fmt.Println(v1, fi39, v40)
                    }
                    switch ("hello" + "hello") {
                    case "a":
                        for fi41 := 0; fi41 < 18; fi41++ {
                            v42 := v4[2147483648]
                            var v43 int64 = int64(v40)
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    default:
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            for fi44 := 0; fi44 < 1; fi44++ {
                                var v45 uint16 = genFn0()
                                _ = fi44
                                _ = v45
                            }
                        case "\n\t":
                            _ = len(v38)
                        default:
                            fmt.Println(v40)
                        }
                    }
                    _ = fi39
                    _ = v40
                }
                v36 += -32768
            }
            v36 += int16(v2)
            _ = v38
        }
        var v46 uint16 = 752
        _ = v46
    } else {
        var v47 uint32 = 544
        if (v47 <= v47) {
            switch "a" {
            case "hello":
                var v48 int8 = -11
                _ = v48
            case "\n\t":
                v49 := GenS2{F0: 42, F1: -9223372036854775808, F2: 0}
                _ = v49
            case "":
                if (0 < -9223372036854775808) {
                    for fi50 := 0; fi50 < 12; fi50++ {
                        v51 := v4[(v47 & v47)]
                        var v52 int64 = -1
                        if false {
                            for fi53 := 0; fi53 < 14; fi53++ {
                                delete(v4, v51)
                                _ = fi53
                            }
                        }
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    if (-1 > 59) {
                        v37 += (v37 / (((v37 % (((v37 >> ((384) & 63))) | 1))) | 1))
                    } else {
                        if (v2 >= -34) {
                            if false {
                                if false {
                                    v54 := GenS2{F0: 0, F1: -9223372036854775808, F2: 0}
                                    v47 -= (v47 | (v47 % (((v47 % ((353) | 1))) | 1)))
                                    _ = v54
                                } else {
                                    if (true && (((((float64(((-2.25 * 1.5) / 3.14159)) * 3.14159) / 1e10) * 0.0) * 1.5) >= (3.14159 + -2.25))) {
                                        var v55 uint32 = v47
                                        var v56 bool = (true && ((0.0 / ((1.5 / 0.1) / -2.25)) < ((0.1 / (0.1 / 1.0)) - (1.5 + float64(3.14159)))))
                                        _ = v55
                                        _ = v56
                                    }
                                    v1 += (v1 >> ((uint(v37)) & 63))
                                }
                                v2 = v2
                            }
                        }
                        if false {
                            fmt.Println(v1)
                            fmt.Println(v36)
                        }
                    }
                    for fi57 := 0; fi57 < 12; fi57++ {
                        v47 = 613
                        fmt.Println(fi57, fi57)
                        _ = fi57
                    }
                }
            }
            if (658 == 0) {
                for fi58 := 0; fi58 < 18; fi58++ {
                    if false {
                        if (!((42 <= 42) && false)) {
                            for fi59 := 0; fi59 < 13; fi59++ {
                                fmt.Println(fi58, v37)
                                if (-128 != 100) {
                                    v36 = 0
                                }
                                if (!(!true)) {
                                    if (!(18446744073709551615 == 333)) {
                                        if ("foo bar" != ("a" + "\n\t")) {
                                            fmt.Println(v36, fi58, v36)
                                            switch ("a" + "hello") {
                                            case "a":
                                                v60 := GenS2{F0: uint16(v37), F1: 1, F2: 748}
                                                _ = v60
                                            case "hello":
                                                v61 := v3[v1]
                                                _ = v61
                                            case "\n\t":
                                                var v62 bool = false
                                                _ = v62
                                            default:
                                                if (255 > 1) {
                                                    var v63 int8 = 1
                                                    _ = v63
                                                }
                                            }
                                            fmt.Println(fi58, fi59, fi59)
                                        } else {
                                            var v64 uint8 = (v37 | (^(v37)))
                                            var v65 int16 = (v36 ^ 32767)
                                            var v66 int = fi59
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                        var v67 float64 = 1.5
                                        v4[715] = v47
                                        _ = v67
                                    } else {
                                        v2 = (v2 ^ (-(v2)))
                                        switch (-(fi58)) {
                                        case -9223372036854775808:
                                            v68 := GenS4{F0: false}
                                            _ = v68
                                        case -1:
                                            fmt.Println(v36)
                                        case -2147483648:
                                            v2 -= (v2 - (v2 % (((^(v2))) | 1)))
                                        default:
                                            var v69 string = "hello"
                                            _ = v69
                                        }
                                    }
                                    v2 = -1
                                    v70 := v3[fi59]
                                    _ = v70
                                }
                                _ = fi59
                            }
                            for fi71 := 0; fi71 < 13; fi71++ {
                                if false {
                                    delete(v4, 2)
                                    if false {
                                        v72 := GenS5{F0: 42}
                                        _ = v72
                                    } else {
                                        var v73 float32 = (float32(0.1) + -2.25)
                                        _ = v73
                                    }
                                }
                                var v74 uint16 = 665
                                _ = fi71
                                _ = v74
                            }
                            v75 := v3[fi58]
                            _ = v75
                        }
                        v36 -= v36
                    } else {
                        for fi76 := 0; fi76 < 16; fi76++ {
                            var v77 bool = false
                            fmt.Println(fi58, v77, v47)
                            v1 -= fi58
                            var v78 uint = 1
                            var v79 int8 = -29
                            _ = fi76
                            _ = v77
                            _ = v78
                            _ = v79
                        }
                    }
                    _ = fi58
                }
                v47 = 0
            } else {
                var v80 bool = true
                v81 := v3[v1]
                _ = v80
                _ = v81
            }
            fmt.Println(v37)
        } else {
            var v82 uint64 = 2
            fmt.Println(v82, v2)
            _ = v82
        }
        v83 := []int{(-(v1)), -1, (-(v1))}
        _ = v47
        _ = v83
    }
    switch "\n\t" {
    case "foo bar":
        var v84 int16 = (v36 >> ((4294967295) & 63))
        _ = v84
    case "a":
        v85 := []int{(v1 | 9223372036854775807)}
        _ = v85
    case "":
        fmt.Println(v37)
    default:
        v1 += 32768
    }
    for fi86 := 0; fi86 < 18; fi86++ {
        var v87 int16 = (v36 ^ (v36 / ((783) | 1)))
        var v88 int16 = v87
        _ = fi86
        _ = v87
        _ = v88
    }
    v89 := v3[v1]
    var v90 string = "\n\t"
    var v91 uint32 = 0
    var v92 uint64 = 4294967295
    if v89 {
        v93 := []int{(v1 + (v1 + (-(v1)))), (v1 * v1)}
        var v94 float64 = 1e10
        v2 -= (v2 / ((-1) | 1))
        _ = v93
        _ = v94
    }
    switch (^(v1)) {
    case -437:
        var v95 uint16 = 255
        _ = v95
    case 32768:
        var v96 int8 = 100
        _ = v96
    }
    v97 := GenS6{F0: (v36 + int16(v91)), F1: (v91 % (((v91 & uint32(v92))) | 1))}
    v37 -= 37
    fmt.Println(v91)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v89
    _ = v90
    _ = v91
    _ = v92
    _ = v97
}
