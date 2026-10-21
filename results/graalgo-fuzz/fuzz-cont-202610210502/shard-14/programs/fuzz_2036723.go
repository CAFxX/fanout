package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 1
    v3 := []string{"a"}
    var v4 int32 = 42
    fmt.Println(v4)
    v5 := GenS1{F0: 2, F1: 1}
    var v6 int32 = v4
    v2 = v2
    v2 += v2
    v4 = 995
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 string = ("a" + "a")
        _ = fi7
        _ = v8
    }
    if true {
        var v9 uint = 65535
        _ = v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
