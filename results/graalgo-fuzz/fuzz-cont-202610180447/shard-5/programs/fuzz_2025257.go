package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }

func main() {
    var v1 bool = false
    var v2 uint64 = 42
    var v3 int8 = -1
    v3 = v3
    v4 := []int{32767, 127, 4294967295}
    v3 = (v3 * 127)
    var v5 float32 = (0.1 - 0.0)
    var v6 uint8 = 58
    var v7 float64 = 1.5
    var v8 bool = true
    _ = len(v4)
    switch "" {
    case "hello":
        v8 = (("\n\t" + "a") == "foo bar")
    }
    var v9 uint8 = 42
    v10 := GenS1{F0: ("a" + "\n\t"), F1: 1, F2: 0}
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v8, v3)
        v8 = v1
        _ = fi11
    }
    v8 = true
    var v12 uint64 = 42
    v8 = false
    var v13 int16 = -962
    var v14 int32 = 42
    switch ((("hello" + "hello") + (("" + "a") + "hello")) + ("a" + "hello")) {
    case "foo bar":
        v13 -= int16(v9)
    case "\n\t":
        if true {
            if v8 {
                var v15 int64 = -395
                _ = v15
            }
            v5 = ((((v5 + (1e10 * (float32(v7) * float32(((v7 + (1e10 - -2.25)) / (1e10 + 1e10)))))) * (float32(0.1) * v5)) - -2.25) / v5)
        } else {
            fmt.Println(v13, v5, v2)
            var v16 float64 = float64(1.5)
            _ = v16
        }
    default:
        var v17 string = "\n\t"
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
