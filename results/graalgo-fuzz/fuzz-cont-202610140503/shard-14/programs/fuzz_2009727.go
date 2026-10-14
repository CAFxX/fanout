package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
func genFn0(p0 int32, p1 bool) uint {
    var v6 int16 = 247
    v7 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 237
}
func genFn1(p0 int, p1 int) uint16 {
    switch p0 {
    case 42:
        v8 := []int{42, (p0 ^ 32768), p0, (p0 ^ -677)}
        _ = v8
    case 0:
        p0 = p0
    case 127:
        v9 := []int{}
        _ = v9
    }
    var v10 int = (p0 - p0)
    var v11 float64 = (float64(0.1) + 1.5)
    var v12 uint8 = 255
    var v13 int16 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 694
    var v3 uint32 = 777
    var v4 float32 = 1e10
    v5 := []int{-2147483648, 1}
    v14 := GenS1{F0: 0, F1: 0, F2: ""}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint64 = v2
        _ = len(v5)
        for fi17 := 0; fi17 < 11; fi17++ {
            v5 = append(v5, (fi15 * fi15))
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    v3 = (v3 * 2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v18 uint32 = (v3 + v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v19 float64 = 1.5
    var v20 uint16 = 65535
    var v21 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
