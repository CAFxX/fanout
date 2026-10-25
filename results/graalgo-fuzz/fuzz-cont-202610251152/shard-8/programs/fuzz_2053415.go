package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint16; F2 uint8 }
func genFn0(p0 uint32) int32 {
    if false {
        v7 := GenS1{F0: 1000, F1: 1, F2: "\n\t"}
        v8 := GenS2{F0: 420, F1: 1, F2: (false || true)}
        _ = v7
        _ = v8
    }
    p0 += p0
    if (2 <= uint8(p0)) {
        var v9 uint16 = 1
        p0 = 578
        p0 -= p0
        _ = v9
    } else {
        fmt.Println(p0)
        p0 = (p0 * 611)
        p0 = p0
    }
    v10 := []int{}
    switch "a" {
    case "foo bar":
        switch "hello" {
        case "hello":
            p0 -= p0
        case "\n\t":
            if len(v10) > 0 {
                v10[0] = -9223372036854775808
            }
        }
    }
    _ = p0
    _ = v10
    return 418
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 49
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 128
    v6 := make(map[string]string)
    v2 = (^(v2))
    for fi11 := 0; fi11 < 9; fi11++ {
        v5 += 127
        _ = fi11
    }
    if ((0.1 / (v1 - float32(0.1))) > float32(-2.25)) {
        fmt.Println(v5, v1)
        switch -685 {
        case -992:
            v2 -= v2
        }
        if true {
            for fi12 := 0; fi12 < 3; fi12++ {
                v13 := GenS1{F0: 399, F1: 0, F2: "hello"}
                v14 := GenS3{F0: fi12}
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 uint64 = 1
            v5 += v5
            _ = v15
        }
    }
    v16 := GenS2{F0: v2, F1: 690, F2: true}
    switch -862 {
    case -1:
        v17 := []int{422, 42, 127}
        _ = v17
    default:
        switch "foo bar" {
        case "":
            var v18 uint64 = v2
            _ = v18
        case "hello":
            var v19 uint64 = (v2 - v2)
            _ = v19
        case "a":
            var v20 uint16 = 2
            _ = v20
        }
    }
    var v21 uint8 = v5
    v22 := []int{128, 9223372036854775807}
    fmt.Println(v2)
    v2 += (v2 ^ (v2 >> ((65535) & 63)))
    v23 := GenS4{F0: 504, F1: 54, F2: v21}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
