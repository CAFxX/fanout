package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 string; F1 int32; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint32; F2 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 bool; F2 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool }
func genFn0() int64 {
    var v4 uint16 = 65535
    var v5 uint16 = v4
    v6 := GenS1{F0: 1, F1: 250}
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 string, p1 int16, p2 int) uint64 {
    var v7 int16 = (^(p1))
    p1 = (p1 / ((p1) | 1))
    var v8 int16 = 1000
    var v9 uint = 173
    switch 42 {
    case 32767:
        switch p0 {
        case "hello":
            for fi10 := 0; fi10 < 8; fi10++ {
                v8 = v7
                v11 := GenS1{F0: (v8 ^ v8), F1: 65535}
                p1 = p1
                _ = fi10
                _ = v11
            }
        case "":
            v12 := GenS1{F0: p1, F1: 465}
            _ = v12
        default:
            var v13 float32 = float32(1e10)
            _ = v13
        }
    case -345:
        if (232 > 404) {
            var v14 uint64 = 2
            v15 := GenS2{F0: -965}
            fmt.Println(v8)
            _ = v14
            _ = v15
        } else {
            p2 += p2
            if true {
                var v16 int64 = -9223372036854775808
                v16 += v16
                if (2 > 714) {
                    v17 := []int{}
                    switch p2 {
                    case 351:
                        switch p0 {
                        case "\n\t":
                            for fi18 := 0; fi18 < 10; fi18++ {
                                if false {
                                    v19 := GenS1{F0: p1, F1: 2147483648}
                                    p0 = "\n\t"
                                    v20 := GenS2{F0: 1}
                                    _ = v19
                                    _ = v20
                                } else {
                                    if true {
                                        switch (p2 * -9223372036854775808) {
                                        case -9223372036854775808:
                                            if len(v17) > 0 {
                                                v17[0] = -2147483648
                                            }
                                        case -1:
                                            v7 += 714
                                        }
                                        var v21 uint = 65535
                                        var v22 int16 = (-(p1))
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v23 uint = 0
                                        _ = v23
                                    }
                                    if len(v17) > 0 {
                                        v17[0] = (fi18 + (fi18 & (p2 << ((v9) & 63))))
                                    }
                                    switch -723 {
                                    case -533:
                                        v17 = append(v17, 32767)
                                    }
                                }
                                switch (fi18 | (fi18 - p2)) {
                                case 32768:
                                    var v24 uint64 = 1
                                    _ = v24
                                case -9223372036854775808:
                                    switch fi18 {
                                    case -1:
                                        for fi25 := 0; fi25 < 12; fi25++ {
                                            _ = len(v17)
                                            fmt.Println(p2, p2)
                                            v16 += (v16 - (v16 + (v16 - v16)))
                                            _ = fi25
                                        }
                                    case 119:
                                        switch -1 {
                                        case 4294967295:
                                            v26 := GenS2{F0: 42}
                                            _ = v26
                                        }
                                    case 32767:
                                        if (-9223372036854775808 == v16) {
                                            for fi27 := 0; fi27 < 15; fi27++ {
                                                p1 -= 0
                                                var v28 int = (p2 << ((255) & 63))
                                                fmt.Println(fi27)
                                                _ = fi27
                                                _ = v28
                                            }
                                        }
                                    }
                                default:
                                    var v29 uint32 = 2147483648
                                    _ = v29
                                }
                                _ = fi18
                            }
                        case "":
                            for fi30 := 0; fi30 < 4; fi30++ {
                                var v31 int = fi30
                                _ = fi30
                                _ = v31
                            }
                        }
                    default:
                        v32 := GenS1{F0: -330, F1: 0}
                        _ = v32
                    }
                    _ = v17
                }
                _ = v16
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                v34 := GenS1{F0: (v8 / (((v8 << (((v9 >> ((961) & 63))) & 63))) | 1)), F1: 2147483648}
                _ = fi33
                _ = v34
            }
        }
    case 128:
        for fi35 := 0; fi35 < 1; fi35++ {
            v7 -= (v8 % ((v7) | 1))
            _ = fi35
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 42
    var v2 string = ("" + "")
    v3 := make(map[uint32]bool)
    v2 = (v2 + (v2 + v2))
    v36 := GenS3{F0: 2147483647, F1: v2}
    var v37 uint64 = 18446744073709551615
    v38 := GenS3{F0: uint32(v37), F1: ("foo bar" + "")}
    v39 := []int{}
    var v40 uint32 = 174
    if ((!(v1 < v1)) || true) {
        v2 = v2
        if false {
            delete(v3, uint32(v37))
            var v41 uint64 = genFn1(v2, 365, 0)
            v40 = 4294967295
            _ = v41
        } else {
            var v42 float32 = 0.1
            var v43 int = 2147483647
            _ = v42
            _ = v43
        }
        if (!((-597 >= 9223372036854775807) && (127 == -874))) {
            for fi44 := 0; fi44 < 14; fi44++ {
                switch "hello" {
                case "a":
                    for fi45 := 0; fi45 < 3; fi45++ {
                        v37 = 2
                        v40 += v40
                        fmt.Println(v37, v37)
                        _ = fi45
                    }
                case "foo bar":
                    v46 := GenS3{F0: v40, F1: (("a" + "foo bar") + v2)}
                    _ = v46
                default:
                    var v47 int = 1
                    _ = v47
                }
                for fi48 := 0; fi48 < 4; fi48++ {
                    v49 := GenS4{F0: "a", F1: 480, F2: genFn1(v2, 618, fi48)}
                    switch v2 {
                    case "\n\t":
                        if (!false) {
                            var v50 int8 = -1
                            v51 := GenS2{F0: 0}
                            var v52 int64 = 9223372036854775807
                            _ = v50
                            _ = v51
                            _ = v52
                        } else {
                            v53 := GenS1{F0: -361, F1: 845}
                            _ = len(v39)
                            v40 = uint32(v37)
                            _ = v53
                        }
                    case "a":
                        var v54 uint64 = uint64(v1)
                        _ = v54
                    }
                    fmt.Println(v40, fi44, v40)
                    _ = fi48
                    _ = v49
                }
                _ = fi44
            }
            for fi55 := 0; fi55 < 13; fi55++ {
                v37 = v37
                for fi56 := 0; fi56 < 14; fi56++ {
                    v39 = append(v39, (fi56 >> ((65535) & 63)))
                    var v57 float64 = ((0.0 - 0.0) - 1e10)
                    if true {
                        v1 += v1
                    } else {
                        if len(v39) >= 2 {
                            v39 = v39[1:len(v39)]
                        }
                        var v58 int = (fi55 % ((int(v40)) | 1))
                        _ = v58
                    }
                    if (-9223372036854775808 != fi55) {
                        if len(v39) > 0 {
                            v39[0] = 1
                        }
                        switch fi55 {
                        case 540:
                            v59 := v3[(v40 + 538)]
                            _ = v59
                        case -9223372036854775808:
                            var v60 uint8 = 125
                            _ = v60
                        case 32767:
                            if true {
                                switch fi56 {
                                case -480:
                                    v2 = v2
                                case -2147483648:
                                    if true {
                                        v2 = (((v2 + "\n\t") + (("\n\t" + v2) + (v2 + v2))) + v2)
                                        v57 = float64(0.1)
                                        if (v57 == (v57 + (-2.25 - v57))) {
                                            var v61 bool = true
                                            v3[v40] = (uint64(v40) >= 1)
                                            v62 := GenS5{F0: 0}
                                            _ = v61
                                            _ = v62
                                        }
                                    }
                                case 42:
                                    v63 := GenS6{F0: v2, F1: (^(v40)), F2: ((-1 == 2147483647) || (!false))}
                                    _ = v63
                                default:
                                    var v64 int32 = 789
                                    _ = v64
                                }
                                if true {
                                    var v65 uint16 = 255
                                    fmt.Println(v57)
                                    fmt.Println(fi55)
                                    v37 -= (v37 & (v37 ^ 11))
                                    _ = v65
                                }
                                switch fi56 {
                                case 11:
                                    v66 := GenS7{F0: true}
                                    _ = v66
                                case 9223372036854775807:
                                    v67 := GenS1{F0: -999, F1: v40}
                                    _ = v67
                                case 248:
                                    if true {
                                        if len(v39) >= 2 {
                                            v39 = v39[1:len(v39)]
                                        }
                                    }
                                }
                                v68 := GenS6{F0: v2, F1: 42, F2: true}
                                _ = v68
                            } else {
                                var v69 int32 = 190
                                var v70 int = 920
                                if false {
                                    v71 := GenS1{F0: -857, F1: 617}
                                    v70 = 917
                                    var v72 int16 = 399
                                    _ = v71
                                    _ = v72
                                }
                                _ = v69
                                _ = v70
                            }
                        default:
                            switch 127 {
                            case 994:
                                var v73 int8 = 100
                                _ = v73
                            default:
                                var v74 string = v2
                                _ = v74
                            }
                        }
                    } else {
                        if len(v39) >= 2 {
                            v39 = v39[1:len(v39)]
                        }
                        var v75 int32 = 647
                        _ = v75
                    }
                    _ = fi56
                    _ = v57
                }
                v37 = (^(v37))
                _ = fi55
            }
        } else {
            v37 -= (v37 * 42)
            fmt.Println(v1, v1)
        }
    }
    v76 := GenS3{F0: (^(v40)), F1: (v2 + (((v2 + "hello") + (v2 + "hello")) + v2))}
    for fi77 := 0; fi77 < 14; fi77++ {
        var v78 int16 = -1
        var v79 uint = 690
        v80 := GenS8{F0: 42, F1: true, F2: -726}
        _ = fi77
        _ = v78
        _ = v79
        _ = v80
    }
    v81 := GenS9{F0: 42}
    v37 = genFn1(v2, -1, -589)
    var v82 int8 = -36
    var v83 int32 = -1
    v84 := GenS2{F0: v83}
    var v85 uint16 = (v1 << ((4294967295) & 63))
    if true {
        var v86 int = 358
        v87 := GenS10{F0: true}
        _ = v86
        _ = v87
    } else {
        for fi88 := 0; fi88 < 19; fi88++ {
            var v89 int = (fi88 * fi88)
            _ = fi88
            _ = v89
        }
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v40:", v40)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v76
    _ = v81
    _ = v82
    _ = v83
    _ = v84
    _ = v85
}
