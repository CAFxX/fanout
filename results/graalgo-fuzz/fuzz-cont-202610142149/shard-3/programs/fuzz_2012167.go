package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint32 = 0
    var v2 float32 = float32(0.1)
    var v3 uint8 = 255
    v4 := GenS1{F0: 2, F1: v1}
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int32 = -2147483648
        _ = fi5
        _ = v6
    }
    fmt.Println(v1, v2, v3)
    v7 := []int{42, 32768, 4294967295, 42}
    var v8 string = ("" + "\n\t")
    for fi9 := 0; fi9 < 9; fi9++ {
        v8 = (v8 + (v8 + v8))
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = len(v7)
        _ = fi9
    }
    fmt.Println(v3, v2)
    switch "hello" {
    case "hello":
        v1 -= (v1 | 1)
    }
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
