package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int16, p1 bool) bool {
    v5 := []int{}
    var v6 uint = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (true && false)
}
func genFn1() int32 {
    v7 := []int{9223372036854775807, -21, 4294967295, 0}
    var v8 int = 604
    if (false && false) {
        for fi9 := 0; fi9 < 10; fi9++ {
            v8 += v8
            var v10 bool = (4294967295 < 680)
            switch (v8 | fi9) {
            case 32768:
                if (true || v10) {
                    if (1 <= 1) {
                        for fi11 := 0; fi11 < 2; fi11++ {
                            var v12 float64 = (1e10 - float64(3.14159))
                            v12 = (v12 * (v12 * 0.0))
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        for fi13 := 0; fi13 < 5; fi13++ {
                            v8 -= fi13
                            var v14 uint16 = 1
                            switch (fi13 % (((-(fi13))) | 1)) {
                            case 127:
                                var v15 float64 = -2.25
                                _ = v15
                            case 775:
                                v14 = (v14 | (v14 | (^(v14))))
                            case 0:
                                v16 := GenS1{F0: 0, F1: 106}
                                _ = v16
                            }
                            _ = fi13
                            _ = v14
                        }
                        if len(v7) > 0 {
                            v7[0] = (v8 | (v8 % ((v8) | 1)))
                        }
                    }
                } else {
                    if (3.14159 < (0.1 * ((0.0 / 1.0) * (0.1 - ((3.14159 + -2.25) - 1.5))))) {
                        v10 = false
                    }
                }
            case 0:
                v8 -= (fi9 + (v8 | -2147483648))
            case 1:
                v10 = false
            }
            _ = fi9
            _ = v10
        }
        if len(v7) > 0 {
            v7[0] = v8
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        if (true && false) {
            if true {
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := GenS1{F0: 0, F1: 18446744073709551615}
                    _ = fi17
                    _ = v18
                }
                v7 = append(v7, (v8 % (((v8 + 127)) | 1)))
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int16 = 32767
                _ = fi19
                _ = v20
            }
            v21 := GenS1{F0: 0, F1: 4294967295}
            _ = v21
        } else {
            var v22 int16 = 65
            var v23 uint32 = 2147483647
            if len(v7) > 0 {
                v7[0] = 2147483647
            }
            _ = v22
            _ = v23
        }
        v7 = append(v7, v8)
        for fi24 := 0; fi24 < 2; fi24++ {
            if false {
                if ("hello" == "\n\t") {
                    var v25 int16 = 543
                    _ = v25
                } else {
                    fmt.Println(fi24, v8)
                    var v26 uint16 = 19
                    var v27 int64 = -139
                    _ = v26
                    _ = v27
                }
                fmt.Println(fi24, v8)
            }
            v7 = append(v7, (fi24 / ((4294967295) | 1)))
            switch fi24 {
            case -1:
                switch (v8 % ((v8) | 1)) {
                case 1:
                    var v28 float64 = 1e10
                    _ = v28
                case -615:
                    fmt.Println(fi24)
                case -219:
                    if false {
                        var v29 float32 = (0.0 + 1.5)
                        if ((!true) || (!(!(!(4 >= 78))))) {
                            v8 = (v8 % ((v8) | 1))
                            var v30 bool = true
                            fmt.Println(fi24, v30)
                            _ = v30
                        } else {
                            var v31 uint = 0
                            _ = v31
                        }
                        switch ("hello" + ((("foo bar" + "hello") + ("hello" + "foo bar")) + ("foo bar" + "a"))) {
                        case "hello":
                            for fi32 := 0; fi32 < 15; fi32++ {
                                var v33 uint = 808
                                _ = fi32
                                _ = v33
                            }
                        case "a":
                            fmt.Println(v8)
                        }
                        _ = v29
                    }
                default:
                    for fi34 := 0; fi34 < 11; fi34++ {
                        if false {
                            if true {
                                var v35 uint8 = 128
                                var v36 int = (fi24 % ((fi34) | 1))
                                _ = v35
                                _ = v36
                            } else {
                                v8 -= 37
                                fmt.Println(v8, fi24)
                                v8 += fi34
                                _ = len(v7)
                                v8 = v8
                            }
                            v37 := GenS1{F0: 42, F1: 42}
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = v37
                        }
                        _ = fi34
                    }
                }
            case 32767:
                fmt.Println(fi24, v8)
            case 1:
                for fi38 := 0; fi38 < 5; fi38++ {
                    v39 := GenS1{F0: -128, F1: 1}
                    var v40 uint64 = 498
                    v8 -= (^(v8))
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
            _ = fi24
        }
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        var v42 int = v8
        _ = len(v7)
        switch "foo bar" {
        case "":
            v42 = v8
        case "foo bar":
            v8 = (v42 >> ((0) & 63))
        case "hello":
            v8 -= (v8 | (v42 << ((65535) & 63)))
        }
        _ = fi41
        _ = v42
    }
    v8 += v8
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 int16 = -1
    v4 := make(map[uint32]uint32)
    for fi43 := 0; fi43 < 16; fi43++ {
        delete(v1, (fi43 - fi43))
        var v44 uint64 = 42
        _ = fi43
        _ = v44
    }
    var v45 int8 = 58
    var v46 uint16 = 331
    var v47 int8 = v45
    v45 = v47
    v48 := GenS2{F0: "", F1: true, F2: 0}
    if false {
        v46 = v46
        var v49 float64 = (-2.25 / ((1e10 / 3.14159) + 1.5))
        _ = v49
    }
    for fi50 := 0; fi50 < 1; fi50++ {
        switch 1 {
        case 127:
            v45 = v45
        case 128:
            switch (-(fi50)) {
            case 128:
                if ((fi50 / ((32767) | 1)) != (fi50 - 42)) {
                    for fi51 := 0; fi51 < 8; fi51++ {
                        if (((!(-9223372036854775808 != 1)) || (true && false)) || false) {
                            if (float64(((0.1 - (-2.25 / 0.1)) * float32(-2.25))) <= (((((1.5 + 3.14159) - -2.25) + (1.5 - 1.5)) + (0.0 * -2.25)) + 0.1)) {
                                switch (fi50 + fi51) {
                                case 16:
                                    v47 -= v45
                                default:
                                    switch "hello" {
                                    case "":
                                        v52 := []int{(fi50 & (fi50 * fi51)), 286, int(v46)}
                                        _ = v52
                                    case "\n\t":
                                        v47 += (v45 & (v47 - -41))
                                    default:
                                        switch (("foo bar" + ("" + "hello")) + "") {
                                        case "foo bar":
                                            var v53 float32 = 0.1
                                            var v54 string = (("hello" + ("\n\t" + "foo bar")) + "hello")
                                            _ = v53
                                            _ = v54
                                        default:
                                            switch fi50 {
                                            case 462:
                                                v55 := []int{fi51, (fi50 ^ fi50), -700}
                                                _ = v55
                                            case -277:
                                                switch fi50 {
                                                case 9223372036854775807:
                                                    var v56 int16 = v3
                                                    v3 = v56
                                                    if (42 <= 4294967295) {
                                                        var v57 float32 = 1.5
                                                        var v58 string = ("\n\t" + "\n\t")
                                                        v45 += v45
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    _ = v56
                                                case 32768:
                                                    if false {
                                                        if (("a" + "a") <= ("foo bar" + "a")) {
                                                            fmt.Println(v47)
                                                            v59 := GenS3{F0: -1, F1: 42}
                                                            v60 := []int{(fi51 * -394)}
                                                            _ = len(v60)
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        var v61 float64 = 0.1
                                                        _ = v61
                                                    }
                                                default:
                                                    fmt.Println(v47, v3, v46)
                                                    if (!true) {
                                                        v62 := GenS1{F0: (-(v47)), F1: 304}
                                                        _ = v62
                                                    } else {
                                                        switch fi51 {
                                                        case -2147483648:
                                                            var v63 uint64 = 2
                                                            _ = v63
                                                        }
                                                    }
                                                    var v64 uint8 = 217
                                                    _ = v64
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                            if (!true) {
                                var v65 int16 = v3
                                var v66 bool = true
                                switch "\n\t" {
                                case "":
                                    v67 := []int{}
                                    _ = v67
                                case "\n\t":
                                    v68 := GenS4{F0: (("foo bar" + "foo bar") + (("hello" + "") + "hello"))}
                                    _ = v68
                                case "a":
                                    if (false || (v3 == v3)) {
                                        var v69 uint16 = v46
                                        v69 = v69
                                        fmt.Println(v45)
                                        _ = v69
                                    }
                                }
                                _ = v65
                                _ = v66
                            } else {
                                delete(v2, 4294967295)
                            }
                            delete(v1, fi50)
                            v70 := []int{}
                            _ = v70
                        }
                        if false {
                            var v71 uint32 = 2147483647
                            fmt.Println(v71, v71)
                            v72 := GenS5{F0: 303}
                            _ = v71
                            _ = v72
                        } else {
                            v47 -= 127
                            switch (fi50 + fi50) {
                            case 853:
                                v47 = v47
                            case 710:
                                var v73 bool = false
                                _ = v73
                            case 878:
                                v74 := v1[2147483647]
                                _ = v74
                            }
                        }
                        _ = fi51
                    }
                    switch "a" {
                    case "":
                        v45 = (v47 * (v47 | 42))
                    case "hello":
                        if false {
                            v46 += (^(v46))
                        }
                    case "\n\t":
                        v45 = 1
                    }
                } else {
                    v47 = (v47 >> ((346) & 63))
                    var v75 uint = 2
                    var v76 uint64 = 701
                    _ = v75
                    _ = v76
                }
            }
        case -944:
            fmt.Println(v47)
        default:
            fmt.Println(v46, v45, v45)
        }
        v46 += (v46 / (((^(v46))) | 1))
        v77 := GenS6{F0: v46}
        _ = fi50
        _ = v77
    }
    for fi78 := 0; fi78 < 10; fi78++ {
        v79 := []int{-9223372036854775808, (fi78 - -2147483648), 2147483647, fi78}
        _ = fi78
        _ = v79
    }
    var v80 uint8 = 255
    v81 := []int{-1, 596, -1}
    switch ("" + ("hello" + "\n\t")) {
    case "\n\t":
        v46 -= (v46 * 255)
    case "hello":
        _ = len(v81)
    case "foo bar":
        fmt.Println(v80)
    default:
        v80 -= v80
    }
    if len(v81) > 0 {
        v81[0] = 128
    }
    var v82 float32 = ((0.1 * (float32(1.5) / ((3.14159 - (0.0 + 1.5)) / 1.0))) + 1e10)
    v82 = (v82 - 3.14159)
    var v83 int16 = (v3 >> ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
