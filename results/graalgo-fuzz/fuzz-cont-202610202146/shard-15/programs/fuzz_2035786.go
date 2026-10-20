package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int }

func main() {
    var v1 int8 = 117
    var v2 int16 = -1
    var v3 uint8 = 197
    var v4 uint64 = 2
    if true {
        v4 = v4
        v5 := GenS1{F0: 0}
        var v6 uint8 = (v3 % ((v3) | 1))
        _ = v5
        _ = v6
    }
    var v7 string = ("foo bar" + "foo bar")
    v8 := GenS2{F0: 968, F1: 0}
    var v9 bool = (true && (false && ((^(v4)) < (v4 / (((v4 / ((4294967295) | 1))) | 1)))))
    v7 = v7
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(fi10)
        switch 839 {
        case 768:
            fmt.Println(v1, v9, v9)
        case 2147483647:
            var v11 uint8 = v3
            _ = v11
        default:
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v7 = v7
        if (!v9) {
            v1 = -128
        }
        v4 = v4
        _ = fi13
    }
    var v14 float32 = -2.25
    v15 := GenS2{F0: 2147483647, F1: -1}
    var v16 uint32 = 0
    v17 := GenS2{F0: v16, F1: 0}
    if v9 {
        v18 := GenS3{F0: 32767}
        v19 := []int{-9223372036854775808, 1, -9223372036854775808}
        _ = v18
        _ = v19
    }
    v20 := []int{0}
    v9 = (1e10 <= (v14 - v14))
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(fi21, v9, v7)
            _ = fi22
        }
        _ = fi21
    }
    var v23 uint8 = uint8(v16)
    var v24 bool = v9
    v25 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
