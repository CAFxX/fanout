package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 uint }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0(p0 uint8, p1 int16, p2 int32) uint64 {
    p0 += (p0 >> ((1) & 63))
    if ((-23 < 127) && (-13 != 127)) {
        var v4 int32 = p2
        v5 := GenS1{F0: 42, F1: 2147483648}
        v6 := GenS2{F0: "hello", F1: 618, F2: 12}
        _ = v4
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 654
}
func genFn1(p0 uint64, p1 uint64) int16 {
    var v7 float64 = 1.5
    fmt.Println(p1)
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = (((1e10 + (float64((0.1 / 1e10)) / -2.25)) - 1.5) - (0.1 - -2.25))
    var v3 uint8 = 136
    var v8 int16 = 1
    v9 := GenS2{F0: (("foo bar" + "hello") + "foo bar"), F1: 767, F2: 65535}
    fmt.Println(v1, v2, v8)
    v10 := []int{-2147483648, -74, 194, 42}
    var v11 bool = (!(127 < 2147483647))
    v2 = v2
    v8 -= 1
    if ("hello" > "") {
        v12 := GenS3{F0: 2, F1: 42}
        var v13 string = "hello"
        v10 = append(v10, -9223372036854775808)
        _ = v12
        _ = v13
    }
    if (!("\n\t" >= ("hello" + ("" + (("hello" + "foo bar") + "\n\t"))))) {
        var v14 int64 = -552
        for fi15 := 0; fi15 < 20; fi15++ {
            v10 = append(v10, (fi15 ^ fi15))
            _ = fi15
        }
        if v11 {
            var v16 uint = 65535
            if v11 {
                if v11 {
                    v3 = (v3 % ((100) | 1))
                } else {
                    if (!v11) {
                        switch "hello" {
                        case "hello":
                            var v17 int8 = -1
                            _ = v17
                        case "foo bar":
                            v18 := GenS4{F0: (v8 % ((v8) | 1)), F1: 1000}
                            _ = v18
                        case "\n\t":
                            v19 := GenS5{F0: (v3 << ((667) & 63)), F1: "a", F2: 1}
                            _ = v19
                        default:
                            if v11 {
                                v20 := GenS2{F0: "hello", F1: v14, F2: 0}
                                switch -9223372036854775808 {
                                case 755:
                                    var v21 int = 127
                                    _ = v21
                                case -666:
                                    for fi22 := 0; fi22 < 6; fi22++ {
                                        if len(v10) > 0 {
                                            v10[0] = fi22
                                        }
                                        var v23 bool = v11
                                        for fi24 := 0; fi24 < 6; fi24++ {
                                            fmt.Println(v16, v23)
                                            v2 -= v2
                                            _ = fi24
                                        }
                                        _ = fi22
                                        _ = v23
                                    }
                                case 32767:
                                    if len(v10) > 0 {
                                        v10[0] = -531
                                    }
                                default:
                                    v10 = append(v10, -513)
                                }
                                _ = v20
                            }
                        }
                    }
                }
            }
            _ = v16
        } else {
            var v25 uint64 = 612
            var v26 uint16 = 2
            if false {
                var v27 bool = false
                v26 -= v26
                _ = v27
            }
            _ = v25
            _ = v26
        }
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
