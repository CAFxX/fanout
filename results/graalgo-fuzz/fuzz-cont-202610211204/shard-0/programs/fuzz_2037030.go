package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int16; F1 uint; F2 uint64 }
func genFn0(p0 string) int8 {
    switch 1 {
    case 778:
        if false {
            fmt.Println(p0)
            p0 = "\n\t"
            v4 := GenS1{F0: 2147483647, F1: 42}
            _ = v4
        } else {
            if true {
                var v5 int8 = 42
                var v6 string = "hello"
                switch 128 {
                case 176:
                    fmt.Println(v6, v5, p0)
                default:
                    var v7 uint16 = 245
                    _ = v7
                }
                _ = v5
                _ = v6
            }
            p0 = (p0 + (p0 + p0))
        }
    case -408:
        var v8 uint64 = 18446744073709551615
        _ = v8
    }
    var v9 uint8 = 202
    v10 := GenS2{F0: 42, F1: 100}
    if false {
        v11 := GenS3{F0: -2147483648, F1: (false || (!true))}
        _ = v11
    } else {
        if (true && (-2147483648 > 127)) {
            v12 := GenS4{F0: 4294967295, F1: 40}
            _ = v12
        } else {
            fmt.Println(v9, p0)
            v13 := []int{-9223372036854775808, 802, 42, -9223372036854775808}
            _ = v13
        }
        v14 := GenS3{F0: 0, F1: false}
        v15 := GenS5{F0: 1000, F1: 0, F2: 4294967295}
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 1, F1: 42}
    _ = p0
    _ = v10
    _ = v16
    _ = v9
    return -1
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{0}
    v3 := []int{128}
    v1 = (v1 + genFn0("foo bar"))
    fmt.Println(v1)
    var v17 string = "a"
    v1 += (-(v1))
    var v18 uint8 = 13
    var v19 int8 = 42
    if (!(42 <= 255)) {
        v18 = (v18 ^ (^(v18)))
        var v20 uint64 = 568
        _ = v20
    }
    if (!true) {
        fmt.Println(v19, v19)
    } else {
        if (562 > 42) {
            var v21 uint = 190
            _ = v21
        }
        _ = len(v2)
        var v22 float32 = (float32(1e10) + -2.25)
        _ = v22
    }
    var v23 bool = true
    if v23 {
        var v24 uint8 = v18
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 uint32 = 4294967295
            fmt.Println(v19)
            _ = fi25
            _ = v26
        }
        fmt.Println(v24)
        _ = v24
    }
    v1 += 0
    var v27 bool = false
    v28 := GenS1{F0: 1, F1: 340}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
}
