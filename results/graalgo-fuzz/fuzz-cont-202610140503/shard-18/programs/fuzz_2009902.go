package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }

func main() {
    var v1 int = -69
    var v2 int64 = 1
    v3 := make(map[int]string)
    var v4 uint64 = 18446744073709551615
    var v5 int8 = 24
    var v6 float32 = 3.14159
    var v7 uint32 = 4294967295
    fmt.Println(v4)
    var v8 int64 = v2
    v9 := GenS1{F0: v5, F1: v7, F2: uint32(v4)}
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{fi10, v1, (fi10 - fi10), fi10}
        _ = fi10
        _ = v11
    }
    if (true && true) {
        v5 -= (v5 >> ((0) & 63))
    } else {
        v1 += v1
        v5 = v5
        v4 -= v4
    }
    var v12 bool = false
    if false {
        v3[(v1 % (((v1 * -704)) | 1))] = "a"
        v13 := GenS2{F0: 787, F1: 441, F2: 128}
        var v14 int16 = 1000
        _ = v13
        _ = v14
    }
    if (!("a" <= ("a" + "\n\t"))) {
        var v15 uint32 = (v7 ^ v7)
        _ = v15
    } else {
        v16 := []int{v1, (v1 % ((v1) | 1)), v1}
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(fi17)
            switch 128 {
            case -1:
                v7 -= 42
            case 1:
                fmt.Println(v4, v1)
            case 4294967295:
                if ((v7 == (^(v7))) || v12) {
                    v5 = -1
                    var v18 int32 = int32(v7)
                    _ = v18
                }
            default:
                _ = len(v3)
            }
            v19 := GenS2{F0: -312, F1: 4294967295, F2: 151}
            _ = fi17
            _ = v19
        }
        _ = v16
    }
    v2 = 9223372036854775807
    var v20 int = (v1 | (v1 << ((0) & 63)))
    var v21 float64 = 0.0
    switch v1 {
    case 65:
        var v22 uint = 2
        _ = v22
    default:
        var v23 int32 = -270
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
