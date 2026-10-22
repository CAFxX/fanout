package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int16; F2 int32 }

func main() {
    var v1 uint8 = 182
    var v2 int = 462
    var v3 uint16 = 65535
    var v4 int = v2
    var v5 uint64 = 0
    if ((-(v2)) >= v4) {
        switch 32768 {
        case 0:
            var v6 uint32 = 2
            _ = v6
        default:
            v3 = 2
        }
        v7 := GenS1{F0: -706, F1: 825, F2: true}
        _ = v7
    }
    switch v2 {
    case 2147483647:
        v8 := []int{(v4 & 720), 2147483647}
        _ = v8
    }
    v5 = (v5 * (v5 * v5))
    var v9 int16 = 1
    v10 := GenS2{F0: 42}
    if ((true && false) && true) {
        v1 -= v1
    } else {
        v11 := GenS3{F0: 1}
        switch (v4 ^ v2) {
        case 4294967295:
            var v12 uint16 = v3
            _ = v12
        case 42:
            var v13 uint64 = v5
            _ = v13
        }
        switch "a" {
        case "a":
            if (false && false) {
                v14 := GenS4{F0: 128, F1: v9, F2: 20}
                _ = v14
            }
        default:
            switch (v2 / ((v2) | 1)) {
            case -9223372036854775808:
                var v15 int32 = 2147483647
                _ = v15
            default:
                fmt.Println(v9, v3)
            }
        }
        _ = v11
    }
    var v16 uint32 = 2147483647
    v17 := []int{-453}
    var v18 bool = true
    var v19 float32 = 0.1
    v3 = (v3 >> ((296) & 63))
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 int32 = 2147483647
        var v22 int8 = 75
        v22 = v22
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 bool = v18
    var v24 uint16 = 65535
    var v25 uint64 = 910
    var v26 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
