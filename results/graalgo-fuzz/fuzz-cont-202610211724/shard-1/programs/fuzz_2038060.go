package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint64; F2 string }
func genFn0(p0 uint32, p1 int16, p2 string) string {
    fmt.Println(p2, p1)
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 54}
    var v8 uint64 = 2
    var v9 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p2 + "\n\t")
}
func genFn1() uint16 {
    var v10 uint = 2
    if true {
        v10 -= (v10 / (((v10 + 282)) | 1))
    }
    _ = v10
    return 325
}

func main() {
    var v1 int8 = 69
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v4 := []uint32{87, 250}
    v5 := make(map[uint32]string)
    var v6 uint64 = 18446744073709551615
    var v11 string = "\n\t"
    switch v11 {
    case "a":
        fmt.Println(v6)
    case "":
        v12 := GenS2{F0: 42, F1: v1}
        _ = v12
    case "foo bar":
        v13 := GenS3{F0: v6, F1: (0 > 2)}
        _ = v13
    }
    switch 32767 {
    case 4294967295:
        if len(v4) > 0 {
            v4[0] = 641
        }
    case -2147483648:
        var v14 float64 = (-2.25 - 1e10)
        _ = v14
    }
    var v15 uint64 = v6
    var v16 uint64 = (v15 | v15)
    v17 := GenS3{F0: (v15 | 18446744073709551615), F1: (2 > 6)}
    var v18 uint8 = 1
    v19 := GenS4{F0: v18}
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := GenS5{F0: false, F1: v6, F2: v11}
            v18 = (^(v18))
            var v23 bool = false
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := GenS2{F0: (v18 ^ 2), F1: v1}
        _ = fi20
        _ = v24
    }
    var v25 uint = 4294967295
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v16 += (v6 + (v6 & (v16 + v6)))
    if false {
        v18 = (v18 >> ((v25) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
