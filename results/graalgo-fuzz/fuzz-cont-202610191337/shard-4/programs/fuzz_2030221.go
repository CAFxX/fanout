package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int32 }
func genFn0() uint {
    var v6 float64 = 3.14159
    v6 -= -2.25
    var v7 float64 = v6
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint32) uint64 {
    var v8 uint8 = 192
    v8 = (v8 * 128)
    v9 := GenS1{F0: 2, F1: true}
    p0 -= p0
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int = -770
    var v3 int64 = 42
    v4 := make(map[int]uint32)
    var v5 float32 = 1e10
    for fi10 := 0; fi10 < 3; fi10++ {
        delete(v4, fi10)
        _ = fi10
    }
    v11 := GenS2{F0: 42, F1: "a"}
    v2 = (v1 | 841)
    v1 = (v1 - v1)
    var v12 uint8 = 0
    var v13 int = int(v3)
    v4[-134] = 674
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: 548, F1: false}
        _ = fi14
        _ = v15
    }
    v13 = -452
    v1 -= -278
    if false {
        var v16 float64 = 3.14159
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            if ((("" <= "a") || false) || false) {
                fmt.Println(v12, v3, v5)
            }
            var v19 int8 = 0
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v20 := GenS3{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
