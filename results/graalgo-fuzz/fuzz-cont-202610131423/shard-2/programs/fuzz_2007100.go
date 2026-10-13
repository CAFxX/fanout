package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int8; F2 int }
type GenS5 struct { F0 int32; F1 uint8; F2 int8 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint; F1 int32; F2 uint8 }
type GenS8 struct { F0 int; F1 uint; F2 int32 }
func genFn0(p0 uint16) string {
    fmt.Println(p0)
    p0 += p0
    fmt.Println(p0)
    _ = p0
    return "hello"
}
func genFn1(p0 uint, p1 uint16) int8 {
    v7 := []int{-1, -1, -946}
    p1 = p1
    p0 = p0
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return -127
}

func main() {
    var v1 float64 = 0.0
    var v2 float64 = v1
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v5 float32 = 1.5
    v6 := make(map[uint32]uint32)
    v8 := GenS1{F0: -485, F1: 42}
    fmt.Println(v2, v1)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint8 = 0
        if true {
            v11 := []int{}
            v10 -= v10
            fmt.Println(fi9, v1)
            _ = v11
        } else {
            var v12 uint16 = 0
            v2 = float64(float32(v2))
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v5, v1)
    v5 = v5
    for fi13 := 0; fi13 < 7; fi13++ {
        if false {
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(v2, fi13, fi14)
                _ = fi14
            }
            var v15 uint64 = 18446744073709551615
            if (false && false) {
                v16 := GenS1{F0: -9223372036854775808, F1: 189}
                if false {
                    var v17 uint8 = 128
                    _ = v17
                } else {
                    var v18 int = (fi13 % (((-(fi13))) | 1))
                    var v19 int8 = 1
                    _ = v18
                    _ = v19
                }
                _ = v16
            } else {
                if (100 == -1) {
                    var v20 bool = true
                    for fi21 := 0; fi21 < 15; fi21++ {
                        var v22 float64 = (3.14159 * v2)
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                } else {
                    delete(v6, 65535)
                    v5 = v5
                }
                v23 := v6[65535]
                v5 += v5
                _ = v23
            }
            _ = v15
        } else {
            fmt.Println(fi13, v1, v1)
        }
        fmt.Println(v5)
        _ = fi13
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS2{F0: 2}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        v2 = 1e10
        v1 += (1.5 + (0.1 / v1))
        if false {
            switch "hello" {
            case "":
                if true {
                    v27 := GenS2{F0: 1}
                    var v28 int32 = 1
                    if true {
                        if (255 == 18446744073709551615) {
                            fmt.Println(fi26, v5, v28)
                            if ((float64((v5 + (3.14159 / (3.14159 / 1.0)))) + (v2 / v1)) < 1.5) {
                                var v29 uint16 = 32767
                                _ = v29
                            } else {
                                switch 802 {
                                case 2147483647:
                                    var v30 uint32 = 65535
                                    _ = v30
                                case 32768:
                                    v31 := GenS2{F0: 2}
                                    _ = v31
                                case 244:
                                    v32 := GenS2{F0: 752}
                                    _ = v32
                                default:
                                    var v33 int64 = 211
                                    _ = v33
                                }
                                switch "\n\t" {
                                case "":
                                    switch "foo bar" {
                                    case "hello":
                                        switch (fi26 ^ 128) {
                                        case -9223372036854775808:
                                            v34 := []int{(fi26 + (fi26 << ((14) & 63)))}
                                            _ = v34
                                        case 235:
                                            v35 := []int{fi26}
                                            _ = v35
                                        }
                                    case "":
                                        switch (-(fi26)) {
                                        case -674:
                                            v36 := GenS2{F0: 24}
                                            _ = v36
                                        case 231:
                                            v28 = v28
                                        case 32768:
                                            for fi37 := 0; fi37 < 9; fi37++ {
                                                v38 := []int{fi37, fi37, fi37, fi26}
                                                _ = fi37
                                                _ = v38
                                            }
                                        }
                                    default:
                                        var v39 float64 = (((v1 + v1) / ((v2 - float64(1e10)) + v1)) + 0.0)
                                        _ = v39
                                    }
                                case "\n\t":
                                    switch (fi26 % ((fi26) | 1)) {
                                    case 2147483647:
                                        var v40 string = "foo bar"
                                        _ = v40
                                    case 4294967295:
                                        for fi41 := 0; fi41 < 2; fi41++ {
                                            v42 := GenS3{F0: 42, F1: 1, F2: 123}
                                            _ = fi41
                                            _ = v42
                                        }
                                    case 644:
                                        switch (fi26 + fi26) {
                                        case -9223372036854775808:
                                            switch "a" {
                                            case "":
                                                v43 := GenS1{F0: (fi26 | (fi26 % ((fi26) | 1))), F1: 138}
                                                _ = v43
                                            case "\n\t":
                                                fmt.Println(v2, v28)
                                            case "foo bar":
                                                switch "a" {
                                                case "hello":
                                                    var v44 int8 = 11
                                                    _ = v44
                                                case "foo bar":
                                                    if (0 <= 2147483647) {
                                                        if (!true) {
                                                            var v45 float32 = 1.5
                                                            var v46 int = fi26
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        for fi47 := 0; fi47 < 2; fi47++ {
                                                            v1 = v2
                                                            v1 -= v1
                                                            if false {
                                                                var v48 int64 = int64(fi47)
                                                                var v49 float64 = (3.14159 / float64(v5))
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                delete(v3, 2147483647)
                                                                v6[827] = 2147483647
                                                                switch fi26 {
                                                                case 683:
                                                                    fmt.Println(v1, fi26)
                                                                case 4294967295:
                                                                    var v50 uint8 = 0
                                                                    _ = v50
                                                                }
                                                                if (17 >= 1) {
                                                                    v1 = ((((3.14159 * 1e10) * 3.14159) / (v2 - v1)) * 1e10)
                                                                } else {
                                                                    v5 += v5
                                                                }
                                                            }
                                                            _ = fi47
                                                        }
                                                        switch (^(fi26)) {
                                                        case -1:
                                                            for fi51 := 0; fi51 < 9; fi51++ {
                                                                v28 = v28
                                                                var v52 uint64 = 2
                                                                _ = fi51
                                                                _ = v52
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    v53 := GenS2{F0: 65}
                                                    _ = v53
                                                default:
                                                    v54 := GenS4{F0: 2, F1: -51, F2: fi26}
                                                    _ = v54
                                                }
                                            }
                                        }
                                    default:
                                        if false {
                                            switch (fi26 * fi26) {
                                            case 42:
                                                if true {
                                                    if true {
                                                        if (!false) {
                                                            var v55 uint8 = 42
                                                            v5 = (((v5 + (float32((float64((0.0 / v5)) + v1)) * v5)) / 0.1) + float32(v2))
                                                            if (2 > 866) {
                                                                fmt.Println(v5, fi26)
                                                                v56 := GenS5{F0: (v28 - (-(v28))), F1: (v55 ^ v55), F2: genFn1(1, 42)}
                                                                v1 = 0.1
                                                                _ = v56
                                                            }
                                                            _ = v55
                                                        }
                                                    } else {
                                                        if (!(true || true)) {
                                                            if (((v5 + ((float32(v1) * v5) * v5)) - float32(3.14159)) > float32(v2)) {
                                                                v5 -= v5
                                                                v57 := []int{fi26, (fi26 % (((fi26 % (((fi26 << ((2) & 63))) | 1))) | 1)), -447, (fi26 ^ (fi26 ^ fi26))}
                                                                if (0 < 2147483647) {
                                                                    v2 = 0.1
                                                                    delete(v3, 1)
                                                                }
                                                                _ = v57
                                                            } else {
                                                                var v58 float32 = (((v5 / (((v5 * (0.0 / (v5 * v5))) * 1e10) * 1e10)) - v5) / v5)
                                                                _ = v58
                                                            }
                                                            switch (^(fi26)) {
                                                            case -306:
                                                                var v59 uint8 = 42
                                                                _ = v59
                                                            case 796:
                                                                v28 = 1
                                                            }
                                                            switch ("" + "\n\t") {
                                                            case "hello":
                                                                v5 = v5
                                                            default:
                                                                v60 := GenS2{F0: 940}
                                                                _ = v60
                                                            }
                                                        }
                                                        if false {
                                                            var v61 int = 32767
                                                            _ = v61
                                                        } else {
                                                            v62 := v4["a"]
                                                            for fi63 := 0; fi63 < 15; fi63++ {
                                                                v64 := GenS5{F0: (-(v28)), F1: 2, F2: 100}
                                                                fmt.Println(v28, fi26)
                                                                _ = fi63
                                                                _ = v64
                                                            }
                                                            _ = v62
                                                        }
                                                        v1 += 1e10
                                                    }
                                                }
                                            case -1:
                                                delete(v6, 578)
                                            case 4294967295:
                                                v2 -= v1
                                            default:
                                                v65 := v3[831]
                                                _ = v65
                                            }
                                            delete(v6, 1)
                                        }
                                    }
                                case "a":
                                    v28 += 1
                                }
                                fmt.Println(v28, v1)
                            }
                            v66 := GenS1{F0: fi26, F1: 166}
                            _ = v66
                        } else {
                            for fi67 := 0; fi67 < 20; fi67++ {
                                v68 := GenS6{F0: "", F1: (("\n\t" + "hello") + "a")}
                                v69 := GenS2{F0: 42}
                                v70 := GenS6{F0: "", F1: ((("hello" + "foo bar") + ("\n\t" + ("\n\t" + ("\n\t" + (("hello" + "") + ("a" + (("\n\t" + "\n\t") + "a"))))))) + "")}
                                v2 = -2.25
                                _ = fi67
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                        }
                        v1 = v2
                    } else {
                        v71 := []int{}
                        for fi72 := 0; fi72 < 15; fi72++ {
                            if len(v71) > 0 {
                                v71[0] = (^(fi26))
                            }
                            var v73 bool = (169 >= 18446744073709551615)
                            _ = fi72
                            _ = v73
                        }
                        var v74 int32 = v28
                        _ = v71
                        _ = v74
                    }
                    _ = v27
                    _ = v28
                } else {
                    _ = len(v4)
                    v75 := GenS3{F0: 1000, F1: 1, F2: 127}
                    if false {
                        fmt.Println(v2, fi26, v1)
                        if (!true) {
                            var v76 uint64 = 42
                            if false {
                                var v77 float32 = -2.25
                                v78 := []int{(fi26 * fi26), (fi26 - (fi26 ^ (fi26 >> ((255) & 63)))), fi26}
                                switch 795 {
                                case -2147483648:
                                    var v79 uint16 = 32767
                                    _ = v79
                                case 127:
                                    v80 := GenS4{F0: 0, F1: -128, F2: fi26}
                                    _ = v80
                                case -9223372036854775808:
                                    var v81 uint8 = 0
                                    _ = v81
                                }
                                _ = v77
                                _ = v78
                            } else {
                                v82 := GenS6{F0: ("hello" + ("foo bar" + "a")), F1: "hello"}
                                _ = v82
                            }
                            _ = v76
                        }
                        for fi83 := 0; fi83 < 13; fi83++ {
                            v4[((("" + "hello") + "foo bar") + "a")] = false
                            _ = fi83
                        }
                    } else {
                        v84 := GenS7{F0: 42, F1: -908, F2: 206}
                        v2 = (3.14159 + -2.25)
                        var v85 string = "a"
                        _ = v84
                        _ = v85
                    }
                    _ = v75
                }
            case "\n\t":
                v1 += -2.25
            default:
                switch (("" + "hello") + "") {
                case "":
                    v5 = (v5 - 0.1)
                case "foo bar":
                    if true {
                        _ = len(v4)
                    } else {
                        v86 := GenS8{F0: 32768, F1: 482, F2: -668}
                        v87 := GenS2{F0: 2}
                        _ = v86
                        _ = v87
                    }
                case "a":
                    v5 = 3.14159
                }
            }
            var v88 int8 = 12
            if true {
                switch (fi26 + (^(fi26))) {
                case 127:
                    if false {
                        var v89 float64 = (0.0 - 1e10)
                        v1 += 1.5
                        v88 = (v88 >> ((43) & 63))
                        _ = v89
                    } else {
                        for fi90 := 0; fi90 < 13; fi90++ {
                            v5 -= 0.0
                            _ = fi90
                        }
                        for fi91 := 0; fi91 < 14; fi91++ {
                            var v92 string = "hello"
                            _ = fi91
                            _ = v92
                        }
                    }
                default:
                    v1 -= (v1 * v2)
                }
                if ((!false) || false) {
                    v93 := GenS7{F0: 505, F1: -2147483648, F2: 73}
                    if true {
                        if false {
                            var v94 int8 = v88
                            var v95 int64 = -956
                            for fi96 := 0; fi96 < 12; fi96++ {
                                v97 := GenS3{F0: 0, F1: -705, F2: 246}
                                var v98 float32 = v5
                                v95 = v95
                                _ = fi96
                                _ = v97
                                _ = v98
                            }
                            _ = v94
                            _ = v95
                        } else {
                            var v99 uint32 = 42
                            _ = v99
                        }
                        fmt.Println(v2)
                        var v100 uint16 = 2
                        _ = v100
                    } else {
                        v1 = (((((((3.14159 * float64((v5 * (-2.25 + v5)))) - v2) - (3.14159 / float64(v5))) / v2) + 0.0) / (3.14159 / v2)) / float64((float32(v2) / v5)))
                    }
                    _ = v93
                } else {
                    v5 -= v5
                }
                switch fi26 {
                case 852:
                    if false {
                        v5 = v5
                    } else {
                        fmt.Println(fi26)
                    }
                case 128:
                    if (32767 == 1000) {
                        v2 -= ((v2 * v1) + (0.1 / v2))
                        v5 = 3.14159
                        delete(v3, 1)
                    }
                case -560:
                    fmt.Println(v2, v88, v88)
                }
            }
            _ = v88
        }
        _ = fi26
    }
    var v101 float32 = (v5 - v5)
    v102 := GenS1{F0: 4294967295, F1: 9223372036854775807}
    v2 -= 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v101
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
