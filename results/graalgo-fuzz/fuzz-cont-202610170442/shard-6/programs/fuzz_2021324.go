package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int32; F1 string; F2 int32 }
func genFn0(p0 int32) int {
    var v7 uint8 = 0
    var v8 uint = 0
    if (!((1 <= 42) || (32768 < -577))) {
        p0 = (-(p0))
    }
    v9 := GenS1{F0: 1, F1: 1}
    v10 := GenS1{F0: 42, F1: 9223372036854775807}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -683
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 bool = (0.1 < 3.14159)
    var v4 float64 = (-2.25 - (1.5 * 1e10))
    var v5 int8 = -1
    var v6 int64 = 1
    fmt.Println(v4)
    if false {
        if false {
            var v11 int8 = -128
            delete(v2, "foo bar")
            v11 = (-(v5))
            _ = v11
        }
    }
    switch -1 {
    case 534:
        v12 := GenS1{F0: 423, F1: v6}
        _ = v12
    case 0:
        v4 = v4
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        switch 128 {
        case 811:
            fmt.Println(v3, fi13)
        }
        _ = fi13
    }
    v14 := GenS2{F0: 258}
    var v15 float64 = v4
    v16 := GenS1{F0: 447, F1: 42}
    v15 += (v4 + (v4 + float64(((0.1 - 1.5) - float32((1e10 + 0.0))))))
    v15 = ((v4 - v15) - v4)
    v17 := GenS3{F0: 116, F1: 65535}
    v18 := GenS1{F0: 2, F1: v6}
    if (1 != 1) {
        fmt.Println(v3, v4)
        fmt.Println(v5, v15)
        switch ("foo bar" + "") {
        case "\n\t":
            fmt.Println(v5)
        default:
            v19 := GenS1{F0: 908, F1: v6}
            _ = v19
        }
    }
    if (1 == 1) {
        if (true && false) {
            v3 = true
            var v20 float64 = float64(float32(v15))
            fmt.Println(v3)
            _ = v20
        }
        v21 := GenS4{F0: 42, F1: "hello", F2: -698}
        v3 = (1 < 0)
        _ = v21
    } else {
        v4 = (v15 - 3.14159)
    }
    v6 += 9223372036854775807
    fmt.Println(v3, v6)
    var v22 uint64 = 259
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
