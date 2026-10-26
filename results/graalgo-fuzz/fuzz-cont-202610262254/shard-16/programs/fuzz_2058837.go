package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    if true {
        v5 := GenS1{F0: 255}
        _ = v5
    }
    var v6 uint16 = 42
    v6 = v6
    v7 := GenS1{F0: 128}
    v6 += 32767
    if true {
        var v8 bool = true
        v8 = (4294967295 < 861)
        _ = v8
    }
    v9 := GenS1{F0: 128}
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
