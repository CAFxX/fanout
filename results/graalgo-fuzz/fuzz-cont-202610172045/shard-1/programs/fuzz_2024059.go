package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 int }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint8; F1 string; F2 uint8 }
func genFn0() int32 {
    switch "\n\t" {
    case "hello":
        var v5 uint16 = 749
        _ = v5
    case "":
        v6 := []int{}
        _ = v6
    default:
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: -1}
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 2147483647}
    var v10 int8 = -128
    v10 -= v10
    var v11 uint64 = 296
    _ = v10
    _ = v11
    _ = v9
    return 42
}
func genFn1(p0 int, p1 int, p2 int64) string {
    if false {
        var v12 uint64 = 494
        var v13 bool = (!true)
        v14 := []int{-456}
        _ = v12
        _ = v13
        _ = v14
    } else {
        if false {
            v15 := GenS1{F0: -2147483648}
            v16 := GenS2{F0: 65535, F1: "foo bar", F2: (p0 * (p0 + (p0 + p0)))}
            var v17 int32 = -2147483648
            _ = v15
            _ = v16
            _ = v17
        } else {
            p1 -= 32768
            switch (p1 ^ (p1 + (p0 ^ p1))) {
            case -662:
                for fi18 := 0; fi18 < 1; fi18++ {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS3{F0: -1, F1: 1000, F2: 0}
                        if (127 != -128) {
                            switch ("" + "\n\t") {
                            case "":
                                fmt.Println(fi18)
                            default:
                                p2 = -9223372036854775808
                            }
                            v21 := []int{}
                            _ = v21
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            case -9223372036854775808:
                for fi22 := 0; fi22 < 17; fi22++ {
                    v23 := GenS2{F0: 2, F1: "\n\t", F2: (p0 - fi22)}
                    v24 := []int{4294967295, (^(p0)), (p0 << ((572) & 63))}
                    v25 := GenS3{F0: 1, F1: 1000, F2: 1}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            default:
                if true {
                    v26 := []int{(-(p0)), p1}
                    p2 += p2
                    p1 += p0
                    _ = v26
                }
            }
            p1 += p1
        }
        p1 -= (p1 << ((2) & 63))
        fmt.Println(p1, p0)
    }
    p0 -= (-(p0))
    var v27 uint = 402
    var v28 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return "hello"
}

func main() {
    var v1 uint8 = 1
    v2 := []int{901, 361, -1}
    var v3 int = -475
    v4 := make(map[uint32]string)
    fmt.Println(v3, v1)
    v1 -= 56
    fmt.Println(v3, v3)
    var v29 int32 = -329
    for fi30 := 0; fi30 < 13; fi30++ {
        fmt.Println(v3)
        v3 -= 42
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 uint16 = 1
            var v33 uint = 18446744073709551615
            if true {
                fmt.Println(fi30, v33, v32)
            }
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    switch "a" {
    case "":
        var v34 bool = (-32768 <= 42)
        _ = v34
    }
    if true {
        for fi35 := 0; fi35 < 10; fi35++ {
            var v36 bool = (!false)
            _ = fi35
            _ = v36
        }
        v37 := GenS1{F0: (v29 % ((genFn0()) | 1))}
        _ = v37
    } else {
        var v38 int16 = -1
        for fi39 := 0; fi39 < 20; fi39++ {
            delete(v4, 42)
            v29 += -2147483648
            v40 := GenS1{F0: (v29 << ((0) & 63))}
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    var v41 uint64 = 1
    if (v3 >= v3) {
        var v42 float32 = (float32(float64(1.5)) * ((((-2.25 + 1e10) / float32(((3.14159 + ((float64(1.5) * (0.0 * (0.1 - (3.14159 / 1.5)))) - ((0.1 * 0.0) * 1.5))) - 1.5))) * 0.1) / 1.0))
        fmt.Println(v29)
        var v43 float64 = -2.25
        _ = v42
        _ = v43
    } else {
        switch ("foo bar" + "") {
        case "hello":
            v1 += v1
        case "":
            v44 := GenS4{F0: 128, F1: 255}
            _ = v44
        case "a":
            var v45 uint = 473
            _ = v45
        default:
            v46 := v4[2]
            _ = v46
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            if (true && true) {
                v41 = 490
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v48 string = ("foo bar" + ("hello" + ""))
                _ = v48
            }
            _ = fi47
        }
    }
    switch "hello" {
    case "foo bar":
        v3 -= (-(v3))
    case "\n\t":
        var v49 int64 = 42
        _ = v49
    default:
        v1 += v1
    }
    fmt.Println(v29, v3, v41)
    for fi50 := 0; fi50 < 9; fi50++ {
        for fi51 := 0; fi51 < 3; fi51++ {
            if (false || ((-517 != 9223372036854775807) || true)) {
                fmt.Println(v1, v1, v41)
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            } else {
                v52 := GenS1{F0: int32(v3)}
                _ = v52
            }
            var v53 uint32 = 2147483648
            if true {
                if true {
                    switch v3 {
                    case 2147483647:
                        if true {
                            fmt.Println(v3, v3)
                            if ((v29 - 455) >= v29) {
                                _ = len(v4)
                            } else {
                                v29 = -539
                                switch (("" + (("foo bar" + "\n\t") + "")) + "foo bar") {
                                case "a":
                                    switch "" {
                                    case "hello":
                                        if true {
                                            v3 -= 0
                                            v1 += v1
                                        } else {
                                            var v54 uint32 = v53
                                            _ = v54
                                        }
                                    case "\n\t":
                                        if false {
                                            v41 += 526
                                        } else {
                                            var v55 int64 = -9223372036854775808
                                            v4[(^(v53))] = ("a" + "foo bar")
                                            fmt.Println(v1, v53, v55)
                                            _ = v55
                                        }
                                    default:
                                        if (432 >= 255) {
                                            v56 := GenS3{F0: int16(v29), F1: -32768, F2: v41}
                                            _ = v56
                                        } else {
                                            var v57 uint16 = 32767
                                            var v58 uint8 = v1
                                            v3 = (v3 >> ((16) & 63))
                                            switch fi51 {
                                            case -518:
                                                v41 = (v41 << ((1) & 63))
                                            default:
                                                v59 := v4[(v53 >> ((255) & 63))]
                                                var v60 float32 = 1e10
                                                _ = v59
                                                _ = v60
                                            }
                                            _ = v57
                                            _ = v58
                                        }
                                    }
                                case "\n\t":
                                    v29 = 42
                                }
                                v61 := GenS5{F0: 89, F1: "hello", F2: 255}
                                var v62 int = fi50
                                v29 = 2147483647
                                fmt.Println(v62, fi50)
                                if false {
                                    v1 += (v1 * (v1 % ((v1) | 1)))
                                    var v63 uint = 363
                                    if ((false && ((1 <= v53) || (1 <= int32(v53)))) || true) {
                                        var v64 float64 = float64((0.0 * 0.1))
                                        v65 := GenS5{F0: v1, F1: (("" + "hello") + ("\n\t" + "hello")), F2: (v1 - 127)}
                                        var v66 uint32 = uint32(v1)
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    } else {
                                        var v67 int32 = (v29 >> (((v63 << (((^(v63))) & 63))) & 63))
                                        fmt.Println(v53, fi51)
                                        switch (v3 / ((v3) | 1)) {
                                        case -1:
                                            v68 := GenS3{F0: 32767, F1: 32767, F2: (^(v41))}
                                            _ = v68
                                        default:
                                            var v69 float32 = (1e10 - ((-2.25 * 3.14159) + 0.1))
                                            _ = v69
                                        }
                                        _ = v67
                                    }
                                    var v70 int16 = int16(v41)
                                    v3 = (v3 - v3)
                                    _ = v63
                                    _ = v70
                                } else {
                                    v1 = 57
                                    var v71 int = fi51
                                    var v72 bool = (1.5 >= -2.25)
                                    _ = v71
                                    _ = v72
                                }
                                var v73 uint8 = v1
                                if false {
                                    v74 := GenS5{F0: 1, F1: "hello", F2: 129}
                                    v3 = (v3 ^ 4294967295)
                                    _ = v74
                                }
                                _ = v61
                                _ = v62
                                _ = v73
                            }
                            _ = len(v2)
                        }
                    case 128:
                        var v75 uint8 = 128
                        _ = v75
                    case 32768:
                        fmt.Println(v3, v53)
                    default:
                        v3 -= fi50
                    }
                    var v76 uint64 = (v41 | 1)
                    _ = v76
                } else {
                    var v77 uint8 = 42
                    var v78 uint16 = 255
                    _ = v77
                    _ = v78
                }
                var v79 int64 = 377
                _ = v79
            } else {
                v53 += 2
                if true {
                    v2 = append(v2, -1)
                    var v80 uint16 = 32767
                    if false {
                        v1 = (v1 * v1)
                    } else {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v3)
                        }
                        fmt.Println(v41)
                        if (true && (1.5 == 0.1)) {
                            var v81 int32 = (v29 / ((-563) | 1))
                            v29 -= v29
                            _ = v81
                        } else {
                            var v82 uint32 = 805
                            var v83 uint = 1
                            _ = v82
                            _ = v83
                        }
                    }
                    _ = v80
                }
                var v84 string = "a"
                _ = v84
            }
            v2 = append(v2, fi50)
            var v85 uint16 = 32767
            if ((v41 << ((1) & 63)) <= v41) {
                v3 -= v3
            }
            _ = len(v2)
            if (true && (!true)) {
                var v86 string = "hello"
                var v87 bool = (9223372036854775807 < -748)
                v88 := GenS5{F0: v1, F1: ("hello" + (v86 + (("a" + "\n\t") + v86))), F2: v1}
                _ = v86
                _ = v87
                _ = v88
            }
            _ = fi51
            _ = v53
            _ = v85
        }
        switch v3 {
        case 128:
            var v89 uint64 = 0
            _ = v89
        case -2147483648:
            v90 := GenS3{F0: 1, F1: -840, F2: (v41 ^ v41)}
            _ = v90
        case 127:
            fmt.Println(v41, v3, fi50)
        default:
            for fi91 := 0; fi91 < 12; fi91++ {
                v1 -= uint8(v41)
                if true {
                    var v92 int64 = -623
                    _ = v92
                } else {
                    v29 += v29
                }
                v93 := GenS2{F0: 265, F1: "\n\t", F2: 388}
                _ = fi91
                _ = v93
            }
        }
        _ = fi50
    }
    if ((v3 & int(v41)) < v3) {
        var v94 uint8 = (v1 << ((56) & 63))
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v41
}
