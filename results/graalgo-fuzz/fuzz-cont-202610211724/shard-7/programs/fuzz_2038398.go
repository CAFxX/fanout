package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 1
    v3 := []string{"foo bar", "\n\t", "hello"}
    var v4 bool = false
    v5 := make(map[uint32]string)
    v6 := []uint32{42, 0, 65535, 1}
    var v7 int32 = 42
    v8 := v1[65535]
    _ = len(v6)
    fmt.Println(v4)
    fmt.Println(v2)
    fmt.Println(v2)
    v4 = (float64(3.14159) < (-2.25 * 3.14159))
    if v4 {
        fmt.Println(v8, v7, v2)
        v9 := GenS1{F0: 130, F1: (v8 & (v2 / ((128) | 1)))}
        if v4 {
            var v10 int32 = v7
            _ = v10
        } else {
            switch (("hello" + "\n\t") + "a") {
            case "":
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            default:
                var v11 float64 = 0.0
                _ = v11
            }
        }
        _ = v9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
