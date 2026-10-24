package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int = 0
    var v4 uint64 = 4294967295
    var v5 bool = false
    v3 = (v3 & 128)
    var v6 bool = v5
    v4 -= v4
    switch v3 {
    case 4294967295:
        var v7 int32 = -1
        _ = v7
    case 0:
        switch v3 {
        case 1:
            if true {
                fmt.Println(v5, v5)
                var v8 uint32 = 42
                _ = v8
            } else {
                if false {
                    v6 = v5
                }
            }
        case -388:
            v9 := GenS1{F0: 342, F1: (v3 | (-(v3)))}
            _ = v9
        default:
            fmt.Println(v4, v6, v3)
        }
    }
    fmt.Println(v5)
    switch (v3 % ((v3) | 1)) {
    case -2147483648:
        v10 := []int{v3, (v3 | (v3 % ((v3) | 1))), -9223372036854775808}
        _ = v10
    case 128:
        switch v3 {
        case 2147483647:
            v5 = ("foo bar" <= "")
        }
    }
    if (0.0 < 1.5) {
        fmt.Println(v3, v3)
        v11 := []int{v3, (-(v3)), (v3 ^ (v3 ^ (^(v3)))), 32768}
        _ = v11
    }
    _ = len(v2)
    var v12 uint = 67
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
