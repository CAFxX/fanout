package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint64; F2 bool }
type GenS7 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 string = ("" + ("hello" + "\n\t"))
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    switch (v1 + ("a" + v1)) {
    case "":
        if false {
            fmt.Println(v1)
            var v4 int64 = 834
            _ = v4
        }
    default:
        var v5 int8 = -128
        _ = v5
    }
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v1, fi7)
        v1 = "hello"
        v8 := GenS1{F0: true, F1: true}
        _ = fi7
        _ = v8
    }
    var v9 uint = 106
    var v10 uint8 = 2
    v10 -= 2
    for fi11 := 0; fi11 < 1; fi11++ {
        _ = len(v2)
        v10 = v6
        _ = fi11
    }
    var v12 int8 = -1
    if ((true && true) || (true && (!((v6 + (^(v6))) > (v10 + 0))))) {
        if ((4294967295 == -673) || (v10 < (v6 / ((152) | 1)))) {
            v1 = (v1 + (v1 + (v1 + "foo bar")))
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint32 = 65535
                v15 := GenS1{F0: (1e10 == 0.0), F1: (true || true)}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        if (671 < 18446744073709551615) {
            _ = len(v3)
            if ((v12 % ((127) | 1)) < 42) {
                fmt.Println(v9)
                if ((1e10 + (3.14159 * 1e10)) == (float64(((1e10 / (-2.25 * (((0.1 - 3.14159) + 1.5) * 1.5))) / ((float32(((3.14159 * (0.0 + 0.0)) - 1.5)) / ((float32((0.1 + (1.5 / float64(float32((1.5 + 3.14159)))))) * float32(1e10)) / 1e10)) * (3.14159 + -2.25)))) + 0.0)) {
                    var v16 float64 = (1.5 / 1.0)
                    switch v1 {
                    case "foo bar":
                        var v17 float32 = 0.0
                        _ = v17
                    case "hello":
                        fmt.Println(v10, v16, v6)
                    case "":
                        if (21 <= 42) {
                            delete(v3, 2)
                            fmt.Println(v12)
                        } else {
                            v18 := GenS2{F0: 2147483647, F1: (v12 / ((v12) | 1))}
                            _ = v18
                        }
                    default:
                        if ((32768 < 127) || (65535 != 65535)) {
                            for fi19 := 0; fi19 < 16; fi19++ {
                                if false {
                                    fmt.Println(v10)
                                    fmt.Println(v10, v1)
                                    fmt.Println(v1, v6)
                                } else {
                                    var v20 bool = true
                                    _ = v20
                                }
                                var v21 int32 = 42
                                _ = fi19
                                _ = v21
                            }
                            var v22 int = -1
                            switch v1 {
                            case "hello":
                                fmt.Println(v6, v16)
                            }
                            _ = v22
                        } else {
                            _ = len(v2)
                        }
                    }
                    var v23 uint64 = 327
                    _ = v16
                    _ = v23
                }
            } else {
                if (v9 == (^(v9))) {
                    var v24 int = int(v9)
                    var v25 uint16 = 137
                    _ = v24
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        if false {
                            for fi27 := 0; fi27 < 12; fi27++ {
                                v28 := GenS2{F0: 2147483647, F1: (-(v12))}
                                var v29 int16 = 42
                                v29 -= 42
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 19; fi30++ {
                                if false {
                                    v31 := GenS2{F0: 517, F1: (v12 * v12)}
                                    _ = v31
                                } else {
                                    if true {
                                        fmt.Println(v6, fi30, fi30)
                                        fmt.Println(v10)
                                        var v32 uint64 = 0
                                        _ = v32
                                    } else {
                                        switch (^(fi30)) {
                                        case 42:
                                            v33 := []int{fi30, fi26, (fi30 / (((-(fi30))) | 1))}
                                            _ = v33
                                        }
                                        switch (fi26 | (-(fi26))) {
                                        case 2147483647:
                                            fmt.Println(fi30, v1, v9)
                                        case 128:
                                            delete(v2, 127)
                                        }
                                        v6 = (v6 % ((v10) | 1))
                                    }
                                    v1 = v1
                                }
                                var v34 int = -1
                                fmt.Println(v34, v12, v34)
                                _ = fi30
                                _ = v34
                            }
                            for fi35 := 0; fi35 < 18; fi35++ {
                                if (float64(((3.14159 + (((-2.25 * 0.0) + 0.0) / (1e10 * 0.1))) * 3.14159)) > 0.0) {
                                    v1 = v1
                                    switch "hello" {
                                    case "\n\t":
                                        v9 -= v9
                                    }
                                    v12 = (v12 / ((127) | 1))
                                }
                                switch fi35 {
                                case -294:
                                    v36 := GenS1{F0: ((457 != 616) || false), F1: (!(1e10 < 0.0))}
                                    _ = v36
                                }
                                _ = fi35
                            }
                        } else {
                            v37 := GenS1{F0: true, F1: true}
                            switch (fi26 >> ((2) & 63)) {
                            case 0:
                                if (0 < 319) {
                                    for fi38 := 0; fi38 < 2; fi38++ {
                                        v9 -= v9
                                        v2[fi26] = 127
                                        if false {
                                            if ((v9 > 953) || false) {
                                                v12 += v12
                                                var v39 int8 = v12
                                                _ = v39
                                            }
                                            var v40 int8 = v12
                                            switch "" {
                                            case "a":
                                                v41 := GenS3{F0: 4294967295, F1: 267}
                                                _ = v41
                                            case "\n\t":
                                                fmt.Println(fi26, fi26, v6)
                                            case "":
                                                v6 -= (^(v6))
                                            default:
                                                v42 := GenS3{F0: 1, F1: (v9 - (^(v9)))}
                                                _ = v42
                                            }
                                            _ = v40
                                        }
                                        _ = fi38
                                    }
                                    v9 = 0
                                    if (false && (true || (("" + v1) == "a"))) {
                                        if true {
                                            v43 := GenS4{F0: v1, F1: 255}
                                            _ = v43
                                        } else {
                                            switch (fi26 % ((fi26) | 1)) {
                                            case -2147483648:
                                                v44 := v2[fi26]
                                                _ = v44
                                            default:
                                                v6 -= (v10 * (^(v10)))
                                            }
                                            v9 = (v9 * (v9 | (v9 / ((v9) | 1))))
                                            v12 = -1
                                        }
                                        if (false && (true && (1 > 1))) {
                                            fmt.Println(v1, v10, fi26)
                                            v45 := GenS2{F0: 340, F1: (v12 & (v12 - (-(v12))))}
                                            fmt.Println(v12, fi26)
                                            _ = v45
                                        }
                                    }
                                } else {
                                    v9 += 255
                                    var v46 int32 = 2147483647
                                    _ = v46
                                }
                            default:
                                for fi47 := 0; fi47 < 1; fi47++ {
                                    v10 += 255
                                    if false {
                                        fmt.Println(v6)
                                        var v48 bool = (1e10 != 0.1)
                                        v6 += 193
                                        _ = v48
                                    }
                                    _ = fi47
                                }
                            }
                            _ = v37
                        }
                        v49 := GenS2{F0: -2147483648, F1: v12}
                        _ = fi26
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 19; fi50++ {
                        _ = len(v2)
                        v51 := GenS2{F0: 2147483647, F1: (v12 | v12)}
                        _ = fi50
                        _ = v51
                    }
                    if true {
                        var v52 int8 = v12
                        for fi53 := 0; fi53 < 14; fi53++ {
                            if true {
                                var v54 uint = v9
                                fmt.Println(v9, v10)
                                _ = v54
                            }
                            v55 := []int{(fi53 ^ 745)}
                            _ = fi53
                            _ = v55
                        }
                        v1 = v1
                        _ = v52
                    } else {
                        if true {
                            v1 = (v1 + (("" + v1) + ("" + (v1 + "a"))))
                        } else {
                            fmt.Println(v10)
                            for fi56 := 0; fi56 < 13; fi56++ {
                                var v57 uint = (v9 | (v9 * (v9 - (^(v9)))))
                                var v58 int16 = 0
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                            var v59 string = (((v1 + (("foo bar" + "\n\t") + "a")) + ("" + ((v1 + "") + (v1 + v1)))) + "a")
                            _ = v59
                        }
                        fmt.Println(v6)
                    }
                }
                switch int(v12) {
                case -9223372036854775808:
                    for fi60 := 0; fi60 < 19; fi60++ {
                        fmt.Println(v9)
                        if true {
                            if (v9 >= v9) {
                                v61 := GenS4{F0: v1, F1: 42}
                                v62 := []int{127, fi60}
                                _ = v61
                                _ = v62
                            }
                            v63 := GenS5{F0: 2147483647, F1: (v6 & (v6 << ((v9) & 63)))}
                            _ = v63
                        }
                        v9 += v9
                        _ = fi60
                    }
                case 32768:
                    for fi64 := 0; fi64 < 2; fi64++ {
                        v1 = "a"
                        switch (fi64 - (fi64 & -2147483648)) {
                        case -137:
                            fmt.Println(v6)
                        case -392:
                            if true {
                                _ = len(v3)
                            }
                        }
                        fmt.Println(v1)
                        _ = fi64
                    }
                default:
                    for fi65 := 0; fi65 < 16; fi65++ {
                        for fi66 := 0; fi66 < 18; fi66++ {
                            switch v1 {
                            case "foo bar":
                                delete(v2, fi66)
                            case "a":
                                v67 := GenS5{F0: 1, F1: v6}
                                _ = v67
                            default:
                                var v68 uint8 = (v6 >> (((v9 + (^(v9)))) & 63))
                                _ = v68
                            }
                            _ = fi66
                        }
                        v69 := GenS3{F0: 4294967295, F1: v9}
                        switch (fi65 * (fi65 << (((v9 % ((42) | 1))) & 63))) {
                        case -9223372036854775808:
                            var v70 int32 = 42
                            _ = v70
                        default:
                            v10 -= 1
                        }
                        _ = fi65
                        _ = v69
                    }
                }
                if false {
                    v71 := GenS6{F0: true, F1: 18446744073709551615, F2: false}
                    v1 = "foo bar"
                    _ = v71
                } else {
                    for fi72 := 0; fi72 < 3; fi72++ {
                        var v73 float64 = ((0.0 + -2.25) / 0.1)
                        fmt.Println(v6, v6)
                        for fi74 := 0; fi74 < 19; fi74++ {
                            v75 := []int{(fi74 % (((fi72 + fi72)) | 1)), -9223372036854775808, fi72, (-(fi72))}
                            if len(v75) > 0 {
                                v75[0] = -9223372036854775808
                            }
                            v76 := GenS7{F0: 42, F1: (v6 ^ (v10 + v6)), F2: 534}
                            if false {
                                v6 += (v10 & (v6 % (((v6 >> (((v9 + v9)) & 63))) | 1)))
                                v6 = (v6 << ((v9) & 63))
                                var v77 uint8 = v6
                                _ = v77
                            } else {
                                switch v1 {
                                case "a":
                                    v73 = v73
                                    v12 -= (-(v12))
                                }
                            }
                            v12 = 42
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                        _ = fi72
                        _ = v73
                    }
                }
            }
            if (!(false || false)) {
                if true {
                    fmt.Println(v1, v9)
                    var v78 uint32 = 2
                    _ = v78
                }
                switch v1 {
                case "\n\t":
                    for fi79 := 0; fi79 < 6; fi79++ {
                        v3[uint32(v9)] = 65535
                        switch v1 {
                        case "\n\t":
                            v80 := GenS5{F0: 432, F1: (v10 * (v10 - (v6 - (v6 * 2))))}
                            _ = v80
                        case "foo bar":
                            v81 := GenS6{F0: (1 > 0), F1: 18446744073709551615, F2: (18446744073709551615 <= 18446744073709551615)}
                            _ = v81
                        case "hello":
                            v12 += v12
                        }
                        _ = fi79
                    }
                case "a":
                    v1 = ""
                default:
                    v9 = (v9 % ((v9) | 1))
                }
            } else {
                fmt.Println(v9, v9, v6)
            }
        }
        switch (v1 + (("foo bar" + v1) + v1)) {
        case "\n\t":
            if true {
                fmt.Println(v10)
                v10 -= (v10 >> (((v9 ^ 853)) & 63))
            } else {
                v82 := v3[826]
                var v83 int8 = 127
                _ = v82
                _ = v83
            }
        case "foo bar":
            v84 := GenS6{F0: true, F1: 0, F2: true}
            _ = v84
        }
    }
    for fi85 := 0; fi85 < 20; fi85++ {
        v86 := []int{}
        _ = fi85
        _ = v86
    }
    v87 := GenS2{F0: -1, F1: 42}
    if ((-52 != v12) || (127 < -1)) {
        for fi88 := 0; fi88 < 11; fi88++ {
            var v89 uint64 = 1
            var v90 uint64 = 2
            var v91 float64 = (-2.25 * 3.14159)
            _ = fi88
            _ = v89
            _ = v90
            _ = v91
        }
        v92 := GenS8{F0: (v6 % ((255) | 1)), F1: 618}
        _ = v92
    }
    v6 += (v6 - 128)
    v9 -= 65535
    v93 := []int{-879, 9223372036854775807, -668, 127}
    v9 += (v9 * (v9 + v9))
    v94 := GenS1{F0: false, F1: (795 < 9223372036854775807)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v87
    _ = v9
    _ = v93
    _ = v94
}
