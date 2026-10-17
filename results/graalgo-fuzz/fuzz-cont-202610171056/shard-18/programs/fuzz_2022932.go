package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int; F1 bool; F2 int16 }
func genFn0(p0 int16, p1 int16) int64 {
    fmt.Println(p1, p0)
    fmt.Println(p0)
    if true {
        var v6 float32 = (1.5 + 3.14159)
        _ = v6
    }
    if false {
        fmt.Println(p0, p1)
        fmt.Println(p1, p1)
        var v7 uint = 42
        _ = v7
    }
    if (65535 == 815) {
        v8 := GenS1{F0: 2147483648, F1: 693, F2: 2147483648}
        if (!false) {
            v9 := GenS2{F0: 321, F1: 1}
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 int32 = 42
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                fmt.Println(p1)
                var v13 int = fi12
                v14 := GenS1{F0: 333, F1: 255, F2: 4294967295}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int8) uint64 {
    if (!(false || false)) {
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS3{F0: (^(fi15)), F1: false, F2: -32768}
            v17 := []int{fi15, (fi15 + fi15)}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 uint16 = 42
        p0 = (p0 & p0)
        _ = v18
    }
    var v19 uint = 1
    fmt.Println(p0, p0)
    _ = p0
    _ = v19
    return 492
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{2147483647, -2147483648, 9223372036854775807}
    var v3 int8 = -128
    var v4 int = -2147483648
    v5 := make(map[string]int)
    v20 := GenS2{F0: 255, F1: -12}
    fmt.Println(v3, v3)
    var v21 uint64 = 903
    if false {
        v21 = 18446744073709551615
    }
    _ = len(v5)
    v21 = v21
    var v22 float64 = 1.5
    var v23 float32 = 1e10
    var v24 int64 = 984
    if true {
        fmt.Println(v3)
        fmt.Println(v22, v21)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
