package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
func genFn0(p0 uint16, p1 uint8, p2 uint32) uint64 {
    if false {
        p2 -= p2
        fmt.Println(p1, p2, p2)
        p1 += 42
    } else {
        p2 -= (^(p2))
    }
    var v5 uint64 = 2
    p1 = 2
    p0 -= (p0 ^ p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (^(v5))
}

func main() {
    v1 := []string{"\n\t", (("a" + ("\n\t" + "")) + "")}
    var v2 bool = false
    v3 := []uint32{2147483647}
    var v4 uint32 = 1
    v2 = (true && (1e10 > 1e10))
    v6 := GenS1{F0: 0, F1: 111, F2: 2}
    fmt.Println(v4)
    if ((((3.14159 * 3.14159) - ((-2.25 * (1.5 * ((((-2.25 + 3.14159) * (-2.25 + 1e10)) / float32(1e10)) + 1e10))) * float32((1e10 / 1.0)))) / 3.14159) <= (float32(3.14159) * 1e10)) {
        v2 = false
    }
    fmt.Println(v2, v2)
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint32 = (v4 << ((533) & 63))
        var v10 string = "\n\t"
        if true {
            var v11 string = ("\n\t" + "hello")
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v4 += v4
    switch 128 {
    case 0:
        v7 += v7
    case 832:
        fmt.Println(v7, v7)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
