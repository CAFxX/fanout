package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int8; F2 int8 }
func genFn0(p0 uint32, p1 int16) uint8 {
    v4 := []int{-9223372036854775808, 32768}
    var v5 uint64 = 155
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int64 = 9223372036854775807
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return uint8(v5)
}
func genFn1(p0 uint8, p1 uint64, p2 uint) uint64 {
    p2 = p2
    var v8 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p1 | p1)
}

func main() {
    var v1 int32 = 445
    v2 := []string{}
    var v3 int64 = 247
    for fi9 := 0; fi9 < 19; fi9++ {
        switch fi9 {
        case -2147483648:
            fmt.Println(v3, v1, v1)
        }
        _ = fi9
    }
    var v10 bool = (!(false && (((0.1 - 0.1) != 1.5) || true)))
    fmt.Println(v3, v1, v1)
    var v11 uint16 = 42
    var v12 uint = 311
    fmt.Println(v12, v10, v1)
    var v13 uint64 = 18446744073709551615
    switch "\n\t" {
    case "a":
        v14 := GenS1{F0: v11, F1: (v12 << ((255) & 63))}
        _ = v14
    case "hello":
        if v10 {
            var v15 int16 = 1000
            _ = v15
        } else {
            var v16 int8 = 42
            if v10 {
                v16 = v16
                v10 = v10
                v1 = 2147483647
            } else {
                fmt.Println(v1, v10, v13)
                switch 274 {
                case 2147483647:
                    if true {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v12 = 637
                            _ = fi17
                        }
                        if ((((v1 / ((-890) | 1)) > 2147483647) && (v13 > v13)) && v10) {
                            var v18 int32 = (v1 / (((v1 % ((v1) | 1))) | 1))
                            fmt.Println(v16)
                            var v19 int32 = v18
                            _ = v18
                            _ = v19
                        } else {
                            var v20 int32 = v1
                            _ = v20
                        }
                    }
                default:
                    switch 1 {
                    case 710:
                        switch 1 {
                        case 804:
                            v1 += v1
                        case 928:
                            v3 = v3
                        default:
                            var v21 bool = (!(255 >= uint(v13)))
                            _ = v21
                        }
                    case 127:
                        if ((v1 - v1) == v1) {
                            if ((^(v3)) >= v3) {
                                v22 := GenS1{F0: (v11 - v11), F1: v12}
                                fmt.Println(v12, v1)
                                v10 = (1 != 251)
                                _ = v22
                            }
                            v16 -= v16
                            v12 = 65535
                        } else {
                            if (!v10) {
                                switch "hello" {
                                case "\n\t":
                                    v13 = v13
                                case "a":
                                    for fi23 := 0; fi23 < 7; fi23++ {
                                        fmt.Println(v1, v11, v10)
                                        v24 := GenS2{F0: fi23}
                                        var v25 float32 = 1.5
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                    }
                                case "foo bar":
                                    fmt.Println(v1)
                                default:
                                    if false {
                                        switch "a" {
                                        case "foo bar":
                                            v12 -= (v12 & 42)
                                        case "hello":
                                            var v26 float32 = -2.25
                                            _ = v26
                                        }
                                    }
                                }
                            }
                        }
                    default:
                        v1 = 1
                    }
                }
            }
            if v10 {
                v27 := GenS1{F0: 42, F1: (v12 ^ 1)}
                _ = v27
            }
            _ = v16
        }
    }
    v12 = 1
    for fi28 := 0; fi28 < 8; fi28++ {
        if (-128 >= -128) {
            v12 = (v12 | (v12 + 2))
        } else {
            for fi29 := 0; fi29 < 19; fi29++ {
                _ = len(v2)
                _ = fi29
            }
            v3 = v3
            var v30 uint64 = (v13 + genFn1(127, v13, v12))
            _ = v30
        }
        _ = fi28
    }
    var v31 uint32 = 42
    v11 = (v11 - (v11 ^ v11))
    v12 = v12
    v3 += v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v12 -= (v12 >> (((v12 - 18446744073709551615)) & 63))
    switch "hello" {
    case "\n\t":
        switch "hello" {
        case "foo bar":
            for fi32 := 0; fi32 < 19; fi32++ {
                _ = len(v2)
                _ = fi32
            }
        case "":
            v3 = (v3 ^ (^(v3)))
        default:
            v33 := GenS2{F0: -326}
            _ = v33
        }
    case "a":
        v34 := GenS3{F0: "foo bar", F1: -110, F2: -128}
        _ = v34
    }
    if v10 {
        v31 += v31
        var v35 string = ("foo bar" + ("foo bar" + (("foo bar" + "") + (("hello" + "hello") + "a"))))
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
}
