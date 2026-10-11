package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int64; F2 uint }

func main() {
    var v1 float32 = 1e10
    var v2 uint = 18446744073709551615
    var v3 bool = (!false)
    v4 := make(map[uint32]int)
    if v3 {
        v1 = float32(0.1)
    }
    if v3 {
        var v5 int16 = 42
        var v6 uint = 867
        switch 0 {
        case -9223372036854775808:
            v6 = (v6 + (v2 - (v2 << ((v6) & 63))))
        }
        _ = v5
        _ = v6
    } else {
        var v7 string = ("" + "hello")
        var v8 float32 = 0.0
        _ = v7
        _ = v8
    }
    if v3 {
        var v9 uint64 = 957
        _ = v9
    }
    switch -1 {
    case 0:
        var v10 string = "foo bar"
        _ = v10
    case 32768:
        v11 := GenS1{F0: "\n\t"}
        _ = v11
    case 4294967295:
        v2 += v2
    }
    _ = len(v4)
    v3 = true
    v2 = (^(v2))
    v12 := GenS1{F0: ("" + ((("a" + "hello") + "hello") + ("a" + "hello")))}
    v3 = v3
    v13 := GenS1{F0: (("hello" + "") + "\n\t")}
    fmt.Println(v1)
    fmt.Println(v1, v2, v1)
    v14 := GenS2{F0: 698, F1: 1, F2: (v2 & (v2 / (((v2 << (((v2 - 546)) & 63))) | 1)))}
    var v15 float32 = v1
    v4[4294967295] = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
