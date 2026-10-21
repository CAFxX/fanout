package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    var v1 int16 = 42
    var v2 int8 = -128
    v3 := []string{"a", (((("foo bar" + "hello") + ("" + "\n\t")) + "foo bar") + ""), "foo bar", "hello"}
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    v6 := make(map[string]string)
    if true {
        v7 := GenS1{F0: 32767}
        _ = v7
    } else {
        v1 += v1
        _ = len(v3)
    }
    v1 = 42
    _ = len(v3)
    v2 = (v2 * (v2 + -110))
    switch ("hello" + "") {
    case "foo bar":
        v8 := GenS1{F0: 32767}
        _ = v8
    case "\n\t":
        v1 = v1
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v1)
    fmt.Println(v1, v2)
    switch 42 {
    case 128:
        v9 := GenS2{F0: 42, F1: ((-778 == v1) && false)}
        _ = v9
    }
    v10 := v6[("\n\t" + "a")]
    var v11 int64 = -630
    var v12 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
