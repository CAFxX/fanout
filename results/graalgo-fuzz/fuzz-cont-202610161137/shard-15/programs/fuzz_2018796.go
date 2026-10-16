package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 4294967295
    var v3 uint8 = 2
    if (true || (!false)) {
        v3 = (v3 / ((v3) | 1))
        v1[-725] = -68
        fmt.Println(v2)
    }
    _ = len(v1)
    fmt.Println(v3)
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{fi4}
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 0, F1: true, F2: false}
    fmt.Println(v3)
    fmt.Println(v3, v2)
    var v7 float64 = 1.5
    v8 := []int{1, 100, 32767}
    v3 -= (v3 >> ((42) & 63))
    for fi9 := 0; fi9 < 9; fi9++ {
        if (4294967295 == 269) {
            delete(v1, 2147483647)
        } else {
            var v10 int32 = 1
            _ = v10
        }
        fmt.Println(v7)
        _ = fi9
    }
    var v11 int = 32767
    v3 = (v3 % ((uint8(v2)) | 1))
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if true {
                var v14 float32 = (float32(v7) - (0.1 + 1e10))
                _ = v14
            } else {
                v15 := GenS2{F0: (v3 - (v3 & v3)), F1: "\n\t"}
                var v16 uint = 65535
                v8 = append(v8, fi12)
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            v7 += (((v7 + v7) + v7) / v7)
            _ = fi17
        }
        v3 += (v3 + (v3 + v3))
        _ = fi12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
