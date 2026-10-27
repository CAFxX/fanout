package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 6
    v3 := make(map[int]int)
    v4 := []string{"", (("a" + ("foo bar" + "\n\t")) + (("" + "\n\t") + ("foo bar" + "foo bar"))), "\n\t"}
    var v5 uint16 = 255
    v5 = (v5 - v5)
    v6 := GenS1{F0: 0}
    v2 = -1
    if (false && true) {
        switch -548 {
        case 1:
            var v7 int32 = -1
            _ = v7
        default:
            fmt.Println(v2)
        }
        if (float32((3.14159 * 0.1)) != 1e10) {
            v5 += v5
            v8 := v1[("hello" + ("" + ""))]
            _ = v8
        }
        var v9 int64 = (-(v2))
        _ = v9
    }
    v5 -= v5
    v10 := GenS1{F0: 0}
    v5 = 42
    var v11 int64 = (-(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
