package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }

func main() {
    var v1 uint8 = 189
    var v2 float32 = 1.5
    var v3 float64 = 0.1
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    var v6 string = "a"
    switch v6 {
    case "":
        _ = len(v4)
    }
    switch 9223372036854775807 {
    case -392:
        v1 += (v1 % (((v1 << ((88) & 63))) | 1))
    case 0:
        v1 += 127
    default:
        fmt.Println(v3, v3, v1)
    }
    v2 -= (3.14159 + 0.0)
    v2 += (v2 / v2)
    switch 127 {
    case 0:
        if (65535 != 1) {
            fmt.Println(v1, v2, v1)
            v5[((v6 + "") + v6)] = (9223372036854775807 < -1)
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: v6, F1: 405, F2: ("\n\t" + "\n\t")}
                v6 = "a"
                _ = fi7
                _ = v8
            }
        } else {
            fmt.Println(v3, v6)
        }
    default:
        delete(v4, 2147483647)
    }
    var v9 int32 = -1
    if false {
        v9 = (v9 * v9)
        var v10 int64 = 1
        _ = v10
    }
    if false {
        var v11 uint32 = uint32(v1)
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
