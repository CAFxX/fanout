package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }

func main() {
    var v1 uint8 = 1
    var v2 uint = 990
    var v3 int32 = 42
    v4 := make(map[uint32]uint32)
    var v5 int64 = 622
    v6 := make(map[string]string)
    var v7 int16 = -821
    var v8 int16 = 1
    var v9 float64 = 1e10
    switch "foo bar" {
    case "a":
        v10 := GenS1{F0: 42, F1: v1}
        _ = v10
    case "foo bar":
        v11 := GenS2{F0: -2147483648, F1: 42}
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        if false {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 int32 = 1
                switch 390 {
                case 443:
                    var v15 uint = (v2 - (v2 ^ v2))
                    _ = v15
                default:
                    fmt.Println(fi12, v14)
                }
                _ = fi13
                _ = v14
            }
            v16 := []int{fi12}
            fmt.Println(v7, v1)
            _ = v16
        }
        v2 = v2
        _ = fi12
    }
    var v17 float64 = ((float64(0.0) + 3.14159) + v9)
    var v18 bool = false
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := []int{fi19, 2147483647, (fi19 % (((fi19 << (((v2 | v2)) & 63))) | 1)), (fi19 | 9223372036854775807)}
        v21 := GenS3{F0: 127, F1: 255, F2: (v3 >> ((v2) & 63))}
        _ = fi19
        _ = v20
        _ = v21
    }
    v7 = (v8 << (((v2 >> ((v2) & 63))) & 63))
    var v22 string = "hello"
    v2 = 1
    delete(v4, 961)
    v7 -= (v8 & v8)
    v23 := v4[0]
    v24 := GenS1{F0: (v5 | (v5 % ((v5) | 1))), F1: (v1 ^ v1)}
    if (float64(0.1) >= v9) {
        v22 = (v22 + v22)
        v4[uint32(v1)] = (^(v23))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
