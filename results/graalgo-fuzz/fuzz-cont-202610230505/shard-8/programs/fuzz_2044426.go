package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int32; F2 int }
func genFn0(p0 uint8, p1 int32, p2 uint) int8 {
    v4 := []int{-2147483648, 1}
    if false {
        var v5 uint16 = 255
        _ = v5
    } else {
        if len(v4) > 0 {
            v4[0] = 32767
        }
        v6 := GenS1{F0: 32767, F1: p0}
        switch -64 {
        case 128:
            for fi7 := 0; fi7 < 4; fi7++ {
                if false {
                    p1 = p1
                } else {
                    var v8 float64 = float64(-2.25)
                    p1 = (p1 | (p1 << (((p2 >> (((p2 & uint(p0))) & 63))) & 63)))
                    p1 += p1
                    _ = v8
                }
                _ = fi7
            }
        default:
            var v9 int = 2147483647
            _ = v9
        }
        _ = v6
    }
    switch 4294967295 {
    case 9223372036854775807:
        if ("" < "hello") {
            v10 := GenS1{F0: 128, F1: p0}
            if true {
                p0 += (p0 ^ (p0 * p0))
                p1 -= p1
                v11 := GenS2{F0: -2147483648}
                _ = v11
            }
            fmt.Println(p0)
            _ = v10
        }
    case 4294967295:
        var v12 int8 = 1
        _ = v12
    case 128:
        v13 := GenS1{F0: 0, F1: 1}
        _ = v13
    }
    switch "a" {
    case "foo bar":
        var v14 float64 = float64(1e10)
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    var v3 int8 = 0
    fmt.Println(v3)
    v15 := GenS1{F0: 127, F1: 191}
    var v16 float64 = 1e10
    if true {
        fmt.Println(v16, v3)
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 uint16 = 0
            v19 := GenS2{F0: 42}
            fmt.Println(v18, v18)
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := GenS2{F0: 9223372036854775807}
        _ = v20
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    _ = len(v1)
    v16 = 0.1
    v16 -= v16
    var v21 int64 = -847
    var v22 int64 = v21
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS3{F0: 65535}
        _ = fi23
        _ = v24
    }
    var v25 uint = 226
    v26 := GenS4{F0: "foo bar", F1: 1, F2: 9223372036854775807}
    var v27 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
