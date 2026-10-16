package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 bool; F2 uint64 }
func genFn0() bool {
    var v4 int8 = 1
    v5 := []int{}
    v4 += (v4 - (-(v4)))
    v4 = (v4 % ((-79) | 1))
    v6 := GenS1{F0: 34, F1: 32767}
    _ = v4
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint16, p1 uint) bool {
    if (true && true) {
        v7 := GenS1{F0: 2, F1: 65535}
        var v8 uint64 = 204
        _ = v7
        _ = v8
    }
    switch 128 {
    case -879:
        switch -920 {
        case -2147483648:
            var v9 float32 = -2.25
            _ = v9
        case 32768:
            var v10 int = 834
            _ = v10
        default:
            v11 := []int{888, 42, 127}
            _ = v11
        }
    }
    v12 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = v12
    return false
}

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 0
    v3 := []int{-550}
    if (!(false && (false && (float32(1e10) >= 1e10)))) {
        if (9223372036854775807 < 32768) {
            var v13 bool = false
            for fi14 := 0; fi14 < 12; fi14++ {
                if len(v3) > 0 {
                    v3[0] = fi14
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi14
            }
            _ = v13
        } else {
            var v15 float64 = 0.1
            v3 = append(v3, 42)
            if (!true) {
                var v16 bool = ((v1 / ((2) | 1)) < 1)
                _ = v16
            } else {
                switch -496 {
                case 1:
                    var v17 uint64 = 4294967295
                    _ = v17
                case -9223372036854775808:
                    fmt.Println(v1, v15)
                case 32768:
                    if true {
                        if false {
                            var v18 string = "hello"
                            for fi19 := 0; fi19 < 8; fi19++ {
                                if len(v3) > 0 {
                                    v3[0] = fi19
                                }
                                v20 := GenS3{F0: -1, F1: v1}
                                if true {
                                    fmt.Println(fi19)
                                }
                                _ = fi19
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 20; fi21++ {
                                v22 := GenS4{F0: 4294967295, F1: false}
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    v2 += 65535
                                    _ = fi23
                                }
                                v24 := GenS2{F0: 671}
                                _ = fi21
                                _ = v22
                                _ = v24
                            }
                            _ = v18
                        } else {
                            v25 := GenS3{F0: 42, F1: 253}
                            for fi26 := 0; fi26 < 10; fi26++ {
                                var v27 int8 = 42
                                _ = fi26
                                _ = v27
                            }
                            if (0 >= 9223372036854775807) {
                                v28 := GenS3{F0: -923, F1: 230}
                                _ = v28
                            }
                            _ = v25
                        }
                    } else {
                        v2 = (v2 % (((v2 / (((^(v2))) | 1))) | 1))
                    }
                }
                v29 := GenS4{F0: 4294967295, F1: false}
                _ = v29
            }
            _ = v15
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                v32 := GenS2{F0: uint32(v2)}
                switch fi31 {
                case 114:
                    v2 = 0
                case 447:
                    fmt.Println(fi31)
                case 4294967295:
                    switch fi30 {
                    case -164:
                        var v33 int8 = int8(fi30)
                        _ = v33
                    default:
                        v34 := GenS5{F0: (("\n\t" + "hello") + ((("hello" + "") + "") + (((((("" + ("\n\t" + ("hello" + "hello"))) + "hello") + (("foo bar" + "\n\t") + "foo bar")) + "a") + "a") + ((("hello" + "") + ("" + "hello")) + ("hello" + ("foo bar" + (((("foo bar" + "foo bar") + ("a" + "foo bar")) + "\n\t") + ("hello" + "hello")))))))), F1: (fi30 * fi31)}
                        _ = v34
                    }
                default:
                    var v35 int32 = -2147483648
                    _ = v35
                }
                switch "" {
                case "foo bar":
                    v2 -= v2
                case "a":
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 = 0
                        _ = len(v3)
                    } else {
                        v2 = 65535
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v36 string = "\n\t"
                        _ = v36
                    }
                case "hello":
                    var v37 int8 = -128
                    _ = v37
                }
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
    }
    if (255 < 322) {
        fmt.Println(v1)
    }
    _ = len(v3)
    if (4294967295 < 612) {
        var v38 float64 = 3.14159
        v39 := GenS6{F0: 0, F1: uint8(v1)}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v38
        _ = v39
    } else {
        if true {
            v40 := GenS7{F0: v2}
            var v41 int16 = 32767
            _ = v40
            _ = v41
        } else {
            switch "" {
            case "hello":
                switch "\n\t" {
                case "foo bar":
                    switch 0 {
                    case 456:
                        fmt.Println(v1)
                    case 796:
                        v2 = uint16(v1)
                    case 42:
                        v2 -= (v2 - (v2 - (v2 + 32767)))
                    default:
                        for fi42 := 0; fi42 < 3; fi42++ {
                            var v43 uint = 42
                            switch (fi42 | 92) {
                            case 914:
                                var v44 uint8 = 255
                                _ = v44
                            case 127:
                                v1 += (v1 ^ (^(v1)))
                            case -1:
                                _ = len(v3)
                            default:
                                v3 = append(v3, fi42)
                            }
                            _ = fi42
                            _ = v43
                        }
                    }
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                case "hello":
                    if true {
                        if (v1 < v1) {
                            for fi45 := 0; fi45 < 10; fi45++ {
                                for fi46 := 0; fi46 < 17; fi46++ {
                                    switch "" {
                                    case "\n\t":
                                        var v47 uint32 = 4294967295
                                        _ = v47
                                    case "a":
                                        var v48 uint64 = v1
                                        _ = v48
                                    }
                                    var v49 float32 = float32(3.14159)
                                    var v50 int8 = -43
                                    _ = fi46
                                    _ = v49
                                    _ = v50
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    var v51 int32 = 587
                                    _ = v51
                                case "a":
                                    var v52 int64 = 42
                                    _ = v52
                                default:
                                    var v53 int8 = 1
                                    _ = v53
                                }
                                _ = len(v3)
                                _ = fi45
                            }
                        }
                        fmt.Println(v2, v2)
                        var v54 int8 = -1
                        _ = v54
                    } else {
                        v1 -= (v1 / (((v1 + v1)) | 1))
                        switch "a" {
                        case "a":
                            var v55 int = 32768
                            _ = v55
                        default:
                            v56 := GenS4{F0: 0, F1: (true && true)}
                            _ = v56
                        }
                    }
                default:
                    var v57 int8 = -128
                    _ = v57
                }
            default:
                var v58 int64 = 1
                _ = v58
            }
            if len(v3) > 0 {
                v3[0] = -1
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v1, v2)
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v59 uint64 = 2
    for fi60 := 0; fi60 < 16; fi60++ {
        switch fi60 {
        case 2147483647:
            fmt.Println(fi60, fi60, v59)
        case 922:
            var v61 float64 = 0.0
            _ = v61
        default:
            var v62 int32 = -2147483648
            _ = v62
        }
        fmt.Println(v59)
        fmt.Println(fi60, v1, v2)
        _ = fi60
    }
    v63 := GenS4{F0: 520, F1: false}
    for fi64 := 0; fi64 < 12; fi64++ {
        v59 -= v1
        var v65 uint64 = 42
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi66 := 0; fi66 < 16; fi66++ {
                v67 := GenS8{F0: v1, F1: false, F2: (v65 % (((v1 % ((1) | 1))) | 1))}
                _ = fi66
                _ = v67
            }
        }
        _ = fi64
        _ = v65
    }
    for fi68 := 0; fi68 < 7; fi68++ {
        v1 += 18446744073709551615
        v69 := GenS2{F0: 4294967295}
        var v70 int16 = 22
        _ = fi68
        _ = v69
        _ = v70
    }
    if true {
        fmt.Println(v2, v2, v2)
        if true {
            for fi71 := 0; fi71 < 13; fi71++ {
                v3 = append(v3, fi71)
                _ = fi71
            }
            if false {
                if false {
                    switch -41 {
                    case -20:
                        switch 42 {
                        case -2147483648:
                            var v72 uint64 = 0
                            _ = v72
                        case 2147483647:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -674:
                            var v73 uint16 = 32767
                            _ = v73
                        }
                    case 0:
                        v74 := GenS2{F0: 65535}
                        _ = v74
                    }
                }
            }
            var v75 uint16 = 2
            _ = v75
        }
    }
    var v76 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v59:", v59)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v63
    _ = v76
}
