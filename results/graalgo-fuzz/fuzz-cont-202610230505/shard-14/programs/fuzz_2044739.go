package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 bool = (v1 && v1)
    v3 = true
    fmt.Println(v1, v1)
    v4 := GenS1{F0: 567, F1: 154, F2: -1}
    fmt.Println(v3)
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 65535, F1: 79, F2: 0}
        _ = fi5
        _ = v6
    }
    v2[2147483647] = (true && v3)
    v2[2147483647] = true
    v7 := GenS1{F0: 0, F1: 65535, F2: -128}
    v8 := GenS1{F0: 543, F1: 797, F2: 114}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
