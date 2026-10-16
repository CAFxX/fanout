package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 int8; F1 int64; F2 uint16 }

func main() {
    var v1 uint = 4294967295
    var v2 uint64 = 1
    var v3 float64 = (0.0 + float64(3.14159))
    var v4 bool = (-791 == 842)
    v5 := GenS1{F0: 233}
    switch ("" + "") {
    case "hello":
        v6 := GenS1{F0: 1}
        _ = v6
    case "":
        var v7 float32 = 0.0
        _ = v7
    }
    switch -411 {
    case 128:
        var v8 int = 567
        _ = v8
    }
    v9 := []int{-896, -9223372036854775808}
    var v10 float32 = (0.1 + (0.0 - float32(3.14159)))
    v3 = (float64(3.14159) + float64(v10))
    v11 := GenS2{F0: 1, F1: (v1 & 255), F2: v4}
    v12 := GenS3{F0: 100, F1: 42, F2: 255}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v4 = (!(v4 || (-32768 < 1)))
        var v14 uint64 = (v2 % (((^(v2))) | 1))
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 255
    _ = len(v9)
    fmt.Println(v4)
    v1 = (v1 % ((65535) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
