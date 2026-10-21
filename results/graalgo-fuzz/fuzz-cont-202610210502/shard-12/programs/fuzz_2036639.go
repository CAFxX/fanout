package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]string)
    var v3 uint = 65535
    var v4 string = "hello"
    var v5 int = -493
    v6 := make(map[int]string)
    var v7 int32 = 1
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(v1)
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
        var v10 uint8 = 255
        _ = fi8
        _ = v10
    }
    v3 -= 0
    switch v5 {
    case 207:
        if false {
            v11 := []int{-1, v5}
            _ = v11
        } else {
            delete(v6, v5)
            switch (("" + v4) + ((("\n\t" + v4) + "") + v4)) {
            case "\n\t":
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := GenS1{F0: uint16(v3), F1: (262 <= 2147483647)}
                    var v14 float64 = (float64(0.0) / 3.14159)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            default:
                v5 += (v5 | v5)
            }
        }
    default:
        var v15 float64 = 3.14159
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := []int{v5, (fi16 & 4294967295), fi16}
        _ = fi16
        _ = v17
    }
    var v18 bool = (!true)
    switch -834 {
    case 0:
        _ = len(v2)
    case 31:
        v18 = ((v4 < v4) && (0.0 == float32(0.0)))
    case -1:
        switch 0 {
        case 2147483647:
            var v19 int16 = 0
            _ = v19
        case 32767:
            var v20 float64 = 1.5
            _ = v20
        }
    default:
        v7 += (v7 | v7)
    }
    var v21 int32 = (v7 & (v7 ^ (v7 / (((v7 ^ 995)) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
