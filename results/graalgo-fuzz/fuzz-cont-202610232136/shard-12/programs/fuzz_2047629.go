package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int32; F1 string }

func main() {
    var v1 int8 = -100
    var v2 uint64 = 18446744073709551615
    var v3 float32 = 0.0
    var v4 float32 = v3
    var v5 int16 = 42
    fmt.Println(v3, v2)
    if false {
        v6 := GenS1{F0: v2}
        _ = v6
    } else {
        if true {
            v7 := GenS1{F0: 2}
            fmt.Println(v3, v4)
            v5 = (v5 << ((2) & 63))
            _ = v7
        } else {
            v3 = ((3.14159 - 1.5) + 1e10)
        }
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint = 18446744073709551615
        var v10 uint8 = 255
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 -= v4
    if (!(v2 < (v2 ^ v2))) {
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(v2)
            v3 += ((v4 - v3) - v3)
            _ = fi11
        }
        v4 -= 3.14159
        v12 := GenS2{F0: 0, F1: (v1 << ((865) & 63))}
        _ = v12
    } else {
        v5 = (v5 - v5)
        var v13 bool = (807 >= 65535)
        _ = v13
    }
    v14 := GenS1{F0: (^(v2))}
    var v15 float64 = 3.14159
    v5 = -1
    for fi16 := 0; fi16 < 10; fi16++ {
        if true {
            fmt.Println(fi16, v1, fi16)
            switch (fi16 * fi16) {
            case 625:
                v17 := []int{fi16, -991, (fi16 | fi16)}
                _ = v17
            case -991:
                var v18 int = (fi16 * (fi16 | fi16))
                _ = v18
            case 42:
                fmt.Println(v15, v3)
            default:
                var v19 float64 = float64(1.5)
                _ = v19
            }
            v20 := GenS3{F0: 580, F1: "\n\t"}
            _ = v20
        }
        v21 := GenS2{F0: 235, F1: (v1 - (v1 & (-(v1))))}
        _ = fi16
        _ = v21
    }
    v4 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
