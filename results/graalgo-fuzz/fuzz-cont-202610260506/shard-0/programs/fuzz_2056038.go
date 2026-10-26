package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 string }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 int8; F1 uint32; F2 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 uint8 }
func genFn0(p0 int8, p1 int32, p2 uint16) uint8 {
    v4 := GenS1{F0: p0}
    v5 := GenS2{F0: p0, F1: 42, F2: -128}
    switch 9223372036854775807 {
    case 32767:
        p1 += 1
    }
    p2 = p2
    var v6 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 92
}

func main() {
    var v1 int32 = 183
    var v2 uint = 4294967295
    v3 := make(map[uint32]bool)
    v2 += v2
    _ = len(v3)
    v1 = v1
    v7 := GenS3{F0: 1, F1: -9223372036854775808, F2: ""}
    if (false && true) {
        switch "" {
        case "foo bar":
            v2 += (v2 << ((v2) & 63))
        case "":
            var v8 uint64 = 1
            _ = v8
        case "hello":
            delete(v3, 17)
        }
        fmt.Println(v2, v1)
    } else {
        switch 0 {
        case 0:
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := []int{fi9, (fi9 ^ (fi9 | (-(fi9))))}
                var v11 int16 = 1000
                _ = fi9
                _ = v10
                _ = v11
            }
        case 128:
            v3[586] = false
        case 2147483647:
            v1 = 42
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 float64 = 0.0
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 string = "\n\t"
                v16 := GenS3{F0: 4294967295, F1: 894, F2: "a"}
                switch (v15 + (v15 + v15)) {
                case "\n\t":
                    fmt.Println(v2, fi12, v2)
                case "foo bar":
                    v1 = -928
                case "a":
                    v17 := []int{-1, int(v2)}
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                v1 += (v1 << ((v2) & 63))
                var v19 int32 = v1
                switch "a" {
                case "foo bar":
                    v20 := []int{}
                    _ = v20
                case "":
                    v21 := []int{(^(fi12)), (fi18 % ((fi18) | 1))}
                    _ = v21
                default:
                    var v22 bool = false
                    switch "foo bar" {
                    case "foo bar":
                        var v23 bool = ((!(!(uint32(v2) > 2147483647))) || v22)
                        _ = v23
                    case "a":
                        if false {
                            v24 := []int{int(v2)}
                            if v22 {
                                v22 = (876 != 2)
                                var v25 uint = (v2 | v2)
                                _ = v25
                            }
                            var v26 uint32 = 2
                            _ = v24
                            _ = v26
                        } else {
                            fmt.Println(v19, fi12, fi18)
                        }
                    case "":
                        if (!((v2 % ((v2) | 1)) == 1)) {
                            switch "hello" {
                            case "":
                                var v27 bool = false
                                _ = v27
                            }
                            var v28 int = (fi12 >> ((2) & 63))
                            v29 := GenS2{F0: 1, F1: 191, F2: 127}
                            _ = v28
                            _ = v29
                        }
                    default:
                        switch ("foo bar" + ("\n\t" + "a")) {
                        case "\n\t":
                            v30 := GenS4{F0: 42, F1: -1, F2: v22}
                            _ = v30
                        case "hello":
                            fmt.Println(v1, fi18, v1)
                            var v31 bool = true
                            _ = len(v3)
                            _ = v31
                        default:
                            switch 2147483647 {
                            case 42:
                                var v32 float64 = v13
                                _ = v32
                            case 838:
                                if ((!("" >= "\n\t")) || (467 <= 4294967295)) {
                                    v33 := v3[2147483648]
                                    var v34 uint16 = 65535
                                    _ = v33
                                    _ = v34
                                }
                            default:
                                v35 := GenS5{F0: 3, F1: 2147483647, F2: v19}
                                _ = v35
                            }
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v13)
                            case "foo bar":
                                v36 := []int{4294967295}
                                _ = v36
                            default:
                                v37 := GenS2{F0: -44, F1: 127, F2: -128}
                                _ = v37
                            }
                            if (-117 <= 100) {
                                v38 := GenS1{F0: -128}
                                var v39 int8 = 47
                                v13 += 0.1
                                _ = v38
                                _ = v39
                            }
                            if (1 < 1) {
                                v19 -= v19
                                v22 = ((((v13 * v13) + (0.0 + v13)) + v13) < 1e10)
                            }
                        }
                    }
                    _ = v22
                }
                v13 = 3.14159
                _ = fi18
                _ = v19
            }
            _ = fi12
            _ = v13
        }
    }
    switch "hello" {
    case "":
        if (0.1 > 0.1) {
            v40 := GenS6{F0: (^(v1))}
            if (1 > 796) {
                v2 -= v2
            }
            _ = v40
        } else {
            v2 = 18446744073709551615
        }
    case "foo bar":
        var v41 uint = 767
        _ = v41
    case "a":
        for fi42 := 0; fi42 < 1; fi42++ {
            delete(v3, 2147483648)
            var v43 uint64 = 724
            for fi44 := 0; fi44 < 18; fi44++ {
                var v45 bool = false
                fmt.Println(v45, v45, v45)
                v46 := []int{(fi42 / ((fi42) | 1)), fi42}
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = fi42
            _ = v43
        }
    }
    v1 += -86
    for fi47 := 0; fi47 < 11; fi47++ {
        var v48 int = (fi47 & fi47)
        for fi49 := 0; fi49 < 13; fi49++ {
            if (true || false) {
                switch v48 {
                case -23:
                    fmt.Println(v2)
                case 1:
                    switch (v48 ^ int(v2)) {
                    case -2147483648:
                        if ((v2 * 966) == 1) {
                            var v50 uint16 = 32767
                            v48 = v48
                            _ = v50
                        } else {
                            var v51 float64 = float64((3.14159 + 0.0))
                            var v52 int = fi47
                            _ = v51
                            _ = v52
                        }
                    case 2147483647:
                        if (!(int64(fi47) <= 42)) {
                            v53 := GenS6{F0: (v1 >> ((v2) & 63))}
                            v54 := []int{(v48 << (((v2 & (v2 % ((v2) | 1)))) & 63))}
                            _ = v53
                            _ = v54
                        }
                    case 42:
                        v55 := GenS7{F0: 0}
                        _ = v55
                    default:
                        switch "a" {
                        case "a":
                            switch (fi47 & fi47) {
                            case -98:
                                switch v48 {
                                case 1:
                                    v48 = (v48 << ((v2) & 63))
                                    var v56 bool = false
                                    var v57 float64 = 0.0
                                    _ = v56
                                    _ = v57
                                default:
                                    var v58 uint64 = 4294967295
                                    _ = v58
                                }
                            case 32767:
                                _ = len(v3)
                            }
                        default:
                            var v59 bool = false
                            _ = v59
                        }
                        var v60 bool = ((false && true) && (75 >= v2))
                        var v61 int8 = -128
                        _ = v60
                        _ = v61
                    }
                case 894:
                    v62 := GenS8{F0: (v1 / ((v1) | 1))}
                    if (!true) {
                        var v63 float32 = (float32(float64(0.1)) + 1.5)
                        _ = v63
                    }
                    v64 := []int{(^(fi47)), -486}
                    _ = v62
                    _ = v64
                }
                v65 := GenS9{F0: (v2 << (((^(v2))) & 63)), F1: uint16(v2)}
                v2 = 65535
                _ = v65
            } else {
                switch (fi49 - v48) {
                case -2147483648:
                    var v66 float64 = -2.25
                    _ = v66
                default:
                    switch 32768 {
                    case 0:
                        var v67 int8 = 100
                        _ = v67
                    case 9223372036854775807:
                        var v68 uint = 764
                        _ = v68
                    case 4294967295:
                        if false {
                            v48 -= 1
                            v69 := GenS1{F0: 0}
                            if false {
                                _ = len(v3)
                            }
                            if false {
                                v1 = int32(fi47)
                            }
                            _ = v69
                        }
                    }
                }
                fmt.Println(v2)
            }
            _ = fi49
        }
        fmt.Println(fi47, v2)
        _ = fi47
        _ = v48
    }
    v70 := GenS1{F0: -128}
    if (true || true) {
        v2 += v2
        if (2147483647 == 32768) {
            v71 := GenS3{F0: 816, F1: -9223372036854775808, F2: "a"}
            _ = v71
        } else {
            switch -2147483648 {
            case -388:
                for fi72 := 0; fi72 < 15; fi72++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v73 uint32 = 2
                        _ = v73
                    case "\n\t":
                        if (false && (0 == 75)) {
                            v2 = (v2 << ((v2) & 63))
                            var v74 uint8 = 33
                            if (!true) {
                                if true {
                                    v75 := GenS4{F0: 2, F1: 127, F2: (((true && (-1 >= -23)) && (!false)) && ((1e10 <= 1.5) || true))}
                                    switch -576 {
                                    case 127:
                                        var v76 uint = uint(v74)
                                        _ = v76
                                    }
                                    v77 := GenS5{F0: -128, F1: 984, F2: (v1 << ((v2) & 63))}
                                    _ = v75
                                    _ = v77
                                }
                                fmt.Println(fi72, v74)
                                delete(v3, 65535)
                            }
                            _ = v74
                        }
                    default:
                        var v78 uint = v2
                        _ = v78
                    }
                    if (v2 == 65535) {
                        if false {
                            if (255 > genFn0(42, v1, 747)) {
                                if ("" < "foo bar") {
                                    var v79 float32 = 1.5
                                    var v80 int = -9223372036854775808
                                    _ = v79
                                    _ = v80
                                }
                                v2 += 2
                            } else {
                                v1 += v1
                            }
                        } else {
                            fmt.Println(v1)
                            if true {
                                switch (fi72 + (fi72 / (((fi72 << ((224) & 63))) | 1))) {
                                case 32768:
                                    v81 := []int{2147483647, (fi72 >> ((v2) & 63)), (-(fi72)), (fi72 * fi72)}
                                    _ = v81
                                default:
                                    if true {
                                        if ((-(v1)) != -2147483648) {
                                            for fi82 := 0; fi82 < 17; fi82++ {
                                                v1 -= v1
                                                v3[0] = true
                                                var v83 uint = 42
                                                _ = fi82
                                                _ = v83
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi72, fi72, fi72)
                                    }
                                }
                                v1 -= v1
                            }
                            v84 := GenS10{F0: 60}
                            _ = v84
                        }
                    }
                    _ = fi72
                }
            }
        }
        v85 := GenS8{F0: (^(v1))}
        _ = v85
    }
    v86 := []int{-817}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v86:", len(v86), v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v70
    _ = v86
}
