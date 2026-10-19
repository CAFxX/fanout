package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 bool; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 uint8; F1 int; F2 int8 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 int8; F1 int16 }
func genFn0(p0 uint8, p1 uint64) int {
    if true {
        fmt.Println(p1, p0)
    }
    if false {
        var v6 uint64 = (p1 % ((p1) | 1))
        _ = v6
    } else {
        if true {
            v7 := GenS1{F0: 939, F1: 802}
            _ = v7
        }
        v8 := GenS1{F0: 0, F1: 42}
        _ = v8
    }
    p1 += p1
    var v9 uint8 = 179
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 1
    v3 := []uint32{}
    var v4 int = 9223372036854775807
    var v5 uint64 = 2
    delete(v1, "a")
    v5 = 0
    if (genFn0(0, v5) > (v4 ^ (-(v4)))) {
        var v10 uint64 = 197
        var v11 int = v4
        _ = v10
        _ = v11
    }
    var v12 uint8 = 1
    v13 := GenS2{F0: 42, F1: v5, F2: v5}
    switch "hello" {
    case "\n\t":
        v3 = append(v3, uint32(v5))
    }
    var v14 uint64 = v5
    switch "foo bar" {
    case "":
        if false {
            v14 -= v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                if true {
                    v16 := GenS2{F0: -608, F1: 2, F2: (^(v5))}
                    switch (((("" + "") + "foo bar") + "hello") + "") {
                    case "":
                        fmt.Println(fi15, v5)
                    case "foo bar":
                        v17 := GenS2{F0: 42, F1: (v5 + v5), F2: (v14 >> ((752) & 63))}
                        _ = v17
                    }
                    v18 := GenS3{F0: v2, F1: false, F2: v2}
                    _ = v16
                    _ = v18
                } else {
                    if false {
                        v12 -= (^(v12))
                        var v19 float32 = 0.1
                        _ = v19
                    }
                    switch fi15 {
                    case -603:
                        switch (-(fi15)) {
                        case -1:
                            fmt.Println(fi15, fi15)
                        case 840:
                            for fi20 := 0; fi20 < 5; fi20++ {
                                _ = len(v1)
                                if true {
                                    if false {
                                        if (!false) {
                                            v12 = (v12 % ((uint8(v5)) | 1))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v21 int8 = 100
                                            _ = v21
                                        } else {
                                            if (!((^(v2)) >= (-(v2)))) {
                                                switch fi15 {
                                                case 9223372036854775807:
                                                    v22 := v1[("" + "\n\t")]
                                                    _ = v22
                                                case 480:
                                                    fmt.Println(fi15)
                                                case 1:
                                                    v5 -= (v14 >> ((1) & 63))
                                                default:
                                                    fmt.Println(fi20)
                                                }
                                                if (!(-1 < -1)) {
                                                    var v23 int8 = 100
                                                    _ = v23
                                                } else {
                                                    v12 = v12
                                                }
                                                v14 = 42
                                            }
                                            var v24 uint = 18446744073709551615
                                            _ = v24
                                        }
                                        v14 -= 0
                                        var v25 string = ("a" + ((("\n\t" + "\n\t") + ("hello" + ("a" + "hello"))) + ""))
                                        _ = v25
                                    } else {
                                        if false {
                                            if (("\n\t" + "a") >= "") {
                                                v26 := GenS1{F0: 2, F1: 360}
                                                if (-32768 > -1) {
                                                    if false {
                                                        if true {
                                                            v27 := v1["\n\t"]
                                                            _ = v27
                                                        }
                                                        v4 -= 4294967295
                                                    } else {
                                                        var v28 int = (-(fi15))
                                                        switch (fi15 ^ (fi20 * (v4 + (^(fi15))))) {
                                                        case 9223372036854775807:
                                                            if (!false) {
                                                                switch (fi20 % (((fi20 % ((fi15) | 1))) | 1)) {
                                                                case 42:
                                                                    var v29 uint64 = (v14 - (v14 % ((v5) | 1)))
                                                                    _ = v29
                                                                }
                                                                var v30 uint64 = (^(v14))
                                                                _ = v30
                                                            }
                                                        case 32768:
                                                            v31 := GenS1{F0: 2147483648, F1: 32767}
                                                            fmt.Println(v12, v4, fi20)
                                                            _ = v31
                                                        }
                                                        _ = v28
                                                    }
                                                    v4 = genFn0(v12, v5)
                                                    switch "a" {
                                                    case "hello":
                                                        if (535 >= 121) {
                                                            delete(v1, ("hello" + ""))
                                                            v32 := GenS4{F0: v12}
                                                            if true {
                                                                fmt.Println(fi15, v14)
                                                                var v33 int32 = 42
                                                                _ = v33
                                                            }
                                                            _ = v32
                                                        } else {
                                                            v34 := v1["hello"]
                                                            v34 = v34
                                                            v35 := GenS2{F0: 32767, F1: 4294967295, F2: v5}
                                                            v34 = ((1e10 + (1e10 - 0.1)) != 1.5)
                                                            var v36 uint64 = (v14 ^ (^(v14)))
                                                            v34 = v34
                                                            if v34 {
                                                                v1["hello"] = v34
                                                                switch (fi15 / ((fi15) | 1)) {
                                                                case 423:
                                                                    v12 -= v12
                                                                    fmt.Println(v12, v2)
                                                                    var v37 uint64 = (^(v5))
                                                                    _ = v37
                                                                default:
                                                                    v38 := v1[("hello" + "")]
                                                                    v39 := GenS1{F0: 42, F1: uint16(v12)}
                                                                    switch "" {
                                                                    case "":
                                                                        v2 -= v2
                                                                    case "hello":
                                                                        v40 := GenS5{F0: (^(v4)), F1: ((56 >= 65535) || v34)}
                                                                        _ = v40
                                                                    case "foo bar":
                                                                        var v41 uint = 2
                                                                        _ = v41
                                                                    default:
                                                                        v36 = v36
                                                                    }
                                                                    _ = v38
                                                                    _ = v39
                                                                }
                                                                v42 := GenS6{F0: 2, F1: 39}
                                                                _ = v42
                                                            } else {
                                                                if (!((3.14159 - 1e10) != 0.1)) {
                                                                    v43 := GenS4{F0: 128}
                                                                    v14 += v5
                                                                    _ = v43
                                                                }
                                                            }
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    case "foo bar":
                                                        switch (v4 | 32768) {
                                                        case 42:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if true {
                                                                    delete(v1, (((("hello" + ("" + "")) + "hello") + (("a" + "hello") + ((("\n\t" + "hello") + ("\n\t" + "hello")) + (("hello" + ("a" + "\n\t")) + "a")))) + "foo bar"))
                                                                    fmt.Println(v5, v14)
                                                                    v14 = (v14 % ((0) | 1))
                                                                }
                                                            case "foo bar":
                                                                if len(v3) > 0 {
                                                                    v3[0] = 2147483648
                                                                }
                                                            case "hello":
                                                                v44 := GenS7{F0: 56}
                                                                _ = v44
                                                            default:
                                                                var v45 int64 = v2
                                                                _ = v45
                                                            }
                                                        case 4294967295:
                                                            var v46 float64 = 1e10
                                                            var v47 int32 = 0
                                                            v2 = v2
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        v5 = uint64(v12)
                                                    default:
                                                        if (!true) {
                                                            if ((v14 == 904) && (65535 == uint(v5))) {
                                                                fmt.Println(v4)
                                                            }
                                                            switch (fi20 ^ v4) {
                                                            case 32768:
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v12 = (v12 - (v12 >> ((4294967295) & 63)))
                                                                }
                                                            case 366:
                                                                v14 += uint64(v12)
                                                            case 888:
                                                                v48 := GenS7{F0: 94}
                                                                _ = v48
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v26
                                            }
                                            v4 = 128
                                        }
                                        switch fi20 {
                                        case -1:
                                            var v49 uint8 = 127
                                            switch "foo bar" {
                                            case "":
                                                delete(v1, "foo bar")
                                                switch "a" {
                                                case "a":
                                                    v1[""] = false
                                                }
                                                if ("a" >= "foo bar") {
                                                    v14 = uint64(v49)
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v3 = append(v3, 124)
                                                    case "":
                                                        var v50 bool = (65535 != 65535)
                                                        _ = v50
                                                    }
                                                    if false {
                                                        var v51 float32 = (0.0 * 1.5)
                                                        if (v14 == (v14 * v5)) {
                                                            v52 := GenS1{F0: 1, F1: 486}
                                                            _ = v52
                                                        }
                                                        _ = v51
                                                    }
                                                }
                                                var v53 uint16 = 255
                                                _ = v53
                                            default:
                                                switch (-(v4)) {
                                                case -104:
                                                    switch fi15 {
                                                    case -9223372036854775808:
                                                        if true {
                                                            v54 := GenS1{F0: 42, F1: 659}
                                                            _ = v54
                                                        } else {
                                                            _ = len(v1)
                                                        }
                                                    case 4294967295:
                                                        delete(v1, "a")
                                                    case 781:
                                                        v55 := GenS3{F0: -757, F1: ((1e10 <= -2.25) || (2 < uint32(v49))), F2: 0}
                                                        _ = v55
                                                    default:
                                                        fmt.Println(fi15)
                                                    }
                                                case -1:
                                                    fmt.Println(v5, v5, v4)
                                                default:
                                                    var v56 float64 = 1.5
                                                    _ = v56
                                                }
                                            }
                                            var v57 int8 = 0
                                            _ = v49
                                            _ = v57
                                        case 42:
                                            var v58 float32 = (0.0 / 1.0)
                                            _ = v58
                                        }
                                        v14 -= 937
                                    }
                                }
                                v59 := GenS4{F0: (v12 * 255)}
                                _ = fi20
                                _ = v59
                            }
                        case 88:
                            var v60 float64 = (0.0 - 1e10)
                            _ = v60
                        default:
                            v61 := GenS3{F0: v2, F1: (v12 != uint8(v5)), F2: v2}
                            _ = v61
                        }
                    default:
                        for fi62 := 0; fi62 < 6; fi62++ {
                            v63 := GenS8{F0: true, F1: -47}
                            switch 1 {
                            case -1:
                                if ((!(367 == 925)) && (true || (255 >= 80))) {
                                    v64 := GenS8{F0: ("hello" > ((((("hello" + ("foo bar" + "\n\t")) + ("a" + "")) + "") + "foo bar") + "\n\t")), F1: -1}
                                    switch genFn0((v12 >> ((1) & 63)), v14) {
                                    case 128:
                                        var v65 uint64 = (v5 % ((147) | 1))
                                        _ = v65
                                    case 32767:
                                        _ = len(v3)
                                    default:
                                        v66 := GenS9{F0: v12, F1: 1}
                                        _ = v66
                                    }
                                    var v67 int64 = (v2 * v2)
                                    _ = v64
                                    _ = v67
                                }
                                if false {
                                    var v68 float32 = float32(3.14159)
                                    v3 = append(v3, 1)
                                    v2 = (v2 & 436)
                                    _ = v68
                                }
                                _ = len(v1)
                                v69 := GenS10{F0: 2147483647, F1: 190}
                                _ = v69
                            case -9223372036854775808:
                                if (false && (!false)) {
                                    switch ("" + "\n\t") {
                                    case "foo bar":
                                        v2 += (v2 + v2)
                                    case "":
                                        delete(v1, "hello")
                                    }
                                }
                            }
                            _ = fi62
                            _ = v63
                        }
                    }
                    for fi70 := 0; fi70 < 8; fi70++ {
                        v14 += 83
                        _ = fi70
                    }
                }
                _ = fi15
            }
            v4 -= (^(v4))
            switch (v4 * genFn0((^(v12)), v5)) {
            case 32768:
                v2 = v2
            case -1:
                if (127 > 0) {
                    for fi71 := 0; fi71 < 1; fi71++ {
                        var v72 int16 = -32768
                        v73 := GenS8{F0: (!((v2 & v2) <= 1)), F1: 0}
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    if (127 != (v12 ^ (^(v12)))) {
                        var v74 bool = true
                        v12 = (v12 | (v12 << ((0) & 63)))
                        _ = v74
                    }
                    for fi75 := 0; fi75 < 7; fi75++ {
                        v2 += (^(v2))
                        v2 += -1
                        _ = fi75
                    }
                } else {
                    fmt.Println(v14, v2)
                    if (768 > 357) {
                        var v76 uint32 = 115
                        delete(v1, ("a" + "foo bar"))
                        v3 = append(v3, v76)
                        _ = v76
                    }
                    var v77 int32 = -1
                    _ = v77
                }
            case 1:
                fmt.Println(v2, v14)
            }
        }
    }
    v78 := GenS11{F0: v12, F1: genFn0(v12, (v5 & v5)), F2: 0}
    fmt.Println(v12, v4, v2)
    v79 := GenS12{F0: v2, F1: 255}
    switch v4 {
    case 325:
        var v80 int32 = -1
        _ = v80
    case 4294967295:
        var v81 int8 = -1
        _ = v81
    case -9223372036854775808:
        for fi82 := 0; fi82 < 15; fi82++ {
            v3 = append(v3, 50)
            v83 := GenS13{F0: -46, F1: 495}
            _ = fi82
            _ = v83
        }
    }
    v84 := GenS7{F0: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v79
    _ = v84
}
