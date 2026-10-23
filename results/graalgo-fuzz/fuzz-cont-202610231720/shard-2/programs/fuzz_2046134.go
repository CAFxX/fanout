package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }

func main() {
    var v1 uint = 42
    var v2 uint16 = 32767
    var v3 int16 = 1
    var v4 uint8 = 0
    v5 := []string{(("foo bar" + "") + "\n\t"), "hello", ("a" + "")}
    v3 = (v3 + 1000)
    var v6 int64 = 42
    v7 := GenS1{F0: v3, F1: (^(v4))}
    var v8 int = -148
    v4 += v4
    fmt.Println(v6, v2, v8)
    var v9 string = ""
    v8 = 32768
    var v10 int64 = v6
    var v11 int64 = v10
    var v12 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
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
