package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []string{(((("\n\t" + ("a" + "")) + "foo bar") + "") + ("\n\t" + "foo bar")), "\n\t", "hello"}
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 float64 = (0.1 + (0.0 * (0.0 + -2.25)))
    v6 := GenS1{F0: "a", F1: -2147483648}
    fmt.Println(v5)
    fmt.Println(v5)
    v7 := v3[-1]
    fmt.Println(v5, v7)
    if (!((!((32767 != -32768) && false)) || (!false))) {
        v1 = append(v1, "")
    }
    v8 := GenS2{F0: -856}
    v2 -= (v2 % ((v2) | 1))
    switch ("" + "") {
    case "hello":
        var v9 int8 = 100
        _ = v9
    case "":
        var v10 string = ""
        _ = v10
    case "\n\t":
        var v11 int16 = 42
        _ = v11
    }
    v12 := GenS1{F0: "", F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
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
