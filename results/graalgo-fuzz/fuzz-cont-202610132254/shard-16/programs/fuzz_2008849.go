package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
type GenS4 struct { F0 int32; F1 bool; F2 int16 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint; F2 int8 }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int8; F1 int; F2 int64 }
func genFn0(p0 uint64, p1 int64, p2 int16) uint {
    var v6 uint = 65535
    v7 := GenS1{F0: -1}
    if ((true || (100 == 1)) || (0.0 >= 3.14159)) {
        var v8 int64 = p1
        _ = v8
    } else {
        for fi9 := 0; fi9 < 12; fi9++ {
            p1 = (p1 % (((p1 - 142)) | 1))
            var v10 float32 = (float32(((float64(-2.25) - 1.5) * 1.5)) + (1.5 / 1.0))
            fmt.Println(v6)
            _ = fi9
            _ = v10
        }
        var v11 uint64 = p0
        var v12 int32 = 1
        _ = v11
        _ = v12
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int32, p1 uint64, p2 uint) uint16 {
    p2 -= (p2 << ((65535) & 63))
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint32 = 65535
        if true {
            if true {
                if (!(42 != -530)) {
                    var v15 int = -682
                    _ = v15
                }
                switch "\n\t" {
                case "":
                    var v16 uint32 = (v14 * (v14 / ((v14) | 1)))
                    _ = v16
                case "foo bar":
                    p2 = p2
                }
            }
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                if (!(2 != 127)) {
                    var v18 int16 = 1
                    _ = v18
                }
                _ = fi17
            }
        }
        var v19 uint16 = 1
        _ = fi13
        _ = v14
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = 1
    var v3 uint8 = 42
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    var v20 uint32 = 0
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int32 = 324
        _ = fi21
        _ = v22
    }
    var v23 uint64 = (v1 / ((v1) | 1))
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := []int{(fi24 & fi24), 9223372036854775807, 128, (^(fi24))}
        if len(v25) > 0 {
            v25[0] = 443
        }
        switch ("a" + ("\n\t" + "")) {
        case "hello":
            if true {
                switch (-(fi24)) {
                case 560:
                    if (0 <= 42) {
                        v26 := GenS1{F0: v2}
                        v27 := GenS1{F0: v2}
                        v20 = (v20 + v20)
                        _ = v26
                        _ = v27
                    } else {
                        var v28 bool = ((true && (!((^(v23)) < v1))) && ((v3 | v3) >= 128))
                        v1 = (v23 << ((589) & 63))
                        _ = v28
                    }
                case -313:
                    if (!false) {
                        for fi29 := 0; fi29 < 17; fi29++ {
                            v25 = append(v25, fi29)
                            _ = fi29
                        }
                        var v30 string = "\n\t"
                        fmt.Println(v30)
                        _ = v30
                    } else {
                        v31 := GenS2{F0: 9223372036854775807}
                        v3 = (v3 | (v3 << ((2) & 63)))
                        v3 = v3
                        _ = v31
                    }
                case 128:
                    v25 = append(v25, fi24)
                }
                if (1 > -1) {
                    if true {
                        if false {
                            for fi32 := 0; fi32 < 17; fi32++ {
                                if len(v25) > 0 {
                                    v25[0] = fi24
                                }
                                v23 = v23
                                _ = fi32
                            }
                            v33 := GenS3{F0: 1, F1: (v20 * (v20 % (((^(v20))) | 1))), F2: -685}
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 8; fi34++ {
                            fmt.Println(v1, fi34, v23)
                            _ = fi34
                        }
                    } else {
                        switch fi24 {
                        case 0:
                            v23 = v23
                        case -38:
                            v35 := GenS3{F0: -771, F1: 1, F2: 0}
                            _ = v35
                        case -1:
                            var v36 int64 = int64(v2)
                            _ = v36
                        default:
                            for fi37 := 0; fi37 < 15; fi37++ {
                                if false {
                                    var v38 bool = true
                                    fmt.Println(v3, fi37)
                                    _ = v38
                                } else {
                                    if true {
                                        if ((v1 % ((v1) | 1)) == v23) {
                                            fmt.Println(v1, fi24)
                                        }
                                        v23 -= (v1 * (v23 >> ((42) & 63)))
                                    }
                                }
                                v39 := GenS1{F0: v2}
                                _ = fi37
                                _ = v39
                            }
                        }
                        var v40 float32 = 0.0
                        v41 := GenS4{F0: -2147483648, F1: true, F2: -32768}
                        _ = v40
                        _ = v41
                    }
                    var v42 bool = (int(v2) < (^(fi24)))
                    fmt.Println(v3, v1)
                    _ = v42
                } else {
                    if (950 >= 0) {
                        if (fi24 > fi24) {
                            if false {
                                if false {
                                    fmt.Println(v23, v2, v20)
                                    var v43 int32 = 1
                                    for fi44 := 0; fi44 < 18; fi44++ {
                                        v45 := GenS5{F0: (^(v1)), F1: -1, F2: (v20 << ((0) & 63))}
                                        v46 := GenS6{F0: 2, F1: 255, F2: v2}
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v43
                                } else {
                                    var v47 int32 = 389
                                    switch (((("foo bar" + ("\n\t" + "\n\t")) + "a") + "foo bar") + "\n\t") {
                                    case "hello":
                                        var v48 string = ("hello" + "hello")
                                        _ = v48
                                    case "a":
                                        if true {
                                            v49 := GenS7{F0: 1, F1: v20}
                                            for fi50 := 0; fi50 < 7; fi50++ {
                                                v4[fi50] = "\n\t"
                                                _ = fi50
                                            }
                                            for fi51 := 0; fi51 < 12; fi51++ {
                                                v1 = (^(v1))
                                                _ = fi51
                                            }
                                            _ = v49
                                        }
                                    case "foo bar":
                                        v52 := GenS8{F0: v20, F1: (v20 & v20)}
                                        _ = v52
                                    default:
                                        fmt.Println(fi24)
                                    }
                                    var v53 uint64 = (v1 - (^(v23)))
                                    _ = v47
                                    _ = v53
                                }
                                var v54 float64 = 3.14159
                                _ = v54
                            }
                            var v55 uint = 4294967295
                            v1 -= 42
                            _ = v55
                        }
                        _ = len(v25)
                    }
                    var v56 int8 = -128
                    v56 -= v2
                    _ = v56
                }
                for fi57 := 0; fi57 < 20; fi57++ {
                    var v58 uint8 = (v3 >> ((847) & 63))
                    _ = fi57
                    _ = v58
                }
            } else {
                for fi59 := 0; fi59 < 4; fi59++ {
                    _ = len(v5)
                    v5[fi59] = (703 == 2)
                    _ = fi59
                }
                if true {
                    v4[fi24] = (("foo bar" + "\n\t") + "foo bar")
                    if false {
                        v3 += v3
                        for fi60 := 0; fi60 < 12; fi60++ {
                            v25 = append(v25, (fi60 | fi24))
                            _ = fi60
                        }
                    }
                    fmt.Println(v3, v1, v1)
                } else {
                    var v61 uint64 = v1
                    for fi62 := 0; fi62 < 11; fi62++ {
                        if true {
                            v61 += uint64(v20)
                        }
                        _ = fi62
                    }
                    v20 += v20
                    _ = v61
                }
            }
        default:
            for fi63 := 0; fi63 < 16; fi63++ {
                var v64 int64 = -9223372036854775808
                var v65 float64 = (3.14159 - 0.0)
                _ = fi63
                _ = v64
                _ = v65
            }
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v20, v1)
    var v66 float64 = 3.14159
    for fi67 := 0; fi67 < 8; fi67++ {
        v4[-39] = ("hello" + "foo bar")
        if false {
            if (int64(fi67) != 0) {
                for fi68 := 0; fi68 < 16; fi68++ {
                    if true {
                        v2 = v2
                    } else {
                        v1 = (v1 - (v1 << ((4294967295) & 63)))
                        fmt.Println(v20)
                        v69 := GenS6{F0: 961, F1: genFn0(4294967295, 0, -32768), F2: v2}
                        _ = v69
                    }
                    _ = fi68
                }
                v20 -= v20
            }
            var v70 uint64 = (v23 << ((1) & 63))
            _ = v70
        }
        _ = fi67
    }
    v23 += v1
    if ((true || true) || (-1 <= 0)) {
        v1 = (v23 ^ (v1 << ((4294967295) & 63)))
        v2 = (v2 % (((v2 << ((65535) & 63))) | 1))
        for fi71 := 0; fi71 < 7; fi71++ {
            switch "" {
            case "a":
                v2 -= v2
            default:
                switch fi71 {
                case 42:
                    if true {
                        v66 -= (-2.25 * float64(0.1))
                    }
                case -1:
                    v20 -= (v20 & v20)
                default:
                    if (!(578 < 0)) {
                        if (0.1 <= v66) {
                            v1 = 423
                            var v72 int8 = -1
                            _ = v72
                        }
                    } else {
                        if false {
                            fmt.Println(v2)
                            var v73 int16 = 1000
                            _ = v73
                        }
                        var v74 int = fi71
                        _ = v74
                    }
                }
            }
            _ = fi71
        }
    } else {
        v2 += (-(v2))
    }
    v75 := GenS2{F0: 271}
    if true {
        var v76 float32 = 1e10
        var v77 uint8 = v3
        _ = v76
        _ = v77
    } else {
        var v78 int = -244
        _ = v78
    }
    var v79 int32 = -2147483648
    if (3.14159 <= ((v66 / v66) * v66)) {
        for fi80 := 0; fi80 < 16; fi80++ {
            v81 := GenS5{F0: 50, F1: (v79 % ((2147483647) | 1)), F2: uint32(v23)}
            if false {
                fmt.Println(v1)
                v23 = (v23 / (((v23 + 2)) | 1))
            } else {
                for fi82 := 0; fi82 < 17; fi82++ {
                    switch "a" {
                    case "\n\t":
                        var v83 int32 = (v79 | v79)
                        _ = v83
                    case "foo bar":
                        switch "\n\t" {
                        case "\n\t":
                            var v84 uint32 = v20
                            _ = v84
                        default:
                            if ((^(v20)) <= v20) {
                                fmt.Println(fi82, v2, v20)
                                switch (("\n\t" + "") + ("" + "foo bar")) {
                                case "\n\t":
                                    var v85 bool = (!true)
                                    _ = v85
                                case "foo bar":
                                    var v86 uint64 = 18446744073709551615
                                    _ = v86
                                case "hello":
                                    var v87 int8 = (v2 * (v2 << ((uint(v1)) & 63)))
                                    _ = v87
                                }
                            } else {
                                var v88 int8 = int8(v3)
                                _ = v88
                            }
                        }
                    default:
                        var v89 uint64 = uint64(v20)
                        _ = v89
                    }
                    v90 := GenS7{F0: (v1 ^ 844), F1: v20}
                    _ = fi82
                    _ = v90
                }
            }
            var v91 uint8 = (^(v3))
            _ = fi80
            _ = v81
            _ = v91
        }
        v92 := GenS4{F0: (v79 | v79), F1: ((1.5 * 3.14159) < (float32(v66) * -2.25)), F2: 0}
        var v93 int32 = (v79 + (v79 - (v79 | int32(v2))))
        _ = v92
        _ = v93
    }
    for fi94 := 0; fi94 < 8; fi94++ {
        var v95 bool = (false && (!(((v3 < v3) || true) && false)))
        v96 := GenS9{F0: (-(v2)), F1: fi94, F2: 979}
        _ = fi94
        _ = v95
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v66:", v66)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v75
    _ = v79
}
