package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
type GenS4 struct { F0 bool; F1 int64; F2 uint64 }
func genFn0() int32 {
    var v5 bool = true
    v5 = (!(1 <= 1))
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 uint8, p2 int64) uint16 {
    v6 := GenS1{F0: (p1 ^ p1)}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 737
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = -1
    var v3 uint16 = 32767
    var v4 int64 = -409
    var v8 uint64 = 807
    if true {
        v4 = v4
        v9 := GenS2{F0: v8}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS2{F0: v8}
            var v12 int16 = 32767
            var v13 int8 = 0
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        switch "hello" {
        case "":
            fmt.Println(v2)
        default:
            v14 := []int{}
            _ = v14
        }
    }
    v2 += v2
    v1 = uint8(v3)
    v3 = 32767
    switch "a" {
    case "\n\t":
        v1 -= v1
    case "a":
        var v15 float32 = 0.1
        _ = v15
    case "foo bar":
        v16 := GenS3{F0: int64(v2), F1: 119, F2: -381}
        _ = v16
    default:
        if true {
            v17 := []int{-2147483648}
            _ = v17
        }
    }
    fmt.Println(v2, v8)
    fmt.Println(v1)
    var v18 uint = uint(v1)
    var v19 int16 = v2
    v19 += (v19 >> ((v18) & 63))
    v3 -= v3
    fmt.Println(v4, v18)
    if (100 >= 127) {
        var v20 uint8 = v1
        for fi21 := 0; fi21 < 11; fi21++ {
            v4 = (v4 / ((v4) | 1))
            _ = fi21
        }
        _ = v20
    } else {
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := GenS1{F0: v1}
            v24 := GenS4{F0: (v3 != (v3 << ((v18) & 63))), F1: v4, F2: (v8 >> (((v18 / (((v18 & 0)) | 1))) & 63))}
            v18 -= (v18 - 0)
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := GenS2{F0: (v8 + v8)}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
