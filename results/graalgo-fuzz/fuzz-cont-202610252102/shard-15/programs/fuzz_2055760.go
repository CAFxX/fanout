package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 int }

func main() {
    var v1 float64 = 3.14159
    var v2 int32 = 42
    v3 := []int{128}
    var v4 int8 = 42
    v5 := []uint32{0}
    var v6 int = 128
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 32767, F1: 850}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 42
    v10 := GenS2{F0: 42, F1: 42, F2: 165}
    if true {
        v11 := GenS2{F0: 32767, F1: 0, F2: v9}
        fmt.Println(v2, v2, v4)
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 float32 = ((float32(v1) * 0.1) - 1e10)
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    var v14 uint8 = v9
    fmt.Println(v2)
    _ = len(v5)
    v2 = (v2 | (v2 * v2))
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v15 uint = 18446744073709551615
        v6 -= (-(v6))
        _ = v15
    }
    var v16 int8 = (v4 << ((4294967295) & 63))
    var v17 uint = 606
    var v18 uint32 = 2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v5) > 0 {
        v5[0] = (v18 | 2)
    }
    switch (^(v6)) {
    case 32768:
        fmt.Println(v14)
    }
    v19 := GenS3{F0: 100, F1: -60, F2: (v6 << ((v17) & 63))}
    v20 := GenS2{F0: 32767, F1: 524, F2: (^(v9))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
