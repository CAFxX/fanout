package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 string; F2 int16 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 int32) int8 {
    var v4 int8 = 100
    var v5 uint = 2
    var v6 string = "a"
    if false {
        fmt.Println(p2)
        var v7 uint = v5
        _ = v7
    } else {
        p0 += p0
        for fi8 := 0; fi8 < 18; fi8++ {
            if (false && ((fi8 > (fi8 >> ((uint(p0)) & 63))) && false)) {
                var v9 uint = (v5 * uint(p1))
                _ = v9
            }
            fmt.Println(v6)
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return -128
}
func genFn1(p0 uint32) uint32 {
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(p0)
        v11 := []int{526, fi10, fi10}
        p0 = 42
        _ = fi10
        _ = v11
    }
    if ((float32(0.1) - 1e10) > -2.25) {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: ((("" + "") + "a") + "hello")}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            if (true && (0.0 != 0.1)) {
                if ((1e10 / 1.0) == (1e10 + (3.14159 * 1.5))) {
                    v15 := GenS1{F0: ("\n\t" + "a")}
                    v16 := GenS1{F0: "foo bar"}
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS1{F0: "\n\t"}
                    v18 := []int{fi14, 128, (fi14 & 128), 928}
                    p0 -= (p0 ^ (p0 << ((255) & 63)))
                    _ = v17
                    _ = v18
                }
                v19 := []int{-1, fi14}
                fmt.Println(p0, p0)
                _ = v19
            } else {
                var v20 int8 = -1
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := GenS1{F0: ""}
                    _ = fi21
                    _ = v22
                }
                fmt.Println(fi14, fi14)
                _ = v20
            }
            _ = fi14
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := GenS1{F0: "foo bar"}
            var v25 uint64 = 972
            p0 -= 254
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        var v26 uint16 = 97
        v26 += 1
        _ = v26
    }
    var v27 uint8 = 199
    _ = p0
    _ = v27
    return p0
}

func main() {
    var v1 int = 128
    var v2 int = v1
    v3 := make(map[int]uint32)
    v2 = (v1 >> ((255) & 63))
    delete(v3, v1)
    v28 := GenS1{F0: "hello"}
    var v29 bool = (42 != 2147483647)
    var v30 string = "foo bar"
    var v31 float32 = (0.1 / (1e10 - 3.14159))
    v30 = (v30 + v30)
    switch 2147483647 {
    case 2147483647:
        switch (v1 | v1) {
        case -542:
            switch v1 {
            case 42:
                for fi32 := 0; fi32 < 18; fi32++ {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 int32 = 390
                        var v35 int64 = -884
                        if (v30 != (("hello" + v30) + "\n\t")) {
                            v36 := GenS1{F0: (v30 + v30)}
                            _ = v36
                        } else {
                            v37 := GenS2{F0: 35}
                            v38 := []int{int(v34), fi32, (-(fi32)), -705}
                            fmt.Println(fi32)
                            _ = v37
                            _ = v38
                        }
                        switch (fi33 ^ (fi33 & 0)) {
                        case 128:
                            v39 := GenS3{F0: 2147483648}
                            _ = v39
                        case -2147483648:
                            v40 := GenS2{F0: 1}
                            _ = v40
                        case 4294967295:
                            if (v29 && v29) {
                                v41 := []int{(fi32 << ((1) & 63))}
                                _ = v41
                            }
                        default:
                            delete(v3, (v1 - 115))
                        }
                        v29 = (18446744073709551615 >= 65535)
                        var v42 uint = 18446744073709551615
                        if v29 {
                            var v43 int8 = 100
                            _ = v43
                        }
                        var v44 uint16 = 1
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v42
                        _ = v44
                    }
                    _ = fi32
                }
            case -627:
                v1 += v1
            }
        case 127:
            v45 := GenS2{F0: 188}
            _ = v45
        default:
            v30 = v30
        }
    default:
        v46 := GenS4{F0: 799}
        _ = v46
    }
    v47 := GenS5{F0: -2147483648, F1: v30, F2: -1}
    switch "\n\t" {
    case "":
        for fi48 := 0; fi48 < 9; fi48++ {
            switch "\n\t" {
            case "":
                if (2147483647 >= 1) {
                    v49 := GenS5{F0: 0, F1: v30, F2: int16(v1)}
                    for fi50 := 0; fi50 < 19; fi50++ {
                        if (!(32767 <= 2)) {
                            var v51 int8 = -101
                            fmt.Println(v29)
                            v52 := GenS1{F0: ""}
                            v29 = false
                            var v53 string = "a"
                            v54 := []int{765, 413}
                            _ = len(v3)
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        if (!v29) {
                            v29 = (v31 > v31)
                        }
                        _ = fi50
                    }
                    _ = v49
                } else {
                    fmt.Println(fi48, fi48, v30)
                    v31 += v31
                    if (genFn0(255, 32767, -2147483648) <= -7) {
                        for fi55 := 0; fi55 < 11; fi55++ {
                            v1 = v1
                            delete(v3, -9223372036854775808)
                            _ = fi55
                        }
                        for fi56 := 0; fi56 < 11; fi56++ {
                            if v29 {
                                var v57 uint8 = 42
                                fmt.Println(v57, v2)
                                _ = v57
                            }
                            switch ((v30 + v30) + "a") {
                            case "foo bar":
                                var v58 uint32 = 42
                                _ = v58
                            case "a":
                                _ = len(v3)
                            }
                            switch ((v30 + "a") + ("\n\t" + (v30 + v30))) {
                            case "foo bar":
                                var v59 int64 = 1
                                _ = v59
                            case "a":
                                var v60 int8 = 0
                                v61 := GenS3{F0: 4294967295}
                                _ = v60
                                _ = v61
                            }
                            _ = fi56
                        }
                    }
                }
            case "a":
                var v62 uint32 = 662
                _ = v62
            case "hello":
                v63 := []int{v2, 2147483647, v1, 971}
                _ = v63
            }
            _ = fi48
        }
    case "foo bar":
        v64 := GenS6{F0: 1, F1: 319}
        _ = v64
    case "hello":
        switch (v30 + "foo bar") {
        case "a":
            fmt.Println(v2, v31, v30)
        case "":
            fmt.Println(v29)
        }
    }
    delete(v3, (v2 >> ((4294967295) & 63)))
    switch ("hello" + "a") {
    case "\n\t":
        var v65 float64 = (0.1 - (((1.5 * 1.5) / (3.14159 - 1e10)) + ((1e10 - 0.0) - 0.1)))
        _ = v65
    default:
        v66 := v3[442]
        _ = v66
    }
    for fi67 := 0; fi67 < 20; fi67++ {
        switch 32767 {
        case 918:
            switch 127 {
            case -511:
                if false {
                    if (100 >= -90) {
                        v68 := []int{v1, v1}
                        var v69 int64 = 1
                        _ = v68
                        _ = v69
                    } else {
                        v70 := GenS7{F0: 127}
                        _ = v70
                    }
                    var v71 float32 = v31
                    _ = v71
                }
            default:
                var v72 int16 = 32767
                _ = v72
            }
        case 2147483647:
            for fi73 := 0; fi73 < 10; fi73++ {
                _ = len(v3)
                _ = fi73
            }
        case 128:
            fmt.Println(v31, v1, v29)
        }
        _ = fi67
    }
    for fi74 := 0; fi74 < 12; fi74++ {
        switch "foo bar" {
        case "foo bar":
            var v75 int8 = 127
            _ = v75
        case "a":
            v30 = (v30 + ("foo bar" + ((v30 + "\n\t") + v30)))
        case "\n\t":
            v76 := GenS2{F0: 32}
            _ = v76
        }
        switch 1 {
        case 42:
            fmt.Println(v30)
        case -2147483648:
            for fi77 := 0; fi77 < 12; fi77++ {
                v1 = v2
                _ = fi77
            }
        default:
            for fi78 := 0; fi78 < 19; fi78++ {
                switch 1 {
                case 331:
                    v79 := GenS8{F0: 4294967295}
                    _ = v79
                case 854:
                    switch (v30 + v30) {
                    case "hello":
                        fmt.Println(v2, v2, v30)
                        if (!(!v29)) {
                            _ = len(v3)
                        } else {
                            v31 = v31
                        }
                    default:
                        fmt.Println(v29, v1, v1)
                    }
                }
                var v80 int32 = -470
                v80 -= (^(v80))
                _ = fi78
                _ = v80
            }
        }
        v30 = v30
        _ = fi74
    }
    fmt.Println(v29, v2)
    if (788 == 542) {
        switch (v30 + ("hello" + (v30 + v30))) {
        case "":
            if (3.14159 == 3.14159) {
                v29 = v29
                v29 = (!(624 < 0))
            }
        case "\n\t":
            switch (v2 * v2) {
            case 32768:
                for fi81 := 0; fi81 < 8; fi81++ {
                    v82 := []int{128}
                    fmt.Println(v2)
                    switch (v1 / ((v2) | 1)) {
                    case 1:
                        fmt.Println(fi81, v29)
                    case 9223372036854775807:
                        switch v30 {
                        case "":
                            v2 += (v1 % ((v1) | 1))
                        case "foo bar":
                            v82 = append(v82, (fi81 * (v1 & (v1 | (v1 & (v2 + (^(v1))))))))
                        case "hello":
                            var v83 bool = v29
                            _ = v83
                        }
                    case -271:
                        var v84 int64 = -316
                        _ = v84
                    default:
                        v1 -= 32767
                    }
                    _ = fi81
                    _ = v82
                }
            case 597:
                fmt.Println(v31, v29, v1)
            case -2147483648:
                v30 = v30
            default:
                var v85 uint = 18446744073709551615
                _ = v85
            }
        default:
            v29 = v29
        }
        if (656 < 946) {
            var v86 uint64 = 2
            var v87 uint = 0
            v88 := []int{440, -2147483648, (v1 >> ((v87) & 63))}
            _ = v86
            _ = v87
            _ = v88
        }
    } else {
        var v89 float64 = 1e10
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v47
}
