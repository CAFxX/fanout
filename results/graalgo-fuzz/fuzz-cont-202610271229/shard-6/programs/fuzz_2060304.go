package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 string }
type GenS4 struct { F0 uint8; F1 int; F2 uint8 }
func genFn0(p0 int64, p1 int64) uint64 {
    p1 -= (p1 * (p0 << ((543) & 63)))
    switch ("a" + "foo bar") {
    case "hello":
        var v5 float64 = (3.14159 / 1e10)
        _ = v5
    default:
        v6 := []int{-1, 32768, 1, 163}
        _ = v6
    }
    var v7 uint16 = 1
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint64 = 210
    v2 := []uint32{65535, 341, 4294967295, 2}
    var v3 int64 = -1
    var v4 int8 = -128
    var v8 string = "a"
    var v9 uint16 = 0
    var v10 uint8 = 255
    v4 -= 1
    v3 = (^(v3))
    v8 = "\n\t"
    switch (v8 + "a") {
    case "\n\t":
        var v11 uint64 = uint64(v10)
        _ = v11
    case "a":
        if ((^(v9)) != (v9 << ((uint(v1)) & 63))) {
            var v12 int64 = 42
            v9 = (v9 / ((v9) | 1))
            if (1.5 <= ((1e10 - (1.5 * (3.14159 - 1e10))) - (-2.25 + 1e10))) {
                switch 2147483647 {
                case 2147483647:
                    v9 -= (v9 | (v9 / ((v9) | 1)))
                case 127:
                    v13 := GenS1{F0: 391, F1: 0, F2: ("foo bar" + ("foo bar" + "hello"))}
                    _ = v13
                default:
                    v3 += (^(v12))
                }
                var v14 int64 = 0
                fmt.Println(v8, v8)
                _ = v14
            } else {
                v15 := GenS1{F0: 42, F1: -288, F2: "foo bar"}
                _ = v15
            }
            _ = v12
        }
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS1{F0: 255, F1: fi16, F2: ((v8 + v8) + (v8 + ("hello" + ("foo bar" + v8))))}
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v1)
    if true {
        fmt.Println(v8, v10, v10)
        switch "a" {
        case "hello":
            var v18 int16 = 42
            _ = v18
        case "":
            fmt.Println(v3, v4, v1)
        }
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            if true {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 bool = false
                    var v22 uint64 = genFn0((v3 >> ((802) & 63)), (v3 / ((0) | 1)))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                var v23 uint8 = v10
                switch v8 {
                case "\n\t":
                    switch 2147483647 {
                    case 4294967295:
                        if true {
                            v24 := GenS2{F0: (v1 << ((1) & 63))}
                            var v25 int32 = -577
                            _ = v24
                            _ = v25
                        }
                    case 32767:
                        switch v8 {
                        case "a":
                            v2 = append(v2, 65535)
                        case "":
                            v26 := GenS2{F0: genFn0(-9223372036854775808, 532)}
                            _ = v26
                        default:
                            var v27 uint = 255
                            _ = v27
                        }
                    case -9223372036854775808:
                        for fi28 := 0; fi28 < 17; fi28++ {
                            switch v8 {
                            case "foo bar":
                                v2 = append(v2, 767)
                            case "":
                                var v29 int = 1
                                _ = v29
                            case "a":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = fi28
                        }
                    }
                default:
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v31 := GenS3{F0: 1, F1: 1, F2: v8}
                        v32 := GenS3{F0: 2147483648, F1: v1, F2: v8}
                        v33 := GenS4{F0: v23, F1: fi30, F2: 128}
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
                _ = v23
            }
            _ = fi19
        }
        if (true || (!(-1 <= v3))) {
            v8 = v8
            v34 := GenS3{F0: 1, F1: (^(v1)), F2: v8}
            var v35 int = 127
            _ = v34
            _ = v35
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, 625)
            var v36 uint8 = (v10 / ((v10) | 1))
            _ = v36
        }
        var v37 uint16 = v9
        _ = v37
    }
    if ((-(v3)) <= v3) {
        var v38 int = -9223372036854775808
        _ = v38
    }
    fmt.Println(v8, v4)
    _ = len(v2)
    v10 = 231
    fmt.Println(v1, v8, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
