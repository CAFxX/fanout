package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 string }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int64; F1 uint; F2 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 uint8; F1 uint32 }
func genFn0(p0 string, p1 int64) uint64 {
    fmt.Println(p1, p0)
    var v5 uint64 = 42
    switch p0 {
    case "hello":
        switch int(p1) {
        case 786:
            for fi6 := 0; fi6 < 15; fi6++ {
                for fi7 := 0; fi7 < 14; fi7++ {
                    v8 := GenS1{F0: true, F1: -1, F2: 1}
                    var v9 uint = 435
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                _ = fi6
            }
        case 0:
            var v10 uint8 = 128
            _ = v10
        }
    case "\n\t":
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint32 = 4294967295
            v13 := GenS2{F0: (-(fi11)), F1: (1 < -365)}
            if ((!true) && (!true)) {
                var v14 bool = (147 > 65535)
                _ = v14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    case "a":
        if true {
            for fi15 := 0; fi15 < 12; fi15++ {
                fmt.Println(p1)
                _ = fi15
            }
            v16 := GenS3{F0: 522, F1: 2, F2: p0}
            _ = v16
        } else {
            if (147 > 42) {
                switch p0 {
                case "foo bar":
                    v17 := GenS2{F0: 2147483647, F1: (false && true)}
                    _ = v17
                case "\n\t":
                    v18 := GenS3{F0: 32767, F1: 4294967295, F2: "a"}
                    _ = v18
                }
                if true {
                    var v19 string = "\n\t"
                    if false {
                        v19 = ""
                        v5 = v5
                    } else {
                        p0 = (p0 + p0)
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v21 := GenS3{F0: 0, F1: 42, F2: p0}
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = v19
                }
                switch ("" + p0) {
                case "a":
                    v22 := []int{}
                    _ = v22
                default:
                    var v23 int64 = p1
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 14; fi24++ {
                p0 = p0
                fmt.Println(fi24, v5)
                _ = fi24
            }
        }
    }
    v5 = 481
    p0 = (((p0 + p0) + p0) + ("" + p0))
    _ = p0
    _ = p1
    _ = v5
    return (v5 + v5)
}
func genFn1(p0 uint16, p1 string) int8 {
    switch -460 {
    case -924:
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 int = fi25
            _ = fi25
            _ = v26
        }
    case 1:
        var v27 bool = (1.5 >= 1e10)
        _ = v27
    case -286:
        fmt.Println(p1, p1)
    default:
        p0 -= p0
    }
    v28 := GenS4{F0: 65535, F1: false}
    p1 = p1
    _ = p0
    _ = p1
    _ = v28
    return 1
}

func main() {
    var v1 uint8 = 182
    var v2 uint16 = 32767
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    if false {
        if true {
            var v29 int16 = 42
            v30 := GenS5{F0: -709, F1: 0, F2: -9223372036854775808}
            _ = v29
            _ = v30
        } else {
            switch "foo bar" {
            case "\n\t":
                if (244 > 1) {
                    v2 -= (^(v2))
                } else {
                    v31 := v4[2147483647]
                    if (!(0 > v1)) {
                        v4[728] = v31
                        for fi32 := 0; fi32 < 13; fi32++ {
                            v2 = 306
                            _ = fi32
                        }
                        v33 := GenS6{F0: 0}
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 20; fi34++ {
                            v31 = v31
                            var v35 float32 = 1e10
                            var v36 uint16 = v2
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 17; fi37++ {
                            var v38 int16 = 1000
                            var v39 float32 = -2.25
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                    _ = v31
                }
            case "":
                v40 := GenS6{F0: 0}
                _ = v40
            default:
                v3[2] = false
            }
            if ((-2.25 + ((0.0 / (0.1 + 0.0)) / (-2.25 - 1.5))) >= 3.14159) {
                v41 := []int{4294967295, 32768, 2147483647, -2147483648}
                var v42 uint16 = (v2 % (((v2 & v2)) | 1))
                _ = v41
                _ = v42
            }
            v2 = (v2 / ((708) | 1))
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            if ((fi43 / (((^(fi43))) | 1)) <= -2147483648) {
                delete(v3, 26)
                _ = len(v3)
                if (!false) {
                    for fi44 := 0; fi44 < 17; fi44++ {
                        fmt.Println(fi43)
                        switch fi44 {
                        case -9223372036854775808:
                            if true {
                                var v45 bool = false
                                _ = v45
                            } else {
                                v2 = (v2 % (((v2 | v2)) | 1))
                                var v46 int = fi43
                                v1 = v1
                                _ = v46
                            }
                        case -1:
                            v1 -= (v1 / ((0) | 1))
                        case 9223372036854775807:
                            v1 = (v1 << ((2) & 63))
                        default:
                            fmt.Println(fi43)
                        }
                        v2 += v2
                        _ = fi44
                    }
                }
            }
            var v47 string = "hello"
            v48 := GenS3{F0: v2, F1: 234, F2: (v47 + "a")}
            _ = fi43
            _ = v47
            _ = v48
        }
    } else {
        v1 = (v1 % ((v1) | 1))
        v49 := []int{}
        if false {
            for fi50 := 0; fi50 < 10; fi50++ {
                for fi51 := 0; fi51 < 1; fi51++ {
                    v52 := GenS7{F0: -128, F1: genFn0((((((("hello" + ("" + "a")) + "hello") + (("a" + "") + "foo bar")) + "\n\t") + (("foo bar" + "") + ("" + ("foo bar" + "foo bar")))) + ""), 9223372036854775807)}
                    var v53 uint8 = v1
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                switch "hello" {
                case "hello":
                    for fi54 := 0; fi54 < 2; fi54++ {
                        switch (fi54 % ((fi50) | 1)) {
                        case 0:
                            v55 := GenS2{F0: 42, F1: (-31 > 42)}
                            var v56 float32 = 0.1
                            v57 := GenS7{F0: 0, F1: 849}
                            _ = v55
                            _ = v56
                            _ = v57
                        case 127:
                            if true {
                                var v58 int64 = 1
                                _ = v58
                            }
                        case 948:
                            v1 += v1
                        }
                        _ = fi54
                    }
                case "foo bar":
                    var v59 uint16 = 1
                    _ = v59
                case "\n\t":
                    switch "a" {
                    case "foo bar":
                        for fi60 := 0; fi60 < 20; fi60++ {
                            v61 := GenS7{F0: -87, F1: 109}
                            _ = fi60
                            _ = v61
                        }
                    case "hello":
                        var v62 uint = 0
                        _ = v62
                    default:
                        switch -9223372036854775808 {
                        case 573:
                            fmt.Println(v1)
                        case -651:
                            if len(v49) > 0 {
                                v49[0] = (^(fi50))
                            }
                        case 32767:
                            v63 := GenS1{F0: false, F1: -1, F2: 27}
                            _ = v63
                        }
                    }
                default:
                    if true {
                        for fi64 := 0; fi64 < 17; fi64++ {
                            switch "\n\t" {
                            case "foo bar":
                                v49 = append(v49, 45)
                            case "":
                                v65 := v4[2147483648]
                                _ = v65
                            default:
                                v1 -= v1
                            }
                            _ = fi64
                        }
                        switch (fi50 + (fi50 >> ((18446744073709551615) & 63))) {
                        case -91:
                            if (!true) {
                                for fi66 := 0; fi66 < 20; fi66++ {
                                    var v67 bool = true
                                    v1 = 42
                                    if false {
                                        switch "" {
                                        case "a":
                                            var v68 uint8 = 127
                                            _ = v68
                                        default:
                                            var v69 uint = 18446744073709551615
                                            _ = v69
                                        }
                                        fmt.Println(v67, fi66)
                                    } else {
                                        var v70 string = "foo bar"
                                        fmt.Println(fi50)
                                        if v67 {
                                            v1 = (v1 * (v1 - uint8(v2)))
                                        }
                                        _ = v70
                                    }
                                    _ = fi66
                                    _ = v67
                                }
                                var v71 int32 = -2147483648
                                _ = v71
                            }
                        case 9223372036854775807:
                            v2 -= (v2 | v2)
                        }
                        v72 := v4[0]
                        _ = v72
                    } else {
                        v2 = 1
                    }
                }
                if len(v49) >= 2 {
                    v49 = v49[1:len(v49)]
                }
                _ = fi50
            }
            for fi73 := 0; fi73 < 8; fi73++ {
                switch "" {
                case "foo bar":
                    v74 := GenS7{F0: 100, F1: 1}
                    _ = v74
                case "hello":
                    var v75 float64 = float64(1.5)
                    _ = v75
                case "\n\t":
                    var v76 uint32 = uint32(v1)
                    _ = v76
                }
                _ = fi73
            }
            v77 := GenS4{F0: 531, F1: (!true)}
            _ = v77
        }
        _ = v49
    }
    switch -1 {
    case -213:
        v1 += (v1 >> ((537) & 63))
    case -1:
        var v78 bool = ((false && (4294967295 <= 2)) && true)
        _ = v78
    case 2147483647:
        for fi79 := 0; fi79 < 20; fi79++ {
            switch "\n\t" {
            case "":
                v80 := GenS5{F0: 9223372036854775807, F1: 4294967295, F2: (fi79 + (fi79 % ((2147483647) | 1)))}
                _ = v80
            default:
                if (true && false) {
                    v2 = (^(v2))
                    v1 -= v1
                    v1 += (v1 ^ (v1 | v1))
                }
            }
            fmt.Println(fi79, v1, v2)
            var v81 int16 = 32767
            _ = fi79
            _ = v81
        }
    default:
        var v82 float32 = 1e10
        _ = v82
    }
    for fi83 := 0; fi83 < 8; fi83++ {
        for fi84 := 0; fi84 < 11; fi84++ {
            _ = len(v4)
            fmt.Println(fi83, fi83, v2)
            v2 += v2
            v85 := GenS8{F0: 1, F1: 2}
            _ = fi84
            _ = v85
        }
        _ = fi83
    }
    fmt.Println(v2, v1)
    v1 = (v1 * v1)
    if false {
        v1 = 249
    }
    _ = len(v4)
    v86 := []int{127}
    if (-2147483648 <= 42) {
        if len(v86) >= 2 {
            v86 = v86[1:len(v86)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v86:", len(v86), v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v86
}
