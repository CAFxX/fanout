package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 uint16; F1 bool; F2 bool }

func main() {
    var v1 uint16 = 381
    var v2 float64 = 0.0
    var v3 int8 = 0
    var v4 int32 = -14
    v1 = 2
    fmt.Println(v2, v1, v2)
    switch 127 {
    case 0:
        v5 := GenS1{F0: 65535, F1: 2}
        _ = v5
    case -834:
        v3 -= (^(v3))
    }
    v6 := []int{0, 32767, 32767, 977}
    var v7 uint = 65535
    var v8 int32 = 2147483647
    v4 += (v4 & -208)
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 string = "foo bar"
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: 514, F1: 400}
            if (((^(v1)) < (^(v1))) && false) {
                v13 := GenS1{F0: 56, F1: 42}
                _ = v13
            } else {
                var v14 float64 = v2
                fmt.Println(v4, v7, v3)
                if (v14 <= 0.1) {
                    fmt.Println(v14, v14)
                    v15 := GenS1{F0: 42, F1: 42}
                    fmt.Println(fi11, v10)
                    _ = v15
                }
                v3 += int8(v8)
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    switch ("a" + "a") {
    case "hello":
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS2{F0: 65535, F1: v4}
            _ = fi16
            _ = v17
        }
    default:
        if len(v6) > 0 {
            v6[0] = 0
        }
    }
    if (!((1.5 != (3.14159 - (0.0 / ((3.14159 - float32(v2)) * (0.1 / -2.25))))) && (65535 < 1))) {
        var v18 uint64 = 4294967295
        _ = v18
    } else {
        switch -2147483648 {
        case 534:
            fmt.Println(v8)
        case 2147483647:
            v19 := GenS1{F0: 42, F1: 0}
            _ = v19
        case 32767:
            if true {
                var v20 float32 = 1e10
                fmt.Println(v1)
                _ = v20
            }
        default:
            var v21 int8 = (-(v3))
            _ = v21
        }
        if len(v6) > 0 {
            v6[0] = 127
        }
    }
    v7 = v7
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            switch "\n\t" {
            case "":
                var v24 uint = 2
                _ = v24
            case "a":
                v2 -= 1.5
            case "\n\t":
                if false {
                    if true {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(fi22, fi22, v4)
                        case "":
                            v6 = append(v6, -606)
                        case "foo bar":
                            v7 = (v7 | v7)
                        default:
                            fmt.Println(v1)
                        }
                        switch fi22 {
                        case 1:
                            var v25 int16 = 0
                            _ = v25
                        default:
                            switch "\n\t" {
                            case "":
                                var v26 float64 = v2
                                _ = v26
                            case "\n\t":
                                var v27 float64 = 3.14159
                                _ = v27
                            default:
                                v28 := GenS3{F0: v1, F1: v4, F2: v3}
                                _ = v28
                            }
                        }
                        var v29 int = (fi22 / (((fi23 * (fi23 << ((v7) & 63)))) | 1))
                        _ = v29
                    }
                    var v30 float64 = v2
                    _ = v30
                } else {
                    if (((true || (0.1 <= ((3.14159 * 3.14159) / float32(float64((0.0 * float32(((0.0 - -2.25) / (v2 + (0.0 / v2)))))))))) && (true && (1.5 < 1.5))) && (("" + "a") != ("a" + ("hello" + ("a" + "hello"))))) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    } else {
                        _ = len(v6)
                        var v31 uint16 = (v1 * 488)
                        _ = v31
                    }
                    switch ((("a" + "foo bar") + "foo bar") + ("hello" + ("\n\t" + "\n\t"))) {
                    case "\n\t":
                        v3 -= 42
                    case "a":
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            if (!false) {
                                if true {
                                    switch -609 {
                                    case 42:
                                        v6 = append(v6, (^(fi23)))
                                    case 128:
                                        fmt.Println(v3, fi22, v3)
                                    case 32768:
                                        v8 -= 42
                                    default:
                                        _ = len(v6)
                                    }
                                    if (2 >= 128) {
                                        fmt.Println(v7, fi22)
                                    } else {
                                        if (0.0 != -2.25) {
                                            fmt.Println(v4, fi22)
                                        } else {
                                            var v32 uint64 = 958
                                            _ = v32
                                        }
                                    }
                                    v8 += 410
                                }
                                var v33 uint64 = 1
                                _ = v33
                            } else {
                                var v34 bool = false
                                var v35 int = int(v7)
                                v36 := GenS4{F0: v1, F1: ((false || ((v1 - 628) > (^(v1)))) && v34), F2: v34}
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            switch "" {
                            case "foo bar":
                                switch "hello" {
                                case "a":
                                    var v37 int16 = -43
                                    _ = v37
                                case "":
                                    if false {
                                        v2 += float64(-2.25)
                                        if false {
                                            fmt.Println(v8)
                                        } else {
                                            v4 += 485
                                            var v38 uint64 = 482
                                            v1 -= v1
                                            _ = v38
                                        }
                                        v39 := GenS2{F0: 2147483648, F1: v8}
                                        _ = v39
                                    }
                                }
                            }
                            v40 := GenS1{F0: 2147483648, F1: 364}
                            _ = v40
                        case "foo bar":
                            fmt.Println(v2, v1, v3)
                        default:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    case "hello":
                        var v41 int64 = -9223372036854775808
                        _ = v41
                    default:
                        if true {
                            var v42 uint8 = 1
                            _ = v42
                        }
                    }
                }
            default:
                if len(v6) > 0 {
                    v6[0] = fi22
                }
            }
            v1 = v1
            _ = fi23
        }
        v43 := GenS2{F0: 42, F1: v4}
        _ = fi22
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
