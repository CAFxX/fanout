package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 bool; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    var v1 int32 = 1
    var v2 int = 1
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v5 int32 = -840
    v6 := GenS1{F0: 0}
    var v7 int16 = -986
    var v8 string = ("hello" + (("\n\t" + "a") + "hello"))
    var v9 bool = false
    fmt.Println(v2)
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS2{F0: ((^(v7)) >= v7), F1: v5}
        v12 := GenS3{F0: false, F1: 4294967295, F2: 42}
        v13 := GenS3{F0: v9, F1: 820, F2: -28}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint8 = 0
    v2 -= -2147483648
    if true {
        delete(v4, 1)
    } else {
        if true {
            fmt.Println(v7, v5)
        } else {
            var v15 int64 = 1
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 float32 = (1e10 * 0.1)
                _ = fi16
                _ = v17
            }
            var v18 uint64 = 18446744073709551615
            _ = v15
            _ = v18
        }
    }
    if v9 {
        if (73 != 2) {
            fmt.Println(v8, v9)
        } else {
            v19 := v4[2147483648]
            for fi20 := 0; fi20 < 7; fi20++ {
                switch (v2 ^ fi20) {
                case 32767:
                    v21 := GenS4{F0: 4294967295, F1: 1}
                    _ = v21
                case 0:
                    var v22 uint16 = 0
                    _ = v22
                }
                _ = fi20
            }
            v5 += 1
            _ = v19
        }
    } else {
        fmt.Println(v2, v2, v14)
    }
    fmt.Println(v8, v9, v14)
    var v23 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
