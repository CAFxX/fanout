package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int8 = 1
    var v2 uint8 = 1
    var v3 float64 = 3.14159
    v1 += v1
    v4 := GenS1{F0: 65535}
    var v5 uint8 = (^(v2))
    var v6 int16 = -32768
    var v7 uint8 = 255
    switch -2147483648 {
    case 32767:
        var v8 int32 = 185
        _ = v8
    }
    if (v7 <= (v2 / ((v2) | 1))) {
        if false {
            for fi9 := 0; fi9 < 3; fi9++ {
                v6 -= 707
                v2 -= v7
                _ = fi9
            }
            v5 += (v5 ^ v5)
            v10 := GenS2{F0: -477}
            _ = v10
        } else {
            var v11 int16 = (v6 - v6)
            v12 := GenS1{F0: 4294967295}
            _ = v11
            _ = v12
        }
        if (true && true) {
            var v13 int8 = v1
            v2 = (^(v2))
            v6 -= (^(v6))
            _ = v13
        }
    }
    v14 := GenS1{F0: 1}
    v15 := []int{}
    var v16 float32 = (0.1 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
