package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 bool }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int64; F2 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint }
func genFn0() int64 {
    var v7 int16 = 42
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int32 = 2147483647
        _ = fi8
        _ = v9
    }
    fmt.Println(v7)
    var v10 bool = false
    _ = v10
    _ = v7
    return 42
}

func main() {
    var v1 int = 42
    var v2 uint8 = 127
    v3 := []int{(-(v1)), 0}
    var v4 uint64 = 1
    var v5 uint32 = 4294967295
    var v6 int64 = -9223372036854775808
    v2 += v2
    if false {
        v11 := GenS1{F0: 771}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                _ = len(v3)
                v14 := GenS2{F0: 65535}
                _ = fi13
                _ = v14
            }
            v15 := GenS2{F0: 358}
            _ = fi12
            _ = v15
        }
    }
    v16 := GenS3{F0: 877, F1: -741, F2: ((v6 | (^(v6))) != (-(v6)))}
    for fi17 := 0; fi17 < 14; fi17++ {
        v2 = (v2 ^ v2)
        _ = fi17
    }
    switch "" {
    case "":
        v6 = genFn0()
    case "foo bar":
        v5 += v5
    default:
        if ((-(v6)) >= genFn0()) {
            if len(v3) > 0 {
                v3[0] = v1
            }
            v18 := GenS1{F0: 915}
            for fi19 := 0; fi19 < 15; fi19++ {
                v1 -= (fi19 | (v1 * (fi19 % ((v1) | 1))))
                for fi20 := 0; fi20 < 20; fi20++ {
                    fmt.Println(v6, v5, v6)
                    v5 += (v5 * (v5 & v5))
                    fmt.Println(v5, v4)
                    _ = fi20
                }
                fmt.Println(fi19, v5, v4)
                _ = fi19
            }
            _ = v18
        }
    }
    v21 := GenS1{F0: 42}
    v22 := GenS2{F0: 255}
    v4 = v4
    var v23 uint8 = v2
    if (v1 != (-(v1))) {
        for fi24 := 0; fi24 < 5; fi24++ {
            if false {
                v3 = append(v3, -9223372036854775808)
            }
            switch ("a" + "a") {
            case "\n\t":
                if false {
                    v2 = 2
                    v25 := GenS4{F0: v1, F1: 876, F2: 967}
                    _ = v25
                } else {
                    fmt.Println(v6, v2)
                    for fi26 := 0; fi26 < 3; fi26++ {
                        var v27 bool = (v4 == uint64(v5))
                        _ = fi26
                        _ = v27
                    }
                    switch "\n\t" {
                    case "a":
                        v23 -= (^(v2))
                    case "foo bar":
                        var v28 int64 = genFn0()
                        _ = v28
                    default:
                        if len(v3) > 0 {
                            v3[0] = int(v6)
                        }
                    }
                }
            case "":
                for fi29 := 0; fi29 < 2; fi29++ {
                    v30 := GenS5{F0: 1}
                    var v31 int8 = 127
                    var v32 float64 = 0.0
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
            default:
                v1 = (fi24 & (v1 & (v1 | (fi24 / (((v1 << ((1) & 63))) | 1)))))
            }
            _ = fi24
        }
        if false {
            switch ("\n\t" + "foo bar") {
            case "hello":
                if false {
                    v23 -= v2
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 float32 = 1e10
                        fmt.Println(v4)
                        v35 := GenS3{F0: 430, F1: 2147483647, F2: ((^(v4)) != v4)}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                } else {
                    var v36 int = (^(v1))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi37 := 0; fi37 < 10; fi37++ {
                        v38 := GenS3{F0: 108, F1: 0, F2: true}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v39 string = "foo bar"
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = v36
                }
            case "\n\t":
                _ = len(v3)
            }
        } else {
            if true {
                v23 = (v23 | (v23 & v2))
            }
        }
        var v40 uint8 = 235
        _ = v40
    } else {
        v41 := GenS2{F0: 832}
        switch v1 {
        case -1:
            v6 -= (v6 * genFn0())
        }
        if (1.5 > 0.0) {
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 int8 = 127
                _ = fi42
                _ = v43
            }
            switch (^(v1)) {
            case 9223372036854775807:
                _ = len(v3)
            case 2147483647:
                fmt.Println(v6, v23)
            case -2147483648:
                for fi44 := 0; fi44 < 15; fi44++ {
                    switch (v1 | fi44) {
                    case 867:
                        switch v1 {
                        case -1:
                            for fi45 := 0; fi45 < 11; fi45++ {
                                v46 := GenS6{F0: 1}
                                fmt.Println(v4, v1, v23)
                                var v47 uint16 = 1
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                        case -508:
                            var v48 uint = 341
                            _ = v48
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case 0:
                        fmt.Println(v4)
                    case 150:
                        switch ("" + (("foo bar" + ("foo bar" + "\n\t")) + "\n\t")) {
                        case "a":
                            fmt.Println(v5)
                        }
                    default:
                        if (0 == 0) {
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = 127
                                }
                            } else {
                                var v49 uint64 = v4
                                for fi50 := 0; fi50 < 2; fi50++ {
                                    var v51 int16 = 32767
                                    _ = len(v3)
                                    _ = fi50
                                    _ = v51
                                }
                                _ = v49
                            }
                            switch ((("\n\t" + "a") + "foo bar") + (("foo bar" + ("foo bar" + "\n\t")) + "foo bar")) {
                            case "foo bar":
                                for fi52 := 0; fi52 < 10; fi52++ {
                                    if true {
                                        v53 := GenS7{F0: 296, F1: -259, F2: 100}
                                        _ = v53
                                    }
                                    switch v1 {
                                    case 128:
                                        if len(v3) > 0 {
                                            v3[0] = fi44
                                        }
                                    }
                                    if false {
                                        v3 = append(v3, fi52)
                                        v2 += (v23 >> ((uint(v4)) & 63))
                                        v1 -= (fi44 & fi44)
                                    } else {
                                        if false {
                                            v6 = -1
                                            var v54 float64 = 3.14159
                                            fmt.Println(v23)
                                            _ = v54
                                        }
                                        if true {
                                            var v55 bool = true
                                            v1 = -684
                                            _ = v55
                                        }
                                        v56 := GenS2{F0: 18446744073709551615}
                                        _ = v56
                                    }
                                    _ = fi52
                                }
                            case "\n\t":
                                v5 += 2
                            case "a":
                                for fi57 := 0; fi57 < 16; fi57++ {
                                    var v58 int16 = 1000
                                    v58 = (-(v58))
                                    var v59 int16 = v58
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                            }
                            var v60 uint8 = v23
                            _ = v60
                        }
                    }
                    fmt.Println(v6, v5)
                    _ = fi44
                }
            }
            if len(v3) > 0 {
                v3[0] = v1
            }
        }
        _ = v41
    }
    if (!(false || true)) {
        var v61 float64 = 0.1
        if (!true) {
            for fi62 := 0; fi62 < 2; fi62++ {
                v63 := GenS3{F0: 203, F1: 79, F2: (!true)}
                if false {
                    var v64 uint32 = uint32(v4)
                    _ = v64
                } else {
                    v2 += (^(v2))
                    var v65 string = "a"
                    _ = v65
                }
                v23 = 2
                _ = fi62
                _ = v63
            }
            for fi66 := 0; fi66 < 18; fi66++ {
                if true {
                    v1 = (v1 - -928)
                    v3 = append(v3, fi66)
                }
                switch (fi66 >> ((803) & 63)) {
                case -270:
                    if true {
                        if len(v3) > 0 {
                            v3[0] = fi66
                        }
                    } else {
                        for fi67 := 0; fi67 < 4; fi67++ {
                            v2 = v23
                            switch fi67 {
                            case 9223372036854775807:
                                v3 = append(v3, 396)
                            case 2147483647:
                                v4 = v4
                            case 32767:
                                if false {
                                    fmt.Println(v5, v6, fi67)
                                    if false {
                                        var v68 uint64 = 2
                                        v1 = 601
                                        _ = v68
                                    } else {
                                        var v69 int = -813
                                        var v70 int8 = 1
                                        switch fi67 {
                                        case 32767:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = len(v3)
                                        case -901:
                                            switch fi66 {
                                            case -1:
                                                switch int(v5) {
                                                case -1:
                                                    var v71 float64 = (v61 - (v61 * v61))
                                                    _ = v71
                                                case 9223372036854775807:
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v72 float32 = ((0.1 / 0.1) / 1.0)
                                                        _ = v72
                                                    }
                                                    var v73 float32 = 0.1
                                                    if (((((v73 * (3.14159 - (0.0 - (float32(-2.25) + v73)))) + v73) / v73) * v73) < v73) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v70 = (^(v70))
                                                        var v74 string = "\n\t"
                                                        _ = v74
                                                    }
                                                    _ = v73
                                                case -2147483648:
                                                    v70 = 100
                                                }
                                            case 0:
                                                switch fi66 {
                                                case 346:
                                                    v75 := GenS4{F0: fi66, F1: 947, F2: (v5 & 830)}
                                                    _ = v75
                                                case 2147483647:
                                                    v76 := GenS6{F0: 65535}
                                                    _ = v76
                                                default:
                                                    switch "" {
                                                    case "\n\t":
                                                        fmt.Println(v5)
                                                    case "hello":
                                                        var v77 uint8 = v23
                                                        _ = v77
                                                    case "foo bar":
                                                        v70 = v70
                                                        fmt.Println(v2, v2, v69)
                                                        fmt.Println(fi67)
                                                    }
                                                }
                                            default:
                                                v4 += (v4 ^ v4)
                                            }
                                        default:
                                            fmt.Println(v6, v5)
                                        }
                                        var v78 int16 = -345
                                        _ = v69
                                        _ = v70
                                        _ = v78
                                    }
                                }
                            }
                            _ = fi67
                        }
                        v79 := GenS8{F0: 0}
                        _ = v79
                    }
                default:
                    for fi80 := 0; fi80 < 9; fi80++ {
                        var v81 int32 = 0
                        v6 -= (v6 / ((9223372036854775807) | 1))
                        v81 = int32(v5)
                        _ = fi80
                        _ = v81
                    }
                }
                for fi82 := 0; fi82 < 5; fi82++ {
                    var v83 int16 = -336
                    var v84 uint32 = 0
                    fmt.Println(v6, fi66)
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                _ = fi66
            }
            for fi85 := 0; fi85 < 1; fi85++ {
                switch "" {
                case "hello":
                    v1 += (fi85 ^ fi85)
                case "":
                    v6 -= (v6 * 556)
                }
                v3 = append(v3, (^(fi85)))
                for fi86 := 0; fi86 < 10; fi86++ {
                    fmt.Println(fi86)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v4, v4, v1)
                    _ = fi86
                }
                _ = fi85
            }
        }
        _ = v61
    }
    if false {
        var v87 float64 = float64(1e10)
        fmt.Println(v87, v6)
        _ = v87
    } else {
        v88 := GenS9{F0: 851}
        _ = v88
    }
    var v89 uint8 = (v2 + (v2 ^ (v2 * (^(v2)))))
    if len(v3) > 0 {
        v3[0] = v1
    }
    var v90 float64 = float64(1e10)
    v91 := GenS8{F0: v6}
    var v92 uint32 = uint32(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
    _ = v92
}
