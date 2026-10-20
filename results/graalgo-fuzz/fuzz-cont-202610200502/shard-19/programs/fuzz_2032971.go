package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int8, p1 uint64, p2 int) uint {
    var v4 uint = 61
    var v5 uint32 = 448
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (^(v4))
}

func main() {
    var v1 int16 = -1
    var v2 uint32 = 4294967295
    var v3 uint64 = 2
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 1, F1: 2}
        _ = fi6
        _ = v7
    }
    var v8 int32 = 1
    v8 -= (v8 | v8)
    var v9 int16 = v1
    if true {
        v8 = v8
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 0, F1: 2}
            var v12 int8 = -128
            fmt.Println(v12, fi10)
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int16 = -1
        _ = v13
    } else {
        v2 -= (v2 % ((v2) | 1))
        v3 = v3
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int8 = 114
            if true {
                var v16 int8 = (-(v15))
                v17 := []int{-2147483648, (fi14 ^ 128), fi14}
                _ = v16
                _ = v17
            }
            v18 := []int{fi14}
            _ = fi14
            _ = v15
            _ = v18
        }
    }
    fmt.Println(v2)
    var v19 int64 = 1
    var v20 int = 9223372036854775807
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := []int{fi21, v20, (-(v20))}
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v24 := []int{}
    v8 += (v8 & v8)
    v24 = append(v24, (v20 ^ (v20 | (v20 + 1))))
    var v25 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v8
    _ = v9
}
