package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 bool }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int64 }

func main() {
    var v1 float32 = 3.14159
    var v2 int16 = -32768
    var v3 uint8 = 2
    var v4 int = -709
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 11; fi6++ {
        v1 = -2.25
        v1 = (3.14159 - v1)
        _ = fi6
    }
    v4 = (v4 - -2147483648)
    var v7 float32 = 3.14159
    var v8 int16 = v2
    if (false || (0.0 > ((0.0 + 0.0) * 0.1))) {
        var v9 int32 = 2147483647
        v10 := GenS1{F0: 42, F1: 507, F2: 1}
        v1 = ((((v1 - ((0.0 / v1) / (((v1 + v1) - (1e10 - v7)) + 0.1))) * v7) - (1.5 - float32(((0.0 / float64(v1)) * (1e10 - 0.1))))) - 3.14159)
        _ = v10
        _ = v9
    } else {
        v3 = 127
        if true {
            v8 = (^(v8))
            switch ((("" + "") + "a") + "a") {
            case "":
                for fi11 := 0; fi11 < 8; fi11++ {
                    fmt.Println(v8)
                    _ = fi11
                }
            case "hello":
                if ((0.1 + (((1e10 - 0.1) + 1.5) - 3.14159)) <= -2.25) {
                    v3 = v5
                }
            default:
                var v12 int16 = v2
                _ = v12
            }
        } else {
            v13 := []int{(-(v4))}
            v7 = ((v7 - (v7 * (-2.25 * v1))) / v7)
            _ = v13
        }
        switch (v4 >> ((0) & 63)) {
        case -1:
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := []int{fi14}
                _ = fi14
                _ = v15
            }
        }
    }
    if ((42 > 251) && false) {
        if (763 != -1) {
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int = (-(fi16))
                _ = fi16
                _ = v17
            }
            var v18 bool = false
            v19 := GenS2{F0: 2, F1: 305, F2: 452}
            _ = v18
            _ = v19
        }
        v20 := GenS3{F0: (v4 + 880), F1: 2, F2: (18446744073709551615 != 0)}
        fmt.Println(v1)
        _ = v20
    } else {
        v7 -= 1e10
        v1 += 1e10
    }
    switch v4 {
    case -2147483648:
        var v21 int8 = -1
        _ = v21
    case -1:
        var v22 string = ""
        _ = v22
    case -423:
        v23 := []int{32767, (v4 * (v4 + v4)), (v4 * v4), (v4 + (v4 / ((123) | 1)))}
        _ = v23
    default:
        if false {
            var v24 string = ("" + ("hello" + ((("hello" + "") + "\n\t") + ((("\n\t" + (("hello" + "foo bar") + "hello")) + "\n\t") + "\n\t"))))
            if (v4 == 32767) {
                v4 += (-(v4))
                if (!true) {
                    v25 := GenS4{F0: 42, F1: (v24 + v24)}
                    var v26 uint8 = (v5 & v5)
                    _ = v25
                    _ = v26
                }
            } else {
                v27 := []int{9223372036854775807}
                _ = v27
            }
            v1 = v7
            _ = v24
        } else {
            var v28 uint = 255
            v29 := []int{(v4 - -1), v4, v4, (v4 << ((v28) & 63))}
            if false {
                for fi30 := 0; fi30 < 19; fi30++ {
                    v31 := GenS1{F0: -32768, F1: 2, F2: v5}
                    _ = fi30
                    _ = v31
                }
                _ = len(v29)
            }
            _ = v28
            _ = v29
        }
    }
    v32 := []int{9223372036854775807, v4, v4}
    switch 9223372036854775807 {
    case 32767:
        var v33 uint = 2
        _ = v33
    case -218:
        if ((("\n\t" + "foo bar") + "") != "a") {
            switch "\n\t" {
            case "\n\t":
                for fi34 := 0; fi34 < 19; fi34++ {
                    var v35 uint32 = uint32(v5)
                    fmt.Println(v8, v3, v35)
                    if (false || true) {
                        var v36 int16 = (v8 - (v2 ^ v8))
                        for fi37 := 0; fi37 < 8; fi37++ {
                            var v38 float64 = (1e10 + 0.0)
                            v2 -= -1
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                            _ = fi37
                            _ = v38
                        }
                        var v39 int64 = 928
                        _ = v36
                        _ = v39
                    } else {
                        switch int(v5) {
                        case -969:
                            var v40 int16 = v2
                            _ = v40
                        default:
                            fmt.Println(fi34, v7)
                        }
                        v32 = append(v32, (fi34 - fi34))
                        var v41 int32 = 0
                        _ = v41
                    }
                    _ = fi34
                    _ = v35
                }
            }
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            v4 -= -855
        } else {
            for fi42 := 0; fi42 < 5; fi42++ {
                v43 := GenS5{F0: 356}
                v4 = fi42
                v44 := GenS6{F0: 0, F1: (v2 * (v8 / ((v2) | 1)))}
                _ = fi42
                _ = v43
                _ = v44
            }
            fmt.Println(v8)
            switch "\n\t" {
            case "a":
                for fi45 := 0; fi45 < 20; fi45++ {
                    if (false || true) {
                        var v46 float32 = 1.5
                        for fi47 := 0; fi47 < 17; fi47++ {
                            var v48 uint = 42
                            _ = fi47
                            _ = v48
                        }
                        v3 = (v5 << ((988) & 63))
                        _ = v46
                    } else {
                        switch (fi45 % ((fi45) | 1)) {
                        case 2147483647:
                            var v49 uint64 = 2
                            _ = v49
                        case 127:
                            v32 = append(v32, 1)
                        }
                    }
                    if true {
                        var v50 bool = true
                        _ = v50
                    }
                    switch (fi45 - v4) {
                    case 127:
                        var v51 uint8 = (^(v5))
                        _ = v51
                    default:
                        v52 := GenS7{F0: 147}
                        _ = v52
                    }
                    _ = fi45
                }
            case "\n\t":
                v1 -= 3.14159
            default:
                v53 := GenS5{F0: -1}
                _ = v53
            }
        }
    case 519:
        var v54 int16 = v8
        _ = v54
    default:
        if len(v32) > 0 {
            v32[0] = (v4 >> ((164) & 63))
        }
    }
    v55 := GenS7{F0: 1}
    var v56 bool = (!false)
    fmt.Println(v4)
    fmt.Println(v3, v5)
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    if true {
        if false {
            switch (^(v4)) {
            case 1:
                v57 := GenS3{F0: v4, F1: 255, F2: (!v56)}
                _ = v57
            case -1:
                if ((float64(v7) - 0.0) == ((3.14159 * ((1e10 + 1e10) - 1.5)) / (1e10 / 0.1))) {
                    v5 += v3
                    if v56 {
                        v58 := GenS7{F0: 0}
                        if (true || (1 >= (v2 - v2))) {
                            if len(v32) > 0 {
                                v32[0] = v4
                            }
                        }
                        fmt.Println(v4)
                        _ = v58
                    }
                    v2 -= v2
                }
            default:
                v5 += (v3 * (v5 >> ((4294967295) & 63)))
            }
            if v56 {
                var v59 uint64 = 4294967295
                _ = v59
            } else {
                v8 -= (v8 / (((v2 & -148)) | 1))
            }
        } else {
            v8 = -32768
            if true {
                switch -1 {
                case 1:
                    for fi60 := 0; fi60 < 13; fi60++ {
                        v32 = append(v32, fi60)
                        v32 = append(v32, -653)
                        fmt.Println(v3)
                        _ = fi60
                    }
                case 128:
                    switch (("\n\t" + "hello") + ("\n\t" + ("\n\t" + ("" + "\n\t")))) {
                    case "a":
                        v4 = (v4 & (^(v4)))
                    case "\n\t":
                        for fi61 := 0; fi61 < 13; fi61++ {
                            _ = len(v32)
                            _ = fi61
                        }
                    case "hello":
                        var v62 uint = 652
                        _ = v62
                    }
                default:
                    var v63 float64 = 1e10
                    _ = v63
                }
                switch 0 {
                case 42:
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        for fi64 := 0; fi64 < 13; fi64++ {
                            fmt.Println(v2)
                            v4 = 32767
                            if len(v32) > 0 {
                                v32[0] = (fi64 ^ (v4 | (fi64 + fi64)))
                            }
                            _ = fi64
                        }
                    case "\n\t":
                        if v56 {
                            if false {
                                if (-2147483648 != -1) {
                                    v4 -= 2147483647
                                } else {
                                    fmt.Println(v56)
                                }
                                var v65 uint16 = 0
                                _ = v65
                            }
                            var v66 uint64 = 469
                            _ = v66
                        } else {
                            var v67 float32 = 1.5
                            v3 = (^(v5))
                            for fi68 := 0; fi68 < 7; fi68++ {
                                v32 = append(v32, -642)
                                v1 = (0.0 - (-2.25 * (((v7 * 1.5) / 3.14159) + float32(3.14159))))
                                _ = fi68
                            }
                            _ = v67
                        }
                    case "a":
                        switch (v4 | v4) {
                        case 0:
                            _ = len(v32)
                        case 294:
                            fmt.Println(v3, v8)
                        case 127:
                            if false {
                                v5 -= v5
                                v69 := GenS8{F0: 0}
                                _ = v69
                            } else {
                                if true {
                                    for fi70 := 0; fi70 < 19; fi70++ {
                                        v5 += v3
                                        fmt.Println(v56)
                                        _ = fi70
                                    }
                                    if (-128 <= 0) {
                                        v3 -= v3
                                        if (-128 <= 35) {
                                            fmt.Println(v5, v7)
                                        }
                                        v71 := GenS6{F0: 255, F1: v8}
                                        _ = v71
                                    } else {
                                        for fi72 := 0; fi72 < 2; fi72++ {
                                            if len(v32) >= 2 {
                                                v32 = v32[1:len(v32)]
                                            }
                                            _ = fi72
                                        }
                                        var v73 int8 = 1
                                        var v74 uint = 255
                                        _ = v73
                                        _ = v74
                                    }
                                    v8 = (v2 >> ((42) & 63))
                                }
                                fmt.Println(v1, v4)
                            }
                        }
                    default:
                        v3 -= v5
                    }
                case 4294967295:
                    switch "hello" {
                    case "":
                        var v75 float32 = v7
                        _ = v75
                    case "hello":
                        var v76 uint16 = 32767
                        _ = v76
                    case "a":
                        fmt.Println(v2, v1)
                    }
                }
                if true {
                    var v77 int = (v4 ^ (v4 & (-(v4))))
                    for fi78 := 0; fi78 < 11; fi78++ {
                        for fi79 := 0; fi79 < 20; fi79++ {
                            fmt.Println(v5, v4)
                            _ = len(v32)
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                            _ = fi79
                        }
                        for fi80 := 0; fi80 < 19; fi80++ {
                            if false {
                                var v81 float32 = (v7 + (v1 / float32(3.14159)))
                                _ = v81
                            }
                            var v82 float32 = float32(float64(((float32((1.5 + 0.0)) * -2.25) - v1)))
                            var v83 int16 = v2
                            _ = fi80
                            _ = v82
                            _ = v83
                        }
                        _ = fi78
                    }
                    var v84 uint64 = 0
                    _ = v77
                    _ = v84
                }
            }
        }
        v85 := GenS9{F0: ("" + ("hello" + "\n\t")), F1: 1}
        _ = v85
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v7
    _ = v8
}
