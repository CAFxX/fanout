package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 bool; F1 string; F2 bool }
type GenS5 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int8; F1 int64; F2 uint64 }
func genFn0(p0 uint32, p1 int64) int16 {
    fmt.Println(p1, p1)
    v6 := []int{127, -2147483648, 32768, 127}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 uint16 = 197
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 uint16) string {
    for fi8 := 0; fi8 < 8; fi8++ {
        p1 = p0
        fmt.Println(fi8)
        v9 := GenS1{F0: genFn0(4294967295, -9223372036854775808), F1: 2147483647}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 2147483647, F1: 32768}
    fmt.Println(p1, p1)
    var v11 uint64 = 346
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return ("a" + "\n\t")
}

func main() {
    var v1 string = (("\n\t" + ("\n\t" + ("a" + "\n\t"))) + ("" + ("" + ("a" + ("hello" + (("\n\t" + "a") + "foo bar"))))))
    var v2 uint64 = 42
    v3 := make(map[int]bool)
    var v4 int64 = -9223372036854775808
    var v5 uint16 = 1
    if (v5 == (^(v5))) {
        v4 = v4
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v4)
            fmt.Println(v5, v1, v4)
            _ = fi12
        }
    }
    v13 := []int{0}
    v13 = append(v13, 42)
    switch (v1 + v1) {
    case "hello":
        fmt.Println(v2)
    }
    if ((float32(0.0) == 0.1) || (0.0 <= 0.1)) {
        switch ((v1 + v1) + "a") {
        case "":
            var v14 float32 = 1.5
            _ = v14
        default:
            switch "foo bar" {
            case "hello":
                if true {
                    var v15 float32 = 0.1
                    v1 = ("foo bar" + v1)
                    _ = v15
                }
            }
        }
        fmt.Println(v2)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if false {
            v3[(fi16 % ((fi16) | 1))] = true
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint8 = 127
                _ = fi17
                _ = v18
            }
            v4 += (v4 | v4)
        } else {
            var v19 float64 = (1e10 * 1.5)
            delete(v3, (fi16 << ((1) & 63)))
            _ = v19
        }
        switch "a" {
        case "foo bar":
            v20 := GenS2{F0: -648, F1: (fi16 << ((963) & 63))}
            _ = v20
        case "\n\t":
            if (!false) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    switch fi16 {
                    case -684:
                        var v22 bool = (0.1 == 0.1)
                        _ = v22
                    case -1:
                        if len(v13) > 0 {
                            v13[0] = (fi21 % ((-802) | 1))
                        }
                    default:
                        var v23 int32 = 0
                        v5 = 42
                        var v24 uint16 = 1
                        _ = v23
                        _ = v24
                    }
                    switch "" {
                    case "foo bar":
                        fmt.Println(v4, v4)
                        v13 = append(v13, -9223372036854775808)
                    case "hello":
                        delete(v3, 4294967295)
                    case "\n\t":
                        var v25 string = (v1 + v1)
                        _ = v25
                    default:
                        if (65535 <= 924) {
                            if (uint(v2) < 435) {
                                if false {
                                    v13 = append(v13, 4294967295)
                                    fmt.Println(v5, v4)
                                }
                                var v26 int16 = -688
                                if (true && false) {
                                    var v27 uint8 = 255
                                    _ = v27
                                } else {
                                    var v28 uint = 924
                                    fmt.Println(v26)
                                    if false {
                                        v26 = (v26 + v26)
                                        fmt.Println(fi21)
                                    } else {
                                        if (v28 == (v28 / ((244) | 1))) {
                                            var v29 int32 = 2147483647
                                            _ = v29
                                        }
                                        v26 += (-(v26))
                                        v30 := GenS3{F0: 42, F1: v1}
                                        var v31 uint8 = 2
                                        v32 := GenS2{F0: 2147483647, F1: (fi16 * (fi21 << ((v28) & 63)))}
                                        var v33 string = v1
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v28
                                }
                                v4 = -936
                                v2 -= 42
                                _ = v26
                            }
                        } else {
                            if false {
                                var v34 string = v1
                                fmt.Println(v2, v1)
                                v4 += v4
                                _ = v34
                            } else {
                                fmt.Println(fi16)
                            }
                            v1 = "foo bar"
                        }
                    }
                    v35 := GenS4{F0: (!false), F1: "foo bar", F2: (249 < 1)}
                    _ = fi21
                    _ = v35
                }
                fmt.Println(v5, v1, v1)
                var v36 int32 = 880
                _ = v36
            }
        case "hello":
            v37 := v3[fi16]
            _ = v37
        default:
            switch (v1 + "\n\t") {
            case "\n\t":
                for fi38 := 0; fi38 < 12; fi38++ {
                    v1 = v1
                    _ = fi38
                }
            case "hello":
                v39 := GenS5{F0: 93, F1: 0, F2: 2}
                _ = v39
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    v40 := GenS6{F0: -1, F1: v4, F2: (v2 | (v2 + 0))}
                    _ = v40
                }
            default:
                var v41 float32 = (float32(3.14159) - (-2.25 - (3.14159 / 1e10)))
                _ = v41
            }
        }
        _ = fi16
    }
    var v42 bool = ("" < v1)
    if len(v13) > 0 {
        v13[0] = -534
    }
    switch 579 {
    case -132:
        v1 = (v1 + "\n\t")
    }
    var v43 int = 0
    v2 = v2
    var v44 string = (v1 + "a")
    v5 -= (^(v5))
    for fi45 := 0; fi45 < 14; fi45++ {
        v46 := GenS2{F0: 548, F1: (fi45 % (((v43 << ((698) & 63))) | 1))}
        _ = fi45
        _ = v46
    }
    var v47 bool = (false || v42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v5
}
