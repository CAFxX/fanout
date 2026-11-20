package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 string) uint64 {
    v7 := GenS1{F0: 32767, F1: true}
    if (9223372036854775807 <= -9223372036854775808) {
        fmt.Println(p0, p0)
    }
    p0 = (p0 / ((0) | 1))
    p1 = "\n\t"
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 65
    var v3 float32 = (0.0 - (3.14159 + 1e10))
    var v4 int = 1
    var v5 int = (v4 * v4)
    var v6 int8 = -128
    var v8 string = ("a" + "foo bar")
    v5 = v5
    v9 := GenS1{F0: 1, F1: ((true && (!false)) && true)}
    v10 := []int{v5, 4294967295, (v5 | (v5 / (((v4 | (-(v4)))) | 1)))}
    if len(v10) > 0 {
        v10[0] = (^(v4))
    }
    v8 = ((v8 + "\n\t") + "\n\t")
    switch (v4 & v5) {
    case 42:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    default:
        for fi11 := 0; fi11 < 11; fi11++ {
            if len(v10) > 0 {
                v10[0] = -9223372036854775808
            }
            v12 := GenS2{F0: 1}
            var v13 int = (v4 << ((2) & 63))
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
