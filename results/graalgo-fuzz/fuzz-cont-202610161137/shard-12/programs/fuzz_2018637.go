package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int; F2 string }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint }
type GenS7 struct { F0 int }
func genFn0(p0 int32) uint {
    switch 831 {
    case 0:
        var v7 int64 = -1
        _ = v7
    case 2147483647:
        var v8 int32 = (-(p0))
        _ = v8
    default:
        v9 := GenS1{F0: 154}
        _ = v9
    }
    if false {
        v10 := []int{1, -9223372036854775808}
        switch -365 {
        case 32768:
            if (0.1 >= 0.0) {
                v11 := GenS1{F0: 2147483647}
                _ = v11
            } else {
                p0 -= p0
            }
        case 9223372036854775807:
            v10 = append(v10, 1)
        case -2147483648:
            var v12 string = (("foo bar" + ("hello" + "")) + "a")
            _ = v12
        default:
            fmt.Println(p0)
        }
        _ = v10
    } else {
        p0 += p0
    }
    var v13 int16 = -54
    if false {
        v14 := GenS2{F0: 437}
        fmt.Println(p0, v13)
        v13 += 930
        _ = v14
    }
    _ = p0
    _ = v13
    return 255
}
func genFn1(p0 int16) uint64 {
    var v15 float64 = 1.5
    if true {
        if false {
            v16 := []int{128, 128}
            if (-1 <= 42) {
                v16 = append(v16, 0)
                v15 += float64(0.0)
                var v17 uint16 = 1
                _ = v17
            } else {
                switch "foo bar" {
                case "hello":
                    switch 815 {
                    case 2147483647:
                        fmt.Println(v15, p0)
                    case -2147483648:
                        switch 0 {
                        case -695:
                            if true {
                                if len(v16) > 0 {
                                    v16[0] = -337
                                }
                            }
                        case 10:
                            var v18 bool = (!false)
                            _ = v18
                        case 542:
                            v19 := GenS1{F0: 263}
                            _ = v19
                        default:
                            v16 = append(v16, 127)
                        }
                    }
                case "a":
                    var v20 string = ("" + "foo bar")
                    _ = v20
                case "\n\t":
                    v16 = append(v16, 32768)
                }
                v16 = append(v16, 93)
                var v21 float32 = 1.5
                _ = v21
            }
            var v22 float64 = ((v15 - (3.14159 * -2.25)) - v15)
            _ = v16
            _ = v22
        }
    }
    if (-2147483648 == -2147483648) {
        v23 := GenS2{F0: 609}
        _ = v23
    }
    switch 1 {
    case 409:
        p0 = p0
    }
    _ = p0
    _ = v15
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]int)
    var v3 int32 = 42
    v4 := []string{"hello", ""}
    var v5 string = ("a" + ("a" + "\n\t"))
    var v6 int16 = -20
    var v24 int8 = 39
    for fi25 := 0; fi25 < 11; fi25++ {
        if (!true) {
            v1 += 1
            if (true || true) {
                var v26 string = "\n\t"
                switch v26 {
                case "hello":
                    var v27 int64 = v1
                    _ = v27
                case "":
                    var v28 uint8 = 1
                    _ = v28
                }
                _ = v26
            }
            switch "foo bar" {
            case "foo bar":
                var v29 int64 = v1
                _ = v29
            case "":
                v3 = (-(v3))
            case "\n\t":
                if ((v1 & (v1 ^ v1)) >= 1) {
                    v30 := GenS3{F0: v6}
                    switch "a" {
                    case "foo bar":
                        if (!(!true)) {
                            v6 += v6
                        } else {
                            if true {
                                v31 := GenS3{F0: -1}
                                _ = v31
                            }
                        }
                    default:
                        _ = len(v4)
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v30
                } else {
                    v24 -= v24
                }
            default:
                v32 := GenS3{F0: int16(fi25)}
                _ = v32
            }
        }
        v33 := GenS4{F0: (v3 << ((genFn0((-(v3)))) & 63))}
        _ = fi25
        _ = v33
    }
    if (!(9223372036854775807 == (v1 | -9223372036854775808))) {
        var v34 int64 = v1
        var v35 int64 = (v1 % ((1) | 1))
        for fi36 := 0; fi36 < 12; fi36++ {
            v2[fi36] = (fi36 * fi36)
            v37 := v2[fi36]
            var v38 uint16 = 179
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = v34
        _ = v35
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        if false {
            if (42 <= 2) {
                var v40 uint64 = 2
                v41 := GenS5{F0: v6, F1: fi39, F2: "foo bar"}
                _ = v40
                _ = v41
            }
            v42 := GenS6{F0: 1, F1: 42, F2: 1}
            _ = v42
        } else {
            if false {
                if false {
                    _ = len(v4)
                    var v43 int = fi39
                    v24 = v24
                    _ = v43
                }
            } else {
                for fi44 := 0; fi44 < 5; fi44++ {
                    if (!true) {
                        switch v5 {
                        case "a":
                            v6 = v6
                        case "hello":
                            v45 := GenS7{F0: (fi39 | (fi39 % (((fi39 << ((4294967295) & 63))) | 1)))}
                            _ = v45
                        default:
                            v24 = v24
                        }
                    } else {
                        var v46 float32 = -2.25
                        v47 := v2[(fi44 | 4294967295)]
                        _ = v46
                        _ = v47
                    }
                    var v48 float64 = (3.14159 + 1e10)
                    v49 := GenS1{F0: 875}
                    _ = fi44
                    _ = v48
                    _ = v49
                }
            }
            for fi50 := 0; fi50 < 10; fi50++ {
                var v51 int8 = (v24 / (((^(v24))) | 1))
                fmt.Println(v24, v6)
                fmt.Println(v24, fi50)
                _ = fi50
                _ = v51
            }
            v24 += v24
        }
        var v52 string = (v5 + "\n\t")
        _ = fi39
        _ = v52
    }
    v6 -= v6
    v53 := v2[0]
    fmt.Println(v5, v24)
    var v54 uint32 = 97
    var v55 uint64 = 542
    var v56 int16 = 545
    var v57 int16 = v56
    v3 = (v3 ^ (-(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
