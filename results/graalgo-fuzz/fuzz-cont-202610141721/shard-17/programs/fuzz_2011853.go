package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS4 struct { F0 int32; F1 uint8; F2 bool }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS8 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint32, p1 int64, p2 uint32) uint64 {
    fmt.Println(p1, p0)
    switch "hello" {
    case "a":
        fmt.Println(p1, p0)
    case "":
        if (2 <= 127) {
            var v7 uint32 = 555
            _ = v7
        } else {
            var v8 uint16 = 42
            switch 32768 {
            case 9223372036854775807:
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint = 0
                    v11 := GenS1{F0: (v10 | v10)}
                    var v12 int64 = p1
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                p1 += 9223372036854775807
                v8 += uint16(p2)
                _ = fi13
            }
            _ = v8
        }
    default:
        if ((!(732 <= 0)) || false) {
            v14 := []int{344}
            if (-1 > -992) {
                for fi15 := 0; fi15 < 7; fi15++ {
                    p0 -= (p2 / ((p0) | 1))
                    p0 -= 1
                    _ = fi15
                }
                var v16 string = "a"
                p2 = (p2 - p2)
                _ = v16
            } else {
                var v17 uint = 65535
                var v18 uint = (v17 << (((v17 | v17)) & 63))
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                p1 -= -9223372036854775808
                p2 -= (p0 ^ p2)
                _ = fi19
            }
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 736
}
func genFn1() uint {
    var v20 int = 0
    if (!true) {
        var v21 bool = false
        var v22 int32 = 380
        _ = v21
        _ = v22
    } else {
        v20 -= (v20 | -2147483648)
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 float64 = 0.1
            v20 = 1
            if false {
                v24 += v24
                fmt.Println(v20, v20)
                v24 = 0.1
            }
            _ = fi23
            _ = v24
        }
        v20 = -1
    }
    _ = v20
    return 18446744073709551615
}

func main() {
    var v1 int64 = 888
    v2 := make(map[string]bool)
    var v3 int32 = 0
    var v4 int8 = -119
    var v5 float32 = 3.14159
    var v6 bool = false
    v25 := []int{42}
    var v26 uint16 = 0
    if (!v6) {
        var v27 uint16 = (v26 - (v26 ^ (v26 << ((4294967295) & 63))))
        fmt.Println(v26, v3)
        _ = v27
    } else {
        if true {
            var v28 float32 = v5
            _ = v28
        }
        v6 = v6
        if (128 < 42) {
            if true {
                fmt.Println(v26)
                v29 := GenS2{F0: 1}
                if (v6 && (231 > 95)) {
                    switch -9223372036854775808 {
                    case -182:
                        switch ("foo bar" + ("\n\t" + ("" + "foo bar"))) {
                        case "hello":
                            v1 -= 104
                        case "":
                            for fi30 := 0; fi30 < 20; fi30++ {
                                if false {
                                    if v6 {
                                        v3 = 0
                                        v26 += 32767
                                    } else {
                                        v2["a"] = v6
                                    }
                                    if v6 {
                                        _ = len(v25)
                                    } else {
                                        delete(v2, ("foo bar" + "foo bar"))
                                        for fi31 := 0; fi31 < 4; fi31++ {
                                            v26 -= v26
                                            _ = fi31
                                        }
                                        if v6 {
                                            if v6 {
                                                delete(v2, "hello")
                                            }
                                            switch -705 {
                                            case 653:
                                                v32 := GenS2{F0: 1}
                                                _ = v32
                                            default:
                                                fmt.Println(fi30, v1, fi30)
                                            }
                                            var v33 float64 = (3.14159 * (float64((v5 * v5)) - (1e10 - (-2.25 * 1.5))))
                                            _ = v33
                                        }
                                    }
                                }
                                _ = fi30
                            }
                        case "foo bar":
                            fmt.Println(v6)
                        }
                    case -548:
                        var v34 bool = true
                        _ = v34
                    default:
                        if (!(65535 < genFn1())) {
                            switch -27 {
                            case 146:
                                v1 = (-(v1))
                            case 0:
                                var v35 int32 = v3
                                _ = v35
                            default:
                                if (genFn1() <= uint(v26)) {
                                    fmt.Println(v6, v4)
                                    if v6 {
                                        if false {
                                            for fi36 := 0; fi36 < 2; fi36++ {
                                                fmt.Println(v4, v4)
                                                if true {
                                                    var v37 string = "hello"
                                                    if ((v5 <= v5) && v6) {
                                                        var v38 string = v37
                                                        _ = v38
                                                    }
                                                    _ = v37
                                                }
                                                v3 += 1
                                                _ = fi36
                                            }
                                        } else {
                                            var v39 uint32 = 202
                                            fmt.Println(v1, v6)
                                            if v6 {
                                                switch "a" {
                                                case "foo bar":
                                                    v40 := GenS2{F0: v3}
                                                    _ = v40
                                                case "hello":
                                                    v41 := GenS3{F0: 18446744073709551615, F1: 255, F2: 1000}
                                                    _ = v41
                                                case "\n\t":
                                                    if v6 {
                                                        v42 := GenS4{F0: (v3 | int32(v39)), F1: 76, F2: (!(-1 <= -1))}
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v43 int16 = 32767
                                                            _ = v43
                                                        case "":
                                                            var v44 uint32 = 2147483647
                                                            _ = v44
                                                        default:
                                                            v2["hello"] = (1e10 >= float32(3.14159))
                                                        }
                                                        _ = v42
                                                    } else {
                                                        var v45 float64 = -2.25
                                                        var v46 string = ""
                                                        v4 += v4
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                default:
                                                    var v47 uint32 = 653
                                                    _ = v47
                                                }
                                            }
                                            _ = v39
                                        }
                                        var v48 bool = true
                                        for fi49 := 0; fi49 < 16; fi49++ {
                                            var v50 uint = genFn1()
                                            v48 = v6
                                            _ = fi49
                                            _ = v50
                                        }
                                        _ = v48
                                    }
                                    var v51 int32 = (v3 >> ((genFn1()) & 63))
                                    _ = v51
                                } else {
                                    if v6 {
                                        if v6 {
                                            if v6 {
                                                switch 1 {
                                                case -829:
                                                    for fi52 := 0; fi52 < 16; fi52++ {
                                                        switch -19 {
                                                        case 0:
                                                            if v6 {
                                                                v53 := GenS5{F0: 2, F1: 2, F2: 506}
                                                                if (!v6) {
                                                                    if ((0.0 / 1.5) >= -2.25) {
                                                                        v54 := GenS6{F0: (v4 % ((0) | 1))}
                                                                        _ = v54
                                                                    } else {
                                                                        v5 += v5
                                                                        v55 := GenS2{F0: (^(v3))}
                                                                        _ = v55
                                                                    }
                                                                    for fi56 := 0; fi56 < 20; fi56++ {
                                                                        v26 += v26
                                                                        v6 = ((128 != 232) && ((v1 % (((v1 + -513)) | 1)) <= (v1 ^ v1)))
                                                                        v57 := GenS3{F0: 720, F1: 94, F2: -1}
                                                                        _ = fi56
                                                                        _ = v57
                                                                    }
                                                                    var v58 float64 = (1e10 + 1e10)
                                                                    _ = v58
                                                                } else {
                                                                    v59 := GenS7{F0: genFn0(0, v1, 2147483648), F1: 1, F2: v1}
                                                                    _ = v59
                                                                }
                                                                if v6 {
                                                                    if v6 {
                                                                        var v60 uint64 = 18446744073709551615
                                                                        v61 := GenS2{F0: v3}
                                                                        for fi62 := 0; fi62 < 3; fi62++ {
                                                                            if true {
                                                                                fmt.Println(v5, v5)
                                                                                v4 = v4
                                                                            } else {
                                                                                var v63 uint64 = 269
                                                                                if false {
                                                                                    switch int(v63) {
                                                                                    case 42:
                                                                                        var v64 bool = (false && (!v6))
                                                                                        _ = v64
                                                                                    case 914:
                                                                                        fmt.Println(v3, v5)
                                                                                    case 127:
                                                                                        v65 := GenS8{F0: v3, F1: 42, F2: 127}
                                                                                        _ = v65
                                                                                    }
                                                                                    fmt.Println(fi52)
                                                                                }
                                                                                _ = v63
                                                                            }
                                                                            if (869 != 971) {
                                                                                if len(v25) > 0 {
                                                                                    v25[0] = 1
                                                                                }
                                                                                fmt.Println(v1, v6)
                                                                            } else {
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v66 := v2["\n\t"]
                                                                                    _ = v66
                                                                                case "a":
                                                                                    if v6 {
                                                                                        var v67 bool = v6
                                                                                        v26 += (v26 >> ((18446744073709551615) & 63))
                                                                                        _ = v67
                                                                                    } else {
                                                                                        v68 := GenS6{F0: v4}
                                                                                        var v69 string = "a"
                                                                                        var v70 int64 = v1
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                        _ = v70
                                                                                    }
                                                                                default:
                                                                                    if ((fi62 + fi62) != int(v60)) {
                                                                                        fmt.Println(fi62, v1, v1)
                                                                                        v25 = append(v25, -778)
                                                                                        v71 := GenS3{F0: genFn0(2147483647, v1, 585), F1: 64, F2: 889}
                                                                                        _ = v71
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(v6)
                                                                            v6 = (!((1 <= 555) && (v60 > (v60 * (^(v60))))))
                                                                            _ = fi62
                                                                        }
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                }
                                                                _ = v53
                                                            } else {
                                                                v72 := v2["a"]
                                                                var v73 int = 9223372036854775807
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        case 42:
                                                            var v74 int32 = -1
                                                            _ = v74
                                                        default:
                                                            switch fi52 {
                                                            case -1:
                                                                var v75 int = (fi52 + 1)
                                                                _ = v75
                                                            case 397:
                                                                for fi76 := 0; fi76 < 6; fi76++ {
                                                                    v6 = false
                                                                    v6 = (0 == 1)
                                                                    fmt.Println(fi76, v26)
                                                                    _ = fi76
                                                                }
                                                            }
                                                        }
                                                        v77 := GenS8{F0: 1, F1: 353, F2: 2}
                                                        switch fi52 {
                                                        case 128:
                                                            if v6 {
                                                                fmt.Println(v3, v26)
                                                                fmt.Println(fi52)
                                                                var v78 uint8 = 0
                                                                _ = v78
                                                            }
                                                        default:
                                                            fmt.Println(v1, v6)
                                                        }
                                                        _ = fi52
                                                        _ = v77
                                                    }
                                                case 9223372036854775807:
                                                    v5 += float32(((float64(1e10) / (((float64(3.14159) + 1e10) / float64(((v5 * ((float32(-2.25) - float32(3.14159)) - (0.1 / 1e10))) / 3.14159))) / 0.1)) * (0.0 * (1e10 + (1e10 + 0.0)))))
                                                case 127:
                                                    if v6 {
                                                        if len(v25) > 0 {
                                                            v25[0] = -1
                                                        }
                                                        v3 = (v3 & 1)
                                                        for fi79 := 0; fi79 < 19; fi79++ {
                                                            v80 := GenS6{F0: (v4 >> ((genFn1()) & 63))}
                                                            _ = fi79
                                                            _ = v80
                                                        }
                                                    } else {
                                                        v81 := GenS7{F0: 2, F1: 225, F2: v1}
                                                        _ = v81
                                                    }
                                                }
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                v4 -= v4
                                            case "hello":
                                                v1 = (v1 ^ (v1 ^ 1))
                                            case "\n\t":
                                                switch ("" + "\n\t") {
                                                case "a":
                                                    v82 := v2[("a" + "\n\t")]
                                                    _ = v82
                                                case "hello":
                                                    var v83 uint = 955
                                                    _ = v83
                                                case "\n\t":
                                                    if len(v25) > 0 {
                                                        v25[0] = 9223372036854775807
                                                    }
                                                }
                                            default:
                                                switch 9223372036854775807 {
                                                case 9223372036854775807:
                                                    switch 32768 {
                                                    case -9223372036854775808:
                                                        for fi84 := 0; fi84 < 4; fi84++ {
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                            _ = fi84
                                                        }
                                                    default:
                                                        fmt.Println(v5)
                                                    }
                                                case 127:
                                                    fmt.Println(v3, v3, v1)
                                                case 2147483647:
                                                    _ = len(v2)
                                                default:
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                }
                                            }
                                            for fi85 := 0; fi85 < 7; fi85++ {
                                                v86 := GenS8{F0: 0, F1: 2147483647, F2: 127}
                                                _ = fi85
                                                _ = v86
                                            }
                                        } else {
                                            for fi87 := 0; fi87 < 1; fi87++ {
                                                if (false && v6) {
                                                    v4 += v4
                                                }
                                                var v88 bool = v6
                                                var v89 int16 = -32768
                                                _ = fi87
                                                _ = v88
                                                _ = v89
                                            }
                                        }
                                        var v90 bool = (4294967295 <= 0)
                                        _ = v90
                                    } else {
                                        v91 := GenS2{F0: (v3 / ((v3) | 1))}
                                        v5 = v5
                                        _ = v91
                                    }
                                    _ = len(v2)
                                    var v92 uint32 = 6
                                    _ = v92
                                }
                            }
                            if len(v25) > 0 {
                                v25[0] = -9223372036854775808
                            }
                        }
                    }
                }
                _ = v29
            } else {
                v93 := GenS8{F0: -2147483648, F1: 0, F2: 2}
                v94 := GenS9{F0: (("foo bar" + "hello") + "hello"), F1: 249}
                _ = v93
                _ = v94
            }
            _ = len(v2)
        } else {
            v6 = v6
            v3 = 1
        }
    }
    delete(v2, (("a" + "\n\t") + ""))
    if v6 {
        if len(v25) > 0 {
            v25[0] = 2147483647
        }
        fmt.Println(v6, v1, v26)
        fmt.Println(v4, v3, v3)
    } else {
        var v95 uint16 = (v26 >> ((4294967295) & 63))
        delete(v2, ("" + "\n\t"))
        _ = v95
    }
    var v96 int32 = v3
    v96 -= (v3 * (v96 ^ (v3 & -525)))
    for fi97 := 0; fi97 < 11; fi97++ {
        if v6 {
            var v98 uint64 = 281
            _ = v98
        } else {
            if (!((-(v4)) <= (v4 % ((v4) | 1)))) {
                switch "\n\t" {
                case "a":
                    v5 = v5
                case "\n\t":
                    v6 = v6
                case "foo bar":
                    v99 := GenS3{F0: 42, F1: 128, F2: 196}
                    _ = v99
                default:
                    for fi100 := 0; fi100 < 11; fi100++ {
                        v2["hello"] = (!v6)
                        _ = fi100
                    }
                }
                _ = len(v25)
            } else {
                v101 := GenS10{F0: 114}
                v26 -= v26
                _ = v101
            }
            fmt.Println(v5, v5)
            v102 := v2["\n\t"]
            _ = v102
        }
        switch "\n\t" {
        case "foo bar":
            v96 += -1
        case "\n\t":
            switch (fi97 / ((42) | 1)) {
            case 32767:
                for fi103 := 0; fi103 < 14; fi103++ {
                    var v104 string = ("foo bar" + ((("hello" + ("a" + "hello")) + (("" + "a") + "hello")) + "foo bar"))
                    var v105 int = 631
                    _ = fi103
                    _ = v104
                    _ = v105
                }
            }
        }
        _ = fi97
    }
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    v25 = append(v25, 2147483647)
    v106 := GenS3{F0: 642, F1: 127, F2: -32768}
    v4 = v4
    var v107 uint16 = (v26 | 32767)
    fmt.Println(v1)
    if false {
        var v108 uint16 = v107
        v26 = (v107 + (^(v107)))
        var v109 string = ""
        _ = v108
        _ = v109
    }
    delete(v2, ("hello" + "\n\t"))
    _ = len(v2)
    switch int(v1) {
    case 1:
        switch 0 {
        case 32767:
            switch 9223372036854775807 {
            case 429:
                v110 := GenS11{F0: (!true), F1: 2}
                _ = v110
            default:
                fmt.Println(v96, v5, v96)
            }
        default:
            var v111 int64 = (v1 & (v1 >> ((uint(v107)) & 63)))
            _ = v111
        }
    case -113:
        var v112 uint64 = 123
        _ = v112
    case 9223372036854775807:
        v1 -= int64(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v96
}
