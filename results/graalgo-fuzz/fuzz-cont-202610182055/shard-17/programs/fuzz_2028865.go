package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int) uint16 {
    v4 := []int{}
    p0 += p0
    _ = p0
    _ = v4
    return 255
}

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 32767
    var v3 uint64 = 0
    v1 -= (v1 / (((v1 + v1)) | 1))
    fmt.Println(v3, v1, v1)
    var v5 uint16 = (v2 % ((v2) | 1))
    var v6 float32 = -2.25
    v7 := GenS1{F0: 42, F1: -128}
    var v8 string = "\n\t"
    var v9 uint64 = (^(v3))
    v5 = 65535
    var v10 int64 = -829
    fmt.Println(v9, v10)
    v11 := []int{959, -2147483648, 394}
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(v8, v5)
        for fi13 := 0; fi13 < 14; fi13++ {
            v6 = ((((v6 + -2.25) / v6) * float32(1e10)) * (1e10 / ((0.1 * (v6 * -2.25)) + 3.14159)))
            _ = fi13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: (^(v3)), F1: int(v1), F2: 4294967295}
        fmt.Println(v8)
        v16 := GenS3{F0: uint64(v5), F1: 42}
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v9 = (v9 / ((v3) | 1))
        _ = fi17
    }
    v18 := GenS2{F0: (v3 ^ v3), F1: 128, F2: (v1 * (v1 ^ (v1 * (^(v1)))))}
    v19 := GenS1{F0: -386, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
