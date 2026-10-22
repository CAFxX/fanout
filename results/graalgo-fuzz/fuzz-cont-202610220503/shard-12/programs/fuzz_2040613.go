package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint64; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 uint = 184
    v4 := make(map[string]int)
    var v5 int16 = -1
    v6 := GenS1{F0: false}
    v5 -= 0
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint8 = 0
        fmt.Println(fi7, v8)
        var v9 int8 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    v2[2147483647] = ""
    var v10 float32 = (3.14159 * (1e10 + 1e10))
    switch 66 {
    case 127:
        v11 := v4[("hello" + "hello")]
        _ = v11
    }
    v12 := GenS2{F0: false, F1: 42, F2: 4294967295}
    switch -463 {
    case 514:
        v13 := GenS2{F0: true, F1: 398, F2: 455}
        _ = v13
    case 654:
        _ = len(v4)
    default:
        v14 := GenS1{F0: ((("" + "") + ("foo bar" + "foo bar")) > "")}
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
