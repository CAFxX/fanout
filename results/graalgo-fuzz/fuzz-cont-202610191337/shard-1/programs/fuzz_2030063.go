package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 2
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 float32 = ((0.0 / 0.1) * 0.1)
    v6 := []string{}
    v7 := GenS1{F0: 128, F1: 0, F2: ("" + "foo bar")}
    switch ("\n\t" + "") {
    case "hello":
        v8 := GenS1{F0: 127, F1: 608, F2: ""}
        _ = v8
    default:
        v2 = v2
    }
    v9 := GenS1{F0: 2147483647, F1: 2, F2: "a"}
    v5 += v5
    var v10 uint8 = 2
    fmt.Println(v10, v2)
    if len(v6) > 0 {
        v6[0] = ""
    }
    v11 := GenS2{F0: 2147483647, F1: v10, F2: 60}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int32 = 0
        v2 += (v2 % ((2) | 1))
        _ = fi12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
