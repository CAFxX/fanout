package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32; F1 bool; F2 int64 }
type GenS4 struct { F0 uint8; F1 string; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int32; F2 int32 }
type GenS6 struct { F0 uint; F1 uint; F2 int8 }
type GenS7 struct { F0 int32; F1 int; F2 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint; F1 string; F2 uint8 }
type GenS10 struct { F0 uint64; F1 int16; F2 uint }
type GenS11 struct { F0 int; F1 uint16; F2 uint64 }
type GenS12 struct { F0 int16; F1 uint32 }
type GenS13 struct { F0 string; F1 uint16 }
type GenS14 struct { F0 uint; F1 uint32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int8; F1 int8; F2 int }
func genFn0() bool {
    var v5 uint = 42
    switch -9223372036854775808 {
    case 128:
        v5 = (v5 - v5)
    }
    _ = v5
    return (32768 > 127)
}
func genFn1(p0 int, p1 int, p2 int32) string {
    v6 := []int{}
    var v7 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = v3
    v8 := []int{127, 0}
    for fi9 := 0; fi9 < 17; fi9++ {
        switch "" {
        case "a":
            fmt.Println(v3, v3, fi9)
        case "hello":
            v10 := v1[4294967295]
            _ = v10
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            if len(v8) > 0 {
                v8[0] = fi11
            }
            _ = fi11
        }
        _ = fi9
    }
    switch 32768 {
    case 32767:
        v12 := GenS1{F0: 4294967295, F1: 2, F2: 2147483647}
        _ = v12
    case 532:
        v3 = v3
    case 127:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        fmt.Println(v3, v3)
    }
    fmt.Println(v4, v4)
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 string = ("foo bar" + "\n\t")
            var v16 int = fi14
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    v4 = v4
    for fi17 := 0; fi17 < 18; fi17++ {
        switch "a" {
        case "a":
            v1[325] = "\n\t"
        case "hello":
            var v18 float64 = 1.5
            _ = v18
        case "foo bar":
            if true {
                for fi19 := 0; fi19 < 17; fi19++ {
                    if (!(-2.25 >= (1.5 * 0.1))) {
                        var v20 uint16 = 32767
                        _ = v20
                    } else {
                        if true {
                            var v21 uint32 = 517
                            _ = v21
                        }
                    }
                    v3 += (^(v4))
                    var v22 uint16 = 383
                    v23 := GenS2{F0: v22, F1: -1}
                    _ = fi19
                    _ = v22
                    _ = v23
                }
            } else {
                if (!(0.0 == -2.25)) {
                    v24 := GenS3{F0: 42, F1: (!true), F2: 42}
                    var v25 float64 = 1e10
                    v26 := GenS2{F0: uint16(v3), F1: 118}
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    switch fi17 {
                    case 128:
                        if true {
                            fmt.Println(fi17, v3, v3)
                            v3 = v3
                            v3 -= 18446744073709551615
                        }
                    case 32768:
                        var v27 int32 = -1
                        _ = v27
                    case 1:
                        v28 := GenS2{F0: 32767, F1: 42}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 6; fi29++ {
                        fmt.Println(fi29, fi17)
                        fmt.Println(v4, v4)
                        v30 := GenS4{F0: 127, F1: ("foo bar" + "hello"), F2: 91}
                        if (!true) {
                            fmt.Println(v3)
                            var v31 uint16 = 32767
                            _ = v31
                        } else {
                            v32 := v1[2147483648]
                            if (0.1 <= 0.0) {
                                if (1 >= 365) {
                                    if (((1.5 / 3.14159) / 1.5) <= 0.1) {
                                        v3 = v3
                                        fmt.Println(v32, v3)
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = (-(fi17))
                                    }
                                    fmt.Println(fi17)
                                } else {
                                    fmt.Println(fi17, v4, fi17)
                                    var v33 int64 = 0
                                    v34 := GenS1{F0: fi17, F1: 227, F2: 1}
                                    _ = v33
                                    _ = v34
                                }
                                delete(v2, v32)
                            }
                            _ = v32
                        }
                        var v35 int = (fi17 & fi29)
                        _ = fi29
                        _ = v30
                        _ = v35
                    }
                    if false {
                        var v36 float32 = 0.0
                        switch "a" {
                        case "hello":
                            v3 -= v4
                        case "":
                            v36 = v36
                        }
                        _ = v36
                    }
                }
            }
        }
        if (((!false) && (2 == 32767)) || true) {
            v4 -= (^(v3))
        } else {
            v8 = append(v8, fi17)
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 int8 = 42
                if true {
                    v3 = (v4 & 42)
                    if true {
                        switch "\n\t" {
                        case "\n\t":
                            v39 := GenS5{F0: 322, F1: -277, F2: -1}
                            _ = v39
                        case "foo bar":
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        case "hello":
                            if (fi17 != (fi17 - (fi17 << ((0) & 63)))) {
                                var v40 bool = (295 != 2)
                                if (1 != 65535) {
                                    delete(v2, ("a" + ("" + "foo bar")))
                                    var v41 int = (fi17 & (fi17 / (((fi17 >> ((18446744073709551615) & 63))) | 1)))
                                    _ = v41
                                } else {
                                    var v42 bool = ((0 != 2147483647) && v40)
                                    var v43 int16 = 0
                                    fmt.Println(fi17, v43, fi37)
                                    if v40 {
                                        v38 = v38
                                    } else {
                                        if len(v8) > 0 {
                                            v8[0] = fi37
                                        }
                                        var v44 uint32 = 0
                                        if (0 == (v43 >> ((413) & 63))) {
                                            if len(v8) > 0 {
                                                v8[0] = (^(fi17))
                                            }
                                            var v45 bool = true
                                            v46 := GenS2{F0: 0, F1: v38}
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v44 -= (v44 % ((v44) | 1))
                                        }
                                        _ = v44
                                    }
                                    switch (fi17 - (fi17 / (((fi37 >> ((0) & 63))) | 1))) {
                                    case -9223372036854775808:
                                        v8 = append(v8, int(v4))
                                    }
                                    _ = v42
                                    _ = v43
                                }
                                v47 := GenS3{F0: 436, F1: (352 >= 706), F2: -534}
                                if (1 < -9223372036854775808) {
                                    if v40 {
                                        v3 = v3
                                        if ((1 < 368) || v40) {
                                            if v40 {
                                                v3 -= v4
                                                var v48 uint32 = 2147483648
                                                v49 := GenS2{F0: 65, F1: v38}
                                                v50 := GenS3{F0: 191, F1: (-1 != 1), F2: -54}
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            }
                                            switch "a" {
                                            case "":
                                                if (!false) {
                                                    v51 := GenS6{F0: 212, F1: 65535, F2: v38}
                                                    var v52 bool = (0 <= 1000)
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    v53 := GenS7{F0: 2147483647, F1: 32767, F2: (-953 != 9223372036854775807)}
                                                    fmt.Println(v3)
                                                    _ = v53
                                                }
                                            case "hello":
                                                var v54 float32 = 1e10
                                                _ = v54
                                            }
                                            if (2 != 42) {
                                                if (2 < 65535) {
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if v40 {
                                                            v3 = 4294967295
                                                            v38 = (-(v38))
                                                            var v55 int16 = 32767
                                                            _ = v55
                                                        }
                                                    case "foo bar":
                                                        v4 += (v4 | (v4 ^ 0))
                                                    default:
                                                        v3 -= v3
                                                    }
                                                    v40 = (!v40)
                                                } else {
                                                    _ = len(v1)
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    if false {
                                                        switch "" {
                                                        case "":
                                                            v56 := GenS6{F0: 18446744073709551615, F1: 848, F2: 1}
                                                            _ = v56
                                                        }
                                                    }
                                                }
                                                switch fi17 {
                                                case -2147483648:
                                                    if v40 {
                                                        var v57 int16 = 0
                                                        v8 = append(v8, fi37)
                                                        v40 = (1 != -517)
                                                        _ = v57
                                                    } else {
                                                        fmt.Println(fi37)
                                                        v3 += v4
                                                        v38 = v38
                                                    }
                                                case 537:
                                                    if len(v8) > 0 {
                                                        v8[0] = (fi17 << ((4294967295) & 63))
                                                    }
                                                case -1:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if v40 {
                                                            v58 := v1[47]
                                                            fmt.Println(v38)
                                                            if false {
                                                                fmt.Println(v58, v58, fi17)
                                                            } else {
                                                                v58 = v58
                                                                v3 = 42
                                                                v59 := GenS3{F0: -2147483648, F1: v40, F2: int64(v3)}
                                                                v3 -= 4294967295
                                                                v4 = v4
                                                                _ = v59
                                                            }
                                                            _ = v58
                                                        }
                                                    case "hello":
                                                        if (!("a" < "a")) {
                                                            v60 := GenS8{F0: (v3 ^ 4294967295)}
                                                            var v61 int64 = 0
                                                            v62 := GenS9{F0: 0, F1: "foo bar", F2: 1}
                                                            v63 := GenS9{F0: 996, F1: "hello", F2: 127}
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    case "foo bar":
                                                        v64 := GenS7{F0: -2147483648, F1: (fi37 & fi37), F2: ("hello" > "")}
                                                        _ = v64
                                                    default:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if (!(v40 && v40)) {
                                                                v3 += v3
                                                                v65 := GenS10{F0: v3, F1: -907, F2: 65535}
                                                                var v66 int32 = -1
                                                                fmt.Println(v40)
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                        case "":
                                                            v4 = 2
                                                        case "a":
                                                            var v67 float32 = 1.5
                                                            _ = v67
                                                        }
                                                    }
                                                }
                                                v40 = v40
                                            }
                                        } else {
                                            v40 = (v40 || v40)
                                            if v40 {
                                                var v68 int64 = -924
                                                v38 = (v38 + 0)
                                                v69 := v2["hello"]
                                                _ = v68
                                                _ = v69
                                            } else {
                                                fmt.Println(fi17)
                                                if (v4 > v4) {
                                                    var v70 uint64 = (v3 | 18446744073709551615)
                                                    v71 := GenS3{F0: -2147483648, F1: v40, F2: -1}
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    if (!((-(fi37)) >= 128)) {
                                                        v72 := GenS1{F0: (-(fi17)), F1: 802, F2: 1}
                                                        _ = v72
                                                    }
                                                }
                                            }
                                            v8 = append(v8, (fi37 + 9223372036854775807))
                                        }
                                        v4 = v4
                                    } else {
                                        v40 = (fi17 < 128)
                                        var v73 int32 = -1
                                        v74 := GenS11{F0: fi37, F1: 42, F2: 31}
                                        var v75 uint8 = 50
                                        switch "a" {
                                        case "a":
                                            v76 := GenS3{F0: (v73 / (((v73 + v73)) | 1)), F1: true, F2: 9223372036854775807}
                                            _ = v76
                                        case "\n\t":
                                            if len(v8) > 0 {
                                                v8[0] = (fi37 - (fi17 + 2147483647))
                                            }
                                        default:
                                            v77 := GenS12{F0: 42, F1: 810}
                                            _ = v77
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            var v78 bool = (true || v40)
                                            _ = v78
                                        case "\n\t":
                                            v79 := GenS13{F0: "\n\t", F1: 810}
                                            _ = v79
                                        }
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    }
                                    if (2147483647 < -2147483648) {
                                        var v80 int8 = 100
                                        var v81 uint32 = 2147483647
                                        _ = v80
                                        _ = v81
                                    }
                                    fmt.Println(v38, fi37, fi17)
                                } else {
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v40, v4, fi37)
                                        var v82 int16 = 32767
                                        v83 := GenS14{F0: 1, F1: 4294967295}
                                        _ = v82
                                        _ = v83
                                    default:
                                        var v84 uint8 = 1
                                        _ = v84
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = (fi17 + 2147483647)
                                    }
                                    v4 += (v3 & 2)
                                    if true {
                                        var v85 float32 = 3.14159
                                        var v86 string = ("" + "hello")
                                        fmt.Println(fi37, fi17)
                                        _ = v85
                                        _ = v86
                                    }
                                    v87 := v1[481]
                                    v87 = v87
                                    _ = v87
                                }
                                delete(v2, "a")
                                _ = v40
                                _ = v47
                            } else {
                                v88 := GenS12{F0: -1, F1: 1}
                                _ = v88
                            }
                        }
                        v3 += (v4 | 42)
                    } else {
                        switch (fi37 >> ((255) & 63)) {
                        case 654:
                            switch "" {
                            case "\n\t":
                                var v89 int32 = -696
                                _ = v89
                            case "a":
                                v90 := GenS14{F0: 2, F1: 2}
                                _ = v90
                            case "foo bar":
                                var v91 uint16 = 32767
                                _ = v91
                            }
                        }
                        var v92 uint = 2
                        v92 = v92
                        v38 = (v38 + (v38 >> (((^(v92))) & 63)))
                        var v93 int = (fi37 / ((-9223372036854775808) | 1))
                        _ = v92
                        _ = v93
                    }
                }
                switch fi37 {
                case 9223372036854775807:
                    _ = len(v8)
                    var v94 float32 = 3.14159
                    _ = v94
                case 2147483647:
                    var v95 float32 = (1e10 / 1.0)
                    _ = v95
                default:
                    var v96 uint32 = 65535
                    _ = v96
                }
                _ = fi37
                _ = v38
            }
            fmt.Println(fi17, v3)
        }
        _ = fi17
    }
    var v97 int64 = 0
    v4 -= (v3 ^ v4)
    v8 = append(v8, -9223372036854775808)
    v98 := GenS5{F0: 2, F1: 0, F2: 1}
    if true {
        if (!false) {
            for fi99 := 0; fi99 < 15; fi99++ {
                for fi100 := 0; fi100 < 5; fi100++ {
                    v101 := GenS14{F0: 0, F1: 42}
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi100
                    _ = v101
                }
                switch ("a" + "foo bar") {
                case "":
                    v4 = (v3 & (^(v4)))
                case "hello":
                    v102 := GenS15{F0: v4}
                    _ = v102
                case "\n\t":
                    var v103 uint16 = 2
                    _ = v103
                default:
                    var v104 int8 = 94
                    _ = v104
                }
                _ = fi99
            }
        } else {
            if ((!(32767 <= 42)) && false) {
                var v105 int64 = 0
                v97 = (v105 | (v97 / ((v105) | 1)))
                _ = v105
            }
            v106 := GenS4{F0: 1, F1: "a", F2: 2147483647}
            v3 = v4
            _ = v106
        }
    }
    v4 += (v4 / (((v4 ^ 18446744073709551615)) | 1))
    switch ((("" + ("foo bar" + ("foo bar" + "\n\t"))) + "\n\t") + ("foo bar" + "")) {
    case "a":
        v4 += v4
    case "\n\t":
        delete(v1, 42)
    case "foo bar":
        v3 = v4
    }
    v107 := GenS16{F0: -128, F1: 127, F2: 128}
    var v108 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v97
    _ = v98
}
