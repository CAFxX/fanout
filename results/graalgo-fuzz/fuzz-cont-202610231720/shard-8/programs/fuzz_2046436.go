package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 42
    v3 := make(map[string]uint32)
    var v4 float32 = 0.0
    var v5 string = (("a" + "hello") + ("hello" + ("\n\t" + "")))
    v6 := []int{127}
    var v7 uint8 = 188
    fmt.Println(v4, v5, v5)
    v8 := GenS1{F0: ((("hello" + v5) + ("a" + v5)) + v5)}
    v9 := GenS1{F0: "hello"}
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    if (9223372036854775807 == -992) {
        v10 := GenS1{F0: (v5 + v5)}
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int16 = 1
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: true, F1: 32767}
    switch (v5 + v5) {
    case "foo bar":
        v2 += v2
    default:
        var v14 uint32 = uint32(v2)
        _ = v14
    }
    var v15 uint16 = 346
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
