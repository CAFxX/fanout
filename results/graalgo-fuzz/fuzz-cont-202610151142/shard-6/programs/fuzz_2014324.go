package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int; F1 uint }
func genFn0() uint32 {
    v7 := GenS1{F0: "foo bar", F1: 2, F2: 357}
    switch "" {
    case "hello":
        switch "\n\t" {
        case "hello":
            v8 := GenS1{F0: "a", F1: 42, F2: 2}
            _ = v8
        case "\n\t":
            var v9 uint16 = 255
            _ = v9
        default:
            for fi10 := 0; fi10 < 13; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 uint = 705
                    if (!false) {
                        v12 = 0
                    } else {
                        v13 := []int{(fi10 + 2147483647)}
                        _ = v13
                    }
                    v14 := GenS2{F0: 9223372036854775807, F1: -629}
                    _ = fi11
                    _ = v12
                    _ = v14
                }
                _ = fi10
            }
        }
    }
    v15 := GenS2{F0: 0, F1: -300}
    var v16 uint32 = 4294967295
    v16 = v16
    _ = v15
    _ = v16
    _ = v7
    return (v16 ^ (v16 * 2147483647))
}

func main() {
    var v1 uint32 = 2
    var v2 uint32 = 812
    v3 := make(map[int]uint32)
    v4 := []int{-9223372036854775808, 128}
    var v5 string = "hello"
    var v6 uint32 = (v2 + 2)
    v2 -= (^(v2))
    v17 := GenS1{F0: v5, F1: 18446744073709551615, F2: 0}
    v5 = v5
    for fi18 := 0; fi18 < 3; fi18++ {
        if (-32768 < 0) {
            var v19 uint32 = v2
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 int8 = 19
                _ = fi20
                _ = v21
            }
            var v22 uint = 0
            _ = v19
            _ = v22
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            v6 = v6
            v24 := GenS2{F0: -988, F1: 42}
            _ = fi23
            _ = v24
        }
        switch "hello" {
        case "hello":
            var v25 uint64 = 42
            _ = v25
        case "":
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 int16 = 32767
                if (!false) {
                    switch v5 {
                    case "":
                        var v28 uint16 = 43
                        _ = v28
                    case "\n\t":
                        v29 := GenS3{F0: (fi18 << ((0) & 63)), F1: 2}
                        _ = v29
                    case "foo bar":
                        v2 += 0
                    default:
                        var v30 int32 = -2147483648
                        _ = v30
                    }
                    switch v5 {
                    case "hello":
                        var v31 float32 = 1.5
                        _ = v31
                    case "a":
                        fmt.Println(v5, v27, fi18)
                        var v32 uint16 = 231
                        _ = v32
                    default:
                        v33 := GenS3{F0: -414, F1: 604}
                        _ = v33
                    }
                }
                _ = fi26
                _ = v27
            }
        default:
            var v34 float64 = (0.1 / -2.25)
            _ = v34
        }
        _ = fi18
    }
    v35 := GenS2{F0: 9223372036854775807, F1: -32768}
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 float32 = (1.5 - (0.0 / 1.0))
        fmt.Println(v1)
        _ = fi36
        _ = v37
    }
    delete(v3, 294)
    v2 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
