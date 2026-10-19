package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint; F1 int32; F2 int64 }
type GenS8 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS9 struct { F0 int16; F1 uint64; F2 uint }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 int8; F1 uint64; F2 int64 }

func main() {
    v1 := []int{127, 2147483647, -2147483648, 9223372036854775807}
    v2 := make(map[uint32]bool)
    var v3 uint64 = 2
    var v4 uint8 = 42
    var v5 string = ""
    var v6 uint16 = 2
    if ((^(v6)) >= (v6 * v6)) {
        fmt.Println(v5, v4, v5)
        v3 = v3
    } else {
        fmt.Println(v4, v5, v3)
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        delete(v2, 1)
        _ = len(v1)
        v8 := v2[2147483648]
        _ = fi7
        _ = v8
    }
    _ = len(v1)
    for fi9 := 0; fi9 < 8; fi9++ {
        v4 += (v4 + 128)
        v6 = v6
        for fi10 := 0; fi10 < 1; fi10++ {
            switch fi10 {
            case -2147483648:
                var v11 bool = ((false || true) && (!true))
                _ = v11
            case -9223372036854775808:
                var v12 uint64 = v3
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    v4 += (v4 / ((42) | 1))
    var v13 uint64 = v3
    if true {
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
        v6 -= (v6 & 2)
        v14 := GenS1{F0: v6, F1: -1}
        _ = v14
    } else {
        v15 := GenS1{F0: (^(v6)), F1: 1000}
        v4 = (v4 >> ((42) & 63))
        _ = v15
    }
    v3 = 42
    switch 4294967295 {
    case 32768:
        v3 += (v13 - (v13 & uint64(v4)))
    case -9223372036854775808:
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int8 = -1
            _ = fi16
            _ = v17
        }
    default:
        if len(v1) > 0 {
            v1[0] = -151
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        if false {
            v19 := v2[907]
            if (v5 < v5) {
                v13 = 0
                for fi20 := 0; fi20 < 20; fi20++ {
                    switch v5 {
                    case "":
                        switch 168 {
                        case 0:
                            var v21 uint = 608
                            _ = v21
                        default:
                            v3 = uint64(v6)
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v22 float32 = 0.0
                        _ = v22
                    }
                    v5 = (v5 + v5)
                    _ = fi20
                }
            }
            fmt.Println(v13, v13, v19)
            _ = v19
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(fi18, v13, fi18)
            if true {
                if true {
                    switch (fi18 | (fi23 + fi18)) {
                    case -1:
                        v3 += v13
                    case 161:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 4294967295:
                        if true {
                            v3 -= v13
                            var v24 uint16 = 75
                            var v25 int16 = 1000
                            _ = v24
                            _ = v25
                        } else {
                            var v26 int8 = 1
                            _ = v26
                        }
                    default:
                        _ = len(v1)
                    }
                }
                v5 = (("hello" + "foo bar") + "\n\t")
            } else {
                v5 = ("a" + "a")
            }
            _ = fi23
        }
        _ = fi18
    }
    v27 := GenS1{F0: (^(v6)), F1: -1}
    v28 := GenS2{F0: (("hello" + "foo bar") + "\n\t"), F1: -63}
    fmt.Println(v5, v3, v5)
    switch 457 {
    case -712:
        if false {
            v6 = v6
            var v29 int32 = 2147483647
            switch 686 {
            case -2147483648:
                if (100 >= -128) {
                    v6 = uint16(v4)
                    v6 += (v6 >> ((65535) & 63))
                }
            case 0:
                for fi30 := 0; fi30 < 1; fi30++ {
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        if false {
                            var v31 float32 = (((0.1 - (((-2.25 * 3.14159) - 0.0) - 0.1)) / -2.25) / 0.1)
                            _ = v31
                        }
                    case 2147483647:
                        if false {
                            var v32 bool = false
                            _ = v32
                        }
                    case -725:
                        if (35 < 65535) {
                            switch 1 {
                            case -797:
                                _ = len(v1)
                            case 9223372036854775807:
                                fmt.Println(fi30, v5, v6)
                            case -2147483648:
                                v33 := GenS2{F0: "hello", F1: fi30}
                                _ = v33
                            }
                        }
                    }
                    var v34 int64 = -211
                    _ = fi30
                    _ = v34
                }
            case -9223372036854775808:
                for fi35 := 0; fi35 < 14; fi35++ {
                    v36 := GenS3{F0: 719}
                    if len(v1) > 0 {
                        v1[0] = (fi35 ^ fi35)
                    }
                    _ = fi35
                    _ = v36
                }
            default:
                v37 := GenS4{F0: 9223372036854775807}
                _ = v37
            }
            _ = v29
        } else {
            for fi38 := 0; fi38 < 18; fi38++ {
                v13 += v13
                v6 = uint16(v3)
                v5 = ("" + "hello")
                _ = fi38
            }
        }
    case 0:
        switch (v5 + v5) {
        case "foo bar":
            var v39 int16 = 0
            _ = v39
        case "":
            if true {
                var v40 uint64 = (v13 + (^(v13)))
                _ = v40
            } else {
                if false {
                    _ = len(v2)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    if ((v6 ^ v6) > 1) {
                        if true {
                            if true {
                                var v41 uint8 = (v4 * (^(v4)))
                                _ = v41
                            } else {
                                fmt.Println(v6, v13, v3)
                            }
                            v42 := GenS2{F0: "a", F1: 449}
                            _ = v42
                        } else {
                            if ((0.1 + 1e10) != 1.5) {
                                v43 := GenS1{F0: 65535, F1: int16(v3)}
                                var v44 uint8 = (^(v4))
                                _ = v43
                                _ = v44
                            } else {
                                v3 -= v13
                                switch 504 {
                                case 128:
                                    v4 += (v4 + (v4 * 127))
                                }
                            }
                        }
                    } else {
                        v5 = (v5 + v5)
                        if true {
                            switch 916 {
                            case -836:
                                if (!(true || false)) {
                                    switch -743 {
                                    case 744:
                                        if (v5 > v5) {
                                            switch -9223372036854775808 {
                                            case 0:
                                                if (255 >= 18446744073709551615) {
                                                    v45 := GenS5{F0: true, F1: (v4 ^ 123), F2: 128}
                                                    _ = v45
                                                } else {
                                                    v6 = 2
                                                    var v46 uint8 = v4
                                                    fmt.Println(v13, v3)
                                                    _ = v46
                                                }
                                            case 42:
                                                v6 -= 2
                                            default:
                                                switch ("foo bar" + (v5 + ("hello" + "\n\t"))) {
                                                case "\n\t":
                                                    for fi47 := 0; fi47 < 14; fi47++ {
                                                        v6 = 42
                                                        _ = fi47
                                                    }
                                                default:
                                                    v13 += v13
                                                }
                                            }
                                            for fi48 := 0; fi48 < 13; fi48++ {
                                                if (!((^(fi48)) <= fi48)) {
                                                    v49 := GenS6{F0: -128, F1: fi48}
                                                    _ = v49
                                                }
                                                for fi50 := 0; fi50 < 5; fi50++ {
                                                    v51 := GenS7{F0: 255, F1: -2147483648, F2: 1}
                                                    v52 := GenS8{F0: (v4 * uint8(v3)), F1: 1, F2: v4}
                                                    _ = fi50
                                                    _ = v51
                                                    _ = v52
                                                }
                                                switch ("foo bar" + (v5 + (v5 + "foo bar"))) {
                                                case "a":
                                                    v53 := GenS4{F0: 536}
                                                    _ = v53
                                                case "hello":
                                                    v54 := GenS4{F0: 372}
                                                    _ = v54
                                                }
                                                _ = fi48
                                            }
                                            v3 -= v3
                                        }
                                    case 306:
                                        var v55 uint16 = (v6 ^ v6)
                                        _ = v55
                                    default:
                                        _ = len(v2)
                                    }
                                    var v56 uint16 = v6
                                    v57 := GenS3{F0: 18446744073709551615}
                                    _ = v56
                                    _ = v57
                                } else {
                                    v58 := GenS9{F0: 32767, F1: 42, F2: uint(v4)}
                                    for fi59 := 0; fi59 < 3; fi59++ {
                                        var v60 int32 = -1
                                        fmt.Println(v6)
                                        _ = fi59
                                        _ = v60
                                    }
                                    var v61 int = 1
                                    _ = v58
                                    _ = v61
                                }
                            case 42:
                                v62 := GenS10{F0: v4, F1: uint16(v3)}
                                _ = v62
                            case 1:
                                for fi63 := 0; fi63 < 8; fi63++ {
                                    for fi64 := 0; fi64 < 1; fi64++ {
                                        v65 := GenS4{F0: 42}
                                        var v66 uint16 = (v6 & v6)
                                        fmt.Println(v13, fi63, v13)
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                    v67 := GenS2{F0: v5, F1: (^(fi63))}
                                    _ = fi63
                                    _ = v67
                                }
                            }
                            var v68 bool = ((((!false) && (9223372036854775807 <= 382)) || true) || (949 < uint32(v4)))
                            var v69 int16 = 112
                            _ = v68
                            _ = v69
                        } else {
                            var v70 int8 = -15
                            fmt.Println(v5, v5, v4)
                            v4 = v4
                            _ = v70
                        }
                        if true {
                            if (1000 <= 42) {
                                for fi71 := 0; fi71 < 10; fi71++ {
                                    if (255 >= 255) {
                                        fmt.Println(fi71)
                                    } else {
                                        if true {
                                            v72 := GenS9{F0: 32767, F1: v3, F2: 18446744073709551615}
                                            _ = v72
                                        } else {
                                            v73 := GenS10{F0: v4, F1: 864}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v13 += v13
                                            _ = v73
                                        }
                                        v3 += 2
                                    }
                                    _ = fi71
                                }
                                if false {
                                    v74 := GenS7{F0: 42, F1: 706, F2: 42}
                                    switch 32768 {
                                    case 32767:
                                        v4 += 42
                                    default:
                                        var v75 float32 = float32((0.1 * ((float64(1e10) * (1e10 / 1e10)) + ((1.5 / float64(((1.5 * float32((1e10 - (0.0 + (((float64((-2.25 * 1.5)) - (3.14159 * 3.14159)) * float64(1.5)) * (0.1 * 0.1)))))) + 0.1))) * 1e10))))
                                        _ = v75
                                    }
                                    _ = v74
                                } else {
                                    v2[356] = (32768 != -200)
                                    v76 := GenS11{F0: -128, F1: v13, F2: 932}
                                    var v77 int8 = 127
                                    _ = v76
                                    _ = v77
                                }
                                v78 := GenS1{F0: (v6 / ((v6) | 1)), F1: 32767}
                                _ = v78
                            }
                            if len(v1) > 0 {
                                v1[0] = 894
                            }
                        }
                    }
                    delete(v2, 445)
                    v5 = "hello"
                }
                var v79 float32 = -2.25
                _ = v79
            }
        }
    case -140:
        var v80 int32 = 2147483647
        _ = v80
    default:
        var v81 string = (((((v5 + v5) + v5) + v5) + (v5 + ((v5 + v5) + "\n\t"))) + (("hello" + (v5 + ("foo bar" + v5))) + v5))
        _ = v81
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
