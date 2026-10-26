package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    var v1 string = ""
    v2 := []int{-9223372036854775808, 128}
    var v3 int32 = 260
    var v4 uint64 = 1
    v3 = v3
    var v5 int = 32768
    var v6 int = (-(v5))
    var v7 int64 = 0
    v7 = (v7 ^ (v7 * v7))
    fmt.Println(v4)
    v8 := GenS1{F0: 42, F1: 42}
    v5 = v5
    if (!(0 >= 42)) {
        if (145 >= 42) {
            switch v5 {
            case 1:
                var v9 int16 = -1
                _ = v9
            case -1:
                v2 = append(v2, v5)
            case 32768:
                fmt.Println(v6, v1, v5)
            }
            if (-1 >= -18) {
                var v10 int64 = v7
                _ = v10
            }
        }
        if len(v2) > 0 {
            v2[0] = (-(v6))
        }
        if true {
            var v11 int = (v5 ^ v6)
            _ = v11
        }
    } else {
        fmt.Println(v4, v7, v6)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v2 = append(v2, 128)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
