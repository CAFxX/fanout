package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1000
    v3 := []int{int(v2), -71}
    var v4 uint64 = 937
    var v5 float64 = float64((-2.25 + (0.1 * (3.14159 / 1.0))))
    if len(v3) > 0 {
        v3[0] = 32768
    }
    v6 := GenS1{F0: -933, F1: 32767}
    var v7 int = 1
    if true {
        fmt.Println(v5, v4, v4)
        var v8 float32 = -2.25
        _ = len(v1)
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if (!(!((fi9 >> ((1) & 63)) <= -2147483648))) {
            v5 += float64((float32(v5) - 0.1))
            v2 = (v2 % ((0) | 1))
        }
        v4 = 185
        _ = fi9
    }
    var v10 string = (("a" + ("foo bar" + ("" + (("a" + ("" + ("\n\t" + "hello"))) + ("foo bar" + ("a" + "a")))))) + "")
    v10 = (v10 + v10)
    var v11 uint8 = 2
    v5 += 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
