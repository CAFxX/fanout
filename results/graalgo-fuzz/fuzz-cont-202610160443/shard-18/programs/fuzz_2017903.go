package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int; F1 int8; F2 int32 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint64; F1 uint; F2 uint8 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{fi7}
        if (0 < -185) {
            _ = len(v8)
            var v9 float32 = 3.14159
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int16 = 787
                var v12 uint64 = 18446744073709551615
                v11 = (v11 + v11)
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            _ = len(v8)
            var v13 int8 = -128
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    var v14 uint16 = 1
    var v15 uint8 = 42
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = -2.25
    var v3 bool = true
    v4 := []int{9223372036854775807, 128, 2147483647}
    v5 := make(map[int]bool)
    var v6 int8 = -1
    if v3 {
        fmt.Println(v3, v6, v6)
        if len(v4) > 0 {
            v4[0] = 130
        }
    }
    fmt.Println(v2, v3)
    v2 = 1e10
    if v3 {
        v16 := GenS1{F0: 42}
        var v17 int8 = v6
        v18 := GenS2{F0: "\n\t", F1: -152}
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint64 = 4294967295
    v1[0] = ("foo bar" + (("\n\t" + ("foo bar" + ("\n\t" + (("foo bar" + "hello") + ((((("hello" + "foo bar") + "a") + ("hello" + "foo bar")) + "hello") + "hello"))))) + "\n\t"))
    var v20 uint = 42
    var v21 int16 = 1000
    if (!false) {
        for fi22 := 0; fi22 < 3; fi22++ {
            if (true && v3) {
                switch ("" + ("hello" + ("" + (("" + (("\n\t" + "a") + ("" + "\n\t"))) + ("a" + "a"))))) {
                case "a":
                    v19 = v19
                default:
                    fmt.Println(v21)
                }
            }
            _ = fi22
        }
        switch "" {
        case "foo bar":
            var v23 int64 = 9223372036854775807
            _ = v23
        case "":
            var v24 float64 = v2
            _ = v24
        default:
            if v3 {
                v6 = (v6 ^ (v6 * -116))
            } else {
                if ("foo bar" != ("a" + "foo bar")) {
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            if (-2147483648 > 0) {
                                var v25 float32 = ((0.0 + 3.14159) * 0.0)
                                fmt.Println(v20, v21)
                                _ = v25
                            }
                        case "foo bar":
                            switch 127 {
                            case 42:
                                var v26 string = ("hello" + "")
                                _ = v26
                            case 1:
                                v27 := GenS2{F0: "\n\t", F1: -979}
                                _ = v27
                            case -275:
                                v21 -= v21
                            default:
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    switch fi28 {
                                    case -390:
                                        v29 := GenS3{F0: -32768, F1: 255}
                                        _ = v29
                                    case 128:
                                        if false {
                                            v30 := GenS2{F0: "foo bar", F1: (fi28 & fi28)}
                                            _ = v30
                                        }
                                    case 2147483647:
                                        var v31 float64 = (((float64(((3.14159 + ((1.5 - 1.5) * -2.25)) * float32(0.1))) + -2.25) - v2) * 0.0)
                                        _ = v31
                                    }
                                    _ = fi28
                                }
                            }
                        }
                    case "foo bar":
                        for fi32 := 0; fi32 < 12; fi32++ {
                            var v33 uint16 = 1
                            delete(v1, 2147483648)
                            if false {
                                switch fi32 {
                                case -851:
                                    v20 -= 403
                                case 0:
                                    v19 = v19
                                case -416:
                                    fmt.Println(v20)
                                default:
                                    v19 = 1
                                }
                                switch "foo bar" {
                                case "a":
                                    _ = len(v5)
                                case "\n\t":
                                    v34 := GenS4{F0: (fi32 << (((v20 / ((v20) | 1))) & 63)), F1: 0, F2: 42}
                                    _ = v34
                                case "foo bar":
                                    var v35 bool = v3
                                    _ = v35
                                }
                                fmt.Println(v2, v20, v6)
                            } else {
                                v36 := GenS5{F0: "", F1: "\n\t"}
                                if (1 >= 42) {
                                    var v37 uint8 = 2
                                    v38 := GenS5{F0: "\n\t", F1: ("\n\t" + ("" + "foo bar"))}
                                    _ = v37
                                    _ = v38
                                }
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v40 uint64 = (^(v19))
                                        _ = v40
                                    case "foo bar":
                                        switch fi32 {
                                        case -928:
                                            var v41 uint8 = 42
                                            _ = v41
                                        case 2147483647:
                                            v21 = -1
                                        case -469:
                                            switch "a" {
                                            case "a":
                                                if false {
                                                    v42 := GenS3{F0: v21, F1: 134}
                                                    v43 := GenS2{F0: "foo bar", F1: (fi39 * fi32)}
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    v1[1] = ""
                                                }
                                            case "":
                                                v5[fi32] = (v3 || (genFn0() < -221))
                                            case "foo bar":
                                                var v44 float32 = ((float32(v2) / 1.5) * 0.1)
                                                _ = v44
                                            }
                                            v3 = v3
                                            v20 += v20
                                        default:
                                            fmt.Println(v2)
                                        }
                                    default:
                                        v2 = (v2 + -2.25)
                                    }
                                    _ = fi39
                                }
                                _ = v36
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                }
                if v3 {
                    if false {
                        var v45 uint16 = 372
                        _ = v45
                    } else {
                        fmt.Println(v2, v20, v3)
                        for fi46 := 0; fi46 < 7; fi46++ {
                            var v47 uint16 = 2
                            fmt.Println(v47, v6, v6)
                            switch "" {
                            case "\n\t":
                                switch (fi46 * (^(fi46))) {
                                case 4294967295:
                                    var v48 uint = v20
                                    _ = v48
                                case -1:
                                    fmt.Println(v19, v2, v21)
                                case 32767:
                                    for fi49 := 0; fi49 < 15; fi49++ {
                                        var v50 int = -2147483648
                                        fmt.Println(v6)
                                        _ = fi49
                                        _ = v50
                                    }
                                }
                            case "":
                                if (630 != 65535) {
                                    var v51 int32 = -2147483648
                                    _ = v51
                                }
                            case "hello":
                                v5[fi46] = v3
                            default:
                                v5[fi46] = true
                            }
                            _ = fi46
                            _ = v47
                        }
                    }
                    var v52 int8 = v6
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v19, v20)
                    default:
                        v53 := GenS1{F0: int(v21)}
                        _ = v53
                    }
                    _ = v52
                }
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    v4 = append(v4, -326)
    for fi54 := 0; fi54 < 15; fi54++ {
        var v55 uint16 = 0
        if (!(248 <= 128)) {
            var v56 float32 = 0.0
            switch (-(fi54)) {
            case 32767:
                fmt.Println(fi54, v56, v2)
            case 128:
                v57 := GenS1{F0: fi54}
                _ = v57
            case -1:
                switch (fi54 - fi54) {
                case 735:
                    v3 = (((v19 & 42) <= 420) || v3)
                default:
                    v19 = (v19 - v19)
                }
            default:
                switch 408 {
                case -9223372036854775808:
                    var v58 float64 = (1.5 + (v2 * ((v2 * v2) * 3.14159)))
                    _ = v58
                case 349:
                    v59 := GenS3{F0: 569, F1: 42}
                    _ = v59
                case 4294967295:
                    v55 -= 0
                }
            }
            _ = v56
        }
        v60 := GenS5{F0: "a", F1: "foo bar"}
        _ = fi54
        _ = v55
        _ = v60
    }
    if false {
        var v61 uint32 = 65535
        v19 -= v19
        _ = v61
    }
    v62 := GenS3{F0: v21, F1: 2}
    v63 := GenS2{F0: "foo bar", F1: -9223372036854775808}
    v64 := GenS6{F0: v19, F1: (v20 - 4294967295), F2: 223}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v64
}
