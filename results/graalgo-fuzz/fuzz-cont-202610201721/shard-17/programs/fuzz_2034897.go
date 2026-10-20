package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint8; F1 int8; F2 bool }
type GenS8 struct { F0 uint; F1 bool; F2 uint32 }
type GenS9 struct { F0 int32; F1 int8; F2 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 int32; F2 uint8 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]int)
    var v3 uint8 = 136
    var v4 int64 = 42
    var v5 string = ""
    v6 := make(map[uint32]uint32)
    if (false && false) {
        var v7 int16 = 1000
        if false {
            switch -1 {
            case -2147483648:
                v8 := []int{}
                _ = v8
            case 1:
                v9 := GenS1{F0: (v7 >> ((18446744073709551615) & 63))}
                _ = v9
            }
            v7 = (v7 * (v7 + 42))
        }
        v10 := GenS1{F0: (^(v7))}
        _ = v10
        _ = v7
    }
    var v11 string = v5
    for fi12 := 0; fi12 < 14; fi12++ {
        v4 -= 42
        v13 := v2[1]
        _ = fi12
        _ = v13
    }
    v11 = v11
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int32 = 2147483647
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := []int{}
            var v18 int64 = (-(v4))
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    delete(v2, 4294967295)
    switch ("a" + (v5 + v5)) {
    case "":
        fmt.Println(v3, v5, v11)
    default:
        var v19 uint64 = 4294967295
        _ = v19
    }
    if false {
        if ("foo bar" < ((v11 + v11) + v11)) {
            v4 += 0
        }
    }
    v20 := []int{1, 4294967295, 4294967295, -683}
    _ = len(v2)
    v21 := v6[979]
    fmt.Println(v21, v3, v3)
    if ((!false) || true) {
        v22 := GenS2{F0: -887, F1: v1, F2: 42}
        v1 -= (^(v1))
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 uint32 = v21
                _ = fi24
                _ = v25
            }
            v26 := GenS2{F0: (fi23 / ((fi23) | 1)), F1: (v1 & 65535), F2: 1}
            var v27 int64 = (v4 + v4)
            _ = fi23
            _ = v26
            _ = v27
        }
        _ = v22
    } else {
        fmt.Println(v1)
    }
    if true {
        switch "\n\t" {
        case "\n\t":
            v6[(^(v21))] = 2147483647
        case "a":
            if (!false) {
                var v28 int32 = 42
                _ = v28
            } else {
                fmt.Println(v3)
                for fi29 := 0; fi29 < 4; fi29++ {
                    if (-1 < -1) {
                        for fi30 := 0; fi30 < 14; fi30++ {
                            var v31 int = fi29
                            var v32 uint64 = 2
                            var v33 uint64 = 752
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        fmt.Println(v21, v3)
                        if len(v20) > 0 {
                            v20[0] = fi29
                        }
                    }
                    v11 = v5
                    v21 = (v21 % ((714) | 1))
                    _ = fi29
                }
                switch v5 {
                case "\n\t":
                    fmt.Println(v5, v3)
                case "hello":
                    var v34 int = -697
                    _ = v34
                case "foo bar":
                    _ = len(v20)
                }
            }
        case "hello":
            for fi35 := 0; fi35 < 7; fi35++ {
                if true {
                    switch "hello" {
                    case "a":
                        var v36 uint64 = 18446744073709551615
                        _ = v36
                    case "foo bar":
                        v5 = ("\n\t" + v11)
                    default:
                        v20 = append(v20, (fi35 & 42))
                    }
                    if (!(v1 != 2)) {
                        var v37 int32 = 1
                        if (-1 > -76) {
                            var v38 int8 = 102
                            var v39 bool = (2 <= 42)
                            v11 = v5
                            _ = v38
                            _ = v39
                        } else {
                            v20 = append(v20, (fi35 & (fi35 ^ 2147483647)))
                        }
                        v40 := GenS2{F0: (fi35 * (fi35 - fi35)), F1: (v1 * v1), F2: 255}
                        _ = v37
                        _ = v40
                    } else {
                        v4 -= (v4 << ((uint(v3)) & 63))
                    }
                    switch (v5 + (v5 + "a")) {
                    case "a":
                        var v41 uint = 255
                        _ = v41
                    case "hello":
                        if (-1 == int16(v1)) {
                            fmt.Println(v5, v4)
                        } else {
                            if (!false) {
                                var v42 string = (v11 + ((("a" + v11) + "hello") + v5))
                                v3 += (^(v3))
                                fmt.Println(fi35)
                                _ = v42
                            } else {
                                v43 := GenS1{F0: -924}
                                for fi44 := 0; fi44 < 9; fi44++ {
                                    var v45 int64 = -633
                                    _ = fi44
                                    _ = v45
                                }
                                v46 := GenS3{F0: (v3 * (v3 | v3))}
                                _ = v43
                                _ = v46
                            }
                            fmt.Println(v4, v5)
                            fmt.Println(fi35, v11, v3)
                        }
                    default:
                        for fi47 := 0; fi47 < 19; fi47++ {
                            switch fi47 {
                            case -1:
                                v48 := GenS4{F0: 2}
                                _ = v48
                            case -250:
                                v1 += v1
                            }
                            v49 := GenS5{F0: (-(fi35))}
                            var v50 int32 = -643
                            _ = fi47
                            _ = v49
                            _ = v50
                        }
                    }
                }
                _ = fi35
            }
        }
        fmt.Println(v4)
    } else {
        for fi51 := 0; fi51 < 13; fi51++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            if (42 != 1) {
                v5 = ("hello" + v5)
                if false {
                    v5 = ("hello" + v11)
                    fmt.Println(v4, fi51, fi51)
                } else {
                    for fi52 := 0; fi52 < 3; fi52++ {
                        if len(v20) > 0 {
                            v20[0] = (fi52 | (fi52 % (((^(fi52))) | 1)))
                        }
                        if true {
                            var v53 uint8 = v3
                            var v54 int8 = 42
                            _ = v53
                            _ = v54
                        } else {
                            if (true || true) {
                                fmt.Println(v11)
                                var v55 float32 = 0.1
                                var v56 uint16 = v1
                                _ = v55
                                _ = v56
                            }
                            fmt.Println(v21, v5, fi52)
                            var v57 float64 = (0.0 * 0.0)
                            _ = v57
                        }
                        _ = fi52
                    }
                }
                for fi58 := 0; fi58 < 19; fi58++ {
                    v59 := GenS6{F0: v11, F1: 4294967295, F2: v1}
                    v3 += (^(v3))
                    var v60 uint = 0
                    v3 -= v3
                    var v61 int64 = 1
                    var v62 string = (v5 + "")
                    _ = fi58
                    _ = v59
                    _ = v60
                    _ = v61
                    _ = v62
                }
            } else {
                if (false || ((((^(v4)) != v4) && false) && (42 > 1))) {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    if false {
                        v21 -= (v21 << ((4294967295) & 63))
                        for fi63 := 0; fi63 < 1; fi63++ {
                            switch "foo bar" {
                            case "hello":
                                var v64 int32 = 42
                                _ = v64
                            case "a":
                                if len(v20) > 0 {
                                    v20[0] = fi51
                                }
                            }
                            _ = len(v20)
                            var v65 uint = 255
                            _ = fi63
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 8; fi66++ {
                            var v67 int = fi66
                            _ = fi66
                            _ = v67
                        }
                    } else {
                        v5 = v11
                    }
                    switch fi51 {
                    case 32767:
                        var v68 uint64 = 4294967295
                        _ = v68
                    }
                } else {
                    for fi69 := 0; fi69 < 12; fi69++ {
                        v70 := GenS6{F0: v11, F1: 1, F2: v1}
                        v6[(v21 | (v21 * 65535))] = (v21 + v21)
                        if true {
                            var v71 int32 = -2147483648
                            _ = v71
                        } else {
                            v72 := GenS6{F0: "foo bar", F1: 830, F2: 2}
                            v3 -= (^(v3))
                            _ = v72
                        }
                        _ = fi69
                        _ = v70
                    }
                }
                var v73 float32 = (((0.0 * 0.1) * 1e10) / 1e10)
                v5 = "foo bar"
                _ = v73
            }
            _ = fi51
        }
        for fi74 := 0; fi74 < 6; fi74++ {
            var v75 float64 = ((0.1 / (0.1 / 1e10)) + 1.5)
            for fi76 := 0; fi76 < 2; fi76++ {
                v3 -= v3
                switch fi74 {
                case 128:
                    if false {
                        if (183 == v4) {
                            if true {
                                var v77 string = "hello"
                                _ = v77
                            } else {
                                v78 := GenS7{F0: v3, F1: -121, F2: (-1 < 1000)}
                                v4 += -1
                                switch (fi74 * (fi74 / (((fi74 - int(v21))) | 1))) {
                                case 1:
                                    if true {
                                        var v79 int32 = 1
                                        var v80 uint = 1
                                        switch 0 {
                                        case 191:
                                            var v81 float32 = (float32(float64(1e10)) * (0.1 * (1.5 / (0.1 * 1e10))))
                                            var v82 string = v5
                                            var v83 uint = (v80 | v80)
                                            _ = v81
                                            _ = v82
                                            _ = v83
                                        case -59:
                                            if true {
                                                v84 := GenS8{F0: (v80 * v80), F1: true, F2: (v21 >> (((^(v80))) & 63))}
                                                v11 = (v11 + v5)
                                                v85 := GenS9{F0: (v79 - -393), F1: -1, F2: (v4 * (-(v4)))}
                                                _ = v84
                                                _ = v85
                                            } else {
                                                v75 -= 1e10
                                            }
                                        default:
                                            if ((!(false || true)) && true) {
                                                if ((v21 / ((v21) | 1)) > v21) {
                                                    v20 = append(v20, 127)
                                                    v86 := GenS10{F0: (v80 % ((v80) | 1))}
                                                    var v87 int64 = int64(v1)
                                                    _ = v86
                                                    _ = v87
                                                }
                                                v88 := GenS7{F0: v3, F1: 94, F2: ((v21 + (v21 % ((uint32(v1)) | 1))) <= 42)}
                                                _ = v88
                                            } else {
                                                delete(v2, 2147483648)
                                                var v89 string = v11
                                                _ = v89
                                            }
                                        }
                                        _ = v79
                                        _ = v80
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            if (((1e10 / 0.1) / (1.5 / (1.5 + (-2.25 + 1.5)))) > (1e10 + (3.14159 + (float32(-2.25) * (float32(v75) + 1e10))))) {
                                                var v90 uint64 = 741
                                                v91 := GenS3{F0: (v3 ^ uint8(v1))}
                                                if (!(3.14159 > float64(((float32(1e10) + (float32(0.1) + -2.25)) + 0.1)))) {
                                                    v92 := GenS1{F0: 0}
                                                    _ = v92
                                                } else {
                                                    var v93 uint8 = (v3 + v3)
                                                    v11 = (v11 + "a")
                                                    _ = v93
                                                }
                                                _ = v90
                                                _ = v91
                                            }
                                            if (false || (1 != -1)) {
                                                fmt.Println(v11, v5)
                                                switch (fi74 % ((32767) | 1)) {
                                                case 127:
                                                    if true {
                                                        v21 = v21
                                                    }
                                                case 1:
                                                    if false {
                                                        var v94 uint32 = (v21 >> ((621) & 63))
                                                        var v95 bool = false
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        v75 = 0.0
                                                        _ = len(v20)
                                                        var v96 int16 = 644
                                                        _ = v96
                                                    }
                                                }
                                            }
                                            var v97 string = v11
                                            _ = v97
                                        default:
                                            v1 = (^(v1))
                                        }
                                        v3 -= 181
                                    }
                                case 32768:
                                    v98 := GenS11{F0: 1000, F1: 42, F2: (v3 + 55)}
                                    _ = v98
                                case -9223372036854775808:
                                    fmt.Println(v5, v11, fi76)
                                default:
                                    v4 -= -452
                                }
                                _ = v78
                            }
                            var v99 int = fi76
                            if ((-9223372036854775808 == v4) && (((-1 <= -1) || true) && false)) {
                                if len(v20) > 0 {
                                    v20[0] = fi74
                                }
                                v4 -= 9223372036854775807
                                v21 = (v21 - (v21 / ((v21) | 1)))
                                fmt.Println(v5, v11)
                            }
                            _ = v99
                        } else {
                            var v100 bool = true
                            _ = v100
                        }
                        v11 = ""
                        if (3.14159 != 3.14159) {
                            var v101 uint64 = 2
                            _ = v101
                        } else {
                            v11 = v11
                        }
                    }
                case 127:
                    v3 += uint8(v1)
                case 32767:
                    var v102 string = v5
                    _ = v102
                default:
                    fmt.Println(v3, v75, v21)
                }
                _ = fi76
            }
            _ = fi74
            _ = v75
        }
        v103 := GenS9{F0: -2147483648, F1: 57, F2: 9223372036854775807}
        _ = v103
    }
    var v104 int32 = -1
    var v105 string = v11
    for fi106 := 0; fi106 < 13; fi106++ {
        var v107 uint32 = (^(v21))
        _ = fi106
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v104
    _ = v105
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
