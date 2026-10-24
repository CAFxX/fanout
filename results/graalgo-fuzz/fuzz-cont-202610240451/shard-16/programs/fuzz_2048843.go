package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int, p1 int, p2 uint64) int8 {
    p2 = (p2 ^ (p2 >> ((342) & 63)))
    v4 := GenS1{F0: p1}
    v5 := GenS1{F0: (p0 & 4294967295)}
    v6 := []int{(p1 | 1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 uint16 = 65535
    var v2 float64 = float64(((((float32(1.5) * (3.14159 * 0.1)) / (((1e10 * (1e10 * -2.25)) / 1.5) * float32(0.1))) + 1e10) + (0.1 * 1.5)))
    v3 := []uint32{}
    _ = len(v3)
    v7 := GenS1{F0: 32768}
    if len(v3) > 0 {
        v3[0] = 122
    }
    v2 = ((1.5 * ((-2.25 + 3.14159) * ((v2 * v2) * v2))) * float64(1.5))
    var v8 int8 = genFn0(127, -165, uint64(v1))
    switch "\n\t" {
    case "":
        switch 249 {
        case 127:
            if (3.14159 >= 1e10) {
                var v9 float32 = (1e10 / float32(v2))
                for fi10 := 0; fi10 < 18; fi10++ {
                    v3 = append(v3, 695)
                    _ = fi10
                }
                _ = v9
            }
        default:
            if (true || (!((v2 * v2) < (v2 * v2)))) {
                var v11 int16 = -1
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    fmt.Println(v1)
                default:
                    v1 = (v1 + v1)
                }
                v12 := GenS1{F0: 42}
                _ = v11
                _ = v12
            }
        }
    case "\n\t":
        fmt.Println(v1, v2, v1)
    case "a":
        var v13 uint = 2
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint = 791
        _ = fi14
        _ = v15
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
