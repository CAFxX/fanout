package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]string)
    v3 := []uint32{0, 202, 0, 4294967295}
    switch "hello" {
    case "\n\t":
        var v4 bool = (true || true)
        _ = v4
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: -465, F1: "", F2: 109}
        _ = len(v2)
        _ = fi5
        _ = v6
    }
    _ = len(v2)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v7 int64 = 575
    var v8 string = "foo bar"
    v9 := GenS1{F0: (v1 % ((380) | 1)), F1: "hello", F2: 81}
    fmt.Println(v8, v8, v8)
    if (!false) {
        fmt.Println(v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
