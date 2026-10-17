package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 int }
type GenS4 struct { F0 int8; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint16; F1 int32; F2 int }
type GenS6 struct { F0 bool; F1 uint64; F2 int }
type GenS7 struct { F0 uint; F1 int16; F2 uint }
type GenS8 struct { F0 uint16; F1 string; F2 uint8 }
type GenS9 struct { F0 int8; F1 int8; F2 int8 }
func genFn0() uint64 {
    v4 := GenS1{F0: 32767, F1: 0, F2: -476}
    v5 := GenS1{F0: -691, F1: 9223372036854775807, F2: 9223372036854775807}
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 int) int8 {
    if ((3.14159 * ((1e10 - 0.0) - (((3.14159 + float64(((-2.25 - float32(1.5)) + 0.0))) / ((((1.5 + -2.25) + float64(1e10)) - 1.5) / 1.0)) / 0.1))) <= 3.14159) {
        for fi6 := 0; fi6 < 2; fi6++ {
            switch (p0 ^ (fi6 & p0)) {
            case 2147483647:
                v7 := GenS1{F0: fi6, F1: 1, F2: -9223372036854775808}
                _ = v7
            case 941:
                for fi8 := 0; fi8 < 2; fi8++ {
                    if (100 > -128) {
                        switch "hello" {
                        case "foo bar":
                            switch ("hello" + "hello") {
                            case "\n\t":
                                v9 := []int{(^(p0)), (p0 - (^(fi8)))}
                                _ = v9
                            default:
                                v10 := []int{664, p0, (fi8 * -1)}
                                _ = v10
                            }
                            switch 256 {
                            case 1:
                                var v11 float32 = -2.25
                                _ = v11
                            case 2147483647:
                                if false {
                                    var v12 float32 = 0.1
                                    _ = v12
                                } else {
                                    var v13 int = (fi6 * (p0 % ((fi6) | 1)))
                                    var v14 int = 128
                                    _ = v13
                                    _ = v14
                                }
                            default:
                                if (!true) {
                                    switch 4294967295 {
                                    case -2147483648:
                                        p0 = -6
                                    }
                                    p0 = 602
                                    var v15 uint64 = 4294967295
                                    _ = v15
                                }
                            }
                            var v16 int32 = -2147483648
                            v17 := []int{(fi6 / ((-1) | 1)), 42, p0}
                            v18 := GenS1{F0: int(v16), F1: 955, F2: 0}
                            _ = v16
                            _ = v17
                            _ = v18
                        case "\n\t":
                            fmt.Println(fi6, p0)
                        }
                        p0 = (fi8 * (-(fi6)))
                    }
                    if false {
                        fmt.Println(fi6)
                        var v19 uint16 = 42
                        fmt.Println(v19)
                        _ = v19
                    } else {
                        fmt.Println(p0, fi8, p0)
                    }
                    v20 := GenS1{F0: 1, F1: 117, F2: 1}
                    _ = fi8
                    _ = v20
                }
            default:
                v21 := GenS1{F0: 32767, F1: 42, F2: 1}
                _ = v21
            }
            _ = fi6
        }
    } else {
        p0 -= p0
        var v22 uint64 = 42
        if false {
            v22 -= 183
            v23 := []int{-2147483648, (-(p0)), p0}
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 float64 = 1.5
                var v26 float64 = ((-2.25 * v25) + v25)
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := GenS1{F0: p0, F1: -663, F2: 42}
                    fmt.Println(v26)
                    switch -2147483648 {
                    case 32767:
                        var v29 float32 = 3.14159
                        _ = v29
                    case 0:
                        v22 = (^(v22))
                    case -219:
                        var v30 uint16 = 32767
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        } else {
            v31 := GenS2{F0: 1, F1: 4294967295, F2: 1}
            v32 := GenS3{F0: -1, F1: 255, F2: 9223372036854775807}
            _ = v31
            _ = v32
        }
        _ = v22
    }
    var v33 int = -974
    switch ("" + "\n\t") {
    case "\n\t":
        for fi34 := 0; fi34 < 8; fi34++ {
            fmt.Println(fi34, fi34, p0)
            var v35 uint16 = 32767
            fmt.Println(v33)
            _ = fi34
            _ = v35
        }
    case "a":
        v36 := GenS4{F0: 100, F1: 0, F2: -78}
        _ = v36
    default:
        var v37 int64 = -9223372036854775808
        _ = v37
    }
    v38 := GenS5{F0: 370, F1: 623, F2: (-(v33))}
    for fi39 := 0; fi39 < 8; fi39++ {
        var v40 int32 = 42
        p0 += 0
        v41 := []int{p0}
        _ = fi39
        _ = v40
        _ = v41
    }
    _ = p0
    _ = v33
    _ = v38
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "hello", "\n\t", "hello"}
    var v3 uint8 = 255
    switch ("" + ("a" + "foo bar")) {
    case "\n\t":
        v3 -= 15
    default:
        v42 := GenS5{F0: 65535, F1: 0, F2: -9223372036854775808}
        _ = v42
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        v44 := GenS6{F0: (v3 <= 176), F1: 2, F2: -971}
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 15; fi45++ {
        v2 = append(v2, ("hello" + (("\n\t" + "hello") + "")))
        for fi46 := 0; fi46 < 13; fi46++ {
            var v47 int = fi45
            var v48 uint = 255
            if (false || true) {
                var v49 int16 = -1
                _ = v49
            } else {
                fmt.Println(v3)
            }
            _ = fi46
            _ = v47
            _ = v48
        }
        v50 := GenS7{F0: 65535, F1: -464, F2: 65535}
        _ = fi45
        _ = v50
    }
    v51 := GenS8{F0: 456, F1: "foo bar", F2: v3}
    v3 -= v3
    v2 = append(v2, ((("" + "a") + "\n\t") + "a"))
    var v52 uint32 = 4294967295
    fmt.Println(v3)
    var v53 float64 = 1e10
    switch "hello" {
    case "hello":
        if true {
            fmt.Println(v3, v3)
            _ = len(v1)
        } else {
            fmt.Println(v3, v53, v52)
        }
    case "a":
        for fi54 := 0; fi54 < 17; fi54++ {
            fmt.Println(v52, v53, v52)
            _ = fi54
        }
    case "foo bar":
        v3 = (v3 | v3)
    default:
        if true {
            v53 += ((v53 * 0.0) + v53)
        } else {
            var v55 uint32 = v52
            _ = v55
        }
    }
    fmt.Println(v53)
    if (float32((v53 / (v53 / (v53 * v53)))) < 0.1) {
        for fi56 := 0; fi56 < 17; fi56++ {
            var v57 uint32 = v52
            if true {
                v52 -= v57
                if (!true) {
                    switch ("" + "foo bar") {
                    case "a":
                        v57 = 42
                    default:
                        v58 := GenS5{F0: 559, F1: -1, F2: fi56}
                        _ = v58
                    }
                }
                for fi59 := 0; fi59 < 17; fi59++ {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    v57 = (v57 | v57)
                    _ = fi59
                }
            } else {
                v60 := v1[""]
                v3 = (v3 & 128)
                _ = v60
            }
            _ = fi56
            _ = v57
        }
        v3 = 31
    } else {
        v61 := GenS2{F0: 9223372036854775807, F1: 0, F2: v3}
        _ = v61
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        v3 = (v3 >> ((18446744073709551615) & 63))
    case -735:
        switch "foo bar" {
        case "a":
            v2 = append(v2, "a")
        case "hello":
            var v62 uint16 = 0
            _ = v62
        case "":
            v2 = append(v2, (("hello" + ((("a" + "") + "hello") + "foo bar")) + ""))
        default:
            for fi63 := 0; fi63 < 19; fi63++ {
                fmt.Println(v52, v52, fi63)
                fmt.Println(v52)
                _ = fi63
            }
        }
    default:
        v64 := GenS9{F0: genFn1(1), F1: -104, F2: -17}
        _ = v64
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println(v53)
    for fi65 := 0; fi65 < 3; fi65++ {
        for fi66 := 0; fi66 < 8; fi66++ {
            switch fi65 {
            case 128:
                if false {
                    var v67 bool = (2147483647 >= -9223372036854775808)
                    var v68 int32 = 2147483647
                    v53 -= (v53 - (0.0 * (v53 * v53)))
                    _ = v67
                    _ = v68
                } else {
                    v53 += (v53 + float64(1e10))
                    fmt.Println(fi66)
                }
            default:
                v52 += (v52 % ((v52) | 1))
            }
            switch "\n\t" {
            case "hello":
                var v69 int8 = 1
                _ = v69
            default:
                v1["\n\t"] = 476
            }
            fmt.Println(v3, fi66, v3)
            _ = len(v1)
            var v70 uint = uint(v3)
            var v71 uint = uint(v52)
            _ = fi66
            _ = v70
            _ = v71
        }
        _ = fi65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
}
