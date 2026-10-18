package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string }
func genFn0(p0 uint32) int64 {
    switch "a" {
    case "\n\t":
        v5 := []int{}
        _ = v5
    default:
        p0 += (^(p0))
    }
    v6 := []int{935, -1, 0, -9223372036854775808}
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 bool = (32767 != 1)
    v2 := make(map[int]bool)
    var v3 string = ("a" + "hello")
    var v4 uint64 = 936
    v7 := GenS1{F0: 114, F1: 0, F2: 18446744073709551615}
    var v8 int64 = genFn0(65535)
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int64 = (v8 % ((66) | 1))
        var v11 uint8 = 127
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v4 -= (^(v4))
        v13 := GenS2{F0: (v4 << ((255) & 63)), F1: 1}
        _ = fi12
        _ = v13
    }
    var v14 int64 = v8
    v4 = v4
    fmt.Println(v1, v1, v1)
    fmt.Println(v1, v1, v14)
    var v15 uint16 = 0
    var v16 uint16 = 65535
    v17 := GenS1{F0: 127, F1: -912, F2: 65535}
    v18 := GenS3{F0: "hello"}
    v19 := GenS1{F0: -1, F1: 32767, F2: 2}
    var v20 int32 = 42
    v20 -= 235
    v21 := []int{32767, 32768, -9223372036854775808, -741}
    var v22 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
