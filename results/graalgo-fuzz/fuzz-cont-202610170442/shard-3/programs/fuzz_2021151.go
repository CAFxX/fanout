package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint16; F2 int8 }

func main() {
    var v1 string = "hello"
    v2 := make(map[int]int)
    var v3 string = v1
    v4 := make(map[uint32]bool)
    fmt.Println(v3)
    fmt.Println(v1)
    v3 = "a"
    v1 = ""
    var v5 uint8 = 2
    var v6 int8 = -89
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(v1, fi7)
        v8 := []int{128}
        for fi9 := 0; fi9 < 17; fi9++ {
            if len(v8) > 0 {
                v8[0] = 42
            }
            var v10 bool = true
            switch -2147483648 {
            case 9223372036854775807:
                fmt.Println(v5)
            case 127:
                if v10 {
                    if true {
                        fmt.Println(v1)
                        var v11 int8 = v6
                        switch "hello" {
                        case "foo bar":
                            v12 := GenS1{F0: 0}
                            _ = v12
                        case "a":
                            if v10 {
                                var v13 uint8 = (v5 ^ 105)
                                _ = v13
                            }
                        default:
                            fmt.Println(v3, fi9)
                        }
                        _ = v11
                    }
                    switch 42 {
                    case -9223372036854775808:
                        v3 = (("a" + v3) + "hello")
                    default:
                        switch fi7 {
                        case 0:
                            if (2147483648 >= 915) {
                                v14 := GenS1{F0: 18446744073709551615}
                                if v10 {
                                    v8 = append(v8, 128)
                                    delete(v2, (fi7 ^ (fi7 << ((0) & 63))))
                                    switch v3 {
                                    case "":
                                        var v15 uint64 = 42
                                        _ = v15
                                    case "\n\t":
                                        if false {
                                            v6 = v6
                                        }
                                    }
                                } else {
                                    v6 = (-(v6))
                                    var v16 int16 = -32768
                                    if false {
                                        var v17 uint64 = 4294967295
                                        _ = v17
                                    }
                                    if (317 > 946) {
                                        fmt.Println(fi7, fi9, v16)
                                        var v18 uint8 = v5
                                        switch fi9 {
                                        case -2147483648:
                                            var v19 bool = (-1 > 42)
                                            var v20 string = "\n\t"
                                            v21 := GenS2{F0: 42}
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        case 128:
                                            v10 = true
                                        case 865:
                                            var v22 bool = true
                                            _ = v22
                                        }
                                        _ = v18
                                    } else {
                                        fmt.Println(v10, v6, v1)
                                        v16 += (v16 - v16)
                                        v23 := GenS3{F0: 166, F1: (^(v5))}
                                        _ = v23
                                    }
                                    _ = v16
                                }
                                _ = v14
                            }
                        case 127:
                            v3 = v1
                        }
                    }
                    v5 = (v5 - (v5 + (v5 ^ v5)))
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    delete(v2, 128)
    v6 -= (-(v6))
    var v24 float64 = 0.0
    v25 := GenS4{F0: (v5 | 8)}
    var v26 uint = 394
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := []int{(fi27 ^ fi27), fi27, (fi27 + -380), 42}
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 string = v1
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    v31 := GenS5{F0: (v26 & v26), F1: 187, F2: (v6 ^ v6)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
