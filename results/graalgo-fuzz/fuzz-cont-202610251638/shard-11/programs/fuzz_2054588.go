package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int8 }
func genFn0(p0 int32, p1 uint32, p2 uint16) uint32 {
    var v6 int16 = 735
    fmt.Println(v6)
    v7 := GenS1{F0: 2147483647, F1: p0}
    switch "hello" {
    case "a":
        if true {
            for fi8 := 0; fi8 < 6; fi8++ {
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 uint = 18446744073709551615
                    fmt.Println(v10, p1, fi9)
                    _ = fi9
                    _ = v10
                }
                switch "hello" {
                case "hello":
                    p0 = p0
                case "foo bar":
                    if ((("hello" + "foo bar") > "hello") || ((p2 | (p2 >> ((221) & 63))) > p2)) {
                        switch "foo bar" {
                        case "":
                            var v11 int = fi8
                            _ = v11
                        case "\n\t":
                            v12 := GenS2{F0: p2}
                            _ = v12
                        }
                        fmt.Println(p2, p1)
                        v13 := GenS2{F0: 65535}
                        _ = v13
                    }
                case "":
                    switch ("hello" + ("hello" + "a")) {
                    case "":
                        var v14 int16 = -32768
                        _ = v14
                    default:
                        for fi15 := 0; fi15 < 3; fi15++ {
                            var v16 int = fi15
                            v17 := []int{}
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                default:
                    switch 42 {
                    case -2147483648:
                        p1 += p1
                    }
                }
                switch ("a" + ("hello" + "hello")) {
                case "a":
                    fmt.Println(p1, v6, p2)
                }
                _ = fi8
            }
            switch -200 {
            case 32767:
                fmt.Println(p2, p1, p2)
            }
        }
    default:
        var v18 int16 = -1
        _ = v18
    }
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}
func genFn1() uint64 {
    var v19 int32 = 91
    var v20 uint = 65535
    switch 310 {
    case 328:
        if (("foo bar" > "a") && true) {
            v19 = v19
            v21 := GenS1{F0: -1, F1: 323}
            _ = v21
        }
    case -2147483648:
        v22 := GenS3{F0: "a"}
        _ = v22
    default:
        switch "\n\t" {
        case "foo bar":
            v23 := GenS3{F0: "a"}
            _ = v23
        case "a":
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS3{F0: "foo bar"}
                switch -688 {
                case 127:
                    fmt.Println(fi24)
                case -9223372036854775808:
                    var v26 uint32 = 65535
                    _ = v26
                case -38:
                    switch fi24 {
                    case -600:
                        fmt.Println(fi24, v20)
                    }
                }
                var v27 float32 = -2.25
                _ = fi24
                _ = v25
                _ = v27
            }
        case "\n\t":
            fmt.Println(v19)
        default:
            var v28 uint64 = 128
            _ = v28
        }
    }
    var v29 int32 = v19
    _ = v19
    _ = v20
    _ = v29
    return 222
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 654
    v3 := []string{""}
    var v4 uint = 453
    var v5 uint64 = 0
    var v30 uint = uint(v5)
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 float64 = (3.14159 / -2.25)
        var v33 float32 = (((-2.25 / float32(v32)) * (float32((((float64(((1.5 / (((((3.14159 + 3.14159) - 1.5) + 3.14159) * float32(((v32 * (v32 / v32)) / float64(3.14159)))) - 0.0)) + 0.1)) / v32) / v32) / (1e10 / v32))) / (float32(1.5) / 0.1))) * -2.25)
        _ = fi31
        _ = v32
        _ = v33
    }
    switch (v2 + 9223372036854775807) {
    case -2147483648:
        if (!true) {
            var v34 float64 = (0.0 + -2.25)
            v35 := GenS4{F0: 100, F1: "hello"}
            _ = v34
            _ = v35
        } else {
            v5 += v5
            v36 := GenS3{F0: ("" + "\n\t")}
            _ = v36
        }
    case 42:
        v37 := GenS5{F0: (v4 ^ 255), F1: 1, F2: 4294967295}
        _ = v37
    }
    var v38 uint = 511
    v38 = v30
    var v39 uint16 = 1
    fmt.Println(v2, v39, v4)
    switch "\n\t" {
    case "":
        switch v2 {
        case 128:
            var v40 uint32 = 965
            _ = v40
        case 32767:
            v41 := GenS6{F0: int8(v5), F1: v38}
            _ = v41
        case -9223372036854775808:
            var v42 uint = 42
            _ = v42
        }
    case "hello":
        var v43 uint64 = (v5 | (v5 - genFn1()))
        _ = v43
    case "\n\t":
        fmt.Println(v4)
    default:
        for fi44 := 0; fi44 < 13; fi44++ {
            if (false || true) {
                if (-56 < 42) {
                    var v45 uint = 466
                    _ = v45
                }
            } else {
                var v46 string = "foo bar"
                if (-2.25 > 0.0) {
                    v3 = append(v3, v46)
                    v47 := GenS1{F0: 449, F1: 0}
                    for fi48 := 0; fi48 < 7; fi48++ {
                        fmt.Println(v2)
                        v39 -= 82
                        v39 += 894
                        _ = fi48
                    }
                    _ = v47
                }
                if false {
                    v2 -= v2
                    if (v39 <= v39) {
                        if (65535 >= 2) {
                            v30 = v38
                            for fi49 := 0; fi49 < 4; fi49++ {
                                if (((!((-1 > -195) && ((fi44 + v2) > (fi44 | (v2 >> (((v4 >> ((v38) & 63))) & 63)))))) || true) && true) {
                                    v50 := GenS7{F0: "a", F1: fi44}
                                    if false {
                                        switch v46 {
                                        case "":
                                            var v51 uint16 = (v39 >> ((v30) & 63))
                                            _ = v51
                                        default:
                                            v2 = fi49
                                        }
                                    }
                                    _ = v50
                                }
                                _ = fi49
                            }
                        } else {
                            v5 = 4294967295
                            var v52 uint = (v4 & (v4 ^ (v4 - v38)))
                            _ = v52
                        }
                    }
                }
                _ = v46
            }
            if true {
                if (true && true) {
                    var v53 uint = 2
                    _ = v53
                } else {
                    var v54 float32 = ((3.14159 - -2.25) / 1.0)
                    v4 = v30
                    v3 = append(v3, "a")
                    _ = v54
                }
            } else {
                for fi55 := 0; fi55 < 17; fi55++ {
                    var v56 float64 = (float64(1.5) / 1.5)
                    var v57 uint32 = 2147483648
                    _ = fi55
                    _ = v56
                    _ = v57
                }
            }
            _ = fi44
        }
    }
    fmt.Println(v39)
    v58 := GenS4{F0: -1, F1: "hello"}
    fmt.Println(v5, v38)
    v5 += (v5 << ((v38) & 63))
    v4 = (v38 * (v38 % (((v38 + (v4 * v4))) | 1)))
    switch "\n\t" {
    case "\n\t":
        var v59 uint8 = 2
        _ = v59
    case "":
        for fi60 := 0; fi60 < 3; fi60++ {
            for fi61 := 0; fi61 < 4; fi61++ {
                if (((0.0 == (0.0 + 3.14159)) && (genFn1() >= (v5 | 18446744073709551615))) && false) {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    if true {
                        var v62 uint64 = (v5 % (((v5 | (v5 / ((v5) | 1)))) | 1))
                        _ = v62
                    }
                } else {
                    var v63 uint8 = 40
                    v39 = (v39 << (((^(v30))) & 63))
                    v4 = (v4 >> ((255) & 63))
                    v63 += 128
                    _ = v63
                }
                var v64 int = 128
                _ = fi61
                _ = v64
            }
            v1[v2] = v2
            v39 -= (v39 >> ((v30) & 63))
            _ = fi60
        }
    case "hello":
        var v65 int = (v2 % (((^(v2))) | 1))
        _ = v65
    }
    switch "foo bar" {
    case "hello":
        v5 = (v5 * v5)
    case "":
        var v66 uint16 = uint16(v38)
        _ = v66
    case "a":
        fmt.Println(v4, v2, v4)
    }
    v67 := GenS8{F0: 95}
    var v68 int8 = int8(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v58
    _ = v67
    _ = v68
}
