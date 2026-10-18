package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool; F1 int; F2 int8 }
type GenS4 struct { F0 string; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int; F2 int64 }
type GenS6 struct { F0 int; F1 int64; F2 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int16, p1 uint32, p2 bool) uint64 {
    v6 := []int{}
    if p2 {
        v7 := GenS1{F0: 2, F1: (0 != -1), F2: -1}
        var v8 int = 42
        _ = v7
        _ = v8
    } else {
        var v9 uint16 = 678
        _ = v9
    }
    v10 := GenS2{F0: 1, F1: p0}
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return 22
}
func genFn1() int {
    v11 := []int{}
    v12 := GenS3{F0: (874 == 1), F1: 2147483647, F2: -70}
    var v13 bool = (!(175 < 128))
    if ((!false) || v13) {
        v14 := GenS4{F0: "\n\t", F1: 372, F2: 0}
        if v13 {
            if v13 {
                v13 = ("hello" != ("" + ("" + "a")))
                v13 = v13
                v13 = (!v13)
            } else {
                switch "foo bar" {
                case "foo bar":
                    v15 := GenS5{F0: 1, F1: 42, F2: -775}
                    _ = v15
                case "":
                    v16 := GenS6{F0: -768, F1: -1, F2: 1}
                    _ = v16
                default:
                    if ((!(v13 && v13)) && (2 < 2)) {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            if (127 < 255) {
                                v13 = v13
                                fmt.Println(fi17)
                                switch 9223372036854775807 {
                                case 127:
                                    for fi18 := 0; fi18 < 20; fi18++ {
                                        switch (fi17 >> ((2) & 63)) {
                                        case 32768:
                                            var v19 int = (fi17 >> ((163) & 63))
                                            _ = v19
                                        default:
                                            switch "\n\t" {
                                            case "hello":
                                                v13 = (!false)
                                            default:
                                                v13 = (65535 <= 580)
                                            }
                                        }
                                        fmt.Println(v13, fi18, v13)
                                        _ = fi18
                                    }
                                default:
                                    var v20 int8 = 100
                                    _ = v20
                                }
                            }
                            if (426 != 1) {
                                switch "a" {
                                case "a":
                                    _ = len(v11)
                                }
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    var v22 uint32 = 42
                                    fmt.Println(v22)
                                    v22 = (v22 / (((v22 + v22)) | 1))
                                    _ = fi21
                                    _ = v22
                                }
                            }
                            _ = fi17
                        }
                    }
                }
                v11 = append(v11, 1)
                v23 := GenS7{F0: 708}
                _ = v23
            }
            v13 = (587 < 2)
        }
        var v24 float64 = 0.1
        _ = v14
        _ = v24
    } else {
        v25 := GenS3{F0: (42 == 42), F1: 127, F2: 71}
        var v26 int16 = 42
        if (32767 > 32767) {
            switch 2147483647 {
            case 1:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            default:
                if len(v11) > 0 {
                    v11[0] = -342
                }
            }
        }
        _ = v25
        _ = v26
    }
    if v13 {
        if v13 {
            _ = len(v11)
            switch 4294967295 {
            case -9223372036854775808:
                var v27 uint64 = 689
                _ = v27
            case 393:
                if len(v11) > 0 {
                    v11[0] = 4294967295
                }
            case -701:
                v13 = (-32768 == 32767)
            default:
                var v28 float64 = (((1.5 + (1.5 / 1e10)) / ((3.14159 * 0.1) - (1e10 + (((float64((((-2.25 / (0.1 - 3.14159)) * float32(0.0)) + ((0.1 * 1.5) + 3.14159))) - float64((1.5 + 0.0))) + 1.5) * 0.1)))) / float64((float32(float64(0.1)) + 0.0)))
                _ = v28
            }
            switch 128 {
            case 0:
                fmt.Println(v13)
            case -2147483648:
                _ = len(v11)
            case 42:
                var v29 uint32 = 594
                _ = v29
            }
        } else {
            if (!(42 > 32767)) {
                v11 = append(v11, 942)
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 int = fi30
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = fi30
                    _ = v31
                }
                v32 := GenS5{F0: 51, F1: 128, F2: 42}
                _ = v32
            }
            var v33 uint64 = 18446744073709551615
            v33 = v33
            _ = v33
        }
        switch (("" + "foo bar") + "") {
        case "\n\t":
            var v34 int16 = 186
            _ = v34
        case "a":
            for fi35 := 0; fi35 < 8; fi35++ {
                if (((((1e10 * (-2.25 + 1.5)) / 1e10) * (((1e10 / ((float32((0.1 / ((float64((float32((1e10 - 0.1)) / 1.0)) + 1e10) / (0.1 / 3.14159)))) * 1e10) - (-2.25 * float32((-2.25 - 0.1))))) - (0.1 + ((3.14159 + (0.0 + (0.1 + 0.0))) + float32((-2.25 / (1.5 - -2.25)))))) * -2.25)) + (-2.25 / 1.5)) != ((float32(0.0) + float32(float64((1e10 * 3.14159)))) * 0.0)) {
                    v13 = (0 >= 4294967295)
                    v13 = true
                } else {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    var v36 uint16 = 32767
                    _ = v36
                }
                var v37 string = ("" + "a")
                _ = fi35
                _ = v37
            }
        case "foo bar":
            fmt.Println(v13)
        }
        if v13 {
            v13 = v13
            v13 = (18446744073709551615 > 42)
        }
    }
    _ = v11
    _ = v12
    _ = v13
    return 32767
}

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 32767
    var v3 float64 = (-2.25 / 1.5)
    var v4 int64 = -897
    v5 := make(map[string]int)
    if false {
        v4 = 693
    } else {
        switch "\n\t" {
        case "a":
            v38 := GenS1{F0: 87, F1: false, F2: (v4 * (v4 << (((v1 / ((1) | 1))) & 63)))}
            _ = v38
        default:
            for fi39 := 0; fi39 < 4; fi39++ {
                fmt.Println(v3, v4, v2)
                var v40 int64 = 0
                v41 := GenS8{F0: 568, F1: 982}
                _ = fi39
                _ = v40
                _ = v41
            }
        }
    }
    var v42 int16 = int16(v1)
    switch 32767 {
    case 1:
        var v43 int64 = 42
        _ = v43
    case -2147483648:
        for fi44 := 0; fi44 < 12; fi44++ {
            v45 := v5["foo bar"]
            v3 += v3
            v46 := []int{(fi44 * fi44), (-(v45)), (fi44 % ((v45) | 1))}
            _ = fi44
            _ = v45
            _ = v46
        }
    default:
        if false {
            fmt.Println(v3, v3)
            var v47 uint16 = 797
            _ = len(v5)
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 6; fi48++ {
        v1 = 4294967295
        for fi49 := 0; fi49 < 11; fi49++ {
            var v50 int64 = (v4 % ((-9223372036854775808) | 1))
            v51 := GenS2{F0: 744, F1: v42}
            v42 -= (v2 - v2)
            _ = fi49
            _ = v50
            _ = v51
        }
        for fi52 := 0; fi52 < 14; fi52++ {
            v3 = float64(-2.25)
            fmt.Println(v4, fi48)
            var v53 uint32 = 2
            _ = fi52
            _ = v53
        }
        _ = fi48
    }
    v54 := GenS4{F0: "a", F1: 1, F2: 2147483647}
    var v55 int8 = 121
    var v56 string = ("hello" + ("hello" + ("foo bar" + "foo bar")))
    v57 := []int{}
    var v58 float64 = 0.0
    if false {
        v59 := v5[""]
        _ = v59
    } else {
        fmt.Println(v55)
        _ = len(v5)
        fmt.Println(v42, v2, v1)
    }
    if (v2 >= v42) {
        v60 := v5[(((v56 + "hello") + v56) + "\n\t")]
        v55 -= v55
        for fi61 := 0; fi61 < 5; fi61++ {
            if len(v57) >= 2 {
                v57 = v57[1:len(v57)]
            }
            _ = fi61
        }
        _ = v60
    } else {
        if (false && false) {
            if false {
                if false {
                    var v62 int32 = 42
                    v55 = v55
                    v1 = 539
                    _ = v62
                } else {
                    for fi63 := 0; fi63 < 12; fi63++ {
                        fmt.Println(v1, fi63)
                        _ = len(v5)
                        var v64 int64 = 42
                        _ = fi63
                        _ = v64
                    }
                }
            } else {
                var v65 int8 = (v55 << ((v1) & 63))
                _ = v65
            }
            v66 := v5[(v56 + ("" + v56))]
            _ = v66
        }
        var v67 float64 = (v3 + v3)
        _ = v67
    }
    v1 -= v1
    if (0.0 < (v58 - v58)) {
        switch int(v4) {
        case 9223372036854775807:
            var v68 uint8 = 128
            _ = v68
        case 42:
            if (42 == 42) {
                var v69 int32 = 42
                var v70 int16 = 1
                _ = v69
                _ = v70
            }
        case -1:
            var v71 float32 = -2.25
            _ = v71
        default:
            var v72 string = "a"
            _ = v72
        }
        switch ((v56 + v56) + (v56 + v56)) {
        case "":
            if (true && ((2 == 1) || ((!true) || true))) {
                fmt.Println(v2, v1, v2)
            }
        case "\n\t":
            var v73 uint16 = 32767
            _ = v73
        case "foo bar":
            for fi74 := 0; fi74 < 3; fi74++ {
                var v75 uint8 = 255
                if ((v42 ^ (v42 << ((v1) & 63))) >= -208) {
                    var v76 string = (v56 + v56)
                    if (!(!false)) {
                        v56 = "a"
                    } else {
                        if false {
                            var v77 bool = true
                            var v78 uint = 934
                            var v79 int32 = 2147483647
                            _ = v77
                            _ = v78
                            _ = v79
                        } else {
                            v80 := GenS6{F0: (fi74 >> (((v1 % (((v1 & v1)) | 1))) & 63)), F1: int64(v55), F2: (fi74 - (fi74 % (((^(fi74))) | 1)))}
                            fmt.Println(v58, v42)
                            _ = v80
                        }
                    }
                    v56 = v56
                    _ = v76
                } else {
                    if false {
                        for fi81 := 0; fi81 < 11; fi81++ {
                            v3 -= 1.5
                            _ = fi81
                        }
                        v82 := GenS3{F0: (42 == 255), F1: (fi74 + genFn1()), F2: -7}
                        v5["foo bar"] = fi74
                        _ = v82
                    } else {
                        v83 := v5[v56]
                        _ = v83
                    }
                }
                v3 = v58
                _ = fi74
                _ = v75
            }
        default:
            fmt.Println(v1, v3)
        }
    }
    var v84 int8 = -5
    v84 -= (v55 + (v55 / (((v55 + (v55 - v84))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v84
}
