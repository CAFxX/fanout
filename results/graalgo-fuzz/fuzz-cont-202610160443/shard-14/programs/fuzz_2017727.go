package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint64 = 712
    var v2 uint64 = v1
    v3 := make(map[string]int)
    var v4 uint16 = 0
    v5 := []string{"foo bar", "hello"}
    var v6 int64 = -10
    var v7 float32 = 1e10
    v5 = append(v5, "\n\t")
    v8 := GenS1{F0: 255}
    v6 += int64(v4)
    var v9 int64 = (v6 - 9223372036854775807)
    _ = len(v5)
    var v10 float32 = (v7 / v7)
    var v11 uint = 18446744073709551615
    if true {
        var v12 bool = (true || (0 < 69))
        _ = v12
    }
    fmt.Println(v6, v10, v6)
    fmt.Println(v2, v4, v11)
    var v13 int8 = 42
    if len(v5) > 0 {
        v5[0] = "a"
    }
    fmt.Println(v11, v11)
    _ = len(v3)
    var v14 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
