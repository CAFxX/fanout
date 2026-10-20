package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int8; F1 bool; F2 int64 }
type GenS6 struct { F0 string; F1 int64; F2 int8 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 bool; F1 int64; F2 int32 }
type GenS9 struct { F0 string; F1 bool; F2 string }
func genFn0(p0 uint8, p1 int8, p2 uint32) int32 {
    v5 := []int{-545, 710}
    var v6 string = "hello"
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(p1, p0, p2)
        for fi8 := 0; fi8 < 11; fi8++ {
            p0 = p0
            if true {
                switch v6 {
                case "foo bar":
                    v9 := GenS1{F0: p1}
                    _ = v9
                case "hello":
                    if ((p1 < (p1 | (p1 * (p1 + (^(p1)))))) && (((fi8 | (fi7 % (((fi7 / (((^(fi7))) | 1))) | 1))) != -778) || (-87 != int8(p0)))) {
                        var v10 int16 = 0
                        _ = v10
                    } else {
                        v5 = append(v5, (fi7 << ((255) & 63)))
                    }
                }
            } else {
                if (((^(p2)) >= p2) || (0 < 32767)) {
                    if true {
                        var v11 uint32 = 2
                        var v12 int16 = 1000
                        p1 += p1
                        v13 := GenS1{F0: (-(p1))}
                        v12 = v12
                        v5 = append(v5, fi8)
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
                var v14 float64 = (0.1 - ((float64(3.14159) - 0.1) / 0.1))
                var v15 int = 580
                _ = v14
                _ = v15
            }
            _ = fi8
        }
        switch v6 {
        case "hello":
            v16 := GenS1{F0: -128}
            _ = v16
        default:
            p2 = 2147483647
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{13, 4294967295, 115, 805}
    var v2 int32 = 42
    var v3 uint32 = 2147483647
    var v4 int = 222
    var v17 uint16 = 2
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            if (!((false && (32767 >= fi18)) && true)) {
                fmt.Println(fi19, v4)
            } else {
                if true {
                    var v20 uint16 = (^(v17))
                    switch 32768 {
                    case 128:
                        switch 32767 {
                        case -116:
                            if (629 == 42) {
                                v21 := GenS2{F0: 42, F1: 785}
                                _ = v21
                            } else {
                                _ = len(v1)
                                if (0 < -1) {
                                    var v22 int32 = (-(v2))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v22
                                } else {
                                    var v23 uint32 = v3
                                    if false {
                                        v17 += (^(v20))
                                        var v24 float32 = ((0.0 * (0.1 * ((1e10 + float32(((0.1 - 3.14159) - 3.14159))) / ((float32(3.14159) * 0.0) + (0.1 * -2.25))))) - ((3.14159 - -2.25) - 0.0))
                                        _ = v24
                                    } else {
                                        v4 -= fi19
                                        if false {
                                            v25 := GenS3{F0: (^(fi19))}
                                            switch (-(v4)) {
                                            case -2147483648:
                                                v26 := GenS1{F0: 73}
                                                _ = v26
                                            case 32767:
                                                v3 = (^(v23))
                                            }
                                            var v27 string = "foo bar"
                                            _ = v25
                                            _ = v27
                                        }
                                    }
                                    switch 42 {
                                    case 728:
                                        if (0 >= 842) {
                                            switch ("foo bar" + "hello") {
                                            case "a":
                                                v28 := GenS3{F0: fi18}
                                                _ = v28
                                            }
                                            fmt.Println(v2)
                                        } else {
                                            var v29 int64 = -9223372036854775808
                                            _ = v29
                                        }
                                    case 127:
                                        v17 += (v17 * v20)
                                    default:
                                        _ = len(v1)
                                    }
                                    _ = v23
                                }
                            }
                        default:
                            v20 -= (v17 + (v17 | (v20 % (((v20 >> ((1) & 63))) | 1))))
                        }
                    case 807:
                        v3 += (^(v3))
                    }
                    _ = v20
                } else {
                    var v30 float64 = float64(-2.25)
                    var v31 float64 = 3.14159
                    var v32 string = ""
                    fmt.Println(v3)
                    v33 := GenS2{F0: 118, F1: 929}
                    v34 := GenS2{F0: 76, F1: 255}
                    if true {
                        _ = len(v1)
                        var v35 float32 = -2.25
                        _ = v35
                    }
                    _ = v30
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                switch fi19 {
                case -27:
                    fmt.Println(fi18, v2)
                case 9223372036854775807:
                    v36 := GenS1{F0: 127}
                    v1 = append(v1, (^(v3)))
                    _ = v36
                }
                var v37 int16 = 0
                _ = v37
            }
            if ((!false) && (2 > 127)) {
                var v38 int = fi19
                fmt.Println(v38, fi19, v2)
                v4 += int(v2)
                _ = v38
            } else {
                var v39 uint64 = 1
                var v40 uint64 = uint64(v3)
                _ = v39
                _ = v40
            }
            if ((v2 * (-(v2))) < v2) {
                var v41 int32 = (v2 | (v2 % ((genFn0(uint8(v17), 1, v3)) | 1)))
                v42 := GenS2{F0: 128, F1: v17}
                if len(v1) > 0 {
                    v1[0] = 42
                }
                v43 := GenS4{F0: (v41 % (((^(v41))) | 1)), F1: v3}
                _ = v41
                _ = v42
                _ = v43
            } else {
                if true {
                    var v44 int16 = 42
                    switch (v4 << ((65535) & 63)) {
                    case 4294967295:
                        fmt.Println(fi18, v44)
                    default:
                        if ((^(v3)) < (v3 * v3)) {
                            v44 = 1000
                            var v45 int64 = -9223372036854775808
                            v46 := GenS4{F0: v2, F1: (v3 | v3)}
                            var v47 int32 = (v2 ^ 653)
                            _ = v45
                            _ = v46
                            _ = v47
                        } else {
                            fmt.Println(v4)
                        }
                    }
                    _ = v44
                }
                v1 = append(v1, v3)
            }
            _ = fi19
        }
        _ = fi18
    }
    v48 := GenS5{F0: 1, F1: false, F2: 9223372036854775807}
    if (uint16(v3) != v17) {
        if (356 <= 1) {
            for fi49 := 0; fi49 < 4; fi49++ {
                var v50 uint8 = uint8(v3)
                fmt.Println(fi49, v17, v17)
                fmt.Println(v2, v4)
                _ = fi49
                _ = v50
            }
        } else {
            fmt.Println(v2, v2, v17)
        }
        if (-128 >= 100) {
            v17 -= v17
            v51 := GenS6{F0: "hello", F1: -57, F2: 127}
            _ = v51
        }
        fmt.Println(v17, v2, v17)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v4, v2, v17)
    fmt.Println(v4, v4)
    var v52 int16 = 1000
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = (v2 - v2)
        v53 := GenS7{F0: -96, F1: ((v17 == v17) || (-9223372036854775808 == 1))}
        _ = v53
    }
    v54 := GenS7{F0: 127, F1: ((v3 | v3) <= 2147483647)}
    for fi55 := 0; fi55 < 19; fi55++ {
        v4 = (v4 / (((v4 << ((4294967295) & 63))) | 1))
        v52 = (v52 | v52)
        switch "hello" {
        case "foo bar":
            for fi56 := 0; fi56 < 9; fi56++ {
                if true {
                    var v57 uint8 = 255
                    if false {
                        switch (v4 - (fi56 / ((v4) | 1))) {
                        case -410:
                            if len(v1) > 0 {
                                v1[0] = (v3 ^ (v3 % (((v3 - v3)) | 1)))
                            }
                            v58 := GenS2{F0: v57, F1: 42}
                            v52 = v52
                            _ = v58
                        case 127:
                            var v59 int64 = -433
                            _ = v59
                        case 309:
                            switch (v4 * (^(fi56))) {
                            case 0:
                                _ = len(v1)
                            default:
                                var v60 float64 = (1.5 + float64(1.5))
                                _ = v60
                            }
                            switch (v4 % (((v4 - (fi55 - fi56))) | 1)) {
                            case 211:
                                var v61 uint64 = 477
                                var v62 uint8 = v57
                                _ = len(v1)
                                _ = v61
                                _ = v62
                            case -1:
                                v57 -= 2
                            case -615:
                                var v63 bool = (-2147483648 > (v4 | (fi55 >> ((48) & 63))))
                                _ = v63
                            default:
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                            }
                        }
                        fmt.Println(v3, v3, v17)
                        var v64 string = "\n\t"
                        _ = v64
                    } else {
                        var v65 int8 = int8(v2)
                        fmt.Println(v52, fi55)
                        v4 = fi55
                        _ = v65
                    }
                    fmt.Println(v17)
                    var v66 float32 = 1.5
                    _ = v57
                    _ = v66
                } else {
                    v67 := GenS6{F0: "a", F1: 0, F2: 1}
                    if ((-2.25 == (1e10 + (1.5 + 3.14159))) && false) {
                        fmt.Println(fi56, fi55, v52)
                        v4 = (fi56 % ((fi55) | 1))
                    } else {
                        var v68 uint8 = 2
                        var v69 int32 = (v2 & (v2 ^ 0))
                        _ = v68
                        _ = v69
                    }
                    v2 = genFn0(uint8(v17), 127, (v3 ^ (v3 - 25)))
                    _ = v67
                }
                switch fi56 {
                case -2147483648:
                    var v70 uint16 = (v17 % ((v17) | 1))
                    _ = v70
                case 770:
                    if true {
                        v71 := GenS8{F0: (v52 == (^(v52))), F1: -1, F2: -253}
                        _ = v71
                    }
                case 4294967295:
                    fmt.Println(fi55)
                default:
                    v72 := GenS4{F0: genFn0(240, -1, 271), F1: (^(v3))}
                    _ = v72
                }
                if true {
                    v73 := GenS9{F0: "hello", F1: true, F2: ("foo bar" + ("hello" + ("a" + ((("" + ("hello" + ("\n\t" + ("foo bar" + "foo bar")))) + ("\n\t" + "")) + ""))))}
                    v2 = (-(v2))
                    fmt.Println(fi56)
                    _ = v73
                }
                _ = fi56
            }
        case "a":
            for fi74 := 0; fi74 < 17; fi74++ {
                v52 -= (v52 & v52)
                _ = fi74
            }
        case "":
            var v75 uint16 = 385
            _ = v75
        default:
            v1 = append(v1, (v3 - v3))
        }
        _ = fi55
    }
    var v76 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v52
    _ = v54
    _ = v76
}
