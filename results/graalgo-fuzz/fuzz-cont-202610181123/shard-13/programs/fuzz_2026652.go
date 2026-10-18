package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int64; F1 int; F2 uint16 }
type GenS4 struct { F0 int64; F1 int64; F2 uint }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint; F1 string; F2 bool }
type GenS12 struct { F0 int8; F1 int; F2 uint }
type GenS13 struct { F0 int8; F1 int32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int; F1 string }

func main() {
    var v1 uint16 = 255
    var v2 int = -1
    var v3 int = (v2 | (v2 + (v2 | v2)))
    var v4 uint = 226
    if false {
        v4 = (v4 / (((v4 ^ v4)) | 1))
        v5 := []int{}
        _ = v5
    }
    v6 := []int{(v2 | (v3 | v2)), (v2 * v2), v2, v2}
    v7 := GenS1{F0: "", F1: -1, F2: 1}
    var v8 int8 = 112
    if true {
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                if (0 >= 217) {
                    v1 = (v1 - uint16(v4))
                    var v11 int32 = 42
                    _ = v11
                } else {
                    var v12 int32 = 1
                    v12 -= v12
                    _ = len(v6)
                    switch "\n\t" {
                    case "\n\t":
                        var v13 bool = true
                        _ = v13
                    }
                    if true {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if false {
                            if len(v6) > 0 {
                                v6[0] = -700
                            }
                        }
                        v8 = -128
                        switch "a" {
                        case "\n\t":
                            v14 := GenS1{F0: (("a" + ("a" + "hello")) + "hello"), F1: -1, F2: 1000}
                            v15 := GenS2{F0: true, F1: int(v8)}
                            _ = v14
                            _ = v15
                        default:
                            v1 = (v1 - (v1 ^ v1))
                        }
                    } else {
                        fmt.Println(v1, v2, fi10)
                        var v16 int = fi10
                        _ = v16
                    }
                    fmt.Println(v2, v1, v2)
                    v2 = (v2 >> ((v4) & 63))
                    _ = v12
                }
                _ = fi10
            }
            v17 := GenS3{F0: -1, F1: -85, F2: (v1 | 42)}
            v3 -= v2
            _ = fi9
            _ = v17
        }
        v18 := GenS3{F0: 9223372036854775807, F1: v3, F2: 183}
        for fi19 := 0; fi19 < 8; fi19++ {
            switch (("hello" + "hello") + "foo bar") {
            case "a":
                if (-246 < -2147483648) {
                    v20 := GenS4{F0: 364, F1: -9223372036854775808, F2: v4}
                    v21 := GenS4{F0: 9223372036854775807, F1: -9223372036854775808, F2: 2}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v20
                    _ = v21
                } else {
                    fmt.Println(v1, v8)
                    var v22 int = 678
                    _ = v22
                }
            }
            v23 := GenS5{F0: v3, F1: 1}
            if ("" >= ("\n\t" + "a")) {
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 uint16 = v1
                    fmt.Println(fi19)
                    switch fi24 {
                    case 4294967295:
                        fmt.Println(v25, fi19)
                    case -2147483648:
                        v26 := GenS6{F0: 0, F1: 42}
                        _ = v26
                    case 128:
                        var v27 uint64 = 0
                        if (!false) {
                            v28 := GenS2{F0: (194 > 2), F1: fi24}
                            _ = len(v6)
                            _ = v28
                        } else {
                            fmt.Println(v2)
                        }
                        v6 = append(v6, (^(fi19)))
                        if false {
                            var v29 bool = true
                            if v29 {
                                v25 += (v1 >> ((2) & 63))
                            } else {
                                var v30 int16 = -1
                                switch "\n\t" {
                                case "\n\t":
                                    v31 := GenS7{F0: false, F1: 836}
                                    switch "foo bar" {
                                    case "hello":
                                        v1 += v1
                                    case "foo bar":
                                        switch ("\n\t" + ("\n\t" + ((("hello" + (("hello" + ("hello" + ("" + "foo bar"))) + ("" + ("" + "\n\t")))) + (("foo bar" + "hello") + "foo bar")) + "a"))) {
                                        case "\n\t":
                                            if v29 {
                                                switch "" {
                                                case "hello":
                                                    v3 = v3
                                                }
                                                fmt.Println(v25, fi24, v3)
                                            }
                                        case "hello":
                                            v32 := GenS8{F0: v29}
                                            _ = v32
                                        case "foo bar":
                                            switch "" {
                                            case "":
                                                v30 -= (v30 - (v30 >> ((v4) & 63)))
                                            case "hello":
                                                v4 = (v4 - 255)
                                            }
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (v29 && (!true)) {
                                                fmt.Println(v27, v4, v1)
                                                if v29 {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v33 := GenS9{F0: 237, F1: 255, F2: -192}
                                                        _ = v33
                                                    default:
                                                        switch fi19 {
                                                        case 1:
                                                            if v29 {
                                                                if ((!(4294967295 > 529)) && (v1 < (^(v25)))) {
                                                                    var v34 int64 = -9223372036854775808
                                                                    var v35 bool = v29
                                                                    if len(v6) > 0 {
                                                                        v6[0] = (v3 & v2)
                                                                    }
                                                                    _ = v34
                                                                    _ = v35
                                                                } else {
                                                                    if true {
                                                                        v1 = uint16(v4)
                                                                    }
                                                                    v3 -= 1
                                                                }
                                                            }
                                                        case 0:
                                                            if ((v27 < v27) || v29) {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                var v36 int32 = int32(v1)
                                                                _ = v36
                                                            }
                                                        case -683:
                                                            v27 = (v27 << ((uint(v1)) & 63))
                                                        default:
                                                            if true {
                                                                v2 += fi24
                                                                v1 += 42
                                                            } else {
                                                                v27 = (v27 >> ((v4) & 63))
                                                                switch -25 {
                                                                case -2147483648:
                                                                    var v37 string = ("" + "foo bar")
                                                                    _ = v37
                                                                case -472:
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                case 97:
                                                                    if v29 {
                                                                        if (true && (!(((v4 << ((v4) & 63)) < v4) || false))) {
                                                                            var v38 uint = 18446744073709551615
                                                                            if v29 {
                                                                                var v39 uint = 65535
                                                                                v6 = append(v6, v2)
                                                                                fmt.Println(v38, v39, fi24)
                                                                                var v40 bool = v29
                                                                                _ = v39
                                                                                _ = v40
                                                                            }
                                                                            v41 := GenS6{F0: 2, F1: -2147483648}
                                                                            _ = v38
                                                                            _ = v41
                                                                        }
                                                                        _ = len(v6)
                                                                        v42 := GenS10{F0: 0}
                                                                        _ = v42
                                                                    } else {
                                                                        v25 = v25
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(fi24, v3)
                                                    if v29 {
                                                        if false {
                                                            v2 = (fi19 & 42)
                                                            var v43 uint8 = 2
                                                            var v44 uint64 = (^(v27))
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            var v45 int = (fi24 & 840)
                                                            _ = v45
                                                        }
                                                        fmt.Println(v30)
                                                    }
                                                }
                                                if (75 == (-(v30))) {
                                                    if (!v29) {
                                                        if ((1e10 / (0.1 / ((1.5 + 0.0) / ((float64((1.5 * float32(1e10))) + 0.1) * ((0.1 - (float64(1e10) * (3.14159 / 1.5))) + 1e10))))) > 0.1) {
                                                            if ((282 != 1) || v29) {
                                                                fmt.Println(v25)
                                                                var v46 int16 = 0
                                                                _ = v46
                                                            } else {
                                                                var v47 uint8 = 108
                                                                v48 := GenS11{F0: 255, F1: ("\n\t" + "a"), F2: (-374 < -850)}
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            _ = len(v6)
                                                        }
                                                        if true {
                                                            _ = len(v6)
                                                        } else {
                                                            if v29 {
                                                                v25 -= (v1 * 73)
                                                                var v49 int32 = 2147483647
                                                                _ = v49
                                                            }
                                                            v50 := GenS7{F0: v29, F1: (v25 ^ v25)}
                                                            _ = v50
                                                        }
                                                    } else {
                                                        if v29 {
                                                            if len(v6) > 0 {
                                                                v6[0] = fi19
                                                            }
                                                        }
                                                        var v51 int16 = v30
                                                        _ = v51
                                                    }
                                                    var v52 uint = (v4 - v4)
                                                    _ = v52
                                                } else {
                                                    if v29 {
                                                        if v29 {
                                                            fmt.Println(fi19)
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "":
                                                        if v29 {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        }
                                                    default:
                                                        var v53 int32 = -2147483648
                                                        _ = v53
                                                    }
                                                }
                                                if ((v8 | (v8 + v8)) != (v8 / ((-14) | 1))) {
                                                    if v29 {
                                                        v25 += v25
                                                        v25 = (v1 << (((v4 / (((v4 / (((v4 ^ 462)) | 1))) | 1))) & 63))
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = -380
                                                    }
                                                }
                                                switch v3 {
                                                case 127:
                                                    if v29 {
                                                        var v54 int64 = 0
                                                        fmt.Println(v27, v4)
                                                        _ = v54
                                                    }
                                                case -176:
                                                    if len(v6) > 0 {
                                                        v6[0] = v3
                                                    }
                                                case 1:
                                                    var v55 int32 = 975
                                                    _ = v55
                                                default:
                                                    switch "" {
                                                    case "hello":
                                                        fmt.Println(v29, v4)
                                                    case "foo bar":
                                                        v6 = append(v6, fi19)
                                                    default:
                                                        v56 := GenS12{F0: (v8 + (v8 % (((v8 ^ -40)) | 1))), F1: (^(v3)), F2: 65535}
                                                        _ = v56
                                                    }
                                                }
                                            } else {
                                                if (float64(0.1) != 1.5) {
                                                    var v57 float32 = (3.14159 + (-2.25 / 1.5))
                                                    var v58 uint8 = 255
                                                    v6 = append(v6, 4294967295)
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    v1 = (v25 | (v1 | (v1 & (v25 >> (((v4 >> ((65535) & 63))) & 63)))))
                                                    v6 = append(v6, (-(v3)))
                                                }
                                            }
                                        case "a":
                                            _ = len(v6)
                                        case "hello":
                                            if (v29 && (true || v29)) {
                                                v30 = 0
                                                var v59 uint64 = (v27 * v27)
                                                v60 := GenS13{F0: int8(v25), F1: 1}
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                    }
                                    if false {
                                        var v61 int64 = 1
                                        var v62 uint64 = (v27 ^ v27)
                                        var v63 uint8 = 244
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v31
                                case "a":
                                    v6 = append(v6, (fi19 * v2))
                                case "hello":
                                    if (((((("hello" + "") + ("a" + ("foo bar" + "\n\t"))) + ((("hello" + "a") + "hello") + ("hello" + ("" + "foo bar")))) + "a") + ("hello" + "foo bar")) >= ("foo bar" + (("" + "") + ("hello" + ("\n\t" + "foo bar"))))) {
                                        v64 := GenS13{F0: 94, F1: -1}
                                        v29 = ((206 >= 42) || true)
                                        _ = v64
                                    }
                                default:
                                    var v65 int32 = 42
                                    _ = v65
                                }
                                v66 := GenS14{F0: v4}
                                _ = v30
                                _ = v66
                            }
                            _ = v29
                        }
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
                switch (("hello" + "hello") + "\n\t") {
                case "":
                    if len(v6) > 0 {
                        v6[0] = 845
                    }
                case "hello":
                    fmt.Println(v3)
                default:
                    for fi67 := 0; fi67 < 12; fi67++ {
                        v8 = v8
                        var v68 uint64 = 42
                        _ = fi67
                        _ = v68
                    }
                }
                v8 = (v8 | (v8 >> ((v4) & 63)))
            } else {
                v69 := GenS15{F0: v8}
                switch (v3 & v2) {
                case 42:
                    switch (-(fi19)) {
                    case 32768:
                        var v70 float64 = float64(0.0)
                        _ = v70
                    default:
                        var v71 uint8 = 0
                        _ = v71
                    }
                case 4294967295:
                    if len(v6) > 0 {
                        v6[0] = v3
                    }
                case 32768:
                    v1 += v1
                default:
                    for fi72 := 0; fi72 < 19; fi72++ {
                        var v73 uint64 = 681
                        _ = len(v6)
                        _ = fi72
                        _ = v73
                    }
                }
                fmt.Println(v2)
                _ = v69
            }
            _ = fi19
            _ = v23
        }
        _ = v18
    } else {
        v4 = v4
        if ((-2.25 * 3.14159) < 1.5) {
            for fi74 := 0; fi74 < 9; fi74++ {
                fmt.Println(fi74, v2)
                var v75 bool = true
                if (4294967295 >= 1) {
                    for fi76 := 0; fi76 < 19; fi76++ {
                        fmt.Println(fi76, v3, fi76)
                        _ = fi76
                    }
                    _ = len(v6)
                } else {
                    v6 = append(v6, (^(v2)))
                    var v77 uint8 = 0
                    _ = v77
                }
                _ = fi74
                _ = v75
            }
            v1 -= v1
            if len(v6) > 0 {
                v6[0] = (v3 / (((^(v3))) | 1))
            }
        }
    }
    v1 = v1
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        var v78 uint16 = 42
        _ = v78
    case "a":
        v8 = -99
    }
    switch v3 {
    case 410:
        v2 = (v3 % (((^(v2))) | 1))
    case -67:
        for fi79 := 0; fi79 < 20; fi79++ {
            _ = len(v6)
            v8 = (v8 << (((v4 - v4)) & 63))
            _ = fi79
        }
    default:
        v80 := GenS13{F0: v8, F1: -1}
        _ = v80
    }
    fmt.Println(v8, v8)
    var v81 uint = 4294967295
    switch int(v8) {
    case 151:
        for fi82 := 0; fi82 < 9; fi82++ {
            var v83 uint16 = v1
            for fi84 := 0; fi84 < 20; fi84++ {
                var v85 float32 = (-2.25 / 1.5)
                _ = fi84
                _ = v85
            }
            _ = fi82
            _ = v83
        }
    case -1:
        _ = len(v6)
    case 153:
        var v86 uint16 = (v1 >> ((18446744073709551615) & 63))
        _ = v86
    }
    v87 := GenS3{F0: -249, F1: v2, F2: 65535}
    var v88 int32 = -1
    var v89 float64 = 1.5
    for fi90 := 0; fi90 < 14; fi90++ {
        for fi91 := 0; fi91 < 13; fi91++ {
            switch (fi91 >> (((v4 ^ v81)) & 63)) {
            case 9223372036854775807:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 32767:
                if false {
                    v3 += 228
                    if false {
                        v8 += (v8 | 1)
                        if len(v6) > 0 {
                            v6[0] = (v3 / ((-9223372036854775808) | 1))
                        }
                    } else {
                        v2 = (fi91 / ((-100) | 1))
                    }
                } else {
                    v1 -= 132
                }
                if (128 > int(v1)) {
                    var v92 int8 = v8
                    fmt.Println(v1, v92)
                    _ = v92
                }
                v3 = (^(fi91))
                v93 := GenS8{F0: (42 < 1)}
                if true {
                    v88 -= v88
                } else {
                    v94 := GenS16{F0: -663, F1: ("" + ("a" + ("foo bar" + "\n\t")))}
                    _ = v94
                }
                var v95 float64 = v89
                _ = v93
                _ = v95
            case 0:
                fmt.Println(v4, v88, v1)
            }
            var v96 int8 = (v8 * (^(v8)))
            _ = fi91
            _ = v96
        }
        v97 := GenS5{F0: v3, F1: 515}
        fmt.Println(v1)
        _ = fi90
        _ = v97
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v81:", v81)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v81
    _ = v87
    _ = v88
    _ = v89
}
