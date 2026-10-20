package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int8; F1 uint64 }
func genFn0(p0 int16) uint {
    if true {
        var v4 float32 = 1e10
        v5 := GenS1{F0: (p0 << ((4294967295) & 63)), F1: -2147483648}
        var v6 int16 = (p0 * (p0 + 1000))
        _ = v4
        _ = v5
        _ = v6
    }
    var v7 float64 = ((0.0 * 3.14159) - 0.1)
    var v8 uint8 = 2
    v9 := GenS1{F0: p0, F1: 32767}
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := []int{(fi10 ^ (fi10 % (((fi10 * (^(fi10)))) | 1))), (fi10 + 2147483647), 98}
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 int32 = 1
            fmt.Println(v13, v8, fi10)
            _ = len(v11)
            _ = fi12
            _ = v13
        }
        v8 += v8
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 218
}
func genFn1(p0 int, p1 int8) uint {
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            if ((((-2.25 + float32(-2.25)) / 1.5) * 1.5) != 3.14159) {
                var v16 int8 = p1
                p1 = (p1 * (p1 | (^(v16))))
                _ = v16
            }
            p1 = p1
            var v17 int32 = -28
            _ = fi15
            _ = v17
        }
        _ = fi14
    }
    p1 += (p1 / ((76) | 1))
    v18 := []int{(p0 ^ p0)}
    if false {
        v19 := GenS2{F0: 32767}
        for fi20 := 0; fi20 < 9; fi20++ {
            p1 += (p1 / ((p1) | 1))
            fmt.Println(p1, p1, p0)
            v21 := GenS3{F0: 4294967295}
            _ = fi20
            _ = v21
        }
        switch (("a" + "hello") + ("" + ("a" + "hello"))) {
        case "\n\t":
            v22 := GenS4{F0: p0}
            _ = v22
        case "":
            fmt.Println(p1, p0)
        default:
            p0 = p0
        }
        _ = v19
    } else {
        v23 := GenS4{F0: (p0 / (((p0 + 395)) | 1))}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v18
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{1, 2, 1}
    var v3 int32 = 0
    for fi24 := 0; fi24 < 9; fi24++ {
        v3 += (v3 << ((0) & 63))
        v25 := GenS2{F0: 0}
        _ = fi24
        _ = v25
    }
    v3 = v3
    v3 += (v3 << ((255) & 63))
    if (4294967295 <= 42) {
        v26 := GenS5{F0: 42, F1: 134, F2: -1}
        if (("" + "hello") <= (("hello" + "hello") + "\n\t")) {
            var v27 int8 = 112
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 uint16 = 534
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        _ = v26
    } else {
        fmt.Println(v3)
        v30 := GenS1{F0: -407, F1: 749}
        v31 := GenS6{F0: 127}
        _ = v30
        _ = v31
    }
    v3 += -247
    if (42 <= 9223372036854775807) {
        v3 = (-(v3))
        delete(v1, 65535)
    } else {
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 bool = false
            fmt.Println(v33, v3, v3)
            _ = fi32
            _ = v33
        }
        var v34 uint8 = 9
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 int16 = 32767
        _ = fi35
        _ = v36
    }
    delete(v1, 1)
    switch -847 {
    case 9223372036854775807:
        if (v3 > v3) {
            v37 := GenS7{F0: (v3 * -39)}
            _ = v37
        }
    case 1:
        _ = len(v1)
    case 32768:
        if ((v3 & 1) <= 42) {
            switch "foo bar" {
            case "a":
                switch 9223372036854775807 {
                case -2147483648:
                    for fi38 := 0; fi38 < 14; fi38++ {
                        var v39 int64 = 9223372036854775807
                        switch ("" + "foo bar") {
                        case "a":
                            v3 += (-(v3))
                        case "\n\t":
                            var v40 int64 = v39
                            _ = v40
                        }
                        if (32767 < -32768) {
                            v41 := GenS4{F0: (fi38 >> ((2) & 63))}
                            _ = v41
                        }
                        _ = fi38
                        _ = v39
                    }
                case 32768:
                    fmt.Println(v3)
                case 2147483647:
                    switch 127 {
                    case 42:
                        if (false || (true || (2 == 329))) {
                            v2 = append(v2, 882)
                            fmt.Println(v3)
                            var v42 string = "foo bar"
                            _ = v42
                        }
                    }
                }
            case "":
                switch "foo bar" {
                case "foo bar":
                    v3 = 42
                case "a":
                    v3 = v3
                }
            case "foo bar":
                if (false || true) {
                    v3 = (v3 & 1)
                    v43 := GenS2{F0: -32768}
                    var v44 uint64 = 4294967295
                    _ = v43
                    _ = v44
                } else {
                    v3 = v3
                    _ = len(v1)
                    for fi45 := 0; fi45 < 11; fi45++ {
                        v3 = (v3 + v3)
                        v3 = (^(v3))
                        _ = fi45
                    }
                }
            }
        } else {
            var v46 string = "a"
            v46 = v46
            var v47 int64 = 0
            _ = v46
            _ = v47
        }
    default:
        var v48 int8 = 0
        _ = v48
    }
    switch (("\n\t" + "hello") + "foo bar") {
    case "foo bar":
        v3 += v3
    default:
        fmt.Println(v3)
    }
    if (!true) {
        for fi49 := 0; fi49 < 6; fi49++ {
            for fi50 := 0; fi50 < 1; fi50++ {
                if false {
                    v51 := GenS3{F0: 2}
                    switch (fi50 - fi49) {
                    case -1:
                        var v52 uint32 = 4294967295
                        _ = v52
                    default:
                        if false {
                            v53 := GenS3{F0: 18446744073709551615}
                            v3 = (v3 * (v3 | v3))
                            if (true || true) {
                                var v54 uint64 = 42
                                v3 = (v3 & (v3 % (((v3 * v3)) | 1)))
                                v1[4294967295] = fi50
                                _ = v54
                            }
                            _ = v53
                        }
                    }
                    var v55 int8 = 1
                    _ = v51
                    _ = v55
                } else {
                    var v56 int8 = -128
                    fmt.Println(fi49)
                    if true {
                        var v57 string = "foo bar"
                        var v58 uint = genFn1(fi50, v56)
                        switch (fi49 << ((v58) & 63)) {
                        case 42:
                            var v59 int16 = -817
                            _ = v59
                        case 9223372036854775807:
                            v57 = "\n\t"
                        case 32767:
                            fmt.Println(v3, fi50)
                        }
                        _ = v57
                        _ = v58
                    } else {
                        v2 = append(v2, 198)
                    }
                    _ = v56
                }
                if false {
                    if true {
                        var v60 uint = 755
                        var v61 string = (((("" + "\n\t") + "foo bar") + "") + ("hello" + ""))
                        v61 = "\n\t"
                        var v62 uint32 = 2147483647
                        if (false && true) {
                            v63 := GenS8{F0: 0, F1: 89, F2: 223}
                            _ = v63
                        }
                        v64 := v1[(^(v62))]
                        _ = v60
                        _ = v61
                        _ = v62
                        _ = v64
                    }
                    fmt.Println(fi50, fi50)
                    if (1 > v3) {
                        var v65 uint8 = 127
                        v65 = v65
                        _ = v65
                    }
                }
                if (!(fi49 >= fi49)) {
                    v66 := GenS1{F0: 793, F1: fi49}
                    _ = v66
                } else {
                    if (42 <= -128) {
                        v67 := GenS5{F0: 65535, F1: 0, F2: 1}
                        switch fi50 {
                        case 9223372036854775807:
                            var v68 int64 = 24
                            _ = v68
                        }
                        _ = v67
                    }
                    v69 := GenS4{F0: 9223372036854775807}
                    switch -1 {
                    case 9223372036854775807:
                        v70 := GenS7{F0: -2147483648}
                        _ = v70
                    case 1:
                        _ = len(v2)
                    default:
                        var v71 float32 = (-2.25 / 1.0)
                        _ = v71
                    }
                    _ = v69
                }
                _ = fi50
            }
            if true {
                for fi72 := 0; fi72 < 7; fi72++ {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi49, fi49, v3)
                    }
                    if (v3 == v3) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v73 int = (^(fi72))
                        if (1 >= 42) {
                            var v74 uint64 = 18446744073709551615
                            var v75 uint16 = 868
                            _ = v74
                            _ = v75
                        }
                        _ = v73
                    }
                    var v76 int = 2147483647
                    _ = fi72
                    _ = v76
                }
            } else {
                v3 += (-(v3))
                if true {
                    if (!true) {
                        if false {
                            v3 += -549
                            if true {
                                v3 += -2147483648
                            } else {
                                for fi77 := 0; fi77 < 19; fi77++ {
                                    if false {
                                        var v78 float32 = 3.14159
                                        var v79 float64 = (1.5 / float64((v78 / v78)))
                                        switch (fi49 | fi77) {
                                        case 1:
                                            var v80 int32 = (v3 - -1)
                                            _ = v80
                                        case -341:
                                            fmt.Println(fi77, fi49, v3)
                                        }
                                        _ = v78
                                        _ = v79
                                    } else {
                                        var v81 uint = genFn1(fi49, -128)
                                        v81 -= (v81 >> ((0) & 63))
                                        v3 = v3
                                        _ = v81
                                    }
                                    _ = fi77
                                }
                            }
                            v3 -= 681
                        } else {
                            if false {
                                v82 := GenS3{F0: 442}
                                v3 += v3
                                _ = v82
                            }
                            var v83 uint = 2
                            _ = v83
                        }
                        var v84 uint8 = 255
                        v3 = (v3 >> ((538) & 63))
                        _ = v84
                    }
                    v3 = v3
                } else {
                    for fi85 := 0; fi85 < 20; fi85++ {
                        v86 := GenS2{F0: -32768}
                        _ = fi85
                        _ = v86
                    }
                    switch "hello" {
                    case "\n\t":
                        v2 = append(v2, 1)
                    case "":
                        var v87 int16 = -1
                        _ = v87
                    }
                    v88 := GenS9{F0: 42, F1: 42}
                    _ = v88
                }
                v3 = (v3 % ((42) | 1))
            }
            _ = fi49
        }
        v3 = -2147483648
        _ = len(v1)
    } else {
        for fi89 := 0; fi89 < 19; fi89++ {
            _ = len(v1)
            _ = fi89
        }
        for fi90 := 0; fi90 < 6; fi90++ {
            var v91 float32 = (1.5 * 1e10)
            for fi92 := 0; fi92 < 16; fi92++ {
                _ = len(v2)
                _ = fi92
            }
            if (true && true) {
                var v93 int = fi90
                if len(v2) > 0 {
                    v2[0] = 2
                }
                _ = v93
            } else {
                if (4294967295 >= 255) {
                    v91 = (0.1 / 1.5)
                    var v94 uint64 = 4294967295
                    _ = v94
                } else {
                    fmt.Println(v91)
                }
                _ = len(v1)
                if true {
                    v91 = (float32((-2.25 / float64(v91))) / v91)
                } else {
                    v95 := GenS10{F0: -97, F1: 0}
                    var v96 float32 = (3.14159 - (-2.25 / -2.25))
                    fmt.Println(v3)
                    _ = v95
                    _ = v96
                }
            }
            _ = fi90
            _ = v91
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
