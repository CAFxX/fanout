package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{170, 65535, 889}
    v3 := make(map[string]int)
    v4 := []int{127}
    _ = len(v2)
    v5 := GenS1{F0: 65535, F1: -453}
    var v6 uint32 = 646
    var v7 int = 42
    v8 := GenS1{F0: 255, F1: -2147483648}
    delete(v3, "")
    v9 := GenS1{F0: 4294967295, F1: 0}
    if true {
        v10 := v3[(("foo bar" + "a") + ("\n\t" + ""))]
        _ = v10
    }
    v4 = append(v4, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
