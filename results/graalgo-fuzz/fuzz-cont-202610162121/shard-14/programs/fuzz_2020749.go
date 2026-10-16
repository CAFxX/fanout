package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 string }
func genFn0(p0 int) uint8 {
    p0 = (p0 | (^(p0)))
    p0 = 1
    switch ("" + "") {
    case "foo bar":
        for fi6 := 0; fi6 < 2; fi6++ {
            switch p0 {
            case 42:
                v7 := []int{83, (fi6 ^ (^(p0))), fi6}
                _ = v7
            }
            fmt.Println(fi6, p0)
            _ = fi6
        }
    case "":
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                fmt.Println(p0, fi8, fi9)
                fmt.Println(fi8, p0, fi8)
                var v10 uint64 = 244
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 0}
            _ = fi8
            _ = v11
        }
    }
    _ = p0
    return 0
}
func genFn1(p0 string, p1 uint32) string {
    for fi12 := 0; fi12 < 5; fi12++ {
        if (42 < 31) {
            var v13 int8 = -1
            if false {
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS2{F0: 380, F1: genFn0(fi14), F2: 42}
                    v16 := GenS3{F0: 0, F1: (4294967295 != 65535)}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 float32 = 1.5
                switch (fi12 % (((^(fi12))) | 1)) {
                case -9223372036854775808:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v13 += (v13 * v13)
                        v19 := []int{-505}
                        _ = fi18
                        _ = v19
                    }
                case 1:
                    if false {
                        fmt.Println(v17, v17, v17)
                        var v20 int = 9223372036854775807
                        for fi21 := 0; fi21 < 9; fi21++ {
                            var v22 int = v20
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                case 42:
                    v23 := GenS4{F0: 382, F1: genFn0((fi12 % ((fi12) | 1)))}
                    _ = v23
                }
                _ = v17
            } else {
                for fi24 := 0; fi24 < 8; fi24++ {
                    fmt.Println(v13)
                    var v25 int8 = (^(v13))
                    if false {
                        v26 := GenS3{F0: 32767, F1: false}
                        v27 := []int{-166, fi12, fi12}
                        var v28 uint8 = 5
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                }
                v29 := GenS4{F0: 42, F1: 55}
                if true {
                    if true {
                        switch p0 {
                        case "a":
                            switch p0 {
                            case "a":
                                for fi30 := 0; fi30 < 18; fi30++ {
                                    var v31 uint8 = uint8(p1)
                                    if ((float32((1e10 + -2.25)) + 0.0) < 0.0) {
                                        v32 := GenS2{F0: 832, F1: (v31 << ((1) & 63)), F2: -2147483648}
                                        v31 += 128
                                        _ = v32
                                    } else {
                                        if (2147483648 < p1) {
                                            v33 := GenS3{F0: 1000, F1: (false || true)}
                                            _ = v33
                                        }
                                        p1 += (^(p1))
                                    }
                                    v34 := []int{}
                                    _ = fi30
                                    _ = v31
                                    _ = v34
                                }
                            case "foo bar":
                                for fi35 := 0; fi35 < 3; fi35++ {
                                    p1 += p1
                                    _ = fi35
                                }
                            case "":
                                v36 := GenS1{F0: 42}
                                _ = v36
                            }
                        case "\n\t":
                            var v37 uint64 = 42
                            _ = v37
                        default:
                            v38 := GenS5{F0: -420}
                            _ = v38
                        }
                    }
                    fmt.Println(p1)
                } else {
                    if false {
                        v39 := []int{fi12, fi12, -796}
                        var v40 float32 = 0.0
                        _ = v39
                        _ = v40
                    }
                    if true {
                        var v41 uint8 = 209
                        v42 := []int{}
                        _ = v41
                        _ = v42
                    }
                    var v43 uint64 = 4294967295
                    _ = v43
                }
                _ = v29
            }
            _ = v13
        }
        if true {
            fmt.Println(fi12, fi12, p0)
            switch fi12 {
            case -2147483648:
                v44 := GenS5{F0: 396}
                _ = v44
            case 127:
                switch (fi12 | (fi12 & fi12)) {
                case -9223372036854775808:
                    if true {
                        var v45 float32 = 1.5
                        for fi46 := 0; fi46 < 7; fi46++ {
                            v47 := []int{(fi46 * 4294967295)}
                            _ = fi46
                            _ = v47
                        }
                        _ = v45
                    }
                case -2147483648:
                    switch p0 {
                    case "\n\t":
                        var v48 string = p0
                        _ = v48
                    default:
                        if false {
                            v49 := []int{-2147483648}
                            if (((fi12 - fi12) > (fi12 + (fi12 << ((2) & 63)))) || (true && true)) {
                                switch ("hello" + "") {
                                case "":
                                    if false {
                                        v49 = append(v49, 1)
                                    } else {
                                        if len(v49) >= 2 {
                                            v49 = v49[1:len(v49)]
                                        }
                                    }
                                case "hello":
                                    p1 = (^(p1))
                                case "\n\t":
                                    p0 = "\n\t"
                                }
                                if len(v49) > 0 {
                                    v49[0] = (fi12 >> ((649) & 63))
                                }
                            } else {
                                p1 += p1
                                p0 = p0
                            }
                            if true {
                                switch p0 {
                                case "foo bar":
                                    for fi50 := 0; fi50 < 8; fi50++ {
                                        if false {
                                            fmt.Println(p1)
                                        } else {
                                            if len(v49) >= 2 {
                                                v49 = v49[1:len(v49)]
                                            }
                                            _ = len(v49)
                                            switch "a" {
                                            case "hello":
                                                p0 = p0
                                            case "a":
                                                p0 = p0
                                            case "":
                                                v51 := GenS6{F0: 42, F1: 42}
                                                _ = v51
                                            }
                                        }
                                        if len(v49) >= 2 {
                                            v49 = v49[1:len(v49)]
                                        }
                                        v49 = append(v49, fi50)
                                        switch (^(fi50)) {
                                        case -889:
                                            fmt.Println(fi50, p0)
                                        default:
                                            switch 9223372036854775807 {
                                            case 9223372036854775807:
                                                var v52 int16 = 1
                                                _ = v52
                                            case -2147483648:
                                                p1 -= p1
                                            }
                                        }
                                        p1 += p1
                                        _ = fi50
                                    }
                                default:
                                    var v53 float32 = 1e10
                                    _ = v53
                                }
                                for fi54 := 0; fi54 < 1; fi54++ {
                                    var v55 float32 = 0.0
                                    if (3.14159 != 0.0) {
                                        var v56 uint = 292
                                        p0 = p0
                                        _ = v56
                                    }
                                    fmt.Println(v55)
                                    v57 := GenS5{F0: -1}
                                    v58 := GenS2{F0: 4294967295, F1: 59, F2: 361}
                                    if (!false) {
                                        var v59 uint32 = (p1 & (^(p1)))
                                        if ((1 > -2147483648) || false) {
                                            fmt.Println(v55, v55, p0)
                                            var v60 int = fi54
                                            var v61 uint32 = p1
                                            v49 = append(v49, -195)
                                            var v62 uint8 = genFn0((fi12 / ((fi54) | 1)))
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        } else {
                                            if (0 == -1) {
                                                v63 := GenS7{F0: -117, F1: -1}
                                                v64 := GenS4{F0: 65535, F1: 2}
                                                p0 = ("foo bar" + p0)
                                                switch 2147483647 {
                                                case -247:
                                                    if true {
                                                        p1 -= p1
                                                    } else {
                                                        var v65 uint16 = 2
                                                        var v66 uint64 = 4294967295
                                                        var v67 int64 = 42
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                }
                                                _ = v63
                                                _ = v64
                                            }
                                            v49 = append(v49, (fi54 - fi12))
                                        }
                                        _ = v59
                                    }
                                    _ = fi54
                                    _ = v55
                                    _ = v57
                                    _ = v58
                                }
                                fmt.Println(p0, p1)
                            } else {
                                v68 := GenS8{F0: true}
                                v69 := GenS9{F0: -2147483648, F1: 1}
                                _ = v68
                                _ = v69
                            }
                            _ = v49
                        } else {
                            var v70 uint = 949
                            fmt.Println(v70)
                            _ = v70
                        }
                    }
                }
            }
            var v71 int64 = 9223372036854775807
            _ = v71
        }
        _ = fi12
    }
    var v72 float64 = ((0.1 + 0.0) / (1.5 - ((3.14159 - (0.0 + (float64(-2.25) + 3.14159))) / -2.25)))
    for fi73 := 0; fi73 < 12; fi73++ {
        v72 = v72
        v74 := GenS8{F0: true}
        _ = fi73
        _ = v74
    }
    _ = p0
    _ = p1
    _ = v72
    return "a"
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 int8 = -76
    var v5 float32 = 0.0
    v5 = (v5 - (v5 - 0.1))
    v75 := GenS10{F0: ("\n\t" + "a")}
    if false {
        v1 += 42
    } else {
        v1 = v1
        v4 += v4
    }
    var v76 int8 = v4
    var v77 int = -9223372036854775808
    _ = len(v3)
    v5 = v5
    var v78 int8 = (-(v4))
    for fi79 := 0; fi79 < 12; fi79++ {
        var v80 uint8 = 44
        _ = fi79
        _ = v80
    }
    switch (v77 >> ((0) & 63)) {
    case 465:
        fmt.Println(v5, v4)
    case -44:
        v81 := v3[("" + (((("\n\t" + "") + "\n\t") + "foo bar") + "foo bar"))]
        _ = v81
    default:
        var v82 uint16 = 32767
        _ = v82
    }
    if ("hello" >= ("hello" + (("foo bar" + ("hello" + "a")) + "a"))) {
        delete(v3, "a")
    } else {
        fmt.Println(v76, v5, v5)
        var v83 int16 = -32768
        v84 := GenS10{F0: "a"}
        _ = v83
        _ = v84
    }
    v85 := GenS1{F0: 32767}
    var v86 int32 = -544
    if ((false && (!(0.1 >= 1.5))) && false) {
        if true {
            v3[(("a" + "") + "hello")] = ("\n\t" + "hello")
        }
    } else {
        for fi87 := 0; fi87 < 12; fi87++ {
            if (1 > -32768) {
                fmt.Println(fi87)
                if true {
                    v88 := []int{v77, (fi87 % (((v77 ^ (-(v77)))) | 1))}
                    for fi89 := 0; fi89 < 18; fi89++ {
                        var v90 float64 = 0.0
                        switch "" {
                        case "":
                            var v91 float64 = float64((((1e10 - v5) / float32((float64(3.14159) + v90))) * float32(1.5)))
                            _ = v91
                        case "hello":
                            fmt.Println(v4, v1, fi87)
                        }
                        _ = fi89
                        _ = v90
                    }
                    v1 = v1
                    _ = v88
                } else {
                    for fi92 := 0; fi92 < 19; fi92++ {
                        v93 := GenS6{F0: v1, F1: 0}
                        _ = fi92
                        _ = v93
                    }
                    fmt.Println(v1)
                }
            }
            _ = fi87
        }
    }
    v86 -= 2147483647
    switch "a" {
    case "a":
        switch ((("\n\t" + "foo bar") + "") + "\n\t") {
        case "\n\t":
            var v94 int32 = int32(v1)
            _ = v94
        case "a":
            if (((v78 >> ((65535) & 63)) != v76) || false) {
                for fi95 := 0; fi95 < 12; fi95++ {
                    var v96 float32 = float32((3.14159 + 0.0))
                    fmt.Println(v4, v4, v5)
                    _ = fi95
                    _ = v96
                }
                if true {
                    v97 := GenS3{F0: 1000, F1: true}
                    v77 = v77
                    fmt.Println(v76, v1)
                    _ = v97
                } else {
                    v86 += 1
                    switch "" {
                    case "":
                        for fi98 := 0; fi98 < 19; fi98++ {
                            var v99 int8 = v76
                            v78 = v76
                            var v100 bool = (true || false)
                            _ = fi98
                            _ = v100
                            _ = v99
                        }
                    }
                }
            }
        case "hello":
            v101 := GenS5{F0: (v1 ^ (v1 >> ((65535) & 63)))}
            _ = v101
        }
    default:
        delete(v3, "a")
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v85
    _ = v86
}
