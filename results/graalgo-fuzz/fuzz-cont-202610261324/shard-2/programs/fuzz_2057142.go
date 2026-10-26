package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }

func main() {
    var v1 float64 = 0.0
    var v2 float32 = 0.1
    v3 := make(map[int]string)
    v4 := []uint32{253, 2}
    var v5 int16 = -32768
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v1, v5)
    case "hello":
        v5 = v5
    default:
        var v6 int16 = 531
        _ = v6
    }
    v2 -= 0.0
    switch 0 {
    case 755:
        v7 := GenS1{F0: -9223372036854775808, F1: 42}
        _ = v7
    }
    if true {
        fmt.Println(v2, v5)
        v8 := GenS1{F0: 9223372036854775807, F1: 509}
        v9 := GenS2{F0: 330}
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 42
        v10 -= 0
        _ = v10
    }
    var v11 int = 127
    v4 = append(v4, 1)
    for fi12 := 0; fi12 < 4; fi12++ {
        _ = len(v4)
        _ = fi12
    }
    var v13 int64 = -1
    v14 := GenS3{F0: v5, F1: v5, F2: 205}
    var v15 uint32 = 933
    var v16 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
