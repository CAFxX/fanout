package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0() int16 {
    var v4 int = -586
    var v5 uint64 = 513
    _ = v4
    _ = v5
    return -442
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -9223372036854775808
    var v3 uint64 = v1
    v2 = v2
    var v6 uint32 = 2
    var v7 int32 = -1
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int16 = 1
            var v10 uint16 = 255
            var v11 int16 = genFn0()
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    } else {
        v1 = v3
        switch -153 {
        case 128:
            var v12 int16 = -315
            _ = v12
        case 32767:
            v13 := []int{42}
            _ = v13
        }
        fmt.Println(v1)
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{(fi14 & (fi14 + (fi14 | 32768))), fi14}
        _ = fi14
        _ = v15
    }
    v16 := []int{-924}
    v7 -= (v7 % ((-1) | 1))
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS1{F0: -60, F1: v3}
        _ = fi17
        _ = v18
    }
    var v19 float64 = 0.0
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 int8 = 72
        _ = fi20
        _ = v21
    }
    var v22 float32 = 1e10
    v23 := GenS1{F0: 0, F1: (v1 * (v3 + (v1 / ((v1) | 1))))}
    v22 = v22
    fmt.Println(v22)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
