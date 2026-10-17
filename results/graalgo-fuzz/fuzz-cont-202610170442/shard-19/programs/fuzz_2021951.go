package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 bool = false
    var v2 int = 2147483647
    var v3 uint16 = 1
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: uint32(v3)}
        _ = fi4
        _ = v5
    }
    fmt.Println(v1)
    var v6 float64 = (-2.25 / 1.0)
    v7 := []int{}
    var v8 bool = (!true)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint64 = 42
        fmt.Println(v10, fi9)
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int = -1
            if (-670 == 2147483647) {
                v10 = (v10 << ((658) & 63))
                if (!false) {
                    v7 = append(v7, fi9)
                    var v13 float64 = v6
                    v3 = uint16(v10)
                    _ = v13
                }
                fmt.Println(v1)
            } else {
                fmt.Println(v2)
                fmt.Println(v6)
                var v14 uint8 = 212
                _ = v14
            }
            var v15 int64 = 132
            _ = fi11
            _ = v12
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    v2 = (v2 ^ -2147483648)
    fmt.Println(v2, v8, v6)
    v16 := GenS2{F0: -547, F1: -419, F2: 1}
    var v17 int32 = 0
    v3 += (v3 - v3)
    v18 := GenS1{F0: 1}
    v19 := GenS3{F0: 1}
    for fi20 := 0; fi20 < 9; fi20++ {
        v2 -= 4294967295
        _ = fi20
    }
    v8 = v1
    v17 += int32(v3)
    var v21 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
