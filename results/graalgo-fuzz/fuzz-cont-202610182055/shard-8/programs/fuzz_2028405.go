package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }

func main() {
    var v1 uint32 = 997
    var v2 float64 = 0.1
    v3 := make(map[uint32]string)
    var v4 bool = true
    var v5 uint = 0
    var v6 uint = v5
    v5 -= (v5 ^ v5)
    v6 = v6
    if v4 {
        v7 := []int{204, 128, -279}
        _ = v7
    }
    var v8 int8 = 1
    if (0 > 18446744073709551615) {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: 0}
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 string = ""
                if v4 {
                    _ = len(v3)
                }
                v13 := []int{(fi11 ^ fi9), -801}
                var v14 int8 = v8
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 int8 = (v8 >> ((v6) & 63))
            _ = fi9
            _ = v10
            _ = v15
        }
        v8 = (v8 / ((v8) | 1))
        v16 := GenS2{F0: v1, F1: -1}
        _ = v16
    } else {
        v17 := GenS1{F0: 32767}
        fmt.Println(v6)
        _ = v17
    }
    v4 = false
    for fi18 := 0; fi18 < 3; fi18++ {
        _ = len(v3)
        _ = fi18
    }
    v19 := GenS3{F0: 0, F1: v8, F2: 42}
    for fi20 := 0; fi20 < 16; fi20++ {
        v4 = v4
        var v21 int16 = 32767
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: 1, F1: v8, F2: 333}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
