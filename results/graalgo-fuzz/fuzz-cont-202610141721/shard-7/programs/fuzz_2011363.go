package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint) uint {
    p0 += p0
    v6 := GenS1{F0: 2, F1: -1, F2: 0}
    if true {
        p0 -= (p0 & 964)
        if (0 <= 0) {
            fmt.Println(p0)
            p0 -= p0
        } else {
            var v7 int16 = 1000
            _ = v7
        }
        v8 := []int{-769}
        _ = v8
    }
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 int32 = 427
    v2 := make(map[int]int)
    var v3 bool = (!false)
    var v4 uint32 = 548
    v5 := []int{128}
    v9 := GenS1{F0: 1, F1: (v1 / ((v1) | 1)), F2: (v4 * 2)}
    v4 = (v4 | (^(v4)))
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint32 = 285
        if (42 > 823) {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint = 0
                v11 = v11
                fmt.Println(v13, v13)
                var v14 int8 = 127
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        v1 += 0
        _ = fi10
        _ = v11
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS2{F0: 255, F1: 1, F2: 1}
        v17 := GenS3{F0: 88, F1: 255}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = 0.0
    var v19 uint16 = 65535
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v3 {
        v19 = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
