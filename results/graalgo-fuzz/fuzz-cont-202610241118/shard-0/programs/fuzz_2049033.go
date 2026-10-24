package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 int32; F2 bool }
type GenS7 struct { F0 int8; F1 int32; F2 int8 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 int32; F2 int32 }
type GenS11 struct { F0 uint8; F1 uint32; F2 int32 }
func genFn0(p0 uint16) uint16 {
    var v5 uint16 = p0
    var v6 uint = 1
    if (false || (true && false)) {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 68, F1: 0}
            v9 := GenS1{F0: v6, F1: 2}
            for fi10 := 0; fi10 < 15; fi10++ {
                p0 = p0
                var v11 uint = 0
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint64 = 18446744073709551615
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int64 = 0
            if false {
                var v16 uint32 = 2
                _ = v16
            } else {
                var v17 int = fi14
                switch "\n\t" {
                case "":
                    var v18 float32 = float32(1.5)
                    _ = v18
                case "a":
                    v19 := GenS2{F0: 32767, F1: 144, F2: 127}
                    _ = v19
                }
                _ = v17
            }
            fmt.Println(v15)
            _ = fi14
            _ = v15
        }
    }
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 255
    var v3 float64 = 1.5
    var v4 int = 32767
    var v20 uint64 = 18446744073709551615
    v21 := GenS1{F0: 799, F1: 839}
    for fi22 := 0; fi22 < 16; fi22++ {
        v3 -= float64(0.0)
        if (0.1 > -2.25) {
            var v23 bool = true
            _ = v23
        }
        v24 := GenS3{F0: 2, F1: 0}
        _ = fi22
        _ = v24
    }
    switch "hello" {
    case "hello":
        var v25 uint16 = 42
        _ = v25
    default:
        var v26 float32 = (float32((float64(0.1) / v3)) / 0.1)
        _ = v26
    }
    if true {
        var v27 int = (v4 % ((v4) | 1))
        _ = v27
    } else {
        v28 := []int{v4, v4, (v4 + v4)}
        if (true || true) {
            var v29 int64 = 42
            _ = v29
        }
        var v30 uint64 = uint64(v2)
        _ = v28
        _ = v30
    }
    var v31 string = "hello"
    fmt.Println(v1)
    var v32 int64 = 0
    v33 := []int{(v4 / ((v4) | 1)), 128, -593, v4}
    if ((^(v4)) != v4) {
        v32 = v32
        var v34 uint = 255
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 uint8 = v2
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    if false {
        var v37 uint32 = 2147483648
        _ = len(v33)
        fmt.Println(v3, v2)
        _ = v37
    } else {
        var v38 float64 = v3
        _ = v38
    }
    v39 := GenS4{F0: 2147483647}
    v40 := GenS5{F0: -934}
    for fi41 := 0; fi41 < 12; fi41++ {
        v2 = (v2 >> ((517) & 63))
        v32 = v32
        switch "" {
        case "a":
            v42 := GenS6{F0: int32(v1), F1: 1, F2: (!((true || false) && true))}
            _ = v42
        case "\n\t":
            var v43 bool = false
            _ = v43
        case "foo bar":
            if (!true) {
                switch (fi41 >> ((645) & 63)) {
                case -240:
                    for fi44 := 0; fi44 < 13; fi44++ {
                        if (764 < 65535) {
                            if true {
                                v31 = v31
                                var v45 uint64 = (v20 % ((v20) | 1))
                                _ = v45
                            }
                            if (true && true) {
                                var v46 int32 = 1
                                _ = v46
                            } else {
                                var v47 int64 = (v32 * int64(v2))
                                var v48 uint32 = 4294967295
                                if len(v33) >= 2 {
                                    v33 = v33[1:len(v33)]
                                }
                                if (false || false) {
                                    var v49 uint16 = 32767
                                    _ = v49
                                }
                                v31 = ((v31 + v31) + v31)
                                _ = v47
                                _ = v48
                            }
                        } else {
                            if true {
                                _ = len(v33)
                                _ = len(v33)
                                if len(v33) > 0 {
                                    v33[0] = 771
                                }
                                if len(v33) >= 2 {
                                    v33 = v33[1:len(v33)]
                                }
                                var v50 uint = 4294967295
                                _ = v50
                            }
                        }
                        _ = fi44
                    }
                case 642:
                    fmt.Println(v31, v31, v3)
                case 809:
                    if ((((0.0 / 1e10) - 1e10) > 1e10) && true) {
                        fmt.Println(v20, v20, v3)
                        switch v31 {
                        case "foo bar":
                            switch v31 {
                            case "a":
                                if ((v4 / ((128) | 1)) <= (v4 - 32768)) {
                                    var v51 uint32 = 0
                                    _ = v51
                                }
                            case "foo bar":
                                v2 += 60
                            case "":
                                v52 := GenS5{F0: 1}
                                _ = v52
                            }
                        case "hello":
                            for fi53 := 0; fi53 < 14; fi53++ {
                                if true {
                                    var v54 float32 = (1e10 / float32((1.5 / v3)))
                                    if false {
                                        var v55 uint8 = (v2 + v2)
                                        if false {
                                            switch (fi53 >> ((768) & 63)) {
                                            case 664:
                                                v55 += 102
                                            case 9223372036854775807:
                                                v55 += 212
                                            case -2147483648:
                                                if (!true) {
                                                    v56 := GenS7{F0: -37, F1: 0, F2: -69}
                                                    _ = v56
                                                } else {
                                                    if ((int16(v2) <= 42) || false) {
                                                        switch 623 {
                                                        case -418:
                                                            v1 -= v1
                                                        case -1:
                                                            var v57 uint8 = (v55 << ((uint(v55)) & 63))
                                                            _ = v57
                                                        case -9223372036854775808:
                                                            if len(v33) > 0 {
                                                                v33[0] = -611
                                                            }
                                                        default:
                                                            v54 += 1.5
                                                        }
                                                        var v58 uint64 = 0
                                                        switch v31 {
                                                        case "\n\t":
                                                            var v59 int8 = 1
                                                            _ = v59
                                                        case "hello":
                                                            _ = len(v33)
                                                        case "":
                                                            v60 := GenS3{F0: 0, F1: 762}
                                                            _ = v60
                                                        }
                                                        if (uint32(v2) >= 964) {
                                                            _ = len(v33)
                                                        }
                                                        v61 := GenS8{F0: true, F1: 42}
                                                        _ = v58
                                                        _ = v61
                                                    }
                                                }
                                            default:
                                                if (true && (847 > 2)) {
                                                    if (-32768 < v1) {
                                                        v62 := GenS9{F0: (1 <= uint64(v2))}
                                                        _ = v62
                                                    } else {
                                                        if true {
                                                            v54 -= ((float32(1.5) - float32(v3)) * (v54 - (v54 * (v54 * v54))))
                                                            var v63 float32 = v54
                                                            var v64 float32 = (float32((v3 * v3)) - ((1.5 + ((float32(v3) + (((v63 / ((v63 - (v63 * 0.1)) / v54)) + float32(v3)) + v63)) * v54)) / float32(1.5)))
                                                            v32 -= 15
                                                            v1 += (v1 << ((0) & 63))
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            v33 = append(v33, -175)
                                                            v33 = append(v33, (^(fi53)))
                                                            v65 := GenS6{F0: 42, F1: 359, F2: true}
                                                            _ = v65
                                                        }
                                                        v1 += (v1 >> ((0) & 63))
                                                    }
                                                    v66 := GenS4{F0: 2147483647}
                                                    v32 += v32
                                                    _ = v66
                                                } else {
                                                    v67 := GenS10{F0: (v1 >> ((53) & 63)), F1: 1, F2: 2147483647}
                                                    if len(v33) >= 2 {
                                                        v33 = v33[1:len(v33)]
                                                    }
                                                    v55 = v55
                                                    v2 = 42
                                                    v4 += fi53
                                                    _ = v67
                                                }
                                            }
                                        }
                                        _ = len(v33)
                                        _ = v55
                                    } else {
                                        fmt.Println(fi53, v54)
                                        var v68 float32 = (v54 - 3.14159)
                                        fmt.Println(v31, v68, fi41)
                                        _ = v68
                                    }
                                    v1 = v1
                                    var v69 uint16 = 255
                                    _ = v54
                                    _ = v69
                                }
                                _ = fi53
                            }
                        default:
                            if (100 >= 0) {
                                for fi70 := 0; fi70 < 2; fi70++ {
                                    var v71 int = (v4 >> ((uint(v20)) & 63))
                                    var v72 float32 = -2.25
                                    v33 = append(v33, 4294967295)
                                    v4 = 32768
                                    _ = fi70
                                    _ = v71
                                    _ = v72
                                }
                                fmt.Println(v31, v1, v31)
                            } else {
                                if len(v33) > 0 {
                                    v33[0] = fi41
                                }
                                if false {
                                    if len(v33) >= 2 {
                                        v33 = v33[1:len(v33)]
                                    }
                                    if len(v33) > 0 {
                                        v33[0] = 127
                                    }
                                } else {
                                    v2 -= v2
                                    var v73 uint64 = 0
                                    _ = v73
                                }
                                v74 := GenS9{F0: true}
                                _ = v74
                            }
                        }
                    }
                }
            }
        default:
            var v75 bool = (0.0 <= v3)
            _ = v75
        }
        _ = fi41
    }
    for fi76 := 0; fi76 < 12; fi76++ {
        var v77 uint16 = 255
        switch (v4 / ((v4) | 1)) {
        case -702:
            if (-386 >= v1) {
                var v78 int8 = 42
                var v79 uint16 = v77
                switch (-(v4)) {
                case 32767:
                    fmt.Println(v31)
                case 128:
                    var v80 int8 = v78
                    _ = v80
                default:
                    for fi81 := 0; fi81 < 11; fi81++ {
                        var v82 uint64 = uint64(v2)
                        v2 = (v2 & 128)
                        _ = fi81
                        _ = v82
                    }
                }
                _ = v78
                _ = v79
            } else {
                for fi83 := 0; fi83 < 10; fi83++ {
                    switch -2147483648 {
                    case 4294967295:
                        v1 += (v1 + 42)
                    case 0:
                        if true {
                            v3 = (((1.5 - v3) / ((v3 + ((v3 * float64(1e10)) - v3)) * (float64(0.0) + v3))) * 3.14159)
                        }
                    default:
                        v2 += (v2 + 128)
                    }
                    _ = fi83
                }
            }
        case -2147483648:
            var v84 string = v31
            _ = v84
        case 757:
            var v85 int16 = -527
            _ = v85
        default:
            v86 := GenS11{F0: 27, F1: 4294967295, F2: 278}
            _ = v86
        }
        _ = fi76
        _ = v77
    }
    var v87 int16 = (^(v1))
    v88 := GenS7{F0: 127, F1: 1, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v87
    _ = v88
}
