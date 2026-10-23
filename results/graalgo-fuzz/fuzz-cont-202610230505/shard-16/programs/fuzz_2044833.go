package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint32; F1 string; F2 uint64 }
type GenS4 struct { F0 string; F1 string; F2 int }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 int; F2 uint32 }
type GenS9 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 string = (("hello" + ("" + "")) + "foo bar")
    var v4 int = 127
    switch v3 {
    case "a":
        switch v4 {
        case 901:
            v3 = v3
        default:
            v4 = v4
        }
    case "hello":
        var v5 float64 = float64(3.14159)
        _ = v5
    default:
        var v6 int64 = 42
        _ = v6
    }
    delete(v1, v4)
    v3 = v3
    if (!(0 > 42)) {
        v4 = (^(v4))
        v7 := []int{198, v4, v4}
        v4 = v4
        _ = v7
    } else {
        if (127 <= 255) {
            for fi8 := 0; fi8 < 10; fi8++ {
                v4 = v4
                switch v4 {
                case -665:
                    v9 := GenS1{F0: 0}
                    _ = v9
                case 4294967295:
                    v3 = v3
                default:
                    v10 := []int{v4, v4, v4}
                    _ = v10
                }
                fmt.Println(fi8)
                _ = fi8
            }
            for fi11 := 0; fi11 < 18; fi11++ {
                v4 += fi11
                _ = fi11
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                switch v4 {
                case 32767:
                    switch v3 {
                    case "foo bar":
                        v13 := GenS2{F0: -1, F1: 9223372036854775807}
                        _ = v13
                    }
                case 519:
                    switch v3 {
                    case "a":
                        if false {
                            var v14 int32 = 42
                            var v15 int8 = int8(v14)
                            _ = v14
                            _ = v15
                        }
                    case "hello":
                        var v16 int64 = 0
                        _ = v16
                    }
                default:
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := v1[fi17]
                        v4 = (v4 % ((fi12) | 1))
                        if (-1 == 42) {
                            var v19 uint16 = 255
                            switch 755 {
                            case 32767:
                                fmt.Println(fi17, fi17, v3)
                            case -184:
                                if (0 < 1000) {
                                    v3 = v3
                                    if (42 != 100) {
                                        switch (("" + v3) + (v3 + "\n\t")) {
                                        case "\n\t":
                                            fmt.Println(v19, v3, v4)
                                        }
                                        var v20 uint64 = 167
                                        var v21 uint64 = (^(v20))
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 uint16 = (v19 % (((v19 / ((v19) | 1))) | 1))
                                    switch ("" + ((v3 + v3) + v3)) {
                                    case "\n\t":
                                        var v23 float64 = ((1.5 * (-2.25 + -2.25)) - (-2.25 + (1.5 * 3.14159)))
                                        _ = v23
                                    }
                                    _ = v22
                                }
                            default:
                                v19 = 255
                            }
                            var v24 bool = true
                            fmt.Println(v3, v18, v24)
                            _ = v19
                            _ = v24
                        } else {
                            if false {
                                v4 += -2147483648
                            }
                            v25 := []int{(fi17 << ((386) & 63))}
                            v1[-1] = -113
                            _ = v25
                        }
                        _ = fi17
                        _ = v18
                    }
                }
                for fi26 := 0; fi26 < 16; fi26++ {
                    fmt.Println(v4, v3)
                    switch ((v3 + "hello") + "foo bar") {
                    case "":
                        fmt.Println(fi12, fi12, v4)
                    case "a":
                        var v27 float64 = 1.5
                        _ = v27
                    case "hello":
                        if false {
                            if (0 <= 1) {
                                if (-128 != 42) {
                                    if false {
                                        var v28 uint = 599
                                        v29 := []int{(fi12 % ((-2147483648) | 1)), (-(fi26))}
                                        fmt.Println(v4, fi12)
                                        _ = v28
                                        _ = v29
                                    } else {
                                        v3 = ""
                                        v30 := []int{fi26, v4, 1}
                                        v4 = (fi26 / (((v4 | (fi12 * fi26))) | 1))
                                        _ = v30
                                    }
                                }
                                var v31 int16 = 42
                                fmt.Println(v3, v3)
                                if ((1.5 < ((-2.25 - (((0.0 + 0.0) + 3.14159) / 0.1)) * ((3.14159 * 0.1) * 0.0))) || (-32768 != (v31 | (-(v31))))) {
                                    fmt.Println(fi26, v3)
                                } else {
                                    var v32 uint16 = 255
                                    delete(v2, fi26)
                                    _ = v32
                                }
                                _ = v31
                            } else {
                                v3 = "hello"
                                switch 117 {
                                case 127:
                                    switch ("foo bar" + ((v3 + v3) + ((v3 + ((v3 + v3) + ("\n\t" + "hello"))) + v3))) {
                                    case "foo bar":
                                        _ = len(v1)
                                    case "a":
                                        v4 -= (-(fi26))
                                    case "":
                                        v33 := GenS1{F0: 2}
                                        _ = v33
                                    default:
                                        var v34 uint64 = 330
                                        _ = v34
                                    }
                                case 1:
                                    v35 := GenS1{F0: 418}
                                    _ = v35
                                }
                            }
                            v36 := GenS2{F0: -1, F1: 1}
                            _ = v36
                        } else {
                            switch v4 {
                            case -2147483648:
                                fmt.Println(fi12)
                            case 512:
                                v4 -= v4
                            case 603:
                                v37 := []int{(fi26 + 0)}
                                _ = v37
                            default:
                                if (533 != 18446744073709551615) {
                                    v38 := GenS2{F0: -1, F1: -967}
                                    v3 = (v3 + ("hello" + "hello"))
                                    var v39 uint64 = 622
                                    _ = v38
                                    _ = v39
                                } else {
                                    if ((10 >= 127) || true) {
                                        v4 -= v4
                                        fmt.Println(v3, fi26, v4)
                                        fmt.Println(v3, fi26)
                                    }
                                    v4 = (fi12 ^ (fi26 + v4))
                                }
                            }
                            switch (fi26 - 42) {
                            case 32767:
                                delete(v1, 42)
                            case 381:
                                if (18446744073709551615 > 2) {
                                    v3 = (v3 + (v3 + (v3 + "hello")))
                                    switch v3 {
                                    case "hello":
                                        switch (fi12 >> ((18446744073709551615) & 63)) {
                                        case -1:
                                            switch (v4 ^ (v4 & 190)) {
                                            case 9223372036854775807:
                                                var v40 int8 = -1
                                                _ = v40
                                            default:
                                                switch (fi12 & (-(v4))) {
                                                case -355:
                                                    v4 += (v4 | 2147483647)
                                                case 127:
                                                    v4 = fi12
                                                case -9223372036854775808:
                                                    v41 := GenS2{F0: -1, F1: 861}
                                                    _ = v41
                                                default:
                                                    v42 := GenS3{F0: 42, F1: "foo bar", F2: 42}
                                                    _ = v42
                                                }
                                            }
                                        case 0:
                                            v43 := GenS1{F0: 4294967295}
                                            _ = v43
                                        case -9223372036854775808:
                                            var v44 int8 = 60
                                            _ = v44
                                        default:
                                            if (!(true || true)) {
                                                if (0 != 0) {
                                                    var v45 int = fi26
                                                    var v46 int = (fi26 | 1)
                                                    v3 = ("hello" + ("hello" + v3))
                                                    v47 := GenS4{F0: "a", F1: (v3 + v3), F2: v46}
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                }
                                                if false {
                                                    delete(v1, v4)
                                                    switch fi26 {
                                                    case 127:
                                                        v48 := GenS1{F0: 2}
                                                        var v49 float32 = 0.0
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    fmt.Println(v4, fi12, fi26)
                                                }
                                                switch "" {
                                                case "a":
                                                    var v50 uint64 = 18446744073709551615
                                                    _ = v50
                                                default:
                                                    v4 = fi12
                                                    if false {
                                                        var v51 uint = 74
                                                        var v52 uint8 = 140
                                                        switch (((((v3 + v3) + v3) + "") + v3) + "a") {
                                                        case "foo bar":
                                                            var v53 int8 = 100
                                                            _ = v53
                                                        case "a":
                                                            v54 := []int{1}
                                                            _ = v54
                                                        }
                                                        if ((true || ((v3 + "") >= (v3 + ""))) && (!false)) {
                                                            var v55 uint32 = 1
                                                            v56 := v1[(fi26 - v4)]
                                                            v4 = -1
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            fmt.Println(v51, v52)
                                                        }
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                }
                                            }
                                            switch fi26 {
                                            case 9223372036854775807:
                                                v57 := GenS2{F0: -1, F1: -9223372036854775808}
                                                var v58 uint8 = 0
                                                if true {
                                                    v59 := []int{}
                                                    delete(v2, 683)
                                                    switch v3 {
                                                    case "\n\t":
                                                        v58 = v58
                                                    }
                                                    _ = v59
                                                }
                                                fmt.Println(v3)
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                        switch (v3 + (v3 + "hello")) {
                                        case "a":
                                            var v60 bool = (!(((((1.5 / (-2.25 - float64(0.1))) / (3.14159 * 3.14159)) / float64((1.5 + -2.25))) - (((0.1 * -2.25) - 3.14159) + 3.14159)) >= (float64(float32(1e10)) / 1e10)))
                                            _ = v60
                                        case "":
                                            v61 := GenS5{F0: 1, F1: 1, F2: 266}
                                            _ = v61
                                        case "\n\t":
                                            v62 := GenS6{F0: v3}
                                            _ = v62
                                        default:
                                            v63 := GenS1{F0: 2}
                                            _ = v63
                                        }
                                    case "\n\t":
                                        v64 := []int{}
                                        _ = v64
                                    default:
                                        v65 := GenS2{F0: 1000, F1: 214}
                                        _ = v65
                                    }
                                }
                            default:
                                switch (((("a" + v3) + ("hello" + v3)) + (v3 + "foo bar")) + (v3 + v3)) {
                                case "\n\t":
                                    if (0 > 255) {
                                        v66 := GenS2{F0: 0, F1: 980}
                                        v67 := []int{v4, v4}
                                        _ = len(v2)
                                        _ = v66
                                        _ = v67
                                    }
                                }
                            }
                        }
                    default:
                        switch fi12 {
                        case 759:
                            v68 := v1[fi26]
                            _ = v68
                        case 568:
                            var v69 int32 = -509
                            _ = v69
                        case -1:
                            v70 := GenS7{F0: 2147483647}
                            _ = v70
                        }
                    }
                    _ = fi26
                }
                _ = fi12
            }
        } else {
            v3 = ((v3 + v3) + "a")
            if (float64(-2.25) < (float64(float32(((1.5 - ((1.5 + float64((0.0 * 1e10))) * 0.1)) * 1e10))) + 0.1)) {
                var v71 uint16 = 0
                _ = v71
            } else {
                v3 = "hello"
                for fi72 := 0; fi72 < 5; fi72++ {
                    if false {
                        v4 -= (v4 & (v4 - 42))
                    } else {
                        v73 := []int{}
                        _ = v73
                    }
                    _ = fi72
                }
            }
        }
        switch v4 {
        case -125:
            switch v3 {
            case "foo bar":
                fmt.Println(v4)
            }
        case 2147483647:
            switch -1 {
            case -598:
                for fi74 := 0; fi74 < 7; fi74++ {
                    v75 := v1[(-(v4))]
                    v75 -= v4
                    _ = fi74
                    _ = v75
                }
            }
        }
    }
    var v76 uint8 = 2
    fmt.Println(v76)
    var v77 uint8 = v76
    var v78 float32 = (float32(1e10) - 0.0)
    v4 = (v4 & (v4 & (v4 & v4)))
    if true {
        v3 = v3
        v79 := []int{4294967295, v4}
        v80 := GenS1{F0: 78}
        _ = v79
        _ = v80
    }
    var v81 uint32 = 2
    switch v4 {
    case 128:
        var v82 int8 = -128
        _ = v82
    case 32768:
        delete(v2, v4)
    case 9223372036854775807:
        for fi83 := 0; fi83 < 17; fi83++ {
            var v84 float64 = -2.25
            for fi85 := 0; fi85 < 14; fi85++ {
                var v86 int32 = 42
                if false {
                    var v87 int16 = 1000
                    v88 := []int{}
                    v89 := GenS8{F0: 2, F1: fi83, F2: v81}
                    _ = v87
                    _ = v88
                    _ = v89
                } else {
                    v3 = (v3 + (v3 + v3))
                }
                _ = fi85
                _ = v86
            }
            _ = fi83
            _ = v84
        }
    default:
        v90 := GenS9{F0: 1000}
        _ = v90
    }
    switch v3 {
    case "foo bar":
        v91 := []int{v4, v4, int(v81), int(v81)}
        _ = v91
    case "\n\t":
        switch 32768 {
        case 1:
            var v92 uint64 = 880
            _ = v92
        case 4294967295:
            fmt.Println(v3, v78, v3)
        default:
            if (42 == -186) {
                fmt.Println(v81)
            }
        }
    case "hello":
        for fi93 := 0; fi93 < 3; fi93++ {
            var v94 uint = 18446744073709551615
            v77 += v76
            _ = fi93
            _ = v94
        }
    }
    v95 := GenS8{F0: 701, F1: (v4 / (((v4 - (v4 / (((v4 + v4)) | 1)))) | 1)), F2: v81}
    v96 := []int{(v4 + (v4 * -2147483648)), (v4 * (-(v4))), (v4 - (v4 / ((v4) | 1))), v4}
    v77 = (^(v76))
    v76 += 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    fmt.Println("slicelen v96:", len(v96), v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v76
    _ = v77
    _ = v78
    _ = v81
    _ = v95
    _ = v96
}
