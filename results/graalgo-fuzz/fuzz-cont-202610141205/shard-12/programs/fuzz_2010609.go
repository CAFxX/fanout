package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 float64 = 3.14159
    v4 := make(map[int]uint32)
    v5 := []int{2147483647}
    var v6 int16 = 783
    var v7 int = 2147483647
    v7 = v7
    v8 := GenS1{F0: -1, F1: 2, F2: 3}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 += v3
    switch ((("hello" + ((("\n\t" + (("hello" + ("foo bar" + ("\n\t" + "a"))) + (("foo bar" + "hello") + (("hello" + "") + "hello")))) + "hello") + ("foo bar" + "\n\t"))) + "") + "hello") {
    case "hello":
        var v9 string = ("" + "")
        _ = v9
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            v3 += v3
            _ = fi10
        }
    }
    v6 += 32767
    v6 = -419
    if len(v5) > 0 {
        v5[0] = (^(v7))
    }
    v3 -= (v3 - v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
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
}
