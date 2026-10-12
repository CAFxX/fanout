package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    var v1 int = 1
    var v2 uint32 = 1
    var v3 int16 = 0
    var v4 float32 = 0.1
    var v5 uint64 = 4294967295
    var v6 int16 = (v3 | 32767)
    if (((178 >= (v6 >> ((255) & 63))) || false) || true) {
        fmt.Println(v1, v4, v2)
    }
    var v7 uint8 = 127
    v8 := []int{(v1 % ((v1) | 1)), -9223372036854775808, int(v2), (v1 % (((v1 / (((v1 + (v1 ^ v1))) | 1))) | 1))}
    _ = len(v8)
    switch v1 {
    case -862:
        var v9 string = (("\n\t" + ("foo bar" + ("\n\t" + "a"))) + ("\n\t" + "a"))
        _ = v9
    default:
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 float32 = ((v4 * float32(0.1)) / (v4 / 3.14159))
            _ = fi10
            _ = v11
        }
    }
    v2 += v2
    v4 += float32(0.0)
    v12 := GenS1{F0: -579, F1: -174}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
