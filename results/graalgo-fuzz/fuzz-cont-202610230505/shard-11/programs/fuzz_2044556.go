package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint; F2 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int8; F2 int32 }
type GenS8 struct { F0 uint8; F1 int16; F2 int32 }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint64; F1 uint32 }
type GenS13 struct { F0 uint64; F1 uint8 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int32; F1 bool; F2 string }
type GenS16 struct { F0 uint16; F1 int; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 1
    var v3 uint = 255
    var v4 uint8 = 127
    v5 := make(map[uint32]uint32)
    var v6 string = ("\n\t" + ("hello" + "a"))
    switch int(v2) {
    case 32768:
        v7 := GenS1{F0: -128, F1: v3, F2: -128}
        _ = v7
    case 127:
        for fi8 := 0; fi8 < 13; fi8++ {
            if false {
                var v9 int = 2147483647
                if (!((fi8 & (v9 | v9)) > -2147483648)) {
                    v10 := []int{fi8, (fi8 + (fi8 % (((^(v9))) | 1))), (fi8 % (((v9 / (((v9 & v9)) | 1))) | 1)), fi8}
                    for fi11 := 0; fi11 < 19; fi11++ {
                        _ = len(v5)
                        switch v6 {
                        case "":
                            if false {
                                v10 = append(v10, fi11)
                                var v12 string = ((("\n\t" + "foo bar") + v6) + "foo bar")
                                _ = v12
                            } else {
                                if (v2 != v2) {
                                    fmt.Println(v9, fi8, fi11)
                                }
                                fmt.Println(fi11)
                            }
                        case "foo bar":
                            var v13 int8 = 127
                            _ = v13
                        default:
                            v14 := GenS1{F0: -128, F1: v3, F2: 0}
                            _ = v14
                        }
                        _ = fi11
                    }
                    _ = v10
                }
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        var v15 float32 = 1.5
                        _ = v15
                    case "a":
                        var v16 uint32 = 2
                        _ = v16
                    default:
                        v17 := GenS1{F0: -128, F1: uint(v2), F2: 0}
                        _ = v17
                    }
                    v9 += v9
                }
                _ = v9
            }
            var v18 bool = (-847 <= -1)
            _ = fi8
            _ = v18
        }
    case 32767:
        fmt.Println(v4)
    }
    var v19 uint32 = 374
    delete(v1, -17)
    if ((!true) || (1 == 184)) {
        if (float32(0.0) == 0.1) {
            for fi20 := 0; fi20 < 9; fi20++ {
                switch "" {
                case "hello":
                    v6 = ("hello" + "hello")
                }
                if false {
                    var v21 int8 = -128
                    var v22 uint16 = 1
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
            if false {
                if true {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        if true {
                            var v24 bool = (!true)
                            _ = v24
                        }
                        _ = fi23
                    }
                    v3 += 0
                } else {
                    v4 -= uint8(v3)
                    v25 := GenS2{F0: "a", F1: 921}
                    v26 := GenS3{F0: 2, F1: (v19 << ((v3) & 63)), F2: v6}
                    _ = v25
                    _ = v26
                }
            }
            var v27 float64 = 3.14159
            _ = v27
        } else {
            var v28 uint64 = 4294967295
            for fi29 := 0; fi29 < 18; fi29++ {
                switch v6 {
                case "foo bar":
                    if (-9223372036854775808 != 0) {
                        v30 := []int{fi29, -1, 32768}
                        var v31 float64 = -2.25
                        _ = v30
                        _ = v31
                    } else {
                        var v32 uint = (v3 / ((65535) | 1))
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v34 := []int{1, 262}
                            if (!(1000 == 1000)) {
                                v19 += (v19 % (((v19 - (^(v19)))) | 1))
                            } else {
                                var v35 float32 = 0.0
                                v28 = (^(v28))
                                fmt.Println(v3, v28)
                                _ = v35
                            }
                            v36 := GenS4{F0: 0}
                            _ = fi33
                            _ = v34
                            _ = v36
                        }
                        var v37 int64 = 368
                        _ = v32
                        _ = v37
                    }
                case "a":
                    var v38 int64 = 202
                    _ = v38
                default:
                    if true {
                        var v39 uint8 = 255
                        _ = v39
                    } else {
                        v40 := v5[(v19 ^ 2147483648)]
                        _ = v40
                    }
                }
                if false {
                    v28 = 1
                } else {
                    delete(v5, v19)
                }
                _ = fi29
            }
            v4 += v4
            _ = v28
        }
        var v41 bool = true
        var v42 int8 = int8(v19)
        _ = v41
        _ = v42
    }
    v4 = v4
    v43 := []int{32768}
    if (42 <= -2147483648) {
        v44 := GenS5{F0: (v6 + ((((v6 + "hello") + v6) + v6) + v6)), F1: (v3 | (v3 % (((v3 / (((v3 << ((v3) & 63))) | 1))) | 1))), F2: v19}
        switch 409 {
        case 32767:
            switch v6 {
            case "":
                v6 = v6
            case "\n\t":
                fmt.Println(v6, v19)
            case "foo bar":
                v45 := GenS6{F0: 2}
                _ = v45
            default:
                if true {
                    v46 := GenS7{F0: 42, F1: -111, F2: -467}
                    if (!true) {
                        for fi47 := 0; fi47 < 16; fi47++ {
                            v48 := GenS8{F0: v4, F1: 42, F2: 992}
                            _ = fi47
                            _ = v48
                        }
                    }
                    _ = v46
                } else {
                    v19 -= 65535
                    switch v6 {
                    case "a":
                        if false {
                            fmt.Println(v4, v6, v19)
                            v43 = append(v43, 127)
                        }
                    case "\n\t":
                        v6 = "a"
                    case "":
                        delete(v5, 25)
                    }
                    var v49 int8 = -41
                    _ = v49
                }
            }
        default:
            v50 := GenS3{F0: 118, F1: v19, F2: ("hello" + "foo bar")}
            _ = v50
        }
        if false {
            v51 := GenS5{F0: "a", F1: (^(v3)), F2: v19}
            v52 := GenS9{F0: ((0.1 + (-2.25 - 0.1)) == 1e10), F1: (^(v19))}
            _ = v51
            _ = v52
        } else {
            for fi53 := 0; fi53 < 2; fi53++ {
                v6 = v6
                _ = fi53
            }
        }
        _ = v44
    } else {
        if (100 < 77) {
            v19 -= 42
        } else {
            v54 := GenS10{F0: 95}
            _ = v54
        }
        v55 := GenS11{F0: 452, F1: v6}
        _ = v55
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        for fi57 := 0; fi57 < 14; fi57++ {
            if (fi57 != fi57) {
                if false {
                    if true {
                        v58 := GenS1{F0: 1, F1: (v3 >> (((v3 + 255)) & 63)), F2: -128}
                        if (uint16(v3) >= (^(v2))) {
                            v2 = v2
                        } else {
                            if len(v43) > 0 {
                                v43[0] = 4294967295
                            }
                            if true {
                                var v59 int = fi56
                                v4 = v4
                                _ = v59
                            } else {
                                fmt.Println(v4)
                                if false {
                                    if (!(false || ((!false) && false))) {
                                        var v60 uint32 = uint32(v3)
                                        v60 -= (v19 << (((v3 % ((v3) | 1))) & 63))
                                        v6 = v6
                                        _ = v60
                                    }
                                } else {
                                    if len(v43) > 0 {
                                        v43[0] = fi57
                                    }
                                    var v61 int64 = -866
                                    var v62 int32 = 2147483647
                                    _ = v61
                                    _ = v62
                                }
                                v63 := v5[65535]
                                _ = v63
                            }
                            v4 = 206
                        }
                        var v64 bool = true
                        _ = v58
                        _ = v64
                    } else {
                        v3 += (v3 * 65535)
                        if ((^(v19)) != 4294967295) {
                            v4 += v4
                            switch (v6 + v6) {
                            case "hello":
                                var v65 int8 = 100
                                _ = v65
                            case "a":
                                v66 := GenS12{F0: 560, F1: 0}
                                _ = v66
                            case "":
                                v4 += v4
                            }
                            switch 440 {
                            case 1:
                                v3 += v3
                            case -60:
                                v2 += uint16(v4)
                            }
                        }
                        v67 := GenS10{F0: int8(v3)}
                        _ = v67
                    }
                    switch fi56 {
                    case 1:
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v3, v6, fi57)
                        case "\n\t":
                            v19 = (v19 ^ v19)
                        }
                    default:
                        var v68 float64 = 0.0
                        _ = v68
                    }
                    fmt.Println(fi56)
                } else {
                    var v69 int32 = 42
                    if true {
                        var v70 int64 = 9223372036854775807
                        var v71 float32 = 3.14159
                        v72 := GenS10{F0: 100}
                        _ = v70
                        _ = v71
                        _ = v72
                    }
                    _ = v69
                }
                v2 = uint16(v19)
                _ = len(v5)
                var v73 bool = (100 < -2)
                v74 := GenS13{F0: 4294967295, F1: v4}
                _ = v73
                _ = v74
            }
            _ = fi57
        }
        var v75 int64 = 42
        _ = fi56
        _ = v75
    }
    switch -9223372036854775808 {
    case 571:
        delete(v5, (v19 << (((v3 ^ (v3 * (v3 % ((943) | 1))))) & 63)))
    case 4294967295:
        v76 := GenS14{F0: -32768}
        _ = v76
    case 32767:
        v77 := GenS9{F0: true, F1: (v19 ^ (v19 * v19))}
        _ = v77
    default:
        if (!false) {
            var v78 string = (v6 + (v6 + v6))
            _ = v78
        }
    }
    delete(v5, uint32(v3))
    v19 = (^(v19))
    for fi79 := 0; fi79 < 16; fi79++ {
        v4 = v4
        switch (fi79 ^ -203) {
        case 1:
            v80 := GenS4{F0: 32767}
            _ = v80
        case 32767:
            var v81 uint32 = 2147483647
            _ = v81
        case 127:
            if false {
                if true {
                    var v82 uint = v3
                    _ = v82
                }
                if (32767 >= v2) {
                    for fi83 := 0; fi83 < 6; fi83++ {
                        if (0.1 < (1.5 * ((0.0 * 0.1) / 1e10))) {
                            if ((!false) || true) {
                                v2 += (^(v2))
                                v84 := GenS8{F0: v4, F1: 32767, F2: 2147483647}
                                _ = v84
                            } else {
                                switch (fi79 % (((fi83 ^ (^(fi83)))) | 1)) {
                                case 907:
                                    v2 = v2
                                case 1:
                                    v6 = ("a" + ("hello" + "\n\t"))
                                case 599:
                                    var v85 int8 = 42
                                    _ = v85
                                default:
                                    if len(v43) > 0 {
                                        v43[0] = (fi83 / (((fi79 * (fi83 * 354))) | 1))
                                    }
                                }
                                switch fi83 {
                                case 1:
                                    fmt.Println(v4)
                                case -9223372036854775808:
                                    var v86 string = v6
                                    _ = v86
                                case 32767:
                                    var v87 uint = 18446744073709551615
                                    _ = v87
                                }
                                v19 += (v19 * (^(v19)))
                            }
                            if false {
                                if (true && false) {
                                    if false {
                                        v88 := GenS15{F0: 0, F1: (!(v2 >= 32767)), F2: "hello"}
                                        _ = v88
                                    }
                                    v89 := GenS13{F0: 4294967295, F1: uint8(v3)}
                                    _ = v89
                                }
                                v90 := GenS5{F0: ((("hello" + "\n\t") + "foo bar") + ("\n\t" + v6)), F1: v3, F2: v19}
                                _ = v90
                            }
                        } else {
                            if true {
                                var v91 int8 = 0
                                v5[v19] = (v19 & v19)
                                _ = v91
                            } else {
                                var v92 int8 = -128
                                var v93 int64 = -248
                                _ = v92
                                _ = v93
                            }
                            v94 := GenS12{F0: 520, F1: v19}
                            v3 = 4294967295
                            _ = v94
                        }
                        fmt.Println(v19, v3)
                        v2 -= (^(v2))
                        _ = fi83
                    }
                } else {
                    for fi95 := 0; fi95 < 8; fi95++ {
                        var v96 int = fi95
                        _ = fi95
                        _ = v96
                    }
                    var v97 int8 = 127
                    _ = v97
                }
            }
        default:
            if (!(98 > 42)) {
                var v98 float32 = 0.1
                fmt.Println(v4, v3)
                _ = v98
            }
        }
        delete(v5, (^(v19)))
        _ = fi79
    }
    fmt.Println(v3, v19, v6)
    v19 -= uint32(v4)
    if ((!false) && true) {
        if (!(!(-9223372036854775808 <= 32768))) {
            v2 = (v2 - 0)
            if len(v43) > 0 {
                v43[0] = 128
            }
        } else {
            for fi99 := 0; fi99 < 3; fi99++ {
                if (!false) {
                    v2 -= uint16(v3)
                } else {
                    v2 += v2
                    fmt.Println(v3, fi99, v6)
                }
                v100 := GenS16{F0: v2, F1: -496, F2: -442}
                _ = fi99
                _ = v100
            }
            v2 = (v2 << ((18446744073709551615) & 63))
            var v101 int16 = -1
            _ = v101
        }
        v19 -= uint32(v4)
    } else {
        fmt.Println(v2, v2, v6)
        var v102 float32 = 1.5
        for fi103 := 0; fi103 < 14; fi103++ {
            var v104 int64 = 1
            _ = fi103
            _ = v104
        }
        _ = v102
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
