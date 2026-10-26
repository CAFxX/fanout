package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }

func main() {
    var v1 int32 = -747
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 uint16 = 842
    if (true && true) {
        if (42 == 895) {
            for fi5 := 0; fi5 < 20; fi5++ {
                v4 -= v4
                if true {
                    v1 = -1
                    v4 = v4
                }
                _ = fi5
            }
        }
        v6 := []int{-784, 992}
        _ = v6
    }
    var v7 uint16 = 32767
    v1 = v1
    _ = len(v2)
    v8 := GenS1{F0: -278, F1: true}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint8 = 42
        switch 301 {
        case -2147483648:
            if true {
                var v11 int16 = 1000
                var v12 uint = 866
                var v13 int = (fi9 | fi9)
                _ = v11
                _ = v12
                _ = v13
            }
        case 42:
            if ((-(v1)) == v1) {
                var v14 uint16 = 167
                v10 = (v10 | v10)
                var v15 uint16 = (v14 | (v4 & (v4 ^ 21)))
                _ = v14
                _ = v15
            }
        case 551:
            fmt.Println(v1, v10)
        default:
            fmt.Println(v10)
        }
        var v16 int16 = -1
        _ = fi9
        _ = v10
        _ = v16
    }
    fmt.Println(v7, v1)
    var v17 uint64 = 4294967295
    fmt.Println(v7, v4, v4)
    v4 = v4
    if false {
        var v18 int32 = 42
        _ = v18
    }
    v19 := GenS2{F0: v4, F1: 1, F2: -317}
    if (65535 >= 755) {
        v20 := []int{}
        _ = v20
    } else {
        var v21 int64 = 42
        v22 := GenS2{F0: 42, F1: -128, F2: v21}
        var v23 float32 = 1.5
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 1, F1: 1, F2: -9223372036854775808}
    var v25 int8 = 0
    switch -785 {
    case 9223372036854775807:
        v26 := []int{-623}
        _ = v26
    default:
        fmt.Println(v4)
    }
    if false {
        var v27 int32 = v1
        _ = v27
    } else {
        v28 := GenS1{F0: 42, F1: false}
        v29 := GenS1{F0: -693, F1: true}
        v30 := GenS2{F0: v7, F1: v25, F2: -873}
        _ = v28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
