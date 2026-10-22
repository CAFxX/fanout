package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 uint }

func main() {
    var v1 int64 = 42
    var v2 float32 = 0.0
    v3 := make(map[uint32]int)
    var v4 uint64 = 4294967295
    for fi5 := 0; fi5 < 19; fi5++ {
        delete(v3, 65535)
        _ = fi5
    }
    v6 := GenS1{F0: (147 >= 0)}
    v7 := GenS1{F0: true}
    v8 := GenS1{F0: false}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 float64 = 3.14159
        fmt.Println(v4)
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v2)
    v2 = v2
    v11 := GenS2{F0: "foo bar", F1: 157}
    fmt.Println(v2, v4)
    v1 = v1
    v12 := GenS3{F0: 343, F1: 839}
    v13 := GenS2{F0: ("\n\t" + "hello"), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
