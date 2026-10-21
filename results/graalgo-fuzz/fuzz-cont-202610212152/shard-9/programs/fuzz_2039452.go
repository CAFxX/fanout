package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int8; F2 uint64 }
func genFn0(p0 int64, p1 uint) uint64 {
    v7 := GenS1{F0: 387}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint64 = 1
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 4294967295}
    v11 := GenS2{F0: -968}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    return 4294967295
}
func genFn1() int32 {
    var v12 uint64 = 34
    if (-47 > 0) {
        var v13 int8 = 30
        if true {
            v14 := GenS3{F0: 0}
            v15 := []int{367}
            var v16 int32 = -1
            _ = v14
            _ = v15
            _ = v16
        }
        v13 = v13
        _ = v13
    }
    v12 += (^(v12))
    var v17 bool = false
    _ = v12
    _ = v17
    return 0
}

func main() {
    var v1 int = 42
    var v2 int16 = -31
    v3 := make(map[uint32]bool)
    var v4 uint = 42
    var v5 int32 = -2147483648
    v6 := []uint32{}
    v5 = v5
    for fi18 := 0; fi18 < 8; fi18++ {
        v4 += v4
        if true {
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 int16 = -985
                v3[2] = true
                _ = fi19
                _ = v20
            }
            var v21 float32 = 1e10
            v22 := v3[42]
            _ = v21
            _ = v22
        }
        v23 := GenS4{F0: (^(v4)), F1: 0, F2: 42}
        _ = fi18
        _ = v23
    }
    fmt.Println(v5, v2, v4)
    fmt.Println(v1)
    switch "hello" {
    case "hello":
        v2 = (v2 >> ((1) & 63))
    case "":
        var v24 int32 = v5
        _ = v24
    }
    v1 = (v1 ^ v1)
    var v25 float32 = 0.1
    v25 -= 0.0
    delete(v3, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
