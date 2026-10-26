package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int; F1 int16; F2 int32 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8; F1 bool; F2 uint32 }
func genFn0() uint {
    v7 := GenS1{F0: -1}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 string = "\n\t"
        switch (fi8 - (fi8 & fi8)) {
        case 2147483647:
            fmt.Println(v9, v9)
        case 157:
            v10 := []int{fi8, 28, (fi8 * (fi8 - fi8)), 128}
            _ = v10
        case 128:
            fmt.Println(fi8)
        default:
            v9 = v9
        }
        _ = fi8
        _ = v9
    }
    var v11 uint = 2
    if false {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float64 = 1e10
            v14 := GenS2{F0: 2147483648}
            v15 := []int{fi12, fi12}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    switch "hello" {
    case "a":
        fmt.Println(v11)
    case "hello":
        v11 = (v11 * (v11 % ((2) | 1)))
    case "":
        v11 -= (v11 * (v11 / ((v11) | 1)))
    default:
        v11 -= v11
    }
    _ = v11
    _ = v7
    return (v11 | v11)
}
func genFn1(p0 uint16) bool {
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS2{F0: 0}
        p0 -= (p0 - (p0 + 65535))
        var v18 float32 = 3.14159
        _ = fi16
        _ = v17
        _ = v18
    }
    p0 = 2
    p0 = p0
    _ = p0
    return false
}

func main() {
    var v1 uint32 = 365
    v2 := []uint32{v1, (^(v1))}
    var v3 string = ("\n\t" + (("hello" + "hello") + ("\n\t" + "a")))
    var v4 int = -9223372036854775808
    v5 := make(map[string]int)
    var v6 float32 = 1e10
    switch -933 {
    case 0:
        v4 += v4
    default:
        var v19 int16 = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS3{F0: v3}
        _ = fi20
        _ = v21
    }
    v22 := GenS4{F0: 724, F1: -1}
    v23 := GenS2{F0: v1}
    _ = len(v2)
    var v24 uint64 = 409
    v25 := GenS5{F0: 549, F1: 1, F2: (v1 * (v1 ^ (^(v1))))}
    switch (v4 & (^(v4))) {
    case 718:
        var v26 uint = 18446744073709551615
        _ = v26
    default:
        switch v4 {
        case 128:
            switch v4 {
            case 9223372036854775807:
                v6 -= v6
            case -73:
                switch v3 {
                case "a":
                    v27 := GenS6{F0: 32767, F1: 964, F2: -564}
                    _ = v27
                case "hello":
                    if (0 == 42) {
                        switch (v3 + v3) {
                        case "":
                            v4 = (v4 - (v4 | (-(v4))))
                        case "hello":
                            var v28 float32 = ((v6 - float32(float64((v6 * (v6 - v6))))) - float32(0.0))
                            _ = v28
                        default:
                            fmt.Println(v3, v6)
                        }
                        var v29 int8 = 1
                        for fi30 := 0; fi30 < 6; fi30++ {
                            var v31 uint64 = 450
                            _ = fi30
                            _ = v31
                        }
                        _ = v29
                    }
                case "":
                    delete(v5, v3)
                default:
                    if true {
                        for fi32 := 0; fi32 < 16; fi32++ {
                            var v33 uint8 = 36
                            _ = fi32
                            _ = v33
                        }
                        if (-1 >= 103) {
                            v3 = (v3 + ("\n\t" + v3))
                            var v34 int32 = -904
                            for fi35 := 0; fi35 < 8; fi35++ {
                                var v36 uint = 645
                                var v37 string = ("\n\t" + (("a" + v3) + v3))
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            _ = v34
                        }
                    } else {
                        switch v3 {
                        case "foo bar":
                            var v38 uint16 = 78
                            _ = v38
                        default:
                            if true {
                                v39 := GenS4{F0: v24, F1: 0}
                                if len(v2) > 0 {
                                    v2[0] = (v1 % ((835) | 1))
                                }
                                v40 := GenS7{F0: -992, F1: (v24 | 28)}
                                _ = v39
                                _ = v40
                            }
                        }
                        var v41 uint = 4294967295
                        v42 := GenS7{F0: -9223372036854775808, F1: v24}
                        _ = v41
                        _ = v42
                    }
                }
            case -2147483648:
                if true {
                    for fi43 := 0; fi43 < 14; fi43++ {
                        v44 := GenS8{F0: 127, F1: (v24 | 2)}
                        _ = fi43
                        _ = v44
                    }
                    var v45 int16 = 42
                    _ = v45
                }
            }
        }
    }
    var v46 uint = 1
    var v47 bool = ((v3 + v3) > v3)
    var v48 int16 = int16(v4)
    if v47 {
        fmt.Println(v47, v3, v47)
    } else {
        fmt.Println(v47, v24)
        switch "hello" {
        case "a":
            var v49 int32 = -98
            _ = v49
        case "\n\t":
            var v50 uint8 = 2
            _ = v50
        case "foo bar":
            if (32767 < 32767) {
                v48 = int16(v1)
            } else {
                for fi51 := 0; fi51 < 7; fi51++ {
                    v52 := GenS9{F0: -128, F1: uint8(v46)}
                    _ = fi51
                    _ = v52
                }
            }
        default:
            fmt.Println(v1, v47)
        }
    }
    fmt.Println(v47, v3)
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 float64 = 1e10
        var v55 uint8 = 128
        _ = fi53
        _ = v54
        _ = v55
    }
    if false {
        if ((1 == 0) && ((v6 >= ((float32(1e10) + v6) / (v6 * ((float32(0.0) + v6) + ((v6 * -2.25) / ((v6 - v6) / (v6 * v6))))))) || v47)) {
            if true {
                delete(v5, v3)
                fmt.Println(v24)
            } else {
                for fi56 := 0; fi56 < 5; fi56++ {
                    for fi57 := 0; fi57 < 17; fi57++ {
                        var v58 float64 = 3.14159
                        fmt.Println(v46)
                        fmt.Println(v58, v3)
                        if v47 {
                            if ((v24 | 2) == (v24 | v24)) {
                                v48 = v48
                                var v59 bool = (!v47)
                                _ = v59
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    if v47 {
                                        var v60 int = 1
                                        v4 -= v60
                                        _ = v60
                                    } else {
                                        v61 := GenS5{F0: -1, F1: 127, F2: 0}
                                        _ = v61
                                    }
                                case "a":
                                    v62 := GenS7{F0: 1, F1: 10}
                                    _ = v62
                                case "":
                                    var v63 float32 = (1e10 / (0.0 * v6))
                                    _ = v63
                                }
                                v48 = -448
                                v64 := GenS6{F0: (^(fi56)), F1: (v48 << ((363) & 63)), F2: 343}
                                if v47 {
                                    switch v4 {
                                    case -323:
                                        v3 = v3
                                    case 83:
                                        var v65 int8 = 80
                                        _ = v65
                                    case 0:
                                        v66 := GenS10{F0: -625}
                                        _ = v66
                                    default:
                                        var v67 int32 = -1
                                        _ = v67
                                    }
                                    if (!false) {
                                        v4 -= 1
                                    }
                                    fmt.Println(fi57, v48)
                                }
                                if v47 {
                                    switch "a" {
                                    case "hello":
                                        v68 := GenS6{F0: (v4 | v4), F1: (v48 - (v48 / ((v48) | 1))), F2: -2147483648}
                                        switch v3 {
                                        case "foo bar":
                                            v69 := GenS2{F0: uint32(v24)}
                                            _ = v69
                                        default:
                                            v4 -= (fi57 - 589)
                                        }
                                        v6 -= v6
                                        _ = v68
                                    case "\n\t":
                                        var v70 int64 = -413
                                        _ = v70
                                    default:
                                        v71 := GenS2{F0: (v1 & 1)}
                                        switch v3 {
                                        case "":
                                            v48 = (v48 << ((v46) & 63))
                                        case "hello":
                                            v1 -= (v1 ^ 4294967295)
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        var v72 int = v4
                                        _ = v71
                                        _ = v72
                                    }
                                    var v73 float64 = (v58 + 0.0)
                                    _ = v73
                                }
                                fmt.Println(v3, v6)
                                delete(v5, (v3 + ""))
                                _ = v64
                            }
                            var v74 uint = genFn0()
                            _ = v74
                        } else {
                            var v75 uint64 = v24
                            var v76 int32 = 1
                            fmt.Println(v4)
                            _ = v75
                            _ = v76
                        }
                        _ = fi57
                        _ = v58
                    }
                    _ = fi56
                }
                if v47 {
                    if v47 {
                        v24 += v24
                        var v77 int = v4
                        _ = v77
                    } else {
                        var v78 int16 = (v48 & 0)
                        _ = v78
                    }
                }
                v79 := GenS11{F0: 0, F1: (v6 != v6), F2: (v1 ^ v1)}
                _ = v79
            }
            v80 := GenS6{F0: (v4 | (-(v4))), F1: 32767, F2: -2147483648}
            fmt.Println(v3, v24)
            _ = v80
        } else {
            var v81 uint16 = 42
            v3 = (v3 + v3)
            v81 -= (v81 ^ 468)
            _ = v81
        }
    } else {
        switch v3 {
        case "a":
            v82 := GenS11{F0: 1, F1: (!true), F2: v1}
            _ = v82
        case "\n\t":
            v83 := GenS9{F0: int8(v1), F1: 35}
            _ = v83
        default:
            for fi84 := 0; fi84 < 15; fi84++ {
                fmt.Println(v46)
                fmt.Println(fi84, v46, v48)
                var v85 uint = v46
                _ = fi84
                _ = v85
            }
        }
        switch v4 {
        case 2147483647:
            switch int(v46) {
            case 2147483647:
                v2 = append(v2, 65535)
            case 127:
                var v86 int = (v4 ^ (v4 * (v4 * v4)))
                _ = v86
            }
        case 850:
            switch -9223372036854775808 {
            case 2147483647:
                _ = len(v5)
            case 32768:
                if (v47 && (int16(v1) == v48)) {
                    if (0 <= 127) {
                        var v87 string = ((v3 + v3) + (v3 + v3))
                        v1 = v1
                        _ = v87
                    } else {
                        var v88 int64 = 711
                        var v89 int16 = v48
                        var v90 int32 = -1
                        _ = v88
                        _ = v89
                        _ = v90
                    }
                    fmt.Println(v6, v46)
                }
            default:
                v1 -= (v1 / (((v1 >> (((v46 + (v46 >> (((v46 % (((v46 ^ (v46 / (((v46 | (^(v46)))) | 1)))) | 1))) & 63)))) & 63))) | 1))
            }
        default:
            v91 := GenS9{F0: 116, F1: 0}
            _ = v91
        }
    }
    v92 := v5[(v3 + v3)]
    v92 -= (^(v4))
    if ((v4 * 32767) < (v4 % ((32768) | 1))) {
        for fi93 := 0; fi93 < 6; fi93++ {
            fmt.Println(v92)
            _ = fi93
        }
        v94 := GenS2{F0: (^(v1))}
        v1 = v1
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v92
}
