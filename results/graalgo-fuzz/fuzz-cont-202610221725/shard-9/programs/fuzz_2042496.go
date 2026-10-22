package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint }
func genFn0(p0 int64, p1 int64, p2 int8) int32 {
    p1 += (-(p0))
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 += p0
        var v8 float64 = (-2.25 / 1.0)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = (0.0 * 0.0)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 408
    v3 := make(map[int]bool)
    var v4 uint8 = 42
    var v5 int64 = 978
    var v6 bool = ((3.14159 - 0.0) > (1e10 + 1e10))
    v11 := v3[-287]
    v12 := GenS1{F0: 42}
    var v13 float64 = 1.5
    fmt.Println(v11)
    switch "foo bar" {
    case "a":
        fmt.Println(v1, v4)
    case "\n\t":
        fmt.Println(v11, v6, v6)
    case "hello":
        fmt.Println(v6, v6)
    }
    v14 := GenS1{F0: 1}
    v15 := GenS2{F0: 1, F1: -711, F2: 255}
    v4 = v4
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint64 = 22
            var v19 string = ""
            var v20 int32 = -77
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 float64 = 3.14159
        var v22 int32 = genFn0((v5 << ((4294967295) & 63)), v5, 127)
        _ = fi16
        _ = v21
        _ = v22
    }
    v6 = (1 > 0)
    v1 = (v1 & v5)
    var v23 int32 = 42
    delete(v3, 672)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
