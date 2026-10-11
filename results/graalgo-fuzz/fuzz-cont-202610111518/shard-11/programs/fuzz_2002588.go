package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint, p1 int) int8 {
    fmt.Println(p1, p1)
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = v7
    return 3
}
func genFn1(p0 uint16, p1 int) uint {
    p0 -= (p0 - (p0 + 385))
    var v8 uint16 = (p0 * 0)
    _ = p0
    _ = p1
    _ = v8
    return 200
}

func main() {
    var v1 uint8 = 111
    v2 := []int{42, 128, 127, 4294967295}
    var v3 uint8 = 2
    v4 := make(map[int]bool)
    var v5 uint16 = 65535
    v6 := []uint32{}
    v9 := GenS1{F0: -858, F1: 19}
    var v10 uint8 = 1
    var v11 uint64 = 363
    v10 = (^(v3))
    v4[0] = (float32((((1e10 + 1.5) + 1.5) * -2.25)) >= ((((3.14159 - 0.1) - (1e10 * 3.14159)) / 1.0) - 3.14159))
    v12 := GenS2{F0: 42}
    if true {
        v13 := v4[16]
        _ = v13
    }
    v14 := GenS3{F0: -777}
    var v15 uint = 169
    if true {
        var v16 uint = genFn1((v5 >> (((v15 >> (((v15 << ((v15) & 63))) & 63))) & 63)), 32767)
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        if (278 > -482) {
            fmt.Println(fi17)
            v18 := GenS2{F0: (^(v1))}
            _ = v18
        }
        _ = fi17
    }
    v19 := GenS4{F0: 1}
    v20 := GenS1{F0: -1, F1: 1}
    fmt.Println(v3, v11, v1)
    v21 := v4[-67]
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
