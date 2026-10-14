package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }

func main() {
    var v1 int = 32767
    var v2 float64 = float64(1e10)
    v3 := []uint32{65535}
    var v4 bool = ((!(42 > 42)) && (396 <= 2))
    v5 := GenS1{F0: "hello", F1: (("" + ((("\n\t" + "a") + ((("a" + "\n\t") + ("hello" + "hello")) + "a")) + ("" + "hello"))) + "foo bar"), F2: 0}
    fmt.Println(v4)
    v2 += (((v2 * (v2 / (((v2 * 1e10) / float64(((0.0 + float32(v2)) + 0.0))) + v2))) + v2) * ((1e10 / float64(1e10)) * (0.0 - 3.14159)))
    var v6 int32 = 2147483647
    var v7 int8 = -1
    var v8 int64 = 42
    var v9 uint8 = 255
    fmt.Println(v7, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
