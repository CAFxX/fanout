package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    var v1 uint = 168
    var v2 int16 = 1000
    v3 := []string{("a" + (("foo bar" + ("foo bar" + (("hello" + "hello") + "hello"))) + "a")), "", ((("\n\t" + ("" + "hello")) + "hello") + "foo bar")}
    var v4 int64 = 0
    var v5 int8 = -128
    v6 := []string{}
    v7 := GenS1{F0: (("foo bar" + "hello") + "hello"), F1: 0}
    var v8 uint32 = 357
    var v9 uint8 = 128
    fmt.Println(v5, v1, v2)
    var v10 uint8 = (v9 * (v9 / ((127) | 1)))
    switch "foo bar" {
    case "hello":
        v10 += (^(v9))
    }
    var v11 float64 = 1.5
    v2 += 42
    v11 += v11
    var v12 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
