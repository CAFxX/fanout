package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint64; F2 string }
type GenS3 struct { F0 int; F1 uint8; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64; F1 uint16; F2 int8 }
type GenS12 struct { F0 int32; F1 uint64; F2 uint16 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]uint32)
    var v3 int8 = 1
    v4 := make(map[string]int)
    v5 := v4["foo bar"]
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                fmt.Println(v1)
                if false {
                    var v8 int16 = 562
                    v9 := []int{v5, v5, -494}
                    if (847 != 565) {
                        v3 = (-(v3))
                        switch (^(fi6)) {
                        case -302:
                            v3 -= (^(v3))
                        case -1:
                            v10 := GenS1{F0: v3}
                            _ = v10
                        case 4294967295:
                            var v11 float32 = (1.5 - ((v1 / 0.1) * ((v1 + ((v1 * v1) - v1)) - v1)))
                            _ = v11
                        default:
                            switch (^(v5)) {
                            case 2147483647:
                                switch "foo bar" {
                                case "foo bar":
                                    var v12 float32 = (v1 + (v1 * float32((-2.25 / (-2.25 / 0.1)))))
                                    _ = v12
                                case "\n\t":
                                    fmt.Println(v3, v8)
                                }
                            case -9223372036854775808:
                                if len(v9) > 0 {
                                    v9[0] = v5
                                }
                            }
                            switch "a" {
                            case "foo bar":
                                _ = len(v9)
                            case "a":
                                var v13 bool = true
                                _ = v13
                            case "":
                                _ = len(v4)
                                var v14 int32 = 42
                                _ = v14
                            }
                        }
                        _ = len(v9)
                    }
                    v5 = (fi6 << ((4294967295) & 63))
                    _ = v8
                    _ = v9
                }
                if false {
                    if false {
                        var v15 string = ""
                        v16 := GenS2{F0: -128, F1: 18446744073709551615, F2: "hello"}
                        switch fi7 {
                        case 0:
                            var v17 bool = true
                            v18 := GenS3{F0: v5, F1: 40, F2: fi6}
                            v19 := []int{v5, v5, -9223372036854775808}
                            _ = v17
                            _ = v18
                            _ = v19
                        case 128:
                            fmt.Println(fi7)
                        case 418:
                            v3 = v3
                        default:
                            if (4294967295 != 459) {
                                var v20 uint32 = 902
                                v21 := []int{fi7, 1, (v5 - (fi7 - (v5 * (fi7 - -526))))}
                                if false {
                                    fmt.Println(v1, v20, v3)
                                    var v22 int32 = int32(v3)
                                    v20 -= 809
                                    _ = v22
                                }
                                _ = v20
                                _ = v21
                            }
                        }
                        var v23 int32 = -147
                        var v24 uint16 = 0
                        _ = v15
                        _ = v16
                        _ = v23
                        _ = v24
                    } else {
                        if (65535 == 1) {
                            var v25 float64 = float64(((((float32(0.1) / (v1 * v1)) * (3.14159 + v1)) - (((0.0 * (v1 / v1)) + (v1 / 0.1)) * 1.5)) / 1.0))
                            if (42 >= 42) {
                                v1 = float32(0.0)
                                switch (fi6 ^ v5) {
                                case 4294967295:
                                    v26 := GenS4{F0: 2}
                                    _ = v26
                                default:
                                    v5 = 9223372036854775807
                                    fmt.Println(v25, fi6)
                                    delete(v4, ((("\n\t" + "hello") + "foo bar") + (("" + "\n\t") + ("foo bar" + "foo bar"))))
                                }
                                var v27 int = fi7
                                switch (fi6 | (fi7 - fi7)) {
                                case 32768:
                                    v28 := GenS5{F0: 1}
                                    _ = v28
                                case 1:
                                    var v29 int32 = -2147483648
                                    _ = v29
                                case 42:
                                    fmt.Println(v3, fi6)
                                default:
                                    v3 = -128
                                }
                                _ = v27
                            } else {
                                v30 := []int{-9223372036854775808, (fi7 / ((v5) | 1))}
                                if (1000 <= 694) {
                                    if true {
                                        var v31 uint32 = 2
                                        v32 := GenS6{F0: ((((v31 & (v31 * v31)) > v31) || (false && (751 == uint(v31)))) && ((true || (true && false)) || (!(((1e10 + v1) / 0.1) != (v1 - (v1 / (v1 + (0.0 + v1)))))))), F1: 142, F2: 678}
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                var v33 int64 = int64(v3)
                                _ = v30
                                _ = v33
                            }
                            switch fi7 {
                            case -1:
                                fmt.Println(v5, v25)
                            case 108:
                                v34 := GenS7{F0: -9223372036854775808, F1: 42}
                                _ = v34
                            case -2147483648:
                                var v35 int32 = -1
                                _ = v35
                            }
                            _ = v25
                        }
                        v36 := v2[(fi7 | (fi6 / ((127) | 1)))]
                        _ = v36
                    }
                    v37 := GenS3{F0: -802, F1: 86, F2: (fi7 ^ fi6)}
                    _ = v37
                }
                _ = fi7
            }
            _ = fi6
        }
        fmt.Println(v5)
    } else {
        for fi38 := 0; fi38 < 1; fi38++ {
            v39 := GenS1{F0: int8(fi38)}
            v4[(("\n\t" + "foo bar") + "a")] = fi38
            _ = fi38
            _ = v39
        }
        if (-466 != -685) {
            switch "foo bar" {
            case "":
                var v40 int32 = -400
                _ = v40
            case "\n\t":
                switch (v5 * (v5 << ((1) & 63))) {
                case -2147483648:
                    for fi41 := 0; fi41 < 20; fi41++ {
                        var v42 float32 = ((0.0 + 3.14159) / v1)
                        v5 = (-(fi41))
                        for fi43 := 0; fi43 < 4; fi43++ {
                            v5 += 1
                            switch (-(fi41)) {
                            case 702:
                                var v44 float64 = float64(v1)
                                _ = v44
                            case 1:
                                v45 := GenS3{F0: (^(v5)), F1: 128, F2: (fi41 % (((fi43 + fi43)) | 1))}
                                _ = v45
                            case -555:
                                v46 := []int{}
                                _ = v46
                            }
                            _ = fi43
                        }
                        _ = fi41
                        _ = v42
                    }
                case 9223372036854775807:
                    switch 1 {
                    case -1:
                        v47 := GenS3{F0: v5, F1: 42, F2: v5}
                        _ = v47
                    case 1:
                        var v48 bool = (v1 > float32((-2.25 / ((-2.25 / 3.14159) / 3.14159))))
                        _ = v48
                    default:
                        var v49 uint16 = 32767
                        _ = v49
                    }
                case -9223372036854775808:
                    fmt.Println(v1)
                }
            }
        }
    }
    if false {
        for fi50 := 0; fi50 < 3; fi50++ {
            var v51 float32 = float32(((((((1e10 * (3.14159 / 1.0)) - 1.5) / (-2.25 / 1e10)) - ((3.14159 * 1.5) + 0.1)) / 1.5) - (-2.25 * 1e10)))
            for fi52 := 0; fi52 < 17; fi52++ {
                var v53 string = "\n\t"
                _ = fi52
                _ = v53
            }
            fmt.Println(v3)
            _ = fi50
            _ = v51
        }
    } else {
        for fi54 := 0; fi54 < 2; fi54++ {
            delete(v4, "foo bar")
            _ = fi54
        }
        v1 += v1
        v55 := []int{(v5 & v5), (v5 + -110), (v5 + (v5 / ((v5) | 1)))}
        _ = v55
    }
    v56 := GenS8{F0: (v5 - (v5 >> ((921) & 63)))}
    v57 := []int{(v5 + (-(v5))), -2147483648, (v5 & v5)}
    if (42 != 0) {
        v5 = (-(v5))
        fmt.Println(v5, v3)
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        for fi59 := 0; fi59 < 13; fi59++ {
            var v60 int = -962
            var v61 int32 = -950
            v61 -= v61
            _ = fi59
            _ = v60
            _ = v61
        }
        var v62 int32 = 376
        _ = fi58
        _ = v62
    }
    var v63 int32 = 1
    if ((-1 > (v63 / ((1) | 1))) && true) {
        fmt.Println(v63, v3)
        for fi64 := 0; fi64 < 6; fi64++ {
            for fi65 := 0; fi65 < 19; fi65++ {
                var v66 uint8 = 132
                _ = fi65
                _ = v66
            }
            switch (fi64 << ((2) & 63)) {
            case -9223372036854775808:
                _ = len(v57)
            case 2147483647:
                var v67 int16 = -486
                _ = v67
            case 273:
                fmt.Println(v63, fi64, fi64)
            default:
                if ((65535 > 42) && true) {
                    if (0.1 == 3.14159) {
                        v3 += v3
                        v68 := GenS6{F0: true, F1: 495, F2: 280}
                        _ = v68
                    }
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        _ = len(v57)
                    case "":
                        v63 -= (v63 % ((0) | 1))
                    default:
                        if len(v57) > 0 {
                            v57[0] = fi64
                        }
                    }
                } else {
                    delete(v4, "")
                    var v69 float64 = 1.5
                    v70 := GenS8{F0: fi64}
                    _ = v69
                    _ = v70
                }
            }
            var v71 uint64 = 177
            _ = fi64
            _ = v71
        }
        v3 = v3
    } else {
        for fi72 := 0; fi72 < 9; fi72++ {
            if true {
                for fi73 := 0; fi73 < 18; fi73++ {
                    switch ("foo bar" + ("foo bar" + "hello")) {
                    case "\n\t":
                        switch 0 {
                        case 796:
                            var v74 int = v5
                            _ = v74
                        case -1:
                            if true {
                                if true {
                                    fmt.Println(fi73)
                                    v5 = fi73
                                    var v75 float64 = (3.14159 / float64(v1))
                                    _ = v75
                                }
                            }
                        default:
                            switch "foo bar" {
                            case "a":
                                v76 := GenS9{F0: ((false && false) || (true && true))}
                                _ = v76
                            case "foo bar":
                                if (2 < 273) {
                                    v5 = 32768
                                    v5 += 42
                                }
                            }
                        }
                        var v77 uint32 = 962
                        _ = v77
                    }
                    if true {
                        if (!true) {
                            if false {
                                v78 := GenS10{F0: 2}
                                if true {
                                    var v79 bool = ((713 >= 514) || true)
                                    var v80 uint16 = 80
                                    _ = v79
                                    _ = v80
                                }
                                var v81 float64 = ((float64(1.5) + (float64(float32((float64(float32(0.1)) + (3.14159 * -2.25)))) / (1e10 + 3.14159))) * 1e10)
                                _ = v78
                                _ = v81
                            } else {
                                v1 = v1
                                var v82 int = v5
                                if false {
                                    v4["\n\t"] = 128
                                    var v83 int16 = 1
                                    v84 := GenS4{F0: 2}
                                    _ = v83
                                    _ = v84
                                } else {
                                    var v85 uint16 = 32767
                                    _ = v85
                                }
                                _ = v82
                            }
                            if true {
                                fmt.Println(v63, fi72, fi72)
                                v63 = v63
                                _ = len(v2)
                            } else {
                                fmt.Println(v1, v5, v3)
                            }
                            v63 = v63
                            var v86 int64 = 9223372036854775807
                            _ = v86
                        } else {
                            var v87 float64 = float64((v1 * (v1 / 1e10)))
                            switch 851 {
                            case 128:
                                if true {
                                    _ = len(v2)
                                    if ((("hello" > "a") || (!false)) || ((127 >= 85) && false)) {
                                        v87 += float64(v1)
                                        var v88 int16 = -32768
                                        v87 += float64((v1 / (3.14159 + 3.14159)))
                                        _ = v88
                                    } else {
                                        v5 -= 32768
                                        if false {
                                            if true {
                                                v1 = (((((((v1 + -2.25) / v1) / v1) + (v1 / v1)) * -2.25) - ((1e10 / 3.14159) / v1)) - (((float32(((float64(float32(v87)) + 0.0) / 0.1)) - v1) * (float32(v87) + v1)) * -2.25))
                                                var v89 uint64 = 889
                                                _ = v89
                                            }
                                            if len(v57) > 0 {
                                                v57[0] = (v5 ^ -2147483648)
                                            }
                                        } else {
                                            v3 -= v3
                                            v90 := GenS1{F0: (v3 * v3)}
                                            if ("" != "hello") {
                                                if true {
                                                    v5 = fi72
                                                    delete(v2, fi72)
                                                    _ = len(v57)
                                                    var v91 uint16 = 1
                                                    _ = v91
                                                } else {
                                                    switch "a" {
                                                    case "hello":
                                                        switch fi72 {
                                                        case -990:
                                                            if false {
                                                                if false {
                                                                    v87 += v87
                                                                }
                                                                v92 := GenS10{F0: 489}
                                                                _ = v92
                                                            } else {
                                                                v3 = v3
                                                                switch "a" {
                                                                case "":
                                                                    if false {
                                                                        fmt.Println(v63, fi73)
                                                                        v57 = append(v57, 42)
                                                                        var v93 uint = 65535
                                                                        v57 = append(v57, (fi72 / ((-492) | 1)))
                                                                        _ = v93
                                                                    } else {
                                                                        var v94 uint16 = 327
                                                                        v3 += v3
                                                                        v1 = v1
                                                                        _ = v94
                                                                    }
                                                                default:
                                                                    var v95 int8 = -1
                                                                    _ = v95
                                                                }
                                                                v63 -= 849
                                                            }
                                                        case 4294967295:
                                                            var v96 string = (("hello" + "hello") + ("a" + "foo bar"))
                                                            var v97 uint64 = 0
                                                            v98 := v4[(v96 + v96)]
                                                            v5 = (fi72 % ((fi72) | 1))
                                                            _ = v96
                                                            _ = v97
                                                            _ = v98
                                                        default:
                                                            v3 -= v3
                                                        }
                                                    }
                                                    if (true && (!(!(false && false)))) {
                                                        v63 -= v63
                                                        var v99 float64 = 3.14159
                                                        _ = v99
                                                    } else {
                                                        v100 := GenS11{F0: 108, F1: 32767, F2: (-(v3))}
                                                        _ = v100
                                                    }
                                                    v3 = v3
                                                }
                                            } else {
                                                v101 := v4[("hello" + "")]
                                                _ = v101
                                            }
                                            _ = v90
                                        }
                                    }
                                }
                            case 2147483647:
                                var v102 uint64 = 4294967295
                                _ = v102
                            default:
                                if (true || (v3 > -1)) {
                                    if (1 == 0) {
                                        _ = len(v57)
                                        fmt.Println(fi72)
                                        var v103 float64 = v87
                                        _ = v103
                                    }
                                    switch "a" {
                                    case "hello":
                                        v5 += (v5 | v5)
                                    case "\n\t":
                                        var v104 bool = (true || true)
                                        _ = v104
                                    }
                                }
                                var v105 int32 = (v63 << ((42) & 63))
                                _ = v105
                            }
                            v63 += 1
                            v106 := GenS12{F0: v63, F1: 42, F2: 133}
                            _ = v106
                            _ = v87
                        }
                    }
                    v5 += fi72
                    _ = fi73
                }
            }
            v107 := GenS2{F0: v3, F1: 4294967295, F2: (("\n\t" + ("\n\t" + "\n\t")) + ("hello" + ("\n\t" + "foo bar")))}
            _ = fi72
            _ = v107
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v63
}
