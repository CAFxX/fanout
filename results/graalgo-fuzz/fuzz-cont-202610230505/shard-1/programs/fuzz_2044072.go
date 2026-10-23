package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int32; F1 string; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 int8; F1 int32; F2 int32 }
type GenS13 struct { F0 int16 }
func genFn0() int {
    if (!true) {
        v4 := GenS1{F0: -1, F1: 976, F2: -71}
        if true {
            var v5 uint8 = 8
            v6 := []int{32767, -814}
            _ = v5
            _ = v6
        }
        var v7 uint32 = 2
        _ = v4
        _ = v7
    } else {
        var v8 uint64 = 18446744073709551615
        v9 := GenS2{F0: v8, F1: 152}
        var v10 uint64 = 877
        _ = v10
        _ = v8
        _ = v9
    }
    switch "" {
    case "foo bar":
        var v11 uint16 = 0
        _ = v11
    default:
        if false {
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 string = "foo bar"
                _ = fi12
                _ = v13
            }
            if (!true) {
                if ("a" != "hello") {
                    v14 := []int{}
                    var v15 string = ("" + "")
                    _ = v14
                    _ = v15
                }
            } else {
                var v16 int32 = -2147483648
                _ = v16
            }
        }
    }
    v17 := []int{128, 940}
    for fi18 := 0; fi18 < 6; fi18++ {
        if false {
            var v19 int64 = -36
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int16 = 234
                fmt.Println(fi20, fi18)
                _ = fi20
                _ = v21
            }
            _ = v19
        } else {
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 bool = (-391 <= 2147483647)
                v17 = append(v17, 2147483647)
                v24 := GenS1{F0: 100, F1: 516, F2: -6}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        _ = fi18
    }
    _ = v17
    return 42
}
func genFn1() int32 {
    switch "\n\t" {
    case "hello":
        v25 := []int{9223372036854775807, 9223372036854775807, genFn0()}
        _ = v25
    default:
        v26 := GenS3{F0: 1}
        _ = v26
    }
    var v27 float32 = 1e10
    _ = v27
    return -56
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    v3 := []string{"foo bar"}
    for fi28 := 0; fi28 < 20; fi28++ {
        v2 = ((!(((v2 && false) && (-1 >= 276)) || (32 == 2))) || (0 <= 0))
        _ = fi28
    }
    v29 := GenS4{F0: "", F1: 65535}
    v2 = true
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS5{F0: 150, F1: 744}
        _ = fi30
        _ = v31
    }
    var v32 uint16 = 503
    v33 := GenS3{F0: -124}
    if len(v3) > 0 {
        v3[0] = ("a" + ("" + "a"))
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        v32 += (v32 * (v32 >> ((714) & 63)))
        switch "hello" {
        case "foo bar":
            if ((1 > -9223372036854775808) || false) {
                v2 = v2
                if (128 == genFn0()) {
                    var v35 float64 = 3.14159
                    delete(v1, "hello")
                    _ = v35
                } else {
                    var v36 uint16 = v32
                    _ = v36
                }
            }
        default:
            v37 := GenS6{F0: 0, F1: ((("\n\t" + "") + (("" + ("foo bar" + "foo bar")) + "foo bar")) + "a"), F2: 1}
            _ = v37
        }
        v2 = (v2 && false)
        _ = fi34
    }
    v38 := GenS4{F0: (("foo bar" + "\n\t") + ("foo bar" + "a")), F1: 4294967295}
    for fi39 := 0; fi39 < 14; fi39++ {
        if v2 {
            switch 4294967295 {
            case 9223372036854775807:
                if (32 > 42) {
                    for fi40 := 0; fi40 < 16; fi40++ {
                        switch "\n\t" {
                        case "a":
                            delete(v1, "")
                        case "":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        default:
                            v32 += (v32 & v32)
                        }
                        if ((((1e10 + 0.0) + (1e10 + -2.25)) + ((0.0 - (1e10 / 1.0)) * -2.25)) <= ((0.1 - 3.14159) - 1.5)) {
                            v2 = false
                        } else {
                            fmt.Println(fi40)
                        }
                        v32 = 689
                        _ = fi40
                    }
                } else {
                    for fi41 := 0; fi41 < 4; fi41++ {
                        if (-1 >= 9223372036854775807) {
                            if v2 {
                                var v42 int8 = 127
                                _ = v42
                            }
                            v3 = append(v3, "hello")
                        }
                        _ = fi41
                    }
                }
            default:
                for fi43 := 0; fi43 < 11; fi43++ {
                    v2 = (-128 > 127)
                    _ = fi43
                }
            }
        } else {
            v2 = (v32 < (v32 % ((355) | 1)))
            for fi44 := 0; fi44 < 3; fi44++ {
                v32 = (v32 % ((v32) | 1))
                switch "a" {
                case "":
                    v45 := GenS7{F0: v32}
                    _ = v45
                case "a":
                    v32 += 1
                    var v46 uint32 = 0
                    _ = v46
                default:
                    var v47 float32 = 1.5
                    _ = v47
                }
                _ = fi44
            }
        }
        v32 = v32
        _ = fi39
    }
    if (-1 <= 100) {
        var v48 uint16 = v32
        if v2 {
            if ((-9223372036854775808 >= -944) || v2) {
                if (v2 || ((0.0 - (0.1 + (3.14159 / ((1.5 + 0.0) - (1.5 - 0.1))))) != 1e10)) {
                    delete(v1, "hello")
                }
                var v49 int8 = -60
                _ = v49
            }
            var v50 uint = 42
            _ = v50
        } else {
            fmt.Println(v48, v2)
            for fi51 := 0; fi51 < 14; fi51++ {
                fmt.Println(fi51, v48, v48)
                for fi52 := 0; fi52 < 1; fi52++ {
                    v2 = true
                    v53 := GenS7{F0: v32}
                    _ = fi52
                    _ = v53
                }
                var v54 float32 = 0.1
                _ = fi51
                _ = v54
            }
        }
        _ = v48
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        if (188 == 1) {
            fmt.Println(fi55)
            for fi56 := 0; fi56 < 16; fi56++ {
                if true {
                    if (1 == 1) {
                        if false {
                            var v57 int = 372
                            var v58 uint32 = 2147483648
                            v59 := GenS8{F0: 127, F1: -1}
                            if v2 {
                                if v2 {
                                    v57 = (fi56 * (fi55 ^ v57))
                                    v58 = v58
                                } else {
                                    var v60 int8 = 1
                                    _ = v60
                                }
                                v61 := GenS9{F0: int(v32), F1: 42}
                                fmt.Println(v57)
                                _ = v61
                            } else {
                                v1[("\n\t" + ("" + ((("hello" + "") + ("a" + "a")) + "\n\t")))] = 128
                            }
                            v62 := GenS9{F0: fi56, F1: 0}
                            if (18446744073709551615 < 42) {
                                v63 := GenS9{F0: v57, F1: 65535}
                                _ = v63
                            } else {
                                v57 = fi55
                            }
                            var v64 int8 = -49
                            _ = v57
                            _ = v58
                            _ = v59
                            _ = v62
                            _ = v64
                        }
                        if v2 {
                            if v2 {
                                v65 := GenS5{F0: (v32 % ((v32) | 1)), F1: 279}
                                _ = len(v1)
                                var v66 int = fi56
                                _ = v65
                                _ = v66
                            }
                            var v67 bool = v2
                            if (genFn0() >= (fi55 << ((1) & 63))) {
                                v32 -= 0
                                if v2 {
                                    v68 := GenS9{F0: (fi55 + (-(fi55))), F1: 255}
                                    var v69 uint32 = 360
                                    if (2 != uint16(v69)) {
                                        v69 -= (v69 >> ((2) & 63))
                                        if (629 < 603) {
                                            fmt.Println(fi55, v69)
                                            v32 += (^(v32))
                                            _ = len(v3)
                                        }
                                    }
                                    _ = v68
                                    _ = v69
                                } else {
                                    switch ("a" + "foo bar") {
                                    case "\n\t":
                                        var v70 int = 739
                                        _ = v70
                                    case "hello":
                                        v1["a"] = (fi56 / ((fi56) | 1))
                                    }
                                    var v71 uint = 465
                                    _ = v71
                                }
                                v67 = (164 <= 132)
                                v72 := GenS1{F0: -82, F1: (v32 & (v32 & 587)), F2: 0}
                                _ = v72
                            }
                            _ = v67
                        }
                    }
                    if false {
                        v73 := GenS4{F0: "a", F1: 42}
                        switch "" {
                        case "foo bar":
                            var v74 uint8 = 94
                            _ = v74
                        case "hello":
                            v75 := GenS10{F0: 42, F1: 834}
                            _ = v75
                        case "\n\t":
                            v32 += 2
                        }
                        _ = v73
                    } else {
                        switch (^(fi55)) {
                        case -295:
                            var v76 int32 = genFn1()
                            _ = v76
                        case -2147483648:
                            v32 += (v32 - 255)
                        case 39:
                            fmt.Println(fi56, fi55)
                        default:
                            if (!true) {
                                v32 = v32
                            } else {
                                v77 := v1[(((("a" + "foo bar") + "\n\t") + "foo bar") + (("a" + "") + "\n\t"))]
                                var v78 uint = 42
                                _ = v77
                                _ = v78
                            }
                        }
                        switch fi55 {
                        case -1:
                            v79 := GenS9{F0: genFn0(), F1: 75}
                            _ = v79
                        case 32768:
                            v2 = (96 >= 255)
                        case 429:
                            var v80 int8 = 100
                            _ = v80
                        default:
                            v1["hello"] = (-(fi55))
                        }
                        fmt.Println(v32, fi56)
                        var v81 int = 4294967295
                        var v82 int32 = genFn1()
                        var v83 bool = (!false)
                        if v83 {
                            v84 := v1[""]
                            v85 := GenS11{F0: "foo bar", F1: 127}
                            v84 = fi55
                            v82 = -1
                            var v86 int8 = 100
                            v3 = append(v3, "hello")
                            v87 := GenS12{F0: v86, F1: genFn1(), F2: v82}
                            _ = v84
                            _ = v85
                            _ = v86
                            _ = v87
                        } else {
                            if len(v3) > 0 {
                                v3[0] = (("hello" + (("" + ("\n\t" + "")) + "a")) + "a")
                            }
                            v2 = v2
                            var v88 uint64 = 796
                            _ = v88
                        }
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                    var v89 float64 = 0.1
                    _ = v89
                }
                var v90 uint32 = 1
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = fi56
                _ = v90
            }
        } else {
            for fi91 := 0; fi91 < 16; fi91++ {
                v92 := GenS13{F0: 42}
                _ = fi91
                _ = v92
            }
        }
        _ = fi55
    }
    v32 = v32
    for fi93 := 0; fi93 < 9; fi93++ {
        var v94 int = (fi93 | genFn0())
        _ = fi93
        _ = v94
    }
    v32 = (^(v32))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v38
}
