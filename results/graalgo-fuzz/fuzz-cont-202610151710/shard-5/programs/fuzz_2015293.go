package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 int }
type GenS3 struct { F0 uint; F1 bool; F2 uint }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int64; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS9 struct { F0 bool; F1 int; F2 string }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 int64; F1 int16 }
type GenS12 struct { F0 string }

func main() {
    var v1 uint8 = 255
    var v2 float64 = 1.5
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    var v5 int32 = 1
    var v6 uint32 = 2147483648
    v7 := []int{127}
    fmt.Println(v1)
    var v8 float32 = 1e10
    if (-1 <= -1) {
        if false {
            switch -2147483648 {
            case 2147483647:
                if true {
                    var v9 float32 = float32((1e10 * v2))
                    v10 := GenS1{F0: 280, F1: 167, F2: 0}
                    _ = v10
                    _ = v9
                } else {
                    v5 = 0
                }
            default:
                v11 := GenS1{F0: 293, F1: 740, F2: 1}
                _ = v11
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(v5)
                fmt.Println(v5, fi12, v2)
                _ = fi12
            }
            v4[0] = false
        } else {
            if (!true) {
                v13 := GenS2{F0: 0, F1: v5, F2: 434}
                for fi14 := 0; fi14 < 5; fi14++ {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 uint = 2
                        _ = fi15
                        _ = v16
                    }
                    v2 += v2
                    var v17 bool = (!true)
                    _ = fi14
                    _ = v17
                }
                _ = v13
            } else {
                switch 759 {
                case 128:
                    var v18 float32 = ((0.0 * v8) * float32(float64((0.0 * (-2.25 + v8)))))
                    _ = v18
                default:
                    v1 = v1
                }
                fmt.Println(v5)
            }
        }
        if false {
            var v19 int32 = 2147483647
            v6 = (v6 ^ 4294967295)
            _ = v19
        } else {
            v1 += 0
            v5 -= v5
            var v20 uint64 = 798
            _ = v20
        }
        fmt.Println(v5, v6)
    }
    var v21 float64 = (1e10 / 3.14159)
    var v22 uint64 = 18446744073709551615
    if (!(0.1 < 0.0)) {
        v4[(v6 << ((4294967295) & 63))] = true
        v23 := GenS2{F0: -1, F1: (v5 | (^(v5))), F2: 42}
        _ = v23
    } else {
        switch ("a" + "") {
        case "hello":
            v8 = (1e10 - 1.5)
        default:
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 uint16 = 219
                var v26 uint16 = (^(v25))
                if true {
                    var v27 bool = true
                    if (v27 && (42 < 9223372036854775807)) {
                        v28 := v4[v6]
                        _ = v28
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    _ = v27
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        if len(v7) > 0 {
            v7[0] = 128
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            fmt.Println(v1, v21, v2)
            v21 += (v21 - float64((v8 / (1e10 + v8))))
            v30 := v3[-1]
            _ = fi29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        if (-32768 >= 1) {
            v6 = v6
            var v32 uint32 = v6
            _ = v32
        } else {
            if len(v7) > 0 {
                v7[0] = 42
            }
        }
        if false {
            switch ((("\n\t" + "") + "a") + "\n\t") {
            case "a":
                switch "a" {
                case "foo bar":
                    if (v21 <= (v2 / (float64(v8) * 0.1))) {
                        v1 -= 42
                        var v33 int64 = 42
                        if false {
                            switch (("\n\t" + ((((("" + "foo bar") + "\n\t") + "hello") + (("a" + "") + "foo bar")) + "hello")) + "hello") {
                            case "foo bar":
                                if (1000 < -1) {
                                    v5 = v5
                                    var v34 uint16 = uint16(v22)
                                    v35 := GenS3{F0: 1, F1: true, F2: 4294967295}
                                    _ = v34
                                    _ = v35
                                }
                            case "a":
                                for fi36 := 0; fi36 < 15; fi36++ {
                                    fmt.Println(v2, v21)
                                    var v37 uint8 = (v1 * v1)
                                    v38 := GenS4{F0: v33, F1: fi36}
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            }
                            fmt.Println(v6, v33)
                            var v39 int8 = -114
                            _ = v39
                        }
                        _ = v33
                    }
                default:
                    v21 = (float64(3.14159) - v2)
                }
            case "foo bar":
                fmt.Println(fi31)
            case "\n\t":
                v22 += (^(v22))
            default:
                var v40 int8 = 23
                _ = v40
            }
            switch fi31 {
            case 565:
                if true {
                    v7 = append(v7, 2147483647)
                    delete(v3, (fi31 % (((fi31 - fi31)) | 1)))
                    for fi41 := 0; fi41 < 9; fi41++ {
                        if true {
                            v1 -= (v1 << ((59) & 63))
                        } else {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                        var v42 int8 = -1
                        _ = fi41
                        _ = v42
                    }
                }
            case 32768:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case 419:
                for fi43 := 0; fi43 < 3; fi43++ {
                    switch "hello" {
                    case "foo bar":
                        v22 = v22
                    case "":
                        if false {
                            v44 := GenS4{F0: -9223372036854775808, F1: -744}
                            if false {
                                if true {
                                    v45 := GenS1{F0: 0, F1: fi43, F2: 100}
                                    switch fi43 {
                                    case 9223372036854775807:
                                        v7 = append(v7, fi31)
                                    case 32767:
                                        var v46 uint = 2
                                        _ = v46
                                    case 42:
                                        v8 += 3.14159
                                    }
                                    switch (^(fi43)) {
                                    case 32768:
                                        var v47 bool = ((v5 >= int32(fi31)) && ("a" == "a"))
                                        _ = v47
                                    }
                                    v6 += v6
                                    var v48 int8 = 79
                                    _ = v45
                                    _ = v48
                                } else {
                                    switch ((("\n\t" + ("\n\t" + ("a" + "\n\t"))) + ("\n\t" + "\n\t")) + "foo bar") {
                                    case "\n\t":
                                        _ = len(v7)
                                    case "":
                                        var v49 bool = (!(false && true))
                                        _ = v49
                                    }
                                }
                                v50 := GenS4{F0: -759, F1: 1}
                                _ = v50
                            } else {
                                v51 := GenS1{F0: (fi43 % ((296) | 1)), F1: 32768, F2: -128}
                                if false {
                                    v2 = v21
                                    if false {
                                        _ = len(v3)
                                        fmt.Println(fi43, v6)
                                    }
                                    v52 := GenS5{F0: "foo bar", F1: -356}
                                    _ = v52
                                } else {
                                    var v53 bool = false
                                    var v54 int64 = 460
                                    _ = v53
                                    _ = v54
                                }
                                v55 := GenS6{F0: 75}
                                _ = v51
                                _ = v55
                            }
                            _ = v44
                        } else {
                            fmt.Println(v2, v8, fi31)
                            v56 := GenS7{F0: (fi43 * (-(fi43))), F1: -629, F2: -32768}
                            v7 = append(v7, fi43)
                            _ = v56
                        }
                    case "\n\t":
                        switch 1 {
                        case 9223372036854775807:
                            v57 := v3[-282]
                            _ = v57
                        default:
                            var v58 string = "hello"
                            _ = v58
                        }
                    }
                    switch "a" {
                    case "":
                        var v59 uint8 = (v1 << ((18446744073709551615) & 63))
                        _ = v59
                    case "a":
                        if (((v8 + v8) == (1e10 * v8)) || (!(!false))) {
                            v3[(-(fi31))] = 2147483648
                        }
                    case "\n\t":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    default:
                        var v60 int8 = 100
                        _ = v60
                    }
                    _ = fi43
                }
            default:
                if (v22 > 18446744073709551615) {
                    for fi61 := 0; fi61 < 16; fi61++ {
                        v62 := GenS3{F0: 42, F1: (!true), F2: 42}
                        v63 := GenS7{F0: fi31, F1: -295, F2: 0}
                        v64 := GenS2{F0: -32768, F1: v5, F2: 4294967295}
                        _ = fi61
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    if false {
                        if (uint16(v6) >= 0) {
                            var v65 int64 = -1
                            v66 := GenS8{F0: v6, F1: v22, F2: 825}
                            _ = v65
                            _ = v66
                        } else {
                            for fi67 := 0; fi67 < 17; fi67++ {
                                fmt.Println(v8)
                                var v68 float64 = 1e10
                                if true {
                                    v69 := GenS9{F0: ((0 != 0) || false), F1: 4294967295, F2: "a"}
                                    fmt.Println(fi67, v1, v5)
                                    var v70 string = ("" + ("" + ("foo bar" + ("foo bar" + ("\n\t" + (("" + "") + ("a" + ("\n\t" + "a"))))))))
                                    var v71 int16 = 156
                                    fmt.Println(v68, v22, v6)
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                                _ = fi67
                                _ = v68
                            }
                            if (42 <= 2) {
                                v3[(fi31 | (fi31 - (-(fi31))))] = v6
                                for fi72 := 0; fi72 < 18; fi72++ {
                                    v2 += float64(v8)
                                    var v73 float64 = v2
                                    _ = fi72
                                    _ = v73
                                }
                                v4[(v6 % (((v6 % ((74) | 1))) | 1))] = true
                            }
                        }
                        var v74 int8 = 48
                        _ = v74
                    } else {
                        v1 -= v1
                    }
                } else {
                    v75 := GenS10{F0: "\n\t", F1: (v5 - (v5 / ((int32(v6)) | 1)))}
                    v76 := GenS11{F0: 9223372036854775807, F1: 1}
                    var v77 int64 = 1
                    _ = v75
                    _ = v76
                    _ = v77
                }
            }
            fmt.Println(v2)
        }
        if true {
            v6 = v6
        }
        _ = fi31
    }
    switch "" {
    case "":
        fmt.Println(v5, v22)
    case "a":
        v8 = 1.5
    case "hello":
        fmt.Println(v8, v1, v21)
    }
    v6 -= v6
    v5 = 1
    if (-1 >= 9223372036854775807) {
        if len(v7) > 0 {
            v7[0] = 42
        }
    }
    fmt.Println(v8)
    switch 1 {
    case -9223372036854775808:
        switch "hello" {
        case "foo bar":
            for fi78 := 0; fi78 < 19; fi78++ {
                if false {
                    var v79 uint = 2
                    _ = v79
                }
                _ = fi78
            }
        case "\n\t":
            var v80 int32 = 938
            _ = v80
        case "":
            fmt.Println(v2)
        default:
            fmt.Println(v2, v21, v2)
        }
    }
    var v81 uint8 = (v1 % (((v1 >> ((0) & 63))) | 1))
    v6 += (^(v6))
    for fi82 := 0; fi82 < 13; fi82++ {
        if false {
            switch "hello" {
            case "hello":
                var v83 int16 = 384
                _ = v83
            case "":
                var v84 uint64 = (v22 << ((0) & 63))
                _ = v84
            case "\n\t":
                if (112 != 42) {
                    v6 = (v6 * 933)
                    var v85 uint64 = (v22 << ((18446744073709551615) & 63))
                    for fi86 := 0; fi86 < 2; fi86++ {
                        if true {
                            v87 := GenS12{F0: "\n\t"}
                            _ = v87
                        }
                        _ = fi86
                    }
                    _ = v85
                } else {
                    var v88 string = "foo bar"
                    var v89 string = ((v88 + (v88 + v88)) + "hello")
                    v2 += v2
                    _ = v88
                    _ = v89
                }
            default:
                for fi90 := 0; fi90 < 5; fi90++ {
                    switch "\n\t" {
                    case "foo bar":
                        var v91 uint = 4294967295
                        _ = v91
                    case "\n\t":
                        if false {
                            var v92 string = ("" + "")
                            _ = v92
                        } else {
                            v22 = v22
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(fi82, v2, v1)
                            case "foo bar":
                                var v93 bool = true
                                _ = v93
                            case "a":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            v7 = append(v7, fi90)
                        }
                    case "":
                        if true {
                            var v94 int16 = 32767
                            _ = v94
                        } else {
                            if len(v7) > 0 {
                                v7[0] = (fi82 & (fi90 ^ fi82))
                            }
                            if false {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            } else {
                                fmt.Println(v8, v5)
                                if true {
                                    v95 := GenS11{F0: 42, F1: 32767}
                                    v96 := v4[(v6 - v6)]
                                    _ = v95
                                    _ = v96
                                }
                                var v97 int32 = (v5 ^ (v5 / ((42) | 1)))
                                _ = v97
                            }
                            if false {
                                var v98 int64 = -1
                                fmt.Println(v21)
                                _ = v98
                            }
                        }
                    default:
                        v99 := GenS1{F0: (fi90 & -248), F1: fi82, F2: 42}
                        _ = v99
                    }
                    v5 += v5
                    fmt.Println(v6)
                    _ = fi90
                }
            }
            var v100 uint64 = (^(v22))
            delete(v4, 890)
            _ = v100
        } else {
            v5 = v5
            var v101 uint16 = 641
            _ = v101
        }
        _ = fi82
    }
    v7 = append(v7, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v81
}
