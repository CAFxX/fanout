package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (^(v1))
    var v3 uint32 = 2147483648
    v2 = 0
    v2 = v1
    v4 := GenS1{F0: v3, F1: -2147483648}
    v1 += 4294967295
    var v5 float64 = 1e10
    var v6 float32 = 0.1
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v2 += v2
            _ = fi8
        }
        _ = fi7
    }
    v3 = 0
    switch "" {
    case "":
        v9 := GenS2{F0: (42 == 2147483647), F1: ((!true) && ("" >= ("foo bar" + (((("foo bar" + "") + "hello") + ("a" + "")) + "a"))))}
        _ = v9
    case "hello":
        var v10 bool = ((float64(0.1) <= (1e10 - (v5 * 1e10))) || false)
        _ = v10
    case "a":
        v11 := []int{42, 32768, 0}
        _ = v11
    }
    v12 := GenS3{F0: v3, F1: -128}
    v13 := GenS2{F0: false, F1: ((!false) && false)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
