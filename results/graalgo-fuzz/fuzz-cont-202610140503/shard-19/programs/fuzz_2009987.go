package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint16; F1 int8; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 uint64 }
func genFn0(p0 int64, p1 int) int {
    var v6 uint = 42
    v7 := []int{p1, (p1 + p1), p1, (p1 ^ -897)}
    v8 := GenS1{F0: (2147483647 == uint32(v6)), F1: (p1 - (p1 << (((v6 % ((2) | 1))) & 63)))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 int = 128
    v2 := []uint32{}
    var v3 bool = (65535 > 1)
    var v4 int32 = -279
    var v5 int8 = 1
    switch (v1 / ((v1) | 1)) {
    case -127:
        if (-569 <= -278) {
            v9 := GenS1{F0: v3, F1: (v1 & v1)}
            v10 := GenS2{F0: (v5 ^ v5)}
            if len(v2) > 0 {
                v2[0] = 963
            }
            _ = v10
            _ = v9
        }
    default:
        v2 = append(v2, 42)
    }
    fmt.Println(v1, v5, v5)
    v1 += genFn0(1, genFn0(0, (v1 >> ((255) & 63))))
    if ((((v5 ^ v5) <= 123) || (false && v3)) || (v3 && false)) {
        var v11 int16 = -32768
        v12 := GenS2{F0: v5}
        _ = v11
        _ = v12
    } else {
        if ("hello" <= "hello") {
            if v3 {
                if (v3 || true) {
                    var v13 int64 = -771
                    v2 = append(v2, 0)
                    _ = v13
                } else {
                    v3 = (1 > 42)
                    v14 := GenS3{F0: 42, F1: v5, F2: 1}
                    _ = len(v2)
                    _ = v14
                }
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if (!true) {
            v3 = (!false)
            v5 = 42
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint = 866
                v2 = append(v2, 2)
                switch 42 {
                case 32768:
                    for fi17 := 0; fi17 < 17; fi17++ {
                        if true {
                            var v18 uint = 4294967295
                            fmt.Println(v16, v16, v3)
                            switch (("\n\t" + "a") + "a") {
                            case "a":
                                fmt.Println(v18)
                            case "":
                                if v3 {
                                    v3 = v3
                                    var v19 bool = ((fi15 & fi17) == fi15)
                                    _ = v19
                                } else {
                                    v5 = v5
                                }
                                switch v1 {
                                case -330:
                                    if len(v2) > 0 {
                                        v2[0] = 2147483648
                                    }
                                default:
                                    var v20 uint = 1
                                    _ = v20
                                }
                            default:
                                var v21 float64 = float64(1e10)
                                _ = v21
                            }
                            _ = v18
                        } else {
                            fmt.Println(v3)
                        }
                        v22 := GenS2{F0: v5}
                        _ = fi17
                        _ = v22
                    }
                case 4294967295:
                    for fi23 := 0; fi23 < 2; fi23++ {
                        var v24 int16 = 1
                        _ = len(v2)
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi15
                _ = v16
            }
        } else {
            for fi25 := 0; fi25 < 17; fi25++ {
                switch (v1 + (fi25 * 127)) {
                case 9223372036854775807:
                    v2 = append(v2, 0)
                case -673:
                    v26 := GenS3{F0: 1, F1: 1, F2: -128}
                    _ = v26
                }
                v27 := GenS3{F0: int16(v5), F1: v5, F2: (v5 | v5)}
                _ = fi25
                _ = v27
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if (v3 || (((v5 ^ v5) > (v5 >> ((4294967295) & 63))) && v3)) {
        var v28 float32 = 1.5
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = v28
    }
    v2 = append(v2, 2147483647)
    if (!((true || (!false)) && ((!(!true)) && v3))) {
        var v29 uint = 65535
        var v30 int32 = 0
        switch "foo bar" {
        case "hello":
            v4 += -2147483648
        case "a":
            v29 = 167
        default:
            if false {
                for fi31 := 0; fi31 < 10; fi31++ {
                    v32 := GenS1{F0: true, F1: (v1 - (fi31 << (((v29 >> ((669) & 63))) & 63)))}
                    v1 -= int(v5)
                    v30 += (v30 & v4)
                    _ = fi31
                    _ = v32
                }
                v5 = -128
            } else {
                var v33 uint = (v29 + (v29 - v29))
                fmt.Println(v30, v5, v29)
                _ = v33
            }
        }
        _ = v29
        _ = v30
    } else {
        fmt.Println(v4)
        if (float64((((-2.25 - 0.1) / float32(float64(0.1))) / (-2.25 + 1.5))) != 1.5) {
            _ = len(v2)
            v4 += (v4 - v4)
        } else {
            _ = len(v2)
        }
    }
    switch (v1 + 0) {
    case 2147483647:
        var v34 uint64 = 1
        _ = v34
    case 1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v5 += 100
    }
    v5 = (v5 - v5)
    fmt.Println(v1)
    switch 1 {
    case 91:
        var v35 float64 = 1e10
        _ = v35
    case 42:
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 uint8 = 194
            v3 = v3
            _ = fi36
            _ = v37
        }
    case 130:
        for fi38 := 0; fi38 < 6; fi38++ {
            if v3 {
                v39 := GenS2{F0: v5}
                if (58 < 545) {
                    switch "a" {
                    case "foo bar":
                        v40 := GenS4{F0: (v5 << ((4294967295) & 63)), F1: 18446744073709551615, F2: 2}
                        _ = v40
                    case "":
                        if len(v2) > 0 {
                            v2[0] = 366
                        }
                    }
                    for fi41 := 0; fi41 < 10; fi41++ {
                        if (-2.25 >= float32(0.1)) {
                            v3 = false
                        } else {
                            v1 += -2147483648
                            v42 := GenS5{F0: 32767, F1: (v5 * -128), F2: -9223372036854775808}
                            _ = v42
                        }
                        v1 -= (^(v1))
                        v5 = v5
                        _ = fi41
                    }
                } else {
                    var v43 float64 = float64(3.14159)
                    if v3 {
                        v3 = (true || (v3 && v3))
                        for fi44 := 0; fi44 < 8; fi44++ {
                            v43 += (v43 + v43)
                            v45 := GenS6{F0: 748, F1: 519}
                            _ = fi44
                            _ = v45
                        }
                        v1 -= 1
                    }
                    _ = v43
                }
                for fi46 := 0; fi46 < 13; fi46++ {
                    v1 = fi46
                    _ = fi46
                }
                _ = v39
            }
            v47 := GenS7{F0: 32768}
            switch "foo bar" {
            case "a":
                v1 = (v1 & 1)
            }
            _ = fi38
            _ = v47
        }
    default:
        if ((v1 % ((v1) | 1)) > v1) {
            var v48 uint16 = 972
            var v49 float64 = ((float64((float32(0.0) + (float32((0.1 * 1e10)) - 3.14159))) * -2.25) + 3.14159)
            if true {
                fmt.Println(v49)
                if v3 {
                    v1 = -9223372036854775808
                } else {
                    if false {
                        for fi50 := 0; fi50 < 8; fi50++ {
                            var v51 uint = 255
                            var v52 int64 = 931
                            switch "hello" {
                            case "\n\t":
                                v1 += (fi50 << ((91) & 63))
                            default:
                                v53 := GenS1{F0: (858 <= uint(v48)), F1: fi50}
                                _ = v53
                            }
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        if v3 {
                            switch ((("hello" + ("" + "foo bar")) + "\n\t") + "hello") {
                            case "hello":
                                v48 += v48
                            case "a":
                                v49 = (1.5 * v49)
                            }
                            fmt.Println(v49)
                            fmt.Println(v4)
                        } else {
                            if (2 >= v48) {
                                if true {
                                    for fi54 := 0; fi54 < 1; fi54++ {
                                        v49 = 0.1
                                        if len(v2) > 0 {
                                            v2[0] = 91
                                        }
                                        var v55 int8 = 100
                                        _ = fi54
                                        _ = v55
                                    }
                                    switch "hello" {
                                    case "":
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                    case "foo bar":
                                        var v56 uint = 2
                                        _ = v56
                                    }
                                    switch (v1 + v1) {
                                    case 317:
                                        switch v1 {
                                        case 0:
                                            v48 -= v48
                                        case -9223372036854775808:
                                            fmt.Println(v3, v5)
                                        }
                                    case 42:
                                        v2 = append(v2, 1)
                                    default:
                                        fmt.Println(v48)
                                    }
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = len(v2)
                                    for fi57 := 0; fi57 < 15; fi57++ {
                                        v49 = v49
                                        for fi58 := 0; fi58 < 2; fi58++ {
                                            var v59 int = (fi58 + v1)
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch 9223372036854775807 {
                                                case -9223372036854775808:
                                                    var v60 float64 = -2.25
                                                    _ = v60
                                                }
                                            }
                                            var v61 bool = (true || (2 <= 0))
                                            _ = fi58
                                            _ = v59
                                            _ = v61
                                        }
                                        var v62 uint = 65535
                                        _ = fi57
                                        _ = v62
                                    }
                                }
                            } else {
                                v5 -= (v5 - (v5 << ((18446744073709551615) & 63)))
                            }
                            if false {
                                switch (v1 ^ (v1 % ((9223372036854775807) | 1))) {
                                case 147:
                                    var v63 int = v1
                                    _ = v63
                                case 128:
                                    fmt.Println(v49, v4)
                                default:
                                    var v64 int32 = (v4 + (v4 ^ 1))
                                    _ = v64
                                }
                                v49 += v49
                            }
                            var v65 int32 = (v4 / ((539) | 1))
                            _ = v65
                        }
                    }
                    for fi66 := 0; fi66 < 7; fi66++ {
                        if false {
                            v67 := GenS8{F0: (fi66 >> ((2) & 63))}
                            switch "" {
                            case "":
                                var v68 uint32 = 0
                                _ = v68
                            }
                            v69 := GenS9{F0: -335}
                            _ = v67
                            _ = v69
                        }
                        var v70 uint8 = 128
                        v4 = (v4 ^ (-(v4)))
                        _ = fi66
                        _ = v70
                    }
                    v71 := GenS10{F0: (^(v4)), F1: v5}
                    _ = v71
                }
            }
            _ = v48
            _ = v49
        } else {
            if (v3 && (880 < 357)) {
                if (!true) {
                    switch v1 {
                    case 0:
                        for fi72 := 0; fi72 < 9; fi72++ {
                            v5 -= v5
                            if ((1 > 0) || (false || (v3 || v3))) {
                                var v73 float32 = -2.25
                                v3 = (!v3)
                                if v3 {
                                    v2 = append(v2, 2147483648)
                                    switch genFn0(42, fi72) {
                                    case 32767:
                                        v2 = append(v2, 4294967295)
                                    case 32768:
                                        var v74 float32 = v73
                                        _ = v74
                                    case 4294967295:
                                        _ = len(v2)
                                    }
                                }
                                _ = v73
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                            }
                            var v75 uint16 = 675
                            _ = fi72
                            _ = v75
                        }
                    case 814:
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                    }
                } else {
                    var v76 uint8 = 42
                    var v77 int16 = 461
                    if v3 {
                        v4 += -1
                        fmt.Println(v5)
                        var v78 float32 = 1e10
                        _ = v78
                    }
                    _ = v76
                    _ = v77
                }
                for fi79 := 0; fi79 < 1; fi79++ {
                    v5 = v5
                    _ = fi79
                }
            } else {
                for fi80 := 0; fi80 < 19; fi80++ {
                    if len(v2) > 0 {
                        v2[0] = 539
                    }
                    if ((v3 && (32767 <= 1)) && (1 < 934)) {
                        if v3 {
                            var v81 uint64 = 1
                            _ = v81
                        }
                        for fi82 := 0; fi82 < 20; fi82++ {
                            if v3 {
                                v83 := GenS4{F0: (v5 + v5), F1: 18446744073709551615, F2: 42}
                                _ = v83
                            }
                            if len(v2) > 0 {
                                v2[0] = 184
                            }
                            var v84 int64 = 1
                            _ = fi82
                            _ = v84
                        }
                    } else {
                        v4 -= (v4 << ((0) & 63))
                        if len(v2) > 0 {
                            v2[0] = 163
                        }
                        v85 := GenS6{F0: 887, F1: 4294967295}
                        _ = v85
                    }
                    for fi86 := 0; fi86 < 6; fi86++ {
                        v87 := GenS2{F0: (^(v5))}
                        v88 := GenS11{F0: 5}
                        _ = len(v2)
                        _ = fi86
                        _ = v87
                        _ = v88
                    }
                    _ = fi80
                }
            }
            if v3 {
                for fi89 := 0; fi89 < 20; fi89++ {
                    v1 -= (v1 * (fi89 ^ (v1 % (((v1 ^ genFn0(-835, (fi89 | (v1 ^ genFn0(9223372036854775807, v1)))))) | 1))))
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        for fi90 := 0; fi90 < 19; fi90++ {
                            v91 := GenS9{F0: -6}
                            switch v1 {
                            case 127:
                                v5 += (v5 / ((-128) | 1))
                            default:
                                var v92 float64 = 0.1
                                _ = v92
                            }
                            var v93 float32 = 1.5
                            _ = fi90
                            _ = v91
                            _ = v93
                        }
                    default:
                        v5 -= (v5 - v5)
                    }
                    _ = fi89
                }
            }
            v94 := GenS4{F0: -80, F1: 1, F2: 2147483647}
            _ = v94
        }
    }
    switch "hello" {
    case "a":
        if (v3 && v3) {
            for fi95 := 0; fi95 < 20; fi95++ {
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                v2 = append(v2, 68)
                _ = fi95
            }
            var v96 uint = 0
            _ = v96
        } else {
            _ = len(v2)
            v3 = v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "hello":
        if ((false || (!((v1 <= v1) && (!true)))) || false) {
            if len(v2) > 0 {
                v2[0] = 2
            }
            v3 = (!v3)
        }
    case "foo bar":
        fmt.Println(v3, v5)
    default:
        v97 := GenS1{F0: ((!(!(18446744073709551615 == 18446744073709551615))) && (v1 != -297)), F1: (v1 + (v1 * v1))}
        _ = v97
    }
    fmt.Println(v5)
    v98 := GenS9{F0: 303}
    var v99 uint16 = 783
    var v100 float32 = 0.1
    for fi101 := 0; fi101 < 11; fi101++ {
        for fi102 := 0; fi102 < 1; fi102++ {
            var v103 uint16 = (v99 % (((v99 + v99)) | 1))
            v2 = append(v2, 2)
            _ = fi102
            _ = v103
        }
        _ = fi101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v98
    _ = v99
}
