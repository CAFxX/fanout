package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
func genFn0(p0 int8) int16 {
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p0)
    default:
        var v4 uint8 = 255
        _ = v4
    }
    if ((((("" + "foo bar") + "hello") + ("foo bar" + "hello")) + ("" + "a")) < "\n\t") {
        v5 := GenS1{F0: -128, F1: 2147483647}
        _ = v5
    } else {
        fmt.Println(p0)
    }
    var v6 int16 = 1
    v6 += (v6 & -201)
    _ = p0
    _ = v6
    return (v6 >> ((4294967295) & 63))
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 string = (("a" + "foo bar") + "hello")
    fmt.Println(v3)
    v3 = "foo bar"
    var v7 int16 = 32767
    var v8 float32 = -2.25
    fmt.Println(v8, v3)
    v7 += genFn0(1)
    v8 = float32(0.0)
    v3 = ("hello" + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
