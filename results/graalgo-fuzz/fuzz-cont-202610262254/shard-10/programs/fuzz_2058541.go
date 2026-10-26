package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0(p0 bool, p1 int16) uint32 {
    fmt.Println(p1, p0)
    p0 = p0
    switch -1 {
    case 128:
        var v7 float64 = (3.14159 + (0.0 / 1.0))
        _ = v7
    case -9223372036854775808:
        fmt.Println(p0)
    case 42:
        p1 = (p1 - (^(p1)))
    }
    p0 = p0
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint = 65535
    var v2 uint8 = 1
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 255
    var v5 string = (("\n\t" + "a") + "hello")
    v6 := make(map[uint32]uint32)
    var v8 int16 = 0
    v9 := []int{4294967295, -2147483648, -2147483648}
    switch 4294967295 {
    case -2147483648:
        _ = len(v6)
    default:
        v1 = (v1 / (((v1 >> (((v1 >> (((^(v1))) & 63))) & 63))) | 1))
    }
    var v10 uint32 = uint32(v3)
    v2 -= v2
    if len(v9) > 0 {
        v9[0] = 4294967295
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: v3, F1: (v8 % (((v8 & v8)) | 1)), F2: 925}
        v2 += (v2 - v2)
        v10 = (v10 & 65535)
        _ = fi11
        _ = v12
    }
    var v13 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
