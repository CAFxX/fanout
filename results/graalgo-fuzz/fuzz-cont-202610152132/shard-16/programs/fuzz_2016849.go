package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8; F1 int8 }

func main() {
    var v1 int = 0
    var v2 float64 = 0.1
    v3 := make(map[int]string)
    v4 := []int{(v1 ^ (v1 >> ((65535) & 63))), v1, -786, -2147483648}
    v5 := GenS1{F0: 32767, F1: (true || true), F2: 42}
    var v6 uint8 = 255
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 1; fi8++ {
        _ = len(v4)
        v1 += 32768
        if true {
            var v9 int32 = 0
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(fi8, v1)
                _ = fi10
            }
            v2 += (1.5 + v2)
            _ = v9
        } else {
            v1 = -9223372036854775808
            v1 = (fi8 >> ((18446744073709551615) & 63))
        }
        _ = fi8
    }
    var v11 uint16 = 508
    var v12 uint8 = 81
    v13 := GenS2{F0: ((v1 >> ((62) & 63)) < (-(v1))), F1: (v7 % (((v7 - v7)) | 1))}
    var v14 uint8 = 2
    v14 += (v6 ^ (v6 >> ((0) & 63)))
    var v15 float32 = 1.5
    v11 = v11
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := GenS3{F0: -32, F1: int8(v12)}
            v1 = int(v12)
            _ = fi17
            _ = v18
        }
        var v19 int16 = 1000
        _ = fi16
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
