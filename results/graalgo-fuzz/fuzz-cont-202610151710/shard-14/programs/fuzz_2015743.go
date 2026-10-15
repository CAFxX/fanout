package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint8 }

func main() {
    v1 := []uint32{0, 891}
    var v2 float64 = (1.5 * (1.5 / (1e10 / 1.0)))
    var v3 uint8 = 2
    var v4 float32 = float32(v2)
    var v5 string = (("" + "a") + "")
    _ = len(v1)
    v6 := GenS1{F0: 805, F1: 171, F2: v3}
    v4 += ((v4 + (v4 * v4)) * float32(((3.14159 / v2) * (3.14159 / float64(0.1)))))
    var v7 float32 = -2.25
    var v8 int64 = -1
    v5 = v5
    _ = len(v1)
    var v9 float32 = v7
    if (-817 == 1000) {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch "a" {
            case "hello":
                fmt.Println(v7)
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                v11 := GenS1{F0: 601, F1: 255, F2: v3}
                _ = v11
            default:
                var v12 uint16 = 0
                _ = v12
            }
            _ = fi10
        }
    }
    v3 += v3
    v9 = 0.0
    var v13 int8 = 1
    var v14 bool = true
    var v15 uint64 = 0
    switch v5 {
    case "\n\t":
        var v16 int32 = 118
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
