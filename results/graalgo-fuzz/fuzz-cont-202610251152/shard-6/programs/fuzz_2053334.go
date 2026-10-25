package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 bool; F1 uint32; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 int64; F2 string }
func genFn0() int {
    var v4 int16 = 32767
    v5 := GenS1{F0: 10, F1: 9223372036854775807}
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 int8 = -106
    v2 := []uint32{}
    var v3 bool = true
    if v3 {
        for fi6 := 0; fi6 < 7; fi6++ {
            switch "\n\t" {
            case "":
                for fi7 := 0; fi7 < 16; fi7++ {
                    v3 = ((!v3) || v3)
                    var v8 float64 = 1e10
                    _ = fi7
                    _ = v8
                }
            case "hello":
                v1 = v1
            case "a":
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            default:
                v9 := GenS2{F0: ""}
                _ = v9
            }
            if true {
                if (!v3) {
                    if v3 {
                        v1 = -46
                    } else {
                        switch "hello" {
                        case "hello":
                            switch "" {
                            case "a":
                                switch genFn0() {
                                case 2147483647:
                                    v10 := GenS2{F0: "a"}
                                    _ = v10
                                case 128:
                                    var v11 float64 = 3.14159
                                    _ = v11
                                default:
                                    var v12 uint32 = 4
                                    _ = v12
                                }
                            case "\n\t":
                                v13 := GenS3{F0: 128, F1: -1, F2: v3}
                                _ = v13
                            }
                        case "a":
                            v14 := GenS3{F0: 1, F1: -9223372036854775808, F2: false}
                            _ = v14
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                } else {
                    fmt.Println(fi6, v1)
                    _ = len(v2)
                    var v15 int8 = 127
                    _ = v15
                }
                var v16 int16 = -786
                _ = v16
            }
            var v17 int64 = -1
            _ = fi6
            _ = v17
        }
        var v18 uint8 = 128
        v18 = v18
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi19
    }
    fmt.Println(v1, v3)
    if v3 {
        if ((-733 <= 102) && (((-2.25 - 0.0) != 0.1) && true)) {
            v1 = v1
            fmt.Println(v3, v1)
        }
        v1 += 0
    } else {
        v20 := GenS2{F0: ""}
        _ = len(v2)
        if v3 {
            v2 = append(v2, 872)
            for fi21 := 0; fi21 < 14; fi21++ {
                v3 = v3
                _ = fi21
            }
        } else {
            v3 = v3
            switch "\n\t" {
            case "foo bar":
                var v22 bool = v3
                _ = v22
            case "a":
                for fi23 := 0; fi23 < 13; fi23++ {
                    if ((-728 >= 498) || (((!((("hello" <= ("" + "foo bar")) || (227 != 109)) || ((2 != 4294967295) || false))) || (!(!true))) || (-706 >= 0))) {
                        v1 -= v1
                        if (fi23 > (fi23 - -513)) {
                            for fi24 := 0; fi24 < 14; fi24++ {
                                v3 = (((((!v3) || v3) || (451 < int16(v1))) || (((fi23 / (((fi23 << ((1) & 63))) | 1)) > (fi24 | (fi23 / ((genFn0()) | 1)))) || v3)) && v3)
                                _ = fi24
                            }
                        } else {
                            _ = len(v2)
                            if (2 != 58) {
                                v25 := GenS4{F0: false, F1: 2147483648, F2: ("a" + "\n\t")}
                                _ = v25
                            }
                            var v26 int8 = 0
                            _ = v26
                        }
                        fmt.Println(fi23, v1)
                    } else {
                        v2 = append(v2, 29)
                        switch (fi23 & (fi23 ^ fi23)) {
                        case -1:
                            v1 = (-(v1))
                        default:
                            switch genFn0() {
                            case 32767:
                                v3 = ((!(!(-1 >= 9223372036854775807))) && v3)
                            case -9223372036854775808:
                                v3 = (!(42 >= 4294967295))
                            default:
                                v1 = v1
                            }
                        }
                    }
                    v27 := GenS1{F0: 255, F1: 42}
                    _ = fi23
                    _ = v27
                }
            default:
                var v28 int = 258
                _ = v28
            }
            if len(v2) > 0 {
                v2[0] = 1
            }
        }
        _ = v20
    }
    switch 2147483647 {
    case 9223372036854775807:
        fmt.Println(v3)
    case 2147483647:
        v29 := GenS5{F0: 20}
        _ = v29
    default:
        if (v3 && v3) {
            switch "\n\t" {
            case "\n\t":
                var v30 string = "a"
                _ = v30
            }
        }
    }
    v31 := GenS1{F0: 0, F1: -185}
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 uint64 = 42
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v3)
    _ = len(v2)
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS6{F0: "", F1: -9223372036854775808, F2: ((("" + "a") + "foo bar") + ("a" + ("\n\t" + (("\n\t" + "a") + (("a" + "hello") + "foo bar")))))}
        v3 = ("" < "a")
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
