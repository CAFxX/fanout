package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 string; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int8; F1 int32; F2 int8 }
type GenS8 struct { F0 uint16; F1 int32; F2 string }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 int; F1 int64; F2 int32 }
type GenS11 struct { F0 bool; F1 int64; F2 uint64 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint = 0
    v3 := []int{v1, v1}
    var v4 int8 = -1
    v5 := []int{127, v1, v1}
    var v6 uint64 = 909
    if false {
        v7 := GenS1{F0: -1, F1: 601, F2: 255}
        _ = v7
    } else {
        v4 -= v4
        _ = len(v5)
    }
    _ = len(v3)
    v6 += v6
    if false {
        var v8 bool = false
        _ = v8
    }
    switch (-(v1)) {
    case -2147483648:
        var v9 int64 = -9223372036854775808
        _ = v9
    case -9223372036854775808:
        v10 := GenS1{F0: 1, F1: -754, F2: 49}
        _ = v10
    case 162:
        var v11 uint16 = 63
        _ = v11
    default:
        v1 = (v1 ^ 32767)
    }
    if false {
        switch -9223372036854775808 {
        case 9223372036854775807:
            v12 := GenS1{F0: 1, F1: 1000, F2: 42}
            _ = v12
        case 32768:
            for fi13 := 0; fi13 < 12; fi13++ {
                if (3.14159 < float64((3.14159 - 1.5))) {
                    v6 = v6
                    switch "" {
                    case "a":
                        switch 0 {
                        case -9223372036854775808:
                            if (746 <= 255) {
                                fmt.Println(v2, v4)
                            } else {
                                v14 := GenS1{F0: 43, F1: 1000, F2: 87}
                                fmt.Println(v6)
                                v15 := GenS2{F0: 67}
                                _ = v14
                                _ = v15
                            }
                        case -2147483648:
                            var v16 bool = (!(-2.25 == (1e10 * 1.5)))
                            _ = v16
                        case 42:
                            fmt.Println(fi13)
                        default:
                            v17 := GenS3{F0: 578, F1: "foo bar", F2: 42}
                            _ = v17
                        }
                    default:
                        for fi18 := 0; fi18 < 10; fi18++ {
                            fmt.Println(v2, v1, v4)
                            var v19 string = (("\n\t" + (("foo bar" + ("hello" + "hello")) + "hello")) + "\n\t")
                            v19 = v19
                            if false {
                                v6 = v6
                            }
                            var v20 int16 = -32768
                            var v21 float64 = 0.1
                            v22 := GenS2{F0: v20}
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                    v4 = (v4 << ((v2) & 63))
                }
                if (45 < 331) {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        v2 -= (v2 >> (((^(v2))) & 63))
                        _ = fi23
                    }
                } else {
                    if true {
                        v4 = -53
                        var v24 uint16 = 42
                        _ = v24
                    }
                    v2 = (v2 * 255)
                    v25 := GenS4{F0: 0}
                    _ = v25
                }
                _ = fi13
            }
        case 4294967295:
            v2 = v2
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v5 = append(v5, 9223372036854775807)
    } else {
        var v26 uint8 = 98
        var v27 uint64 = 346
        _ = v26
        _ = v27
    }
    if (1 <= 580) {
        v28 := GenS5{F0: 9223372036854775807, F1: -581}
        fmt.Println(v6)
        _ = v28
    }
    switch (v1 % (((-(v1))) | 1)) {
    case -9223372036854775808:
        fmt.Println(v4)
    }
    v29 := GenS2{F0: 1000}
    switch v1 {
    case 4294967295:
        for fi30 := 0; fi30 < 19; fi30++ {
            if true {
                for fi31 := 0; fi31 < 17; fi31++ {
                    fmt.Println(v4)
                    v32 := GenS2{F0: 0}
                    if (1e10 >= ((((0.0 + 1.5) + 0.0) + 3.14159) / -2.25)) {
                        switch ("foo bar" + ("hello" + "foo bar")) {
                        case "hello":
                            v2 = (^(v2))
                        }
                        var v33 float32 = 0.0
                        var v34 uint32 = 4294967295
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                v3 = append(v3, v1)
                v3 = append(v3, 32767)
            } else {
                v4 = -128
                v6 = (v6 << ((v2) & 63))
                switch v1 {
                case -626:
                    for fi35 := 0; fi35 < 3; fi35++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi35
                    }
                case 4294967295:
                    if (false || (-366 > 9223372036854775807)) {
                        for fi36 := 0; fi36 < 9; fi36++ {
                            var v37 int32 = 1
                            _ = fi36
                            _ = v37
                        }
                        v1 = -351
                        v38 := GenS6{F0: -396, F1: "a"}
                        _ = v38
                    } else {
                        if false {
                            var v39 float32 = 0.0
                            _ = v39
                        }
                        if false {
                            for fi40 := 0; fi40 < 1; fi40++ {
                                if true {
                                    if (true && false) {
                                        var v41 int = (fi30 & v1)
                                        fmt.Println(fi40, fi30, fi30)
                                        _ = v41
                                    }
                                    if true {
                                        if len(v5) > 0 {
                                            v5[0] = -756
                                        }
                                        v2 = v2
                                    } else {
                                        var v42 uint32 = 675
                                        _ = v42
                                    }
                                    if true {
                                        v2 += (v2 - v2)
                                    } else {
                                        v4 = 0
                                    }
                                } else {
                                    _ = len(v3)
                                    var v43 int64 = 42
                                    if len(v5) > 0 {
                                        v5[0] = int(v4)
                                    }
                                    _ = v43
                                }
                                if (((((1.5 - (1e10 * float32(((1e10 * -2.25) + 0.1)))) * 1e10) / 1.0) + -2.25) <= 1e10) {
                                    fmt.Println(fi30, v4, v6)
                                    v44 := GenS3{F0: 1, F1: "hello", F2: 2147483647}
                                    v2 = 38
                                    v4 = v4
                                    var v45 uint = (v2 - v2)
                                    fmt.Println(v2, fi30, fi40)
                                    var v46 int16 = 687
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                                var v47 int64 = 843
                                _ = fi40
                                _ = v47
                            }
                            v4 = (v4 * -128)
                            fmt.Println(v4)
                        } else {
                            if true {
                                var v48 int64 = -1
                                _ = v48
                            } else {
                                if true {
                                    for fi49 := 0; fi49 < 20; fi49++ {
                                        if ((42 < 439) || true) {
                                            var v50 int64 = 0
                                            _ = v50
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = fi49
                                    }
                                    if ("hello" >= ("a" + "\n\t")) {
                                        for fi51 := 0; fi51 < 4; fi51++ {
                                            var v52 int8 = v4
                                            var v53 int32 = -2147483648
                                            v6 += 4294967295
                                            _ = fi51
                                            _ = v52
                                            _ = v53
                                        }
                                        fmt.Println(fi30, v4, v6)
                                    }
                                    switch v1 {
                                    case 128:
                                        for fi54 := 0; fi54 < 5; fi54++ {
                                            v2 -= (v2 / ((4294967295) | 1))
                                            v55 := GenS4{F0: 580}
                                            fmt.Println(fi54, v2)
                                            v5 = append(v5, 718)
                                            _ = fi54
                                            _ = v55
                                        }
                                    default:
                                        var v56 int64 = 42
                                        _ = v56
                                    }
                                }
                                fmt.Println(fi30, v6)
                            }
                            _ = len(v5)
                        }
                    }
                case 0:
                    v5 = append(v5, v1)
                }
            }
            _ = fi30
        }
    case 2147483647:
        var v57 bool = (false || (42 > 4294967295))
        _ = v57
    case -595:
        if ((v1 * v1) < -88) {
            v6 = 566
            switch (v1 + v1) {
            case -445:
                v2 = v2
            }
            switch "foo bar" {
            case "\n\t":
                v2 -= v2
            case "foo bar":
                fmt.Println(v4, v4, v2)
            }
        } else {
            if (!false) {
                v58 := GenS7{F0: 42, F1: 1, F2: (v4 % (((^(v4))) | 1))}
                _ = v58
            } else {
                v4 -= (v4 >> ((65535) & 63))
            }
        }
    default:
        switch (^(v1)) {
        case 2147483647:
            if true {
                var v59 uint = v2
                for fi60 := 0; fi60 < 14; fi60++ {
                    v61 := GenS8{F0: 844, F1: -564, F2: "hello"}
                    _ = fi60
                    _ = v61
                }
                if len(v5) > 0 {
                    v5[0] = 32767
                }
                _ = v59
            }
        case -2147483648:
            v62 := GenS1{F0: -1, F1: 42, F2: 3}
            _ = v62
        default:
            v1 += -373
        }
    }
    fmt.Println(v1, v4)
    switch v1 {
    case -9223372036854775808:
        if false {
            v6 -= (v6 | (v6 | (v6 >> (((v2 | 2)) & 63))))
        }
    case 32768:
        for fi63 := 0; fi63 < 5; fi63++ {
            var v64 int16 = 1
            var v65 int32 = -1
            v66 := GenS6{F0: -576, F1: "\n\t"}
            _ = fi63
            _ = v64
            _ = v65
            _ = v66
        }
    default:
        if true {
            v67 := GenS9{F0: true, F1: 42}
            if true {
                fmt.Println(v1, v4, v4)
                var v68 int64 = 9223372036854775807
                _ = v68
            } else {
                if false {
                    for fi69 := 0; fi69 < 16; fi69++ {
                        fmt.Println(v6)
                        v2 -= v2
                        _ = fi69
                    }
                    v2 += v2
                } else {
                    for fi70 := 0; fi70 < 18; fi70++ {
                        fmt.Println(v2, v1)
                        var v71 float64 = ((0.1 / 1e10) * -2.25)
                        for fi72 := 0; fi72 < 2; fi72++ {
                            if (float32(((float64((1.5 - float32(v71))) * (v71 + 1e10)) / 1e10)) > 3.14159) {
                                var v73 int32 = 42
                                _ = v73
                            } else {
                                if false {
                                    if true {
                                        switch "foo bar" {
                                        case "a":
                                            v6 = 85
                                        case "\n\t":
                                            _ = len(v3)
                                        }
                                        v6 += 0
                                        v1 -= (fi70 % (((fi70 / ((fi72) | 1))) | 1))
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v74 uint16 = 295
                                        _ = v74
                                    }
                                    var v75 bool = (!((276 != 0) || true))
                                    _ = v75
                                } else {
                                    if (-1 <= 92) {
                                        v1 += (fi70 * fi70)
                                    }
                                    var v76 uint64 = v6
                                    _ = v76
                                }
                                v77 := GenS2{F0: -32768}
                                v78 := GenS10{F0: v1, F1: -1, F2: 0}
                                _ = v77
                                _ = v78
                            }
                            v71 += 0.1
                            var v79 uint64 = (v6 | v6)
                            _ = fi72
                            _ = v79
                        }
                        _ = fi70
                        _ = v71
                    }
                    for fi80 := 0; fi80 < 17; fi80++ {
                        v81 := GenS1{F0: 32767, F1: -1, F2: 255}
                        switch v1 {
                        case 424:
                            var v82 int8 = v4
                            _ = v82
                        default:
                            if true {
                                v83 := GenS1{F0: 42, F1: 0, F2: 0}
                                if false {
                                    if true {
                                        v84 := GenS9{F0: true, F1: 1}
                                        _ = v84
                                    }
                                    fmt.Println(v6, v6, v6)
                                    var v85 uint8 = 2
                                    _ = v85
                                } else {
                                    v1 = (fi80 >> (((v2 ^ v2)) & 63))
                                    v86 := GenS7{F0: 0, F1: -858, F2: (v4 % ((v4) | 1))}
                                    v4 += v4
                                    _ = v86
                                }
                                for fi87 := 0; fi87 < 8; fi87++ {
                                    var v88 uint64 = (v6 % ((4294967295) | 1))
                                    _ = fi87
                                    _ = v88
                                }
                                _ = v83
                            }
                        }
                        v89 := GenS11{F0: ((1 == 1) || false), F1: 903, F2: (v6 & 0)}
                        _ = fi80
                        _ = v81
                        _ = v89
                    }
                    v4 = -1
                }
            }
            for fi90 := 0; fi90 < 12; fi90++ {
                switch ("foo bar" + "foo bar") {
                case "":
                    if true {
                        v4 = -128
                    }
                default:
                    fmt.Println(v4, fi90, fi90)
                }
                v91 := GenS4{F0: 1}
                _ = fi90
                _ = v91
            }
            _ = v67
        } else {
            if (!true) {
                var v92 int16 = 32767
                _ = v92
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            var v93 uint64 = v6
            if (float64(0.0) <= (3.14159 + 3.14159)) {
                v1 -= (v1 ^ 42)
            } else {
                var v94 int16 = 0
                for fi95 := 0; fi95 < 2; fi95++ {
                    if false {
                        var v96 uint8 = 28
                        _ = v96
                    } else {
                        for fi97 := 0; fi97 < 7; fi97++ {
                            v98 := GenS12{F0: 65535}
                            v93 = 18446744073709551615
                            v99 := GenS10{F0: fi97, F1: 621, F2: 86}
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                        var v100 uint16 = uint16(v6)
                        switch (("foo bar" + "\n\t") + "foo bar") {
                        case "":
                            if len(v3) > 0 {
                                v3[0] = fi95
                            }
                        }
                        _ = v100
                    }
                    _ = fi95
                }
                _ = v94
            }
            _ = v93
        }
    }
    _ = len(v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 = append(v3, v1)
    v5 = append(v5, (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
