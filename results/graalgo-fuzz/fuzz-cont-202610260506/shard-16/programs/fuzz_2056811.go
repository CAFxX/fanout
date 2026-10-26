package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint) int16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int = fi6
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: (-(p0)), F1: 1, F2: 255}
            p1 = 841
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int64 = 42
            v7 = 9223372036854775807
            _ = fi10
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    p1 = 4294967295
    _ = p0
    _ = p1
    return (^(p0))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 uint8 = 2
    var v4 uint64 = 278
    v5 := make(map[int]int)
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            v4 -= (v4 & (^(v4)))
            _ = fi13
        }
        _ = fi12
    }
    v14 := []int{-9223372036854775808, -2147483648, 2147483647, 2147483647}
    switch "a" {
    case "foo bar":
        fmt.Println(v4, v3)
    case "\n\t":
        if (((1e10 * -2.25) - (0.0 + ((0.1 / (3.14159 - (0.0 * -2.25))) - float32(3.14159)))) > float32(0.0)) {
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS1{F0: genFn0(-902, 255), F1: 9223372036854775807, F2: uint16(v3)}
                for fi17 := 0; fi17 < 9; fi17++ {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    switch (fi15 * (fi17 - 2147483647)) {
                    case 128:
                        fmt.Println(fi17, fi17)
                    case 797:
                        v4 -= (v4 | 2)
                    case -9223372036854775808:
                        if true {
                            var v18 uint32 = 42
                            _ = len(v14)
                            fmt.Println(fi15, v18)
                            _ = len(v14)
                            _ = v18
                        }
                    default:
                        v4 = (v4 % ((v4) | 1))
                    }
                    _ = len(v14)
                    if len(v14) > 0 {
                        v14[0] = fi15
                    }
                    _ = fi17
                }
                _ = fi15
                _ = v16
            }
        } else {
            v19 := GenS1{F0: genFn0(genFn0(32767, 4294967295), 2), F1: 42, F2: 65535}
            if ("hello" < (("a" + "foo bar") + "hello")) {
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 uint = 468
                    v4 += (v4 | (v4 + v4))
                    v21 = (v21 % (((^(v21))) | 1))
                    _ = fi20
                    _ = v21
                }
                switch 127 {
                case 128:
                    switch "" {
                    case "":
                        v4 = 2
                    case "\n\t":
                        var v22 bool = (false || true)
                        _ = v22
                    default:
                        if len(v14) > 0 {
                            v14[0] = 32767
                        }
                    }
                default:
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                var v23 int8 = 127
                _ = v23
            } else {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = v19
        }
    case "":
        var v24 float32 = -2.25
        _ = v24
    default:
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS2{F0: 83, F1: 263, F2: 428}
            for fi27 := 0; fi27 < 10; fi27++ {
                if false {
                    v28 := GenS2{F0: 551, F1: -2147483648, F2: 219}
                    v5[127] = fi27
                    switch ("foo bar" + "") {
                    case "a":
                        fmt.Println(v3)
                    case "hello":
                        v29 := GenS3{F0: -547, F1: 388}
                        v4 -= v4
                        _ = v29
                    }
                    _ = v28
                }
                delete(v2, (-(fi27)))
                switch (fi27 | fi27) {
                case 9223372036854775807:
                    v4 += 906
                case -666:
                    if (!true) {
                        var v30 string = ("a" + "a")
                        v31 := GenS2{F0: 32767, F1: 2147483647, F2: 0}
                        fmt.Println(fi27)
                        _ = v30
                        _ = v31
                    } else {
                        delete(v1, (("foo bar" + "hello") + ""))
                        var v32 int64 = -580
                        _ = v32
                    }
                case -2147483648:
                    switch (fi27 << ((978) & 63)) {
                    case 4294967295:
                        switch -1 {
                        case -1:
                            var v33 uint32 = 42
                            _ = v33
                        case 127:
                            v4 = v4
                        case 2147483647:
                            v34 := GenS1{F0: -1, F1: 104, F2: 0}
                            _ = v34
                        }
                    default:
                        switch "foo bar" {
                        case "":
                            var v35 string = ""
                            _ = v35
                        case "hello":
                            if true {
                                v3 = 128
                                var v36 int16 = -1
                                _ = v36
                            }
                        case "foo bar":
                            v4 -= (^(v4))
                        }
                    }
                    if (1e10 == (0.1 - 0.1)) {
                        fmt.Println(v4)
                        fmt.Println(v3)
                        var v37 int16 = -175
                        _ = v37
                    }
                }
                _ = fi27
            }
            fmt.Println(fi25)
            _ = fi25
            _ = v26
        }
    }
    if false {
        var v38 int = 632
        _ = v38
    }
    v4 -= 1
    for fi39 := 0; fi39 < 11; fi39++ {
        v4 = 2
        _ = fi39
    }
    v3 += 1
    v40 := GenS4{F0: "a", F1: -62}
    switch 0 {
    case 2147483647:
        if true {
            v4 += (v4 / ((v4) | 1))
        }
    case 32767:
        fmt.Println(v3)
    default:
        v4 = (v4 / (((v4 >> ((uint(v3)) & 63))) | 1))
    }
    v1[(("\n\t" + ("a" + "a")) + (("" + "\n\t") + (("\n\t" + "hello") + ("" + "a"))))] = "\n\t"
    v4 += 4294967295
    _ = len(v14)
    v41 := GenS5{F0: 0}
    v42 := GenS2{F0: 255, F1: 42, F2: 255}
    var v43 uint = 145
    v44 := GenS3{F0: -2147483648, F1: 4294967295}
    if (((1.5 / float64(3.14159)) / 1e10) > (float64(3.14159) + 1e10)) {
        var v45 uint32 = 42
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
