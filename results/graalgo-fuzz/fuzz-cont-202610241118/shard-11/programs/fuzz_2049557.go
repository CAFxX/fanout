package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int32; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 bool; F2 string }
type GenS8 struct { F0 uint32; F1 int32; F2 int8 }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 int32; F1 int64; F2 int64 }
type GenS11 struct { F0 int; F1 uint }
func genFn0(p0 uint64, p1 int) int64 {
    var v7 uint8 = 127
    p0 = (p0 & p0)
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{(^(fi8))}
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(fi10, p0, p0)
            if false {
                if (true && ((0 > 65535) || true)) {
                    if false {
                        var v11 float32 = -2.25
                        _ = v11
                    }
                }
                var v12 float64 = (1.5 / 1.0)
                v13 := GenS1{F0: 2147483647, F1: 1000}
                _ = v12
                _ = v13
            }
            var v14 uint = 255
            _ = fi10
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    if (true && (2 <= 4)) {
        p1 -= (^(p1))
    }
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint) string {
    v15 := []int{-2147483648, 128, -1}
    if true {
        if (-128 == 0) {
            fmt.Println(p0)
            var v16 int8 = 0
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint64 = 2
                for fi19 := 0; fi19 < 4; fi19++ {
                    v20 := GenS2{F0: fi17, F1: 2147483647, F2: 1}
                    v21 := GenS3{F0: 0, F1: 1, F2: 0}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        p0 -= p0
    }
    var v22 string = ""
    v23 := GenS4{F0: 1000}
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 uint16 = 863
        if (!true) {
            if len(v15) > 0 {
                v15[0] = (-(fi24))
            }
            var v26 int = (fi24 / ((fi24) | 1))
            _ = v26
        } else {
            if ((42 == -1) || ((v25 > 1) && (42 >= -473))) {
                if true {
                    var v27 uint32 = 65535
                    var v28 string = (v22 + (("foo bar" + (v22 + v22)) + (((v22 + "") + v22) + v22)))
                    _ = v27
                    _ = v28
                } else {
                    if (!false) {
                        for fi29 := 0; fi29 < 8; fi29++ {
                            v22 = v22
                            if true {
                                v25 = v25
                            }
                            v30 := GenS2{F0: -505, F1: 65535, F2: 1}
                            _ = fi29
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 20; fi31++ {
                            if false {
                                var v32 string = ""
                                v33 := GenS3{F0: 0, F1: int16(p0), F2: 1}
                                p0 = 255
                                _ = v32
                                _ = v33
                            }
                            v22 = ""
                            var v34 float64 = (-2.25 - 0.1)
                            _ = fi31
                            _ = v34
                        }
                        var v35 float64 = 3.14159
                        _ = v35
                    }
                    switch (fi24 | fi24) {
                    case 2147483647:
                        for fi36 := 0; fi36 < 8; fi36++ {
                            fmt.Println(v22, fi24, v25)
                            v22 = v22
                            _ = fi36
                        }
                    default:
                        var v37 int32 = -1
                        _ = v37
                    }
                    v38 := GenS3{F0: 255, F1: 877, F2: 42}
                    _ = v38
                }
                v39 := GenS4{F0: 32767}
                var v40 int = (fi24 + -558)
                _ = v39
                _ = v40
            }
            if len(v15) > 0 {
                v15[0] = (fi24 % ((-9223372036854775808) | 1))
            }
            var v41 uint16 = (^(v25))
            _ = v41
        }
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = v15
    _ = v22
    _ = v23
    return v22
}

func main() {
    v1 := []int{}
    var v2 string = ("a" + "")
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    v2 = ("a" + (v2 + (v2 + (v2 + "\n\t"))))
    v2 = ("a" + v2)
    if false {
        v2 = (v2 + (("\n\t" + "hello") + v2))
        v2 = "hello"
        v2 = "\n\t"
    } else {
        if false {
            v42 := GenS5{F0: 701, F1: 42, F2: 42}
            _ = v42
        }
    }
    v43 := GenS5{F0: 9, F1: 418, F2: 1}
    fmt.Println(v2)
    for fi44 := 0; fi44 < 2; fi44++ {
        for fi45 := 0; fi45 < 3; fi45++ {
            v2 = "a"
            var v46 int = fi44
            if true {
                v46 = 32768
                v46 += -2147483648
                v2 = (("" + v2) + "")
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v47 uint32 = 279
                switch 32768 {
                case 4294967295:
                    if false {
                        if true {
                            v46 -= -2147483648
                            var v48 uint64 = 1
                            fmt.Println(v48, v47)
                            _ = v48
                        }
                        if true {
                            fmt.Println(v2, fi45, fi45)
                            if (((42 == 64) && (!(!(0 == 2)))) && false) {
                                v46 += fi45
                                var v49 uint16 = 368
                                var v50 uint16 = 927
                                if false {
                                    var v51 string = v2
                                    var v52 int8 = 12
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 int64 = int64(v46)
                                    if ((3.14159 - float32(1e10)) >= 1e10) {
                                        v49 = v49
                                        v54 := GenS4{F0: -230}
                                        _ = v54
                                    }
                                    var v55 uint16 = 859
                                    var v56 float64 = (0.0 / ((1e10 - 1.5) / -2.25))
                                    _ = v53
                                    _ = v55
                                    _ = v56
                                }
                                _ = v49
                                _ = v50
                            }
                            v57 := GenS2{F0: (^(fi45)), F1: (v47 << ((255) & 63)), F2: -32768}
                            _ = v57
                        }
                        var v58 uint64 = 154
                        v46 = (v46 % ((2147483647) | 1))
                        _ = v58
                    }
                case 42:
                    fmt.Println(fi45, fi45)
                default:
                    v47 = (v47 ^ v47)
                    var v59 uint8 = 148
                    var v60 int8 = 100
                    var v61 float64 = 0.1
                    _ = v59
                    _ = v60
                    _ = v61
                }
                v2 = "hello"
                v47 = (v47 >> ((42) & 63))
                v3[(v47 ^ (v47 * (v47 + 419)))] = false
                _ = v47
            }
            v2 = v2
            _ = fi45
            _ = v46
        }
        for fi62 := 0; fi62 < 12; fi62++ {
            if true {
                fmt.Println(fi44, fi62, fi44)
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi62
        }
        if false {
            for fi63 := 0; fi63 < 19; fi63++ {
                var v64 bool = false
                var v65 uint32 = 18
                _ = fi63
                _ = v64
                _ = v65
            }
        } else {
            v66 := GenS6{F0: 0}
            for fi67 := 0; fi67 < 6; fi67++ {
                v68 := GenS1{F0: 0, F1: 32767}
                _ = fi67
                _ = v68
            }
            if true {
                v2 = "hello"
                var v69 int64 = -9223372036854775808
                var v70 float32 = 1e10
                _ = v69
                _ = v70
            } else {
                v3[2147483648] = false
                switch ("a" + v2) {
                case "":
                    if (true && true) {
                        if true {
                            v71 := GenS7{F0: "\n\t", F1: (932 <= 255), F2: ((v2 + (v2 + (v2 + v2))) + "")}
                            v1 = append(v1, fi44)
                            v72 := GenS8{F0: 2147483647, F1: -652, F2: 56}
                            _ = v71
                            _ = v72
                        } else {
                            if (0 < 190) {
                                var v73 int = fi44
                                _ = v73
                            }
                            switch 42 {
                            case 1:
                                v1 = append(v1, -2147483648)
                            case 315:
                                switch (^(fi44)) {
                                case -626:
                                    if len(v1) > 0 {
                                        v1[0] = -1
                                    }
                                case 42:
                                    for fi74 := 0; fi74 < 14; fi74++ {
                                        fmt.Println(v2)
                                        _ = fi74
                                    }
                                case -956:
                                    v75 := GenS9{F0: 0, F1: -27}
                                    _ = v75
                                }
                            case 9223372036854775807:
                                _ = len(v3)
                            default:
                                var v76 uint64 = 42
                                _ = v76
                            }
                        }
                        v2 = (v2 + v2)
                        v77 := GenS10{F0: 42, F1: 1, F2: 0}
                        _ = v77
                    } else {
                        _ = len(v1)
                        for fi78 := 0; fi78 < 10; fi78++ {
                            if (!true) {
                                v79 := GenS10{F0: 42, F1: 9223372036854775807, F2: -9223372036854775808}
                                if ((754 < -1) && (!true)) {
                                    var v80 uint32 = 2147483647
                                    _ = v80
                                } else {
                                    var v81 int = -9223372036854775808
                                    if len(v1) > 0 {
                                        v1[0] = (v81 + 127)
                                    }
                                    v82 := GenS8{F0: 2, F1: 2147483647, F2: 74}
                                    if (127 <= -120) {
                                        v2 = "\n\t"
                                        var v83 uint = 21
                                        v84 := GenS9{F0: 1, F1: 42}
                                        _ = v83
                                        _ = v84
                                    }
                                    v2 = "\n\t"
                                    if false {
                                        v85 := GenS11{F0: v81, F1: 157}
                                        var v86 int8 = 100
                                        _ = len(v6)
                                        var v87 bool = ((1e10 + -2.25) <= 0.0)
                                        var v88 float64 = 3.14159
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                    var v89 int16 = -32768
                                    switch (fi78 ^ (fi44 * (fi78 % ((0) | 1)))) {
                                    case 32768:
                                        v81 = v81
                                    case -9223372036854775808:
                                        if true {
                                            v90 := GenS4{F0: (v89 | (v89 & (v89 % ((1000) | 1))))}
                                            var v91 string = v2
                                            _ = v90
                                            _ = v91
                                        } else {
                                            _ = len(v1)
                                            v81 = (-(v81))
                                        }
                                    case 87:
                                        var v92 uint8 = 74
                                        _ = v92
                                    default:
                                        v89 += -103
                                    }
                                    _ = v81
                                    _ = v82
                                    _ = v89
                                }
                                _ = v79
                            } else {
                                v93 := GenS7{F0: "", F1: (!(!true)), F2: "\n\t"}
                                v1 = append(v1, (-(fi78)))
                                _ = len(v1)
                                _ = v93
                            }
                            _ = fi78
                        }
                    }
                default:
                    switch (fi44 | fi44) {
                    case 4294967295:
                        v2 = v2
                    case -186:
                        switch fi44 {
                        case 32767:
                            for fi94 := 0; fi94 < 13; fi94++ {
                                var v95 uint = 834
                                v2 = (v2 + v2)
                                switch "" {
                                case "":
                                    var v96 uint64 = 0
                                    _ = v96
                                }
                                _ = fi94
                                _ = v95
                            }
                        default:
                            var v97 string = ""
                            _ = v97
                        }
                    default:
                        var v98 uint = 18446744073709551615
                        _ = v98
                    }
                }
            }
            _ = v66
        }
        _ = fi44
    }
    var v99 int8 = 35
    var v100 float64 = ((3.14159 * 3.14159) + ((1.5 * (-2.25 * (0.1 * 1.5))) - (((0.0 + 0.1) + (1e10 - -2.25)) - 1.5)))
    var v101 uint32 = 2147483648
    v2 = ((("a" + ((v2 + ((((v2 + "foo bar") + v2) + "") + ("foo bar" + ((v2 + ((v2 + "foo bar") + v2)) + v2)))) + "")) + ((v2 + v2) + v2)) + v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v99
}
