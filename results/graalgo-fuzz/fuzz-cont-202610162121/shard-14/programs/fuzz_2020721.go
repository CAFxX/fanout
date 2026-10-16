package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint; F1 bool; F2 uint64 }

func main() {
    v1 := []string{("a" + "\n\t"), ("foo bar" + "\n\t"), "foo bar", "\n\t"}
    v2 := make(map[string]uint32)
    var v3 float32 = (float32(0.0) + 1e10)
    var v4 int32 = -2147483648
    v5 := make(map[uint32]int)
    var v6 uint64 = 939
    var v7 uint = 255
    switch 32767 {
    case 983:
        v5[65535] = 233
    case 42:
        switch "a" {
        case "hello":
            var v8 int32 = 2147483647
            _ = v8
        case "\n\t":
            fmt.Println(v4, v7)
        case "a":
            var v9 string = "hello"
            _ = v9
        }
    case -9223372036854775808:
        for fi10 := 0; fi10 < 10; fi10++ {
            v1 = append(v1, (("\n\t" + ("" + "\n\t")) + "hello"))
            _ = fi10
        }
    }
    if false {
        v11 := GenS1{F0: (v7 ^ 273)}
        _ = v11
    }
    var v12 string = "\n\t"
    if (v4 < (v4 << (((v7 | v7)) & 63))) {
        var v13 uint = v7
        _ = v13
    }
    if (100 < 20) {
        delete(v2, v12)
        var v14 float64 = 0.1
        delete(v5, 650)
        _ = v14
    }
    var v15 int8 = 1
    if true {
        v16 := GenS2{F0: v12, F1: v4, F2: (v4 & 2147483647)}
        fmt.Println(v6)
        _ = v16
    } else {
        var v17 uint64 = 126
        switch 993 {
        case 42:
            if (0.0 != -2.25) {
                if false {
                    if (9223372036854775807 <= 42) {
                        fmt.Println(v3)
                    } else {
                        v18 := GenS1{F0: (v7 >> ((65535) & 63))}
                        for fi19 := 0; fi19 < 1; fi19++ {
                            fmt.Println(fi19, v6)
                            _ = fi19
                        }
                        _ = v18
                    }
                    if true {
                        fmt.Println(v17)
                        switch "" {
                        case "foo bar":
                            for fi20 := 0; fi20 < 6; fi20++ {
                                v12 = v12
                                v12 = v12
                                _ = len(v1)
                                _ = fi20
                            }
                        case "a":
                            for fi21 := 0; fi21 < 6; fi21++ {
                                v7 = v7
                                _ = fi21
                            }
                        default:
                            v22 := GenS3{F0: 42, F1: (^(v7))}
                            _ = v22
                        }
                        v17 = 427
                    } else {
                        fmt.Println(v17, v6)
                        if (9223372036854775807 > -510) {
                            var v23 string = ""
                            fmt.Println(v7)
                            _ = v23
                        } else {
                            v15 -= v15
                            v15 -= v15
                        }
                    }
                    switch "a" {
                    case "hello":
                        v3 -= float32(0.1)
                    case "\n\t":
                        for fi24 := 0; fi24 < 20; fi24++ {
                            var v25 uint64 = 4294967295
                            _ = len(v2)
                            _ = fi24
                            _ = v25
                        }
                    default:
                        v26 := GenS2{F0: (((v12 + v12) + (v12 + (v12 + (v12 + v12)))) + "hello"), F1: (v4 >> ((4294967295) & 63)), F2: (v4 & v4)}
                        _ = v26
                    }
                }
                var v27 int32 = v4
                _ = v27
            } else {
                fmt.Println(v15)
            }
        case 9223372036854775807:
            v28 := GenS4{F0: uint8(v6), F1: v12}
            _ = v28
        default:
            fmt.Println(v6, v6, v15)
        }
        if false {
            v17 += 4294967295
        } else {
            var v29 uint = v7
            v7 = v7
            _ = v29
        }
        _ = v17
    }
    if true {
        if false {
            v15 -= (v15 - v15)
        }
    }
    _ = len(v1)
    if false {
        v3 = v3
        v12 = "a"
        switch "" {
        case "hello":
            v4 -= v4
        case "\n\t":
            if (false && false) {
                v30 := GenS4{F0: 120, F1: v12}
                v31 := GenS5{F0: (v7 >> (((v7 >> (((v7 & 842)) & 63))) & 63)), F1: true, F2: (v6 >> (((^(v7))) & 63))}
                _ = v30
                _ = v31
            } else {
                switch (v12 + v12) {
                case "a":
                    v2[(("hello" + v12) + "")] = 65535
                case "\n\t":
                    switch (("a" + v12) + v12) {
                    case "":
                        v32 := GenS1{F0: (v7 >> ((v7) & 63))}
                        _ = v32
                    case "foo bar":
                        var v33 bool = (1000 >= 0)
                        _ = v33
                    default:
                        if true {
                            for fi34 := 0; fi34 < 9; fi34++ {
                                if ((309 > 894) || true) {
                                    for fi35 := 0; fi35 < 6; fi35++ {
                                        v7 = v7
                                        _ = fi35
                                    }
                                    v7 -= 1
                                    var v36 int16 = 42
                                    _ = v36
                                } else {
                                    var v37 uint32 = 148
                                    _ = v37
                                }
                                switch (v12 + v12) {
                                case "\n\t":
                                    _ = len(v1)
                                case "foo bar":
                                    v38 := GenS5{F0: v7, F1: false, F2: v6}
                                    _ = v38
                                case "a":
                                    for fi39 := 0; fi39 < 4; fi39++ {
                                        v40 := GenS5{F0: (^(v7)), F1: (v7 < v7), F2: v6}
                                        var v41 uint32 = 2
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                _ = fi34
                            }
                            for fi42 := 0; fi42 < 17; fi42++ {
                                v4 = (-(v4))
                                v43 := v5[4294967295]
                                _ = fi42
                                _ = v43
                            }
                            v44 := GenS1{F0: v7}
                            _ = v44
                        }
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        case "foo bar":
            v4 += 2147483647
        }
    } else {
        _ = len(v1)
        for fi45 := 0; fi45 < 2; fi45++ {
            var v46 float64 = (-2.25 - (-2.25 * 1e10))
            var v47 string = v12
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
