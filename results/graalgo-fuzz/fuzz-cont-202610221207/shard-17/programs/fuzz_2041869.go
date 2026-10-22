package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 uint }
type GenS4 struct { F0 int32; F1 int32; F2 uint16 }
type GenS5 struct { F0 bool; F1 int; F2 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0(p0 int) int64 {
    switch "foo bar" {
    case "":
        switch "foo bar" {
        case "\n\t":
            if (44 > 0) {
                p0 += p0
            } else {
                var v4 int8 = 0
                var v5 int8 = v4
                if false {
                    var v6 int = (-(p0))
                    _ = v6
                }
                _ = v4
                _ = v5
            }
        case "foo bar":
            v7 := GenS1{F0: -112}
            _ = v7
        default:
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := []int{(fi8 >> ((2) & 63)), p0, (p0 / ((fi8) | 1))}
                var v10 float32 = (3.14159 * ((float32(0.0) / 0.1) / 0.1))
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    case "a":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 bool = true
            _ = fi11
            _ = v12
        }
    case "hello":
        v13 := []int{p0, 32768, 4294967295}
        _ = v13
    }
    if false {
        p0 = -372
        for fi14 := 0; fi14 < 18; fi14++ {
            p0 = p0
            p0 = p0
            switch 4294967295 {
            case 42:
                v15 := GenS1{F0: 0}
                _ = v15
            }
            _ = fi14
        }
        fmt.Println(p0)
    }
    _ = p0
    return -1
}
func genFn1(p0 int64) uint {
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int8 = -1
        v18 := GenS1{F0: (v17 | v17)}
        _ = fi16
        _ = v17
        _ = v18
    }
    if (!true) {
        v19 := GenS1{F0: 0}
        v20 := []int{}
        _ = v19
        _ = v20
    } else {
        p0 -= 9223372036854775807
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(p0, p0)
            p0 += (p0 & (p0 % ((158) | 1)))
            _ = fi21
        }
        v22 := []int{1, -844, 9223372036854775807, -9223372036854775808}
        _ = v22
    }
    p0 -= -742
    var v23 uint32 = 2147483648
    p0 = genFn0(189)
    _ = p0
    _ = v23
    return 105
}

func main() {
    var v1 uint32 = 488
    v2 := make(map[int]uint32)
    var v3 uint16 = 32767
    switch "\n\t" {
    case "\n\t":
        if (42 <= 1) {
            v24 := []int{}
            for fi25 := 0; fi25 < 13; fi25++ {
                v3 = (v3 % (((v3 ^ (v3 + 684))) | 1))
                _ = fi25
            }
            var v26 uint32 = v1
            _ = v24
            _ = v26
        }
    default:
        delete(v2, 4294967295)
    }
    var v27 uint8 = 2
    fmt.Println(v27)
    var v28 int64 = -913
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS1{F0: 100}
        var v31 uint64 = 4294967295
        if false {
            v32 := GenS1{F0: 8}
            for fi33 := 0; fi33 < 11; fi33++ {
                switch -1 {
                case -910:
                    v34 := GenS1{F0: -128}
                    _ = v34
                case 32767:
                    v27 -= (v27 + v27)
                case 731:
                    var v35 int64 = v28
                    _ = v35
                default:
                    switch fi29 {
                    case 1:
                        v36 := GenS1{F0: 61}
                        _ = v36
                    case 0:
                        v28 -= int64(fi33)
                    }
                }
                v37 := []int{fi33}
                _ = fi33
                _ = v37
            }
            switch (fi29 / (((fi29 - fi29)) | 1)) {
            case -9223372036854775808:
                var v38 uint8 = v27
                _ = v38
            case 0:
                var v39 string = "a"
                _ = v39
            case 128:
                v40 := GenS2{F0: ((fi29 >> ((genFn1(v28)) & 63)) != fi29), F1: v28}
                _ = v40
            default:
                var v41 bool = true
                _ = v41
            }
            _ = v32
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi42 := 0; fi42 < 11; fi42++ {
        var v43 float32 = (1e10 * 0.0)
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        if true {
            fmt.Println(fi44, fi44)
            if false {
                for fi45 := 0; fi45 < 17; fi45++ {
                    var v46 bool = ((v27 > v27) || (!(!(42 != 613))))
                    _ = fi45
                    _ = v46
                }
            }
        } else {
            var v47 float64 = (1e10 * 0.1)
            _ = v47
        }
        fmt.Println(v3, v28, fi44)
        _ = fi44
    }
    v48 := GenS3{F0: 0, F1: 32768, F2: 255}
    v49 := v2[-2147483648]
    v50 := GenS4{F0: 0, F1: 42, F2: (v3 << ((4294967295) & 63))}
    v51 := v2[2147483647]
    v52 := GenS5{F0: true, F1: -9223372036854775808, F2: 470}
    v53 := GenS6{F0: ("hello" + "\n\t")}
    switch -2147483648 {
    case 42:
        if true {
            v54 := v2[9223372036854775807]
            if (0 < -358) {
                switch ("foo bar" + "hello") {
                case "foo bar":
                    for fi55 := 0; fi55 < 19; fi55++ {
                        v56 := GenS7{F0: (!true), F1: 65535}
                        v57 := GenS6{F0: ("a" + ("" + "hello"))}
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                }
            }
            _ = v54
        } else {
            var v58 float64 = ((3.14159 * (0.1 + 0.1)) * 3.14159)
            fmt.Println(v51)
            if true {
                for fi59 := 0; fi59 < 3; fi59++ {
                    for fi60 := 0; fi60 < 9; fi60++ {
                        if false {
                            v61 := GenS6{F0: "foo bar"}
                            fmt.Println(v58, fi59)
                            v49 -= 2147483647
                            v62 := GenS8{F0: (v49 ^ uint32(v27)), F1: (v51 % ((v51) | 1)), F2: 65535}
                            if false {
                                v63 := []int{205, fi60, int(v1)}
                                _ = v63
                            }
                            var v64 string = "\n\t"
                            _ = v61
                            _ = v62
                            _ = v64
                        }
                        var v65 int64 = v28
                        fmt.Println(v28, v58)
                        var v66 float32 = 3.14159
                        _ = fi60
                        _ = v65
                        _ = v66
                    }
                    _ = len(v2)
                    for fi67 := 0; fi67 < 20; fi67++ {
                        v68 := GenS3{F0: uint64(v27), F1: fi67, F2: 18446744073709551615}
                        if (65535 != v3) {
                            v69 := []int{(fi67 | (fi67 % (((fi59 / (((fi59 / (((fi67 + -803)) | 1))) | 1))) | 1))), fi59, int(v51)}
                            _ = v69
                        }
                        fmt.Println(fi67, v27)
                        var v70 int32 = 0
                        v51 += (^(v1))
                        v27 = v27
                        v71 := []int{127, (fi59 % ((int(v27)) | 1)), fi59, fi59}
                        if true {
                            switch fi67 {
                            case 9223372036854775807:
                                if true {
                                    var v72 int16 = 926
                                    _ = v72
                                } else {
                                    if len(v71) >= 2 {
                                        v71 = v71[1:len(v71)]
                                    }
                                    var v73 uint = 18446744073709551615
                                    _ = v73
                                }
                            }
                            v1 = v49
                            var v74 uint16 = v3
                            v58 = v58
                            _ = v74
                        }
                        _ = fi67
                        _ = v68
                        _ = v70
                        _ = v71
                    }
                    _ = fi59
                }
            }
            _ = v58
        }
    case 127:
        v75 := []int{4294967295, 4294967295}
        _ = v75
    }
    v2[2147483647] = (v49 << ((99) & 63))
    var v76 float64 = 1.5
    var v77 uint64 = uint64(v27)
    var v78 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v76
    _ = v77
    _ = v78
}
