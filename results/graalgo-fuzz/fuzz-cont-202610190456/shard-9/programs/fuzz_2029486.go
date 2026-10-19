package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint64 = 1
    var v2 int8 = 0
    var v3 float32 = ((-2.25 - float32(((1e10 + 1.5) + (float64(-2.25) - 1e10)))) * (0.1 / 0.1))
    if (false || false) {
        if true {
            v3 += v3
            v3 -= v3
        }
    } else {
        fmt.Println(v3)
        var v4 string = "a"
        _ = v4
    }
    var v5 int = -1
    v6 := []int{128, 32767}
    var v7 uint32 = 0
    var v8 uint = 382
    v9 := GenS1{F0: 339, F1: 42}
    var v10 int16 = -421
    if (true && (3.14159 == (0.1 + ((0.1 / 0.1) * 0.0)))) {
        v11 := GenS1{F0: (v7 + (v7 << ((v8) & 63))), F1: -1}
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        if len(v6) > 0 {
            v6[0] = (v5 >> (((v8 & v8)) & 63))
        }
        if (((0.0 / v3) - float32(0.1)) <= ((3.14159 - 1e10) - 1.5)) {
            var v13 bool = false
            _ = v13
        } else {
            v7 = 2147483648
            if false {
                var v14 int16 = v10
                _ = v14
            }
        }
        fmt.Println(v7, v3)
        _ = fi12
    }
    v15 := GenS2{F0: 65535, F1: 50}
    for fi16 := 0; fi16 < 5; fi16++ {
        v8 -= 42
        var v17 bool = ((fi16 ^ fi16) != fi16)
        _ = fi16
        _ = v17
    }
    v7 -= (v7 | (v7 - (v7 + v7)))
    var v18 int8 = v2
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v19 := GenS3{F0: (v8 << (((v8 >> (((v8 / ((0) | 1))) & 63))) & 63)), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
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
