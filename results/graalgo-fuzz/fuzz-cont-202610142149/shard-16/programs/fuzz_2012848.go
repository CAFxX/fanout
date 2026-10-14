package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    var v1 int8 = -1
    var v2 int = 32767
    v3 := make(map[string]uint32)
    v4 := []uint32{0, 2147483647}
    v5 := []string{"foo bar", "\n\t"}
    fmt.Println(v2)
    v1 = (v1 ^ v1)
    var v6 int = (-(v2))
    if true {
        v7 := GenS1{F0: -32768}
        v1 -= v1
        _ = v7
    }
    v5 = append(v5, "foo bar")
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    switch ("foo bar" + "") {
    case "":
        fmt.Println(v2, v6)
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = 0
        }
    case "a":
        v2 = v2
    default:
        v8 := GenS2{F0: 922, F1: 1}
        _ = v8
    }
    v9 := v3["\n\t"]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
