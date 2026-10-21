package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0() uint64 {
    v6 := []int{-2147483648}
    var v7 bool = (-560 > 42)
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int = -1
    var v2 string = (("\n\t" + "") + "")
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 65535
    var v5 int64 = 238
    v8 := []int{(v1 << ((36) & 63))}
    var v9 uint32 = 1
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v10 uint16 = 42
    var v11 uint8 = 1
    v1 -= 838
    v2 = "hello"
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint8 = (^(v11))
            v15 := GenS1{F0: 32767, F1: v10}
            if true {
                fmt.Println(v14, v1)
                v16 := GenS1{F0: int(v10), F1: v10}
                fmt.Println(v9, v2, v5)
                _ = v16
            } else {
                if ((1 <= -1) || false) {
                    v17 := GenS2{F0: 619, F1: v10}
                    var v18 int8 = 42
                    var v19 float32 = 3.14159
                    fmt.Println(fi12, v10)
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    if false {
                        v9 = (v9 & uint32(v11))
                        v20 := GenS1{F0: -2147483648, F1: v10}
                        _ = v20
                    } else {
                        var v21 int8 = 95
                        _ = v21
                    }
                    switch fi12 {
                    case 128:
                        var v22 float64 = float64(1.5)
                        _ = v22
                    case -9223372036854775808:
                        v23 := GenS3{F0: uint8(v10)}
                        var v24 int64 = (v5 * v5)
                        _ = v23
                        _ = v24
                    }
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        v11 -= (v11 & v11)
        v5 -= (v5 - v5)
        _ = fi12
    }
    v25 := GenS4{F0: (v9 + (^(v4))), F1: false}
    var v26 int = v1
    var v27 uint64 = 395
    v5 -= 866
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
