package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int64; F1 int32; F2 string }

func main() {
    var v1 bool = false
    var v2 string = "hello"
    var v3 float64 = 1e10
    v4 := make(map[uint32]uint32)
    v4[2] = 350
    switch 32768 {
    case 2147483647:
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := GenS1{F0: 255, F1: 42}
            v3 = v3
            var v7 int32 = -29
            _ = fi5
            _ = v6
            _ = v7
        }
    case -995:
        v8 := GenS1{F0: 1, F1: 558}
        _ = v8
    }
    if (520 <= 2) {
        v2 = (v2 + v2)
    }
    var v9 int8 = 0
    var v10 float32 = -2.25
    v11 := []int{4294967295, -64, 445, -2147483648}
    v12 := GenS2{F0: -1, F1: 1, F2: v2}
    if (!false) {
        if (1 > 1) {
            fmt.Println(v10)
        } else {
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(v2, v2, v9)
                var v14 bool = v1
                _ = fi13
                _ = v14
            }
        }
        switch "foo bar" {
        case "a":
            fmt.Println(v2, v9)
        }
        v1 = (495 < 78)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
