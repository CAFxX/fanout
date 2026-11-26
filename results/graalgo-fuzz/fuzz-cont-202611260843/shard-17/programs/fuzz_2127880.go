package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = ((-2.25 / 1.0) / (0.0 + (3.14159 / 1e10)))
    var v3 int = 2147483647
    delete(v1, v3)
    v4 := []int{(v3 << ((42) & 63)), (v3 / (((v3 - v3)) | 1))}
    v5 := GenS1{F0: 128, F1: 867, F2: 2147483648}
    if (!((0 != 65535) && (493 < 610))) {
        v6 := GenS2{F0: 1, F1: 98, F2: -480}
        _ = v6
    } else {
        if (false && true) {
            fmt.Println(v3, v2)
        } else {
            v2 += (v2 * (v2 + v2))
            if false {
                fmt.Println(v3)
                var v7 uint64 = 18446744073709551615
                _ = v7
            } else {
                var v8 bool = (false && false)
                v9 := GenS2{F0: 18446744073709551615, F1: 65535, F2: 42}
                var v10 int8 = -1
                _ = v10
                _ = v8
                _ = v9
            }
        }
    }
    v2 = 1e10
    v2 += v2
    switch "\n\t" {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        switch "a" {
        case "hello":
            v2 = ((v2 + v2) / 0.1)
        default:
            v3 -= (v3 + v3)
        }
    default:
        var v11 bool = false
        _ = v11
    }
    v3 -= 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
