package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint8; F1 uint; F2 int64 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS10 struct { F0 int; F1 uint16; F2 int }
func genFn0(p0 int32, p1 bool, p2 uint8) int16 {
    var v6 string = "hello"
    v7 := []int{4294967295}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 string, p1 string, p2 uint16) int16 {
    if true {
        v9 := GenS1{F0: ("" + p1)}
        for fi10 := 0; fi10 < 2; fi10++ {
            p0 = p1
            _ = fi10
        }
        _ = v9
    }
    var v11 string = (((((((p1 + (p0 + p0)) + p1) + (p1 + p0)) + "\n\t") + ((p1 + (((p0 + p0) + p1) + p1)) + "a")) + "") + (p1 + ("hello" + (("" + ((("\n\t" + p0) + p0) + p0)) + ((p0 + ("hello" + p1)) + p0)))))
    var v12 int32 = -880
    for fi13 := 0; fi13 < 3; fi13++ {
        v12 = v12
        v14 := GenS1{F0: p0}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            switch v11 {
            case "\n\t":
                v11 = ((("a" + "hello") + v11) + ("a" + p1))
            case "foo bar":
                v11 = v11
            case "":
                v12 = 0
            default:
                v17 := []int{fi16, fi15}
                _ = v17
            }
            switch (("hello" + "foo bar") + p1) {
            case "\n\t":
                if true {
                    var v18 int32 = 1
                    if false {
                        v19 := []int{(fi16 & -1), (fi16 & fi15)}
                        if (!(!false)) {
                            switch "a" {
                            case "foo bar":
                                v12 = (v18 % (((v18 ^ v18)) | 1))
                            default:
                                v20 := GenS1{F0: v11}
                                _ = v20
                            }
                            p1 = p1
                        }
                        var v21 int64 = -9223372036854775808
                        switch p1 {
                        case "hello":
                            var v22 uint8 = 128
                            _ = v22
                        case "a":
                            if false {
                                var v23 uint32 = 2
                                var v24 string = p1
                                _ = v23
                                _ = v24
                            }
                        default:
                            v12 = (-(v12))
                        }
                        fmt.Println(v12)
                        _ = v19
                        _ = v21
                    } else {
                        if false {
                            v11 = ("hello" + "\n\t")
                            var v25 uint32 = 775
                            v12 = -120
                            v18 -= 0
                            _ = v25
                        }
                    }
                    fmt.Println(v12)
                    v26 := []int{-192}
                    _ = v18
                    _ = v26
                } else {
                    if (42 != -654) {
                        if (945 > -574) {
                            var v27 int = 9223372036854775807
                            v12 -= -281
                            _ = v27
                        }
                        if false {
                            p1 = "a"
                            v28 := GenS1{F0: ("a" + "a")}
                            if (1 <= 0) {
                                v29 := GenS1{F0: "\n\t"}
                                v30 := []int{fi16, (fi16 >> ((102) & 63))}
                                _ = v29
                                _ = v30
                            }
                            _ = v28
                        }
                    } else {
                        var v31 int64 = -569
                        var v32 string = p1
                        switch v11 {
                        case "a":
                            var v33 int = 42
                            _ = v33
                        case "\n\t":
                            var v34 int64 = (v31 + v31)
                            _ = v34
                        default:
                            p1 = p0
                        }
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(p2)
                    var v35 float32 = (float32(float64(0.0)) - 3.14159)
                    _ = v35
                }
            case "":
                if (true || (!true)) {
                    fmt.Println(v12, v11, fi16)
                    v36 := []int{fi15, -9223372036854775808}
                    _ = v36
                } else {
                    fmt.Println(fi15, p2, p1)
                    var v37 string = ((("foo bar" + (v11 + "")) + (p0 + (p1 + (("" + ("a" + p0)) + p1)))) + (p0 + ""))
                    switch p1 {
                    case "foo bar":
                        fmt.Println(v12)
                        if false {
                            v38 := []int{}
                            _ = v38
                        } else {
                            var v39 bool = true
                            v40 := GenS1{F0: p1}
                            _ = v39
                            _ = v40
                        }
                        p1 = (v11 + v37)
                    default:
                        v41 := []int{(fi15 & 2147483647), (fi15 >> ((1) & 63)), -974, (^(fi16))}
                        _ = v41
                    }
                    _ = v37
                }
            }
            switch "hello" {
            case "":
                if false {
                    fmt.Println(fi16, p1, p0)
                    switch fi16 {
                    case 139:
                        v42 := []int{(fi15 - fi16)}
                        var v43 uint64 = 18446744073709551615
                        v12 -= v12
                        _ = v42
                        _ = v43
                    case 127:
                        var v44 int64 = 9223372036854775807
                        _ = v44
                    case -959:
                        var v45 uint = 42
                        _ = v45
                    }
                } else {
                    v46 := []int{(fi16 + fi16), fi16, (fi15 / (((-(fi16))) | 1))}
                    _ = v46
                }
            case "foo bar":
                var v47 uint32 = 1
                _ = v47
            case "hello":
                var v48 float64 = 0.0
                if true {
                    if false {
                        v12 = (-(v12))
                        var v49 int64 = 9223372036854775807
                        _ = v49
                    } else {
                        v50 := GenS2{F0: 0}
                        if ((1.5 - -2.25) != 0.1) {
                            switch fi16 {
                            case 697:
                                p2 = p2
                            default:
                                var v51 int64 = 0
                                _ = v51
                            }
                        } else {
                            if true {
                                fmt.Println(p1, fi15)
                                fmt.Println(p2, v12)
                            }
                            if (855 != 4294967295) {
                                var v52 uint8 = 255
                                var v53 float64 = float64(float32(1e10))
                                _ = v52
                                _ = v53
                            }
                        }
                        _ = v50
                    }
                    v48 += (v48 * 1.5)
                }
                var v54 string = ""
                v55 := []int{(-(fi16)), (-(fi16)), fi15, fi15}
                p0 = (v54 + "foo bar")
                _ = v48
                _ = v54
                _ = v55
            }
            v56 := []int{}
            if (v11 > p1) {
                if (v12 == (v12 ^ 42)) {
                    switch 128 {
                    case 32767:
                        var v57 int16 = -1
                        _ = v57
                    default:
                        fmt.Println(fi15, fi16, p1)
                    }
                    var v58 int8 = 42
                    _ = v58
                }
                if false {
                    var v59 int = 4294967295
                    v60 := GenS3{F0: 9223372036854775807, F1: true}
                    _ = v59
                    _ = v60
                } else {
                    v12 += (v12 * (v12 | v12))
                    if ("hello" > p0) {
                        if false {
                            v61 := GenS4{F0: 191, F1: 996}
                            _ = v61
                        } else {
                            v62 := GenS5{F0: 42, F1: 759}
                            _ = v62
                        }
                    } else {
                        if true {
                            v56 = append(v56, 128)
                            fmt.Println(p0, v11)
                        } else {
                            p2 += 65535
                        }
                    }
                }
            }
            _ = fi16
            _ = v56
        }
        for fi63 := 0; fi63 < 16; fi63++ {
            v64 := []int{(^(fi15)), fi63, (-(fi15))}
            _ = fi63
            _ = v64
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -543
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = ((0.1 * 1e10) / ((-2.25 / -2.25) + 0.0))
    var v3 uint = 18446744073709551615
    var v4 uint8 = 128
    var v5 uint32 = 2147483648
    v4 -= v4
    for fi65 := 0; fi65 < 13; fi65++ {
        for fi66 := 0; fi66 < 11; fi66++ {
            v67 := GenS6{F0: 255, F1: v3, F2: 867}
            _ = fi66
            _ = v67
        }
        if false {
            v3 = v3
            switch "" {
            case "":
                for fi68 := 0; fi68 < 9; fi68++ {
                    v69 := GenS3{F0: -679, F1: (v2 <= (v2 - 0.1))}
                    v2 -= ((((float32(1.5) / v2) / (v2 + ((v2 - (((-2.25 - 3.14159) - 3.14159) / v2)) + v2))) / float32(0.1)) * -2.25)
                    _ = fi68
                    _ = v69
                }
            case "a":
                var v70 int = (fi65 >> (((v3 << (((^(v3))) & 63))) & 63))
                _ = v70
            case "foo bar":
                v71 := GenS7{F0: 1, F1: (v3 - v3)}
                _ = v71
            default:
                v5 -= (v5 / (((^(v5))) | 1))
            }
            for fi72 := 0; fi72 < 14; fi72++ {
                var v73 uint = 4294967295
                _ = fi72
                _ = v73
            }
        }
        _ = fi65
    }
    v2 += (1e10 - (v2 / v2))
    if false {
        var v74 uint16 = 32767
        v1[32767] = ("" + "a")
        for fi75 := 0; fi75 < 2; fi75++ {
            var v76 int64 = 1
            v77 := GenS7{F0: 42, F1: v3}
            _ = fi75
            _ = v76
            _ = v77
        }
        _ = v74
    }
    switch -2147483648 {
    case 9223372036854775807:
        switch "foo bar" {
        case "\n\t":
            for fi78 := 0; fi78 < 9; fi78++ {
                v79 := GenS8{F0: v3, F1: 127}
                if (false || false) {
                    fmt.Println(v3, v3, v5)
                    if (1 >= 2) {
                        v4 = uint8(v5)
                        fmt.Println(v4, v2)
                        var v80 int = (fi78 + fi78)
                        _ = v80
                    }
                }
                v81 := GenS5{F0: 2, F1: 65535}
                _ = fi78
                _ = v79
                _ = v81
            }
        }
    case 127:
        fmt.Println(v2)
    case 1:
        var v82 uint = (v3 % ((678) | 1))
        _ = v82
    }
    delete(v1, -1)
    var v83 uint = v3
    v83 -= (v3 + v3)
    var v84 float64 = 0.0
    fmt.Println(v84)
    v2 = (float32((1.5 + v84)) / float32(v84))
    fmt.Println(v83, v84, v84)
    delete(v1, 9223372036854775807)
    if true {
        v85 := GenS2{F0: (v83 / (((^(v83))) | 1))}
        v2 = v2
        _ = v85
    }
    v86 := GenS5{F0: 18, F1: 1}
    if (127 < 1) {
        v84 = v84
        for fi87 := 0; fi87 < 20; fi87++ {
            switch (fi87 / (((fi87 % ((fi87) | 1))) | 1)) {
            case -827:
                switch (fi87 - fi87) {
                case -2147483648:
                    v88 := GenS9{F0: v4, F1: 335, F2: 944}
                    _ = v88
                default:
                    v89 := GenS6{F0: 2, F1: 4294967295, F2: 9223372036854775807}
                    _ = v89
                }
            case -773:
                v5 -= (v5 / ((2147483647) | 1))
            default:
                v83 -= (v83 % ((v3) | 1))
            }
            v90 := []int{-9223372036854775808, -9223372036854775808}
            _ = fi87
            _ = v90
        }
        fmt.Println(v2, v5)
    }
    v91 := GenS10{F0: 298, F1: 184, F2: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v86
    _ = v91
}
