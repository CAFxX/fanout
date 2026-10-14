package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int16; F1 bool; F2 int16 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint8, p2 int16) int {
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(p2, fi6, p2)
        _ = fi6
    }
    v7 := GenS1{F0: (p0 & (p0 & p0)), F1: (false && false), F2: 104}
    var v8 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -493
}
func genFn1(p0 uint32, p1 uint64, p2 uint64) uint {
    if false {
        var v9 string = (("a" + "a") + "")
        p0 = 201
        _ = v9
    } else {
        p1 = (p1 & (p2 % (((^(p2))) | 1)))
        p0 += p0
        switch "a" {
        case "\n\t":
            var v10 uint8 = 128
            _ = v10
        case "foo bar":
            var v11 uint64 = (p1 % (((^(p2))) | 1))
            _ = v11
        case "hello":
            fmt.Println(p1, p2)
        }
    }
    v12 := []int{genFn0((^(p0)), 1, -431), -107, 4294967295, 707}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 float64 = 0.1
    var v2 int = -238
    var v3 float32 = 1.5
    var v4 int16 = 1
    var v5 bool = (172 > 1)
    var v13 uint64 = 682
    switch ("foo bar" + ("hello" + "a")) {
    case "foo bar":
        v14 := GenS2{F0: 255, F1: v13, F2: 255}
        _ = v14
    case "\n\t":
        v5 = (-2147483648 > -2147483648)
    case "":
        var v15 uint64 = v13
        _ = v15
    }
    switch (^(v2)) {
    case -763:
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                if v5 {
                    if false {
                        v18 := GenS2{F0: 2, F1: v13, F2: 32767}
                        v3 += 0.1
                        if (!(!(v4 <= (-(v4))))) {
                            v19 := []int{fi16, 127}
                            v4 = v4
                            v20 := GenS3{F0: -1, F1: (("foo bar" + "hello") + "\n\t")}
                            _ = v19
                            _ = v20
                        } else {
                            v21 := []int{fi17, genFn0(2147483647, 128, 42)}
                            var v22 string = (("a" + ("" + (("a" + ("\n\t" + "foo bar")) + "a"))) + "\n\t")
                            _ = v21
                            _ = v22
                        }
                        _ = v18
                    } else {
                        if (!true) {
                            switch "hello" {
                            case "hello":
                                var v23 int64 = 0
                                _ = v23
                            }
                        } else {
                            v1 = v1
                            fmt.Println(v4)
                            v13 -= v13
                            v24 := GenS1{F0: 51, F1: v5, F2: -128}
                            v1 += v1
                            _ = v24
                        }
                        if v5 {
                            v4 = v4
                            if v5 {
                                var v25 float32 = v3
                                v26 := GenS3{F0: 7, F1: ""}
                                _ = v25
                                _ = v26
                            } else {
                                v1 -= (float64(v3) + v1)
                                switch (fi16 / ((genFn0(766, 2, -32768)) | 1)) {
                                case 1:
                                    v27 := GenS3{F0: -65, F1: ("hello" + "foo bar")}
                                    _ = v27
                                case 2147483647:
                                    v3 -= v3
                                default:
                                    if (273 != 9223372036854775807) {
                                        var v28 uint16 = 42
                                        _ = v28
                                    }
                                }
                            }
                            var v29 int = (^(fi16))
                            _ = v29
                        } else {
                            var v30 float64 = v1
                            fmt.Println(v4)
                            var v31 uint8 = 0
                            _ = v30
                            _ = v31
                        }
                        var v32 int8 = 0
                        _ = v32
                    }
                } else {
                    if (!((v5 || (2 >= 255)) || true)) {
                        var v33 uint32 = 811
                        var v34 uint16 = 2
                        var v35 float32 = -2.25
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                v36 := GenS3{F0: -118, F1: "\n\t"}
                if ((1 > v4) || v5) {
                    var v37 int16 = 0
                    _ = v37
                }
                _ = fi17
                _ = v36
            }
            v38 := []int{}
            _ = fi16
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        fmt.Println(v5)
        for fi40 := 0; fi40 < 4; fi40++ {
            v2 = fi40
            _ = fi40
        }
        var v41 float64 = v1
        _ = fi39
        _ = v41
    }
    v42 := GenS2{F0: 623, F1: 18446744073709551615, F2: 2}
    if v5 {
        v43 := GenS4{F0: (v13 >> ((42) & 63)), F1: 18446744073709551615}
        var v44 int32 = -2147483648
        var v45 float64 = (0.0 + v1)
        _ = v43
        _ = v44
        _ = v45
    } else {
        var v46 string = (("foo bar" + "\n\t") + "foo bar")
        v46 = ("" + ((v46 + ((v46 + v46) + "\n\t")) + "foo bar"))
        fmt.Println(v3)
        _ = v46
    }
    if (v5 && (100 > 81)) {
        var v47 int32 = 32
        _ = v47
    } else {
        v4 = v4
    }
    if (v4 > (v4 - (v4 / ((-32768) | 1)))) {
        v3 = v3
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        switch "" {
        case "hello":
            v13 = (v13 & (v13 + 2))
        case "foo bar":
            var v49 int = 9223372036854775807
            _ = v49
        case "a":
            if (!(v13 >= (v13 / ((0) | 1)))) {
                v50 := []int{v2, fi48, fi48, 128}
                _ = v50
            } else {
                switch 32767 {
                case 4294967295:
                    switch genFn0(984, 192, (v4 >> ((42) & 63))) {
                    case -2147483648:
                        var v51 uint64 = v13
                        _ = v51
                    case 32768:
                        v52 := GenS3{F0: 79, F1: ("hello" + "foo bar")}
                        _ = v52
                    case 32767:
                        if ((false || ((v4 | v4) > -1)) && ("hello" != "\n\t")) {
                            v53 := []int{-556, (v2 * -1), genFn0(2147483647, 121, 0), genFn0(65535, 127, (v4 >> ((259) & 63)))}
                            if ((((("" + ("foo bar" + "a")) + ("hello" + "a")) + "a") != "a") || v5) {
                                if len(v53) > 0 {
                                    v53[0] = genFn0(19, 255, -341)
                                }
                                switch 35 {
                                case 127:
                                    var v54 int32 = 174
                                    _ = v54
                                }
                            }
                            _ = v53
                        }
                    default:
                        v4 -= -1
                    }
                case 600:
                    fmt.Println(v13, v5)
                }
            }
        default:
            switch "foo bar" {
            case "":
                v55 := []int{genFn0(2147483648, 138, (v4 * v4))}
                _ = v55
            }
        }
        if (2147483647 >= 1) {
            v1 -= 3.14159
        } else {
            for fi56 := 0; fi56 < 18; fi56++ {
                v57 := []int{}
                _ = fi56
                _ = v57
            }
        }
        var v58 uint32 = 42
        _ = fi48
        _ = v58
    }
    var v59 string = "a"
    fmt.Println(v59, v4, v5)
    v5 = v5
    v59 = "a"
    fmt.Println(v59)
    v13 = (v13 << ((0) & 63))
    switch "a" {
    case "foo bar":
        v2 += genFn0(1, 0, (v4 << ((65535) & 63)))
    case "":
        if false {
            v60 := GenS5{F0: 220, F1: 70}
            for fi61 := 0; fi61 < 16; fi61++ {
                v13 = v13
                _ = fi61
            }
            if (42 >= 0) {
                fmt.Println(v3)
            } else {
                fmt.Println(v4, v4, v2)
            }
            _ = v60
        } else {
            if (!v5) {
                var v62 string = (v59 + ("foo bar" + (v59 + (v59 + v59))))
                v63 := GenS1{F0: 650, F1: (!true), F2: -28}
                v64 := []int{}
                _ = v62
                _ = v63
                _ = v64
            } else {
                v65 := []int{int(v13), (v2 ^ v2), (v2 + 9223372036854775807), v2}
                _ = v65
            }
            var v66 float64 = v1
            _ = v66
        }
    case "\n\t":
        v67 := GenS6{F0: (v4 + (v4 >> ((2) & 63))), F1: (v1 == v1), F2: 1000}
        _ = v67
    default:
        if v5 {
            if (v5 || v5) {
                v68 := GenS7{F0: 690}
                v59 = ("foo bar" + "foo bar")
                v69 := GenS1{F0: uint32(v13), F1: (1 >= -300), F2: -128}
                _ = v68
                _ = v69
            } else {
                if v5 {
                    fmt.Println(v2, v3)
                }
                fmt.Println(v3, v1, v13)
            }
            for fi70 := 0; fi70 < 14; fi70++ {
                fmt.Println(v2, fi70, v59)
                v71 := []int{(^(fi70)), fi70}
                _ = fi70
                _ = v71
            }
            var v72 int64 = 481
            _ = v72
        } else {
            v73 := []int{(v2 << ((956) & 63)), v2, (v2 - v2)}
            _ = v73
        }
    }
    v74 := []int{v2, v2, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v59
    _ = v74
}
