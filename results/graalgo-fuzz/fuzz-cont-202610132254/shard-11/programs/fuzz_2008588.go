package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8; F1 int8; F2 bool }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }

func main() {
    var v1 int64 = -1
    v2 := []uint32{2147483647}
    var v3 uint64 = 704
    var v4 uint = 42
    v1 += v1
    switch 0 {
    case 32767:
        v3 -= v3
    case -2147483648:
        var v5 uint16 = 847
        _ = v5
    default:
        switch 42 {
        case 605:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            fmt.Println(v4, v4, v1)
        }
    }
    fmt.Println(v4, v4, v1)
    var v6 float64 = 0.0
    switch ("foo bar" + "") {
    case "\n\t":
        switch 128 {
        case 32767:
            v7 := GenS1{F0: 282, F1: 245}
            _ = v7
        default:
            _ = len(v2)
        }
    }
    v4 += v4
    v8 := GenS2{F0: -98, F1: 110, F2: true}
    var v9 float64 = v6
    switch -9223372036854775808 {
    case 9:
        v1 += v1
    case 128:
        var v10 string = "\n\t"
        _ = v10
    case 2147483647:
        if true {
            if true {
                var v11 uint32 = 317
                _ = v11
            } else {
                if len(v2) > 0 {
                    v2[0] = 42
                }
            }
        } else {
            _ = len(v2)
            switch -2147483648 {
            case 128:
                var v12 bool = false
                _ = v12
            case 2147483647:
                if (2 <= 42) {
                    var v13 int32 = -720
                    var v14 int8 = 0
                    if true {
                        _ = len(v2)
                        var v15 uint32 = 441
                        _ = v15
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            for fi16 := 0; fi16 < 1; fi16++ {
                                if false {
                                    v13 = -493
                                    if false {
                                        fmt.Println(v4, v4)
                                        v3 = (v3 * v3)
                                        var v17 int16 = 42
                                        _ = v17
                                    } else {
                                        var v18 string = "foo bar"
                                        _ = v18
                                    }
                                    v19 := GenS3{F0: (fi16 >> ((18446744073709551615) & 63)), F1: 42, F2: (v14 - 1)}
                                    _ = v19
                                } else {
                                    var v20 int = fi16
                                    v13 = v13
                                    switch "foo bar" {
                                    case "":
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v6, v3, v9)
                                        case "hello":
                                            fmt.Println(v13, v3)
                                        case "a":
                                            var v21 int16 = 32767
                                            _ = v21
                                        default:
                                            var v22 uint64 = (v3 - v3)
                                            _ = v22
                                        }
                                    case "hello":
                                        var v23 uint64 = (v3 / (((v3 % (((^(v3))) | 1))) | 1))
                                        _ = v23
                                    case "foo bar":
                                        if (255 < 167) {
                                            var v24 float64 = 0.1
                                            _ = v24
                                        }
                                    }
                                    _ = v20
                                }
                                var v25 int = (fi16 + fi16)
                                _ = fi16
                                _ = v25
                            }
                        }
                        var v26 int64 = (^(v1))
                        v27 := GenS1{F0: 42, F1: 9223372036854775807}
                        _ = v26
                        _ = v27
                    }
                    _ = v13
                    _ = v14
                }
            default:
                for fi28 := 0; fi28 < 5; fi28++ {
                    switch "hello" {
                    case "":
                        fmt.Println(v4, v4, v9)
                    case "hello":
                        v2 = append(v2, 4294967295)
                    case "foo bar":
                        if (true || (-128 == -102)) {
                            v9 += ((0.1 * 0.0) * v6)
                            v29 := GenS4{F0: (0 == 1)}
                            _ = v29
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v2 = append(v2, 1)
                            var v30 int8 = 57
                            _ = v30
                        }
                    default:
                        fmt.Println(v3, fi28, v9)
                    }
                    _ = len(v2)
                    _ = fi28
                }
            }
        }
    default:
        if len(v2) > 0 {
            v2[0] = 481
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= 426
    v1 -= (v1 ^ (v1 & (v1 & -9223372036854775808)))
    v31 := GenS5{F0: -1}
    for fi32 := 0; fi32 < 20; fi32++ {
        for fi33 := 0; fi33 < 5; fi33++ {
            var v34 int16 = 1
            if len(v2) > 0 {
                v2[0] = 2
            }
            v6 = (v6 + (0.0 / 0.1))
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
