package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int16; F1 bool; F2 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 bool; F2 int16 }
type GenS7 struct { F0 int8; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = -2.25
    v3 := make(map[uint32]string)
    var v4 bool = false
    v5 := make(map[uint32]uint32)
    var v6 int8 = 52
    var v7 int32 = -556
    v7 -= 2147483647
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{(fi8 / (((fi8 - (-(fi8)))) | 1))}
        fmt.Println(v7, fi8)
        if (2 > 0) {
            v2 += -2.25
        } else {
            v2 -= (((v2 + (float64(-2.25) - v2)) / ((v2 * (v2 / (((v2 + 0.1) * (v2 / v2)) * (v2 / (v2 * (v2 / (v2 + 1.5))))))) / (v2 * 3.14159))) / -2.25)
        }
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 1
    v11 := GenS1{F0: (^(v10)), F1: 796}
    v10 -= 1
    v12 := GenS1{F0: v10, F1: -1}
    var v13 float64 = v2
    for fi14 := 0; fi14 < 4; fi14++ {
        if true {
            if false {
                var v15 uint = 4294967295
                _ = len(v3)
                _ = v15
            } else {
                v16 := []int{}
                _ = v16
            }
            v7 += (v7 / ((-732) | 1))
        }
        _ = fi14
    }
    if ((!true) && (v4 || (uint(v10) < 0))) {
        for fi17 := 0; fi17 < 9; fi17++ {
            v2 = v13
            v18 := []int{(fi17 % (((fi17 - fi17)) | 1)), 1, (fi17 - (fi17 * (fi17 ^ fi17)))}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 int8 = (v6 | v6)
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 float32 = -2.25
                fmt.Println(v4)
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        if false {
            for fi23 := 0; fi23 < 17; fi23++ {
                v24 := []int{-38, (fi23 >> ((600) & 63)), 2147483647, 466}
                if true {
                    fmt.Println(v7, fi23)
                    v25 := GenS1{F0: v10, F1: 1000}
                    v7 -= (-(v7))
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
            fmt.Println(v4, v6, v13)
            var v26 uint16 = 139
            _ = v26
        } else {
            switch "foo bar" {
            case "":
                v27 := GenS1{F0: (v10 & v10), F1: 1}
                _ = v27
            case "a":
                v28 := GenS1{F0: v10, F1: 1}
                _ = v28
            case "\n\t":
                switch 513 {
                case 973:
                    fmt.Println(v2, v4, v7)
                case 32768:
                    v29 := GenS2{F0: 82}
                    _ = v29
                default:
                    if (2 < 289) {
                        v10 = (^(v10))
                        var v30 int64 = 1
                        _ = v30
                    } else {
                        v2 -= v2
                        fmt.Println(v6, v6, v4)
                        for fi31 := 0; fi31 < 5; fi31++ {
                            fmt.Println(v4)
                            _ = fi31
                        }
                    }
                }
            default:
                v32 := []int{-450, 2147483647, -2147483648, 4294967295}
                _ = v32
            }
            switch "" {
            case "a":
                for fi33 := 0; fi33 < 6; fi33++ {
                    if ((!v4) || (852 == 1)) {
                        if false {
                            switch (("" + "\n\t") + ((("" + "foo bar") + ("foo bar" + "foo bar")) + "a")) {
                            case "foo bar":
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    v35 := []int{fi34, (fi34 << ((4294967295) & 63))}
                                    switch "a" {
                                    case "\n\t":
                                        v10 = v10
                                    case "hello":
                                        if (true && (891 == 1)) {
                                            fmt.Println(v7)
                                            v7 -= (v7 + v7)
                                            switch (fi34 & (fi33 + (fi33 % ((fi34) | 1)))) {
                                            case 70:
                                                fmt.Println(fi34)
                                            case 128:
                                                v36 := GenS3{F0: 2, F1: uint8(v10)}
                                                _ = v36
                                            case -81:
                                                v5[2] = 2147483647
                                            default:
                                                delete(v1, ("\n\t" + "a"))
                                                v2 = float64(0.0)
                                            }
                                            v37 := GenS4{F0: int16(v10), F1: v4, F2: fi34}
                                            if v4 {
                                                switch "" {
                                                case "foo bar":
                                                    v38 := GenS5{F0: 179}
                                                    _ = v38
                                                case "\n\t":
                                                    if v4 {
                                                        v39 := GenS1{F0: 42, F1: 42}
                                                        if len(v35) > 0 {
                                                            v35[0] = fi34
                                                        }
                                                        fmt.Println(v6, v6)
                                                        _ = v39
                                                    }
                                                case "a":
                                                    switch (fi34 - -226) {
                                                    case 32768:
                                                        fmt.Println(v6, fi34)
                                                        var v40 int16 = -32768
                                                        switch fi34 {
                                                        case 127:
                                                            fmt.Println(v10)
                                                        default:
                                                            v13 += (1e10 * -2.25)
                                                        }
                                                        _ = v40
                                                    case 225:
                                                        var v41 uint16 = 42
                                                        _ = v41
                                                    case 143:
                                                        v10 -= (v10 >> ((2) & 63))
                                                    default:
                                                        if v4 {
                                                            switch "a" {
                                                            case "hello":
                                                                var v42 bool = v4
                                                                _ = v42
                                                            }
                                                        } else {
                                                            var v43 int = fi34
                                                            switch (v43 % (((v43 * v43)) | 1)) {
                                                            case 186:
                                                                var v44 int64 = -1
                                                                _ = v44
                                                            default:
                                                                v45 := GenS5{F0: 2147483648}
                                                                _ = v45
                                                            }
                                                            _ = v43
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(fi34, fi33, v2)
                                                }
                                                var v46 float64 = v13
                                                if ((v6 & v6) <= 42) {
                                                    if len(v35) > 0 {
                                                        v35[0] = (fi33 % ((228) | 1))
                                                    }
                                                }
                                                _ = v46
                                            } else {
                                                var v47 int = fi33
                                                _ = v47
                                            }
                                            _ = v37
                                        } else {
                                            var v48 uint8 = 1
                                            _ = v48
                                        }
                                    }
                                    v4 = v4
                                    v4 = v4
                                    v49 := GenS2{F0: 173}
                                    var v50 bool = (373 >= 1)
                                    _ = fi34
                                    _ = v35
                                    _ = v49
                                    _ = v50
                                }
                            case "hello":
                                var v51 uint32 = 1
                                _ = v51
                            }
                            v52 := GenS6{F0: 68, F1: (638 > (v10 / (((^(v10))) | 1))), F2: 175}
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 14; fi53++ {
                            fmt.Println(v7)
                            v6 += (^(v6))
                            v10 += (v10 >> ((396) & 63))
                            _ = fi53
                        }
                    } else {
                        var v54 int32 = 0
                        switch 127 {
                        case 4294967295:
                            v55 := GenS2{F0: 0}
                            _ = v55
                        }
                        _ = v54
                    }
                    if v4 {
                        fmt.Println(v13, v2, v6)
                        v56 := GenS7{F0: (v6 >> ((450) & 63)), F1: (v6 / ((71) | 1)), F2: 255}
                        _ = v56
                    }
                    if false {
                        fmt.Println(v6, v2)
                    }
                    _ = fi33
                }
            case "foo bar":
                v2 = (3.14159 / v2)
            }
            v3[65535] = (("a" + ((("" + "a") + "\n\t") + "")) + "a")
        }
    }
    v7 = (v7 | -1)
    var v57 bool = (!(v7 < v7))
    delete(v1, "")
    switch 2147483647 {
    case 857:
        switch 1 {
        case -745:
            var v58 uint = 42
            _ = v58
        case 963:
            fmt.Println(v7, v7, v7)
        case 127:
            v10 += (v10 | (v10 * v10))
        }
    case -424:
        if false {
            if (1 > 1) {
                for fi59 := 0; fi59 < 1; fi59++ {
                    v60 := GenS3{F0: 65535, F1: 127}
                    _ = fi59
                    _ = v60
                }
            } else {
                fmt.Println(v13, v4)
                v61 := []int{32767}
                if ((v10 * v10) <= v10) {
                    _ = len(v1)
                    fmt.Println(v6, v7, v10)
                }
                _ = v61
            }
            fmt.Println(v7)
            if (145 >= 2) {
                v2 = (3.14159 - 3.14159)
                if (!(42 < -1)) {
                    fmt.Println(v7, v57)
                }
                var v62 uint32 = 4294967295
                _ = v62
            }
        }
    }
    var v63 int32 = v7
    for fi64 := 0; fi64 < 5; fi64++ {
        for fi65 := 0; fi65 < 18; fi65++ {
            fmt.Println(v13, v7)
            var v66 uint8 = 69
            _ = fi65
            _ = v66
        }
        var v67 bool = v57
        var v68 int = (fi64 + fi64)
        _ = fi64
        _ = v67
        _ = v68
    }
    var v69 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v63
    _ = v69
    _ = v7
}
