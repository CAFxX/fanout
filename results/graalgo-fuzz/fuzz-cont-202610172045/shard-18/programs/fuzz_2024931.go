package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 string }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    switch 42 {
    case 0:
        fmt.Println(v1)
    default:
        if (0 < 32767) {
            fmt.Println(v1)
        }
    }
    v2[("\n\t" + ("foo bar" + "\n\t"))] = "foo bar"
    var v4 uint8 = 42
    var v5 uint8 = (^(v4))
    v6 := []int{-197}
    _ = len(v2)
    var v7 uint32 = 1
    switch 0 {
    case -568:
        v8 := GenS1{F0: 1, F1: 0}
        _ = v8
    }
    var v9 uint64 = 2
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS2{F0: -1, F1: 4294967295, F2: ("\n\t" + "a")}
        _ = fi10
        _ = v11
    }
    v7 += (v7 | 2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
