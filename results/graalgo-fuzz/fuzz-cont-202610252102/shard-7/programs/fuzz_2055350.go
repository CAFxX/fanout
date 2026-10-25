package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{"a", ("a" + (("hello" + "foo bar") + "foo bar")), ("" + "\n\t"), "hello"}
    var v3 bool = (51 > 127)
    var v4 uint64 = 2
    var v5 float32 = 3.14159
    var v6 uint32 = 777
    v3 = v3
    v7 := GenS1{F0: 128}
    v4 -= (v4 >> ((uint(v6)) & 63))
    var v8 uint32 = v6
    v4 += (v4 * v4)
    _ = len(v2)
    v6 += v6
    var v9 uint32 = (v6 - 1)
    switch "\n\t" {
    case "foo bar":
        delete(v1, 0)
    }
    var v10 int16 = -32768
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
