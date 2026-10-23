package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 string }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 uint; F1 int; F2 int8 }
type GenS10 struct { F0 int64; F1 string; F2 int16 }
func genFn0(p0 uint64, p1 int) int64 {
    fmt.Println(p1, p0)
    p0 = 2
    v4 := []int{(p1 | (^(p1))), p1}
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 = p0
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 692
}

func main() {
    var v1 uint8 = 42
    var v2 int8 = 16
    v3 := make(map[string]string)
    if (-2 == 32767) {
        if true {
            v6 := []int{128, 1, 128, 32768}
            for fi7 := 0; fi7 < 13; fi7++ {
                fmt.Println(fi7, v1)
                v6 = append(v6, 1)
                _ = fi7
            }
            _ = v6
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            switch "hello" {
            case "a":
                switch -593 {
                case 0:
                    var v9 int16 = 743
                    _ = v9
                case 42:
                    v10 := GenS1{F0: 536, F1: "\n\t", F2: v2}
                    _ = v10
                default:
                    v2 = v2
                }
            case "hello":
                v11 := v3["hello"]
                _ = v11
            case "\n\t":
                v2 = v2
            default:
                for fi12 := 0; fi12 < 13; fi12++ {
                    _ = len(v3)
                    _ = fi12
                }
            }
            if true {
                v13 := GenS1{F0: 2147483647, F1: "foo bar", F2: -128}
                if true {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        var v15 int64 = genFn0(4294967295, -234)
                        v16 := GenS2{F0: false, F1: -797, F2: 660}
                        v15 = (-(v15))
                        v1 -= (v1 % ((255) | 1))
                        var v17 uint16 = 42
                        var v18 uint = 42
                        v19 := GenS1{F0: 2, F1: "a", F2: 100}
                        if (2147483647 != 312) {
                            switch ("\n\t" + "\n\t") {
                            case "a":
                                var v20 uint64 = 42
                                fmt.Println(v2, fi14, v17)
                                _ = v20
                            case "hello":
                                var v21 int64 = (v15 >> ((uint(v1)) & 63))
                                _ = v21
                            case "":
                                v22 := GenS2{F0: true, F1: 679, F2: 2}
                                _ = v22
                            default:
                                if (!(((fi8 != fi8) && false) || (1e10 > (3.14159 / float64(float32((3.14159 - 1.5))))))) {
                                    v15 = (v15 << (((v18 << ((v18) & 63))) & 63))
                                } else {
                                    v1 -= (v1 >> (((v18 * v18)) & 63))
                                    v2 = v2
                                }
                            }
                            var v23 int16 = -537
                            switch "" {
                            case "a":
                                v3["\n\t"] = (("" + "hello") + "foo bar")
                            }
                            _ = v23
                        }
                        v24 := GenS3{F0: genFn0(1, fi14), F1: 895, F2: "hello"}
                        fmt.Println(fi8)
                        v1 = uint8(v18)
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v24
                    }
                    if true {
                        v2 -= (^(v2))
                    } else {
                        v25 := []int{(fi8 % ((fi8) | 1)), (fi8 >> ((511) & 63))}
                        fmt.Println(fi8)
                        _ = v25
                    }
                    var v26 uint32 = 2147483648
                    _ = v26
                }
                _ = v13
            }
            fmt.Println(fi8, v1, fi8)
            _ = fi8
        }
    }
    if true {
        v27 := GenS1{F0: 636, F1: ("" + ""), F2: v2}
        switch 42 {
        case 32768:
            v28 := GenS4{F0: 255, F1: true}
            _ = v28
        }
        var v29 int8 = v2
        _ = v27
        _ = v29
    } else {
        if (false && false) {
            if false {
                if false {
                    _ = len(v3)
                } else {
                    switch ("hello" + "a") {
                    case "":
                        var v30 int16 = 0
                        _ = v30
                    case "foo bar":
                        for fi31 := 0; fi31 < 10; fi31++ {
                            v2 = -128
                            switch "\n\t" {
                            case "hello":
                                var v32 int16 = 1
                                _ = v32
                            }
                            fmt.Println(v2, v2, v1)
                            _ = fi31
                        }
                    }
                }
                for fi33 := 0; fi33 < 1; fi33++ {
                    if (genFn0(18446744073709551615, 127) <= genFn0(747, fi33)) {
                        var v34 int16 = 32767
                        fmt.Println(v34, v1)
                        for fi35 := 0; fi35 < 15; fi35++ {
                            v1 -= v1
                            v2 = v2
                            switch "a" {
                            case "foo bar":
                                v1 -= 0
                            case "":
                                v36 := GenS5{F0: ("" + ("" + "hello"))}
                                _ = v36
                            default:
                                var v37 int16 = v34
                                _ = v37
                            }
                            _ = fi35
                        }
                        _ = v34
                    }
                    v2 += v2
                    if ((false && true) || (255 <= 0)) {
                        if (true && true) {
                            var v38 uint = 337
                            _ = v38
                        } else {
                            for fi39 := 0; fi39 < 11; fi39++ {
                                v40 := GenS3{F0: genFn0(486, (fi39 * fi39)), F1: 1, F2: (("hello" + ("hello" + (("hello" + "\n\t") + ("a" + "a")))) + "a")}
                                _ = fi39
                                _ = v40
                            }
                            var v41 float32 = 1.5
                            fmt.Println(fi33, fi33)
                            _ = v41
                        }
                        switch 9223372036854775807 {
                        case 128:
                            var v42 uint8 = v1
                            _ = v42
                        case 127:
                            if (true || ("a" > "hello")) {
                                v2 = -17
                                for fi43 := 0; fi43 < 11; fi43++ {
                                    v44 := []int{-1000}
                                    v45 := GenS6{F0: -576, F1: 255, F2: 42}
                                    v2 = (v2 >> ((65535) & 63))
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                }
                            } else {
                                var v46 uint = 745
                                _ = v46
                            }
                        default:
                            v1 -= (v1 + (^(v1)))
                        }
                        var v47 int8 = v2
                        _ = v47
                    }
                    _ = fi33
                }
            }
            v48 := []int{int(v2), 2147483647, 4294967295, -170}
            if (0 >= 0) {
                if false {
                    v49 := GenS2{F0: true, F1: 1, F2: 42}
                    switch 2147483647 {
                    case -176:
                        var v50 float64 = float64(0.0)
                        _ = v50
                    case 127:
                        fmt.Println(v2)
                    case 4294967295:
                        for fi51 := 0; fi51 < 12; fi51++ {
                            fmt.Println(v2, fi51, fi51)
                            switch "a" {
                            case "hello":
                                fmt.Println(fi51, v1)
                            case "":
                                v3["foo bar"] = ""
                            }
                            v2 = (v2 | (-(v2)))
                            _ = fi51
                        }
                    default:
                        for fi52 := 0; fi52 < 10; fi52++ {
                            if false {
                                for fi53 := 0; fi53 < 3; fi53++ {
                                    if true {
                                        if false {
                                            var v54 uint = 660
                                            if len(v48) > 0 {
                                                v48[0] = fi53
                                            }
                                            v55 := GenS1{F0: 2147483647, F1: "\n\t", F2: (v2 + (v2 << ((4294967295) & 63)))}
                                            _ = v54
                                            _ = v55
                                        } else {
                                            var v56 bool = false
                                            _ = v56
                                        }
                                        var v57 uint64 = 242
                                        _ = v57
                                    } else {
                                        fmt.Println(v2)
                                        v58 := GenS3{F0: -264, F1: 0, F2: "\n\t"}
                                        fmt.Println(v1, v1)
                                        switch fi53 {
                                        case 9223372036854775807:
                                            if false {
                                                if len(v48) >= 2 {
                                                    v48 = v48[1:len(v48)]
                                                }
                                                v59 := GenS1{F0: 907, F1: "\n\t", F2: v2}
                                                switch 796 {
                                                case 2147483647:
                                                    var v60 int = (^(fi53))
                                                    _ = v60
                                                default:
                                                    var v61 int32 = -67
                                                    _ = v61
                                                }
                                                _ = v59
                                            } else {
                                                var v62 uint = 1
                                                _ = v62
                                            }
                                        case -1:
                                            if (42 <= 505) {
                                                _ = len(v48)
                                            }
                                        }
                                        v1 = v1
                                        _ = v58
                                    }
                                    var v63 uint16 = 0
                                    _ = fi53
                                    _ = v63
                                }
                                var v64 int32 = 0
                                _ = v64
                            } else {
                                for fi65 := 0; fi65 < 18; fi65++ {
                                    v2 -= (^(v2))
                                    switch fi65 {
                                    case 319:
                                        if false {
                                            v48 = append(v48, fi65)
                                        }
                                    case -9223372036854775808:
                                        _ = len(v3)
                                    case 1:
                                        delete(v3, "\n\t")
                                    default:
                                        if len(v48) >= 2 {
                                            v48 = v48[1:len(v48)]
                                        }
                                    }
                                    _ = fi65
                                }
                                v66 := GenS4{F0: 4294967295, F1: (-687 > (fi52 / (((fi52 ^ 179)) | 1)))}
                                _ = v66
                            }
                            _ = fi52
                        }
                    }
                    var v67 int16 = 32767
                    _ = v49
                    _ = v67
                }
                var v68 int16 = -432
                var v69 uint32 = 1
                _ = v68
                _ = v69
            } else {
                if (32767 < -1) {
                    v70 := GenS1{F0: 4294967295, F1: (("" + ("hello" + "foo bar")) + "\n\t"), F2: 30}
                    for fi71 := 0; fi71 < 5; fi71++ {
                        fmt.Println(fi71, v2, fi71)
                        v1 = (v1 ^ v1)
                        _ = fi71
                    }
                    _ = v70
                } else {
                    if false {
                        var v72 float32 = 0.1
                        v1 += v1
                        _ = v72
                    } else {
                        v73 := GenS7{F0: 1}
                        _ = v73
                    }
                    v74 := GenS8{F0: 256, F1: 2147483647}
                    v75 := v3[("foo bar" + ("a" + "a"))]
                    _ = v74
                    _ = v75
                }
                switch 128 {
                case -173:
                    var v76 float32 = 0.1
                    _ = v76
                case -1:
                    var v77 int8 = 127
                    _ = v77
                case 4294967295:
                    v78 := GenS9{F0: 18446744073709551615, F1: 32768, F2: 1}
                    _ = v78
                }
                v79 := v3[(("\n\t" + ("a" + "")) + ((("hello" + ("" + "\n\t")) + ("\n\t" + "foo bar")) + ""))]
                _ = v79
            }
            _ = v48
        } else {
            switch -120 {
            case 9223372036854775807:
                var v80 int32 = -949
                _ = v80
            default:
                if true {
                    switch ("" + ("foo bar" + "a")) {
                    case "":
                        v2 = v2
                    case "foo bar":
                        for fi81 := 0; fi81 < 19; fi81++ {
                            var v82 uint8 = v1
                            _ = fi81
                            _ = v82
                        }
                    case "hello":
                        if (false && (!(255 >= 32767))) {
                            fmt.Println(v1, v1)
                            v83 := GenS10{F0: 42, F1: ("a" + ("\n\t" + "")), F2: 0}
                            _ = v83
                        }
                    default:
                        v2 = 1
                    }
                }
            }
            fmt.Println(v1, v1)
            var v84 uint32 = 1
            _ = v84
        }
        if (2 < 4294967295) {
            v1 -= v1
            v1 = (v1 - 238)
            v1 += 106
        }
        var v85 int = 4294967295
        _ = v85
    }
    v86 := GenS1{F0: 4294967295, F1: "hello", F2: (v2 * (v2 + v2))}
    v2 = v2
    v2 -= v2
    var v87 uint32 = 42
    var v88 uint64 = 4294967295
    switch "a" {
    case "hello":
        if false {
            if true {
                if (-1 <= -1) {
                    var v89 uint16 = 780
                    _ = v89
                } else {
                    switch "hello" {
                    case "hello":
                        var v90 int32 = 1
                        _ = v90
                    }
                    var v91 uint64 = uint64(v1)
                    _ = v91
                }
                var v92 int32 = 1
                _ = v92
            } else {
                v1 -= v1
                if ("" > (((("" + ("foo bar" + "foo bar")) + ((("foo bar" + ("\n\t" + ("a" + "a"))) + "\n\t") + "a")) + "a") + "\n\t")) {
                    if true {
                        if (true && (false && ((true || false) || true))) {
                            switch 128 {
                            case 1:
                                v3["hello"] = ""
                            case 32768:
                                for fi93 := 0; fi93 < 19; fi93++ {
                                    var v94 int32 = 753
                                    _ = fi93
                                    _ = v94
                                }
                            }
                            v1 = 2
                            var v95 int16 = -1
                            _ = v95
                        } else {
                            switch ((("" + "\n\t") + ((("a" + "foo bar") + "\n\t") + "a")) + ("a" + ("a" + ("hello" + ("\n\t" + "hello"))))) {
                            case "\n\t":
                                v96 := []int{189, 32767}
                                _ = v96
                            }
                        }
                        v87 = (v87 | (v87 - v87))
                    } else {
                        if true {
                            var v97 float32 = -2.25
                            for fi98 := 0; fi98 < 8; fi98++ {
                                v99 := GenS6{F0: 1, F1: (^(v1)), F2: 15}
                                _ = fi98
                                _ = v99
                            }
                            _ = v97
                        }
                        v100 := GenS7{F0: v88}
                        v3["a"] = ("\n\t" + ("\n\t" + ""))
                        _ = v100
                    }
                    switch 32767 {
                    case -304:
                        delete(v3, ("\n\t" + "\n\t"))
                    }
                    var v101 uint32 = v87
                    _ = v101
                } else {
                    var v102 uint64 = 0
                    v103 := []int{-23}
                    _ = v102
                    _ = v103
                }
            }
        } else {
            var v104 int64 = 42
            v105 := GenS10{F0: (v104 - (v104 | (v104 & genFn0(v88, -9223372036854775808)))), F1: ("a" + ""), F2: -11}
            _ = v104
            _ = v105
        }
    case "a":
        v88 = v88
    default:
        var v106 int8 = v2
        _ = v106
    }
    var v107 uint16 = 42
    v108 := []int{32768, -2147483648, 2147483647}
    v2 = (v2 << ((718) & 63))
    v107 = 1
    for fi109 := 0; fi109 < 11; fi109++ {
        v1 -= (^(v1))
        _ = fi109
    }
    v110 := GenS10{F0: 1, F1: ("foo bar" + "a"), F2: 1000}
    for fi111 := 0; fi111 < 3; fi111++ {
        var v112 uint32 = (^(v87))
        _ = fi111
        _ = v112
    }
    switch -1 {
    case -2147483648:
        var v113 uint16 = (v107 | 0)
        _ = v113
    }
    v88 = 1
    var v114 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v108:", len(v108), v108)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v107
    _ = v108
    _ = v110
    _ = v114
    _ = v2
    _ = v3
    _ = v86
    _ = v87
    _ = v88
}
