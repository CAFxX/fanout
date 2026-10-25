package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    var v3 int = 32768
    var v4 int8 = -128
    v5 := make(map[int]bool)
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 += fi7
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 float64 = float64(((3.14159 * v2) / v2))
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 bool = (15 > v4)
    v11 := GenS1{F0: (v6 / ((v6) | 1)), F1: 42}
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v3, v2)
        if true {
            var v13 float32 = float32(0.0)
            for fi14 := 0; fi14 < 4; fi14++ {
                if true {
                    v15 := GenS2{F0: -153}
                    fmt.Println(v10)
                    v16 := v5[fi14]
                    _ = v15
                    _ = v16
                }
                var v17 int32 = -1
                v3 -= fi12
                _ = fi14
                _ = v17
            }
            v18 := GenS2{F0: -1}
            _ = v13
            _ = v18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            if v10 {
                fmt.Println(v4, v2)
                v4 += 100
                var v20 uint16 = 448
                var v21 uint8 = (^(v6))
                _ = v20
                _ = v21
            } else {
                var v22 int8 = (v4 % (((^(v4))) | 1))
                if v10 {
                    var v23 uint32 = 0
                    v22 = v22
                    _ = v23
                }
                _ = v22
            }
            v6 = v6
            switch ("foo bar" + ("a" + "")) {
            case "\n\t":
                if false {
                    _ = len(v1)
                    if true {
                        v3 = (v3 << ((42) & 63))
                        v5[fi12] = v10
                        fmt.Println(fi12)
                    } else {
                        if false {
                            if v10 {
                                fmt.Println(fi12, v3)
                                var v24 uint8 = 0
                                _ = v24
                            } else {
                                if (0 < -9223372036854775808) {
                                    var v25 bool = true
                                    _ = v25
                                } else {
                                    v6 += (^(v6))
                                    if true {
                                        v3 = (v3 % ((fi12) | 1))
                                        if (-4 == 1000) {
                                            var v26 float32 = (v2 + float32((-2.25 / 3.14159)))
                                            if v10 {
                                                var v27 int16 = -785
                                                v28 := GenS1{F0: 2, F1: 255}
                                                _ = v27
                                                _ = v28
                                            }
                                            v10 = true
                                            v2 -= v2
                                            _ = v26
                                        } else {
                                            if v10 {
                                                v29 := v1[42]
                                                _ = v29
                                            }
                                            switch (v3 >> ((1) & 63)) {
                                            case 2147483647:
                                                var v30 int8 = 42
                                                _ = v30
                                            }
                                        }
                                        fmt.Println(v4, v6, fi19)
                                    }
                                    v10 = false
                                }
                            }
                        }
                        switch fi19 {
                        case 32768:
                            v31 := []int{(v3 * (fi12 * (v3 / ((1) | 1))))}
                            _ = v31
                        case 2147483647:
                            v32 := GenS2{F0: 1}
                            _ = v32
                        case 5:
                            v33 := []int{}
                            _ = v33
                        default:
                            v34 := GenS2{F0: 669}
                            _ = v34
                        }
                    }
                    v2 += (1e10 - (((1e10 + 3.14159) + v2) + (0.0 * float32(3.14159))))
                } else {
                    switch (^(fi12)) {
                    case -924:
                        if (!(-2.25 == 1e10)) {
                            v35 := GenS2{F0: 1}
                            _ = v35
                        }
                    case -9223372036854775808:
                        v10 = false
                    case 2147483647:
                        v1[(v3 - (v3 & -725))] = "\n\t"
                    }
                    var v36 float32 = 1e10
                    _ = v36
                }
            default:
                fmt.Println(fi19, v3)
            }
            _ = fi19
        }
        _ = fi12
    }
    v37 := GenS1{F0: v6, F1: 255}
    v38 := GenS3{F0: -9223372036854775808, F1: 127, F2: true}
    if (255 <= 137) {
        v39 := []int{}
        _ = v39
    }
    var v40 string = (("hello" + (("a" + "\n\t") + "foo bar")) + "foo bar")
    v41 := v5[v3]
    v42 := GenS3{F0: v3, F1: (v3 * 1), F2: (1 < -966)}
    switch (v40 + ("\n\t" + "")) {
    case "a":
        var v43 bool = v10
        _ = v43
    case "\n\t":
        switch (^(v3)) {
        case 0:
            v44 := GenS2{F0: 0}
            _ = v44
        case 128:
            if v41 {
                if (1000 <= 1) {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        var v46 int = v3
                        _ = fi45
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 7; fi47++ {
                        for fi48 := 0; fi48 < 15; fi48++ {
                            switch (^(v3)) {
                            case 538:
                                if true {
                                    v4 += -1
                                    v40 = v40
                                } else {
                                    v49 := GenS4{F0: "hello"}
                                    fmt.Println(fi48, v40, v6)
                                    _ = v49
                                }
                            case 9223372036854775807:
                                fmt.Println(v6, v4, v4)
                                var v50 int32 = 372
                                v2 += (0.1 - ((v2 + (0.0 / 1.5)) - v2))
                                _ = v50
                            case 773:
                                var v51 int8 = (-(v4))
                                _ = v51
                            }
                            var v52 uint16 = 0
                            _ = fi48
                            _ = v52
                        }
                        switch v40 {
                        case "\n\t":
                            switch v40 {
                            case "a":
                                v2 += v2
                            case "\n\t":
                                v53 := GenS3{F0: 0, F1: v3, F2: v41}
                                _ = v53
                            case "":
                                for fi54 := 0; fi54 < 18; fi54++ {
                                    fmt.Println(v40, v40, v10)
                                    _ = fi54
                                }
                            }
                        case "a":
                            fmt.Println(v3, v3)
                        }
                        _ = fi47
                    }
                    v55 := GenS5{F0: 119, F1: 65535, F2: 4294967295}
                    _ = v55
                } else {
                    v56 := GenS2{F0: 1}
                    fmt.Println(v10)
                    if v10 {
                        v4 = (v4 ^ (v4 & v4))
                        v1[(^(v3))] = v40
                    } else {
                        switch (v3 << ((1) & 63)) {
                        case 710:
                            fmt.Println(v3, v40)
                        case 22:
                            var v57 float64 = 1e10
                            _ = v57
                        case -299:
                            v58 := []int{v3}
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 16; fi59++ {
                            v40 = "\n\t"
                            v10 = (-652 >= 625)
                            if ((-2.25 != -2.25) || ((v4 * v4) != (v4 % ((v4) | 1)))) {
                                if false {
                                    if (1 < 42) {
                                        v10 = v10
                                    } else {
                                        for fi60 := 0; fi60 < 11; fi60++ {
                                            fmt.Println(fi60, v3, v2)
                                            _ = fi60
                                        }
                                        fmt.Println(v10, v4)
                                        fmt.Println(v3)
                                    }
                                }
                            } else {
                                var v61 int8 = 43
                                for fi62 := 0; fi62 < 4; fi62++ {
                                    var v63 uint8 = v6
                                    v10 = (!(-1 <= 1000))
                                    if (v4 <= (-(v61))) {
                                        if false {
                                            v63 += v6
                                        } else {
                                            switch v3 {
                                            case -287:
                                                v2 += v2
                                            case 1:
                                                v64 := GenS6{F0: 65535}
                                                _ = v64
                                            }
                                            var v65 int = 2147483647
                                            v66 := []int{42, fi62, (v65 * -130), 728}
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                    _ = fi62
                                    _ = v63
                                }
                                if true {
                                    v40 = "hello"
                                } else {
                                    for fi67 := 0; fi67 < 4; fi67++ {
                                        v41 = v10
                                        v68 := []int{fi59}
                                        v10 = (float64(v2) > -2.25)
                                        _ = fi67
                                        _ = v68
                                    }
                                }
                                _ = v61
                            }
                            _ = fi59
                        }
                    }
                    _ = v56
                }
                switch int(v4) {
                case -9223372036854775808:
                    for fi69 := 0; fi69 < 3; fi69++ {
                        var v70 uint32 = 2147483647
                        switch (v40 + v40) {
                        case "hello":
                            v70 -= v70
                        }
                        _ = fi69
                        _ = v70
                    }
                case 0:
                    v71 := GenS1{F0: v6, F1: 27}
                    _ = v71
                }
            }
        case 819:
            v72 := GenS5{F0: -1, F1: 628, F2: 2}
            _ = v72
        }
    case "":
        v3 -= v3
    default:
        v41 = (!(1 != 2))
    }
    var v73 uint32 = 2
    if v41 {
        v74 := GenS4{F0: (v40 + (("hello" + ((("hello" + v40) + ((("\n\t" + v40) + ("" + ("a" + v40))) + (("\n\t" + ("foo bar" + "")) + "foo bar"))) + v40)) + (((v40 + (v40 + v40)) + "\n\t") + v40)))}
        _ = len(v1)
        for fi75 := 0; fi75 < 12; fi75++ {
            v76 := GenS4{F0: (v40 + v40)}
            switch v40 {
            case "\n\t":
                fmt.Println(v41, v41)
            case "foo bar":
                switch v40 {
                case "foo bar":
                    if (v2 >= (v2 * (v2 * v2))) {
                        v77 := GenS7{F0: (-(v3))}
                        v78 := []int{v3, v3}
                        _ = v77
                        _ = v78
                    }
                case "\n\t":
                    var v79 int64 = 151
                    _ = v79
                }
            case "a":
                v2 = (v2 / 0.1)
            }
            switch "" {
            case "":
                var v80 int8 = (v4 % (((v4 | v4)) | 1))
                _ = v80
            case "\n\t":
                delete(v5, (v3 + -2147483648))
            }
            _ = fi75
            _ = v76
        }
        _ = v74
    }
    switch (v3 >> ((18446744073709551615) & 63)) {
    case 9223372036854775807:
        v2 = v2
    case 2147483647:
        for fi81 := 0; fi81 < 18; fi81++ {
            var v82 uint32 = (^(v73))
            v73 -= (v82 ^ (v82 | v82))
            _ = fi81
            _ = v82
        }
    case -614:
        v83 := GenS8{F0: (v6 / (((v6 - v6)) | 1)), F1: 1}
        _ = v83
    default:
        fmt.Println(v41, v73)
    }
    v10 = true
    if (true && (v10 && (1 == 469))) {
        var v84 int32 = -2147483648
        _ = v84
    }
    for fi85 := 0; fi85 < 6; fi85++ {
        if true {
            v86 := GenS9{F0: -579}
            for fi87 := 0; fi87 < 19; fi87++ {
                if (1.5 == (1.5 / 1.5)) {
                    v88 := []int{fi85, (fi85 - fi85)}
                    _ = v88
                } else {
                    v89 := v1[fi87]
                    v41 = v10
                    v90 := []int{(fi87 ^ (fi87 ^ v3)), 103, -149, (v3 >> ((uint(v6)) & 63))}
                    _ = v89
                    _ = v90
                }
                v91 := GenS10{F0: 18446744073709551615}
                v6 = v6
                _ = fi87
                _ = v91
            }
            fmt.Println(v73, v73)
            _ = v86
        }
        v5[32768] = (!v10)
        _ = fi85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v73
}
