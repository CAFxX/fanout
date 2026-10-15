package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint8, p1 int64) uint16 {
    var v4 int = 9223372036854775807
    v5 := GenS1{F0: (p0 >> ((663) & 63))}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v6 int32 = -207
    v7 := []int{326, 1, 2147483647, 128}
    if len(v7) > 0 {
        v7[0] = -630
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        if (278 == 42) {
            if ("\n\t" <= "\n\t") {
                if len(v7) > 0 {
                    v7[0] = fi8
                }
            }
            v6 = v6
        } else {
            v6 -= (-(v6))
            switch "" {
            case "":
                for fi9 := 0; fi9 < 12; fi9++ {
                    v6 = v6
                    v7 = append(v7, fi9)
                    _ = fi9
                }
            case "a":
                v7 = append(v7, 513)
            default:
                _ = len(v1)
            }
            v6 += (v6 % ((v6) | 1))
        }
        var v10 int64 = 42
        v10 = -9223372036854775808
        _ = fi8
        _ = v10
    }
    switch -2147483648 {
    case -693:
        v6 = v6
    default:
        var v11 float64 = (0.1 / 3.14159)
        _ = v11
    }
    v6 += (v6 % ((v6) | 1))
    v6 += v6
    v12 := GenS2{F0: 0, F1: 4294967295}
    v13 := v2[0]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
