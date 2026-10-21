package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint64; F1 bool; F2 int8 }
type GenS5 struct { F0 uint; F1 int8; F2 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 int32; F2 uint }
func genFn0(p0 int32, p1 int16, p2 int8) int8 {
    fmt.Println(p2)
    v5 := []int{32768, 42, -1}
    p1 = -32768
    if (-500 < 926) {
        for fi6 := 0; fi6 < 16; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: (p2 >> ((2) & 63))}
                var v9 int8 = 29
                var v10 uint = 65535
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            v5 = append(v5, -1)
            var v11 int64 = 1
            _ = fi6
            _ = v11
        }
        if len(v5) > 0 {
            v5[0] = -893
        }
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            if (733 <= 42) {
                v13 := GenS2{F0: 18446744073709551615}
                var v14 string = ("hello" + "")
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        switch 127 {
        case 79:
            switch "\n\t" {
            case "a":
                p1 += (p1 - (p1 * (p1 - 32767)))
            case "foo bar":
                var v15 uint16 = 0
                _ = v15
            case "hello":
                p2 += p2
            }
        case -881:
            p1 = (p1 << ((883) & 63))
        }
    }
    if (!((3.14159 + float32((0.0 - 0.1))) < 1.5)) {
        var v16 uint = 255
        var v17 bool = false
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 int32, p1 int, p2 uint32) int {
    fmt.Println(p2, p2, p0)
    switch (p1 << ((304) & 63)) {
    case -463:
        fmt.Println(p1, p0, p2)
    case 32768:
        if false {
            v18 := []int{p1, p1, p1}
            v19 := GenS1{F0: -2}
            if false {
                var v20 float64 = float64(1e10)
                _ = v20
            }
            _ = v18
            _ = v19
        }
    }
    fmt.Println(p2, p2, p0)
    if true {
        v21 := GenS1{F0: -128}
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := []int{1}
            switch "foo bar" {
            case "a":
                p1 += (fi22 / ((-2147483648) | 1))
            case "\n\t":
                switch "hello" {
                case "hello":
                    var v24 int = 127
                    _ = v24
                case "a":
                    if (!false) {
                        v25 := GenS3{F0: 32767, F1: uint16(p2)}
                        p1 -= p1
                        _ = v25
                    }
                }
            }
            p1 += (p1 << ((18446744073709551615) & 63))
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        var v26 uint8 = 128
        v27 := []int{p1, (p1 % ((128) | 1)), (p1 * (p1 / ((-9223372036854775808) | 1))), (p1 & (p1 - 32768))}
        if (p1 <= (^(p1))) {
            fmt.Println(p0, p0)
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            v27 = append(v27, p1)
        } else {
            v28 := GenS1{F0: 42}
            v27 = append(v27, (-(p1)))
            _ = v28
        }
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 ^ 9223372036854775807)
}

func main() {
    v1 := []string{"", "a", "hello", ("foo bar" + "")}
    var v2 uint8 = 128
    v3 := []uint32{0, uint32(v2), 4294967295}
    v4 := make(map[uint32]bool)
    fmt.Println(v2)
    fmt.Println(v2)
    for fi29 := 0; fi29 < 12; fi29++ {
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int = fi29
            v31 = fi29
            _ = fi30
            _ = v31
        }
        if (30 == 18446744073709551615) {
            fmt.Println(fi29)
            v2 = (v2 % (((v2 * (v2 % ((1) | 1)))) | 1))
        } else {
            v2 += (v2 & (v2 / (((v2 % ((v2) | 1))) | 1)))
            v2 = 229
        }
        var v32 uint = 0
        _ = fi29
        _ = v32
    }
    v2 = v2
    if (42 <= int(v2)) {
        if true {
            for fi33 := 0; fi33 < 18; fi33++ {
                fmt.Println(fi33, v2)
                _ = fi33
            }
        }
        var v34 uint = 4294967295
        _ = v34
    } else {
        for fi35 := 0; fi35 < 13; fi35++ {
            if true {
                if (-2.25 == 0.0) {
                    var v36 string = ("hello" + (("\n\t" + ("\n\t" + ("" + "a"))) + ""))
                    v37 := GenS4{F0: 1, F1: false, F2: 1}
                    _ = v36
                    _ = v37
                }
                v2 = (v2 >> ((42) & 63))
                if true {
                    v2 = v2
                    for fi38 := 0; fi38 < 18; fi38++ {
                        switch 265 {
                        case 0:
                            v4[42] = (!(530 < -32768))
                        default:
                            if (2147483647 < -2147483648) {
                                _ = len(v3)
                                fmt.Println(v2, fi35, fi35)
                            }
                        }
                        v2 = (v2 * v2)
                        if true {
                            v39 := GenS3{F0: -32768, F1: 0}
                            v2 += v2
                            _ = v39
                        } else {
                            v40 := GenS4{F0: 18446744073709551615, F1: (3.14159 > 1.5), F2: -128}
                            v2 = (v2 + v2)
                            fmt.Println(fi38)
                            v41 := GenS2{F0: 4294967295}
                            v2 = (v2 - (v2 >> ((194) & 63)))
                            v42 := v4[65535]
                            var v43 uint = 284
                            _ = v40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        _ = fi38
                    }
                } else {
                    fmt.Println(v2, fi35)
                }
            }
            for fi44 := 0; fi44 < 16; fi44++ {
                v2 = (v2 << ((1) & 63))
                var v45 uint8 = v2
                v46 := GenS5{F0: 4294967295, F1: -81, F2: -42}
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = fi35
        }
        fmt.Println(v2)
        v2 -= v2
    }
    switch 127 {
    case -9223372036854775808:
        for fi47 := 0; fi47 < 8; fi47++ {
            v2 = (^(v2))
            _ = fi47
        }
    case -2147483648:
        if ((v2 == 121) && false) {
            _ = len(v4)
            v48 := GenS2{F0: 718}
            v2 -= 175
            _ = v48
        }
    case -717:
        if true {
            var v49 int16 = 32767
            _ = v49
        } else {
            fmt.Println(v2)
            for fi50 := 0; fi50 < 3; fi50++ {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                if ((2 >= 354) && false) {
                    if (439 <= 4294967295) {
                        var v51 uint8 = 96
                        v52 := GenS6{F0: 42}
                        v53 := GenS1{F0: -1}
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    v2 = (v2 + (v2 >> ((505) & 63)))
                } else {
                    var v54 int64 = 703
                    _ = v54
                }
                _ = fi50
            }
        }
    }
    var v55 int16 = 87
    for fi56 := 0; fi56 < 4; fi56++ {
        var v57 int64 = 0
        switch genFn1(-1, (fi56 ^ -2147483648), 4294967295) {
        case 42:
            for fi58 := 0; fi58 < 19; fi58++ {
                if (786 <= 0) {
                    if true {
                        var v59 bool = (v55 < v55)
                        switch 32768 {
                        case -1:
                            var v60 float32 = (-2.25 * 3.14159)
                            _ = v60
                        case 32767:
                            v1 = append(v1, "")
                        case -172:
                            if (3.14159 != float32((0.1 / 0.1))) {
                                switch (fi56 ^ genFn1(0, fi56, 2147483648)) {
                                case 1:
                                    v57 = 9223372036854775807
                                default:
                                    var v61 float64 = 1e10
                                    _ = v61
                                }
                                v1 = append(v1, "\n\t")
                            }
                        default:
                            if true {
                                v2 += v2
                                v57 += v57
                            }
                        }
                        _ = v59
                    }
                } else {
                    var v62 int16 = (v55 >> ((150) & 63))
                    _ = v62
                }
                var v63 float64 = (0.1 * (0.1 * ((1.5 / 1.0) * 3.14159)))
                v64 := GenS7{F0: true, F1: 1, F2: 65535}
                _ = fi58
                _ = v63
                _ = v64
            }
        case 32768:
            for fi65 := 0; fi65 < 10; fi65++ {
                switch (^(fi56)) {
                case 9223372036854775807:
                    v57 += (^(v57))
                case -9223372036854775808:
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                }
                _ = fi65
            }
        case 1:
            if true {
                switch fi56 {
                case 4294967295:
                    for fi66 := 0; fi66 < 4; fi66++ {
                        var v67 int = fi56
                        var v68 bool = (2 == 0)
                        var v69 int = fi66
                        _ = fi66
                        _ = v67
                        _ = v68
                        _ = v69
                    }
                case 32767:
                    v55 += -638
                case 128:
                    if true {
                        var v70 uint16 = uint16(v2)
                        _ = v70
                    }
                }
                if (2147483647 >= 874) {
                    v3 = append(v3, 1)
                    for fi71 := 0; fi71 < 6; fi71++ {
                        if false {
                            var v72 uint32 = 562
                            v2 -= (v2 / ((255) | 1))
                            var v73 int64 = (v57 >> ((255) & 63))
                            _ = v72
                            _ = v73
                        }
                        v74 := GenS2{F0: 0}
                        if ((fi56 + (fi71 / (((fi56 ^ 42)) | 1))) <= fi71) {
                            if ((false || false) && true) {
                                var v75 bool = true
                                if v75 {
                                    fmt.Println(v55, v57, v57)
                                } else {
                                    if v75 {
                                        var v76 uint8 = v2
                                        v57 -= (v57 % (((v57 << ((767) & 63))) | 1))
                                        _ = v76
                                    } else {
                                        fmt.Println(fi56)
                                        v2 = 0
                                        v75 = (false && v75)
                                    }
                                    v75 = (v2 < 1)
                                    var v77 float64 = (1.5 - -2.25)
                                    _ = v77
                                }
                                var v78 float32 = -2.25
                                _ = v75
                                _ = v78
                            } else {
                                fmt.Println(v2)
                                _ = len(v1)
                            }
                            switch ("" + ("" + "foo bar")) {
                            case "hello":
                                fmt.Println(v2, fi71, v2)
                            case "\n\t":
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                            }
                            v79 := v4[309]
                            _ = v79
                        } else {
                            _ = len(v4)
                            var v80 int64 = int64(v55)
                            _ = v80
                        }
                        _ = fi71
                        _ = v74
                    }
                } else {
                    fmt.Println(v2, v2)
                    v55 = (v55 - (-(v55)))
                }
                for fi81 := 0; fi81 < 3; fi81++ {
                    var v82 string = "hello"
                    _ = fi81
                    _ = v82
                }
            } else {
                for fi83 := 0; fi83 < 4; fi83++ {
                    fmt.Println(fi83, v57)
                    v4[246] = (!false)
                    var v84 float64 = 1e10
                    _ = fi83
                    _ = v84
                }
                if true {
                    _ = len(v1)
                    v55 = (v55 >> ((65535) & 63))
                }
            }
        }
        var v85 int8 = genFn0(847, (v55 - v55), 1)
        _ = fi56
        _ = v57
        _ = v85
    }
    v4[582] = true
    var v86 string = "a"
    var v87 int16 = v55
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v86
    _ = v87
}
