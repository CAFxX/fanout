package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int32; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 bool; F2 int64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS13 struct { F0 uint64 }
func genFn0(p0 string, p1 int32, p2 bool) int16 {
    v5 := []int{32767, 32768, 42, -2147483648}
    p0 = p0
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: -1, F1: 65535, F2: 65535}
            p0 = "a"
            _ = fi7
            _ = v8
        }
        p0 = p0
        if (!(42 < 1000)) {
            var v9 uint8 = 127
            if true {
                p0 = p0
                v5 = append(v5, fi6)
                fmt.Println(p2, p1, p2)
            }
            _ = v9
        } else {
            var v10 int8 = -76
            var v11 float64 = (0.1 / float64((3.14159 - (1e10 / -2.25))))
            var v12 int16 = 32767
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 string = ("foo bar" + "hello")
    var v4 int16 = 42
    var v13 int = -9223372036854775808
    var v14 uint16 = 779
    var v15 uint32 = 176
    switch -546 {
    case -1:
        for fi16 := 0; fi16 < 7; fi16++ {
            if (!false) {
                fmt.Println(v4, fi16)
            } else {
                for fi17 := 0; fi17 < 12; fi17++ {
                    v14 = v14
                    _ = fi17
                }
                switch 381 {
                case 127:
                    for fi18 := 0; fi18 < 17; fi18++ {
                        if true {
                            fmt.Println(v3, v4, v14)
                            fmt.Println(v3)
                        }
                        v14 += 99
                        v19 := GenS2{F0: -28, F1: (v15 ^ 0)}
                        _ = fi18
                        _ = v19
                    }
                case 32767:
                    var v20 uint64 = 4294967295
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := v2[v3]
                fmt.Println(fi21, fi16, v22)
                v13 -= fi16
                _ = fi21
                _ = v22
            }
            _ = fi16
        }
    case 127:
        v23 := v2[""]
        _ = v23
    default:
        delete(v1, 684)
    }
    switch ("foo bar" + v3) {
    case "foo bar":
        fmt.Println(v14)
    }
    switch v13 {
    case 2147483647:
        fmt.Println(v15, v3, v4)
    case 42:
        v4 = v4
    default:
        if (false || (2147483648 == v15)) {
            v24 := GenS1{F0: -128, F1: 2, F2: (v15 & (v15 & 2))}
            v3 = "hello"
            _ = v24
        } else {
            for fi25 := 0; fi25 < 15; fi25++ {
                switch (("\n\t" + "foo bar") + v3) {
                case "hello":
                    for fi26 := 0; fi26 < 3; fi26++ {
                        if false {
                            v2[(((v3 + v3) + "\n\t") + (v3 + v3))] = (52 != 42)
                            v4 += int16(v14)
                            fmt.Println(v3, fi25)
                            var v27 int32 = 0
                            var v28 uint8 = 255
                            _ = v27
                            _ = v28
                        }
                        fmt.Println(v15, v4)
                        _ = fi26
                    }
                case "":
                    for fi29 := 0; fi29 < 14; fi29++ {
                        if true {
                            switch "" {
                            case "a":
                                fmt.Println(v13)
                            case "":
                                if (!(9223372036854775807 > 9223372036854775807)) {
                                    v30 := GenS3{F0: 42}
                                    v4 = 1
                                    switch v3 {
                                    case "\n\t":
                                        v31 := GenS2{F0: -128, F1: v15}
                                        _ = v31
                                    default:
                                        var v32 float64 = (((((((1e10 * 1.5) - (0.0 + (-2.25 + -2.25))) - 1e10) - 0.0) * ((0.0 + 0.1) * 1e10)) + (1.5 * float64(1.5))) * float64(3.14159))
                                        v3 = v3
                                        if ((v15 >= (v15 / ((724) | 1))) || true) {
                                            v13 = fi29
                                            if (v4 != v4) {
                                                switch v3 {
                                                case "foo bar":
                                                    var v33 uint16 = (v14 - (v14 - 255))
                                                    _ = v33
                                                }
                                            }
                                            if false {
                                                v34 := GenS3{F0: 2147483647}
                                                if (-808 < -2147483648) {
                                                    var v35 int = fi29
                                                    var v36 int64 = 152
                                                    _ = v35
                                                    _ = v36
                                                }
                                                if (9223372036854775807 < -563) {
                                                    v4 -= (v4 * 1000)
                                                    switch (("foo bar" + (v3 + v3)) + (v3 + ((v3 + (v3 + "\n\t")) + v3))) {
                                                    case "hello":
                                                        v15 += v15
                                                    case "\n\t":
                                                        var v37 string = v3
                                                        _ = v37
                                                    case "a":
                                                        var v38 int = (-(v13))
                                                        _ = v38
                                                    }
                                                }
                                                _ = v34
                                            } else {
                                                var v39 uint16 = 65535
                                                _ = v39
                                            }
                                            switch (v3 + v3) {
                                            case "foo bar":
                                                switch int(v4) {
                                                case 32768:
                                                    v40 := GenS4{F0: v4, F1: 2147483647, F2: -686}
                                                    _ = v40
                                                case -645:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v41 := GenS1{F0: 0, F1: (^(v14)), F2: (v15 << ((18446744073709551615) & 63))}
                                                        _ = v41
                                                    case "":
                                                        var v42 bool = (613 < 2)
                                                        v4 = v4
                                                        v43 := GenS2{F0: 123, F1: 780}
                                                        if v42 {
                                                            v44 := GenS2{F0: int8(v4), F1: v15}
                                                            _ = v44
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    case "a":
                                                        v13 = (fi25 | (fi25 | (v13 / (((fi29 * v13)) | 1))))
                                                    }
                                                default:
                                                    v45 := GenS2{F0: -1, F1: (v15 + (v15 + (v15 * 42)))}
                                                    _ = v45
                                                }
                                            default:
                                                var v46 int64 = 169
                                                _ = v46
                                            }
                                            v47 := GenS5{F0: int64(fi29)}
                                            _ = v47
                                        } else {
                                            _ = len(v1)
                                            v32 = (v32 - v32)
                                            switch fi25 {
                                            case 431:
                                                _ = len(v1)
                                            case 9223372036854775807:
                                                v48 := v2[v3]
                                                _ = v48
                                            case 2147483647:
                                                v15 = 42
                                            default:
                                                if ((true && false) || (42 != 255)) {
                                                    var v49 int64 = -202
                                                    _ = v49
                                                }
                                            }
                                        }
                                        _ = v32
                                    }
                                    _ = v30
                                } else {
                                    v15 = (v15 >> ((355) & 63))
                                    var v50 uint8 = 42
                                    _ = v50
                                }
                            }
                        } else {
                            v51 := GenS6{F0: 65535}
                            var v52 float32 = (-2.25 + 0.0)
                            switch int(v4) {
                            case -2147483648:
                                v3 = ("\n\t" + ((((v3 + (((((((v3 + "") + v3) + (v3 + v3)) + "foo bar") + (v3 + "\n\t")) + v3) + v3)) + v3) + (((((v3 + (v3 + v3)) + ("" + v3)) + "") + v3) + "a")) + ("foo bar" + "\n\t")))
                            case 32768:
                                switch 9223372036854775807 {
                                case -1:
                                    v3 = "hello"
                                case 714:
                                    var v53 int16 = (v4 & (v4 - 0))
                                    _ = v53
                                case -623:
                                    if (("hello" + v3) > v3) {
                                        switch v3 {
                                        case "\n\t":
                                            switch (fi25 - 42) {
                                            case 183:
                                                fmt.Println(v13, v52)
                                            }
                                            v54 := GenS6{F0: (v15 + v15)}
                                            v55 := []int{(fi25 << ((275) & 63)), (fi25 / ((0) | 1)), (fi29 / (((fi25 * int(v15))) | 1)), (fi25 % (((fi29 << ((18446744073709551615) & 63))) | 1))}
                                            _ = v54
                                            _ = v55
                                        default:
                                            v56 := GenS7{F0: 1, F1: ((v3 + ("a" + "a")) != "hello"), F2: -766}
                                            _ = v56
                                        }
                                        if (!(true || (817 > 398))) {
                                            v3 = ((v3 + "") + v3)
                                        } else {
                                            var v57 uint32 = (v15 >> ((4294967295) & 63))
                                            v58 := GenS4{F0: genFn0("hello", -1, true), F1: -1, F2: 0}
                                            if (false && ((v3 + v3) == "a")) {
                                                v59 := []int{(v13 | (^(v13))), (^(fi29)), v13}
                                                v57 = (v57 | (v15 / (((^(v57))) | 1)))
                                                var v60 int = fi25
                                                _ = v59
                                                _ = v60
                                            } else {
                                                if false {
                                                    v14 = (v14 << ((308) & 63))
                                                    if true {
                                                        var v61 float32 = 1.5
                                                        _ = v61
                                                    } else {
                                                        if (986 != 42) {
                                                            var v62 float32 = 1e10
                                                            _ = v62
                                                        }
                                                    }
                                                    v4 = v4
                                                    switch v3 {
                                                    case "a":
                                                        if false {
                                                            v63 := GenS8{F0: 0}
                                                            if false {
                                                                switch ("a" + v3) {
                                                                case "a":
                                                                    var v64 int64 = 0
                                                                    _ = v64
                                                                case "hello":
                                                                    v52 += (((1e10 + 0.1) - v52) * v52)
                                                                case "foo bar":
                                                                    delete(v1, (v13 - v13))
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                                v65 := []int{(v13 | 1), fi25, (v13 ^ fi29)}
                                                                v3 = (((("hello" + (v3 + "\n\t")) + v3) + (("hello" + v3) + v3)) + "\n\t")
                                                                _ = v65
                                                            } else {
                                                                v14 = v14
                                                                v66 := GenS9{F0: 281, F1: 0, F2: 0}
                                                                v15 = (^(v15))
                                                                _ = v66
                                                            }
                                                            var v67 bool = false
                                                            _ = v63
                                                            _ = v67
                                                        }
                                                    case "foo bar":
                                                        v14 -= v14
                                                    case "":
                                                        var v68 uint = 172
                                                        _ = v68
                                                    }
                                                    var v69 int64 = 9223372036854775807
                                                    var v70 uint = 0
                                                    var v71 uint = v70
                                                    v72 := []int{fi25, (fi29 ^ (fi29 ^ 1)), fi29}
                                                    _ = v69
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                }
                                                switch v3 {
                                                case "a":
                                                    var v73 int16 = genFn0(((v3 + v3) + ((v3 + v3) + ("a" + "a"))), 411, (!((true && (1 == 42)) || true)))
                                                    _ = v73
                                                case "":
                                                    v3 = v3
                                                default:
                                                    v15 -= 1
                                                }
                                            }
                                            _ = v57
                                            _ = v58
                                        }
                                        var v74 int32 = 975
                                        switch fi25 {
                                        case -909:
                                            if true {
                                                _ = len(v2)
                                                if true {
                                                    switch (v3 + (v3 + v3)) {
                                                    case "a":
                                                        v75 := []int{(-(v13)), -221, (v13 | -2147483648), int(v15)}
                                                        _ = v75
                                                    case "":
                                                        fmt.Println(fi29)
                                                    default:
                                                        v76 := []int{}
                                                        _ = v76
                                                    }
                                                    v14 = 296
                                                    v77 := GenS9{F0: (v14 + v14), F1: 1, F2: 255}
                                                    _ = v77
                                                } else {
                                                    switch (v13 - fi29) {
                                                    case 9223372036854775807:
                                                        switch v13 {
                                                        case -2147483648:
                                                            v14 = (v14 ^ 65535)
                                                        default:
                                                            v78 := []int{(fi29 & 42), (fi25 / (((fi25 & fi29)) | 1))}
                                                            _ = v78
                                                        }
                                                    case 1:
                                                        fmt.Println(v74, fi25)
                                                    default:
                                                        var v79 uint16 = (v14 + (v14 - 1))
                                                        _ = v79
                                                    }
                                                    fmt.Println(v52, v15, v3)
                                                    if true {
                                                        if (!((-9223372036854775808 <= 1) && true)) {
                                                            fmt.Println(v4)
                                                        } else {
                                                            var v80 int = fi29
                                                            v81 := GenS10{F0: 161}
                                                            switch fi29 {
                                                            case -51:
                                                                v14 -= 65535
                                                            case 127:
                                                                var v82 uint32 = (v15 - (^(v15)))
                                                                _ = v82
                                                            default:
                                                                fmt.Println(v80, v3, v4)
                                                            }
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        v3 = v3
                                                    } else {
                                                        v83 := GenS2{F0: 1, F1: 328}
                                                        v4 -= v4
                                                        _ = v83
                                                    }
                                                    if true {
                                                        var v84 int64 = 9223372036854775807
                                                        v3 = "a"
                                                        fmt.Println(v15, v52)
                                                        _ = v84
                                                    }
                                                    var v85 float32 = v52
                                                    v4 = (v4 / (((v4 ^ v4)) | 1))
                                                    var v86 bool = true
                                                    _ = v85
                                                    _ = v86
                                                }
                                                switch v3 {
                                                case "":
                                                    var v87 uint = 18446744073709551615
                                                    _ = v87
                                                case "foo bar":
                                                    var v88 uint8 = 42
                                                    _ = v88
                                                case "hello":
                                                    var v89 int32 = -2147483648
                                                    _ = v89
                                                default:
                                                    var v90 int8 = -128
                                                    _ = v90
                                                }
                                                v74 = v74
                                            }
                                        case 686:
                                            switch "" {
                                            case "":
                                                v91 := GenS6{F0: v15}
                                                _ = v91
                                            default:
                                                if (!false) {
                                                    var v92 uint32 = 4294967295
                                                    _ = v92
                                                } else {
                                                    v2[(v3 + "")] = (!((1e10 - 1e10) < (0.1 + (-2.25 / 1.0))))
                                                    if (v74 < 0) {
                                                        v15 -= v15
                                                    } else {
                                                        fmt.Println(v74, fi25)
                                                        v93 := GenS11{F0: 4294967295}
                                                        _ = v93
                                                    }
                                                }
                                            }
                                        case 620:
                                            var v94 uint64 = 1
                                            _ = v94
                                        default:
                                            var v95 uint8 = 90
                                            _ = v95
                                        }
                                        fmt.Println(v74, v52)
                                        _ = v74
                                    }
                                }
                            default:
                                if (float32((3.14159 * (1e10 / 1.0))) >= v52) {
                                    v14 = v14
                                    delete(v1, (fi25 + (v13 % ((fi25) | 1))))
                                    v15 = (v15 ^ (v15 + v15))
                                } else {
                                    var v96 int16 = genFn0(("\n\t" + ((v3 + ((v3 + "hello") + (v3 + v3))) + "\n\t")), -1, false)
                                    var v97 uint8 = 42
                                    var v98 uint = 508
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                }
                            }
                            _ = v51
                            _ = v52
                        }
                        _ = fi29
                    }
                case "a":
                    for fi99 := 0; fi99 < 14; fi99++ {
                        if (!((false || (0 > 2)) || false)) {
                            fmt.Println(fi25, v3)
                            v100 := GenS12{F0: 9223372036854775807, F1: v15, F2: 42}
                            fmt.Println(fi25, fi25, v3)
                            _ = v100
                        }
                        v101 := GenS13{F0: 177}
                        _ = fi99
                        _ = v101
                    }
                }
                _ = fi25
            }
            v102 := GenS1{F0: -1, F1: 32767, F2: v15}
            _ = v102
        }
    }
    for fi103 := 0; fi103 < 17; fi103++ {
        var v104 float64 = 0.0
        switch 4294967295 {
        case 32767:
            for fi105 := 0; fi105 < 10; fi105++ {
                v106 := GenS3{F0: 952}
                fmt.Println(v14, v14, v104)
                if true {
                    var v107 uint32 = v15
                    _ = v107
                }
                _ = fi105
                _ = v106
            }
        default:
            fmt.Println(v4)
        }
        _ = fi103
        _ = v104
    }
    var v108 int16 = (v4 / ((v4) | 1))
    v15 = v15
    v109 := GenS13{F0: 4294967295}
    var v110 float64 = 3.14159
    for fi111 := 0; fi111 < 19; fi111++ {
        v1[v13] = ("a" + "foo bar")
        var v112 uint32 = v15
        _ = fi111
        _ = v112
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v110:", v110)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
