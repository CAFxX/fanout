package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 string; F2 int64 }
type GenS5 struct { F0 uint64; F1 int8; F2 uint }
type GenS6 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint32; F1 int64; F2 uint }
type GenS9 struct { F0 string; F1 uint8 }
func genFn0(p0 int8, p1 int, p2 uint64) uint {
    for fi5 := 0; fi5 < 17; fi5++ {
        p2 = (p2 ^ (p2 % (((^(p2))) | 1)))
        _ = fi5
    }
    var v6 int = (p1 - (p1 % ((p1) | 1)))
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            if true {
                v9 := []int{fi8, v6}
                _ = v9
            } else {
                if true {
                    v10 := GenS1{F0: p2, F1: p2, F2: 2}
                    v6 = -629
                    p1 = 2147483647
                    _ = v10
                }
                fmt.Println(v6, v6, fi7)
                v11 := GenS1{F0: 749, F1: 18446744073709551615, F2: 2147483647}
                var v12 uint64 = (p2 - 869)
                v13 := []int{(^(fi8)), -2147483648, (fi8 & 2147483647), (^(fi7))}
                _ = v11
                _ = v12
                _ = v13
            }
            fmt.Println(fi7)
            v6 = -928
            v14 := GenS2{F0: 65535}
            _ = fi8
            _ = v14
        }
        var v15 string = (((((("hello" + "a") + "\n\t") + "\n\t") + ("foo bar" + "hello")) + "hello") + "foo bar")
        var v16 uint16 = 65535
        _ = fi7
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 float64 = (0.1 * (3.14159 + 1e10))
    var v2 uint64 = 2
    v3 := []string{(("\n\t" + "foo bar") + "")}
    v4 := []int{9223372036854775807, 2147483647, 9223372036854775807, -109}
    var v17 uint = 65535
    v17 = (v17 | (v17 - v17))
    switch ((("a" + "a") + "") + "hello") {
    case "foo bar":
        v1 -= float64(0.1)
    case "hello":
        v17 += 18446744073709551615
    case "\n\t":
        for fi18 := 0; fi18 < 13; fi18++ {
            v3 = append(v3, "hello")
            if (!true) {
                v19 := GenS3{F0: -372}
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 bool = false
                    var v22 uint8 = 0
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                switch "hello" {
                case "foo bar":
                    if (1 < uint32(v17)) {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            v24 := GenS1{F0: uint64(v17), F1: v2, F2: 1}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi23
                            _ = v24
                        }
                        v1 = v1
                        switch 1 {
                        case 71:
                            if (((("" + ("\n\t" + "a")) + ("\n\t" + "")) <= ("" + ("\n\t" + "foo bar"))) && false) {
                                var v25 int16 = -32768
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 1; fi26++ {
                                    v27 := GenS4{F0: 382, F1: "hello", F2: 962}
                                    _ = fi26
                                    _ = v27
                                }
                            }
                        }
                    } else {
                        if (256 <= 1) {
                            for fi28 := 0; fi28 < 6; fi28++ {
                                var v29 uint8 = 128
                                _ = fi28
                                _ = v29
                            }
                            if (42 != (^(v2))) {
                                if false {
                                    v30 := GenS1{F0: v2, F1: 430, F2: 2147483648}
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    _ = v30
                                } else {
                                    switch fi18 {
                                    case 373:
                                        var v31 float32 = 0.0
                                        _ = v31
                                    case 2147483647:
                                        v2 = v2
                                    case -9223372036854775808:
                                        for fi32 := 0; fi32 < 14; fi32++ {
                                            v1 -= v1
                                            var v33 int16 = 32767
                                            v17 += (v17 % (((^(v17))) | 1))
                                            _ = fi32
                                            _ = v33
                                        }
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                    }
                                }
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "":
                                    var v34 int16 = -32768
                                    _ = v34
                                default:
                                    v1 = v1
                                }
                            }
                            if (!(!true)) {
                                var v35 uint64 = v2
                                _ = v35
                            }
                        }
                    }
                case "\n\t":
                    if false {
                        fmt.Println(v17, v1)
                    }
                case "hello":
                    var v36 float32 = 1e10
                    _ = v36
                default:
                    v37 := GenS3{F0: 0}
                    _ = v37
                }
                _ = v19
            }
            _ = fi18
        }
    }
    fmt.Println(v2, v2)
    switch ("" + "\n\t") {
    case "\n\t":
        var v38 int16 = -1
        _ = v38
    case "":
        var v39 int8 = 121
        _ = v39
    case "foo bar":
        v40 := GenS5{F0: (^(v2)), F1: 42, F2: v17}
        _ = v40
    default:
        var v41 int = -1
        _ = v41
    }
    if (219 != 127) {
        v17 = (v17 / (((v17 * (v17 - v17))) | 1))
        v42 := GenS6{F0: uint8(v2), F1: 32767, F2: 42}
        _ = v42
    } else {
        fmt.Println(v17)
        if false {
            for fi43 := 0; fi43 < 18; fi43++ {
                v17 -= v17
                _ = fi43
            }
        }
        if true {
            var v44 int32 = 187
            var v45 uint16 = 851
            _ = v44
            _ = v45
        }
    }
    if true {
        var v46 uint16 = 32767
        switch 4294967295 {
        case 0:
            for fi47 := 0; fi47 < 10; fi47++ {
                fmt.Println(v17, fi47)
                _ = fi47
            }
        case 9223372036854775807:
            for fi48 := 0; fi48 < 6; fi48++ {
                var v49 int32 = int32(fi48)
                _ = fi48
                _ = v49
            }
        case -2147483648:
            if len(v3) > 0 {
                v3[0] = (("\n\t" + (("foo bar" + "") + "a")) + ("foo bar" + ((("foo bar" + (("\n\t" + "a") + (("hello" + (("foo bar" + "") + "")) + ((("hello" + ((("\n\t" + "") + ((("\n\t" + "foo bar") + (("foo bar" + "foo bar") + ("hello" + ""))) + ((("a" + "a") + ((("hello" + "hello") + "") + "\n\t")) + "a"))) + "\n\t")) + "foo bar") + "\n\t")))) + "foo bar") + "a")))
            }
        default:
            switch 128 {
            case 463:
                v50 := GenS6{F0: 207, F1: 42, F2: (^(v46))}
                _ = v50
            case 32768:
                var v51 uint8 = 0
                _ = v51
            case 128:
                for fi52 := 0; fi52 < 7; fi52++ {
                    fmt.Println(v17)
                    v46 += v46
                    _ = fi52
                }
            }
        }
        v46 -= v46
        _ = v46
    } else {
        for fi53 := 0; fi53 < 4; fi53++ {
            if len(v3) > 0 {
                v3[0] = ""
            }
            switch 32768 {
            case 4294967295:
                for fi54 := 0; fi54 < 2; fi54++ {
                    if true {
                        var v55 int8 = -126
                        var v56 int32 = 1
                        if false {
                            var v57 uint = genFn0((v55 * (v55 ^ (v55 - v55))), -945, v2)
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = 928
                                }
                            } else {
                                var v58 uint16 = 0
                                if ("foo bar" != "a") {
                                    switch "a" {
                                    case "foo bar":
                                        v59 := GenS7{F0: 1, F1: ""}
                                        _ = v59
                                    case "hello":
                                        v60 := GenS8{F0: 1, F1: 9223372036854775807, F2: v57}
                                        _ = v60
                                    case "a":
                                        if (!false) {
                                            switch ("a" + (("\n\t" + ("a" + "foo bar")) + "\n\t")) {
                                            case "\n\t":
                                                if false {
                                                    var v61 float64 = ((3.14159 - (v1 - 3.14159)) + (v1 / v1))
                                                    _ = v61
                                                }
                                            case "":
                                                var v62 uint = genFn0((-(v55)), int(v55), v2)
                                                _ = v62
                                            case "a":
                                                v2 = (v2 >> ((genFn0(v55, (fi53 << ((v17) & 63)), v2)) & 63))
                                            }
                                            _ = len(v4)
                                        } else {
                                            fmt.Println(v1, fi53, v17)
                                            v63 := GenS3{F0: 0}
                                            _ = v63
                                        }
                                    }
                                }
                                if true {
                                    var v64 uint8 = 127
                                    _ = v64
                                } else {
                                    if false {
                                        var v65 uint32 = 4294967295
                                        var v66 bool = true
                                        var v67 uint16 = v58
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    switch -848 {
                                    case 32767:
                                        v55 += (v55 / (((v55 - v55)) | 1))
                                    case 752:
                                        v58 += (v58 ^ (v58 | (v58 ^ 255)))
                                    case 0:
                                        var v68 string = "a"
                                        _ = v68
                                    default:
                                        fmt.Println(v2)
                                    }
                                    v1 = 3.14159
                                }
                                _ = v58
                            }
                            v69 := GenS6{F0: 42, F1: 1000, F2: 255}
                            _ = v57
                            _ = v69
                        } else {
                            if true {
                                var v70 uint16 = uint16(v2)
                                var v71 uint = genFn0(v55, fi54, 42)
                                var v72 uint16 = (v70 % ((v70) | 1))
                                _ = v70
                                _ = v71
                                _ = v72
                            }
                            if ((!(208 == 208)) && (-862 != 172)) {
                                var v73 uint64 = v2
                                v74 := GenS9{F0: "", F1: 2}
                                if (float32((v1 - (v1 - (0.1 - float64(float32(0.1)))))) >= float32((1e10 + v1))) {
                                    v56 = (^(v56))
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        v75 := GenS4{F0: 32767, F1: ((("a" + "\n\t") + (("hello" + ("a" + ("\n\t" + ""))) + "hello")) + (("a" + "foo bar") + (("a" + "") + ("a" + "a")))), F2: -917}
                                        _ = v75
                                    default:
                                        if false {
                                            var v76 uint32 = 93
                                            _ = v76
                                        } else {
                                            v1 = v1
                                            switch 4294967295 {
                                            case -1:
                                                if len(v3) > 0 {
                                                    v3[0] = "a"
                                                }
                                            }
                                        }
                                    }
                                    v55 = (v55 + v55)
                                }
                                _ = v73
                                _ = v74
                            } else {
                                fmt.Println(v2, v2)
                                if true {
                                    v56 += int32(v17)
                                    fmt.Println(v1)
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = fi54
                                    }
                                }
                            }
                        }
                        _ = v55
                        _ = v56
                    }
                    if true {
                        v17 -= 1
                    }
                    if true {
                        switch 2147483647 {
                        case 0:
                            _ = len(v3)
                        case -227:
                            v1 -= v1
                        case 32768:
                            var v77 float32 = 3.14159
                            _ = v77
                        }
                        if (756 >= 446) {
                            var v78 int = (fi53 - fi54)
                            v1 -= v1
                            fmt.Println(fi53, v17)
                            _ = v78
                        } else {
                            var v79 float64 = 1.5
                            _ = v79
                        }
                        v1 -= (v1 / v1)
                        if (!((fi54 ^ (fi54 & fi54)) <= 2147483647)) {
                            v3 = append(v3, ("\n\t" + "hello"))
                            fmt.Println(v17, v17)
                        }
                        switch 32768 {
                        case -412:
                            var v80 int64 = -934
                            _ = v80
                        }
                        v17 = (v17 & v17)
                        var v81 int8 = 100
                        _ = v81
                    } else {
                        var v82 int16 = 42
                        v2 -= 1
                        _ = v82
                    }
                    _ = fi54
                }
            case 42:
                v2 -= 676
            case -534:
                var v83 uint8 = 1
                _ = v83
            }
            _ = fi53
        }
    }
    v17 = (v17 / ((genFn0(100, -376, 4294967295)) | 1))
    v1 = ((1e10 / (-2.25 * float64(((1e10 / (0.1 + 1.5)) + (1e10 * 0.0))))) - (-2.25 / ((v1 / 3.14159) - ((0.0 * v1) * 0.0))))
    var v84 uint32 = 1
    if (!true) {
        var v85 uint64 = (v2 * v2)
        _ = v85
    } else {
        var v86 uint = v17
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v84
}
