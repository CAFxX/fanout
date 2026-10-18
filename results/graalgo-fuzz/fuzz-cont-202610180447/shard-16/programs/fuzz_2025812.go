package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 868
    var v3 int16 = -1
    v4 := GenS1{F0: 55, F1: ""}
    fmt.Println(v2, v2)
    if (-2147483648 <= 32767) {
        v2 -= 42
        v2 = (v2 & v2)
        v3 = -170
    } else {
        v3 = (v3 << ((65535) & 63))
        if false {
            var v5 uint16 = 2
            v2 -= 1
            _ = v5
        }
        switch "\n\t" {
        case "a":
            var v6 uint8 = 128
            _ = v6
        default:
            fmt.Println(v3, v3)
        }
    }
    v7 := GenS2{F0: "\n\t", F1: 14}
    v8 := []int{4294967295, 1, 127}
    if len(v8) > 0 {
        v8[0] = 1
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint16 = 482
        _ = fi9
        _ = v10
    }
    var v11 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
