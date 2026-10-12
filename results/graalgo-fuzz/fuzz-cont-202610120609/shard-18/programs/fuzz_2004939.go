package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 string; F2 string }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 int, p1 uint8, p2 uint64) string {
    v4 := GenS1{F0: 0, F1: p2}
    var v5 int16 = -32768
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return ""
}
func genFn1() int16 {
    var v6 int64 = 1
    v6 = v6
    v6 = (v6 >> ((42) & 63))
    switch "hello" {
    case "\n\t":
        fmt.Println(v6)
    case "":
        v6 -= v6
    case "a":
        v6 = v6
    }
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 255
    var v3 int16 = 681
    var v7 uint8 = 157
    switch "hello" {
    case "\n\t":
        if (true || true) {
            v8 := GenS1{F0: 1, F1: 121}
            _ = v8
        } else {
            var v9 uint32 = 2147483647
            _ = v9
        }
    }
    v10 := []int{}
    v11 := GenS1{F0: v7, F1: 0}
    v7 = v7
    switch "" {
    case "\n\t":
        switch 271 {
        case 4294967295:
            var v12 float32 = 3.14159
            _ = v12
        case 2147483647:
            var v13 uint8 = (^(v7))
            _ = v13
        case 9223372036854775807:
            fmt.Println(v3, v2, v2)
        }
    case "":
        var v14 float64 = (-2.25 + -2.25)
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v10 = append(v10, int(v3))
        v16 := GenS1{F0: (v7 / ((v7) | 1)), F1: 42}
        if (!(true && ((0.1 != 0.1) || (((18446744073709551615 <= (v2 / (((v2 / ((v2) | 1))) | 1))) && (!(2147483648 < uint32(v2)))) || (!false))))) {
            v17 := GenS1{F0: v7, F1: 42}
            _ = v17
        } else {
            if true {
                switch "" {
                case "hello":
                    fmt.Println(fi15, fi15)
                case "a":
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if false {
                            if false {
                                v19 := GenS2{F0: -2147483648}
                                var v20 int64 = -285
                                _ = v19
                                _ = v20
                            }
                            var v21 int64 = 9223372036854775807
                            v1[fi15] = 2147483647
                            _ = v21
                        } else {
                            var v22 uint8 = 2
                            _ = v22
                        }
                        var v23 uint64 = 1
                        _ = fi18
                        _ = v23
                    }
                case "foo bar":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        fmt.Println(v3, v7)
                        v7 += v7
                        _ = fi24
                    }
                default:
                    v2 -= v2
                }
                v25 := v1[fi15]
                _ = v25
            } else {
                var v26 int8 = 127
                fmt.Println(v26, fi15)
                var v27 string = "foo bar"
                _ = v26
                _ = v27
            }
            delete(v1, (fi15 / (((fi15 >> ((v2) & 63))) | 1)))
        }
        _ = fi15
        _ = v16
    }
    if true {
        v28 := GenS3{F0: 140, F1: "", F2: "foo bar"}
        _ = v28
    } else {
        v29 := GenS4{F0: 0, F1: (v2 & (v2 - v2))}
        _ = v29
    }
    var v30 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v7
}
