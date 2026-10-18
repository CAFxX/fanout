package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint16, p1 bool) int8 {
    v5 := []int{127, 379}
    _ = len(v5)
    p0 += p0
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    var v1 uint = 255
    var v2 uint64 = 614
    var v3 float32 = float32((0.0 + 3.14159))
    v4 := make(map[uint32]string)
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 = 0.1
        v1 = (v1 ^ (^(v1)))
        v2 = v2
        _ = fi6
    }
    var v7 uint64 = (^(v2))
    v1 = (^(v1))
    v8 := GenS1{F0: 1, F1: 0, F2: (-2.25 < 0.1)}
    var v9 int64 = 212
    v10 := GenS1{F0: 189, F1: 1, F2: true}
    switch "" {
    case "\n\t":
        if (((("\n\t" + "a") + "hello") + ("\n\t" + "a")) < ("" + "hello")) {
            v11 := GenS2{F0: -2147483648, F1: -835}
            _ = v11
        }
    case "hello":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint32 = 2147483648
            v14 := []int{42, int(v1), (fi12 ^ int(v1)), -775}
            v15 := GenS3{F0: 2, F1: (v13 * (v13 | (v13 % (((v13 - (v13 - (v13 / (((v13 + 1)) | 1))))) | 1))))}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    default:
        var v16 float32 = (v3 * v3)
        _ = v16
    }
    v9 -= v9
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int16 = 32767
        v18 = (v18 & 32767)
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
