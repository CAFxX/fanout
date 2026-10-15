package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int32; F1 int; F2 uint32 }
type GenS3 struct { F0 int32; F1 string; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint32; F1 uint8; F2 bool }
type GenS8 struct { F0 uint8; F1 int16 }
func genFn0(p0 int16, p1 uint16, p2 int32) uint64 {
    v6 := []int{389, -2147483648, 42, 0}
    if false {
        v7 := GenS1{F0: 65535, F1: p0}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int16 = 1
            switch fi8 {
            case 275:
                var v10 string = "a"
                _ = v10
            case 0:
                v11 := GenS1{F0: 2147483647, F1: 1000}
                _ = v11
            case 4294967295:
                if (4294967295 >= 504) {
                    p0 -= (p0 * -32768)
                    p1 -= (p1 << ((873) & 63))
                }
            }
            _ = fi8
            _ = v9
        }
        var v12 string = ("" + "a")
        switch "\n\t" {
        case "foo bar":
            var v13 uint64 = 755
            _ = v13
        case "\n\t":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "":
            var v14 uint16 = 269
            _ = v14
        default:
            p1 -= p1
        }
        _ = v12
    }
    fmt.Println(p1)
    var v15 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return 2
}
func genFn1() uint16 {
    var v16 uint64 = 2
    if true {
        if (true && true) {
            v16 = v16
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 string = ""
                _ = fi17
                _ = v18
            }
        } else {
            switch 473 {
            case 1:
                for fi19 := 0; fi19 < 18; fi19++ {
                    v16 = genFn0(-1, 255, 1)
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 uint16 = 32767
                        v22 := []int{fi20}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = fi19
                }
            case 128:
                var v23 uint16 = 32767
                _ = v23
            case 127:
                v24 := GenS1{F0: 65535, F1: 411}
                _ = v24
            }
            if (0 < 816) {
                v16 = v16
                var v25 uint8 = 206
                _ = v25
            }
            fmt.Println(v16)
        }
    } else {
        for fi26 := 0; fi26 < 13; fi26++ {
            if (!true) {
                if true {
                    switch (fi26 & fi26) {
                    case 32768:
                        v16 -= v16
                    case -2147483648:
                        if false {
                            for fi27 := 0; fi27 < 15; fi27++ {
                                v16 = (v16 / ((genFn0(int16(fi26), 0, 353)) | 1))
                                _ = fi27
                            }
                        }
                    default:
                        fmt.Println(fi26, v16)
                    }
                    v16 += (v16 ^ v16)
                }
                v28 := []int{474, 2147483647, (fi26 | (fi26 ^ -1))}
                _ = v28
            } else {
                v29 := []int{fi26, (fi26 ^ 2147483647)}
                fmt.Println(v16, fi26)
                for fi30 := 0; fi30 < 12; fi30++ {
                    if len(v29) > 0 {
                        v29[0] = (fi26 + (fi26 + fi30))
                    }
                    if (!true) {
                        v16 += v16
                    } else {
                        if len(v29) > 0 {
                            v29[0] = 914
                        }
                        fmt.Println(fi30, fi26)
                    }
                    if false {
                        switch "" {
                        case "hello":
                            v16 = v16
                        default:
                            fmt.Println(fi30)
                        }
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        v31 := GenS1{F0: 4294967295, F1: -301}
                        v32 := GenS1{F0: 0, F1: -1}
                        var v33 float64 = 1.5
                        switch fi26 {
                        case -9223372036854775808:
                            if true {
                                if false {
                                    var v34 int16 = 42
                                    _ = v34
                                }
                                if (!("foo bar" >= ("\n\t" + (((((("hello" + "a") + "") + "a") + ((("" + "\n\t") + "\n\t") + "\n\t")) + "a") + "")))) {
                                    var v35 int64 = 1
                                    _ = v35
                                } else {
                                    var v36 int32 = 806
                                    v33 -= v33
                                    _ = v36
                                }
                                var v37 int64 = 42
                                _ = v37
                            }
                        case 2147483647:
                            var v38 bool = false
                            _ = v38
                        default:
                            var v39 int = fi26
                            _ = v39
                        }
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        var v40 int = (fi26 | fi30)
                        v16 = 132
                        _ = v40
                    }
                    switch 307 {
                    case 127:
                        if (3.14159 != (1e10 * 0.1)) {
                            var v41 float64 = -2.25
                            _ = v41
                        }
                    }
                    var v42 uint64 = v16
                    v42 -= genFn0(1000, 255, 2147483647)
                    _ = fi30
                    _ = v42
                }
                _ = v29
            }
            v43 := []int{}
            fmt.Println(fi26, fi26)
            _ = fi26
            _ = v43
        }
        for fi44 := 0; fi44 < 16; fi44++ {
            v45 := []int{fi44}
            _ = fi44
            _ = v45
        }
    }
    _ = v16
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 127
    v3 := []uint32{65535, 2147483648}
    var v4 float64 = 1e10
    var v5 uint32 = 2147483648
    for fi46 := 0; fi46 < 17; fi46++ {
        v47 := GenS2{F0: -999, F1: 42, F2: (^(v5))}
        v4 = (v4 * (v4 + v4))
        v1[v5] = false
        _ = fi46
        _ = v47
    }
    if false {
        v2 += v2
        for fi48 := 0; fi48 < 2; fi48++ {
            if (!(263 == 668)) {
                v49 := GenS3{F0: -1, F1: "hello", F2: 590}
                _ = v49
            } else {
                v2 -= v2
                v4 = -2.25
            }
            _ = fi48
        }
        v4 -= (1.5 * 3.14159)
    }
    v50 := GenS2{F0: 772, F1: 127, F2: (v5 | v5)}
    v5 += 741
    if ((v5 % ((v5) | 1)) == (^(v5))) {
        var v51 int16 = 42
        var v52 bool = (v4 < (v4 + 0.0))
        v5 += v5
        _ = v51
        _ = v52
    } else {
        v2 = 2
        var v53 string = "hello"
        _ = v53
    }
    if (32767 < 42) {
        var v54 int8 = -128
        _ = v54
    }
    v2 -= 42
    for fi55 := 0; fi55 < 15; fi55++ {
        v56 := GenS4{F0: 0, F1: 171}
        if false {
            v5 = v5
        } else {
            for fi57 := 0; fi57 < 14; fi57++ {
                var v58 bool = true
                v4 = 0.1
                _ = fi57
                _ = v58
            }
            if ((false || false) || false) {
                var v59 uint8 = (v2 + v2)
                var v60 uint16 = 1
                _ = v59
                _ = v60
            } else {
                for fi61 := 0; fi61 < 19; fi61++ {
                    v5 = v5
                    _ = fi61
                }
                fmt.Println(v4, v2)
            }
            v5 += (^(v5))
        }
        _ = fi55
        _ = v56
    }
    fmt.Println(v5)
    for fi62 := 0; fi62 < 20; fi62++ {
        if true {
            v2 -= v2
        } else {
            if false {
                var v63 int64 = -9223372036854775808
                _ = v63
            } else {
                v64 := GenS2{F0: -169, F1: (fi62 + fi62), F2: (^(v5))}
                switch "\n\t" {
                case "\n\t":
                    switch fi62 {
                    case 1:
                        var v65 uint16 = 255
                        _ = v65
                    case 2147483647:
                        var v66 uint = 255
                        _ = v66
                    case 762:
                        _ = len(v1)
                    }
                case "":
                    v4 = v4
                case "a":
                    switch -479 {
                    case 0:
                        for fi67 := 0; fi67 < 13; fi67++ {
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                var v68 uint = 1
                                _ = v68
                            case "foo bar":
                                delete(v1, (^(v5)))
                            case "":
                                if true {
                                    v4 = 0.1
                                    var v69 float64 = 1e10
                                    var v70 int8 = 1
                                    if (834 < 4294967295) {
                                        v71 := GenS5{F0: "\n\t"}
                                        fmt.Println(v5, v69, fi62)
                                        if true {
                                            v69 = v69
                                            fmt.Println(fi62, fi62, v5)
                                            var v72 bool = (false || (!(v5 >= v5)))
                                            _ = v72
                                        } else {
                                            if (!true) {
                                                v70 += -128
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    var v73 uint8 = (^(v2))
                                                    _ = v73
                                                case "a":
                                                    v70 += (v70 >> ((2) & 63))
                                                    fmt.Println(fi67, v69, v5)
                                                    if false {
                                                        fmt.Println(v5, fi67, v5)
                                                        if (uint32(v2) != 1) {
                                                            switch (fi62 | fi67) {
                                                            case 4294967295:
                                                                var v74 int32 = -922
                                                                _ = v74
                                                            default:
                                                                v2 = (v2 / ((255) | 1))
                                                            }
                                                            v3 = append(v3, v5)
                                                        } else {
                                                            var v75 uint32 = (v5 & 0)
                                                            var v76 float32 = 1e10
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = v5
                                                        }
                                                    }
                                                case "hello":
                                                    v77 := GenS6{F0: v5, F1: 874, F2: (v70 & v70)}
                                                    _ = v77
                                                default:
                                                    var v78 uint64 = 865
                                                    _ = v78
                                                }
                                            }
                                            var v79 float64 = 1.5
                                            _ = v79
                                        }
                                        var v80 int8 = 1
                                        _ = v71
                                        _ = v80
                                    }
                                    var v81 int16 = -32768
                                    v82 := GenS2{F0: int32(v5), F1: (fi62 | fi62), F2: (v5 % ((v5) | 1))}
                                    _ = v69
                                    _ = v70
                                    _ = v81
                                    _ = v82
                                } else {
                                    switch fi62 {
                                    case 32768:
                                        fmt.Println(fi67)
                                        if true {
                                            var v83 int = fi62
                                            _ = len(v1)
                                            _ = v83
                                        }
                                        if true {
                                            v84 := GenS7{F0: v5, F1: 255, F2: (65535 == 42)}
                                            _ = v84
                                        }
                                    }
                                    if (42 > 4294967295) {
                                        var v85 int = (fi62 | (fi62 * fi67))
                                        var v86 uint16 = 32767
                                        _ = v85
                                        _ = v86
                                    }
                                }
                            default:
                                fmt.Println(v5, fi62, v5)
                            }
                            var v87 float64 = 0.0
                            _ = fi67
                            _ = v87
                        }
                    }
                }
                _ = v64
            }
        }
        switch "" {
        case "a":
            v4 = v4
        case "\n\t":
            var v88 bool = false
            _ = v88
        case "hello":
            for fi89 := 0; fi89 < 12; fi89++ {
                var v90 int64 = -1
                _ = fi89
                _ = v90
            }
        default:
            fmt.Println(v4, fi62, fi62)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi62
    }
    for fi91 := 0; fi91 < 6; fi91++ {
        if true {
            v2 += (v2 + 116)
        } else {
            var v92 string = "a"
            _ = v92
        }
        _ = len(v3)
        v2 = (v2 & uint8(v5))
        _ = fi91
    }
    fmt.Println(v4)
    v93 := GenS8{F0: v2, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v93
}
