package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }

func main() {
    var v1 uint = 2
    var v2 float64 = 3.14159
    var v3 int8 = 1
    v4 := make(map[int]bool)
    v3 = (v3 * (-(v3)))
    var v5 int32 = 42
    var v6 int32 = (v5 ^ v5)
    switch 402 {
    case 1:
        v7 := []int{9223372036854775807}
        _ = v7
    case 127:
        switch 127 {
        case -2147483648:
            switch ((("\n\t" + (("" + "") + "\n\t")) + "\n\t") + "foo bar") {
            case "":
                var v8 int32 = v5
                _ = v8
            }
        }
    case -2147483648:
        fmt.Println(v2)
    default:
        v6 = v5
    }
    v9 := GenS1{F0: (v6 % ((v6) | 1)), F1: int64(v6), F2: v6}
    v4[int(v3)] = (false && ((428 >= 1) && true))
    _ = len(v4)
    var v10 uint32 = 2
    var v11 int = 128
    var v12 uint32 = (v10 % ((v10) | 1))
    v11 = 999
    switch "a" {
    case "a":
        v12 = v12
    default:
        var v13 float32 = 0.0
        _ = v13
    }
    v11 -= v11
    if (v10 == (v10 / (((v12 % (((v12 << (((v1 & 715)) & 63))) | 1))) | 1))) {
        v14 := GenS1{F0: v5, F1: 583, F2: (v5 >> ((v1) & 63))}
        var v15 float32 = (1e10 * 0.0)
        var v16 uint8 = 159
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint16 = 531
    var v18 float64 = v2
    if false {
        v3 = v3
    }
    if (42 > 18446744073709551615) {
        fmt.Println(v12, v11)
    } else {
        var v19 uint64 = 2
        delete(v4, (v11 / (((v11 | v11)) | 1)))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
