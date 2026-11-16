package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 uint64) uint16 {
    if true {
        v4 := GenS1{F0: 42, F1: 1}
        _ = v4
    } else {
        var v5 uint8 = 43
        v5 += (v5 | (v5 * (^(v5))))
        for fi6 := 0; fi6 < 18; fi6++ {
            p1 = p1
            _ = fi6
        }
        _ = v5
    }
    fmt.Println(p1, p1)
    var v7 string = "hello"
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1(p0 uint64, p1 int16, p2 int64) int {
    v8 := GenS2{F0: 1}
    p2 += (p2 & p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 int8 = -5
    var v2 uint64 = 0
    var v3 uint = 2
    v2 -= 18446744073709551615
    var v9 uint = 2
    v10 := GenS3{F0: 869}
    v11 := []int{-9223372036854775808, -1}
    v12 := GenS3{F0: 20}
    v1 = (v1 & 1)
    var v13 float32 = 3.14159
    switch 4294967295 {
    case 42:
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint16 = genFn0((fi14 << (((v3 / (((v3 >> (((v3 - 255)) & 63))) | 1))) & 63)), (v2 - 749))
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint = 2
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    default:
        fmt.Println(v9)
    }
    switch 4294967295 {
    case 32767:
        fmt.Println(v1)
    default:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
