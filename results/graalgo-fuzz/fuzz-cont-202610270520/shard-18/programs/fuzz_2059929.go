package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 int; F1 int32; F2 int64 }
type GenS4 struct { F0 int16; F1 int8; F2 int }
type GenS5 struct { F0 uint8; F1 int; F2 uint16 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 uint32; F2 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0(p0 string) int64 {
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: 100, F1: 0}
        if (!true) {
            var v7 string = ((("" + ((p0 + (("a" + "hello") + p0)) + (("foo bar" + "foo bar") + "a"))) + p0) + ((("a" + "\n\t") + p0) + p0))
            _ = v7
        }
        var v8 int16 = 42
        _ = fi5
        _ = v6
        _ = v8
    }
    var v9 float32 = float32(1.5)
    _ = p0
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{146, 2, 366, 1}
    var v2 int32 = 2147483647
    var v3 uint64 = 918
    v4 := make(map[string]bool)
    delete(v4, "a")
    fmt.Println(v3)
    v4[""] = (1 < -703)
    v2 = v2
    var v10 uint = 421
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            if (true || (((-82 <= -63) || ((^(v3)) < v3)) || (1 <= 142))) {
                var v13 uint = v10
                if true {
                    if len(v1) > 0 {
                        v1[0] = 898
                    }
                } else {
                    if true {
                        if false {
                            v13 = v10
                            switch (-(fi11)) {
                            case -1:
                                v4[""] = (1 >= 496)
                            default:
                                var v14 uint16 = 2
                                _ = v14
                            }
                            var v15 float64 = 0.1
                            _ = v15
                        } else {
                            fmt.Println(fi11, v13)
                            fmt.Println(fi12, v13)
                            v4[("a" + ("foo bar" + "a"))] = true
                        }
                        if true {
                            switch (fi11 - (fi12 + (-(fi11)))) {
                            case 128:
                                var v16 int64 = 9223372036854775807
                                _ = v16
                            case 32767:
                                if false {
                                    if (133 < 1) {
                                        var v17 int32 = v2
                                        _ = v17
                                    }
                                } else {
                                    v10 += (v10 + (^(v10)))
                                    v4[("a" + ("hello" + "foo bar"))] = ((-201 == 461) && true)
                                    fmt.Println(fi12, v3)
                                }
                            case 4294967295:
                                if (!(32767 < 255)) {
                                    v2 = v2
                                    if (1e10 < float64((1e10 - 0.1))) {
                                        v18 := GenS2{F0: -128, F1: -1, F2: v2}
                                        var v19 uint16 = 645
                                        v20 := GenS3{F0: 9223372036854775807, F1: -1, F2: 1}
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    } else {
                                        var v21 int64 = -1
                                        if false {
                                            switch "" {
                                            case "\n\t":
                                                var v22 string = "hello"
                                                _ = v22
                                            case "a":
                                                var v23 int = fi11
                                                _ = v23
                                            case "hello":
                                                v1 = append(v1, 770)
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 1000
                                            }
                                        }
                                        _ = v21
                                    }
                                    v2 = (^(v2))
                                } else {
                                    var v24 int = fi11
                                    _ = v24
                                }
                            }
                        } else {
                            switch (fi11 << ((309) & 63)) {
                            case 389:
                                v2 = (-(v2))
                            default:
                                if false {
                                    fmt.Println(fi11, v13)
                                    if true {
                                        switch (("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + "foo bar")) + "hello") {
                                        case "\n\t":
                                            fmt.Println(fi11, v3, v13)
                                        }
                                    } else {
                                        var v25 int = fi11
                                        _ = v25
                                    }
                                } else {
                                    _ = len(v1)
                                }
                            }
                            var v26 int32 = (v2 + v2)
                            _ = v26
                        }
                        var v27 int32 = (v2 % ((v2) | 1))
                        v28 := v4[""]
                        switch ("\n\t" + ("" + "hello")) {
                        case "hello":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "":
                            switch (((("hello" + "foo bar") + "\n\t") + "foo bar") + "a") {
                            case "foo bar":
                                v29 := GenS4{F0: 42, F1: 42, F2: (fi12 * fi12)}
                                _ = v29
                            case "\n\t":
                                v10 += (^(v10))
                            case "hello":
                                _ = len(v1)
                            default:
                                v30 := GenS4{F0: 42, F1: -100, F2: fi11}
                                _ = v30
                            }
                        case "foo bar":
                            v2 = (v27 & v2)
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(v27, v13, v27)
                        }
                        _ = v27
                        _ = v28
                    } else {
                        v31 := GenS5{F0: 17, F1: 127, F2: 32767}
                        v2 = v2
                        _ = v31
                    }
                }
                _ = v13
            } else {
                v2 = int32(v3)
                var v32 uint16 = 2
                _ = v32
            }
            var v33 uint32 = 634
            _ = fi12
            _ = v33
        }
        switch ("\n\t" + ("foo bar" + "a")) {
        case "foo bar":
            v3 = v3
        case "hello":
            _ = len(v1)
        }
        _ = fi11
    }
    if false {
        for fi34 := 0; fi34 < 15; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                if (((fi34 | (fi34 & (fi34 + int(v3)))) > (fi34 / ((fi35) | 1))) && (291 > 65535)) {
                    v36 := GenS6{F0: genFn0(""), F1: 220}
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    _ = v36
                } else {
                    fmt.Println(fi35)
                    v10 += 2
                }
                var v37 int = fi35
                v38 := GenS3{F0: (^(fi35)), F1: -886, F2: int64(v10)}
                fmt.Println(v3, v37, fi35)
                v10 -= (v10 & (v10 ^ v10))
                _ = fi35
                _ = v37
                _ = v38
            }
            if len(v1) > 0 {
                v1[0] = 429
            }
            if true {
                for fi39 := 0; fi39 < 16; fi39++ {
                    switch "" {
                    case "":
                        fmt.Println(v10, v2)
                    }
                    v40 := GenS6{F0: int64(v2), F1: 239}
                    v41 := GenS7{F0: 0}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                v2 += v2
                v42 := GenS8{F0: 104, F1: 65535, F2: -1}
                _ = v42
            } else {
                switch (fi34 + -9223372036854775808) {
                case -282:
                    v10 = 65535
                case 876:
                    for fi43 := 0; fi43 < 3; fi43++ {
                        switch -1 {
                        case 427:
                            var v44 float64 = 0.0
                            _ = v44
                        case 668:
                            if (float64(0.1) != (-2.25 / float64((1.5 * -2.25)))) {
                                if false {
                                    v45 := GenS6{F0: -1, F1: 127}
                                    var v46 float64 = (0.0 + -2.25)
                                    _ = v45
                                    _ = v46
                                } else {
                                    v10 -= 42
                                    v47 := GenS1{F0: 100, F1: 255}
                                    _ = v47
                                }
                                if ((v2 >= v2) || (v3 >= (v3 * (v3 + (v3 & (v3 ^ (v3 / ((v3) | 1)))))))) {
                                    fmt.Println(v2)
                                    if (!(false && (42 <= 2))) {
                                        var v48 uint = v10
                                        switch (("" + "a") + (((("foo bar" + ("a" + "a")) + ((("" + "foo bar") + "hello") + (("" + ("a" + "a")) + ""))) + (("\n\t" + (("a" + ("\n\t" + ("foo bar" + ""))) + ("" + ("hello" + ("foo bar" + ""))))) + ("" + "a"))) + "")) {
                                        case "\n\t":
                                            if (!true) {
                                                v10 = v48
                                                var v49 int64 = -959
                                                _ = v49
                                            }
                                        case "":
                                            v48 += 255
                                        }
                                        fmt.Println(fi43)
                                        _ = v48
                                    } else {
                                        var v50 int8 = 100
                                        switch fi34 {
                                        case -668:
                                            v51 := GenS2{F0: 0, F1: 0, F2: (v2 >> (((v10 + 131)) & 63))}
                                            _ = v51
                                        case 0:
                                            v52 := GenS3{F0: (fi34 - -834), F1: v2, F2: -526}
                                            _ = v52
                                        default:
                                            v50 = v50
                                        }
                                        v53 := GenS9{F0: "a"}
                                        _ = v50
                                        _ = v53
                                    }
                                    switch fi34 {
                                    case 128:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case 0:
                                        v54 := GenS2{F0: 120, F1: 715, F2: (v2 / ((v2) | 1))}
                                        _ = v54
                                    case 2147483647:
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(v3, fi34, fi43)
                                        case "foo bar":
                                            v3 -= uint64(v10)
                                        default:
                                            v55 := GenS8{F0: 125, F1: 889, F2: v2}
                                            _ = v55
                                        }
                                    }
                                }
                                if (("" + "a") == "hello") {
                                    var v56 int8 = -1
                                    _ = v56
                                } else {
                                    v57 := v4["foo bar"]
                                    v1 = append(v1, 994)
                                    if (!true) {
                                        if v57 {
                                            v10 = v10
                                        }
                                        v58 := GenS8{F0: -128, F1: 65535, F2: (v2 % ((550) | 1))}
                                        _ = v58
                                    }
                                    _ = v57
                                }
                            }
                        case 127:
                            if (v10 != v10) {
                                v59 := GenS1{F0: 103, F1: 233}
                                _ = v59
                            }
                        }
                        v60 := GenS1{F0: -124, F1: 29}
                        var v61 int = 9223372036854775807
                        _ = fi43
                        _ = v60
                        _ = v61
                    }
                }
                var v62 uint64 = (v3 ^ (v3 + v3))
                for fi63 := 0; fi63 < 15; fi63++ {
                    v3 = 42
                    _ = fi63
                }
                _ = v62
            }
            _ = fi34
        }
    } else {
        if true {
            var v64 uint8 = 0
            v65 := GenS6{F0: 42, F1: (v64 | v64)}
            _ = v64
            _ = v65
        }
        if (-2.25 > 1.5) {
            if true {
                v66 := GenS3{F0: 0, F1: -2147483648, F2: 42}
                if false {
                    v2 = v2
                    var v67 float32 = (0.0 + -2.25)
                    var v68 uint32 = uint32(v3)
                    _ = v67
                    _ = v68
                } else {
                    var v69 uint8 = 42
                    if false {
                        switch 0 {
                        case -2147483648:
                            fmt.Println(v69, v69, v10)
                        }
                    }
                    v70 := GenS10{F0: v2, F1: (!((v10 / ((2) | 1)) <= v10)), F2: uint64(v69)}
                    _ = v69
                    _ = v70
                }
                _ = v66
            } else {
                for fi71 := 0; fi71 < 14; fi71++ {
                    fmt.Println(v10)
                    fmt.Println(v2, fi71)
                    _ = fi71
                }
                var v72 int16 = 32767
                _ = v72
            }
            v2 = v2
            for fi73 := 0; fi73 < 19; fi73++ {
                v10 -= (v10 - (v10 & (v10 ^ 188)))
                var v74 uint16 = 255
                switch fi73 {
                case 32767:
                    if true {
                        var v75 uint16 = 342
                        for fi76 := 0; fi76 < 9; fi76++ {
                            if false {
                                switch int(v10) {
                                case -9223372036854775808:
                                    _ = len(v1)
                                case 766:
                                    v77 := GenS7{F0: 2}
                                    _ = v77
                                case 32768:
                                    v75 -= 2
                                }
                                var v78 int32 = v2
                                var v79 int32 = v2
                                _ = v78
                                _ = v79
                            } else {
                                var v80 string = "a"
                                _ = v80
                            }
                            var v81 float64 = (((1.5 / 1e10) * float64(1.5)) + ((float64(1.5) - 1e10) * 1e10))
                            _ = fi76
                            _ = v81
                        }
                        _ = v75
                    } else {
                        v4[("hello" + ("foo bar" + ("a" + "hello")))] = true
                        for fi82 := 0; fi82 < 14; fi82++ {
                            switch (fi73 + (^(fi82))) {
                            case 0:
                                fmt.Println(v74, fi73)
                            }
                            var v83 float64 = (3.14159 * ((0.0 * 1.5) + float64(1.5)))
                            _ = fi82
                            _ = v83
                        }
                    }
                case 0:
                    if (42 < 42) {
                        for fi84 := 0; fi84 < 15; fi84++ {
                            v85 := GenS9{F0: "foo bar"}
                            v2 += v2
                            _ = fi84
                            _ = v85
                        }
                        var v86 uint16 = 42
                        _ = v86
                    }
                }
                _ = fi73
                _ = v74
            }
        }
        var v87 uint16 = 42
        _ = v87
    }
    if true {
        var v88 uint32 = 995
        var v89 int64 = 42
        _ = v88
        _ = v89
    } else {
        delete(v4, "\n\t")
        switch 134 {
        case 9223372036854775807:
            if false {
                if (true && true) {
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                }
                var v90 uint16 = 0
                _ = len(v4)
                _ = v90
            } else {
                fmt.Println(v10)
            }
        case -1:
            fmt.Println(v3, v10, v10)
        default:
            v10 += 4294967295
        }
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
