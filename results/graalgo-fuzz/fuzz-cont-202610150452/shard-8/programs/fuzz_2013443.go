package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 uint32 = 171
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := GenS1{F0: -2147483648, F1: -32768}
        v3 -= v3
        fmt.Println(fi4, v3)
        _ = fi4
        _ = v5
    }
    v3 -= 2
    _ = len(v1)
    _ = len(v2)
    switch "a" {
    case "a":
        if false {
            var v6 int = 9223372036854775807
            v3 = (v3 | v3)
            var v7 uint8 = 248
            _ = v6
            _ = v7
        }
    case "\n\t":
        fmt.Println(v3)
    }
    fmt.Println(v3)
    fmt.Println(v3)
    delete(v2, -1)
    switch 32767 {
    case -1:
        if false {
            var v8 string = "hello"
            if (-940 <= -9223372036854775808) {
                v9 := GenS2{F0: 29, F1: -1}
                _ = v9
            }
            _ = v8
        } else {
            fmt.Println(v3)
            var v10 int = 4294967295
            fmt.Println(v10, v3)
            _ = v10
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
