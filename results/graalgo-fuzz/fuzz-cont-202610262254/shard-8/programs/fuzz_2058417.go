package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 string; F1 int; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 int16 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]string)
    var v3 string = "foo bar"
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            fmt.Println(fi4)
            switch v3 {
            case "a":
                switch (fi5 ^ (fi5 + fi4)) {
                case -1:
                    fmt.Println(fi5)
                case -2147483648:
                    var v6 int8 = 0
                    _ = v6
                case -542:
                    var v7 int16 = 1
                    _ = v7
                }
            default:
                v1 += ((v1 - v1) * v1)
                v1 -= (1.5 / v1)
            }
            var v8 int64 = 1
            _ = fi5
            _ = v8
        }
        _ = fi4
    }
    switch 9223372036854775807 {
    case 42:
        switch -1 {
        case -2147483648:
            switch -782 {
            case -965:
                var v9 float64 = -2.25
                _ = v9
            case -785:
                v10 := GenS1{F0: 1, F1: 4294967295, F2: 42}
                _ = v10
            default:
                var v11 int = 4294967295
                _ = v11
            }
        case 0:
            v12 := GenS1{F0: 0, F1: 209, F2: -2147483648}
            _ = v12
        default:
            fmt.Println(v1)
        }
    default:
        if true {
            v2[("" + v3)] = v3
            var v13 int64 = -1
            _ = v13
        }
    }
    var v14 float32 = (1e10 * 0.1)
    v1 = v1
    v2["a"] = v3
    switch (v3 + "hello") {
    case "hello":
        for fi15 := 0; fi15 < 5; fi15++ {
            if true {
                if false {
                    switch (fi15 - fi15) {
                    case -921:
                        switch (fi15 >> ((2) & 63)) {
                        case 27:
                            v14 += (0.1 / v14)
                        }
                    case -641:
                        v14 = (v14 * v1)
                    default:
                        var v16 uint8 = 255
                        _ = v16
                    }
                    fmt.Println(v14, v1, fi15)
                }
                fmt.Println(v1, fi15)
            }
            v14 = v1
            _ = fi15
        }
    case "\n\t":
        _ = len(v2)
    case "a":
        v1 += v14
    default:
        v17 := []int{-51, 42, 237}
        _ = v17
    }
    v14 -= float32(1.5)
    v18 := []int{128, -1, 32768, -2147483648}
    fmt.Println(v3, v3, v14)
    if true {
        v1 += ((v1 * (1e10 * float32(0.0))) + 3.14159)
        switch 128 {
        case 128:
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS1{F0: 1, F1: 715, F2: 961}
                _ = fi19
                _ = v20
            }
        case -571:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case 32768:
            v3 = v3
        default:
            switch 128 {
            case 42:
                if (v1 < (1e10 + (v1 * (((1e10 + 1e10) / (v14 + (v1 + v1))) + v1)))) {
                    v1 += float32(0.0)
                    v14 -= v14
                } else {
                    v21 := GenS1{F0: 243, F1: 4294967295, F2: 1}
                    delete(v2, ((v3 + v3) + v3))
                    var v22 int32 = 1
                    _ = v21
                    _ = v22
                }
            case 172:
                var v23 uint8 = 128
                _ = v23
            default:
                for fi24 := 0; fi24 < 20; fi24++ {
                    if true {
                        switch fi24 {
                        case -1:
                            switch ((v3 + "hello") + (v3 + "foo bar")) {
                            case "a":
                                if false {
                                    v25 := GenS1{F0: 413, F1: 42, F2: -137}
                                    fmt.Println(fi24, v1, v1)
                                    _ = v25
                                }
                            case "foo bar":
                                fmt.Println(fi24)
                            case "hello":
                                for fi26 := 0; fi26 < 2; fi26++ {
                                    var v27 uint64 = 18446744073709551615
                                    var v28 int8 = -1
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            }
                        case 32768:
                            v14 += 3.14159
                        default:
                            fmt.Println(fi24, v3, v14)
                        }
                        if (false && false) {
                            for fi29 := 0; fi29 < 2; fi29++ {
                                if (!true) {
                                    v14 = (1e10 / (v14 + v1))
                                }
                                v2["\n\t"] = ("a" + v3)
                                if (v14 < (v14 * float32(0.0))) {
                                    if ("\n\t" <= v3) {
                                        var v30 float32 = -2.25
                                        if false {
                                            v14 = v14
                                            _ = len(v18)
                                        }
                                        _ = v30
                                    } else {
                                        var v31 uint = 508
                                        _ = v31
                                    }
                                    v14 = (v14 + v1)
                                } else {
                                    if (true || false) {
                                        var v32 float64 = ((3.14159 + 0.1) * ((1e10 - 1.5) - ((float64((1.5 + float32(((0.1 - (3.14159 * 1e10)) * 0.0)))) - (1.5 * (3.14159 / 1.5))) - 1.5)))
                                        _ = v32
                                    }
                                    v14 = ((v1 / (v14 * 0.1)) + 1.5)
                                    var v33 float64 = ((0.0 - 0.0) - 1.5)
                                    _ = v33
                                }
                                v3 = ("" + v3)
                                _ = fi29
                            }
                            fmt.Println(v3)
                        } else {
                            var v34 int16 = int16(fi24)
                            var v35 int64 = -1
                            switch (v3 + v3) {
                            case "\n\t":
                                for fi36 := 0; fi36 < 2; fi36++ {
                                    var v37 string = v3
                                    v3 = (v3 + v37)
                                    _ = fi36
                                    _ = v37
                                }
                            }
                            _ = v34
                            _ = v35
                        }
                        for fi38 := 0; fi38 < 1; fi38++ {
                            v1 = v1
                            if true {
                                if true {
                                    _ = len(v18)
                                }
                                var v39 int64 = -525
                                if false {
                                    fmt.Println(v3, v1, fi38)
                                    switch v3 {
                                    case "hello":
                                        v40 := GenS1{F0: v39, F1: 2, F2: 2147483647}
                                        _ = v40
                                    }
                                    var v41 int64 = 0
                                    _ = len(v18)
                                    if (v3 < "a") {
                                        var v42 uint64 = 202
                                        switch fi38 {
                                        case 137:
                                            if true {
                                                switch (v3 + v3) {
                                                case "foo bar":
                                                    fmt.Println(fi24)
                                                case "\n\t":
                                                    if false {
                                                        v42 += 423
                                                        if (true || true) {
                                                            var v43 float32 = (3.14159 - 0.1)
                                                            v39 = 42
                                                            _ = len(v18)
                                                            _ = v43
                                                        } else {
                                                            if true {
                                                                v3 = ((("\n\t" + "foo bar") + v3) + (v3 + (v3 + v3)))
                                                                switch fi24 {
                                                                case 2147483647:
                                                                    v44 := GenS2{F0: 100, F1: (!true)}
                                                                    _ = v44
                                                                case 32767:
                                                                    v42 = 18446744073709551615
                                                                case 128:
                                                                    v2[(((v3 + v3) + v3) + v3)] = ("" + ("hello" + (v3 + v3)))
                                                                }
                                                                switch fi24 {
                                                                case 1:
                                                                    v1 -= -2.25
                                                                }
                                                            }
                                                            var v45 int64 = (v41 - -241)
                                                            _ = v45
                                                        }
                                                    } else {
                                                        v18 = append(v18, fi38)
                                                    }
                                                case "":
                                                    if (v3 >= "a") {
                                                        v46 := GenS3{F0: 100}
                                                        v47 := GenS1{F0: v41, F1: 42, F2: 186}
                                                        v42 = v42
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        var v48 bool = true
                                                        _ = v48
                                                    }
                                                }
                                                v42 = 275
                                                if true {
                                                    var v49 uint16 = 65535
                                                    _ = v49
                                                } else {
                                                    if (!(!false)) {
                                                        if (!true) {
                                                            v42 = 1
                                                        }
                                                        v50 := GenS1{F0: 593, F1: 4294967295, F2: int32(v42)}
                                                        var v51 uint32 = 65535
                                                        fmt.Println(v14, v1, v14)
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        fmt.Println(v41)
                                                        var v52 uint = 255
                                                        v53 := GenS1{F0: v41, F1: 65535, F2: 603}
                                                        v42 = v42
                                                        v54 := GenS2{F0: v52, F1: false}
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    v55 := GenS4{F0: 1}
                                                    var v56 int8 = 42
                                                    if false {
                                                        var v57 bool = false
                                                        if true {
                                                            var v58 float32 = v1
                                                            fmt.Println(v41, v56)
                                                            var v59 int = (fi38 ^ (fi24 - (fi24 + fi38)))
                                                            if v57 {
                                                                v60 := GenS3{F0: v56}
                                                                fmt.Println(v58, v57)
                                                                _ = v60
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            if (uint32(v42) > 4294967295) {
                                                                v42 = v42
                                                                switch (fi38 * 1) {
                                                                case 379:
                                                                    switch v3 {
                                                                    case "foo bar":
                                                                        if v57 {
                                                                            v61 := GenS5{F0: 32767, F1: (v39 >> ((65535) & 63))}
                                                                            _ = v61
                                                                        }
                                                                        var v62 uint64 = v42
                                                                        delete(v2, v3)
                                                                        _ = v62
                                                                    }
                                                                case 42:
                                                                    var v63 float64 = -2.25
                                                                    _ = v63
                                                                case -359:
                                                                    v14 -= 3.14159
                                                                default:
                                                                    delete(v2, v3)
                                                                }
                                                                _ = len(v18)
                                                            }
                                                        }
                                                        _ = v57
                                                    } else {
                                                        delete(v2, v3)
                                                        var v64 int8 = 1
                                                        _ = v64
                                                    }
                                                    v65 := v2[(((((v3 + ("hello" + v3)) + ((v3 + v3) + ("hello" + "a"))) + "\n\t") + v3) + v3)]
                                                    var v66 string = "a"
                                                    fmt.Println(v3)
                                                    _ = v55
                                                    _ = v56
                                                    _ = v65
                                                    _ = v66
                                                }
                                                if true {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v39 += v39
                                                    case "":
                                                        switch ("hello" + (("a" + v3) + "\n\t")) {
                                                        case "hello":
                                                            if false {
                                                                v1 += v14
                                                                if len(v18) >= 2 {
                                                                    v18 = v18[1:len(v18)]
                                                                }
                                                                switch v3 {
                                                                case "a":
                                                                    v1 += v14
                                                                default:
                                                                    v67 := GenS6{F0: (("foo bar" + (v3 + v3)) + (v3 + "\n\t")), F1: -1, F2: 335}
                                                                    var v68 int64 = (v39 % ((v41) | 1))
                                                                    if (556 != 42) {
                                                                        var v69 uint16 = 0
                                                                        _ = v69
                                                                    }
                                                                    if len(v18) > 0 {
                                                                        v18[0] = fi38
                                                                    }
                                                                    v18 = append(v18, fi24)
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                            } else {
                                                                if (532 < 70) {
                                                                    var v70 uint = 669
                                                                    _ = v70
                                                                }
                                                                switch v3 {
                                                                case "foo bar":
                                                                    if false {
                                                                        v71 := GenS4{F0: 1000}
                                                                        var v72 string = "hello"
                                                                        if (((v42 | (^(v42))) <= 0) || (!true)) {
                                                                            var v73 int = fi38
                                                                            _ = v73
                                                                        } else {
                                                                            var v74 uint32 = 466
                                                                            v75 := GenS7{F0: (253 <= 42)}
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                case "":
                                                                    switch (((v3 + (v3 + v3)) + "") + "a") {
                                                                    case "\n\t":
                                                                        var v76 float64 = 1.5
                                                                        _ = v76
                                                                    default:
                                                                        fmt.Println(v39)
                                                                    }
                                                                case "\n\t":
                                                                    var v77 int32 = -2147483648
                                                                    _ = v77
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(fi38, v14, v42)
                                                        }
                                                        v1 = (1.5 / 1.0)
                                                    case "\n\t":
                                                        var v78 float32 = 0.1
                                                        _ = v78
                                                    default:
                                                        v3 = v3
                                                    }
                                                    var v79 int8 = 81
                                                    v3 = "\n\t"
                                                    _ = v79
                                                }
                                            } else {
                                                _ = len(v18)
                                                var v80 uint8 = 255
                                                _ = v80
                                            }
                                        case 42:
                                            v81 := v2[v3]
                                            _ = v81
                                        default:
                                            var v82 string = ("a" + ((v3 + v3) + ("foo bar" + (v3 + (v3 + v3)))))
                                            _ = v82
                                        }
                                        v41 += (v39 >> ((813) & 63))
                                        if true {
                                            var v83 int32 = 1
                                            _ = v83
                                        } else {
                                            switch (^(fi38)) {
                                            case 128:
                                                v39 = -121
                                            case 1:
                                                var v84 int64 = (v41 | 0)
                                                _ = v84
                                            case -678:
                                                fmt.Println(v1)
                                            default:
                                                var v85 float32 = v1
                                                _ = v85
                                            }
                                            v18 = append(v18, (fi24 << ((2) & 63)))
                                        }
                                        _ = v42
                                    }
                                    _ = v41
                                } else {
                                    v86 := GenS8{F0: 532, F1: (2147483647 != 1)}
                                    v14 = v1
                                    var v87 uint = 627
                                    _ = v86
                                    _ = v87
                                }
                                _ = v39
                            } else {
                                var v88 float64 = 0.1
                                fmt.Println(v14, v88)
                                switch v3 {
                                case "hello":
                                    v3 = ("" + ((v3 + ("\n\t" + "")) + ""))
                                case "\n\t":
                                    v89 := GenS9{F0: 0, F1: 448}
                                    _ = v89
                                }
                                _ = v88
                            }
                            _ = fi38
                        }
                    } else {
                        fmt.Println(v14, fi24)
                    }
                    v1 = v1
                    if true {
                        var v90 int32 = 954
                        _ = v90
                    } else {
                        _ = len(v2)
                        if false {
                            fmt.Println(v3, fi24)
                            for fi91 := 0; fi91 < 14; fi91++ {
                                if true {
                                    if false {
                                        var v92 float64 = float64(v1)
                                        fmt.Println(v92)
                                        v93 := GenS10{F0: 772}
                                        _ = v92
                                        _ = v93
                                    }
                                }
                                _ = fi91
                            }
                        }
                        var v94 float64 = 0.1
                        _ = v94
                    }
                    _ = fi24
                }
            }
        }
        for fi95 := 0; fi95 < 19; fi95++ {
            v18 = append(v18, 119)
            _ = fi95
        }
    }
    v96 := GenS5{F0: 1, F1: -1}
    fmt.Println(v14)
    fmt.Println(v1, v14)
    v97 := GenS10{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v96
    _ = v97
}
