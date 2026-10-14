package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }

func main() {
    var v1 int8 = -1
    var v2 uint8 = 127
    var v3 float32 = 1.5
    v4 := GenS1{F0: "a", F1: 0, F2: v1}
    for fi5 := 0; fi5 < 15; fi5++ {
        v2 = v2
        v6 := GenS1{F0: "\n\t", F1: (^(fi5)), F2: v1}
        _ = fi5
        _ = v6
    }
    v1 = 10
    var v7 uint = 635
    v1 = v1
    fmt.Println(v3, v1)
    v8 := GenS1{F0: "", F1: 32767, F2: 1}
    v7 += (v7 & 42)
    var v9 uint = v7
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
