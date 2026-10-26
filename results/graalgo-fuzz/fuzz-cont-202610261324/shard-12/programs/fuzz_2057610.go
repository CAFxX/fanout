package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = 2147483647
    var v3 uint16 = uint16(v1)
    v4 := make(map[string]int)
    v1 += 4294967295
    var v5 float32 = ((((-2.25 - (1e10 - (3.14159 + 0.1))) * 0.0) * 1e10) + float32(0.1))
    fmt.Println(v1, v1, v3)
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int64 = 9223372036854775807
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := []int{(-(v2))}
    if (true && false) {
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 uint8 = 2
                switch "" {
                case "":
                    if (!(561 == 1)) {
                        if (!((1.5 / -2.25) != ((1.5 - (float64((((((0.1 * 1.5) / (v5 + -2.25)) - v5) - v5) - float32((0.1 * (0.0 / -2.25))))) - (0.1 + 1e10))) * 0.1))) {
                            if len(v9) > 0 {
                                v9[0] = fi11
                            }
                        } else {
                            _ = len(v9)
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v13 uint8 = 128
                            if true {
                                v14 := GenS1{F0: v3, F1: 0}
                                v15 := GenS1{F0: (v3 / ((v3) | 1)), F1: (v13 / ((v13) | 1))}
                                _ = v14
                                _ = v15
                            }
                            _ = v13
                        }
                    }
                }
                var v16 uint32 = uint32(v3)
                _ = fi11
                _ = v12
                _ = v16
            }
            v2 += v2
            _ = fi10
        }
        if ((1.5 < (3.14159 - v5)) || (-1 <= 42)) {
            switch (v2 / (((v2 % (((-(v2))) | 1))) | 1)) {
            case 0:
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 float64 = (3.14159 + float64(0.0))
                    var v19 string = "foo bar"
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case -157:
                v20 := GenS1{F0: (^(v3)), F1: 162}
                _ = v20
            }
            fmt.Println(v3, v5, v5)
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := GenS2{F0: -32768}
                v3 = (v3 * v3)
                _ = fi21
                _ = v22
            }
        } else {
            _ = len(v9)
        }
    } else {
        _ = len(v9)
        if false {
            var v23 int16 = -654
            v24 := GenS1{F0: v3, F1: 255}
            _ = v23
            _ = v24
        }
        v4[(("foo bar" + ("hello" + ("foo bar" + ("foo bar" + "a")))) + "a")] = 0
    }
    switch v2 {
    case 9223372036854775807:
        for fi25 := 0; fi25 < 19; fi25++ {
            fmt.Println(v5, v2, v5)
            _ = fi25
        }
    case 128:
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := GenS1{F0: (v3 - v3), F1: 0}
            var v28 int8 = -89
            switch (fi26 % ((fi26) | 1)) {
            case 127:
                v29 := v4["hello"]
                _ = v29
            case -647:
                fmt.Println(fi26, fi26, v5)
            case 42:
                switch (fi26 >> ((0) & 63)) {
                case -9223372036854775808:
                    v28 -= -1
                default:
                    if (true && true) {
                        var v30 int8 = -128
                        var v31 uint32 = 2147483648
                        for fi32 := 0; fi32 < 7; fi32++ {
                            v31 -= (v31 + 1)
                            if len(v9) > 0 {
                                v9[0] = -251
                            }
                            _ = fi32
                        }
                        _ = v30
                        _ = v31
                    }
                }
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    case 0:
        v33 := GenS2{F0: 302}
        _ = v33
    }
    v34 := GenS3{F0: 100, F1: 0}
    if len(v9) > 0 {
        v9[0] = 361
    }
    v2 = (v2 % ((42) | 1))
    var v35 int64 = 409
    fmt.Println(v3)
    if (!(9223372036854775807 > (v35 % (((v35 >> ((42) & 63))) | 1)))) {
        v3 = (^(v3))
    } else {
        if (v1 >= (v1 | (v1 | (v1 & (^(v1)))))) {
            if false {
                for fi36 := 0; fi36 < 8; fi36++ {
                    if true {
                        if false {
                            if (!(582 < 2147483648)) {
                                v37 := GenS4{F0: -1}
                                _ = v37
                            }
                            fmt.Println(v35, v35)
                        }
                        _ = len(v4)
                    } else {
                        v5 = ((float32((1e10 / (0.1 / 3.14159))) * (v5 - v5)) - v5)
                    }
                    v1 += 0
                    if len(v9) > 0 {
                        v9[0] = fi36
                    }
                    _ = fi36
                }
                fmt.Println(v1, v35)
                if (!false) {
                    fmt.Println(v2, v3)
                    v1 += (v1 >> ((1) & 63))
                } else {
                    v38 := GenS4{F0: 87}
                    fmt.Println(v5)
                    _ = v38
                }
            } else {
                fmt.Println(v1, v5, v35)
                v39 := GenS2{F0: 32767}
                fmt.Println(v35, v35, v35)
                _ = v39
            }
            v35 += v35
        } else {
            var v40 int32 = -1
            switch v2 {
            case 127:
                v3 = (v3 - 65535)
            case 4294967295:
                var v41 bool = (false && true)
                _ = v41
            default:
                for fi42 := 0; fi42 < 19; fi42++ {
                    var v43 uint64 = 18446744073709551615
                    if (false || false) {
                        if true {
                            v2 += (v2 - (^(fi42)))
                            delete(v4, ("" + "\n\t"))
                        } else {
                            v44 := GenS3{F0: -128, F1: 42}
                            v43 = (v1 * (v43 % ((938) | 1)))
                            _ = v44
                        }
                        v4["foo bar"] = (v2 | (v2 ^ 804))
                    } else {
                        if (9223372036854775807 > v35) {
                            var v45 int8 = -1
                            if (810 != 2147483647) {
                                if (("\n\t" + "a") > ("a" + "hello")) {
                                    for fi46 := 0; fi46 < 10; fi46++ {
                                        var v47 string = ("a" + (("\n\t" + "foo bar") + "foo bar"))
                                        delete(v4, v47)
                                        _ = fi46
                                        _ = v47
                                    }
                                    for fi48 := 0; fi48 < 12; fi48++ {
                                        fmt.Println(v35, v45)
                                        v3 += (v3 + 473)
                                        var v49 int32 = -957
                                        switch v2 {
                                        case 32767:
                                            switch (fi42 & (v2 * 128)) {
                                            case -16:
                                                var v50 int64 = (v35 >> ((328) & 63))
                                                _ = v50
                                            case 9223372036854775807:
                                                v40 = v40
                                            case 0:
                                                var v51 float64 = 0.1
                                                _ = v51
                                            default:
                                                var v52 int8 = v45
                                                _ = v52
                                            }
                                        default:
                                            var v53 uint64 = (v43 - 18446744073709551615)
                                            _ = v53
                                        }
                                        _ = fi48
                                        _ = v49
                                    }
                                    if ((!((128 < 127) && ((v1 <= 18446744073709551615) || false))) && ((^(v2)) >= 2147483647)) {
                                        if true {
                                            var v54 int32 = 2147483647
                                            fmt.Println(v54)
                                            v5 += 3.14159
                                            _ = v54
                                        }
                                        if true {
                                            fmt.Println(fi42, v40)
                                            if (1 != 1000) {
                                                v43 -= v43
                                            } else {
                                                var v55 int64 = v35
                                                for fi56 := 0; fi56 < 10; fi56++ {
                                                    if ((((^(v35)) != (v35 ^ (v55 ^ v55))) && (v40 >= (v40 - -2147483648))) && (!((!false) && (18446744073709551615 == 1)))) {
                                                        fmt.Println(v43, v55, fi56)
                                                        var v57 uint8 = 0
                                                        _ = v57
                                                    }
                                                    var v58 uint32 = 2147483648
                                                    _ = fi56
                                                    _ = v58
                                                }
                                                v3 -= (v3 - v3)
                                                _ = v55
                                            }
                                            v5 = v5
                                        } else {
                                            if (3.14159 != -2.25) {
                                                var v59 int64 = (v35 >> ((0) & 63))
                                                switch "a" {
                                                case "foo bar":
                                                    for fi60 := 0; fi60 < 16; fi60++ {
                                                        v1 += 461
                                                        if true {
                                                            var v61 int = int(v59)
                                                            if len(v9) > 0 {
                                                                v9[0] = 32768
                                                            }
                                                            _ = v61
                                                        } else {
                                                            fmt.Println(v45, v59)
                                                            v45 = (v45 & -1)
                                                        }
                                                        _ = fi60
                                                    }
                                                }
                                                _ = v59
                                            } else {
                                                var v62 float64 = 1e10
                                                for fi63 := 0; fi63 < 10; fi63++ {
                                                    var v64 uint16 = 2
                                                    switch fi42 {
                                                    case -150:
                                                        v35 = (v35 & (v35 + v35))
                                                        if true {
                                                            fmt.Println(fi63)
                                                        }
                                                    case -2147483648:
                                                        v65 := GenS3{F0: (v45 + int8(v40)), F1: 1000}
                                                        _ = v65
                                                    case 2147483647:
                                                        if true {
                                                            fmt.Println(v64, v35)
                                                            delete(v4, ("a" + "a"))
                                                        } else {
                                                            var v66 int8 = (-(v45))
                                                            _ = v66
                                                        }
                                                    }
                                                    _ = fi63
                                                    _ = v64
                                                }
                                                fmt.Println(v45)
                                                _ = v62
                                            }
                                            var v67 string = "a"
                                            _ = v67
                                        }
                                        var v68 uint64 = 274
                                        _ = v68
                                    } else {
                                        switch fi42 {
                                        case 0:
                                            for fi69 := 0; fi69 < 16; fi69++ {
                                                var v70 uint8 = 147
                                                v3 = v3
                                                _ = fi69
                                                _ = v70
                                            }
                                        case 128:
                                            var v71 bool = (!false)
                                            _ = v71
                                        case 32767:
                                            v35 -= v35
                                        default:
                                            if ((3.14159 * (0.1 - (1e10 + (0.1 * 0.0)))) == float64((((1e10 + (0.0 * (0.1 + 0.1))) + v5) / v5))) {
                                                if (!("\n\t" == "")) {
                                                    var v72 int = (v2 - 2147483647)
                                                    v1 -= 4294967295
                                                    _ = v72
                                                } else {
                                                    if true {
                                                        v73 := v4["foo bar"]
                                                        v2 += -621
                                                        _ = v73
                                                    } else {
                                                        if len(v9) > 0 {
                                                            v9[0] = v2
                                                        }
                                                    }
                                                    var v74 int64 = (-(v35))
                                                    var v75 int32 = v40
                                                    _ = v74
                                                    _ = v75
                                                }
                                                v3 = v3
                                            } else {
                                                v45 += v45
                                                var v76 uint64 = 2
                                                v1 = v43
                                                _ = v76
                                            }
                                        }
                                        v77 := GenS3{F0: v45, F1: -1}
                                        _ = v77
                                    }
                                } else {
                                    v3 = (v3 % ((2) | 1))
                                }
                                v3 = 0
                            } else {
                                v43 = (v43 + 137)
                                v43 += v1
                            }
                            v9 = append(v9, (v2 * fi42))
                            _ = v45
                        } else {
                            v5 = (0.0 * v5)
                            if ((!(!false)) || ((false || false) && (0 < 740))) {
                                var v78 string = (("" + "foo bar") + ("\n\t" + (("" + "a") + ("foo bar" + ("hello" + ("\n\t" + "a"))))))
                                _ = v78
                            }
                        }
                    }
                    _ = fi42
                    _ = v43
                }
            }
            v9 = append(v9, 42)
            _ = v40
        }
    }
    switch "\n\t" {
    case "\n\t":
        v3 -= (v3 - (v3 ^ (v3 & v3)))
    case "a":
        fmt.Println(v1, v5, v5)
    default:
        var v79 int8 = -128
        _ = v79
    }
    var v80 float64 = float64(((float32(float64(3.14159)) - v5) + (v5 + (v5 + v5))))
    fmt.Println(v35)
    _ = len(v9)
    if false {
        var v81 int8 = -1
        var v82 int8 = 127
        _ = v81
        _ = v82
    } else {
        v1 = 58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v80
    _ = v9
}
