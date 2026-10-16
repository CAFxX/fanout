package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int32; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 bool; F1 int32; F2 uint32 }
func genFn0() uint {
    v7 := []int{127, -504}
    var v8 int32 = 1
    v8 += v8
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: -804}
        if false {
            v8 = -471
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v8)
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string, p1 uint) uint16 {
    if ((!false) && false) {
        for fi11 := 0; fi11 < 12; fi11++ {
            p1 = (p1 | p1)
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint = 2
                if true {
                    p0 = (p0 + p0)
                    var v14 int = fi12
                    _ = v14
                } else {
                    switch 904 {
                    case 675:
                        if (-2.25 > (3.14159 - (((1.5 + 0.1) / (1.5 * 1e10)) / (3.14159 / ((3.14159 / (((0.1 + 0.1) + 1e10) + 1.5)) * (float64((-2.25 * (3.14159 / -2.25))) - 1.5)))))) {
                            var v15 float64 = -2.25
                            v13 += (p1 % ((p1) | 1))
                            _ = v15
                        }
                    case 9223372036854775807:
                        v13 = (v13 | p1)
                    case -766:
                        fmt.Println(fi12, fi11)
                    }
                }
                fmt.Println(p0)
                _ = fi12
                _ = v13
            }
            var v16 int8 = int8(fi11)
            _ = fi11
            _ = v16
        }
        if (p1 >= p1) {
            v17 := []int{-762, -581}
            _ = v17
        }
        for fi18 := 0; fi18 < 20; fi18++ {
            fmt.Println(p1)
            if (false && (((-2.25 - (-2.25 - 1.5)) / -2.25) >= -2.25)) {
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 uint8 = 1
                    var v21 string = "foo bar"
                    v22 := GenS2{F0: 32767, F1: 42, F2: 4294967295}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            } else {
                p0 = p0
                p0 = (p0 + "hello")
            }
            _ = fi18
        }
    }
    switch "foo bar" {
    case "":
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS1{F0: -1}
            p0 = "a"
            _ = fi23
            _ = v24
        }
    case "\n\t":
        v25 := GenS3{F0: 42, F1: 765}
        _ = v25
    default:
        for fi26 := 0; fi26 < 14; fi26++ {
            p1 = (p1 & (p1 | p1))
            _ = fi26
        }
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []uint32{42, 42, 4294967295, 2147483648}
    v2 := []int{1, 138, -809, 2147483647}
    var v3 int16 = 42
    var v4 bool = (42 > 2)
    var v5 uint = 102
    var v6 uint = 18446744073709551615
    fmt.Println(v5, v6, v6)
    switch 32768 {
    case 4294967295:
        fmt.Println(v4, v3)
    }
    switch "foo bar" {
    case "":
        v5 = v6
    case "\n\t":
        fmt.Println(v5)
    case "hello":
        for fi27 := 0; fi27 < 10; fi27++ {
            for fi28 := 0; fi28 < 12; fi28++ {
                switch fi27 {
                case 32767:
                    v1 = append(v1, 2147483648)
                }
                var v29 uint8 = 2
                v3 = int16(v6)
                _ = fi28
                _ = v29
            }
            var v30 int8 = 42
            v5 = (v6 >> ((v6) & 63))
            _ = fi27
            _ = v30
        }
    }
    v4 = (65535 <= 1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "a" {
    case "":
        _ = len(v2)
    case "a":
        var v31 uint8 = 128
        _ = v31
    default:
        switch "" {
        case "\n\t":
            v6 -= (v6 * 65535)
        case "":
            if false {
                var v32 uint32 = 2
                v6 -= (v5 / ((v6) | 1))
                if true {
                    if ((1e10 - 0.0) == 0.1) {
                        v5 = (v5 - (v6 % ((v6) | 1)))
                        var v33 uint32 = 0
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v32 -= 2
                        _ = fi34
                    }
                } else {
                    if true {
                        if (!(1 != 664)) {
                            var v35 uint8 = 207
                            _ = v35
                        } else {
                            if v4 {
                                v36 := GenS2{F0: 2, F1: 65535, F2: 4294967295}
                                v5 -= 4294967295
                                _ = v36
                            } else {
                                if ((v4 || true) && v4) {
                                    fmt.Println(v4, v6, v6)
                                } else {
                                    var v37 int8 = -1
                                    _ = v37
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        v38 := GenS4{F0: 42, F1: -2147483648, F2: 4294967295}
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 14; fi39++ {
                            v5 = (^(v5))
                            switch ("hello" + ("a" + "a")) {
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "hello":
                                for fi40 := 0; fi40 < 1; fi40++ {
                                    v41 := GenS2{F0: 1, F1: genFn1(("" + ("foo bar" + "\n\t")), (v5 & 42)), F2: 495}
                                    _ = fi40
                                    _ = v41
                                }
                            default:
                                v42 := GenS4{F0: 2147483647, F1: int32(v3), F2: (v32 & 4294967295)}
                                _ = v42
                            }
                            var v43 uint = v5
                            _ = fi39
                            _ = v43
                        }
                        v44 := GenS5{F0: 255, F1: (-682 > -1)}
                        _ = v44
                    }
                }
                _ = v32
            } else {
                var v45 float64 = 1e10
                v4 = (v3 == (v3 | v3))
                _ = v45
            }
        }
    }
    v46 := GenS4{F0: -2147483648, F1: 286, F2: 2147483647}
    v5 = v5
    _ = len(v2)
    var v47 int = 2147483647
    var v48 uint8 = 187
    v48 -= 255
    var v49 uint8 = 166
    if (763 > 0) {
        v50 := GenS6{F0: v4, F1: -2147483648, F2: 921}
        v51 := GenS3{F0: v47, F1: 111}
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = v50
        _ = v51
    } else {
        var v52 uint8 = 0
        var v53 uint32 = 621
        if (v47 != (v47 % ((-256) | 1))) {
            switch "" {
            case "\n\t":
                var v54 float64 = ((0.1 - 1.5) / 3.14159)
                _ = v54
            case "hello":
                fmt.Println(v48)
            }
            for fi55 := 0; fi55 < 1; fi55++ {
                var v56 string = "\n\t"
                _ = fi55
                _ = v56
            }
            if false {
                for fi57 := 0; fi57 < 13; fi57++ {
                    switch "foo bar" {
                    case "a":
                        for fi58 := 0; fi58 < 6; fi58++ {
                            v53 = (^(v53))
                            _ = fi58
                        }
                    default:
                        v6 -= genFn0()
                    }
                    var v59 bool = ((!(!true)) || false)
                    _ = fi57
                    _ = v59
                }
                fmt.Println(v52, v49)
                switch ("a" + "\n\t") {
                case "foo bar":
                    var v60 string = "hello"
                    _ = v60
                case "a":
                    for fi61 := 0; fi61 < 7; fi61++ {
                        v4 = (!(!true))
                        if ((1.5 >= 1.5) || (!(!(v4 && (42 == 1))))) {
                            fmt.Println(v53, v4)
                        } else {
                            switch fi61 {
                            case 32768:
                                v6 += 4294967295
                            case -1:
                                if true {
                                    if false {
                                        v52 += 0
                                        if v4 {
                                            v2 = append(v2, -923)
                                            fmt.Println(v6, v49, v53)
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            for fi62 := 0; fi62 < 12; fi62++ {
                                                var v63 int16 = 0
                                                if (0 == v49) {
                                                    var v64 string = ("hello" + "hello")
                                                    _ = v64
                                                }
                                                v47 += (v47 - fi62)
                                                _ = fi62
                                                _ = v63
                                            }
                                        case "":
                                            if (0 == 31) {
                                                v6 = (v6 + v5)
                                            } else {
                                                switch (^(fi61)) {
                                                case 127:
                                                    var v65 int16 = (v3 / (((v3 & v3)) | 1))
                                                    _ = v65
                                                case 52:
                                                    v4 = (-2147483648 > -574)
                                                default:
                                                    var v66 uint16 = 65535
                                                    _ = v66
                                                }
                                            }
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            v53 -= v53
                                        default:
                                            v67 := GenS4{F0: -271, F1: 42, F2: uint32(v48)}
                                            _ = v67
                                        }
                                        for fi68 := 0; fi68 < 3; fi68++ {
                                            v69 := GenS1{F0: (v3 ^ (v3 / (((v3 / ((v3) | 1))) | 1)))}
                                            var v70 uint = genFn0()
                                            var v71 int32 = 42
                                            fmt.Println(v53, v47)
                                            var v72 int32 = -2147483648
                                            fmt.Println(v53, v52)
                                            _ = fi68
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        }
                                    }
                                    var v73 int = fi61
                                    _ = v73
                                } else {
                                    v6 = 255
                                    v47 -= (^(v47))
                                    switch v47 {
                                    case -830:
                                        fmt.Println(v47)
                                    case -1:
                                        for fi74 := 0; fi74 < 2; fi74++ {
                                            v3 -= (v3 ^ v3)
                                            v5 = v6
                                            _ = fi74
                                        }
                                    default:
                                        for fi75 := 0; fi75 < 16; fi75++ {
                                            v76 := GenS5{F0: (v52 << ((v5) & 63)), F1: true}
                                            if ((!v4) && false) {
                                                var v77 int = (fi75 - fi75)
                                                _ = v77
                                            }
                                            _ = fi75
                                            _ = v76
                                        }
                                    }
                                }
                            }
                            var v78 bool = v4
                            _ = v78
                        }
                        switch "" {
                        case "a":
                            v79 := GenS4{F0: 2147483647, F1: -2147483648, F2: v53}
                            _ = v79
                        case "foo bar":
                            v5 += v6
                        default:
                            for fi80 := 0; fi80 < 16; fi80++ {
                                var v81 uint = 2
                                v5 += uint(v53)
                                v82 := GenS4{F0: 42, F1: 42, F2: (v53 & v53)}
                                v52 = (v52 >> ((v6) & 63))
                                v81 = v81
                                if ((4294967295 > 2) || false) {
                                    switch "foo bar" {
                                    case "a":
                                        var v83 float64 = 0.0
                                        v84 := GenS3{F0: (v47 / ((v47) | 1)), F1: 81}
                                        var v85 int32 = -2147483648
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    case "\n\t":
                                        v3 -= v3
                                    default:
                                        var v86 string = ""
                                        _ = v86
                                    }
                                    fmt.Println(v47, v6)
                                    var v87 int = (fi80 - 42)
                                    _ = v87
                                }
                                _ = fi80
                                _ = v81
                                _ = v82
                            }
                        }
                        _ = fi61
                    }
                case "\n\t":
                    v4 = (v3 >= (v3 & int16(v48)))
                default:
                    v4 = v4
                }
            } else {
                var v88 int64 = 946
                fmt.Println(v88, v3)
                _ = v88
            }
        } else {
            for fi89 := 0; fi89 < 17; fi89++ {
                for fi90 := 0; fi90 < 1; fi90++ {
                    if v4 {
                        v91 := GenS1{F0: (^(v3))}
                        switch ("" + ("hello" + "foo bar")) {
                        case "hello":
                            v5 = genFn0()
                        default:
                            var v92 uint8 = (v52 + (^(v49)))
                            _ = v92
                        }
                        _ = v91
                    } else {
                        var v93 uint64 = 4294967295
                        var v94 bool = v4
                        var v95 int32 = -1
                        _ = v93
                        _ = v94
                        _ = v95
                    }
                    var v96 uint8 = 0
                    var v97 uint64 = 2
                    _ = fi90
                    _ = v96
                    _ = v97
                }
                v49 -= (v49 >> ((v6) & 63))
                _ = fi89
            }
        }
        _ = v52
        _ = v53
    }
    if (42 < 0) {
        for fi98 := 0; fi98 < 9; fi98++ {
            var v99 uint8 = (^(v48))
            var v100 int64 = 9223372036854775807
            _ = fi98
            _ = v100
            _ = v99
        }
    }
    if (((("foo bar" + ("\n\t" + "hello")) + "a") + "a") != "hello") {
        for fi101 := 0; fi101 < 13; fi101++ {
            v2 = append(v2, (-(v47)))
            _ = fi101
        }
        switch -7 {
        case 0:
            var v102 uint8 = 255
            _ = v102
        default:
            fmt.Println(v3)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
