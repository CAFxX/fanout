package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint8 = 255
    v2 := []int{261, 0}
    v3 := make(map[int]int)
    var v4 float64 = (float64(1.5) + (-2.25 - 0.1))
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 float64 = v4
        _ = fi5
        _ = v6
    }
    v4 -= (v4 + v4)
    v4 = v4
    v7 := GenS1{F0: 0, F1: -425}
    var v8 uint = 2
    fmt.Println(v8)
    v4 = 0.1
    for fi9 := 0; fi9 < 12; fi9++ {
        if (!(true || (!true))) {
            if (1000 <= 0) {
                v8 = (v8 - v8)
                var v10 int = int(v8)
                _ = v10
            }
            var v11 bool = (!false)
            fmt.Println(fi9)
            _ = v11
        } else {
            var v12 uint32 = 2147483647
            if true {
                switch fi9 {
                case -9223372036854775808:
                    v12 = 83
                }
                v13 := GenS1{F0: 255, F1: -2147483648}
                _ = v13
            } else {
                v2 = append(v2, (fi9 / ((-9223372036854775808) | 1)))
                var v14 int16 = 32767
                var v15 uint8 = (v1 / ((0) | 1))
                _ = v14
                _ = v15
            }
            v16 := GenS1{F0: 496, F1: -1}
            _ = v12
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 float32 = (-2.25 + (0.1 + 0.1))
            _ = fi17
            _ = v18
        }
        var v19 uint16 = 65535
        _ = fi9
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
