package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"", "hello", (("foo bar" + ("hello" + ("hello" + ("a" + "foo bar")))) + "a")}
    v3 := make(map[int]bool)
    var v4 string = ""
    v5 := []uint32{1, 684, 971}
    v6 := []uint32{4294967295}
    fmt.Println(v4)
    _ = len(v1)
    fmt.Println(v4)
    if (!false) {
        v4 = "foo bar"
    }
    var v7 int16 = -1
    v7 = (^(v7))
    v8 := v1[v4]
    var v9 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: v7, F1: 65535, F2: 255}
        fmt.Println(v4, v9)
        _ = fi10
        _ = v11
    }
    v8 = (v8 - 0)
    fmt.Println(v7)
    var v12 int16 = (^(v7))
    fmt.Println(v12, v9, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
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
