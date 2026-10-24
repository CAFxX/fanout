package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
func genFn0() int16 {
    var v5 int32 = 948
    fmt.Println(v5)
    var v6 uint64 = 18446744073709551615
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 bool) int8 {
    v7 := GenS1{F0: -1, F1: 732}
    fmt.Println(p0)
    switch "foo bar" {
    case "hello":
        for fi8 := 0; fi8 < 1; fi8++ {
            switch (fi8 >> ((255) & 63)) {
            case -1:
                p0 = p0
            case 32767:
                var v9 uint64 = 4294967295
                _ = v9
            case -254:
                v10 := []int{(fi8 * 42), 32767, fi8}
                _ = v10
            default:
                fmt.Println(p0)
            }
            fmt.Println(p0, p0)
            _ = fi8
        }
    }
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 42
    var v3 uint = 0
    var v4 uint8 = 127
    v2 -= (v2 - 815)
    v4 = 1
    v11 := GenS1{F0: 0, F1: 150}
    v2 = v2
    if false {
        v12 := GenS2{F0: uint(v4), F1: "hello"}
        _ = v12
    }
    v4 += (v4 & (v4 ^ v4))
    var v13 uint64 = v2
    var v14 uint32 = 4294967295
    v2 -= (v2 | (v2 % ((v13) | 1)))
    var v15 int8 = 100
    v16 := []int{32767, 32768, 23}
    fmt.Println(v15, v2)
    switch 42 {
    case 355:
        fmt.Println(v4)
    case 2147483647:
        v15 = (v15 * (v15 - 8))
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS3{F0: 38, F1: fi17, F2: 0}
        _ = fi17
        _ = v18
    }
    var v19 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
