package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint8; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int8; F2 string }
type GenS9 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -1
    var v3 uint32 = 1
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{(fi5 % ((fi5) | 1)), fi5, fi5}
        switch "" {
        case "hello":
            switch (-(fi5)) {
            case 4294967295:
                var v7 string = "hello"
                _ = v7
            default:
                var v8 int = fi5
                _ = v8
            }
        case "a":
            var v9 string = "\n\t"
            _ = v9
        case "\n\t":
            var v10 string = "foo bar"
            _ = v10
        }
        switch "hello" {
        case "\n\t":
            if true {
                if false {
                    var v11 uint16 = 2
                    if (-1 >= -9223372036854775808) {
                        v12 := GenS1{F0: (fi5 >> ((126) & 63)), F1: -54, F2: false}
                        _ = v12
                    } else {
                        var v13 uint = 394
                        _ = v13
                    }
                    v3 -= 2147483648
                    _ = v11
                } else {
                    if false {
                        v2 = (^(v2))
                        var v14 int8 = v2
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            if (18446744073709551615 < (v4 / ((0) | 1))) {
                                v16 := GenS2{F0: 1, F1: -987}
                                v17 := GenS3{F0: 0, F1: v4}
                                var v18 string = "hello"
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            _ = len(v1)
                            var v19 bool = (44 == 0)
                            _ = fi15
                            _ = v19
                        }
                        var v20 int64 = -1
                        _ = v20
                    }
                    v21 := v1[1]
                    _ = v21
                }
                _ = len(v1)
            }
        case "hello":
            switch (fi5 | fi5) {
            case 1:
                _ = len(v1)
            }
        default:
            for fi22 := 0; fi22 < 5; fi22++ {
                v2 += v2
                switch "" {
                case "foo bar":
                    v2 = 127
                case "a":
                    switch (^(fi22)) {
                    case 1:
                        v23 := GenS4{F0: true, F1: 375}
                        _ = v23
                    case -1:
                        var v24 float64 = (-2.25 / (-2.25 - float64(0.0)))
                        _ = v24
                    }
                }
                switch (("a" + "hello") + "hello") {
                case "foo bar":
                    v25 := GenS5{F0: false}
                    _ = v25
                case "\n\t":
                    var v26 int32 = -249
                    _ = v26
                default:
                    if len(v6) > 0 {
                        v6[0] = (fi22 & 127)
                    }
                }
                _ = fi22
            }
        }
        _ = fi5
        _ = v6
    }
    v27 := GenS4{F0: (2 == 1), F1: -1}
    if (true && true) {
        var v28 string = ""
        var v29 int8 = -128
        _ = v28
        _ = v29
    } else {
        if (float32((1e10 - 0.0)) > 1.5) {
            for fi30 := 0; fi30 < 15; fi30++ {
                fmt.Println(v2, fi30)
                var v31 uint32 = v3
                v32 := GenS6{F0: ("" + "foo bar"), F1: 174, F2: "foo bar"}
                _ = fi30
                _ = v31
                _ = v32
            }
            v2 -= 42
            v2 = v2
        }
        var v33 uint32 = 2147483647
        _ = v33
    }
    switch "\n\t" {
    case "foo bar":
        var v34 uint64 = v4
        _ = v34
    }
    var v35 float32 = (1.5 - ((1.5 / float32(1e10)) / 1.0))
    switch 1 {
    case 4294967295:
        if (-2.25 < (-2.25 / v35)) {
            v2 -= v2
            for fi36 := 0; fi36 < 4; fi36++ {
                v35 = ((v35 * (float32(3.14159) / (v35 + v35))) - v35)
                if (v2 < (v2 | 50)) {
                    v37 := GenS6{F0: "foo bar", F1: 127, F2: "hello"}
                    fmt.Println(v4)
                    v4 = 0
                    _ = v37
                }
                _ = fi36
            }
        }
    default:
        v38 := GenS5{F0: false}
        _ = v38
    }
    switch "a" {
    case "a":
        fmt.Println(v3, v35)
    case "":
        v3 = (^(v3))
    case "\n\t":
        var v39 uint64 = 18446744073709551615
        _ = v39
    default:
        var v40 float64 = 0.1
        _ = v40
    }
    if false {
        var v41 bool = false
        if (0 <= (v2 - (v2 - v2))) {
            v42 := GenS7{F0: -1}
            fmt.Println(v41)
            if (-18 >= (v2 + -128)) {
                v41 = (1 <= 1000)
                if (v41 && v41) {
                    v35 = (float32((((0.0 + 0.0) * 1e10) - 0.1)) / -2.25)
                }
            } else {
                v41 = ((v3 + (v3 / ((v3) | 1))) <= v3)
                v43 := v1[32768]
                switch "a" {
                case "a":
                    v35 = 1.5
                default:
                    switch 866 {
                    case -357:
                        fmt.Println(v3)
                    case 42:
                        var v44 int64 = -767
                        _ = v44
                    }
                }
                _ = v43
            }
            _ = v42
        }
        fmt.Println(v3, v4, v2)
        _ = v41
    } else {
        v3 = (^(v3))
        fmt.Println(v2, v3, v2)
    }
    v45 := GenS8{F0: 1, F1: 127, F2: ("\n\t" + "foo bar")}
    switch ((("hello" + "") + "") + "a") {
    case "\n\t":
        for fi46 := 0; fi46 < 20; fi46++ {
            v47 := []int{-9223372036854775808}
            _ = fi46
            _ = v47
        }
    case "foo bar":
        var v48 uint64 = 2
        _ = v48
    default:
        var v49 int = 32767
        _ = v49
    }
    v2 += (v2 & 64)
    switch "" {
    case "hello":
        switch -1 {
        case 480:
            for fi50 := 0; fi50 < 5; fi50++ {
                v3 -= 2147483648
                v51 := GenS3{F0: uint8(v4), F1: (v4 | v4)}
                _ = fi50
                _ = v51
            }
        case -492:
            if false {
                var v52 bool = ((-308 == -9223372036854775808) || true)
                if v52 {
                    switch -1 {
                    case 128:
                        if v52 {
                            if v52 {
                                for fi53 := 0; fi53 < 8; fi53++ {
                                    v52 = v52
                                    _ = fi53
                                }
                                if true {
                                    for fi54 := 0; fi54 < 10; fi54++ {
                                        var v55 bool = (1 > 567)
                                        _ = fi54
                                        _ = v55
                                    }
                                    v2 -= v2
                                }
                                var v56 float32 = v35
                                _ = v56
                            } else {
                                v35 = v35
                            }
                            if (!v52) {
                                if true {
                                    var v57 int16 = -594
                                    switch 32767 {
                                    case -9223372036854775808:
                                        v58 := []int{32767, -474, 4294967295, -375}
                                        _ = v58
                                    }
                                    _ = v57
                                }
                                var v59 uint = 572
                                for fi60 := 0; fi60 < 11; fi60++ {
                                    var v61 uint16 = 923
                                    if false {
                                        v2 = -1
                                    }
                                    for fi62 := 0; fi62 < 10; fi62++ {
                                        switch ("a" + "a") {
                                        case "hello":
                                            v63 := []int{int(v61)}
                                            _ = v63
                                        case "a":
                                            v64 := []int{(fi62 + (fi62 >> (((v59 & v59)) & 63)))}
                                            _ = v64
                                        case "\n\t":
                                            v2 -= (-(v2))
                                        default:
                                            if (uint16(v3) >= (v61 >> ((uint(v61)) & 63))) {
                                                var v65 uint8 = 127
                                                if v52 {
                                                    v66 := []int{2147483647}
                                                    v1[fi62] = v3
                                                    switch fi60 {
                                                    case 32768:
                                                        switch ("" + "") {
                                                        case "\n\t":
                                                            v59 += v59
                                                        case "":
                                                            var v67 bool = (v52 && (-226 <= int16(v61)))
                                                            _ = v67
                                                        case "hello":
                                                            v68 := v1[fi62]
                                                            _ = v68
                                                        }
                                                    default:
                                                        var v69 uint64 = (v4 | (^(v4)))
                                                        _ = v69
                                                    }
                                                    _ = v66
                                                }
                                                _ = v65
                                            }
                                        }
                                        v70 := []int{127, fi62, (^(fi62)), (-(fi62))}
                                        var v71 uint32 = (v3 - 4294967295)
                                        _ = fi62
                                        _ = v70
                                        _ = v71
                                    }
                                    _ = fi60
                                    _ = v61
                                }
                                _ = v59
                            } else {
                                if v52 {
                                    var v72 uint64 = v4
                                    _ = v72
                                } else {
                                    for fi73 := 0; fi73 < 3; fi73++ {
                                        v52 = (!true)
                                        v52 = v52
                                        _ = fi73
                                    }
                                    var v74 float32 = 0.0
                                    _ = v74
                                }
                                fmt.Println(v3, v3, v4)
                            }
                            fmt.Println(v35)
                        } else {
                            if (!v52) {
                                v75 := []int{}
                                _ = v75
                            }
                            fmt.Println(v35, v52, v35)
                        }
                    case -1:
                        if v52 {
                            switch "foo bar" {
                            case "\n\t":
                                var v76 float32 = (3.14159 * 0.0)
                                _ = v76
                            case "foo bar":
                                for fi77 := 0; fi77 < 19; fi77++ {
                                    var v78 uint32 = 65535
                                    v78 = 506
                                    _ = fi77
                                    _ = v78
                                }
                            }
                            var v79 bool = (v52 || v52)
                            _ = v79
                        }
                    case 127:
                        v52 = v52
                    }
                    v3 += 0
                    for fi80 := 0; fi80 < 12; fi80++ {
                        var v81 int16 = -32768
                        fmt.Println(v2, v81, v4)
                        _ = fi80
                        _ = v81
                    }
                } else {
                    for fi82 := 0; fi82 < 19; fi82++ {
                        if (((fi82 % ((-954) | 1)) > (fi82 / ((978) | 1))) && v52) {
                            switch ("foo bar" + "foo bar") {
                            case "":
                                fmt.Println(v3, fi82)
                            default:
                                v4 -= v4
                            }
                            var v83 int32 = -1
                            _ = v83
                        }
                        switch 345 {
                        case 120:
                            v35 = (v35 + 0.0)
                        }
                        v2 = (v2 + (v2 % (((v2 ^ 100)) | 1)))
                        _ = fi82
                    }
                    var v84 int32 = 0
                    v52 = v52
                    _ = v84
                }
                _ = v52
            } else {
                v85 := GenS7{F0: 127}
                _ = v85
            }
        case -221:
            var v86 int = 32767
            _ = v86
        default:
            for fi87 := 0; fi87 < 9; fi87++ {
                var v88 bool = (65535 == 438)
                _ = fi87
                _ = v88
            }
        }
    case "":
        var v89 int16 = -74
        _ = v89
    case "foo bar":
        fmt.Println(v35, v2, v4)
    }
    for fi90 := 0; fi90 < 6; fi90++ {
        if false {
            v2 += v2
            var v91 float32 = (v35 * float32((0.1 / 1.5)))
            v92 := GenS8{F0: -388, F1: v2, F2: ""}
            _ = v91
            _ = v92
        }
        _ = fi90
    }
    for fi93 := 0; fi93 < 18; fi93++ {
        v94 := []int{}
        if true {
            switch "hello" {
            case "foo bar":
                v95 := GenS1{F0: 42, F1: -257, F2: ((849 <= 32767) && true)}
                _ = v95
            case "hello":
                v4 -= 767
            case "":
                v3 = (v3 / (((^(v3))) | 1))
            default:
                fmt.Println(v2)
            }
            v3 = 65535
            fmt.Println(v35)
        }
        _ = fi93
        _ = v94
    }
    v3 = v3
    if (-9223372036854775808 < 42) {
        v96 := GenS9{F0: 181, F1: -1}
        v97 := GenS4{F0: (!(255 == 2)), F1: 1}
        fmt.Println(v2)
        _ = v96
        _ = v97
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v45
}
