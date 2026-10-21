package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    var v1 uint32 = 2
    var v2 uint32 = v1
    v3 := make(map[string]int)
    if false {
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 uint8 = 255
            var v6 uint16 = 255
            _ = fi4
            _ = v5
            _ = v6
        }
        switch "foo bar" {
        case "a":
            if (((float32(float64(3.14159)) / 1.0) - 0.1) > 0.1) {
                v2 = (v2 << ((116) & 63))
                delete(v3, "")
            }
        }
    }
    v2 = (v1 % ((1) | 1))
    if (!false) {
        v1 -= 65535
    }
    v7 := []int{1, 9223372036854775807}
    fmt.Println(v1, v2)
    _ = len(v7)
    _ = len(v3)
    fmt.Println(v1)
    v1 += v2
    var v8 uint8 = 0
    fmt.Println(v8, v1, v1)
    v9 := GenS1{F0: -9223372036854775808, F1: 42}
    if true {
        v1 = v2
    } else {
        if false {
            var v10 uint64 = 915
            var v11 uint16 = 0
            _ = v10
            _ = v11
        }
        var v12 int8 = -128
        _ = v12
    }
    if false {
        var v13 int16 = 42
        _ = v13
    }
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
