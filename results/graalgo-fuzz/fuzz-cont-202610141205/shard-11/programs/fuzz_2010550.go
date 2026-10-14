package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
type GenS2 struct { F0 string; F1 int; F2 uint }
type GenS3 struct { F0 int32; F1 string; F2 uint }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0(p0 uint8, p1 int64, p2 uint64) uint64 {
    v4 := []int{-208, 2147483647, 127}
    p1 -= p1
    if false {
        v5 := GenS1{F0: 0, F1: 2147483647, F2: p1}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 19; fi6++ {
            if false {
                for fi7 := 0; fi7 < 6; fi7++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v8 uint64 = (p2 << ((250) & 63))
                    if len(v4) > 0 {
                        v4[0] = (fi6 | int(v8))
                    }
                    _ = fi7
                    _ = v8
                }
                for fi9 := 0; fi9 < 11; fi9++ {
                    p1 += int64(p2)
                    _ = fi9
                }
            } else {
                p0 += p0
                if (!true) {
                    if ((!((true && (!true)) && false)) || true) {
                        if (!true) {
                            var v10 int = fi6
                            for fi11 := 0; fi11 < 15; fi11++ {
                                var v12 uint16 = 23
                                _ = fi11
                                _ = v12
                            }
                            for fi13 := 0; fi13 < 20; fi13++ {
                                v14 := GenS2{F0: "a", F1: 32768, F2: 42}
                                _ = fi13
                                _ = v14
                            }
                            _ = v10
                        } else {
                            switch "hello" {
                            case "\n\t":
                                var v15 float32 = (0.0 - (((float32(0.0) / -2.25) * 1e10) + 1.5))
                                _ = v15
                            case "foo bar":
                                p2 -= 2
                            }
                            fmt.Println(p2, p1, p0)
                            var v16 string = "foo bar"
                            _ = v16
                        }
                        var v17 uint16 = 32767
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 7; fi18++ {
                            _ = len(v4)
                            if len(v4) > 0 {
                                v4[0] = fi6
                            }
                            var v19 float64 = -2.25
                            _ = fi18
                            _ = v19
                        }
                    }
                    fmt.Println(fi6, fi6, p1)
                }
                var v20 string = ("hello" + "")
                _ = v20
            }
            var v21 int16 = -1
            _ = fi6
            _ = v21
        }
        var v22 uint16 = 985
        var v23 int8 = 56
        _ = v22
        _ = v23
    }
    p0 -= (p0 % ((p0) | 1))
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    v1 := []string{""}
    var v2 string = "foo bar"
    var v3 int = -192
    var v24 int16 = 1000
    v25 := GenS2{F0: (v2 + v2), F1: 128, F2: 408}
    if false {
        fmt.Println(v3)
        for fi26 := 0; fi26 < 18; fi26++ {
            switch v2 {
            case "a":
                v27 := GenS1{F0: 363, F1: 42, F2: 5}
                _ = v27
            case "hello":
                if (9223372036854775807 != -9223372036854775808) {
                    if ((("foo bar" + "") + "a") != (("a" + "") + "hello")) {
                        if true {
                            v2 = "hello"
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v28 := GenS3{F0: -2147483648, F1: v2, F2: 69}
                    _ = v28
                } else {
                    v3 = fi26
                    if ((false && true) && false) {
                        fmt.Println(v2, v24, v24)
                        v3 = (v3 >> ((2) & 63))
                        v2 = v2
                    }
                    var v29 int16 = v24
                    _ = v29
                }
            }
            _ = fi26
        }
        for fi30 := 0; fi30 < 3; fi30++ {
            v2 = (("a" + v2) + "a")
            _ = fi30
        }
    }
    v24 = (v24 | -394)
    v31 := GenS1{F0: 42, F1: -522, F2: -1}
    v32 := GenS2{F0: ("\n\t" + ((v2 + (v2 + v2)) + v2)), F1: -586, F2: 18446744073709551615}
    if (false || (0.0 <= 0.0)) {
        v33 := GenS4{F0: (187 != 0), F1: 32767, F2: ((("hello" + v2) + ("foo bar" + ("" + v2))) + v2)}
        for fi34 := 0; fi34 < 2; fi34++ {
            switch fi34 {
            case 9223372036854775807:
                if (!true) {
                    v24 = v24
                    v35 := GenS2{F0: (("\n\t" + ((v2 + (v2 + ("" + ("foo bar" + v2)))) + (v2 + ""))) + (("a" + v2) + "a")), F1: (-(fi34)), F2: 0}
                    if true {
                        for fi36 := 0; fi36 < 7; fi36++ {
                            if true {
                                switch ((v2 + v2) + v2) {
                                case "foo bar":
                                    v37 := GenS4{F0: (2147483647 < 0), F1: 255, F2: "foo bar"}
                                    _ = v37
                                case "\n\t":
                                    v1 = append(v1, "a")
                                }
                                var v38 int32 = -60
                                var v39 uint8 = 65
                                if true {
                                    if (!(false || ((v24 & (v24 | 1000)) == int16(v39)))) {
                                        var v40 bool = (float32((-2.25 / -2.25)) == 0.1)
                                        _ = v40
                                    } else {
                                        switch fi36 {
                                        case -9223372036854775808:
                                            fmt.Println(fi34)
                                        }
                                        var v41 uint64 = 4294967295
                                        var v42 uint8 = 2
                                        _ = len(v1)
                                        fmt.Println(v2, v41)
                                        var v43 uint32 = 0
                                        if (!(v2 == ((((v2 + ((v2 + v2) + (v2 + ("\n\t" + ("" + ""))))) + "") + (v2 + v2)) + (v2 + "foo bar")))) {
                                            if (!true) {
                                                var v44 uint = 65535
                                                switch (v2 + "a") {
                                                case "":
                                                    fmt.Println(v2, v3, v43)
                                                case "foo bar":
                                                    var v45 int64 = -229
                                                    _ = v45
                                                default:
                                                    var v46 int8 = 100
                                                    _ = v46
                                                }
                                                if ((true || false) || (-1 < -459)) {
                                                    v47 := GenS4{F0: (0.1 > 0.0), F1: 666, F2: v2}
                                                    v44 -= v44
                                                    v41 += (^(v41))
                                                    _ = v47
                                                }
                                                _ = v44
                                            }
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (v2 + "a")
                                        }
                                        switch v3 {
                                        case 32767:
                                            if (!true) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v48 := GenS2{F0: "\n\t", F1: (-(v3)), F2: 4294967295}
                                                _ = v48
                                            }
                                        case -445:
                                            v39 -= v39
                                        case -2147483648:
                                            v49 := GenS4{F0: (927 <= int16(v41)), F1: 0, F2: (v2 + v2)}
                                            _ = v49
                                        default:
                                            fmt.Println(v42)
                                        }
                                        var v50 uint = 65535
                                        var v51 int = v3
                                        v1 = append(v1, "")
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                        _ = v50
                                        _ = v51
                                    }
                                    fmt.Println(v24)
                                    var v52 string = (("\n\t" + v2) + v2)
                                    _ = v52
                                }
                                _ = v38
                                _ = v39
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = ((v2 + (v2 + v2)) + v2)
                                }
                                var v53 int8 = 37
                                var v54 int32 = -114
                                v53 = v53
                                _ = v53
                                _ = v54
                            }
                            _ = fi36
                        }
                        v55 := GenS3{F0: -949, F1: (v2 + "\n\t"), F2: 499}
                        _ = v55
                    } else {
                        var v56 bool = true
                        v3 += (fi34 << ((2) & 63))
                        v56 = (!(!(100 >= 81)))
                        _ = v56
                    }
                    _ = v35
                } else {
                    for fi57 := 0; fi57 < 20; fi57++ {
                        if (!(2 >= 403)) {
                            switch "a" {
                            case "\n\t":
                                switch v2 {
                                case "":
                                    v2 = "foo bar"
                                case "\n\t":
                                    var v58 int16 = v24
                                    _ = v58
                                case "foo bar":
                                    switch v2 {
                                    case "hello":
                                        if (v2 < ((v2 + v2) + v2)) {
                                            if len(v1) > 0 {
                                                v1[0] = v2
                                            }
                                            if true {
                                                v24 = (^(v24))
                                                fmt.Println(fi57, v3, v2)
                                                fmt.Println(v3)
                                            }
                                        }
                                    default:
                                        var v59 int = fi57
                                        _ = v59
                                    }
                                default:
                                    v1 = append(v1, v2)
                                    var v60 int32 = 0
                                    var v61 uint8 = 57
                                    _ = v60
                                    _ = v61
                                }
                            }
                        }
                        var v62 int8 = 1
                        _ = fi57
                        _ = v62
                    }
                    v1 = append(v1, ("foo bar" + v2))
                }
            case 84:
                v63 := GenS5{F0: (v24 % ((v24) | 1))}
                _ = v63
            }
            var v64 uint16 = 1
            if true {
                if (0.0 < ((0.0 * (0.0 + (0.0 / -2.25))) - (-2.25 / (0.0 + 1e10)))) {
                    var v65 uint32 = 587
                    _ = v65
                }
                for fi66 := 0; fi66 < 20; fi66++ {
                    v67 := GenS5{F0: v24}
                    v68 := GenS6{F0: -1}
                    v69 := GenS1{F0: (v64 % ((32767) | 1)), F1: -835, F2: 1}
                    _ = fi66
                    _ = v67
                    _ = v68
                    _ = v69
                }
                _ = len(v1)
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi70 := 0; fi70 < 18; fi70++ {
                    switch fi34 {
                    case 9223372036854775807:
                        if (-1 == 0) {
                            if true {
                                v64 += (v64 + v64)
                                if (!true) {
                                    var v71 int = fi34
                                    v72 := GenS7{F0: 696, F1: (v64 % (((v64 & 0)) | 1)), F2: -128}
                                    _ = v71
                                    _ = v72
                                } else {
                                    if false {
                                        _ = len(v1)
                                        if true {
                                            v2 = ((v2 + "foo bar") + v2)
                                        } else {
                                            switch -2147483648 {
                                            case -611:
                                                var v73 uint64 = 188
                                                _ = v73
                                            default:
                                                var v74 int64 = -9223372036854775808
                                                _ = v74
                                            }
                                            if true {
                                                if (886 > 42) {
                                                    var v75 uint32 = 2
                                                    _ = v75
                                                }
                                            }
                                            fmt.Println(v3, fi34)
                                            if false {
                                                var v76 uint64 = 0
                                                _ = v76
                                            }
                                            v2 = "a"
                                            if false {
                                                fmt.Println(fi70)
                                            } else {
                                                fmt.Println(v24, fi34, v2)
                                                v24 += v24
                                                fmt.Println(fi34, v24, v64)
                                            }
                                        }
                                        if false {
                                            var v77 uint16 = (v64 * (v64 - v64))
                                            _ = v77
                                        } else {
                                            switch (v3 << ((2) & 63)) {
                                            case 1:
                                                fmt.Println(v2)
                                            case 127:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                var v78 float64 = 0.1
                                                _ = v78
                                            }
                                            v24 += v24
                                        }
                                    } else {
                                        v79 := GenS3{F0: int32(v64), F1: v2, F2: 0}
                                        var v80 int64 = 708
                                        v80 = v80
                                        _ = v79
                                        _ = v80
                                    }
                                }
                                if false {
                                    v64 = 255
                                }
                            }
                        }
                    case 0:
                        var v81 uint8 = 1
                        _ = v81
                    case 32768:
                        var v82 uint = 0
                        _ = v82
                    }
                    var v83 uint16 = v64
                    v24 += (v24 & v24)
                    _ = fi70
                    _ = v83
                }
            }
            _ = fi34
            _ = v64
        }
        var v84 int8 = 42
        _ = v33
        _ = v84
    } else {
        fmt.Println(v3, v3)
        v85 := GenS5{F0: 1}
        var v86 float64 = -2.25
        _ = v85
        _ = v86
    }
    fmt.Println(v3)
    for fi87 := 0; fi87 < 11; fi87++ {
        var v88 int = v3
        if (!((!true) && (((!(1 == 301)) && (1 >= 0)) || false))) {
            v1 = append(v1, v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            fmt.Println(fi87, v24)
            switch (fi87 - 4294967295) {
            case 32768:
                if len(v1) > 0 {
                    v1[0] = (v2 + ((((v2 + v2) + v2) + v2) + v2))
                }
            }
            v24 = (v24 | (v24 | 0))
        }
        if len(v1) > 0 {
            v1[0] = "a"
        }
        _ = fi87
        _ = v88
    }
    v3 += 42
    switch (v3 << ((760) & 63)) {
    case 4294967295:
        v1 = append(v1, "a")
    case 128:
        for fi89 := 0; fi89 < 5; fi89++ {
            var v90 float32 = 0.1
            _ = fi89
            _ = v90
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
}
