package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint64; F2 int }
func genFn0(p0 string, p1 bool, p2 int32) int16 {
    fmt.Println(p2)
    v5 := []int{}
    var v6 uint = 0
    switch p0 {
    case "hello":
        var v7 uint8 = 88
        _ = v7
    case "foo bar":
        if ((-2.25 - 0.0) <= 0.1) {
            p1 = (p1 || p1)
            for fi8 := 0; fi8 < 8; fi8++ {
                if (!p1) {
                    fmt.Println(fi8)
                    v6 += 255
                }
                _ = fi8
            }
        } else {
            v9 := GenS1{F0: 443, F1: 2147483647}
            for fi10 := 0; fi10 < 2; fi10++ {
                switch 1 {
                case 32767:
                    var v11 int64 = 42
                    _ = v11
                case -994:
                    v12 := GenS2{F0: 1, F1: 173}
                    _ = v12
                case -9223372036854775808:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    switch p0 {
                    case "\n\t":
                        p1 = (255 != 940)
                    case "foo bar":
                        var v13 int8 = 42
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 13; fi14++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if p1 {
                                fmt.Println(p2, p1, fi14)
                            }
                            switch fi14 {
                            case -473:
                                p0 = p0
                            case 42:
                                if (p1 && p1) {
                                    _ = len(v5)
                                } else {
                                    var v15 uint16 = 65535
                                    var v16 uint = (v6 & (v6 - v6))
                                    if (!false) {
                                        var v17 int = fi10
                                        var v18 uint32 = 738
                                        var v19 uint = 2
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if ((p2 << ((42) & 63)) >= p2) {
                                            if false {
                                                var v20 uint32 = 393
                                                _ = v20
                                            }
                                        } else {
                                            p1 = ((2 != 4294967295) && p1)
                                            switch "\n\t" {
                                            case "hello":
                                                var v21 float64 = 1.5
                                                _ = v21
                                            case "\n\t":
                                                p1 = (291 < 4294967295)
                                            default:
                                                var v22 string = (p0 + ("a" + ((("" + p0) + p0) + "foo bar")))
                                                _ = v22
                                            }
                                            var v23 float64 = 1e10
                                            _ = v23
                                        }
                                        var v24 uint64 = 4294967295
                                        v25 := GenS2{F0: v6, F1: 1000}
                                        var v26 uint32 = 4294967295
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v15
                                    _ = v16
                                }
                            }
                            fmt.Println(p2, fi10, fi14)
                            _ = fi14
                        }
                    }
                }
                _ = fi10
            }
            p0 = "a"
            _ = v9
        }
    case "a":
        for fi27 := 0; fi27 < 2; fi27++ {
            p1 = (p0 < (((p0 + p0) + "") + (p0 + p0)))
            v28 := GenS2{F0: (v6 | v6), F1: -1}
            _ = fi27
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 6; fi29++ {
            var v30 uint64 = 2
            for fi31 := 0; fi31 < 9; fi31++ {
                if (67 != 255) {
                    if (float64(1.5) >= 1.5) {
                        v32 := GenS3{F0: p1, F1: -976}
                        var v33 float64 = (float64(((float32(0.1) + 1e10) / float32(((0.1 * 1.5) - (((3.14159 + float64(0.0)) + -2.25) + 3.14159))))) * float64((((1.5 / 1.0) / ((1.5 + 0.0) * 1e10)) * 0.0)))
                        switch 0 {
                        case 2147483647:
                            fmt.Println(fi29, p1)
                        case 127:
                            if len(v5) > 0 {
                                v5[0] = fi31
                            }
                        }
                        _ = v32
                        _ = v33
                    } else {
                        switch -240 {
                        case 42:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 2147483647:
                            v5 = append(v5, -359)
                        default:
                            var v34 float32 = 0.0
                            v35 := GenS4{F0: 0}
                            v30 += (v30 ^ (v30 - v30))
                            _ = v34
                            _ = v35
                        }
                    }
                    v36 := GenS3{F0: (!(32767 != 32767)), F1: -1}
                    _ = v36
                } else {
                    v6 = 42
                }
                v37 := GenS3{F0: (!(p1 || (v6 >= 899))), F1: -9223372036854775808}
                _ = fi31
                _ = v37
            }
            var v38 int16 = 0
            _ = fi29
            _ = v30
            _ = v38
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 string) uint16 {
    for fi39 := 0; fi39 < 20; fi39++ {
        v40 := GenS1{F0: 1, F1: -807}
        _ = fi39
        _ = v40
    }
    switch "a" {
    case "foo bar":
        var v41 uint8 = 1
        _ = v41
    }
    var v42 float32 = 1e10
    for fi43 := 0; fi43 < 17; fi43++ {
        if (false && true) {
            v42 += (v42 + (((v42 * (v42 - v42)) + (float32(float64((v42 / v42))) / v42)) / 1e10))
            v44 := GenS5{F0: -9223372036854775808, F1: 979, F2: fi43}
            for fi45 := 0; fi45 < 7; fi45++ {
                if ((((0.0 * float32(-2.25)) - (v42 / (v42 + ((((v42 * float32((1e10 * 1.5))) - 0.1) * 1e10) * 3.14159)))) / (0.0 - 1e10)) >= v42) {
                    v46 := []int{}
                    v46 = append(v46, fi45)
                    switch p0 {
                    case "hello":
                        var v47 int64 = -1
                        _ = v47
                    case "foo bar":
                        if true {
                            v42 += v42
                            p0 = ("hello" + ((p0 + p0) + ((p0 + p0) + p0)))
                        } else {
                            var v48 float64 = 1.5
                            if len(v46) > 0 {
                                v46[0] = fi45
                            }
                            _ = v48
                        }
                    default:
                        _ = len(v46)
                        var v49 int16 = 32767
                        _ = v49
                    }
                    fmt.Println(p0, fi43, p0)
                    _ = v46
                } else {
                    var v50 uint64 = 2
                    _ = v50
                }
                _ = fi45
            }
            _ = v44
        }
        if false {
            switch 671 {
            case 0:
                fmt.Println(p0)
            default:
                v51 := []int{(fi43 ^ -2147483648), 127, (^(fi43)), fi43}
                _ = v51
            }
        } else {
            v52 := []int{fi43, (fi43 % (((fi43 | fi43)) | 1))}
            var v53 bool = (6 == 114)
            _ = v52
            _ = v53
        }
        var v54 string = p0
        _ = fi43
        _ = v54
    }
    var v55 int16 = 1000
    _ = p0
    _ = v42
    _ = v55
    return 865
}

func main() {
    var v1 float32 = 3.14159
    var v2 int32 = -2147483648
    v3 := []string{"a"}
    var v4 int64 = 9223372036854775807
    var v56 int16 = -1
    if (false && (!false)) {
        v4 += 0
        v56 += v56
        var v57 uint64 = 2
        _ = v57
    } else {
        var v58 int = 32767
        for fi59 := 0; fi59 < 7; fi59++ {
            v1 += float32(0.0)
            for fi60 := 0; fi60 < 4; fi60++ {
                switch "" {
                case "foo bar":
                    fmt.Println(v2, v58, fi59)
                }
                _ = fi60
            }
            _ = fi59
        }
        v2 = v2
        _ = v58
    }
    v61 := GenS1{F0: 894, F1: 127}
    for fi62 := 0; fi62 < 16; fi62++ {
        v63 := GenS4{F0: 880}
        _ = fi62
        _ = v63
    }
    v4 -= v4
    v2 = (^(v2))
    for fi64 := 0; fi64 < 17; fi64++ {
        var v65 uint64 = 279
        if false {
            if true {
                v65 = v65
                switch "" {
                case "a":
                    for fi66 := 0; fi66 < 14; fi66++ {
                        v1 -= (0.1 / ((float32((-2.25 + 0.1)) / ((v1 * v1) + v1)) * ((float32(((1e10 - (0.1 + (1.5 / ((1e10 - 0.0) * (3.14159 + 3.14159))))) + 3.14159)) + 1e10) / v1)))
                        v56 += (v56 * 1000)
                        _ = fi66
                    }
                }
            }
        } else {
            if (32767 > 32767) {
                switch 32768 {
                case 1:
                    if (0.1 <= (1.5 + 1e10)) {
                        v4 = (v4 | (v4 + (v4 / (((-(v4))) | 1))))
                        v2 = v2
                    }
                }
            }
        }
        _ = fi64
        _ = v65
    }
    var v67 float32 = 1e10
    if len(v3) > 0 {
        v3[0] = ""
    }
    for fi68 := 0; fi68 < 16; fi68++ {
        var v69 uint16 = 228
        _ = fi68
        _ = v69
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi70 := 0; fi70 < 9; fi70++ {
        if true {
            var v71 int32 = v2
            fmt.Println(v4)
            v71 += (v71 % (((-(v2))) | 1))
            _ = v71
        } else {
            v3 = append(v3, "")
            if true {
                v2 = v2
            } else {
                v2 = 42
            }
        }
        var v72 int16 = (^(v56))
        _ = fi70
        _ = v72
    }
    var v73 uint = 255
    var v74 uint = v73
    for fi75 := 0; fi75 < 20; fi75++ {
        v56 = genFn0("hello", (!(802 != (fi75 ^ (fi75 << ((v74) & 63))))), -800)
        _ = fi75
    }
    var v76 uint32 = 626
    var v77 float32 = v67
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v67:", v67)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v61
    _ = v67
    _ = v73
    _ = v74
    _ = v76
    _ = v77
}
