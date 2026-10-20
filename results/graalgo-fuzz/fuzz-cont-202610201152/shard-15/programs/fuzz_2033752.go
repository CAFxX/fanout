package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 string = ("hello" + ("hello" + ("foo bar" + "")))
    var v2 uint32 = 639
    var v3 bool = (0 > 696)
    var v4 uint32 = (v2 * v2)
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 255
    v7 := []int{32767, 0, -2147483648}
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 5; fi9++ {
        v8 -= -2.25
        for fi10 := 0; fi10 < 16; fi10++ {
            if (-130 > 42) {
                if v3 {
                    var v11 int16 = 42
                    v12 := GenS1{F0: -532, F1: v6, F2: 110}
                    _ = v11
                    _ = v12
                }
                var v13 uint64 = 42
                var v14 int16 = 0
                _ = v13
                _ = v14
            }
            v8 = (-2.25 - (v8 * v8))
            v15 := GenS2{F0: uint16(v6)}
            _ = fi10
            _ = v15
        }
        _ = fi9
    }
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    var v16 uint64 = 4294967295
    var v17 uint64 = uint64(v4)
    v18 := GenS1{F0: -906, F1: v6, F2: 0}
    var v19 int8 = -1
    fmt.Println(v8)
    v20 := GenS2{F0: 259}
    v21 := GenS3{F0: int16(v6)}
    var v22 string = ((v1 + v1) + "a")
    v8 = (v8 + float64(1.5))
    v23 := GenS3{F0: 1}
    var v24 int8 = v19
    v25 := GenS4{F0: v6}
    var v26 int32 = -868
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
