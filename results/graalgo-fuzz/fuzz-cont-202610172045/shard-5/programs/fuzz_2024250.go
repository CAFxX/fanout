package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint32, p1 int64, p2 bool) uint32 {
    v6 := GenS1{F0: 4294967295, F1: p0}
    v7 := []int{9223372036854775807, 976}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 / ((p0) | 1))
}
func genFn1(p0 uint8) string {
    v8 := GenS1{F0: 18446744073709551615, F1: 1}
    p0 -= (p0 | p0)
    _ = p0
    _ = v8
    return ("" + "a")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 42
    var v3 uint64 = 1
    var v4 bool = ((-2.25 <= (3.14159 / ((1e10 / 1e10) + 3.14159))) && (42 == -32768))
    v5 := make(map[string]bool)
    var v9 uint = (v2 | (v2 ^ (^(v2))))
    v10 := GenS1{F0: (v9 * v2), F1: 4294967295}
    fmt.Println(v9)
    var v11 uint16 = 42
    var v12 float32 = 3.14159
    var v13 int64 = -615
    for fi14 := 0; fi14 < 1; fi14++ {
        switch "foo bar" {
        case "\n\t":
            var v15 int32 = -2147483648
            _ = v15
        default:
            for fi16 := 0; fi16 < 9; fi16++ {
                v13 -= 42
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v11 = 2
                case "hello":
                    _ = len(v5)
                default:
                    v9 += 42
                }
                _ = fi16
            }
        }
        _ = fi14
    }
    v17 := GenS2{F0: uint32(v11), F1: 0}
    v9 = (v2 * (v9 << ((v2) & 63)))
    v18 := []int{32768, 555, 128}
    var v19 float64 = (1.5 - (3.14159 * 0.0))
    v20 := GenS1{F0: (v9 >> ((255) & 63)), F1: 2147483648}
    v21 := GenS2{F0: 42, F1: 870}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
