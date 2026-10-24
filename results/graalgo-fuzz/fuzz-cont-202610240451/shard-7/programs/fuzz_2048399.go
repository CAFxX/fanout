package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS7 struct { F0 uint; F1 uint64; F2 int16 }
type GenS8 struct { F0 int; F1 int32; F2 bool }
type GenS9 struct { F0 int; F1 string; F2 uint32 }

func main() {
    var v1 uint8 = 0
    v2 := []uint32{}
    var v3 int16 = 123
    if (285 != 42) {
        _ = len(v2)
        _ = len(v2)
    }
    switch 491 {
    case -590:
        fmt.Println(v1)
    }
    if (!(!true)) {
        v4 := GenS1{F0: 35, F1: 0, F2: v3}
        switch 128 {
        case 128:
            v5 := GenS1{F0: -1, F1: 65535, F2: (v3 & v3)}
            _ = v5
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v4
    } else {
        var v6 bool = (2 >= v1)
        var v7 int = -722
        v8 := GenS1{F0: -51, F1: 2, F2: v3}
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int8 = -128
    var v10 uint = 42
    if true {
        v10 = (v10 * (v10 ^ 2))
        if false {
            var v11 int64 = 42
            var v12 int = -1
            if true {
                switch v12 {
                case 660:
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                case 32768:
                    var v13 uint = (v10 % ((v10) | 1))
                    _ = v13
                case 0:
                    var v14 uint64 = 287
                    _ = v14
                default:
                    v1 = (v1 ^ 127)
                }
                if (!true) {
                    if len(v2) > 0 {
                        v2[0] = 208
                    }
                    if true {
                        v1 -= uint8(v10)
                        fmt.Println(v3)
                    }
                } else {
                    v12 += v12
                    var v15 string = ("\n\t" + "")
                    _ = len(v2)
                    _ = v15
                }
                if ((v3 >= v3) && false) {
                    v9 -= (v9 + v9)
                }
            }
            _ = v11
            _ = v12
        } else {
            v16 := GenS1{F0: v9, F1: v10, F2: v3}
            var v17 uint16 = uint16(v1)
            _ = v16
            _ = v17
        }
    } else {
        v10 = 65535
        if ((v9 & (-(v9))) != v9) {
            for fi18 := 0; fi18 < 10; fi18++ {
                v10 = 1
                var v19 uint32 = 1
                fmt.Println(v9, v1)
                _ = fi18
                _ = v19
            }
            if ((1e10 + 3.14159) != 0.0) {
                var v20 int8 = (v9 % (((v9 ^ int8(v3))) | 1))
                fmt.Println(v1, v10)
                if true {
                    var v21 float64 = 0.1
                    v22 := GenS2{F0: v9}
                    if (!false) {
                        v9 += (^(v9))
                        if false {
                            var v23 uint8 = v1
                            for fi24 := 0; fi24 < 4; fi24++ {
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    var v26 int16 = (-(v3))
                                    fmt.Println(fi25)
                                    _ = fi25
                                    _ = v26
                                }
                                var v27 uint32 = 0
                                _ = fi24
                                _ = v27
                            }
                            if true {
                                var v28 uint = (^(v10))
                                var v29 int32 = 1
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int32 = -717
                                _ = v30
                            }
                            _ = v23
                        }
                    } else {
                        if true {
                            if false {
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v32 := GenS3{F0: false}
                                    if false {
                                        for fi33 := 0; fi33 < 7; fi33++ {
                                            v34 := GenS4{F0: v1, F1: (^(v3))}
                                            _ = fi33
                                            _ = v34
                                        }
                                        var v35 float32 = float32(((v21 * (float64((((-2.25 / (float32(v21) / ((float32(v21) / ((-2.25 / (((0.1 / float32(v21)) * ((0.1 - (3.14159 + float32(1e10))) * float32((v21 * v21)))) * (((0.0 * (3.14159 / 1e10)) + 0.1) + 1e10))) / -2.25)) - (3.14159 * (3.14159 + (-2.25 / 3.14159)))))) / (3.14159 + float32(1.5))) - 0.1)) + -2.25)) - -2.25))
                                        v3 += (v3 + 1000)
                                        _ = v35
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            }
                            v1 += 151
                        }
                        var v36 string = "foo bar"
                        _ = v36
                    }
                    _ = v21
                    _ = v22
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v20
            }
            _ = len(v2)
        }
        if true {
            var v37 float64 = 0.0
            var v38 int = -516
            var v39 int8 = v9
            _ = v37
            _ = v38
            _ = v39
        }
    }
    if true {
        var v40 int8 = v9
        if false {
            if (!true) {
                v40 = 1
                fmt.Println(v1)
            } else {
                switch ("foo bar" + (("\n\t" + "a") + ("\n\t" + ""))) {
                case "":
                    v40 = (v40 ^ -128)
                case "hello":
                    var v41 int8 = 42
                    _ = v41
                }
                for fi42 := 0; fi42 < 17; fi42++ {
                    fmt.Println(v9, v40)
                    _ = fi42
                }
            }
            _ = len(v2)
        }
        _ = v40
    } else {
        if true {
            switch -1 {
            case 32767:
                fmt.Println(v9, v3, v1)
            case 2147483647:
                switch "hello" {
                case "":
                    switch 4294967295 {
                    case -9223372036854775808:
                        v3 = (v3 << ((42) & 63))
                    case 2147483647:
                        v2 = append(v2, 217)
                    default:
                        if true {
                            var v43 uint16 = 65535
                            v3 = (v3 * 1)
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                            _ = v43
                        }
                    }
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                default:
                    fmt.Println(v3, v10)
                }
            }
        } else {
            v2 = append(v2, 65535)
            v2 = append(v2, 450)
        }
        fmt.Println(v3, v3)
    }
    v44 := GenS5{F0: ("" + "a"), F1: 0}
    v3 = (v3 % ((int16(v1)) | 1))
    v1 += (v1 * (v1 ^ (v1 ^ v1)))
    fmt.Println(v10)
    fmt.Println(v10, v1)
    var v45 float32 = 3.14159
    var v46 int = 1
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    switch "" {
    case "hello":
        if ((!false) || false) {
            v9 += (v9 << ((v10) & 63))
            if (501 < 65535) {
                fmt.Println(v46)
                for fi47 := 0; fi47 < 11; fi47++ {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        v10 -= (^(v10))
                        _ = fi48
                    }
                    v2 = append(v2, 0)
                    _ = fi47
                }
                var v49 uint32 = uint32(v1)
                _ = v49
            } else {
                switch "\n\t" {
                case "hello":
                    v50 := GenS6{F0: 42, F1: -27, F2: uint32(v1)}
                    _ = v50
                case "\n\t":
                    for fi51 := 0; fi51 < 2; fi51++ {
                        _ = len(v2)
                        _ = fi51
                    }
                }
                if (765 <= 852) {
                    var v52 uint64 = 18446744073709551615
                    if true {
                        var v53 bool = (65535 > 42)
                        fmt.Println(v45)
                        _ = v53
                    }
                    _ = v52
                }
                var v54 uint32 = 196
                _ = v54
            }
            if true {
                if true {
                    v45 += v45
                } else {
                    v10 += v10
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v45 += float32(1.5)
            }
        }
    case "a":
        if ((v1 >= 255) && false) {
            v9 = v9
            for fi55 := 0; fi55 < 13; fi55++ {
                _ = len(v2)
                v3 += (^(v3))
                _ = fi55
            }
        } else {
            var v56 int = 127
            for fi57 := 0; fi57 < 13; fi57++ {
                for fi58 := 0; fi58 < 12; fi58++ {
                    v9 -= 100
                    _ = fi58
                }
                v10 = v10
                switch fi57 {
                case -605:
                    v3 = v3
                case 2147483647:
                    if false {
                        v9 -= (v9 % (((v9 ^ (v9 + -17))) | 1))
                        fmt.Println(v3, v56)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case -425:
                    v2 = append(v2, 2147483648)
                default:
                    var v59 uint64 = uint64(v10)
                    _ = v59
                }
                _ = fi57
            }
            fmt.Println(v56, v46, v9)
            _ = v56
        }
    }
    if (785 != 0) {
        if false {
            for fi60 := 0; fi60 < 15; fi60++ {
                for fi61 := 0; fi61 < 12; fi61++ {
                    if true {
                        var v62 uint64 = 2
                        if (true && (false || ((!false) || (0 != 65535)))) {
                            var v63 float32 = 0.1
                            _ = v63
                        } else {
                            _ = len(v2)
                            var v64 uint16 = 770
                            v2 = append(v2, 984)
                            _ = v64
                        }
                        _ = v62
                    }
                    if false {
                        if true {
                            switch ("foo bar" + "hello") {
                            case "foo bar":
                                fmt.Println(fi61)
                                v1 -= v1
                            case "hello":
                                v3 = (-(v3))
                                v65 := GenS7{F0: 1, F1: 2, F2: 0}
                                _ = v65
                            case "\n\t":
                                v1 -= v1
                            }
                            v66 := GenS8{F0: (v46 << ((253) & 63)), F1: 2147483647, F2: true}
                            fmt.Println(fi61, fi61)
                            _ = v66
                        }
                        var v67 uint32 = 4294967295
                        if (2 < 2147483648) {
                            v46 = v46
                            v9 += v9
                            v68 := GenS7{F0: (v10 | (v10 << (((^(v10))) & 63))), F1: 611, F2: v3}
                            _ = v68
                        } else {
                            var v69 int = (fi61 + (-(fi60)))
                            v46 = fi60
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v69
                        }
                        _ = v67
                    }
                    var v70 int8 = v9
                    switch fi60 {
                    case 2147483647:
                        if false {
                            var v71 uint16 = 42
                            fmt.Println(v45, v70, v3)
                            if (!false) {
                                v3 = 1
                                switch fi60 {
                                case 0:
                                    v72 := GenS8{F0: 2147483647, F1: -189, F2: (true && false)}
                                    _ = v72
                                case 1:
                                    _ = len(v2)
                                default:
                                    var v73 int8 = 42
                                    _ = v73
                                }
                                var v74 int = int(v9)
                                switch 1 {
                                case 42:
                                    v75 := GenS9{F0: (v46 * v46), F1: "\n\t", F2: 4294967295}
                                    _ = v75
                                case -1:
                                    switch 32767 {
                                    case 476:
                                        var v76 uint = (v10 / ((v10) | 1))
                                        _ = v76
                                    }
                                    var v77 uint32 = 2147483647
                                    _ = v77
                                case 32768:
                                    if (!(1 != 983)) {
                                        var v78 uint64 = 678
                                        var v79 float32 = (v45 + v45)
                                        fmt.Println(v3, v70, v78)
                                        _ = v78
                                        _ = v79
                                    }
                                }
                                _ = v74
                            } else {
                                if false {
                                    if (!(-891 != v3)) {
                                        v46 -= (-(fi60))
                                        if (!true) {
                                            v80 := GenS8{F0: (v46 & fi60), F1: -2147483648, F2: (((^(v10)) >= 795) || false)}
                                            v46 = fi61
                                            _ = v80
                                        } else {
                                            var v81 uint16 = v71
                                            _ = v81
                                        }
                                    } else {
                                        var v82 bool = (v1 < v1)
                                        switch "a" {
                                        case "foo bar":
                                            var v83 uint16 = (^(v71))
                                            _ = v83
                                        }
                                        v2 = append(v2, 2147483647)
                                        v1 = 1
                                        _ = v82
                                    }
                                    var v84 uint = v10
                                    _ = v84
                                } else {
                                    fmt.Println(v70, v71, fi61)
                                    if (v45 <= float32(((1e10 + 1e10) - 0.0))) {
                                        fmt.Println(v70)
                                    }
                                }
                                var v85 uint64 = 0
                                v85 -= (v85 ^ v85)
                                _ = v85
                            }
                            _ = v71
                        } else {
                            v2 = append(v2, 65535)
                            var v86 int64 = 1
                            _ = v86
                        }
                    case 32767:
                        if (!(!(!true))) {
                            _ = len(v2)
                            v87 := GenS3{F0: true}
                            _ = v87
                        }
                    case 0:
                        var v88 int = fi60
                        _ = v88
                    }
                    _ = fi61
                    _ = v70
                }
                switch fi60 {
                case 32768:
                    fmt.Println(v1, v10)
                }
                if ((!(("" + "") != "foo bar")) || true) {
                    v45 += v45
                } else {
                    v9 = (v9 + v9)
                }
                _ = fi60
            }
        } else {
            var v89 uint64 = 18446744073709551615
            _ = v89
        }
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v90 := GenS2{F0: (-(v9))}
        _ = v90
    } else {
        fmt.Println(v3, v10, v46)
    }
    for fi91 := 0; fi91 < 3; fi91++ {
        v45 = v45
        _ = fi91
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v9
}
