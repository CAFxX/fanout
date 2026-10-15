package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0(p0 uint32, p1 uint8, p2 uint8) uint8 {
    if (("a" + ("hello" + "foo bar")) < ((("a" + "") + "hello") + ((("hello" + "") + "a") + "\n\t"))) {
        v4 := GenS1{F0: -128, F1: 1000, F2: (p0 | p0)}
        v5 := []int{}
        _ = v4
        _ = v5
    } else {
        var v6 uint64 = 850
        _ = v6
    }
    p1 = p2
    v7 := []int{-2147483648, -923, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}

func main() {
    var v1 string = ""
    var v2 int = -420
    var v3 uint = 833
    v8 := []int{(v2 << (((v3 | v3)) & 63)), (v2 | (v2 >> ((v3) & 63)))}
    v8 = append(v8, (v2 << ((4294967295) & 63)))
    var v9 float32 = (0.0 * float32((((((-2.25 * 1.5) / 1.0) - (-2.25 + (1.5 - ((3.14159 / 3.14159) - (1e10 - 1.5))))) * (0.0 / 1e10)) - 0.0)))
    v10 := GenS1{F0: 23, F1: 34, F2: 117}
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v9, v1, fi12)
        _ = fi12
    }
    var v13 uint8 = 42
    if (((v2 < v2) && true) && ((((((((v9 + 1.5) - -2.25) + v9) / v9) + v9) * v9) * (v9 * (v9 / (float32(0.0) * (0.1 - v9))))) == (((v9 - v9) + v9) * v9))) {
        switch v2 {
        case 307:
            v2 += v2
        }
        v14 := GenS1{F0: 42, F1: 1000, F2: 381}
        _ = v14
    } else {
        var v15 int = v2
        _ = v15
    }
    v2 += (v2 << (((v3 % (((v3 / ((v3) | 1))) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
