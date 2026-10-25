package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 string = (("foo bar" + ("" + "")) + (("a" + "foo bar") + "hello"))
    v4 := make(map[string]string)
    var v5 bool = (!(127 < 100))
    v6 := []int{517}
    v7 := GenS1{F0: 723, F1: 42, F2: v3}
    switch -693 {
    case 42:
        v3 = v3
    case 128:
        var v8 int = 127
        _ = v8
    case -427:
        fmt.Println(v5, v5)
    }
    v9 := GenS2{F0: 4294967295, F1: 79}
    v10 := GenS2{F0: 334, F1: 2147483647}
    fmt.Println(v5, v5)
    var v11 uint32 = 42
    var v12 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
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
    _ = v9
}
