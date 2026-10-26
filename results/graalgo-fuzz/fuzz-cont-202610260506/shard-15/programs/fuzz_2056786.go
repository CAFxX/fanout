package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }

func main() {
    var v1 int64 = -784
    var v2 float32 = 1e10
    v3 := make(map[string]bool)
    var v4 uint16 = 2
    var v5 uint64 = 0
    v6 := []int{496, -490, -627}
    var v7 int = 486
    var v8 string = ("\n\t" + "\n\t")
    fmt.Println(v2, v5)
    v8 = ((("\n\t" + ("foo bar" + "foo bar")) + v8) + "\n\t")
    fmt.Println(v8)
    v7 += v7
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v9 := GenS1{F0: 255, F1: v4, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
