package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint32, p1 int8) uint {
    var v5 bool = (1e10 <= float64(float32((1e10 + ((1e10 * 3.14159) - 0.0)))))
    v6 := GenS1{F0: p0, F1: (^(p0))}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 878
}
func genFn1(p0 int, p1 uint8) uint64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(p0, p0, p1)
        _ = fi7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 = fi8
        p0 = p0
        if false {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := []int{(fi9 << ((4294967295) & 63)), fi9, fi9, (p0 | (p0 * (fi9 >> ((403) & 63))))}
                _ = fi9
                _ = v10
            }
        }
        _ = fi8
    }
    v11 := []int{(p0 * (p0 >> ((316) & 63))), p0}
    switch "\n\t" {
    case "hello":
        var v12 int = p0
        _ = v12
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = v11
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 0
    v3 := make(map[int]uint32)
    var v4 int16 = 32767
    var v13 uint = 0
    for fi14 := 0; fi14 < 8; fi14++ {
        v13 += v13
        _ = fi14
    }
    v2 = v2
    fmt.Println(v13)
    fmt.Println(v13)
    var v15 uint32 = 4294967295
    v15 -= (v15 % ((v15) | 1))
    var v16 float64 = float64(1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
