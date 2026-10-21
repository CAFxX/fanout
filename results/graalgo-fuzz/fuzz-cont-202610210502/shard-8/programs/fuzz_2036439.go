package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 int64 = 9223372036854775807
    if (42 < 128) {
        var v4 uint = 2
        var v5 uint64 = 432
        v6 := GenS1{F0: 255, F1: 2}
        _ = v4
        _ = v5
        _ = v6
    }
    fmt.Println(v3)
    v7 := []int{128, -899, -163, 32768}
    v3 = v3
    v3 = v3
    v3 -= (v3 ^ 1)
    v3 = (-(v3))
    v8 := GenS1{F0: 42, F1: 1}
    if false {
        switch 308 {
        case 32767:
            fmt.Println(v3)
        default:
            var v9 int64 = -326
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int32 = -2147483648
            v11 = (v11 % (((v11 - v11)) | 1))
            _ = fi10
            _ = v11
        }
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
    }
    v7 = append(v7, -1)
    v3 -= 42
    fmt.Println(v3)
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v3)
        _ = fi12
    }
    delete(v1, ("" + "\n\t"))
    v3 += 212
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
