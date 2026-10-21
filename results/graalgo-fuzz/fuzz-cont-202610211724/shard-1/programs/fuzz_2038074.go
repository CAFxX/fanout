package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int8 }
type GenS5 struct { F0 string; F1 uint64; F2 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
func genFn0(p0 uint8) int32 {
    p0 = 0
    var v7 uint32 = 1
    switch "" {
    case "hello":
        v7 -= (v7 >> ((65535) & 63))
    default:
        if (!false) {
            for fi8 := 0; fi8 < 2; fi8++ {
                if (false && true) {
                    v9 := GenS1{F0: 0, F1: (fi8 >> ((0) & 63)), F2: 4294967295}
                    _ = v9
                } else {
                    var v10 int32 = 0
                    fmt.Println(p0, v10)
                    for fi11 := 0; fi11 < 11; fi11++ {
                        if (uint16(v7) >= 1) {
                            v12 := []int{}
                            v10 -= v10
                            fmt.Println(v7)
                            _ = v12
                        } else {
                            switch "a" {
                            case "":
                                var v13 int16 = -1
                                _ = v13
                            }
                            v14 := GenS1{F0: 238, F1: fi8, F2: 2}
                            v7 = v7
                            _ = v14
                        }
                        _ = fi11
                    }
                    _ = v10
                }
                fmt.Println(v7)
                _ = fi8
            }
            var v15 int64 = 0
            _ = v15
        } else {
            var v16 int16 = 1000
            v17 := []int{32768, 32768, 0, 32768}
            v16 = 780
            _ = v16
            _ = v17
        }
    }
    v7 += 1
    p0 -= (p0 << ((898) & 63))
    _ = p0
    _ = v7
    return -2147483648
}
func genFn1() int32 {
    var v18 uint8 = 127
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v18, v18)
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := []int{(fi19 + fi20), (fi20 / ((fi20) | 1)), (fi20 + fi19)}
            v22 := GenS1{F0: 2, F1: (fi19 + (fi20 / ((fi19) | 1))), F2: uint64(v18)}
            _ = fi20
            _ = v21
            _ = v22
        }
        v23 := GenS2{F0: 472, F1: 0}
        _ = fi19
        _ = v23
    }
    var v24 int = 32767
    v18 += v18
    _ = v18
    _ = v24
    return 2147483647
}

func main() {
    var v1 uint8 = 217
    v2 := make(map[uint32]int)
    var v3 uint32 = 966
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    v6 := []int{-317, -446}
    if true {
        fmt.Println(v1)
        var v25 uint = 811
        _ = v25
    } else {
        if (-104 >= -106) {
            var v26 uint8 = v1
            v27 := GenS1{F0: 470, F1: 32768, F2: 1}
            _ = v26
            _ = v27
        } else {
            switch 445 {
            case 573:
                v28 := GenS3{F0: "a", F1: genFn1()}
                _ = v28
            case 2147483647:
                var v29 int32 = 42
                _ = v29
            case 1:
                if true {
                    v1 -= 87
                    var v30 string = ("" + "\n\t")
                    _ = v30
                }
            default:
                v1 = v1
            }
        }
        v1 = v1
        for fi31 := 0; fi31 < 6; fi31++ {
            v6 = append(v6, fi31)
            _ = fi31
        }
    }
    v1 -= 42
    var v32 uint8 = v1
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        if (("\n\t" == "\n\t") && false) {
            var v34 string = "hello"
            _ = v34
        }
        delete(v2, (v3 % (((v3 / ((65535) | 1))) | 1)))
        _ = fi33
    }
    var v35 float64 = 1.5
    v6 = append(v6, -672)
    if false {
        v36 := v5[v3]
        v37 := GenS1{F0: uint(v3), F1: 951, F2: 4294967295}
        _ = v36
        _ = v37
    } else {
        v6 = append(v6, -9223372036854775808)
        for fi38 := 0; fi38 < 4; fi38++ {
            var v39 float64 = float64(0.1)
            _ = fi38
            _ = v39
        }
    }
    if ((v32 ^ v1) >= v1) {
        var v40 uint16 = 0
        var v41 uint32 = 198
        if (128 == 2) {
            fmt.Println(v35)
            var v42 uint16 = (^(v40))
            _ = v42
        } else {
            if ("a" == ("a" + "hello")) {
                var v43 string = "foo bar"
                fmt.Println(v35, v41)
                _ = v43
            }
            var v44 uint16 = (v40 & (^(v40)))
            _ = v44
        }
        _ = v40
        _ = v41
    }
    if true {
        var v45 float32 = 0.0
        _ = v45
    } else {
        var v46 int32 = 571
        v47 := GenS1{F0: 4294967295, F1: 0, F2: uint64(v3)}
        _ = len(v2)
        _ = v46
        _ = v47
    }
    var v48 int = 9223372036854775807
    if false {
        if (!false) {
            var v49 bool = (false || ((v48 >> ((4294967295) & 63)) >= (-(v48))))
            v50 := GenS1{F0: 2, F1: v48, F2: 42}
            _ = v49
            _ = v50
        } else {
            fmt.Println(v35)
            for fi51 := 0; fi51 < 16; fi51++ {
                v52 := GenS2{F0: v3, F1: -2}
                v2[v3] = 42
                for fi53 := 0; fi53 < 18; fi53++ {
                    fmt.Println(v1, v32, v1)
                    v54 := GenS2{F0: 0, F1: 42}
                    if true {
                        fmt.Println(v35, v32, fi53)
                        var v55 uint8 = v32
                        _ = v55
                    }
                    _ = fi53
                    _ = v54
                }
                _ = fi51
                _ = v52
            }
            v56 := GenS4{F0: v3, F1: 1, F2: 1}
            _ = v56
        }
        for fi57 := 0; fi57 < 1; fi57++ {
            for fi58 := 0; fi58 < 1; fi58++ {
                var v59 int8 = -65
                _ = fi58
                _ = v59
            }
            _ = len(v6)
            for fi60 := 0; fi60 < 2; fi60++ {
                v61 := v4["hello"]
                _ = fi60
                _ = v61
            }
            _ = fi57
        }
        for fi62 := 0; fi62 < 18; fi62++ {
            v63 := GenS5{F0: ((("\n\t" + "foo bar") + "") + ("foo bar" + "hello")), F1: 42, F2: ((1.5 - -2.25) >= 3.14159)}
            var v64 int32 = genFn0(v1)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi62
            _ = v63
            _ = v64
        }
    }
    v65 := GenS2{F0: 65535, F1: int32(v1)}
    switch "a" {
    case "foo bar":
        fmt.Println(v1)
    default:
        for fi66 := 0; fi66 < 3; fi66++ {
            if true {
                var v67 uint32 = v3
                var v68 int8 = 127
                _ = v67
                _ = v68
            } else {
                delete(v2, v3)
                switch "hello" {
                case "hello":
                    for fi69 := 0; fi69 < 11; fi69++ {
                        switch "foo bar" {
                        case "hello":
                            var v70 uint8 = (v1 - (v32 * (v32 << ((4294967295) & 63))))
                            _ = v70
                        case "":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "foo bar":
                            v35 += 0.1
                        default:
                            v71 := GenS6{F0: 100}
                            _ = v71
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = len(v5)
                        _ = fi69
                    }
                case "foo bar":
                    fmt.Println(v48)
                case "a":
                    switch ("foo bar" + ((((("foo bar" + ("foo bar" + "\n\t")) + "foo bar") + "a") + "") + "")) {
                    case "":
                        for fi72 := 0; fi72 < 5; fi72++ {
                            if false {
                                switch 2147483647 {
                                case 128:
                                    if ((v35 < v35) && (47 <= 0)) {
                                        v73 := GenS7{F0: 42, F1: (v3 + 0)}
                                        if false {
                                            fmt.Println(v1, fi66, v1)
                                        }
                                        _ = v73
                                    }
                                }
                                if (false && true) {
                                    v74 := GenS2{F0: (v3 ^ 944), F1: 955}
                                    _ = v74
                                } else {
                                    v32 = (v1 << ((18446744073709551615) & 63))
                                }
                                v1 -= 255
                                v4[("\n\t" + "hello")] = (v3 * v3)
                                if (-128 >= 0) {
                                    switch ((("hello" + "hello") + "") + ("" + (("a" + "") + ("hello" + (((("foo bar" + "hello") + "a") + ("" + "foo bar")) + ""))))) {
                                    case "":
                                        var v75 uint16 = 32767
                                        _ = v75
                                    default:
                                        var v76 float32 = 0.0
                                        _ = v76
                                    }
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v48 -= (fi66 * (^(fi66)))
                                }
                                if false {
                                    v48 += fi72
                                }
                            } else {
                                switch "a" {
                                case "\n\t":
                                    var v77 int32 = 42
                                    _ = v77
                                case "a":
                                    var v78 int32 = 76
                                    fmt.Println(v1, v78, v48)
                                    switch "hello" {
                                    case "\n\t":
                                        delete(v2, v3)
                                        delete(v2, v3)
                                    case "":
                                        v6 = append(v6, fi72)
                                    case "a":
                                        v79 := GenS6{F0: -125}
                                        v3 -= (v3 >> ((65535) & 63))
                                        _ = v79
                                    default:
                                        if true {
                                            v3 -= 854
                                            if (127 <= 106) {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v1 = 2
                                            }
                                        }
                                        if (true || ((v78 + (v78 >> ((65535) & 63))) <= (v78 >> ((65535) & 63)))) {
                                            var v80 uint8 = 126
                                            _ = len(v6)
                                            var v81 float64 = (0.0 / (v35 * 0.0))
                                            v32 += v80
                                            if (42 <= 4294967295) {
                                                v1 -= 235
                                                v82 := GenS8{F0: 2147483647}
                                                _ = v82
                                            }
                                            switch (-(fi66)) {
                                            case 4294967295:
                                                v83 := GenS7{F0: 1, F1: v3}
                                                _ = v83
                                            default:
                                                v84 := GenS2{F0: 366, F1: (-(v78))}
                                                _ = v84
                                            }
                                            v85 := v2[2147483648]
                                            if true {
                                                v86 := GenS7{F0: 42, F1: (v3 - (v3 + v3))}
                                                _ = v86
                                            }
                                            _ = v80
                                            _ = v81
                                            _ = v85
                                        }
                                    }
                                    _ = v78
                                case "hello":
                                    var v87 int = fi66
                                    v35 += 0.1
                                    _ = v87
                                }
                                var v88 string = ""
                                var v89 string = "\n\t"
                                _ = v88
                                _ = v89
                            }
                            v3 += 2147483648
                            _ = fi72
                        }
                    }
                default:
                    var v90 uint = 0
                    _ = v90
                }
                var v91 uint64 = 18446744073709551615
                _ = v91
            }
            if false {
                fmt.Println(v48, v48, fi66)
            }
            _ = fi66
        }
    }
    var v92 int8 = 42
    switch "\n\t" {
    case "foo bar":
        for fi93 := 0; fi93 < 6; fi93++ {
            for fi94 := 0; fi94 < 16; fi94++ {
                fmt.Println(v1)
                if true {
                    if true {
                        switch "" {
                        case "hello":
                            v95 := GenS9{F0: 42}
                            _ = v95
                        case "foo bar":
                            var v96 uint64 = 0
                            _ = v96
                        default:
                            v4[(("" + "") + ("foo bar" + ""))] = (^(v3))
                        }
                        v97 := v5[v3]
                        _ = v97
                    } else {
                        if (false && false) {
                            v98 := GenS5{F0: "", F1: 2, F2: true}
                            switch ("a" + "a") {
                            case "a":
                                v35 -= 1e10
                            case "foo bar":
                                fmt.Println(fi94)
                            case "\n\t":
                                var v99 uint16 = 909
                                _ = v99
                            default:
                                var v100 float64 = (float64(3.14159) + v35)
                                _ = v100
                            }
                            v1 += v1
                            _ = v98
                        } else {
                            v101 := GenS5{F0: "foo bar", F1: 1, F2: true}
                            _ = v101
                        }
                    }
                    var v102 uint32 = (v3 + (v3 & (v3 & 4294967295)))
                    _ = v102
                } else {
                    v92 += 0
                    if true {
                        v103 := GenS8{F0: fi93}
                        var v104 uint32 = v3
                        _ = v103
                        _ = v104
                    }
                    fmt.Println(v35, fi94, v35)
                    fmt.Println(v48, v1)
                    v92 = v92
                }
                if (-792 != 1) {
                    v4["foo bar"] = v3
                    var v105 int = (^(fi94))
                    _ = v105
                } else {
                    if (v35 > ((v35 * 0.0) * (v35 + v35))) {
                        var v106 int16 = 32767
                        var v107 int32 = 0
                        _ = v106
                        _ = v107
                    }
                    if (!true) {
                        v35 = -2.25
                    }
                    if len(v6) > 0 {
                        v6[0] = -176
                    }
                }
                _ = fi94
            }
            v35 = 0.0
            v48 += (fi93 & (fi93 << ((255) & 63)))
            _ = fi93
        }
    default:
        v92 -= v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v65
    _ = v92
}
