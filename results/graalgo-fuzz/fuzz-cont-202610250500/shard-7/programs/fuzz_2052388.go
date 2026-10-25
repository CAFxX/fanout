package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{(("foo bar" + (("\n\t" + ("hello" + "hello")) + "hello")) + ""), (("hello" + "foo bar") + ("foo bar" + (((("hello" + "") + "") + (("foo bar" + ("" + "")) + "foo bar")) + "")))}
    var v3 int16 = 42
    var v4 int64 = -9223372036854775808
    delete(v1, 127)
    v4 = (v4 - 1)
    var v5 int16 = (v3 | v3)
    if (!true) {
        var v6 float32 = (float32(1e10) - float32(float64(-2.25)))
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v5 = v5
        _ = fi7
    }
    fmt.Println(v3)
    v8 := GenS1{F0: "a", F1: 1, F2: 696}
    v2 = append(v2, ("" + "a"))
    fmt.Println(v3, v5, v5)
    for fi9 := 0; fi9 < 9; fi9++ {
        v5 = v5
        _ = len(v2)
        _ = fi9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
