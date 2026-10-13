package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 uint16; F2 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 bool }

func main() {
    var v1 string = "\n\t"
    var v2 uint = 557
    var v3 string = v1
    v4 := make(map[int]int)
    v5 := []int{}
    var v6 int16 = -228
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 128, F1: 0}
        _ = fi7
        _ = v8
    }
    if true {
        v9 := GenS2{F0: int(v6), F1: 9223372036854775807}
        v10 := GenS1{F0: 255, F1: 930}
        _ = v10
        _ = v9
    }
    if true {
        var v11 bool = (2147483647 != 405)
        v6 = v6
        _ = v11
    } else {
        v12 := GenS2{F0: -9223372036854775808, F1: -931}
        var v13 uint32 = uint32(v2)
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v2 += v2
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int = (fi14 ^ (fi14 ^ fi14))
            v3 = "a"
            _ = fi15
            _ = v16
        }
        v1 = v1
        _ = fi14
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        fmt.Println(fi17, v2, v6)
        for fi18 := 0; fi18 < 9; fi18++ {
            v3 = ("a" + (((v1 + v1) + "foo bar") + v1))
            v19 := GenS1{F0: 42, F1: 252}
            v3 = (v1 + ((v1 + v3) + (v1 + v1)))
            _ = fi18
            _ = v19
        }
        var v20 uint8 = 153
        _ = fi17
        _ = v20
    }
    v2 = v2
    v2 += v2
    if false {
        var v21 int8 = 42
        v6 = v6
        _ = v21
    }
    v22 := v4[32767]
    v2 -= (^(v2))
    v4[(v22 + v22)] = (v22 << ((v2) & 63))
    fmt.Println(v3, v22)
    for fi23 := 0; fi23 < 9; fi23++ {
        v1 = ""
        if (0 != -1) {
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 uint16 = 42
                _ = fi24
                _ = v25
            }
            if false {
                v3 = v3
            }
        } else {
            var v26 uint8 = 2
            _ = v26
        }
        if (((v3 + "") >= ("hello" + v1)) || true) {
            v27 := GenS3{F0: "", F1: 65535}
            v2 = v2
            v28 := GenS4{F0: 255}
            _ = v27
            _ = v28
        } else {
            var v29 uint64 = 2
            switch v3 {
            case "foo bar":
                v3 = (v3 + v1)
            case "\n\t":
                if true {
                    if true {
                        var v30 uint = (v2 / (((^(v2))) | 1))
                        switch (((v3 + "a") + "") + v3) {
                        case "\n\t":
                            for fi31 := 0; fi31 < 8; fi31++ {
                                var v32 uint16 = 0
                                _ = fi31
                                _ = v32
                            }
                        case "":
                            v33 := v4[v22]
                            _ = v33
                        default:
                            var v34 int64 = 64
                            _ = v34
                        }
                        switch (v1 + ((v1 + "") + v3)) {
                        case "\n\t":
                            var v35 uint = (^(v30))
                            _ = v35
                        }
                        _ = v30
                    } else {
                        v2 += (v2 - (v2 ^ v2))
                        for fi36 := 0; fi36 < 11; fi36++ {
                            if ((v2 / ((v2) | 1)) == v2) {
                                if true {
                                    if true {
                                        v37 := GenS4{F0: 128}
                                        _ = v37
                                    } else {
                                        v38 := v4[(fi23 ^ v22)]
                                        v2 -= v2
                                        fmt.Println(fi36)
                                        _ = v38
                                    }
                                }
                            } else {
                                fmt.Println(fi23)
                                v1 = v3
                            }
                            var v39 uint64 = v29
                            v40 := GenS3{F0: v3, F1: 2147483647}
                            var v41 float32 = float32((-2.25 * ((-2.25 - 3.14159) + 3.14159)))
                            v42 := GenS5{F0: fi36, F1: v6}
                            if (!("foo bar" <= ("" + v1))) {
                                v1 = ""
                            } else {
                                var v43 float32 = ((v41 + (v41 * v41)) - (float32(1.5) / v41))
                                v29 -= v29
                                var v44 uint32 = 789
                                _ = v43
                                _ = v44
                            }
                            if len(v5) > 0 {
                                v5[0] = (^(fi23))
                            }
                            v1 = (((v3 + (v3 + "\n\t")) + "a") + "hello")
                            _ = fi36
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        _ = len(v5)
                    }
                    for fi45 := 0; fi45 < 4; fi45++ {
                        fmt.Println(fi23, v2)
                        if (!false) {
                            var v46 int8 = 1
                            v29 += (v29 * 1)
                            if false {
                                var v47 uint = (v2 * 176)
                                if (true || (2 < 65535)) {
                                    var v48 uint = (v47 / (((v47 + 0)) | 1))
                                    _ = v48
                                }
                                _ = v47
                            }
                            _ = v46
                        } else {
                            if (v29 >= 1) {
                                v22 += v22
                                fmt.Println(v6, fi23, v2)
                                switch (("\n\t" + v1) + "a") {
                                case "\n\t":
                                    v3 = (v1 + "a")
                                case "a":
                                    v3 = v3
                                case "hello":
                                    if len(v5) > 0 {
                                        v5[0] = (fi23 % ((v22) | 1))
                                    }
                                default:
                                    v4[fi23] = -9223372036854775808
                                }
                            } else {
                                if true {
                                    v5 = append(v5, fi23)
                                } else {
                                    v5 = append(v5, fi45)
                                    if true {
                                        switch v22 {
                                        case -9223372036854775808:
                                            delete(v4, 2147483647)
                                        }
                                        v3 = (v1 + "\n\t")
                                    }
                                    var v49 float32 = ((3.14159 + ((float32(1e10) - (-2.25 / 1e10)) * 0.0)) + (3.14159 * ((float32(0.1) * -2.25) - 0.1)))
                                    _ = v49
                                }
                            }
                            v22 = -643
                        }
                        switch v3 {
                        case "\n\t":
                            v50 := GenS2{F0: 127, F1: -70}
                            _ = v50
                        }
                        _ = fi45
                    }
                    v51 := GenS5{F0: 741, F1: (v6 | int16(v2))}
                    _ = v51
                } else {
                    _ = len(v4)
                    var v52 uint32 = 4294967295
                    _ = v52
                }
            }
            _ = v29
        }
        _ = fi23
    }
    if (!(false || ((v2 ^ (v2 / ((v2) | 1))) > (v2 << ((1) & 63))))) {
        if false {
            fmt.Println(v3, v2)
            var v53 int64 = 1
            if ((!false) && ((1.5 > 0.0) || (!(!(!(1.5 >= 1e10)))))) {
                v6 += v6
                switch "\n\t" {
                case "\n\t":
                    if (1 <= 65535) {
                        v2 -= (v2 << (((v2 | (v2 / (((v2 + (v2 * v2))) | 1)))) & 63))
                    } else {
                        v54 := GenS6{F0: 0}
                        if (v6 >= (v6 | (v6 / ((-940) | 1)))) {
                            fmt.Println(v3, v1, v1)
                        } else {
                            v2 = (v2 >> (((^(v2))) & 63))
                            var v55 float64 = 0.1
                            _ = v55
                        }
                        var v56 uint64 = 491
                        _ = v54
                        _ = v56
                    }
                case "a":
                    for fi57 := 0; fi57 < 17; fi57++ {
                        switch v1 {
                        case "hello":
                            var v58 int32 = 42
                            _ = v58
                        case "foo bar":
                            fmt.Println(v22, v6)
                        default:
                            var v59 int8 = -39
                            _ = v59
                        }
                        v53 += -9223372036854775808
                        if false {
                            var v60 uint16 = 39
                            v2 += v2
                            _ = v60
                        } else {
                            if true {
                                v3 = ("" + "a")
                                v61 := GenS7{F0: v2, F1: uint16(v2), F2: -128}
                                _ = v61
                            } else {
                                switch "" {
                                case "a":
                                    delete(v4, v22)
                                }
                                var v62 int64 = v53
                                var v63 float64 = 0.1
                                _ = v62
                                _ = v63
                            }
                            if len(v5) > 0 {
                                v5[0] = 1
                            }
                        }
                        _ = fi57
                    }
                }
                v64 := GenS8{F0: -131}
                _ = v64
            } else {
                v1 = v3
            }
            _ = v53
        }
        for fi65 := 0; fi65 < 16; fi65++ {
            if (!(!(-946 >= 0))) {
                fmt.Println(v1, v6)
            } else {
                switch 168 {
                case 2147483647:
                    v66 := GenS8{F0: -2147483648}
                    _ = v66
                default:
                    v67 := GenS5{F0: -1, F1: (v6 + (v6 | v6))}
                    _ = v67
                }
                var v68 bool = true
                switch fi65 {
                case 32768:
                    v1 = v1
                case 0:
                    v69 := GenS1{F0: 0, F1: 311}
                    _ = v69
                case 9223372036854775807:
                    var v70 uint32 = 42
                    _ = v70
                default:
                    _ = len(v5)
                }
                _ = v68
            }
            _ = fi65
        }
    } else {
        if (-9223372036854775808 == 9223372036854775807) {
            var v71 int32 = -2147483648
            for fi72 := 0; fi72 < 15; fi72++ {
                var v73 int64 = -1
                var v74 int8 = 100
                _ = fi72
                _ = v73
                _ = v74
            }
            _ = v71
        }
        if false {
            v75 := GenS5{F0: (v22 - v22), F1: -714}
            if (v22 != (-(v22))) {
                v5 = append(v5, v22)
                var v76 int = v22
                v76 = v22
                _ = v76
            } else {
                var v77 uint8 = 255
                _ = v77
            }
            _ = v75
        }
        v22 = (v22 + (v22 ^ v22))
    }
    v78 := GenS9{F0: 42, F1: false}
    var v79 float32 = 1.5
    if false {
        var v80 uint16 = 255
        fmt.Println(v6, v22, v2)
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
}
