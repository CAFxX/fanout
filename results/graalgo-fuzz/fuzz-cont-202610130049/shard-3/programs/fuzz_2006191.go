package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := []int{-105, 128}
    var v2 int64 = 338
    v3 := make(map[string]uint32)
    var v4 uint16 = 0
    v1 = append(v1, -1)
    v5 := GenS1{F0: 1, F1: 2}
    v6 := GenS1{F0: 4294967295, F1: 2}
    v7 := GenS1{F0: 422, F1: 42}
    var v8 float32 = (-2.25 / 3.14159)
    delete(v3, "")
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 4294967295, F1: 45}
        _ = fi10
        _ = v11
    }
    if len(v1) > 0 {
        v1[0] = -789
    }
    switch 128 {
    case -309:
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS2{F0: "\n\t", F1: 42}
            v9 -= (v9 >> ((0) & 63))
            v8 += (1e10 * ((v8 / v8) / float32(((1.5 / 3.14159) + 3.14159))))
            _ = fi12
            _ = v13
        }
    case 1:
        v8 -= 0.1
    case 9223372036854775807:
        v3[(((("a" + "hello") + ("foo bar" + ("hello" + "a"))) + "\n\t") + "foo bar")] = 122
    default:
        v4 -= (v4 ^ 642)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
