package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 1
    v3 := []string{"a", "hello", ((("\n\t" + "a") + "a") + "foo bar"), "hello"}
    v4 := []int{-9223372036854775808, 636, -9223372036854775808, 127}
    v5 := []string{"foo bar", "hello", "a"}
    v2 = (v2 >> ((2) & 63))
    delete(v1, -946)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: -9223372036854775808, F1: 411}
            var v9 int64 = 0
            v2 -= (v2 + v2)
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: -1, F1: 255}
        _ = fi6
        _ = v10
    }
    if (128 >= 42) {
        v11 := GenS1{F0: 1, F1: 435}
        _ = v11
    }
    v2 = v2
    var v12 uint64 = 727
    v2 -= v2
    var v13 int64 = -1
    v4 = append(v4, 127)
    fmt.Println(v12, v12, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
