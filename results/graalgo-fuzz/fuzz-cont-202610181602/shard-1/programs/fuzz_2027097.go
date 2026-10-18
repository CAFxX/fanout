package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }

func main() {
    var v1 uint32 = 578
    v2 := make(map[string]string)
    var v3 uint64 = 246
    v4 := GenS1{F0: true, F1: true}
    _ = len(v2)
    fmt.Println(v1)
    v1 += v1
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int16 = -983
        var v7 float32 = 0.0
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: true, F1: true}
    v2["a"] = ""
    v9 := GenS2{F0: -804, F1: v1, F2: -1}
    var v10 string = ("foo bar" + ("foo bar" + ("hello" + "hello")))
    v1 = 1
    fmt.Println(v1, v1, v1)
    for fi11 := 0; fi11 < 14; fi11++ {
        v10 = (((("" + v10) + v10) + (v10 + v10)) + "hello")
        var v12 int32 = 42
        var v13 uint8 = 42
        _ = fi11
        _ = v12
        _ = v13
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
