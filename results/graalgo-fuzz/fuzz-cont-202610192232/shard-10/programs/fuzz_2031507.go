package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]string)
    var v3 int64 = 9223372036854775807
    v3 -= (v3 >> ((115) & 63))
    var v4 int = -1
    v5 := GenS1{F0: 100}
    var v6 uint = 959
    fmt.Println(v3, v6)
    fmt.Println(v3)
    fmt.Println(v4)
    v7 := []int{v4, int(v6)}
    if (v6 <= v6) {
        v3 += 1
        var v8 float32 = (((1.5 - ((1.5 - ((-2.25 / 1e10) / (((3.14159 + (-2.25 + 1e10)) * 3.14159) / 1.0))) - 1.5)) / (0.1 + float32(0.0))) + (0.0 - -2.25))
        _ = v8
    }
    var v9 int16 = v1
    v3 = v3
    fmt.Println(v4, v9)
    v9 = (v1 ^ (v9 / ((178) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
