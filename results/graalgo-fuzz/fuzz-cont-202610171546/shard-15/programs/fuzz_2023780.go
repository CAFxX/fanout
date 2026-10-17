package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int = 127
    v3 = -1
    var v4 bool = false
    fmt.Println(v4, v3)
    if v4 {
        v4 = v4
        var v5 int64 = 9223372036854775807
        _ = v5
    } else {
        switch v3 {
        case -9223372036854775808:
            if ((!(120 != -66)) || v4) {
                fmt.Println(v3, v3)
            } else {
                v6 := GenS1{F0: 2147483647}
                _ = v6
            }
        case 1:
            v7 := GenS2{F0: -1, F1: 18446744073709551615}
            _ = v7
        case 336:
            var v8 int64 = 0
            _ = v8
        }
        v4 = v4
        v3 += 476
    }
    if false {
        if v4 {
            v9 := GenS2{F0: 9223372036854775807, F1: 542}
            _ = v9
        } else {
            var v10 uint32 = 4294967295
            _ = v10
        }
        var v11 string = "a"
        delete(v2, "a")
        _ = v11
    }
    var v12 bool = v4
    var v13 uint16 = 65535
    var v14 bool = v4
    fmt.Println(v4)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
