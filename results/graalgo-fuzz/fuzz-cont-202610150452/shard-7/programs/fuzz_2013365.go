package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 bool }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 int; F2 string }

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 190
    var v3 uint32 = 2147483648
    var v4 string = ("foo bar" + "hello")
    var v5 uint64 = 1
    switch ("" + ("" + v4)) {
    case "a":
        if false {
            v6 := GenS1{F0: 0}
            var v7 bool = false
            switch "a" {
            case "hello":
                switch v4 {
                case "a":
                    v8 := []int{0, 30, int(v5), 4294967295}
                    _ = v8
                }
            case "":
                var v9 int32 = 1
                _ = v9
            }
            _ = v6
            _ = v7
        } else {
            if true {
                var v10 uint16 = (^(v1))
                _ = v10
            }
        }
    case "hello":
        if true {
            v11 := GenS2{F0: -9223372036854775808, F1: 45}
            v12 := []int{1, int(v2), 32768}
            _ = v11
            _ = v12
        }
    case "":
        switch 127 {
        case 32768:
            var v13 int16 = 694
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            if false {
                switch fi14 {
                case -204:
                    v15 := []int{}
                    _ = v15
                case -52:
                    var v16 int64 = -9223372036854775808
                    _ = v16
                case 4294967295:
                    var v17 int = -666
                    _ = v17
                }
                fmt.Println(fi14, fi14)
                switch (fi14 ^ fi14) {
                case 9223372036854775807:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if true {
                            var v19 uint8 = v2
                            switch 2147483647 {
                            case 32768:
                                var v20 int64 = 152
                                _ = v20
                            case 128:
                                if true {
                                    v5 -= uint64(v1)
                                    v21 := GenS3{F0: (v3 - (^(v3)))}
                                    var v22 int64 = 568
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v23 uint32 = 2147483647
                                        _ = v23
                                    }
                                    v24 := GenS2{F0: fi14, F1: 42}
                                    _ = v21
                                    _ = v22
                                    _ = v24
                                } else {
                                    fmt.Println(v19, v19)
                                    v25 := GenS3{F0: 4294967295}
                                    switch (v4 + v4) {
                                    case "a":
                                        var v26 uint64 = v5
                                        _ = v26
                                    case "":
                                        v5 = (v5 / (((v5 % ((v5) | 1))) | 1))
                                    default:
                                        var v27 int8 = 42
                                        _ = v27
                                    }
                                    if false {
                                        fmt.Println(v5)
                                        var v28 int32 = 685
                                        _ = v28
                                    }
                                    _ = v25
                                }
                            case 127:
                                v5 -= 1
                            default:
                                fmt.Println(v5)
                            }
                            _ = v19
                        } else {
                            v29 := []int{fi14, int(v5), (fi14 << ((18446744073709551615) & 63))}
                            var v30 uint16 = uint16(v5)
                            if false {
                                if len(v29) > 0 {
                                    v29[0] = (fi18 * (fi18 % (((fi18 >> ((42) & 63))) | 1)))
                                }
                            }
                            _ = v29
                            _ = v30
                        }
                        switch ((v4 + v4) + v4) {
                        case "hello":
                            v2 = v2
                        case "\n\t":
                            v3 += v3
                            v3 += (v3 & 42)
                            var v31 float64 = ((-2.25 + 0.0) + 0.1)
                            _ = v31
                        case "":
                            var v32 uint = 1
                            _ = v32
                        }
                        _ = fi18
                    }
                case 32767:
                    var v33 int16 = 1000
                    _ = v33
                case 32768:
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v2 = v2
                        v35 := GenS1{F0: -9223372036854775808}
                        var v36 int = -772
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            } else {
                switch fi14 {
                case 0:
                    v3 -= (v3 & uint32(v5))
                case 127:
                    var v37 int32 = 42
                    _ = v37
                }
                var v38 int = (fi14 * fi14)
                var v39 int = 32768
                _ = v38
                _ = v39
            }
            _ = fi14
        }
    }
    switch 32767 {
    case 700:
        var v40 bool = true
        _ = v40
    case 314:
        for fi41 := 0; fi41 < 11; fi41++ {
            if false {
                v4 = "hello"
                for fi42 := 0; fi42 < 12; fi42++ {
                    v2 -= 2
                    var v43 bool = false
                    switch v4 {
                    case "\n\t":
                        v5 = v5
                    case "a":
                        v44 := []int{}
                        _ = v44
                    default:
                        v45 := []int{}
                        _ = v45
                    }
                    _ = fi42
                    _ = v43
                }
                var v46 int8 = 11
                _ = v46
            } else {
                switch (fi41 << ((65535) & 63)) {
                case 0:
                    var v47 uint16 = 255
                    _ = v47
                }
                for fi48 := 0; fi48 < 2; fi48++ {
                    var v49 int8 = 86
                    if true {
                        v3 -= (^(v3))
                        v3 = (v3 * (^(v3)))
                    }
                    _ = fi48
                    _ = v49
                }
                var v50 float64 = 0.0
                _ = v50
            }
            _ = fi41
        }
    }
    v51 := GenS1{F0: 684}
    for fi52 := 0; fi52 < 15; fi52++ {
        var v53 float32 = 1e10
        _ = fi52
        _ = v53
    }
    if false {
        for fi54 := 0; fi54 < 8; fi54++ {
            var v55 float64 = 3.14159
            v56 := []int{int(v5), fi54}
            _ = fi54
            _ = v55
            _ = v56
        }
        for fi57 := 0; fi57 < 20; fi57++ {
            fmt.Println(v3)
            if (0.0 <= (((float64(0.1) * (-2.25 / (1e10 - (1.5 * (1e10 * 3.14159))))) - (((0.0 + (0.1 + 0.0)) + (((-2.25 - 0.0) / 3.14159) + 0.0)) / ((0.1 / -2.25) * -2.25))) + float64(((1.5 / ((1.5 / -2.25) * -2.25)) / 0.1)))) {
                v58 := GenS4{F0: v3, F1: v3, F2: ((fi57 & 32767) < 175)}
                if (-1 > 42) {
                    v2 = 128
                    var v59 int = fi57
                    _ = v59
                }
                _ = v58
            } else {
                for fi60 := 0; fi60 < 9; fi60++ {
                    v2 += (^(v2))
                    if true {
                        v4 = "foo bar"
                    }
                    _ = fi60
                }
            }
            for fi61 := 0; fi61 < 7; fi61++ {
                v4 = (v4 + (v4 + "a"))
                _ = fi61
            }
            _ = fi57
        }
        for fi62 := 0; fi62 < 3; fi62++ {
            for fi63 := 0; fi63 < 16; fi63++ {
                fmt.Println(v1, v1, fi63)
                v4 = "foo bar"
                var v64 uint8 = uint8(v5)
                _ = fi63
                _ = v64
            }
            var v65 int16 = -143
            _ = fi62
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 12; fi66++ {
        var v67 uint64 = v5
        if (true && (0.0 == (((float32((0.1 + 1.5)) * 1e10) - 0.0) / 0.1))) {
            if (42 < 42) {
                v4 = v4
                v68 := []int{}
                v5 = v67
                _ = v68
            } else {
                var v69 string = ("foo bar" + v4)
                _ = v69
            }
            for fi70 := 0; fi70 < 10; fi70++ {
                if (244 > 2147483647) {
                    v4 = ((("hello" + "\n\t") + (v4 + ("a" + v4))) + v4)
                    switch fi66 {
                    case 4294967295:
                        v2 += v2
                        v71 := GenS5{F0: (v67 * (v67 / (((v67 | 1)) | 1))), F1: (v2 | v2), F2: 2}
                        _ = v71
                    default:
                        v2 -= (v2 ^ (v2 & 255))
                        v72 := GenS6{F0: 287, F1: v4}
                        var v73 int8 = -121
                        v74 := GenS7{F0: v3}
                        switch "\n\t" {
                        case "":
                            v4 = v4
                            v2 += v2
                            v75 := []int{fi70, (fi70 % (((-(fi66))) | 1))}
                            _ = v75
                        case "foo bar":
                            var v76 bool = ((fi66 | fi70) == fi70)
                            _ = v76
                        default:
                            var v77 bool = false
                            _ = v77
                        }
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                } else {
                    v1 = v1
                    fmt.Println(v5, v1)
                    v78 := []int{(fi70 + fi70), (fi66 - fi70), (fi70 % ((32767) | 1))}
                    var v79 uint32 = (v3 + v3)
                    if len(v78) > 0 {
                        v78[0] = -9223372036854775808
                    }
                    _ = v78
                    _ = v79
                }
                fmt.Println(v3, v1, v3)
                _ = fi70
            }
            v80 := GenS7{F0: 2147483647}
            _ = v80
        } else {
            if (-54 >= 0) {
                v81 := []int{fi66, fi66, (fi66 ^ 9223372036854775807), (fi66 >> ((18446744073709551615) & 63))}
                if false {
                    v67 -= v67
                }
                switch v4 {
                case "hello":
                    fmt.Println(v1)
                default:
                    var v82 uint = 2
                    _ = v82
                }
                _ = v81
            }
            fmt.Println(v67, v1)
            if (false && true) {
                switch v4 {
                case "foo bar":
                    var v83 uint = 255
                    _ = v83
                case "":
                    var v84 int = fi66
                    _ = v84
                default:
                    fmt.Println(v4, v2, v3)
                }
                v85 := GenS1{F0: 9223372036854775807}
                _ = v85
            } else {
                if true {
                    fmt.Println(v4)
                    if (0 > -32768) {
                        for fi86 := 0; fi86 < 12; fi86++ {
                            fmt.Println(fi66, v67)
                            v87 := GenS8{F0: -1, F1: fi86, F2: v4}
                            _ = fi86
                            _ = v87
                        }
                        switch v4 {
                        case "foo bar":
                            var v88 int8 = 62
                            _ = v88
                        case "\n\t":
                            v3 = uint32(v1)
                        case "a":
                            v89 := []int{-487, (fi66 % ((-233) | 1))}
                            _ = v89
                        }
                        v90 := []int{fi66, fi66, (fi66 >> ((65535) & 63))}
                        _ = v90
                    } else {
                        var v91 int16 = 32767
                        _ = v91
                    }
                    for fi92 := 0; fi92 < 9; fi92++ {
                        v3 += (v3 * 651)
                        v93 := GenS7{F0: v3}
                        _ = fi92
                        _ = v93
                    }
                }
            }
        }
        _ = fi66
        _ = v67
    }
    v94 := []int{42}
    var v95 bool = (!(18446744073709551615 == (v5 << ((0) & 63))))
    var v96 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v94:", len(v94), v94)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v94
    _ = v95
    _ = v96
}
