package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 int; F2 uint16 }

func main() {
    var v1 string = (("hello" + ((("foo bar" + ("a" + "foo bar")) + ("a" + "a")) + "\n\t")) + ("\n\t" + "hello"))
    v2 := make(map[int]string)
    var v3 bool = (381 >= 1)
    v4 := []uint32{1, 1, 65535}
    var v5 int32 = -821
    switch "foo bar" {
    case "foo bar":
        var v6 string = v1
        _ = v6
    default:
        if len(v4) > 0 {
            v4[0] = 976
        }
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v3)
        var v8 bool = v3
        _ = fi7
        _ = v8
    }
    v4 = append(v4, 2147483647)
    v2[915] = ("" + (("\n\t" + v1) + (v1 + ((v1 + v1) + "foo bar"))))
    v3 = v3
    var v9 int32 = (v5 & (v5 & (-(v5))))
    if false {
        if v3 {
            v9 = (v9 | 1)
        } else {
            fmt.Println(v1, v1)
        }
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                switch (v1 + (("a" + (v1 + "\n\t")) + "\n\t")) {
                case "":
                    v5 = (v5 ^ (v5 * v9))
                case "\n\t":
                    var v12 uint8 = 2
                    _ = v12
                case "foo bar":
                    if (v3 && false) {
                        var v13 uint8 = 2
                        switch (-(fi11)) {
                        case 2147483647:
                            if (float64((-2.25 * 0.0)) <= -2.25) {
                                fmt.Println(v13)
                            } else {
                                var v14 int16 = 0
                                v9 -= v5
                                fmt.Println(v3, v3, v9)
                                _ = len(v4)
                                _ = v14
                            }
                            if v3 {
                                v4 = append(v4, 203)
                                if (!(!v3)) {
                                    var v15 uint32 = uint32(v13)
                                    var v16 float32 = (0.1 / 0.1)
                                    fmt.Println(v9, v13, v15)
                                    _ = v15
                                    _ = v16
                                } else {
                                    if (4294967295 < 943) {
                                        v17 := GenS1{F0: 1, F1: 0, F2: (^(v13))}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v17
                                    }
                                    _ = len(v2)
                                }
                            }
                        case 127:
                            var v18 int32 = (v5 * v9)
                            _ = v18
                        default:
                            var v19 int16 = 32767
                            _ = v19
                        }
                        _ = v13
                    } else {
                        v20 := GenS1{F0: 0, F1: 455, F2: 128}
                        _ = v20
                    }
                default:
                    if true {
                        var v21 uint = 255
                        switch ((("foo bar" + "\n\t") + "hello") + "foo bar") {
                        case "foo bar":
                            var v22 bool = (0 >= 1)
                            _ = v22
                        default:
                            switch v1 {
                            case "\n\t":
                                v4 = append(v4, 2)
                            case "foo bar":
                                fmt.Println(fi11, v5, fi11)
                            default:
                                fmt.Println(v3)
                            }
                        }
                        var v23 uint16 = 358
                        _ = v21
                        _ = v23
                    }
                    if v3 {
                        fmt.Println(v9)
                    } else {
                        var v24 bool = (!((v9 / ((2147483647) | 1)) > v5))
                        _ = len(v2)
                        fmt.Println(fi11, v1, fi11)
                        _ = len(v4)
                        if v3 {
                            var v25 int64 = -727
                            var v26 int8 = 100
                            v27 := v2[fi11]
                            if v3 {
                                var v28 uint = 2
                                var v29 int = fi10
                                _ = len(v2)
                                v25 = 1
                                fmt.Println(fi10, v25, fi11)
                                _ = v28
                                _ = v29
                            }
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        switch v1 {
                        case "":
                            var v30 int = 28
                            _ = v30
                        case "\n\t":
                            if v24 {
                                if false {
                                    v1 = "foo bar"
                                    var v31 int32 = -315
                                    v9 += -279
                                    if (((false && v24) || (77 >= fi10)) || (!(0 != 75))) {
                                        if len(v4) > 0 {
                                            v4[0] = 0
                                        }
                                    } else {
                                        fmt.Println(v24)
                                        fmt.Println(v3, fi11, v1)
                                        v24 = (1.5 > (0.1 * -2.25))
                                    }
                                    v32 := GenS1{F0: 2147483648, F1: 136, F2: 155}
                                    v24 = (false || v3)
                                    _ = v31
                                    _ = v32
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        _ = v24
                    }
                    v9 = 2147483647
                    if (float32(-2.25) < float32(float64((0.1 / 1e10)))) {
                        var v33 uint8 = 2
                        var v34 int = (fi10 % ((fi11) | 1))
                        _ = v33
                        _ = v34
                    }
                    switch v1 {
                    case "a":
                        fmt.Println(v5)
                    case "foo bar":
                        var v35 int8 = 127
                        _ = v35
                    case "\n\t":
                        v36 := GenS2{F0: 0}
                        _ = v36
                    default:
                        var v37 float64 = 0.1
                        fmt.Println(fi10)
                        v38 := GenS3{F0: 979, F1: 2}
                        _ = v37
                        _ = v38
                    }
                    if ((-2.25 >= 1e10) || v3) {
                        switch (v1 + (v1 + v1)) {
                        case "a":
                            var v39 float64 = 0.0
                            _ = v39
                        }
                        _ = len(v2)
                    }
                    v9 += v9
                }
                v5 = v5
                _ = fi11
            }
            v1 = "foo bar"
            _ = fi10
        }
    }
    v3 = (!(18446744073709551615 == 1))
    var v40 int32 = v9
    switch "hello" {
    case "a":
        v3 = true
    case "hello":
        v5 = v5
    }
    fmt.Println(v40, v3)
    v9 = (v9 % ((v5) | 1))
    v3 = (117 >= 255)
    v3 = v3
    if false {
        fmt.Println(v5, v1, v1)
        switch 9223372036854775807 {
        case 4294967295:
            for fi41 := 0; fi41 < 16; fi41++ {
                v1 = ("foo bar" + v1)
                v5 = v5
                if (v3 && (!v3)) {
                    var v42 uint8 = 255
                    switch ("a" + (v1 + "a")) {
                    case "a":
                        var v43 bool = (!((!(v3 || (((v3 || v3) || (false || v3)) && v3))) || v3))
                        _ = v43
                    case "foo bar":
                        v3 = true
                    default:
                        v1 = "foo bar"
                    }
                    _ = v42
                } else {
                    switch "hello" {
                    case "":
                        fmt.Println(v9, v9)
                    case "hello":
                        var v44 int = fi41
                        _ = v44
                    default:
                        fmt.Println(v9, fi41)
                    }
                    fmt.Println(v9)
                    if false {
                        v45 := GenS1{F0: 832, F1: 1, F2: 2}
                        v5 += v40
                        if (!v3) {
                            var v46 float64 = (((float64(1e10) + 0.0) - 1.5) / 1.0)
                            v1 = (v1 + v1)
                            if v3 {
                                _ = len(v4)
                                if (v3 || false) {
                                    switch 128 {
                                    case 32768:
                                        v1 = ((v1 + (v1 + ("\n\t" + v1))) + "\n\t")
                                    }
                                    v47 := GenS2{F0: 1}
                                    delete(v2, fi41)
                                    _ = v47
                                } else {
                                    if v3 {
                                        if false {
                                            switch 128 {
                                            case 32767:
                                                var v48 int64 = 45
                                                _ = v48
                                            case -607:
                                                v49 := GenS4{F0: 959, F1: (fi41 + (fi41 + -1)), F2: -20}
                                                _ = v49
                                            default:
                                                if v3 {
                                                    if (("" + "\n\t") == "foo bar") {
                                                        fmt.Println(fi41, v46)
                                                        var v50 int8 = 100
                                                        if (v3 || v3) {
                                                            fmt.Println(v9, fi41, v1)
                                                            v50 = v50
                                                            var v51 int = -492
                                                            _ = v51
                                                        } else {
                                                            v40 = (v40 >> ((4294967295) & 63))
                                                            v3 = v3
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        _ = v50
                                                    }
                                                    v9 = v5
                                                } else {
                                                    fmt.Println(fi41)
                                                    fmt.Println(fi41, v46, v9)
                                                }
                                            }
                                        }
                                        if (v1 == v1) {
                                            v46 = 1.5
                                            var v52 float64 = 0.1
                                            _ = len(v2)
                                            _ = v52
                                        } else {
                                            v53 := GenS5{F0: 1}
                                            if (-54 >= 100) {
                                                v1 = ""
                                                for fi54 := 0; fi54 < 14; fi54++ {
                                                    switch 362 {
                                                    case 42:
                                                        _ = len(v2)
                                                    case 9223372036854775807:
                                                        var v55 int16 = int16(fi54)
                                                        _ = v55
                                                    case -984:
                                                        v56 := GenS5{F0: 18446744073709551615}
                                                        _ = v56
                                                    default:
                                                        var v57 int16 = 32767
                                                        _ = v57
                                                    }
                                                    if (((1.5 / 3.14159) - 3.14159) > 0.1) {
                                                        var v58 uint64 = 220
                                                        if false {
                                                            if false {
                                                                if v3 {
                                                                    v59 := GenS6{F0: (^(v58)), F1: -771, F2: 0}
                                                                    fmt.Println(v40, v1, v1)
                                                                    var v60 uint32 = 25
                                                                    _ = v59
                                                                    _ = v60
                                                                } else {
                                                                    var v61 int16 = -32768
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = v61
                                                                }
                                                                var v62 uint16 = 434
                                                                _ = v62
                                                            } else {
                                                                v46 = 0.1
                                                            }
                                                            fmt.Println(fi41)
                                                            v58 = v58
                                                        } else {
                                                            v3 = v3
                                                        }
                                                        v63 := GenS1{F0: 88, F1: (v58 ^ 0), F2: 0}
                                                        _ = v58
                                                        _ = v63
                                                    }
                                                    v3 = v3
                                                    if true {
                                                        switch "hello" {
                                                        case "":
                                                            v5 += (v5 + v40)
                                                        case "a":
                                                            if v3 {
                                                                fmt.Println(fi41, v1)
                                                                v64 := GenS1{F0: 2, F1: 0, F2: 255}
                                                                _ = v64
                                                            } else {
                                                                v5 = 0
                                                            }
                                                        }
                                                        switch (fi54 | fi54) {
                                                        case 1:
                                                            v40 = -1
                                                            var v65 float64 = float64((0.1 - 3.14159))
                                                            var v66 int = fi54
                                                            _ = v65
                                                            _ = v66
                                                        case 32767:
                                                            var v67 float32 = (((3.14159 + (float32(v46) - 0.0)) - 3.14159) - 0.0)
                                                            _ = v67
                                                        default:
                                                            v68 := GenS2{F0: -30}
                                                            _ = v68
                                                        }
                                                        if false {
                                                            switch (fi54 << ((18446744073709551615) & 63)) {
                                                            case 128:
                                                                var v69 uint8 = 0
                                                                _ = v69
                                                            default:
                                                                if ((v40 % ((v9) | 1)) == (v40 + v40)) {
                                                                    v70 := GenS3{F0: 742, F1: 262}
                                                                    var v71 uint32 = 65535
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            }
                                                        }
                                                        v4 = append(v4, 42)
                                                    } else {
                                                        v72 := GenS4{F0: 2147483647, F1: fi41, F2: 42}
                                                        _ = v72
                                                    }
                                                    if v3 {
                                                        v46 = (1e10 / 1.0)
                                                        v3 = v3
                                                    } else {
                                                        var v73 string = "hello"
                                                        _ = len(v4)
                                                        var v74 uint16 = 630
                                                        v1 = v73
                                                        var v75 string = (v73 + (((((v73 + ((v73 + "a") + "\n\t")) + "hello") + ("\n\t" + (v73 + v1))) + ("a" + v1)) + v1))
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    _ = fi54
                                                }
                                                var v76 uint32 = 584
                                                _ = v76
                                            } else {
                                                for fi77 := 0; fi77 < 19; fi77++ {
                                                    v78 := GenS7{F0: -1}
                                                    if false {
                                                        if (-1 >= 331) {
                                                            v79 := GenS1{F0: 1, F1: 42, F2: 2}
                                                            v9 = v9
                                                            _ = v79
                                                        } else {
                                                            fmt.Println(v40)
                                                            fmt.Println(v1, v46)
                                                        }
                                                    } else {
                                                        if true {
                                                            v4 = append(v4, 0)
                                                            v4 = append(v4, 2147483648)
                                                            if ("\n\t" <= v1) {
                                                                if len(v4) > 0 {
                                                                    v4[0] = 42
                                                                }
                                                            } else {
                                                                fmt.Println(v3, v46, v3)
                                                                var v80 float64 = ((float64(-2.25) + ((3.14159 * ((0.0 / float64(3.14159)) - v46)) + (float64((0.0 * 0.0)) + (((1.5 * 3.14159) / v46) / ((float64(1e10) * (v46 - 1.5)) / v46))))) / v46)
                                                                _ = v80
                                                            }
                                                        }
                                                    }
                                                    _ = fi77
                                                    _ = v78
                                                }
                                            }
                                            _ = v53
                                        }
                                    } else {
                                        v81 := GenS7{F0: (v9 * v40)}
                                        for fi82 := 0; fi82 < 16; fi82++ {
                                            var v83 bool = ("a" <= v1)
                                            if len(v4) > 0 {
                                                v4[0] = 42
                                            }
                                            fmt.Println(v9, v46, v83)
                                            v5 -= (v40 - (v40 >> ((4294967295) & 63)))
                                            switch (^(fi41)) {
                                            case -2147483648:
                                                var v84 uint = 566
                                                _ = v84
                                            case 0:
                                                if v83 {
                                                    switch ("\n\t" + ("hello" + (v1 + v1))) {
                                                    case "":
                                                        if v83 {
                                                            fmt.Println(fi41, fi41, v9)
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v85 int16 = -1
                                                                _ = v85
                                                            case "\n\t":
                                                                var v86 uint32 = 2147483647
                                                                _ = v86
                                                            }
                                                        } else {
                                                            var v87 uint32 = 4294967295
                                                            v88 := GenS1{F0: v87, F1: 2, F2: 1}
                                                            v40 += -318
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v9, fi82)
                                                    default:
                                                        v2[fi82] = (v1 + (v1 + ""))
                                                    }
                                                    var v89 int32 = v40
                                                    v90 := GenS8{F0: v1, F1: fi82, F2: 255}
                                                    _ = v89
                                                    _ = v90
                                                } else {
                                                    v1 = v1
                                                    switch (fi41 >> ((4294967295) & 63)) {
                                                    case 127:
                                                        v9 += v5
                                                    default:
                                                        v4 = append(v4, 42)
                                                    }
                                                }
                                            }
                                            fmt.Println(v83, v5, v1)
                                            _ = fi82
                                            _ = v83
                                        }
                                        _ = v81
                                    }
                                    switch ("foo bar" + "a") {
                                    case "foo bar":
                                        switch "a" {
                                        case "\n\t":
                                            var v91 uint = 621
                                            _ = v91
                                        case "foo bar":
                                            v5 = (v5 / ((v5) | 1))
                                        case "":
                                            for fi92 := 0; fi92 < 18; fi92++ {
                                                v5 = v5
                                                var v93 int16 = 32767
                                                var v94 float64 = (float64(-2.25) / float64(0.1))
                                                v5 -= (v9 | (v40 % ((v5) | 1)))
                                                v3 = v3
                                                var v95 int32 = int32(fi92)
                                                var v96 int64 = 42
                                                var v97 bool = (v93 >= v93)
                                                _ = fi92
                                                _ = v93
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                                _ = v97
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                    case "\n\t":
                                        if len(v4) > 0 {
                                            v4[0] = 2147483647
                                        }
                                    default:
                                        v98 := v2[-451]
                                        _ = v98
                                    }
                                }
                                for fi99 := 0; fi99 < 13; fi99++ {
                                    fmt.Println(v46, v3, v5)
                                    if false {
                                        var v100 int = 128
                                        v2[9223372036854775807] = ""
                                        var v101 bool = true
                                        _ = v100
                                        _ = v101
                                    }
                                    _ = fi99
                                }
                            }
                            _ = v46
                        }
                        _ = v45
                    }
                }
                _ = fi41
            }
        case -887:
            for fi102 := 0; fi102 < 15; fi102++ {
                fmt.Println(v40)
                v1 = (v1 + (v1 + (("" + v1) + (v1 + v1))))
                _ = fi102
            }
        case 9223372036854775807:
            v5 = (v40 * v5)
        default:
            v40 = v5
        }
        v103 := GenS7{F0: v40}
        _ = v103
    } else {
        var v104 int8 = -1
        _ = v104
    }
    v40 = 42
    var v105 float64 = (0.1 - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
