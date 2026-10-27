package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 string; F2 int }
type GenS5 struct { F0 bool; F1 int16; F2 int16 }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 int16; F2 uint32 }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 uint32; F1 int8; F2 bool }

func main() {
    var v1 int32 = 42
    var v2 uint8 = 255
    var v3 int8 = 20
    var v4 uint32 = 4294967295
    v5 := make(map[uint32]string)
    v6 := []uint32{v4, 0}
    if (!true) {
        v4 -= 2
    } else {
        switch -9223372036854775808 {
        case 32768:
            if true {
                if (!false) {
                    v7 := GenS1{F0: 0}
                    var v8 int8 = -128
                    var v9 uint = 255
                    _ = v7
                    _ = v8
                    _ = v9
                }
                fmt.Println(v3, v2, v2)
            } else {
                if ((!(((v4 + v4) < (v4 | (^(v4)))) && false)) && true) {
                    if (18446744073709551615 >= 0) {
                        var v10 int8 = v3
                        v4 = 2
                        var v11 uint8 = (v2 * (v2 | v2))
                        _ = v10
                        _ = v11
                    }
                    v2 = v2
                    v1 += (v1 * v1)
                }
                if true {
                    v1 -= (-(v1))
                    v5[v4] = "foo bar"
                } else {
                    if true {
                        var v12 int32 = 482
                        _ = v12
                    }
                    switch 0 {
                    case 0:
                        v13 := GenS2{F0: 1, F1: 0, F2: 65535}
                        _ = v13
                    case -9223372036854775808:
                        var v14 uint64 = 147
                        _ = v14
                    }
                }
            }
        case 9223372036854775807:
            var v15 uint8 = v2
            _ = v15
        default:
            fmt.Println(v2)
        }
        fmt.Println(v2, v3, v3)
        fmt.Println(v1, v2, v4)
    }
    v5[v4] = "foo bar"
    var v16 int = 1
    delete(v5, (v4 | (v4 / ((v4) | 1))))
    var v17 uint16 = 42
    switch "\n\t" {
    case "\n\t":
        for fi18 := 0; fi18 < 2; fi18++ {
            v1 = (v1 | (v1 + (^(v1))))
            _ = fi18
        }
    case "":
        for fi19 := 0; fi19 < 1; fi19++ {
            v16 = v16
            if false {
                var v20 float32 = float32((0.0 + ((0.1 - 1e10) - 3.14159)))
                v2 = uint8(v4)
                _ = len(v5)
                _ = v20
            }
            var v21 int32 = v1
            _ = fi19
            _ = v21
        }
    case "a":
        var v22 float64 = (float64((-2.25 / 3.14159)) * (((0.0 / (float64(3.14159) + 0.0)) * 1e10) - -2.25))
        _ = v22
    }
    if true {
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := GenS2{F0: v17, F1: 42, F2: 65535}
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint32 = 42
                v17 += (v17 << ((42) & 63))
                _ = fi25
                _ = v26
            }
            if (-2.25 <= 0.0) {
                v16 -= (fi23 | -246)
                switch "hello" {
                case "foo bar":
                    if false {
                        if ((72 == 32767) || true) {
                            var v27 uint8 = (v2 - (^(v2)))
                            _ = v27
                        } else {
                            var v28 string = ("foo bar" + ("hello" + "foo bar"))
                            var v29 bool = (!((42 < 731) && false))
                            _ = v28
                            _ = v29
                        }
                        v1 -= v1
                        fmt.Println(v17, v17, v3)
                    } else {
                        if true {
                            if len(v6) > 0 {
                                v6[0] = 0
                            }
                        }
                        fmt.Println(v1, v2, v2)
                    }
                case "\n\t":
                    v30 := GenS2{F0: 338, F1: 255, F2: 0}
                    _ = v30
                case "":
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v32 := GenS1{F0: 4294967295}
                        v33 := GenS3{F0: (v1 | int32(v2))}
                        v34 := GenS1{F0: 2}
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                default:
                    switch (-(v16)) {
                    case 128:
                        _ = len(v5)
                    case -1:
                        fmt.Println(v2, v17, v16)
                    case 4294967295:
                        v16 = fi23
                    }
                }
            } else {
                switch "\n\t" {
                case "a":
                    switch v16 {
                    case -668:
                        v16 = fi23
                    case 4294967295:
                        v1 = (-(v1))
                    default:
                        for fi35 := 0; fi35 < 2; fi35++ {
                            fmt.Println(v4, v16)
                            v36 := GenS3{F0: v1}
                            _ = fi35
                            _ = v36
                        }
                    }
                case "\n\t":
                    delete(v5, v4)
                case "hello":
                    var v37 int16 = -670
                    _ = v37
                }
                var v38 int16 = -94
                _ = v38
            }
            _ = fi23
            _ = v24
        }
        v17 -= uint16(v2)
    } else {
        fmt.Println(v17, v2)
    }
    switch -9223372036854775808 {
    case 42:
        var v39 uint = 1
        _ = v39
    default:
        v5[v4] = ((("hello" + "") + ("a" + "")) + "\n\t")
    }
    switch "hello" {
    case "a":
        if true {
            var v40 uint8 = 42
            _ = v40
        }
    case "hello":
        for fi41 := 0; fi41 < 16; fi41++ {
            var v42 uint8 = (^(v2))
            _ = fi41
            _ = v42
        }
    case "foo bar":
        var v43 int64 = -1
        _ = v43
    default:
        var v44 uint64 = uint64(v2)
        _ = v44
    }
    var v45 uint = 1
    switch int(v17) {
    case 32768:
        var v46 string = "a"
        _ = v46
    case -488:
        _ = len(v5)
    case -400:
        if false {
            for fi47 := 0; fi47 < 7; fi47++ {
                v5[v4] = ("" + "hello")
                _ = len(v5)
                v2 += v2
                _ = fi47
            }
        } else {
            for fi48 := 0; fi48 < 16; fi48++ {
                v49 := v5[uint32(v45)]
                v16 -= fi48
                var v50 int = v16
                _ = fi48
                _ = v49
                _ = v50
            }
            v51 := GenS2{F0: 65535, F1: (v45 - uint(v17)), F2: (v17 | (v17 - (v17 ^ 42)))}
            _ = v51
        }
    }
    var v52 bool = true
    switch (v16 << ((v45) & 63)) {
    case 32768:
        for fi53 := 0; fi53 < 19; fi53++ {
            if v52 {
                _ = len(v6)
                v54 := GenS1{F0: 779}
                v55 := v5[uint32(v2)]
                _ = v54
                _ = v55
            } else {
                fmt.Println(fi53, v3, v2)
                if (-9223372036854775808 > 42) {
                    v6 = append(v6, v4)
                } else {
                    var v56 float64 = ((3.14159 * (0.1 - ((float64(-2.25) / 1.0) / 1.0))) / 3.14159)
                    _ = v56
                }
                if (!(0.1 >= 0.1)) {
                    fmt.Println(fi53, fi53, v3)
                    v57 := GenS1{F0: 2}
                    _ = v57
                } else {
                    if len(v6) > 0 {
                        v6[0] = v4
                    }
                }
            }
            delete(v5, (v4 * 709))
            if ((!false) || v52) {
                if ("foo bar" != "hello") {
                    fmt.Println(v17)
                    v16 -= int(v17)
                }
                var v58 int8 = 127
                _ = v58
            }
            _ = fi53
        }
    case -2147483648:
        var v59 uint8 = v2
        _ = v59
    default:
        var v60 int32 = (v1 % ((-1) | 1))
        _ = v60
    }
    for fi61 := 0; fi61 < 20; fi61++ {
        if false {
            if v52 {
                if len(v6) > 0 {
                    v6[0] = v4
                }
                var v62 uint = (v45 - 2)
                _ = v62
            } else {
                var v63 int = 0
                var v64 int64 = int64(v3)
                var v65 uint = (v45 | (v45 | v45))
                _ = v63
                _ = v64
                _ = v65
            }
            fmt.Println(v45, v45)
            var v66 string = "a"
            _ = v66
        }
        v67 := GenS4{F0: int32(v17), F1: "a", F2: (fi61 % (((fi61 + v16)) | 1))}
        for fi68 := 0; fi68 < 4; fi68++ {
            if v52 {
                fmt.Println(v45)
                delete(v5, (v4 * v4))
                v69 := GenS5{F0: (v2 != 67), F1: 0, F2: 1}
                v6 = append(v6, 2147483647)
                _ = v69
            }
            v16 = fi61
            if v52 {
                v4 = 803
                v70 := GenS6{F0: v1, F1: v4}
                if (!(v45 != (^(v45)))) {
                    v17 += (v17 - v17)
                } else {
                    if v52 {
                        var v71 int16 = 1000
                        if v52 {
                            if true {
                                var v72 int16 = (v71 / (((^(v71))) | 1))
                                _ = v72
                            } else {
                                if true {
                                    var v73 uint16 = v17
                                    v74 := GenS7{F0: v17}
                                    if len(v6) > 0 {
                                        v6[0] = uint32(v2)
                                    }
                                    _ = v73
                                    _ = v74
                                }
                                if (!((1e10 / 1e10) < 1.5)) {
                                    v1 = v1
                                    v1 = v1
                                }
                                v4 = (^(v4))
                            }
                            var v75 string = ("foo bar" + "")
                            fmt.Println(v75)
                            var v76 uint64 = 0
                            _ = v75
                            _ = v76
                        } else {
                            if ("\n\t" <= ("hello" + "")) {
                                v77 := GenS8{F0: (!false), F1: 0, F2: 393}
                                v78 := GenS9{F0: -43, F1: 336}
                                v5[v4] = "hello"
                                _ = v77
                                _ = v78
                            } else {
                                v45 += (v45 / (((v45 << ((v45) & 63))) | 1))
                                v17 = (v17 % ((v17) | 1))
                            }
                            var v79 uint = 142
                            v3 += v3
                            _ = v79
                        }
                        _ = v71
                    } else {
                        fmt.Println(v16, v1)
                        switch ("foo bar" + (("foo bar" + ("foo bar" + "\n\t")) + "a")) {
                        case "hello":
                            v3 = 31
                        case "foo bar":
                            v1 += v1
                        default:
                            v80 := GenS8{F0: v52, F1: -906, F2: (^(v4))}
                            v1 = v1
                            v45 -= (^(v45))
                            v3 += (v3 | v3)
                            if true {
                                if ((!(!v52)) && false) {
                                    v81 := GenS2{F0: v17, F1: v45, F2: v17}
                                    _ = v81
                                }
                                v82 := v5[v4]
                                _ = v82
                            }
                            _ = v80
                        }
                    }
                    v5[(^(v4))] = "hello"
                }
                _ = v70
            } else {
                v1 -= (v1 / ((v1) | 1))
                if v52 {
                    var v83 uint64 = uint64(v2)
                    switch ("" + "\n\t") {
                    case "":
                        fmt.Println(v1, fi68, fi68)
                    case "foo bar":
                        if (1e10 == (1.5 + 0.0)) {
                            var v84 string = ("a" + "hello")
                            v85 := GenS3{F0: int32(v2)}
                            _ = v84
                            _ = v85
                        }
                    }
                    _ = v83
                } else {
                    if ((v2 % ((v2) | 1)) <= v2) {
                        if v52 {
                            var v86 int64 = 42
                            var v87 uint32 = (v4 ^ (v4 / (((v4 & v4)) | 1)))
                            if (v52 && true) {
                                if v52 {
                                    v88 := GenS10{F0: (v4 & (v4 >> (((^(v45))) & 63))), F1: -128, F2: (-73 < (^(v3)))}
                                    var v89 int8 = (^(v3))
                                    v90 := GenS8{F0: true, F1: int16(v4), F2: v4}
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                } else {
                                    fmt.Println(v87, v52, v1)
                                    v4 += (v87 - v87)
                                    v52 = v52
                                }
                                var v91 int = (v16 & -61)
                                fmt.Println(v16)
                                _ = v91
                            }
                            _ = v86
                            _ = v87
                        }
                    }
                }
            }
            _ = fi68
        }
        _ = fi61
        _ = v67
    }
    for fi92 := 0; fi92 < 15; fi92++ {
        if v52 {
            delete(v5, v4)
            _ = len(v6)
        }
        _ = fi92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v52
    _ = v6
}
