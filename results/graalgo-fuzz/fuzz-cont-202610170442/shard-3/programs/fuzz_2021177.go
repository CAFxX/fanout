package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0(p0 uint64) uint64 {
    if true {
        p0 -= (p0 * 1)
    } else {
        p0 = 42
    }
    p0 = (p0 << ((18446744073709551615) & 63))
    var v4 bool = ((90 >= 42) || false)
    var v5 int64 = 9223372036854775807
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(p0, v4)
        var v7 int32 = 0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint64) string {
    switch "hello" {
    case "foo bar":
        p0 += p0
    case "hello":
        p0 -= (p0 << ((119) & 63))
    }
    var v8 int = -1
    p0 -= p0
    _ = p0
    _ = v8
    return "foo bar"
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]bool)
    var v3 int16 = (v1 ^ v1)
    v3 = 32767
    v9 := GenS1{F0: ((!(false && false)) || true), F1: "\n\t"}
    if (true || (float32(-2.25) > 1e10)) {
        v10 := GenS1{F0: true, F1: "a"}
        v11 := []int{32768}
        var v12 uint32 = 0
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{-467, 1}
    v14 := GenS1{F0: false, F1: "hello"}
    fmt.Println(v1, v3)
    var v15 int8 = 127
    var v16 int64 = 302
    fmt.Println(v16, v3, v15)
    var v17 float32 = ((3.14159 * 0.0) / 1.0)
    var v18 float64 = 0.1
    v15 -= 104
    v18 = v18
    switch 32768 {
    case 4294967295:
        var v19 bool = (v15 <= v15)
        _ = v19
    case -928:
        _ = len(v13)
    case 981:
        var v20 int = 521
        _ = v20
    default:
        v21 := GenS1{F0: true, F1: "\n\t"}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
