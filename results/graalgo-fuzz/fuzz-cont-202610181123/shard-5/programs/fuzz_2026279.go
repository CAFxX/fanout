package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0() int64 {
    var v6 int32 = 0
    var v7 int32 = v6
    _ = v6
    _ = v7
    return -753
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = ((-956 == 42) && ((1 < -49) && (-9223372036854775808 == 42)))
    v3 := make(map[uint32]string)
    v4 := []int{418, 32768, -9223372036854775808, -1}
    var v5 uint = 4294967295
    var v8 uint32 = 2
    v9 := GenS1{F0: v8}
    if v2 {
        switch -758 {
        case -754:
            fmt.Println(v5)
        }
        v1 -= float32(0.1)
    }
    if (((v5 * (v5 << ((v5) & 63))) <= 1) && (((v5 != 65535) || v2) || v2)) {
        var v10 int8 = 86
        if v2 {
            v10 += 0
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(fi11)
                _ = fi11
            }
            var v12 int16 = 441
            _ = v12
        } else {
            v1 += ((1.5 / v1) - 1e10)
            var v13 uint32 = v8
            _ = v13
        }
        if (!false) {
            v2 = v2
        } else {
            v2 = (128 >= 42)
        }
        _ = v10
    } else {
        v14 := GenS2{F0: v2}
        _ = v14
    }
    var v15 int = 42
    fmt.Println(v15, v2)
    v4 = append(v4, (v15 | v15))
    var v16 int16 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS2{F0: (!(100 > -128))}
        switch fi17 {
        case 2147483647:
            var v19 int16 = v16
            _ = v19
        case -963:
            var v20 uint = (v5 - v5)
            _ = v20
        case 32767:
            v1 = v1
        }
        switch (^(fi17)) {
        case 9223372036854775807:
            var v21 uint8 = 0
            _ = v21
        case 127:
            v1 = float32(0.0)
        case 42:
            fmt.Println(fi17, v15)
        default:
            v8 += (v8 & v8)
        }
        _ = fi17
        _ = v18
    }
    switch "hello" {
    case "":
        v8 = v8
    case "\n\t":
        v22 := GenS3{F0: 1}
        _ = v22
    case "hello":
        v23 := GenS4{F0: 217, F1: 44}
        _ = v23
    }
    var v24 int64 = 416
    v25 := GenS1{F0: (^(v8))}
    v26 := GenS4{F0: 42, F1: 97}
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 uint8 = 205
        for fi29 := 0; fi29 < 5; fi29++ {
            v5 = v5
            v5 -= 615
            _ = fi29
        }
        fmt.Println(v5, v24, fi27)
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
