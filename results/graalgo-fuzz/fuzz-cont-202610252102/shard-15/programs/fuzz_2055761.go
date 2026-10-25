package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int; F1 int32; F2 uint }
type GenS6 struct { F0 int; F1 int; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int64; F2 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS10 struct { F0 int16; F1 uint16; F2 int16 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(fi6)
        fmt.Println(fi6)
        var v7 uint8 = 42
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 620
    v8 = (v8 & (v8 | (v8 ^ (v8 - (v8 | (v8 ^ (v8 + 14)))))))
    v8 = 2
    v9 := GenS1{F0: 4294967295}
    _ = v8
    _ = v9
    return 2
}
func genFn1() int8 {
    v10 := GenS2{F0: -381, F1: 480, F2: 402}
    var v11 string = ""
    for fi12 := 0; fi12 < 4; fi12++ {
        switch (v11 + v11) {
        case "hello":
            var v13 int32 = -98
            _ = v13
        case "foo bar":
            if true {
                var v14 int16 = -32768
                _ = v14
            } else {
                var v15 string = v11
                _ = v15
            }
        case "a":
            v16 := GenS3{F0: v11}
            _ = v16
        default:
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 uint64 = genFn0()
                _ = fi17
                _ = v18
            }
        }
        v11 = v11
        v11 = "\n\t"
        _ = fi12
    }
    _ = v10
    _ = v11
    return 103
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 float32 = (0.1 + 3.14159)
    v5 := make(map[uint32]uint32)
    v19 := GenS4{F0: genFn0(), F1: 127}
    var v20 uint32 = 342
    var v21 int32 = 466
    delete(v5, v20)
    fmt.Println(v20, v4, v20)
    v21 -= 1
    if ((3.14159 * v4) > float32((1e10 - 0.0))) {
        if false {
            v22 := []int{2147483647}
            v20 -= v20
            if ((0 > 255) && (1.5 >= 1e10)) {
                switch "foo bar" {
                case "hello":
                    var v23 uint64 = 4294967295
                    _ = v23
                default:
                    v4 = 0.0
                }
                if ((!(!(((0 >= 16) && true) || true))) && (((0 != 255) && ((((v21 == (v21 / (((v21 % (((v21 - v21)) | 1))) | 1))) && true) && true) && false)) || (false && ((v21 & v21) <= (v21 - 42))))) {
                    switch "" {
                    case "hello":
                        v1[v20] = false
                    }
                    v21 += v21
                    v21 = v21
                }
            }
            _ = v22
        }
        switch -2147483648 {
        case 0:
            v4 = v4
        case 127:
            v1[(^(v20))] = (true && (true && false))
        }
        delete(v3, (^(v20)))
    }
    if (128 == 0) {
        var v24 uint16 = 0
        switch -376 {
        case 980:
            var v25 string = "hello"
            _ = v25
        case 128:
            for fi26 := 0; fi26 < 18; fi26++ {
                v20 = (v20 << ((65535) & 63))
                _ = fi26
            }
        case 874:
            for fi27 := 0; fi27 < 2; fi27++ {
                switch (fi27 + 1) {
                case -1:
                    var v28 int = fi27
                    _ = v28
                case 128:
                    var v29 uint8 = 0
                    _ = v29
                case 539:
                    fmt.Println(v4, fi27)
                }
                _ = fi27
            }
        }
        _ = v24
    } else {
        v30 := []int{32767}
        if (!(-1 < int8(v21))) {
            if true {
                var v31 float32 = 3.14159
                _ = v31
            } else {
                switch ("foo bar" + "foo bar") {
                case "a":
                    for fi32 := 0; fi32 < 20; fi32++ {
                        fmt.Println(v20)
                        v33 := GenS3{F0: "foo bar"}
                        _ = fi32
                        _ = v33
                    }
                case "hello":
                    v4 = (((v4 / (v4 - (-2.25 + v4))) * v4) - 0.1)
                case "foo bar":
                    var v34 float32 = (v4 + 1.5)
                    _ = v34
                default:
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                }
                var v35 int64 = 42
                _ = v35
            }
            if len(v30) > 0 {
                v30[0] = 696
            }
        } else {
            if (!(42 != 4294967295)) {
                v4 += v4
                v21 -= (v21 / ((-1) | 1))
                for fi36 := 0; fi36 < 13; fi36++ {
                    var v37 int = int(v21)
                    fmt.Println(v20)
                    v38 := v3[v20]
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            } else {
                v21 += 2147483647
                for fi39 := 0; fi39 < 19; fi39++ {
                    if (((255 < 1) || true) && (((3.14159 / 1e10) / (float64(0.0) - 1.5)) <= float64(((3.14159 - v4) * (0.1 + 3.14159))))) {
                        for fi40 := 0; fi40 < 19; fi40++ {
                            v21 = v21
                            _ = fi40
                        }
                        if true {
                            v21 = v21
                            if true {
                                v20 += 2147483647
                                _ = len(v30)
                                var v41 int32 = (v21 - (^(v21)))
                                _ = v41
                            }
                        }
                    } else {
                        if (((106 != 1) && (genFn1() > -128)) && false) {
                            if ((("" + "\n\t") + (("hello" + ("" + "foo bar")) + "foo bar")) > "hello") {
                                for fi42 := 0; fi42 < 12; fi42++ {
                                    var v43 float32 = 0.0
                                    if ((v21 ^ 2147483647) >= 2147483647) {
                                        switch fi39 {
                                        case -1:
                                            var v44 int64 = 1
                                            fmt.Println(fi39, fi39, v21)
                                            v43 -= v43
                                            _ = v44
                                        case 128:
                                            switch "\n\t" {
                                            case "\n\t":
                                                v4 = v4
                                                if true {
                                                    v2[-1] = (v20 % ((62) | 1))
                                                    if (false || false) {
                                                        v20 += (v20 * v20)
                                                        v43 -= float32(-2.25)
                                                        switch (fi42 & (fi42 / (((fi42 / ((-26) | 1))) | 1))) {
                                                        case 42:
                                                            var v45 uint16 = 32767
                                                            _ = v45
                                                        case -1:
                                                            switch -9223372036854775808 {
                                                            case 9223372036854775807:
                                                                v46 := GenS5{F0: -284, F1: (-(v21)), F2: 1}
                                                                _ = v46
                                                            case 1:
                                                                v21 -= (v21 | v21)
                                                            }
                                                        case -9223372036854775808:
                                                            if (true && (true || ((v21 * (v21 % ((v21) | 1))) > (v21 >> ((279) & 63))))) {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    var v47 int64 = -1
                                                                    _ = v47
                                                                }
                                                                var v48 uint16 = 624
                                                                var v49 float32 = (v43 - ((-2.25 + (v4 / 1.5)) * 3.14159))
                                                                switch "a" {
                                                                case "\n\t":
                                                                    if len(v30) >= 2 {
                                                                        v30 = v30[1:len(v30)]
                                                                    }
                                                                    v50 := GenS3{F0: "hello"}
                                                                    v30 = append(v30, fi42)
                                                                    _ = len(v2)
                                                                    _ = v50
                                                                case "hello":
                                                                    var v51 int8 = genFn1()
                                                                    _ = v51
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                switch fi39 {
                                                                case 1:
                                                                    delete(v2, (fi42 * (fi42 + (-(fi39)))))
                                                                case 383:
                                                                    switch (("" + "a") + "\n\t") {
                                                                    case "a":
                                                                        v52 := v5[(v20 & v20)]
                                                                        if (false || (4294967295 != 184)) {
                                                                            if true {
                                                                                if ((v20 % ((506) | 1)) < (^(v20))) {
                                                                                    v53 := GenS2{F0: fi42, F1: (^(v52)), F2: 42}
                                                                                    if ((!(144 != 127)) && true) {
                                                                                        if (genFn0() != 617) {
                                                                                            v54 := GenS5{F0: 32767, F1: (v21 & (v21 >> ((42) & 63))), F2: 4294967295}
                                                                                            v20 += (^(v52))
                                                                                            _ = v54
                                                                                        } else {
                                                                                            if true {
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    v30 = append(v30, 128)
                                                                                                    switch fi39 {
                                                                                                    case 42:
                                                                                                        fmt.Println(v4)
                                                                                                        var v55 int64 = -424
                                                                                                        v56 := GenS6{F0: fi39, F1: fi39, F2: 127}
                                                                                                        _ = v55
                                                                                                        _ = v56
                                                                                                    case 1:
                                                                                                        v57 := GenS3{F0: (("\n\t" + "hello") + "a")}
                                                                                                        _ = v57
                                                                                                    case -2147483648:
                                                                                                        if true {
                                                                                                            v20 = v52
                                                                                                        } else {
                                                                                                            v58 := GenS3{F0: ""}
                                                                                                            _ = v58
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v52 -= 2147483648
                                                                                                var v59 float64 = ((1.5 - 0.1) + 3.14159)
                                                                                                _ = v59
                                                                                            }
                                                                                            v60 := GenS2{F0: fi39, F1: v52, F2: 957}
                                                                                            _ = len(v5)
                                                                                            _ = v60
                                                                                        }
                                                                                    }
                                                                                    v61 := GenS7{F0: 4294967295, F1: -640, F2: -9223372036854775808}
                                                                                    v52 += (^(v52))
                                                                                    _ = v53
                                                                                    _ = v61
                                                                                } else {
                                                                                    v62 := GenS8{F0: false}
                                                                                    if ((((fi42 / (((-(fi39))) | 1)) != -9223372036854775808) && true) || ((243 == 69) || false)) {
                                                                                        v4 += (v43 - v4)
                                                                                        var v63 uint = uint(v20)
                                                                                        _ = v63
                                                                                    }
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                            v64 := GenS6{F0: (fi39 ^ fi42), F1: (^(fi42)), F2: 239}
                                                                            var v65 int64 = 88
                                                                            v52 = 0
                                                                            v30 = append(v30, fi39)
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                        _ = v52
                                                                    case "":
                                                                        var v66 string = "\n\t"
                                                                        _ = v66
                                                                    case "hello":
                                                                        var v67 uint = 2
                                                                        _ = v67
                                                                    }
                                                                }
                                                                v43 += v43
                                                                var v68 uint = 882
                                                                v20 = 405
                                                                v43 = 1e10
                                                                _ = v68
                                                            }
                                                        default:
                                                            v43 += ((v4 / v4) + v4)
                                                        }
                                                        fmt.Println(v21, v43)
                                                        fmt.Println(v4, v4, v43)
                                                    }
                                                } else {
                                                    var v69 float32 = (float32((1.5 * 1e10)) / 0.1)
                                                    var v70 uint32 = 0
                                                    fmt.Println(v70, fi42, fi42)
                                                    _ = v69
                                                    _ = v70
                                                }
                                                if len(v30) > 0 {
                                                    v30[0] = 4294967295
                                                }
                                            case "hello":
                                                delete(v3, (v20 % (((^(v20))) | 1)))
                                            }
                                        }
                                    } else {
                                        if true {
                                            var v71 float32 = v4
                                            switch -818 {
                                            case 295:
                                                _ = len(v3)
                                            case 42:
                                                v72 := GenS9{F0: 42, F1: 42, F2: 150}
                                                _ = v72
                                            case 4294967295:
                                                var v73 int32 = v21
                                                _ = v73
                                            }
                                            var v74 int32 = (v21 * (v21 << ((274) & 63)))
                                            _ = v71
                                            _ = v74
                                        }
                                        _ = len(v3)
                                        if (!(0 >= 42)) {
                                            fmt.Println(fi42)
                                            v21 = (v21 + v21)
                                        } else {
                                            v75 := v5[(v20 / ((v20) | 1))]
                                            _ = v75
                                        }
                                    }
                                    _ = fi42
                                    _ = v43
                                }
                                v76 := GenS5{F0: fi39, F1: (-(v21)), F2: 1}
                                if (false && (1.5 < float64(v4))) {
                                    if (true && false) {
                                        switch ("a" + "\n\t") {
                                        case "hello":
                                            fmt.Println(v4, v4, v4)
                                        default:
                                            var v77 uint64 = 4294967295
                                            _ = v77
                                        }
                                        switch 955 {
                                        case 9223372036854775807:
                                            v78 := GenS10{F0: 1000, F1: 2, F2: -32768}
                                            _ = v78
                                        case -74:
                                            fmt.Println(v21)
                                        }
                                    }
                                }
                                _ = v76
                            } else {
                                var v79 float32 = 1e10
                                _ = v79
                            }
                        }
                        var v80 uint8 = 127
                        _ = v80
                    }
                    var v81 float64 = 3.14159
                    _ = fi39
                    _ = v81
                }
            }
        }
        _ = v30
    }
    v4 = (-2.25 + 1e10)
    v21 -= (v21 ^ (v21 - 42))
    for fi82 := 0; fi82 < 6; fi82++ {
        fmt.Println(fi82)
        _ = fi82
    }
    v4 = v4
    var v83 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v83
}
