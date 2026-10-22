package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int64 = 42
    var v2 bool = false
    v3 := []string{"foo bar", "\n\t", ("foo bar" + "foo bar")}
    var v4 float64 = 1.5
    var v5 int16 = -32768
    var v6 float32 = 0.0
    var v7 int16 = -1
    v4 -= (-2.25 * v4)
    v3 = append(v3, "a")
    v8 := GenS1{F0: v7, F1: (1 < 127), F2: 65535}
    var v9 int = 32767
    v2 = v2
    var v10 int32 = 2147483647
    var v11 float64 = (1.5 / 1e10)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 := GenS2{F0: 0}
    if true {
        var v13 bool = (v2 && (("hello" + ("foo bar" + "foo bar")) != "hello"))
        v1 = v1
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
