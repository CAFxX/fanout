package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 int; F2 bool }
type GenS8 struct { F0 int16; F1 int; F2 uint64 }
func genFn0(p0 int8, p1 int32) int32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{fi6, 32768, (fi6 & 128), -2147483648}
        v8 := GenS1{F0: (p1 / ((-2147483648) | 1)), F1: (p1 ^ p1), F2: p0}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch int(p1) {
    case -2147483648:
        v9 := GenS1{F0: p1, F1: p1, F2: -128}
        _ = v9
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(p1, p0)
    case "\n\t":
        var v10 int16 = 1
        _ = v10
    case "a":
        switch -9223372036854775808 {
        case 32767:
            v11 := GenS1{F0: 1, F1: 386, F2: p0}
            _ = v11
        default:
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 float64 = ((((0.1 / ((((0.0 - 0.0) - (3.14159 * (0.1 + float64(0.1)))) + 3.14159) + 1e10)) + 3.14159) + 1e10) / 1.0)
                _ = fi12
                _ = v13
            }
        }
    default:
        for fi14 := 0; fi14 < 17; fi14++ {
            p1 = -472
            var v15 float32 = (3.14159 - 0.0)
            var v16 uint64 = 2
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 uint = 18446744073709551615
    v4 := make(map[int]uint32)
    var v5 uint8 = 0
    switch ((("foo bar" + ("a" + "")) + ("\n\t" + "\n\t")) + ("" + (((("a" + "a") + ("hello" + "\n\t")) + "hello") + "hello"))) {
    case "":
        var v17 uint32 = 65535
        _ = v17
    case "foo bar":
        var v18 float32 = 1.5
        _ = v18
    }
    v3 -= v3
    _ = len(v4)
    v19 := GenS2{F0: 42}
    v3 = (v3 / ((v3) | 1))
    v1 = false
    fmt.Println(v3, v5, v1)
    v20 := GenS3{F0: 9223372036854775807, F1: -1}
    v1 = v1
    v1 = (2 != 265)
    if true {
        v3 -= v3
    }
    v21 := GenS4{F0: 65535, F1: -9223372036854775808, F2: 0}
    fmt.Println(v3, v3)
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch ("\n\t" + "") {
            case "hello":
                var v24 int = (fi23 - fi22)
                _ = v24
            case "\n\t":
                if (uint(v5) == v3) {
                    v3 += (v3 | 65535)
                } else {
                    if false {
                        var v25 int32 = -1
                        var v26 float32 = 1.5
                        var v27 int64 = 42
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    if (!(!v1)) {
                        v1 = (-1 >= -128)
                        switch fi22 {
                        case 0:
                            if (-2147483648 <= -675) {
                                v5 += v5
                                if v1 {
                                    v1 = (255 < 2)
                                }
                                switch "a" {
                                case "":
                                    switch "" {
                                    case "a":
                                        v4[(fi22 - 935)] = 943
                                    }
                                case "foo bar":
                                    switch fi23 {
                                    case 32767:
                                        v1 = false
                                    case 2147483647:
                                        v28 := GenS4{F0: 65535, F1: -730, F2: uint16(v3)}
                                        var v29 float64 = 0.1
                                        var v30 string = "a"
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    case 9223372036854775807:
                                        v31 := GenS4{F0: 665, F1: -9223372036854775808, F2: 475}
                                        _ = v31
                                    default:
                                        delete(v2, "hello")
                                    }
                                default:
                                    var v32 bool = false
                                    _ = v32
                                }
                                v3 = v3
                            } else {
                                fmt.Println(fi22, v5, v1)
                                var v33 uint16 = 42
                                _ = v33
                            }
                            v34 := GenS3{F0: 1, F1: 9223372036854775807}
                            v35 := []int{fi22, (fi23 + (fi23 & fi23))}
                            _ = v34
                            _ = v35
                        case -2147483648:
                            v3 = 255
                        }
                    }
                }
            case "":
                fmt.Println(v3)
            default:
                fmt.Println(v1, v5)
            }
            v5 -= 106
            v3 += (v3 / ((v3) | 1))
            _ = fi23
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            if (!v1) {
                switch (^(fi22)) {
                case -117:
                    if (!(v1 && (0 <= 437))) {
                        var v37 float64 = 3.14159
                        _ = v37
                    } else {
                        switch (fi36 | (fi36 ^ 2147483647)) {
                        case 2147483647:
                            if v1 {
                                v38 := GenS5{F0: 2147483648, F1: -2147483648, F2: v5}
                                _ = v38
                            }
                        case 42:
                            if v1 {
                                if false {
                                    _ = len(v2)
                                    v39 := v2["a"]
                                    _ = v39
                                } else {
                                    if (!(636 >= 2)) {
                                        v40 := GenS6{F0: 1}
                                        _ = v40
                                    }
                                }
                            }
                            fmt.Println(v1, v1, v3)
                        case 9223372036854775807:
                            _ = len(v2)
                        default:
                            switch (fi22 * (fi36 * fi36)) {
                            case 2147483647:
                                v1 = (!(!(!(!v1))))
                            case -932:
                                v41 := v2["\n\t"]
                                _ = v41
                            case 9223372036854775807:
                                v42 := []int{(fi22 - fi22), 127, fi36, fi36}
                                _ = v42
                            default:
                                v43 := GenS5{F0: 1, F1: 2147483647, F2: 0}
                                _ = v43
                            }
                        }
                        fmt.Println(fi36, v1, fi36)
                        switch fi22 {
                        case 0:
                            var v44 uint16 = 42
                            _ = v44
                        case 127:
                            v45 := []int{(fi36 + (fi36 / (((fi36 & fi36)) | 1)))}
                            if len(v45) >= 2 {
                                v45 = v45[1:len(v45)]
                            }
                            v1 = ((v3 % ((v3) | 1)) <= (v3 % (((v3 % ((65535) | 1))) | 1)))
                            _ = v45
                        }
                    }
                case 42:
                    fmt.Println(fi22, v5)
                case -2147483648:
                    v5 += (^(v5))
                    var v46 string = ("\n\t" + (("hello" + "foo bar") + "a"))
                    var v47 float64 = (float64((1.5 + 3.14159)) / (float64(((-2.25 - 1.5) + -2.25)) * 3.14159))
                    _ = v46
                    _ = v47
                default:
                    fmt.Println(v3, v5)
                }
            }
            fmt.Println(v1, v5, v1)
            fmt.Println(v3)
            _ = fi36
        }
        _ = fi22
    }
    if v1 {
        if (!((0.0 > 1.5) || (1 != -92))) {
            if v1 {
                v3 = (v3 >> ((v3) & 63))
                for fi48 := 0; fi48 < 15; fi48++ {
                    var v49 uint64 = uint64(v3)
                    if v1 {
                        switch 4294967295 {
                        case -690:
                            v3 = 0
                        case 32767:
                            var v50 float32 = ((0.1 * 1.5) * 1.5)
                            _ = v50
                        default:
                            fmt.Println(v3)
                        }
                    }
                    fmt.Println(fi48, v5)
                    _ = fi48
                    _ = v49
                }
                var v51 uint32 = 2147483647
                _ = v51
            } else {
                v52 := []int{915}
                if false {
                    fmt.Println(v1, v1, v5)
                    var v53 int64 = 616
                    switch int(v53) {
                    case -2147483648:
                        switch "" {
                        case "hello":
                            _ = len(v4)
                        default:
                            switch int(v5) {
                            case -2147483648:
                                switch "hello" {
                                case "hello":
                                    v5 = 128
                                default:
                                    v2[""] = -2147483648
                                }
                            case 42:
                                v54 := GenS3{F0: v53, F1: (v53 << ((18446744073709551615) & 63))}
                                _ = v54
                            case 431:
                                if (!v1) {
                                    if len(v52) >= 2 {
                                        v52 = v52[1:len(v52)]
                                    }
                                    if (!(42 > 0)) {
                                        _ = len(v4)
                                    }
                                } else {
                                    v1 = (65535 > 255)
                                    var v55 int8 = 0
                                    var v56 int16 = -539
                                    _ = v55
                                    _ = v56
                                }
                            }
                        }
                    default:
                        switch "" {
                        case "":
                            v53 = (v53 * -9223372036854775808)
                        default:
                            var v57 float64 = ((((0.0 * ((0.1 * 1e10) - float64((-2.25 - 1e10)))) - 1e10) / (0.0 + 3.14159)) - 1e10)
                            _ = v57
                        }
                    }
                    _ = v53
                }
                _ = v52
            }
            if false {
                var v58 int64 = 1
                v1 = true
                for fi59 := 0; fi59 < 15; fi59++ {
                    switch "" {
                    case "":
                        for fi60 := 0; fi60 < 15; fi60++ {
                            switch "\n\t" {
                            case "a":
                                v1 = (!((!false) && v1))
                            }
                            if (42 < 745) {
                                v61 := GenS7{F0: 121, F1: -8, F2: true}
                                var v62 float64 = 0.1
                                _ = v61
                                _ = v62
                            } else {
                                v63 := GenS8{F0: 1, F1: 32767, F2: 1}
                                _ = v63
                            }
                            v58 = (^(v58))
                            _ = fi60
                        }
                    case "a":
                        v58 -= v58
                    case "hello":
                        fmt.Println(fi59, fi59)
                    default:
                        v64 := []int{int(v58)}
                        _ = v64
                    }
                    _ = fi59
                }
                _ = v58
            } else {
                fmt.Println(v1, v3, v5)
                var v65 uint16 = 924
                _ = v65
            }
            var v66 float32 = 1e10
            _ = v66
        }
        v3 -= v3
        v5 = (v5 * (v5 << ((0) & 63)))
    }
    v67 := GenS2{F0: 93}
    fmt.Println(v5)
    v1 = (false || (!((v5 != (v5 & v5)) || v1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v67
}
