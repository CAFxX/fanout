package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
func genFn0(p0 int64, p1 int8) int16 {
    v7 := []int{677}
    p0 -= 9223372036854775807
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p1 += 42
    _ = p0
    _ = p1
    _ = v7
    return int16(p0)
}

func main() {
    var v1 float32 = ((-2.25 + (((float32((1e10 - 0.0)) / 3.14159) / 3.14159) / 1.0)) / (0.1 - 1e10))
    v2 := []uint32{}
    var v3 float32 = (1e10 * v1)
    var v4 uint32 = 65535
    v5 := []string{}
    var v6 int32 = -1
    v8 := GenS1{F0: "\n\t", F1: 4294967295, F2: -28}
    v6 -= 2147483647
    var v9 int = 32768
    v5 = append(v5, "hello")
    v3 = v3
    v10 := GenS1{F0: "foo bar", F1: v9, F2: -84}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "" {
    case "hello":
        var v11 uint = 2
        _ = v11
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        v6 = 0
    }
    switch "hello" {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = (("hello" + ("foo bar" + "hello")) + "foo bar")
        }
    case "a":
        var v12 uint = 42
        _ = v12
    default:
        switch ((("foo bar" + "foo bar") + "hello") + ("foo bar" + ("" + "\n\t"))) {
        case "hello":
            var v13 int32 = (v6 % ((-438) | 1))
            _ = v13
        case "\n\t":
            v1 = v3
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
