package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }
func genFn0(p0 uint, p1 uint32) string {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int = (fi7 & 0)
        _ = fi7
        _ = v8
    }
    fmt.Println(p1)
    if ((1 != p0) || ((1 >= 1) && false)) {
        p1 += (p1 * (p1 - (p1 + (p1 ^ 2))))
        v9 := GenS1{F0: -1, F1: 1}
        _ = v9
    }
    _ = p0
    _ = p1
    return (("a" + "foo bar") + (("\n\t" + "foo bar") + "a"))
}

func main() {
    var v1 float64 = 0.1
    var v2 int = 32767
    v3 := make(map[string]string)
    v4 := []uint32{4294967295, 42}
    v5 := []int{-2147483648, -586, 1}
    var v6 int16 = 330
    switch v2 {
    case -584:
        v6 -= (v6 ^ (v6 * v6))
    case -9223372036854775808:
        fmt.Println(v1, v1)
    }
    switch ("hello" + ("hello" + "\n\t")) {
    case "foo bar":
        var v10 int16 = (v6 % ((0) | 1))
        _ = v10
    case "hello":
        var v11 int8 = 42
        _ = v11
    }
    var v12 uint64 = 147
    var v13 uint16 = 1
    var v14 uint32 = 42
    fmt.Println(v1, v2)
    fmt.Println(v12)
    v12 = v12
    v15 := GenS1{F0: -9223372036854775808, F1: -1}
    _ = len(v3)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint8 = 103
        fmt.Println(v1, v1, v13)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint = 42
        v2 -= 309
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: (v14 + (^(v14))), F1: -243, F2: uint32(v12)}
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 float32 = (float32(1.5) * ((-2.25 - ((1e10 + 3.14159) / 1.0)) / (1.5 + 3.14159)))
        v22 = (v22 + 0.1)
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
