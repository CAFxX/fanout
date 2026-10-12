package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 int64 }
type GenS6 struct { F0 uint; F1 uint16; F2 string }
func genFn0(p0 int, p1 bool, p2 bool) int32 {
    var v6 int8 = 127
    var v7 uint = 255
    p0 = p0
    if ((!p2) && (int(v7) >= (p0 | (p0 << ((18446744073709551615) & 63))))) {
        if p2 {
            v7 -= (^(v7))
            p0 -= 2147483647
            v6 -= (v6 / (((^(v6))) | 1))
        }
        v7 += (v7 >> (((^(v7))) & 63))
    } else {
        p1 = (42 == 1)
        p2 = (false && p2)
    }
    switch (p0 >> ((v7) & 63)) {
    case 4294967295:
        var v8 uint16 = 254
        _ = v8
    case 128:
        v7 = (v7 % (((v7 / (((v7 - v7)) | 1))) | 1))
    case -1:
        if (uint8(v7) >= 255) {
            var v9 uint32 = 466
            _ = v9
        } else {
            var v10 float64 = 3.14159
            v11 := []int{}
            p1 = (v6 != (v6 | v6))
            _ = v10
            _ = v11
        }
    default:
        var v12 bool = ((^(p0)) < p0)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return int32(p0)
}
func genFn1(p0 int16) uint16 {
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int = 124
        _ = fi13
        _ = v14
    }
    if true {
        switch 32768 {
        case -305:
            v15 := GenS1{F0: false, F1: 0, F2: "hello"}
            _ = v15
        case 32768:
            fmt.Println(p0)
        }
        p0 = 362
        fmt.Println(p0)
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := []int{}
        if (0 < 255) {
            for fi18 := 0; fi18 < 15; fi18++ {
                switch (fi16 % (((fi18 * fi18)) | 1)) {
                case 32767:
                    fmt.Println(p0, fi18, fi16)
                }
                _ = fi18
            }
            p0 = -557
            var v19 int64 = 9223372036854775807
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    var v20 bool = (!false)
    fmt.Println(p0)
    _ = p0
    _ = v20
    return 353
}

func main() {
    var v1 uint8 = 53
    var v2 int64 = 1
    v3 := []string{("\n\t" + "a")}
    var v4 int16 = 385
    var v5 int64 = v2
    fmt.Println(v1)
    for fi21 := 0; fi21 < 16; fi21++ {
        if true {
            fmt.Println(v4)
        } else {
            switch fi21 {
            case -1:
                switch (^(fi21)) {
                case 4294967295:
                    fmt.Println(fi21, fi21)
                case 683:
                    fmt.Println(v2, v1)
                case -9223372036854775808:
                    var v22 uint64 = 234
                    _ = v22
                }
            case 0:
                fmt.Println(v2, v4, v2)
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int64 = v5
                v3 = append(v3, "\n\t")
                var v25 uint8 = (v1 >> ((2) & 63))
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi26
            }
        }
        v1 -= 31
        switch "foo bar" {
        case "":
            for fi27 := 0; fi27 < 14; fi27++ {
                v1 = v1
                if (42 == v2) {
                    switch "foo bar" {
                    case "\n\t":
                        var v28 string = "a"
                        _ = v28
                    }
                }
                _ = len(v3)
                _ = fi27
            }
        }
        _ = fi21
    }
    var v29 uint64 = 692
    for fi30 := 0; fi30 < 4; fi30++ {
        switch fi30 {
        case 1:
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 int64 = -9223372036854775808
                v29 += (^(v29))
                _ = fi31
                _ = v32
            }
        }
        v2 -= -481
        switch (fi30 << ((4294967295) & 63)) {
        case 851:
            v3 = append(v3, ("" + ""))
        default:
            var v33 int16 = (v4 - v4)
            _ = v33
        }
        _ = fi30
    }
    switch "a" {
    case "":
        if false {
            for fi34 := 0; fi34 < 18; fi34++ {
                for fi35 := 0; fi35 < 5; fi35++ {
                    var v36 uint8 = (v1 / (((v1 << ((823) & 63))) | 1))
                    v1 += (v36 & (v1 * v1))
                    _ = fi35
                    _ = v36
                }
                if ((v29 * v29) == 0) {
                    v4 = v4
                    switch (fi34 & (fi34 & fi34)) {
                    case 2147483647:
                        if false {
                            var v37 string = ""
                            _ = v37
                        }
                    case 158:
                        switch fi34 {
                        case 2147483647:
                            v29 += v29
                        case -2147483648:
                            var v38 bool = (2147483647 != -2147483648)
                            _ = v38
                        default:
                            v2 += -603
                        }
                    case 128:
                        fmt.Println(v2, v4)
                    }
                } else {
                    v39 := GenS2{F0: v1}
                    switch ("foo bar" + ("foo bar" + ("foo bar" + ""))) {
                    case "":
                        fmt.Println(v2, v29)
                    }
                    _ = v39
                }
                _ = fi34
            }
            var v40 int16 = int16(v29)
            v1 = v1
            _ = v40
        } else {
            switch 128 {
            case 9223372036854775807:
                for fi41 := 0; fi41 < 8; fi41++ {
                    for fi42 := 0; fi42 < 2; fi42++ {
                        var v43 uint16 = 2
                        _ = len(v3)
                        _ = fi42
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 4; fi44++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi44
                    }
                    v2 = 0
                    _ = fi41
                }
            }
        }
    case "a":
        var v45 bool = true
        _ = v45
    }
    v1 -= 2
    switch 42 {
    case 2147483647:
        v1 += v1
    }
    if (((-(v2)) == v2) || false) {
        if (true || (false || ((871 != 65535) && true))) {
            var v46 float64 = 0.1
            _ = v46
        }
        for fi47 := 0; fi47 < 4; fi47++ {
            v4 -= 1000
            for fi48 := 0; fi48 < 6; fi48++ {
                v29 = (v29 | v29)
                if (v29 != (v29 & 4294967295)) {
                    var v49 int8 = 62
                    v50 := GenS3{F0: 65535, F1: v1}
                    fmt.Println(v29, v1)
                    _ = v49
                    _ = v50
                } else {
                    fmt.Println(fi48, v1, fi47)
                    var v51 int64 = -9223372036854775808
                    v5 += v51
                    _ = v51
                }
                v29 += (v29 & v29)
                switch "foo bar" {
                case "a":
                    var v52 int32 = int32(v29)
                    v5 += -1
                    switch (fi48 & (-(fi48))) {
                    case 393:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 363:
                        fmt.Println(v29, fi47, v4)
                    case 429:
                        var v53 int8 = -1
                        _ = v53
                    }
                    _ = v52
                case "foo bar":
                    switch int(v4) {
                    case -321:
                        v29 -= v29
                        v4 -= v4
                    case -2147483648:
                        v1 = 2
                    default:
                        fmt.Println(fi48, fi48)
                        if false {
                            if (2 <= 2147483648) {
                                switch (("" + "a") + "a") {
                                case "hello":
                                    v29 += v29
                                default:
                                    switch 711 {
                                    case 0:
                                        var v54 int64 = -9223372036854775808
                                        _ = v54
                                    case 489:
                                        fmt.Println(v5, v5)
                                    case 42:
                                        var v55 string = "\n\t"
                                        _ = v55
                                    default:
                                        var v56 uint = 0
                                        _ = v56
                                    }
                                }
                                v1 -= (v1 + (v1 ^ v1))
                                fmt.Println(v1, v5)
                            }
                            if (!true) {
                                switch (fi47 - 32768) {
                                case 4294967295:
                                    v2 -= v5
                                case 989:
                                    v4 = 412
                                }
                                v2 = (v5 + v5)
                                var v57 bool = false
                                _ = v57
                            } else {
                                var v58 uint32 = 1
                                _ = len(v3)
                                _ = v58
                            }
                            var v59 uint = 2
                            _ = v59
                        }
                        if true {
                            switch fi47 {
                            case 128:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 0:
                                v2 = v5
                            }
                            v4 = v4
                            v1 = (v1 ^ 183)
                        }
                    }
                case "":
                    switch "foo bar" {
                    case "a":
                        v2 -= (v5 * (v5 / ((v2) | 1)))
                    default:
                        v2 -= 840
                    }
                }
                switch 127 {
                case -1:
                    var v60 uint = 398
                    _ = v60
                case 42:
                    _ = len(v3)
                case 1:
                    fmt.Println(v2, fi48, v29)
                default:
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                }
                fmt.Println(v4, fi48, v1)
                _ = fi48
            }
            _ = fi47
        }
    } else {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            if ((0.1 / (1e10 - -2.25)) >= 1e10) {
                v4 -= (v4 / ((1) | 1))
            }
        case "hello":
            v2 = -807
        case "":
            v61 := GenS1{F0: true, F1: 0, F2: "foo bar"}
            _ = v61
        default:
            var v62 int16 = v4
            _ = v62
        }
        v63 := GenS4{F0: 128}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v63
    }
    if true {
        v64 := GenS2{F0: v1}
        for fi65 := 0; fi65 < 6; fi65++ {
            fmt.Println(v1, v1)
            var v66 int64 = -1
            _ = fi65
            _ = v66
        }
        _ = v64
    }
    var v67 int = 791
    var v68 int64 = (v5 & (v2 + (v5 * v5)))
    v3 = append(v3, "\n\t")
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if (19 <= 65535) {
        v68 -= v2
    } else {
        if false {
            switch 895 {
            case 4294967295:
                v69 := GenS5{F0: 18446744073709551615, F1: (v4 & 1000), F2: 0}
                _ = v69
            case 0:
                var v70 int8 = -7
                _ = v70
            case 543:
                v1 += (v1 % (((^(v1))) | 1))
            }
            v4 += v4
        }
        _ = len(v3)
        v71 := GenS6{F0: 0, F1: 750, F2: ("foo bar" + "\n\t")}
        _ = v71
    }
    if true {
        var v72 float64 = -2.25
        v68 += -9223372036854775808
        if (v68 != (-(v5))) {
            v73 := GenS4{F0: 200}
            for fi74 := 0; fi74 < 19; fi74++ {
                fmt.Println(v1, v2)
                _ = fi74
            }
            _ = v73
        } else {
            fmt.Println(v1, v4, v2)
            v75 := GenS5{F0: (v29 | 977), F1: v4, F2: 0}
            _ = v75
        }
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
}
