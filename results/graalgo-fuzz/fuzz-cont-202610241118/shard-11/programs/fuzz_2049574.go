package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int32; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint }
type GenS6 struct { F0 uint32; F1 string; F2 int32 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint) int32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        p0 = p0
        p0 = 201
        p0 -= p0
        _ = fi4
    }
    var v5 int64 = -518
    if false {
        switch -9223372036854775808 {
        case 1:
            v5 += v5
        default:
            var v6 int64 = (v5 & -9223372036854775808)
            _ = v6
        }
        v7 := []int{1}
        _ = v7
    } else {
        fmt.Println(v5)
        v8 := GenS1{F0: 2147483648, F1: 32767, F2: 18446744073709551615}
        _ = v8
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := []int{-285, -342, -1}
    v2 := make(map[int]string)
    var v3 float64 = 3.14159
    v9 := GenS2{F0: 398}
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 float64 = 1e10
        v11 = ((float64(-2.25) * (v11 + (((float64(((0.0 / ((1e10 + (1.5 + 1e10)) + (1.5 - 1e10))) - 1e10)) - v11) * (v11 / 0.1)) * v11))) / 1.5)
        _ = fi10
        _ = v11
    }
    var v12 uint32 = 2
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS3{F0: 42, F1: -871, F2: 32767}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    var v16 int32 = genFn0(2)
    if true {
        v16 -= (v16 << ((4294967295) & 63))
        fmt.Println(v12, v16, v16)
        switch 42 {
        case 2147483647:
            if false {
                v17 := GenS3{F0: 1, F1: genFn0(712), F2: 255}
                v16 += (v16 ^ 1)
                v12 = (^(v12))
                _ = v17
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := GenS2{F0: (v12 * 885)}
                    var v20 uint8 = 1
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v21 := GenS2{F0: (v12 >> ((0) & 63))}
                _ = v21
            }
        case 127:
            v16 -= v16
        case -639:
            for fi22 := 0; fi22 < 1; fi22++ {
                if false {
                    var v23 uint16 = 941
                    var v24 uint64 = 42
                    var v25 uint = 803
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    var v26 int64 = 799
                    _ = v26
                }
                for fi27 := 0; fi27 < 6; fi27++ {
                    if len(v1) > 0 {
                        v1[0] = (fi27 + fi27)
                    }
                    v28 := GenS3{F0: v16, F1: (v16 - int32(fi27)), F2: 65535}
                    var v29 int64 = 42
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi22
            }
        default:
            if ((false && true) || false) {
                v12 += v12
                var v30 string = "\n\t"
                for fi31 := 0; fi31 < 2; fi31++ {
                    v30 = v30
                    if len(v1) > 0 {
                        v1[0] = fi31
                    }
                    switch -833 {
                    case 4294967295:
                        var v32 int64 = 980
                        _ = v32
                    case 42:
                        switch "" {
                        case "hello":
                            switch fi31 {
                            case 42:
                                for fi33 := 0; fi33 < 9; fi33++ {
                                    v34 := GenS3{F0: genFn0(787), F1: genFn0(1), F2: 1}
                                    _ = fi33
                                    _ = v34
                                }
                            case 0:
                                for fi35 := 0; fi35 < 8; fi35++ {
                                    if ((1.5 <= float32((v3 * (v3 / float64(-2.25))))) && false) {
                                        v36 := GenS4{F0: v16, F1: 0}
                                        v30 = "foo bar"
                                        _ = v36
                                    } else {
                                        var v37 int = (fi35 * 493)
                                        if (false && ((3.14159 - (v3 + v3)) == float64(float32(1.5)))) {
                                            var v38 uint = 42
                                            var v39 uint16 = 215
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = v37
                                    }
                                    switch "hello" {
                                    case "":
                                        var v40 uint32 = (v12 & 2)
                                        _ = v40
                                    case "a":
                                        var v41 int = (fi31 >> ((18446744073709551615) & 63))
                                        _ = v41
                                    }
                                    v12 += v12
                                    _ = fi35
                                }
                            default:
                                for fi42 := 0; fi42 < 8; fi42++ {
                                    var v43 int32 = v16
                                    if true {
                                        switch (("a" + "foo bar") + "\n\t") {
                                        case "a":
                                            v44 := GenS2{F0: (v12 << ((4294967295) & 63))}
                                            _ = v44
                                        case "":
                                            var v45 uint = 383
                                            _ = v45
                                        default:
                                            v43 -= (v16 ^ (-(v43)))
                                        }
                                        if true {
                                            v46 := GenS5{F0: 42, F1: v16, F2: 65535}
                                            _ = v46
                                        }
                                        fmt.Println(v12, fi31, v3)
                                    }
                                    _ = fi42
                                    _ = v43
                                }
                            }
                        }
                    default:
                        v16 = v16
                    }
                    _ = fi31
                }
                _ = v30
            }
        }
    } else {
        fmt.Println(v16)
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        switch (fi47 ^ fi47) {
        case 9223372036854775807:
            for fi48 := 0; fi48 < 8; fi48++ {
                delete(v2, fi48)
                v12 = 42
                _ = fi48
            }
        case -247:
            var v49 uint = 721
            _ = v49
        default:
            v12 = (v12 << ((1) & 63))
        }
        _ = fi47
    }
    var v50 int64 = int64(v16)
    v1 = append(v1, 1)
    var v51 int64 = (^(v50))
    if false {
        fmt.Println(v50)
        var v52 uint16 = 32767
        _ = v52
    }
    var v53 uint32 = (v12 * v12)
    if ("a" >= "\n\t") {
        if ((v50 << ((578) & 63)) <= v50) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v54 int = -1
        var v55 bool = (-2147483648 != 0)
        _ = v54
        _ = v55
    } else {
        delete(v2, 32767)
        v56 := GenS6{F0: v53, F1: "a", F2: (v16 * v16)}
        _ = v56
    }
    if (42 == 18446744073709551615) {
        for fi57 := 0; fi57 < 20; fi57++ {
            for fi58 := 0; fi58 < 8; fi58++ {
                v59 := GenS3{F0: -2147483648, F1: (v16 ^ genFn0(0)), F2: 32767}
                _ = fi58
                _ = v59
            }
            _ = fi57
        }
    } else {
        v50 = (v51 | -1)
    }
    fmt.Println(v50, v50)
    switch 2147483647 {
    case 802:
        if false {
            if true {
                fmt.Println(v16, v12)
                v50 = v50
            } else {
                switch 1 {
                case 4294967295:
                    if false {
                        if false {
                            v12 = (v53 - (v53 % ((42) | 1)))
                            var v60 int64 = -1
                            v60 += (^(v60))
                            _ = v60
                        } else {
                            v51 = v51
                            if (22 <= -128) {
                                v53 = v12
                                if false {
                                    fmt.Println(v51)
                                    v16 -= (v16 << ((18446744073709551615) & 63))
                                } else {
                                    for fi61 := 0; fi61 < 3; fi61++ {
                                        v50 = v50
                                        for fi62 := 0; fi62 < 12; fi62++ {
                                            var v63 uint64 = 0
                                            _ = fi62
                                            _ = v63
                                        }
                                        var v64 int8 = -61
                                        _ = fi61
                                        _ = v64
                                    }
                                    var v65 uint16 = 753
                                    switch 0 {
                                    case -9223372036854775808:
                                        var v66 float64 = -2.25
                                        _ = v66
                                    case 127:
                                        if (2147483648 < (v12 - 65535)) {
                                            if ((0 < int16(v65)) || (true && false)) {
                                                v16 += (v16 & (v16 | v16))
                                            } else {
                                                v51 = v51
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        } else {
                                            v67 := GenS1{F0: v12, F1: -935, F2: 58}
                                            _ = v67
                                        }
                                    default:
                                        var v68 int16 = -32768
                                        _ = v68
                                    }
                                    _ = v65
                                }
                            }
                        }
                        if (!(0.1 <= (0.0 / 1.5))) {
                            var v69 int8 = 1
                            _ = v69
                        }
                    }
                }
                if (((v3 - v3) - v3) < v3) {
                    var v70 int16 = 1000
                    switch "a" {
                    case "foo bar":
                        var v71 int8 = 42
                        _ = v71
                    case "\n\t":
                        switch -1 {
                        case 9223372036854775807:
                            if (!(23 < v50)) {
                                for fi72 := 0; fi72 < 8; fi72++ {
                                    v73 := GenS3{F0: -877, F1: (v16 >> ((255) & 63)), F2: 112}
                                    _ = fi72
                                    _ = v73
                                }
                            } else {
                                var v74 uint16 = 2
                                var v75 int16 = v70
                                _ = v74
                                _ = v75
                            }
                        case 158:
                            var v76 uint64 = 42
                            _ = v76
                        default:
                            if false {
                                for fi77 := 0; fi77 < 8; fi77++ {
                                    var v78 float64 = (((0.0 / ((0.1 + (float64(-2.25) * v3)) / v3)) - float64((1.5 * -2.25))) + (v3 * v3))
                                    var v79 bool = (false && ((-128 > 93) && (false && (!true))))
                                    switch "a" {
                                    case "foo bar":
                                        v2[42] = "a"
                                    case "":
                                        for fi80 := 0; fi80 < 17; fi80++ {
                                            v81 := GenS6{F0: (^(v53)), F1: "hello", F2: (v16 - (v16 % (((v16 >> ((uint(v12)) & 63))) | 1)))}
                                            var v82 bool = v79
                                            _ = fi80
                                            _ = v81
                                            _ = v82
                                        }
                                    case "hello":
                                        var v83 int = (fi77 + (fi77 % (((fi77 ^ -2147483648)) | 1)))
                                        _ = v83
                                    default:
                                        var v84 bool = false
                                        _ = v84
                                    }
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                }
                            }
                        }
                    default:
                        v85 := v2[int(v16)]
                        _ = v85
                    }
                    _ = v70
                } else {
                    for fi86 := 0; fi86 < 10; fi86++ {
                        v87 := GenS6{F0: (v12 % ((v12) | 1)), F1: "hello", F2: (v16 ^ 42)}
                        fmt.Println(v12)
                        v88 := GenS4{F0: (v16 & (v16 >> ((844) & 63))), F1: 947}
                        _ = fi86
                        _ = v87
                        _ = v88
                    }
                    v16 = v16
                    switch "" {
                    case "\n\t":
                        var v89 int = 1
                        _ = v89
                    default:
                        var v90 uint8 = 255
                        _ = v90
                    }
                }
                if (false || true) {
                    var v91 int8 = 42
                    _ = v91
                } else {
                    var v92 uint16 = 42
                    for fi93 := 0; fi93 < 19; fi93++ {
                        v94 := GenS7{F0: 42, F1: 128, F2: 1}
                        if len(v1) > 0 {
                            v1[0] = 32768
                        }
                        _ = fi93
                        _ = v94
                    }
                    _ = v92
                }
            }
            v95 := GenS8{F0: 1, F1: v51}
            _ = v95
        } else {
            v1 = append(v1, 32768)
        }
    case 1:
        fmt.Println(v3, v50, v50)
    }
    v51 -= v50
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v53
    _ = v9
}
