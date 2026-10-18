package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
func genFn0(p0 int16, p1 uint8, p2 int16) int16 {
    var v5 uint = 4294967295
    v5 = (v5 + (v5 >> ((2) & 63)))
    var v6 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1000
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -128
    var v3 int = 9223372036854775807
    v4 := []int{}
    fmt.Println(v3, v2)
    if false {
        var v7 int32 = -987
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint64 = 4294967295
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        var v10 uint8 = 128
        v2 = 100
        if true {
            for fi11 := 0; fi11 < 5; fi11++ {
                v10 = (v10 % (((v10 & v10)) | 1))
                v3 = (v3 - 32767)
                v3 -= int(v2)
                _ = fi11
            }
            v1[""] = ""
        } else {
            v10 += (v10 >> ((259) & 63))
        }
        _ = v10
    }
    var v12 int = (v3 << ((3) & 63))
    _ = len(v1)
    for fi13 := 0; fi13 < 20; fi13++ {
        v3 = (v3 / ((v12) | 1))
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int16 = 252
            var v16 uint64 = 18446744073709551615
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    v1["\n\t"] = "hello"
    _ = len(v4)
    v1[("\n\t" + ("foo bar" + "foo bar"))] = "hello"
    if (!((!(1 == v2)) && true)) {
        var v17 uint64 = 2
        switch "hello" {
        case "\n\t":
            fmt.Println(v17, v2)
        default:
            for fi18 := 0; fi18 < 2; fi18++ {
                delete(v1, (("a" + "\n\t") + ""))
                v19 := GenS1{F0: (127 >= 255)}
                _ = fi18
                _ = v19
            }
        }
        _ = v17
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        if false {
            v21 := GenS2{F0: 32767, F1: 42, F2: 524}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 -= v2
            _ = v21
        } else {
            for fi22 := 0; fi22 < 12; fi22++ {
                switch "" {
                case "\n\t":
                    v23 := v1["\n\t"]
                    _ = v23
                case "hello":
                    v4 = append(v4, 0)
                    if (genFn0(761, 2, 42) > genFn0(1, 0, 531)) {
                        var v24 uint64 = 2
                        v25 := GenS2{F0: 2, F1: 32767, F2: (v24 >> ((65535) & 63))}
                        if true {
                            v2 -= (v2 - (v2 / ((v2) | 1)))
                            switch (fi22 >> ((284) & 63)) {
                            case 0:
                                var v26 int64 = -9223372036854775808
                                var v27 float64 = (3.14159 * 1.5)
                                _ = v26
                                _ = v27
                            case 9223372036854775807:
                                v12 -= 127
                            case -515:
                                var v28 uint = 77
                                _ = v28
                            }
                            v12 = (v3 - (^(v3)))
                        }
                        _ = v24
                        _ = v25
                    } else {
                        var v29 int16 = -844
                        delete(v1, "a")
                        v12 += (^(fi22))
                        _ = v29
                    }
                    if (-2147483648 == 1) {
                        if true {
                            fmt.Println(fi20)
                            v30 := GenS2{F0: 830, F1: 1000, F2: 4294967295}
                            v12 += v3
                            _ = v30
                        } else {
                            v31 := GenS3{F0: 1, F1: 18446744073709551615}
                            var v32 float64 = (((1.5 + ((3.14159 / (1e10 - float64((float32((-2.25 / 0.1)) / 3.14159)))) - 0.0)) * 0.1) * (1e10 * 0.0))
                            v33 := GenS4{F0: 42, F1: 1, F2: 32767}
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        if true {
                            var v34 int8 = (^(v2))
                            var v35 uint32 = 381
                            switch (("foo bar" + "") + "hello") {
                            case "a":
                                switch "" {
                                case "hello":
                                    fmt.Println(fi20)
                                default:
                                    var v36 string = "foo bar"
                                    _ = v36
                                }
                            }
                            _ = v34
                            _ = v35
                        }
                    }
                }
                switch "a" {
                case "a":
                    if false {
                        v2 -= v2
                        v4 = append(v4, (^(fi20)))
                    }
                default:
                    var v37 uint16 = 378
                    delete(v1, ("hello" + ""))
                    _ = v37
                }
                var v38 uint8 = 205
                var v39 float32 = 1e10
                _ = len(v1)
                _ = fi22
                _ = v38
                _ = v39
            }
            switch "hello" {
            case "":
                for fi40 := 0; fi40 < 19; fi40++ {
                    switch v3 {
                    case -638:
                        v41 := GenS2{F0: 429, F1: genFn0(42, 42, 32767), F2: 2}
                        _ = v41
                    case 42:
                        fmt.Println(fi20, v3)
                    }
                    _ = fi40
                }
            case "foo bar":
                for fi42 := 0; fi42 < 7; fi42++ {
                    if false {
                        fmt.Println(v3, v12)
                        v43 := GenS2{F0: 42, F1: 503, F2: 723}
                        _ = v43
                    }
                    var v44 float64 = (3.14159 - (((1e10 - (float64(3.14159) / (float64((-2.25 - 0.0)) - float64(float32(1e10))))) - 0.0) * float64(((3.14159 + -2.25) * ((1.5 - ((0.1 + (0.1 * 1.5)) * 0.1)) + -2.25)))))
                    _ = fi42
                    _ = v44
                }
            case "\n\t":
                var v45 uint16 = 867
                _ = v45
            default:
                if (212 < -2147483648) {
                    v12 += (v3 * (v12 + (v12 + (^(v3)))))
                    var v46 int64 = 238
                    _ = v46
                } else {
                    if ((1 < 42) || false) {
                        fmt.Println(v2, v3)
                        v47 := GenS1{F0: (2147483647 < 326)}
                        for fi48 := 0; fi48 < 6; fi48++ {
                            var v49 uint = 4294967295
                            var v50 uint8 = 128
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        _ = v47
                    } else {
                        if true {
                            for fi51 := 0; fi51 < 3; fi51++ {
                                var v52 uint16 = 2
                                fmt.Println(v52, v2)
                                _ = fi51
                                _ = v52
                            }
                            v53 := GenS5{F0: 2}
                            _ = v53
                        }
                        if len(v4) > 0 {
                            v4[0] = (v3 / ((2147483647) | 1))
                        }
                        _ = len(v4)
                    }
                }
            }
        }
        _ = fi20
    }
    delete(v1, "\n\t")
    var v54 int16 = 579
    if (0.1 != -2.25) {
        v4 = append(v4, (v12 << ((2) & 63)))
    } else {
        for fi55 := 0; fi55 < 15; fi55++ {
            for fi56 := 0; fi56 < 1; fi56++ {
                v57 := GenS3{F0: 175, F1: 669}
                _ = fi56
                _ = v57
            }
            _ = len(v4)
            _ = fi55
        }
    }
    if (18446744073709551615 >= 2) {
        v2 = (v2 + v2)
        if false {
            var v58 uint32 = 296
            _ = v58
        } else {
            var v59 uint16 = 1
            var v60 float64 = -2.25
            for fi61 := 0; fi61 < 17; fi61++ {
                for fi62 := 0; fi62 < 12; fi62++ {
                    fmt.Println(v3, v60, v12)
                    v3 = fi61
                    _ = fi62
                }
                _ = fi61
            }
            _ = v59
            _ = v60
        }
        for fi63 := 0; fi63 < 4; fi63++ {
            v2 += v2
            if false {
                var v64 int64 = -1
                switch v3 {
                case 1:
                    var v65 uint8 = 128
                    _ = v65
                default:
                    var v66 int32 = -854
                    _ = v66
                }
                for fi67 := 0; fi67 < 13; fi67++ {
                    var v68 int = 1
                    v64 -= (^(v64))
                    fmt.Println(v2, v3, v64)
                    _ = fi67
                    _ = v68
                }
                _ = v64
            } else {
                if true {
                    for fi69 := 0; fi69 < 1; fi69++ {
                        if len(v4) > 0 {
                            v4[0] = v3
                        }
                        fmt.Println(v54)
                        var v70 int64 = -9223372036854775808
                        _ = fi69
                        _ = v70
                    }
                }
                v12 -= v3
                for fi71 := 0; fi71 < 18; fi71++ {
                    v72 := GenS6{F0: (!((65535 > 61) || true))}
                    _ = fi71
                    _ = v72
                }
            }
            var v73 int = -842
            _ = fi63
            _ = v73
        }
    }
    switch "foo bar" {
    case "a":
        var v74 float32 = (0.0 * (float32(float64(1e10)) / ((-2.25 + (-2.25 / 1.0)) * (float32((1e10 * 1e10)) * (float32(1e10) - 0.0)))))
        _ = v74
    }
    v54 = v54
    for fi75 := 0; fi75 < 5; fi75++ {
        if (true || true) {
            if false {
                var v76 uint32 = 781
                v76 -= (v76 + v76)
                _ = v76
            } else {
                v77 := GenS2{F0: 921, F1: (v54 - v54), F2: 0}
                v2 = (v2 % (((^(v2))) | 1))
                _ = v77
            }
            v2 += (^(v2))
        }
        _ = fi75
    }
    for fi78 := 0; fi78 < 5; fi78++ {
        v79 := GenS6{F0: (false && false)}
        _ = fi78
        _ = v79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v54
}
