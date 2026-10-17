package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int32; F1 uint }

func main() {
    var v1 string = "a"
    v2 := []string{"foo bar", v1}
    v3 := make(map[int]uint32)
    v1 = ("hello" + v1)
    if false {
        v1 = ""
    } else {
        switch 32768 {
        case 42:
            if false {
                v1 = "\n\t"
                fmt.Println(v1)
            }
        case 128:
            var v4 float32 = (0.0 + 1.5)
            _ = v4
        }
    }
    var v5 float32 = 0.0
    if false {
        fmt.Println(v1, v5)
        if false {
            v1 = v1
            var v6 uint8 = 128
            _ = v6
        } else {
            switch 4294967295 {
            case 539:
                switch v1 {
                case "hello":
                    v7 := GenS1{F0: (((("a" + (v1 + v1)) + ("" + (v1 + (((("hello" + "a") + v1) + (v1 + v1)) + v1)))) + "a") + v1), F1: 400, F2: 18446744073709551615}
                    _ = v7
                }
            case -1:
                switch "a" {
                case "hello":
                    if false {
                        fmt.Println(v5, v1)
                    }
                default:
                    for fi8 := 0; fi8 < 14; fi8++ {
                        v5 = v5
                        _ = fi8
                    }
                }
            }
            fmt.Println(v5)
            v5 = 0.0
        }
    }
    switch -228 {
    case 127:
        fmt.Println(v1, v1)
    }
    v3[-2147483648] = 65535
    var v9 float32 = float32(float64(-2.25))
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                _ = len(v2)
                v9 = (0.1 / 1.0)
            }
            var v12 uint8 = 2
            _ = fi11
            _ = v12
        }
        var v13 int32 = -76
        v14 := v3[(fi10 % (((fi10 | (fi10 & (fi10 + 127)))) | 1))]
        _ = fi10
        _ = v13
        _ = v14
    }
    if true {
        v5 += v5
        if (1 != 319) {
            v15 := GenS1{F0: v1, F1: 716, F2: 0}
            var v16 int64 = 0
            v5 = v9
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: v1, F1: 2, F2: 1}
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int = fi18
            if false {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v2 = append(v2, ((v1 + v1) + (v1 + ("foo bar" + v1))))
                    if ((65535 > 67) || false) {
                        v2 = append(v2, "hello")
                        if false {
                            v5 -= (v9 + v9)
                            fmt.Println(fi20, fi20)
                        } else {
                            v19 = fi18
                        }
                        v1 = ("a" + "hello")
                    }
                    _ = fi20
                }
                v9 -= float32(((((float64(v5) - float64(((float32(-2.25) / (((0.1 / 1e10) + v9) / v9)) - v9))) * 1.5) / 3.14159) + ((3.14159 * 0.1) * 0.0)))
                var v21 uint16 = 65535
                _ = v21
            } else {
                var v22 int64 = 42
                switch v1 {
                case "":
                    var v23 int8 = 1
                    _ = v23
                }
                v19 = 354
                _ = v22
            }
            v5 -= v5
            _ = fi18
            _ = v19
        }
    }
    if (354 < 4294967295) {
        if (((((v9 / 1e10) + 3.14159) == (v9 / 1.0)) && true) || (v9 > (v5 / -2.25))) {
            switch v1 {
            case "":
                v24 := GenS1{F0: "a", F1: 65535, F2: 42}
                _ = v24
            case "\n\t":
                v25 := GenS1{F0: (v1 + v1), F1: 255, F2: 949}
                _ = v25
            case "hello":
                v5 -= (float32(3.14159) / 3.14159)
            default:
                v26 := GenS1{F0: v1, F1: 1, F2: 42}
                _ = v26
            }
            if ((v1 + ("hello" + v1)) == v1) {
                if (0.0 > float64(((v5 + v9) + v9))) {
                    if false {
                        var v27 float64 = 1e10
                        _ = v27
                    }
                    if false {
                        if (!true) {
                            var v28 int64 = 213
                            var v29 uint32 = 245
                            v29 = v29
                            _ = v28
                            _ = v29
                        }
                    }
                    v5 = ((float32((0.0 + 1e10)) - v9) + v5)
                }
                v30 := GenS1{F0: v1, F1: 42, F2: 18446744073709551615}
                v9 -= 1e10
                _ = v30
            } else {
                if false {
                    if (91 != 42) {
                        v31 := GenS2{F0: v1, F1: 2}
                        _ = v31
                    }
                    v1 = ((v1 + v1) + "foo bar")
                    for fi32 := 0; fi32 < 19; fi32++ {
                        var v33 uint = 4294967295
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi32
                        _ = v33
                    }
                } else {
                    if false {
                        var v34 int32 = -1
                        if (v9 >= (float32(-2.25) + v5)) {
                            v9 -= 1.5
                            for fi35 := 0; fi35 < 7; fi35++ {
                                switch (("\n\t" + (v1 + (("\n\t" + ("a" + ((v1 + v1) + v1))) + (v1 + v1)))) + "a") {
                                case "a":
                                    for fi36 := 0; fi36 < 9; fi36++ {
                                        var v37 int = fi35
                                        fmt.Println(fi36, fi36, v5)
                                        _ = fi36
                                        _ = v37
                                    }
                                }
                                _ = fi35
                            }
                        } else {
                            for fi38 := 0; fi38 < 7; fi38++ {
                                v5 = v9
                                _ = fi38
                            }
                        }
                        _ = len(v2)
                        _ = v34
                    }
                    v9 = v9
                    switch 1 {
                    case 9223372036854775807:
                        if ((!true) && false) {
                            switch -566 {
                            case 9223372036854775807:
                                if (v1 != ((v1 + v1) + "")) {
                                    if false {
                                        for fi39 := 0; fi39 < 2; fi39++ {
                                            v1 = ("" + v1)
                                            v9 = 3.14159
                                            var v40 uint32 = 962
                                            _ = fi39
                                            _ = v40
                                        }
                                        switch 1 {
                                        case 127:
                                            delete(v3, 845)
                                        case 128:
                                            var v41 int8 = 100
                                            _ = v41
                                        }
                                        v5 = v9
                                    } else {
                                        v9 += float32(3.14159)
                                    }
                                    var v42 string = v1
                                    _ = len(v3)
                                    _ = v42
                                } else {
                                    v1 = v1
                                }
                            case 42:
                                v1 = "hello"
                            case 16:
                                switch -9223372036854775808 {
                                case -726:
                                    v5 -= v5
                                default:
                                    v43 := GenS3{F0: (!false)}
                                    _ = v43
                                }
                            }
                        } else {
                            if ((138 > -410) || false) {
                                v44 := v3[32768]
                                _ = v44
                            }
                            v5 = v9
                            delete(v3, -187)
                        }
                    }
                }
            }
            fmt.Println(v9, v9, v5)
        } else {
            v45 := GenS4{F0: 0, F1: false}
            fmt.Println(v5)
            _ = v45
        }
    }
    fmt.Println(v5, v9, v5)
    switch "foo bar" {
    case "\n\t":
        var v46 bool = true
        _ = v46
    case "":
        switch -9223372036854775808 {
        case -680:
            fmt.Println(v1, v1, v9)
        case 34:
            v47 := GenS2{F0: v1, F1: 18446744073709551615}
            _ = v47
        case -83:
            for fi48 := 0; fi48 < 9; fi48++ {
                v49 := GenS5{F0: -507, F1: 896}
                v9 = v9
                _ = fi48
                _ = v49
            }
        }
    case "hello":
        v50 := GenS1{F0: v1, F1: 808, F2: 4294967295}
        _ = v50
    }
    var v51 uint8 = 255
    v52 := GenS4{F0: 916, F1: ((true && (42 <= 2147483647)) && (127 != 4294967295))}
    var v53 uint8 = v51
    var v54 bool = (false && false)
    var v55 int64 = -660
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v9
}
