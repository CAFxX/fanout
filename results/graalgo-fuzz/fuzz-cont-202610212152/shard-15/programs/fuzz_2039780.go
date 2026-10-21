package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 string }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 string; F1 int; F2 int64 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint32 = (v1 << ((255) & 63))
    v3 := make(map[int]string)
    v4 := []string{}
    var v5 uint = 18446744073709551615
    if (true || (-1 <= 100)) {
        v6 := GenS1{F0: 627, F1: 65535, F2: int8(v5)}
        _ = v6
    } else {
        fmt.Println(v2, v5)
        var v7 int32 = -744
        _ = v7
    }
    v1 = uint32(v5)
    v5 -= 18446744073709551615
    if (0 <= -807) {
        v8 := GenS2{F0: 255, F1: 65535, F2: "a"}
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int32 = -2147483648
                _ = fi9
                _ = v10
            }
            if false {
                v11 := GenS1{F0: 42, F1: 42, F2: 99}
                v5 -= (v5 * (v5 >> (((v5 | v5)) & 63)))
                _ = v11
            } else {
                if true {
                    var v12 float64 = 3.14159
                    for fi13 := 0; fi13 < 12; fi13++ {
                        var v14 int16 = 0
                        v5 += (v5 >> ((v5) & 63))
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                } else {
                    v15 := GenS3{F0: 738, F1: 32767}
                    v3[4294967295] = ("" + ((("" + "hello") + "\n\t") + ("foo bar" + ("a" + ""))))
                    _ = v15
                }
                for fi16 := 0; fi16 < 10; fi16++ {
                    v2 -= v2
                    _ = fi16
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = len(v4)
        }
        _ = v8
    } else {
        for fi17 := 0; fi17 < 15; fi17++ {
            if (false && (42 != uint16(v5))) {
                if false {
                    switch "hello" {
                    case "":
                        var v18 int16 = 703
                        _ = v18
                    case "hello":
                        fmt.Println(v2, v5, v5)
                    default:
                        v19 := GenS2{F0: 432, F1: 65535, F2: "hello"}
                        _ = v19
                    }
                }
                fmt.Println(v5, fi17)
            } else {
                v20 := GenS4{F0: (v5 & v5)}
                var v21 bool = (((0.1 * (0.1 - (float32(((float64(((-2.25 / (float32((0.0 + float64((1.5 + ((0.1 - (0.1 + 1e10)) * (-2.25 * (-2.25 - float32(0.0)))))))) / -2.25)) * (float32((0.0 - ((float64(((0.0 * float32((((-2.25 - (0.0 - -2.25)) - (-2.25 * 1.5)) * 1e10))) + (1.5 + 1.5))) + 1.5) + 1.5))) + 1.5))) / (1.5 * (1.5 - (0.0 - ((3.14159 + 3.14159) * 0.1))))) * 1.5)) * -2.25))) + 0.1) == (3.14159 * -2.25))
                var v22 uint64 = 669
                _ = v20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 2; fi23++ {
                switch 32768 {
                case -830:
                    if (!false) {
                        var v24 int16 = 0
                        v25 := v3[(fi17 << ((1) & 63))]
                        _ = v24
                        _ = v25
                    } else {
                        fmt.Println(v1, fi17, v2)
                    }
                case -448:
                    var v26 uint16 = 0
                    _ = v26
                case 524:
                    var v27 bool = (v2 != (v2 * (v1 ^ v2)))
                    _ = v27
                }
                var v28 int8 = 100
                _ = fi23
                _ = v28
            }
            _ = fi17
        }
    }
    _ = len(v3)
    var v29 uint16 = 32767
    v1 -= v1
    v30 := GenS4{F0: v5}
    switch "\n\t" {
    case "hello":
        v1 += 429
    case "\n\t":
        if (("hello" + "a") > "hello") {
            v31 := GenS5{F0: -128, F1: 0}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 int16 = -577
                v33 = -32768
                fmt.Println(fi32)
                _ = fi32
                _ = v33
            }
            if false {
                for fi34 := 0; fi34 < 3; fi34++ {
                    v5 = v5
                    if (1000 == 0) {
                        var v35 uint16 = (v29 + v29)
                        _ = v35
                    } else {
                        if ((-2.25 >= -2.25) || ((128 != fi34) && true)) {
                            v3[fi34] = "foo bar"
                            if (((2147483647 != 354) && (false || (4 <= 0))) && ((v1 << (((^(v5))) & 63)) > 358)) {
                                v5 -= v5
                                var v36 int32 = int32(v2)
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 17; fi37++ {
                                var v38 int64 = 42
                                v29 += (v29 * 665)
                                _ = fi37
                                _ = v38
                            }
                        }
                        _ = len(v4)
                        v29 += v29
                    }
                    _ = fi34
                }
                v29 = 255
                if (false || ("foo bar" >= "\n\t")) {
                    var v39 float64 = 0.1
                    switch "\n\t" {
                    case "foo bar":
                        for fi40 := 0; fi40 < 14; fi40++ {
                            var v41 uint = v5
                            if false {
                                if (-544 == 0) {
                                    for fi42 := 0; fi42 < 2; fi42++ {
                                        var v43 uint = v41
                                        var v44 uint8 = uint8(v29)
                                        _ = fi42
                                        _ = v43
                                        _ = v44
                                    }
                                    switch fi40 {
                                    case 128:
                                        for fi45 := 0; fi45 < 6; fi45++ {
                                            var v46 uint8 = 224
                                            v1 = (v1 << ((v41) & 63))
                                            v39 = v39
                                            var v47 uint64 = 2
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                    default:
                                        v39 = (v39 + ((-2.25 / -2.25) / 1.5))
                                    }
                                    _ = len(v4)
                                } else {
                                    var v48 uint16 = (v29 ^ v29)
                                    v39 -= (v39 + v39)
                                    _ = v48
                                }
                                switch (fi40 % (((-(fi40))) | 1)) {
                                case 4294967295:
                                    if (!(1.5 >= 0.0)) {
                                        v41 = 23
                                        v41 = v41
                                    } else {
                                        v39 = float64(((-2.25 + 1.5) - (1e10 * float32((1e10 / v39)))))
                                        if (65535 > (^(v1))) {
                                            v2 = (^(v1))
                                        }
                                    }
                                case 2147483647:
                                    var v49 int16 = 42
                                    _ = v49
                                }
                            }
                            delete(v3, 32767)
                            _ = fi40
                            _ = v41
                        }
                    case "hello":
                        if (2 != 0) {
                            var v50 uint8 = 2
                            v51 := GenS2{F0: v29, F1: (^(v29)), F2: "hello"}
                            _ = v50
                            _ = v51
                        }
                    default:
                        if false {
                            v29 += v29
                        }
                    }
                    _ = v39
                } else {
                    var v52 int = -177
                    _ = v52
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            }
        }
    case "":
        v1 = (v1 * v2)
    }
    fmt.Println(v29, v5, v1)
    var v53 float32 = float32(1.5)
    var v54 string = ("a" + "a")
    v29 += v29
    var v55 string = (v54 + v54)
    if true {
        var v56 int16 = -1
        if false {
            fmt.Println(v56, v1, v55)
        }
        var v57 uint8 = 0
        _ = v56
        _ = v57
    } else {
        v55 = v54
    }
    var v58 float64 = (1.5 * 3.14159)
    var v59 uint64 = 2
    if (v1 != (v2 ^ (v2 - (v2 % ((v2) | 1))))) {
        for fi60 := 0; fi60 < 15; fi60++ {
            if true {
                v61 := GenS4{F0: 65535}
                if (((float64(((((-2.25 * float32(v58)) / -2.25) + (-2.25 + float32(v58))) / v53)) / 1.0) - ((float64(0.0) + (v58 - float64(-2.25))) * 1e10)) > (((((v58 - v58) / v58) - float64(3.14159)) * (1.5 * v58)) + ((v58 / v58) * 1e10))) {
                    v54 = "foo bar"
                } else {
                    v4 = append(v4, "\n\t")
                }
                if false {
                    fmt.Println(fi60)
                    switch (-(fi60)) {
                    case 127:
                        v29 = (v29 - v29)
                    case 0:
                        var v62 string = (("a" + v54) + v55)
                        _ = v62
                    case 32767:
                        delete(v3, fi60)
                    }
                } else {
                    fmt.Println(v58)
                    v63 := GenS6{F0: (("foo bar" + ((("foo bar" + v54) + "") + (v54 + "hello"))) + (v54 + "a")), F1: 1, F2: int64(v59)}
                    _ = v63
                }
                _ = v61
            }
            for fi64 := 0; fi64 < 10; fi64++ {
                v65 := GenS7{F0: 32767}
                fmt.Println(v5, v58)
                _ = fi64
                _ = v65
            }
            _ = fi60
        }
        _ = len(v3)
        if true {
            for fi66 := 0; fi66 < 14; fi66++ {
                var v67 int32 = 42
                _ = fi66
                _ = v67
            }
        } else {
            v55 = v54
            _ = len(v3)
        }
    } else {
        for fi68 := 0; fi68 < 5; fi68++ {
            if false {
                if true {
                    v1 = 1
                }
                _ = len(v4)
                switch v54 {
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    for fi69 := 0; fi69 < 11; fi69++ {
                        _ = len(v4)
                        _ = fi69
                    }
                default:
                    v53 -= v53
                }
            } else {
                v3[(fi68 + fi68)] = (((v55 + v54) + ("\n\t" + v54)) + v55)
            }
            v29 = (v29 ^ v29)
            var v70 bool = (fi68 > fi68)
            _ = fi68
            _ = v70
        }
        v71 := GenS2{F0: (v29 & (^(v29))), F1: (v29 % (((v29 & (v29 / ((v29) | 1)))) | 1)), F2: v55}
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v58
    _ = v59
}
