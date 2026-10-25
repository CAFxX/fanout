package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16; F1 int64; F2 int32 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 int16 }
func genFn0(p0 uint16, p1 int) int8 {
    v7 := GenS1{F0: 109, F1: -128}
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{(^(fi8))}
        _ = fi8
        _ = v9
    }
    if (0 == p1) {
        v10 := []int{p1, p1}
        _ = v10
    } else {
        var v11 int64 = -1
        var v12 int32 = 2147483647
        p0 = 996
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 128, F1: 127}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return -128
}

func main() {
    v1 := []string{"", ("hello" + (("a" + "foo bar") + ""))}
    v2 := []uint32{4294967295, 2147483648, 919, 2147483648}
    var v3 int = -1
    var v4 float32 = float32((((-2.25 - 1.5) - 1e10) + (0.1 + ((1.5 - float64(3.14159)) / (((0.1 - float64(0.1)) + -2.25) - 0.1)))))
    v5 := []uint32{1, 2147483647, 42, 2147483647}
    var v6 uint32 = 602
    var v14 uint8 = 2
    if true {
        v15 := GenS3{F0: -32768, F1: -9223372036854775808, F2: -634}
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                v3 = v3
                v6 += 4294967295
                _ = fi17
            }
            v18 := GenS2{F0: 42, F1: 42}
            _ = fi16
            _ = v18
        }
        var v19 uint64 = 333
        _ = v15
        _ = v19
    } else {
        var v20 uint8 = v14
        _ = v20
    }
    v21 := GenS4{F0: -508, F1: uint64(v6)}
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (false && (v4 >= (0.0 / float32(((3.14159 / (-2.25 + 0.0)) - 0.0))))) {
            switch v3 {
            case -1:
                fmt.Println(v4)
            case 78:
                var v22 float32 = v4
                _ = v22
            default:
                for fi23 := 0; fi23 < 1; fi23++ {
                    v6 += v6
                    fmt.Println(fi23)
                    var v24 int64 = 9223372036854775807
                    _ = fi23
                    _ = v24
                }
            }
            var v25 string = "a"
            _ = v25
        }
        var v26 float64 = 0.1
        _ = v26
    } else {
        var v27 int8 = 127
        _ = v27
    }
    v28 := GenS3{F0: 1, F1: -1, F2: 0}
    if true {
        for fi29 := 0; fi29 < 4; fi29++ {
            switch "\n\t" {
            case "":
                var v30 uint8 = (v14 ^ v14)
                _ = v30
            case "\n\t":
                v14 = 181
            default:
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 int32 = 0
                    var v33 float64 = (0.0 - 1.5)
                    var v34 uint32 = v6
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            }
            switch (("hello" + "a") + "\n\t") {
            case "foo bar":
                v3 -= (v3 * 128)
            case "hello":
                if (!(42 != 0)) {
                    var v35 bool = true
                    for fi36 := 0; fi36 < 12; fi36++ {
                        if (0 == 1) {
                            if v35 {
                                fmt.Println(fi36, fi36, v35)
                                var v37 int32 = 911
                                v38 := GenS5{F0: 1}
                                _ = v37
                                _ = v38
                            } else {
                                v3 = 245
                                v39 := GenS3{F0: 32767, F1: 9223372036854775807, F2: 2147483647}
                                var v40 float32 = 1.5
                                _ = v39
                                _ = v40
                            }
                        }
                        _ = fi36
                    }
                    _ = v35
                }
            }
            _ = fi29
        }
    }
    var v41 int32 = 2147483647
    var v42 int16 = -1
    for fi43 := 0; fi43 < 5; fi43++ {
        fmt.Println(v14, v42, fi43)
        _ = fi43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        if true {
            if true {
                var v45 uint = 42
                var v46 string = "hello"
                var v47 uint8 = v14
                _ = v45
                _ = v46
                _ = v47
            }
            for fi48 := 0; fi48 < 18; fi48++ {
                v49 := GenS1{F0: 100, F1: -128}
                if len(v5) > 0 {
                    v5[0] = (v6 ^ 871)
                }
                if true {
                    switch fi48 {
                    case 4294967295:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case -948:
                        v42 += 42
                    case -1:
                        if (v3 <= (fi48 ^ (fi48 << ((0) & 63)))) {
                            if false {
                                if (!true) {
                                    fmt.Println(v14)
                                    v4 = v4
                                }
                                var v50 int8 = 31
                                var v51 float64 = 0.0
                                _ = v50
                                _ = v51
                            }
                        } else {
                            var v52 uint32 = v6
                            _ = v52
                        }
                    default:
                        if false {
                            _ = len(v1)
                            var v53 int16 = 963
                            _ = v53
                        }
                    }
                    if true {
                        var v54 float64 = float64(v4)
                        switch 32768 {
                        case 32767:
                            var v55 float32 = ((v4 * float32(1.5)) * v4)
                            _ = v55
                        case -571:
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                v6 = v6
                            } else {
                                var v56 int8 = 127
                                var v57 bool = true
                                _ = len(v5)
                                _ = v56
                                _ = v57
                            }
                        default:
                            if len(v2) > 0 {
                                v2[0] = v6
                            }
                        }
                        switch "a" {
                        case "a":
                            var v58 bool = (false || (0 <= (v14 >> ((165) & 63))))
                            _ = v58
                        case "hello":
                            var v59 int8 = genFn0(1, (v3 / ((fi48) | 1)))
                            _ = v59
                        }
                        _ = v54
                    }
                }
                _ = fi48
                _ = v49
            }
        }
        switch "\n\t" {
        case "\n\t":
            v4 -= v4
        case "":
            for fi60 := 0; fi60 < 2; fi60++ {
                switch int(v14) {
                case -464:
                    v61 := GenS5{F0: 127}
                    _ = v61
                case 42:
                    var v62 uint32 = v6
                    var v63 float64 = float64((1.5 - 1.5))
                    if (439 != 1) {
                        var v64 uint64 = uint64(v6)
                        fmt.Println(v42, v6)
                        var v65 uint32 = 2147483648
                        _ = v64
                        _ = v65
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                    }
                    var v66 uint16 = 42
                    v2 = append(v2, v6)
                    _ = v62
                    _ = v63
                    _ = v66
                case 32767:
                    var v67 int32 = (v41 * (^(v41)))
                    _ = v67
                default:
                    v68 := GenS2{F0: v14, F1: 127}
                    _ = v68
                }
                if len(v2) > 0 {
                    v2[0] = (v6 & v6)
                }
                fmt.Println(v41, v4, v6)
                var v69 int64 = -1
                _ = fi60
                _ = v69
            }
        default:
            var v70 int = int(v6)
            _ = v70
        }
        if false {
            if true {
                v3 -= v3
                fmt.Println(v41)
                fmt.Println(v41, fi44)
            }
            if false {
                v71 := GenS4{F0: 9223372036854775807, F1: 731}
                _ = v71
            } else {
                if (!false) {
                    v6 += (v6 * 2)
                    fmt.Println(fi44, v14)
                    switch fi44 {
                    case 1:
                        v72 := GenS6{F0: 2}
                        _ = v72
                    case -2147483648:
                        v73 := GenS7{F0: "\n\t", F1: 1000}
                        _ = v73
                    case 474:
                        var v74 string = "\n\t"
                        _ = v74
                    }
                }
                fmt.Println(v3, v4)
            }
        } else {
            v14 = 128
        }
        _ = fi44
    }
    switch "hello" {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v3 = (v3 % ((2147483647) | 1))
    }
    v42 = (v42 ^ (^(v42)))
    var v75 int64 = 813
    v75 -= -9223372036854775808
    if true {
        var v76 int32 = v41
        _ = v76
    } else {
        if false {
            v75 = (v75 >> ((1) & 63))
            if false {
                var v77 uint64 = 965
                var v78 float32 = ((3.14159 * 0.0) - (((v4 + (((float32(0.1) / float32(-2.25)) - float32((0.0 / 1.0))) - v4)) - ((0.1 + float32(float64(v4))) * v4)) - v4))
                _ = v77
                _ = v78
            } else {
                var v79 int8 = genFn0(32767, v3)
                _ = v79
            }
            var v80 int64 = (^(v75))
            _ = v80
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v75
}
