package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []string{"a", "foo bar"}
    var v2 uint = 1
    var v3 uint = (^(v2))
    v4 := make(map[string]uint32)
    var v5 uint = (^(v3))
    v2 -= (v5 & 723)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 := GenS1{F0: 127, F1: -205}
    var v7 int32 = -869
    fmt.Println(v2, v2, v2)
    v3 = v2
    v3 = 18446744073709551615
    var v8 bool = (-2.25 < 3.14159)
    v5 += 255
    var v9 int8 = -1
    v10 := GenS2{F0: v8}
    fmt.Println(v7)
    v7 += (v7 / ((v7) | 1))
    for fi11 := 0; fi11 < 1; fi11++ {
        switch ("hello" + ((("a" + "hello") + "hello") + "hello")) {
        case "a":
            v3 = v3
        case "\n\t":
            v12 := GenS3{F0: v7, F1: 1}
            _ = v12
        default:
            v7 = v7
        }
        var v13 int16 = -1
        _ = fi11
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 int16 = -354
        v16 := GenS4{F0: 255}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 string = ("hello" + ("hello" + "foo bar"))
    switch "a" {
    case "\n\t":
        v2 += 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
