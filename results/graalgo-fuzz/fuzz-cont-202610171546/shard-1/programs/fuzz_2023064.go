package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 int }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]int)
    var v3 uint64 = 4294967295
    var v4 bool = (false || (18446744073709551615 != v3))
    v5 := []uint32{2}
    v6 := make(map[string]string)
    v7 := GenS1{F0: 255, F1: -1}
    if (1 != 1) {
        switch 0 {
        case 0:
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v4 = v4
                    var v10 int8 = 83
                    v3 = (^(v3))
                    var v11 float64 = 3.14159
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
        case 127:
            delete(v2, "foo bar")
        case 4294967295:
            fmt.Println(v1, v4)
        }
    }
    v12 := GenS2{F0: 42, F1: 2, F2: -2147483648}
    v1 = (v1 & 42)
    if false {
        _ = len(v6)
        var v13 int16 = -194
        _ = v13
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            v3 = (v3 / ((18446744073709551615) | 1))
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint16 = 234
                if v4 {
                    v1 += v1
                    v4 = v4
                    var v17 uint16 = v16
                    _ = v17
                }
                var v18 uint64 = v3
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch "foo bar" {
                case "\n\t":
                    if v4 {
                        var v19 uint32 = 2147483648
                        v20 := GenS1{F0: uint8(v16), F1: 23}
                        switch (fi15 - 341) {
                        case 128:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case -1:
                            v21 := GenS3{F0: v16, F1: 0}
                            var v22 uint16 = (v16 % (((v16 >> ((uint(v19)) & 63))) | 1))
                            var v23 int64 = -1
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = v19
                        _ = v20
                    } else {
                        v24 := GenS2{F0: 9223372036854775807, F1: v16, F2: 127}
                        fmt.Println(v1, v18)
                        fmt.Println(v16, v16)
                        var v25 uint16 = (v16 | uint16(v18))
                        _ = v24
                        _ = v25
                    }
                case "foo bar":
                    if len(v5) > 0 {
                        v5[0] = 224
                    }
                }
                _ = fi15
                _ = v16
                _ = v18
            }
            v1 = v1
            _ = fi14
        }
        var v26 uint = 0
        _ = v26
    }
    v1 += v1
    var v27 uint64 = v3
    for fi28 := 0; fi28 < 9; fi28++ {
        if (true || (-1 <= 1000)) {
            var v29 int64 = 1
            _ = v29
        }
        _ = fi28
    }
    var v30 int64 = 318
    fmt.Println(v4)
    v1 += (^(v1))
    var v31 uint = 4294967295
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "foo bar" {
    case "":
        if (!true) {
            v32 := v6[("\n\t" + ("foo bar" + "foo bar"))]
            _ = v32
        } else {
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 int32 = 113
                switch (("hello" + "foo bar") + "a") {
                case "hello":
                    delete(v2, "a")
                case "foo bar":
                    var v35 string = "\n\t"
                    _ = v35
                default:
                    fmt.Println(v1, v3, v1)
                }
                _ = fi33
                _ = v34
            }
        }
    case "\n\t":
        switch "foo bar" {
        case "hello":
            var v36 uint = (v31 % ((2) | 1))
            _ = v36
        }
    case "foo bar":
        for fi37 := 0; fi37 < 11; fi37++ {
            if ((128 >= 0) && v4) {
                fmt.Println(v27, fi37, v30)
            } else {
                var v38 int16 = 506
                v39 := GenS4{F0: 1}
                _ = v38
                _ = v39
            }
            fmt.Println(v27, v30)
            _ = fi37
        }
    default:
        v4 = (v4 && true)
    }
    v31 = v31
    v40 := GenS1{F0: 30, F1: 1}
    var v41 uint32 = 538
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
