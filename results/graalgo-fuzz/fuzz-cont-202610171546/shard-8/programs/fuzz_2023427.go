package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (1e10 / 1.5)
    var v3 uint64 = 922
    var v4 int32 = -232
    var v5 uint64 = v3
    v6 := GenS1{F0: uint32(v5)}
    v3 -= v5
    fmt.Println(v4)
    fmt.Println(v4, v4, v3)
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (fi7 >> ((0) & 63)) {
        case 32768:
            v3 += v3
        case 334:
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := v1[fi8]
                if (0.1 == -2.25) {
                    _ = len(v1)
                    v5 -= v5
                } else {
                    var v10 uint64 = v5
                    _ = v10
                }
                var v11 int8 = 127
                _ = fi8
                _ = v11
                _ = v9
            }
        case -847:
            v12 := v1[(fi7 & fi7)]
            _ = v12
        default:
            v2 = float64((((float32(float64(float32(v2))) / 3.14159) * (0.1 + 1.5)) * (-2.25 + (3.14159 / 1.5))))
        }
        var v13 uint64 = v3
        switch "hello" {
        case "a":
            v4 += (v4 << ((699) & 63))
        default:
            var v14 string = "a"
            _ = v14
        }
        _ = fi7
        _ = v13
    }
    if (uint(v5) <= 4294967295) {
        fmt.Println(v2, v4, v2)
        if (-2.25 >= -2.25) {
            v15 := v1[32767]
            if (0 <= 162) {
                var v16 bool = ((v3 >> ((0) & 63)) >= 42)
                _ = v16
            } else {
                switch (v15 & (-(v15))) {
                case -2147483648:
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if false {
                            v2 -= float64(0.0)
                            for fi18 := 0; fi18 < 14; fi18++ {
                                v1[(v15 << ((4294967295) & 63))] = fi18
                                _ = fi18
                            }
                            var v19 float32 = ((0.0 * -2.25) + 0.0)
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 13; fi20++ {
                                v2 += (1.5 / (v2 / float64(((3.14159 - 3.14159) + 1.5))))
                                fmt.Println(v15)
                                var v21 float64 = (-2.25 - float64(1e10))
                                _ = fi20
                                _ = v21
                            }
                            v22 := GenS2{F0: 42, F1: 127}
                            v23 := GenS1{F0: 42}
                            _ = v22
                            _ = v23
                        }
                        v4 -= 826
                        for fi24 := 0; fi24 < 12; fi24++ {
                            var v25 uint16 = 42
                            _ = fi24
                            _ = v25
                        }
                        _ = fi17
                    }
                }
                v26 := GenS3{F0: 42}
                _ = v26
            }
            _ = v15
        } else {
            for fi27 := 0; fi27 < 2; fi27++ {
                if false {
                    if false {
                        v28 := GenS2{F0: 88, F1: -90}
                        v3 += 4294967295
                        if true {
                            v29 := []int{}
                            _ = v29
                        }
                        _ = v28
                    } else {
                        if true {
                            var v30 int64 = int64(fi27)
                            var v31 uint16 = 2
                            _ = v30
                            _ = v31
                        } else {
                            v2 -= v2
                        }
                    }
                    if (14 == 1) {
                        switch "foo bar" {
                        case "":
                            for fi32 := 0; fi32 < 13; fi32++ {
                                var v33 uint8 = 92
                                _ = fi32
                                _ = v33
                            }
                        }
                    } else {
                        v34 := GenS4{F0: (v5 * (v5 % ((v3) | 1))), F1: 0, F2: 1}
                        if (v3 >= (v3 << ((255) & 63))) {
                            v35 := GenS2{F0: 0, F1: -1}
                            delete(v1, 1)
                            _ = v35
                        }
                        var v36 int32 = -2147483648
                        _ = v34
                        _ = v36
                    }
                    var v37 uint = 906
                    _ = v37
                }
                v38 := GenS4{F0: (v3 / ((v5) | 1)), F1: -1, F2: 396}
                _ = fi27
                _ = v38
            }
        }
        v39 := GenS5{F0: v3, F1: ("" + ("foo bar" + "hello"))}
        _ = v39
    }
    v40 := GenS3{F0: -9223372036854775808}
    v4 += v4
    for fi41 := 0; fi41 < 14; fi41++ {
        for fi42 := 0; fi42 < 19; fi42++ {
            if (42 != 1) {
                var v43 uint16 = 1
                _ = v43
            }
            _ = fi42
        }
        switch ("foo bar" + ("foo bar" + ("\n\t" + ("\n\t" + "foo bar")))) {
        case "a":
            switch ("foo bar" + "\n\t") {
            case "a":
                v1[2147483647] = fi41
            case "hello":
                v2 = (v2 + v2)
            }
        case "":
            for fi44 := 0; fi44 < 15; fi44++ {
                fmt.Println(v2, v5, v4)
                v2 += v2
                fmt.Println(v2, v4)
                _ = fi44
            }
        }
        v45 := GenS4{F0: v3, F1: 776, F2: 909}
        _ = fi41
        _ = v45
    }
    var v46 uint64 = 1
    for fi47 := 0; fi47 < 16; fi47++ {
        switch "foo bar" {
        case "a":
            switch fi47 {
            case 42:
                _ = len(v1)
            default:
                for fi48 := 0; fi48 < 9; fi48++ {
                    fmt.Println(fi47, v4)
                    if true {
                        v2 = v2
                        v49 := []int{-425}
                        if false {
                            var v50 float64 = v2
                            v46 -= (v3 % ((4294967295) | 1))
                            v5 += v46
                            var v51 uint32 = 4294967295
                            var v52 int32 = (v4 << ((255) & 63))
                            _ = v50
                            _ = v51
                            _ = v52
                        } else {
                            switch "\n\t" {
                            case "a":
                                v1[32768] = fi47
                            case "\n\t":
                                v2 = float64(((1.5 + 3.14159) * 0.1))
                            case "foo bar":
                                v3 -= 1
                            default:
                                var v53 float32 = 1e10
                                var v54 uint16 = 42
                                _ = v53
                                _ = v54
                            }
                            var v55 string = "\n\t"
                            _ = len(v1)
                            switch (fi48 * (fi48 % ((int(v4)) | 1))) {
                            case 436:
                                switch v55 {
                                case "":
                                    v46 = 4294967295
                                case "hello":
                                    fmt.Println(v3, fi48)
                                case "foo bar":
                                    fmt.Println(v2, v55)
                                }
                                var v56 int = fi48
                                v57 := GenS6{F0: -626}
                                _ = v56
                                _ = v57
                            case 1:
                                switch fi47 {
                                case -9223372036854775808:
                                    v58 := GenS7{F0: 42, F1: "\n\t"}
                                    _ = v58
                                default:
                                    var v59 string = v55
                                    _ = v59
                                }
                            case -9223372036854775808:
                                v5 += 1
                            default:
                                fmt.Println(fi47)
                            }
                            fmt.Println(fi48, fi47)
                            v55 = (v55 + v55)
                            _ = v55
                        }
                        _ = v49
                    }
                    var v60 int16 = 338
                    _ = fi48
                    _ = v60
                }
            }
        default:
            v4 = 0
        }
        for fi61 := 0; fi61 < 13; fi61++ {
            v62 := GenS8{F0: 4294967295}
            if (!(!false)) {
                var v63 int = 2147483647
                switch v63 {
                case -367:
                    v64 := GenS6{F0: -9223372036854775808}
                    _ = v64
                case 420:
                    var v65 bool = ((4294967295 > 9223372036854775807) && true)
                    _ = v65
                }
                _ = v63
            } else {
                v3 = (v5 / ((v5) | 1))
                v3 = (v5 | (v5 * v5))
                v66 := []int{(fi61 - fi61), (fi47 & fi61)}
                _ = v66
            }
            if (!(v4 != (v4 / ((v4) | 1)))) {
                v2 = v2
                var v67 int8 = 0
                switch fi47 {
                case 32767:
                    v5 += v5
                default:
                    v68 := GenS4{F0: (^(v46)), F1: 9223372036854775807, F2: 230}
                    _ = v68
                }
                _ = v67
            } else {
                v69 := []int{fi61, (fi61 | (fi61 & -147))}
                _ = v69
            }
            v2 = 1.5
            var v70 int64 = int64(fi47)
            _ = fi61
            _ = v62
            _ = v70
        }
        _ = len(v1)
        _ = fi47
    }
    if (!true) {
        if ((v3 % (((v3 | 2)) | 1)) == 0) {
            v1[406] = -1
            fmt.Println(v46)
        } else {
            v71 := GenS5{F0: v46, F1: ""}
            _ = v71
        }
        fmt.Println(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v5
    _ = v6
}
