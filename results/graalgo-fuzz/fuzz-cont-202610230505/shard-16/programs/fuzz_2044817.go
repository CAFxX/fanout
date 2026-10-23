package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 string }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int16; F1 int8; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int32; F1 bool; F2 bool }
type GenS10 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 uint64; F1 int64 }
type GenS13 struct { F0 bool; F1 uint16 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int64 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch (-(fi7)) {
        case -578:
            var v8 uint8 = 245
            _ = v8
        case 32767:
            if (!(((!false) && false) && (311 >= 1))) {
                var v9 bool = false
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 14; fi10++ {
                switch "\n\t" {
                case "foo bar":
                    if (923 == 2) {
                        var v11 int8 = -1
                        _ = v11
                    }
                case "hello":
                    switch (^(fi7)) {
                    case -9223372036854775808:
                        var v12 int8 = 42
                        _ = v12
                    }
                    var v13 int32 = -2147483648
                    if false {
                        v14 := GenS1{F0: v13, F1: 42, F2: 126}
                        _ = v14
                    }
                    _ = v13
                }
                switch "" {
                case "":
                    var v15 int8 = 119
                    _ = v15
                case "hello":
                    var v16 uint32 = 37
                    _ = v16
                default:
                    var v17 string = ""
                    _ = v17
                }
                _ = fi10
            }
        }
        var v18 bool = false
        var v19 bool = (!(-2.25 <= 1.5))
        _ = fi7
        _ = v18
        _ = v19
    }
    var v20 float32 = 3.14159
    v21 := []int{-2147483648}
    var v22 float32 = ((v20 * ((3.14159 + (float32(0.1) / ((float32(1.5) - v20) - float32((1e10 - -2.25))))) + (-2.25 / 0.1))) * 1.5)
    _ = v20
    _ = v21
    _ = v22
    return -1
}

func main() {
    var v1 float32 = (1.5 - (3.14159 + -2.25))
    var v2 float32 = v1
    var v3 int64 = -6
    var v4 float64 = 1e10
    v5 := make(map[string]string)
    var v6 uint16 = 65535
    _ = len(v5)
    v23 := []int{0, 42}
    delete(v5, (("\n\t" + ("\n\t" + "\n\t")) + "a"))
    var v24 uint64 = 980
    if (!(2 == v24)) {
        fmt.Println(v2, v1, v2)
        if (v3 < (v3 ^ (^(v3)))) {
            fmt.Println(v2)
        }
    } else {
        for fi25 := 0; fi25 < 9; fi25++ {
            delete(v5, "\n\t")
            fmt.Println(fi25, fi25, v2)
            _ = fi25
        }
        var v26 uint8 = 127
        switch "\n\t" {
        case "\n\t":
            v27 := GenS1{F0: -1, F1: 127, F2: (v6 & v6)}
            _ = v27
        case "":
            var v28 int32 = -987
            _ = v28
        case "a":
            if true {
                for fi29 := 0; fi29 < 9; fi29++ {
                    v5["\n\t"] = "a"
                    switch (fi29 + fi29) {
                    case -612:
                        if (!true) {
                            v24 = (v24 >> ((4294967295) & 63))
                            if (42 == 0) {
                                for fi30 := 0; fi30 < 4; fi30++ {
                                    var v31 uint16 = 42
                                    _ = fi30
                                    _ = v31
                                }
                                v2 = 3.14159
                            }
                            var v32 uint32 = 229
                            _ = v32
                        }
                    case 127:
                        var v33 uint8 = v26
                        _ = v33
                    case -9223372036854775808:
                        if (-1 > -252) {
                            v34 := GenS2{F0: 1}
                            v35 := GenS3{F0: -957, F1: 42, F2: "foo bar"}
                            var v36 uint32 = 128
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    default:
                        if (398 != (v6 % ((0) | 1))) {
                            v24 = (v24 % ((v24) | 1))
                            _ = len(v5)
                        }
                    }
                    _ = fi29
                }
                for fi37 := 0; fi37 < 10; fi37++ {
                    v38 := GenS4{F0: v6, F1: -741}
                    var v39 uint64 = (v24 ^ (^(v24)))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            } else {
                if true {
                    fmt.Println(v2)
                }
                if (!false) {
                    switch "foo bar" {
                    case "hello":
                        var v40 int = 700
                        _ = v40
                    case "":
                        var v41 int32 = 536
                        _ = v41
                    }
                }
            }
        }
        _ = v26
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        for fi43 := 0; fi43 < 7; fi43++ {
            v6 = uint16(v24)
            _ = fi43
        }
        fmt.Println(v1, v6)
        _ = fi42
    }
    v24 = 2
    v44 := GenS3{F0: 1000, F1: 959, F2: ("foo bar" + (("a" + "a") + "a"))}
    v1 = v2
    v45 := GenS5{F0: -42, F1: -1, F2: -27}
    for fi46 := 0; fi46 < 19; fi46++ {
        fmt.Println(v6, v2, v2)
        v6 += 65535
        v47 := GenS6{F0: "a"}
        _ = fi46
        _ = v47
    }
    v48 := GenS7{F0: 65535, F1: 423}
    var v49 int16 = 0
    if (!(v6 == 951)) {
        switch 42 {
        case 641:
            v49 += v49
        case 2147483647:
            v50 := GenS4{F0: 750, F1: v49}
            _ = v50
        }
        for fi51 := 0; fi51 < 8; fi51++ {
            fmt.Println(v24, v49, v3)
            if true {
                var v52 int64 = (v3 * (v3 >> ((18446744073709551615) & 63)))
                if false {
                    if true {
                        for fi53 := 0; fi53 < 13; fi53++ {
                            v54 := GenS8{F0: 2147483647, F1: false}
                            var v55 int8 = -128
                            if true {
                                var v56 int16 = v49
                                var v57 uint8 = 71
                                _ = v56
                                _ = v57
                            } else {
                                v58 := GenS7{F0: 883, F1: 686}
                                delete(v5, "hello")
                                _ = v58
                            }
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        v5[("" + "a")] = ("foo bar" + (("hello" + (("" + "hello") + "\n\t")) + "hello"))
                        switch "foo bar" {
                        case "\n\t":
                            if (false || (true || false)) {
                                v4 += ((v4 + (float64(float32(((3.14159 / v4) + v4))) * ((0.0 - (0.1 / v4)) + 1e10))) + float64((v1 / ((v1 + v1) + v2))))
                            } else {
                                if len(v23) > 0 {
                                    v23[0] = fi51
                                }
                            }
                        case "foo bar":
                            if ((1 <= (v24 << ((4294967295) & 63))) || ((int32(v6) != 883) || true)) {
                                var v59 float32 = 1.5
                                var v60 uint16 = (v6 & (v6 + (v6 << ((42) & 63))))
                                _ = v59
                                _ = v60
                            } else {
                                if (158 > 2) {
                                    for fi61 := 0; fi61 < 15; fi61++ {
                                        v3 = (v52 >> ((620) & 63))
                                        switch (-(fi61)) {
                                        case -455:
                                            _ = len(v5)
                                        case -2147483648:
                                            v49 = v49
                                        }
                                        _ = fi61
                                    }
                                }
                                if ((v49 * (v49 + v49)) == (v49 - v49)) {
                                    v49 -= 0
                                } else {
                                    var v62 uint = 65535
                                    for fi63 := 0; fi63 < 3; fi63++ {
                                        v1 += v2
                                        _ = len(v23)
                                        _ = fi63
                                    }
                                    _ = v62
                                }
                                if false {
                                    var v64 float32 = (v2 * float32(v4))
                                    v65 := GenS8{F0: fi51, F1: true}
                                    _ = v64
                                    _ = v65
                                } else {
                                    var v66 string = "\n\t"
                                    v67 := GenS9{F0: 42, F1: true, F2: false}
                                    for fi68 := 0; fi68 < 16; fi68++ {
                                        if true {
                                            _ = len(v23)
                                            v69 := GenS10{F0: 128, F1: 933, F2: (v24 + 625)}
                                            v52 = v52
                                            if (!(100 < 0)) {
                                                v52 = (v3 << ((355) & 63))
                                            }
                                            _ = v69
                                        }
                                        v70 := GenS11{F0: 128, F1: false}
                                        _ = fi68
                                        _ = v70
                                    }
                                    _ = v66
                                    _ = v67
                                }
                            }
                        default:
                            switch "" {
                            case "a":
                                if true {
                                    var v71 uint32 = 316
                                    fmt.Println(v24, v24)
                                    v72 := GenS10{F0: 255, F1: 1, F2: (^(v24))}
                                    _ = v71
                                    _ = v72
                                }
                            default:
                                delete(v5, "\n\t")
                            }
                        }
                    } else {
                        var v73 uint16 = v6
                        for fi74 := 0; fi74 < 15; fi74++ {
                            fmt.Println(v1)
                            var v75 uint32 = 94
                            _ = fi74
                            _ = v75
                        }
                        _ = v73
                    }
                    if (v6 > (v6 - (v6 * v6))) {
                        v52 -= 891
                    }
                }
                _ = v52
            } else {
                if true {
                    v6 += v6
                    v24 = 4294967295
                    delete(v5, "a")
                }
                var v76 float64 = 1e10
                v77 := GenS12{F0: 585, F1: 0}
                _ = v76
                _ = v77
            }
            _ = fi51
        }
    } else {
        fmt.Println(v4)
        if false {
            if true {
                var v78 int16 = int16(v6)
                for fi79 := 0; fi79 < 5; fi79++ {
                    var v80 string = "\n\t"
                    fmt.Println(v6, v80)
                    _ = fi79
                    _ = v80
                }
                _ = v78
            }
            fmt.Println(v1, v3)
            for fi81 := 0; fi81 < 11; fi81++ {
                var v82 uint32 = 315
                switch (^(fi81)) {
                case -9223372036854775808:
                    v82 -= (v82 << ((263) & 63))
                case 1:
                    fmt.Println(v6, v2, v1)
                case 9223372036854775807:
                    var v83 uint32 = (v82 - 42)
                    _ = v83
                default:
                    switch (fi81 % ((-771) | 1)) {
                    case 519:
                        v84 := v5[((("a" + ("a" + "\n\t")) + ("hello" + "a")) + "foo bar")]
                        _ = v84
                    default:
                        if false {
                            var v85 int64 = int64(v24)
                            _ = v85
                        } else {
                            switch (("a" + "") + "hello") {
                            case "hello":
                                if (false || (!(false || true))) {
                                    var v86 int32 = 2147483647
                                    var v87 int16 = (v49 >> ((65535) & 63))
                                    _ = v86
                                    _ = v87
                                } else {
                                    for fi88 := 0; fi88 < 12; fi88++ {
                                        fmt.Println(v4)
                                        switch "foo bar" {
                                        case "":
                                            switch "foo bar" {
                                            case "a":
                                                var v89 int16 = v49
                                                _ = v89
                                            case "\n\t":
                                                if true {
                                                    if true {
                                                        var v90 float64 = ((1e10 / v4) - v4)
                                                        _ = v90
                                                    } else {
                                                        v49 -= -791
                                                    }
                                                } else {
                                                    v23 = append(v23, fi88)
                                                    _ = len(v5)
                                                }
                                            default:
                                                if (-653 < (v3 % ((genFn0()) | 1))) {
                                                    fmt.Println(v49, v24, v82)
                                                }
                                            }
                                            _ = len(v23)
                                        case "hello":
                                            var v91 uint32 = v82
                                            _ = v91
                                        case "foo bar":
                                            var v92 uint8 = 42
                                            _ = v92
                                        default:
                                            v93 := GenS13{F0: false, F1: v6}
                                            _ = v93
                                        }
                                        var v94 int64 = (v3 ^ genFn0())
                                        switch fi81 {
                                        case -2147483648:
                                            if true {
                                                var v95 uint32 = (^(v82))
                                                v3 = (-(v3))
                                                _ = v95
                                            }
                                        }
                                        _ = fi88
                                        _ = v94
                                    }
                                    delete(v5, "hello")
                                    v1 += float32(3.14159)
                                }
                            case "\n\t":
                                fmt.Println(v24, v2)
                            default:
                                for fi96 := 0; fi96 < 4; fi96++ {
                                    var v97 int64 = 1
                                    _ = fi96
                                    _ = v97
                                }
                            }
                        }
                    }
                }
                _ = fi81
                _ = v82
            }
        } else {
            v98 := GenS7{F0: 2147483648, F1: 684}
            _ = v98
        }
        v99 := GenS14{F0: (v3 / (((v3 - v3)) | 1))}
        _ = v99
    }
    if false {
        if true {
            v1 += 3.14159
            fmt.Println(v1, v1)
        } else {
            for fi100 := 0; fi100 < 19; fi100++ {
                if true {
                    var v101 uint16 = (v6 / ((897) | 1))
                    v5[("" + (("\n\t" + ("\n\t" + "")) + "\n\t"))] = "\n\t"
                    _ = v101
                }
                var v102 int8 = 127
                _ = fi100
                _ = v102
            }
            switch 127 {
            case 2147483647:
                v49 = v49
            case 9223372036854775807:
                for fi103 := 0; fi103 < 4; fi103++ {
                    switch fi103 {
                    case 9223372036854775807:
                        v104 := GenS6{F0: "a"}
                        _ = v104
                    default:
                        v105 := GenS13{F0: false, F1: 508}
                        _ = v105
                    }
                    _ = fi103
                }
            }
            if true {
                v24 += v24
            } else {
                if len(v23) > 0 {
                    v23[0] = -9223372036854775808
                }
                v106 := GenS14{F0: (-(v3))}
                var v107 uint8 = 124
                _ = v106
                _ = v107
            }
        }
        v2 = v1
        v108 := GenS15{F0: genFn0()}
        _ = v108
    }
    _ = len(v5)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
