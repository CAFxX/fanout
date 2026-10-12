package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int64; F2 int64 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 string }
func genFn0(p0 int64, p1 uint8, p2 int8) int16 {
    v6 := GenS1{F0: 864, F1: 0}
    v7 := GenS2{F0: (p0 % ((-190) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int = 32767
    var v2 uint = 0
    var v3 bool = false
    var v4 uint8 = 1
    var v5 uint64 = 149
    v8 := []int{v1, 128}
    if ((^(v1)) >= (v1 % (((v1 ^ (^(v1)))) | 1))) {
        switch v1 {
        case 42:
            v5 = v5
        case 32767:
            if v3 {
                v1 -= v1
                var v9 uint64 = v5
                switch "\n\t" {
                case "\n\t":
                    v10 := GenS2{F0: 9223372036854775807}
                    _ = v10
                case "hello":
                    for fi11 := 0; fi11 < 15; fi11++ {
                        if v3 {
                            if true {
                                v9 += v5
                            } else {
                                var v12 int8 = -44
                                for fi13 := 0; fi13 < 2; fi13++ {
                                    v14 := GenS3{F0: int64(v2)}
                                    _ = fi13
                                    _ = v14
                                }
                                v5 = (v9 >> ((uint(v9)) & 63))
                                _ = v12
                            }
                        } else {
                            if v3 {
                                fmt.Println(v3, fi11, v1)
                                fmt.Println(v1, v1, v9)
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    if (!v3) {
                                        fmt.Println(v1, fi11)
                                    }
                                    _ = fi15
                                }
                            }
                            if len(v8) > 0 {
                                v8[0] = v1
                            }
                        }
                        switch "foo bar" {
                        case "\n\t":
                            _ = len(v8)
                        default:
                            switch "" {
                            case "\n\t":
                                if false {
                                    v16 := GenS2{F0: -360}
                                    _ = v16
                                }
                            case "a":
                                v17 := GenS3{F0: -469}
                                _ = v17
                            default:
                                var v18 uint32 = 1
                                _ = v18
                            }
                        }
                        _ = fi11
                    }
                case "foo bar":
                    if (4294967295 == 65535) {
                        switch ((((("foo bar" + ("" + "a")) + "a") + "") + "") + ((("foo bar" + "foo bar") + "foo bar") + "foo bar")) {
                        case "a":
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    if v3 {
                                        for fi20 := 0; fi20 < 10; fi20++ {
                                            var v21 string = "\n\t"
                                            _ = fi20
                                            _ = v21
                                        }
                                        v22 := GenS1{F0: 4294967295, F1: 544}
                                        _ = v22
                                    } else {
                                        fmt.Println(v5)
                                        for fi23 := 0; fi23 < 7; fi23++ {
                                            v3 = v3
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            var v24 uint64 = 18446744073709551615
                                            _ = fi23
                                            _ = v24
                                        }
                                        v25 := GenS4{F0: uint64(v4), F1: 712}
                                        _ = v25
                                    }
                                    for fi26 := 0; fi26 < 17; fi26++ {
                                        v1 += fi26
                                        _ = fi26
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = fi19
                                    }
                                    _ = fi19
                                }
                            case "\n\t":
                                v8 = append(v8, v1)
                            default:
                                var v27 int16 = -1
                                _ = v27
                            }
                        default:
                            if v3 {
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    var v29 int8 = -1
                                    for fi30 := 0; fi30 < 2; fi30++ {
                                        v31 := GenS5{F0: (v4 & v4)}
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                switch 4294967295 {
                                case 432:
                                    v32 := GenS6{F0: 2, F1: 0, F2: int64(v1)}
                                    _ = v32
                                case -2147483648:
                                    v33 := GenS3{F0: -28}
                                    _ = v33
                                default:
                                    if len(v8) > 0 {
                                        v8[0] = 127
                                    }
                                }
                            } else {
                                fmt.Println(v3)
                                v8 = append(v8, (v1 / (((v1 << ((v2) & 63))) | 1)))
                                v4 += (v4 + 2)
                            }
                        }
                    }
                default:
                    v34 := GenS6{F0: 2, F1: 0, F2: 42}
                    _ = v34
                }
                _ = v9
            } else {
                if v3 {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        switch (v1 / (((v1 | 4294967295)) | 1)) {
                        case -9223372036854775808:
                            v1 += (v1 << ((v2) & 63))
                        case 2147483647:
                            v4 += (v4 >> ((v2) & 63))
                        case 4294967295:
                            fmt.Println(v1, v2)
                        default:
                            v1 = 127
                        }
                        v5 = (v5 * v5)
                        _ = fi35
                    }
                    switch "a" {
                    case "\n\t":
                        var v36 uint = (v2 - (v2 % ((4294967295) | 1)))
                        _ = v36
                    case "foo bar":
                        switch (v1 % ((int(v2)) | 1)) {
                        case -2147483648:
                            v4 -= (v4 ^ v4)
                        case -654:
                            fmt.Println(v3)
                        case -1:
                            fmt.Println(v2)
                        default:
                            switch "a" {
                            case "hello":
                                for fi37 := 0; fi37 < 10; fi37++ {
                                    v38 := GenS2{F0: 9223372036854775807}
                                    switch "\n\t" {
                                    case "foo bar":
                                        for fi39 := 0; fi39 < 5; fi39++ {
                                            if true {
                                                v4 = (v4 << (((v2 & (v2 | 18446744073709551615))) & 63))
                                                var v40 int8 = 42
                                                _ = v40
                                            } else {
                                                v41 := GenS3{F0: 1}
                                                v42 := GenS7{F0: 100, F1: 127}
                                                _ = v41
                                                _ = v42
                                            }
                                            _ = fi39
                                        }
                                    case "a":
                                        var v43 uint16 = 255
                                        _ = v43
                                    case "\n\t":
                                        for fi44 := 0; fi44 < 1; fi44++ {
                                            var v45 int64 = 0
                                            if (128 == 1) {
                                                var v46 bool = v3
                                                fmt.Println(v2)
                                                _ = v46
                                            } else {
                                                var v47 int64 = (v45 >> (((v2 ^ (v2 / ((v2) | 1)))) & 63))
                                                _ = v47
                                            }
                                            v2 += uint(v5)
                                            _ = fi44
                                            _ = v45
                                        }
                                    }
                                    _ = fi37
                                    _ = v38
                                }
                            case "":
                                v4 -= (v4 / ((255) | 1))
                            default:
                                switch v1 {
                                case -9223372036854775808:
                                    for fi48 := 0; fi48 < 9; fi48++ {
                                        v4 = (^(v4))
                                        fmt.Println(v4, v2, v1)
                                        for fi49 := 0; fi49 < 4; fi49++ {
                                            v1 = fi48
                                            _ = fi49
                                        }
                                        _ = fi48
                                    }
                                default:
                                    var v50 string = "\n\t"
                                    _ = v50
                                }
                            }
                        }
                    case "hello":
                        _ = len(v8)
                    }
                } else {
                    var v51 bool = true
                    _ = v51
                }
            }
        case 289:
            for fi52 := 0; fi52 < 11; fi52++ {
                var v53 int64 = 9223372036854775807
                var v54 uint = v2
                _ = fi52
                _ = v53
                _ = v54
            }
        }
        v55 := GenS8{F0: 42, F1: 65535}
        _ = v55
    }
    v56 := GenS9{F0: 65535}
    for fi57 := 0; fi57 < 12; fi57++ {
        if v3 {
            v3 = (32767 >= 1)
            v2 -= v2
        }
        _ = fi57
    }
    if ((1 >= 1) || (((-47 > 67) && v3) || true)) {
        switch (v1 - v1) {
        case 32768:
            if (v3 || v3) {
                var v58 int8 = -74
                if ((0.1 - 1.5) != 1e10) {
                    if (713 > 32767) {
                        if ((!(((18446744073709551615 > v5) && true) && v3)) || v3) {
                            v8 = append(v8, 2147483647)
                            for fi59 := 0; fi59 < 2; fi59++ {
                                fmt.Println(v2)
                                _ = fi59
                            }
                        }
                    }
                    v1 = v1
                    if (false && true) {
                        switch "hello" {
                        case "":
                            var v60 uint64 = v5
                            _ = v60
                        default:
                            v61 := GenS10{F0: 42, F1: -1, F2: v5}
                            _ = v61
                        }
                        v1 = -21
                    } else {
                        for fi62 := 0; fi62 < 5; fi62++ {
                            v2 += (v2 << (((v2 << ((v2) & 63))) & 63))
                            if true {
                                var v63 string = "foo bar"
                                v4 = 145
                                for fi64 := 0; fi64 < 1; fi64++ {
                                    if v3 {
                                        if v3 {
                                            if v3 {
                                                fmt.Println(fi62)
                                                var v65 uint16 = uint16(v5)
                                                v4 = v4
                                                _ = v65
                                            }
                                            switch v63 {
                                            case "hello":
                                                v63 = ((v63 + "") + v63)
                                            case "foo bar":
                                                v66 := GenS11{F0: (^(v1)), F1: (v1 & (-(fi62)))}
                                                if (1e10 == float64((-2.25 / float32((1e10 + 1.5))))) {
                                                    var v67 uint16 = 65535
                                                    v1 -= (^(fi62))
                                                    _ = v67
                                                }
                                                v63 = "\n\t"
                                                _ = v66
                                            case "\n\t":
                                                var v68 float64 = 3.14159
                                                v2 = 2
                                                _ = v68
                                            default:
                                                v69 := GenS12{F0: 1}
                                                _ = v69
                                            }
                                            v4 += v4
                                            if v3 {
                                                switch 127 {
                                                case 9223372036854775807:
                                                    v1 += (v1 % ((451) | 1))
                                                case 128:
                                                    if (!(0.0 != 1.5)) {
                                                        fmt.Println(v63)
                                                        if ((v2 + (^(v2))) <= v2) {
                                                            fmt.Println(fi62, v1, v3)
                                                            v58 += v58
                                                        }
                                                    } else {
                                                        if v3 {
                                                            switch v1 {
                                                            case -9223372036854775808:
                                                                switch v63 {
                                                                case "":
                                                                    v1 += 128
                                                                default:
                                                                    v4 = 2
                                                                }
                                                            default:
                                                                if true {
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    _ = len(v8)
                                                                    var v70 bool = ((!false) && (!(fi62 <= fi64)))
                                                                    _ = v70
                                                                } else {
                                                                    if (fi64 > v1) {
                                                                        v71 := GenS13{F0: 66}
                                                                        _ = v71
                                                                    }
                                                                    var v72 int16 = -966
                                                                    _ = v72
                                                                }
                                                                switch v63 {
                                                                case "":
                                                                    var v73 int8 = v58
                                                                    _ = v73
                                                                default:
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                }
                                                                fmt.Println(fi62, v3)
                                                            }
                                                        }
                                                        v74 := GenS4{F0: v5, F1: 65535}
                                                        v2 = 255
                                                        _ = v74
                                                    }
                                                case 32768:
                                                    v75 := GenS7{F0: (v58 & v58), F1: 127}
                                                    _ = v75
                                                }
                                                var v76 float32 = 1e10
                                                _ = v76
                                            } else {
                                                v1 = fi62
                                            }
                                        }
                                    }
                                    var v77 float64 = 0.1
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = fi64
                                    _ = v77
                                }
                                _ = v63
                            }
                            var v78 int8 = (v58 % ((v58) | 1))
                            _ = fi62
                            _ = v78
                        }
                        v3 = v3
                    }
                } else {
                    var v79 int = (v1 % (((^(v1))) | 1))
                    if v3 {
                        fmt.Println(v1, v3)
                        var v80 int = 32767
                        v1 = (v79 ^ v80)
                        _ = v80
                    } else {
                        fmt.Println(v5)
                        _ = len(v8)
                    }
                    _ = v79
                }
                if ((v1 * (v1 * v1)) >= -2147483648) {
                    v81 := GenS10{F0: 65535, F1: 2147483647, F2: v5}
                    _ = v81
                }
                _ = v58
            } else {
                var v82 float64 = (1e10 * ((-2.25 / 0.1) - (-2.25 - 0.1)))
                var v83 float64 = float64((((1.5 / (-2.25 * 3.14159)) + 0.1) / (1.5 * (3.14159 / 0.1))))
                v82 += float64(1e10)
                _ = v82
                _ = v83
            }
        case 4294967295:
            v4 = (^(v4))
        case -9223372036854775808:
            var v84 int64 = 407
            _ = v84
        }
    } else {
        var v85 int8 = -1
        for fi86 := 0; fi86 < 11; fi86++ {
            var v87 bool = true
            _ = fi86
            _ = v87
        }
        v88 := GenS14{F0: -2147483648, F1: "foo bar"}
        _ = v85
        _ = v88
    }
    switch "" {
    case "foo bar":
        var v89 string = ""
        _ = v89
    case "hello":
        fmt.Println(v2)
    default:
        fmt.Println(v4, v5)
    }
    v3 = (true || (0.1 == (0.1 + 0.0)))
    _ = len(v8)
    _ = len(v8)
    v2 = (v2 ^ (v2 << (((v2 - v2)) & 63)))
    v3 = true
    var v90 int32 = 0
    v2 += (v2 << (((v2 | v2)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v8
    _ = v90
}
