package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    v4 := []int{42, 127}
    v3 = (-(v3))
    var v5 string = "hello"
    switch "hello" {
    case "hello":
        var v6 int8 = -38
        _ = v6
    default:
        v5 = v5
    }
    switch ("hello" + (v5 + (v5 + v5))) {
    case "hello":
        v7 := GenS1{F0: v3, F1: 777}
        _ = v7
    default:
        var v8 int32 = 2147483647
        _ = v8
    }
    switch "a" {
    case "a":
        v3 = 9223372036854775807
    case "":
        for fi9 := 0; fi9 < 12; fi9++ {
            v5 = "foo bar"
            _ = fi9
        }
    default:
        v10 := GenS2{F0: 42, F1: (!v1)}
        _ = v10
    }
    v5 = "hello"
    var v11 int8 = -32
    switch (("\n\t" + "hello") + "foo bar") {
    case "foo bar":
        switch (v5 + v5) {
        case "":
            v1 = (3.14159 < (((0.1 + 1.5) * (1e10 - 0.0)) + 3.14159))
        default:
            _ = len(v2)
        }
    case "a":
        fmt.Println(v5, v1, v1)
    }
    fmt.Println(v11, v11)
    v5 = (("a" + "hello") + (("\n\t" + (v5 + v5)) + ("foo bar" + ("\n\t" + ("\n\t" + "")))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
