package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int16, p1 int8, p2 bool) uint16 {
    p0 += (p0 & (p0 & p0))
    var v4 int16 = (p0 & p0)
    v5 := GenS1{F0: 9223372036854775807}
    switch "\n\t" {
    case "\n\t":
        for fi6 := 0; fi6 < 1; fi6++ {
            v4 -= p0
            _ = fi6
        }
    case "hello":
        p0 += (v4 & (p0 & p0))
    case "foo bar":
        p2 = (p2 || (((0.1 * 0.0) / 3.14159) < 0.0))
    }
    var v7 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v7
    return v7
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]bool)
    var v3 int = 32767
    v3 -= (v3 * -9223372036854775808)
    var v8 float32 = 0.0
    v9 := GenS1{F0: v3}
    v3 -= 42
    _ = len(v2)
    var v10 float64 = (1e10 + 1e10)
    var v11 bool = (0 != 888)
    var v12 uint = 869
    if true {
        v1 = v1
        if ((744 > 2) || (!(v8 == (v8 * v8)))) {
            v13 := GenS1{F0: v3}
            if (255 != 198) {
                var v14 uint16 = uint16(v12)
                _ = v14
            } else {
                v11 = true
                var v15 uint = 0
                v10 -= 0.1
                _ = v15
            }
            _ = v13
        } else {
            switch "a" {
            case "\n\t":
                _ = len(v2)
            default:
                switch v3 {
                case -279:
                    v8 -= 1.5
                default:
                    fmt.Println(v10, v10, v11)
                }
            }
        }
    }
    v16 := GenS2{F0: v12}
    var v17 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
