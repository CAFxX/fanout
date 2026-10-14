package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 string, p1 uint64) uint16 {
    if (-1 <= -128) {
        var v4 uint32 = 1
        for fi5 := 0; fi5 < 9; fi5++ {
            fmt.Println(v4, p0)
            _ = fi5
        }
        fmt.Println(p1, v4, p0)
        _ = v4
    }
    v6 := []int{1, 4294967295, 1}
    if ((p1 + p1) == p1) {
        v7 := GenS1{F0: 1000, F1: -32768}
        _ = v7
    }
    var v8 bool = true
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 827
    var v3 uint = 42
    v2 += 2147483647
    v3 += (^(v3))
    var v9 int8 = 100
    var v10 int16 = -232
    switch 0 {
    case 1:
        delete(v1, 2147483648)
    case 270:
        var v11 int64 = 9223372036854775807
        _ = v11
    case 32768:
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS2{F0: "hello"}
            _ = fi12
            _ = v13
        }
    }
    _ = len(v1)
    v10 -= (v10 << (((v3 & 255)) & 63))
    _ = len(v1)
    var v14 string = "hello"
    fmt.Println(v9, v3, v14)
    fmt.Println(v10)
    v2 -= 0
    var v15 int64 = 363
    v16 := []int{9223372036854775807, 128}
    var v17 string = v14
    var v18 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
