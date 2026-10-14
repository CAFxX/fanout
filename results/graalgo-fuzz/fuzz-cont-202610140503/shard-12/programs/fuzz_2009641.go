package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int16; F1 uint; F2 uint }
func genFn0(p0 uint, p1 int32) int16 {
    v7 := GenS1{F0: 32768}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int = 4294967295
        switch "a" {
        case "":
            v10 := GenS1{F0: fi8}
            _ = v10
        case "a":
            if false {
                v11 := []int{(fi8 + fi8), (fi8 + int(p0)), v9, v9}
                for fi12 := 0; fi12 < 5; fi12++ {
                    v13 := GenS1{F0: 0}
                    var v14 int8 = -128
                    if (2147483647 != 2147483647) {
                        v9 -= 42
                        var v15 int8 = v14
                        v11 = append(v11, (v9 * (-(v9))))
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                if len(v11) > 0 {
                    v11[0] = (v9 * (fi8 + (fi8 + (v9 | (fi8 ^ 821)))))
                }
                _ = v11
            }
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 float32 = (-2.25 - (3.14159 + (0.1 + (((0.1 * (3.14159 * 3.14159)) / (1.5 / 1.5)) + float32(1.5)))))
    var v2 int16 = 404
    var v3 uint64 = 4294967295
    v4 := []int{-9223372036854775808, 1, 0}
    var v5 int32 = 0
    var v6 bool = false
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 float32 = v1
        v6 = ((false && v6) || (v6 && (255 < 816)))
        _ = fi16
        _ = v17
    }
    v5 = 466
    v6 = v6
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v2, fi18, v2)
            _ = fi19
        }
        v3 = 897
        for fi20 := 0; fi20 < 6; fi20++ {
            if (42 != 42) {
                var v21 float32 = v1
                _ = len(v4)
                _ = v21
            } else {
                fmt.Println(v5, fi18, v1)
                v2 = 1
                v6 = (v1 >= float32(float64((v1 * 0.0))))
            }
            v4 = append(v4, fi20)
            _ = fi20
        }
        _ = fi18
    }
    var v22 uint64 = (v3 >> ((232) & 63))
    for fi23 := 0; fi23 < 19; fi23++ {
        if (!v6) {
            v6 = v6
            v24 := GenS2{F0: v5}
            for fi25 := 0; fi25 < 19; fi25++ {
                if (!(1 <= 289)) {
                    if (-17 < int8(v2)) {
                        var v26 int = int(v2)
                        var v27 string = ("foo bar" + "a")
                        _ = v26
                        _ = v27
                    } else {
                        var v28 bool = (((float64(v1) / 1e10) * float64(v1)) > ((0.0 - 1e10) / -2.25))
                        if (!v6) {
                            v29 := GenS1{F0: (fi25 + (fi25 * fi25))}
                            _ = v29
                        } else {
                            var v30 string = ""
                            switch (("hello" + (v30 + (v30 + v30))) + ("a" + v30)) {
                            case "":
                                v22 = (^(v22))
                            case "\n\t":
                                var v31 int16 = (v2 * 1000)
                                _ = v31
                            }
                            var v32 uint = 65535
                            _ = v30
                            _ = v32
                        }
                        _ = v28
                    }
                    switch "\n\t" {
                    case "foo bar":
                        v4 = append(v4, (fi25 - (^(fi23))))
                    case "":
                        v3 = v22
                    default:
                        v33 := GenS3{F0: v6, F1: v5}
                        v6 = v6
                        _ = v33
                    }
                    if (!false) {
                        if len(v4) > 0 {
                            v4[0] = fi23
                        }
                        fmt.Println(fi23, v22)
                        var v34 uint8 = 1
                        v2 = (v2 ^ genFn0(362, (v5 + v5)))
                        _ = v34
                    }
                } else {
                    if (1.5 > 1e10) {
                        fmt.Println(v6, v5, fi23)
                        var v35 int16 = (v2 ^ v2)
                        v6 = ((float32(float64(3.14159)) > v1) && (42 <= 2147483647))
                        _ = v35
                    } else {
                        switch "" {
                        case "a":
                            switch -9223372036854775808 {
                            case -42:
                                v4 = append(v4, 1)
                            case -9223372036854775808:
                                var v36 string = "a"
                                switch fi25 {
                                case 9223372036854775807:
                                    v37 := GenS1{F0: (fi23 ^ fi23)}
                                    _ = v37
                                case 0:
                                    fmt.Println(v36, fi25)
                                default:
                                    v1 = v1
                                }
                                var v38 int16 = -1
                                _ = v36
                                _ = v38
                            case 128:
                                if true {
                                    fmt.Println(fi25, v3)
                                }
                            default:
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    fmt.Println(v1, v3)
                                case "foo bar":
                                    fmt.Println(v22, v22, v3)
                                case "hello":
                                    fmt.Println(fi23, v2, v3)
                                }
                            }
                        case "":
                            var v39 uint64 = 1
                            _ = v39
                        case "foo bar":
                            fmt.Println(v1, v2)
                        }
                        if ((^(fi25)) <= fi23) {
                            switch fi23 {
                            case 32767:
                                _ = len(v4)
                            case 866:
                                var v40 int8 = 42
                                _ = v40
                            case 128:
                                v6 = true
                            default:
                                switch "hello" {
                                case "foo bar":
                                    switch ("foo bar" + "\n\t") {
                                    case "":
                                        if len(v4) > 0 {
                                            v4[0] = int(v3)
                                        }
                                    default:
                                        _ = len(v4)
                                    }
                                }
                            }
                            var v41 int = (fi23 & 986)
                            v42 := GenS1{F0: fi25}
                            _ = v41
                            _ = v42
                        } else {
                            v1 = 0.0
                            if ((v5 | v5) >= 42) {
                                var v43 float64 = 1.5
                                if (v6 || v6) {
                                    v4 = append(v4, (fi25 * (fi23 / ((2147483647) | 1))))
                                }
                                _ = v43
                            }
                            if (2 != 0) {
                                var v44 uint64 = v22
                                _ = v44
                            }
                        }
                    }
                }
                switch "hello" {
                case "foo bar":
                    switch "" {
                    case "":
                        fmt.Println(fi25, v6, v3)
                    default:
                        v1 += (v1 - v1)
                    }
                    v45 := GenS4{F0: 151}
                    if true {
                        fmt.Println(fi25, fi23, v5)
                        v22 -= (v22 | v3)
                        if true {
                            v1 = v1
                            var v46 uint16 = 235
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v46
                        } else {
                            v3 = v3
                            if v6 {
                                fmt.Println(v1)
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (!(((-(v2)) >= -426) && v6)) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v47 uint16 = 450
                                        _ = v47
                                    case "a":
                                        if v6 {
                                            v48 := GenS2{F0: (v5 + (-(v5)))}
                                            v3 = 329
                                            switch (fi23 & 0) {
                                            case -126:
                                                v49 := GenS1{F0: (fi25 / ((fi25) | 1))}
                                                _ = v49
                                            default:
                                                v2 -= genFn0(2, (-(v5)))
                                            }
                                            _ = len(v4)
                                            v1 = (-2.25 + 1.5)
                                            fmt.Println(v6, v3, v22)
                                            _ = v48
                                        }
                                    case "":
                                        v50 := GenS5{F0: (fi23 >> ((65535) & 63))}
                                        _ = v50
                                    }
                                } else {
                                    v51 := GenS3{F0: (65535 != 0), F1: (-(v5))}
                                    v2 = v2
                                    switch "" {
                                    case "\n\t":
                                        var v52 int32 = (v5 | v5)
                                        _ = v52
                                    case "":
                                        if v6 {
                                            v2 += v2
                                            fmt.Println(v22, fi25, v1)
                                        }
                                    case "a":
                                        v22 = (^(v3))
                                    default:
                                        fmt.Println(v2, v6)
                                    }
                                    _ = v51
                                }
                                v5 -= (v5 % (((-(v5))) | 1))
                            }
                        }
                    } else {
                        v1 = v1
                        var v53 uint16 = 917
                        _ = v53
                    }
                    _ = v45
                }
                if v6 {
                    var v54 uint32 = 111
                    switch "a" {
                    case "":
                        fmt.Println(v6, fi23, v6)
                    case "hello":
                        v55 := GenS5{F0: (fi25 << ((2) & 63))}
                        _ = v55
                    case "foo bar":
                        if (117 < 38) {
                            v2 = 0
                        } else {
                            v3 = (v3 / ((v22) | 1))
                        }
                        fmt.Println(v22, v22, v1)
                    default:
                        var v56 uint = 1
                        _ = v56
                    }
                    _ = v54
                } else {
                    var v57 uint16 = 42
                    if len(v4) > 0 {
                        v4[0] = fi23
                    }
                    var v58 int32 = v5
                    _ = v57
                    _ = v58
                }
                _ = fi25
            }
            _ = v24
        }
        if (v6 || (32767 <= 2)) {
            if v6 {
                var v59 int = (fi23 | (fi23 - 4294967295))
                _ = v59
            } else {
                switch (fi23 >> ((286) & 63)) {
                case -449:
                    fmt.Println(v3, fi23)
                default:
                    for fi60 := 0; fi60 < 3; fi60++ {
                        switch "a" {
                        case "\n\t":
                            v3 = (^(v3))
                            switch "\n\t" {
                            case "\n\t":
                                v4 = append(v4, (fi60 << ((42) & 63)))
                            case "":
                                switch 127 {
                                case 32768:
                                    var v61 int8 = -1
                                    _ = v61
                                default:
                                    var v62 float32 = v1
                                    _ = v62
                                }
                            case "foo bar":
                                switch fi60 {
                                case 315:
                                    if (1 != 853) {
                                        v63 := GenS6{F0: 255, F1: 0}
                                        v4 = append(v4, fi60)
                                        v64 := GenS7{F0: 1}
                                        _ = v63
                                        _ = v64
                                    } else {
                                        if v6 {
                                            v2 = v2
                                        }
                                    }
                                case 594:
                                    var v65 string = "foo bar"
                                    _ = v65
                                default:
                                    var v66 float64 = (float64((((1e10 / (v1 + float32(0.1))) - 1e10) + v1)) - 0.1)
                                    _ = v66
                                }
                            default:
                                if v6 {
                                    fmt.Println(fi60, v22, fi23)
                                }
                            }
                            if v6 {
                                if v6 {
                                    var v67 uint64 = v3
                                    _ = v67
                                }
                            }
                        case "hello":
                            v68 := GenS3{F0: false, F1: (v5 >> ((28) & 63))}
                            _ = v68
                        }
                        fmt.Println(v5)
                        _ = fi60
                    }
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if v6 {
                var v69 int64 = 105
                _ = v69
            }
        } else {
            for fi70 := 0; fi70 < 7; fi70++ {
                if v6 {
                    var v71 float64 = -2.25
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v71
                }
                fmt.Println(v5, v1, v5)
                var v72 float32 = (v1 - (v1 + v1))
                _ = fi70
                _ = v72
            }
            v1 -= 0.1
            v73 := GenS8{F0: 42, F1: v2}
            _ = v73
        }
        v74 := GenS9{F0: 65535, F1: 0}
        _ = fi23
        _ = v74
    }
    for fi75 := 0; fi75 < 16; fi75++ {
        for fi76 := 0; fi76 < 20; fi76++ {
            if (v6 && true) {
                var v77 uint32 = 2147483647
                _ = v77
            } else {
                var v78 uint8 = 162
                var v79 uint = uint(v78)
                v4 = append(v4, (fi75 >> ((18446744073709551615) & 63)))
                _ = v78
                _ = v79
            }
            var v80 uint64 = (^(v22))
            v22 = 42
            v1 += v1
            v3 = v3
            _ = fi76
            _ = v80
        }
        _ = fi75
    }
    fmt.Println(v6, v5)
    v81 := GenS6{F0: 0, F1: 9223372036854775807}
    var v82 string = "hello"
    v83 := GenS9{F0: 2147483647, F1: 91}
    v84 := GenS1{F0: -155}
    for fi85 := 0; fi85 < 13; fi85++ {
        var v86 string = v82
        v87 := GenS5{F0: (^(fi85))}
        switch "foo bar" {
        case "foo bar":
            for fi88 := 0; fi88 < 17; fi88++ {
                _ = len(v4)
                _ = fi88
            }
        case "a":
            if ((fi85 << ((0) & 63)) < (fi85 & fi85)) {
                v89 := GenS1{F0: (fi85 - fi85)}
                if v6 {
                    v82 = ("foo bar" + "foo bar")
                    v90 := GenS10{F0: (v2 | (v2 & -754)), F1: 177, F2: 42}
                    _ = v90
                }
                _ = v89
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi85
        _ = v86
        _ = v87
    }
    v3 = v22
    for fi91 := 0; fi91 < 4; fi91++ {
        fmt.Println(v82, v2, v3)
        _ = fi91
    }
    for fi92 := 0; fi92 < 14; fi92++ {
        for fi93 := 0; fi93 < 11; fi93++ {
            if len(v4) > 0 {
                v4[0] = 0
            }
            v2 = (v2 << ((42) & 63))
            _ = fi93
        }
        if len(v4) > 0 {
            v4[0] = (fi92 - fi92)
        }
        _ = fi92
    }
    v5 = v5
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
