package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }

func main() {
    var v1 int32 = -703
    v2 := make(map[uint32]int)
    var v3 int8 = 1
    var v4 int = 42
    v5 := make(map[uint32]uint32)
    v6 := GenS1{F0: 42, F1: 710, F2: 742}
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint32 = 65535
        if (("a" == "foo bar") || false) {
            var v9 float64 = -2.25
            _ = v9
        } else {
            var v10 uint64 = 1
            for fi11 := 0; fi11 < 8; fi11++ {
                switch "" {
                case "hello":
                    if (fi11 <= 32768) {
                        switch "" {
                        case "\n\t":
                            v12 := []int{(v4 | fi11)}
                            _ = v12
                        case "hello":
                            var v13 int16 = -707
                            _ = v13
                        }
                        v8 = 4294967295
                        fmt.Println(v8, v1, v8)
                        delete(v2, (v8 - (v8 ^ v8)))
                        v14 := []int{fi7, (fi11 % ((v4) | 1)), 2147483647}
                        v3 = (v3 - 100)
                        _ = v14
                    } else {
                        v10 += v10
                    }
                case "foo bar":
                    switch 4294967295 {
                    case 9223372036854775807:
                        var v15 uint64 = 6
                        _ = v15
                    case 493:
                        fmt.Println(fi11)
                    case 127:
                        var v16 int64 = 42
                        _ = v16
                    }
                case "\n\t":
                    if (194 == 591) {
                        v17 := v2[(v8 + v8)]
                        var v18 uint32 = v8
                        v19 := v2[458]
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi11
            }
            for fi20 := 0; fi20 < 18; fi20++ {
                if true {
                    var v21 int64 = 1
                    switch "" {
                    case "foo bar":
                        v4 = -883
                    case "a":
                        fmt.Println(fi20, v4, v3)
                    default:
                        v22 := v2[v8]
                        _ = v22
                    }
                    v3 = 37
                    _ = v21
                }
                _ = fi20
            }
            _ = v10
        }
        if (42 >= 2147483647) {
            v3 += int8(v8)
            v23 := []int{fi7, fi7, int(v8)}
            var v24 uint32 = 502
            _ = v23
            _ = v24
        }
        _ = fi7
        _ = v8
    }
    var v25 int64 = -952
    v26 := GenS1{F0: v25, F1: 226, F2: 819}
    v4 = v4
    v27 := v5[4294967295]
    v28 := GenS1{F0: (v25 >> ((792) & 63)), F1: 65535, F2: 28}
    fmt.Println(v25, v3, v4)
    var v29 uint16 = 65535
    for fi30 := 0; fi30 < 5; fi30++ {
        v4 += (v4 ^ 461)
        fmt.Println(v29, fi30, fi30)
        for fi31 := 0; fi31 < 7; fi31++ {
            fmt.Println(fi31, fi30)
            var v32 float64 = (0.1 * 1e10)
            v32 -= ((v32 * v32) * v32)
            var v33 string = "hello"
            v4 = -9223372036854775808
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    v34 := []int{-9223372036854775808, v4, (v4 + v4)}
    for fi35 := 0; fi35 < 1; fi35++ {
        _ = len(v34)
        _ = len(v34)
        var v36 int64 = 786
        _ = fi35
        _ = v36
    }
    var v37 float64 = 0.0
    if false {
        var v38 float32 = (0.1 + 1e10)
        var v39 int32 = v1
        _ = v38
        _ = v39
    } else {
        v2[4294967295] = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
