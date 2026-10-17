package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }

func main() {
    v1 := []string{(("hello" + "hello") + "foo bar")}
    var v2 bool = false
    v3 := []string{("foo bar" + "foo bar"), ""}
    v4 := make(map[int]int)
    var v5 float64 = 0.1
    var v6 string = (((("foo bar" + "hello") + "hello") + ("" + "hello")) + "foo bar")
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: (fi7 - (^(fi7))), F1: 32767, F2: (3.14159 != ((float32(v5) + 1e10) - 0.0))}
        v9 := GenS2{F0: 1, F1: (fi7 + (fi7 * (fi7 / ((fi7) | 1)))), F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = 128
    var v11 float32 = ((float32(float64(1e10)) / (((3.14159 + 1e10) * ((1.5 - 0.1) / -2.25)) - 3.14159)) * float32(float64(0.0)))
    v6 = v6
    v5 = float64(float32(v5))
    var v12 int64 = 42
    fmt.Println(v5)
    var v13 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
