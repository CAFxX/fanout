package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint }

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = -1
    v3 := []string{("\n\t" + ("foo bar" + ("\n\t" + (("" + "\n\t") + "\n\t")))), "hello"}
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint16 = 547
        var v8 uint16 = v7
        var v9 string = ""
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int8 = (-(v2))
    v2 = v10
    var v11 uint8 = 2
    fmt.Println(v1, v11)
    v11 += v11
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: ((v1 >> ((2) & 63)) < v1), F1: 0}
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 4294967295, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
