package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int8 }

func main() {
    var v1 uint8 = 1
    var v2 uint16 = 179
    v3 := make(map[int]bool)
    v4 := []int{-1}
    v5 := GenS1{F0: 127, F1: 42, F2: 127}
    if false {
        var v6 string = "hello"
        _ = v6
    } else {
        if false {
            var v7 uint32 = 0
            _ = v7
        }
    }
    v1 += (v1 / ((0) | 1))
    if (1 > 932) {
        var v8 uint32 = 2147483647
        var v9 int16 = 1000
        _ = v8
        _ = v9
    }
    switch 32767 {
    case 2147483647:
        _ = len(v3)
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: fi10, F1: 1, F2: 0}
        if (-981 < -1) {
            switch "\n\t" {
            case "":
                v1 = 255
            default:
                if (253 == -1) {
                    var v12 int32 = -1
                    var v13 uint64 = uint64(v2)
                    v12 += v12
                    _ = v12
                    _ = v13
                } else {
                    v1 = (v1 / (((v1 * v1)) | 1))
                    var v14 float32 = -2.25
                    var v15 uint = 2
                    _ = v14
                    _ = v15
                }
            }
            fmt.Println(v2, v2)
            var v16 float64 = (1.5 / 0.1)
            _ = v16
        }
        _ = fi10
        _ = v11
    }
    v3[94] = ((v1 & v1) > (v1 - 42))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
