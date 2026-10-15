package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0() uint64 {
    var v5 uint32 = 1
    v6 := []int{4294967295, 32767, 841}
    _ = v5
    _ = v6
    return 533
}
func genFn1(p0 uint, p1 uint16, p2 string) uint {
    var v7 bool = (true || ((!(true && false)) && true))
    var v8 float64 = float64(1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 919
    v3 := make(map[string]string)
    var v4 uint8 = 128
    v2 -= v2
    delete(v3, ("hello" + "hello"))
    if len(v1) > 0 {
        v1[0] = (v2 & v2)
    }
    var v9 int = 9223372036854775807
    v10 := GenS1{F0: genFn0(), F1: v9}
    if (false && true) {
        if ((false && (true || (-28 > 0))) || (false || (42 < genFn0()))) {
            var v11 string = ""
            v12 := GenS1{F0: genFn0(), F1: v9}
            v13 := GenS2{F0: -471, F1: (v4 / (((v4 % (((v4 >> ((42) & 63))) | 1))) | 1))}
            _ = v11
            _ = v12
            _ = v13
        }
        v4 = (v4 + v4)
    }
    if (!false) {
        var v14 uint8 = (^(v4))
        fmt.Println(v9)
        _ = v14
    } else {
        var v15 uint32 = v2
        _ = v15
    }
    fmt.Println(v2, v4, v9)
    var v16 int = -505
    v17 := GenS2{F0: 1000, F1: (v4 | 255)}
    v18 := GenS3{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
