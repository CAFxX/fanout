package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 string; F1 int32; F2 bool }
type GenS5 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint }
type GenS8 struct { F0 int8 }
func genFn0() int {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 0, F1: 955, F2: 2}
        switch "foo bar" {
        case "\n\t":
            var v7 uint64 = 941
            _ = v7
        default:
            var v8 int32 = -2147483648
            _ = v8
        }
        var v9 uint16 = 599
        _ = fi5
        _ = v6
        _ = v9
    }
    var v10 int8 = -88
    fmt.Println(v10)
    if true {
        if false {
            for fi11 := 0; fi11 < 4; fi11++ {
                v10 += 100
                var v12 int16 = 1
                v13 := GenS2{F0: 65535, F1: ((("foo bar" + "foo bar") + ("\n\t" + "foo bar")) + ((("hello" + "") + "") + "hello")), F2: (v10 % ((v10) | 1))}
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 int32 = -957
            _ = v14
        } else {
            fmt.Println(v10)
            var v15 uint32 = 274
            if false {
                var v16 bool = (true || (((v10 & v10) != 127) || ((4294967295 > 4294967295) && (-1 < v10))))
                _ = v16
            } else {
                switch ("foo bar" + ("\n\t" + (("foo bar" + "hello") + (((((("a" + ("\n\t" + "\n\t")) + "foo bar") + "") + "a") + "") + (((("a" + (("foo bar" + (("a" + "hello") + "hello")) + "a")) + "") + ((("hello" + ("a" + "a")) + "\n\t") + "")) + ((("hello" + "hello") + ("" + ("foo bar" + "a"))) + "a")))))) {
                case "hello":
                    v17 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807, F2: 4294967295}
                    _ = v17
                case "":
                    v10 = v10
                default:
                    fmt.Println(v10, v10)
                }
                switch "hello" {
                case "":
                    fmt.Println(v15, v15)
                case "hello":
                    v18 := []int{9223372036854775807, 32767, -2147483648}
                    _ = v18
                }
            }
            _ = v15
        }
    }
    if true {
        switch "\n\t" {
        case "a":
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := GenS2{F0: 198, F1: "", F2: v10}
                if true {
                    var v21 uint64 = 268
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
        default:
            v22 := GenS2{F0: 2, F1: "hello", F2: (v10 / ((127) | 1))}
            _ = v22
        }
    } else {
        var v23 int32 = 42
        fmt.Println(v23, v10)
        v24 := GenS2{F0: 118, F1: "hello", F2: 0}
        _ = v23
        _ = v24
    }
    _ = v10
    return -9223372036854775808
}
func genFn1(p0 string, p1 uint8) uint {
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS1{F0: 65535, F1: -9223372036854775808, F2: 1}
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: 2, F1: ((p0 + ("" + "hello")) + (p0 + p0)), F2: -10}
    p1 = p1
    if (!(int64(p1) <= 0)) {
        p0 = (p0 + (p0 + p0))
    } else {
        switch 0 {
        case 0:
            fmt.Println(p0)
        case 42:
            var v28 int16 = 32767
            _ = v28
        }
    }
    _ = p0
    _ = p1
    _ = v27
    return 183
}

func main() {
    var v1 uint = 799
    var v2 bool = (!((!(32768 > 128)) || true))
    v3 := []uint32{1, 573, 2, 65535}
    v4 := []int{32767, 128, 26}
    for fi29 := 0; fi29 < 17; fi29++ {
        switch (-(fi29)) {
        case 9223372036854775807:
            v1 = v1
        case 2147483647:
            for fi30 := 0; fi30 < 9; fi30++ {
                if v2 {
                    var v31 uint16 = 65535
                    switch (fi30 ^ fi30) {
                    case 679:
                        v32 := GenS3{F0: 1, F1: 758}
                        _ = v32
                    case 363:
                        if v2 {
                            v2 = (384 <= 2147483648)
                        } else {
                            if v2 {
                                var v33 bool = (((1.5 > 0.0) || (!(v2 && false))) && v2)
                                v33 = (!false)
                                _ = len(v4)
                                _ = v33
                            } else {
                                if (1 == 1) {
                                    var v34 uint64 = 4294967295
                                    if true {
                                        v2 = v2
                                    } else {
                                        var v35 uint8 = 128
                                        var v36 uint16 = (v31 | (v31 % (((v31 << (((^(v1))) & 63))) | 1)))
                                        var v37 uint16 = uint16(v34)
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    if v2 {
                                        if (!(v31 == (v31 / ((v31) | 1)))) {
                                            var v38 bool = (v2 && v2)
                                            if v2 {
                                                if len(v3) > 0 {
                                                    v3[0] = 4294967295
                                                }
                                            }
                                            fmt.Println(fi29)
                                            _ = v38
                                        }
                                    } else {
                                        v39 := GenS2{F0: v1, F1: ("a" + ("" + "foo bar")), F2: 42}
                                        _ = v39
                                    }
                                    _ = v34
                                }
                                if len(v4) > 0 {
                                    v4[0] = fi30
                                }
                                var v40 uint32 = 988
                                if v2 {
                                    v31 -= (^(v31))
                                    v41 := GenS2{F0: 255, F1: "\n\t", F2: 61}
                                    v1 = 4294967295
                                    _ = v41
                                }
                                v40 = 12
                                _ = v40
                            }
                            v1 += 0
                            v1 = 1
                        }
                    default:
                        if (true && ((v31 - (v31 & 42)) <= v31)) {
                            _ = len(v4)
                            var v42 uint8 = 42
                            if false {
                                _ = len(v3)
                            } else {
                                fmt.Println(v2, v42)
                                var v43 string = "a"
                                _ = v43
                            }
                            _ = v42
                        }
                    }
                    switch fi30 {
                    case -1:
                        if len(v4) > 0 {
                            v4[0] = 32767
                        }
                    case -2147483648:
                        var v44 uint16 = v31
                        _ = v44
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = v31
                }
                switch ("foo bar" + "") {
                case "hello":
                    if (!v2) {
                        v45 := GenS1{F0: v1, F1: -456, F2: 0}
                        var v46 int8 = -78
                        switch ("foo bar" + ((("\n\t" + "a") + "a") + "\n\t")) {
                        case "\n\t":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            var v47 int = fi30
                            _ = v47
                        }
                        _ = v45
                        _ = v46
                    } else {
                        _ = len(v3)
                        var v48 string = "\n\t"
                        _ = v48
                    }
                default:
                    var v49 float32 = 3.14159
                    _ = v49
                }
                v4 = append(v4, (fi29 & genFn0()))
                _ = fi30
            }
        }
        _ = fi29
    }
    var v50 int8 = 42
    fmt.Println(v1, v1, v50)
    for fi51 := 0; fi51 < 4; fi51++ {
        if (42 > -2147483648) {
            if (2147483647 < 824) {
                switch fi51 {
                case 474:
                    for fi52 := 0; fi52 < 3; fi52++ {
                        v53 := GenS2{F0: (v1 % ((v1) | 1)), F1: "hello", F2: v50}
                        fmt.Println(v1, v50)
                        if v2 {
                            switch -9223372036854775808 {
                            case 127:
                                v54 := GenS2{F0: (v1 ^ (v1 << ((v1) & 63))), F1: ("a" + ("hello" + "")), F2: v50}
                                _ = v54
                            case -1:
                                v50 = v50
                            }
                            v55 := GenS4{F0: "foo bar", F1: 0, F2: true}
                            if (2147483647 < -2147483648) {
                                if (4294967295 == 0) {
                                    fmt.Println(v1, v1, v1)
                                    var v56 string = ""
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    _ = v56
                                }
                            }
                            _ = v55
                        }
                        _ = fi52
                        _ = v53
                    }
                case 32767:
                    if (false && ((uint32(v1) > uint32(v1)) || v2)) {
                        if (("a" == "foo bar") || v2) {
                            var v57 int32 = -14
                            if v2 {
                                v57 -= (^(v57))
                            }
                            for fi58 := 0; fi58 < 3; fi58++ {
                                var v59 int8 = 39
                                _ = fi58
                                _ = v59
                            }
                            _ = v57
                        } else {
                            var v60 string = "\n\t"
                            _ = v60
                        }
                    }
                default:
                    for fi61 := 0; fi61 < 7; fi61++ {
                        fmt.Println(fi51)
                        if v2 {
                            v62 := GenS1{F0: genFn1("a", uint8(v1)), F1: 1, F2: 0}
                            if ((-174 <= -197) || v2) {
                                var v63 int32 = -2147483648
                                _ = v63
                            } else {
                                v2 = v2
                                v50 = (v50 / ((v50) | 1))
                            }
                            _ = len(v3)
                            _ = v62
                        }
                        _ = fi61
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v2, fi51, v2)
            } else {
                var v64 bool = (fi51 == genFn0())
                v65 := GenS3{F0: -1, F1: 1}
                var v66 uint8 = 2
                _ = v64
                _ = v65
                _ = v66
            }
            v3 = append(v3, 1)
        }
        v67 := GenS5{F0: false, F1: 2147483648, F2: 1}
        _ = fi51
        _ = v67
    }
    var v68 uint8 = 0
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v69 int16 = 42
        var v70 int64 = 0
        _ = v69
        _ = v70
    } else {
        switch "hello" {
        case "a":
            var v71 string = ("a" + "foo bar")
            _ = v71
        case "":
            v72 := GenS5{F0: v2, F1: 323, F2: 42}
            _ = v72
        default:
            var v73 uint32 = 2
            _ = v73
        }
    }
    v2 = true
    for fi74 := 0; fi74 < 2; fi74++ {
        v2 = (-52 == -573)
        fmt.Println(v50, v1)
        _ = fi74
    }
    switch "foo bar" {
    case "foo bar":
        if v2 {
            var v75 int8 = 1
            _ = v75
        }
    default:
        var v76 int16 = -32768
        _ = v76
    }
    var v77 uint = 4294967295
    for fi78 := 0; fi78 < 7; fi78++ {
        v4 = append(v4, (fi78 & 32767))
        v50 += (v50 & v50)
        _ = fi78
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v2 {
        if (!true) {
            v79 := GenS2{F0: v1, F1: ("hello" + "hello"), F2: v50}
            if v2 {
                var v80 uint16 = 0
                if v2 {
                    v77 -= uint(v68)
                    var v81 int16 = -1
                    v50 = -28
                    _ = v81
                }
                if v2 {
                    fmt.Println(v50, v68)
                    for fi82 := 0; fi82 < 15; fi82++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v80 -= (v80 << ((v77) & 63))
                        _ = fi82
                    }
                    switch 9223372036854775807 {
                    case 1:
                        v83 := GenS6{F0: (^(v1)), F1: -1}
                        _ = v83
                    case 9223372036854775807:
                        for fi84 := 0; fi84 < 3; fi84++ {
                            var v85 float32 = (float32(1.5) * 0.0)
                            var v86 string = "hello"
                            v85 = v85
                            _ = fi84
                            _ = v85
                            _ = v86
                        }
                    case 127:
                        switch "foo bar" {
                        case "a":
                            switch 505 {
                            case 0:
                                v50 = (v50 * 0)
                            case 42:
                                fmt.Println(v2, v80, v77)
                            default:
                                fmt.Println(v2, v68, v77)
                            }
                        case "\n\t":
                            var v87 string = "a"
                            _ = v87
                        default:
                            if v2 {
                                if v2 {
                                    switch ("a" + (("" + ("\n\t" + "foo bar")) + "")) {
                                    case "foo bar":
                                        if (v2 || v2) {
                                            v1 = v77
                                            for fi88 := 0; fi88 < 13; fi88++ {
                                                fmt.Println(v1)
                                                _ = fi88
                                            }
                                        }
                                    }
                                    fmt.Println(v1, v1, v77)
                                }
                                v3 = append(v3, 544)
                                v89 := GenS2{F0: v1, F1: ("\n\t" + "hello"), F2: (v50 | v50)}
                                _ = v89
                            }
                        }
                    default:
                        v80 -= v80
                    }
                }
                _ = v80
            } else {
                var v90 int16 = 42
                var v91 int64 = -1
                for fi92 := 0; fi92 < 19; fi92++ {
                    var v93 int = (fi92 * fi92)
                    _ = fi92
                    _ = v93
                }
                _ = v90
                _ = v91
            }
            v94 := GenS3{F0: 0, F1: 65535}
            _ = v79
            _ = v94
        } else {
            fmt.Println(v50, v1, v2)
        }
        for fi95 := 0; fi95 < 15; fi95++ {
            v1 -= (v1 * v77)
            v50 -= (v50 | v50)
            switch fi95 {
            case -1:
                var v96 int64 = -1
                _ = v96
            case 696:
                v97 := GenS7{F0: (^(v68)), F1: v2, F2: (v77 | (v1 & v1))}
                _ = v97
            default:
                for fi98 := 0; fi98 < 10; fi98++ {
                    if ((0.1 > -2.25) && (!(!(false || v2)))) {
                        var v99 uint8 = (v68 % (((v68 ^ (v68 & v68))) | 1))
                        var v100 uint8 = v99
                        fmt.Println(v1, v68)
                        _ = v100
                        _ = v99
                    } else {
                        fmt.Println(v2, fi95)
                        v2 = (!v2)
                        v101 := GenS8{F0: (v50 - (v50 << (((v1 % ((255) | 1))) & 63)))}
                        v102 := GenS4{F0: "hello", F1: -2147483648, F2: true}
                        v77 = genFn1(("" + (("hello" + "hello") + ("hello" + "foo bar"))), (v68 % (((v68 - 127)) | 1)))
                        if (1e10 <= 1e10) {
                            if v2 {
                                v103 := GenS5{F0: (!(892 > 1)), F1: 65535, F2: 1}
                                _ = v103
                            } else {
                                var v104 bool = false
                                _ = v104
                            }
                            switch fi95 {
                            case 42:
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                            default:
                                var v105 uint8 = v68
                                _ = v105
                            }
                            v77 += v77
                        }
                        _ = v101
                        _ = v102
                    }
                    _ = fi98
                }
            }
            _ = fi95
        }
        if (v2 && (0 > -606)) {
            var v106 int16 = 188
            switch 1 {
            case 868:
                var v107 uint16 = 2
                _ = v107
            }
            v108 := GenS3{F0: (v106 + -32768), F1: 1}
            _ = v106
            _ = v108
        } else {
            v1 = v1
            if len(v4) > 0 {
                v4[0] = genFn0()
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v68:", v68)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v68
    _ = v77
}
