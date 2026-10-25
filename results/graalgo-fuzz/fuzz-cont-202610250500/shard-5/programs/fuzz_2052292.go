package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 float32 = 0.0
    v2 := []int{0, 127, -742}
    var v3 uint8 = 101
    var v4 float64 = float64((3.14159 - (v1 - 0.0)))
    v5 := []uint32{204, 42, 2, 913}
    v6 := GenS1{F0: -1}
    fmt.Println(v1, v1, v4)
    var v7 int16 = -32768
    fmt.Println(v7)
    v8 := GenS2{F0: true, F1: -654, F2: 9223372036854775807}
    var v9 int = -9223372036854775808
    var v10 uint = 0
    var v11 uint8 = (v3 + (v3 - v3))
    v2 = append(v2, v9)
    var v12 int32 = 1
    v2 = append(v2, (v9 % (((v9 & (v9 | 32768))) | 1)))
    _ = len(v5)
    for fi13 := 0; fi13 < 8; fi13++ {
        _ = len(v2)
        v7 += -1
        _ = fi13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 int64 = 129
        fmt.Println(v7, v11)
        _ = fi14
        _ = v15
    }
    v12 = (v12 - v12)
    if (v7 > int16(v10)) {
        switch 1 {
        case -2147483648:
            var v16 uint16 = 2
            _ = v16
        default:
            v17 := GenS3{F0: true}
            _ = v17
        }
        v18 := GenS4{F0: v3, F1: v9}
        _ = v18
    } else {
        if (!(!(v10 >= 144))) {
            v12 = v12
            switch "\n\t" {
            case "":
                v11 = 0
            case "a":
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 uint8 = (v3 >> ((v10) & 63))
                    _ = fi19
                    _ = v20
                }
            }
            var v21 uint8 = (v11 >> ((4294967295) & 63))
            _ = v21
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            v3 = 128
            fmt.Println(v9, v10, v4)
            _ = fi22
        }
    }
    fmt.Println(v7, v3)
    var v23 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
