package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
func genFn0(p0 int, p1 uint) uint8 {
    fmt.Println(p1, p1)
    v7 := GenS1{F0: -517, F1: -57, F2: 42}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1(p0 uint16) uint {
    v8 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 450
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1e10
    var v4 int32 = 500
    var v5 float64 = (0.0 * 1.5)
    v6 := make(map[uint32]bool)
    v9 := GenS1{F0: 32767, F1: 0, F2: 42}
    if false {
        v10 := GenS1{F0: 1000, F1: (v4 | 0), F2: -487}
        _ = v10
    } else {
        var v11 uint32 = 531
        var v12 uint32 = v11
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS2{F0: (1 > (v4 ^ v4))}
        _ = fi13
        _ = v14
    }
    v5 = (((v5 + 3.14159) - (((v5 - float64(v3)) - (v5 * v5)) * 1.5)) / v5)
    switch -534 {
    case 364:
        v3 = -2.25
    case 500:
        var v15 uint32 = 65535
        _ = v15
    case 128:
        v4 -= (^(v4))
    default:
        v16 := []int{-9223372036854775808, -429, 218}
        _ = v16
    }
    fmt.Println(v3, v4, v5)
    var v17 int16 = -1
    var v18 int16 = (v17 ^ (-(v17)))
    v19 := GenS2{F0: true}
    switch 362 {
    case 127:
        var v20 uint8 = 255
        _ = v20
    case 0:
        v21 := GenS3{F0: 128, F1: 544, F2: 571}
        _ = v21
    case 9223372036854775807:
        v4 += (-(v4))
    }
    if (1 <= 136) {
        var v22 uint16 = 728
        v17 = (v17 / (((v18 - 800)) | 1))
        if true {
            switch "foo bar" {
            case "hello":
                if true {
                    v22 = v22
                } else {
                    var v23 uint = 195
                    _ = v23
                }
            case "foo bar":
                v24 := []int{-881, 212, -435}
                _ = v24
            }
        }
        _ = v22
    } else {
        var v25 int8 = 42
        delete(v6, 838)
        var v26 float64 = (v5 / 1.0)
        _ = v25
        _ = v26
    }
    switch ("a" + "hello") {
    case "foo bar":
        fmt.Println(v18, v5, v17)
    default:
        v27 := GenS3{F0: 815, F1: 42, F2: 2147483648}
        _ = v27
    }
    var v28 string = "hello"
    var v29 uint16 = 355
    var v30 float32 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
