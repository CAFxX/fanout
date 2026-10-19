package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := []int{}
    v2 := []int{42, 0}
    var v3 int = 4294967295
    v4 := make(map[string]bool)
    var v5 float32 = ((0.0 - 1e10) / -2.25)
    delete(v4, "a")
    var v6 int16 = 32767
    var v7 int32 = 331
    var v8 int16 = v6
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 string = "hello"
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 16
    var v12 int8 = 127
    v8 -= 42
    switch v3 {
    case 42:
        v13 := GenS1{F0: v11, F1: (v11 * v11)}
        _ = v13
    }
    fmt.Println(v11)
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
