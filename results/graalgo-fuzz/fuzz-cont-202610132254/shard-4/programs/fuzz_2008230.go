package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8) string {
    fmt.Println(p0)
    p0 -= (p0 & p0)
    _ = p0
    return "a"
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 42
    var v3 int64 = 9223372036854775807
    var v4 float32 = float32(1e10)
    var v5 int32 = 360
    var v6 int32 = v5
    if false {
        var v7 uint64 = 18446744073709551615
        v2 = v2
        _ = v7
    }
    v8 := GenS1{F0: -109}
    fmt.Println(v3, v3)
    var v9 int32 = 42
    switch -993 {
    case 42:
        v10 := GenS1{F0: -128}
        _ = v10
    case 494:
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int8 = -1
            v13 := GenS1{F0: v12}
            v9 += v5
            _ = fi11
            _ = v12
            _ = v13
        }
    case 9223372036854775807:
        v14 := []int{-9223372036854775808, 128}
        _ = v14
    default:
        v5 = v5
    }
    fmt.Println(v4, v2)
    if true {
        v15 := GenS2{F0: -32768}
        fmt.Println(v3)
        if (v3 == v3) {
            if true {
                switch 128 {
                case -1:
                    fmt.Println(v2, v4, v6)
                case 0:
                    fmt.Println(v1, v5, v2)
                default:
                    var v16 float32 = float32((1e10 - (0.0 * (float64(1e10) - (-2.25 * 0.1)))))
                    _ = v16
                }
            }
        }
        _ = v15
    }
    if false {
        var v17 uint = 1
        v18 := []int{0, 239}
        v6 = v9
        _ = v17
        _ = v18
    }
    switch 2147483647 {
    case 4294967295:
        if ((v1 | v1) <= (v1 | (v1 - (v1 | (^(v1)))))) {
            var v19 uint64 = 18446744073709551615
            v20 := []int{960, 0, 32768, 127}
            if len(v20) > 0 {
                v20[0] = -2147483648
            }
            _ = v19
            _ = v20
        } else {
            var v21 int8 = 1
            _ = v21
        }
    case -221:
        var v22 uint64 = 702
        _ = v22
    default:
        v23 := []int{32768}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
