package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 bool = false
    var v4 int32 = 1
    var v5 uint32 = 0
    v5 += v5
    v5 = 42
    _ = len(v2)
    v4 = (^(v4))
    fmt.Println(v5)
    v6 := GenS1{F0: v5, F1: 1}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint64 = 2
        _ = fi7
        _ = v8
    }
    if false {
        v5 = 2147483647
        v4 = -1
        switch -9223372036854775808 {
        case 737:
            v9 := []int{1, int(v4), 9223372036854775807}
            _ = v9
        case 9223372036854775807:
            v5 -= v5
        default:
            if (!((!v3) || v3)) {
                var v10 uint = 1
                fmt.Println(v3, v3)
                for fi11 := 0; fi11 < 11; fi11++ {
                    fmt.Println(v5, fi11, v10)
                    _ = fi11
                }
                _ = v10
            } else {
                var v12 int16 = 1
                var v13 int32 = v4
                _ = v12
                _ = v13
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
