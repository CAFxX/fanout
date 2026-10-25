package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint32 {
    switch "foo bar" {
    case "":
        v4 := []int{-238}
        _ = v4
    default:
        if (84 < 42) {
            if (3.14159 != -2.25) {
                for fi5 := 0; fi5 < 11; fi5++ {
                    var v6 float32 = (float32(1e10) / 1.0)
                    _ = fi5
                    _ = v6
                }
                var v7 uint64 = 18446744073709551615
                switch -2147483648 {
                case -124:
                    v8 := GenS1{F0: 65535}
                    _ = v8
                default:
                    var v9 int32 = 0
                    _ = v9
                }
                _ = v7
            } else {
                var v10 float64 = 1e10
                var v11 uint32 = 65535
                _ = v10
                _ = v11
            }
            var v12 float32 = (0.0 * ((0.0 * float32(-2.25)) * -2.25))
            v12 -= float32(0.0)
            _ = v12
        }
    }
    var v13 uint32 = 0
    v13 = v13
    _ = v13
    return v13
}

func main() {
    v1 := []int{9223372036854775807, 1}
    var v2 int8 = -54
    v3 := []int{127}
    v2 = (v2 + v2)
    switch 2147483647 {
    case -817:
        v2 = v2
    }
    if (true || (!true)) {
        switch 802 {
        case 0:
            if (("a" + "") == "a") {
                v2 = (v2 - (-(v2)))
            } else {
                var v14 uint16 = 261
                _ = v14
            }
        case 32768:
            v15 := GenS1{F0: 65535}
            _ = v15
        default:
            fmt.Println(v2)
        }
    }
    var v16 uint = 42
    fmt.Println(v2, v2)
    var v17 uint64 = 1
    v2 += v2
    v18 := GenS1{F0: 0}
    fmt.Println(v2)
    var v19 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
