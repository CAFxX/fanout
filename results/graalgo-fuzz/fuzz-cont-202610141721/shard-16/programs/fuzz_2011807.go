package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16) uint32 {
    v7 := []int{}
    if ((p0 * 137) != p0) {
        if false {
            v8 := GenS1{F0: 100}
            _ = v8
        } else {
            fmt.Println(p0)
        }
    } else {
        var v9 int = 747
        v10 := GenS2{F0: 32767, F1: -1}
        v11 := GenS2{F0: 1000, F1: 405}
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint32 = 613
    v2 := make(map[uint32]bool)
    var v3 bool = false
    var v4 string = "hello"
    var v5 uint64 = 4294967295
    v6 := make(map[string]string)
    fmt.Println(v5, v5, v3)
    fmt.Println(v4, v4, v1)
    var v12 uint64 = v5
    delete(v6, ((v4 + v4) + "a"))
    var v13 bool = v3
    v14 := GenS1{F0: -48}
    switch v4 {
    case "":
        var v15 float64 = 0.0
        _ = v15
    default:
        var v16 uint16 = 32767
        _ = v16
    }
    var v17 uint8 = 128
    v1 = v1
    v18 := GenS3{F0: -1}
    var v19 string = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
