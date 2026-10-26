package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 3.14159
    var v3 uint = 4294967295
    var v4 float32 = (0.0 * -2.25)
    var v5 int16 = 32767
    var v6 uint64 = 367
    v3 += v3
    v6 -= v6
    fmt.Println(v2)
    var v7 uint64 = 2
    v5 = v5
    v4 = (v4 * (v4 + -2.25))
    v8 := []int{-1, -767, -1}
    v5 += (v5 | 42)
    if (164 != 127) {
        v4 += 0.1
    } else {
        switch "\n\t" {
        case "":
            switch -639 {
            case -9223372036854775808:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
        }
        v9 := GenS1{F0: 1, F1: v1}
        fmt.Println(v3, v2)
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    var v11 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
