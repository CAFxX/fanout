package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint8; F1 int32; F2 int32 }
type GenS6 struct { F0 uint8; F1 int8; F2 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = 3.14159
    var v3 int8 = -1
    var v4 bool = true
    var v5 int = 823
    var v6 int64 = 1
    v6 += -513
    v7 := GenS1{F0: 42}
    v2 = ((float64(1.5) + (v2 * 3.14159)) - 0.1)
    v8 := []int{v5, v5, (v5 + v5)}
    for fi9 := 0; fi9 < 13; fi9++ {
        switch (("a" + "foo bar") + ("foo bar" + "\n\t")) {
        case "\n\t":
            var v10 uint8 = 166
            _ = v10
        case "hello":
            if len(v8) > 0 {
                v8[0] = -2147483648
            }
        case "":
            if (-1 != 42) {
                var v11 int32 = 391
                for fi12 := 0; fi12 < 2; fi12++ {
                    switch 32767 {
                    case 4294967295:
                        var v13 bool = v4
                        _ = v13
                    case 624:
                        if v4 {
                            var v14 int16 = 0
                            _ = v14
                        }
                    default:
                        v15 := GenS1{F0: 1}
                        var v16 uint8 = 128
                        _ = v15
                        _ = v16
                    }
                    var v17 int64 = (v6 ^ v6)
                    _ = fi12
                    _ = v17
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v11
            }
        }
        var v18 float32 = 3.14159
        _ = fi9
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS2{F0: (v5 / (((fi19 + fi19)) | 1))}
        if (v4 || v4) {
            if true {
                _ = len(v8)
                v8 = append(v8, fi19)
            }
            v6 = 496
            v2 = ((0.1 / 1.0) + 0.0)
        }
        if (!(4294967295 <= 128)) {
            v5 -= (fi19 / (((-(v5))) | 1))
            v21 := GenS2{F0: v5}
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    if ((false && (v1 >= 687)) && false) {
        v22 := GenS3{F0: 121, F1: 884, F2: 182}
        if (1.5 >= 1e10) {
            if (!(42 > v6)) {
                var v23 uint32 = 0
                var v24 bool = (255 > 42)
                _ = v23
                _ = v24
            } else {
                switch "a" {
                case "":
                    fmt.Println(v3)
                default:
                    var v25 uint64 = 42
                    _ = v25
                }
                var v26 uint64 = v1
                v6 += v6
                _ = v26
            }
            v4 = (!v4)
        } else {
            v3 += 127
        }
        fmt.Println(v5, v6)
        _ = v22
    } else {
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 float64 = (v2 + 1e10)
            v29 := GenS4{F0: 0, F1: "\n\t"}
            _ = fi27
            _ = v28
            _ = v29
        }
        v6 += v6
        switch "foo bar" {
        case "foo bar":
            for fi30 := 0; fi30 < 17; fi30++ {
                if (v5 <= v5) {
                    fmt.Println(v2, v5)
                } else {
                    if (((v5 * 4294967295) > (v5 + (fi30 >> ((uint(v1)) & 63)))) && (v4 && v4)) {
                        v1 = v1
                        v31 := GenS2{F0: (fi30 - v5)}
                        _ = v31
                    } else {
                        v6 = (v6 + v6)
                    }
                }
                var v32 float32 = -2.25
                _ = fi30
                _ = v32
            }
        case "hello":
            switch "hello" {
            case "foo bar":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case "\n\t":
                v2 = 1.5
            }
        }
    }
    var v33 uint64 = v1
    var v34 int8 = v3
    if len(v8) > 0 {
        v8[0] = 335
    }
    v34 += (v3 % (((v34 % ((22) | 1))) | 1))
    for fi35 := 0; fi35 < 13; fi35++ {
        for fi36 := 0; fi36 < 5; fi36++ {
            v4 = true
            var v37 float32 = (0.1 + 3.14159)
            v5 += fi35
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        if (1 < 0) {
            if true {
                fmt.Println(fi38, v5)
                v39 := GenS5{F0: 2, F1: 0, F2: 173}
                v4 = (0 >= 1000)
                _ = v39
            } else {
                for fi40 := 0; fi40 < 8; fi40++ {
                    v1 = (v1 / ((0) | 1))
                    switch ("a" + "\n\t") {
                    case "hello":
                        var v41 int32 = int32(v6)
                        _ = v41
                    case "":
                        if ((v4 || (!v4)) || (!(!v4))) {
                            switch (fi40 * (fi38 * v5)) {
                            case -9223372036854775808:
                                switch 127 {
                                case 631:
                                    v8 = append(v8, 32768)
                                case 128:
                                    switch "a" {
                                    case "\n\t":
                                        var v42 int8 = v34
                                        _ = v42
                                    case "foo bar":
                                        var v43 bool = (9223372036854775807 < v5)
                                        _ = v43
                                    case "":
                                        v2 = (1e10 - 0.1)
                                    }
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    v3 = (v3 >> ((4294967295) & 63))
                                default:
                                    var v44 uint32 = 1
                                    _ = v44
                                }
                            case 32768:
                                if (!(int32(v6) == -1)) {
                                    if (2 < 1) {
                                        v45 := GenS3{F0: 18446744073709551615, F1: fi40, F2: 426}
                                        var v46 float64 = (v2 * (1.5 * 1e10))
                                        if false {
                                            if (-755 >= 544) {
                                                if len(v8) > 0 {
                                                    v8[0] = (^(v5))
                                                }
                                                v33 = v33
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                            }
                                            v8 = append(v8, 175)
                                        }
                                        _ = v45
                                        _ = v46
                                    }
                                    v6 = -9223372036854775808
                                }
                            case -972:
                                var v47 float64 = 1e10
                                _ = v47
                            }
                        } else {
                            _ = len(v8)
                            v2 = v2
                        }
                    case "foo bar":
                        v48 := GenS2{F0: v5}
                        _ = v48
                    default:
                        switch "foo bar" {
                        case "":
                            v8 = append(v8, v5)
                        case "\n\t":
                            v4 = v4
                        case "hello":
                            fmt.Println(fi40, v2, v34)
                        default:
                            var v49 int8 = v3
                            _ = v49
                        }
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    v34 = (v3 << ((42) & 63))
                    var v50 bool = (!v4)
                    if (1 == 128) {
                        switch "hello" {
                        case "":
                            var v51 uint = 323
                            _ = v51
                        }
                        var v52 uint = 0
                        _ = v52
                    } else {
                        switch "\n\t" {
                        case "a":
                            if ((1 < 146) || (255 == 255)) {
                                switch 2147483647 {
                                case -548:
                                    var v53 uint32 = 42
                                    _ = v53
                                default:
                                    v5 -= int(v3)
                                    if v50 {
                                        if (v4 || (!(!v50))) {
                                            var v54 string = "hello"
                                            if ((v6 < (v6 | v6)) && v50) {
                                                if (!v50) {
                                                    v55 := GenS5{F0: 1, F1: -324, F2: 0}
                                                    _ = v55
                                                } else {
                                                    var v56 uint64 = 18446744073709551615
                                                    fmt.Println(v1, fi40, fi38)
                                                    v1 = (v33 | 498)
                                                    fmt.Println(v3, v6, v1)
                                                    if len(v8) > 0 {
                                                        v8[0] = v5
                                                    }
                                                    v50 = v50
                                                    fmt.Println(v3, v50)
                                                    switch v5 {
                                                    case -825:
                                                        var v57 int8 = (v3 ^ v34)
                                                        _ = v57
                                                    case 32767:
                                                        v3 = v3
                                                    default:
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        v58 := GenS1{F0: 42}
                                                        _ = v58
                                                    }
                                                    if v4 {
                                                        v54 = v54
                                                        fmt.Println(v33, v50)
                                                        var v59 int16 = -1
                                                        _ = v59
                                                    } else {
                                                        v3 = -101
                                                        if len(v8) > 0 {
                                                            v8[0] = (fi38 + (fi40 << ((42) & 63)))
                                                        }
                                                        var v60 int16 = 403
                                                        _ = v60
                                                    }
                                                    _ = v56
                                                }
                                                switch (v54 + (v54 + (v54 + v54))) {
                                                case "a":
                                                    if len(v8) > 0 {
                                                        v8[0] = (fi40 % ((802) | 1))
                                                    }
                                                }
                                            } else {
                                                _ = len(v8)
                                                var v61 int8 = v34
                                                _ = v61
                                            }
                                            _ = v54
                                        } else {
                                            if (0 > 2) {
                                                _ = len(v8)
                                                v5 += v5
                                                v4 = (0 <= 128)
                                            } else {
                                                v6 = (v6 / ((-1) | 1))
                                                fmt.Println(v1)
                                            }
                                        }
                                        switch fi38 {
                                        case 9223372036854775807:
                                            var v62 uint16 = 69
                                            _ = v62
                                        case -1:
                                            v34 = v34
                                            v63 := GenS1{F0: -2147483648}
                                            _ = v63
                                        default:
                                            if (9223372036854775807 >= (-(v5))) {
                                                v1 += (v33 - (v1 % ((v1) | 1)))
                                                v5 = (v5 % (((v5 << ((884) & 63))) | 1))
                                                var v64 uint = uint(v33)
                                                v64 -= (v64 % ((4294967295) | 1))
                                                _ = v64
                                            }
                                        }
                                        v8 = append(v8, (v5 - 1))
                                        var v65 string = "a"
                                        fmt.Println(v6)
                                        v66 := GenS5{F0: 68, F1: 42, F2: -163}
                                        _ = v65
                                        _ = v66
                                    }
                                    if v4 {
                                        v67 := GenS5{F0: 1, F1: 0, F2: -1}
                                        if v50 {
                                            v34 = v34
                                            fmt.Println(fi40, v3)
                                            v68 := GenS1{F0: 42}
                                            _ = v68
                                        } else {
                                            var v69 float32 = 0.1
                                            var v70 uint8 = 160
                                            var v71 string = ("hello" + ("\n\t" + "\n\t"))
                                            v72 := GenS2{F0: fi40}
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v67
                                    } else {
                                        switch "hello" {
                                        case "":
                                            v73 := GenS3{F0: 462, F1: fi40, F2: 42}
                                            _ = v73
                                        case "hello":
                                            var v74 int8 = 42
                                            _ = v74
                                        case "foo bar":
                                            if ("foo bar" <= ("a" + ("a" + (("foo bar" + (((("\n\t" + "\n\t") + "foo bar") + "") + "\n\t")) + "foo bar")))) {
                                                if v50 {
                                                    if len(v8) > 0 {
                                                        v8[0] = fi38
                                                    }
                                                    if false {
                                                        v50 = (15 != -32768)
                                                        if false {
                                                            var v75 uint8 = 169
                                                            _ = v75
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                if (v33 < v33) {
                                                                    var v76 int16 = 32767
                                                                    var v77 int8 = (-(v3))
                                                                    _ = v76
                                                                    _ = v77
                                                                } else {
                                                                    fmt.Println(v33, v1)
                                                                    fmt.Println(v3, v34)
                                                                    v33 = (^(v33))
                                                                }
                                                            default:
                                                                v78 := GenS4{F0: 65535, F1: "a"}
                                                                if true {
                                                                    v3 = (v34 & (v3 * 42))
                                                                } else {
                                                                    v3 = v3
                                                                    v79 := GenS6{F0: uint8(v33), F1: (v34 / ((-1) | 1)), F2: (v6 >> ((255) & 63))}
                                                                    _ = v79
                                                                }
                                                                _ = v78
                                                            }
                                                            v34 -= 0
                                                        }
                                                        var v80 string = ("\n\t" + ("hello" + ""))
                                                        var v81 int8 = (v3 % ((v34) | 1))
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        if v50 {
                                                            v82 := GenS2{F0: (fi40 << ((0) & 63))}
                                                            _ = v82
                                                        } else {
                                                            v1 -= v1
                                                            v3 = (v34 - -90)
                                                            if true {
                                                                v8 = append(v8, 19)
                                                            } else {
                                                                v83 := GenS7{F0: ((("\n\t" + "") + (("" + "foo bar") + "a")) + ((("hello" + ("a" + ("" + "a"))) + "") + "hello"))}
                                                                if (v3 != v3) {
                                                                    v1 += (v33 & (v33 * (v33 << ((2) & 63))))
                                                                    if v4 {
                                                                        switch (v5 % ((fi38) | 1)) {
                                                                        case 127:
                                                                            v84 := GenS8{F0: v5}
                                                                            _ = v84
                                                                        default:
                                                                            v33 = (v33 << ((857) & 63))
                                                                        }
                                                                    }
                                                                    fmt.Println(fi40, v50, v3)
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    var v85 float64 = float64(((1.5 / (1.5 / 0.1)) - float32(float64(0.0))))
                                                                    _ = v85
                                                                case "":
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v34, fi38, v33)
                                                                default:
                                                                    fmt.Println(v34)
                                                                }
                                                                _ = v83
                                                            }
                                                            var v86 int32 = -181
                                                            switch ("" + "a") {
                                                            case "hello":
                                                                var v87 uint16 = 255
                                                                _ = v87
                                                            }
                                                            _ = v86
                                                        }
                                                        switch (^(fi38)) {
                                                        case -892:
                                                            fmt.Println(v34, v34)
                                                        }
                                                        v3 = v3
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        var v88 float32 = 3.14159
                                                        _ = v88
                                                    }
                                                } else {
                                                    if (!(!((!((v6 - (v6 - (-(v6)))) >= int64(fi38))) && true))) {
                                                        if v4 {
                                                            v2 = v2
                                                        } else {
                                                            v6 = (-(v6))
                                                        }
                                                    }
                                                }
                                                _ = len(v8)
                                            }
                                        }
                                        var v89 uint8 = 42
                                        _ = v89
                                    }
                                    if v4 {
                                        var v90 uint64 = v33
                                        var v91 int64 = (v6 >> ((548) & 63))
                                        v5 -= fi38
                                        _ = v90
                                        _ = v91
                                    }
                                    var v92 uint32 = 0
                                    _ = v92
                                }
                                var v93 int16 = 579
                                _ = v93
                            } else {
                                var v94 string = ""
                                switch "" {
                                case "foo bar":
                                    v4 = true
                                default:
                                    var v95 uint8 = 0
                                    _ = v95
                                }
                                switch "" {
                                case "\n\t":
                                    v96 := GenS8{F0: 42}
                                    _ = v96
                                case "a":
                                    var v97 int16 = 559
                                    _ = v97
                                default:
                                    var v98 uint32 = 2
                                    _ = v98
                                }
                                _ = v94
                            }
                        case "":
                            if v4 {
                                v4 = (!(!v4))
                                switch ("\n\t" + ("\n\t" + ("" + "\n\t"))) {
                                case "hello":
                                    var v99 float32 = 1e10
                                    _ = v99
                                case "\n\t":
                                    fmt.Println(v50, v5, v33)
                                default:
                                    v1 -= 18446744073709551615
                                }
                                v33 += 623
                            } else {
                                var v100 uint16 = uint16(v33)
                                _ = v100
                            }
                        case "\n\t":
                            fmt.Println(v5, v6)
                        default:
                            if v4 {
                                fmt.Println(v5, v50, v6)
                                var v101 int64 = v6
                                v2 = float64(0.0)
                                _ = v101
                            } else {
                                if v4 {
                                    v8 = append(v8, v5)
                                    v33 += v1
                                    var v102 int8 = (^(v34))
                                    v1 = (v33 - (v1 - 4294967295))
                                    _ = v102
                                } else {
                                    if v50 {
                                        if true {
                                            v103 := GenS9{F0: -896}
                                            v104 := GenS10{F0: (v5 + (fi40 | fi38)), F1: 32767}
                                            v105 := GenS11{F0: -862}
                                            _ = v103
                                            _ = v104
                                            _ = v105
                                        } else {
                                            v33 = v33
                                        }
                                        v34 += (v34 - v3)
                                        v5 = 32767
                                    } else {
                                        if len(v8) > 0 {
                                            v8[0] = (^(fi40))
                                        }
                                    }
                                    _ = len(v8)
                                    v8 = append(v8, (v5 ^ fi38))
                                }
                                var v106 uint64 = (^(v33))
                                if false {
                                    v1 += v33
                                    var v107 int = 32768
                                    _ = v107
                                }
                                _ = v106
                            }
                        }
                    }
                    v6 -= v6
                    v108 := GenS12{F0: false}
                    _ = fi40
                    _ = v108
                    _ = v50
                }
                for fi109 := 0; fi109 < 1; fi109++ {
                    if false {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        v1 = v1
                        if false {
                            v5 = -905
                            if (!true) {
                                var v110 uint64 = 2
                                _ = v110
                            } else {
                                fmt.Println(v2, fi38)
                                var v111 int64 = (v6 ^ v6)
                                v1 -= v1
                                fmt.Println(v5, v3)
                                _ = v111
                            }
                            if true {
                                fmt.Println(v3)
                                v112 := GenS11{F0: int16(fi38)}
                                _ = v112
                            }
                            if len(v8) > 0 {
                                v8[0] = (^(fi38))
                            }
                            var v113 float64 = 1.5
                            v6 += 292
                            _ = v113
                        }
                    }
                    _ = fi109
                }
                v114 := GenS1{F0: -422}
                _ = v114
            }
            v5 = 635
        }
        v6 += (v6 ^ (v6 & (v6 << ((uint(v33)) & 63))))
        var v115 string = ((("a" + ("a" + "")) + "foo bar") + "\n\t")
        _ = fi38
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
