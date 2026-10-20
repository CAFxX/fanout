package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    v1 := []int{}
    var v2 uint32 = 617
    v3 := []string{""}
    var v4 uint16 = 696
    v2 -= v2
    if ((-1 == 0) || true) {
        v4 = v4
    }
    v5 := GenS1{F0: 100, F1: 2}
    var v6 int8 = 42
    fmt.Println(v6, v4, v6)
    v7 := GenS1{F0: (-(v6)), F1: 255}
    v8 := GenS1{F0: -128, F1: 255}
    if false {
        switch "foo bar" {
        case "":
            var v9 float32 = 0.0
            _ = v9
        case "\n\t":
            _ = len(v1)
        default:
            var v10 int16 = -1
            _ = v10
        }
        fmt.Println(v4, v4)
    } else {
        switch ((((("\n\t" + ((("hello" + "\n\t") + "a") + "")) + "") + "") + "") + "\n\t") {
        case "hello":
            _ = len(v3)
        }
    }
    v4 = (v4 & (v4 | v4))
    switch "hello" {
    case "a":
        var v11 int8 = v6
        _ = v11
    }
    var v12 string = ""
    v1 = append(v1, 4294967295)
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int64 = 677
        _ = fi13
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
