package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{"", "foo bar", ("hello" + ("a" + "\n\t"))}
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 string = (("a" + "\n\t") + ("foo bar" + "a"))
        v7 := GenS1{F0: 460}
        var v8 bool = (-128 <= -128)
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v2 += (v2 + (0.1 * (v2 + v2)))
    v2 -= (1.5 - (v2 + 0.0))
    if (2 == 163) {
        v2 -= v2
    }
    var v9 int64 = 9223372036854775807
    v9 = v9
    v10 := GenS1{F0: v9}
    var v11 float64 = v2
    var v12 uint32 = 500
    v2 = v11
    v13 := GenS2{F0: -443}
    v14 := GenS2{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
