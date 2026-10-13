package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 int; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int64; F1 uint32; F2 int32 }
func genFn0(p0 bool, p1 string, p2 uint8) int16 {
    var v6 uint8 = p2
    switch p1 {
    case "foo bar":
        var v7 bool = true
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 13; fi8++ {
            if p0 {
                var v9 bool = (v6 >= p2)
                v10 := GenS1{F0: -1, F1: 9223372036854775807}
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    }
    switch p1 {
    case "foo bar":
        v11 := []int{42, 127, -1}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 133
    v3 := make(map[int]bool)
    v4 := []string{"\n\t", ("a" + "a"), ("a" + ""), "a"}
    v5 := make(map[int]string)
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := v3[fi12]
        fmt.Println(v1, v13)
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int16 = -1
            v16 := GenS1{F0: 1, F1: 0}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    if false {
        var v17 int32 = 675
        var v18 int64 = -9223372036854775808
        _ = v17
        _ = v18
    } else {
        fmt.Println(v1, v2)
    }
    v1 += v1
    v19 := GenS1{F0: 1, F1: 1}
    if false {
        v20 := GenS1{F0: -1, F1: 42}
        _ = v20
    } else {
        v1 += v1
        var v21 uint64 = (v1 - (v1 / ((649) | 1)))
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(fi22, v1, v2)
        _ = fi22
    }
    v2 = v2
    for fi23 := 0; fi23 < 13; fi23++ {
        delete(v3, (^(fi23)))
        _ = fi23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        if ((42 < 1) && false) {
            v1 -= (^(v1))
            for fi25 := 0; fi25 < 14; fi25++ {
                if (-981 < 355) {
                    if (18446744073709551615 <= (v1 & v1)) {
                        v1 = v1
                    } else {
                        v2 += 376
                        if true {
                            fmt.Println(fi24)
                        } else {
                            switch "hello" {
                            case "foo bar":
                                var v26 uint8 = 128
                                _ = v26
                            case "a":
                                v27 := GenS1{F0: -1, F1: 0}
                                _ = v27
                            }
                            v1 = 334
                            v1 = v1
                            if true {
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    v28 := GenS1{F0: -1, F1: -1}
                                    v29 := GenS2{F0: 23}
                                    _ = v28
                                    _ = v29
                                }
                                switch "foo bar" {
                                case "hello":
                                    var v30 uint16 = v2
                                    _ = v30
                                }
                                if false {
                                    if false {
                                        v4 = append(v4, "")
                                    } else {
                                        v4 = append(v4, ("\n\t" + "\n\t"))
                                        var v31 int8 = 42
                                        var v32 int = (fi24 ^ (^(fi24)))
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 float32 = 1.5
                                    _ = v33
                                }
                            } else {
                                v34 := GenS3{F0: -2147483648}
                                var v35 uint16 = (v2 * 32767)
                                var v36 uint32 = 42
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                        }
                        _ = len(v5)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v1 -= v1
                    var v37 uint = 1
                    _ = v37
                } else {
                    v1 -= (^(v1))
                    v4 = append(v4, "hello")
                    fmt.Println(v2, v2, fi25)
                }
                switch "" {
                case "":
                    if true {
                        if false {
                            var v38 float64 = 3.14159
                            var v39 float64 = v38
                            switch (^(fi24)) {
                            case 32767:
                                v40 := GenS3{F0: 2147483647}
                                _ = v40
                            case 431:
                                if true {
                                    v1 -= v1
                                } else {
                                    v38 = ((v38 * 1.5) / v38)
                                    var v41 uint16 = (v2 << ((42) & 63))
                                    _ = v41
                                }
                                v3[128] = ((fi25 | (^(fi25))) != fi24)
                                v39 = v39
                            case 127:
                                fmt.Println(fi25, fi25)
                            default:
                                if false {
                                    v1 = v1
                                    v39 += -2.25
                                    v38 += v39
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                }
                            }
                            if (true || (1 > -76)) {
                                var v42 float64 = v39
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (!(!(235 > 2))) {
                                    var v43 bool = true
                                    v39 += (v38 * v42)
                                    var v44 int8 = -34
                                    _ = v43
                                    _ = v44
                                }
                                fmt.Println(fi25, fi24)
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                                _ = v42
                            } else {
                                v39 += 0.1
                                v2 += 42
                                var v45 int = fi25
                                _ = v45
                            }
                            v46 := GenS3{F0: 0}
                            _ = v38
                            _ = v39
                            _ = v46
                        } else {
                            v1 = (^(v1))
                            var v47 uint32 = 42
                            switch fi24 {
                            case 32768:
                                switch (("hello" + "foo bar") + "") {
                                case "foo bar":
                                    var v48 float64 = 0.1
                                    _ = v48
                                case "a":
                                    v4 = append(v4, "hello")
                                }
                            }
                            _ = v47
                        }
                        v49 := GenS2{F0: 100}
                        if (false || (!true)) {
                            v50 := GenS4{F0: 127}
                            v51 := GenS5{F0: 42}
                            if (!((-1 >= 42) && (true || false))) {
                                v2 = (v2 >> ((125) & 63))
                                fmt.Println(v1, fi25, fi25)
                                v1 = v1
                                v52 := GenS6{F0: 0, F1: 2}
                                _ = v52
                            }
                            switch "" {
                            case "\n\t":
                                var v53 float32 = 3.14159
                                _ = v53
                            case "foo bar":
                                v2 -= v2
                            default:
                                v54 := GenS5{F0: 176}
                                _ = v54
                            }
                            _ = v50
                            _ = v51
                        }
                        _ = v49
                    }
                default:
                    v2 -= (v2 + (v2 << ((42) & 63)))
                }
                _ = fi25
            }
        } else {
            if true {
                var v55 float32 = 1.5
                _ = v55
            } else {
                for fi56 := 0; fi56 < 8; fi56++ {
                    fmt.Println(v2, v2)
                    v57 := GenS1{F0: -2147483648, F1: 456}
                    var v58 uint16 = v2
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                var v59 float64 = ((0.0 / 1e10) - 0.0)
                for fi60 := 0; fi60 < 9; fi60++ {
                    v4 = append(v4, "\n\t")
                    var v61 uint = 1
                    _ = fi60
                    _ = v61
                }
                _ = v59
            }
            _ = len(v4)
        }
        fmt.Println(v1)
        v2 = 115
        _ = fi24
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        for fi63 := 0; fi63 < 19; fi63++ {
            var v64 uint32 = uint32(v1)
            _ = fi63
            _ = v64
        }
        v2 += 255
        _ = fi62
    }
    for fi65 := 0; fi65 < 14; fi65++ {
        if (162 != 255) {
            var v66 int64 = -577
            if (!false) {
                var v67 int16 = 600
                var v68 uint32 = 65535
                _ = v67
                _ = v68
            } else {
                var v69 uint64 = (v1 ^ (v1 >> ((2) & 63)))
                var v70 int64 = (v66 / ((v66) | 1))
                _ = v69
                _ = v70
            }
            v2 = (v2 + uint16(v1))
            _ = v66
        }
        _ = fi65
    }
    for fi71 := 0; fi71 < 1; fi71++ {
        delete(v5, 42)
        _ = fi71
    }
    v2 = v2
    var v72 uint8 = 255
    if (-57 > -56) {
        v1 += (v1 * v1)
        var v73 string = "\n\t"
        for fi74 := 0; fi74 < 9; fi74++ {
            switch fi74 {
            case 127:
                if false {
                    fmt.Println(v72, v1)
                    for fi75 := 0; fi75 < 5; fi75++ {
                        if (2147483647 < -1) {
                            var v76 uint32 = 319
                            if ((fi74 | fi75) > int(v72)) {
                                if ((v2 + (^(v2))) != 65535) {
                                    if len(v4) > 0 {
                                        v4[0] = v73
                                    }
                                } else {
                                    if false {
                                        switch 513 {
                                        case 78:
                                            var v77 string = v73
                                            _ = v77
                                        default:
                                            v2 = v2
                                        }
                                        if (100 < 123) {
                                            v78 := GenS7{F0: true, F1: 3}
                                            _ = v78
                                        } else {
                                            v73 = ("" + ((("" + (v73 + v73)) + "") + ((("foo bar" + v73) + "") + "")))
                                            var v79 int = (^(fi74))
                                            if false {
                                                v72 = v72
                                                v2 = (v2 | (v2 * (v2 * (^(v2)))))
                                                fmt.Println(v73)
                                            } else {
                                                var v80 uint64 = v1
                                                switch v73 {
                                                case "hello":
                                                    if ((int8(v80) <= 42) || (false || ((!(0 <= int16(v72))) || true))) {
                                                        var v81 int16 = -792
                                                        v82 := GenS6{F0: fi75, F1: 0}
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                case "foo bar":
                                                    if len(v4) > 0 {
                                                        v4[0] = v73
                                                    }
                                                default:
                                                    switch fi74 {
                                                    case -1:
                                                        switch fi74 {
                                                        case 9223372036854775807:
                                                            v83 := v5[(fi74 ^ v79)]
                                                            _ = v83
                                                        case -832:
                                                            fmt.Println(fi75)
                                                        case -2147483648:
                                                            v84 := GenS2{F0: 0}
                                                            _ = v84
                                                        default:
                                                            v79 += (fi74 & fi74)
                                                        }
                                                    }
                                                }
                                                if ((!(true && true)) && (false || true)) {
                                                    v1 = (^(v80))
                                                    v85 := GenS8{F0: 9223372036854775807, F1: v72, F2: 0}
                                                    fmt.Println(v80, v72, v79)
                                                    _ = v85
                                                } else {
                                                    var v86 int16 = genFn0(((true || false) && (false || false)), "", v72)
                                                    _ = v86
                                                }
                                                var v87 float32 = (3.14159 / (1e10 / float32(3.14159)))
                                                _ = v80
                                                _ = v87
                                            }
                                            _ = v79
                                        }
                                        v2 += v2
                                        delete(v5, -558)
                                        var v88 float32 = 3.14159
                                        _ = v88
                                    }
                                    var v89 bool = false
                                    var v90 int16 = 1
                                    v91 := GenS9{F0: 42, F1: 0, F2: 1}
                                    if len(v4) > 0 {
                                        v4[0] = ("a" + v73)
                                    }
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                            } else {
                                _ = len(v5)
                            }
                            var v92 int32 = -1
                            _ = v76
                            _ = v92
                        }
                        _ = len(v5)
                        v2 -= (v2 | (^(v2)))
                        _ = fi75
                    }
                    v2 -= (^(v2))
                } else {
                    for fi93 := 0; fi93 < 3; fi93++ {
                        var v94 int32 = 2147483647
                        _ = fi93
                        _ = v94
                    }
                    v72 -= (v72 & v72)
                }
            case -647:
                var v95 float64 = float64(float32(0.0))
                _ = v95
            }
            var v96 int16 = 0
            _ = fi74
            _ = v96
        }
        _ = v73
    } else {
        fmt.Println(v2, v1, v1)
        var v97 uint8 = (v72 - v72)
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
}
