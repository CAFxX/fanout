package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 int64; F1 uint16; F2 int }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint; F1 int32 }

func main() {
    var v1 string = "a"
    var v2 uint = 4294967295
    var v3 uint8 = 78
    var v4 uint8 = v3
    var v5 uint64 = 535
    var v6 uint32 = 606
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    v3 = 111
    if true {
        var v9 int16 = 993
        _ = v9
    }
    v5 -= (v5 | (^(v5)))
    v10 := []int{32767, 268, int(v2), 0}
    _ = len(v10)
    v11 := GenS1{F0: -178, F1: 168}
    if (true && true) {
        v5 -= (v5 >> ((4294967295) & 63))
    } else {
        var v12 float64 = (1e10 * 1.5)
        if true {
            var v13 int = 806
            if (42 != 1) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    if false {
                        v15 := GenS1{F0: 9223372036854775807, F1: 128}
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v4 += (v3 ^ uint8(v5))
                            v17 := GenS1{F0: -1, F1: uint8(v2)}
                            fmt.Println(v5, v12, fi14)
                            _ = fi16
                            _ = v17
                        }
                        fmt.Println(v3)
                    }
                    if false {
                        fmt.Println(v5)
                        fmt.Println(v3, fi14, v4)
                        var v18 uint16 = uint16(v2)
                        _ = v18
                    } else {
                        v1 = (("\n\t" + "") + v1)
                        v19 := GenS1{F0: 960, F1: v3}
                        _ = v19
                    }
                    if false {
                        switch ((v1 + v1) + (("" + ("" + (("" + v1) + "foo bar"))) + (("a" + (v1 + v1)) + v1))) {
                        case "foo bar":
                            v20 := GenS2{F0: (v4 ^ (v4 << ((v2) & 63))), F1: 602, F2: (v5 / ((v5) | 1))}
                            _ = v20
                        case "\n\t":
                            switch fi14 {
                            case 32768:
                                for fi21 := 0; fi21 < 20; fi21++ {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    _ = fi21
                                }
                            case 1:
                                var v22 float32 = -2.25
                                _ = v22
                            case -259:
                                switch -939 {
                                case 127:
                                    fmt.Println(v3, v3)
                                case 32767:
                                    _ = len(v10)
                                case -9223372036854775808:
                                    v6 = 155
                                }
                            default:
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                        }
                    }
                    _ = fi14
                }
                var v23 int8 = int8(v6)
                if len(v10) > 0 {
                    v10[0] = v13
                }
                _ = v23
            }
            _ = v13
        }
        _ = v12
    }
    var v24 uint = uint(v5)
    v25 := GenS1{F0: 42, F1: v3}
    switch v1 {
    case "":
        var v26 int16 = 1
        _ = v26
    }
    if (v3 != 255) {
        v6 = (^(v6))
    } else {
        v27 := GenS3{F0: (^(v6)), F1: (1 != 0)}
        _ = v27
    }
    var v28 uint8 = v4
    fmt.Println(v4)
    _ = len(v10)
    switch v1 {
    case "a":
        for fi29 := 0; fi29 < 17; fi29++ {
            var v30 int32 = 0
            v30 -= (v30 - int32(v5))
            v1 = v1
            _ = fi29
            _ = v30
        }
    case "\n\t":
        var v31 bool = true
        _ = v31
    case "":
        fmt.Println(v28, v2, v1)
    default:
        for fi32 := 0; fi32 < 11; fi32++ {
            v1 = v1
            switch (fi32 & fi32) {
            case -611:
                switch fi32 {
                case -2147483648:
                    for fi33 := 0; fi33 < 19; fi33++ {
                        if (!(127 <= 1)) {
                            v4 = uint8(v5)
                            v1 = v1
                            v4 = 2
                        } else {
                            if true {
                                v34 := GenS4{F0: v1}
                                v1 = "foo bar"
                                _ = v34
                            } else {
                                if len(v10) > 0 {
                                    v10[0] = (fi32 / (((fi32 * fi33)) | 1))
                                }
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                if ((!false) && false) {
                                    v2 = v24
                                    if (1 >= (v28 & v28)) {
                                        if (v24 <= (v2 & (^(v2)))) {
                                            _ = len(v10)
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                if true {
                                                    v3 -= v28
                                                    if false {
                                                        switch -576 {
                                                        case 1:
                                                            v6 -= 1
                                                        case -228:
                                                            var v35 int = (fi32 | 42)
                                                            _ = v35
                                                        case 32768:
                                                            v24 -= uint(v3)
                                                        }
                                                        fmt.Println(v2)
                                                        v36 := GenS5{F0: 18446744073709551615}
                                                        _ = v36
                                                    } else {
                                                        if len(v10) > 0 {
                                                            v10[0] = fi33
                                                        }
                                                    }
                                                    v24 += 42
                                                    v37 := GenS4{F0: v1}
                                                    _ = v37
                                                } else {
                                                    v38 := GenS4{F0: ""}
                                                    v5 = 1
                                                    _ = v38
                                                }
                                            case "\n\t":
                                                v10 = append(v10, fi32)
                                            }
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            if true {
                                                var v39 int16 = 32767
                                                switch "" {
                                                case "\n\t":
                                                    v10 = append(v10, (fi33 | int(v24)))
                                                default:
                                                    var v40 uint64 = v5
                                                    _ = v40
                                                }
                                                var v41 float64 = 0.1
                                                _ = v39
                                                _ = v41
                                            } else {
                                                fmt.Println(v3, v4, v28)
                                            }
                                            var v42 int = fi32
                                            if (1 >= -9223372036854775808) {
                                                v1 = (v1 + ((v1 + v1) + v1))
                                                var v43 int8 = 127
                                                var v44 int8 = v43
                                                _ = v43
                                                _ = v44
                                            } else {
                                                if (-720 < -1) {
                                                    v24 -= 42
                                                    v45 := GenS5{F0: (v5 + 42)}
                                                    _ = v45
                                                }
                                            }
                                            _ = v42
                                        }
                                        _ = len(v10)
                                    }
                                    switch 0 {
                                    case 32767:
                                        var v46 uint = v2
                                        _ = v46
                                    case 127:
                                        fmt.Println(v1, v6, v5)
                                    case 479:
                                        v10 = append(v10, fi32)
                                    }
                                    switch (v1 + ((((("\n\t" + v1) + ("" + v1)) + v1) + "hello") + "hello")) {
                                    case "\n\t":
                                        var v47 bool = (v6 < (v6 ^ v6))
                                        _ = v47
                                    }
                                    switch fi33 {
                                    case -931:
                                        if (-9223372036854775808 >= 409) {
                                            v48 := GenS6{F0: (^(v28)), F1: 0}
                                            _ = v48
                                        } else {
                                            if false {
                                                switch fi32 {
                                                case 128:
                                                    if len(v10) > 0 {
                                                        v10[0] = (-(fi33))
                                                    }
                                                case 2147483647:
                                                    fmt.Println(v3)
                                                case 931:
                                                    v2 = 4294967295
                                                }
                                                v6 = 0
                                                v3 -= 128
                                                switch 127 {
                                                case 126:
                                                    if (42 > -1) {
                                                        if len(v10) > 0 {
                                                            v10[0] = (-(fi32))
                                                        }
                                                        if (v5 != (v5 * (^(v5)))) {
                                                            v3 -= (^(v3))
                                                            v2 -= (v24 ^ v2)
                                                            if (!false) {
                                                                var v49 uint = v24
                                                                if true {
                                                                    fmt.Println(fi32, v5)
                                                                }
                                                                _ = v49
                                                            } else {
                                                                var v50 int = (fi32 * (fi32 ^ fi32))
                                                                var v51 uint16 = 1
                                                                v52 := GenS7{F0: 0, F1: (v5 <= (^(v5)))}
                                                                _ = v50
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            v53 := GenS4{F0: "a"}
                                                            switch fi32 {
                                                            case -9223372036854775808:
                                                                var v54 bool = true
                                                                _ = v54
                                                            case -452:
                                                                _ = len(v10)
                                                            default:
                                                                v55 := GenS8{F0: 1, F1: 515, F2: fi33}
                                                                _ = v55
                                                            }
                                                            _ = v53
                                                        }
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                    }
                                                case -47:
                                                    var v56 uint32 = v6
                                                    _ = v56
                                                case 650:
                                                    var v57 float64 = (3.14159 - float64(0.1))
                                                    _ = v57
                                                default:
                                                    if ((-2.25 <= float32(1e10)) || false) {
                                                        if (((v1 + "a") > v1) && (-1 >= 42)) {
                                                            var v58 int32 = -969
                                                            v59 := GenS9{F0: v6}
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        fmt.Println(v24, fi32)
                                                    } else {
                                                        var v60 float64 = 1.5
                                                        var v61 int32 = 42
                                                        switch "" {
                                                        case "a":
                                                            v24 += 2
                                                        case "":
                                                            if (true && (false || true)) {
                                                                v1 = "\n\t"
                                                                var v62 bool = true
                                                                var v63 int16 = -36
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                        case "\n\t":
                                                            v1 = v1
                                                        default:
                                                            v64 := GenS7{F0: -1, F1: false}
                                                            if len(v10) > 0 {
                                                                v10[0] = (fi32 & (fi32 / (((fi32 * (fi33 % ((fi33) | 1)))) | 1)))
                                                            }
                                                            _ = v64
                                                        }
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                }
                                                v5 -= (v5 << (((v2 | v24)) & 63))
                                            }
                                            v24 = 65535
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            switch v1 {
                                            case "":
                                                switch (fi32 % (((fi32 * (fi33 + -1))) | 1)) {
                                                case 127:
                                                    var v65 float64 = ((0.1 - (-2.25 / (float64(((0.0 + -2.25) * ((0.0 * 0.1) - ((1.5 + 1.5) + 0.1)))) + ((0.0 * (float64(1e10) * 0.0)) * -2.25)))) / 3.14159)
                                                    _ = v65
                                                }
                                                if (!false) {
                                                    var v66 int32 = -516
                                                    v4 -= (^(v4))
                                                    _ = len(v10)
                                                    _ = v66
                                                }
                                            case "a":
                                                fmt.Println(v6, v24)
                                            case "hello":
                                                var v67 int = fi32
                                                _ = v67
                                            default:
                                                fmt.Println(v6)
                                            }
                                        }
                                    default:
                                        v3 += v28
                                    }
                                    v68 := GenS5{F0: (v5 ^ 0)}
                                    _ = v68
                                }
                            }
                            if true {
                                var v69 uint = v24
                                var v70 float64 = 1.5
                                fmt.Println(v6)
                                _ = v69
                                _ = v70
                            } else {
                                v1 = (v1 + ("" + ("hello" + ("foo bar" + ""))))
                                var v71 bool = ((^(v28)) >= (v3 * uint8(v5)))
                                _ = v71
                            }
                        }
                        var v72 int16 = -1
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        _ = fi33
                        _ = v72
                    }
                case -528:
                    for fi73 := 0; fi73 < 7; fi73++ {
                        v74 := GenS10{F0: 1}
                        _ = fi73
                        _ = v74
                    }
                default:
                    v5 -= (v5 / ((v5) | 1))
                }
            default:
                v75 := GenS11{F0: v2, F1: -520}
                _ = v75
            }
            _ = fi32
        }
    }
    var v76 uint16 = 236
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
