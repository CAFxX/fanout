package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = false
    var v3 uint32 = 42
    var v4 uint16 = 1
    var v5 float64 = 3.14159
    v6 := make(map[int]string)
    var v7 uint = 851
    var v8 uint64 = (v1 | v1)
    var v9 uint64 = (^(v1))
    switch ("hello" + ("" + ("hello" + "hello"))) {
    case "":
        if (!v2) {
            v6[int(v4)] = ""
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        fmt.Println(fi12)
        if false {
            v13 := GenS1{F0: (^(v7))}
            v14 := GenS2{F0: v4, F1: uint8(v9)}
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(v3, v3, v3)
                if true {
                    switch "" {
                    case "foo bar":
                        var v16 string = "a"
                        v9 = (v9 - (v1 << ((v7) & 63)))
                        _ = v16
                    default:
                        v17 := v6[fi12]
                        _ = v17
                    }
                } else {
                    v4 = v4
                    fmt.Println(v1, v8)
                }
                _ = fi15
            }
            v18 := v6[(fi12 ^ fi12)]
            var v19 int64 = -1
            _ = v18
            _ = v19
        }
        _ = fi12
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if (v9 == 1) {
            if (179 <= 72) {
                var v21 int64 = 42
                for fi22 := 0; fi22 < 13; fi22++ {
                    if true {
                        v21 = (v21 >> ((v7) & 63))
                    }
                    var v23 int32 = 0
                    delete(v6, 1)
                    if (1.5 == v5) {
                        v21 = (v21 * (v21 & (v21 >> ((v7) & 63))))
                        var v24 uint64 = v1
                        _ = v24
                    } else {
                        v9 -= 0
                        v21 += v21
                        v25 := GenS1{F0: v7}
                        _ = v25
                    }
                    if ((0.1 - v5) <= v5) {
                        delete(v6, int(v7))
                        var v26 int64 = (v21 + v21)
                        _ = v26
                    } else {
                        var v27 float32 = 0.1
                        v1 += (v1 - (v8 << (((v7 + (^(v7)))) & 63)))
                        _ = v27
                    }
                    v28 := GenS1{F0: (v7 & v7)}
                    var v29 int8 = -103
                    _ = fi22
                    _ = v23
                    _ = v28
                    _ = v29
                }
                _ = v21
            }
            _ = len(v6)
        }
        var v30 int16 = int16(v3)
        _ = fi20
        _ = v30
    }
    delete(v6, 128)
    v31 := GenS2{F0: v4, F1: 0}
    v5 = v5
    var v32 int = -971
    v33 := []int{(v32 % ((32768) | 1)), v32, v32}
    fmt.Println(v9)
    v7 -= (v7 >> ((v7) & 63))
    v32 += -453
    fmt.Println(v3)
    var v34 int32 = 2147483647
    v35 := GenS3{F0: v34, F1: v4, F2: 348}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
