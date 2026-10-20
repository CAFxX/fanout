package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int16; F1 uint; F2 uint64 }
type GenS3 struct { F0 int8; F1 int32; F2 string }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int8; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 string; F1 int; F2 uint64 }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 uint16; F1 int64 }
type GenS12 struct { F0 int16; F1 int32 }
func genFn0(p0 string) int64 {
    v5 := GenS1{F0: 1, F1: 0}
    var v6 uint = 18446744073709551615
    var v7 int64 = 1
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint16 = 255
            v10 := GenS1{F0: -1, F1: (v7 & v7)}
            _ = fi8
            _ = v10
            _ = v9
        }
        switch p0 {
        case "foo bar":
            var v11 int64 = -9223372036854775808
            _ = v11
        case "hello":
            switch -2147483648 {
            case 492:
                v6 = (^(v6))
            }
        case "":
            for fi12 := 0; fi12 < 11; fi12++ {
                switch (fi12 >> (((v6 / (((^(v6))) | 1))) & 63)) {
                case 9223372036854775807:
                    v7 += v7
                case 4294967295:
                    p0 = ((((p0 + ("" + (p0 + p0))) + p0) + ((("foo bar" + "foo bar") + p0) + "\n\t")) + "foo bar")
                }
                _ = fi12
            }
        }
    } else {
        p0 = "\n\t"
        if (!(false || true)) {
            v13 := []int{128, 0, 9223372036854775807}
            _ = v13
        } else {
            v14 := []int{1, 42, -9223372036854775808, 1}
            if (0.1 <= 0.1) {
                fmt.Println(v6, p0, v6)
                switch "" {
                case "":
                    switch ((p0 + ((p0 + (p0 + "\n\t")) + "foo bar")) + p0) {
                    case "foo bar":
                        v14 = append(v14, 0)
                    case "":
                        for fi15 := 0; fi15 < 1; fi15++ {
                            if false {
                                fmt.Println(p0, fi15)
                                for fi16 := 0; fi16 < 4; fi16++ {
                                    var v17 uint64 = 18446744073709551615
                                    p0 = p0
                                    v18 := GenS2{F0: -442, F1: 4294967295, F2: v17}
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                                p0 = ((p0 + ("" + "hello")) + ("hello" + ""))
                            }
                            _ = fi15
                        }
                    case "a":
                        for fi19 := 0; fi19 < 13; fi19++ {
                            if (int8(v7) <= 1) {
                                var v20 uint16 = 2
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 6; fi21++ {
                                if (844 == 42) {
                                    if false {
                                        var v22 string = p0
                                        if (519 > 376) {
                                            fmt.Println(v22, v22, v22)
                                            switch (v22 + p0) {
                                            case "\n\t":
                                                if true {
                                                    v23 := GenS1{F0: -2147483648, F1: (v7 >> ((65535) & 63))}
                                                    if (v6 == v6) {
                                                        var v24 int = (fi21 >> ((v6) & 63))
                                                        v7 += (v7 * (v7 - (v7 >> ((v6) & 63))))
                                                        v25 := GenS3{F0: 42, F1: 2147483647, F2: v22}
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                    _ = v23
                                                } else {
                                                    fmt.Println(fi21)
                                                    var v26 string = "hello"
                                                    fmt.Println(v22, v26, v26)
                                                    _ = v26
                                                }
                                            case "":
                                                if (0.1 >= 1e10) {
                                                    v14 = append(v14, fi21)
                                                } else {
                                                    _ = len(v14)
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v27 uint32 = 65535
                                                        _ = v27
                                                    case "a":
                                                        v28 := GenS4{F0: 255, F1: 149}
                                                        _ = v28
                                                    case "":
                                                        var v29 uint64 = 18446744073709551615
                                                        _ = v29
                                                    default:
                                                        v14 = append(v14, 4294967295)
                                                    }
                                                }
                                            }
                                        }
                                        _ = v22
                                    }
                                } else {
                                    p0 = p0
                                    if len(v14) > 0 {
                                        v14[0] = fi19
                                    }
                                    var v30 string = p0
                                    if len(v14) > 0 {
                                        v14[0] = fi19
                                    }
                                    var v31 uint = (v6 / (((v6 - v6)) | 1))
                                    switch p0 {
                                    case "\n\t":
                                        if ((v7 << (((^(v31))) & 63)) == v7) {
                                            v6 = 668
                                            v32 := GenS2{F0: 43, F1: (v6 ^ 42), F2: 0}
                                            var v33 uint64 = 4294967295
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                    var v34 string = p0
                                    _ = v30
                                    _ = v31
                                    _ = v34
                                }
                                p0 = (p0 + "")
                                v6 = v6
                                v14 = append(v14, 32767)
                                var v35 int8 = 42
                                _ = fi21
                                _ = v35
                            }
                            var v36 int = 9223372036854775807
                            _ = fi19
                            _ = v36
                        }
                    default:
                        var v37 uint64 = 18446744073709551615
                        _ = v37
                    }
                case "a":
                    var v38 uint16 = 65535
                    _ = v38
                case "foo bar":
                    v7 = 0
                }
                var v39 int8 = -91
                _ = v39
            }
            if false {
                v40 := GenS2{F0: 32767, F1: v6, F2: 0}
                for fi41 := 0; fi41 < 12; fi41++ {
                    if true {
                        p0 = ("a" + p0)
                        var v42 float64 = 0.1
                        var v43 int8 = 100
                        _ = v42
                        _ = v43
                    } else {
                        v44 := GenS5{F0: v6, F1: 1}
                        var v45 int16 = -311
                        _ = v44
                        _ = v45
                    }
                    _ = fi41
                }
                _ = v40
            }
            _ = v14
        }
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 uint64) uint64 {
    v46 := GenS6{F0: 42, F1: -9223372036854775808, F2: 65535}
    var v47 uint64 = (p0 - 4294967295)
    _ = p0
    _ = v46
    _ = v47
    return v47
}

func main() {
    var v1 float64 = (0.0 + 3.14159)
    var v2 int64 = -9223372036854775808
    var v3 int64 = v2
    var v4 int = -2147483648
    for fi48 := 0; fi48 < 15; fi48++ {
        if (((false && true) && false) && false) {
            var v49 bool = ((1 >= 42) && true)
            for fi50 := 0; fi50 < 11; fi50++ {
                v2 += (v2 + (v3 << ((65535) & 63)))
                _ = fi50
            }
            _ = v49
        } else {
            v51 := []int{}
            switch -2147483648 {
            case 32768:
                switch fi48 {
                case 128:
                    v2 = (v2 ^ v2)
                case 32767:
                    var v52 uint = 42
                    _ = v52
                default:
                    v51 = append(v51, (fi48 % (((v4 >> ((65535) & 63))) | 1)))
                }
            default:
                fmt.Println(v3, v1)
            }
            _ = v51
        }
        for fi53 := 0; fi53 < 19; fi53++ {
            var v54 int = fi48
            v55 := GenS7{F0: 0}
            _ = fi53
            _ = v54
            _ = v55
        }
        _ = fi48
    }
    if (true || false) {
        for fi56 := 0; fi56 < 6; fi56++ {
            var v57 float32 = -2.25
            _ = fi56
            _ = v57
        }
    }
    v58 := []int{v4}
    for fi59 := 0; fi59 < 20; fi59++ {
        var v60 int8 = 100
        for fi61 := 0; fi61 < 13; fi61++ {
            v60 = (v60 / (((v60 << ((65535) & 63))) | 1))
            switch fi59 {
            case 320:
                switch "hello" {
                case "":
                    v4 = 557
                case "a":
                    v62 := GenS1{F0: -637, F1: (v3 + genFn0("hello"))}
                    _ = v62
                default:
                    var v63 uint32 = 2147483648
                    _ = v63
                }
                v2 += (-(v3))
                if false {
                    v58 = append(v58, v4)
                    var v64 int8 = (v60 - (v60 / ((v60) | 1)))
                    v65 := GenS8{F0: 4294967295, F1: 61}
                    _ = v64
                    _ = v65
                }
                v58 = append(v58, -58)
            }
            var v66 int32 = -1
            _ = fi61
            _ = v66
        }
        _ = fi59
        _ = v60
    }
    var v67 int8 = 1
    var v68 float64 = 3.14159
    var v69 uint32 = 713
    for fi70 := 0; fi70 < 6; fi70++ {
        if false {
            v4 -= (v4 & 0)
        } else {
            switch 127 {
            case 0:
                v71 := GenS6{F0: v67, F1: genFn0("foo bar"), F2: 964}
                _ = v71
            case 32767:
                fmt.Println(v2, v4)
            case 2147483647:
                var v72 int = fi70
                _ = v72
            }
            for fi73 := 0; fi73 < 9; fi73++ {
                _ = len(v58)
                var v74 uint = 1
                _ = fi73
                _ = v74
            }
            var v75 uint8 = 250
            _ = v75
        }
        v69 += v69
        _ = fi70
    }
    if (42 > -889) {
        for fi76 := 0; fi76 < 14; fi76++ {
            if ((v69 + v69) >= v69) {
                var v77 int16 = -32768
                var v78 int = fi76
                for fi79 := 0; fi79 < 10; fi79++ {
                    if false {
                        var v80 uint16 = 1
                        _ = v80
                    } else {
                        fmt.Println(v3, v2)
                    }
                    _ = fi79
                }
                _ = v77
                _ = v78
            } else {
                v81 := GenS7{F0: 253}
                if (!(1 != 18446744073709551615)) {
                    _ = len(v58)
                } else {
                    v68 -= v1
                    var v82 float64 = 3.14159
                    _ = v82
                }
                _ = v81
            }
            _ = fi76
        }
        v83 := GenS9{F0: (("foo bar" + ("\n\t" + "a")) + (("a" + ("" + ("foo bar" + "hello"))) + ("" + "foo bar"))), F1: (v4 | v4), F2: 469}
        fmt.Println(v69)
        _ = v83
    }
    var v84 uint32 = (^(v69))
    if (267 >= v69) {
        for fi85 := 0; fi85 < 7; fi85++ {
            v58 = append(v58, 0)
            var v86 int16 = -273
            _ = fi85
            _ = v86
        }
    } else {
        v84 -= v69
        switch ("\n\t" + "\n\t") {
        case "hello":
            switch (v4 | (v4 << ((928) & 63))) {
            case -1:
                switch 2147483647 {
                case 647:
                    switch (v4 - (v4 << ((255) & 63))) {
                    case 157:
                        for fi87 := 0; fi87 < 11; fi87++ {
                            v88 := GenS10{F0: true, F1: v69}
                            if false {
                                v89 := GenS3{F0: v67, F1: -2147483648, F2: ""}
                                v69 -= (v69 * v69)
                                fmt.Println(fi87, v2)
                                _ = v89
                            }
                            for fi90 := 0; fi90 < 20; fi90++ {
                                var v91 uint64 = 18446744073709551615
                                _ = fi90
                                _ = v91
                            }
                            _ = fi87
                            _ = v88
                        }
                    case 42:
                        switch "foo bar" {
                        case "hello":
                            if false {
                                if len(v58) >= 2 {
                                    v58 = v58[1:len(v58)]
                                }
                                if len(v58) > 0 {
                                    v58[0] = v4
                                }
                            }
                        case "foo bar":
                            v92 := GenS11{F0: 2, F1: (v2 - (v2 | v2))}
                            _ = v92
                        }
                    case 857:
                        if true {
                            var v93 string = "hello"
                            v93 = v93
                            if len(v58) > 0 {
                                v58[0] = v4
                            }
                            _ = v93
                        } else {
                            v94 := GenS10{F0: false, F1: (^(v84))}
                            for fi95 := 0; fi95 < 6; fi95++ {
                                var v96 uint8 = 2
                                var v97 string = (("\n\t" + (("foo bar" + (((("a" + "\n\t") + "") + "hello") + "\n\t")) + "foo bar")) + ("hello" + "a"))
                                for fi98 := 0; fi98 < 7; fi98++ {
                                    v96 -= v96
                                    _ = fi98
                                }
                                _ = fi95
                                _ = v96
                                _ = v97
                            }
                            _ = v94
                        }
                    default:
                        if (931 == 0) {
                            fmt.Println(v69, v67)
                            switch "a" {
                            case "hello":
                                for fi99 := 0; fi99 < 4; fi99++ {
                                    switch -2147483648 {
                                    case 32767:
                                        var v100 int32 = 2147483647
                                        _ = v100
                                    case 692:
                                        for fi101 := 0; fi101 < 6; fi101++ {
                                            v58 = append(v58, (fi101 + (^(v4))))
                                            _ = fi101
                                        }
                                    }
                                    _ = fi99
                                }
                            case "a":
                                for fi102 := 0; fi102 < 7; fi102++ {
                                    if (105 == 2) {
                                        if len(v58) > 0 {
                                            v58[0] = (fi102 >> ((437) & 63))
                                        }
                                    } else {
                                        for fi103 := 0; fi103 < 6; fi103++ {
                                            v58 = append(v58, (fi102 << ((65535) & 63)))
                                            if ((2 < 18446744073709551615) && true) {
                                                fmt.Println(v68, fi103, fi102)
                                            }
                                            _ = fi103
                                        }
                                        v69 += 65535
                                        for fi104 := 0; fi104 < 15; fi104++ {
                                            var v105 float64 = v68
                                            if false {
                                                if len(v58) > 0 {
                                                    v58[0] = int(v2)
                                                }
                                                var v106 uint32 = (v69 + (v69 | (^(v84))))
                                                _ = v106
                                            } else {
                                                v84 = (v69 & 65535)
                                            }
                                            _ = fi104
                                            _ = v105
                                        }
                                    }
                                    v84 -= v69
                                    for fi107 := 0; fi107 < 13; fi107++ {
                                        var v108 int8 = (v67 / (((v67 % (((v67 | v67)) | 1))) | 1))
                                        v58 = append(v58, fi102)
                                        _ = fi107
                                        _ = v108
                                    }
                                    _ = fi102
                                }
                            case "foo bar":
                                v4 -= v4
                            }
                        } else {
                            for fi109 := 0; fi109 < 20; fi109++ {
                                v2 -= (v3 << ((18446744073709551615) & 63))
                                fmt.Println(v3, v69, v1)
                                _ = fi109
                            }
                            v110 := GenS12{F0: -1, F1: 0}
                            if len(v58) >= 2 {
                                v58 = v58[1:len(v58)]
                            }
                            _ = v110
                        }
                    }
                case 178:
                    fmt.Println(v4, v84, v84)
                case 127:
                    v111 := GenS11{F0: 223, F1: v3}
                    _ = v111
                }
            default:
                var v112 int = (v4 * (v4 | (^(v4))))
                _ = v112
            }
        default:
            switch 128 {
            case 9223372036854775807:
                fmt.Println(v2, v3)
            case -1:
                var v113 int64 = 9223372036854775807
                _ = v113
            }
        }
    }
    v58 = append(v58, (v4 + v4))
    switch v4 {
    case 1:
        var v114 int16 = 1
        _ = v114
    case 32767:
        var v115 uint32 = 73
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v67
    _ = v68
    _ = v69
    _ = v84
}
