package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 bool; F2 int64 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 string }
func genFn0(p0 int32) int16 {
    var v7 int32 = (p0 >> ((1) & 63))
    p0 -= v7
    switch 32768 {
    case 32768:
        p0 -= v7
    case 128:
        var v8 uint64 = 18446744073709551615
        _ = v8
    }
    v9 := []int{69, 1, -1}
    for fi10 := 0; fi10 < 11; fi10++ {
        switch "hello" {
        case "":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        case "a":
            _ = len(v9)
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint = 42
            var v13 float64 = 1e10
            var v14 float64 = (v13 + (((v13 + 0.1) * v13) / v13))
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    _ = p0
    _ = v7
    _ = v9
    return 42
}

func main() {
    var v1 int = -888
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v4 uint64 = 230
    v5 := make(map[uint32]string)
    v6 := make(map[int]int)
    switch "foo bar" {
    case "a":
        if false {
            for fi15 := 0; fi15 < 2; fi15++ {
                for fi16 := 0; fi16 < 13; fi16++ {
                    fmt.Println(v1)
                    v1 = fi15
                    var v17 uint16 = 65535
                    var v18 int16 = 42
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = len(v5)
                if true {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 uint64 = (v4 + (v4 | (v4 % ((189) | 1))))
                        _ = fi19
                        _ = v20
                    }
                }
                _ = fi15
            }
        } else {
            switch "" {
            case "foo bar":
                if false {
                    v21 := GenS1{F0: -1, F1: 9223372036854775807}
                    var v22 float32 = 3.14159
                    _ = v21
                    _ = v22
                }
            case "hello":
                fmt.Println(v4, v1)
            }
            switch "a" {
            case "":
                v23 := []int{(v1 >> ((4294967295) & 63)), v1, v1}
                _ = v23
            case "\n\t":
                switch ("a" + "\n\t") {
                case "foo bar":
                    if (!true) {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            switch (v1 / ((670) | 1)) {
                            case 1:
                                _ = len(v2)
                            default:
                                v4 -= v4
                            }
                            _ = fi24
                        }
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v26 := GenS2{F0: ("\n\t" + "foo bar")}
                            fmt.Println(v1, v1)
                            v27 := GenS3{F0: "\n\t", F1: 0}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    } else {
                        var v28 int8 = -128
                        var v29 uint64 = (v4 & 4294967295)
                        _ = v28
                        _ = v29
                    }
                default:
                    v30 := GenS4{F0: v1}
                    _ = v30
                }
            case "hello":
                if (1e10 <= 0.1) {
                    fmt.Println(v4, v1)
                    v31 := []int{}
                    var v32 int16 = int16(v1)
                    _ = v31
                    _ = v32
                } else {
                    var v33 uint = 65535
                    for fi34 := 0; fi34 < 10; fi34++ {
                        var v35 int16 = 0
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
            default:
                for fi36 := 0; fi36 < 4; fi36++ {
                    v37 := GenS1{F0: -56, F1: v1}
                    v1 -= 2147483647
                    _ = fi36
                    _ = v37
                }
            }
        }
    default:
        var v38 uint = 42
        _ = v38
    }
    var v39 uint32 = 31
    var v40 bool = false
    v39 += v39
    v1 += 128
    v41 := GenS1{F0: -1, F1: 188}
    var v42 uint16 = 32767
    for fi43 := 0; fi43 < 13; fi43++ {
        if false {
            for fi44 := 0; fi44 < 20; fi44++ {
                var v45 uint = 1
                _ = fi44
                _ = v45
            }
        }
        fmt.Println(fi43, v4)
        delete(v3, 2147483647)
        _ = fi43
    }
    for fi46 := 0; fi46 < 6; fi46++ {
        var v47 float64 = -2.25
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        var v49 float32 = 0.1
        v4 = 0
        _ = fi48
        _ = v49
    }
    var v50 bool = (-645 == 42)
    switch 755 {
    case -793:
        switch v1 {
        case 9223372036854775807:
            v42 = v42
        case -9223372036854775808:
            switch ("foo bar" + "hello") {
            case "a":
                v50 = v50
            default:
                var v51 uint8 = 255
                _ = v51
            }
        case 42:
            switch "hello" {
            case "\n\t":
                switch v1 {
                case 2147483647:
                    switch -9223372036854775808 {
                    case 127:
                        for fi52 := 0; fi52 < 13; fi52++ {
                            var v53 int8 = 42
                            v54 := []int{-946}
                            if v50 {
                                for fi55 := 0; fi55 < 13; fi55++ {
                                    v40 = ((1e10 + (3.14159 - 0.1)) >= 0.1)
                                    var v56 float64 = 3.14159
                                    _ = fi55
                                    _ = v56
                                }
                                v4 -= v4
                                fmt.Println(fi52, v53)
                            }
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                    }
                case 9223372036854775807:
                    if true {
                        var v57 uint32 = 2
                        v58 := GenS4{F0: v1}
                        if true {
                            fmt.Println(v40, v39, v57)
                        } else {
                            v59 := GenS3{F0: "foo bar", F1: -977}
                            for fi60 := 0; fi60 < 11; fi60++ {
                                if v50 {
                                    v61 := []int{}
                                    v62 := GenS1{F0: -104, F1: (v1 >> ((42) & 63))}
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := GenS5{F0: 2147483648, F1: v50, F2: 873}
                                    v42 -= (v42 / (((v42 >> ((52) & 63))) | 1))
                                    v64 := v2[(^(fi60))]
                                    _ = v63
                                    _ = v64
                                }
                                if (9223372036854775807 != -605) {
                                    if ((!((("foo bar" + "a") == "\n\t") || v40)) || false) {
                                        for fi65 := 0; fi65 < 2; fi65++ {
                                            if true {
                                                v66 := GenS6{F0: ("" + "a"), F1: (65535 < (^(v42)))}
                                                var v67 bool = false
                                                _ = v66
                                                _ = v67
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    var v68 uint32 = v39
                                                    _ = v68
                                                default:
                                                    switch (v1 % ((0) | 1)) {
                                                    case 2147483647:
                                                        v69 := GenS3{F0: "\n\t", F1: 1}
                                                        _ = v69
                                                    case -110:
                                                        switch (v1 + (-(fi60))) {
                                                        case -61:
                                                            if false {
                                                                switch "" {
                                                                case "":
                                                                    fmt.Println(v42)
                                                                case "hello":
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        var v70 float32 = 0.0
                                                                        _ = v70
                                                                    }
                                                                }
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v39 -= (v39 | v39)
                                                                }
                                                            }
                                                        case -990:
                                                            delete(v2, (fi65 ^ -537))
                                                        case 461:
                                                            var v71 uint8 = 169
                                                            _ = v71
                                                        }
                                                    }
                                                }
                                                if v50 {
                                                    var v72 uint8 = 47
                                                    v42 += uint16(v72)
                                                    if v50 {
                                                        switch "\n\t" {
                                                        case "a":
                                                            switch (fi65 ^ v1) {
                                                            case 32767:
                                                                v4 -= 42
                                                            case -904:
                                                                var v73 int16 = 415
                                                                v1 = v1
                                                                var v74 string = "a"
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        case "\n\t":
                                                            if (-32768 < 42) {
                                                                var v75 uint = 1
                                                                _ = v75
                                                            } else {
                                                                v72 += (^(v72))
                                                                var v76 int = (fi60 + (v1 + 128))
                                                                _ = v76
                                                            }
                                                        case "":
                                                            v77 := GenS1{F0: 1, F1: fi65}
                                                            _ = v77
                                                        default:
                                                            var v78 string = (("foo bar" + "a") + "a")
                                                            _ = v78
                                                        }
                                                        v50 = v50
                                                    } else {
                                                        fmt.Println(v39, v4)
                                                        v50 = ((^(v42)) < 255)
                                                    }
                                                    _ = v72
                                                } else {
                                                    var v79 uint = 903
                                                    _ = v79
                                                }
                                            }
                                            var v80 float32 = 1.5
                                            _ = fi65
                                            _ = v80
                                        }
                                        for fi81 := 0; fi81 < 4; fi81++ {
                                            var v82 uint8 = 255
                                            if false {
                                                fmt.Println(v50)
                                            }
                                            _ = fi81
                                            _ = v82
                                        }
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        switch "a" {
                                        case "a":
                                            for fi83 := 0; fi83 < 2; fi83++ {
                                                v1 = 2147483647
                                                v84 := []int{(v1 / ((314) | 1)), fi60}
                                                if (fi60 == fi83) {
                                                    fmt.Println(v50, v57)
                                                }
                                                _ = fi83
                                                _ = v84
                                            }
                                        case "foo bar":
                                            v85 := GenS7{F0: 0, F1: v42}
                                            _ = v85
                                        case "\n\t":
                                            v42 += 1
                                        }
                                    case "":
                                        v40 = v40
                                    default:
                                        v86 := GenS8{F0: -928, F1: ((fi60 >> ((18446744073709551615) & 63)) >= (fi60 >> ((4294967295) & 63)))}
                                        _ = v86
                                    }
                                }
                                if false {
                                    v57 -= 614
                                    delete(v5, v39)
                                }
                                _ = fi60
                            }
                            var v87 uint64 = v4
                            _ = v59
                            _ = v87
                        }
                        _ = v57
                        _ = v58
                    } else {
                        var v88 int = (v1 | v1)
                        v89 := GenS9{F0: 0}
                        v90 := v6[v88]
                        _ = v88
                        _ = v89
                        _ = v90
                    }
                case 307:
                    fmt.Println(v40, v4)
                default:
                    var v91 float64 = float64((1.5 * ((3.14159 * 1e10) + 1.5)))
                    _ = v91
                }
            }
        default:
            var v92 uint64 = (v4 % ((v4) | 1))
            _ = v92
        }
    case 54:
        fmt.Println(v42)
    case 899:
        var v93 float32 = ((-2.25 - 0.0) * 0.0)
        _ = v93
    default:
        switch (^(v1)) {
        case 128:
            var v94 float32 = -2.25
            _ = v94
        case 2147483647:
            fmt.Println(v50)
        default:
            if true {
                var v95 int64 = -9223372036854775808
                _ = v95
            }
        }
    }
    v40 = v40
    v96 := GenS10{F0: 1, F1: "a"}
    v97 := []int{-2147483648, (v1 >> ((4294967295) & 63)), int(v39)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v97:", len(v97), v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v50
    _ = v6
    _ = v96
    _ = v97
}
