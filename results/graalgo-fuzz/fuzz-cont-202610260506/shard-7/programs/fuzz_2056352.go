package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
func genFn0(p0 int32, p1 int, p2 int8) string {
    var v5 int32 = p0
    p0 -= (p0 / ((-2147483648) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "foo bar"
}
func genFn1(p0 int, p1 int32, p2 uint64) int32 {
    var v6 int32 = -2147483648
    var v7 int8 = 104
    if (uint32(p2) > 2147483648) {
        var v8 int8 = v7
        p1 = (-(v6))
        _ = v8
    } else {
        var v9 uint8 = 127
        for fi10 := 0; fi10 < 18; fi10++ {
            v9 += (v9 + 2)
            fmt.Println(p0)
            var v11 float64 = (((1e10 + -2.25) - 0.1) / (((0.0 + 0.0) / (1e10 / 1e10)) + (1e10 * 3.14159)))
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := []int{-9223372036854775808, 32768, 2147483647, 128}
    v2 := make(map[uint32]int)
    var v3 string = ((("hello" + ("a" + "foo bar")) + ("hello" + "\n\t")) + ("" + ""))
    var v4 int64 = 42
    var v12 uint64 = 169
    v13 := GenS1{F0: 255, F1: (("foo bar" + v3) + "a"), F2: 128}
    v14 := GenS1{F0: 42, F1: v3, F2: 127}
    switch v3 {
    case "":
        var v15 uint32 = 372
        _ = v15
    case "\n\t":
        v3 = ""
    case "foo bar":
        _ = len(v1)
    default:
        switch "a" {
        case "hello":
            var v16 int32 = genFn1(42, 2147483647, (v12 ^ v12))
            _ = v16
        case "":
            var v17 uint = 18446744073709551615
            _ = v17
        }
    }
    if (-2.25 < ((((-2.25 * 0.1) + 1e10) - 1.5) / 1.5)) {
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v12)
            _ = fi18
        }
        var v19 uint64 = v12
        var v20 int8 = 0
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 250, F1: "foo bar", F2: 164}
    _ = len(v1)
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := GenS1{F0: 43, F1: "\n\t", F2: 147}
        _ = fi22
        _ = v23
    }
    var v24 int32 = -177
    var v25 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
