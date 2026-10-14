package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 string, p1 uint16, p2 uint16) uint64 {
    v4 := []int{32768, -9223372036854775808, -1}
    v5 := GenS1{F0: 128}
    switch p0 {
    case "\n\t":
        fmt.Println(p2, p0)
    case "a":
        var v6 int8 = -128
        _ = v6
    default:
        p2 = 65535
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 625
}

func main() {
    var v1 float64 = (((0.0 / 1e10) + 1.5) / (-2.25 - 0.0))
    v2 := []uint32{998, 42, 871}
    var v3 float64 = v1
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 128}
            _ = fi7
            _ = v8
        }
        var v9 int64 = 42
        _ = v9
    }
    switch "foo bar" {
    case "":
        v3 = (1.5 * v1)
    default:
        switch 415 {
        case -202:
            v10 := GenS1{F0: 192}
            _ = v10
        case 32767:
            v11 := GenS1{F0: 1}
            _ = v11
        default:
            var v12 float32 = float32(v1)
            _ = v12
        }
    }
    var v13 uint32 = 1
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            if (235 != 255) {
                v16 := GenS1{F0: 255}
                var v17 int64 = -9223372036854775808
                _ = v16
                _ = v17
            } else {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v18 := GenS1{F0: 26}
                    v13 = v13
                    var v19 int16 = 42
                    _ = v18
                    _ = v19
                }
                v2 = append(v2, (v13 & (^(v13))))
            }
            var v20 uint64 = 509
            _ = fi15
            _ = v20
        }
        v13 = (v13 & v13)
        v1 = float64((0.1 - 0.1))
        _ = fi14
    }
    if true {
        v3 -= (0.0 * (v3 * float64((0.1 + -2.25))))
        switch -2147483648 {
        case 354:
            fmt.Println(v3, v13)
        case 4294967295:
            for fi21 := 0; fi21 < 12; fi21++ {
                switch "a" {
                case "\n\t":
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := GenS2{F0: 42, F1: (3.14159 <= 0.0)}
                        _ = fi22
                        _ = v23
                    }
                case "a":
                    v24 := GenS2{F0: 1, F1: false}
                    _ = v24
                case "hello":
                    var v25 int = fi21
                    _ = v25
                }
                fmt.Println(fi21, v3)
                _ = fi21
            }
        }
        if (false || false) {
            var v26 string = "foo bar"
            _ = v26
        }
    } else {
        if (785 >= 32768) {
            var v27 uint16 = 0
            var v28 int32 = 237
            if (912 > 379) {
                for fi29 := 0; fi29 < 17; fi29++ {
                    if len(v2) > 0 {
                        v2[0] = (v13 * v13)
                    }
                    _ = fi29
                }
                v30 := GenS2{F0: v28, F1: true}
                _ = v30
            } else {
                var v31 uint8 = 0
                var v32 uint8 = 127
                v31 = v32
                _ = v31
                _ = v32
            }
            _ = v27
            _ = v28
        } else {
            for fi33 := 0; fi33 < 12; fi33++ {
                v3 += ((float64(1e10) + v1) / ((v1 + float64((-2.25 / -2.25))) * 1e10))
                var v34 uint = 2
                v13 = 0
                _ = fi33
                _ = v34
            }
            if false {
                v35 := GenS3{F0: 0, F1: -1, F2: (v13 / ((v13) | 1))}
                var v36 int16 = 0
                _ = v35
                _ = v36
            }
        }
        var v37 uint32 = (^(v13))
        var v38 int64 = 1
        _ = v37
        _ = v38
    }
    switch 0 {
    case 9223372036854775807:
        for fi39 := 0; fi39 < 12; fi39++ {
            if ((float64(1e10) - v1) > v1) {
                if (1.5 >= 0.1) {
                    _ = len(v2)
                    v13 = 2147483647
                } else {
                    fmt.Println(fi39, v13, fi39)
                }
                v1 = float64(-2.25)
                for fi40 := 0; fi40 < 2; fi40++ {
                    if len(v2) > 0 {
                        v2[0] = (v13 << ((255) & 63))
                    }
                    _ = fi40
                }
            } else {
                if (true || true) {
                    v1 -= v1
                } else {
                    var v41 uint = 0
                    fmt.Println(v41, v41)
                    _ = v41
                }
                _ = len(v2)
                v1 -= v3
            }
            _ = fi39
        }
    case 636:
        v42 := GenS2{F0: 42, F1: (!true)}
        _ = v42
    case -164:
        if (((32767 == 892) && (652 > v13)) && false) {
            v43 := GenS2{F0: 42, F1: true}
            if false {
                if true {
                    for fi44 := 0; fi44 < 7; fi44++ {
                        v45 := GenS3{F0: 0, F1: 913, F2: (v13 % ((0) | 1))}
                        fmt.Println(v13, fi44, v13)
                        if (true || ((42 != 246) && true)) {
                            v13 += v13
                            v13 += 2
                            switch "" {
                            case "":
                                fmt.Println(v1)
                            case "hello":
                                v1 = ((v3 + ((float64((3.14159 + 1e10)) - -2.25) - (v3 / v3))) / (float64(-2.25) - 0.0))
                            }
                        }
                        _ = fi44
                        _ = v45
                    }
                    var v46 int64 = 616
                    _ = v46
                } else {
                    switch 1 {
                    case 128:
                        switch "hello" {
                        case "\n\t":
                            var v47 int32 = 704
                            _ = v47
                        case "hello":
                            fmt.Println(v13, v1, v3)
                        default:
                            v3 -= v1
                        }
                    case 32768:
                        v13 = (v13 << ((4294967295) & 63))
                    default:
                        for fi48 := 0; fi48 < 19; fi48++ {
                            var v49 int16 = -1
                            v13 = v13
                            _ = fi48
                            _ = v49
                        }
                    }
                }
            }
            _ = v43
        }
    }
    v50 := GenS4{F0: 255, F1: 32767}
    var v51 uint64 = 827
    v3 -= v3
    for fi52 := 0; fi52 < 7; fi52++ {
        var v53 int = (-(fi52))
        _ = fi52
        _ = v53
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        v13 = v13
        v55 := GenS3{F0: -758, F1: -909, F2: (v13 + 351)}
        _ = fi54
        _ = v55
    }
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (-1000 == 9223372036854775807) {
        v51 = (^(v51))
        if (4294967295 <= v13) {
            v51 -= (v51 | v51)
            v1 += v3
            var v56 int16 = 32767
            _ = v56
        } else {
            var v57 bool = ("a" != "hello")
            _ = v57
        }
    }
    var v58 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v58
}
