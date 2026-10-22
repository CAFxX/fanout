package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int64; F1 string; F2 int8 }

func main() {
    var v1 int16 = 198
    v2 := []uint32{42, 427, 2147483648, 2147483647}
    v3 := make(map[uint32]string)
    var v4 float64 = float64(-2.25)
    var v5 float64 = 3.14159
    var v6 uint = 65535
    v6 -= (^(v6))
    var v7 int32 = -692
    v5 -= 0.1
    v8 := GenS1{F0: 0, F1: 1, F2: v7}
    switch 0 {
    case 32767:
        if false {
            if true {
                if (7 > 1) {
                    if (!false) {
                        for fi9 := 0; fi9 < 19; fi9++ {
                            for fi10 := 0; fi10 < 2; fi10++ {
                                v5 -= v4
                                _ = len(v3)
                                _ = fi10
                            }
                            _ = fi9
                        }
                        _ = len(v2)
                    }
                    var v11 bool = (2 >= 42)
                    if (128 != 2) {
                        for fi12 := 0; fi12 < 20; fi12++ {
                            v13 := v3[2147483648]
                            _ = fi12
                            _ = v13
                        }
                        var v14 int16 = v1
                        v15 := v3[422]
                        _ = v14
                        _ = v15
                    }
                    _ = v11
                }
                v4 += v5
                v6 += v6
            } else {
                var v16 uint32 = 2147483648
                if (0 != 0) {
                    v4 = v4
                    var v17 int = 1
                    v7 = (v7 / ((v7) | 1))
                    _ = v17
                } else {
                    switch ("" + "\n\t") {
                    case "foo bar":
                        for fi18 := 0; fi18 < 1; fi18++ {
                            if true {
                                fmt.Println(v4, v6)
                            } else {
                                var v19 uint64 = 4294967295
                                v20 := GenS2{F0: 32768}
                                v21 := GenS2{F0: fi18}
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            _ = fi18
                        }
                    case "\n\t":
                        _ = len(v3)
                    case "a":
                        var v22 uint8 = 148
                        _ = v22
                    }
                    var v23 uint8 = 42
                    _ = v23
                }
                _ = v16
            }
        }
    }
    if (v7 > 2147483647) {
        if (1 >= 32767) {
            v3[2147483648] = "\n\t"
        } else {
            v24 := GenS1{F0: 127, F1: 105, F2: (v7 + (v7 * (v7 << ((v6) & 63))))}
            fmt.Println(v6, v6)
            if true {
                if false {
                    v1 -= (v1 - (v1 / ((-32768) | 1)))
                    fmt.Println(v4, v5, v7)
                }
                if false {
                    v2 = append(v2, 342)
                } else {
                    v25 := GenS3{F0: false, F1: 789}
                    _ = v25
                }
            }
            _ = v24
        }
    }
    fmt.Println(v1)
    if (!true) {
        v1 = 0
    }
    switch ("hello" + (("foo bar" + ("hello" + "hello")) + ("" + ""))) {
    case "a":
        v26 := v3[2147483648]
        _ = v26
    default:
        fmt.Println(v6, v1, v5)
    }
    switch "\n\t" {
    case "hello":
        switch "hello" {
        case "a":
            v27 := v3[431]
            _ = v27
        case "":
            switch "foo bar" {
            case "a":
                v7 -= (-(v7))
            case "foo bar":
                _ = len(v3)
            case "":
                v2 = append(v2, uint32(v6))
            }
        default:
            if true {
                var v28 int64 = 1
                _ = v28
            }
        }
    }
    v3[349] = (("\n\t" + "") + "")
    v29 := GenS3{F0: true, F1: 2}
    v30 := GenS4{F0: 435, F1: "\n\t", F2: 0}
    var v31 int = 9223372036854775807
    v4 += v4
    _ = len(v3)
    var v32 bool = false
    for fi33 := 0; fi33 < 9; fi33++ {
        _ = len(v2)
        v7 = (v7 ^ v7)
        _ = len(v2)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
