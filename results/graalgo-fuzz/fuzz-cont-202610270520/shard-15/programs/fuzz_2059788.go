package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS10 struct { F0 uint16; F1 uint8; F2 int64 }
func genFn0(p0 int8, p1 uint, p2 bool) string {
    p0 = p0
    v4 := GenS1{F0: 128}
    if p2 {
        var v5 int = 127
        v6 := GenS1{F0: v5}
        if p2 {
            v7 := GenS1{F0: v5}
            _ = v7
        }
        _ = v5
        _ = v6
    }
    switch "a" {
    case "\n\t":
        for fi8 := 0; fi8 < 11; fi8++ {
            if (!(!p2)) {
                p2 = p2
                fmt.Println(fi8)
            } else {
                if false {
                    var v9 int = fi8
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v9 += v9
                        _ = fi10
                    }
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := []int{-159, v9, -344, fi8}
                        _ = fi11
                        _ = v12
                    }
                    _ = v9
                }
                fmt.Println(p2)
            }
            p1 += p1
            _ = fi8
        }
    default:
        p0 = 0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ("\n\t" + "foo bar")
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 193
    v3 := []int{127, -2147483648, 9223372036854775807, -1}
    if false {
        v2 += v2
        v13 := GenS1{F0: -1}
        _ = v13
    } else {
        var v14 uint16 = 1
        v14 += (^(v14))
        for fi15 := 0; fi15 < 5; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                v14 = 0
                if ((((float64(1e10) - -2.25) / 1.0) + -2.25) != (1e10 + 0.0)) {
                    var v17 int64 = -1
                    if true {
                        switch 128 {
                        case -1:
                            v2 -= v2
                        }
                        v18 := GenS2{F0: 635, F1: -656}
                        _ = v18
                    }
                    v2 = 42
                    var v19 uint64 = 884
                    _ = v17
                    _ = v19
                } else {
                    switch (^(fi16)) {
                    case 1:
                        if (false && (true && true)) {
                            if (!false) {
                                var v20 int8 = 116
                                v14 = 2
                                _ = v20
                            }
                            var v21 int = 4294967295
                            v22 := GenS1{F0: fi16}
                            _ = v21
                            _ = v22
                        } else {
                            if (2147483647 <= 0) {
                                switch (fi15 + (^(fi15))) {
                                case 127:
                                    if (127 < 100) {
                                        v2 = (v2 + 128)
                                        delete(v1, fi15)
                                        v23 := GenS3{F0: int8(v2), F1: 65535}
                                        v24 := GenS4{F0: 42, F1: 270, F2: (v2 / ((v2) | 1))}
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 float64 = 1e10
                                        if len(v3) > 0 {
                                            v3[0] = (fi15 ^ fi15)
                                        }
                                        var v26 float64 = v25
                                        _ = v25
                                        _ = v26
                                    }
                                default:
                                    switch fi15 {
                                    case 9223372036854775807:
                                        var v27 int32 = 2147483647
                                        _ = v27
                                    case 75:
                                        var v28 bool = false
                                        _ = v28
                                    case 32768:
                                        v29 := GenS5{F0: 2}
                                        _ = v29
                                    default:
                                        if false {
                                            if (956 >= 42) {
                                                if (fi15 > fi15) {
                                                    fmt.Println(fi16, fi15)
                                                } else {
                                                    var v30 float64 = 0.1
                                                    var v31 int64 = -820
                                                    v32 := GenS6{F0: 1, F1: ""}
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                }
                                                v14 += v14
                                            } else {
                                                fmt.Println(fi16, fi15, v2)
                                                fmt.Println(v2)
                                            }
                                            fmt.Println(fi16, v14, fi15)
                                            if len(v3) > 0 {
                                                v3[0] = (fi16 - -1)
                                            }
                                        }
                                    }
                                }
                                var v33 int8 = -128
                                _ = v33
                            }
                            v34 := GenS7{F0: 149}
                            _ = v34
                        }
                        var v35 int64 = 9223372036854775807
                        _ = v35
                    default:
                        v36 := v1[-9223372036854775808]
                        _ = v36
                    }
                    v37 := GenS5{F0: 4294967295}
                    _ = v37
                }
                _ = fi16
            }
            _ = fi15
        }
        _ = v14
    }
    v2 = v2
    fmt.Println(v2)
    fmt.Println(v2)
    if (true && false) {
        if (18446744073709551615 <= 255) {
            var v38 float32 = (float32(0.1) / 1.0)
            for fi39 := 0; fi39 < 17; fi39++ {
                v40 := v1[475]
                fmt.Println(v2, fi39)
                _ = fi39
                _ = v40
            }
            switch 809 {
            case 947:
                v38 += v38
            case 32768:
                if len(v3) > 0 {
                    v3[0] = 127
                }
            case 787:
                var v41 uint = 268
                _ = v41
            default:
                fmt.Println(v38, v2)
            }
            _ = v38
        }
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        for fi42 := 0; fi42 < 10; fi42++ {
            switch fi42 {
            case 34:
                switch (-(fi42)) {
                case 1:
                    var v43 float32 = 1e10
                    _ = v43
                case -1:
                    switch (fi42 / (((fi42 | (fi42 % ((fi42) | 1)))) | 1)) {
                    case 128:
                        v2 -= v2
                    case 42:
                        fmt.Println(fi42)
                    }
                case 32767:
                    v2 -= v2
                }
            case -989:
                v2 = (^(v2))
            case -1:
                var v44 int = (fi42 % ((9223372036854775807) | 1))
                _ = v44
            }
            v2 = (v2 * (^(v2)))
            for fi45 := 0; fi45 < 11; fi45++ {
                if true {
                    fmt.Println(fi45)
                    v2 = v2
                    v46 := GenS1{F0: fi42}
                    _ = v46
                } else {
                    var v47 bool = ((!true) || (1 > 743))
                    switch ("" + "\n\t") {
                    case "\n\t":
                        v2 = 2
                    case "":
                        switch ("a" + "foo bar") {
                        case "":
                            v48 := GenS8{F0: 1}
                            _ = v48
                        case "a":
                            var v49 uint64 = 2
                            _ = v49
                        default:
                            v2 = v2
                        }
                    case "a":
                        if (true && false) {
                            v2 = (v2 ^ (v2 * (v2 * (v2 | v2))))
                        } else {
                            v2 += v2
                        }
                    }
                    fmt.Println(fi42, fi42, v2)
                    var v50 uint = 2
                    _ = v47
                    _ = v50
                }
                var v51 float64 = (float64((((1.5 / ((3.14159 / 1.0) * 0.1)) / -2.25) / 1.0)) * float64(3.14159))
                _ = fi45
                _ = v51
            }
            _ = fi42
        }
    }
    fmt.Println(v2)
    _ = len(v1)
    v2 += (v2 | v2)
    if (18446744073709551615 > 2) {
        if false {
            var v52 float64 = (1.5 * 1.5)
            _ = v52
        } else {
            var v53 uint64 = 1
            _ = v53
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "a" {
        case "a":
            v2 = v2
        }
    }
    if false {
        v2 += v2
    } else {
        if true {
            if (!true) {
                if true {
                    v2 = v2
                    var v54 float32 = (3.14159 + 0.0)
                    _ = v54
                } else {
                    v3 = append(v3, 127)
                    for fi55 := 0; fi55 < 5; fi55++ {
                        if (true || false) {
                            v2 = v2
                        }
                        _ = fi55
                    }
                }
                fmt.Println(v2)
                if false {
                    var v56 uint8 = 147
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 6; fi57++ {
                        if true {
                            fmt.Println(fi57)
                            v2 += 42
                            if (0.0 < -2.25) {
                                fmt.Println(v2, fi57)
                                var v58 uint8 = v2
                                _ = v58
                            }
                        } else {
                            for fi59 := 0; fi59 < 17; fi59++ {
                                _ = len(v3)
                                if (!true) {
                                    var v60 int8 = 42
                                    v2 = v2
                                    _ = v60
                                } else {
                                    v2 = v2
                                    fmt.Println(fi57)
                                    v61 := GenS9{F0: 2, F1: 16, F2: v2}
                                    _ = v61
                                }
                                _ = fi59
                            }
                        }
                        _ = fi57
                    }
                }
            } else {
                switch "hello" {
                case "a":
                    var v62 float32 = (0.0 / (0.1 * 1.5))
                    _ = v62
                }
            }
            v2 += v2
            var v63 uint32 = 4294967295
            _ = v63
        }
        if ("hello" >= "hello") {
            if false {
                for fi64 := 0; fi64 < 11; fi64++ {
                    for fi65 := 0; fi65 < 8; fi65++ {
                        if len(v3) > 0 {
                            v3[0] = fi64
                        }
                        var v66 bool = false
                        var v67 uint = uint(v2)
                        v68 := GenS10{F0: 973, F1: (v2 % (((v2 | (v2 & (v2 + 110)))) | 1)), F2: 858}
                        v67 += 65535
                        _ = fi65
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    _ = fi64
                }
                switch "foo bar" {
                case "":
                    v2 = (v2 | v2)
                default:
                    var v69 string = "a"
                    _ = v69
                }
                var v70 string = ""
                _ = v70
            }
        }
        if (true || (208 == 65535)) {
            v2 -= (v2 / (((v2 - v2)) | 1))
            var v71 uint8 = 42
            for fi72 := 0; fi72 < 13; fi72++ {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    v71 = v71
                default:
                    var v73 int32 = -49
                    _ = v73
                }
                var v74 int8 = 0
                _ = fi72
                _ = v74
            }
            _ = v71
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
}
