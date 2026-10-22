package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int16 }

func main() {
    v1 := []int{128}
    var v2 int8 = -128
    var v3 uint64 = 2
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    v2 = 42
    if (((1.5 - 0.1) != (1e10 * (-2.25 * (-2.25 - -2.25)))) && false) {
        switch ("foo bar" + "a") {
        case "foo bar":
            var v7 string = (((("hello" + "foo bar") + "\n\t") + ("" + "hello")) + "hello")
            _ = v7
        default:
            _ = len(v1)
        }
        var v8 float64 = -2.25
        _ = v8
    }
    var v9 uint8 = 42
    v10 := GenS1{F0: -39, F1: (-(v2)), F2: 1}
    fmt.Println(v9, v9)
    var v11 uint32 = 4294967295
    for fi12 := 0; fi12 < 19; fi12++ {
        v11 = (v11 - v11)
        var v13 uint16 = 262
        _ = fi12
        _ = v13
    }
    var v14 uint8 = (v9 / ((v9) | 1))
    var v15 int32 = -800
    var v16 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
