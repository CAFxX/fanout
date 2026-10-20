package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) uint32 {
    var v5 string = "a"
    var v6 uint16 = 2
    var v7 uint = 2
    var v8 float32 = -2.25
    v5 = v5
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []int{-661, -804}
    v2 := []int{9223372036854775807, 0, -9223372036854775808}
    var v3 float32 = 0.1
    var v4 uint32 = 65535
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v4, fi9)
        var v10 int16 = 1000
        var v11 float64 = 3.14159
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int64 = 0
    if true {
        var v13 bool = (2147483647 >= 4294967295)
        v3 = v3
        var v14 uint16 = 42
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 bool = (true && false)
        v12 = v12
        _ = fi15
        _ = v16
    }
    var v17 string = "\n\t"
    switch 4294967295 {
    case 96:
        var v18 string = v17
        _ = v18
    case 32767:
        switch ((((v17 + v17) + v17) + "a") + v17) {
        case "\n\t":
            v3 -= (v3 * v3)
        case "hello":
            _ = len(v2)
        }
    case 127:
        var v19 uint32 = (v4 % ((v4) | 1))
        _ = v19
    }
    switch "\n\t" {
    case "hello":
        var v20 int64 = (v12 + -1)
        _ = v20
    case "foo bar":
        v21 := GenS1{F0: v17, F1: 42, F2: 15}
        _ = v21
    case "a":
        v22 := GenS2{F0: 0, F1: 0}
        _ = v22
    }
    _ = len(v1)
    v4 = (^(v4))
    fmt.Println(v3, v4)
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if (!(4294967295 >= 42)) {
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 float32 = v3
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    } else {
        switch "" {
        case "foo bar":
            v4 = (v4 & (^(v4)))
        }
    }
    var v26 uint = 65
    v3 = v3
    v27 := GenS3{F0: 2147483647}
    v26 -= (v26 >> ((v26) & 63))
    switch v17 {
    case "":
        switch "foo bar" {
        case "foo bar":
            var v28 uint64 = 2
            _ = v28
        }
    case "hello":
        v3 -= v3
    default:
        v26 = (v26 - (v26 ^ v26))
    }
    var v29 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
