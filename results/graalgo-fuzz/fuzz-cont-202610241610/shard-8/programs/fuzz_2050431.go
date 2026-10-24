package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 uint; F1 int16 }
type GenS10 struct { F0 int; F1 int16; F2 uint16 }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 uint32; F1 uint64; F2 int16 }
func genFn0(p0 uint8, p1 uint, p2 int64) int64 {
    p2 = (-(p2))
    v7 := []int{}
    p0 = (p0 & 2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int32 = -2147483648
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 bool = true
    var v6 uint16 = 1
    v3["\n\t"] = "hello"
    for fi8 := 0; fi8 < 7; fi8++ {
        v5 = false
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: 18446744073709551615, F1: (("" + "foo bar") + ("\n\t" + "foo bar")), F2: (v6 << ((v1) & 63))}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS1{F0: 1, F1: "a", F2: (v6 / ((v6) | 1))}
    switch 128 {
    case 127:
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                v2 = 2147483647
                _ = fi13
            }
            v6 = (v6 | (v6 * v6))
            _ = fi12
        }
    }
    var v14 int32 = (v2 << ((v1) & 63))
    fmt.Println(v14, v2)
    for fi15 := 0; fi15 < 5; fi15++ {
        v2 = (^(v14))
        v1 -= (^(v1))
        _ = fi15
    }
    var v16 int = 2147483647
    var v17 bool = (64 != 0)
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint32 = 65535
            _ = fi19
            _ = v20
        }
        if (!v5) {
            var v21 uint16 = (v6 % (((^(v6))) | 1))
            if true {
                if (true && v17) {
                    if (!(0.1 == 1.5)) {
                        v22 := GenS1{F0: 2, F1: "a", F2: v21}
                        _ = v22
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            if (("a" + "hello") > "\n\t") {
                                var v23 float32 = 1.5
                                switch ("" + ("hello" + "")) {
                                case "\n\t":
                                    v1 += v1
                                case "foo bar":
                                    var v24 uint64 = 2
                                    _ = v24
                                case "":
                                    for fi25 := 0; fi25 < 8; fi25++ {
                                        v26 := GenS2{F0: 1}
                                        var v27 int = (fi18 % ((526) | 1))
                                        var v28 int16 = 42
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                default:
                                    var v29 float32 = (v23 * (1e10 * ((float32((1e10 + 0.0)) * v23) - v23)))
                                    _ = v29
                                }
                                _ = v23
                            } else {
                                if (genFn0(1, 184, 0) <= 691) {
                                    v30 := GenS3{F0: 246}
                                    v31 := []int{fi18, (^(fi18)), v16}
                                    var v32 string = "foo bar"
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                var v33 bool = true
                                _ = v33
                            }
                        }
                        if (false || v17) {
                            switch 127 {
                            case -2147483648:
                                v34 := []int{(fi18 >> (((v1 >> (((v1 * v1)) & 63))) & 63)), (^(fi18))}
                                _ = v34
                            default:
                                for fi35 := 0; fi35 < 1; fi35++ {
                                    var v36 string = (("foo bar" + "") + "\n\t")
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        }
                        if false {
                            var v37 uint = (^(v1))
                            for fi38 := 0; fi38 < 11; fi38++ {
                                v39 := v4["foo bar"]
                                _ = fi38
                                _ = v39
                            }
                            _ = v37
                        }
                    }
                    var v40 float32 = (((0.0 / (-2.25 * 0.1)) + 3.14159) - 1e10)
                    _ = v40
                } else {
                    v41 := v4[""]
                    v42 := GenS4{F0: 42, F1: 42, F2: 65535}
                    _ = v41
                    _ = v42
                }
                v43 := GenS5{F0: 28, F1: 1}
                for fi44 := 0; fi44 < 15; fi44++ {
                    if ((v1 - 42) != (^(v1))) {
                        v45 := []int{}
                        _ = v45
                    }
                    _ = fi44
                }
                _ = v43
            } else {
                var v46 uint8 = 0
                _ = v46
            }
            _ = v21
        } else {
            if (v17 && ((128 <= 87) || ("hello" != "\n\t"))) {
                if (100 > 91) {
                    var v47 float32 = 1e10
                    v4[(("" + "") + "\n\t")] = (!(v47 > 0.0))
                    var v48 int64 = -300
                    _ = v47
                    _ = v48
                } else {
                    fmt.Println(v6, v16)
                }
                var v49 uint64 = 42
                v50 := GenS3{F0: 1}
                _ = v49
                _ = v50
            } else {
                v51 := GenS2{F0: 0}
                for fi52 := 0; fi52 < 10; fi52++ {
                    delete(v3, "")
                    var v53 int = 1
                    v54 := []int{}
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = v51
            }
            v55 := []int{int(v1)}
            _ = v55
        }
        var v56 uint32 = 4294967295
        _ = fi18
        _ = v56
    }
    var v57 float64 = (-2.25 / 1.5)
    if false {
        for fi58 := 0; fi58 < 11; fi58++ {
            if (v5 || (543 >= 2147483648)) {
                var v59 float32 = (0.1 - 1.5)
                _ = v59
            }
            switch (fi58 & (-(fi58))) {
            case 2147483647:
                switch ("hello" + "foo bar") {
                case "foo bar":
                    if v5 {
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            fmt.Println(v14, v17)
                        case "hello":
                            for fi60 := 0; fi60 < 8; fi60++ {
                                v57 += (float64(float32((v57 - v57))) / (v57 + (0.0 - v57)))
                                var v61 int64 = 89
                                if v17 {
                                    var v62 float32 = -2.25
                                    _ = v62
                                } else {
                                    var v63 uint = v1
                                    v6 += v6
                                    var v64 int8 = int8(v61)
                                    _ = v63
                                    _ = v64
                                }
                                _ = fi60
                                _ = v61
                            }
                        default:
                            var v65 uint = (^(v1))
                            _ = v65
                        }
                        if (-1 > 952) {
                            fmt.Println(v5, v1)
                        }
                    } else {
                        fmt.Println(v17)
                        v5 = v17
                        fmt.Println(v17)
                    }
                }
            }
            _ = fi58
        }
        for fi66 := 0; fi66 < 17; fi66++ {
            for fi67 := 0; fi67 < 7; fi67++ {
                var v68 uint8 = 128
                _ = fi67
                _ = v68
            }
            for fi69 := 0; fi69 < 10; fi69++ {
                v70 := []int{(-(v16)), (^(fi66)), (^(fi69))}
                fmt.Println(v6, v6)
                switch "" {
                case "hello":
                    if (v5 && (v5 || (!(!(!false))))) {
                        fmt.Println(v16)
                    } else {
                        if v17 {
                            switch (fi69 + int(v2)) {
                            case -1:
                                v71 := v3[(("" + "foo bar") + ((("\n\t" + "a") + ("" + "")) + ""))]
                                _ = v71
                            case 778:
                                v72 := GenS5{F0: 18446744073709551615, F1: -76}
                                _ = v72
                            case 0:
                                v4[("hello" + "hello")] = (false && (18446744073709551615 <= 888))
                                var v73 bool = (!v5)
                                v74 := GenS3{F0: 255}
                                _ = v73
                                _ = v74
                            default:
                                v75 := GenS6{F0: (v6 / ((209) | 1)), F1: (!(int32(v6) != (v14 + v14)))}
                                _ = v75
                            }
                        }
                        v76 := GenS6{F0: (v6 / ((v6) | 1)), F1: (379 >= 1)}
                        if false {
                            if (v17 || v17) {
                                if v5 {
                                    if len(v70) > 0 {
                                        v70[0] = v16
                                    }
                                } else {
                                    v6 += 140
                                    v1 += (v1 - v1)
                                    var v77 int = (^(v16))
                                    _ = v77
                                }
                                if v5 {
                                    v57 += float64(float32(v57))
                                    var v78 int8 = 7
                                    var v79 uint32 = 497
                                    _ = v78
                                    _ = v79
                                } else {
                                    var v80 string = ((("a" + "a") + ("" + "a")) + "")
                                    v81 := GenS7{F0: 42}
                                    _ = v80
                                    _ = v81
                                }
                            } else {
                                var v82 uint64 = 42
                                if v17 {
                                    fmt.Println(fi69, v16, v6)
                                    v14 -= -582
                                    fmt.Println(v17, v17)
                                }
                                _ = v82
                            }
                            if ((fi66 ^ (^(fi69))) <= fi69) {
                                fmt.Println(v14)
                                v16 = 2147483647
                            } else {
                                v6 -= (v6 / (((v6 >> (((v1 & v1)) & 63))) | 1))
                                v16 = (fi69 % (((fi66 / ((v16) | 1))) | 1))
                                if false {
                                    v83 := GenS8{F0: (v6 & (v6 | 2)), F1: (v6 ^ 1)}
                                    v84 := GenS9{F0: (^(v1)), F1: -898}
                                    v85 := GenS6{F0: 0, F1: v5}
                                    v86 := GenS10{F0: (fi66 / ((v16) | 1)), F1: 1, F2: 42}
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                } else {
                                    switch fi69 {
                                    case -516:
                                        v2 = (v14 | (v14 | (v14 & (v14 & 2147483647))))
                                    }
                                    switch "hello" {
                                    case "":
                                        v87 := GenS1{F0: 4294967295, F1: "hello", F2: v6}
                                        _ = v87
                                    case "a":
                                        if (((!(v6 == (v6 << ((65535) & 63)))) || v5) || true) {
                                            if v5 {
                                                if false {
                                                    if (!v17) {
                                                        v88 := GenS3{F0: 255}
                                                        _ = v88
                                                    }
                                                    var v89 int8 = 27
                                                    _ = v89
                                                } else {
                                                    var v90 uint32 = 784
                                                    var v91 string = "hello"
                                                    v92 := GenS11{F0: "hello", F1: 42}
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                }
                                                if (127 >= 167) {
                                                    var v93 uint8 = 223
                                                    v2 = -301
                                                    _ = v93
                                                } else {
                                                    delete(v4, ("hello" + ("a" + "\n\t")))
                                                    if v5 {
                                                        v94 := GenS8{F0: (v6 & v6), F1: v6}
                                                        _ = v94
                                                    } else {
                                                        var v95 uint64 = 4294967295
                                                        var v96 string = "\n\t"
                                                        if v5 {
                                                            v97 := GenS8{F0: 727, F1: (v6 - uint16(v1))}
                                                            if (v17 || (-2.25 > 1.5)) {
                                                                if false {
                                                                    var v98 string = "a"
                                                                    v3[""] = "\n\t"
                                                                    switch v98 {
                                                                    case "\n\t":
                                                                        v14 += v14
                                                                    default:
                                                                        switch fi69 {
                                                                        case -1:
                                                                            if true {
                                                                                switch 0 {
                                                                                case 32767:
                                                                                    var v99 string = ""
                                                                                    _ = v99
                                                                                case 136:
                                                                                    if (false && false) {
                                                                                        if false {
                                                                                            if (-7 != 107) {
                                                                                                v16 = (fi66 * fi66)
                                                                                            } else {
                                                                                                var v100 uint8 = 106
                                                                                                v95 = 2
                                                                                                var v101 bool = (v5 || v17)
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v3[(v98 + v96)] = ("hello" + v98)
                                                                                    }
                                                                                default:
                                                                                    switch v98 {
                                                                                    case "\n\t":
                                                                                        delete(v4, ("hello" + "a"))
                                                                                    }
                                                                                }
                                                                                fmt.Println(v2, v1)
                                                                                var v102 float32 = (-2.25 + ((3.14159 + -2.25) / 0.1))
                                                                                _ = v102
                                                                            } else {
                                                                                v14 = v14
                                                                            }
                                                                        case 1:
                                                                            delete(v3, ((v96 + "a") + v98))
                                                                        case 4294967295:
                                                                            fmt.Println(v57)
                                                                        }
                                                                    }
                                                                    v103 := GenS2{F0: -9223372036854775808}
                                                                    var v104 float64 = v57
                                                                    _ = v103
                                                                    _ = v104
                                                                    _ = v98
                                                                }
                                                            }
                                                            _ = v97
                                                        }
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                }
                                            }
                                            v105 := GenS12{F0: 2147483648, F1: 42, F2: -32768}
                                            var v106 int8 = -1
                                            v1 = (v1 << ((v1) & 63))
                                            _ = v105
                                            _ = v106
                                        }
                                        var v107 int8 = 127
                                        _ = v107
                                    default:
                                        var v108 int8 = 100
                                        _ = v108
                                    }
                                    var v109 float64 = 1e10
                                    _ = v109
                                }
                                v14 -= (v14 % ((v2) | 1))
                            }
                        }
                        if v5 {
                            var v110 uint32 = 2
                            fmt.Println(v14)
                            v14 = v14
                            _ = v110
                        }
                        _ = v76
                    }
                }
                v111 := v4["hello"]
                _ = fi69
                _ = v111
                _ = v70
            }
            _ = fi66
        }
        v112 := v3["a"]
        _ = v112
    } else {
        switch (v16 + 0) {
        case -9223372036854775808:
            var v113 uint = v1
            _ = v113
        case -2147483648:
            for fi114 := 0; fi114 < 2; fi114++ {
                fmt.Println(v57, v57, v17)
                v115 := []int{}
                _ = fi114
                _ = v115
            }
        case -1:
            v5 = v17
        }
        if v5 {
            v57 -= v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
