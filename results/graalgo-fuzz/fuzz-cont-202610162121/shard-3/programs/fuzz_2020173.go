package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 bool; F1 int16; F2 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 int, p1 bool, p2 int16) uint8 {
    var v5 uint64 = 42
    v6 := GenS1{F0: 2}
    var v7 string = "a"
    v8 := GenS1{F0: 0}
    v9 := []int{p0, p0, p0, (p0 % ((p0) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 uint16) uint8 {
    p0 = (p0 >> ((4294967295) & 63))
    p0 = p0
    switch 42 {
    case 32768:
        p0 = p0
    case 128:
        var v10 uint = 65535
        _ = v10
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint32 = 2147483647
            v13 := []int{0}
            var v14 uint64 = 1
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    }
    p0 += 576
    v15 := []int{42, -542}
    _ = p0
    _ = v15
    return 255
}

func main() {
    var v1 uint64 = 42
    var v2 string = ""
    v3 := make(map[string]int)
    var v4 uint32 = 2147483648
    switch -790 {
    case 20:
        var v16 int64 = 9223372036854775807
        _ = v16
    case 128:
        v17 := []int{-287, 9223372036854775807}
        _ = v17
    case 42:
        var v18 uint32 = (v4 << ((255) & 63))
        _ = v18
    default:
        v3[v2] = -9223372036854775808
    }
    switch -2147483648 {
    case 4294967295:
        v2 = "a"
    case 2147483647:
        switch 0 {
        case 127:
            for fi19 := 0; fi19 < 6; fi19++ {
                switch (-(fi19)) {
                case -216:
                    var v20 uint64 = (v1 * 42)
                    _ = v20
                case 32767:
                    switch (fi19 - fi19) {
                    case 42:
                        for fi21 := 0; fi21 < 14; fi21++ {
                            var v22 string = v2
                            v2 = v22
                            if false {
                                switch v2 {
                                case "hello":
                                    v22 = "a"
                                case "":
                                    if (false || (true && true)) {
                                        v3["foo bar"] = 1
                                        v23 := GenS1{F0: 65535}
                                        v4 += v4
                                        _ = v23
                                    } else {
                                        var v24 string = ("hello" + v22)
                                        _ = len(v3)
                                        v22 = (((v2 + v22) + v22) + v2)
                                        fmt.Println(fi19)
                                        v25 := GenS2{F0: (v4 - v4), F1: (^(v1))}
                                        v1 = v1
                                        _ = v24
                                        _ = v25
                                    }
                                }
                            } else {
                                v22 = v2
                            }
                            _ = fi21
                            _ = v22
                        }
                    }
                case 4294967295:
                    v1 += 2
                default:
                    var v26 int32 = 530
                    _ = v26
                }
                switch "a" {
                case "":
                    if (!(((1e10 > (1e10 - (float64(1e10) + 1e10))) && (255 != 4294967295)) && (false || false))) {
                        v27 := GenS2{F0: uint32(v1), F1: (v1 * v1)}
                        if false {
                            switch "\n\t" {
                            case "a":
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    switch (^(fi28)) {
                                    case -753:
                                        fmt.Println(fi19, v4, fi19)
                                    default:
                                        v29 := v3[(("foo bar" + v2) + v2)]
                                        _ = v29
                                    }
                                    _ = fi28
                                }
                            case "":
                                for fi30 := 0; fi30 < 18; fi30++ {
                                    v31 := []int{fi30}
                                    if (612 < 32767) {
                                        var v32 uint64 = (v1 % (((^(v1))) | 1))
                                        v33 := GenS3{F0: 255, F1: 1}
                                        _ = v32
                                        _ = v33
                                    }
                                    fmt.Println(v4)
                                    _ = fi30
                                    _ = v31
                                }
                            }
                        } else {
                            v2 = v2
                        }
                        _ = v27
                    } else {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            var v35 uint8 = 42
                            _ = fi34
                            _ = v35
                        }
                    }
                case "foo bar":
                    v1 -= v1
                }
                if (false || false) {
                    var v36 uint64 = v1
                    var v37 int16 = 1
                    var v38 int64 = -9223372036854775808
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi19
            }
        case -993:
            var v39 int32 = 2147483647
            _ = v39
        case -2147483648:
            for fi40 := 0; fi40 < 9; fi40++ {
                switch fi40 {
                case 9223372036854775807:
                    v41 := GenS4{F0: true, F1: 1}
                    _ = v41
                default:
                    if (((0.1 - 1e10) == 1e10) || false) {
                        if (65535 < 0) {
                            v42 := []int{fi40}
                            switch "foo bar" {
                            case "":
                                if false {
                                    v43 := GenS4{F0: (!(32768 == (-(fi40)))), F1: 51}
                                    v4 = v4
                                    v44 := GenS3{F0: 255, F1: 2147483647}
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 int8 = -79
                                    for fi46 := 0; fi46 < 12; fi46++ {
                                        delete(v3, v2)
                                        _ = fi46
                                    }
                                    v1 += v1
                                    _ = v45
                                }
                            case "a":
                                if (int32(v1) > -309) {
                                    fmt.Println(v1, v1, v1)
                                }
                            case "hello":
                                fmt.Println(v4, v4, fi40)
                            }
                            var v47 uint8 = 1
                            _ = v42
                            _ = v47
                        }
                    }
                }
                switch v2 {
                case "a":
                    for fi48 := 0; fi48 < 10; fi48++ {
                        var v49 uint64 = v1
                        var v50 int16 = 42
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                case "foo bar":
                    var v51 string = "a"
                    _ = v51
                case "hello":
                    fmt.Println(v1, v4)
                }
                fmt.Println(v2, v2)
                _ = fi40
            }
        default:
            if true {
                fmt.Println(v2, v4)
                var v52 uint = 443
                for fi53 := 0; fi53 < 12; fi53++ {
                    var v54 uint16 = 65535
                    v52 -= v52
                    for fi55 := 0; fi55 < 6; fi55++ {
                        v52 += v52
                        _ = fi55
                    }
                    _ = fi53
                    _ = v54
                }
                _ = v52
            }
        }
    default:
        var v56 uint8 = 87
        _ = v56
    }
    fmt.Println(v4, v1, v4)
    var v57 int16 = 1000
    if (false || false) {
        for fi58 := 0; fi58 < 13; fi58++ {
            v59 := GenS5{F0: true, F1: -1, F2: -522}
            _ = fi58
            _ = v59
        }
    } else {
        v60 := GenS5{F0: (79 == 127), F1: v57, F2: 186}
        v1 = v1
        _ = v60
    }
    v61 := []int{-2147483648, 595, -1, -488}
    switch 127 {
    case 2147483647:
        var v62 uint = 9
        _ = v62
    case 32767:
        v63 := GenS6{F0: -2147483648}
        _ = v63
    case -2147483648:
        if (float32((-2.25 * 0.0)) <= (3.14159 * float32((0.1 + -2.25)))) {
            if true {
                v64 := GenS1{F0: 1}
                if (((v4 & (v4 - (v4 >> ((18446744073709551615) & 63)))) <= (v4 - (v4 + v4))) && (42 > 9223372036854775807)) {
                    v57 -= (^(v57))
                    if (!(-1 != -1)) {
                        var v65 uint16 = 704
                        var v66 uint8 = 0
                        _ = v65
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 7; fi67++ {
                        v57 = v57
                        fmt.Println(v2, v4, v1)
                        _ = fi67
                    }
                } else {
                    var v68 uint32 = v4
                    v61 = append(v61, 127)
                    switch v2 {
                    case "hello":
                        if (0 < 666) {
                            v57 += 224
                        }
                    default:
                        v69 := GenS3{F0: genFn0(32768, (v1 <= (v1 >> ((314) & 63))), (v57 / ((0) | 1))), F1: 0}
                        _ = v69
                    }
                    _ = v68
                }
                if true {
                    var v70 int8 = 127
                    for fi71 := 0; fi71 < 14; fi71++ {
                        for fi72 := 0; fi72 < 18; fi72++ {
                            var v73 uint64 = v1
                            if ((v70 >= v70) && true) {
                                delete(v3, v2)
                            }
                            var v74 uint = 0
                            _ = fi72
                            _ = v73
                            _ = v74
                        }
                        switch fi71 {
                        case 127:
                            v4 = (^(v4))
                        }
                        var v75 uint16 = uint16(v1)
                        _ = fi71
                        _ = v75
                    }
                    switch 42 {
                    case 128:
                        fmt.Println(v70, v1, v70)
                    case 1:
                        v4 += 1
                    default:
                        var v76 uint16 = 110
                        _ = v76
                    }
                    _ = v70
                }
                _ = v64
            } else {
                if len(v61) > 0 {
                    v61[0] = 9223372036854775807
                }
                v57 += -1
                if false {
                    _ = len(v61)
                    for fi77 := 0; fi77 < 17; fi77++ {
                        var v78 uint = 2
                        v57 -= (v57 ^ v57)
                        _ = fi77
                        _ = v78
                    }
                    switch 784 {
                    case -1:
                        var v79 int32 = -545
                        _ = v79
                    case -2147483648:
                        var v80 uint16 = 255
                        _ = v80
                    }
                }
            }
            v2 = v2
        } else {
            switch 32768 {
            case 32767:
                for fi81 := 0; fi81 < 14; fi81++ {
                    for fi82 := 0; fi82 < 16; fi82++ {
                        fmt.Println(fi81, v57)
                        _ = fi82
                    }
                    _ = fi81
                }
            case 623:
                switch v2 {
                case "foo bar":
                    fmt.Println(v57, v4, v2)
                case "\n\t":
                    switch v2 {
                    case "\n\t":
                        v1 = (v1 + (v1 | (v1 - (v1 & v1))))
                    case "hello":
                        for fi83 := 0; fi83 < 3; fi83++ {
                            switch (fi83 % ((fi83) | 1)) {
                            case 128:
                                var v84 uint8 = 1
                                _ = v84
                            case -222:
                                for fi85 := 0; fi85 < 20; fi85++ {
                                    switch v2 {
                                    case "":
                                        v57 -= v57
                                    }
                                    v1 = 2
                                    if true {
                                        if true {
                                            v2 = "foo bar"
                                            var v86 uint32 = 2147483647
                                            _ = v86
                                        }
                                        delete(v3, ("\n\t" + "hello"))
                                        var v87 uint16 = 32767
                                        var v88 uint32 = v4
                                        _ = v87
                                        _ = v88
                                    }
                                    _ = fi85
                                }
                            case 4294967295:
                                var v89 uint32 = v4
                                _ = v89
                            }
                            _ = fi83
                        }
                    }
                case "":
                    for fi90 := 0; fi90 < 4; fi90++ {
                        _ = len(v61)
                        _ = fi90
                    }
                }
            }
            var v91 int64 = -477
            var v92 float64 = 1.5
            _ = v91
            _ = v92
        }
    default:
        v1 += (v1 >> ((886) & 63))
    }
    for fi93 := 0; fi93 < 16; fi93++ {
        v4 = v4
        fmt.Println(fi93, v2)
        _ = fi93
    }
    v2 = v2
    var v94 uint8 = 1
    if true {
        fmt.Println(v94, v4, v1)
    }
    var v95 float32 = 0.0
    var v96 float32 = float32(3.14159)
    v97 := GenS4{F0: true, F1: 1}
    fmt.Println(v1, v95)
    v3["foo bar"] = -632
    v98 := GenS5{F0: true, F1: v57, F2: 1}
    var v99 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v61
    _ = v94
    _ = v95
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
