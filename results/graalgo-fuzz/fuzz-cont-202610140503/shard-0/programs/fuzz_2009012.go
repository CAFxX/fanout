package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float32 = (float32(float64(3.14159)) / -2.25)
    v2 := []int{}
    var v3 float64 = 0.0
    var v4 uint32 = 65535
    var v5 uint16 = 190
    var v6 uint32 = (v4 >> ((1) & 63))
    _ = len(v2)
    _ = len(v2)
    var v7 int16 = 817
    switch -1 {
    case 42:
        v6 = v4
    }
    v5 -= v5
    v4 += 0
    v7 += (v7 * v7)
    if true {
        var v8 float64 = (-2.25 + v3)
        _ = v8
    } else {
        v1 = (v1 + v1)
        if len(v2) > 0 {
            v2[0] = 127
        }
        if (!false) {
            if false {
                fmt.Println(v3)
                _ = len(v2)
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    var v9 string = ""
    _ = len(v2)
    var v10 int32 = 0
    _ = len(v2)
    if false {
        switch 4294967295 {
        case 2147483647:
            var v11 int16 = v7
            _ = v11
        case 42:
            if ((v3 - v3) >= (v3 * 0.0)) {
                v12 := GenS1{F0: 42, F1: 1}
                _ = v12
            } else {
                fmt.Println(v7, v5)
            }
        case 1:
            var v13 uint64 = 851
            _ = v13
        default:
            v7 += (v7 + v7)
        }
        var v14 uint8 = 2
        if false {
            switch 127 {
            case -2147483648:
                var v15 string = v9
                _ = v15
            default:
                var v16 uint64 = uint64(v5)
                _ = v16
            }
            v14 += (v14 - v14)
        }
        _ = v14
    } else {
        v17 := GenS2{F0: (1 >= v7)}
        var v18 uint = 255
        _ = v17
        _ = v18
    }
    fmt.Println(v7)
    v6 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
