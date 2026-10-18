package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }
func genFn0(p0 int32, p1 uint16, p2 uint64) int32 {
    p2 = (^(p2))
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int32 = 42
        _ = fi5
        _ = v6
    }
    p1 = 373
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 4
    v3 := make(map[uint32]bool)
    var v4 bool = false
    v7 := []int{9223372036854775807}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 := GenS1{F0: 60, F1: 2147483647, F2: ("hello" + ("hello" + (("hello" + ("hello" + ("foo bar" + "a"))) + ("a" + ("a" + "\n\t")))))}
    v7 = append(v7, 32767)
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS2{F0: (v2 * v2), F1: 634, F2: fi9}
        _ = fi9
        _ = v10
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v2 += v2
    fmt.Println(v2, v4)
    if (!true) {
        var v11 int16 = 32767
        _ = v11
    } else {
        var v12 int8 = 1
        switch "hello" {
        case "":
            var v13 uint64 = 857
            _ = v13
        case "foo bar":
            fmt.Println(v2, v12, v12)
        }
        _ = v12
    }
    v1[954] = 73
    var v14 float32 = 0.1
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int64 = -9223372036854775808
        if (548 >= 4294967295) {
            var v17 int32 = -2147483648
            if v4 {
                switch fi15 {
                case -2147483648:
                    v1[4294967295] = 1
                default:
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        if len(v7) > 0 {
                            v7[0] = (fi15 % ((fi15) | 1))
                        }
                    }
                }
                v17 += int32(fi15)
            }
            if (v16 >= 9223372036854775807) {
                if ((v17 % ((2147483647) | 1)) < v17) {
                    v3[42] = v4
                }
                v1[97] = 65535
            } else {
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS2{F0: (^(v2)), F1: 0, F2: -2147483648}
                    _ = fi18
                    _ = v19
                }
                if v4 {
                    switch "\n\t" {
                    case "a":
                        var v20 string = "\n\t"
                        _ = v20
                    default:
                        switch (("a" + "a") + "foo bar") {
                        case "\n\t":
                            var v21 uint64 = 42
                            _ = v21
                        }
                    }
                } else {
                    fmt.Println(v17, fi15)
                    if true {
                        if (genFn0(-2147483648, 665, 351) >= (v17 << ((2) & 63))) {
                            var v22 uint32 = 70
                            var v23 int = (fi15 % ((0) | 1))
                            _ = v22
                            _ = v23
                        } else {
                            if (1 >= 42) {
                                var v24 int16 = 0
                                var v25 int64 = v16
                                _ = v24
                                _ = v25
                            } else {
                                v16 = 42
                                if true {
                                    v26 := GenS2{F0: v2, F1: 2, F2: fi15}
                                    _ = v26
                                }
                                if (!v4) {
                                    if (!((v16 % ((-158) | 1)) <= (v16 | (v16 + v16)))) {
                                        fmt.Println(v14, v2, v17)
                                    }
                                    var v27 float32 = (v14 / 0.1)
                                    var v28 float64 = ((1e10 / ((((1e10 + 1e10) * (1e10 - (((3.14159 * 1e10) / (0.1 / ((float64((float32((0.0 + (-2.25 - (1e10 + 1.5)))) * v14)) + float64(float32(-2.25))) - (3.14159 / 1.0)))) + float64(v14)))) + 1.5) + 1e10)) / 3.14159)
                                    _ = v27
                                    _ = v28
                                }
                            }
                        }
                        switch "a" {
                        case "hello":
                            fmt.Println(v4)
                        default:
                            v16 += (v16 % ((v16) | 1))
                        }
                        v3[uint32(v2)] = v4
                    }
                }
                fmt.Println(v17, v17)
            }
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    _ = len(v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
