package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    v1 := []uint32{2, 2147483648, 65535}
    var v2 float64 = 1.5
    v3 := make(map[uint32]int)
    if len(v1) > 0 {
        v1[0] = 42
    }
    v4 := v3[2147483648]
    fmt.Println(v2)
    v5 := GenS1{F0: -1, F1: false, F2: 42}
    var v6 uint64 = 797
    if false {
        fmt.Println(v6)
    }
    switch "" {
    case "\n\t":
        fmt.Println(v6, v4, v2)
    case "foo bar":
        v6 -= 2
    case "a":
        v7 := GenS1{F0: 1, F1: ((("\n\t" + "hello") + "a") >= "\n\t"), F2: 1}
        _ = v7
    default:
        var v8 int32 = -651
        _ = v8
    }
    v1 = append(v1, 301)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int64 = 9223372036854775807
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: (v4 >> ((4294967295) & 63)), F1: 100}
    v12 := GenS1{F0: -600, F1: true, F2: 37}
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
