package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 int16; F1 int16; F2 uint16 }
type GenS5 struct { F0 string; F1 string; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 bool }
func genFn0(p0 uint64, p1 uint32) bool {
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(fi6, p1, fi6)
        for fi7 := 0; fi7 < 10; fi7++ {
            p1 -= 2
            _ = fi7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{}
            p0 += p0
            fmt.Println(fi6)
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    var v10 bool = (255 != 1)
    v11 := GenS1{F0: -20, F1: 1}
    p1 = (p1 + 2147483648)
    var v12 uint = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return v10
}

func main() {
    var v1 string = ("a" + "\n\t")
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 float64 = -2.25
    var v5 int8 = 1
    if ((!false) && (0 > 560)) {
        fmt.Println(v1)
        if (!(9223372036854775807 <= 121)) {
            v13 := GenS2{F0: v1, F1: 1}
            var v14 uint64 = 18446744073709551615
            _ = v13
            _ = v14
        }
    }
    if false {
        v4 = 1.5
        if (-128 == (v5 / ((v5) | 1))) {
            for fi15 := 0; fi15 < 14; fi15++ {
                v3[2] = ("\n\t" + v1)
                switch ("\n\t" + ((v1 + "hello") + v1)) {
                case "foo bar":
                    v5 = v5
                default:
                    v4 = v4
                }
                if (!false) {
                    var v16 uint8 = 128
                    fmt.Println(v4, v4, v5)
                    _ = v16
                } else {
                    if (0.0 > 1.5) {
                        v17 := GenS3{F0: 0, F1: ((v1 + "foo bar") + (v1 + (("a" + "\n\t") + v1))), F2: 2147483647}
                        _ = v17
                    } else {
                        var v18 string = v1
                        _ = v18
                    }
                }
                _ = fi15
            }
            v19 := GenS3{F0: -420, F1: "", F2: 1}
            for fi20 := 0; fi20 < 14; fi20++ {
                for fi21 := 0; fi21 < 8; fi21++ {
                    switch (-(fi20)) {
                    case 362:
                        v4 = v4
                    }
                    fmt.Println(fi20, v1, fi21)
                    var v22 float32 = (float32(float64(0.0)) * (((1.5 * (0.0 - (1e10 - 0.1))) + 1.5) - (3.14159 / 0.1)))
                    _ = fi21
                    _ = v22
                }
                v23 := GenS1{F0: (v5 + v5), F1: 140}
                var v24 uint32 = 443
                _ = fi20
                _ = v23
                _ = v24
            }
            _ = v19
        }
    }
    fmt.Println(v1, v4, v5)
    var v25 uint = 42
    switch v1 {
    case "foo bar":
        var v26 int16 = 42
        _ = v26
    case "a":
        v27 := GenS1{F0: 0, F1: 42}
        _ = v27
    default:
        if false {
            v28 := []int{-1, 32768, 4294967295, -1}
            _ = v28
        }
    }
    v29 := GenS1{F0: 1, F1: 128}
    v30 := GenS2{F0: "\n\t", F1: 127}
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := []int{fi31, (fi31 | fi31), (fi31 * fi31), (fi31 + 1)}
        _ = fi31
        _ = v32
    }
    v4 = v4
    v25 += v25
    if true {
        var v33 string = ((v1 + ((v1 + v1) + (v1 + "a"))) + v1)
        _ = v33
    } else {
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 float64 = 1.5
            for fi36 := 0; fi36 < 8; fi36++ {
                switch fi34 {
                case 643:
                    v37 := GenS4{F0: -1, F1: 42, F2: 65535}
                    _ = v37
                case 0:
                    v25 = 742
                }
                _ = fi36
            }
            _ = fi34
            _ = v35
        }
        if true {
            v4 = (((((v4 / ((((1.5 - (((0.0 + v4) + ((-2.25 / ((v4 + v4) - v4)) + (v4 / float64(0.1)))) - v4)) / (float64(3.14159) * ((0.0 + v4) / v4))) - v4) + 0.1)) / 1.0) * -2.25) + 3.14159) * v4)
            var v38 int = -1
            var v39 bool = false
            _ = v38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := GenS5{F0: v1, F1: v1, F2: v1}
        var v42 int64 = -687
        if (true && ((v25 - 4294967295) != 2)) {
            v43 := v3[4294967295]
            _ = v43
        } else {
            switch (v1 + (v1 + "foo bar")) {
            case "hello":
                for fi44 := 0; fi44 < 19; fi44++ {
                    v45 := GenS2{F0: "a", F1: 128}
                    v46 := GenS4{F0: 0, F1: 1, F2: 300}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            default:
                var v47 bool = false
                _ = v47
            }
            v1 = "foo bar"
        }
        _ = fi40
        _ = v41
        _ = v42
    }
    for fi48 := 0; fi48 < 3; fi48++ {
        v4 = (0.1 / ((v4 * 0.1) / v4))
        v49 := GenS3{F0: 1, F1: v1, F2: 367}
        delete(v3, 42)
        _ = fi48
        _ = v49
    }
    if true {
        switch v1 {
        case "hello":
            var v50 bool = (true && true)
            _ = v50
        case "":
            var v51 uint64 = 1
            _ = v51
        case "foo bar":
            for fi52 := 0; fi52 < 7; fi52++ {
                if true {
                    if (238 > 737) {
                        v53 := GenS6{F0: 1000}
                        v54 := GenS3{F0: 58, F1: v1, F2: 2147483647}
                        if true {
                            v3[611] = "a"
                            for fi55 := 0; fi55 < 5; fi55++ {
                                v25 += (v25 >> ((18446744073709551615) & 63))
                                _ = fi55
                            }
                            fmt.Println(v1, v25, v1)
                        }
                        _ = v53
                        _ = v54
                    } else {
                        var v56 string = ((v1 + v1) + v1)
                        v25 -= (v25 & 461)
                        if (false && (255 < 255)) {
                            if true {
                                var v57 uint32 = 942
                                _ = v57
                            }
                            v25 += (v25 << (((^(v25))) & 63))
                            v4 -= v4
                        }
                        _ = v56
                    }
                    v5 = (v5 % (((v5 * v5)) | 1))
                } else {
                    for fi58 := 0; fi58 < 20; fi58++ {
                        var v59 int64 = 582
                        v60 := GenS6{F0: 570}
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                    v61 := []int{fi52, -653}
                    _ = v61
                }
                _ = fi52
            }
        default:
            var v62 uint64 = 472
            _ = v62
        }
        switch v1 {
        case "hello":
            fmt.Println(v4)
        case "foo bar":
            v25 -= (^(v25))
        case "":
            v25 = 65535
        default:
            fmt.Println(v25)
        }
        v5 -= (v5 - 100)
    } else {
        v63 := GenS7{F0: "", F1: true}
        switch "a" {
        case "a":
            switch 240 {
            case 127:
                v64 := v2[v1]
                _ = v64
            case 415:
                var v65 int64 = -9223372036854775808
                _ = v65
            }
        case "foo bar":
            var v66 int32 = 2147483647
            _ = v66
        case "":
            if (3.14159 != 1.5) {
                v67 := GenS2{F0: "a", F1: 2}
                v4 -= v4
                var v68 uint16 = 468
                _ = v67
                _ = v68
            } else {
                v69 := []int{9223372036854775807, -2147483648, 9223372036854775807, -2147483648}
                _ = v69
            }
        default:
            fmt.Println(v5, v4, v4)
        }
        fmt.Println(v4)
        _ = v63
    }
    v70 := []int{9223372036854775807, 32768}
    var v71 uint = (v25 - 186)
    for fi72 := 0; fi72 < 13; fi72++ {
        v4 -= (((v4 / float64((0.0 - 0.1))) / ((v4 + v4) / v4)) * v4)
        _ = fi72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v70
    _ = v71
}
