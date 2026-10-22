package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int8; F1 int16; F2 int32 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string; F1 int64; F2 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 int16; F1 uint64; F2 int16 }
func genFn0() uint32 {
    var v7 uint = 1
    v7 -= 968
    var v8 string = ("a" + "hello")
    v9 := []int{9223372036854775807, -2147483648}
    _ = v7
    _ = v8
    _ = v9
    return 272
}
func genFn1(p0 string) uint32 {
    if false {
        var v10 int64 = 42
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                v10 = (v10 | (v10 * v10))
            }
            var v12 int16 = -1
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        v13 := []int{}
        _ = v13
    }
    var v14 float64 = float64(0.1)
    v15 := []int{-586, 0}
    p0 = p0
    _ = p0
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 int16 = 1
    v4 := []int{2147483647, -9223372036854775808, -411, 32768}
    v5 := make(map[int]int)
    var v6 int16 = v3
    var v16 uint = 4294967295
    if true {
        var v17 uint16 = 222
        if true {
            v18 := GenS1{F0: 1}
            for fi19 := 0; fi19 < 5; fi19++ {
                v20 := GenS1{F0: 4294967295}
                if false {
                    if (false && (!false)) {
                        v2[192] = "foo bar"
                        if true {
                            for fi21 := 0; fi21 < 9; fi21++ {
                                v1 = append(v1, 352)
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                switch fi21 {
                                case -369:
                                    if true {
                                        var v22 int32 = int32(v17)
                                        _ = v22
                                    } else {
                                        var v23 string = ("hello" + (("" + "") + ("\n\t" + "foo bar")))
                                        _ = v23
                                    }
                                    v6 = (v6 >> ((2) & 63))
                                    fmt.Println(v17, v6)
                                case 9223372036854775807:
                                    var v24 uint64 = 2
                                    _ = v24
                                }
                                _ = fi21
                            }
                            v3 += (v6 << ((v16) & 63))
                        }
                    }
                    v25 := GenS2{F0: 65535, F1: (128 != 86)}
                    _ = v25
                }
                _ = fi19
                _ = v20
            }
            fmt.Println(v6, v17)
            _ = v18
        }
        _ = v17
    } else {
        switch "\n\t" {
        case "hello":
            switch 4294967295 {
            case 128:
                var v26 int16 = 32767
                _ = v26
            case -825:
                var v27 uint8 = uint8(v16)
                _ = v27
            default:
                v28 := GenS1{F0: uint64(v16)}
                _ = v28
            }
        case "foo bar":
            _ = len(v5)
        default:
            for fi29 := 0; fi29 < 19; fi29++ {
                fmt.Println(v6, v16)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v30 bool = false
                _ = fi29
                _ = v30
            }
        }
        if (!(2 == 65535)) {
            var v31 uint64 = 389
            _ = v31
        }
        var v32 uint32 = 0
        _ = v32
    }
    var v33 uint8 = 106
    v3 = (v6 | (v6 & 42))
    switch "\n\t" {
    case "a":
        fmt.Println(v6, v6)
    case "hello":
        switch 127 {
        case 128:
            if true {
                if false {
                    v34 := GenS1{F0: 42}
                    _ = v34
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v35 int32 = 0
                    v36 := v5[42]
                    _ = v35
                    _ = v36
                }
                for fi37 := 0; fi37 < 9; fi37++ {
                    var v38 int16 = (v3 / ((v3) | 1))
                    var v39 uint64 = 747
                    v40 := GenS1{F0: (v39 + (^(v39)))}
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
            } else {
                v41 := GenS2{F0: 961, F1: (!(((0.0 + 0.1) < 3.14159) && true))}
                _ = v41
            }
        case 42:
            var v42 float32 = ((0.1 - (0.0 - float32(0.1))) * 0.0)
            _ = v42
        }
    case "":
        if true {
            if (true || (v33 >= v33)) {
                if true {
                    var v43 int16 = -32768
                    _ = v43
                } else {
                    v44 := GenS2{F0: genFn0(), F1: (3.14159 < (1.5 / (-2.25 - (1e10 * float32(-2.25)))))}
                    var v45 uint = 2
                    _ = v44
                    _ = v45
                }
            }
            fmt.Println(v3)
        }
    }
    fmt.Println(v6)
    var v46 int16 = v3
    v16 = v16
    for fi47 := 0; fi47 < 15; fi47++ {
        v6 = 1
        _ = fi47
    }
    v6 = (v46 << (((v16 & 0)) & 63))
    for fi48 := 0; fi48 < 6; fi48++ {
        if (!(42 <= 32767)) {
            v46 = v3
            if (823 >= 255) {
                for fi49 := 0; fi49 < 14; fi49++ {
                    var v50 float32 = 0.0
                    v6 -= (^(v3))
                    var v51 int = fi49
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                for fi52 := 0; fi52 < 11; fi52++ {
                    switch (fi52 >> (((v16 % ((v16) | 1))) & 63)) {
                    case 127:
                        var v53 uint32 = 0
                        _ = v53
                    case -9223372036854775808:
                        if false {
                            v54 := GenS1{F0: 2}
                            var v55 int = 285
                            _ = v54
                            _ = v55
                        } else {
                            v46 = -32768
                            fmt.Println(fi52)
                        }
                    case 9223372036854775807:
                        v6 -= (^(v46))
                    }
                    fmt.Println(v3, fi48, v16)
                    fmt.Println(v3, v33)
                    _ = fi52
                }
            } else {
                v56 := GenS2{F0: 834, F1: false}
                for fi57 := 0; fi57 < 14; fi57++ {
                    if false {
                        v6 += 32767
                    } else {
                        v58 := v5[fi48]
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                        _ = v58
                    }
                    var v59 int32 = -705
                    _ = fi57
                    _ = v59
                }
                _ = v56
            }
            if true {
                for fi60 := 0; fi60 < 18; fi60++ {
                    v61 := GenS1{F0: 18446744073709551615}
                    if (833 < 1) {
                        if (true || (fi60 >= 0)) {
                            var v62 float32 = (((-2.25 + 0.0) + (0.1 / -2.25)) * 0.1)
                            _ = v62
                        } else {
                            v16 = 255
                            var v63 uint16 = 65535
                            if false {
                                v64 := GenS3{F0: ((fi48 / (((fi48 * fi60)) | 1)) == (fi48 ^ fi60)), F1: v16}
                                if true {
                                    v6 += 32767
                                    var v65 int64 = 9223372036854775807
                                    var v66 float32 = 0.1
                                    if len(v4) > 0 {
                                        v4[0] = -2147483648
                                    }
                                    _ = v65
                                    _ = v66
                                } else {
                                    v3 = (^(v6))
                                    switch fi48 {
                                    case 169:
                                        fmt.Println(fi48, v6, fi48)
                                    case 2147483647:
                                        var v67 int = (fi60 & -1)
                                        _ = v67
                                    case 32767:
                                        var v68 int = 9223372036854775807
                                        _ = v68
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            var v69 int8 = 42
                                            _ = v69
                                        case "a":
                                            v70 := v5[(fi48 ^ (fi60 ^ -9223372036854775808))]
                                            _ = v70
                                        default:
                                            switch ("" + "hello") {
                                            case "hello":
                                                var v71 string = ""
                                                _ = v71
                                            case "":
                                                delete(v2, 1)
                                            }
                                            if (-128 == -128) {
                                                var v72 uint8 = (v33 | (v33 >> ((1) & 63)))
                                                if false {
                                                    v6 += int16(v63)
                                                    v73 := GenS2{F0: 871, F1: (!true)}
                                                    _ = v73
                                                } else {
                                                    var v74 string = "hello"
                                                    _ = v74
                                                }
                                                v75 := GenS2{F0: 65535, F1: true}
                                                _ = v72
                                                _ = v75
                                            }
                                        }
                                    }
                                }
                                _ = v64
                            } else {
                                v63 = (^(v63))
                                var v76 int32 = 42
                                v46 = (v6 & (v6 & (v46 - 960)))
                                _ = v76
                            }
                            _ = v63
                        }
                    }
                    v6 = v3
                    _ = fi60
                    _ = v61
                }
            } else {
                switch (fi48 ^ fi48) {
                case 9223372036854775807:
                    var v77 int16 = v6
                    _ = v77
                default:
                    fmt.Println(fi48, fi48)
                }
                switch fi48 {
                case 818:
                    if true {
                        fmt.Println(v46, fi48, v33)
                        if ((((^(v33)) >= (v33 >> ((0) & 63))) || (!(("hello" <= ("" + (("a" + ("foo bar" + "")) + ""))) || (!false)))) && true) {
                            v33 = 134
                            v3 -= v46
                            v3 = 32767
                        }
                    } else {
                        switch (fi48 % (((^(fi48))) | 1)) {
                        case 1:
                            switch "hello" {
                            case "hello":
                                if (v46 > (v3 / (((v3 << (((v16 - 367)) & 63))) | 1))) {
                                    if true {
                                        v78 := GenS4{F0: 503, F1: 32767}
                                        v79 := GenS5{F0: 0, F1: 651, F2: -2147483648}
                                        var v80 int16 = v3
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                    } else {
                                        v16 -= (v16 / (((v16 >> ((v16) & 63))) | 1))
                                        switch (fi48 * fi48) {
                                        case -2147483648:
                                            for fi81 := 0; fi81 < 14; fi81++ {
                                                v82 := GenS6{F0: int8(v33), F1: 42}
                                                var v83 float64 = 3.14159
                                                fmt.Println(v3, fi48, fi48)
                                                fmt.Println(v3, v33, v6)
                                                v84 := GenS6{F0: 86, F1: 1}
                                                _ = fi81
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                            }
                                        case 127:
                                            delete(v5, fi48)
                                        default:
                                            v85 := GenS7{F0: true}
                                            _ = v85
                                        }
                                    }
                                    var v86 int16 = (v46 * 0)
                                    _ = v86
                                }
                            }
                        case 416:
                            var v87 uint32 = uint32(v33)
                            _ = v87
                        }
                    }
                case 4294967295:
                    v3 += v46
                }
            }
        } else {
            var v88 int = (fi48 | fi48)
            _ = v88
        }
        _ = fi48
    }
    v89 := v5[2147483647]
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 259 {
    case -9223372036854775808:
        fmt.Println(v33, v16)
    case 974:
        if true {
            for fi90 := 0; fi90 < 12; fi90++ {
                v91 := GenS8{F0: (("\n\t" + "") + "a"), F1: 580, F2: v46}
                var v92 uint8 = v33
                _ = fi90
                _ = v91
                _ = v92
            }
            var v93 float32 = ((3.14159 * (3.14159 / -2.25)) * 0.0)
            v94 := GenS6{F0: -25, F1: 584}
            _ = v93
            _ = v94
        } else {
            switch ("a" + ("\n\t" + "a")) {
            case "":
                delete(v5, 128)
            case "a":
                v16 = v16
            case "\n\t":
                v3 += (v46 * v3)
            default:
                for fi95 := 0; fi95 < 2; fi95++ {
                    var v96 uint = v16
                    _ = fi95
                    _ = v96
                }
            }
            v33 = v33
        }
    case 2147483647:
        if (false && (1 <= 1)) {
            var v97 int = (v89 - (^(v89)))
            var v98 uint16 = 32767
            _ = v97
            _ = v98
        }
    default:
        if (genFn0() <= 65535) {
            v99 := GenS9{F0: (("foo bar" + "hello") + "")}
            v100 := GenS5{F0: -128, F1: int16(v16), F2: -755}
            _ = v100
            _ = v99
        } else {
            _ = len(v1)
            v101 := GenS8{F0: "a", F1: 42, F2: 42}
            _ = len(v5)
            _ = v101
        }
    }
    v102 := GenS3{F0: (2 >= 164), F1: 373}
    switch ("a" + ("\n\t" + "a")) {
    case "\n\t":
        var v103 int8 = 100
        _ = v103
    case "hello":
        if true {
            v104 := GenS7{F0: (42 < 251)}
            var v105 int = int(v3)
            _ = v104
            _ = v105
        }
    case "foo bar":
        v106 := GenS10{F0: true, F1: 42}
        _ = v106
    }
    for fi107 := 0; fi107 < 12; fi107++ {
        _ = len(v2)
        v3 -= (v46 + 1000)
        v108 := GenS11{F0: v3, F1: 1, F2: (v6 / ((32767) | 1))}
        _ = fi107
        _ = v108
    }
    v109 := GenS6{F0: 127, F1: 671}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v102
    _ = v109
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v89
}
