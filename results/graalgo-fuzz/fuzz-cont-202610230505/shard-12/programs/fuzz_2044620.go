package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 int32, p1 string) int32 {
    p0 += p0
    var v6 int64 = 642
    _ = p0
    _ = p1
    _ = v6
    return 8
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1000
    var v3 bool = false
    var v4 uint = 42
    v5 := make(map[int]uint32)
    var v7 uint8 = 255
    var v8 int = 32767
    v8 += v8
    switch v8 {
    case 128:
        switch v8 {
        case 32767:
            v9 := GenS1{F0: v7}
            _ = v9
        }
    case 4294967295:
        fmt.Println(v7)
    case 127:
        var v10 uint16 = 657
        _ = v10
    default:
        switch "hello" {
        case "hello":
            var v11 uint = v4
            _ = v11
        case "\n\t":
            v3 = (!(v3 && true))
        case "":
            v12 := []int{}
            _ = v12
        }
    }
    var v13 float32 = 3.14159
    var v14 uint8 = v7
    var v15 uint = v4
    var v16 uint64 = 1
    v3 = (32767 <= 65535)
    if false {
        v17 := GenS1{F0: (v7 & 112)}
        var v18 uint16 = 42
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: (!(0.1 == (0.0 - ((0.0 * -2.25) + float64(1.5)))))}
    fmt.Println(v7, v2)
    fmt.Println(v7)
    v20 := GenS3{F0: 100, F1: false}
    var v21 int64 = 233
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
