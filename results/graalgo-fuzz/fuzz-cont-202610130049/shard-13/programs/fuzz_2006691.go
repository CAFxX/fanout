package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 int; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int8 = -128
    var v2 uint32 = 42
    var v3 float64 = (1e10 / float64(0.1))
    var v4 uint32 = v2
    var v5 int8 = (v1 + 121)
    v6 := []int{128, 127, 32768, -811}
    var v7 uint64 = 4294967295
    switch (("a" + "\n\t") + "foo bar") {
    case "\n\t":
        if true {
            if false {
                switch ("foo bar" + "hello") {
                case "hello":
                    v7 = v7
                case "\n\t":
                    v8 := GenS1{F0: 2}
                    _ = v8
                }
                v9 := GenS1{F0: (^(v4))}
                _ = v9
            } else {
                if true {
                    v2 = (v2 * (v2 >> ((18446744073709551615) & 63)))
                    v10 := GenS1{F0: 4294967295}
                    _ = v10
                }
                if true {
                    v4 += v2
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        if (!(!(-2147483648 <= 0))) {
                            var v12 int = 512
                            v2 += v2
                            v4 -= v4
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 18; fi13++ {
                                if false {
                                    v5 -= v1
                                    if (!false) {
                                        fmt.Println(v2)
                                        switch (^(fi11)) {
                                        case -2147483648:
                                            var v14 uint8 = uint8(v7)
                                            _ = v14
                                        case 4294967295:
                                            var v15 int64 = -1
                                            _ = v15
                                        }
                                        v7 -= (v7 % ((v7) | 1))
                                        var v16 uint32 = 589
                                        v1 = -1
                                        _ = v16
                                    }
                                    var v17 int32 = 1
                                    _ = v17
                                } else {
                                    v4 = v4
                                }
                                _ = fi13
                            }
                            v7 = 42
                        }
                        var v18 uint16 = 1
                        _ = fi11
                        _ = v18
                    }
                    fmt.Println(v4, v4, v2)
                    v2 = 578
                }
                fmt.Println(v7)
            }
            v6 = append(v6, 128)
        } else {
            var v19 string = "a"
            v2 -= v2
            _ = v19
        }
    case "foo bar":
        if (90 >= 146) {
            for fi20 := 0; fi20 < 8; fi20++ {
                _ = len(v6)
                if true {
                    v21 := GenS2{F0: -1, F1: (v7 | v7)}
                    switch "hello" {
                    case "hello":
                        var v22 int8 = (v1 | int8(v7))
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v5 = v5
                            _ = len(v6)
                            if false {
                                v24 := GenS2{F0: 933, F1: 2}
                                _ = v24
                            } else {
                                switch int(v2) {
                                case 4294967295:
                                    switch "" {
                                    case "":
                                        v25 := GenS3{F0: -1, F1: ("" != (("" + "") + "hello")), F2: fi23}
                                        _ = v25
                                    case "\n\t":
                                        v3 = 1.5
                                    }
                                case 32768:
                                    v26 := GenS3{F0: -9223372036854775808, F1: (true || true), F2: (-(fi23))}
                                    _ = v26
                                }
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v7, v2)
                                default:
                                    var v27 uint8 = 255
                                    _ = v27
                                }
                                if (!true) {
                                    if (v3 == ((((1e10 * float64(0.0)) / 1.5) * v3) / (v3 - v3))) {
                                        v6 = append(v6, int(v2))
                                    } else {
                                        if ((v7 & 4294967295) < uint64(v2)) {
                                            var v28 uint16 = 116
                                            _ = v28
                                        }
                                        if (0 == 331) {
                                            fmt.Println(v5)
                                        } else {
                                            var v29 uint = 0
                                            _ = v29
                                        }
                                    }
                                }
                                switch int(v1) {
                                case 9223372036854775807:
                                    v30 := GenS4{F0: (fi20 - fi20), F1: 2147483647, F2: 42}
                                    _ = v30
                                default:
                                    if false {
                                        if true {
                                            var v31 int = 2147483647
                                            _ = v31
                                        } else {
                                            v6 = append(v6, fi20)
                                            v32 := GenS1{F0: 2147483647}
                                            if false {
                                                var v33 string = "a"
                                                switch (-(fi23)) {
                                                case -9223372036854775808:
                                                    v4 += 42
                                                }
                                                var v34 uint16 = 228
                                                _ = v33
                                                _ = v34
                                            } else {
                                                var v35 uint64 = v7
                                                _ = v35
                                            }
                                            _ = v32
                                        }
                                        fmt.Println(v3)
                                        fmt.Println(v2, v3)
                                    } else {
                                        var v36 string = ("a" + ("" + "a"))
                                        _ = v36
                                    }
                                }
                                if false {
                                    v6 = append(v6, -477)
                                    var v37 int32 = -1
                                    _ = v37
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        switch (fi20 * 4294967295) {
                                        case -9223372036854775808:
                                            v38 := GenS1{F0: (v2 >> ((1) & 63))}
                                            _ = v38
                                        default:
                                            switch "\n\t" {
                                            case "":
                                                var v39 string = "a"
                                                _ = v39
                                            }
                                        }
                                    case "\n\t":
                                        fmt.Println(v1)
                                    default:
                                        v40 := GenS4{F0: (fi20 / ((fi23) | 1)), F1: 2147483647, F2: 127}
                                        _ = v40
                                    }
                                    var v41 int = fi20
                                    var v42 int32 = -2147483648
                                    var v43 int8 = (v5 * -128)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if (fi23 > fi23) {
                                        var v44 uint8 = 128
                                        _ = v44
                                    }
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            }
                            v6 = append(v6, fi23)
                            _ = fi23
                        }
                    }
                    if (false && false) {
                        var v45 float64 = (1e10 / 3.14159)
                        _ = v45
                    } else {
                        fmt.Println(v7, v1, fi20)
                        var v46 int16 = int16(v7)
                        _ = v46
                    }
                    _ = v21
                } else {
                    var v47 float64 = v3
                    v6 = append(v6, fi20)
                    if true {
                        var v48 int = fi20
                        v7 -= v7
                        if (false || (v48 > fi20)) {
                            var v49 uint32 = (^(v2))
                            _ = v49
                        }
                        _ = v48
                    }
                    _ = v47
                }
                _ = fi20
            }
            var v50 int16 = 1
            fmt.Println(v1, v4)
            _ = v50
        }
    default:
        var v51 string = ""
        _ = v51
    }
    fmt.Println(v5, v2)
    fmt.Println(v2, v1, v7)
    var v52 uint8 = uint8(v2)
    switch -22 {
    case 4294967295:
        var v53 float32 = 1e10
        _ = v53
    case 165:
        if (1 >= 1) {
            v54 := GenS5{F0: (v52 & v52)}
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
            _ = v54
        } else {
            v5 = v5
            v6 = append(v6, 0)
        }
    case 127:
        if ((3.14159 + 1e10) == 0.0) {
            switch 2147483647 {
            case -2147483648:
                if (v3 < (float64(float32(v3)) * (0.0 / v3))) {
                    if (1 > -32768) {
                        v55 := GenS2{F0: -476, F1: (v7 & v7)}
                        for fi56 := 0; fi56 < 10; fi56++ {
                            if true {
                                v57 := GenS5{F0: 61}
                                var v58 int64 = -392
                                _ = v57
                                _ = v58
                            }
                            v5 += v1
                            _ = fi56
                        }
                        _ = v55
                    }
                    for fi59 := 0; fi59 < 13; fi59++ {
                        v7 -= (v7 - v7)
                        _ = fi59
                    }
                    var v60 int8 = 100
                    _ = v60
                }
            default:
                v61 := GenS1{F0: v4}
                _ = v61
            }
        }
    }
    if ((true || (false || false)) && (42 < 42)) {
        if len(v6) > 0 {
            v6[0] = 979
        }
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v7
}
