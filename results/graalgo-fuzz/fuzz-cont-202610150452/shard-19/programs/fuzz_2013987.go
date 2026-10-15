package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 461
    var v3 uint8 = 42
    v4 := make(map[int]int)
    var v5 uint = 1
    var v6 int = -1
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int16 = int16(v5)
        switch "hello" {
        case "hello":
            fmt.Println(v3)
        case "foo bar":
            for fi9 := 0; fi9 < 4; fi9++ {
                fmt.Println(fi9, fi9, fi9)
                v8 -= -32768
                v10 := GenS1{F0: (42 >= 0)}
                v11 := GenS2{F0: (v6 | 127), F1: 0, F2: 832}
                _ = fi9
                _ = v10
                _ = v11
            }
        case "a":
            var v12 uint8 = (^(v3))
            _ = v12
        default:
            fmt.Println(v2, v3)
        }
        fmt.Println(v2, v8)
        _ = fi7
        _ = v8
    }
    fmt.Println(v2)
    for fi13 := 0; fi13 < 17; fi13++ {
        if (!false) {
            var v14 int8 = -128
            var v15 float32 = 3.14159
            _ = v14
            _ = v15
        }
        var v16 uint8 = v3
        _ = fi13
        _ = v16
    }
    switch (v6 >> ((42) & 63)) {
    case 9223372036854775807:
        var v17 int = (v6 | (v6 - (v6 | 539)))
        _ = v17
    case 288:
        _ = len(v4)
    default:
        var v18 int8 = 119
        _ = v18
    }
    if (1 <= 42) {
        v1["a"] = uint32(v3)
        switch (v6 << (((v5 << ((v5) & 63))) & 63)) {
        case 32767:
            var v19 int32 = -1
            _ = v19
        default:
            if true {
                v20 := GenS3{F0: (v2 | v2)}
                _ = v20
            } else {
                var v21 uint64 = uint64(v5)
                for fi22 := 0; fi22 < 12; fi22++ {
                    v23 := GenS1{F0: true}
                    for fi24 := 0; fi24 < 20; fi24++ {
                        switch "hello" {
                        case "hello":
                            var v25 uint32 = 117
                            _ = v25
                        case "":
                            switch "\n\t" {
                            case "foo bar":
                                var v26 uint16 = 398
                                _ = v26
                            case "hello":
                                v6 -= v6
                            case "a":
                                var v27 bool = (fi22 > (fi24 >> ((2) & 63)))
                                _ = v27
                            default:
                                if (2 != (v21 / ((v21) | 1))) {
                                    v28 := GenS2{F0: 1, F1: 9223372036854775807, F2: 517}
                                    _ = v28
                                }
                            }
                        default:
                            v6 = 32767
                        }
                        if (-2.25 == 3.14159) {
                            var v29 int = (fi22 - (fi24 << ((1) & 63)))
                            var v30 float32 = 1.5
                            switch int(v21) {
                            case 2147483647:
                                var v31 int32 = -363
                                _ = v31
                            case 9223372036854775807:
                                v32 := GenS1{F0: false}
                                _ = v32
                            case 32768:
                                _ = len(v1)
                            default:
                                fmt.Println(v21, fi22)
                            }
                            _ = v29
                            _ = v30
                        } else {
                            v33 := GenS4{F0: 1000, F1: (^(v2)), F2: 65535}
                            _ = v33
                        }
                        var v34 float64 = 0.1
                        if (fi24 != -9223372036854775808) {
                            v5 = (v5 - v5)
                            var v35 uint64 = v21
                            _ = v35
                        }
                        if true {
                            if (-1 == 127) {
                                fmt.Println(v2, fi24)
                                v36 := GenS2{F0: (v6 / ((22) | 1)), F1: 42, F2: 0}
                                v21 += uint64(v5)
                                _ = v36
                            } else {
                                var v37 int32 = -2147483648
                                switch 127 {
                                case 210:
                                    v38 := []int{v6}
                                    _ = v38
                                }
                                v39 := GenS1{F0: (false && true)}
                                _ = v37
                                _ = v39
                            }
                        }
                        _ = fi24
                        _ = v34
                    }
                    for fi40 := 0; fi40 < 10; fi40++ {
                        switch fi22 {
                        case 127:
                            delete(v1, "foo bar")
                        case 614:
                            fmt.Println(v21)
                        case 0:
                            v41 := GenS5{F0: v2}
                            _ = v41
                        default:
                            v42 := GenS3{F0: 2}
                            _ = v42
                        }
                        if (!((427 != 975) && (v3 == 95))) {
                            if true {
                                v43 := GenS1{F0: true}
                                switch -40 {
                                case 4294967295:
                                    if false {
                                        fmt.Println(v2, v21)
                                    }
                                case 42:
                                    fmt.Println(v3)
                                }
                                var v44 int64 = -9223372036854775808
                                _ = v43
                                _ = v44
                            }
                        }
                        v45 := v4[fi22]
                        _ = fi40
                        _ = v45
                    }
                    _ = fi22
                    _ = v23
                }
                v21 = 18446744073709551615
                _ = v21
            }
        }
        for fi46 := 0; fi46 < 12; fi46++ {
            var v47 uint = v5
            var v48 float64 = (0.1 + float64((-2.25 - (1e10 / float32(((((-2.25 / 1.0) + (-2.25 * 0.0)) / (0.1 / float64(1e10))) - 1.5))))))
            switch (-(fi46)) {
            case 1:
                var v49 uint64 = (v2 | v2)
                _ = v49
            default:
                v50 := []int{(v6 ^ fi46), (fi46 % (((v6 / ((-2147483648) | 1))) | 1)), (fi46 ^ (fi46 - fi46)), (v6 >> ((v47) & 63))}
                _ = v50
            }
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    if true {
        for fi51 := 0; fi51 < 20; fi51++ {
            v52 := GenS4{F0: -1, F1: (^(v2)), F2: 2}
            fmt.Println(v3)
            v6 += fi51
            _ = fi51
            _ = v52
        }
        v1[("a" + "\n\t")] = 42
        v5 = (v5 >> (((^(v5))) & 63))
    }
    if (!false) {
        fmt.Println(v3, v2, v3)
        fmt.Println(v5)
    }
    var v53 uint32 = 2147483647
    v2 -= 827
    if ((^(v53)) != 488) {
        v54 := GenS1{F0: ((!("\n\t" != "a")) || true)}
        for fi55 := 0; fi55 < 17; fi55++ {
            switch "hello" {
            case "foo bar":
                if (!true) {
                    v56 := []int{4294967295}
                    _ = v56
                } else {
                    if false {
                        if false {
                            var v57 string = "\n\t"
                            if (3.14159 < float32((0.1 * 0.1))) {
                                if true {
                                    v58 := []int{9223372036854775807, v6}
                                    _ = v58
                                } else {
                                    var v59 int = -742
                                    _ = v59
                                }
                                fmt.Println(v5)
                                v53 = uint32(v5)
                            } else {
                                var v60 float64 = 0.0
                                var v61 string = v57
                                v2 += v2
                                _ = v60
                                _ = v61
                            }
                            _ = v57
                        }
                        if false {
                            if ((1e10 / 3.14159) < 3.14159) {
                                v62 := GenS2{F0: (fi55 & -56), F1: -17, F2: v53}
                                _ = v62
                            } else {
                                fmt.Println(v2, v6)
                                v63 := []int{v6, fi55, (fi55 + -745), (v6 / ((fi55) | 1))}
                                if ((^(v53)) >= (v53 % (((v53 + (v53 & v53))) | 1))) {
                                    for fi64 := 0; fi64 < 2; fi64++ {
                                        v65 := GenS3{F0: (v2 ^ (v2 / ((v2) | 1)))}
                                        var v66 uint = v5
                                        v2 += v2
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                } else {
                                    for fi67 := 0; fi67 < 14; fi67++ {
                                        var v68 uint64 = (v2 << ((v5) & 63))
                                        _ = fi67
                                        _ = v68
                                    }
                                    if (367 < v5) {
                                        v3 = (v3 | (^(v3)))
                                    }
                                }
                                _ = v63
                            }
                            var v69 uint16 = 0
                            _ = v69
                        } else {
                            switch (("a" + "\n\t") + ("foo bar" + "")) {
                            case "foo bar":
                                switch "" {
                                case "":
                                    if false {
                                        switch ("\n\t" + "a") {
                                        case "":
                                            v70 := GenS1{F0: (!(!false))}
                                            _ = v70
                                        case "hello":
                                            var v71 int32 = -520
                                            _ = v71
                                        case "\n\t":
                                            var v72 uint64 = 0
                                            _ = v72
                                        default:
                                            var v73 int64 = 667
                                            _ = v73
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(fi55)
                                case "\n\t":
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v74 uint16 = 172
                                        _ = v74
                                    case "hello":
                                        for fi75 := 0; fi75 < 20; fi75++ {
                                            var v76 uint = v5
                                            v2 -= v2
                                            _ = fi75
                                            _ = v76
                                        }
                                    }
                                default:
                                    v53 -= v53
                                }
                            case "\n\t":
                                v77 := GenS5{F0: (v2 * (v2 / ((0) | 1)))}
                                _ = v77
                            }
                            v78 := GenS6{F0: "foo bar", F1: uint64(v5)}
                            _ = v78
                        }
                    } else {
                        v79 := GenS1{F0: false}
                        _ = v79
                    }
                    var v80 uint32 = (v53 | 635)
                    _ = v80
                }
            case "\n\t":
                var v81 int8 = -58
                _ = v81
            case "a":
                var v82 float64 = (float64(1.5) * (-2.25 * (1.5 / 0.1)))
                _ = v82
            default:
                v83 := GenS1{F0: true}
                _ = v83
            }
            fmt.Println(fi55, v3)
            _ = fi55
        }
        v84 := GenS7{F0: 9}
        _ = v54
        _ = v84
    } else {
        var v85 int64 = 42
        fmt.Println(v85, v5, v3)
        _ = v85
    }
    switch "" {
    case "a":
        fmt.Println(v3, v5)
    case "":
        _ = len(v1)
    case "hello":
        for fi86 := 0; fi86 < 15; fi86++ {
            switch (^(fi86)) {
            case -9223372036854775808:
                fmt.Println(v6, v53)
            }
            _ = fi86
        }
    }
    fmt.Println(v5, v2, v53)
    fmt.Println(v3)
    var v87 bool = (0.1 > float64((1.5 - 0.1)))
    switch v6 {
    case 4294967295:
        v87 = (v87 || (32767 != 1000))
    case 2147483647:
        v2 -= (v2 << ((964) & 63))
    case 127:
        v6 = -1
    default:
        if v87 {
            var v88 int32 = 42
            for fi89 := 0; fi89 < 8; fi89++ {
                var v90 int32 = (-(v88))
                v91 := GenS4{F0: 650, F1: 228, F2: 0}
                var v92 float64 = 1.5
                _ = fi89
                _ = v90
                _ = v91
                _ = v92
            }
            _ = v88
        } else {
            _ = len(v4)
            fmt.Println(v53, v87)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v87
}
