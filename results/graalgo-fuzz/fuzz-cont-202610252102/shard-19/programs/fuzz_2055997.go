package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0(p0 int16, p1 uint8, p2 uint32) bool {
    p1 = (p1 ^ (^(p1)))
    v5 := GenS1{F0: int(p1), F1: -342}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (-128 <= 0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (false || (255 < 200))
    var v3 int64 = 494
    var v4 int64 = (v3 / (((v3 & (v3 + -1))) | 1))
    fmt.Println(v4, v4, v2)
    v2 = v2
    switch "\n\t" {
    case "a":
        if ((true || v2) || (-2147483648 != 42)) {
            var v6 float64 = 0.1
            var v7 uint8 = 183
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := []int{32768, (fi8 - (fi8 >> ((255) & 63))), fi8, (^(fi8))}
                if len(v9) > 0 {
                    v9[0] = (fi8 << ((0) & 63))
                }
                _ = fi8
                _ = v9
            }
            _ = v6
            _ = v7
        } else {
            v10 := GenS1{F0: -1, F1: 32768}
            if (0 == -1) {
                v3 = (v3 << ((103) & 63))
                if false {
                    if true {
                        var v11 uint16 = 65535
                        _ = v11
                    }
                    v12 := GenS1{F0: 9223372036854775807, F1: 42}
                    _ = v12
                }
            }
            _ = v10
        }
    case "foo bar":
        var v13 uint = 18446744073709551615
        _ = v13
    case "hello":
        _ = len(v1)
    default:
        fmt.Println(v4, v2)
    }
    v14 := GenS1{F0: 2147483647, F1: 57}
    var v15 uint16 = 130
    var v16 uint64 = 4294967295
    if ((-2.25 == (3.14159 + ((float32(3.14159) - 0.1) * 0.0))) && (18446744073709551615 > 18446744073709551615)) {
        v4 -= -758
        var v17 uint32 = 2147483648
        _ = v17
    }
    var v18 float64 = -2.25
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS1{F0: fi19, F1: (fi20 - (fi19 - fi19))}
            var v22 int = 0
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    fmt.Println(v3, v2)
    var v23 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
