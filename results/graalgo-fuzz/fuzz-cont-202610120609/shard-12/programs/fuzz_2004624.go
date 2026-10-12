package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int32; F1 int64; F2 int32 }
type GenS6 struct { F0 int8; F1 string; F2 bool }
type GenS7 struct { F0 int32 }
func genFn0(p0 int8, p1 int) bool {
    var v4 bool = true
    for fi5 := 0; fi5 < 3; fi5++ {
        switch "\n\t" {
        case "\n\t":
            if (1 == -1) {
                if (p1 < (fi5 * -765)) {
                    fmt.Println(v4, v4)
                    var v6 uint8 = 93
                    _ = v6
                }
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 float64 = 0.0
                    if (42 == 0) {
                        switch "" {
                        case "foo bar":
                            var v9 uint = 1
                            _ = v9
                        case "hello":
                            v10 := []int{}
                            _ = v10
                        case "":
                            fmt.Println(v8)
                        default:
                            v11 := []int{}
                            _ = v11
                        }
                        if (v4 && (v4 || (65535 == 1))) {
                            v12 := GenS1{F0: 1}
                            v13 := GenS2{F0: 2147483647, F1: -645}
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS2{F0: 0, F1: fi7}
                        v15 := GenS3{F0: 1, F1: fi5}
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int32 = -1
                        v4 = (0 < 1)
                        switch (fi5 << ((42) & 63)) {
                        case -2147483648:
                            var v17 int8 = p0
                            _ = v17
                        case 4294967295:
                            fmt.Println(v16, p0, fi5)
                        case 972:
                            fmt.Println(p1, v8, v16)
                        default:
                            var v18 uint64 = 85
                            _ = v18
                        }
                        if true {
                            var v19 float32 = 3.14159
                            v19 = 1.5
                            var v20 bool = v4
                            _ = v19
                            _ = v20
                        }
                        v21 := []int{fi5, 674, (fi7 ^ (p1 | fi5)), 2147483647}
                        _ = v16
                        _ = v21
                    }
                    _ = fi7
                    _ = v8
                }
            }
        default:
            switch "hello" {
            case "":
                var v22 string = ("a" + "foo bar")
                _ = v22
            case "\n\t":
                switch fi5 {
                case 0:
                    switch p1 {
                    case 219:
                        fmt.Println(v4)
                    }
                case 9223372036854775807:
                    v23 := []int{-9223372036854775808, p1}
                    _ = v23
                case -993:
                    switch (p1 ^ (p1 - (-(fi5)))) {
                    case 218:
                        v24 := GenS4{F0: 4294967295, F1: 127}
                        _ = v24
                    case 1:
                        switch ("" + "a") {
                        case "hello":
                            if (1.5 == (0.0 + 1e10)) {
                                switch ("foo bar" + ("" + ("foo bar" + ("hello" + ("\n\t" + (((("\n\t" + "hello") + "a") + (("hello" + ("\n\t" + ("hello" + (("hello" + (("foo bar" + ("a" + "\n\t")) + "a")) + "foo bar")))) + "\n\t")) + "a")))))) {
                                case "a":
                                    v25 := []int{-430, 32768, (fi5 | fi5)}
                                    _ = v25
                                }
                                switch p1 {
                                case -296:
                                    if (1 < 0) {
                                        v26 := []int{}
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        _ = v26
                                    } else {
                                        for fi27 := 0; fi27 < 11; fi27++ {
                                            p1 = p1
                                            v4 = false
                                            p1 += 32767
                                            _ = fi27
                                        }
                                        v28 := GenS4{F0: 18446744073709551615, F1: 85}
                                        var v29 uint32 = 876
                                        _ = v28
                                        _ = v29
                                    }
                                case 32768:
                                    var v30 int8 = p0
                                    _ = v30
                                }
                                var v31 float64 = -2.25
                                _ = v31
                            } else {
                                var v32 float32 = (-2.25 + float32(0.1))
                                var v33 int8 = (p0 >> ((18446744073709551615) & 63))
                                var v34 uint = 1
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        case "a":
                            switch p1 {
                            case -9223372036854775808:
                                fmt.Println(p1, v4)
                            }
                        case "foo bar":
                            if ("a" <= "") {
                                fmt.Println(v4, p1, p1)
                                v35 := GenS2{F0: -1, F1: p1}
                                _ = v35
                            } else {
                                if v4 {
                                    v4 = ((-2.25 + 1.5) > 0.0)
                                }
                            }
                        }
                    case 42:
                        var v36 float64 = (((1.5 - 0.1) * 3.14159) - 3.14159)
                        _ = v36
                    }
                }
            case "foo bar":
                var v37 float64 = 1e10
                _ = v37
            }
        }
        _ = fi5
    }
    if v4 {
        if (!v4) {
            p1 -= int(p0)
        }
    } else {
        var v38 int64 = -472
        _ = v38
    }
    _ = p0
    _ = p1
    _ = v4
    return (false && false)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v39 int8 = -1
    var v40 int8 = (v39 % ((1) | 1))
    v39 = 1
    var v41 int32 = -1
    var v42 uint64 = 625
    v41 += (v41 % (((v41 << ((4294967295) & 63))) | 1))
    if true {
        fmt.Println(v39, v40, v42)
    }
    v42 += (^(v42))
    v43 := v3[("a" + "hello")]
    v3[v43] = v43
    if (761 <= 1) {
        for fi44 := 0; fi44 < 6; fi44++ {
            for fi45 := 0; fi45 < 15; fi45++ {
                switch 1 {
                case -686:
                    v42 += (v42 / ((v42) | 1))
                    fmt.Println(v41, fi45, fi44)
                    switch "a" {
                    case "":
                        v41 -= (v41 + v41)
                    case "\n\t":
                        if false {
                            v40 = (v39 & (^(v40)))
                        } else {
                            switch (fi44 % ((-346) | 1)) {
                            case 875:
                                var v46 int = 127
                                _ = v46
                            case 9223372036854775807:
                                v47 := []int{fi45, (fi45 + fi44), (fi45 | fi45), fi44}
                                _ = v47
                            case -9223372036854775808:
                                v39 = v40
                            }
                            if (-1 == 9223372036854775807) {
                                fmt.Println(v40, fi44)
                                v48 := []int{fi44, fi45}
                                _ = v48
                            }
                            fmt.Println(v41, v39)
                        }
                    default:
                        v49 := GenS5{F0: (v41 << ((4294967295) & 63)), F1: -9223372036854775808, F2: -1}
                        _ = v49
                    }
                default:
                    if (!(456 != 484)) {
                        switch ("foo bar" + "a") {
                        case "hello":
                            v50 := GenS1{F0: -198}
                            _ = v50
                        case "":
                            v40 += 42
                        default:
                            _ = len(v1)
                        }
                        v51 := GenS4{F0: 1, F1: 81}
                        _ = v51
                    }
                }
                if true {
                    fmt.Println(fi45, v40, v40)
                } else {
                    v52 := GenS6{F0: (v39 | v39), F1: v43, F2: (4294967295 > v42)}
                    var v53 string = "a"
                    _ = v52
                    _ = v53
                }
                v54 := []int{-2147483648, (fi45 >> ((1) & 63)), fi45, (^(fi44))}
                _ = fi45
                _ = v54
            }
            delete(v3, "a")
            _ = fi44
        }
        var v55 int64 = 9223372036854775807
        for fi56 := 0; fi56 < 7; fi56++ {
            var v57 int64 = (v55 / ((v55) | 1))
            switch fi56 {
            case 32767:
                v58 := GenS2{F0: (v41 / ((v41) | 1)), F1: (fi56 * (fi56 ^ fi56))}
                _ = v58
            case -1:
                var v59 int32 = (^(v41))
                _ = v59
            }
            _ = fi56
            _ = v57
        }
        _ = v55
    }
    var v60 int8 = 0
    v61 := GenS7{F0: v41}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v60
    _ = v61
}
