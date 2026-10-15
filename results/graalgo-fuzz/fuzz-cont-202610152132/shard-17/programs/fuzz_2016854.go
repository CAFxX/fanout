package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int8; F2 uint8 }

func main() {
    var v1 uint16 = 32767
    var v2 bool = true
    v3 := make(map[uint32]int)
    var v4 uint8 = 111
    v5 := make(map[string]int)
    v6 := []int{-2147483648}
    v6 = append(v6, -2147483648)
    v4 -= v4
    var v7 int32 = -1
    v3[0] = 0
    v8 := GenS1{F0: 132, F1: 9223372036854775807}
    v1 -= v1
    switch 379 {
    case 1:
        if v2 {
            _ = len(v5)
            fmt.Println(v7, v2, v2)
            v1 = (v1 % ((0) | 1))
        } else {
            var v9 uint16 = v1
            _ = v9
        }
    case 32768:
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 string = "\n\t"
            _ = len(v6)
            if v2 {
                var v12 float32 = 1e10
                fmt.Println(v7, v7)
                _ = v12
            } else {
                var v13 string = (v11 + v11)
                var v14 uint32 = 2
                v15 := GenS2{F0: "\n\t"}
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
        }
    case 4294967295:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        v2 = v2
    }
    fmt.Println(v7, v7)
    var v16 int = 9223372036854775807
    v17 := GenS3{F0: 2}
    v18 := GenS4{F0: 4294967295, F1: -77, F2: 127}
    var v19 int = (v16 * (v16 + (v16 + (v16 % ((v16) | 1)))))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
