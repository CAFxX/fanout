package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint64, p1 string, p2 int16) uint8 {
    v5 := GenS1{F0: 105}
    switch 0 {
    case -521:
        p2 -= p2
    case -9223372036854775808:
        var v6 float32 = (3.14159 * ((((0.0 * 3.14159) / 1.5) + 0.1) * float32((-2.25 / ((float64(float32(float64(0.1))) + (3.14159 - -2.25)) + 1.5)))))
        _ = v6
    case 128:
        v7 := GenS1{F0: 950}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 float64 = ((3.14159 / 3.14159) * 1e10)
    var v3 uint16 = 2
    var v4 uint16 = (v3 * (v3 << ((65535) & 63)))
    var v8 uint64 = 42
    if true {
        var v9 int8 = 127
        v3 = (v3 % ((65535) | 1))
        _ = v9
    }
    v4 -= (v3 - (v4 << ((4294967295) & 63)))
    switch 477 {
    case -1:
        v8 = v8
    case 244:
        v10 := []int{1, 32767, 42, 32768}
        _ = v10
    case 4294967295:
        switch "a" {
        case "\n\t":
            if false {
                var v11 float32 = 3.14159
                v2 += 3.14159
                if ((0.0 + (v2 - v2)) >= (v2 / (v2 * v2))) {
                    v12 := []int{127}
                    _ = v12
                }
                _ = v11
            }
        default:
            var v13 uint = 2
            _ = v13
        }
    default:
        switch 2147483647 {
        case 1:
            fmt.Println(v3, v2, v8)
        case 0:
            v3 += (v3 & (v3 * v3))
        case 2147483647:
            v2 += 3.14159
        }
    }
    if (1 >= 42) {
        fmt.Println(v1, v2)
        v2 = (v2 / float64(-2.25))
    } else {
        var v14 int64 = v1
        var v15 int8 = 29
        _ = v14
        _ = v15
    }
    v2 -= v2
    v1 = v1
    if true {
        v16 := GenS1{F0: 2147483647}
        _ = v16
    } else {
        v17 := GenS1{F0: 643}
        v18 := []int{-287, 42}
        var v19 uint32 = 2147483647
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
