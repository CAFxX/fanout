package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0() int64 {
    var v5 int8 = 100
    if false {
        v6 := []int{-2147483648, 4294967295}
        if true {
            v7 := GenS1{F0: 2, F1: -1, F2: 732}
            _ = v7
        }
        var v8 uint16 = 965
        _ = v6
        _ = v8
    }
    var v9 int16 = -1
    _ = v5
    _ = v9
    return 42
}

func main() {
    var v1 float32 = float32(3.14159)
    var v2 uint8 = 255
    var v3 uint32 = 2147483647
    var v4 uint16 = 101
    v2 += (v2 | v2)
    fmt.Println(v2)
    v10 := []int{32767, 128}
    var v11 int32 = 0
    v4 = 1
    switch "hello" {
    case "a":
        var v12 uint = 1
        _ = v12
    case "":
        v3 = (v3 ^ v3)
    case "hello":
        v2 = (^(v2))
    default:
        v10 = append(v10, -9223372036854775808)
    }
    v13 := GenS1{F0: (^(v2)), F1: 0, F2: (v4 & (^(v4)))}
    switch 1 {
    case 0:
        v10 = append(v10, -177)
    case 246:
        v14 := GenS1{F0: v2, F1: -1, F2: (^(v4))}
        _ = v14
    }
    v15 := GenS1{F0: (v2 ^ (v2 / (((v2 | v2)) | 1))), F1: -1, F2: v4}
    var v16 uint32 = v3
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
