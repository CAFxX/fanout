package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int16; F1 uint64; F2 uint }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int32 }
func genFn0() int {
    if (1 == 0) {
        switch "" {
        case "hello":
            v7 := []int{104, 32768, -84, -299}
            _ = v7
        case "\n\t":
            if (42 < 2) {
                var v8 string = ((((("\n\t" + "a") + "hello") + "foo bar") + "foo bar") + "\n\t")
                v8 = v8
                _ = v8
            } else {
                var v9 int = 9223372036854775807
                v9 += v9
                var v10 uint16 = 255
                _ = v10
                _ = v9
            }
        }
        var v11 int16 = -32768
        _ = v11
    }
    v12 := []int{276, -358, -9223372036854775808}
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 bool = true
        var v15 uint8 = 255
        if ("\n\t" != ((("foo bar" + ("a" + "")) + ("" + "")) + "\n\t")) {
            if v14 {
                for fi16 := 0; fi16 < 18; fi16++ {
                    if len(v12) > 0 {
                        v12[0] = fi13
                    }
                    v15 = v15
                    _ = fi16
                }
                if true {
                    fmt.Println(v15, fi13, v15)
                } else {
                    if (1 < -102) {
                        switch "hello" {
                        case "hello":
                            switch "a" {
                            case "hello":
                                if (("hello" + "a") >= (("foo bar" + (("foo bar" + "hello") + "")) + ("" + ("foo bar" + "a")))) {
                                    fmt.Println(v15)
                                    if (2147483647 < 4294967295) {
                                        var v17 int16 = -623
                                        if (324 == 0) {
                                            var v18 float32 = 3.14159
                                            if len(v12) > 0 {
                                                v12[0] = (fi13 * (fi13 | fi13))
                                            }
                                            if (v14 || v14) {
                                                v19 := GenS1{F0: 789, F1: 948}
                                                if (0 == 0) {
                                                    for fi20 := 0; fi20 < 12; fi20++ {
                                                        v21 := GenS1{F0: 125, F1: 870}
                                                        var v22 float64 = ((3.14159 / float64(-2.25)) + -2.25)
                                                        _ = fi20
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                    v17 -= (^(v17))
                                                    fmt.Println(v15, v15, v15)
                                                }
                                                var v23 uint = 110
                                                _ = v19
                                                _ = v23
                                            } else {
                                                v24 := GenS2{F0: v15, F1: 2, F2: 71}
                                                var v25 bool = (1.5 < -2.25)
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = v18
                                        } else {
                                            fmt.Println(v15)
                                            for fi26 := 0; fi26 < 14; fi26++ {
                                                _ = len(v12)
                                                switch "foo bar" {
                                                case "":
                                                    v27 := GenS3{F0: 65535}
                                                    _ = v27
                                                }
                                                v15 = v15
                                                _ = fi26
                                            }
                                            v17 = (^(v17))
                                        }
                                        _ = v17
                                    } else {
                                        fmt.Println(v15, fi13)
                                    }
                                    var v28 int32 = 1
                                    _ = v28
                                } else {
                                    if true {
                                        v15 = (v15 - v15)
                                        _ = len(v12)
                                        v15 = (v15 / ((0) | 1))
                                    } else {
                                        _ = len(v12)
                                    }
                                }
                            case "foo bar":
                                var v29 float32 = -2.25
                                _ = v29
                            }
                        case "foo bar":
                            for fi30 := 0; fi30 < 19; fi30++ {
                                if v14 {
                                    _ = len(v12)
                                    if (uint16(v15) <= 2) {
                                        v31 := GenS2{F0: 127, F1: 4294967295, F2: (fi30 + (fi13 * (fi13 - (fi30 * fi13))))}
                                        _ = v31
                                    } else {
                                        var v32 bool = true
                                        if len(v12) > 0 {
                                            v12[0] = 907
                                        }
                                        v15 = v15
                                        _ = v32
                                    }
                                }
                                v14 = (4294967295 <= 2147483648)
                                _ = fi30
                            }
                        }
                    }
                }
            } else {
                v15 = (v15 % ((v15) | 1))
                for fi33 := 0; fi33 < 16; fi33++ {
                    if len(v12) > 0 {
                        v12[0] = 2147483647
                    }
                    _ = fi33
                }
            }
            var v34 float32 = 1.5
            _ = v34
        } else {
            for fi35 := 0; fi35 < 7; fi35++ {
                var v36 bool = true
                switch "" {
                case "a":
                    switch fi13 {
                    case 1:
                        v14 = true
                    case 42:
                        fmt.Println(v14)
                        v37 := GenS3{F0: 65535}
                        _ = v37
                    }
                }
                _ = fi35
                _ = v36
            }
            if len(v12) > 0 {
                v12[0] = fi13
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    if true {
        v38 := GenS3{F0: 522}
        var v39 uint32 = 674
        _ = v38
        _ = v39
    }
    _ = len(v12)
    _ = v12
    return -2147483648
}

func main() {
    v1 := []uint32{}
    var v2 string = "\n\t"
    v3 := []uint32{65535, 4294967295}
    var v4 int16 = -32768
    var v5 int16 = v4
    var v6 string = (v2 + v2)
    v1 = append(v1, 617)
    var v40 uint64 = 42
    var v41 uint8 = 2
    v6 = v2
    var v42 uint16 = 32767
    v5 = v5
    v41 -= (v41 % ((v41) | 1))
    v43 := GenS4{F0: -976}
    switch v6 {
    case "a":
        var v44 int = 32767
        _ = v44
    case "hello":
        if ((1.5 - 1.5) == (-2.25 - 3.14159)) {
            if ((633 < 2) || (-2.25 != 0.1)) {
                var v45 uint32 = 65535
                _ = v45
            }
            v4 = int16(v42)
            v42 = 431
        } else {
            if ((^(v42)) != (^(v42))) {
                fmt.Println(v42)
                for fi46 := 0; fi46 < 11; fi46++ {
                    v1 = append(v1, uint32(v40))
                    for fi47 := 0; fi47 < 10; fi47++ {
                        switch genFn0() {
                        case -705:
                            fmt.Println(v42)
                        case -2147483648:
                            if true {
                                v48 := GenS3{F0: 42}
                                v6 = ("" + ((v2 + v2) + (v6 + "a")))
                                if false {
                                    if ((0.1 - (((((0.0 + ((1e10 + (0.1 * 3.14159)) / 0.1)) / 1.0) + 1.5) + (0.0 + float32(1e10))) + (3.14159 / -2.25))) <= float32(1.5)) {
                                        if true {
                                            v49 := GenS5{F0: 1, F1: v2}
                                            _ = v49
                                        } else {
                                            if false {
                                                var v50 int64 = 9223372036854775807
                                                switch (("foo bar" + "") + v2) {
                                                case "a":
                                                    v51 := GenS6{F0: (^(v5)), F1: uint64(v41), F2: 2}
                                                    _ = v51
                                                }
                                                v52 := GenS2{F0: v41, F1: 928, F2: fi46}
                                                _ = v50
                                                _ = v52
                                            } else {
                                                if ((float64(((((float32(-2.25) + float32((0.1 - ((float64(0.1) - (0.0 / 1e10)) - -2.25)))) + 3.14159) / 1.5) / 1.0)) + 1e10) >= 0.0) {
                                                    var v53 bool = false
                                                    _ = v53
                                                } else {
                                                    _ = len(v1)
                                                    v4 = 42
                                                }
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v54 := GenS6{F0: (v5 & -1), F1: v40, F2: 4294967295}
                                            _ = v54
                                        }
                                        v2 = v2
                                    } else {
                                        var v55 float32 = (0.0 * 3.14159)
                                        v4 -= -276
                                        _ = v55
                                    }
                                }
                                _ = v48
                            }
                        case -493:
                            fmt.Println(fi47, v41, v5)
                        }
                        if ((((0.0 + 0.1) + (float64(1.5) + 1e10)) - float64(1.5)) > 0.1) {
                            v56 := GenS2{F0: (v41 * v41), F1: 2147483648, F2: fi47}
                            var v57 uint = uint(v40)
                            _ = v56
                            _ = v57
                        } else {
                            var v58 int32 = -2147483648
                            v59 := GenS1{F0: 65535, F1: uint(v42)}
                            if (v41 != 128) {
                                _ = len(v1)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v58
                            _ = v59
                        }
                        if ((1 >= 216) || (!true)) {
                            switch v6 {
                            case "\n\t":
                                var v60 int8 = int8(v42)
                                _ = v60
                            case "foo bar":
                                switch fi47 {
                                case 127:
                                    v42 += 42
                                case 2147483647:
                                    _ = len(v3)
                                default:
                                    switch (fi46 >> ((18446744073709551615) & 63)) {
                                    case 9223372036854775807:
                                        var v61 float64 = -2.25
                                        _ = v61
                                    default:
                                        var v62 float64 = 1e10
                                        _ = v62
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                }
                            default:
                                v42 = (v42 | v42)
                            }
                        } else {
                            v42 -= 65535
                            if false {
                                if ((v4 >> ((1) & 63)) <= (^(v5))) {
                                    var v63 int32 = -63
                                    var v64 uint32 = 235
                                    _ = v63
                                    _ = v64
                                } else {
                                    switch v6 {
                                    case "foo bar":
                                        if true {
                                            v41 -= (v41 << ((433) & 63))
                                            switch fi46 {
                                            case -9223372036854775808:
                                                var v65 float64 = 0.0
                                                _ = v65
                                            case 268:
                                                var v66 bool = (fi46 < fi46)
                                                _ = v66
                                            case 42:
                                                v4 += -1
                                            default:
                                                v6 = v6
                                            }
                                            var v67 string = "hello"
                                            _ = v67
                                        } else {
                                            var v68 uint32 = 2147483648
                                            v68 = v68
                                            _ = v68
                                        }
                                    case "":
                                        var v69 uint64 = v40
                                        _ = v69
                                    case "\n\t":
                                        var v70 float64 = (((0.1 / (3.14159 - 0.1)) - ((1.5 + (0.0 * -2.25)) + float64(0.1))) * (float64(1.5) + ((((3.14159 - (1e10 / (-2.25 - (0.1 + 1.5)))) - 3.14159) + (3.14159 - 3.14159)) / 1.0)))
                                        _ = v70
                                    default:
                                        if (v40 > (^(v40))) {
                                            v41 += 200
                                        } else {
                                            var v71 int8 = 42
                                            _ = v71
                                        }
                                    }
                                }
                                var v72 uint64 = 2
                                _ = v72
                            }
                        }
                        if (0 < 1) {
                            v3 = append(v3, 567)
                        }
                        _ = fi47
                    }
                    if ((v41 << ((42) & 63)) != (v41 - (v41 / ((v41) | 1)))) {
                        v6 = (v6 + "foo bar")
                    } else {
                        var v73 int64 = -1
                        _ = v73
                    }
                    _ = fi46
                }
            } else {
                for fi74 := 0; fi74 < 8; fi74++ {
                    if false {
                        v1 = append(v1, 0)
                        if true {
                            v75 := GenS7{F0: 1, F1: 0}
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                            v76 := GenS3{F0: 65535}
                            _ = v75
                            _ = v76
                        }
                    } else {
                        var v77 int = fi74
                        v40 = v40
                        _ = len(v3)
                        _ = v77
                    }
                    v2 = (v6 + (("hello" + v6) + "hello"))
                    _ = fi74
                }
            }
        }
    default:
        switch (("hello" + v2) + "") {
        case "foo bar":
            switch "" {
            case "\n\t":
                fmt.Println(v42, v6)
            case "hello":
                var v78 uint = 255
                _ = v78
            case "":
                var v79 int8 = -1
                _ = v79
            default:
                v80 := GenS6{F0: (v4 + 0), F1: (^(v40)), F2: 1}
                _ = v80
            }
        case "hello":
            v2 = v2
        default:
            v4 = v4
        }
    }
    var v81 uint = 240
    fmt.Println(v42)
    for fi82 := 0; fi82 < 13; fi82++ {
        v41 = (v41 % ((34) | 1))
        _ = fi82
    }
    switch ((v6 + "foo bar") + (v2 + ("hello" + v6))) {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v83 := GenS8{F0: 2147483647}
    if ((false || ((v42 / ((v42) | 1)) < v42)) || true) {
        v81 = (^(v81))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v81
    _ = v83
}
