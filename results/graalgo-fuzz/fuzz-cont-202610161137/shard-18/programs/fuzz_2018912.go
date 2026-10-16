package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint32 }
func genFn0(p0 bool, p1 uint32, p2 string) uint32 {
    fmt.Println(p0, p0, p2)
    v7 := []int{4294967295, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1(p0 uint64, p1 uint8, p2 uint) int8 {
    var v8 int16 = -32768
    v9 := []int{-161, -899, 32768, 705}
    v8 = (v8 * 1000)
    switch "hello" {
    case "hello":
        p1 += p1
    case "":
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: uint8(p2), F1: true, F2: 30}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int = 32767
    v2 := []uint32{}
    var v3 uint32 = 514
    var v4 uint16 = 15
    var v5 uint = 606
    var v6 int8 = 69
    var v12 uint32 = 58
    fmt.Println(v3, v5)
    var v13 string = ("\n\t" + ("" + "hello"))
    var v14 int32 = 42
    v6 = -92
    if false {
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: 255, F1: (42 != (v14 * v14)), F2: 2147483648}
            v17 := GenS1{F0: 42, F1: (-86 == (v6 >> (((v5 * (v5 ^ 948))) & 63))), F2: (v12 ^ v3)}
            _ = fi15
            _ = v16
            _ = v17
        }
        v12 += (^(v12))
        fmt.Println(v12, v6)
    }
    var v18 int64 = int64(v1)
    fmt.Println(v5, v5, v6)
    var v19 uint8 = 0
    v3 += genFn0(false, (v12 << (((v5 & (v5 + 18446744073709551615))) & 63)), "a")
    fmt.Println(v6, v14, v4)
    v20 := GenS2{F0: 65535, F1: v5, F2: v3}
    var v21 uint8 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
