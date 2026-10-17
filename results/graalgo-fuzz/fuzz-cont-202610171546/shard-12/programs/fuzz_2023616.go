package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 uint; F1 int16; F2 bool }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64 }

func main() {
    var v1 bool = ((float32(float64((3.14159 + -2.25))) * 0.1) < (-2.25 + (1.5 / (((float32(1.5) + 1e10) / 3.14159) / 3.14159))))
    var v2 uint64 = 1
    v3 := make(map[uint32]int)
    v4 := []string{("" + "foo bar"), "a"}
    var v5 float64 = (((1.5 - ((3.14159 / 1.0) - 1.5)) + float64((1.5 * 0.1))) * (1e10 - -2.25))
    var v6 int16 = 32767
    if (42 >= 2147483647) {
        var v7 bool = false
        var v8 int64 = 120
        var v9 int32 = -2147483648
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 bool = (0.1 == float32((v5 - 3.14159)))
        v11 := GenS1{F0: 127}
        for fi12 := 0; fi12 < 15; fi12++ {
            switch -2147483648 {
            case 42:
                fmt.Println(v10, v6)
            case -1:
                v6 += v6
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    switch -971 {
    case -9223372036854775808:
        v4 = append(v4, "foo bar")
    case 0:
        var v13 float64 = v5
        _ = v13
    }
    var v14 string = "hello"
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            v14 = "foo bar"
            v17 := GenS1{F0: 1}
            var v18 int = 2147483647
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 int16 = (v6 / (((v6 - v6)) | 1))
        _ = fi15
        _ = v19
    }
    v5 += ((v5 - -2.25) * v5)
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS2{F0: 805}
        for fi22 := 0; fi22 < 20; fi22++ {
            fmt.Println(fi22, v2)
            _ = fi22
        }
        switch fi20 {
        case 9223372036854775807:
            _ = len(v4)
        case 42:
            for fi23 := 0; fi23 < 17; fi23++ {
                if len(v4) > 0 {
                    v4[0] = (v14 + v14)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v5, fi23, fi20)
                _ = fi23
            }
        case -374:
            v24 := GenS3{F0: (v6 ^ v6), F1: 2}
            _ = v24
        default:
            if v1 {
                delete(v3, 2)
                switch (v14 + "\n\t") {
                case "a":
                    if (!(v1 && (!v1))) {
                        if len(v4) > 0 {
                            v4[0] = (v14 + "hello")
                        }
                        fmt.Println(v6)
                        v25 := GenS4{F0: 855, F1: -2147483648}
                        _ = v25
                    } else {
                        var v26 int8 = 1
                        fmt.Println(v6)
                        if v1 {
                            switch "" {
                            case "foo bar":
                                v3[65535] = fi20
                            case "":
                                delete(v3, 2)
                            default:
                                for fi27 := 0; fi27 < 16; fi27++ {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (470 > 513) {
                                        v6 = v6
                                        v26 = v26
                                        v28 := v3[53]
                                        _ = v28
                                    } else {
                                        fmt.Println(fi20, v2, v14)
                                        if ((v14 + (v14 + "hello")) >= v14) {
                                            switch (("foo bar" + "foo bar") + v14) {
                                            case "\n\t":
                                                v5 = 1.5
                                            case "foo bar":
                                                v29 := GenS4{F0: 65535, F1: -566}
                                                _ = v29
                                            default:
                                                if ((0.1 < 0.1) || v1) {
                                                    if len(v4) > 0 {
                                                        v4[0] = v14
                                                    }
                                                    v30 := GenS3{F0: (v6 >> ((499) & 63)), F1: 255}
                                                    if (v14 >= v14) {
                                                        if false {
                                                            fmt.Println(fi27, fi27, v6)
                                                        } else {
                                                            v4 = append(v4, v14)
                                                            if v1 {
                                                                fmt.Println(v2)
                                                                v31 := v3[2147483648]
                                                                _ = v31
                                                            }
                                                        }
                                                    }
                                                    _ = v30
                                                }
                                            }
                                            v6 = v6
                                        } else {
                                            v32 := GenS3{F0: 42, F1: 1}
                                            if ((v6 > v6) || (0.0 > (0.1 / (float32(1e10) * float32(float64(0.1)))))) {
                                                fmt.Println(v2, fi27)
                                                var v33 int16 = v6
                                                _ = v33
                                            } else {
                                                var v34 string = (v14 + "\n\t")
                                                _ = v34
                                            }
                                            var v35 float64 = v5
                                            _ = v32
                                            _ = v35
                                        }
                                    }
                                    _ = fi27
                                }
                            }
                            v36 := GenS5{F0: 977, F1: v6, F2: (v1 && (65535 >= 4294967295))}
                            if len(v4) > 0 {
                                v4[0] = v14
                            }
                            _ = v36
                        }
                        _ = v26
                    }
                }
            }
        }
        _ = fi20
        _ = v21
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        if len(v4) > 0 {
            v4[0] = v14
        }
        var v38 float32 = ((0.1 * 0.0) + ((0.1 * 0.1) / ((((-2.25 - (((float32(v5) * 0.1) / (1.5 + (3.14159 / ((3.14159 / 1.0) / ((1e10 + (1.5 * float32(1e10))) - 1.5))))) * 1e10)) - float32((-2.25 * 3.14159))) + 1e10) * (-2.25 * 0.0))))
        if true {
            v38 = (v38 + v38)
            v2 += v2
        }
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        v40 := GenS6{F0: (v14 + "hello"), F1: uint(v2)}
        v2 += (v2 + (^(v2)))
        _ = fi39
        _ = v40
    }
    if v1 {
        v6 = 481
    }
    if v1 {
        for fi41 := 0; fi41 < 10; fi41++ {
            for fi42 := 0; fi42 < 20; fi42++ {
                v43 := GenS7{F0: uint16(v2)}
                var v44 uint16 = 0
                _ = fi42
                _ = v43
                _ = v44
            }
            for fi45 := 0; fi45 < 9; fi45++ {
                switch fi45 {
                case -2147483648:
                    fmt.Println(fi41, fi41)
                case 2147483647:
                    delete(v3, 42)
                default:
                    _ = len(v3)
                    v46 := GenS6{F0: v14, F1: 18446744073709551615}
                    var v47 bool = v1
                    _ = v46
                    _ = v47
                }
                v2 = 1
                _ = fi45
            }
            _ = fi41
        }
    } else {
        for fi48 := 0; fi48 < 8; fi48++ {
            var v49 float64 = 3.14159
            _ = fi48
            _ = v49
        }
        var v50 int64 = -351
        for fi51 := 0; fi51 < 17; fi51++ {
            v50 += v50
            _ = fi51
        }
        _ = v50
    }
    for fi52 := 0; fi52 < 9; fi52++ {
        switch (^(fi52)) {
        case 128:
            if (42 != 1) {
                if (281 <= 249) {
                    var v53 int8 = 1
                    _ = v53
                } else {
                    v54 := GenS6{F0: v14, F1: 283}
                    switch 789 {
                    case 127:
                        var v55 uint32 = 2147483648
                        _ = v55
                    case 32767:
                        if (v1 || true) {
                            v2 += (v2 / ((4294967295) | 1))
                        }
                    default:
                        if (!v1) {
                            switch "\n\t" {
                            case "foo bar":
                                v5 = ((v5 - v5) + 0.1)
                            case "\n\t":
                                switch -172 {
                                case -954:
                                    if v1 {
                                        var v56 int = (fi52 ^ (fi52 + fi52))
                                        v6 += (v6 * 32767)
                                        v5 += float64(((1.5 / (0.1 + (0.1 + 0.0))) * 1.5))
                                        _ = v56
                                    } else {
                                        if false {
                                            v14 = v14
                                        }
                                        v5 = (((-2.25 * 1e10) - v5) * float64(float32((v5 * (v5 - v5)))))
                                        v5 = v5
                                    }
                                }
                            case "a":
                                v3[2147483647] = (fi52 - (^(fi52)))
                            default:
                                v2 = (^(v2))
                            }
                            v57 := v3[1]
                            _ = v57
                        }
                    }
                    _ = v54
                }
                var v58 int8 = 0
                v1 = true
                _ = v58
            }
        case -458:
            v5 = ((v5 * v5) - v5)
        default:
            fmt.Println(v5, v5, v14)
        }
        v59 := GenS8{F0: 940}
        _ = fi52
        _ = v59
    }
    for fi60 := 0; fi60 < 13; fi60++ {
        if v1 {
            var v61 uint64 = (^(v2))
            v3[2] = (fi60 - fi60)
            _ = v61
        } else {
            v62 := GenS5{F0: 4294967295, F1: v6, F2: false}
            var v63 int16 = (v6 + (v6 * v6))
            for fi64 := 0; fi64 < 15; fi64++ {
                var v65 uint64 = 360
                switch (v14 + "hello") {
                case "hello":
                    v65 -= (v65 | v2)
                case "\n\t":
                    v1 = v1
                case "foo bar":
                    fmt.Println(v63, v2, v1)
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi64
                _ = v65
            }
            _ = v62
            _ = v63
        }
        _ = fi60
    }
    v5 = 1e10
    v5 = (float64((0.0 + -2.25)) + (((0.1 - 3.14159) + v5) / v5))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
